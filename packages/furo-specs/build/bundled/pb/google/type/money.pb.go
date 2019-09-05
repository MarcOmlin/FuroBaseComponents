// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/type/money.proto

package google_type

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Represents an amount of money with its currency type. https://github.com/googleapis/googleapis/blob/master/google/money.proto
type Money struct {
	// The 3-letter currency code defined in ISO 4217.
	CurrencyCode string `protobuf:"bytes,2,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	// String representation of money entity
	DisplayName string `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Number of nano (10^-9) units of the amount. For example $-1.75 is represented as `units`=-1 and `nanos`=-750,000,000.
	Nanos int64 `protobuf:"varint,4,opt,name=nanos,proto3" json:"nanos,omitempty"`
	// The whole units of the amount.
	Units                int64    `protobuf:"varint,3,opt,name=units,proto3" json:"units,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Money) Reset()         { *m = Money{} }
func (m *Money) String() string { return proto.CompactTextString(m) }
func (*Money) ProtoMessage()    {}
func (*Money) Descriptor() ([]byte, []int) {
	return fileDescriptor_f093d3d05ab4bbee, []int{0}
}

func (m *Money) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Money.Unmarshal(m, b)
}
func (m *Money) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Money.Marshal(b, m, deterministic)
}
func (m *Money) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Money.Merge(m, src)
}
func (m *Money) XXX_Size() int {
	return xxx_messageInfo_Money.Size(m)
}
func (m *Money) XXX_DiscardUnknown() {
	xxx_messageInfo_Money.DiscardUnknown(m)
}

var xxx_messageInfo_Money proto.InternalMessageInfo

func (m *Money) GetCurrencyCode() string {
	if m != nil {
		return m.CurrencyCode
	}
	return ""
}

func (m *Money) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Money) GetNanos() int64 {
	if m != nil {
		return m.Nanos
	}
	return 0
}

func (m *Money) GetUnits() int64 {
	if m != nil {
		return m.Units
	}
	return 0
}

func init() {
	proto.RegisterType((*Money)(nil), "google.type.Money")
}

func init() { proto.RegisterFile("google/type/money.proto", fileDescriptor_f093d3d05ab4bbee) }

var fileDescriptor_f093d3d05ab4bbee = []byte{
	// 156 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0xcf, 0xcd, 0xcf, 0x4b, 0xad, 0xd4, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x86, 0x48, 0xe8, 0x81, 0x24, 0x94, 0xaa, 0xb9, 0x58, 0x7d, 0x41,
	0x72, 0x42, 0xca, 0x5c, 0xbc, 0xc9, 0xa5, 0x45, 0x45, 0xa9, 0x79, 0xc9, 0x95, 0xf1, 0xc9, 0xf9,
	0x29, 0xa9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x3c, 0x30, 0x41, 0xe7, 0xfc, 0x94, 0x54,
	0x21, 0x45, 0x2e, 0x9e, 0x94, 0xcc, 0xe2, 0x82, 0x9c, 0xc4, 0xca, 0xf8, 0xbc, 0xc4, 0xdc, 0x54,
	0x09, 0x46, 0xb0, 0x1a, 0x6e, 0xa8, 0x98, 0x5f, 0x62, 0x6e, 0xaa, 0x90, 0x08, 0x17, 0x6b, 0x5e,
	0x62, 0x5e, 0x7e, 0xb1, 0x04, 0x8b, 0x02, 0xa3, 0x06, 0x73, 0x10, 0x84, 0x03, 0x12, 0x2d, 0xcd,
	0xcb, 0x2c, 0x29, 0x96, 0x60, 0x86, 0x88, 0x82, 0x39, 0x49, 0x6c, 0x60, 0x07, 0x19, 0x03, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x17, 0x1b, 0xa1, 0x7c, 0xab, 0x00, 0x00, 0x00,
}
