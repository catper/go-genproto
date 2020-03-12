// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/api/expr/v1alpha1/cel_service.proto

package expr

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("google/api/expr/v1alpha1/cel_service.proto", fileDescriptor_f35b2125e64b6d66)
}

var fileDescriptor_f35b2125e64b6d66 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x31, 0x4b, 0x03, 0x31,
	0x14, 0xc7, 0xad, 0x5a, 0x87, 0x88, 0x0a, 0x01, 0x51, 0x0e, 0x27, 0xc1, 0x0a, 0x0e, 0x09, 0xad,
	0xa3, 0x2e, 0xd7, 0xe2, 0x7e, 0xe8, 0xd6, 0xa5, 0xc4, 0xf8, 0xcc, 0x1d, 0xe6, 0xf2, 0x62, 0x72,
	0x3d, 0xfc, 0x8c, 0x8e, 0x8e, 0xf7, 0x49, 0x1c, 0x25, 0x77, 0x29, 0x55, 0xe4, 0xac, 0x63, 0x92,
	0xdf, 0xfb, 0xff, 0xc8, 0xfb, 0x93, 0x2b, 0x85, 0xa8, 0x34, 0x70, 0x61, 0x0b, 0x0e, 0x6f, 0xd6,
	0xf1, 0x7a, 0x2c, 0xb4, 0xcd, 0xc5, 0x98, 0x4b, 0xd0, 0x0b, 0x0f, 0xae, 0x2e, 0x24, 0x30, 0xeb,
	0xb0, 0x42, 0x7a, 0xda, 0xb1, 0x4c, 0xd8, 0x82, 0x05, 0x96, 0xad, 0xd8, 0xe4, 0xe4, 0x5b, 0x8a,
	0xd4, 0x05, 0x98, 0xaa, 0x1b, 0x49, 0x26, 0xfd, 0xf1, 0x68, 0x9e, 0xd1, 0x95, 0xc2, 0x48, 0xf8,
	0xa9, 0x99, 0x7c, 0x6c, 0x13, 0x32, 0x03, 0xfd, 0xd0, 0x5d, 0x52, 0x45, 0x86, 0x99, 0x70, 0x1e,
	0xe8, 0x88, 0xf5, 0xf9, 0x59, 0x0b, 0xdc, 0xc3, 0xeb, 0x12, 0x7c, 0x95, 0x5c, 0x6e, 0xe4, 0xbc,
	0x45, 0xe3, 0xe1, 0xfc, 0xa0, 0x49, 0x49, 0xfb, 0x4b, 0x5c, 0x3a, 0x09, 0x34, 0x27, 0xc3, 0x59,
	0x0e, 0xf2, 0xe5, 0x2f, 0x51, 0x0b, 0xfc, 0x43, 0x14, 0xb9, 0x28, 0x3a, 0x6c, 0xd2, 0x7d, 0x1b,
	0xd4, 0x4f, 0x8b, 0x80, 0xd1, 0x39, 0xd9, 0xbd, 0xab, 0x85, 0xa6, 0x17, 0xfd, 0x01, 0xe1, 0x7d,
	0xe5, 0x19, 0x6d, 0xc2, 0xa2, 0x66, 0xa7, 0x49, 0xb7, 0x92, 0xe3, 0xf7, 0x94, 0x4a, 0xd0, 0x71,
	0x44, 0xd8, 0xc2, 0x33, 0x89, 0xe5, 0xd4, 0x91, 0x33, 0x89, 0x65, 0x6f, 0xd0, 0xf4, 0x68, 0xbd,
	0xf1, 0x2c, 0xb4, 0x90, 0x0d, 0xe6, 0xb7, 0x11, 0x56, 0xa8, 0x85, 0x51, 0x0c, 0x9d, 0xe2, 0x0a,
	0x4c, 0xdb, 0x11, 0x5f, 0xa7, 0xff, 0xae, 0xf6, 0x26, 0x9c, 0x3e, 0x07, 0x83, 0xc7, 0xbd, 0x96,
	0xbd, 0xfe, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x25, 0x59, 0x35, 0x64, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CelServiceClient is the client API for CelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CelServiceClient interface {
	// Transforms CEL source text into a parsed representation.
	Parse(ctx context.Context, in *ParseRequest, opts ...grpc.CallOption) (*ParseResponse, error)
	// Runs static checks on a parsed CEL representation and return
	// an annotated representation, or a set of issues.
	Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error)
	// Evaluates a parsed or annotation CEL representation given
	// values of external bindings.
	Eval(ctx context.Context, in *EvalRequest, opts ...grpc.CallOption) (*EvalResponse, error)
}

type celServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCelServiceClient(cc grpc.ClientConnInterface) CelServiceClient {
	return &celServiceClient{cc}
}

func (c *celServiceClient) Parse(ctx context.Context, in *ParseRequest, opts ...grpc.CallOption) (*ParseResponse, error) {
	out := new(ParseResponse)
	err := c.cc.Invoke(ctx, "/google.api.expr.v1alpha1.CelService/Parse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *celServiceClient) Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, "/google.api.expr.v1alpha1.CelService/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *celServiceClient) Eval(ctx context.Context, in *EvalRequest, opts ...grpc.CallOption) (*EvalResponse, error) {
	out := new(EvalResponse)
	err := c.cc.Invoke(ctx, "/google.api.expr.v1alpha1.CelService/Eval", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CelServiceServer is the server API for CelService service.
type CelServiceServer interface {
	// Transforms CEL source text into a parsed representation.
	Parse(context.Context, *ParseRequest) (*ParseResponse, error)
	// Runs static checks on a parsed CEL representation and return
	// an annotated representation, or a set of issues.
	Check(context.Context, *CheckRequest) (*CheckResponse, error)
	// Evaluates a parsed or annotation CEL representation given
	// values of external bindings.
	Eval(context.Context, *EvalRequest) (*EvalResponse, error)
}

// UnimplementedCelServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCelServiceServer struct {
}

func (*UnimplementedCelServiceServer) Parse(ctx context.Context, req *ParseRequest) (*ParseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Parse not implemented")
}
func (*UnimplementedCelServiceServer) Check(ctx context.Context, req *CheckRequest) (*CheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (*UnimplementedCelServiceServer) Eval(ctx context.Context, req *EvalRequest) (*EvalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Eval not implemented")
}

func RegisterCelServiceServer(s *grpc.Server, srv CelServiceServer) {
	s.RegisterService(&_CelService_serviceDesc, srv)
}

func _CelService_Parse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CelServiceServer).Parse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.api.expr.v1alpha1.CelService/Parse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CelServiceServer).Parse(ctx, req.(*ParseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CelService_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CelServiceServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.api.expr.v1alpha1.CelService/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CelServiceServer).Check(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CelService_Eval_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EvalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CelServiceServer).Eval(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.api.expr.v1alpha1.CelService/Eval",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CelServiceServer).Eval(ctx, req.(*EvalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CelService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.api.expr.v1alpha1.CelService",
	HandlerType: (*CelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Parse",
			Handler:    _CelService_Parse_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _CelService_Check_Handler,
		},
		{
			MethodName: "Eval",
			Handler:    _CelService_Eval_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/api/expr/v1alpha1/cel_service.proto",
}
