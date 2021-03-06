// Code generated by protoc-gen-go. DO NOT EDIT.
// source: menu/menu.proto

package menu

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

// Item of a contextual menu
type Menuitem struct {
	// String representation of the menu item action
	Action string `protobuf:"bytes,5,opt,name=action,proto3" json:"action,omitempty"`
	// Children of this item
	Children []*Menuitem `protobuf:"bytes,7,rep,name=children,proto3" json:"children,omitempty"`
	// Keyboard command hint
	Command string `protobuf:"bytes,4,opt,name=command,proto3" json:"command,omitempty"`
	// Display actions as disabled when they can only be used sometimes, under certain conditions. They should be displayed as disabled rather than removing them.
	Disabled bool `protobuf:"varint,3,opt,name=disabled,proto3" json:"disabled,omitempty"`
	// String representation of the menu item. Menu item text
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Attribute flags e.g. important, negative, positive
	Flags []string `protobuf:"bytes,8,rep,name=flags,proto3" json:"flags,omitempty"`
	// Leading icon of the menu
	Icon string `protobuf:"bytes,1,opt,name=icon,proto3" json:"icon,omitempty"`
	// Item has a leading divider line
	LeadingDivider bool `protobuf:"varint,6,opt,name=leading_divider,json=leadingDivider,proto3" json:"leading_divider,omitempty"`
	// Optional payload
	Payload              []*any.Any `protobuf:"bytes,9,rep,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Menuitem) Reset()         { *m = Menuitem{} }
func (m *Menuitem) String() string { return proto.CompactTextString(m) }
func (*Menuitem) ProtoMessage()    {}
func (*Menuitem) Descriptor() ([]byte, []int) {
	return fileDescriptor_2659839f6edb8cc2, []int{0}
}

func (m *Menuitem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Menuitem.Unmarshal(m, b)
}
func (m *Menuitem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Menuitem.Marshal(b, m, deterministic)
}
func (m *Menuitem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Menuitem.Merge(m, src)
}
func (m *Menuitem) XXX_Size() int {
	return xxx_messageInfo_Menuitem.Size(m)
}
func (m *Menuitem) XXX_DiscardUnknown() {
	xxx_messageInfo_Menuitem.DiscardUnknown(m)
}

var xxx_messageInfo_Menuitem proto.InternalMessageInfo

func (m *Menuitem) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *Menuitem) GetChildren() []*Menuitem {
	if m != nil {
		return m.Children
	}
	return nil
}

func (m *Menuitem) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *Menuitem) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func (m *Menuitem) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Menuitem) GetFlags() []string {
	if m != nil {
		return m.Flags
	}
	return nil
}

func (m *Menuitem) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *Menuitem) GetLeadingDivider() bool {
	if m != nil {
		return m.LeadingDivider
	}
	return false
}

func (m *Menuitem) GetPayload() []*any.Any {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*Menuitem)(nil), "menu.Menuitem")
}

func init() { proto.RegisterFile("menu/menu.proto", fileDescriptor_2659839f6edb8cc2) }

var fileDescriptor_2659839f6edb8cc2 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x90, 0xcb, 0x6a, 0xf3, 0x30,
	0x10, 0x85, 0x49, 0xe2, 0xf8, 0x32, 0xf9, 0xff, 0x04, 0x44, 0x28, 0x6a, 0x56, 0x6e, 0x37, 0x35,
	0x5d, 0xc8, 0xd0, 0x3e, 0x41, 0x42, 0xb7, 0xed, 0xc2, 0x2f, 0x10, 0xc6, 0x92, 0xe2, 0x0a, 0x74,
	0x31, 0xbe, 0x14, 0xfc, 0x4a, 0x7d, 0xca, 0x62, 0xd9, 0xce, 0x46, 0xcc, 0x39, 0x9a, 0x6f, 0x86,
	0x39, 0x70, 0x30, 0xd2, 0xf6, 0xf9, 0xf8, 0xb0, 0xba, 0x71, 0x9d, 0x23, 0xc1, 0x58, 0x9f, 0x1e,
	0x2b, 0xe7, 0x2a, 0x2d, 0x73, 0xef, 0x95, 0xfd, 0x2d, 0x47, 0x3b, 0x4c, 0x0d, 0xcf, 0xbf, 0x6b,
	0x88, 0x3f, 0xa5, 0xed, 0x55, 0x27, 0x0d, 0x79, 0x80, 0x10, 0x79, 0xa7, 0x9c, 0xa5, 0xdb, 0x74,
	0x95, 0x25, 0xc5, 0xac, 0xc8, 0x2b, 0xc4, 0xfc, 0x5b, 0x69, 0xd1, 0x48, 0x4b, 0xa3, 0x74, 0x93,
	0xed, 0xde, 0xf6, 0xcc, 0x2f, 0x59, 0xc8, 0xe2, 0xfe, 0x4f, 0x28, 0x44, 0xdc, 0x19, 0x83, 0x56,
	0xd0, 0xc0, 0x0f, 0x59, 0x24, 0x39, 0x41, 0x2c, 0x54, 0x8b, 0xa5, 0x96, 0x82, 0x6e, 0xd2, 0x55,
	0x16, 0x17, 0x77, 0x4d, 0x9e, 0xe0, 0x9f, 0x50, 0x6d, 0xad, 0x71, 0xb8, 0x5a, 0x34, 0x92, 0xae,
	0x3d, 0xba, 0x9b, 0xbd, 0x2f, 0x34, 0x92, 0x1c, 0x61, 0x7b, 0xd3, 0x58, 0xb5, 0x34, 0x4e, 0x37,
	0x59, 0x52, 0x4c, 0x82, 0x10, 0x08, 0x14, 0x77, 0x96, 0xae, 0x3c, 0xe0, 0x6b, 0xf2, 0x02, 0x07,
	0x2d, 0x51, 0x28, 0x5b, 0x5d, 0x85, 0xfa, 0x51, 0x42, 0x36, 0x34, 0xf4, 0xfb, 0xf6, 0xb3, 0xfd,
	0x31, 0xb9, 0x84, 0x41, 0x54, 0xe3, 0xa0, 0x1d, 0x0a, 0x9a, 0xf8, 0xb3, 0x8e, 0x6c, 0x4a, 0x8a,
	0x2d, 0x49, 0xb1, 0xb3, 0x1d, 0x8a, 0xa5, 0xe9, 0x42, 0xe1, 0x3f, 0x77, 0x86, 0x21, 0x37, 0xd2,
	0xdf, 0x7f, 0x89, 0xc6, 0x00, 0xce, 0xb5, 0x2a, 0x43, 0x0f, 0xbc, 0xff, 0x05, 0x00, 0x00, 0xff,
	0xff, 0x65, 0xb0, 0x70, 0x9a, 0x81, 0x01, 0x00, 0x00,
}
