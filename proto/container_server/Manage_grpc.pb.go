// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package Manage

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ManagerClient is the client API for Manager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManagerClient interface {
	//Container utils
	ListContainers(ctx context.Context, in *ListContainers_Request, opts ...grpc.CallOption) (*ListContainers_Response, error)
	PruneContainers(ctx context.Context, in *PruneContainers_Request, opts ...grpc.CallOption) (*PruneContainers_Response, error)
	GetContainer(ctx context.Context, in *GetContainer_Request, opts ...grpc.CallOption) (*GetContainer_Response, error)
	StartContainer(ctx context.Context, in *StartContainer_Request, opts ...grpc.CallOption) (*StartContainer_Response, error)
	StopContainer(ctx context.Context, in *StopContainer_Request, opts ...grpc.CallOption) (*StopContainer_Response, error)
	RestartContainer(ctx context.Context, in *RestartContainer_Request, opts ...grpc.CallOption) (*RestartContainer_Response, error)
	RemoveContainer(ctx context.Context, in *RemoveContainer_Request, opts ...grpc.CallOption) (*RemoveContainer_Response, error)
	CreateContainer(ctx context.Context, in *CreateContainer_Request, opts ...grpc.CallOption) (*CreateContainer_Response, error)
	//所有的data为tar
	GetFile(ctx context.Context, in *GetFile_Request, opts ...grpc.CallOption) (Manager_GetFileClient, error)
	UpdateFile(ctx context.Context, in *UpdateFile_Request, opts ...grpc.CallOption) (*UpdateFile_Response, error)
	ListFile(ctx context.Context, in *ListFile_Request, opts ...grpc.CallOption) (*ListFile_Response, error)
	//Image utils
	PruneImages(ctx context.Context, in *PruneImages_Request, opts ...grpc.CallOption) (*PruneImages_Response, error)
	ListImages(ctx context.Context, in *ListImages_Request, opts ...grpc.CallOption) (*ListImages_Response, error)
	PullImage(ctx context.Context, in *PullImage_Request, opts ...grpc.CallOption) (*PullImage_Response, error)
	BuildImage(ctx context.Context, in *BuildImage_Request, opts ...grpc.CallOption) (*BuildImage_Response, error)
	LoadImage(ctx context.Context, opts ...grpc.CallOption) (Manager_LoadImageClient, error)
	GetImage(ctx context.Context, in *GetImage_Request, opts ...grpc.CallOption) (*GetImage_Response, error)
	RemoveImage(ctx context.Context, in *RemoveImage_Request, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type managerClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerClient(cc grpc.ClientConnInterface) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) ListContainers(ctx context.Context, in *ListContainers_Request, opts ...grpc.CallOption) (*ListContainers_Response, error) {
	out := new(ListContainers_Response)
	err := c.cc.Invoke(ctx, "/Manager/ListContainers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) PruneContainers(ctx context.Context, in *PruneContainers_Request, opts ...grpc.CallOption) (*PruneContainers_Response, error) {
	out := new(PruneContainers_Response)
	err := c.cc.Invoke(ctx, "/Manager/PruneContainers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetContainer(ctx context.Context, in *GetContainer_Request, opts ...grpc.CallOption) (*GetContainer_Response, error) {
	out := new(GetContainer_Response)
	err := c.cc.Invoke(ctx, "/Manager/GetContainer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) StartContainer(ctx context.Context, in *StartContainer_Request, opts ...grpc.CallOption) (*StartContainer_Response, error) {
	out := new(StartContainer_Response)
	err := c.cc.Invoke(ctx, "/Manager/StartContainer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) StopContainer(ctx context.Context, in *StopContainer_Request, opts ...grpc.CallOption) (*StopContainer_Response, error) {
	out := new(StopContainer_Response)
	err := c.cc.Invoke(ctx, "/Manager/StopContainer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) RestartContainer(ctx context.Context, in *RestartContainer_Request, opts ...grpc.CallOption) (*RestartContainer_Response, error) {
	out := new(RestartContainer_Response)
	err := c.cc.Invoke(ctx, "/Manager/RestartContainer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) RemoveContainer(ctx context.Context, in *RemoveContainer_Request, opts ...grpc.CallOption) (*RemoveContainer_Response, error) {
	out := new(RemoveContainer_Response)
	err := c.cc.Invoke(ctx, "/Manager/RemoveContainer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CreateContainer(ctx context.Context, in *CreateContainer_Request, opts ...grpc.CallOption) (*CreateContainer_Response, error) {
	out := new(CreateContainer_Response)
	err := c.cc.Invoke(ctx, "/Manager/CreateContainer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetFile(ctx context.Context, in *GetFile_Request, opts ...grpc.CallOption) (Manager_GetFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &Manager_ServiceDesc.Streams[0], "/Manager/GetFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &managerGetFileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Manager_GetFileClient interface {
	Recv() (*GetFile_Response, error)
	grpc.ClientStream
}

type managerGetFileClient struct {
	grpc.ClientStream
}

func (x *managerGetFileClient) Recv() (*GetFile_Response, error) {
	m := new(GetFile_Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *managerClient) UpdateFile(ctx context.Context, in *UpdateFile_Request, opts ...grpc.CallOption) (*UpdateFile_Response, error) {
	out := new(UpdateFile_Response)
	err := c.cc.Invoke(ctx, "/Manager/UpdateFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ListFile(ctx context.Context, in *ListFile_Request, opts ...grpc.CallOption) (*ListFile_Response, error) {
	out := new(ListFile_Response)
	err := c.cc.Invoke(ctx, "/Manager/ListFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) PruneImages(ctx context.Context, in *PruneImages_Request, opts ...grpc.CallOption) (*PruneImages_Response, error) {
	out := new(PruneImages_Response)
	err := c.cc.Invoke(ctx, "/Manager/PruneImages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ListImages(ctx context.Context, in *ListImages_Request, opts ...grpc.CallOption) (*ListImages_Response, error) {
	out := new(ListImages_Response)
	err := c.cc.Invoke(ctx, "/Manager/ListImages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) PullImage(ctx context.Context, in *PullImage_Request, opts ...grpc.CallOption) (*PullImage_Response, error) {
	out := new(PullImage_Response)
	err := c.cc.Invoke(ctx, "/Manager/PullImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) BuildImage(ctx context.Context, in *BuildImage_Request, opts ...grpc.CallOption) (*BuildImage_Response, error) {
	out := new(BuildImage_Response)
	err := c.cc.Invoke(ctx, "/Manager/BuildImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) LoadImage(ctx context.Context, opts ...grpc.CallOption) (Manager_LoadImageClient, error) {
	stream, err := c.cc.NewStream(ctx, &Manager_ServiceDesc.Streams[1], "/Manager/LoadImage", opts...)
	if err != nil {
		return nil, err
	}
	x := &managerLoadImageClient{stream}
	return x, nil
}

type Manager_LoadImageClient interface {
	Send(*LoadImage_Request) error
	CloseAndRecv() (*LoadImage_Response, error)
	grpc.ClientStream
}

type managerLoadImageClient struct {
	grpc.ClientStream
}

func (x *managerLoadImageClient) Send(m *LoadImage_Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *managerLoadImageClient) CloseAndRecv() (*LoadImage_Response, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(LoadImage_Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *managerClient) GetImage(ctx context.Context, in *GetImage_Request, opts ...grpc.CallOption) (*GetImage_Response, error) {
	out := new(GetImage_Response)
	err := c.cc.Invoke(ctx, "/Manager/GetImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) RemoveImage(ctx context.Context, in *RemoveImage_Request, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Manager/RemoveImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerServer is the server API for Manager service.
// All implementations must embed UnimplementedManagerServer
// for forward compatibility
type ManagerServer interface {
	//Container utils
	ListContainers(context.Context, *ListContainers_Request) (*ListContainers_Response, error)
	PruneContainers(context.Context, *PruneContainers_Request) (*PruneContainers_Response, error)
	GetContainer(context.Context, *GetContainer_Request) (*GetContainer_Response, error)
	StartContainer(context.Context, *StartContainer_Request) (*StartContainer_Response, error)
	StopContainer(context.Context, *StopContainer_Request) (*StopContainer_Response, error)
	RestartContainer(context.Context, *RestartContainer_Request) (*RestartContainer_Response, error)
	RemoveContainer(context.Context, *RemoveContainer_Request) (*RemoveContainer_Response, error)
	CreateContainer(context.Context, *CreateContainer_Request) (*CreateContainer_Response, error)
	//所有的data为tar
	GetFile(*GetFile_Request, Manager_GetFileServer) error
	UpdateFile(context.Context, *UpdateFile_Request) (*UpdateFile_Response, error)
	ListFile(context.Context, *ListFile_Request) (*ListFile_Response, error)
	//Image utils
	PruneImages(context.Context, *PruneImages_Request) (*PruneImages_Response, error)
	ListImages(context.Context, *ListImages_Request) (*ListImages_Response, error)
	PullImage(context.Context, *PullImage_Request) (*PullImage_Response, error)
	BuildImage(context.Context, *BuildImage_Request) (*BuildImage_Response, error)
	LoadImage(Manager_LoadImageServer) error
	GetImage(context.Context, *GetImage_Request) (*GetImage_Response, error)
	RemoveImage(context.Context, *RemoveImage_Request) (*emptypb.Empty, error)
	mustEmbedUnimplementedManagerServer()
}

// UnimplementedManagerServer must be embedded to have forward compatible implementations.
type UnimplementedManagerServer struct {
}

func (UnimplementedManagerServer) ListContainers(context.Context, *ListContainers_Request) (*ListContainers_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListContainers not implemented")
}
func (UnimplementedManagerServer) PruneContainers(context.Context, *PruneContainers_Request) (*PruneContainers_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PruneContainers not implemented")
}
func (UnimplementedManagerServer) GetContainer(context.Context, *GetContainer_Request) (*GetContainer_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContainer not implemented")
}
func (UnimplementedManagerServer) StartContainer(context.Context, *StartContainer_Request) (*StartContainer_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartContainer not implemented")
}
func (UnimplementedManagerServer) StopContainer(context.Context, *StopContainer_Request) (*StopContainer_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopContainer not implemented")
}
func (UnimplementedManagerServer) RestartContainer(context.Context, *RestartContainer_Request) (*RestartContainer_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RestartContainer not implemented")
}
func (UnimplementedManagerServer) RemoveContainer(context.Context, *RemoveContainer_Request) (*RemoveContainer_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveContainer not implemented")
}
func (UnimplementedManagerServer) CreateContainer(context.Context, *CreateContainer_Request) (*CreateContainer_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateContainer not implemented")
}
func (UnimplementedManagerServer) GetFile(*GetFile_Request, Manager_GetFileServer) error {
	return status.Errorf(codes.Unimplemented, "method GetFile not implemented")
}
func (UnimplementedManagerServer) UpdateFile(context.Context, *UpdateFile_Request) (*UpdateFile_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFile not implemented")
}
func (UnimplementedManagerServer) ListFile(context.Context, *ListFile_Request) (*ListFile_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFile not implemented")
}
func (UnimplementedManagerServer) PruneImages(context.Context, *PruneImages_Request) (*PruneImages_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PruneImages not implemented")
}
func (UnimplementedManagerServer) ListImages(context.Context, *ListImages_Request) (*ListImages_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListImages not implemented")
}
func (UnimplementedManagerServer) PullImage(context.Context, *PullImage_Request) (*PullImage_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PullImage not implemented")
}
func (UnimplementedManagerServer) BuildImage(context.Context, *BuildImage_Request) (*BuildImage_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuildImage not implemented")
}
func (UnimplementedManagerServer) LoadImage(Manager_LoadImageServer) error {
	return status.Errorf(codes.Unimplemented, "method LoadImage not implemented")
}
func (UnimplementedManagerServer) GetImage(context.Context, *GetImage_Request) (*GetImage_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImage not implemented")
}
func (UnimplementedManagerServer) RemoveImage(context.Context, *RemoveImage_Request) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveImage not implemented")
}
func (UnimplementedManagerServer) mustEmbedUnimplementedManagerServer() {}

// UnsafeManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManagerServer will
// result in compilation errors.
type UnsafeManagerServer interface {
	mustEmbedUnimplementedManagerServer()
}

func RegisterManagerServer(s grpc.ServiceRegistrar, srv ManagerServer) {
	s.RegisterService(&Manager_ServiceDesc, srv)
}

func _Manager_ListContainers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListContainers_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ListContainers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Manager/ListContainers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ListContainers(ctx, req.(*ListContainers_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_PruneContainers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PruneContainers_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).PruneContainers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Manager/PruneContainers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).PruneContainers(ctx, req.(*PruneContainers_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContainer_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Manager/GetContainer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetContainer(ctx, req.(*GetContainer_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_StartContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartContainer_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).StartContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Manager/StartContainer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).StartContainer(ctx, req.(*StartContainer_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_StopContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopContainer_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).StopContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Manager/StopContainer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).StopContainer(ctx, req.(*StopContainer_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_RestartContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestartContainer_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).RestartContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Manager/RestartContainer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).RestartContainer(ctx, req.(*RestartContainer_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_RemoveContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveContainer_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).RemoveContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Manager/RemoveContainer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).RemoveContainer(ctx, req.(*RemoveContainer_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CreateContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateContainer_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Manager/CreateContainer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateContainer(ctx, req.(*CreateContainer_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetFile_Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ManagerServer).GetFile(m, &managerGetFileServer{stream})
}

type Manager_GetFileServer interface {
	Send(*GetFile_Response) error
	grpc.ServerStream
}

type managerGetFileServer struct {
	grpc.ServerStream
}

func (x *managerGetFileServer) Send(m *GetFile_Response) error {
	return x.ServerStream.SendMsg(m)
}

func _Manager_UpdateFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFile_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).UpdateFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Manager/UpdateFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).UpdateFile(ctx, req.(*UpdateFile_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ListFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFile_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ListFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Manager/ListFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ListFile(ctx, req.(*ListFile_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_PruneImages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PruneImages_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).PruneImages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Manager/PruneImages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).PruneImages(ctx, req.(*PruneImages_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ListImages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListImages_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ListImages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Manager/ListImages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ListImages(ctx, req.(*ListImages_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_PullImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PullImage_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).PullImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Manager/PullImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).PullImage(ctx, req.(*PullImage_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_BuildImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildImage_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).BuildImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Manager/BuildImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).BuildImage(ctx, req.(*BuildImage_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_LoadImage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ManagerServer).LoadImage(&managerLoadImageServer{stream})
}

type Manager_LoadImageServer interface {
	SendAndClose(*LoadImage_Response) error
	Recv() (*LoadImage_Request, error)
	grpc.ServerStream
}

type managerLoadImageServer struct {
	grpc.ServerStream
}

func (x *managerLoadImageServer) SendAndClose(m *LoadImage_Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *managerLoadImageServer) Recv() (*LoadImage_Request, error) {
	m := new(LoadImage_Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Manager_GetImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetImage_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Manager/GetImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetImage(ctx, req.(*GetImage_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_RemoveImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveImage_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).RemoveImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Manager/RemoveImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).RemoveImage(ctx, req.(*RemoveImage_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Manager_ServiceDesc is the grpc.ServiceDesc for Manager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Manager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListContainers",
			Handler:    _Manager_ListContainers_Handler,
		},
		{
			MethodName: "PruneContainers",
			Handler:    _Manager_PruneContainers_Handler,
		},
		{
			MethodName: "GetContainer",
			Handler:    _Manager_GetContainer_Handler,
		},
		{
			MethodName: "StartContainer",
			Handler:    _Manager_StartContainer_Handler,
		},
		{
			MethodName: "StopContainer",
			Handler:    _Manager_StopContainer_Handler,
		},
		{
			MethodName: "RestartContainer",
			Handler:    _Manager_RestartContainer_Handler,
		},
		{
			MethodName: "RemoveContainer",
			Handler:    _Manager_RemoveContainer_Handler,
		},
		{
			MethodName: "CreateContainer",
			Handler:    _Manager_CreateContainer_Handler,
		},
		{
			MethodName: "UpdateFile",
			Handler:    _Manager_UpdateFile_Handler,
		},
		{
			MethodName: "ListFile",
			Handler:    _Manager_ListFile_Handler,
		},
		{
			MethodName: "PruneImages",
			Handler:    _Manager_PruneImages_Handler,
		},
		{
			MethodName: "ListImages",
			Handler:    _Manager_ListImages_Handler,
		},
		{
			MethodName: "PullImage",
			Handler:    _Manager_PullImage_Handler,
		},
		{
			MethodName: "BuildImage",
			Handler:    _Manager_BuildImage_Handler,
		},
		{
			MethodName: "GetImage",
			Handler:    _Manager_GetImage_Handler,
		},
		{
			MethodName: "RemoveImage",
			Handler:    _Manager_RemoveImage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetFile",
			Handler:       _Manager_GetFile_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "LoadImage",
			Handler:       _Manager_LoadImage_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/container_server/Manage.proto",
}
