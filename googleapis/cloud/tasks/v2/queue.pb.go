// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/tasks/v2/queue.proto

package tasks

import (
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
	duration "github.com/catper/protobuf/ptypes/duration"
	timestamp "github.com/catper/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// State of the queue.
type Queue_State int32

const (
	// Unspecified state.
	Queue_STATE_UNSPECIFIED Queue_State = 0
	// The queue is running. Tasks can be dispatched.
	//
	// If the queue was created using Cloud Tasks and the queue has
	// had no activity (method calls or task dispatches) for 30 days,
	// the queue may take a few minutes to re-activate. Some method
	// calls may return [NOT_FOUND][google.rpc.Code.NOT_FOUND] and
	// tasks may not be dispatched for a few minutes until the queue
	// has been re-activated.
	Queue_RUNNING Queue_State = 1
	// Tasks are paused by the user. If the queue is paused then Cloud
	// Tasks will stop delivering tasks from it, but more tasks can
	// still be added to it by the user.
	Queue_PAUSED Queue_State = 2
	// The queue is disabled.
	//
	// A queue becomes `DISABLED` when
	// [queue.yaml](https://cloud.google.com/appengine/docs/python/config/queueref)
	// or
	// [queue.xml](https://cloud.google.com/appengine/docs/standard/java/config/queueref)
	// is uploaded which does not contain the queue. You cannot directly disable
	// a queue.
	//
	// When a queue is disabled, tasks can still be added to a queue
	// but the tasks are not dispatched.
	//
	// To permanently delete this queue and all of its tasks, call
	// [DeleteQueue][google.cloud.tasks.v2.CloudTasks.DeleteQueue].
	Queue_DISABLED Queue_State = 3
)

var Queue_State_name = map[int32]string{
	0: "STATE_UNSPECIFIED",
	1: "RUNNING",
	2: "PAUSED",
	3: "DISABLED",
}

var Queue_State_value = map[string]int32{
	"STATE_UNSPECIFIED": 0,
	"RUNNING":           1,
	"PAUSED":            2,
	"DISABLED":          3,
}

func (x Queue_State) String() string {
	return proto.EnumName(Queue_State_name, int32(x))
}

func (Queue_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a4a1833e2495b95c, []int{0, 0}
}

// A queue is a container of related tasks. Queues are configured to manage
// how those tasks are dispatched. Configurable properties include rate limits,
// retry options, queue types, and others.
type Queue struct {
	// Caller-specified and required in [CreateQueue][google.cloud.tasks.v2.CloudTasks.CreateQueue],
	// after which it becomes output only.
	//
	// The queue name.
	//
	// The queue name must have the following format:
	// `projects/PROJECT_ID/locations/LOCATION_ID/queues/QUEUE_ID`
	//
	// * `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]),
	//    hyphens (-), colons (:), or periods (.).
	//    For more information, see
	//    [Identifying
	//    projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects)
	// * `LOCATION_ID` is the canonical ID for the queue's location.
	//    The list of available locations can be obtained by calling
	//    [ListLocations][google.cloud.location.Locations.ListLocations].
	//    For more information, see https://cloud.google.com/about/locations/.
	// * `QUEUE_ID` can contain letters ([A-Za-z]), numbers ([0-9]), or
	//   hyphens (-). The maximum length is 100 characters.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Overrides for
	// [task-level app_engine_routing][google.cloud.tasks.v2.AppEngineHttpRequest.app_engine_routing].
	// These settings apply only to
	// [App Engine tasks][google.cloud.tasks.v2.AppEngineHttpRequest] in this queue.
	// [Http tasks][google.cloud.tasks.v2.HttpRequest] are not affected.
	//
	// If set, `app_engine_routing_override` is used for all
	// [App Engine tasks][google.cloud.tasks.v2.AppEngineHttpRequest] in the queue, no matter what the
	// setting is for the
	// [task-level app_engine_routing][google.cloud.tasks.v2.AppEngineHttpRequest.app_engine_routing].
	AppEngineRoutingOverride *AppEngineRouting `protobuf:"bytes,2,opt,name=app_engine_routing_override,json=appEngineRoutingOverride,proto3" json:"app_engine_routing_override,omitempty"`
	// Rate limits for task dispatches.
	//
	// [rate_limits][google.cloud.tasks.v2.Queue.rate_limits] and [retry_config][google.cloud.tasks.v2.Queue.retry_config] are
	// related because they both control task attempts. However they control task
	// attempts in different ways:
	//
	// * [rate_limits][google.cloud.tasks.v2.Queue.rate_limits] controls the total rate of
	//   dispatches from a queue (i.e. all traffic dispatched from the
	//   queue, regardless of whether the dispatch is from a first
	//   attempt or a retry).
	// * [retry_config][google.cloud.tasks.v2.Queue.retry_config] controls what happens to
	//   particular a task after its first attempt fails. That is,
	//   [retry_config][google.cloud.tasks.v2.Queue.retry_config] controls task retries (the
	//   second attempt, third attempt, etc).
	//
	// The queue's actual dispatch rate is the result of:
	//
	// * Number of tasks in the queue
	// * User-specified throttling: [rate_limits][google.cloud.tasks.v2.Queue.rate_limits],
	//   [retry_config][google.cloud.tasks.v2.Queue.retry_config], and the
	//   [queue's state][google.cloud.tasks.v2.Queue.state].
	// * System throttling due to `429` (Too Many Requests) or `503` (Service
	//   Unavailable) responses from the worker, high error rates, or to smooth
	//   sudden large traffic spikes.
	RateLimits *RateLimits `protobuf:"bytes,3,opt,name=rate_limits,json=rateLimits,proto3" json:"rate_limits,omitempty"`
	// Settings that determine the retry behavior.
	//
	// * For tasks created using Cloud Tasks: the queue-level retry settings
	//   apply to all tasks in the queue that were created using Cloud Tasks.
	//   Retry settings cannot be set on individual tasks.
	// * For tasks created using the App Engine SDK: the queue-level retry
	//   settings apply to all tasks in the queue which do not have retry settings
	//   explicitly set on the task and were created by the App Engine SDK. See
	//   [App Engine
	//   documentation](https://cloud.google.com/appengine/docs/standard/python/taskqueue/push/retrying-tasks).
	RetryConfig *RetryConfig `protobuf:"bytes,4,opt,name=retry_config,json=retryConfig,proto3" json:"retry_config,omitempty"`
	// Output only. The state of the queue.
	//
	// `state` can only be changed by called
	// [PauseQueue][google.cloud.tasks.v2.CloudTasks.PauseQueue],
	// [ResumeQueue][google.cloud.tasks.v2.CloudTasks.ResumeQueue], or uploading
	// [queue.yaml/xml](https://cloud.google.com/appengine/docs/python/config/queueref).
	// [UpdateQueue][google.cloud.tasks.v2.CloudTasks.UpdateQueue] cannot be used to change `state`.
	State Queue_State `protobuf:"varint,5,opt,name=state,proto3,enum=google.cloud.tasks.v2.Queue_State" json:"state,omitempty"`
	// Output only. The last time this queue was purged.
	//
	// All tasks that were [created][google.cloud.tasks.v2.Task.create_time] before this time
	// were purged.
	//
	// A queue can be purged using [PurgeQueue][google.cloud.tasks.v2.CloudTasks.PurgeQueue], the
	// [App Engine Task Queue SDK, or the Cloud
	// Console](https://cloud.google.com/appengine/docs/standard/python/taskqueue/push/deleting-tasks-and-queues#purging_all_tasks_from_a_queue).
	//
	// Purge time will be truncated to the nearest microsecond. Purge
	// time will be unset if the queue has never been purged.
	PurgeTime *timestamp.Timestamp `protobuf:"bytes,6,opt,name=purge_time,json=purgeTime,proto3" json:"purge_time,omitempty"`
	// Configuration options for writing logs to
	// [Stackdriver Logging](https://cloud.google.com/logging/docs/). If this
	// field is unset, then no logs are written.
	StackdriverLoggingConfig *StackdriverLoggingConfig `protobuf:"bytes,9,opt,name=stackdriver_logging_config,json=stackdriverLoggingConfig,proto3" json:"stackdriver_logging_config,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                  `json:"-"`
	XXX_unrecognized         []byte                    `json:"-"`
	XXX_sizecache            int32                     `json:"-"`
}

func (m *Queue) Reset()         { *m = Queue{} }
func (m *Queue) String() string { return proto.CompactTextString(m) }
func (*Queue) ProtoMessage()    {}
func (*Queue) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4a1833e2495b95c, []int{0}
}

func (m *Queue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Queue.Unmarshal(m, b)
}
func (m *Queue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Queue.Marshal(b, m, deterministic)
}
func (m *Queue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Queue.Merge(m, src)
}
func (m *Queue) XXX_Size() int {
	return xxx_messageInfo_Queue.Size(m)
}
func (m *Queue) XXX_DiscardUnknown() {
	xxx_messageInfo_Queue.DiscardUnknown(m)
}

var xxx_messageInfo_Queue proto.InternalMessageInfo

func (m *Queue) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Queue) GetAppEngineRoutingOverride() *AppEngineRouting {
	if m != nil {
		return m.AppEngineRoutingOverride
	}
	return nil
}

func (m *Queue) GetRateLimits() *RateLimits {
	if m != nil {
		return m.RateLimits
	}
	return nil
}

func (m *Queue) GetRetryConfig() *RetryConfig {
	if m != nil {
		return m.RetryConfig
	}
	return nil
}

func (m *Queue) GetState() Queue_State {
	if m != nil {
		return m.State
	}
	return Queue_STATE_UNSPECIFIED
}

func (m *Queue) GetPurgeTime() *timestamp.Timestamp {
	if m != nil {
		return m.PurgeTime
	}
	return nil
}

func (m *Queue) GetStackdriverLoggingConfig() *StackdriverLoggingConfig {
	if m != nil {
		return m.StackdriverLoggingConfig
	}
	return nil
}

// Rate limits.
//
// This message determines the maximum rate that tasks can be dispatched by a
// queue, regardless of whether the dispatch is a first task attempt or a retry.
//
// Note: The debugging command, [RunTask][google.cloud.tasks.v2.CloudTasks.RunTask], will run a task
// even if the queue has reached its [RateLimits][google.cloud.tasks.v2.RateLimits].
type RateLimits struct {
	// The maximum rate at which tasks are dispatched from this queue.
	//
	// If unspecified when the queue is created, Cloud Tasks will pick the
	// default.
	//
	// * The maximum allowed value is 500.
	//
	//
	// This field has the same meaning as
	// [rate in
	// queue.yaml/xml](https://cloud.google.com/appengine/docs/standard/python/config/queueref#rate).
	MaxDispatchesPerSecond float64 `protobuf:"fixed64,1,opt,name=max_dispatches_per_second,json=maxDispatchesPerSecond,proto3" json:"max_dispatches_per_second,omitempty"`
	// Output only. The max burst size.
	//
	// Max burst size limits how fast tasks in queue are processed when
	// many tasks are in the queue and the rate is high. This field
	// allows the queue to have a high rate so processing starts shortly
	// after a task is enqueued, but still limits resource usage when
	// many tasks are enqueued in a short period of time.
	//
	// The [token bucket](https://wikipedia.org/wiki/Token_Bucket)
	// algorithm is used to control the rate of task dispatches. Each
	// queue has a token bucket that holds tokens, up to the maximum
	// specified by `max_burst_size`. Each time a task is dispatched, a
	// token is removed from the bucket. Tasks will be dispatched until
	// the queue's bucket runs out of tokens. The bucket will be
	// continuously refilled with new tokens based on
	// [max_dispatches_per_second][google.cloud.tasks.v2.RateLimits.max_dispatches_per_second].
	//
	// Cloud Tasks will pick the value of `max_burst_size` based on the
	// value of
	// [max_dispatches_per_second][google.cloud.tasks.v2.RateLimits.max_dispatches_per_second].
	//
	// For queues that were created or updated using
	// `queue.yaml/xml`, `max_burst_size` is equal to
	// [bucket_size](https://cloud.google.com/appengine/docs/standard/python/config/queueref#bucket_size).
	// Since `max_burst_size` is output only, if
	// [UpdateQueue][google.cloud.tasks.v2.CloudTasks.UpdateQueue] is called on a queue
	// created by `queue.yaml/xml`, `max_burst_size` will be reset based
	// on the value of
	// [max_dispatches_per_second][google.cloud.tasks.v2.RateLimits.max_dispatches_per_second],
	// regardless of whether
	// [max_dispatches_per_second][google.cloud.tasks.v2.RateLimits.max_dispatches_per_second]
	// is updated.
	//
	MaxBurstSize int32 `protobuf:"varint,2,opt,name=max_burst_size,json=maxBurstSize,proto3" json:"max_burst_size,omitempty"`
	// The maximum number of concurrent tasks that Cloud Tasks allows
	// to be dispatched for this queue. After this threshold has been
	// reached, Cloud Tasks stops dispatching tasks until the number of
	// concurrent requests decreases.
	//
	// If unspecified when the queue is created, Cloud Tasks will pick the
	// default.
	//
	//
	// The maximum allowed value is 5,000.
	//
	//
	// This field has the same meaning as
	// [max_concurrent_requests in
	// queue.yaml/xml](https://cloud.google.com/appengine/docs/standard/python/config/queueref#max_concurrent_requests).
	MaxConcurrentDispatches int32    `protobuf:"varint,3,opt,name=max_concurrent_dispatches,json=maxConcurrentDispatches,proto3" json:"max_concurrent_dispatches,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *RateLimits) Reset()         { *m = RateLimits{} }
func (m *RateLimits) String() string { return proto.CompactTextString(m) }
func (*RateLimits) ProtoMessage()    {}
func (*RateLimits) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4a1833e2495b95c, []int{1}
}

func (m *RateLimits) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimits.Unmarshal(m, b)
}
func (m *RateLimits) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimits.Marshal(b, m, deterministic)
}
func (m *RateLimits) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimits.Merge(m, src)
}
func (m *RateLimits) XXX_Size() int {
	return xxx_messageInfo_RateLimits.Size(m)
}
func (m *RateLimits) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimits.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimits proto.InternalMessageInfo

func (m *RateLimits) GetMaxDispatchesPerSecond() float64 {
	if m != nil {
		return m.MaxDispatchesPerSecond
	}
	return 0
}

func (m *RateLimits) GetMaxBurstSize() int32 {
	if m != nil {
		return m.MaxBurstSize
	}
	return 0
}

func (m *RateLimits) GetMaxConcurrentDispatches() int32 {
	if m != nil {
		return m.MaxConcurrentDispatches
	}
	return 0
}

// Retry config.
//
// These settings determine when a failed task attempt is retried.
type RetryConfig struct {
	// Number of attempts per task.
	//
	// Cloud Tasks will attempt the task `max_attempts` times (that is, if the
	// first attempt fails, then there will be `max_attempts - 1` retries). Must
	// be >= -1.
	//
	// If unspecified when the queue is created, Cloud Tasks will pick the
	// default.
	//
	// -1 indicates unlimited attempts.
	//
	// This field has the same meaning as
	// [task_retry_limit in
	// queue.yaml/xml](https://cloud.google.com/appengine/docs/standard/python/config/queueref#retry_parameters).
	MaxAttempts int32 `protobuf:"varint,1,opt,name=max_attempts,json=maxAttempts,proto3" json:"max_attempts,omitempty"`
	// If positive, `max_retry_duration` specifies the time limit for
	// retrying a failed task, measured from when the task was first
	// attempted. Once `max_retry_duration` time has passed *and* the
	// task has been attempted [max_attempts][google.cloud.tasks.v2.RetryConfig.max_attempts]
	// times, no further attempts will be made and the task will be
	// deleted.
	//
	// If zero, then the task age is unlimited.
	//
	// If unspecified when the queue is created, Cloud Tasks will pick the
	// default.
	//
	//
	// `max_retry_duration` will be truncated to the nearest second.
	//
	// This field has the same meaning as
	// [task_age_limit in
	// queue.yaml/xml](https://cloud.google.com/appengine/docs/standard/python/config/queueref#retry_parameters).
	MaxRetryDuration *duration.Duration `protobuf:"bytes,2,opt,name=max_retry_duration,json=maxRetryDuration,proto3" json:"max_retry_duration,omitempty"`
	// A task will be [scheduled][google.cloud.tasks.v2.Task.schedule_time] for retry between
	// [min_backoff][google.cloud.tasks.v2.RetryConfig.min_backoff] and
	// [max_backoff][google.cloud.tasks.v2.RetryConfig.max_backoff] duration after it fails,
	// if the queue's [RetryConfig][google.cloud.tasks.v2.RetryConfig] specifies that the task should be
	// retried.
	//
	// If unspecified when the queue is created, Cloud Tasks will pick the
	// default.
	//
	//
	// `min_backoff` will be truncated to the nearest second.
	//
	// This field has the same meaning as
	// [min_backoff_seconds in
	// queue.yaml/xml](https://cloud.google.com/appengine/docs/standard/python/config/queueref#retry_parameters).
	MinBackoff *duration.Duration `protobuf:"bytes,3,opt,name=min_backoff,json=minBackoff,proto3" json:"min_backoff,omitempty"`
	// A task will be [scheduled][google.cloud.tasks.v2.Task.schedule_time] for retry between
	// [min_backoff][google.cloud.tasks.v2.RetryConfig.min_backoff] and
	// [max_backoff][google.cloud.tasks.v2.RetryConfig.max_backoff] duration after it fails,
	// if the queue's [RetryConfig][google.cloud.tasks.v2.RetryConfig] specifies that the task should be
	// retried.
	//
	// If unspecified when the queue is created, Cloud Tasks will pick the
	// default.
	//
	//
	// `max_backoff` will be truncated to the nearest second.
	//
	// This field has the same meaning as
	// [max_backoff_seconds in
	// queue.yaml/xml](https://cloud.google.com/appengine/docs/standard/python/config/queueref#retry_parameters).
	MaxBackoff *duration.Duration `protobuf:"bytes,4,opt,name=max_backoff,json=maxBackoff,proto3" json:"max_backoff,omitempty"`
	// The time between retries will double `max_doublings` times.
	//
	// A task's retry interval starts at
	// [min_backoff][google.cloud.tasks.v2.RetryConfig.min_backoff], then doubles
	// `max_doublings` times, then increases linearly, and finally
	// retries retries at intervals of
	// [max_backoff][google.cloud.tasks.v2.RetryConfig.max_backoff] up to
	// [max_attempts][google.cloud.tasks.v2.RetryConfig.max_attempts] times.
	//
	// For example, if [min_backoff][google.cloud.tasks.v2.RetryConfig.min_backoff] is 10s,
	// [max_backoff][google.cloud.tasks.v2.RetryConfig.max_backoff] is 300s, and
	// `max_doublings` is 3, then the a task will first be retried in
	// 10s. The retry interval will double three times, and then
	// increase linearly by 2^3 * 10s.  Finally, the task will retry at
	// intervals of [max_backoff][google.cloud.tasks.v2.RetryConfig.max_backoff] until the
	// task has been attempted [max_attempts][google.cloud.tasks.v2.RetryConfig.max_attempts]
	// times. Thus, the requests will retry at 10s, 20s, 40s, 80s, 160s,
	// 240s, 300s, 300s, ....
	//
	// If unspecified when the queue is created, Cloud Tasks will pick the
	// default.
	//
	//
	// This field has the same meaning as
	// [max_doublings in
	// queue.yaml/xml](https://cloud.google.com/appengine/docs/standard/python/config/queueref#retry_parameters).
	MaxDoublings         int32    `protobuf:"varint,5,opt,name=max_doublings,json=maxDoublings,proto3" json:"max_doublings,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RetryConfig) Reset()         { *m = RetryConfig{} }
func (m *RetryConfig) String() string { return proto.CompactTextString(m) }
func (*RetryConfig) ProtoMessage()    {}
func (*RetryConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4a1833e2495b95c, []int{2}
}

func (m *RetryConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetryConfig.Unmarshal(m, b)
}
func (m *RetryConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetryConfig.Marshal(b, m, deterministic)
}
func (m *RetryConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetryConfig.Merge(m, src)
}
func (m *RetryConfig) XXX_Size() int {
	return xxx_messageInfo_RetryConfig.Size(m)
}
func (m *RetryConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RetryConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RetryConfig proto.InternalMessageInfo

func (m *RetryConfig) GetMaxAttempts() int32 {
	if m != nil {
		return m.MaxAttempts
	}
	return 0
}

func (m *RetryConfig) GetMaxRetryDuration() *duration.Duration {
	if m != nil {
		return m.MaxRetryDuration
	}
	return nil
}

func (m *RetryConfig) GetMinBackoff() *duration.Duration {
	if m != nil {
		return m.MinBackoff
	}
	return nil
}

func (m *RetryConfig) GetMaxBackoff() *duration.Duration {
	if m != nil {
		return m.MaxBackoff
	}
	return nil
}

func (m *RetryConfig) GetMaxDoublings() int32 {
	if m != nil {
		return m.MaxDoublings
	}
	return 0
}

// Configuration options for writing logs to
// [Stackdriver Logging](https://cloud.google.com/logging/docs/).
type StackdriverLoggingConfig struct {
	// Specifies the fraction of operations to write to
	// [Stackdriver Logging](https://cloud.google.com/logging/docs/).
	// This field may contain any value between 0.0 and 1.0, inclusive.
	// 0.0 is the default and means that no operations are logged.
	SamplingRatio        float64  `protobuf:"fixed64,1,opt,name=sampling_ratio,json=samplingRatio,proto3" json:"sampling_ratio,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StackdriverLoggingConfig) Reset()         { *m = StackdriverLoggingConfig{} }
func (m *StackdriverLoggingConfig) String() string { return proto.CompactTextString(m) }
func (*StackdriverLoggingConfig) ProtoMessage()    {}
func (*StackdriverLoggingConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4a1833e2495b95c, []int{3}
}

func (m *StackdriverLoggingConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StackdriverLoggingConfig.Unmarshal(m, b)
}
func (m *StackdriverLoggingConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StackdriverLoggingConfig.Marshal(b, m, deterministic)
}
func (m *StackdriverLoggingConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StackdriverLoggingConfig.Merge(m, src)
}
func (m *StackdriverLoggingConfig) XXX_Size() int {
	return xxx_messageInfo_StackdriverLoggingConfig.Size(m)
}
func (m *StackdriverLoggingConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_StackdriverLoggingConfig.DiscardUnknown(m)
}

var xxx_messageInfo_StackdriverLoggingConfig proto.InternalMessageInfo

func (m *StackdriverLoggingConfig) GetSamplingRatio() float64 {
	if m != nil {
		return m.SamplingRatio
	}
	return 0
}

func init() {
	proto.RegisterEnum("google.cloud.tasks.v2.Queue_State", Queue_State_name, Queue_State_value)
	proto.RegisterType((*Queue)(nil), "google.cloud.tasks.v2.Queue")
	proto.RegisterType((*RateLimits)(nil), "google.cloud.tasks.v2.RateLimits")
	proto.RegisterType((*RetryConfig)(nil), "google.cloud.tasks.v2.RetryConfig")
	proto.RegisterType((*StackdriverLoggingConfig)(nil), "google.cloud.tasks.v2.StackdriverLoggingConfig")
}

func init() {
	proto.RegisterFile("google/cloud/tasks/v2/queue.proto", fileDescriptor_a4a1833e2495b95c)
}

var fileDescriptor_a4a1833e2495b95c = []byte{
	// 755 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xd1, 0x4e, 0xe3, 0x46,
	0x14, 0xad, 0x01, 0xd3, 0x72, 0x03, 0x28, 0x1d, 0x89, 0xd6, 0x49, 0xab, 0x02, 0x69, 0xab, 0xf2,
	0x64, 0x4b, 0x54, 0xaa, 0x4a, 0xfa, 0x94, 0x10, 0x17, 0x45, 0x42, 0x69, 0x6a, 0xc3, 0x43, 0xab,
	0x95, 0xac, 0x89, 0x33, 0xf1, 0xce, 0x92, 0x99, 0xf1, 0xce, 0x8c, 0xa3, 0x2c, 0x88, 0xdf, 0xd8,
	0xfd, 0x86, 0xfd, 0xb5, 0xfd, 0x8a, 0x95, 0xc7, 0x63, 0x82, 0x58, 0xb2, 0xfb, 0x94, 0xf1, 0xbd,
	0xe7, 0x9c, 0x7b, 0x67, 0xee, 0xb9, 0x81, 0xe3, 0x4c, 0x88, 0x6c, 0x4e, 0x82, 0x74, 0x2e, 0x8a,
	0x69, 0xa0, 0xb1, 0xba, 0x51, 0xc1, 0xe2, 0x34, 0x78, 0x5d, 0x90, 0x82, 0xf8, 0xb9, 0x14, 0x5a,
	0xa0, 0x83, 0x0a, 0xe2, 0x1b, 0x88, 0x6f, 0x20, 0xfe, 0xe2, 0xb4, 0xdd, 0xb2, 0x4c, 0x9c, 0xd3,
	0x40, 0x12, 0x25, 0x0a, 0x99, 0x5a, 0x46, 0xbb, 0xf3, 0xbc, 0xa8, 0xc6, 0x32, 0x23, 0xda, 0x62,
	0x7e, 0xb2, 0x18, 0xf3, 0x35, 0x29, 0x66, 0xc1, 0xb4, 0x90, 0x58, 0x53, 0xc1, 0x6d, 0xfe, 0xf0,
	0x69, 0x5e, 0x53, 0x46, 0x94, 0xc6, 0x2c, 0xb7, 0x80, 0x1f, 0x1f, 0xd5, 0xc7, 0x9c, 0x0b, 0x6d,
	0xd8, 0xaa, 0xca, 0x76, 0xde, 0xb9, 0xe0, 0xfe, 0x5b, 0x5e, 0x02, 0x21, 0xd8, 0xe2, 0x98, 0x11,
	0xcf, 0x39, 0x72, 0x4e, 0x76, 0x22, 0x73, 0x46, 0x33, 0xf8, 0x01, 0xe7, 0x79, 0x42, 0x78, 0x46,
	0x39, 0x49, 0xa4, 0x28, 0x34, 0xe5, 0x59, 0x22, 0x16, 0x44, 0x4a, 0x3a, 0x25, 0xde, 0xc6, 0x91,
	0x73, 0xd2, 0x38, 0xfd, 0xcd, 0x7f, 0xf6, 0xe2, 0x7e, 0x2f, 0xcf, 0x43, 0x43, 0x8c, 0x2a, 0x5e,
	0xe4, 0xe1, 0x27, 0x91, 0x7f, 0xac, 0x10, 0xea, 0x43, 0x43, 0x62, 0x4d, 0x92, 0x39, 0x65, 0x54,
	0x2b, 0x6f, 0xd3, 0xe8, 0x1e, 0xaf, 0xd1, 0x8d, 0xb0, 0x26, 0x97, 0x06, 0x18, 0x81, 0x7c, 0x38,
	0xa3, 0x10, 0x76, 0x25, 0xd1, 0xf2, 0x4d, 0x92, 0x0a, 0x3e, 0xa3, 0x99, 0xb7, 0x65, 0x44, 0x3a,
	0xeb, 0x44, 0x4a, 0xe8, 0xb9, 0x41, 0x46, 0x0d, 0xb9, 0xfa, 0x40, 0x7f, 0x82, 0xab, 0x34, 0xd6,
	0xc4, 0x73, 0x8f, 0x9c, 0x93, 0xfd, 0xb5, 0x7c, 0xf3, 0x66, 0x7e, 0x5c, 0x22, 0xa3, 0x8a, 0x80,
	0xce, 0x00, 0xf2, 0x42, 0x66, 0x24, 0x29, 0x27, 0xe0, 0x6d, 0x9b, 0xf2, 0xed, 0x9a, 0x5e, 0x8f,
	0xc7, 0xbf, 0xaa, 0xc7, 0x13, 0xed, 0x18, 0x74, 0xf9, 0x8d, 0x18, 0xb4, 0x95, 0xc6, 0xe9, 0xcd,
	0x54, 0xd2, 0x05, 0x91, 0xc9, 0x5c, 0x64, 0x59, 0xf9, 0xd0, 0xf6, 0x26, 0x3b, 0x46, 0x2a, 0x58,
	0xd3, 0x49, 0xbc, 0x22, 0x5e, 0x56, 0x3c, 0x7b, 0x2d, 0x4f, 0xad, 0xc9, 0x74, 0x42, 0x70, 0x4d,
	0xe7, 0xe8, 0x00, 0xbe, 0x8d, 0xaf, 0x7a, 0x57, 0x61, 0x72, 0x3d, 0x8a, 0xc7, 0xe1, 0xf9, 0xf0,
	0xef, 0x61, 0x38, 0x68, 0x7e, 0x85, 0x1a, 0xf0, 0x75, 0x74, 0x3d, 0x1a, 0x0d, 0x47, 0x17, 0x4d,
	0x07, 0x01, 0x6c, 0x8f, 0x7b, 0xd7, 0x71, 0x38, 0x68, 0x6e, 0xa0, 0x5d, 0xf8, 0x66, 0x30, 0x8c,
	0x7b, 0xfd, 0xcb, 0x70, 0xd0, 0xdc, 0xec, 0xbe, 0xf8, 0xd0, 0xfb, 0x0f, 0x0e, 0x4d, 0x3b, 0x55,
	0x37, 0x55, 0x83, 0x38, 0xa7, 0xca, 0x4f, 0x05, 0x0b, 0x2a, 0x5f, 0xfd, 0x91, 0x4b, 0xf1, 0x8a,
	0xa4, 0x5a, 0x05, 0x77, 0xf6, 0x74, 0x1f, 0xcc, 0x45, 0x5a, 0xb9, 0x30, 0xb8, 0xab, 0x8f, 0xf7,
	0xd5, 0x2e, 0xa9, 0xe0, 0xce, 0xfc, 0xde, 0x77, 0xde, 0x3b, 0x00, 0xab, 0x51, 0xa3, 0x33, 0x68,
	0x31, 0xbc, 0x4c, 0xa6, 0x54, 0xe5, 0x58, 0xa7, 0x2f, 0x89, 0x4a, 0x72, 0x22, 0x13, 0x45, 0x52,
	0xc1, 0xa7, 0xc6, 0xb3, 0x4e, 0xf4, 0x1d, 0xc3, 0xcb, 0xc1, 0x43, 0x7e, 0x4c, 0x64, 0x6c, 0xb2,
	0xe8, 0x17, 0xd8, 0x2f, 0xa9, 0x93, 0x42, 0x2a, 0x9d, 0x28, 0x7a, 0x5b, 0x19, 0xd7, 0x8d, 0x76,
	0x19, 0x5e, 0xf6, 0xcb, 0x60, 0x4c, 0x6f, 0x09, 0xea, 0x56, 0x05, 0x52, 0xc1, 0xd3, 0x42, 0x4a,
	0xc2, 0xf5, 0xa3, 0x5a, 0xc6, 0x91, 0x6e, 0xf4, 0x3d, 0xc3, 0xcb, 0xf3, 0x87, 0xfc, 0xaa, 0x54,
	0xe7, 0xed, 0x06, 0x34, 0x1e, 0x39, 0x0a, 0x1d, 0x43, 0xa9, 0x9d, 0x60, 0xad, 0x09, 0xcb, 0xb5,
	0x32, 0xfd, 0xb9, 0x51, 0x83, 0xe1, 0x65, 0xcf, 0x86, 0xd0, 0x05, 0xa0, 0x12, 0x52, 0x59, 0xb6,
	0xde, 0x69, 0xbb, 0x51, 0xad, 0x4f, 0x5c, 0x33, 0xb0, 0x80, 0xa8, 0xc9, 0xf0, 0xd2, 0x54, 0xaa,
	0x23, 0xa8, 0x0b, 0x0d, 0x46, 0x79, 0x32, 0xc1, 0xe9, 0x8d, 0x98, 0xcd, 0xec, 0xee, 0x7c, 0x46,
	0x01, 0x18, 0xe5, 0xfd, 0x0a, 0x6c, 0xb8, 0xe5, 0xcb, 0x58, 0xee, 0xd6, 0x97, 0xb9, 0x78, 0x59,
	0x73, 0x7f, 0x86, 0x3d, 0x33, 0x10, 0x51, 0x4c, 0xe6, 0x94, 0x67, 0xca, 0x2c, 0x4c, 0xf5, 0xa8,
	0x83, 0x3a, 0xd6, 0xe9, 0x81, 0xb7, 0xce, 0x9f, 0xe8, 0x57, 0xd8, 0x57, 0x98, 0xe5, 0x25, 0x30,
	0x31, 0xfa, 0x76, 0x8c, 0x7b, 0x75, 0x34, 0x2a, 0x83, 0x7d, 0x02, 0xad, 0x54, 0xb0, 0xe7, 0xcd,
	0xdf, 0x07, 0xe3, 0xb1, 0x71, 0xd9, 0xe8, 0xd8, 0xf9, 0xbf, 0x6b, 0x41, 0x99, 0x98, 0x63, 0x9e,
	0xf9, 0x42, 0x66, 0x41, 0x46, 0xb8, 0xb9, 0x46, 0xb0, 0xf2, 0xe6, 0x93, 0xbf, 0xdb, 0xbf, 0xcc,
	0x61, 0xb2, 0x6d, 0x60, 0xbf, 0x7f, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x54, 0x5b, 0x4f, 0xe9,
	0x05, 0x00, 0x00,
}
