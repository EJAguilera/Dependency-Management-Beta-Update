// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: packagebirdservices.proto

package services

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

// PackagebirdServicesClient is the client API for PackagebirdServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PackagebirdServicesClient interface {
	// Create Operations
	CreateProject(ctx context.Context, in *ProjectRequest, opts ...grpc.CallOption) (*OperationResponse, error)
	CreatePackage(ctx context.Context, in *PackageRequest, opts ...grpc.CallOption) (*OperationResponse, error)
	// Dependency Operations
	AddPackage(ctx context.Context, in *PackageRequest, opts ...grpc.CallOption) (*OperationResponse, error)
	RemovePackage(ctx context.Context, in *PackageRequest, opts ...grpc.CallOption) (*OperationResponse, error)
	// Package Operations
	BuildPackage(ctx context.Context, in *PackageRequest, opts ...grpc.CallOption) (*OperationResponse, error)
	TestPackage(ctx context.Context, in *PackageRequest, opts ...grpc.CallOption) (*OperationResponse, error)
	// Get Operations
	GetPackages(ctx context.Context, in *Blank, opts ...grpc.CallOption) (*PackageList, error)
	GetProjects(ctx context.Context, in *Blank, opts ...grpc.CallOption) (*ProjectList, error)
	// File Transfer Operations
	DownloadFile(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (PackagebirdServices_DownloadFileClient, error)
	UploadFile(ctx context.Context, opts ...grpc.CallOption) (PackagebirdServices_UploadFileClient, error)
}

type packagebirdServicesClient struct {
	cc grpc.ClientConnInterface
}

func NewPackagebirdServicesClient(cc grpc.ClientConnInterface) PackagebirdServicesClient {
	return &packagebirdServicesClient{cc}
}

func (c *packagebirdServicesClient) CreateProject(ctx context.Context, in *ProjectRequest, opts ...grpc.CallOption) (*OperationResponse, error) {
	out := new(OperationResponse)
	err := c.cc.Invoke(ctx, "/services.PackagebirdServices/CreateProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packagebirdServicesClient) CreatePackage(ctx context.Context, in *PackageRequest, opts ...grpc.CallOption) (*OperationResponse, error) {
	out := new(OperationResponse)
	err := c.cc.Invoke(ctx, "/services.PackagebirdServices/CreatePackage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packagebirdServicesClient) AddPackage(ctx context.Context, in *PackageRequest, opts ...grpc.CallOption) (*OperationResponse, error) {
	out := new(OperationResponse)
	err := c.cc.Invoke(ctx, "/services.PackagebirdServices/AddPackage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packagebirdServicesClient) RemovePackage(ctx context.Context, in *PackageRequest, opts ...grpc.CallOption) (*OperationResponse, error) {
	out := new(OperationResponse)
	err := c.cc.Invoke(ctx, "/services.PackagebirdServices/RemovePackage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packagebirdServicesClient) BuildPackage(ctx context.Context, in *PackageRequest, opts ...grpc.CallOption) (*OperationResponse, error) {
	out := new(OperationResponse)
	err := c.cc.Invoke(ctx, "/services.PackagebirdServices/BuildPackage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packagebirdServicesClient) TestPackage(ctx context.Context, in *PackageRequest, opts ...grpc.CallOption) (*OperationResponse, error) {
	out := new(OperationResponse)
	err := c.cc.Invoke(ctx, "/services.PackagebirdServices/TestPackage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packagebirdServicesClient) GetPackages(ctx context.Context, in *Blank, opts ...grpc.CallOption) (*PackageList, error) {
	out := new(PackageList)
	err := c.cc.Invoke(ctx, "/services.PackagebirdServices/GetPackages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packagebirdServicesClient) GetProjects(ctx context.Context, in *Blank, opts ...grpc.CallOption) (*ProjectList, error) {
	out := new(ProjectList)
	err := c.cc.Invoke(ctx, "/services.PackagebirdServices/GetProjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packagebirdServicesClient) DownloadFile(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (PackagebirdServices_DownloadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &PackagebirdServices_ServiceDesc.Streams[0], "/services.PackagebirdServices/DownloadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &packagebirdServicesDownloadFileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PackagebirdServices_DownloadFileClient interface {
	Recv() (*File, error)
	grpc.ClientStream
}

type packagebirdServicesDownloadFileClient struct {
	grpc.ClientStream
}

func (x *packagebirdServicesDownloadFileClient) Recv() (*File, error) {
	m := new(File)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *packagebirdServicesClient) UploadFile(ctx context.Context, opts ...grpc.CallOption) (PackagebirdServices_UploadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &PackagebirdServices_ServiceDesc.Streams[1], "/services.PackagebirdServices/UploadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &packagebirdServicesUploadFileClient{stream}
	return x, nil
}

type PackagebirdServices_UploadFileClient interface {
	Send(*File) error
	CloseAndRecv() (*OperationResponse, error)
	grpc.ClientStream
}

type packagebirdServicesUploadFileClient struct {
	grpc.ClientStream
}

func (x *packagebirdServicesUploadFileClient) Send(m *File) error {
	return x.ClientStream.SendMsg(m)
}

func (x *packagebirdServicesUploadFileClient) CloseAndRecv() (*OperationResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(OperationResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PackagebirdServicesServer is the server API for PackagebirdServices service.
// All implementations must embed UnimplementedPackagebirdServicesServer
// for forward compatibility
type PackagebirdServicesServer interface {
	// Create Operations
	CreateProject(context.Context, *ProjectRequest) (*OperationResponse, error)
	CreatePackage(context.Context, *PackageRequest) (*OperationResponse, error)
	// Dependency Operations
	AddPackage(context.Context, *PackageRequest) (*OperationResponse, error)
	RemovePackage(context.Context, *PackageRequest) (*OperationResponse, error)
	// Package Operations
	BuildPackage(context.Context, *PackageRequest) (*OperationResponse, error)
	TestPackage(context.Context, *PackageRequest) (*OperationResponse, error)
	// Get Operations
	GetPackages(context.Context, *Blank) (*PackageList, error)
	GetProjects(context.Context, *Blank) (*ProjectList, error)
	// File Transfer Operations
	DownloadFile(*DownloadRequest, PackagebirdServices_DownloadFileServer) error
	UploadFile(PackagebirdServices_UploadFileServer) error
	mustEmbedUnimplementedPackagebirdServicesServer()
}

// UnimplementedPackagebirdServicesServer must be embedded to have forward compatible implementations.
type UnimplementedPackagebirdServicesServer struct {
}

func (UnimplementedPackagebirdServicesServer) CreateProject(context.Context, *ProjectRequest) (*OperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProject not implemented")
}
func (UnimplementedPackagebirdServicesServer) CreatePackage(context.Context, *PackageRequest) (*OperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePackage not implemented")
}
func (UnimplementedPackagebirdServicesServer) AddPackage(context.Context, *PackageRequest) (*OperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPackage not implemented")
}
func (UnimplementedPackagebirdServicesServer) RemovePackage(context.Context, *PackageRequest) (*OperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemovePackage not implemented")
}
func (UnimplementedPackagebirdServicesServer) BuildPackage(context.Context, *PackageRequest) (*OperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuildPackage not implemented")
}
func (UnimplementedPackagebirdServicesServer) TestPackage(context.Context, *PackageRequest) (*OperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestPackage not implemented")
}
func (UnimplementedPackagebirdServicesServer) GetPackages(context.Context, *Blank) (*PackageList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPackages not implemented")
}
func (UnimplementedPackagebirdServicesServer) GetProjects(context.Context, *Blank) (*ProjectList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProjects not implemented")
}
func (UnimplementedPackagebirdServicesServer) DownloadFile(*DownloadRequest, PackagebirdServices_DownloadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method DownloadFile not implemented")
}
func (UnimplementedPackagebirdServicesServer) UploadFile(PackagebirdServices_UploadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedPackagebirdServicesServer) mustEmbedUnimplementedPackagebirdServicesServer() {}

// UnsafePackagebirdServicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PackagebirdServicesServer will
// result in compilation errors.
type UnsafePackagebirdServicesServer interface {
	mustEmbedUnimplementedPackagebirdServicesServer()
}

func RegisterPackagebirdServicesServer(s grpc.ServiceRegistrar, srv PackagebirdServicesServer) {
	s.RegisterService(&PackagebirdServices_ServiceDesc, srv)
}

func _PackagebirdServices_CreateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackagebirdServicesServer).CreateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.PackagebirdServices/CreateProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackagebirdServicesServer).CreateProject(ctx, req.(*ProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PackagebirdServices_CreatePackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackagebirdServicesServer).CreatePackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.PackagebirdServices/CreatePackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackagebirdServicesServer).CreatePackage(ctx, req.(*PackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PackagebirdServices_AddPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackagebirdServicesServer).AddPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.PackagebirdServices/AddPackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackagebirdServicesServer).AddPackage(ctx, req.(*PackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PackagebirdServices_RemovePackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackagebirdServicesServer).RemovePackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.PackagebirdServices/RemovePackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackagebirdServicesServer).RemovePackage(ctx, req.(*PackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PackagebirdServices_BuildPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackagebirdServicesServer).BuildPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.PackagebirdServices/BuildPackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackagebirdServicesServer).BuildPackage(ctx, req.(*PackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PackagebirdServices_TestPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackagebirdServicesServer).TestPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.PackagebirdServices/TestPackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackagebirdServicesServer).TestPackage(ctx, req.(*PackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PackagebirdServices_GetPackages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Blank)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackagebirdServicesServer).GetPackages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.PackagebirdServices/GetPackages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackagebirdServicesServer).GetPackages(ctx, req.(*Blank))
	}
	return interceptor(ctx, in, info, handler)
}

func _PackagebirdServices_GetProjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Blank)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackagebirdServicesServer).GetProjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.PackagebirdServices/GetProjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackagebirdServicesServer).GetProjects(ctx, req.(*Blank))
	}
	return interceptor(ctx, in, info, handler)
}

func _PackagebirdServices_DownloadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PackagebirdServicesServer).DownloadFile(m, &packagebirdServicesDownloadFileServer{stream})
}

type PackagebirdServices_DownloadFileServer interface {
	Send(*File) error
	grpc.ServerStream
}

type packagebirdServicesDownloadFileServer struct {
	grpc.ServerStream
}

func (x *packagebirdServicesDownloadFileServer) Send(m *File) error {
	return x.ServerStream.SendMsg(m)
}

func _PackagebirdServices_UploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PackagebirdServicesServer).UploadFile(&packagebirdServicesUploadFileServer{stream})
}

type PackagebirdServices_UploadFileServer interface {
	SendAndClose(*OperationResponse) error
	Recv() (*File, error)
	grpc.ServerStream
}

type packagebirdServicesUploadFileServer struct {
	grpc.ServerStream
}

func (x *packagebirdServicesUploadFileServer) SendAndClose(m *OperationResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *packagebirdServicesUploadFileServer) Recv() (*File, error) {
	m := new(File)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PackagebirdServices_ServiceDesc is the grpc.ServiceDesc for PackagebirdServices service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PackagebirdServices_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.PackagebirdServices",
	HandlerType: (*PackagebirdServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProject",
			Handler:    _PackagebirdServices_CreateProject_Handler,
		},
		{
			MethodName: "CreatePackage",
			Handler:    _PackagebirdServices_CreatePackage_Handler,
		},
		{
			MethodName: "AddPackage",
			Handler:    _PackagebirdServices_AddPackage_Handler,
		},
		{
			MethodName: "RemovePackage",
			Handler:    _PackagebirdServices_RemovePackage_Handler,
		},
		{
			MethodName: "BuildPackage",
			Handler:    _PackagebirdServices_BuildPackage_Handler,
		},
		{
			MethodName: "TestPackage",
			Handler:    _PackagebirdServices_TestPackage_Handler,
		},
		{
			MethodName: "GetPackages",
			Handler:    _PackagebirdServices_GetPackages_Handler,
		},
		{
			MethodName: "GetProjects",
			Handler:    _PackagebirdServices_GetProjects_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DownloadFile",
			Handler:       _PackagebirdServices_DownloadFile_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "UploadFile",
			Handler:       _PackagebirdServices_UploadFile_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "packagebirdservices.proto",
}
