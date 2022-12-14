// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: shortener/v1alpha1/shortener.proto

package shortener

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ShortServiceClient is the client API for ShortService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShortServiceClient interface {
	Shorten(ctx context.Context, in *ShortenRequest, opts ...grpc.CallOption) (*ShortenResponse, error)
	Lengthen(ctx context.Context, in *LengthenRequest, opts ...grpc.CallOption) (*LengthenResponse, error)
}

type shortServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShortServiceClient(cc grpc.ClientConnInterface) ShortServiceClient {
	return &shortServiceClient{cc}
}

func (c *shortServiceClient) Shorten(ctx context.Context, in *ShortenRequest, opts ...grpc.CallOption) (*ShortenResponse, error) {
	out := new(ShortenResponse)
	err := c.cc.Invoke(ctx, "/shortener.v1alpha1.ShortService/Shorten", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortServiceClient) Lengthen(ctx context.Context, in *LengthenRequest, opts ...grpc.CallOption) (*LengthenResponse, error) {
	out := new(LengthenResponse)
	err := c.cc.Invoke(ctx, "/shortener.v1alpha1.ShortService/Lengthen", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShortServiceServer is the server API for ShortService service.
// All implementations must embed UnimplementedShortServiceServer
// for forward compatibility
type ShortServiceServer interface {
	Shorten(context.Context, *ShortenRequest) (*ShortenResponse, error)
	Lengthen(context.Context, *LengthenRequest) (*LengthenResponse, error)
	mustEmbedUnimplementedShortServiceServer()
}

// UnimplementedShortServiceServer must be embedded to have forward compatible implementations.
type UnimplementedShortServiceServer struct {
}

func (UnimplementedShortServiceServer) Shorten(context.Context, *ShortenRequest) (*ShortenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Shorten not implemented")
}
func (UnimplementedShortServiceServer) Lengthen(context.Context, *LengthenRequest) (*LengthenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Lengthen not implemented")
}
func (UnimplementedShortServiceServer) mustEmbedUnimplementedShortServiceServer() {}

// UnsafeShortServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShortServiceServer will
// result in compilation errors.
type UnsafeShortServiceServer interface {
	mustEmbedUnimplementedShortServiceServer()
}

func RegisterShortServiceServer(s grpc.ServiceRegistrar, srv ShortServiceServer) {
	s.RegisterService(&ShortService_ServiceDesc, srv)
}

func _ShortService_Shorten_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShortenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortServiceServer).Shorten(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortener.v1alpha1.ShortService/Shorten",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortServiceServer).Shorten(ctx, req.(*ShortenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShortService_Lengthen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LengthenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortServiceServer).Lengthen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortener.v1alpha1.ShortService/Lengthen",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortServiceServer).Lengthen(ctx, req.(*LengthenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ShortService_ServiceDesc is the grpc.ServiceDesc for ShortService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShortService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shortener.v1alpha1.ShortService",
	HandlerType: (*ShortServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Shorten",
			Handler:    _ShortService_Shorten_Handler,
		},
		{
			MethodName: "Lengthen",
			Handler:    _ShortService_Lengthen_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shortener/v1alpha1/shortener.proto",
}
