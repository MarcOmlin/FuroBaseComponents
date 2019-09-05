// Code generated by protoc-gen-go. DO NOT EDIT.
// source: furo/property.proto

package furo

import (
	protobuf "../google/protobuf"
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

// Type to define property values with type information
type Property struct {
	// data part of the property
	Data *protobuf.Any `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	// String representation of the property
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Id of the property
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// meta part of the property
	Meta                 *Meta    `protobuf:"bytes,4,opt,name=meta,proto3" json:"meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Property) Reset()         { *m = Property{} }
func (m *Property) String() string { return proto.CompactTextString(m) }
func (*Property) ProtoMessage()    {}
func (*Property) Descriptor() ([]byte, []int) {
	return fileDescriptor_bca1bac703f4631e, []int{0}
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

func (m *Property) GetData() *protobuf.Any {
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
	proto.RegisterType((*Property)(nil), "furo.Property")
}

func init() { proto.RegisterFile("furo/property.proto", fileDescriptor_bca1bac703f4631e) }

var fileDescriptor_bca1bac703f4631e = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0x2b, 0x2d, 0xca,
	0xd7, 0x2f, 0x28, 0xca, 0x2f, 0x48, 0x2d, 0x2a, 0xa9, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x01, 0x09, 0x4a, 0xf1, 0x83, 0xa5, 0x72, 0x53, 0x4b, 0x12, 0x21, 0xc2, 0x52, 0x92, 0xe9,
	0xf9, 0xf9, 0xe9, 0x39, 0xa9, 0xfa, 0x60, 0x5e, 0x52, 0x69, 0x9a, 0x7e, 0x62, 0x1e, 0x54, 0x87,
	0x52, 0x3b, 0x23, 0x17, 0x47, 0x00, 0xd4, 0x10, 0x21, 0x0d, 0x2e, 0x96, 0x94, 0xc4, 0x92, 0x44,
	0x09, 0x66, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x11, 0x3d, 0x88, 0x36, 0x3d, 0x98, 0x36, 0x3d, 0xc7,
	0xbc, 0xca, 0x20, 0xb0, 0x0a, 0x21, 0x45, 0x2e, 0x9e, 0x94, 0xcc, 0xe2, 0x82, 0x9c, 0xc4, 0xca,
	0xf8, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x6e, 0xa8, 0x98, 0x5f,
	0x62, 0x6e, 0xaa, 0x10, 0x1f, 0x17, 0x53, 0x66, 0x8a, 0x04, 0x23, 0x58, 0x82, 0x29, 0x33, 0x45,
	0x48, 0x8e, 0x8b, 0x05, 0xe4, 0x24, 0x09, 0x16, 0xb0, 0xe1, 0x5c, 0x7a, 0x20, 0x47, 0xea, 0xf9,
	0xa6, 0x96, 0x24, 0x06, 0x81, 0xc5, 0x93, 0xd8, 0xc0, 0xd6, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xc5, 0x5a, 0x0f, 0xfc, 0xd9, 0x00, 0x00, 0x00,
}
