// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: ListContents.proto

package listcontent

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

// ListContentServicesClient is the client API for ListContentServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ListContentServicesClient interface {
	GetContent(ctx context.Context, in *ContentRequest, opts ...grpc.CallOption) (*ContentResponse, error)
}

type listContentServicesClient struct {
	cc grpc.ClientConnInterface
}

func NewListContentServicesClient(cc grpc.ClientConnInterface) ListContentServicesClient {
	return &listContentServicesClient{cc}
}

func (c *listContentServicesClient) GetContent(ctx context.Context, in *ContentRequest, opts ...grpc.CallOption) (*ContentResponse, error) {
	out := new(ContentResponse)
	err := c.cc.Invoke(ctx, "/listcontent.ListContentServices/GetContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListContentServicesServer is the server API for ListContentServices service.
// All implementations must embed UnimplementedListContentServicesServer
// for forward compatibility
type ListContentServicesServer interface {
	GetContent(context.Context, *ContentRequest) (*ContentResponse, error)
	mustEmbedUnimplementedListContentServicesServer()
}

// UnimplementedListContentServicesServer must be embedded to have forward compatible implementations.
type UnimplementedListContentServicesServer struct {
}

func (UnimplementedListContentServicesServer) GetContent(context.Context, *ContentRequest) (*ContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContent not implemented")
}
func (UnimplementedListContentServicesServer) mustEmbedUnimplementedListContentServicesServer() {}

// UnsafeListContentServicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ListContentServicesServer will
// result in compilation errors.
type UnsafeListContentServicesServer interface {
	mustEmbedUnimplementedListContentServicesServer()
}

func RegisterListContentServicesServer(s grpc.ServiceRegistrar, srv ListContentServicesServer) {
	s.RegisterService(&ListContentServices_ServiceDesc, srv)
}

func _ListContentServices_GetContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListContentServicesServer).GetContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/listcontent.ListContentServices/GetContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListContentServicesServer).GetContent(ctx, req.(*ContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ListContentServices_ServiceDesc is the grpc.ServiceDesc for ListContentServices service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ListContentServices_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "listcontent.ListContentServices",
	HandlerType: (*ListContentServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetContent",
			Handler:    _ListContentServices_GetContent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ListContents.proto",
}
