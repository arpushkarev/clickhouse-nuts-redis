// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: service.proto

package item

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

// ItemV1Client is the client API for ItemV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ItemV1Client interface {
	Post(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*PostResponse, error)
	Get(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetResponse, error)
	Patch(ctx context.Context, in *PatchRequest, opts ...grpc.CallOption) (*PatchResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type itemV1Client struct {
	cc grpc.ClientConnInterface
}

func NewItemV1Client(cc grpc.ClientConnInterface) ItemV1Client {
	return &itemV1Client{cc}
}

func (c *itemV1Client) Post(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*PostResponse, error) {
	out := new(PostResponse)
	err := c.cc.Invoke(ctx, "/api.item_v1.ItemV1/Post", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemV1Client) Get(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/api.item_v1.ItemV1/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemV1Client) Patch(ctx context.Context, in *PatchRequest, opts ...grpc.CallOption) (*PatchResponse, error) {
	out := new(PatchResponse)
	err := c.cc.Invoke(ctx, "/api.item_v1.ItemV1/Patch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemV1Client) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/api.item_v1.ItemV1/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ItemV1Server is the server API for ItemV1 service.
// All implementations must embed UnimplementedItemV1Server
// for forward compatibility
type ItemV1Server interface {
	Post(context.Context, *PostRequest) (*PostResponse, error)
	Get(context.Context, *emptypb.Empty) (*GetResponse, error)
	Patch(context.Context, *PatchRequest) (*PatchResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	mustEmbedUnimplementedItemV1Server()
}

// UnimplementedItemV1Server must be embedded to have forward compatible implementations.
type UnimplementedItemV1Server struct {
}

func (UnimplementedItemV1Server) Post(context.Context, *PostRequest) (*PostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Post not implemented")
}
func (UnimplementedItemV1Server) Get(context.Context, *emptypb.Empty) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedItemV1Server) Patch(context.Context, *PatchRequest) (*PatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Patch not implemented")
}
func (UnimplementedItemV1Server) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedItemV1Server) mustEmbedUnimplementedItemV1Server() {}

// UnsafeItemV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ItemV1Server will
// result in compilation errors.
type UnsafeItemV1Server interface {
	mustEmbedUnimplementedItemV1Server()
}

func RegisterItemV1Server(s grpc.ServiceRegistrar, srv ItemV1Server) {
	s.RegisterService(&ItemV1_ServiceDesc, srv)
}

func _ItemV1_Post_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemV1Server).Post(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.item_v1.ItemV1/Post",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemV1Server).Post(ctx, req.(*PostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemV1_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemV1Server).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.item_v1.ItemV1/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemV1Server).Get(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemV1_Patch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemV1Server).Patch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.item_v1.ItemV1/Patch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemV1Server).Patch(ctx, req.(*PatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemV1_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemV1Server).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.item_v1.ItemV1/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemV1Server).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ItemV1_ServiceDesc is the grpc.ServiceDesc for ItemV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ItemV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.item_v1.ItemV1",
	HandlerType: (*ItemV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Post",
			Handler:    _ItemV1_Post_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ItemV1_Get_Handler,
		},
		{
			MethodName: "Patch",
			Handler:    _ItemV1_Patch_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ItemV1_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}