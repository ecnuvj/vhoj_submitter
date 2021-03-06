// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package submitterpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// SubmitServiceClient is the client API for SubmitService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubmitServiceClient interface {
	SubmitCode(ctx context.Context, in *SubmitCodeRequest, opts ...grpc.CallOption) (*SubmitCodeResponse, error)
	ReSubmitCode(ctx context.Context, in *ReSubmitCodeRequest, opts ...grpc.CallOption) (*ReSubmitCodeResponse, error)
	ListSubmissions(ctx context.Context, in *ListSubmissionsRequest, opts ...grpc.CallOption) (*ListSubmissionsResponse, error)
	GetSubmission(ctx context.Context, in *GetSubmissionRequest, opts ...grpc.CallOption) (*GetSubmissionResponse, error)
}

type submitServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSubmitServiceClient(cc grpc.ClientConnInterface) SubmitServiceClient {
	return &submitServiceClient{cc}
}

func (c *submitServiceClient) SubmitCode(ctx context.Context, in *SubmitCodeRequest, opts ...grpc.CallOption) (*SubmitCodeResponse, error) {
	out := new(SubmitCodeResponse)
	err := c.cc.Invoke(ctx, "/sdk.SubmitService/SubmitCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *submitServiceClient) ReSubmitCode(ctx context.Context, in *ReSubmitCodeRequest, opts ...grpc.CallOption) (*ReSubmitCodeResponse, error) {
	out := new(ReSubmitCodeResponse)
	err := c.cc.Invoke(ctx, "/sdk.SubmitService/ReSubmitCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *submitServiceClient) ListSubmissions(ctx context.Context, in *ListSubmissionsRequest, opts ...grpc.CallOption) (*ListSubmissionsResponse, error) {
	out := new(ListSubmissionsResponse)
	err := c.cc.Invoke(ctx, "/sdk.SubmitService/ListSubmissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *submitServiceClient) GetSubmission(ctx context.Context, in *GetSubmissionRequest, opts ...grpc.CallOption) (*GetSubmissionResponse, error) {
	out := new(GetSubmissionResponse)
	err := c.cc.Invoke(ctx, "/sdk.SubmitService/GetSubmission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubmitServiceServer is the server API for SubmitService service.
// All implementations must embed UnimplementedSubmitServiceServer
// for forward compatibility
type SubmitServiceServer interface {
	SubmitCode(context.Context, *SubmitCodeRequest) (*SubmitCodeResponse, error)
	ReSubmitCode(context.Context, *ReSubmitCodeRequest) (*ReSubmitCodeResponse, error)
	ListSubmissions(context.Context, *ListSubmissionsRequest) (*ListSubmissionsResponse, error)
	GetSubmission(context.Context, *GetSubmissionRequest) (*GetSubmissionResponse, error)
	mustEmbedUnimplementedSubmitServiceServer()
}

// UnimplementedSubmitServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSubmitServiceServer struct {
}

func (UnimplementedSubmitServiceServer) SubmitCode(context.Context, *SubmitCodeRequest) (*SubmitCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitCode not implemented")
}
func (UnimplementedSubmitServiceServer) ReSubmitCode(context.Context, *ReSubmitCodeRequest) (*ReSubmitCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReSubmitCode not implemented")
}
func (UnimplementedSubmitServiceServer) ListSubmissions(context.Context, *ListSubmissionsRequest) (*ListSubmissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSubmissions not implemented")
}
func (UnimplementedSubmitServiceServer) GetSubmission(context.Context, *GetSubmissionRequest) (*GetSubmissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubmission not implemented")
}
func (UnimplementedSubmitServiceServer) mustEmbedUnimplementedSubmitServiceServer() {}

// UnsafeSubmitServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubmitServiceServer will
// result in compilation errors.
type UnsafeSubmitServiceServer interface {
	mustEmbedUnimplementedSubmitServiceServer()
}

func RegisterSubmitServiceServer(s grpc.ServiceRegistrar, srv SubmitServiceServer) {
	s.RegisterService(&_SubmitService_serviceDesc, srv)
}

func _SubmitService_SubmitCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubmitServiceServer).SubmitCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sdk.SubmitService/SubmitCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubmitServiceServer).SubmitCode(ctx, req.(*SubmitCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubmitService_ReSubmitCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReSubmitCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubmitServiceServer).ReSubmitCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sdk.SubmitService/ReSubmitCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubmitServiceServer).ReSubmitCode(ctx, req.(*ReSubmitCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubmitService_ListSubmissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSubmissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubmitServiceServer).ListSubmissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sdk.SubmitService/ListSubmissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubmitServiceServer).ListSubmissions(ctx, req.(*ListSubmissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubmitService_GetSubmission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubmissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubmitServiceServer).GetSubmission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sdk.SubmitService/GetSubmission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubmitServiceServer).GetSubmission(ctx, req.(*GetSubmissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SubmitService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sdk.SubmitService",
	HandlerType: (*SubmitServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitCode",
			Handler:    _SubmitService_SubmitCode_Handler,
		},
		{
			MethodName: "ReSubmitCode",
			Handler:    _SubmitService_ReSubmitCode_Handler,
		},
		{
			MethodName: "ListSubmissions",
			Handler:    _SubmitService_ListSubmissions_Handler,
		},
		{
			MethodName: "GetSubmission",
			Handler:    _SubmitService_GetSubmission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "submitterpb/service.proto",
}
