// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: proto/protos/wasm/wasm.proto

package wasm

import (
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// QueryAllContractStateRequest is the request type for the
// Query/AllContractState RPC method
type QueryAllContractStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// address is the address of the contract
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// pagination defines an optional pagination for the request.
	Pagination *query.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *QueryAllContractStateRequest) Reset() {
	*x = QueryAllContractStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protos_wasm_wasm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryAllContractStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryAllContractStateRequest) ProtoMessage() {}

func (x *QueryAllContractStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protos_wasm_wasm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryAllContractStateRequest.ProtoReflect.Descriptor instead.
func (*QueryAllContractStateRequest) Descriptor() ([]byte, []int) {
	return file_proto_protos_wasm_wasm_proto_rawDescGZIP(), []int{0}
}

func (x *QueryAllContractStateRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *QueryAllContractStateRequest) GetPagination() *query.PageRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

// QueryAllContractStateResponse is the response type for the
// Query/AllContractState RPC method
type QueryAllContractStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Models []*Model `protobuf:"bytes,1,rep,name=models,proto3" json:"models,omitempty"`
	// pagination defines the pagination in the response.
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *QueryAllContractStateResponse) Reset() {
	*x = QueryAllContractStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protos_wasm_wasm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryAllContractStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryAllContractStateResponse) ProtoMessage() {}

func (x *QueryAllContractStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protos_wasm_wasm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryAllContractStateResponse.ProtoReflect.Descriptor instead.
func (*QueryAllContractStateResponse) Descriptor() ([]byte, []int) {
	return file_proto_protos_wasm_wasm_proto_rawDescGZIP(), []int{1}
}

func (x *QueryAllContractStateResponse) GetModels() []*Model {
	if x != nil {
		return x.Models
	}
	return nil
}

func (x *QueryAllContractStateResponse) GetPagination() *query.PageResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

// Model is a struct that holds a KV pair
type Model struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// hex-encode key to read it better (this is often ascii)
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// base64-encode raw value
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Model) Reset() {
	*x = Model{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protos_wasm_wasm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Model) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Model) ProtoMessage() {}

func (x *Model) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protos_wasm_wasm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Model.ProtoReflect.Descriptor instead.
func (*Model) Descriptor() ([]byte, []int) {
	return file_proto_protos_wasm_wasm_proto_rawDescGZIP(), []int{2}
}

func (x *Model) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *Model) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

// QuerySmartContractStateRequest is the request type for the
// Query/SmartContractState RPC method
type QuerySmartContractStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// address is the address of the contract
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// QueryData contains the query data passed to the contract
	QueryData []byte `protobuf:"bytes,2,opt,name=query_data,json=queryData,proto3" json:"query_data,omitempty"`
}

func (x *QuerySmartContractStateRequest) Reset() {
	*x = QuerySmartContractStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protos_wasm_wasm_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuerySmartContractStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuerySmartContractStateRequest) ProtoMessage() {}

func (x *QuerySmartContractStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protos_wasm_wasm_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuerySmartContractStateRequest.ProtoReflect.Descriptor instead.
func (*QuerySmartContractStateRequest) Descriptor() ([]byte, []int) {
	return file_proto_protos_wasm_wasm_proto_rawDescGZIP(), []int{3}
}

func (x *QuerySmartContractStateRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *QuerySmartContractStateRequest) GetQueryData() []byte {
	if x != nil {
		return x.QueryData
	}
	return nil
}

// QuerySmartContractStateResponse is the response type for the
// Query/SmartContractState RPC method
type QuerySmartContractStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Data contains the json data returned from the smart contract
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *QuerySmartContractStateResponse) Reset() {
	*x = QuerySmartContractStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protos_wasm_wasm_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuerySmartContractStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuerySmartContractStateResponse) ProtoMessage() {}

func (x *QuerySmartContractStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protos_wasm_wasm_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuerySmartContractStateResponse.ProtoReflect.Descriptor instead.
func (*QuerySmartContractStateResponse) Descriptor() ([]byte, []int) {
	return file_proto_protos_wasm_wasm_proto_rawDescGZIP(), []int{4}
}

func (x *QuerySmartContractStateResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_proto_protos_wasm_wasm_proto protoreflect.FileDescriptor

var file_proto_protos_wasm_wasm_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x77,
	0x61, 0x73, 0x6d, 0x2f, 0x77, 0x61, 0x73, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x63, 0x6f, 0x73, 0x6d, 0x77, 0x61, 0x73, 0x6d, 0x2e, 0x77, 0x61, 0x73, 0x6d, 0x2e, 0x76, 0x31,
	0x1a, 0x11, 0x61, 0x6d, 0x69, 0x6e, 0x6f, 0x2f, 0x61, 0x6d, 0x69, 0x6e, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x01, 0x0a, 0x1c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41,
	0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x46, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0a, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xa4, 0x01, 0x0a, 0x1d, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x73,
	0x6d, 0x77, 0x61, 0x73, 0x6d, 0x2e, 0x77, 0x61, 0x73, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x42, 0x09, 0xc8, 0xde, 0x1f, 0x00, 0xa8, 0xe7, 0xb0, 0x2a, 0x01, 0x52, 0x06,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x12, 0x47, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x65, 0x0a, 0x05, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x46, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x34, 0xfa, 0xde, 0x1f, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x62, 0x66, 0x74, 0x2f, 0x63,
	0x6f, 0x6d, 0x65, 0x74, 0x62, 0x66, 0x74, 0x2f, 0x6c, 0x69, 0x62, 0x73, 0x2f, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x2e, 0x48, 0x65, 0x78, 0x42, 0x79, 0x74, 0x65, 0x73, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x71, 0x0a, 0x1e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53,
	0x6d, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x35, 0x0a, 0x0a, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x16, 0xfa, 0xde, 0x1f, 0x12, 0x52, 0x61, 0x77, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x09,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x22, 0x4d, 0x0a, 0x1f, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x16, 0xfa, 0xde, 0x1f, 0x12,
	0x52, 0x61, 0x77, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xfb, 0x01, 0x0a, 0x05, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x12, 0x75, 0x0a, 0x10, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2e, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x77, 0x61, 0x73,
	0x6d, 0x2e, 0x77, 0x61, 0x73, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41,
	0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x77, 0x61, 0x73,
	0x6d, 0x2e, 0x77, 0x61, 0x73, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41,
	0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7b, 0x0a, 0x12, 0x53, 0x6d, 0x61,
	0x72, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x30, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x77, 0x61, 0x73, 0x6d, 0x2e, 0x77, 0x61, 0x73, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x31, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x77, 0x61, 0x73, 0x6d, 0x2e, 0x77, 0x61, 0x73,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x67, 0x67, 0x79, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x7a,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x74, 0x65, 0x6c, 0x6c, 0x61, 0x72, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x77, 0x61, 0x73, 0x6d, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_protos_wasm_wasm_proto_rawDescOnce sync.Once
	file_proto_protos_wasm_wasm_proto_rawDescData = file_proto_protos_wasm_wasm_proto_rawDesc
)

func file_proto_protos_wasm_wasm_proto_rawDescGZIP() []byte {
	file_proto_protos_wasm_wasm_proto_rawDescOnce.Do(func() {
		file_proto_protos_wasm_wasm_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_protos_wasm_wasm_proto_rawDescData)
	})
	return file_proto_protos_wasm_wasm_proto_rawDescData
}

var file_proto_protos_wasm_wasm_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_protos_wasm_wasm_proto_goTypes = []interface{}{
	(*QueryAllContractStateRequest)(nil),    // 0: cosmwasm.wasm.v1.QueryAllContractStateRequest
	(*QueryAllContractStateResponse)(nil),   // 1: cosmwasm.wasm.v1.QueryAllContractStateResponse
	(*Model)(nil),                           // 2: cosmwasm.wasm.v1.Model
	(*QuerySmartContractStateRequest)(nil),  // 3: cosmwasm.wasm.v1.QuerySmartContractStateRequest
	(*QuerySmartContractStateResponse)(nil), // 4: cosmwasm.wasm.v1.QuerySmartContractStateResponse
	(*query.PageRequest)(nil),               // 5: cosmos.base.query.v1beta1.PageRequest
	(*query.PageResponse)(nil),              // 6: cosmos.base.query.v1beta1.PageResponse
}
var file_proto_protos_wasm_wasm_proto_depIdxs = []int32{
	5, // 0: cosmwasm.wasm.v1.QueryAllContractStateRequest.pagination:type_name -> cosmos.base.query.v1beta1.PageRequest
	2, // 1: cosmwasm.wasm.v1.QueryAllContractStateResponse.models:type_name -> cosmwasm.wasm.v1.Model
	6, // 2: cosmwasm.wasm.v1.QueryAllContractStateResponse.pagination:type_name -> cosmos.base.query.v1beta1.PageResponse
	0, // 3: cosmwasm.wasm.v1.Query.AllContractState:input_type -> cosmwasm.wasm.v1.QueryAllContractStateRequest
	3, // 4: cosmwasm.wasm.v1.Query.SmartContractState:input_type -> cosmwasm.wasm.v1.QuerySmartContractStateRequest
	1, // 5: cosmwasm.wasm.v1.Query.AllContractState:output_type -> cosmwasm.wasm.v1.QueryAllContractStateResponse
	4, // 6: cosmwasm.wasm.v1.Query.SmartContractState:output_type -> cosmwasm.wasm.v1.QuerySmartContractStateResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_protos_wasm_wasm_proto_init() }
func file_proto_protos_wasm_wasm_proto_init() {
	if File_proto_protos_wasm_wasm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_protos_wasm_wasm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryAllContractStateRequest); i {
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
		file_proto_protos_wasm_wasm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryAllContractStateResponse); i {
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
		file_proto_protos_wasm_wasm_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Model); i {
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
		file_proto_protos_wasm_wasm_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuerySmartContractStateRequest); i {
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
		file_proto_protos_wasm_wasm_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuerySmartContractStateResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_protos_wasm_wasm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_protos_wasm_wasm_proto_goTypes,
		DependencyIndexes: file_proto_protos_wasm_wasm_proto_depIdxs,
		MessageInfos:      file_proto_protos_wasm_wasm_proto_msgTypes,
	}.Build()
	File_proto_protos_wasm_wasm_proto = out.File
	file_proto_protos_wasm_wasm_proto_rawDesc = nil
	file_proto_protos_wasm_wasm_proto_goTypes = nil
	file_proto_protos_wasm_wasm_proto_depIdxs = nil
}