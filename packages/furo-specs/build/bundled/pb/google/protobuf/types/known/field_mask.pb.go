// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/protobuf/types/known/field_mask.proto

package google_protobuf_types_known

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// A field mask in update operations specifies which fields of the targeted resource are going to be updated. https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/field_mask.proto
type FieldMask struct {
	// The implementation of any API method which has a FieldMask type field in the request should verify the included field paths, and return an `INVALID_ARGUMENT` error if any path is duplicated or unmappable.
	Paths                []string `protobuf:"bytes,1,rep,name=paths,proto3" json:"paths,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FieldMask) Reset()         { *m = FieldMask{} }
func (m *FieldMask) String() string { return proto.CompactTextString(m) }
func (*FieldMask) ProtoMessage()    {}
func (*FieldMask) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f6e6eab01706be6, []int{0}
}

func (m *FieldMask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldMask.Unmarshal(m, b)
}
func (m *FieldMask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldMask.Marshal(b, m, deterministic)
}
func (m *FieldMask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldMask.Merge(m, src)
}
func (m *FieldMask) XXX_Size() int {
	return xxx_messageInfo_FieldMask.Size(m)
}
func (m *FieldMask) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldMask.DiscardUnknown(m)
}

var xxx_messageInfo_FieldMask proto.InternalMessageInfo

func (m *FieldMask) GetPaths() []string {
	if m != nil {
		return m.Paths
	}
	return nil
}

func init() {
	proto.RegisterType((*FieldMask)(nil), "google.protobuf.types.known.FieldMask")
}

func init() {
	proto.RegisterFile("google/protobuf/types/known/field_mask.proto", fileDescriptor_0f6e6eab01706be6)
}

var fileDescriptor_0f6e6eab01706be6 = []byte{
	// 114 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x49, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x2f, 0xa9, 0x2c, 0x48,
	0x2d, 0xd6, 0xcf, 0xce, 0xcb, 0x2f, 0xcf, 0xd3, 0x4f, 0xcb, 0x4c, 0xcd, 0x49, 0x89, 0xcf, 0x4d,
	0x2c, 0xce, 0xd6, 0x03, 0xcb, 0x0b, 0x49, 0x43, 0x54, 0xeb, 0xc1, 0x54, 0xeb, 0x81, 0x55, 0xeb,
	0x81, 0x55, 0x2b, 0x29, 0x72, 0x71, 0xba, 0x81, 0x34, 0xf8, 0x26, 0x16, 0x67, 0x0b, 0x89, 0x70,
	0xb1, 0x16, 0x24, 0x96, 0x64, 0x14, 0x4b, 0x30, 0x2a, 0x30, 0x6b, 0x70, 0x06, 0x41, 0x38, 0x49,
	0x6c, 0x60, 0x8d, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x7e, 0xec, 0x91, 0x76, 0x00,
	0x00, 0x00,
}
