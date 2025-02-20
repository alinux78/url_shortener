// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.6.1
// source: service.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	UrlShortenerService_Shorten_FullMethodName = "/service.UrlShortenerService/Shorten"
	UrlShortenerService_Resolve_FullMethodName = "/service.UrlShortenerService/Resolve"
)

// UrlShortenerServiceClient is the client API for UrlShortenerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UrlShortenerServiceClient interface {
	Shorten(ctx context.Context, in *UrlShortenRequest, opts ...grpc.CallOption) (*UrlShortenResponse, error)
	Resolve(ctx context.Context, in *UrlResolveRequest, opts ...grpc.CallOption) (*UrlResolveResponse, error)
}

type urlShortenerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUrlShortenerServiceClient(cc grpc.ClientConnInterface) UrlShortenerServiceClient {
	return &urlShortenerServiceClient{cc}
}

func (c *urlShortenerServiceClient) Shorten(ctx context.Context, in *UrlShortenRequest, opts ...grpc.CallOption) (*UrlShortenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UrlShortenResponse)
	err := c.cc.Invoke(ctx, UrlShortenerService_Shorten_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *urlShortenerServiceClient) Resolve(ctx context.Context, in *UrlResolveRequest, opts ...grpc.CallOption) (*UrlResolveResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UrlResolveResponse)
	err := c.cc.Invoke(ctx, UrlShortenerService_Resolve_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UrlShortenerServiceServer is the server API for UrlShortenerService service.
// All implementations must embed UnimplementedUrlShortenerServiceServer
// for forward compatibility.
type UrlShortenerServiceServer interface {
	Shorten(context.Context, *UrlShortenRequest) (*UrlShortenResponse, error)
	Resolve(context.Context, *UrlResolveRequest) (*UrlResolveResponse, error)
	mustEmbedUnimplementedUrlShortenerServiceServer()
}

// UnimplementedUrlShortenerServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUrlShortenerServiceServer struct{}

func (UnimplementedUrlShortenerServiceServer) Shorten(context.Context, *UrlShortenRequest) (*UrlShortenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Shorten not implemented")
}
func (UnimplementedUrlShortenerServiceServer) Resolve(context.Context, *UrlResolveRequest) (*UrlResolveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Resolve not implemented")
}
func (UnimplementedUrlShortenerServiceServer) mustEmbedUnimplementedUrlShortenerServiceServer() {}
func (UnimplementedUrlShortenerServiceServer) testEmbeddedByValue()                             {}

// UnsafeUrlShortenerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UrlShortenerServiceServer will
// result in compilation errors.
type UnsafeUrlShortenerServiceServer interface {
	mustEmbedUnimplementedUrlShortenerServiceServer()
}

func RegisterUrlShortenerServiceServer(s grpc.ServiceRegistrar, srv UrlShortenerServiceServer) {
	// If the following call pancis, it indicates UnimplementedUrlShortenerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UrlShortenerService_ServiceDesc, srv)
}

func _UrlShortenerService_Shorten_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UrlShortenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlShortenerServiceServer).Shorten(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UrlShortenerService_Shorten_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlShortenerServiceServer).Shorten(ctx, req.(*UrlShortenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UrlShortenerService_Resolve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UrlResolveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlShortenerServiceServer).Resolve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UrlShortenerService_Resolve_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlShortenerServiceServer).Resolve(ctx, req.(*UrlResolveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UrlShortenerService_ServiceDesc is the grpc.ServiceDesc for UrlShortenerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UrlShortenerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.UrlShortenerService",
	HandlerType: (*UrlShortenerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Shorten",
			Handler:    _UrlShortenerService_Shorten_Handler,
		},
		{
			MethodName: "Resolve",
			Handler:    _UrlShortenerService_Resolve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
