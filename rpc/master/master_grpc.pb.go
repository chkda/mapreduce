// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.20.3
// source: rpc/master/master.proto

package master

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Master_Trigger_FullMethodName            = "/Master/Trigger"
	Master_RegisterWorker_FullMethodName     = "/Master/RegisterWorker"
	Master_UpdateMapResult_FullMethodName    = "/Master/UpdateMapResult"
	Master_UpdateDataNodes_FullMethodName    = "/Master/UpdateDataNodes"
	Master_UpdateReduceResult_FullMethodName = "/Master/UpdateReduceResult"
)

// MasterClient is the client API for Master service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MasterClient interface {
	Trigger(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*MasterAck, error)
	RegisterWorker(ctx context.Context, in *WorkerInfo, opts ...grpc.CallOption) (*MasterAck, error)
	UpdateMapResult(ctx context.Context, in *MapResult, opts ...grpc.CallOption) (*MasterAck, error)
	UpdateDataNodes(ctx context.Context, in *DFSDataNodesInfo, opts ...grpc.CallOption) (*MasterAck, error)
	UpdateReduceResult(ctx context.Context, in *ReduceResult, opts ...grpc.CallOption) (*MasterAck, error)
}

type masterClient struct {
	cc grpc.ClientConnInterface
}

func NewMasterClient(cc grpc.ClientConnInterface) MasterClient {
	return &masterClient{cc}
}

func (c *masterClient) Trigger(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*MasterAck, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MasterAck)
	err := c.cc.Invoke(ctx, Master_Trigger_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) RegisterWorker(ctx context.Context, in *WorkerInfo, opts ...grpc.CallOption) (*MasterAck, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MasterAck)
	err := c.cc.Invoke(ctx, Master_RegisterWorker_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) UpdateMapResult(ctx context.Context, in *MapResult, opts ...grpc.CallOption) (*MasterAck, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MasterAck)
	err := c.cc.Invoke(ctx, Master_UpdateMapResult_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) UpdateDataNodes(ctx context.Context, in *DFSDataNodesInfo, opts ...grpc.CallOption) (*MasterAck, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MasterAck)
	err := c.cc.Invoke(ctx, Master_UpdateDataNodes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) UpdateReduceResult(ctx context.Context, in *ReduceResult, opts ...grpc.CallOption) (*MasterAck, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MasterAck)
	err := c.cc.Invoke(ctx, Master_UpdateReduceResult_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MasterServer is the server API for Master service.
// All implementations must embed UnimplementedMasterServer
// for forward compatibility
type MasterServer interface {
	Trigger(context.Context, *TaskRequest) (*MasterAck, error)
	RegisterWorker(context.Context, *WorkerInfo) (*MasterAck, error)
	UpdateMapResult(context.Context, *MapResult) (*MasterAck, error)
	UpdateDataNodes(context.Context, *DFSDataNodesInfo) (*MasterAck, error)
	UpdateReduceResult(context.Context, *ReduceResult) (*MasterAck, error)
	mustEmbedUnimplementedMasterServer()
}

// UnimplementedMasterServer must be embedded to have forward compatible implementations.
type UnimplementedMasterServer struct {
}

func (UnimplementedMasterServer) Trigger(context.Context, *TaskRequest) (*MasterAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Trigger not implemented")
}
func (UnimplementedMasterServer) RegisterWorker(context.Context, *WorkerInfo) (*MasterAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterWorker not implemented")
}
func (UnimplementedMasterServer) UpdateMapResult(context.Context, *MapResult) (*MasterAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMapResult not implemented")
}
func (UnimplementedMasterServer) UpdateDataNodes(context.Context, *DFSDataNodesInfo) (*MasterAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDataNodes not implemented")
}
func (UnimplementedMasterServer) UpdateReduceResult(context.Context, *ReduceResult) (*MasterAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateReduceResult not implemented")
}
func (UnimplementedMasterServer) mustEmbedUnimplementedMasterServer() {}

// UnsafeMasterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MasterServer will
// result in compilation errors.
type UnsafeMasterServer interface {
	mustEmbedUnimplementedMasterServer()
}

func RegisterMasterServer(s grpc.ServiceRegistrar, srv MasterServer) {
	s.RegisterService(&Master_ServiceDesc, srv)
}

func _Master_Trigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).Trigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Master_Trigger_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).Trigger(ctx, req.(*TaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_RegisterWorker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkerInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).RegisterWorker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Master_RegisterWorker_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).RegisterWorker(ctx, req.(*WorkerInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_UpdateMapResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MapResult)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).UpdateMapResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Master_UpdateMapResult_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).UpdateMapResult(ctx, req.(*MapResult))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_UpdateDataNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DFSDataNodesInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).UpdateDataNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Master_UpdateDataNodes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).UpdateDataNodes(ctx, req.(*DFSDataNodesInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_UpdateReduceResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReduceResult)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).UpdateReduceResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Master_UpdateReduceResult_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).UpdateReduceResult(ctx, req.(*ReduceResult))
	}
	return interceptor(ctx, in, info, handler)
}

// Master_ServiceDesc is the grpc.ServiceDesc for Master service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Master_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Master",
	HandlerType: (*MasterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Trigger",
			Handler:    _Master_Trigger_Handler,
		},
		{
			MethodName: "RegisterWorker",
			Handler:    _Master_RegisterWorker_Handler,
		},
		{
			MethodName: "UpdateMapResult",
			Handler:    _Master_UpdateMapResult_Handler,
		},
		{
			MethodName: "UpdateDataNodes",
			Handler:    _Master_UpdateDataNodes_Handler,
		},
		{
			MethodName: "UpdateReduceResult",
			Handler:    _Master_UpdateReduceResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/master/master.proto",
}
