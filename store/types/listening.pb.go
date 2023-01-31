// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/store/v1beta1/listening.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/tendermint/tendermint/abci/types"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// StoreKVPair is a KVStore KVPair used for listening to state changes (Sets and Deletes)
// It optionally includes the StoreKey for the originating KVStore and a Boolean flag to distinguish between Sets and
// Deletes
//
// Since: cosmos-sdk 0.43
type StoreKVPair struct {
	StoreKey string `protobuf:"bytes,1,opt,name=store_key,json=storeKey,proto3" json:"store_key,omitempty"`
	Delete   bool   `protobuf:"varint,2,opt,name=delete,proto3" json:"delete,omitempty"`
	Key      []byte `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Value    []byte `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *StoreKVPair) Reset()         { *m = StoreKVPair{} }
func (m *StoreKVPair) String() string { return proto.CompactTextString(m) }
func (*StoreKVPair) ProtoMessage()    {}
func (*StoreKVPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6caeb9d7b7c7c10, []int{0}
}
func (m *StoreKVPair) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StoreKVPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StoreKVPair.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StoreKVPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreKVPair.Merge(m, src)
}
func (m *StoreKVPair) XXX_Size() int {
	return m.Size()
}
func (m *StoreKVPair) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreKVPair.DiscardUnknown(m)
}

var xxx_messageInfo_StoreKVPair proto.InternalMessageInfo

func (m *StoreKVPair) GetStoreKey() string {
	if m != nil {
		return m.StoreKey
	}
	return ""
}

func (m *StoreKVPair) GetDelete() bool {
	if m != nil {
		return m.Delete
	}
	return false
}

func (m *StoreKVPair) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *StoreKVPair) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

// BlockMetadata contains all the abci event data of a block
// the file streamer dump them into files together with the state changes.
type BlockMetadata struct {
	RequestBeginBlock  *types.RequestBeginBlock   `protobuf:"bytes,1,opt,name=request_begin_block,json=requestBeginBlock,proto3" json:"request_begin_block,omitempty"`
	ResponseBeginBlock *types.ResponseBeginBlock  `protobuf:"bytes,2,opt,name=response_begin_block,json=responseBeginBlock,proto3" json:"response_begin_block,omitempty"`
	DeliverTxs         []*BlockMetadata_DeliverTx `protobuf:"bytes,3,rep,name=deliver_txs,json=deliverTxs,proto3" json:"deliver_txs,omitempty"`
	RequestEndBlock    *types.RequestEndBlock     `protobuf:"bytes,4,opt,name=request_end_block,json=requestEndBlock,proto3" json:"request_end_block,omitempty"`
	ResponseEndBlock   *types.ResponseEndBlock    `protobuf:"bytes,5,opt,name=response_end_block,json=responseEndBlock,proto3" json:"response_end_block,omitempty"`
	ResponseCommit     *types.ResponseCommit      `protobuf:"bytes,6,opt,name=response_commit,json=responseCommit,proto3" json:"response_commit,omitempty"`
}

func (m *BlockMetadata) Reset()         { *m = BlockMetadata{} }
func (m *BlockMetadata) String() string { return proto.CompactTextString(m) }
func (*BlockMetadata) ProtoMessage()    {}
func (*BlockMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6caeb9d7b7c7c10, []int{1}
}
func (m *BlockMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockMetadata.Merge(m, src)
}
func (m *BlockMetadata) XXX_Size() int {
	return m.Size()
}
func (m *BlockMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_BlockMetadata proto.InternalMessageInfo

func (m *BlockMetadata) GetRequestBeginBlock() *types.RequestBeginBlock {
	if m != nil {
		return m.RequestBeginBlock
	}
	return nil
}

func (m *BlockMetadata) GetResponseBeginBlock() *types.ResponseBeginBlock {
	if m != nil {
		return m.ResponseBeginBlock
	}
	return nil
}

func (m *BlockMetadata) GetDeliverTxs() []*BlockMetadata_DeliverTx {
	if m != nil {
		return m.DeliverTxs
	}
	return nil
}

func (m *BlockMetadata) GetRequestEndBlock() *types.RequestEndBlock {
	if m != nil {
		return m.RequestEndBlock
	}
	return nil
}

func (m *BlockMetadata) GetResponseEndBlock() *types.ResponseEndBlock {
	if m != nil {
		return m.ResponseEndBlock
	}
	return nil
}

func (m *BlockMetadata) GetResponseCommit() *types.ResponseCommit {
	if m != nil {
		return m.ResponseCommit
	}
	return nil
}

// DeliverTx encapulate deliver tx request and response.
type BlockMetadata_DeliverTx struct {
	Request  *types.RequestDeliverTx  `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Response *types.ResponseDeliverTx `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
}

func (m *BlockMetadata_DeliverTx) Reset()         { *m = BlockMetadata_DeliverTx{} }
func (m *BlockMetadata_DeliverTx) String() string { return proto.CompactTextString(m) }
func (*BlockMetadata_DeliverTx) ProtoMessage()    {}
func (*BlockMetadata_DeliverTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6caeb9d7b7c7c10, []int{1, 0}
}
func (m *BlockMetadata_DeliverTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockMetadata_DeliverTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockMetadata_DeliverTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockMetadata_DeliverTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockMetadata_DeliverTx.Merge(m, src)
}
func (m *BlockMetadata_DeliverTx) XXX_Size() int {
	return m.Size()
}
func (m *BlockMetadata_DeliverTx) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockMetadata_DeliverTx.DiscardUnknown(m)
}

var xxx_messageInfo_BlockMetadata_DeliverTx proto.InternalMessageInfo

func (m *BlockMetadata_DeliverTx) GetRequest() *types.RequestDeliverTx {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *BlockMetadata_DeliverTx) GetResponse() *types.ResponseDeliverTx {
	if m != nil {
		return m.Response
	}
	return nil
}

func init() {
	proto.RegisterType((*StoreKVPair)(nil), "cosmos.store.v1beta1.StoreKVPair")
	proto.RegisterType((*BlockMetadata)(nil), "cosmos.store.v1beta1.BlockMetadata")
	proto.RegisterType((*BlockMetadata_DeliverTx)(nil), "cosmos.store.v1beta1.BlockMetadata.DeliverTx")
}

func init() {
	proto.RegisterFile("cosmos/store/v1beta1/listening.proto", fileDescriptor_b6caeb9d7b7c7c10)
}

var fileDescriptor_b6caeb9d7b7c7c10 = []byte{
	// 458 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x9b, 0xb5, 0x2b, 0xed, 0x5b, 0x60, 0xc3, 0x54, 0x28, 0xda, 0xa4, 0x90, 0x15, 0x0e,
	0xbd, 0xe0, 0x68, 0xe5, 0x88, 0xc4, 0xa1, 0x80, 0x84, 0x34, 0xfe, 0x29, 0xfc, 0x39, 0x70, 0x89,
	0x92, 0xe6, 0xd5, 0x64, 0x9a, 0xda, 0xc5, 0xf6, 0xaa, 0xf5, 0x1b, 0x70, 0xe4, 0x63, 0x71, 0xdc,
	0x91, 0x23, 0x6a, 0xbf, 0x07, 0x42, 0xb1, 0xd3, 0x74, 0x19, 0xcd, 0xcd, 0x7e, 0xf2, 0x3c, 0xbf,
	0x3c, 0xaf, 0x2d, 0xc3, 0xe3, 0x89, 0x50, 0x33, 0xa1, 0x02, 0xa5, 0x85, 0xc4, 0x60, 0x71, 0x9a,
	0xa0, 0x8e, 0x4f, 0x83, 0x8c, 0x29, 0x8d, 0x9c, 0xf1, 0x73, 0x3a, 0x97, 0x42, 0x0b, 0xd2, 0xb7,
	0x2e, 0x6a, 0x5c, 0xb4, 0x70, 0x1d, 0x1d, 0x6b, 0xe4, 0x29, 0xca, 0x19, 0xe3, 0x3a, 0x88, 0x93,
	0x09, 0x0b, 0xf4, 0x72, 0x8e, 0xca, 0x46, 0x06, 0xdf, 0xa0, 0xf7, 0x31, 0x77, 0x9f, 0x7d, 0xf9,
	0x10, 0x33, 0x49, 0x8e, 0xa1, 0x6b, 0xc2, 0xd1, 0x14, 0x97, 0xae, 0xe3, 0x3b, 0xc3, 0x6e, 0xd8,
	0x31, 0xc2, 0x19, 0x2e, 0xc9, 0x03, 0x68, 0xa7, 0x98, 0xa1, 0x46, 0x77, 0xcf, 0x77, 0x86, 0x9d,
	0xb0, 0xd8, 0x91, 0x43, 0x68, 0xe6, 0xf6, 0xa6, 0xef, 0x0c, 0x6f, 0x87, 0xf9, 0x92, 0xf4, 0x61,
	0x7f, 0x11, 0x67, 0x17, 0xe8, 0xb6, 0x8c, 0x66, 0x37, 0x83, 0xbf, 0x2d, 0xb8, 0x33, 0xce, 0xc4,
	0x64, 0xfa, 0x16, 0x75, 0x9c, 0xc6, 0x3a, 0x26, 0x21, 0xdc, 0x97, 0xf8, 0xfd, 0x02, 0x95, 0x8e,
	0x12, 0x3c, 0x67, 0x3c, 0x4a, 0xf2, 0xcf, 0xe6, 0xc7, 0xbd, 0xd1, 0x80, 0x6e, 0x8b, 0xd3, 0xbc,
	0x38, 0x0d, 0xad, 0x77, 0x9c, 0x5b, 0x0d, 0x28, 0xbc, 0x27, 0x6f, 0x4a, 0xe4, 0x33, 0xf4, 0x25,
	0xaa, 0xb9, 0xe0, 0x0a, 0x2b, 0xd0, 0x3d, 0x03, 0x7d, 0xb4, 0x03, 0x6a, 0xcd, 0xd7, 0xa8, 0x44,
	0xfe, 0xa7, 0x91, 0x77, 0xd0, 0x4b, 0x31, 0x63, 0x0b, 0x94, 0x91, 0xbe, 0x54, 0x6e, 0xd3, 0x6f,
	0x0e, 0x7b, 0xa3, 0x27, 0x74, 0xd7, 0x89, 0xd3, 0xca, 0x90, 0xf4, 0xa5, 0x8d, 0x7d, 0xba, 0x0c,
	0x21, 0xdd, 0x2c, 0x15, 0x79, 0x03, 0x9b, 0xee, 0x11, 0xf2, 0xb4, 0xe8, 0xd8, 0x32, 0x1d, 0xfd,
	0xba, 0xc1, 0x5f, 0xf1, 0xd4, 0x16, 0x3c, 0x90, 0x55, 0x81, 0xbc, 0x87, 0xb2, 0xf3, 0x35, 0xdc,
	0xbe, 0xc1, 0x9d, 0xd4, 0x8e, 0x5c, 0xf2, 0x0e, 0xe5, 0x0d, 0x85, 0xbc, 0x86, 0x83, 0x12, 0x38,
	0x11, 0xb3, 0x19, 0xd3, 0x6e, 0xdb, 0xd0, 0x1e, 0xd6, 0xd2, 0x5e, 0x18, 0x5b, 0x78, 0x57, 0x56,
	0xf6, 0x47, 0x3f, 0x1c, 0xe8, 0x96, 0x47, 0x40, 0x9e, 0xc1, 0xad, 0xa2, 0x7b, 0x71, 0xcb, 0x27,
	0x75, 0xc3, 0x6e, 0x8f, 0x6d, 0x93, 0x20, 0xcf, 0xa1, 0xb3, 0x81, 0x17, 0xd7, 0x39, 0xa8, 0x6d,
	0xb3, 0x8d, 0x97, 0x99, 0xf1, 0xe8, 0xd7, 0xca, 0x73, 0xae, 0x56, 0x9e, 0xf3, 0x67, 0xe5, 0x39,
	0x3f, 0xd7, 0x5e, 0xe3, 0x6a, 0xed, 0x35, 0x7e, 0xaf, 0xbd, 0xc6, 0x57, 0xd7, 0xde, 0xa3, 0x4a,
	0xa7, 0x94, 0x89, 0xe2, 0x95, 0x99, 0x67, 0x92, 0xb4, 0xcd, 0x3b, 0x79, 0xfa, 0x2f, 0x00, 0x00,
	0xff, 0xff, 0x2b, 0x2f, 0x77, 0x06, 0x82, 0x03, 0x00, 0x00,
}

func (m *StoreKVPair) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StoreKVPair) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StoreKVPair) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintListening(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintListening(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Delete {
		i--
		if m.Delete {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.StoreKey) > 0 {
		i -= len(m.StoreKey)
		copy(dAtA[i:], m.StoreKey)
		i = encodeVarintListening(dAtA, i, uint64(len(m.StoreKey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BlockMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ResponseCommit != nil {
		{
			size, err := m.ResponseCommit.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintListening(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.ResponseEndBlock != nil {
		{
			size, err := m.ResponseEndBlock.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintListening(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.RequestEndBlock != nil {
		{
			size, err := m.RequestEndBlock.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintListening(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.DeliverTxs) > 0 {
		for iNdEx := len(m.DeliverTxs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DeliverTxs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintListening(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.ResponseBeginBlock != nil {
		{
			size, err := m.ResponseBeginBlock.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintListening(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.RequestBeginBlock != nil {
		{
			size, err := m.RequestBeginBlock.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintListening(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BlockMetadata_DeliverTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockMetadata_DeliverTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockMetadata_DeliverTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Response != nil {
		{
			size, err := m.Response.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintListening(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Request != nil {
		{
			size, err := m.Request.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintListening(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintListening(dAtA []byte, offset int, v uint64) int {
	offset -= sovListening(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StoreKVPair) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.StoreKey)
	if l > 0 {
		n += 1 + l + sovListening(uint64(l))
	}
	if m.Delete {
		n += 2
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovListening(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovListening(uint64(l))
	}
	return n
}

func (m *BlockMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RequestBeginBlock != nil {
		l = m.RequestBeginBlock.Size()
		n += 1 + l + sovListening(uint64(l))
	}
	if m.ResponseBeginBlock != nil {
		l = m.ResponseBeginBlock.Size()
		n += 1 + l + sovListening(uint64(l))
	}
	if len(m.DeliverTxs) > 0 {
		for _, e := range m.DeliverTxs {
			l = e.Size()
			n += 1 + l + sovListening(uint64(l))
		}
	}
	if m.RequestEndBlock != nil {
		l = m.RequestEndBlock.Size()
		n += 1 + l + sovListening(uint64(l))
	}
	if m.ResponseEndBlock != nil {
		l = m.ResponseEndBlock.Size()
		n += 1 + l + sovListening(uint64(l))
	}
	if m.ResponseCommit != nil {
		l = m.ResponseCommit.Size()
		n += 1 + l + sovListening(uint64(l))
	}
	return n
}

func (m *BlockMetadata_DeliverTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Request != nil {
		l = m.Request.Size()
		n += 1 + l + sovListening(uint64(l))
	}
	if m.Response != nil {
		l = m.Response.Size()
		n += 1 + l + sovListening(uint64(l))
	}
	return n
}

func sovListening(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozListening(x uint64) (n int) {
	return sovListening(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StoreKVPair) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowListening
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StoreKVPair: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StoreKVPair: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoreKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListening
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthListening
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthListening
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StoreKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delete", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListening
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Delete = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListening
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthListening
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthListening
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListening
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthListening
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthListening
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipListening(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthListening
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BlockMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowListening
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BlockMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestBeginBlock", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListening
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthListening
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthListening
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RequestBeginBlock == nil {
				m.RequestBeginBlock = &types.RequestBeginBlock{}
			}
			if err := m.RequestBeginBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResponseBeginBlock", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListening
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthListening
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthListening
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ResponseBeginBlock == nil {
				m.ResponseBeginBlock = &types.ResponseBeginBlock{}
			}
			if err := m.ResponseBeginBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeliverTxs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListening
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthListening
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthListening
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeliverTxs = append(m.DeliverTxs, &BlockMetadata_DeliverTx{})
			if err := m.DeliverTxs[len(m.DeliverTxs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestEndBlock", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListening
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthListening
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthListening
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RequestEndBlock == nil {
				m.RequestEndBlock = &types.RequestEndBlock{}
			}
			if err := m.RequestEndBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResponseEndBlock", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListening
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthListening
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthListening
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ResponseEndBlock == nil {
				m.ResponseEndBlock = &types.ResponseEndBlock{}
			}
			if err := m.ResponseEndBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResponseCommit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListening
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthListening
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthListening
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ResponseCommit == nil {
				m.ResponseCommit = &types.ResponseCommit{}
			}
			if err := m.ResponseCommit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipListening(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthListening
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BlockMetadata_DeliverTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowListening
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DeliverTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeliverTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Request", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListening
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthListening
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthListening
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Request == nil {
				m.Request = &types.RequestDeliverTx{}
			}
			if err := m.Request.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Response", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListening
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthListening
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthListening
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Response == nil {
				m.Response = &types.ResponseDeliverTx{}
			}
			if err := m.Response.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipListening(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthListening
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipListening(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowListening
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowListening
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowListening
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthListening
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupListening
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthListening
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthListening        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowListening          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupListening = fmt.Errorf("proto: unexpected end of group")
)
