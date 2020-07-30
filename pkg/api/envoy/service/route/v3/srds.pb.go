// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/route/v3/srds.proto

package envoy_service_route_v3

import (
	context "context"
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/datawire/ambassador/pkg/api/envoy/annotations"
	v3 "github.com/datawire/ambassador/pkg/api/envoy/service/discovery/v3"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SrdsDummy struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SrdsDummy) Reset()         { *m = SrdsDummy{} }
func (m *SrdsDummy) String() string { return proto.CompactTextString(m) }
func (*SrdsDummy) ProtoMessage()    {}
func (*SrdsDummy) Descriptor() ([]byte, []int) {
	return fileDescriptor_e461818506941dcb, []int{0}
}

func (m *SrdsDummy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SrdsDummy.Unmarshal(m, b)
}
func (m *SrdsDummy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SrdsDummy.Marshal(b, m, deterministic)
}
func (m *SrdsDummy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SrdsDummy.Merge(m, src)
}
func (m *SrdsDummy) XXX_Size() int {
	return xxx_messageInfo_SrdsDummy.Size(m)
}
func (m *SrdsDummy) XXX_DiscardUnknown() {
	xxx_messageInfo_SrdsDummy.DiscardUnknown(m)
}

var xxx_messageInfo_SrdsDummy proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SrdsDummy)(nil), "envoy.service.route.v3.SrdsDummy")
}

func init() { proto.RegisterFile("envoy/service/route/v3/srds.proto", fileDescriptor_e461818506941dcb) }

var fileDescriptor_e461818506941dcb = []byte{
	// 395 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0xb1, 0xce, 0xd3, 0x30,
	0x14, 0x85, 0x71, 0x7f, 0xa9, 0x08, 0x4f, 0x34, 0x03, 0x42, 0xa1, 0x45, 0xb4, 0x62, 0x40, 0x55,
	0xeb, 0x94, 0x46, 0x62, 0xe8, 0xc0, 0x50, 0x2a, 0xe6, 0xaa, 0x79, 0x02, 0x93, 0x5c, 0x82, 0xa5,
	0xd6, 0x37, 0xd8, 0x8e, 0x21, 0x5b, 0xc5, 0x54, 0xb1, 0x22, 0x31, 0x30, 0x30, 0xf1, 0x14, 0xec,
	0x48, 0xac, 0x88, 0x57, 0xe0, 0x0d, 0xfa, 0x02, 0x28, 0x4e, 0x1a, 0x1a, 0x15, 0x21, 0x58, 0x18,
	0xad, 0xfb, 0xdd, 0x73, 0xee, 0xb1, 0x0e, 0x1d, 0x82, 0xb4, 0x58, 0x04, 0x1a, 0x94, 0x15, 0x31,
	0x04, 0x0a, 0x73, 0x03, 0x81, 0x0d, 0x03, 0xad, 0x12, 0xcd, 0x32, 0x85, 0x06, 0xbd, 0x5b, 0x0e,
	0x61, 0x35, 0xc2, 0x1c, 0xc2, 0x6c, 0xe8, 0x8f, 0xdb, 0xab, 0x89, 0xd0, 0x31, 0x5a, 0x50, 0x45,
	0xb9, 0xde, 0x3c, 0x2a, 0x0d, 0xbf, 0x9f, 0x22, 0xa6, 0x5b, 0x08, 0x78, 0x26, 0x02, 0x2e, 0x25,
	0x1a, 0x6e, 0x04, 0xca, 0xda, 0xc1, 0xbf, 0x57, 0x29, 0x9d, 0x0d, 0x02, 0x05, 0x1a, 0x73, 0x15,
	0x43, 0x4d, 0x0c, 0xf2, 0x24, 0xe3, 0x2d, 0x40, 0x1b, 0x6e, 0xf2, 0x93, 0xc0, 0xf0, 0x62, 0x6c,
	0x41, 0x69, 0x81, 0x52, 0xc8, 0xb4, 0x42, 0x46, 0x63, 0x7a, 0x23, 0x52, 0x89, 0x5e, 0xe5, 0xbb,
	0x5d, 0xb1, 0x18, 0x7c, 0xf8, 0x72, 0xb8, 0x7b, 0x9b, 0xd6, 0xc9, 0x78, 0x26, 0x98, 0x9d, 0xb3,
	0x66, 0x3c, 0x3f, 0x5e, 0xd1, 0x7e, 0x14, 0x63, 0x06, 0xc9, 0xa6, 0x0c, 0xab, 0x57, 0xa7, 0x34,
	0x51, 0x95, 0xd5, 0x7b, 0x45, 0xbd, 0xc8, 0x28, 0xe0, 0xbb, 0x73, 0xca, 0x9b, 0xb0, 0xf6, 0x4f,
	0xfd, 0xfa, 0x04, 0x1b, 0xb2, 0x46, 0x63, 0x03, 0x2f, 0x73, 0xd0, 0xc6, 0x9f, 0xfe, 0x25, 0xad,
	0x33, 0x94, 0x1a, 0x46, 0xd7, 0x1e, 0x90, 0x19, 0xf1, 0xf6, 0x84, 0xf6, 0x56, 0xb0, 0x35, 0xbc,
	0x65, 0xfc, 0xf0, 0x8f, 0x52, 0x25, 0x7e, 0xe1, 0x3e, 0xff, 0x97, 0x95, 0xd6, 0x09, 0x1f, 0x09,
	0xed, 0x3d, 0x05, 0x13, 0xbf, 0xf8, 0x7f, 0xd9, 0x27, 0x6f, 0xbe, 0xff, 0x78, 0xd7, 0x19, 0x8c,
	0xee, 0xb4, 0xaa, 0xb5, 0xd0, 0xee, 0x80, 0xa9, 0x2b, 0xa4, 0x76, 0xc8, 0xd5, 0x82, 0x8c, 0xfd,
	0x47, 0x6f, 0x3f, 0xbd, 0x3f, 0x5e, 0x9f, 0xd1, 0xda, 0x23, 0x46, 0xf9, 0x5c, 0xa4, 0x4d, 0x6d,
	0xd9, 0xd9, 0xd9, 0x4f, 0xdc, 0x28, 0x57, 0xae, 0x33, 0xcb, 0xc7, 0x9f, 0xf7, 0x5f, 0xbf, 0x75,
	0x3b, 0x37, 0x3b, 0xf4, 0xbe, 0xc0, 0x6a, 0x39, 0x53, 0xf8, 0xba, 0x60, 0xbf, 0xef, 0xff, 0xd2,
	0xf5, 0x69, 0x5d, 0x96, 0x6b, 0x4d, 0x0e, 0x84, 0x3c, 0xeb, 0xba, 0xa2, 0x85, 0x3f, 0x03, 0x00,
	0x00, 0xff, 0xff, 0xe4, 0x73, 0xdc, 0xb3, 0x53, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ScopedRoutesDiscoveryServiceClient is the client API for ScopedRoutesDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ScopedRoutesDiscoveryServiceClient interface {
	StreamScopedRoutes(ctx context.Context, opts ...grpc.CallOption) (ScopedRoutesDiscoveryService_StreamScopedRoutesClient, error)
	DeltaScopedRoutes(ctx context.Context, opts ...grpc.CallOption) (ScopedRoutesDiscoveryService_DeltaScopedRoutesClient, error)
	FetchScopedRoutes(ctx context.Context, in *v3.DiscoveryRequest, opts ...grpc.CallOption) (*v3.DiscoveryResponse, error)
}

type scopedRoutesDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewScopedRoutesDiscoveryServiceClient(cc *grpc.ClientConn) ScopedRoutesDiscoveryServiceClient {
	return &scopedRoutesDiscoveryServiceClient{cc}
}

func (c *scopedRoutesDiscoveryServiceClient) StreamScopedRoutes(ctx context.Context, opts ...grpc.CallOption) (ScopedRoutesDiscoveryService_StreamScopedRoutesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ScopedRoutesDiscoveryService_serviceDesc.Streams[0], "/envoy.service.route.v3.ScopedRoutesDiscoveryService/StreamScopedRoutes", opts...)
	if err != nil {
		return nil, err
	}
	x := &scopedRoutesDiscoveryServiceStreamScopedRoutesClient{stream}
	return x, nil
}

type ScopedRoutesDiscoveryService_StreamScopedRoutesClient interface {
	Send(*v3.DiscoveryRequest) error
	Recv() (*v3.DiscoveryResponse, error)
	grpc.ClientStream
}

type scopedRoutesDiscoveryServiceStreamScopedRoutesClient struct {
	grpc.ClientStream
}

func (x *scopedRoutesDiscoveryServiceStreamScopedRoutesClient) Send(m *v3.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *scopedRoutesDiscoveryServiceStreamScopedRoutesClient) Recv() (*v3.DiscoveryResponse, error) {
	m := new(v3.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *scopedRoutesDiscoveryServiceClient) DeltaScopedRoutes(ctx context.Context, opts ...grpc.CallOption) (ScopedRoutesDiscoveryService_DeltaScopedRoutesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ScopedRoutesDiscoveryService_serviceDesc.Streams[1], "/envoy.service.route.v3.ScopedRoutesDiscoveryService/DeltaScopedRoutes", opts...)
	if err != nil {
		return nil, err
	}
	x := &scopedRoutesDiscoveryServiceDeltaScopedRoutesClient{stream}
	return x, nil
}

type ScopedRoutesDiscoveryService_DeltaScopedRoutesClient interface {
	Send(*v3.DeltaDiscoveryRequest) error
	Recv() (*v3.DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type scopedRoutesDiscoveryServiceDeltaScopedRoutesClient struct {
	grpc.ClientStream
}

func (x *scopedRoutesDiscoveryServiceDeltaScopedRoutesClient) Send(m *v3.DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *scopedRoutesDiscoveryServiceDeltaScopedRoutesClient) Recv() (*v3.DeltaDiscoveryResponse, error) {
	m := new(v3.DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *scopedRoutesDiscoveryServiceClient) FetchScopedRoutes(ctx context.Context, in *v3.DiscoveryRequest, opts ...grpc.CallOption) (*v3.DiscoveryResponse, error) {
	out := new(v3.DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.service.route.v3.ScopedRoutesDiscoveryService/FetchScopedRoutes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScopedRoutesDiscoveryServiceServer is the server API for ScopedRoutesDiscoveryService service.
type ScopedRoutesDiscoveryServiceServer interface {
	StreamScopedRoutes(ScopedRoutesDiscoveryService_StreamScopedRoutesServer) error
	DeltaScopedRoutes(ScopedRoutesDiscoveryService_DeltaScopedRoutesServer) error
	FetchScopedRoutes(context.Context, *v3.DiscoveryRequest) (*v3.DiscoveryResponse, error)
}

// UnimplementedScopedRoutesDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedScopedRoutesDiscoveryServiceServer struct {
}

func (*UnimplementedScopedRoutesDiscoveryServiceServer) StreamScopedRoutes(srv ScopedRoutesDiscoveryService_StreamScopedRoutesServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamScopedRoutes not implemented")
}
func (*UnimplementedScopedRoutesDiscoveryServiceServer) DeltaScopedRoutes(srv ScopedRoutesDiscoveryService_DeltaScopedRoutesServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaScopedRoutes not implemented")
}
func (*UnimplementedScopedRoutesDiscoveryServiceServer) FetchScopedRoutes(ctx context.Context, req *v3.DiscoveryRequest) (*v3.DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchScopedRoutes not implemented")
}

func RegisterScopedRoutesDiscoveryServiceServer(s *grpc.Server, srv ScopedRoutesDiscoveryServiceServer) {
	s.RegisterService(&_ScopedRoutesDiscoveryService_serviceDesc, srv)
}

func _ScopedRoutesDiscoveryService_StreamScopedRoutes_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ScopedRoutesDiscoveryServiceServer).StreamScopedRoutes(&scopedRoutesDiscoveryServiceStreamScopedRoutesServer{stream})
}

type ScopedRoutesDiscoveryService_StreamScopedRoutesServer interface {
	Send(*v3.DiscoveryResponse) error
	Recv() (*v3.DiscoveryRequest, error)
	grpc.ServerStream
}

type scopedRoutesDiscoveryServiceStreamScopedRoutesServer struct {
	grpc.ServerStream
}

func (x *scopedRoutesDiscoveryServiceStreamScopedRoutesServer) Send(m *v3.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *scopedRoutesDiscoveryServiceStreamScopedRoutesServer) Recv() (*v3.DiscoveryRequest, error) {
	m := new(v3.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ScopedRoutesDiscoveryService_DeltaScopedRoutes_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ScopedRoutesDiscoveryServiceServer).DeltaScopedRoutes(&scopedRoutesDiscoveryServiceDeltaScopedRoutesServer{stream})
}

type ScopedRoutesDiscoveryService_DeltaScopedRoutesServer interface {
	Send(*v3.DeltaDiscoveryResponse) error
	Recv() (*v3.DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type scopedRoutesDiscoveryServiceDeltaScopedRoutesServer struct {
	grpc.ServerStream
}

func (x *scopedRoutesDiscoveryServiceDeltaScopedRoutesServer) Send(m *v3.DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *scopedRoutesDiscoveryServiceDeltaScopedRoutesServer) Recv() (*v3.DeltaDiscoveryRequest, error) {
	m := new(v3.DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ScopedRoutesDiscoveryService_FetchScopedRoutes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScopedRoutesDiscoveryServiceServer).FetchScopedRoutes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.route.v3.ScopedRoutesDiscoveryService/FetchScopedRoutes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScopedRoutesDiscoveryServiceServer).FetchScopedRoutes(ctx, req.(*v3.DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ScopedRoutesDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.route.v3.ScopedRoutesDiscoveryService",
	HandlerType: (*ScopedRoutesDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchScopedRoutes",
			Handler:    _ScopedRoutesDiscoveryService_FetchScopedRoutes_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamScopedRoutes",
			Handler:       _ScopedRoutesDiscoveryService_StreamScopedRoutes_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeltaScopedRoutes",
			Handler:       _ScopedRoutesDiscoveryService_DeltaScopedRoutes_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/route/v3/srds.proto",
}