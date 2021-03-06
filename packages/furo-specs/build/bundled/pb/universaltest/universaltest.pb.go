// Code generated by protoc-gen-go. DO NOT EDIT.
// source: universaltest/universaltest.proto

package universaltest

import (
	furo "../furo"
	fat "../furo/fat"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// oneof experiment spec for testing
type Universaltest struct {
	// field of a fat bool for the Universaltest
	FatBool *fat.Bool `protobuf:"bytes,10,opt,name=fat_bool,json=fatBool,proto3" json:"fat_bool,omitempty"`
	// field of a fat int32 for the Universaltest
	FatInt32 *fat.Int32 `protobuf:"bytes,7,opt,name=fat_int32,json=fatInt32,proto3" json:"fat_int32,omitempty"`
	// field of a fat string for the Universaltest
	FatString *fat.String `protobuf:"bytes,4,opt,name=fat_string,json=fatString,proto3" json:"fat_string,omitempty"`
	// field of a fat string for the Universaltest
	FatStringList *fat.String `protobuf:"bytes,8,opt,name=fat_string_list,json=fatStringList,proto3" json:"fat_string_list,omitempty"`
	// field of repeated fat string for the Universaltest
	FatStringRepeated []*fat.String `protobuf:"bytes,11,rep,name=fat_string_repeated,json=fatStringRepeated,proto3" json:"fat_string_repeated,omitempty"`
	// Identity of a universaltes type
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// field of a scalar int32 for the Universaltest
	ScalarInt32 int32 `protobuf:"varint,5,opt,name=scalar_int32,json=scalarInt32,proto3" json:"scalar_int32,omitempty"`
	// field of a scalar string for the Universaltest
	ScalarString string `protobuf:"bytes,2,opt,name=scalar_string,json=scalarString,proto3" json:"scalar_string,omitempty"`
	// field of a wrapper boolean for the Universaltest
	WrapperBool *protobuf.BoolValue `protobuf:"bytes,9,opt,name=wrapper_bool,json=wrapperBool,proto3" json:"wrapper_bool,omitempty"`
	// field of a wrapper int32 for the Universaltest
	WrapperInt32 *protobuf.Int32Value `protobuf:"bytes,6,opt,name=wrapper_int32,json=wrapperInt32,proto3" json:"wrapper_int32,omitempty"`
	// field of a wrapper string for the Universaltest
	WrapperString        *protobuf.StringValue `protobuf:"bytes,3,opt,name=wrapper_string,json=wrapperString,proto3" json:"wrapper_string,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Universaltest) Reset()         { *m = Universaltest{} }
func (m *Universaltest) String() string { return proto.CompactTextString(m) }
func (*Universaltest) ProtoMessage()    {}
func (*Universaltest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdab0381820757b0, []int{0}
}

func (m *Universaltest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Universaltest.Unmarshal(m, b)
}
func (m *Universaltest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Universaltest.Marshal(b, m, deterministic)
}
func (m *Universaltest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Universaltest.Merge(m, src)
}
func (m *Universaltest) XXX_Size() int {
	return xxx_messageInfo_Universaltest.Size(m)
}
func (m *Universaltest) XXX_DiscardUnknown() {
	xxx_messageInfo_Universaltest.DiscardUnknown(m)
}

var xxx_messageInfo_Universaltest proto.InternalMessageInfo

func (m *Universaltest) GetFatBool() *fat.Bool {
	if m != nil {
		return m.FatBool
	}
	return nil
}

func (m *Universaltest) GetFatInt32() *fat.Int32 {
	if m != nil {
		return m.FatInt32
	}
	return nil
}

func (m *Universaltest) GetFatString() *fat.String {
	if m != nil {
		return m.FatString
	}
	return nil
}

func (m *Universaltest) GetFatStringList() *fat.String {
	if m != nil {
		return m.FatStringList
	}
	return nil
}

func (m *Universaltest) GetFatStringRepeated() []*fat.String {
	if m != nil {
		return m.FatStringRepeated
	}
	return nil
}

func (m *Universaltest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Universaltest) GetScalarInt32() int32 {
	if m != nil {
		return m.ScalarInt32
	}
	return 0
}

func (m *Universaltest) GetScalarString() string {
	if m != nil {
		return m.ScalarString
	}
	return ""
}

func (m *Universaltest) GetWrapperBool() *protobuf.BoolValue {
	if m != nil {
		return m.WrapperBool
	}
	return nil
}

func (m *Universaltest) GetWrapperInt32() *protobuf.Int32Value {
	if m != nil {
		return m.WrapperInt32
	}
	return nil
}

func (m *Universaltest) GetWrapperString() *protobuf.StringValue {
	if m != nil {
		return m.WrapperString
	}
	return nil
}

// UniversaltestEntity with Universaltest type in data
type UniversaltestEntity struct {
	// contains a universaltest.Universaltest
	Data *Universaltest `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// Hateoas links
	Links []*furo.Link `protobuf:"bytes,2,rep,name=links,proto3" json:"links,omitempty"`
	// Meta for the response
	Meta                 *furo.Meta `protobuf:"bytes,3,opt,name=meta,proto3" json:"meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UniversaltestEntity) Reset()         { *m = UniversaltestEntity{} }
func (m *UniversaltestEntity) String() string { return proto.CompactTextString(m) }
func (*UniversaltestEntity) ProtoMessage()    {}
func (*UniversaltestEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdab0381820757b0, []int{1}
}

func (m *UniversaltestEntity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UniversaltestEntity.Unmarshal(m, b)
}
func (m *UniversaltestEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UniversaltestEntity.Marshal(b, m, deterministic)
}
func (m *UniversaltestEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UniversaltestEntity.Merge(m, src)
}
func (m *UniversaltestEntity) XXX_Size() int {
	return xxx_messageInfo_UniversaltestEntity.Size(m)
}
func (m *UniversaltestEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_UniversaltestEntity.DiscardUnknown(m)
}

var xxx_messageInfo_UniversaltestEntity proto.InternalMessageInfo

func (m *UniversaltestEntity) GetData() *Universaltest {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *UniversaltestEntity) GetLinks() []*furo.Link {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *UniversaltestEntity) GetMeta() *furo.Meta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func init() {
	proto.RegisterType((*Universaltest)(nil), "universaltest.Universaltest")
	proto.RegisterType((*UniversaltestEntity)(nil), "universaltest.UniversaltestEntity")
}

func init() { proto.RegisterFile("universaltest/universaltest.proto", fileDescriptor_fdab0381820757b0) }

var fileDescriptor_fdab0381820757b0 = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x4d, 0x6f, 0xd4, 0x30,
	0x14, 0x54, 0xb6, 0xd9, 0xb6, 0xfb, 0xd2, 0xec, 0x82, 0x7b, 0x89, 0x16, 0x54, 0xa5, 0xe5, 0xb2,
	0x48, 0x28, 0x41, 0xed, 0x85, 0x0b, 0x52, 0x05, 0xe2, 0x80, 0x54, 0x2e, 0x46, 0x70, 0xad, 0xbc,
	0xc4, 0x59, 0x59, 0x35, 0x71, 0x64, 0xbf, 0x80, 0xf8, 0x09, 0xfc, 0x1e, 0xfe, 0x20, 0xf2, 0xb3,
	0xd3, 0x6d, 0x54, 0x38, 0x44, 0x72, 0xc6, 0x33, 0xe3, 0x79, 0x1f, 0x70, 0x3e, 0x74, 0xea, 0x87,
	0xb4, 0x4e, 0x68, 0x94, 0x0e, 0xeb, 0xc9, 0x5f, 0xd5, 0x5b, 0x83, 0x86, 0xe5, 0x13, 0x70, 0x7d,
	0xb6, 0x33, 0x66, 0xa7, 0x65, 0x4d, 0x97, 0xdb, 0xa1, 0xad, 0x7f, 0x5a, 0xd1, 0xf7, 0xd2, 0xba,
	0x40, 0x5f, 0xb3, 0x76, 0xb0, 0xa6, 0x6e, 0x05, 0xfa, 0x2f, 0x62, 0x2b, 0xc2, 0xbe, 0x4b, 0x14,
	0x13, 0x40, 0xab, 0xee, 0x2e, 0x00, 0x17, 0x7f, 0x52, 0xc8, 0xbf, 0x3c, 0x7c, 0x87, 0xbd, 0x84,
	0xe3, 0x56, 0xe0, 0xed, 0xd6, 0x18, 0x5d, 0x40, 0x99, 0x6c, 0xb2, 0xcb, 0x65, 0xe5, 0x55, 0x95,
	0xb7, 0x7d, 0x67, 0x8c, 0xe6, 0x47, 0xad, 0x40, 0x7f, 0x60, 0xaf, 0x60, 0xe1, 0xa9, 0xaa, 0xc3,
	0xab, 0xcb, 0xe2, 0x88, 0xb8, 0xab, 0x3d, 0xf7, 0xa3, 0x87, 0xb9, 0x37, 0xa3, 0x13, 0xab, 0x01,
	0x3c, 0xdb, 0xa1, 0x55, 0xdd, 0xae, 0x48, 0x89, 0xfe, 0x64, 0x4f, 0xff, 0x4c, 0x38, 0xf7, 0x8e,
	0xe1, 0xc8, 0xde, 0xc0, 0x6a, 0x2f, 0xb8, 0xd5, 0xca, 0x61, 0x71, 0xfc, 0x1f, 0x55, 0x7e, 0xaf,
	0xba, 0x51, 0x0e, 0xd9, 0x35, 0x9c, 0x3e, 0x50, 0x5a, 0xd9, 0x4b, 0x81, 0xb2, 0x29, 0xb2, 0xf2,
	0xe0, 0x9f, 0xea, 0xa7, 0xf7, 0x6a, 0x1e, 0xa9, 0x6c, 0x09, 0x33, 0xd5, 0x14, 0x49, 0x99, 0x6c,
	0x16, 0x7c, 0xa6, 0x1a, 0x76, 0x0e, 0x27, 0xee, 0x9b, 0xd0, 0xc2, 0xc6, 0x6a, 0xe7, 0x65, 0xb2,
	0x99, 0xf3, 0x2c, 0x60, 0xa1, 0xbe, 0x17, 0x90, 0x47, 0x4a, 0x2c, 0x71, 0x46, 0xea, 0xa8, 0x8b,
	0x35, 0xbd, 0x85, 0x93, 0x38, 0xb7, 0xd0, 0xe1, 0x05, 0x15, 0xb4, 0xae, 0xc2, 0x70, 0xab, 0x71,
	0xb8, 0xd4, 0xe8, 0xaf, 0x42, 0x0f, 0x92, 0x67, 0x91, 0x4f, 0x1d, 0xbf, 0x86, 0x7c, 0x94, 0x87,
	0x1c, 0x87, 0xa4, 0x7f, 0xf6, 0x48, 0x4f, 0x91, 0x82, 0xc1, 0xf8, 0x60, 0x48, 0xf9, 0x1e, 0x96,
	0xa3, 0x43, 0x8c, 0x79, 0x40, 0x16, 0xcf, 0x1f, 0x59, 0x84, 0xc4, 0xc1, 0x63, 0x7c, 0x35, 0x60,
	0x17, 0xbf, 0x13, 0x38, 0x9d, 0x6c, 0xcd, 0x87, 0x0e, 0x15, 0xfe, 0x62, 0xaf, 0x21, 0x6d, 0x04,
	0x0a, 0xea, 0x9b, 0xb7, 0x9c, 0xae, 0xf5, 0x44, 0xc1, 0x89, 0xc9, 0x4a, 0x98, 0xfb, 0x6d, 0x74,
	0xc5, 0x8c, 0x66, 0x03, 0x61, 0x36, 0x37, 0xaa, 0xbb, 0xe3, 0xe1, 0x82, 0x9d, 0x41, 0xea, 0x17,
	0x38, 0xc6, 0x8c, 0x84, 0x4f, 0x12, 0x05, 0x27, 0x7c, 0x7b, 0x48, 0x81, 0xaf, 0xfe, 0x06, 0x00,
	0x00, 0xff, 0xff, 0x35, 0x5d, 0x85, 0x44, 0x52, 0x03, 0x00, 0x00,
}
