// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package auditv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// AuditAPIClient is the client API for AuditAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuditAPIClient interface {
	GetEvents(ctx context.Context, in *GetEventsRequest, opts ...grpc.CallOption) (*GetEventsResponse, error)
}

type auditAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewAuditAPIClient(cc grpc.ClientConnInterface) AuditAPIClient {
	return &auditAPIClient{cc}
}

func (c *auditAPIClient) GetEvents(ctx context.Context, in *GetEventsRequest, opts ...grpc.CallOption) (*GetEventsResponse, error) {
	out := new(GetEventsResponse)
	err := c.cc.Invoke(ctx, "/clutch.audit.v1.AuditAPI/GetEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuditAPIServer is the server API for AuditAPI service.
// All implementations should embed UnimplementedAuditAPIServer
// for forward compatibility
type AuditAPIServer interface {
	GetEvents(context.Context, *GetEventsRequest) (*GetEventsResponse, error)
}

// UnimplementedAuditAPIServer should be embedded to have forward compatible implementations.
type UnimplementedAuditAPIServer struct {
}

func (UnimplementedAuditAPIServer) GetEvents(context.Context, *GetEventsRequest) (*GetEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvents not implemented")
}

// UnsafeAuditAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuditAPIServer will
// result in compilation errors.
type UnsafeAuditAPIServer interface {
	mustEmbedUnimplementedAuditAPIServer()
}

func RegisterAuditAPIServer(s *grpc.Server, srv AuditAPIServer) {
	s.RegisterService(&_AuditAPI_serviceDesc, srv)
}

func _AuditAPI_GetEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuditAPIServer).GetEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.audit.v1.AuditAPI/GetEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuditAPIServer).GetEvents(ctx, req.(*GetEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuditAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "clutch.audit.v1.AuditAPI",
	HandlerType: (*AuditAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEvents",
			Handler:    _AuditAPI_GetEvents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "audit/v1/audit.proto",
}