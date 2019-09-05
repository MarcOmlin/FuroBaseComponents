// Code generated by protoc-gen-go. DO NOT EDIT.
// source: treeservice/service.proto

package treeservice

import (
	protobuf "../google/protobuf"
	tree "../tree"
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

type CreateTreeServiceRequest struct {
	Data                 *tree.Tree `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CreateTreeServiceRequest) Reset()         { *m = CreateTreeServiceRequest{} }
func (m *CreateTreeServiceRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTreeServiceRequest) ProtoMessage()    {}
func (*CreateTreeServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9dd10294fe0b6ab0, []int{0}
}

func (m *CreateTreeServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTreeServiceRequest.Unmarshal(m, b)
}
func (m *CreateTreeServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTreeServiceRequest.Marshal(b, m, deterministic)
}
func (m *CreateTreeServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTreeServiceRequest.Merge(m, src)
}
func (m *CreateTreeServiceRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTreeServiceRequest.Size(m)
}
func (m *CreateTreeServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTreeServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTreeServiceRequest proto.InternalMessageInfo

func (m *CreateTreeServiceRequest) GetData() *tree.Tree {
	if m != nil {
		return m.Data
	}
	return nil
}

type DeleteTreeServiceRequest struct {
	Tre                  string          `protobuf:"bytes,1,opt,name=tre,proto3" json:"tre,omitempty"`
	Data                 *protobuf.Empty `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *DeleteTreeServiceRequest) Reset()         { *m = DeleteTreeServiceRequest{} }
func (m *DeleteTreeServiceRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteTreeServiceRequest) ProtoMessage()    {}
func (*DeleteTreeServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9dd10294fe0b6ab0, []int{1}
}

func (m *DeleteTreeServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTreeServiceRequest.Unmarshal(m, b)
}
func (m *DeleteTreeServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTreeServiceRequest.Marshal(b, m, deterministic)
}
func (m *DeleteTreeServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTreeServiceRequest.Merge(m, src)
}
func (m *DeleteTreeServiceRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteTreeServiceRequest.Size(m)
}
func (m *DeleteTreeServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTreeServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTreeServiceRequest proto.InternalMessageInfo

func (m *DeleteTreeServiceRequest) GetTre() string {
	if m != nil {
		return m.Tre
	}
	return ""
}

func (m *DeleteTreeServiceRequest) GetData() *protobuf.Empty {
	if m != nil {
		return m.Data
	}
	return nil
}

type GetTreeServiceRequest struct {
	Tre                  string   `protobuf:"bytes,1,opt,name=tre,proto3" json:"tre,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTreeServiceRequest) Reset()         { *m = GetTreeServiceRequest{} }
func (m *GetTreeServiceRequest) String() string { return proto.CompactTextString(m) }
func (*GetTreeServiceRequest) ProtoMessage()    {}
func (*GetTreeServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9dd10294fe0b6ab0, []int{2}
}

func (m *GetTreeServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTreeServiceRequest.Unmarshal(m, b)
}
func (m *GetTreeServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTreeServiceRequest.Marshal(b, m, deterministic)
}
func (m *GetTreeServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTreeServiceRequest.Merge(m, src)
}
func (m *GetTreeServiceRequest) XXX_Size() int {
	return xxx_messageInfo_GetTreeServiceRequest.Size(m)
}
func (m *GetTreeServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTreeServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTreeServiceRequest proto.InternalMessageInfo

func (m *GetTreeServiceRequest) GetTre() string {
	if m != nil {
		return m.Tre
	}
	return ""
}

type ListTreeServiceRequest struct {
	//Partial representation, fields=id,name
	Fields string `protobuf:"bytes,1,opt,name=fields,proto3" json:"fields,omitempty"`
	//*
	// Sort fields, comma separated list for the ordering
	// use **?filter=-display_name** with a dash to sort descending
	// use **?filter=display_name** to sort ascending
	OrderBy string `protobuf:"bytes,2,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	//Filter
	Filter string `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
	//Page number for paginated content. Tipp: follow the HATEOAS next, prev,...
	Page int32 `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	//Number of elements to return per page
	Limit int32 `protobuf:"varint,5,opt,name=limit,proto3" json:"limit,omitempty"`
	//https://cloud.google.com/apis/design/design_patterns#resource_view
	View string `protobuf:"bytes,8,opt,name=view,proto3" json:"view,omitempty"`
	//Query term to search a tree
	Q                    string   `protobuf:"bytes,11,opt,name=q,proto3" json:"q,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTreeServiceRequest) Reset()         { *m = ListTreeServiceRequest{} }
func (m *ListTreeServiceRequest) String() string { return proto.CompactTextString(m) }
func (*ListTreeServiceRequest) ProtoMessage()    {}
func (*ListTreeServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9dd10294fe0b6ab0, []int{3}
}

func (m *ListTreeServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTreeServiceRequest.Unmarshal(m, b)
}
func (m *ListTreeServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTreeServiceRequest.Marshal(b, m, deterministic)
}
func (m *ListTreeServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTreeServiceRequest.Merge(m, src)
}
func (m *ListTreeServiceRequest) XXX_Size() int {
	return xxx_messageInfo_ListTreeServiceRequest.Size(m)
}
func (m *ListTreeServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTreeServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListTreeServiceRequest proto.InternalMessageInfo

func (m *ListTreeServiceRequest) GetFields() string {
	if m != nil {
		return m.Fields
	}
	return ""
}

func (m *ListTreeServiceRequest) GetOrderBy() string {
	if m != nil {
		return m.OrderBy
	}
	return ""
}

func (m *ListTreeServiceRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *ListTreeServiceRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListTreeServiceRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListTreeServiceRequest) GetView() string {
	if m != nil {
		return m.View
	}
	return ""
}

func (m *ListTreeServiceRequest) GetQ() string {
	if m != nil {
		return m.Q
	}
	return ""
}

type UpdateTreeServiceRequest struct {
	Tre                  string     `protobuf:"bytes,1,opt,name=tre,proto3" json:"tre,omitempty"`
	Data                 *tree.Tree `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UpdateTreeServiceRequest) Reset()         { *m = UpdateTreeServiceRequest{} }
func (m *UpdateTreeServiceRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateTreeServiceRequest) ProtoMessage()    {}
func (*UpdateTreeServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9dd10294fe0b6ab0, []int{4}
}

func (m *UpdateTreeServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTreeServiceRequest.Unmarshal(m, b)
}
func (m *UpdateTreeServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTreeServiceRequest.Marshal(b, m, deterministic)
}
func (m *UpdateTreeServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTreeServiceRequest.Merge(m, src)
}
func (m *UpdateTreeServiceRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateTreeServiceRequest.Size(m)
}
func (m *UpdateTreeServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTreeServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTreeServiceRequest proto.InternalMessageInfo

func (m *UpdateTreeServiceRequest) GetTre() string {
	if m != nil {
		return m.Tre
	}
	return ""
}

func (m *UpdateTreeServiceRequest) GetData() *tree.Tree {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateTreeServiceRequest)(nil), "treeservice.CreateTreeServiceRequest")
	proto.RegisterType((*DeleteTreeServiceRequest)(nil), "treeservice.DeleteTreeServiceRequest")
	proto.RegisterType((*GetTreeServiceRequest)(nil), "treeservice.GetTreeServiceRequest")
	proto.RegisterType((*ListTreeServiceRequest)(nil), "treeservice.ListTreeServiceRequest")
	proto.RegisterType((*UpdateTreeServiceRequest)(nil), "treeservice.UpdateTreeServiceRequest")
}

func init() { proto.RegisterFile("treeservice/service.proto", fileDescriptor_9dd10294fe0b6ab0) }

var fileDescriptor_9dd10294fe0b6ab0 = []byte{
	// 481 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0xdb, 0xa4, 0x6d, 0x26, 0x48, 0x2d, 0xa3, 0x36, 0xdd, 0x1a, 0x1a, 0x45, 0xae, 0x40,
	0xa5, 0x07, 0x5b, 0x0a, 0xb7, 0x1e, 0x29, 0x15, 0x97, 0x9e, 0x0c, 0x48, 0xdc, 0x2a, 0x27, 0x9e,
	0x5a, 0x5b, 0x1c, 0xaf, 0xb3, 0xde, 0x16, 0x45, 0x88, 0x0b, 0xbf, 0x80, 0xf8, 0x0e, 0x3e, 0x86,
	0x5f, 0xe0, 0x43, 0xd0, 0x8e, 0x1d, 0x39, 0x49, 0x1d, 0x29, 0x97, 0x64, 0x67, 0x76, 0xe6, 0xbd,
	0xd9, 0x79, 0xcf, 0x70, 0x62, 0x34, 0x51, 0x41, 0xfa, 0x51, 0x8e, 0x29, 0xa8, 0xfe, 0xfd, 0x5c,
	0x2b, 0xa3, 0xb0, 0xbb, 0x70, 0xe5, 0xbe, 0x4c, 0x94, 0x4a, 0x52, 0x0a, 0xa2, 0x5c, 0x06, 0x51,
	0x96, 0x29, 0x13, 0x19, 0xa9, 0xb2, 0xa2, 0x2c, 0x75, 0xf7, 0x6d, 0x69, 0x60, 0x7f, 0xaa, 0xc4,
	0x8b, 0xaa, 0x9c, 0xa3, 0xd1, 0xc3, 0x5d, 0x40, 0x93, 0xdc, 0xcc, 0xca, 0x4b, 0xef, 0x12, 0xc4,
	0x95, 0xa6, 0xc8, 0xd0, 0x27, 0x4d, 0xf4, 0xb1, 0x24, 0x08, 0x69, 0xfa, 0x40, 0x85, 0xc1, 0x3e,
	0xb4, 0xe2, 0xc8, 0x44, 0xc2, 0x19, 0x38, 0xe7, 0xdd, 0x21, 0xf8, 0x8c, 0x69, 0xeb, 0x42, 0xce,
	0x7b, 0x5f, 0x40, 0xbc, 0xa7, 0x94, 0x1a, 0x7b, 0x0f, 0x60, 0xdb, 0x68, 0xe2, 0xd6, 0x4e, 0x68,
	0x8f, 0x78, 0x51, 0xa1, 0x6d, 0x31, 0x5a, 0xcf, 0x2f, 0xa7, 0xf2, 0xe7, 0x53, 0xf9, 0xd7, 0x76,
	0xaa, 0x0a, 0xf9, 0x0d, 0x1c, 0x7d, 0x20, 0xb3, 0x09, 0xac, 0xf7, 0xc7, 0x81, 0xde, 0x8d, 0x2c,
	0x9a, 0x8a, 0x7b, 0xb0, 0x73, 0x27, 0x29, 0x8d, 0x8b, 0xaa, 0xbe, 0x8a, 0xf0, 0x04, 0xf6, 0x94,
	0x8e, 0x49, 0xdf, 0x8e, 0x66, 0x3c, 0x4d, 0x27, 0xdc, 0xe5, 0xf8, 0xdd, 0xac, 0x6c, 0x49, 0x0d,
	0x69, 0xb1, 0x3d, 0x6f, 0xb1, 0x11, 0x22, 0xb4, 0xf2, 0x28, 0x21, 0xd1, 0x1a, 0x38, 0xe7, 0xed,
	0x90, 0xcf, 0x78, 0x08, 0xed, 0x54, 0x4e, 0xa4, 0x11, 0x6d, 0x4e, 0x96, 0x81, 0xad, 0x7c, 0x94,
	0xf4, 0x4d, 0xec, 0x71, 0x3f, 0x9f, 0xf1, 0x19, 0x38, 0x53, 0xd1, 0xe5, 0x84, 0x33, 0xf5, 0x6e,
	0x40, 0x7c, 0xce, 0xe3, 0x68, 0xc3, 0xb5, 0xf5, 0x97, 0xd6, 0xf6, 0x44, 0x84, 0xe1, 0xef, 0x16,
	0x74, 0x17, 0x80, 0x30, 0x06, 0xa8, 0x05, 0xc5, 0x57, 0xfe, 0x82, 0x71, 0xfc, 0x75, 0x4a, 0xbb,
	0x07, 0x35, 0xec, 0x75, 0x66, 0xa4, 0x99, 0x79, 0xa7, 0x3f, 0xff, 0xfe, 0xfb, 0xb5, 0x75, 0xec,
	0xed, 0x07, 0x13, 0x35, 0xfe, 0x6a, 0x99, 0xd8, 0x52, 0xc5, 0x25, 0xb3, 0xe2, 0x3d, 0x40, 0x2d,
	0xfd, 0x0a, 0xcb, 0x3a, 0x4f, 0xb8, 0x6b, 0x34, 0x9f, 0x73, 0x5d, 0x1c, 0xad, 0x70, 0x05, 0xdf,
	0x8d, 0xa6, 0x1f, 0x98, 0xc0, 0x6e, 0x65, 0x06, 0xf4, 0x96, 0x88, 0x1a, 0x2d, 0xd2, 0xf0, 0x96,
	0xd7, 0x8c, 0x3f, 0xc0, 0x7e, 0x23, 0x7e, 0x90, 0x90, 0xf1, 0xef, 0x0b, 0x95, 0xe1, 0x2d, 0x74,
	0xe6, 0x4e, 0x2a, 0xf0, 0x6c, 0x89, 0xaa, 0xd9, 0x61, 0xee, 0x61, 0xcd, 0x75, 0xa5, 0xd2, 0x94,
	0xc6, 0xf6, 0x43, 0xf4, 0x8e, 0x99, 0xef, 0x39, 0xae, 0xee, 0xce, 0x6e, 0xad, 0x56, 0x7e, 0x65,
	0x6b, 0xeb, 0x2c, 0xd1, 0xf0, 0x9e, 0x33, 0xc6, 0x3f, 0x1d, 0x36, 0xef, 0xab, 0x54, 0x68, 0xb4,
	0xc3, 0x4b, 0x7e, 0xfb, 0x3f, 0x00, 0x00, 0xff, 0xff, 0xc2, 0x10, 0xa0, 0x4b, 0x55, 0x04, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TreeServiceClient is the client API for TreeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TreeServiceClient interface {
	// Creates a new Tree
	CreateTree(ctx context.Context, in *CreateTreeServiceRequest, opts ...grpc.CallOption) (*tree.TreeEntity, error)
	// Delete a Tree
	DeleteTree(ctx context.Context, in *DeleteTreeServiceRequest, opts ...grpc.CallOption) (*protobuf.Empty, error)
	// The Get method takes zero or more parameters, and returns a TreeEntity which contains a Tree
	GetTree(ctx context.Context, in *GetTreeServiceRequest, opts ...grpc.CallOption) (*tree.TreeEntity, error)
	// The List method takes zero or more parameters as input, and returns a TreeCollection of TreeEntity that match the input parameters.
	ListTrees(ctx context.Context, in *ListTreeServiceRequest, opts ...grpc.CallOption) (*tree.TreeCollection, error)
	// Updates a Tree, partial updates are supported
	UpdateTree(ctx context.Context, in *UpdateTreeServiceRequest, opts ...grpc.CallOption) (*tree.TreeEntity, error)
}

type treeServiceClient struct {
	cc *grpc.ClientConn
}

func NewTreeServiceClient(cc *grpc.ClientConn) TreeServiceClient {
	return &treeServiceClient{cc}
}

func (c *treeServiceClient) CreateTree(ctx context.Context, in *CreateTreeServiceRequest, opts ...grpc.CallOption) (*tree.TreeEntity, error) {
	out := new(tree.TreeEntity)
	err := c.cc.Invoke(ctx, "/treeservice.TreeService/CreateTree", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *treeServiceClient) DeleteTree(ctx context.Context, in *DeleteTreeServiceRequest, opts ...grpc.CallOption) (*protobuf.Empty, error) {
	out := new(protobuf.Empty)
	err := c.cc.Invoke(ctx, "/treeservice.TreeService/DeleteTree", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *treeServiceClient) GetTree(ctx context.Context, in *GetTreeServiceRequest, opts ...grpc.CallOption) (*tree.TreeEntity, error) {
	out := new(tree.TreeEntity)
	err := c.cc.Invoke(ctx, "/treeservice.TreeService/GetTree", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *treeServiceClient) ListTrees(ctx context.Context, in *ListTreeServiceRequest, opts ...grpc.CallOption) (*tree.TreeCollection, error) {
	out := new(tree.TreeCollection)
	err := c.cc.Invoke(ctx, "/treeservice.TreeService/ListTrees", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *treeServiceClient) UpdateTree(ctx context.Context, in *UpdateTreeServiceRequest, opts ...grpc.CallOption) (*tree.TreeEntity, error) {
	out := new(tree.TreeEntity)
	err := c.cc.Invoke(ctx, "/treeservice.TreeService/UpdateTree", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TreeServiceServer is the server API for TreeService service.
type TreeServiceServer interface {
	// Creates a new Tree
	CreateTree(context.Context, *CreateTreeServiceRequest) (*tree.TreeEntity, error)
	// Delete a Tree
	DeleteTree(context.Context, *DeleteTreeServiceRequest) (*protobuf.Empty, error)
	// The Get method takes zero or more parameters, and returns a TreeEntity which contains a Tree
	GetTree(context.Context, *GetTreeServiceRequest) (*tree.TreeEntity, error)
	// The List method takes zero or more parameters as input, and returns a TreeCollection of TreeEntity that match the input parameters.
	ListTrees(context.Context, *ListTreeServiceRequest) (*tree.TreeCollection, error)
	// Updates a Tree, partial updates are supported
	UpdateTree(context.Context, *UpdateTreeServiceRequest) (*tree.TreeEntity, error)
}

func RegisterTreeServiceServer(s *grpc.Server, srv TreeServiceServer) {
	s.RegisterService(&_TreeService_serviceDesc, srv)
}

func _TreeService_CreateTree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTreeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TreeServiceServer).CreateTree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/treeservice.TreeService/CreateTree",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TreeServiceServer).CreateTree(ctx, req.(*CreateTreeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TreeService_DeleteTree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTreeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TreeServiceServer).DeleteTree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/treeservice.TreeService/DeleteTree",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TreeServiceServer).DeleteTree(ctx, req.(*DeleteTreeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TreeService_GetTree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTreeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TreeServiceServer).GetTree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/treeservice.TreeService/GetTree",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TreeServiceServer).GetTree(ctx, req.(*GetTreeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TreeService_ListTrees_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTreeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TreeServiceServer).ListTrees(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/treeservice.TreeService/ListTrees",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TreeServiceServer).ListTrees(ctx, req.(*ListTreeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TreeService_UpdateTree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTreeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TreeServiceServer).UpdateTree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/treeservice.TreeService/UpdateTree",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TreeServiceServer).UpdateTree(ctx, req.(*UpdateTreeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TreeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "treeservice.TreeService",
	HandlerType: (*TreeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTree",
			Handler:    _TreeService_CreateTree_Handler,
		},
		{
			MethodName: "DeleteTree",
			Handler:    _TreeService_DeleteTree_Handler,
		},
		{
			MethodName: "GetTree",
			Handler:    _TreeService_GetTree_Handler,
		},
		{
			MethodName: "ListTrees",
			Handler:    _TreeService_ListTrees_Handler,
		},
		{
			MethodName: "UpdateTree",
			Handler:    _TreeService_UpdateTree_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "treeservice/service.proto",
}
