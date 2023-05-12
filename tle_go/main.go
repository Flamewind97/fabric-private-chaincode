/*
Copyright IBM Corp. All Rights Reserved.
Copyright 2020 Intel Corporation

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"bytes"
	"fmt"

	pcommon "github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric/bccsp"
	"github.com/hyperledger/fabric/common/flogging"
	"github.com/hyperledger/fabric/common/policies"
	"github.com/hyperledger/fabric/common/semaphore"
	"github.com/hyperledger/fabric/common/util"
	"github.com/hyperledger/fabric/core/committer/txvalidator"
	"github.com/hyperledger/fabric/core/committer/txvalidator/plugin"
	validatorv14 "github.com/hyperledger/fabric/core/committer/txvalidator/v14"
	validatorv20 "github.com/hyperledger/fabric/core/committer/txvalidator/v20"
	"github.com/hyperledger/fabric/core/committer/txvalidator/v20/plugindispatcher"
	vir "github.com/hyperledger/fabric/core/committer/txvalidator/v20/valinforetriever"
	"github.com/hyperledger/fabric/core/peer"
	"github.com/hyperledger/fabric/gossip/common"
	"github.com/hyperledger/fabric/protoutil"
)

func main() {

	// blockValidator := CreateBlockValidator()
	// err := VerifyBlock(blockValidator)

	// if err != nil {
	// 	fmt.Println("ERROR")
	// }

	// err = TxValidator()
	// if err != nil {
	// 	fmt.Println("Tx error")
	// }
}

var tleLogger = flogging.MustGetLogger("tle.verifyblock")

func VerifyBlock(channelPolicyManagerGetter policies.ChannelPolicyManagerGetter, chainID common.ChannelID, seqNum uint64, block *pcommon.Block) error {
	if block.Header == nil {
		return fmt.Errorf("Invalid Block on channel [%s]. Header must be different from nil.", chainID)
	}

	blockSeqNum := block.Header.Number
	if seqNum != blockSeqNum {
		return fmt.Errorf("Claimed seqNum is [%d] but actual seqNum inside block is [%d]", seqNum, blockSeqNum)
	}

	// - Extract channelID and compare with chainID
	channelID, err := protoutil.GetChannelIDFromBlock(block)
	if err != nil {
		return fmt.Errorf("Failed getting channel id from block with id [%d] on channel [%s]: [%s]", block.Header.Number, chainID, err)
	}

	if channelID != string(chainID) {
		return fmt.Errorf("Invalid block's channel id. Expected [%s]. Given [%s]", chainID, channelID)
	}

	// - Unmarshal medatada
	if block.Metadata == nil || len(block.Metadata.Metadata) == 0 {
		return fmt.Errorf("Block with id [%d] on channel [%s] does not have metadata. Block not valid.", block.Header.Number, chainID)
	}

	metadata, err := protoutil.GetMetadataFromBlock(block, pcommon.BlockMetadataIndex_SIGNATURES)
	if err != nil {
		return fmt.Errorf("Failed unmarshalling medatata for signatures [%s]", err)
	}

	// - Verify that Header.DataHash is equal to the hash of block.Data
	// This is to ensure that the header is consistent with the data carried by this block
	if !bytes.Equal(protoutil.BlockDataHash(block.Data), block.Header.DataHash) {
		return fmt.Errorf("Header.DataHash is different from Hash(block.Data) for block with id [%d] on channel [%s]", block.Header.Number, chainID)
	}

	// - Get Policy for block validation

	// Get the policy manager for channelID
	cpm := channelPolicyManagerGetter.Manager(channelID)
	if cpm == nil {
		return fmt.Errorf("Could not acquire policy manager for channel %s", channelID)
	}
	tleLogger.Debugf("Got policy manager for channel [%s]", channelID)

	// Get block validation policy
	policy, ok := cpm.GetPolicy(policies.BlockValidation)
	// ok is true if it was the policy requested, or false if it is the default policy
	tleLogger.Debugf("Got block validation policy for channel [%s] with flag [%t]", channelID, ok)

	// - Prepare SignedData
	signatureSet := []*protoutil.SignedData{}
	for _, metadataSignature := range metadata.Signatures {
		shdr, err := protoutil.UnmarshalSignatureHeader(metadataSignature.SignatureHeader)
		if err != nil {
			return fmt.Errorf("Failed unmarshalling signature header for block with id [%d] on channel [%s]: [%s]", block.Header.Number, chainID, err)
		}
		signatureSet = append(
			signatureSet,
			&protoutil.SignedData{
				Identity:  shdr.Creator,
				Data:      util.ConcatenateBytes(metadata.Value, metadataSignature.SignatureHeader, protoutil.BlockHeaderBytes(block.Header)),
				Signature: metadataSignature.Signature,
			},
		)
	}

	// - Evaluate policy
	return policy.EvaluateSignedData(signatureSet)
}

func CreateTxValidator(
	nWorkers int,
	cid string,
	cryptoProvider bccsp.BCCSP,
	channel *peer.Channel,
	// bundle *channelconfig.Bundle,
	pluginMapper plugin.MapBasedMapper,
	policyManagerFunc func(cid string) policies.Manager,
	legacyLifecycleValidation plugindispatcher.LifecycleResources,
	newLifecycleValidation plugindispatcher.CollectionAndLifecycleResources,
) (*txvalidator.ValidationRouter, error) {

	validationWorkersSemaphore := semaphore.New(nWorkers)

	validator := &txvalidator.ValidationRouter{
		CapabilityProvider: channel,
		V14Validator: validatorv14.NewTxValidator(
			cid,
			validationWorkersSemaphore,
			channel,
			pluginMapper,
			cryptoProvider,
		),
		V20Validator: validatorv20.NewTxValidator(
			cid,
			validationWorkersSemaphore,
			channel,
			channel.Ledger(),
			&vir.ValidationInfoRetrieveShim{
				New:    newLifecycleValidation,
				Legacy: legacyLifecycleValidation,
			},
			&peer.CollectionInfoShim{
				CollectionAndLifecycleResources: newLifecycleValidation,
				// ChannelID:                       bundle.ConfigtxValidator().ChannelID(),
				ChannelID: cid,
			},
			pluginMapper,
			policies.PolicyManagerGetterFunc(policyManagerFunc),
			cryptoProvider,
		),
	}

	return validator, nil
}

func ValidateTx(validator *txvalidator.ValidationRouter, block *pcommon.Block) error {
	return validator.Validate(block)
}
