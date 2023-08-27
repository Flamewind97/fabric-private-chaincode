// Copyright IBM Corp. All Rights Reserved.
// Copyright 2020 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.11.4
// source: fpc/tl_session.proto

package tl_session

import (
	any "github.com/golang/protobuf/ptypes/any"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SessionSetupInitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelId   string `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	ChaincodeId string `protobuf:"bytes,2,opt,name=chaincode_id,json=chaincodeId,proto3" json:"chaincode_id,omitempty"`
	EnclaveId   string `protobuf:"bytes,3,opt,name=enclave_id,json=enclaveId,proto3" json:"enclave_id,omitempty"`
}

func (x *SessionSetupInitRequest) Reset() {
	*x = SessionSetupInitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fpc_tl_session_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionSetupInitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionSetupInitRequest) ProtoMessage() {}

func (x *SessionSetupInitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fpc_tl_session_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionSetupInitRequest.ProtoReflect.Descriptor instead.
func (*SessionSetupInitRequest) Descriptor() ([]byte, []int) {
	return file_fpc_tl_session_proto_rawDescGZIP(), []int{0}
}

func (x *SessionSetupInitRequest) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *SessionSetupInitRequest) GetChaincodeId() string {
	if x != nil {
		return x.ChaincodeId
	}
	return ""
}

func (x *SessionSetupInitRequest) GetEnclaveId() string {
	if x != nil {
		return x.EnclaveId
	}
	return ""
}

type SessionSetupInitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId uint64 `protobuf:"varint,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	// serialized struct sgx_dh_msg1_t
	Msg1 []byte `protobuf:"bytes,2,opt,name=msg1,proto3" json:"msg1,omitempty"`
}

func (x *SessionSetupInitResponse) Reset() {
	*x = SessionSetupInitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fpc_tl_session_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionSetupInitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionSetupInitResponse) ProtoMessage() {}

func (x *SessionSetupInitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fpc_tl_session_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionSetupInitResponse.ProtoReflect.Descriptor instead.
func (*SessionSetupInitResponse) Descriptor() ([]byte, []int) {
	return file_fpc_tl_session_proto_rawDescGZIP(), []int{1}
}

func (x *SessionSetupInitResponse) GetSessionId() uint64 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

func (x *SessionSetupInitResponse) GetMsg1() []byte {
	if x != nil {
		return x.Msg1
	}
	return nil
}

type SessionSetupCompleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId uint64 `protobuf:"varint,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	// serialized struct sgx_dh_msg2_t
	Msg2 []byte `protobuf:"bytes,2,opt,name=msg2,proto3" json:"msg2,omitempty"`
}

func (x *SessionSetupCompleteRequest) Reset() {
	*x = SessionSetupCompleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fpc_tl_session_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionSetupCompleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionSetupCompleteRequest) ProtoMessage() {}

func (x *SessionSetupCompleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fpc_tl_session_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionSetupCompleteRequest.ProtoReflect.Descriptor instead.
func (*SessionSetupCompleteRequest) Descriptor() ([]byte, []int) {
	return file_fpc_tl_session_proto_rawDescGZIP(), []int{2}
}

func (x *SessionSetupCompleteRequest) GetSessionId() uint64 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

func (x *SessionSetupCompleteRequest) GetMsg2() []byte {
	if x != nil {
		return x.Msg2
	}
	return nil
}

type SessionSetupCompleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId uint64 `protobuf:"varint,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	// serialized struct sgx_dh_msg2_t
	Msg3        []byte `protobuf:"bytes,2,opt,name=msg3,proto3" json:"msg3,omitempty"`
	ChannelId   string `protobuf:"bytes,3,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	ChannelHash []byte `protobuf:"bytes,4,opt,name=channel_hash,json=channelHash,proto3" json:"channel_hash,omitempty"`
	ChaincodeId string `protobuf:"bytes,5,opt,name=chaincode_id,json=chaincodeId,proto3" json:"chaincode_id,omitempty"`
	EnclaveId   string `protobuf:"bytes,6,opt,name=enclave_id,json=enclaveId,proto3" json:"enclave_id,omitempty"`
}

func (x *SessionSetupCompleteResponse) Reset() {
	*x = SessionSetupCompleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fpc_tl_session_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionSetupCompleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionSetupCompleteResponse) ProtoMessage() {}

func (x *SessionSetupCompleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fpc_tl_session_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionSetupCompleteResponse.ProtoReflect.Descriptor instead.
func (*SessionSetupCompleteResponse) Descriptor() ([]byte, []int) {
	return file_fpc_tl_session_proto_rawDescGZIP(), []int{3}
}

func (x *SessionSetupCompleteResponse) GetSessionId() uint64 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

func (x *SessionSetupCompleteResponse) GetMsg3() []byte {
	if x != nil {
		return x.Msg3
	}
	return nil
}

func (x *SessionSetupCompleteResponse) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *SessionSetupCompleteResponse) GetChannelHash() []byte {
	if x != nil {
		return x.ChannelHash
	}
	return nil
}

func (x *SessionSetupCompleteResponse) GetChaincodeId() string {
	if x != nil {
		return x.ChaincodeId
	}
	return ""
}

func (x *SessionSetupCompleteResponse) GetEnclaveId() string {
	if x != nil {
		return x.EnclaveId
	}
	return ""
}

type SessionCloseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId uint64 `protobuf:"varint,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
}

func (x *SessionCloseRequest) Reset() {
	*x = SessionCloseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fpc_tl_session_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionCloseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionCloseRequest) ProtoMessage() {}

func (x *SessionCloseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fpc_tl_session_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionCloseRequest.ProtoReflect.Descriptor instead.
func (*SessionCloseRequest) Descriptor() ([]byte, []int) {
	return file_fpc_tl_session_proto_rawDescGZIP(), []int{4}
}

func (x *SessionCloseRequest) GetSessionId() uint64 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

type SessionCloseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId uint64 `protobuf:"varint,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
}

func (x *SessionCloseResponse) Reset() {
	*x = SessionCloseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fpc_tl_session_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionCloseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionCloseResponse) ProtoMessage() {}

func (x *SessionCloseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fpc_tl_session_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionCloseResponse.ProtoReflect.Descriptor instead.
func (*SessionCloseResponse) Descriptor() ([]byte, []int) {
	return file_fpc_tl_session_proto_rawDescGZIP(), []int{5}
}

func (x *SessionCloseResponse) GetSessionId() uint64 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

type SessionTXRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId uint64 `protobuf:"varint,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Nonce     []byte `protobuf:"bytes,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Request   []byte `protobuf:"bytes,3,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *SessionTXRequest) Reset() {
	*x = SessionTXRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fpc_tl_session_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionTXRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionTXRequest) ProtoMessage() {}

func (x *SessionTXRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fpc_tl_session_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionTXRequest.ProtoReflect.Descriptor instead.
func (*SessionTXRequest) Descriptor() ([]byte, []int) {
	return file_fpc_tl_session_proto_rawDescGZIP(), []int{6}
}

func (x *SessionTXRequest) GetSessionId() uint64 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

func (x *SessionTXRequest) GetNonce() []byte {
	if x != nil {
		return x.Nonce
	}
	return nil
}

func (x *SessionTXRequest) GetRequest() []byte {
	if x != nil {
		return x.Request
	}
	return nil
}

type SessionTXResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId uint64 `protobuf:"varint,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	// should be same as in request: somewhat redundant but makes processing of MAC easier/more uniform.
	Nonce     []byte `protobuf:"bytes,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Respoonse []byte `protobuf:"bytes,3,opt,name=respoonse,proto3" json:"respoonse,omitempty"`
}

func (x *SessionTXResponse) Reset() {
	*x = SessionTXResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fpc_tl_session_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionTXResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionTXResponse) ProtoMessage() {}

func (x *SessionTXResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fpc_tl_session_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionTXResponse.ProtoReflect.Descriptor instead.
func (*SessionTXResponse) Descriptor() ([]byte, []int) {
	return file_fpc_tl_session_proto_rawDescGZIP(), []int{7}
}

func (x *SessionTXResponse) GetSessionId() uint64 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

func (x *SessionTXResponse) GetNonce() []byte {
	if x != nil {
		return x.Nonce
	}
	return nil
}

func (x *SessionTXResponse) GetRespoonse() []byte {
	if x != nil {
		return x.Respoonse
	}
	return nil
}

type SessionError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId uint64 `protobuf:"varint,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	ErrorCode int32  `protobuf:"varint,2,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	ErrorMsg  string `protobuf:"bytes,3,opt,name=error_msg,json=errorMsg,proto3" json:"error_msg,omitempty"`
}

func (x *SessionError) Reset() {
	*x = SessionError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fpc_tl_session_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionError) ProtoMessage() {}

func (x *SessionError) ProtoReflect() protoreflect.Message {
	mi := &file_fpc_tl_session_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionError.ProtoReflect.Descriptor instead.
func (*SessionError) Descriptor() ([]byte, []int) {
	return file_fpc_tl_session_proto_rawDescGZIP(), []int{8}
}

func (x *SessionError) GetSessionId() uint64 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

func (x *SessionError) GetErrorCode() int32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

func (x *SessionError) GetErrorMsg() string {
	if x != nil {
		return x.ErrorMsg
	}
	return ""
}

type SessionMsgPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Payload:
	//	*SessionMsgPayload_StpIntReq
	//	*SessionMsgPayload_StpIntRsp
	//	*SessionMsgPayload_StpCmpReq
	//	*SessionMsgPayload_StpCmpRsp
	//	*SessionMsgPayload_ClsReq
	//	*SessionMsgPayload_ClsRsp
	//	*SessionMsgPayload_TxReq
	//	*SessionMsgPayload_TxRsp
	//	*SessionMsgPayload_Error
	Payload isSessionMsgPayload_Payload `protobuf_oneof:"payload"`
}

func (x *SessionMsgPayload) Reset() {
	*x = SessionMsgPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fpc_tl_session_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionMsgPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionMsgPayload) ProtoMessage() {}

func (x *SessionMsgPayload) ProtoReflect() protoreflect.Message {
	mi := &file_fpc_tl_session_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionMsgPayload.ProtoReflect.Descriptor instead.
func (*SessionMsgPayload) Descriptor() ([]byte, []int) {
	return file_fpc_tl_session_proto_rawDescGZIP(), []int{9}
}

func (m *SessionMsgPayload) GetPayload() isSessionMsgPayload_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *SessionMsgPayload) GetStpIntReq() *SessionSetupInitRequest {
	if x, ok := x.GetPayload().(*SessionMsgPayload_StpIntReq); ok {
		return x.StpIntReq
	}
	return nil
}

func (x *SessionMsgPayload) GetStpIntRsp() *SessionSetupInitResponse {
	if x, ok := x.GetPayload().(*SessionMsgPayload_StpIntRsp); ok {
		return x.StpIntRsp
	}
	return nil
}

func (x *SessionMsgPayload) GetStpCmpReq() *SessionSetupCompleteRequest {
	if x, ok := x.GetPayload().(*SessionMsgPayload_StpCmpReq); ok {
		return x.StpCmpReq
	}
	return nil
}

func (x *SessionMsgPayload) GetStpCmpRsp() *SessionSetupCompleteResponse {
	if x, ok := x.GetPayload().(*SessionMsgPayload_StpCmpRsp); ok {
		return x.StpCmpRsp
	}
	return nil
}

func (x *SessionMsgPayload) GetClsReq() *SessionCloseRequest {
	if x, ok := x.GetPayload().(*SessionMsgPayload_ClsReq); ok {
		return x.ClsReq
	}
	return nil
}

func (x *SessionMsgPayload) GetClsRsp() *SessionCloseResponse {
	if x, ok := x.GetPayload().(*SessionMsgPayload_ClsRsp); ok {
		return x.ClsRsp
	}
	return nil
}

func (x *SessionMsgPayload) GetTxReq() *SessionTXRequest {
	if x, ok := x.GetPayload().(*SessionMsgPayload_TxReq); ok {
		return x.TxReq
	}
	return nil
}

func (x *SessionMsgPayload) GetTxRsp() *SessionTXResponse {
	if x, ok := x.GetPayload().(*SessionMsgPayload_TxRsp); ok {
		return x.TxRsp
	}
	return nil
}

func (x *SessionMsgPayload) GetError() *SessionError {
	if x, ok := x.GetPayload().(*SessionMsgPayload_Error); ok {
		return x.Error
	}
	return nil
}

type isSessionMsgPayload_Payload interface {
	isSessionMsgPayload_Payload()
}

type SessionMsgPayload_StpIntReq struct {
	StpIntReq *SessionSetupInitRequest `protobuf:"bytes,1,opt,name=stp_int_req,json=stpIntReq,proto3,oneof"`
}

type SessionMsgPayload_StpIntRsp struct {
	StpIntRsp *SessionSetupInitResponse `protobuf:"bytes,2,opt,name=stp_int_rsp,json=stpIntRsp,proto3,oneof"`
}

type SessionMsgPayload_StpCmpReq struct {
	StpCmpReq *SessionSetupCompleteRequest `protobuf:"bytes,3,opt,name=stp_cmp_req,json=stpCmpReq,proto3,oneof"`
}

type SessionMsgPayload_StpCmpRsp struct {
	StpCmpRsp *SessionSetupCompleteResponse `protobuf:"bytes,4,opt,name=stp_cmp_rsp,json=stpCmpRsp,proto3,oneof"`
}

type SessionMsgPayload_ClsReq struct {
	ClsReq *SessionCloseRequest `protobuf:"bytes,5,opt,name=cls_req,json=clsReq,proto3,oneof"`
}

type SessionMsgPayload_ClsRsp struct {
	ClsRsp *SessionCloseResponse `protobuf:"bytes,6,opt,name=cls_rsp,json=clsRsp,proto3,oneof"`
}

type SessionMsgPayload_TxReq struct {
	TxReq *SessionTXRequest `protobuf:"bytes,7,opt,name=tx_req,json=txReq,proto3,oneof"`
}

type SessionMsgPayload_TxRsp struct {
	TxRsp *SessionTXResponse `protobuf:"bytes,8,opt,name=tx_rsp,json=txRsp,proto3,oneof"`
}

type SessionMsgPayload_Error struct {
	Error *SessionError `protobuf:"bytes,9,opt,name=error,proto3,oneof"`
}

func (*SessionMsgPayload_StpIntReq) isSessionMsgPayload_Payload() {}

func (*SessionMsgPayload_StpIntRsp) isSessionMsgPayload_Payload() {}

func (*SessionMsgPayload_StpCmpReq) isSessionMsgPayload_Payload() {}

func (*SessionMsgPayload_StpCmpRsp) isSessionMsgPayload_Payload() {}

func (*SessionMsgPayload_ClsReq) isSessionMsgPayload_Payload() {}

func (*SessionMsgPayload_ClsRsp) isSessionMsgPayload_Payload() {}

func (*SessionMsgPayload_TxReq) isSessionMsgPayload_Payload() {}

func (*SessionMsgPayload_TxRsp) isSessionMsgPayload_Payload() {}

func (*SessionMsgPayload_Error) isSessionMsgPayload_Payload() {}

type SessionMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SerializedPayload *any.Any `protobuf:"bytes,1,opt,name=serialized_payload,json=serializedPayload,proto3" json:"serialized_payload,omitempty"` // This should be a serialization of SessionMsgPayload
	Mac               []byte   `protobuf:"bytes,2,opt,name=mac,proto3" json:"mac,omitempty"`
}

func (x *SessionMsg) Reset() {
	*x = SessionMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fpc_tl_session_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionMsg) ProtoMessage() {}

func (x *SessionMsg) ProtoReflect() protoreflect.Message {
	mi := &file_fpc_tl_session_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionMsg.ProtoReflect.Descriptor instead.
func (*SessionMsg) Descriptor() ([]byte, []int) {
	return file_fpc_tl_session_proto_rawDescGZIP(), []int{10}
}

func (x *SessionMsg) GetSerializedPayload() *any.Any {
	if x != nil {
		return x.SerializedPayload
	}
	return nil
}

func (x *SessionMsg) GetMac() []byte {
	if x != nil {
		return x.Mac
	}
	return nil
}

var File_fpc_tl_session_proto protoreflect.FileDescriptor

var file_fpc_tl_session_proto_rawDesc = []byte{
	0x0a, 0x14, 0x66, 0x70, 0x63, 0x2f, 0x74, 0x6c, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x6c, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7a, 0x0a,
	0x17, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x75, 0x70, 0x49, 0x6e, 0x69,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x63, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x6e,
	0x63, 0x6c, 0x61, 0x76, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x65, 0x6e, 0x63, 0x6c, 0x61, 0x76, 0x65, 0x49, 0x64, 0x22, 0x4d, 0x0a, 0x18, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x75, 0x70, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x73, 0x67, 0x31, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x6d, 0x73, 0x67, 0x31, 0x22, 0x50, 0x0a, 0x1b, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x75, 0x70, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x73, 0x67, 0x32, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x6d, 0x73, 0x67, 0x32, 0x22, 0xd5, 0x01, 0x0a, 0x1c, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x75, 0x70, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x73,
	0x67, 0x33, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x6d, 0x73, 0x67, 0x33, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64,
	0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x6e, 0x63, 0x6c, 0x61, 0x76, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6e, 0x63, 0x6c, 0x61, 0x76, 0x65,
	0x49, 0x64, 0x22, 0x34, 0x0a, 0x13, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6c, 0x6f,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x35, 0x0a, 0x14, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22,
	0x61, 0x0a, 0x10, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x58, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x66, 0x0a, 0x11, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x58, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x09, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x69, 0x0a, 0x0c, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x4d, 0x73, 0x67, 0x22, 0xde, 0x04, 0x0a, 0x11, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x4d, 0x73, 0x67, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x45, 0x0a, 0x0b, 0x73,
	0x74, 0x70, 0x5f, 0x69, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x74, 0x6c, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x75, 0x70, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x09, 0x73, 0x74, 0x70, 0x49, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x46, 0x0a, 0x0b, 0x73, 0x74, 0x70, 0x5f, 0x69, 0x6e, 0x74, 0x5f, 0x72, 0x73,
	0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x74, 0x6c, 0x5f, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x75,
	0x70, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52,
	0x09, 0x73, 0x74, 0x70, 0x49, 0x6e, 0x74, 0x52, 0x73, 0x70, 0x12, 0x49, 0x0a, 0x0b, 0x73, 0x74,
	0x70, 0x5f, 0x63, 0x6d, 0x70, 0x5f, 0x72, 0x65, 0x71, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x74, 0x6c, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x75, 0x70, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x09, 0x73, 0x74, 0x70, 0x43,
	0x6d, 0x70, 0x52, 0x65, 0x71, 0x12, 0x4a, 0x0a, 0x0b, 0x73, 0x74, 0x70, 0x5f, 0x63, 0x6d, 0x70,
	0x5f, 0x72, 0x73, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x74, 0x6c, 0x5f,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x74, 0x75, 0x70, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x09, 0x73, 0x74, 0x70, 0x43, 0x6d, 0x70, 0x52, 0x73,
	0x70, 0x12, 0x3a, 0x0a, 0x07, 0x63, 0x6c, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x6c, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x06, 0x63, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x12, 0x3b, 0x0a,
	0x07, 0x63, 0x6c, 0x73, 0x5f, 0x72, 0x73, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x74, 0x6c, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x48, 0x00, 0x52, 0x06, 0x63, 0x6c, 0x73, 0x52, 0x73, 0x70, 0x12, 0x35, 0x0a, 0x06, 0x74, 0x78,
	0x5f, 0x72, 0x65, 0x71, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x6c, 0x5f,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54,
	0x58, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x05, 0x74, 0x78, 0x52, 0x65,
	0x71, 0x12, 0x36, 0x0a, 0x06, 0x74, 0x78, 0x5f, 0x72, 0x73, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x6c, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x58, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x48, 0x00, 0x52, 0x05, 0x74, 0x78, 0x52, 0x73, 0x70, 0x12, 0x30, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x6c, 0x5f, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x09, 0x0a, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x63, 0x0a, 0x0a, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x4d, 0x73, 0x67, 0x12, 0x43, 0x0a, 0x12, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a,
	0x65, 0x64, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x11, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a,
	0x65, 0x64, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x63,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6d, 0x61, 0x63, 0x42, 0x4c, 0x5a, 0x4a, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c,
	0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2d, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x2d, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x74,
	0x6c, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_fpc_tl_session_proto_rawDescOnce sync.Once
	file_fpc_tl_session_proto_rawDescData = file_fpc_tl_session_proto_rawDesc
)

func file_fpc_tl_session_proto_rawDescGZIP() []byte {
	file_fpc_tl_session_proto_rawDescOnce.Do(func() {
		file_fpc_tl_session_proto_rawDescData = protoimpl.X.CompressGZIP(file_fpc_tl_session_proto_rawDescData)
	})
	return file_fpc_tl_session_proto_rawDescData
}

var file_fpc_tl_session_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_fpc_tl_session_proto_goTypes = []interface{}{
	(*SessionSetupInitRequest)(nil),      // 0: tl_session.SessionSetupInitRequest
	(*SessionSetupInitResponse)(nil),     // 1: tl_session.SessionSetupInitResponse
	(*SessionSetupCompleteRequest)(nil),  // 2: tl_session.SessionSetupCompleteRequest
	(*SessionSetupCompleteResponse)(nil), // 3: tl_session.SessionSetupCompleteResponse
	(*SessionCloseRequest)(nil),          // 4: tl_session.SessionCloseRequest
	(*SessionCloseResponse)(nil),         // 5: tl_session.SessionCloseResponse
	(*SessionTXRequest)(nil),             // 6: tl_session.SessionTXRequest
	(*SessionTXResponse)(nil),            // 7: tl_session.SessionTXResponse
	(*SessionError)(nil),                 // 8: tl_session.SessionError
	(*SessionMsgPayload)(nil),            // 9: tl_session.SessionMsgPayload
	(*SessionMsg)(nil),                   // 10: tl_session.SessionMsg
	(*any.Any)(nil),                      // 11: google.protobuf.Any
}
var file_fpc_tl_session_proto_depIdxs = []int32{
	0,  // 0: tl_session.SessionMsgPayload.stp_int_req:type_name -> tl_session.SessionSetupInitRequest
	1,  // 1: tl_session.SessionMsgPayload.stp_int_rsp:type_name -> tl_session.SessionSetupInitResponse
	2,  // 2: tl_session.SessionMsgPayload.stp_cmp_req:type_name -> tl_session.SessionSetupCompleteRequest
	3,  // 3: tl_session.SessionMsgPayload.stp_cmp_rsp:type_name -> tl_session.SessionSetupCompleteResponse
	4,  // 4: tl_session.SessionMsgPayload.cls_req:type_name -> tl_session.SessionCloseRequest
	5,  // 5: tl_session.SessionMsgPayload.cls_rsp:type_name -> tl_session.SessionCloseResponse
	6,  // 6: tl_session.SessionMsgPayload.tx_req:type_name -> tl_session.SessionTXRequest
	7,  // 7: tl_session.SessionMsgPayload.tx_rsp:type_name -> tl_session.SessionTXResponse
	8,  // 8: tl_session.SessionMsgPayload.error:type_name -> tl_session.SessionError
	11, // 9: tl_session.SessionMsg.serialized_payload:type_name -> google.protobuf.Any
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_fpc_tl_session_proto_init() }
func file_fpc_tl_session_proto_init() {
	if File_fpc_tl_session_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fpc_tl_session_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionSetupInitRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fpc_tl_session_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionSetupInitResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fpc_tl_session_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionSetupCompleteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fpc_tl_session_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionSetupCompleteResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fpc_tl_session_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionCloseRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fpc_tl_session_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionCloseResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fpc_tl_session_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionTXRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fpc_tl_session_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionTXResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fpc_tl_session_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionError); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fpc_tl_session_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionMsgPayload); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fpc_tl_session_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionMsg); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_fpc_tl_session_proto_msgTypes[9].OneofWrappers = []interface{}{
		(*SessionMsgPayload_StpIntReq)(nil),
		(*SessionMsgPayload_StpIntRsp)(nil),
		(*SessionMsgPayload_StpCmpReq)(nil),
		(*SessionMsgPayload_StpCmpRsp)(nil),
		(*SessionMsgPayload_ClsReq)(nil),
		(*SessionMsgPayload_ClsRsp)(nil),
		(*SessionMsgPayload_TxReq)(nil),
		(*SessionMsgPayload_TxRsp)(nil),
		(*SessionMsgPayload_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_fpc_tl_session_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fpc_tl_session_proto_goTypes,
		DependencyIndexes: file_fpc_tl_session_proto_depIdxs,
		MessageInfos:      file_fpc_tl_session_proto_msgTypes,
	}.Build()
	File_fpc_tl_session_proto = out.File
	file_fpc_tl_session_proto_rawDesc = nil
	file_fpc_tl_session_proto_goTypes = nil
	file_fpc_tl_session_proto_depIdxs = nil
}
