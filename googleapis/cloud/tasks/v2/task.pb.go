// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/tasks/v2/task.proto

package tasks

import (
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
	duration "github.com/catper/protobuf/ptypes/duration"
	timestamp "github.com/catper/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
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

// The view specifies a subset of [Task][google.cloud.tasks.v2.Task] data.
//
// When a task is returned in a response, not all
// information is retrieved by default because some data, such as
// payloads, might be desirable to return only when needed because
// of its large size or because of the sensitivity of data that it
// contains.
type Task_View int32

const (
	// Unspecified. Defaults to BASIC.
	Task_VIEW_UNSPECIFIED Task_View = 0
	// The basic view omits fields which can be large or can contain
	// sensitive data.
	//
	// This view does not include the
	// [body in AppEngineHttpRequest][google.cloud.tasks.v2.AppEngineHttpRequest.body].
	// Bodies are desirable to return only when needed, because they
	// can be large and because of the sensitivity of the data that you
	// choose to store in it.
	Task_BASIC Task_View = 1
	// All information is returned.
	//
	// Authorization for [FULL][google.cloud.tasks.v2.Task.View.FULL] requires
	// `cloudtasks.tasks.fullView` [Google IAM](https://cloud.google.com/iam/)
	// permission on the [Queue][google.cloud.tasks.v2.Queue] resource.
	Task_FULL Task_View = 2
)

var Task_View_name = map[int32]string{
	0: "VIEW_UNSPECIFIED",
	1: "BASIC",
	2: "FULL",
}

var Task_View_value = map[string]int32{
	"VIEW_UNSPECIFIED": 0,
	"BASIC":            1,
	"FULL":             2,
}

func (x Task_View) String() string {
	return proto.EnumName(Task_View_name, int32(x))
}

func (Task_View) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5d55198283613840, []int{0, 0}
}

// A unit of scheduled work.
type Task struct {
	// Optionally caller-specified in [CreateTask][google.cloud.tasks.v2.CloudTasks.CreateTask].
	//
	// The task name.
	//
	// The task name must have the following format:
	// `projects/PROJECT_ID/locations/LOCATION_ID/queues/QUEUE_ID/tasks/TASK_ID`
	//
	// * `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]),
	//    hyphens (-), colons (:), or periods (.).
	//    For more information, see
	//    [Identifying
	//    projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects)
	// * `LOCATION_ID` is the canonical ID for the task's location.
	//    The list of available locations can be obtained by calling
	//    [ListLocations][google.cloud.location.Locations.ListLocations].
	//    For more information, see https://cloud.google.com/about/locations/.
	// * `QUEUE_ID` can contain letters ([A-Za-z]), numbers ([0-9]), or
	//   hyphens (-). The maximum length is 100 characters.
	// * `TASK_ID` can contain only letters ([A-Za-z]), numbers ([0-9]),
	//   hyphens (-), or underscores (_). The maximum length is 500 characters.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The message to send to the worker.
	//
	// Types that are valid to be assigned to MessageType:
	//	*Task_AppEngineHttpRequest
	//	*Task_HttpRequest
	MessageType isTask_MessageType `protobuf_oneof:"message_type"`
	// The time when the task is scheduled to be attempted or retried.
	//
	// `schedule_time` will be truncated to the nearest microsecond.
	ScheduleTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=schedule_time,json=scheduleTime,proto3" json:"schedule_time,omitempty"`
	// Output only. The time that the task was created.
	//
	// `create_time` will be truncated to the nearest second.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The deadline for requests sent to the worker. If the worker does not
	// respond by this deadline then the request is cancelled and the attempt
	// is marked as a `DEADLINE_EXCEEDED` failure. Cloud Tasks will retry the
	// task according to the [RetryConfig][google.cloud.tasks.v2.RetryConfig].
	//
	// Note that when the request is cancelled, Cloud Tasks will stop listing for
	// the response, but whether the worker stops processing depends on the
	// worker. For example, if the worker is stuck, it may not react to cancelled
	// requests.
	//
	// The default and maximum values depend on the type of request:
	//
	// * For [HTTP tasks][google.cloud.tasks.v2.HttpRequest], the default is 10 minutes. The deadline
	//   must be in the interval [15 seconds, 30 minutes].
	//
	// * For [App Engine tasks][google.cloud.tasks.v2.AppEngineHttpRequest], 0 indicates that the
	//   request has the default deadline. The default deadline depends on the
	//   [scaling
	//   type](https://cloud.google.com/appengine/docs/standard/go/how-instances-are-managed#instance_scaling)
	//   of the service: 10 minutes for standard apps with automatic scaling, 24
	//   hours for standard apps with manual and basic scaling, and 60 minutes for
	//   flex apps. If the request deadline is set, it must be in the interval [15
	//   seconds, 24 hours 15 seconds]. Regardless of the task's
	//   `dispatch_deadline`, the app handler will not run for longer than than
	//   the service's timeout. We recommend setting the `dispatch_deadline` to
	//   at most a few seconds more than the app handler's timeout. For more
	//   information see
	//   [Timeouts](https://cloud.google.com/tasks/docs/creating-appengine-handlers#timeouts).
	//
	// `dispatch_deadline` will be truncated to the nearest millisecond. The
	// deadline is an approximate deadline.
	DispatchDeadline *duration.Duration `protobuf:"bytes,6,opt,name=dispatch_deadline,json=dispatchDeadline,proto3" json:"dispatch_deadline,omitempty"`
	// Output only. The number of attempts dispatched.
	//
	// This count includes attempts which have been dispatched but haven't
	// received a response.
	DispatchCount int32 `protobuf:"varint,7,opt,name=dispatch_count,json=dispatchCount,proto3" json:"dispatch_count,omitempty"`
	// Output only. The number of attempts which have received a response.
	ResponseCount int32 `protobuf:"varint,8,opt,name=response_count,json=responseCount,proto3" json:"response_count,omitempty"`
	// Output only. The status of the task's first attempt.
	//
	// Only [dispatch_time][google.cloud.tasks.v2.Attempt.dispatch_time] will be set.
	// The other [Attempt][google.cloud.tasks.v2.Attempt] information is not retained by Cloud Tasks.
	FirstAttempt *Attempt `protobuf:"bytes,9,opt,name=first_attempt,json=firstAttempt,proto3" json:"first_attempt,omitempty"`
	// Output only. The status of the task's last attempt.
	LastAttempt *Attempt `protobuf:"bytes,10,opt,name=last_attempt,json=lastAttempt,proto3" json:"last_attempt,omitempty"`
	// Output only. The view specifies which subset of the [Task][google.cloud.tasks.v2.Task] has
	// been returned.
	View                 Task_View `protobuf:"varint,11,opt,name=view,proto3,enum=google.cloud.tasks.v2.Task_View" json:"view,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d55198283613840, []int{0}
}

func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (m *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(m, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isTask_MessageType interface {
	isTask_MessageType()
}

type Task_AppEngineHttpRequest struct {
	AppEngineHttpRequest *AppEngineHttpRequest `protobuf:"bytes,2,opt,name=app_engine_http_request,json=appEngineHttpRequest,proto3,oneof"`
}

type Task_HttpRequest struct {
	HttpRequest *HttpRequest `protobuf:"bytes,3,opt,name=http_request,json=httpRequest,proto3,oneof"`
}

func (*Task_AppEngineHttpRequest) isTask_MessageType() {}

func (*Task_HttpRequest) isTask_MessageType() {}

func (m *Task) GetMessageType() isTask_MessageType {
	if m != nil {
		return m.MessageType
	}
	return nil
}

func (m *Task) GetAppEngineHttpRequest() *AppEngineHttpRequest {
	if x, ok := m.GetMessageType().(*Task_AppEngineHttpRequest); ok {
		return x.AppEngineHttpRequest
	}
	return nil
}

func (m *Task) GetHttpRequest() *HttpRequest {
	if x, ok := m.GetMessageType().(*Task_HttpRequest); ok {
		return x.HttpRequest
	}
	return nil
}

func (m *Task) GetScheduleTime() *timestamp.Timestamp {
	if m != nil {
		return m.ScheduleTime
	}
	return nil
}

func (m *Task) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Task) GetDispatchDeadline() *duration.Duration {
	if m != nil {
		return m.DispatchDeadline
	}
	return nil
}

func (m *Task) GetDispatchCount() int32 {
	if m != nil {
		return m.DispatchCount
	}
	return 0
}

func (m *Task) GetResponseCount() int32 {
	if m != nil {
		return m.ResponseCount
	}
	return 0
}

func (m *Task) GetFirstAttempt() *Attempt {
	if m != nil {
		return m.FirstAttempt
	}
	return nil
}

func (m *Task) GetLastAttempt() *Attempt {
	if m != nil {
		return m.LastAttempt
	}
	return nil
}

func (m *Task) GetView() Task_View {
	if m != nil {
		return m.View
	}
	return Task_VIEW_UNSPECIFIED
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Task) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Task_AppEngineHttpRequest)(nil),
		(*Task_HttpRequest)(nil),
	}
}

// The status of a task attempt.
type Attempt struct {
	// Output only. The time that this attempt was scheduled.
	//
	// `schedule_time` will be truncated to the nearest microsecond.
	ScheduleTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=schedule_time,json=scheduleTime,proto3" json:"schedule_time,omitempty"`
	// Output only. The time that this attempt was dispatched.
	//
	// `dispatch_time` will be truncated to the nearest microsecond.
	DispatchTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=dispatch_time,json=dispatchTime,proto3" json:"dispatch_time,omitempty"`
	// Output only. The time that this attempt response was received.
	//
	// `response_time` will be truncated to the nearest microsecond.
	ResponseTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=response_time,json=responseTime,proto3" json:"response_time,omitempty"`
	// Output only. The response from the worker for this attempt.
	//
	// If `response_time` is unset, then the task has not been attempted or is
	// currently running and the `response_status` field is meaningless.
	ResponseStatus       *status.Status `protobuf:"bytes,4,opt,name=response_status,json=responseStatus,proto3" json:"response_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Attempt) Reset()         { *m = Attempt{} }
func (m *Attempt) String() string { return proto.CompactTextString(m) }
func (*Attempt) ProtoMessage()    {}
func (*Attempt) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d55198283613840, []int{1}
}

func (m *Attempt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Attempt.Unmarshal(m, b)
}
func (m *Attempt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Attempt.Marshal(b, m, deterministic)
}
func (m *Attempt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Attempt.Merge(m, src)
}
func (m *Attempt) XXX_Size() int {
	return xxx_messageInfo_Attempt.Size(m)
}
func (m *Attempt) XXX_DiscardUnknown() {
	xxx_messageInfo_Attempt.DiscardUnknown(m)
}

var xxx_messageInfo_Attempt proto.InternalMessageInfo

func (m *Attempt) GetScheduleTime() *timestamp.Timestamp {
	if m != nil {
		return m.ScheduleTime
	}
	return nil
}

func (m *Attempt) GetDispatchTime() *timestamp.Timestamp {
	if m != nil {
		return m.DispatchTime
	}
	return nil
}

func (m *Attempt) GetResponseTime() *timestamp.Timestamp {
	if m != nil {
		return m.ResponseTime
	}
	return nil
}

func (m *Attempt) GetResponseStatus() *status.Status {
	if m != nil {
		return m.ResponseStatus
	}
	return nil
}

func init() {
	proto.RegisterEnum("google.cloud.tasks.v2.Task_View", Task_View_name, Task_View_value)
	proto.RegisterType((*Task)(nil), "google.cloud.tasks.v2.Task")
	proto.RegisterType((*Attempt)(nil), "google.cloud.tasks.v2.Attempt")
}

func init() {
	proto.RegisterFile("google/cloud/tasks/v2/task.proto", fileDescriptor_5d55198283613840)
}

var fileDescriptor_5d55198283613840 = []byte{
	// 655 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xdb, 0x6a, 0xdb, 0x4c,
	0x10, 0xc7, 0x23, 0xc7, 0x39, 0x78, 0xed, 0xe4, 0xf3, 0x27, 0x52, 0xa2, 0x84, 0x92, 0x1a, 0x43,
	0xc1, 0x50, 0xd0, 0x52, 0xb7, 0x57, 0xc9, 0x45, 0xb0, 0x1d, 0xa7, 0x31, 0x84, 0x12, 0x94, 0x43,
	0xa1, 0x37, 0x62, 0x23, 0x4d, 0x64, 0x35, 0x92, 0x76, 0xa3, 0x5d, 0x25, 0x14, 0xe3, 0x07, 0xe8,
	0x9b, 0xf5, 0x59, 0xfa, 0x14, 0x45, 0x7b, 0x70, 0x73, 0x32, 0x69, 0xaf, 0x34, 0x3b, 0xf3, 0xff,
	0xff, 0xb4, 0xda, 0x99, 0x15, 0x6a, 0x45, 0x94, 0x46, 0x09, 0xe0, 0x20, 0xa1, 0x45, 0x88, 0x05,
	0xe1, 0xd7, 0x1c, 0xdf, 0x76, 0x65, 0xe0, 0xb2, 0x9c, 0x0a, 0x6a, 0xbf, 0x52, 0x0a, 0x57, 0x2a,
	0x5c, 0xa9, 0x70, 0x6f, 0xbb, 0xdb, 0x5b, 0xda, 0x48, 0x58, 0x8c, 0x73, 0xe0, 0xb4, 0xc8, 0x03,
	0x50, 0x8e, 0xed, 0xf6, 0x3c, 0x66, 0x1e, 0x81, 0xd0, 0x9a, 0x1d, 0xad, 0x91, 0xab, 0xcb, 0xe2,
	0x0a, 0x87, 0x45, 0x4e, 0x44, 0x4c, 0x33, 0x5d, 0x7f, 0xf3, 0xb8, 0x2e, 0xe2, 0x14, 0xb8, 0x20,
	0x29, 0xd3, 0x82, 0x4d, 0x2d, 0xc8, 0x59, 0x80, 0xb9, 0x20, 0xa2, 0xe0, 0xba, 0xf0, 0xfa, 0xde,
	0xc6, 0x48, 0x96, 0x51, 0x21, 0xb1, 0xba, 0xda, 0xfe, 0xb9, 0x8c, 0xaa, 0x67, 0x84, 0x5f, 0xdb,
	0x36, 0xaa, 0x66, 0x24, 0x05, 0xc7, 0x6a, 0x59, 0x9d, 0x9a, 0x27, 0x63, 0x3b, 0x44, 0x9b, 0x84,
	0x31, 0x1f, 0xb2, 0x28, 0xce, 0xc0, 0x1f, 0x0b, 0xc1, 0xfc, 0x1c, 0x6e, 0x0a, 0xe0, 0xc2, 0xa9,
	0xb4, 0xac, 0x4e, 0xbd, 0xfb, 0xce, 0x7d, 0xf6, 0x30, 0xdc, 0x1e, 0x63, 0x43, 0x69, 0x3a, 0x12,
	0x82, 0x79, 0xca, 0x72, 0xb4, 0xe0, 0x6d, 0x90, 0x67, 0xf2, 0xf6, 0x27, 0xd4, 0x78, 0x80, 0x5e,
	0x94, 0xe8, 0xf6, 0x1c, 0xf4, 0x43, 0x62, 0x7d, 0x7c, 0x0f, 0xb4, 0x8f, 0xd6, 0x78, 0x30, 0x86,
	0xb0, 0x48, 0xc0, 0x2f, 0x8f, 0xc7, 0xa9, 0x4a, 0xd2, 0xb6, 0x21, 0x99, 0xb3, 0x73, 0xcf, 0xcc,
	0xd9, 0x79, 0x0d, 0x63, 0x28, 0x53, 0xf6, 0x1e, 0xaa, 0x07, 0x39, 0x10, 0xa1, 0xed, 0x4b, 0x2f,
	0xda, 0x91, 0x92, 0x4b, 0xf3, 0x21, 0xfa, 0x3f, 0x8c, 0x39, 0x23, 0x22, 0x18, 0xfb, 0x21, 0x90,
	0x30, 0x89, 0x33, 0x70, 0x96, 0x25, 0x62, 0xeb, 0x09, 0xe2, 0x40, 0x77, 0xd7, 0x6b, 0x1a, 0xcf,
	0x81, 0xb6, 0xd8, 0x6f, 0xd1, 0xfa, 0x8c, 0x13, 0xd0, 0x22, 0x13, 0xce, 0x4a, 0xcb, 0xea, 0x2c,
	0x79, 0x6b, 0x26, 0x3b, 0x28, 0x93, 0xa5, 0x2c, 0x07, 0xce, 0x68, 0xc6, 0x41, 0xcb, 0x56, 0x95,
	0xcc, 0x64, 0x95, 0x6c, 0x80, 0xd6, 0xae, 0xe2, 0x9c, 0x0b, 0x9f, 0x08, 0x01, 0x29, 0x13, 0x4e,
	0x4d, 0xee, 0x68, 0x67, 0x5e, 0xe3, 0x94, 0xca, 0x6b, 0x48, 0x93, 0x5e, 0xd9, 0x3d, 0xd4, 0x48,
	0xc8, 0x3d, 0x06, 0xfa, 0x2b, 0x46, 0xbd, 0xf4, 0x18, 0xc4, 0x47, 0x54, 0xbd, 0x8d, 0xe1, 0xce,
	0xa9, 0xb7, 0xac, 0xce, 0x7a, 0xb7, 0x35, 0xc7, 0x5a, 0x4e, 0xa2, 0x7b, 0x11, 0xc3, 0x9d, 0x27,
	0xd5, 0xed, 0xf7, 0xa8, 0x5a, 0xae, 0xec, 0x0d, 0xd4, 0xbc, 0x18, 0x0d, 0xbf, 0xf8, 0xe7, 0x9f,
	0x4f, 0x4f, 0x86, 0x83, 0xd1, 0xe1, 0x68, 0x78, 0xd0, 0x5c, 0xb0, 0x6b, 0x68, 0xa9, 0xdf, 0x3b,
	0x1d, 0x0d, 0x9a, 0x96, 0xbd, 0x8a, 0xaa, 0x87, 0xe7, 0xc7, 0xc7, 0xcd, 0xca, 0xee, 0xf8, 0x57,
	0x0f, 0xd0, 0x8e, 0xe4, 0x2a, 0xac, 0x7a, 0x13, 0x61, 0x31, 0x77, 0x03, 0x9a, 0x62, 0x39, 0xec,
	0x03, 0x96, 0xd3, 0x6f, 0x10, 0x08, 0x8e, 0x27, 0x3a, 0x9a, 0xe2, 0x84, 0x06, 0xea, 0x66, 0xe0,
	0x89, 0x09, 0xa7, 0xf8, 0xa6, 0x80, 0x02, 0x38, 0x9e, 0xc8, 0xe7, 0x54, 0xdf, 0xdd, 0x49, 0xf9,
	0x98, 0xf6, 0xd7, 0x51, 0x23, 0x05, 0xce, 0x49, 0x04, 0xbe, 0xf8, 0xce, 0xa0, 0xfd, 0xa3, 0x82,
	0x56, 0xcc, 0xe7, 0x3e, 0x19, 0x45, 0xeb, 0x1f, 0x47, 0x71, 0x1f, 0xcd, 0xfa, 0xad, 0x00, 0x95,
	0x97, 0x01, 0xc6, 0x60, 0x00, 0xb3, 0xf9, 0x90, 0x80, 0xc5, 0x97, 0x01, 0xc6, 0xa0, 0x2f, 0xc3,
	0x7f, 0x33, 0x80, 0xfa, 0xa1, 0xe8, 0xfb, 0x64, 0x1b, 0x44, 0xce, 0x02, 0xf7, 0x54, 0x56, 0xbc,
	0xd9, 0x2c, 0xaa, 0x75, 0x3f, 0x44, 0x5b, 0x01, 0x4d, 0x9f, 0xef, 0x72, 0xbf, 0x56, 0xf6, 0xe0,
	0xa4, 0x7c, 0xff, 0x89, 0xf5, 0x75, 0x57, 0x6b, 0x22, 0x9a, 0x90, 0x2c, 0x72, 0x69, 0x1e, 0xe1,
	0x08, 0x32, 0xb9, 0x3b, 0xfc, 0xa7, 0x75, 0x8f, 0xfe, 0x9d, 0x7b, 0x32, 0xb8, 0x5c, 0x96, 0xb2,
	0x0f, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xed, 0x7e, 0x48, 0x17, 0xb5, 0x05, 0x00, 0x00,
}
