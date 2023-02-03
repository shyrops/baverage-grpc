// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: beveragesmgmt.proto

package beverage_grpc

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

// BeveragesManagementClient is the client API for BeveragesManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BeveragesManagementClient interface {
	CreateBeverage(ctx context.Context, in *CreateBeverageRequest, opts ...grpc.CallOption) (*Beverage, error)
}

type beveragesManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewBeveragesManagementClient(cc grpc.ClientConnInterface) BeveragesManagementClient {
	return &beveragesManagementClient{cc}
}

func (c *beveragesManagementClient) CreateBeverage(ctx context.Context, in *CreateBeverageRequest, opts ...grpc.CallOption) (*Beverage, error) {
	out := new(Beverage)
	err := c.cc.Invoke(ctx, "/api.BeveragesManagement/CreateBeverage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BeveragesManagementServer is the server API for BeveragesManagement service.
// All implementations must embed UnimplementedBeveragesManagementServer
// for forward compatibility
type BeveragesManagementServer interface {
	CreateBeverage(context.Context, *CreateBeverageRequest) (*Beverage, error)
	mustEmbedUnimplementedBeveragesManagementServer()
}

// UnimplementedBeveragesManagementServer must be embedded to have forward compatible implementations.
type UnimplementedBeveragesManagementServer struct {
}

func (UnimplementedBeveragesManagementServer) CreateBeverage(context.Context, *CreateBeverageRequest) (*Beverage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBeverage not implemented")
}
func (UnimplementedBeveragesManagementServer) mustEmbedUnimplementedBeveragesManagementServer() {}

// UnsafeBeveragesManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BeveragesManagementServer will
// result in compilation errors.
type UnsafeBeveragesManagementServer interface {
	mustEmbedUnimplementedBeveragesManagementServer()
}

func RegisterBeveragesManagementServer(s grpc.ServiceRegistrar, srv BeveragesManagementServer) {
	s.RegisterService(&BeveragesManagement_ServiceDesc, srv)
}

func _BeveragesManagement_CreateBeverage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBeverageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeveragesManagementServer).CreateBeverage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.BeveragesManagement/CreateBeverage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeveragesManagementServer).CreateBeverage(ctx, req.(*CreateBeverageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BeveragesManagement_ServiceDesc is the grpc.ServiceDesc for BeveragesManagement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BeveragesManagement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.BeveragesManagement",
	HandlerType: (*BeveragesManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBeverage",
			Handler:    _BeveragesManagement_CreateBeverage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "beveragesmgmt.proto",
}
