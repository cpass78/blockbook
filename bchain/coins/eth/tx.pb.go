// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tx.proto

/*
Package eth is a generated protocol buffer package.

It is generated from these files:
	tx.proto

It has these top-level messages:
	ProtoTransaction
*/
package eth

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ProtoTransaction struct {
	AccountNonce     uint64 `protobuf:"varint,1,opt,name=AccountNonce" json:"AccountNonce,omitempty"`
	Price            []byte `protobuf:"bytes,2,opt,name=Price,proto3" json:"Price,omitempty"`
	GasLimit         uint64 `protobuf:"varint,3,opt,name=GasLimit" json:"GasLimit,omitempty"`
	Value            []byte `protobuf:"bytes,4,opt,name=Value,proto3" json:"Value,omitempty"`
	Payload          []byte `protobuf:"bytes,5,opt,name=Payload,proto3" json:"Payload,omitempty"`
	Hash             []byte `protobuf:"bytes,6,opt,name=Hash,proto3" json:"Hash,omitempty"`
	BlockNumber      uint32 `protobuf:"varint,7,opt,name=BlockNumber" json:"BlockNumber,omitempty"`
	BlockTime        uint64 `protobuf:"varint,8,opt,name=BlockTime" json:"BlockTime,omitempty"`
	To               []byte `protobuf:"bytes,9,opt,name=To,proto3" json:"To,omitempty"`
	From             []byte `protobuf:"bytes,10,opt,name=From,proto3" json:"From,omitempty"`
	TransactionIndex uint32 `protobuf:"varint,11,opt,name=TransactionIndex" json:"TransactionIndex,omitempty"`
	V                []byte `protobuf:"bytes,12,opt,name=V,proto3" json:"V,omitempty"`
	R                []byte `protobuf:"bytes,13,opt,name=R,proto3" json:"R,omitempty"`
	S                []byte `protobuf:"bytes,14,opt,name=S,proto3" json:"S,omitempty"`
}

func (m *ProtoTransaction) Reset()                    { *m = ProtoTransaction{} }
func (m *ProtoTransaction) String() string            { return proto.CompactTextString(m) }
func (*ProtoTransaction) ProtoMessage()               {}
func (*ProtoTransaction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ProtoTransaction) GetAccountNonce() uint64 {
	if m != nil {
		return m.AccountNonce
	}
	return 0
}

func (m *ProtoTransaction) GetPrice() []byte {
	if m != nil {
		return m.Price
	}
	return nil
}

func (m *ProtoTransaction) GetGasLimit() uint64 {
	if m != nil {
		return m.GasLimit
	}
	return 0
}

func (m *ProtoTransaction) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *ProtoTransaction) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *ProtoTransaction) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *ProtoTransaction) GetBlockNumber() uint32 {
	if m != nil {
		return m.BlockNumber
	}
	return 0
}

func (m *ProtoTransaction) GetBlockTime() uint64 {
	if m != nil {
		return m.BlockTime
	}
	return 0
}

func (m *ProtoTransaction) GetTo() []byte {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *ProtoTransaction) GetFrom() []byte {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *ProtoTransaction) GetTransactionIndex() uint32 {
	if m != nil {
		return m.TransactionIndex
	}
	return 0
}

func (m *ProtoTransaction) GetV() []byte {
	if m != nil {
		return m.V
	}
	return nil
}

func (m *ProtoTransaction) GetR() []byte {
	if m != nil {
		return m.R
	}
	return nil
}

func (m *ProtoTransaction) GetS() []byte {
	if m != nil {
		return m.S
	}
	return nil
}

func init() {
	proto.RegisterType((*ProtoTransaction)(nil), "eth.ProtoTransaction")
}

func init() { proto.RegisterFile("tx.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xbd, 0x6a, 0xeb, 0x40,
	0x10, 0x85, 0x59, 0x59, 0xb6, 0xe5, 0xb1, 0x6c, 0xcc, 0x70, 0x8b, 0xe1, 0x92, 0x42, 0xb8, 0x12,
	0x29, 0xd2, 0xe4, 0x09, 0x92, 0x22, 0x3f, 0x10, 0x8c, 0x90, 0x85, 0xfa, 0xf5, 0x7a, 0xc1, 0x22,
	0x92, 0x26, 0x48, 0x2b, 0x70, 0x5e, 0x38, 0xcf, 0x11, 0x76, 0x44, 0x12, 0x87, 0x74, 0xf3, 0x7d,
	0x70, 0xf6, 0x2c, 0x07, 0x22, 0x77, 0xbe, 0x79, 0xeb, 0xd8, 0x31, 0x4e, 0xac, 0x3b, 0x6d, 0x3f,
	0x02, 0xd8, 0x64, 0x1e, 0x8b, 0x4e, 0xb7, 0xbd, 0x36, 0xae, 0xe2, 0x16, 0xb7, 0x10, 0xdf, 0x19,
	0xc3, 0x43, 0xeb, 0x76, 0xdc, 0x1a, 0x4b, 0x2a, 0x51, 0x69, 0x98, 0xff, 0x72, 0xf8, 0x0f, 0xa6,
	0x59, 0x57, 0x19, 0x4b, 0x41, 0xa2, 0xd2, 0x38, 0x1f, 0x01, 0xff, 0x43, 0xf4, 0xa8, 0xfb, 0x97,
	0xaa, 0xa9, 0x1c, 0x4d, 0x24, 0xf5, 0xcd, 0x3e, 0x51, 0xea, 0x7a, 0xb0, 0x14, 0x8e, 0x09, 0x01,
	0x24, 0x98, 0x67, 0xfa, 0xbd, 0x66, 0x7d, 0xa4, 0xa9, 0xf8, 0x2f, 0x44, 0x84, 0xf0, 0x49, 0xf7,
	0x27, 0x9a, 0x89, 0x96, 0x1b, 0x13, 0x58, 0xde, 0xd7, 0x6c, 0x5e, 0x77, 0x43, 0x73, 0xb0, 0x1d,
	0xcd, 0x13, 0x95, 0xae, 0xf2, 0x4b, 0x85, 0x57, 0xb0, 0x10, 0x2c, 0xaa, 0xc6, 0x52, 0x24, 0x5f,
	0xf8, 0x11, 0xb8, 0x86, 0xa0, 0x60, 0x5a, 0xc8, 0x8b, 0x41, 0xc1, 0xbe, 0xe3, 0xa1, 0xe3, 0x86,
	0x60, 0xec, 0xf0, 0x37, 0x5e, 0xc3, 0xe6, 0x62, 0x8c, 0xe7, 0xf6, 0x68, 0xcf, 0xb4, 0x94, 0xa2,
	0x3f, 0x1e, 0x63, 0x50, 0x25, 0xc5, 0x12, 0x56, 0xa5, 0xa7, 0x9c, 0x56, 0x23, 0xe5, 0x9e, 0xf6,
	0xb4, 0x1e, 0x69, 0x7f, 0x98, 0xc9, 0xe8, 0xb7, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x15, 0xc8,
	0xe4, 0x30, 0x80, 0x01, 0x00, 0x00,
}