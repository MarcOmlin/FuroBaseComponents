// Code generated by protoc-gen-go. DO NOT EDIT.
// source: furo/property.proto

package furo

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// String type to use in property
type StringOptionProperty struct {
	// String representation of val
	DisplayName string `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// The value, Id is used to make working with data-inputs easier
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringOptionProperty) Reset()         { *m = StringOptionProperty{} }
func (m *StringOptionProperty) String() string { return proto.CompactTextString(m) }
func (*StringOptionProperty) ProtoMessage()    {}
func (*StringOptionProperty) Descriptor() ([]byte, []int) {
	return fileDescriptor_bca1bac703f4631e, []int{0}
}

func (m *StringOptionProperty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringOptionProperty.Unmarshal(m, b)
}
func (m *StringOptionProperty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringOptionProperty.Marshal(b, m, deterministic)
}
func (m *StringOptionProperty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringOptionProperty.Merge(m, src)
}
func (m *StringOptionProperty) XXX_Size() int {
	return xxx_messageInfo_StringOptionProperty.Size(m)
}
func (m *StringOptionProperty) XXX_DiscardUnknown() {
	xxx_messageInfo_StringOptionProperty.DiscardUnknown(m)
}

var xxx_messageInfo_StringOptionProperty proto.InternalMessageInfo

func (m *StringOptionProperty) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *StringOptionProperty) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// Number type with embedded meta
type NumberProperty struct {
	// data part
	Data                 float32  `protobuf:"fixed32,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NumberProperty) Reset()         { *m = NumberProperty{} }
func (m *NumberProperty) String() string { return proto.CompactTextString(m) }
func (*NumberProperty) ProtoMessage()    {}
func (*NumberProperty) Descriptor() ([]byte, []int) {
	return fileDescriptor_bca1bac703f4631e, []int{1}
}

func (m *NumberProperty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NumberProperty.Unmarshal(m, b)
}
func (m *NumberProperty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NumberProperty.Marshal(b, m, deterministic)
}
func (m *NumberProperty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NumberProperty.Merge(m, src)
}
func (m *NumberProperty) XXX_Size() int {
	return xxx_messageInfo_NumberProperty.Size(m)
}
func (m *NumberProperty) XXX_DiscardUnknown() {
	xxx_messageInfo_NumberProperty.DiscardUnknown(m)
}

var xxx_messageInfo_NumberProperty proto.InternalMessageInfo

func (m *NumberProperty) GetData() float32 {
	if m != nil {
		return m.Data
	}
	return 0
}

// String type to use in property
type StringProperty struct {
	// data part
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringProperty) Reset()         { *m = StringProperty{} }
func (m *StringProperty) String() string { return proto.CompactTextString(m) }
func (*StringProperty) ProtoMessage()    {}
func (*StringProperty) Descriptor() ([]byte, []int) {
	return fileDescriptor_bca1bac703f4631e, []int{2}
}

func (m *StringProperty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringProperty.Unmarshal(m, b)
}
func (m *StringProperty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringProperty.Marshal(b, m, deterministic)
}
func (m *StringProperty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringProperty.Merge(m, src)
}
func (m *StringProperty) XXX_Size() int {
	return xxx_messageInfo_StringProperty.Size(m)
}
func (m *StringProperty) XXX_DiscardUnknown() {
	xxx_messageInfo_StringProperty.DiscardUnknown(m)
}

var xxx_messageInfo_StringProperty proto.InternalMessageInfo

func (m *StringProperty) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

// Integer type with embedded meta
type IntegerProperty struct {
	// Integer data part
	Data                 int32    `protobuf:"varint,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntegerProperty) Reset()         { *m = IntegerProperty{} }
func (m *IntegerProperty) String() string { return proto.CompactTextString(m) }
func (*IntegerProperty) ProtoMessage()    {}
func (*IntegerProperty) Descriptor() ([]byte, []int) {
	return fileDescriptor_bca1bac703f4631e, []int{3}
}

func (m *IntegerProperty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntegerProperty.Unmarshal(m, b)
}
func (m *IntegerProperty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntegerProperty.Marshal(b, m, deterministic)
}
func (m *IntegerProperty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntegerProperty.Merge(m, src)
}
func (m *IntegerProperty) XXX_Size() int {
	return xxx_messageInfo_IntegerProperty.Size(m)
}
func (m *IntegerProperty) XXX_DiscardUnknown() {
	xxx_messageInfo_IntegerProperty.DiscardUnknown(m)
}

var xxx_messageInfo_IntegerProperty proto.InternalMessageInfo

func (m *IntegerProperty) GetData() int32 {
	if m != nil {
		return m.Data
	}
	return 0
}

// Type to define property values with type information
type Property struct {
	// property code for additional settings
	Code string `protobuf:"bytes,5,opt,name=code,proto3" json:"code,omitempty"`
	// data part of the property
	Data *any.Any `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	// String representation of the property
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Id of the property
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Meta for the response
	Meta                 *Meta    `protobuf:"bytes,4,opt,name=meta,proto3" json:"meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Property) Reset()         { *m = Property{} }
func (m *Property) String() string { return proto.CompactTextString(m) }
func (*Property) ProtoMessage()    {}
func (*Property) Descriptor() ([]byte, []int) {
	return fileDescriptor_bca1bac703f4631e, []int{4}
}

func (m *Property) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Property.Unmarshal(m, b)
}
func (m *Property) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Property.Marshal(b, m, deterministic)
}
func (m *Property) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Property.Merge(m, src)
}
func (m *Property) XXX_Size() int {
	return xxx_messageInfo_Property.Size(m)
}
func (m *Property) XXX_DiscardUnknown() {
	xxx_messageInfo_Property.DiscardUnknown(m)
}

var xxx_messageInfo_Property proto.InternalMessageInfo

func (m *Property) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Property) GetData() *any.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Property) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Property) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Property) GetMeta() *Meta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func init() {
	proto.RegisterType((*StringOptionProperty)(nil), "furo.StringOptionProperty")
	proto.RegisterType((*NumberProperty)(nil), "furo.NumberProperty")
	proto.RegisterType((*StringProperty)(nil), "furo.StringProperty")
	proto.RegisterType((*IntegerProperty)(nil), "furo.IntegerProperty")
	proto.RegisterType((*Property)(nil), "furo.Property")
}

func init() { proto.RegisterFile("furo/property.proto", fileDescriptor_bca1bac703f4631e) }

var fileDescriptor_bca1bac703f4631e = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xcd, 0x4a, 0x03, 0x31,
	0x14, 0x85, 0x49, 0x9c, 0x8a, 0xbd, 0x95, 0x16, 0x62, 0x17, 0x63, 0x17, 0x52, 0x8b, 0xc2, 0xac,
	0x32, 0xa0, 0x4f, 0xe0, 0xb2, 0x0b, 0xab, 0x8c, 0x0f, 0x20, 0x19, 0x73, 0x3b, 0x04, 0x3a, 0x49,
	0x88, 0x99, 0xc5, 0x3c, 0x8b, 0x2f, 0x2b, 0xf9, 0xa9, 0x08, 0xda, 0x5d, 0x72, 0xee, 0xb9, 0xdf,
	0xe5, 0x1c, 0xb8, 0xda, 0x0f, 0xce, 0xd4, 0xd6, 0x19, 0x8b, 0xce, 0x8f, 0xdc, 0x3a, 0xe3, 0x0d,
	0x2b, 0x82, 0xb8, 0x5a, 0xc4, 0x51, 0x8f, 0x5e, 0x24, 0x79, 0x75, 0xdd, 0x19, 0xd3, 0x1d, 0xb0,
	0x8e, 0xbf, 0x76, 0xd8, 0xd7, 0x42, 0xe7, 0x8d, 0xcd, 0x16, 0x96, 0x6f, 0xde, 0x29, 0xdd, 0xbd,
	0x58, 0xaf, 0x8c, 0x7e, 0xcd, 0x3c, 0x76, 0x0b, 0x97, 0x52, 0x7d, 0xda, 0x83, 0x18, 0xdf, 0xb5,
	0xe8, 0xb1, 0x24, 0x6b, 0x52, 0x4d, 0x9b, 0x59, 0xd6, 0x76, 0xa2, 0x47, 0x36, 0x07, 0xaa, 0x64,
	0x49, 0xe3, 0x80, 0x2a, 0xb9, 0xb9, 0x83, 0xf9, 0x6e, 0xe8, 0x5b, 0x74, 0x3f, 0x10, 0x06, 0x85,
	0x14, 0x5e, 0xc4, 0x65, 0xda, 0xc4, 0x77, 0x70, 0xa5, 0x83, 0xff, 0xba, 0xa6, 0xd9, 0x75, 0x0f,
	0x8b, 0xad, 0xf6, 0xd8, 0x9d, 0x80, 0x4d, 0xb2, 0xed, 0x8b, 0xc0, 0xc5, 0x6f, 0xc3, 0x87, 0x91,
	0x58, 0x4e, 0x12, 0x27, 0xbc, 0x59, 0x95, 0x97, 0xce, 0xd6, 0xa4, 0x9a, 0x3d, 0x2c, 0x79, 0x2a,
	0x82, 0x1f, 0x8b, 0xe0, 0x4f, 0x7a, 0x4c, 0xa8, 0x3f, 0x81, 0xe9, 0xa9, 0xc0, 0xe4, 0x18, 0x98,
	0xdd, 0x40, 0x11, 0x4a, 0x2e, 0x8b, 0x08, 0x07, 0x1e, 0x6a, 0xe7, 0xcf, 0xe8, 0x45, 0x13, 0xf5,
	0xf6, 0x3c, 0x9e, 0x79, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x96, 0x0c, 0xc9, 0x15, 0xab, 0x01,
	0x00, 0x00,
}
