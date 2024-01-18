// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: proto-blockchain_v2/blockchain.proto

package proto

import (
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

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stx []byte `protobuf:"bytes,1,opt,name=Stx,proto3" json:"Stx,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blockchain_v2_blockchain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blockchain_v2_blockchain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_proto_blockchain_v2_blockchain_proto_rawDescGZIP(), []int{0}
}

func (x *Transaction) GetStx() []byte {
	if x != nil {
		return x.Stx
	}
	return nil
}

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sblock []byte `protobuf:"bytes,1,opt,name=Sblock,proto3" json:"Sblock,omitempty"`
	Addr   string `protobuf:"bytes,2,opt,name=Addr,proto3" json:"Addr,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blockchain_v2_blockchain_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blockchain_v2_blockchain_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_proto_blockchain_v2_blockchain_proto_rawDescGZIP(), []int{1}
}

func (x *Block) GetSblock() []byte {
	if x != nil {
		return x.Sblock
	}
	return nil
}

func (x *Block) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=Result,proto3" json:"Result,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blockchain_v2_blockchain_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blockchain_v2_blockchain_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_blockchain_v2_blockchain_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type ChainRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request string `protobuf:"bytes,1,opt,name=Request,proto3" json:"Request,omitempty"`
}

func (x *ChainRequest) Reset() {
	*x = ChainRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blockchain_v2_blockchain_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChainRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChainRequest) ProtoMessage() {}

func (x *ChainRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blockchain_v2_blockchain_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChainRequest.ProtoReflect.Descriptor instead.
func (*ChainRequest) Descriptor() ([]byte, []int) {
	return file_proto_blockchain_v2_blockchain_proto_rawDescGZIP(), []int{3}
}

func (x *ChainRequest) GetRequest() string {
	if x != nil {
		return x.Request
	}
	return ""
}

type UtxoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request string `protobuf:"bytes,1,opt,name=Request,proto3" json:"Request,omitempty"`
}

func (x *UtxoRequest) Reset() {
	*x = UtxoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blockchain_v2_blockchain_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UtxoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UtxoRequest) ProtoMessage() {}

func (x *UtxoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blockchain_v2_blockchain_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UtxoRequest.ProtoReflect.Descriptor instead.
func (*UtxoRequest) Descriptor() ([]byte, []int) {
	return file_proto_blockchain_v2_blockchain_proto_rawDescGZIP(), []int{4}
}

func (x *UtxoRequest) GetRequest() string {
	if x != nil {
		return x.Request
	}
	return ""
}

type Chain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Schain []byte `protobuf:"bytes,1,opt,name=Schain,proto3" json:"Schain,omitempty"`
}

func (x *Chain) Reset() {
	*x = Chain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blockchain_v2_blockchain_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chain) ProtoMessage() {}

func (x *Chain) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blockchain_v2_blockchain_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chain.ProtoReflect.Descriptor instead.
func (*Chain) Descriptor() ([]byte, []int) {
	return file_proto_blockchain_v2_blockchain_proto_rawDescGZIP(), []int{5}
}

func (x *Chain) GetSchain() []byte {
	if x != nil {
		return x.Schain
	}
	return nil
}

type ProbeBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sprobeblock []byte `protobuf:"bytes,1,opt,name=Sprobeblock,proto3" json:"Sprobeblock,omitempty"`
}

func (x *ProbeBlock) Reset() {
	*x = ProbeBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blockchain_v2_blockchain_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProbeBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProbeBlock) ProtoMessage() {}

func (x *ProbeBlock) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blockchain_v2_blockchain_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProbeBlock.ProtoReflect.Descriptor instead.
func (*ProbeBlock) Descriptor() ([]byte, []int) {
	return file_proto_blockchain_v2_blockchain_proto_rawDescGZIP(), []int{6}
}

func (x *ProbeBlock) GetSprobeblock() []byte {
	if x != nil {
		return x.Sprobeblock
	}
	return nil
}

type Utxo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SUtxo []byte `protobuf:"bytes,1,opt,name=SUtxo,proto3" json:"SUtxo,omitempty"`
}

func (x *Utxo) Reset() {
	*x = Utxo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blockchain_v2_blockchain_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Utxo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Utxo) ProtoMessage() {}

func (x *Utxo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blockchain_v2_blockchain_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Utxo.ProtoReflect.Descriptor instead.
func (*Utxo) Descriptor() ([]byte, []int) {
	return file_proto_blockchain_v2_blockchain_proto_rawDescGZIP(), []int{7}
}

func (x *Utxo) GetSUtxo() []byte {
	if x != nil {
		return x.SUtxo
	}
	return nil
}

type MultipleBlocksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockId int32 `protobuf:"varint,1,opt,name=BlockId,proto3" json:"BlockId,omitempty"`
}

func (x *MultipleBlocksRequest) Reset() {
	*x = MultipleBlocksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blockchain_v2_blockchain_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultipleBlocksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultipleBlocksRequest) ProtoMessage() {}

func (x *MultipleBlocksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blockchain_v2_blockchain_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultipleBlocksRequest.ProtoReflect.Descriptor instead.
func (*MultipleBlocksRequest) Descriptor() ([]byte, []int) {
	return file_proto_blockchain_v2_blockchain_proto_rawDescGZIP(), []int{8}
}

func (x *MultipleBlocksRequest) GetBlockId() int32 {
	if x != nil {
		return x.BlockId
	}
	return 0
}

type ProbeBlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockId int32 `protobuf:"varint,1,opt,name=BlockId,proto3" json:"BlockId,omitempty"`
}

func (x *ProbeBlockRequest) Reset() {
	*x = ProbeBlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blockchain_v2_blockchain_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProbeBlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProbeBlockRequest) ProtoMessage() {}

func (x *ProbeBlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blockchain_v2_blockchain_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProbeBlockRequest.ProtoReflect.Descriptor instead.
func (*ProbeBlockRequest) Descriptor() ([]byte, []int) {
	return file_proto_blockchain_v2_blockchain_proto_rawDescGZIP(), []int{9}
}

func (x *ProbeBlockRequest) GetBlockId() int32 {
	if x != nil {
		return x.BlockId
	}
	return 0
}

type AddrRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request string `protobuf:"bytes,1,opt,name=Request,proto3" json:"Request,omitempty"`
}

func (x *AddrRequest) Reset() {
	*x = AddrRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blockchain_v2_blockchain_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddrRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddrRequest) ProtoMessage() {}

func (x *AddrRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blockchain_v2_blockchain_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddrRequest.ProtoReflect.Descriptor instead.
func (*AddrRequest) Descriptor() ([]byte, []int) {
	return file_proto_blockchain_v2_blockchain_proto_rawDescGZIP(), []int{10}
}

func (x *AddrRequest) GetRequest() string {
	if x != nil {
		return x.Request
	}
	return ""
}

type UpdateAddrRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=Address,proto3" json:"Address,omitempty"`
}

func (x *UpdateAddrRequest) Reset() {
	*x = UpdateAddrRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blockchain_v2_blockchain_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAddrRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAddrRequest) ProtoMessage() {}

func (x *UpdateAddrRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blockchain_v2_blockchain_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAddrRequest.ProtoReflect.Descriptor instead.
func (*UpdateAddrRequest) Descriptor() ([]byte, []int) {
	return file_proto_blockchain_v2_blockchain_proto_rawDescGZIP(), []int{11}
}

func (x *UpdateAddrRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type AddrList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddressList []string `protobuf:"bytes,1,rep,name=AddressList,proto3" json:"AddressList,omitempty"`
}

func (x *AddrList) Reset() {
	*x = AddrList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blockchain_v2_blockchain_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddrList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddrList) ProtoMessage() {}

func (x *AddrList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blockchain_v2_blockchain_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddrList.ProtoReflect.Descriptor instead.
func (*AddrList) Descriptor() ([]byte, []int) {
	return file_proto_blockchain_v2_blockchain_proto_rawDescGZIP(), []int{12}
}

func (x *AddrList) GetAddressList() []string {
	if x != nil {
		return x.AddressList
	}
	return nil
}

var File_proto_blockchain_v2_blockchain_proto protoreflect.FileDescriptor

var file_proto_blockchain_v2_blockchain_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x5f, 0x76, 0x32, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1f, 0x0a,
	0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03,
	0x53, 0x74, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x53, 0x74, 0x78, 0x22, 0x33,
	0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x53, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12,
	0x12, 0x0a, 0x04, 0x41, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x41,
	0x64, 0x64, 0x72, 0x22, 0x22, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x28, 0x0a, 0x0c, 0x43, 0x68, 0x61, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x27, 0x0a, 0x0b, 0x55, 0x74, 0x78, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x1f, 0x0a, 0x05, 0x43, 0x68,
	0x61, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x06, 0x53, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x22, 0x2e, 0x0a, 0x0a, 0x50,
	0x72, 0x6f, 0x62, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x70, 0x72,
	0x6f, 0x62, 0x65, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b,
	0x53, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x1c, 0x0a, 0x04, 0x55,
	0x74, 0x78, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x55, 0x74, 0x78, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x53, 0x55, 0x74, 0x78, 0x6f, 0x22, 0x31, 0x0a, 0x15, 0x4d, 0x75, 0x6c,
	0x74, 0x69, 0x70, 0x6c, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x22, 0x2d, 0x0a, 0x11,
	0x50, 0x72, 0x6f, 0x62, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x22, 0x27, 0x0a, 0x0b, 0x41,
	0x64, 0x64, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x2d, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x64,
	0x64, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x22, 0x2c, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73,
	0x74, 0x32, 0xc3, 0x03, 0x0a, 0x0a, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x68, 0x61, 0x69, 0x6e,
	0x12, 0x35, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x32, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x68,
	0x61, 0x69, 0x6e, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x2d, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x55, 0x74, 0x78,
	0x6f, 0x53, 0x65, 0x74, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x74, 0x78,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x55, 0x74, 0x78, 0x6f, 0x12, 0x3f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4d, 0x75, 0x6c, 0x74,
	0x69, 0x70, 0x6c, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x3e, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x79, 0x49, 0x64, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x62,
	0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x32, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x64, 0x64,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64,
	0x64, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x0e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_blockchain_v2_blockchain_proto_rawDescOnce sync.Once
	file_proto_blockchain_v2_blockchain_proto_rawDescData = file_proto_blockchain_v2_blockchain_proto_rawDesc
)

func file_proto_blockchain_v2_blockchain_proto_rawDescGZIP() []byte {
	file_proto_blockchain_v2_blockchain_proto_rawDescOnce.Do(func() {
		file_proto_blockchain_v2_blockchain_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_blockchain_v2_blockchain_proto_rawDescData)
	})
	return file_proto_blockchain_v2_blockchain_proto_rawDescData
}

var file_proto_blockchain_v2_blockchain_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_proto_blockchain_v2_blockchain_proto_goTypes = []interface{}{
	(*Transaction)(nil),           // 0: proto.Transaction
	(*Block)(nil),                 // 1: proto.Block
	(*Response)(nil),              // 2: proto.response
	(*ChainRequest)(nil),          // 3: proto.ChainRequest
	(*UtxoRequest)(nil),           // 4: proto.UtxoRequest
	(*Chain)(nil),                 // 5: proto.Chain
	(*ProbeBlock)(nil),            // 6: proto.ProbeBlock
	(*Utxo)(nil),                  // 7: proto.Utxo
	(*MultipleBlocksRequest)(nil), // 8: proto.MultipleBlocksRequest
	(*ProbeBlockRequest)(nil),     // 9: proto.ProbeBlockRequest
	(*AddrRequest)(nil),           // 10: proto.AddrRequest
	(*UpdateAddrRequest)(nil),     // 11: proto.UpdateAddrRequest
	(*AddrList)(nil),              // 12: proto.AddrList
}
var file_proto_blockchain_v2_blockchain_proto_depIdxs = []int32{
	0,  // 0: proto.BlockChain.GetTransaction:input_type -> proto.Transaction
	1,  // 1: proto.BlockChain.GetBlock:input_type -> proto.Block
	3,  // 2: proto.BlockChain.GetBlockChain:input_type -> proto.ChainRequest
	4,  // 3: proto.BlockChain.GetUtxoSet:input_type -> proto.UtxoRequest
	8,  // 4: proto.BlockChain.GetMultipleBlocks:input_type -> proto.MultipleBlocksRequest
	9,  // 5: proto.BlockChain.GetOneBlockById:input_type -> proto.ProbeBlockRequest
	10, // 6: proto.BlockChain.GetAddrList:input_type -> proto.AddrRequest
	11, // 7: proto.BlockChain.UpdateAddrList:input_type -> proto.UpdateAddrRequest
	2,  // 8: proto.BlockChain.GetTransaction:output_type -> proto.response
	2,  // 9: proto.BlockChain.GetBlock:output_type -> proto.response
	5,  // 10: proto.BlockChain.GetBlockChain:output_type -> proto.Chain
	7,  // 11: proto.BlockChain.GetUtxoSet:output_type -> proto.Utxo
	5,  // 12: proto.BlockChain.GetMultipleBlocks:output_type -> proto.Chain
	6,  // 13: proto.BlockChain.GetOneBlockById:output_type -> proto.ProbeBlock
	12, // 14: proto.BlockChain.GetAddrList:output_type -> proto.AddrList
	2,  // 15: proto.BlockChain.UpdateAddrList:output_type -> proto.response
	8,  // [8:16] is the sub-list for method output_type
	0,  // [0:8] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_proto_blockchain_v2_blockchain_proto_init() }
func file_proto_blockchain_v2_blockchain_proto_init() {
	if File_proto_blockchain_v2_blockchain_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_blockchain_v2_blockchain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file_proto_blockchain_v2_blockchain_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
		file_proto_blockchain_v2_blockchain_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_proto_blockchain_v2_blockchain_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChainRequest); i {
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
		file_proto_blockchain_v2_blockchain_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UtxoRequest); i {
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
		file_proto_blockchain_v2_blockchain_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chain); i {
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
		file_proto_blockchain_v2_blockchain_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProbeBlock); i {
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
		file_proto_blockchain_v2_blockchain_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Utxo); i {
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
		file_proto_blockchain_v2_blockchain_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultipleBlocksRequest); i {
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
		file_proto_blockchain_v2_blockchain_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProbeBlockRequest); i {
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
		file_proto_blockchain_v2_blockchain_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddrRequest); i {
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
		file_proto_blockchain_v2_blockchain_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAddrRequest); i {
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
		file_proto_blockchain_v2_blockchain_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddrList); i {
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
			RawDescriptor: file_proto_blockchain_v2_blockchain_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_blockchain_v2_blockchain_proto_goTypes,
		DependencyIndexes: file_proto_blockchain_v2_blockchain_proto_depIdxs,
		MessageInfos:      file_proto_blockchain_v2_blockchain_proto_msgTypes,
	}.Build()
	File_proto_blockchain_v2_blockchain_proto = out.File
	file_proto_blockchain_v2_blockchain_proto_rawDesc = nil
	file_proto_blockchain_v2_blockchain_proto_goTypes = nil
	file_proto_blockchain_v2_blockchain_proto_depIdxs = nil
}