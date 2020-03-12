// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/devtools/containeranalysis/v1/containeranalysis.proto

package containeranalysis

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
	_ "github.com/catper/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	v1 "google.golang.org/genproto/googleapis/iam/v1"
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
	proto.RegisterFile("google/devtools/containeranalysis/v1/containeranalysis.proto", fileDescriptor_e74edb4ed33b4f81)
}

var fileDescriptor_e74edb4ed33b4f81 = []byte{
	// 428 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0xc1, 0x8a, 0xd5, 0x30,
	0x14, 0x86, 0x69, 0x45, 0x17, 0xc5, 0x8d, 0x05, 0x37, 0x45, 0x1c, 0x28, 0xb3, 0xd0, 0x2e, 0x12,
	0xaa, 0x6e, 0x8c, 0x8a, 0x8c, 0x57, 0x28, 0xee, 0x06, 0x1d, 0x67, 0x31, 0x1b, 0xc9, 0x64, 0x8e,
	0x25, 0xd2, 0xe4, 0xd4, 0x24, 0x2d, 0x8c, 0x22, 0x88, 0x0b, 0x5f, 0xc0, 0x37, 0x70, 0xe9, 0x4b,
	0xf8, 0x10, 0xfa, 0x08, 0xae, 0xc4, 0x87, 0x90, 0xa6, 0xa9, 0xdc, 0x7b, 0x7b, 0xbd, 0x5c, 0x04,
	0x97, 0x3d, 0xff, 0x7f, 0xfe, 0xfe, 0x1f, 0xe4, 0x24, 0xf7, 0x6b, 0xc4, 0xba, 0x01, 0x7a, 0x06,
	0xbd, 0x43, 0x6c, 0x2c, 0x15, 0xa8, 0x1d, 0x97, 0x1a, 0x0c, 0xd7, 0xbc, 0x39, 0xb7, 0xd2, 0xd2,
	0xbe, 0x9c, 0x0f, 0x49, 0x6b, 0xd0, 0x61, 0xba, 0x3f, 0x6e, 0x93, 0x69, 0x9b, 0xcc, 0x8d, 0x7d,
	0x99, 0x5d, 0x0b, 0xff, 0xe0, 0xad, 0xa4, 0x5c, 0x6b, 0x74, 0xdc, 0x49, 0xd4, 0x21, 0x23, 0xbb,
	0x1e, 0x54, 0xc9, 0xd5, 0xf0, 0x2b, 0xc9, 0xd5, 0x8b, 0x16, 0x1b, 0x29, 0xce, 0x83, 0x9e, 0xad,
	0xea, 0x2b, 0xda, 0x5e, 0xd0, 0xfc, 0xd7, 0x69, 0xf7, 0x92, 0x3a, 0xa9, 0xc0, 0x3a, 0xae, 0xda,
	0xd1, 0x70, 0xeb, 0xe3, 0xc5, 0xe4, 0xca, 0x62, 0xea, 0x74, 0x10, 0x3a, 0xa5, 0x5f, 0xa3, 0xe4,
	0xf2, 0x33, 0x70, 0x4f, 0xb8, 0x3a, 0xf4, 0x69, 0x69, 0x4e, 0x02, 0x88, 0xe4, 0x8a, 0xf4, 0x25,
	0x59, 0x16, 0x9f, 0xc2, 0xeb, 0x0e, 0xac, 0xcb, 0xae, 0xae, 0x79, 0x46, 0x35, 0x77, 0x1f, 0xbe,
	0xfd, 0xf8, 0x14, 0xeb, 0x9c, 0x0c, 0xe5, 0xde, 0x1a, 0xb0, 0xd8, 0x19, 0x01, 0x0f, 0x5a, 0x83,
	0xaf, 0x40, 0x38, 0x4b, 0x0b, 0xaa, 0xd1, 0x81, 0xa5, 0xc5, 0x3b, 0x66, 0x97, 0x52, 0x59, 0x54,
	0x9c, 0xdc, 0xcd, 0xef, 0xfc, 0x75, 0x09, 0x85, 0xe8, 0x8c, 0x01, 0x2d, 0x36, 0xae, 0x7a, 0x82,
	0x6a, 0x1b, 0x41, 0xf5, 0x5f, 0x08, 0xea, 0x7f, 0x27, 0x58, 0x5b, 0x4d, 0x7f, 0x45, 0x49, 0x7a,
	0x04, 0xd6, 0x0f, 0xc1, 0x28, 0x69, 0xed, 0xf0, 0x26, 0xd2, 0x1b, 0x6b, 0x1d, 0xe7, 0x96, 0x89,
	0xe6, 0xe6, 0x0e, 0x4e, 0xdb, 0xa2, 0xb6, 0x90, 0xbf, 0x8f, 0x3c, 0xe2, 0x9b, 0x2d, 0x6d, 0xff,
	0x20, 0xba, 0x59, 0xcc, 0x00, 0xfa, 0x30, 0x67, 0xbb, 0x82, 0x6e, 0x0c, 0x78, 0xf4, 0x3d, 0x4a,
	0xf6, 0x04, 0xaa, 0xa9, 0xf3, 0xa6, 0x3b, 0x39, 0x8c, 0x4e, 0x9e, 0x07, 0xb9, 0xc6, 0x86, 0xeb,
	0x9a, 0xa0, 0xa9, 0x69, 0x0d, 0xda, 0x3f, 0x65, 0x3a, 0x4a, 0xbc, 0x95, 0x76, 0xfb, 0xb1, 0xde,
	0x9b, 0x0d, 0x3f, 0xc7, 0x17, 0xaa, 0xc5, 0xc1, 0x97, 0xb8, 0xa8, 0xc6, 0xf4, 0x45, 0x83, 0xdd,
	0x19, 0x79, 0x0c, 0xfd, 0x91, 0xbf, 0xd9, 0xd9, 0x7d, 0x90, 0xe3, 0xf2, 0x67, 0xbc, 0x3f, 0x9a,
	0x19, 0xf3, 0x6e, 0xc6, 0x66, 0x2e, 0xc6, 0x8e, 0xcb, 0xd3, 0x4b, 0xbe, 0xdb, 0xed, 0xdf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xb6, 0xa4, 0x2b, 0x8c, 0x47, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ContainerAnalysisClient is the client API for ContainerAnalysis service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ContainerAnalysisClient interface {
	// Sets the access control policy on the specified note or occurrence.
	// Requires `containeranalysis.notes.setIamPolicy` or
	// `containeranalysis.occurrences.setIamPolicy` permission if the resource is
	// a note or an occurrence, respectively.
	//
	// The resource takes the format `projects/[PROJECT_ID]/notes/[NOTE_ID]` for
	// notes and `projects/[PROJECT_ID]/occurrences/[OCCURRENCE_ID]` for
	// occurrences.
	SetIamPolicy(ctx context.Context, in *v1.SetIamPolicyRequest, opts ...grpc.CallOption) (*v1.Policy, error)
	// Gets the access control policy for a note or an occurrence resource.
	// Requires `containeranalysis.notes.setIamPolicy` or
	// `containeranalysis.occurrences.setIamPolicy` permission if the resource is
	// a note or occurrence, respectively.
	//
	// The resource takes the format `projects/[PROJECT_ID]/notes/[NOTE_ID]` for
	// notes and `projects/[PROJECT_ID]/occurrences/[OCCURRENCE_ID]` for
	// occurrences.
	GetIamPolicy(ctx context.Context, in *v1.GetIamPolicyRequest, opts ...grpc.CallOption) (*v1.Policy, error)
	// Returns the permissions that a caller has on the specified note or
	// occurrence. Requires list permission on the project (for example,
	// `containeranalysis.notes.list`).
	//
	// The resource takes the format `projects/[PROJECT_ID]/notes/[NOTE_ID]` for
	// notes and `projects/[PROJECT_ID]/occurrences/[OCCURRENCE_ID]` for
	// occurrences.
	TestIamPermissions(ctx context.Context, in *v1.TestIamPermissionsRequest, opts ...grpc.CallOption) (*v1.TestIamPermissionsResponse, error)
}

type containerAnalysisClient struct {
	cc *grpc.ClientConn
}

func NewContainerAnalysisClient(cc *grpc.ClientConn) ContainerAnalysisClient {
	return &containerAnalysisClient{cc}
}

func (c *containerAnalysisClient) SetIamPolicy(ctx context.Context, in *v1.SetIamPolicyRequest, opts ...grpc.CallOption) (*v1.Policy, error) {
	out := new(v1.Policy)
	err := c.cc.Invoke(ctx, "/google.devtools.containeranalysis.v1.ContainerAnalysis/SetIamPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerAnalysisClient) GetIamPolicy(ctx context.Context, in *v1.GetIamPolicyRequest, opts ...grpc.CallOption) (*v1.Policy, error) {
	out := new(v1.Policy)
	err := c.cc.Invoke(ctx, "/google.devtools.containeranalysis.v1.ContainerAnalysis/GetIamPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerAnalysisClient) TestIamPermissions(ctx context.Context, in *v1.TestIamPermissionsRequest, opts ...grpc.CallOption) (*v1.TestIamPermissionsResponse, error) {
	out := new(v1.TestIamPermissionsResponse)
	err := c.cc.Invoke(ctx, "/google.devtools.containeranalysis.v1.ContainerAnalysis/TestIamPermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContainerAnalysisServer is the server API for ContainerAnalysis service.
type ContainerAnalysisServer interface {
	// Sets the access control policy on the specified note or occurrence.
	// Requires `containeranalysis.notes.setIamPolicy` or
	// `containeranalysis.occurrences.setIamPolicy` permission if the resource is
	// a note or an occurrence, respectively.
	//
	// The resource takes the format `projects/[PROJECT_ID]/notes/[NOTE_ID]` for
	// notes and `projects/[PROJECT_ID]/occurrences/[OCCURRENCE_ID]` for
	// occurrences.
	SetIamPolicy(context.Context, *v1.SetIamPolicyRequest) (*v1.Policy, error)
	// Gets the access control policy for a note or an occurrence resource.
	// Requires `containeranalysis.notes.setIamPolicy` or
	// `containeranalysis.occurrences.setIamPolicy` permission if the resource is
	// a note or occurrence, respectively.
	//
	// The resource takes the format `projects/[PROJECT_ID]/notes/[NOTE_ID]` for
	// notes and `projects/[PROJECT_ID]/occurrences/[OCCURRENCE_ID]` for
	// occurrences.
	GetIamPolicy(context.Context, *v1.GetIamPolicyRequest) (*v1.Policy, error)
	// Returns the permissions that a caller has on the specified note or
	// occurrence. Requires list permission on the project (for example,
	// `containeranalysis.notes.list`).
	//
	// The resource takes the format `projects/[PROJECT_ID]/notes/[NOTE_ID]` for
	// notes and `projects/[PROJECT_ID]/occurrences/[OCCURRENCE_ID]` for
	// occurrences.
	TestIamPermissions(context.Context, *v1.TestIamPermissionsRequest) (*v1.TestIamPermissionsResponse, error)
}

// UnimplementedContainerAnalysisServer can be embedded to have forward compatible implementations.
type UnimplementedContainerAnalysisServer struct {
}

func (*UnimplementedContainerAnalysisServer) SetIamPolicy(ctx context.Context, req *v1.SetIamPolicyRequest) (*v1.Policy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetIamPolicy not implemented")
}
func (*UnimplementedContainerAnalysisServer) GetIamPolicy(ctx context.Context, req *v1.GetIamPolicyRequest) (*v1.Policy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIamPolicy not implemented")
}
func (*UnimplementedContainerAnalysisServer) TestIamPermissions(ctx context.Context, req *v1.TestIamPermissionsRequest) (*v1.TestIamPermissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestIamPermissions not implemented")
}

func RegisterContainerAnalysisServer(s *grpc.Server, srv ContainerAnalysisServer) {
	s.RegisterService(&_ContainerAnalysis_serviceDesc, srv)
}

func _ContainerAnalysis_SetIamPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.SetIamPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerAnalysisServer).SetIamPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.containeranalysis.v1.ContainerAnalysis/SetIamPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerAnalysisServer).SetIamPolicy(ctx, req.(*v1.SetIamPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContainerAnalysis_GetIamPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.GetIamPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerAnalysisServer).GetIamPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.containeranalysis.v1.ContainerAnalysis/GetIamPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerAnalysisServer).GetIamPolicy(ctx, req.(*v1.GetIamPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContainerAnalysis_TestIamPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.TestIamPermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerAnalysisServer).TestIamPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.containeranalysis.v1.ContainerAnalysis/TestIamPermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerAnalysisServer).TestIamPermissions(ctx, req.(*v1.TestIamPermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ContainerAnalysis_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.devtools.containeranalysis.v1.ContainerAnalysis",
	HandlerType: (*ContainerAnalysisServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetIamPolicy",
			Handler:    _ContainerAnalysis_SetIamPolicy_Handler,
		},
		{
			MethodName: "GetIamPolicy",
			Handler:    _ContainerAnalysis_GetIamPolicy_Handler,
		},
		{
			MethodName: "TestIamPermissions",
			Handler:    _ContainerAnalysis_TestIamPermissions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/devtools/containeranalysis/v1/containeranalysis.proto",
}
