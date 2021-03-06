// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/managed_placement_view_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

// Request message for
// [ManagedPlacementViewService.GetManagedPlacementView][google.ads.googleads.v1.services.ManagedPlacementViewService.GetManagedPlacementView].
type GetManagedPlacementViewRequest struct {
	// The resource name of the Managed Placement View to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetManagedPlacementViewRequest) Reset()         { *m = GetManagedPlacementViewRequest{} }
func (m *GetManagedPlacementViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetManagedPlacementViewRequest) ProtoMessage()    {}
func (*GetManagedPlacementViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c301455ca7db209f, []int{0}
}

func (m *GetManagedPlacementViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetManagedPlacementViewRequest.Unmarshal(m, b)
}
func (m *GetManagedPlacementViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetManagedPlacementViewRequest.Marshal(b, m, deterministic)
}
func (m *GetManagedPlacementViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetManagedPlacementViewRequest.Merge(m, src)
}
func (m *GetManagedPlacementViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetManagedPlacementViewRequest.Size(m)
}
func (m *GetManagedPlacementViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetManagedPlacementViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetManagedPlacementViewRequest proto.InternalMessageInfo

func (m *GetManagedPlacementViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetManagedPlacementViewRequest)(nil), "google.ads.googleads.v1.services.GetManagedPlacementViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/managed_placement_view_service.proto", fileDescriptor_c301455ca7db209f)
}

var fileDescriptor_c301455ca7db209f = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x3d, 0x6b, 0xdb, 0x40,
	0x18, 0x46, 0x2a, 0x14, 0x2a, 0xda, 0x45, 0x4b, 0x8b, 0x5b, 0x8a, 0x70, 0x3d, 0x14, 0x0f, 0x77,
	0xa8, 0x1d, 0x4c, 0x2e, 0xe4, 0x43, 0x06, 0xe3, 0x2c, 0x09, 0xc6, 0x01, 0x0d, 0x41, 0x20, 0x2e,
	0xd2, 0x8b, 0x10, 0x58, 0x77, 0x8a, 0x5e, 0x59, 0x1e, 0x42, 0x96, 0x2c, 0xf9, 0x01, 0xf9, 0x07,
	0x19, 0xf3, 0x53, 0xb2, 0x86, 0xfc, 0x83, 0x4c, 0xd9, 0xb3, 0x07, 0xf9, 0x7c, 0x02, 0x83, 0x65,
	0x6f, 0x0f, 0x77, 0xcf, 0xc7, 0xbd, 0xcf, 0x7b, 0xd6, 0x28, 0x91, 0x32, 0x99, 0x01, 0xe5, 0x31,
	0x52, 0x05, 0x6b, 0x54, 0xb9, 0x14, 0xa1, 0xa8, 0xd2, 0x08, 0x90, 0x66, 0x5c, 0xf0, 0x04, 0xe2,
	0x30, 0x9f, 0xf1, 0x08, 0x32, 0x10, 0x65, 0x58, 0xa5, 0xb0, 0x08, 0x57, 0xf7, 0x24, 0x2f, 0x64,
	0x29, 0x6d, 0x47, 0x69, 0x09, 0x8f, 0x91, 0x34, 0x36, 0xa4, 0x72, 0x89, 0xb6, 0xe9, 0x1c, 0xb6,
	0x05, 0x15, 0x80, 0x72, 0x5e, 0xb4, 0x27, 0xa9, 0x84, 0xce, 0x2f, 0xad, 0xcf, 0x53, 0xca, 0x85,
	0x90, 0x25, 0x2f, 0x53, 0x29, 0x50, 0xdd, 0x76, 0x47, 0xd6, 0xef, 0x31, 0x94, 0xa7, 0xca, 0x60,
	0xa2, 0xf5, 0x7e, 0x0a, 0x8b, 0x29, 0x5c, 0xcd, 0x01, 0x4b, 0xfb, 0x8f, 0xf5, 0x4d, 0x27, 0x85,
	0x82, 0x67, 0xf0, 0xc3, 0x70, 0x8c, 0xbf, 0x5f, 0xa6, 0x5f, 0xf5, 0xe1, 0x19, 0xcf, 0xe0, 0xdf,
	0xbb, 0x61, 0xfd, 0xdc, 0x64, 0x72, 0xae, 0xa6, 0xb0, 0x5f, 0x0c, 0xeb, 0x7b, 0x4b, 0x8e, 0x7d,
	0x4c, 0x76, 0x75, 0x40, 0xb6, 0x3f, 0xb1, 0x33, 0x68, 0x75, 0x68, 0x3a, 0x22, 0x9b, 0xf4, 0xdd,
	0xa3, 0xdb, 0xe7, 0xd7, 0x7b, 0x73, 0xcf, 0x1e, 0xd4, 0x7d, 0x5e, 0xaf, 0x8d, 0x79, 0x10, 0xcd,
	0xb1, 0x94, 0x19, 0x14, 0x48, 0xfb, 0xba, 0xe0, 0x35, 0x31, 0xd2, 0xfe, 0xcd, 0xf0, 0xce, 0xb4,
	0x7a, 0x91, 0xcc, 0x76, 0x4e, 0x30, 0x74, 0xb6, 0xb4, 0x33, 0xa9, 0x37, 0x31, 0x31, 0x2e, 0x4e,
	0x56, 0x2e, 0x89, 0x9c, 0x71, 0x91, 0x10, 0x59, 0x24, 0x34, 0x01, 0xb1, 0xdc, 0x93, 0xde, 0x7c,
	0x9e, 0x62, 0xfb, 0x8f, 0xdb, 0xd7, 0xe0, 0xc1, 0xfc, 0x34, 0xf6, 0xbc, 0x47, 0xd3, 0x19, 0x2b,
	0x43, 0x2f, 0x46, 0xa2, 0x60, 0x8d, 0x7c, 0x97, 0xac, 0x82, 0xf1, 0x49, 0x53, 0x02, 0x2f, 0xc6,
	0xa0, 0xa1, 0x04, 0xbe, 0x1b, 0x68, 0xca, 0x9b, 0xd9, 0x53, 0xe7, 0x8c, 0x79, 0x31, 0x32, 0xd6,
	0x90, 0x18, 0xf3, 0x5d, 0xc6, 0x34, 0xed, 0xf2, 0xf3, 0xf2, 0x9d, 0xff, 0x3f, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xc9, 0xd9, 0x16, 0xaa, 0x18, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ManagedPlacementViewServiceClient is the client API for ManagedPlacementViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ManagedPlacementViewServiceClient interface {
	// Returns the requested Managed Placement view in full detail.
	GetManagedPlacementView(ctx context.Context, in *GetManagedPlacementViewRequest, opts ...grpc.CallOption) (*resources.ManagedPlacementView, error)
}

type managedPlacementViewServiceClient struct {
	cc *grpc.ClientConn
}

func NewManagedPlacementViewServiceClient(cc *grpc.ClientConn) ManagedPlacementViewServiceClient {
	return &managedPlacementViewServiceClient{cc}
}

func (c *managedPlacementViewServiceClient) GetManagedPlacementView(ctx context.Context, in *GetManagedPlacementViewRequest, opts ...grpc.CallOption) (*resources.ManagedPlacementView, error) {
	out := new(resources.ManagedPlacementView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.ManagedPlacementViewService/GetManagedPlacementView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagedPlacementViewServiceServer is the server API for ManagedPlacementViewService service.
type ManagedPlacementViewServiceServer interface {
	// Returns the requested Managed Placement view in full detail.
	GetManagedPlacementView(context.Context, *GetManagedPlacementViewRequest) (*resources.ManagedPlacementView, error)
}

func RegisterManagedPlacementViewServiceServer(s *grpc.Server, srv ManagedPlacementViewServiceServer) {
	s.RegisterService(&_ManagedPlacementViewService_serviceDesc, srv)
}

func _ManagedPlacementViewService_GetManagedPlacementView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetManagedPlacementViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagedPlacementViewServiceServer).GetManagedPlacementView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.ManagedPlacementViewService/GetManagedPlacementView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagedPlacementViewServiceServer).GetManagedPlacementView(ctx, req.(*GetManagedPlacementViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ManagedPlacementViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.ManagedPlacementViewService",
	HandlerType: (*ManagedPlacementViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetManagedPlacementView",
			Handler:    _ManagedPlacementViewService_GetManagedPlacementView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/managed_placement_view_service.proto",
}
