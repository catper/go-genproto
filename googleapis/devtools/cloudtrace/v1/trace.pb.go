// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/devtools/cloudtrace/v1/trace.proto

package cloudtrace

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
	empty "github.com/catper/protobuf/ptypes/empty"
	timestamp "github.com/catper/protobuf/ptypes/timestamp"
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

// Type of span. Can be used to specify additional relationships between spans
// in addition to a parent/child relationship.
type TraceSpan_SpanKind int32

const (
	// Unspecified.
	TraceSpan_SPAN_KIND_UNSPECIFIED TraceSpan_SpanKind = 0
	// Indicates that the span covers server-side handling of an RPC or other
	// remote network request.
	TraceSpan_RPC_SERVER TraceSpan_SpanKind = 1
	// Indicates that the span covers the client-side wrapper around an RPC or
	// other remote request.
	TraceSpan_RPC_CLIENT TraceSpan_SpanKind = 2
)

var TraceSpan_SpanKind_name = map[int32]string{
	0: "SPAN_KIND_UNSPECIFIED",
	1: "RPC_SERVER",
	2: "RPC_CLIENT",
}

var TraceSpan_SpanKind_value = map[string]int32{
	"SPAN_KIND_UNSPECIFIED": 0,
	"RPC_SERVER":            1,
	"RPC_CLIENT":            2,
}

func (x TraceSpan_SpanKind) String() string {
	return proto.EnumName(TraceSpan_SpanKind_name, int32(x))
}

func (TraceSpan_SpanKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_91f8d9fbc62846cf, []int{2, 0}
}

// Type of data returned for traces in the list.
type ListTracesRequest_ViewType int32

const (
	// Default is `MINIMAL` if unspecified.
	ListTracesRequest_VIEW_TYPE_UNSPECIFIED ListTracesRequest_ViewType = 0
	// Minimal view of the trace record that contains only the project
	// and trace IDs.
	ListTracesRequest_MINIMAL ListTracesRequest_ViewType = 1
	// Root span view of the trace record that returns the root spans along
	// with the minimal trace data.
	ListTracesRequest_ROOTSPAN ListTracesRequest_ViewType = 2
	// Complete view of the trace record that contains the actual trace data.
	// This is equivalent to calling the REST `get` or RPC `GetTrace` method
	// using the ID of each listed trace.
	ListTracesRequest_COMPLETE ListTracesRequest_ViewType = 3
)

var ListTracesRequest_ViewType_name = map[int32]string{
	0: "VIEW_TYPE_UNSPECIFIED",
	1: "MINIMAL",
	2: "ROOTSPAN",
	3: "COMPLETE",
}

var ListTracesRequest_ViewType_value = map[string]int32{
	"VIEW_TYPE_UNSPECIFIED": 0,
	"MINIMAL":               1,
	"ROOTSPAN":              2,
	"COMPLETE":              3,
}

func (x ListTracesRequest_ViewType) String() string {
	return proto.EnumName(ListTracesRequest_ViewType_name, int32(x))
}

func (ListTracesRequest_ViewType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_91f8d9fbc62846cf, []int{3, 0}
}

// A trace describes how long it takes for an application to perform an
// operation. It consists of a set of spans, each of which represent a single
// timed event within the operation.
type Trace struct {
	// Project ID of the Cloud project where the trace data is stored.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// Globally unique identifier for the trace. This identifier is a 128-bit
	// numeric value formatted as a 32-byte hex string. For example,
	// `382d4f4c6b7bb2f4a972559d9085001d`.
	TraceId string `protobuf:"bytes,2,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	// Collection of spans in the trace.
	Spans                []*TraceSpan `protobuf:"bytes,3,rep,name=spans,proto3" json:"spans,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Trace) Reset()         { *m = Trace{} }
func (m *Trace) String() string { return proto.CompactTextString(m) }
func (*Trace) ProtoMessage()    {}
func (*Trace) Descriptor() ([]byte, []int) {
	return fileDescriptor_91f8d9fbc62846cf, []int{0}
}

func (m *Trace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Trace.Unmarshal(m, b)
}
func (m *Trace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Trace.Marshal(b, m, deterministic)
}
func (m *Trace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Trace.Merge(m, src)
}
func (m *Trace) XXX_Size() int {
	return xxx_messageInfo_Trace.Size(m)
}
func (m *Trace) XXX_DiscardUnknown() {
	xxx_messageInfo_Trace.DiscardUnknown(m)
}

var xxx_messageInfo_Trace proto.InternalMessageInfo

func (m *Trace) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *Trace) GetTraceId() string {
	if m != nil {
		return m.TraceId
	}
	return ""
}

func (m *Trace) GetSpans() []*TraceSpan {
	if m != nil {
		return m.Spans
	}
	return nil
}

// List of new or updated traces.
type Traces struct {
	// List of traces.
	Traces               []*Trace `protobuf:"bytes,1,rep,name=traces,proto3" json:"traces,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Traces) Reset()         { *m = Traces{} }
func (m *Traces) String() string { return proto.CompactTextString(m) }
func (*Traces) ProtoMessage()    {}
func (*Traces) Descriptor() ([]byte, []int) {
	return fileDescriptor_91f8d9fbc62846cf, []int{1}
}

func (m *Traces) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Traces.Unmarshal(m, b)
}
func (m *Traces) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Traces.Marshal(b, m, deterministic)
}
func (m *Traces) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Traces.Merge(m, src)
}
func (m *Traces) XXX_Size() int {
	return xxx_messageInfo_Traces.Size(m)
}
func (m *Traces) XXX_DiscardUnknown() {
	xxx_messageInfo_Traces.DiscardUnknown(m)
}

var xxx_messageInfo_Traces proto.InternalMessageInfo

func (m *Traces) GetTraces() []*Trace {
	if m != nil {
		return m.Traces
	}
	return nil
}

// A span represents a single timed event within a trace. Spans can be nested
// and form a trace tree. Often, a trace contains a root span that describes the
// end-to-end latency of an operation and, optionally, one or more subspans for
// its suboperations. Spans do not need to be contiguous. There may be gaps
// between spans in a trace.
type TraceSpan struct {
	// Identifier for the span. Must be a 64-bit integer other than 0 and
	// unique within a trace. For example, `2205310701640571284`.
	SpanId uint64 `protobuf:"fixed64,1,opt,name=span_id,json=spanId,proto3" json:"span_id,omitempty"`
	// Distinguishes between spans generated in a particular context. For example,
	// two spans with the same name may be distinguished using `RPC_CLIENT`
	// and `RPC_SERVER` to identify queueing latency associated with the span.
	Kind TraceSpan_SpanKind `protobuf:"varint,2,opt,name=kind,proto3,enum=google.devtools.cloudtrace.v1.TraceSpan_SpanKind" json:"kind,omitempty"`
	// Name of the span. Must be less than 128 bytes. The span name is sanitized
	// and displayed in the Stackdriver Trace tool in the
	// Google Cloud Platform Console.
	// The name may be a method name or some other per-call site name.
	// For the same executable and the same call point, a best practice is
	// to use a consistent name, which makes it easier to correlate
	// cross-trace spans.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Start time of the span in nanoseconds from the UNIX epoch.
	StartTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// End time of the span in nanoseconds from the UNIX epoch.
	EndTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// Optional. ID of the parent span, if any.
	ParentSpanId uint64 `protobuf:"fixed64,6,opt,name=parent_span_id,json=parentSpanId,proto3" json:"parent_span_id,omitempty"`
	// Collection of labels associated with the span. Label keys must be less than
	// 128 bytes. Label values must be less than 16 kilobytes (10MB for
	// `/stacktrace` values).
	//
	// Some predefined label keys exist, or you may create your own. When creating
	// your own, we recommend the following formats:
	//
	// * `/category/product/key` for agents of well-known products (e.g.
	//   `/db/mongodb/read_size`).
	// * `short_host/path/key` for domain-specific keys (e.g.
	//   `foo.com/myproduct/bar`)
	//
	// Predefined labels include:
	//
	// *   `/agent`
	// *   `/component`
	// *   `/error/message`
	// *   `/error/name`
	// *   `/http/client_city`
	// *   `/http/client_country`
	// *   `/http/client_protocol`
	// *   `/http/client_region`
	// *   `/http/host`
	// *   `/http/method`
	// *   `/http/path`
	// *   `/http/redirected_url`
	// *   `/http/request/size`
	// *   `/http/response/size`
	// *   `/http/route`
	// *   `/http/status_code`
	// *   `/http/url`
	// *   `/http/user_agent`
	// *   `/pid`
	// *   `/stacktrace`
	// *   `/tid`
	Labels               map[string]string `protobuf:"bytes,7,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TraceSpan) Reset()         { *m = TraceSpan{} }
func (m *TraceSpan) String() string { return proto.CompactTextString(m) }
func (*TraceSpan) ProtoMessage()    {}
func (*TraceSpan) Descriptor() ([]byte, []int) {
	return fileDescriptor_91f8d9fbc62846cf, []int{2}
}

func (m *TraceSpan) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TraceSpan.Unmarshal(m, b)
}
func (m *TraceSpan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TraceSpan.Marshal(b, m, deterministic)
}
func (m *TraceSpan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TraceSpan.Merge(m, src)
}
func (m *TraceSpan) XXX_Size() int {
	return xxx_messageInfo_TraceSpan.Size(m)
}
func (m *TraceSpan) XXX_DiscardUnknown() {
	xxx_messageInfo_TraceSpan.DiscardUnknown(m)
}

var xxx_messageInfo_TraceSpan proto.InternalMessageInfo

func (m *TraceSpan) GetSpanId() uint64 {
	if m != nil {
		return m.SpanId
	}
	return 0
}

func (m *TraceSpan) GetKind() TraceSpan_SpanKind {
	if m != nil {
		return m.Kind
	}
	return TraceSpan_SPAN_KIND_UNSPECIFIED
}

func (m *TraceSpan) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TraceSpan) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *TraceSpan) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *TraceSpan) GetParentSpanId() uint64 {
	if m != nil {
		return m.ParentSpanId
	}
	return 0
}

func (m *TraceSpan) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

// The request message for the `ListTraces` method. All fields are required
// unless specified.
type ListTracesRequest struct {
	// Required. ID of the Cloud project where the trace data is stored.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// Optional. Type of data returned for traces in the list. Default is
	// `MINIMAL`.
	View ListTracesRequest_ViewType `protobuf:"varint,2,opt,name=view,proto3,enum=google.devtools.cloudtrace.v1.ListTracesRequest_ViewType" json:"view,omitempty"`
	// Optional. Maximum number of traces to return. If not specified or <= 0, the
	// implementation selects a reasonable value.  The implementation may
	// return fewer traces than the requested page size.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Token identifying the page of results to return. If provided, use the
	// value of the `next_page_token` field from a previous request.
	PageToken string `protobuf:"bytes,4,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Start of the time interval (inclusive) during which the trace data was
	// collected from the application.
	StartTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// End of the time interval (inclusive) during which the trace data was
	// collected from the application.
	EndTime *timestamp.Timestamp `protobuf:"bytes,6,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// Optional. A filter against labels for the request.
	//
	// By default, searches use prefix matching. To specify exact match, prepend
	// a plus symbol (`+`) to the search term.
	// Multiple terms are ANDed. Syntax:
	//
	// *   `root:NAME_PREFIX` or `NAME_PREFIX`: Return traces where any root
	//     span starts with `NAME_PREFIX`.
	// *   `+root:NAME` or `+NAME`: Return traces where any root span's name is
	//     exactly `NAME`.
	// *   `span:NAME_PREFIX`: Return traces where any span starts with
	//     `NAME_PREFIX`.
	// *   `+span:NAME`: Return traces where any span's name is exactly
	//     `NAME`.
	// *   `latency:DURATION`: Return traces whose overall latency is
	//     greater or equal to than `DURATION`. Accepted units are nanoseconds
	//     (`ns`), milliseconds (`ms`), and seconds (`s`). Default is `ms`. For
	//     example, `latency:24ms` returns traces whose overall latency
	//     is greater than or equal to 24 milliseconds.
	// *   `label:LABEL_KEY`: Return all traces containing the specified
	//     label key (exact match, case-sensitive) regardless of the key:value
	//     pair's value (including empty values).
	// *   `LABEL_KEY:VALUE_PREFIX`: Return all traces containing the specified
	//     label key (exact match, case-sensitive) whose value starts with
	//     `VALUE_PREFIX`. Both a key and a value must be specified.
	// *   `+LABEL_KEY:VALUE`: Return all traces containing a key:value pair
	//     exactly matching the specified text. Both a key and a value must be
	//     specified.
	// *   `method:VALUE`: Equivalent to `/http/method:VALUE`.
	// *   `url:VALUE`: Equivalent to `/http/url:VALUE`.
	Filter string `protobuf:"bytes,7,opt,name=filter,proto3" json:"filter,omitempty"`
	// Optional. Field used to sort the returned traces.
	// Can be one of the following:
	//
	// *   `trace_id`
	// *   `name` (`name` field of root span in the trace)
	// *   `duration` (difference between `end_time` and `start_time` fields of
	//      the root span)
	// *   `start` (`start_time` field of the root span)
	//
	// Descending order can be specified by appending `desc` to the sort field
	// (for example, `name desc`).
	//
	// Only one sort field is permitted.
	OrderBy              string   `protobuf:"bytes,8,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTracesRequest) Reset()         { *m = ListTracesRequest{} }
func (m *ListTracesRequest) String() string { return proto.CompactTextString(m) }
func (*ListTracesRequest) ProtoMessage()    {}
func (*ListTracesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_91f8d9fbc62846cf, []int{3}
}

func (m *ListTracesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTracesRequest.Unmarshal(m, b)
}
func (m *ListTracesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTracesRequest.Marshal(b, m, deterministic)
}
func (m *ListTracesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTracesRequest.Merge(m, src)
}
func (m *ListTracesRequest) XXX_Size() int {
	return xxx_messageInfo_ListTracesRequest.Size(m)
}
func (m *ListTracesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTracesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListTracesRequest proto.InternalMessageInfo

func (m *ListTracesRequest) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *ListTracesRequest) GetView() ListTracesRequest_ViewType {
	if m != nil {
		return m.View
	}
	return ListTracesRequest_VIEW_TYPE_UNSPECIFIED
}

func (m *ListTracesRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListTracesRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListTracesRequest) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *ListTracesRequest) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *ListTracesRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *ListTracesRequest) GetOrderBy() string {
	if m != nil {
		return m.OrderBy
	}
	return ""
}

// The response message for the `ListTraces` method.
type ListTracesResponse struct {
	// List of trace records as specified by the view parameter.
	Traces []*Trace `protobuf:"bytes,1,rep,name=traces,proto3" json:"traces,omitempty"`
	// If defined, indicates that there are more traces that match the request
	// and that this value should be passed to the next request to continue
	// retrieving additional traces.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTracesResponse) Reset()         { *m = ListTracesResponse{} }
func (m *ListTracesResponse) String() string { return proto.CompactTextString(m) }
func (*ListTracesResponse) ProtoMessage()    {}
func (*ListTracesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_91f8d9fbc62846cf, []int{4}
}

func (m *ListTracesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTracesResponse.Unmarshal(m, b)
}
func (m *ListTracesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTracesResponse.Marshal(b, m, deterministic)
}
func (m *ListTracesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTracesResponse.Merge(m, src)
}
func (m *ListTracesResponse) XXX_Size() int {
	return xxx_messageInfo_ListTracesResponse.Size(m)
}
func (m *ListTracesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTracesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListTracesResponse proto.InternalMessageInfo

func (m *ListTracesResponse) GetTraces() []*Trace {
	if m != nil {
		return m.Traces
	}
	return nil
}

func (m *ListTracesResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// The request message for the `GetTrace` method.
type GetTraceRequest struct {
	// Required. ID of the Cloud project where the trace data is stored.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// Required. ID of the trace to return.
	TraceId              string   `protobuf:"bytes,2,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTraceRequest) Reset()         { *m = GetTraceRequest{} }
func (m *GetTraceRequest) String() string { return proto.CompactTextString(m) }
func (*GetTraceRequest) ProtoMessage()    {}
func (*GetTraceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_91f8d9fbc62846cf, []int{5}
}

func (m *GetTraceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTraceRequest.Unmarshal(m, b)
}
func (m *GetTraceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTraceRequest.Marshal(b, m, deterministic)
}
func (m *GetTraceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTraceRequest.Merge(m, src)
}
func (m *GetTraceRequest) XXX_Size() int {
	return xxx_messageInfo_GetTraceRequest.Size(m)
}
func (m *GetTraceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTraceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTraceRequest proto.InternalMessageInfo

func (m *GetTraceRequest) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *GetTraceRequest) GetTraceId() string {
	if m != nil {
		return m.TraceId
	}
	return ""
}

// The request message for the `PatchTraces` method.
type PatchTracesRequest struct {
	// Required. ID of the Cloud project where the trace data is stored.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// Required. The body of the message.
	Traces               *Traces  `protobuf:"bytes,2,opt,name=traces,proto3" json:"traces,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PatchTracesRequest) Reset()         { *m = PatchTracesRequest{} }
func (m *PatchTracesRequest) String() string { return proto.CompactTextString(m) }
func (*PatchTracesRequest) ProtoMessage()    {}
func (*PatchTracesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_91f8d9fbc62846cf, []int{6}
}

func (m *PatchTracesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PatchTracesRequest.Unmarshal(m, b)
}
func (m *PatchTracesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PatchTracesRequest.Marshal(b, m, deterministic)
}
func (m *PatchTracesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PatchTracesRequest.Merge(m, src)
}
func (m *PatchTracesRequest) XXX_Size() int {
	return xxx_messageInfo_PatchTracesRequest.Size(m)
}
func (m *PatchTracesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PatchTracesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PatchTracesRequest proto.InternalMessageInfo

func (m *PatchTracesRequest) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *PatchTracesRequest) GetTraces() *Traces {
	if m != nil {
		return m.Traces
	}
	return nil
}

func init() {
	proto.RegisterEnum("google.devtools.cloudtrace.v1.TraceSpan_SpanKind", TraceSpan_SpanKind_name, TraceSpan_SpanKind_value)
	proto.RegisterEnum("google.devtools.cloudtrace.v1.ListTracesRequest_ViewType", ListTracesRequest_ViewType_name, ListTracesRequest_ViewType_value)
	proto.RegisterType((*Trace)(nil), "google.devtools.cloudtrace.v1.Trace")
	proto.RegisterType((*Traces)(nil), "google.devtools.cloudtrace.v1.Traces")
	proto.RegisterType((*TraceSpan)(nil), "google.devtools.cloudtrace.v1.TraceSpan")
	proto.RegisterMapType((map[string]string)(nil), "google.devtools.cloudtrace.v1.TraceSpan.LabelsEntry")
	proto.RegisterType((*ListTracesRequest)(nil), "google.devtools.cloudtrace.v1.ListTracesRequest")
	proto.RegisterType((*ListTracesResponse)(nil), "google.devtools.cloudtrace.v1.ListTracesResponse")
	proto.RegisterType((*GetTraceRequest)(nil), "google.devtools.cloudtrace.v1.GetTraceRequest")
	proto.RegisterType((*PatchTracesRequest)(nil), "google.devtools.cloudtrace.v1.PatchTracesRequest")
}

func init() {
	proto.RegisterFile("google/devtools/cloudtrace/v1/trace.proto", fileDescriptor_91f8d9fbc62846cf)
}

var fileDescriptor_91f8d9fbc62846cf = []byte{
	// 1039 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xcf, 0x6f, 0x1b, 0xc5,
	0x17, 0xff, 0xae, 0x1d, 0xff, 0xc8, 0x4b, 0xbe, 0xa9, 0x3b, 0x50, 0xea, 0xb8, 0xb4, 0x84, 0x15,
	0xa0, 0x14, 0xda, 0x5d, 0x1c, 0xa8, 0x44, 0x02, 0x42, 0x72, 0x92, 0x6d, 0xb4, 0xaa, 0xe3, 0x58,
	0x6b, 0x37, 0x08, 0x14, 0x69, 0x35, 0xf1, 0x4e, 0x9c, 0x25, 0xeb, 0x9d, 0x65, 0x67, 0xec, 0xe0,
	0x94, 0x1e, 0xe0, 0xc8, 0x95, 0x3b, 0x67, 0x14, 0xce, 0xfc, 0x13, 0xe5, 0x06, 0xb7, 0x9c, 0x38,
	0x70, 0xe4, 0x8f, 0x40, 0x33, 0xb3, 0x1b, 0x5b, 0x8e, 0xa8, 0x1d, 0xb8, 0x58, 0xf3, 0x7e, 0xbf,
	0xf7, 0x79, 0x9f, 0x19, 0x2f, 0xdc, 0xef, 0x52, 0xda, 0x0d, 0x88, 0xe9, 0x91, 0x01, 0xa7, 0x34,
	0x60, 0x66, 0x27, 0xa0, 0x7d, 0x8f, 0xc7, 0xb8, 0x43, 0xcc, 0x41, 0xd5, 0x94, 0x07, 0x23, 0x8a,
	0x29, 0xa7, 0xe8, 0xae, 0x72, 0x35, 0x52, 0x57, 0x63, 0xe4, 0x6a, 0x0c, 0xaa, 0x95, 0xd7, 0x93,
	0x4c, 0x38, 0xf2, 0x4d, 0x1c, 0x86, 0x94, 0x63, 0xee, 0xd3, 0x90, 0xa9, 0xe0, 0xca, 0xed, 0x31,
	0x6b, 0x27, 0xf0, 0x49, 0xc8, 0x13, 0xc3, 0x1b, 0x63, 0x86, 0x23, 0x9f, 0x04, 0x9e, 0x7b, 0x48,
	0x8e, 0xf1, 0xc0, 0xa7, 0x71, 0xe2, 0xb0, 0x3c, 0xe6, 0x10, 0x13, 0x46, 0xfb, 0x71, 0xda, 0x51,
	0xe5, 0x4e, 0x62, 0x92, 0xd2, 0x61, 0xff, 0xc8, 0x24, 0xbd, 0x88, 0x0f, 0x27, 0x12, 0x5f, 0x1a,
	0xb9, 0xdf, 0x23, 0x8c, 0xe3, 0x5e, 0xa4, 0x1c, 0xf4, 0x6f, 0x35, 0xc8, 0xb5, 0x45, 0xf7, 0xe8,
	0x2e, 0x40, 0x14, 0xd3, 0x2f, 0x49, 0x87, 0xbb, 0xbe, 0x57, 0xd6, 0x56, 0xb4, 0xd5, 0x79, 0x67,
	0x3e, 0xd1, 0xd8, 0x1e, 0x5a, 0x86, 0xa2, 0x9c, 0x52, 0x18, 0x33, 0xd2, 0x58, 0x90, 0xb2, 0xed,
	0xa1, 0x4f, 0x21, 0xc7, 0x22, 0x1c, 0xb2, 0x72, 0x76, 0x25, 0xbb, 0xba, 0xb0, 0xb6, 0x6a, 0xbc,
	0x14, 0x23, 0x43, 0x96, 0x6b, 0x45, 0x38, 0x74, 0x54, 0x98, 0xfe, 0x18, 0xf2, 0x52, 0xc7, 0xd0,
	0x27, 0x90, 0x97, 0x6e, 0xac, 0xac, 0xc9, 0x54, 0x6f, 0xcd, 0x92, 0xca, 0x49, 0x62, 0xf4, 0xbf,
	0xb2, 0x30, 0x7f, 0x99, 0x1c, 0xdd, 0x86, 0x82, 0x48, 0x9f, 0x0e, 0x93, 0x77, 0xf2, 0x42, 0xb4,
	0x3d, 0x64, 0xc1, 0xdc, 0x89, 0x1f, 0xaa, 0x29, 0x96, 0xd6, 0xaa, 0xb3, 0x76, 0x6b, 0x88, 0x9f,
	0x27, 0x7e, 0xe8, 0x39, 0x32, 0x1c, 0x21, 0x98, 0x0b, 0x71, 0x8f, 0x94, 0xb3, 0x12, 0x0c, 0x79,
	0x46, 0xeb, 0x00, 0x8c, 0xe3, 0x98, 0xbb, 0x02, 0xe6, 0xf2, 0xdc, 0x8a, 0xb6, 0xba, 0xb0, 0x56,
	0x49, 0x0b, 0xa4, 0x3b, 0x30, 0xda, 0xe9, 0x0e, 0x9c, 0x79, 0xe9, 0x2d, 0x64, 0xf4, 0x08, 0x8a,
	0x24, 0xf4, 0x54, 0x60, 0x6e, 0x6a, 0x60, 0x81, 0x84, 0x9e, 0x0c, 0xbb, 0x0f, 0x4b, 0x11, 0x8e,
	0x49, 0xc8, 0xdd, 0x74, 0xd8, 0xbc, 0x18, 0x76, 0x33, 0xfb, 0x47, 0x4d, 0x73, 0x16, 0x95, 0xa9,
	0xa5, 0xe6, 0xae, 0x43, 0x3e, 0xc0, 0x87, 0x24, 0x60, 0xe5, 0x82, 0x04, 0xf7, 0xc3, 0x99, 0x27,
	0xaf, 0xcb, 0x30, 0x2b, 0xe4, 0xf1, 0xd0, 0x49, 0x72, 0x54, 0xd6, 0x61, 0x61, 0x4c, 0x8d, 0x4a,
	0x90, 0x3d, 0x21, 0xc3, 0x84, 0x36, 0xe2, 0x88, 0x5e, 0x85, 0xdc, 0x00, 0x07, 0x7d, 0x92, 0xb0,
	0x45, 0x09, 0x1b, 0x99, 0x8f, 0x34, 0xdd, 0x82, 0x62, 0x8a, 0x25, 0x5a, 0x86, 0x5b, 0xad, 0x66,
	0xad, 0xe1, 0x3e, 0xb1, 0x1b, 0xdb, 0xee, 0xd3, 0x46, 0xab, 0x69, 0x6d, 0xd9, 0x8f, 0x6d, 0x6b,
	0xbb, 0xf4, 0x3f, 0xb4, 0x04, 0xe0, 0x34, 0xb7, 0xdc, 0x96, 0xe5, 0xec, 0x5b, 0x4e, 0x49, 0x4b,
	0xe5, 0xad, 0xba, 0x6d, 0x35, 0xda, 0xa5, 0x8c, 0xfe, 0x6b, 0x16, 0x6e, 0xd6, 0x7d, 0xc6, 0x15,
	0x77, 0x1c, 0xf2, 0x55, 0x9f, 0x30, 0x8e, 0xf4, 0xab, 0x34, 0x16, 0x60, 0x64, 0xc6, 0xb9, 0xec,
	0xc0, 0xdc, 0xc0, 0x27, 0xa7, 0x09, 0x03, 0xd6, 0xa7, 0xe0, 0x70, 0xa5, 0x86, 0xb1, 0xef, 0x93,
	0xd3, 0xf6, 0x30, 0x22, 0x0a, 0x65, 0x99, 0x0b, 0xad, 0xc0, 0x7c, 0x84, 0xbb, 0xc4, 0x65, 0xfe,
	0x99, 0xe2, 0x44, 0x4e, 0x59, 0x8b, 0x42, 0xdb, 0xf2, 0xcf, 0xd4, 0x05, 0x13, 0x1e, 0x9c, 0x9e,
	0x90, 0x50, 0x92, 0x43, 0x5c, 0x30, 0xdc, 0x25, 0x6d, 0xa1, 0x98, 0xe0, 0x4e, 0xee, 0xdf, 0x72,
	0x27, 0x3f, 0x3b, 0x77, 0xee, 0x40, 0xfe, 0xc8, 0x0f, 0x38, 0x89, 0xcb, 0x85, 0x14, 0x26, 0xcd,
	0x49, 0x54, 0xe8, 0x1e, 0x14, 0x69, 0xec, 0x91, 0xd8, 0x3d, 0x1c, 0x96, 0x8b, 0x23, 0x73, 0x41,
	0x2a, 0x37, 0x87, 0x7a, 0x03, 0x8a, 0x29, 0x0c, 0x62, 0x89, 0xfb, 0xb6, 0xf5, 0x99, 0xdb, 0xfe,
	0xbc, 0x69, 0x4d, 0x2c, 0x71, 0x01, 0x0a, 0xbb, 0x76, 0xc3, 0xde, 0xad, 0xd5, 0x4b, 0x1a, 0x5a,
	0x84, 0xa2, 0xb3, 0xb7, 0xd7, 0x16, 0x0b, 0x2f, 0x65, 0x84, 0xb4, 0xb5, 0xb7, 0xdb, 0xac, 0x5b,
	0x6d, 0xab, 0x94, 0xd5, 0xcf, 0x00, 0x8d, 0x03, 0xcd, 0x22, 0x1a, 0x32, 0xf2, 0xdf, 0x1e, 0x04,
	0xf4, 0x0e, 0xdc, 0x08, 0xc9, 0xd7, 0xdc, 0x1d, 0x83, 0x5d, 0x91, 0xf1, 0xff, 0x42, 0xdd, 0x4c,
	0xa1, 0xd7, 0x9f, 0xc2, 0x8d, 0x1d, 0xa2, 0x4a, 0x5f, 0x87, 0x46, 0xf7, 0x26, 0x9f, 0x44, 0xe5,
	0x91, 0xbe, 0x8b, 0xfa, 0x37, 0x80, 0x9a, 0x98, 0x77, 0x8e, 0xaf, 0x4f, 0xd0, 0xcd, 0xcb, 0xb1,
	0x33, 0x72, 0x9d, 0x6f, 0xcf, 0x32, 0x36, 0x53, 0x69, 0x92, 0xc8, 0xb5, 0x5f, 0x72, 0xb0, 0xa8,
	0xae, 0x30, 0x89, 0x07, 0x7e, 0x87, 0xa0, 0x9f, 0x34, 0x80, 0x11, 0xc4, 0xe8, 0xfd, 0xeb, 0xd2,
	0xbe, 0x52, 0xbd, 0x46, 0x84, 0xda, 0x9f, 0xfe, 0xe8, 0xa2, 0x36, 0x36, 0xed, 0x77, 0xbf, 0xff,
	0xf9, 0x43, 0x46, 0x47, 0x2b, 0xe2, 0x4f, 0x35, 0xd1, 0x32, 0xf3, 0xd9, 0xc8, 0xfe, 0xdc, 0x4c,
	0x16, 0xf7, 0xb3, 0x06, 0xc5, 0x74, 0x23, 0xc8, 0x98, 0x52, 0x76, 0x62, 0x75, 0x95, 0x99, 0x38,
	0xa2, 0xdb, 0x17, 0xb5, 0x57, 0x46, 0x95, 0x1f, 0xa4, 0x8b, 0x94, 0x2d, 0x3e, 0x44, 0xef, 0x4d,
	0x6b, 0xd1, 0x7c, 0x96, 0x46, 0x3c, 0x47, 0x3f, 0x6a, 0xb0, 0x30, 0xb6, 0x68, 0x34, 0x0d, 0xa7,
	0xab, 0xa4, 0xa8, 0xbc, 0x76, 0xe5, 0xbe, 0x5a, 0xe2, 0x5f, 0x5c, 0xdf, 0xbe, 0xa8, 0xdd, 0x9c,
	0xec, 0x92, 0xc9, 0x1e, 0xdf, 0x5d, 0x9b, 0x0a, 0xe3, 0x46, 0x42, 0x85, 0xca, 0xb9, 0xf6, 0xa2,
	0xb6, 0x3c, 0xd6, 0x83, 0xaa, 0x85, 0x23, 0x9f, 0x19, 0x1d, 0xda, 0xfb, 0xad, 0xf6, 0xbd, 0x76,
	0xcc, 0x79, 0xc4, 0x36, 0x4c, 0xf3, 0xf4, 0xf4, 0x74, 0xc2, 0x6a, 0xe2, 0x3e, 0x3f, 0x56, 0x1f,
	0x44, 0x0f, 0xa3, 0x00, 0xf3, 0x23, 0x1a, 0xf7, 0x1e, 0x4c, 0x73, 0x57, 0x65, 0x70, 0x14, 0x91,
	0xd0, 0x9b, 0xd1, 0x39, 0x26, 0xd8, 0xa3, 0x61, 0x30, 0xdc, 0x3c, 0xd7, 0xe0, 0xcd, 0x0e, 0xed,
	0xbd, 0x1c, 0xc3, 0x4d, 0x90, 0xf8, 0x35, 0x05, 0x5a, 0x4d, 0xed, 0x8b, 0x9d, 0xc4, 0xb9, 0x4b,
	0x03, 0x1c, 0x76, 0x0d, 0x1a, 0x77, 0xcd, 0x2e, 0x09, 0x25, 0x96, 0xe6, 0xa8, 0xe4, 0x3f, 0x7c,
	0xdf, 0x7d, 0x3c, 0x92, 0xce, 0x33, 0xb7, 0x76, 0x54, 0xa6, 0x2d, 0xa1, 0x53, 0x54, 0x31, 0xf6,
	0xab, 0x2f, 0x52, 0xfd, 0x81, 0xd4, 0x1f, 0x48, 0xfd, 0xc1, 0x7e, 0xf5, 0x30, 0x2f, 0x6b, 0x7c,
	0xf0, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x34, 0xe0, 0x43, 0x75, 0x42, 0x0a, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TraceServiceClient is the client API for TraceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TraceServiceClient interface {
	// Returns of a list of traces that match the specified filter conditions.
	ListTraces(ctx context.Context, in *ListTracesRequest, opts ...grpc.CallOption) (*ListTracesResponse, error)
	// Gets a single trace by its ID.
	GetTrace(ctx context.Context, in *GetTraceRequest, opts ...grpc.CallOption) (*Trace, error)
	// Sends new traces to Stackdriver Trace or updates existing traces. If the ID
	// of a trace that you send matches that of an existing trace, any fields
	// in the existing trace and its spans are overwritten by the provided values,
	// and any new fields provided are merged with the existing trace data. If the
	// ID does not match, a new trace is created.
	PatchTraces(ctx context.Context, in *PatchTracesRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type traceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTraceServiceClient(cc grpc.ClientConnInterface) TraceServiceClient {
	return &traceServiceClient{cc}
}

func (c *traceServiceClient) ListTraces(ctx context.Context, in *ListTracesRequest, opts ...grpc.CallOption) (*ListTracesResponse, error) {
	out := new(ListTracesResponse)
	err := c.cc.Invoke(ctx, "/google.devtools.cloudtrace.v1.TraceService/ListTraces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceServiceClient) GetTrace(ctx context.Context, in *GetTraceRequest, opts ...grpc.CallOption) (*Trace, error) {
	out := new(Trace)
	err := c.cc.Invoke(ctx, "/google.devtools.cloudtrace.v1.TraceService/GetTrace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceServiceClient) PatchTraces(ctx context.Context, in *PatchTracesRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/google.devtools.cloudtrace.v1.TraceService/PatchTraces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TraceServiceServer is the server API for TraceService service.
type TraceServiceServer interface {
	// Returns of a list of traces that match the specified filter conditions.
	ListTraces(context.Context, *ListTracesRequest) (*ListTracesResponse, error)
	// Gets a single trace by its ID.
	GetTrace(context.Context, *GetTraceRequest) (*Trace, error)
	// Sends new traces to Stackdriver Trace or updates existing traces. If the ID
	// of a trace that you send matches that of an existing trace, any fields
	// in the existing trace and its spans are overwritten by the provided values,
	// and any new fields provided are merged with the existing trace data. If the
	// ID does not match, a new trace is created.
	PatchTraces(context.Context, *PatchTracesRequest) (*empty.Empty, error)
}

// UnimplementedTraceServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTraceServiceServer struct {
}

func (*UnimplementedTraceServiceServer) ListTraces(ctx context.Context, req *ListTracesRequest) (*ListTracesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTraces not implemented")
}
func (*UnimplementedTraceServiceServer) GetTrace(ctx context.Context, req *GetTraceRequest) (*Trace, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrace not implemented")
}
func (*UnimplementedTraceServiceServer) PatchTraces(ctx context.Context, req *PatchTracesRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchTraces not implemented")
}

func RegisterTraceServiceServer(s *grpc.Server, srv TraceServiceServer) {
	s.RegisterService(&_TraceService_serviceDesc, srv)
}

func _TraceService_ListTraces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTracesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraceServiceServer).ListTraces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.cloudtrace.v1.TraceService/ListTraces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraceServiceServer).ListTraces(ctx, req.(*ListTracesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TraceService_GetTrace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTraceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraceServiceServer).GetTrace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.cloudtrace.v1.TraceService/GetTrace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraceServiceServer).GetTrace(ctx, req.(*GetTraceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TraceService_PatchTraces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchTracesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraceServiceServer).PatchTraces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.cloudtrace.v1.TraceService/PatchTraces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraceServiceServer).PatchTraces(ctx, req.(*PatchTracesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TraceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.devtools.cloudtrace.v1.TraceService",
	HandlerType: (*TraceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListTraces",
			Handler:    _TraceService_ListTraces_Handler,
		},
		{
			MethodName: "GetTrace",
			Handler:    _TraceService_GetTrace_Handler,
		},
		{
			MethodName: "PatchTraces",
			Handler:    _TraceService_PatchTraces_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/devtools/cloudtrace/v1/trace.proto",
}
