// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/watcher/v1/watch.proto

package watcher

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
	any "github.com/catper/protobuf/ptypes/any"
	_ "github.com/catper/protobuf/ptypes/empty"
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

// A reported value can be in one of the following states:
type Change_State int32

const (
	// The element exists and its full value is included in data.
	Change_EXISTS Change_State = 0
	// The element does not exist.
	Change_DOES_NOT_EXIST Change_State = 1
	// Element may or may not exist. Used only for initial state delivery when
	// the client is not interested in fetching the initial state. See the
	// "Initial State" section above.
	Change_INITIAL_STATE_SKIPPED Change_State = 2
	// The element may exist, but some error has occurred. More information is
	// available in the data field - the value is a serialized Status
	// proto (from [google.rpc.Status][])
	Change_ERROR Change_State = 3
)

var Change_State_name = map[int32]string{
	0: "EXISTS",
	1: "DOES_NOT_EXIST",
	2: "INITIAL_STATE_SKIPPED",
	3: "ERROR",
}

var Change_State_value = map[string]int32{
	"EXISTS":                0,
	"DOES_NOT_EXIST":        1,
	"INITIAL_STATE_SKIPPED": 2,
	"ERROR":                 3,
}

func (x Change_State) String() string {
	return proto.EnumName(Change_State_name, int32(x))
}

func (Change_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fbde036e07af626b, []int{2, 0}
}

// The message used by the client to register interest in an entity.
type Request struct {
	// The `target` value **must** be a valid URL path pointing to an entity
	// to watch. Note that the service name **must** be
	// removed from the target field (e.g., the target field must say
	// "/foo/bar", not "myservice.googleapis.com/foo/bar"). A client is
	// also allowed to pass system-specific parameters in the URL that
	// are only obeyed by some implementations. Some parameters will be
	// implementation-specific. However, some have predefined meaning
	// and are listed here:
	//
	//  * recursive = true|false [default=false]
	//    If set to true, indicates that the client wants to watch all elements
	//    of entities in the subtree rooted at the entity's name in `target`. For
	//    descendants that are not the immediate children of the target, the
	//    `Change.element` will contain slashes.
	//
	//    Note that some namespaces and entities will not support recursive
	//    watching. When watching such an entity, a client must not set recursive
	//    to true. Otherwise, it will receive an `UNIMPLEMENTED` error.
	//
	// Normal URL encoding must be used inside `target`.  For example, if a query
	// parameter name or value, or the non-query parameter portion of `target`
	// contains a special character, it must be %-encoded.  We recommend that
	// clients and servers use their runtime's URL library to produce and consume
	// target values.
	Target string `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	// The `resume_marker` specifies how much of the existing underlying state is
	// delivered to the client when the watch request is received by the
	// system. The client can set this marker in one of the following ways to get
	// different semantics:
	//
	// *   Parameter is not specified or has the value "".
	//     Semantics: Fetch initial state.
	//     The client wants the entity's initial state to be delivered. See the
	//     description in "Initial State".
	//
	// *   Parameter is set to the string "now" (UTF-8 encoding).
	//     Semantics: Fetch new changes only.
	//     The client just wants to get the changes received by the system after
	//     the watch point. The system may deliver changes from before the watch
	//     point as well.
	//
	// *   Parameter is set to a value received in an earlier
	//     `Change.resume_marker` field while watching the same entity.
	//     Semantics: Resume from a specific point.
	//     The client wants to receive the changes from a specific point; this
	//     value must correspond to a value received in the `Change.resume_marker`
	//     field. The system may deliver changes from before the `resume_marker`
	//     as well. If the system cannot resume the stream from this point (e.g.,
	//     if it is too far behind in the stream), it can raise the
	//     `FAILED_PRECONDITION` error.
	//
	// An implementation MUST support an unspecified parameter and the
	// empty string "" marker (initial state fetching) and the "now" marker.
	// It need not support resuming from a specific point.
	ResumeMarker         []byte   `protobuf:"bytes,2,opt,name=resume_marker,json=resumeMarker,proto3" json:"resume_marker,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbde036e07af626b, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *Request) GetResumeMarker() []byte {
	if m != nil {
		return m.ResumeMarker
	}
	return nil
}

// A batch of Change messages.
type ChangeBatch struct {
	// A list of Change messages.
	Changes              []*Change `protobuf:"bytes,1,rep,name=changes,proto3" json:"changes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ChangeBatch) Reset()         { *m = ChangeBatch{} }
func (m *ChangeBatch) String() string { return proto.CompactTextString(m) }
func (*ChangeBatch) ProtoMessage()    {}
func (*ChangeBatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbde036e07af626b, []int{1}
}

func (m *ChangeBatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeBatch.Unmarshal(m, b)
}
func (m *ChangeBatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeBatch.Marshal(b, m, deterministic)
}
func (m *ChangeBatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeBatch.Merge(m, src)
}
func (m *ChangeBatch) XXX_Size() int {
	return xxx_messageInfo_ChangeBatch.Size(m)
}
func (m *ChangeBatch) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeBatch.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeBatch proto.InternalMessageInfo

func (m *ChangeBatch) GetChanges() []*Change {
	if m != nil {
		return m.Changes
	}
	return nil
}

// A Change indicates the most recent state of an element.
type Change struct {
	// Name of the element, interpreted relative to the entity's actual
	// name. "" refers to the entity itself. The element name is a valid
	// UTF-8 string.
	Element string `protobuf:"bytes,1,opt,name=element,proto3" json:"element,omitempty"`
	// The state of the `element`.
	State Change_State `protobuf:"varint,2,opt,name=state,proto3,enum=google.watcher.v1.Change_State" json:"state,omitempty"`
	// The actual change data. This field is present only when `state() == EXISTS`
	// or `state() == ERROR`. Please see
	// [google.protobuf.Any][google.protobuf.Any] about how to use the Any type.
	Data *any.Any `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	// If present, provides a compact representation of all the messages that have
	// been received by the caller for the given entity, e.g., it could be a
	// sequence number or a multi-part timestamp/version vector. This marker can
	// be provided in the Request message, allowing the caller to resume the
	// stream watching at a specific point without fetching the initial state.
	ResumeMarker []byte `protobuf:"bytes,4,opt,name=resume_marker,json=resumeMarker,proto3" json:"resume_marker,omitempty"`
	// If true, this Change is followed by more Changes that are in the same group
	// as this Change.
	Continued            bool     `protobuf:"varint,5,opt,name=continued,proto3" json:"continued,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Change) Reset()         { *m = Change{} }
func (m *Change) String() string { return proto.CompactTextString(m) }
func (*Change) ProtoMessage()    {}
func (*Change) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbde036e07af626b, []int{2}
}

func (m *Change) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Change.Unmarshal(m, b)
}
func (m *Change) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Change.Marshal(b, m, deterministic)
}
func (m *Change) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Change.Merge(m, src)
}
func (m *Change) XXX_Size() int {
	return xxx_messageInfo_Change.Size(m)
}
func (m *Change) XXX_DiscardUnknown() {
	xxx_messageInfo_Change.DiscardUnknown(m)
}

var xxx_messageInfo_Change proto.InternalMessageInfo

func (m *Change) GetElement() string {
	if m != nil {
		return m.Element
	}
	return ""
}

func (m *Change) GetState() Change_State {
	if m != nil {
		return m.State
	}
	return Change_EXISTS
}

func (m *Change) GetData() *any.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Change) GetResumeMarker() []byte {
	if m != nil {
		return m.ResumeMarker
	}
	return nil
}

func (m *Change) GetContinued() bool {
	if m != nil {
		return m.Continued
	}
	return false
}

func init() {
	proto.RegisterEnum("google.watcher.v1.Change_State", Change_State_name, Change_State_value)
	proto.RegisterType((*Request)(nil), "google.watcher.v1.Request")
	proto.RegisterType((*ChangeBatch)(nil), "google.watcher.v1.ChangeBatch")
	proto.RegisterType((*Change)(nil), "google.watcher.v1.Change")
}

func init() {
	proto.RegisterFile("google/watcher/v1/watch.proto", fileDescriptor_fbde036e07af626b)
}

var fileDescriptor_fbde036e07af626b = []byte{
	// 449 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0xdd, 0x92, 0xd2, 0xd3, 0x31, 0x75, 0x16, 0x43, 0x69, 0x19, 0x10, 0x85, 0x9b, 0x5c,
	0x25, 0xac, 0x13, 0x12, 0x12, 0x57, 0x2d, 0x0b, 0x52, 0x04, 0x5b, 0x2b, 0x27, 0x12, 0x13, 0x37,
	0x91, 0x97, 0x99, 0xac, 0xa2, 0xb1, 0x4b, 0xe2, 0x0e, 0xed, 0x96, 0x57, 0x40, 0x3c, 0x19, 0xaf,
	0xc0, 0x83, 0xa0, 0xda, 0x0e, 0x20, 0xb2, 0xde, 0x9d, 0xf3, 0xfd, 0xd8, 0xe7, 0x3b, 0x36, 0x3c,
	0x29, 0x84, 0x28, 0x96, 0x2c, 0xfc, 0x4a, 0x65, 0x7e, 0xcd, 0xaa, 0xf0, 0xe6, 0x58, 0x97, 0xc1,
	0xaa, 0x12, 0x52, 0xe0, 0x03, 0x4d, 0x07, 0x86, 0x0e, 0x6e, 0x8e, 0x47, 0x47, 0xc6, 0x41, 0x57,
	0x8b, 0x90, 0x72, 0x2e, 0x24, 0x95, 0x0b, 0xc1, 0x6b, 0x6d, 0x18, 0x0d, 0x0d, 0xab, 0xba, 0xcb,
	0xf5, 0xa7, 0x90, 0xf2, 0x5b, 0x43, 0x3d, 0xfe, 0x9f, 0x62, 0xe5, 0x4a, 0x1a, 0xd2, 0x7b, 0x0b,
	0x5d, 0xc2, 0xbe, 0xac, 0x59, 0x2d, 0xf1, 0x23, 0xb0, 0x25, 0xad, 0x0a, 0x26, 0x1d, 0xe4, 0x22,
	0xbf, 0x47, 0x4c, 0x87, 0x9f, 0xc3, 0x83, 0x8a, 0xd5, 0xeb, 0x92, 0x65, 0x25, 0xad, 0x3e, 0xb3,
	0xca, 0xe9, 0xb8, 0xc8, 0xdf, 0x23, 0x7b, 0x1a, 0x3c, 0x53, 0x98, 0x37, 0x85, 0xfe, 0x9b, 0x6b,
	0xca, 0x0b, 0x36, 0xdd, 0x4c, 0x8c, 0x4f, 0xa0, 0x9b, 0xab, 0xb6, 0x76, 0x90, 0xbb, 0xe3, 0xf7,
	0xc7, 0xc3, 0xa0, 0x95, 0x28, 0xd0, 0x06, 0xd2, 0x28, 0xbd, 0x1f, 0x1d, 0xb0, 0x35, 0x86, 0x1d,
	0xe8, 0xb2, 0x25, 0x2b, 0x19, 0x6f, 0x86, 0x69, 0x5a, 0xfc, 0x12, 0xac, 0x5a, 0x52, 0xc9, 0xd4,
	0x14, 0xfb, 0xe3, 0x67, 0x5b, 0xcf, 0x0d, 0x92, 0x8d, 0x8c, 0x68, 0x35, 0xf6, 0x61, 0xf7, 0x8a,
	0x4a, 0xea, 0xd8, 0x2e, 0xf2, 0xfb, 0xe3, 0x87, 0x8d, 0xab, 0xd9, 0x49, 0x30, 0xe1, 0xb7, 0x44,
	0x29, 0xda, 0x71, 0x77, 0xdb, 0x71, 0xf1, 0x11, 0xf4, 0x72, 0xc1, 0xe5, 0x82, 0xaf, 0xd9, 0x95,
	0x63, 0xb9, 0xc8, 0xbf, 0x4f, 0xfe, 0x02, 0xde, 0x19, 0x58, 0xea, 0x72, 0x0c, 0x60, 0x47, 0x17,
	0x71, 0x92, 0x26, 0x83, 0x7b, 0x18, 0xc3, 0xfe, 0xe9, 0x2c, 0x4a, 0xb2, 0xf3, 0x59, 0x9a, 0x29,
	0x70, 0x80, 0xf0, 0x10, 0x0e, 0xe3, 0xf3, 0x38, 0x8d, 0x27, 0xef, 0xb3, 0x24, 0x9d, 0xa4, 0x51,
	0x96, 0xbc, 0x8b, 0xe7, 0xf3, 0xe8, 0x74, 0xd0, 0xc1, 0x3d, 0xb0, 0x22, 0x42, 0x66, 0x64, 0xb0,
	0x33, 0xce, 0xa1, 0xfb, 0x41, 0xa7, 0xc3, 0x17, 0x60, 0xa9, 0x12, 0x8f, 0xee, 0xc8, 0x6d, 0x1e,
	0x72, 0xf4, 0x74, 0xeb, 0x4e, 0xd4, 0xe3, 0x78, 0x07, 0xdf, 0x7e, 0xfe, 0xfa, 0xde, 0xe9, 0xe3,
	0xde, 0x9f, 0x5f, 0xf7, 0x02, 0x4d, 0x33, 0x38, 0xcc, 0x45, 0xd9, 0x76, 0x4e, 0x41, 0x5d, 0x38,
	0xdf, 0x2c, 0x6a, 0x8e, 0x3e, 0xbe, 0x32, 0x82, 0x42, 0x2c, 0x29, 0x2f, 0x02, 0x51, 0x15, 0x61,
	0xc1, 0xb8, 0x5a, 0x63, 0xa8, 0x29, 0xba, 0x5a, 0xd4, 0xff, 0x7c, 0xeb, 0xd7, 0xa6, 0xbc, 0xb4,
	0x95, 0xe8, 0xe4, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x48, 0x0a, 0xba, 0x6c, 0xfa, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WatcherClient is the client API for Watcher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WatcherClient interface {
	// Start a streaming RPC to get watch information from the server.
	Watch(ctx context.Context, in *Request, opts ...grpc.CallOption) (Watcher_WatchClient, error)
}

type watcherClient struct {
	cc grpc.ClientConnInterface
}

func NewWatcherClient(cc grpc.ClientConnInterface) WatcherClient {
	return &watcherClient{cc}
}

func (c *watcherClient) Watch(ctx context.Context, in *Request, opts ...grpc.CallOption) (Watcher_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Watcher_serviceDesc.Streams[0], "/google.watcher.v1.Watcher/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &watcherWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Watcher_WatchClient interface {
	Recv() (*ChangeBatch, error)
	grpc.ClientStream
}

type watcherWatchClient struct {
	grpc.ClientStream
}

func (x *watcherWatchClient) Recv() (*ChangeBatch, error) {
	m := new(ChangeBatch)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WatcherServer is the server API for Watcher service.
type WatcherServer interface {
	// Start a streaming RPC to get watch information from the server.
	Watch(*Request, Watcher_WatchServer) error
}

// UnimplementedWatcherServer can be embedded to have forward compatible implementations.
type UnimplementedWatcherServer struct {
}

func (*UnimplementedWatcherServer) Watch(req *Request, srv Watcher_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}

func RegisterWatcherServer(s *grpc.Server, srv WatcherServer) {
	s.RegisterService(&_Watcher_serviceDesc, srv)
}

func _Watcher_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WatcherServer).Watch(m, &watcherWatchServer{stream})
}

type Watcher_WatchServer interface {
	Send(*ChangeBatch) error
	grpc.ServerStream
}

type watcherWatchServer struct {
	grpc.ServerStream
}

func (x *watcherWatchServer) Send(m *ChangeBatch) error {
	return x.ServerStream.SendMsg(m)
}

var _Watcher_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.watcher.v1.Watcher",
	HandlerType: (*WatcherServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _Watcher_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "google/watcher/v1/watch.proto",
}
