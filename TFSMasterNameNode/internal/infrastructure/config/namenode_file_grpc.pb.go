// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: namenode_file.proto

package config

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

// LocateChunkClient is the client API for LocateChunk service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LocateChunkClient interface {
	LocateChunk(ctx context.Context, in *ChunkLocation, opts ...grpc.CallOption) (*Response, error)
}

type locateChunkClient struct {
	cc grpc.ClientConnInterface
}

func NewLocateChunkClient(cc grpc.ClientConnInterface) LocateChunkClient {
	return &locateChunkClient{cc}
}

func (c *locateChunkClient) LocateChunk(ctx context.Context, in *ChunkLocation, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/LocateChunk/LocateChunk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LocateChunkServer is the server API for LocateChunk service.
// All implementations must embed UnimplementedLocateChunkServer
// for forward compatibility
type LocateChunkServer interface {
	LocateChunk(context.Context, *ChunkLocation) (*Response, error)
	mustEmbedUnimplementedLocateChunkServer()
}

// UnimplementedLocateChunkServer must be embedded to have forward compatible implementations.
type UnimplementedLocateChunkServer struct {
}

func (UnimplementedLocateChunkServer) LocateChunk(context.Context, *ChunkLocation) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LocateChunk not implemented")
}
func (UnimplementedLocateChunkServer) mustEmbedUnimplementedLocateChunkServer() {}

// UnsafeLocateChunkServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LocateChunkServer will
// result in compilation errors.
type UnsafeLocateChunkServer interface {
	mustEmbedUnimplementedLocateChunkServer()
}

func RegisterLocateChunkServer(s grpc.ServiceRegistrar, srv LocateChunkServer) {
	s.RegisterService(&LocateChunk_ServiceDesc, srv)
}

func _LocateChunk_LocateChunk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChunkLocation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocateChunkServer).LocateChunk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LocateChunk/LocateChunk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocateChunkServer).LocateChunk(ctx, req.(*ChunkLocation))
	}
	return interceptor(ctx, in, info, handler)
}

// LocateChunk_ServiceDesc is the grpc.ServiceDesc for LocateChunk service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LocateChunk_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "LocateChunk",
	HandlerType: (*LocateChunkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LocateChunk",
			Handler:    _LocateChunk_LocateChunk_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "namenode_file.proto",
}

// HeartBeatClient is the client API for HeartBeat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HeartBeatClient interface {
	HeartBeat(ctx context.Context, in *Socket, opts ...grpc.CallOption) (*Response, error)
}

type heartBeatClient struct {
	cc grpc.ClientConnInterface
}

func NewHeartBeatClient(cc grpc.ClientConnInterface) HeartBeatClient {
	return &heartBeatClient{cc}
}

func (c *heartBeatClient) HeartBeat(ctx context.Context, in *Socket, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/HeartBeat/HeartBeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HeartBeatServer is the server API for HeartBeat service.
// All implementations must embed UnimplementedHeartBeatServer
// for forward compatibility
type HeartBeatServer interface {
	HeartBeat(context.Context, *Socket) (*Response, error)
	mustEmbedUnimplementedHeartBeatServer()
}

// UnimplementedHeartBeatServer must be embedded to have forward compatible implementations.
type UnimplementedHeartBeatServer struct {
}

func (UnimplementedHeartBeatServer) HeartBeat(context.Context, *Socket) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HeartBeat not implemented")
}
func (UnimplementedHeartBeatServer) mustEmbedUnimplementedHeartBeatServer() {}

// UnsafeHeartBeatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HeartBeatServer will
// result in compilation errors.
type UnsafeHeartBeatServer interface {
	mustEmbedUnimplementedHeartBeatServer()
}

func RegisterHeartBeatServer(s grpc.ServiceRegistrar, srv HeartBeatServer) {
	s.RegisterService(&HeartBeat_ServiceDesc, srv)
}

func _HeartBeat_HeartBeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Socket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeartBeatServer).HeartBeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HeartBeat/HeartBeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeartBeatServer).HeartBeat(ctx, req.(*Socket))
	}
	return interceptor(ctx, in, info, handler)
}

// HeartBeat_ServiceDesc is the grpc.ServiceDesc for HeartBeat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HeartBeat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "HeartBeat",
	HandlerType: (*HeartBeatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HeartBeat",
			Handler:    _HeartBeat_HeartBeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "namenode_file.proto",
}

// FetchClient is the client API for Fetch service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FetchClient interface {
	FetchSocket(ctx context.Context, in *VoidRequest, opts ...grpc.CallOption) (*Sockets, error)
	FetchLocations(ctx context.Context, in *VoidRequest, opts ...grpc.CallOption) (*Metadata, error)
	FetchMetadata(ctx context.Context, in *FileMetadataFetch, opts ...grpc.CallOption) (*Response, error)
	DeleteMetadataFetch(ctx context.Context, in *DeleteMetadata, opts ...grpc.CallOption) (*Response, error)
	FetchLocateChunk(ctx context.Context, in *ChunkLocation, opts ...grpc.CallOption) (*Response, error)
	FetchHeartBeat(ctx context.Context, in *Socket, opts ...grpc.CallOption) (*Response, error)
}

type fetchClient struct {
	cc grpc.ClientConnInterface
}

func NewFetchClient(cc grpc.ClientConnInterface) FetchClient {
	return &fetchClient{cc}
}

func (c *fetchClient) FetchSocket(ctx context.Context, in *VoidRequest, opts ...grpc.CallOption) (*Sockets, error) {
	out := new(Sockets)
	err := c.cc.Invoke(ctx, "/Fetch/FetchSocket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fetchClient) FetchLocations(ctx context.Context, in *VoidRequest, opts ...grpc.CallOption) (*Metadata, error) {
	out := new(Metadata)
	err := c.cc.Invoke(ctx, "/Fetch/FetchLocations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fetchClient) FetchMetadata(ctx context.Context, in *FileMetadataFetch, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/Fetch/FetchMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fetchClient) DeleteMetadataFetch(ctx context.Context, in *DeleteMetadata, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/Fetch/DeleteMetadataFetch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fetchClient) FetchLocateChunk(ctx context.Context, in *ChunkLocation, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/Fetch/FetchLocateChunk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fetchClient) FetchHeartBeat(ctx context.Context, in *Socket, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/Fetch/FetchHeartBeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FetchServer is the server API for Fetch service.
// All implementations must embed UnimplementedFetchServer
// for forward compatibility
type FetchServer interface {
	FetchSocket(context.Context, *VoidRequest) (*Sockets, error)
	FetchLocations(context.Context, *VoidRequest) (*Metadata, error)
	FetchMetadata(context.Context, *FileMetadataFetch) (*Response, error)
	DeleteMetadataFetch(context.Context, *DeleteMetadata) (*Response, error)
	FetchLocateChunk(context.Context, *ChunkLocation) (*Response, error)
	FetchHeartBeat(context.Context, *Socket) (*Response, error)
	mustEmbedUnimplementedFetchServer()
}

// UnimplementedFetchServer must be embedded to have forward compatible implementations.
type UnimplementedFetchServer struct {
}

func (UnimplementedFetchServer) FetchSocket(context.Context, *VoidRequest) (*Sockets, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchSocket not implemented")
}
func (UnimplementedFetchServer) FetchLocations(context.Context, *VoidRequest) (*Metadata, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchLocations not implemented")
}
func (UnimplementedFetchServer) FetchMetadata(context.Context, *FileMetadataFetch) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchMetadata not implemented")
}
func (UnimplementedFetchServer) DeleteMetadataFetch(context.Context, *DeleteMetadata) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMetadataFetch not implemented")
}
func (UnimplementedFetchServer) FetchLocateChunk(context.Context, *ChunkLocation) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchLocateChunk not implemented")
}
func (UnimplementedFetchServer) FetchHeartBeat(context.Context, *Socket) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchHeartBeat not implemented")
}
func (UnimplementedFetchServer) mustEmbedUnimplementedFetchServer() {}

// UnsafeFetchServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FetchServer will
// result in compilation errors.
type UnsafeFetchServer interface {
	mustEmbedUnimplementedFetchServer()
}

func RegisterFetchServer(s grpc.ServiceRegistrar, srv FetchServer) {
	s.RegisterService(&Fetch_ServiceDesc, srv)
}

func _Fetch_FetchSocket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FetchServer).FetchSocket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Fetch/FetchSocket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FetchServer).FetchSocket(ctx, req.(*VoidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fetch_FetchLocations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FetchServer).FetchLocations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Fetch/FetchLocations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FetchServer).FetchLocations(ctx, req.(*VoidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fetch_FetchMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileMetadataFetch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FetchServer).FetchMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Fetch/FetchMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FetchServer).FetchMetadata(ctx, req.(*FileMetadataFetch))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fetch_DeleteMetadataFetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMetadata)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FetchServer).DeleteMetadataFetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Fetch/DeleteMetadataFetch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FetchServer).DeleteMetadataFetch(ctx, req.(*DeleteMetadata))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fetch_FetchLocateChunk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChunkLocation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FetchServer).FetchLocateChunk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Fetch/FetchLocateChunk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FetchServer).FetchLocateChunk(ctx, req.(*ChunkLocation))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fetch_FetchHeartBeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Socket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FetchServer).FetchHeartBeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Fetch/FetchHeartBeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FetchServer).FetchHeartBeat(ctx, req.(*Socket))
	}
	return interceptor(ctx, in, info, handler)
}

// Fetch_ServiceDesc is the grpc.ServiceDesc for Fetch service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Fetch_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Fetch",
	HandlerType: (*FetchServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchSocket",
			Handler:    _Fetch_FetchSocket_Handler,
		},
		{
			MethodName: "FetchLocations",
			Handler:    _Fetch_FetchLocations_Handler,
		},
		{
			MethodName: "FetchMetadata",
			Handler:    _Fetch_FetchMetadata_Handler,
		},
		{
			MethodName: "DeleteMetadataFetch",
			Handler:    _Fetch_DeleteMetadataFetch_Handler,
		},
		{
			MethodName: "FetchLocateChunk",
			Handler:    _Fetch_FetchLocateChunk_Handler,
		},
		{
			MethodName: "FetchHeartBeat",
			Handler:    _Fetch_FetchHeartBeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "namenode_file.proto",
}
