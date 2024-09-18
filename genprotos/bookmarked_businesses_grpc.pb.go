// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: bookmarked_businesses.proto

package genprotos

import (
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

// Bookmarked_BusinessesClient is the client API for Bookmarked_Businesses service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Bookmarked_BusinessesClient interface {
}

type bookmarked_BusinessesClient struct {
	cc grpc.ClientConnInterface
}

func NewBookmarked_BusinessesClient(cc grpc.ClientConnInterface) Bookmarked_BusinessesClient {
	return &bookmarked_BusinessesClient{cc}
}

// Bookmarked_BusinessesServer is the server API for Bookmarked_Businesses service.
// All implementations must embed UnimplementedBookmarked_BusinessesServer
// for forward compatibility.
type Bookmarked_BusinessesServer interface {
	mustEmbedUnimplementedBookmarked_BusinessesServer()
}

// UnimplementedBookmarked_BusinessesServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBookmarked_BusinessesServer struct{}

func (UnimplementedBookmarked_BusinessesServer) mustEmbedUnimplementedBookmarked_BusinessesServer() {}
func (UnimplementedBookmarked_BusinessesServer) testEmbeddedByValue()                               {}

// UnsafeBookmarked_BusinessesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Bookmarked_BusinessesServer will
// result in compilation errors.
type UnsafeBookmarked_BusinessesServer interface {
	mustEmbedUnimplementedBookmarked_BusinessesServer()
}

func RegisterBookmarked_BusinessesServer(s grpc.ServiceRegistrar, srv Bookmarked_BusinessesServer) {
	// If the following call pancis, it indicates UnimplementedBookmarked_BusinessesServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Bookmarked_Businesses_ServiceDesc, srv)
}

// Bookmarked_Businesses_ServiceDesc is the grpc.ServiceDesc for Bookmarked_Businesses service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Bookmarked_Businesses_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bookmarked_businesses.Bookmarked_Businesses",
	HandlerType: (*Bookmarked_BusinessesServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "bookmarked_businesses.proto",
}
