// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package projectv1

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

// ProjectAPIClient is the client API for ProjectAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProjectAPIClient interface {
	GetProjects(ctx context.Context, in *GetProjectsRequest, opts ...grpc.CallOption) (*GetProjectsResponse, error)
}

type projectAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectAPIClient(cc grpc.ClientConnInterface) ProjectAPIClient {
	return &projectAPIClient{cc}
}

func (c *projectAPIClient) GetProjects(ctx context.Context, in *GetProjectsRequest, opts ...grpc.CallOption) (*GetProjectsResponse, error) {
	out := new(GetProjectsResponse)
	err := c.cc.Invoke(ctx, "/clutch.project.v1.ProjectAPI/GetProjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectAPIServer is the server API for ProjectAPI service.
// All implementations should embed UnimplementedProjectAPIServer
// for forward compatibility
type ProjectAPIServer interface {
	GetProjects(context.Context, *GetProjectsRequest) (*GetProjectsResponse, error)
}

// UnimplementedProjectAPIServer should be embedded to have forward compatible implementations.
type UnimplementedProjectAPIServer struct {
}

func (UnimplementedProjectAPIServer) GetProjects(context.Context, *GetProjectsRequest) (*GetProjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProjects not implemented")
}

// UnsafeProjectAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProjectAPIServer will
// result in compilation errors.
type UnsafeProjectAPIServer interface {
	mustEmbedUnimplementedProjectAPIServer()
}

func RegisterProjectAPIServer(s grpc.ServiceRegistrar, srv ProjectAPIServer) {
	s.RegisterService(&ProjectAPI_ServiceDesc, srv)
}

func _ProjectAPI_GetProjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectAPIServer).GetProjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.project.v1.ProjectAPI/GetProjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectAPIServer).GetProjects(ctx, req.(*GetProjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProjectAPI_ServiceDesc is the grpc.ServiceDesc for ProjectAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProjectAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "clutch.project.v1.ProjectAPI",
	HandlerType: (*ProjectAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProjects",
			Handler:    _ProjectAPI_GetProjects_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "project/v1/project.proto",
}