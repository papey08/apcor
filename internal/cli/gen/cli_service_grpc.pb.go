// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package gen

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

// CliServiceClient is the client API for CliService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CliServiceClient interface {
	InstallDeps(ctx context.Context, in *InstallDepsRequest, opts ...grpc.CallOption) (*InstallDepsResponse, error)
	ExecPreTasks(ctx context.Context, in *ExecPreTasksRequest, opts ...grpc.CallOption) (*ExecPreTasksResponse, error)
	ExecPostTasks(ctx context.Context, in *ExecPostTasksRequest, opts ...grpc.CallOption) (*ExecPostTasksResponse, error)
	Clone(ctx context.Context, in *CloneRequest, opts ...grpc.CallOption) (*CloneResponse, error)
	Checkout(ctx context.Context, in *CheckoutRequest, opts ...grpc.CallOption) (*CheckoutResponse, error)
	Branch(ctx context.Context, in *BranchRequest, opts ...grpc.CallOption) (*BranchResponse, error)
	Pull(ctx context.Context, in *PullRequest, opts ...grpc.CallOption) (*PullResponse, error)
	PreBuild(ctx context.Context, in *PreBuildRequest, opts ...grpc.CallOption) (*PreBuildResponse, error)
	Build(ctx context.Context, in *BuildRequest, opts ...grpc.CallOption) (*BuildResponse, error)
	Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (*RunResponse, error)
	Debug(ctx context.Context, in *DebugRequest, opts ...grpc.CallOption) (*DebugResponse, error)
	Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error)
}

type cliServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCliServiceClient(cc grpc.ClientConnInterface) CliServiceClient {
	return &cliServiceClient{cc}
}

func (c *cliServiceClient) InstallDeps(ctx context.Context, in *InstallDepsRequest, opts ...grpc.CallOption) (*InstallDepsResponse, error) {
	out := new(InstallDepsResponse)
	err := c.cc.Invoke(ctx, "/apcor.CliService/InstallDeps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cliServiceClient) ExecPreTasks(ctx context.Context, in *ExecPreTasksRequest, opts ...grpc.CallOption) (*ExecPreTasksResponse, error) {
	out := new(ExecPreTasksResponse)
	err := c.cc.Invoke(ctx, "/apcor.CliService/ExecPreTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cliServiceClient) ExecPostTasks(ctx context.Context, in *ExecPostTasksRequest, opts ...grpc.CallOption) (*ExecPostTasksResponse, error) {
	out := new(ExecPostTasksResponse)
	err := c.cc.Invoke(ctx, "/apcor.CliService/ExecPostTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cliServiceClient) Clone(ctx context.Context, in *CloneRequest, opts ...grpc.CallOption) (*CloneResponse, error) {
	out := new(CloneResponse)
	err := c.cc.Invoke(ctx, "/apcor.CliService/Clone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cliServiceClient) Checkout(ctx context.Context, in *CheckoutRequest, opts ...grpc.CallOption) (*CheckoutResponse, error) {
	out := new(CheckoutResponse)
	err := c.cc.Invoke(ctx, "/apcor.CliService/Checkout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cliServiceClient) Branch(ctx context.Context, in *BranchRequest, opts ...grpc.CallOption) (*BranchResponse, error) {
	out := new(BranchResponse)
	err := c.cc.Invoke(ctx, "/apcor.CliService/Branch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cliServiceClient) Pull(ctx context.Context, in *PullRequest, opts ...grpc.CallOption) (*PullResponse, error) {
	out := new(PullResponse)
	err := c.cc.Invoke(ctx, "/apcor.CliService/Pull", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cliServiceClient) PreBuild(ctx context.Context, in *PreBuildRequest, opts ...grpc.CallOption) (*PreBuildResponse, error) {
	out := new(PreBuildResponse)
	err := c.cc.Invoke(ctx, "/apcor.CliService/PreBuild", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cliServiceClient) Build(ctx context.Context, in *BuildRequest, opts ...grpc.CallOption) (*BuildResponse, error) {
	out := new(BuildResponse)
	err := c.cc.Invoke(ctx, "/apcor.CliService/Build", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cliServiceClient) Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (*RunResponse, error) {
	out := new(RunResponse)
	err := c.cc.Invoke(ctx, "/apcor.CliService/Run", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cliServiceClient) Debug(ctx context.Context, in *DebugRequest, opts ...grpc.CallOption) (*DebugResponse, error) {
	out := new(DebugResponse)
	err := c.cc.Invoke(ctx, "/apcor.CliService/Debug", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cliServiceClient) Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error) {
	out := new(StopResponse)
	err := c.cc.Invoke(ctx, "/apcor.CliService/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CliServiceServer is the server API for CliService service.
// All implementations should embed UnimplementedCliServiceServer
// for forward compatibility
type CliServiceServer interface {
	InstallDeps(context.Context, *InstallDepsRequest) (*InstallDepsResponse, error)
	ExecPreTasks(context.Context, *ExecPreTasksRequest) (*ExecPreTasksResponse, error)
	ExecPostTasks(context.Context, *ExecPostTasksRequest) (*ExecPostTasksResponse, error)
	Clone(context.Context, *CloneRequest) (*CloneResponse, error)
	Checkout(context.Context, *CheckoutRequest) (*CheckoutResponse, error)
	Branch(context.Context, *BranchRequest) (*BranchResponse, error)
	Pull(context.Context, *PullRequest) (*PullResponse, error)
	PreBuild(context.Context, *PreBuildRequest) (*PreBuildResponse, error)
	Build(context.Context, *BuildRequest) (*BuildResponse, error)
	Run(context.Context, *RunRequest) (*RunResponse, error)
	Debug(context.Context, *DebugRequest) (*DebugResponse, error)
	Stop(context.Context, *StopRequest) (*StopResponse, error)
}

// UnimplementedCliServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCliServiceServer struct {
}

func (UnimplementedCliServiceServer) InstallDeps(context.Context, *InstallDepsRequest) (*InstallDepsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InstallDeps not implemented")
}
func (UnimplementedCliServiceServer) ExecPreTasks(context.Context, *ExecPreTasksRequest) (*ExecPreTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecPreTasks not implemented")
}
func (UnimplementedCliServiceServer) ExecPostTasks(context.Context, *ExecPostTasksRequest) (*ExecPostTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecPostTasks not implemented")
}
func (UnimplementedCliServiceServer) Clone(context.Context, *CloneRequest) (*CloneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Clone not implemented")
}
func (UnimplementedCliServiceServer) Checkout(context.Context, *CheckoutRequest) (*CheckoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Checkout not implemented")
}
func (UnimplementedCliServiceServer) Branch(context.Context, *BranchRequest) (*BranchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Branch not implemented")
}
func (UnimplementedCliServiceServer) Pull(context.Context, *PullRequest) (*PullResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pull not implemented")
}
func (UnimplementedCliServiceServer) PreBuild(context.Context, *PreBuildRequest) (*PreBuildResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreBuild not implemented")
}
func (UnimplementedCliServiceServer) Build(context.Context, *BuildRequest) (*BuildResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Build not implemented")
}
func (UnimplementedCliServiceServer) Run(context.Context, *RunRequest) (*RunResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Run not implemented")
}
func (UnimplementedCliServiceServer) Debug(context.Context, *DebugRequest) (*DebugResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Debug not implemented")
}
func (UnimplementedCliServiceServer) Stop(context.Context, *StopRequest) (*StopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}

// UnsafeCliServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CliServiceServer will
// result in compilation errors.
type UnsafeCliServiceServer interface {
	mustEmbedUnimplementedCliServiceServer()
}

func RegisterCliServiceServer(s grpc.ServiceRegistrar, srv CliServiceServer) {
	s.RegisterService(&CliService_ServiceDesc, srv)
}

func _CliService_InstallDeps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstallDepsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CliServiceServer).InstallDeps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apcor.CliService/InstallDeps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CliServiceServer).InstallDeps(ctx, req.(*InstallDepsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CliService_ExecPreTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecPreTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CliServiceServer).ExecPreTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apcor.CliService/ExecPreTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CliServiceServer).ExecPreTasks(ctx, req.(*ExecPreTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CliService_ExecPostTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecPostTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CliServiceServer).ExecPostTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apcor.CliService/ExecPostTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CliServiceServer).ExecPostTasks(ctx, req.(*ExecPostTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CliService_Clone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CliServiceServer).Clone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apcor.CliService/Clone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CliServiceServer).Clone(ctx, req.(*CloneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CliService_Checkout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CliServiceServer).Checkout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apcor.CliService/Checkout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CliServiceServer).Checkout(ctx, req.(*CheckoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CliService_Branch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BranchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CliServiceServer).Branch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apcor.CliService/Branch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CliServiceServer).Branch(ctx, req.(*BranchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CliService_Pull_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PullRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CliServiceServer).Pull(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apcor.CliService/Pull",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CliServiceServer).Pull(ctx, req.(*PullRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CliService_PreBuild_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PreBuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CliServiceServer).PreBuild(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apcor.CliService/PreBuild",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CliServiceServer).PreBuild(ctx, req.(*PreBuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CliService_Build_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CliServiceServer).Build(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apcor.CliService/Build",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CliServiceServer).Build(ctx, req.(*BuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CliService_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CliServiceServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apcor.CliService/Run",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CliServiceServer).Run(ctx, req.(*RunRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CliService_Debug_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DebugRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CliServiceServer).Debug(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apcor.CliService/Debug",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CliServiceServer).Debug(ctx, req.(*DebugRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CliService_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CliServiceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apcor.CliService/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CliServiceServer).Stop(ctx, req.(*StopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CliService_ServiceDesc is the grpc.ServiceDesc for CliService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CliService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "apcor.CliService",
	HandlerType: (*CliServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InstallDeps",
			Handler:    _CliService_InstallDeps_Handler,
		},
		{
			MethodName: "ExecPreTasks",
			Handler:    _CliService_ExecPreTasks_Handler,
		},
		{
			MethodName: "ExecPostTasks",
			Handler:    _CliService_ExecPostTasks_Handler,
		},
		{
			MethodName: "Clone",
			Handler:    _CliService_Clone_Handler,
		},
		{
			MethodName: "Checkout",
			Handler:    _CliService_Checkout_Handler,
		},
		{
			MethodName: "Branch",
			Handler:    _CliService_Branch_Handler,
		},
		{
			MethodName: "Pull",
			Handler:    _CliService_Pull_Handler,
		},
		{
			MethodName: "PreBuild",
			Handler:    _CliService_PreBuild_Handler,
		},
		{
			MethodName: "Build",
			Handler:    _CliService_Build_Handler,
		},
		{
			MethodName: "Run",
			Handler:    _CliService_Run_Handler,
		},
		{
			MethodName: "Debug",
			Handler:    _CliService_Debug_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _CliService_Stop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cli_service.proto",
}
