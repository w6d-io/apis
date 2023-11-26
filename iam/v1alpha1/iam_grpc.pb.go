// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1alpha1

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

// IamClient is the client API for Iam service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IamClient interface {
	AddPermission(ctx context.Context, in *Input, opts ...grpc.CallOption) (*Reply, error)
	RemovePermission(ctx context.Context, in *Input, opts ...grpc.CallOption) (*Reply, error)
	ReplacePermission(ctx context.Context, in *Input, opts ...grpc.CallOption) (*Reply, error)
}

type iamClient struct {
	cc grpc.ClientConnInterface
}

func NewIamClient(cc grpc.ClientConnInterface) IamClient {
	return &iamClient{cc}
}

func (c *iamClient) AddPermission(ctx context.Context, in *Input, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/permission.Iam/AddPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iamClient) RemovePermission(ctx context.Context, in *Input, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/permission.Iam/RemovePermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iamClient) ReplacePermission(ctx context.Context, in *Input, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/permission.Iam/ReplacePermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IamServer is the server API for Iam service.
// All implementations must embed UnimplementedIamServer
// for forward compatibility
type IamServer interface {
	AddPermission(context.Context, *Input) (*Reply, error)
	RemovePermission(context.Context, *Input) (*Reply, error)
	ReplacePermission(context.Context, *Input) (*Reply, error)
	mustEmbedUnimplementedIamServer()
}

// UnimplementedIamServer must be embedded to have forward compatible implementations.
type UnimplementedIamServer struct {
}

func (UnimplementedIamServer) AddPermission(context.Context, *Input) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPermission not implemented")
}
func (UnimplementedIamServer) RemovePermission(context.Context, *Input) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemovePermission not implemented")
}
func (UnimplementedIamServer) ReplacePermission(context.Context, *Input) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplacePermission not implemented")
}
func (UnimplementedIamServer) mustEmbedUnimplementedIamServer() {}

// UnsafeIamServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IamServer will
// result in compilation errors.
type UnsafeIamServer interface {
	mustEmbedUnimplementedIamServer()
}

func RegisterIamServer(s grpc.ServiceRegistrar, srv IamServer) {
	s.RegisterService(&Iam_ServiceDesc, srv)
}

func _Iam_AddPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Input)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IamServer).AddPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/permission.Iam/AddPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IamServer).AddPermission(ctx, req.(*Input))
	}
	return interceptor(ctx, in, info, handler)
}

func _Iam_RemovePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Input)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IamServer).RemovePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/permission.Iam/RemovePermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IamServer).RemovePermission(ctx, req.(*Input))
	}
	return interceptor(ctx, in, info, handler)
}

func _Iam_ReplacePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Input)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IamServer).ReplacePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/permission.Iam/ReplacePermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IamServer).ReplacePermission(ctx, req.(*Input))
	}
	return interceptor(ctx, in, info, handler)
}

// Iam_ServiceDesc is the grpc.ServiceDesc for Iam service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Iam_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "permission.Iam",
	HandlerType: (*IamServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddPermission",
			Handler:    _Iam_AddPermission_Handler,
		},
		{
			MethodName: "RemovePermission",
			Handler:    _Iam_RemovePermission_Handler,
		},
		{
			MethodName: "ReplacePermission",
			Handler:    _Iam_ReplacePermission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "iam.proto",
}
