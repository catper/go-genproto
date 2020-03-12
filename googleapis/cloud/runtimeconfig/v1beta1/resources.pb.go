// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/runtimeconfig/v1beta1/resources.proto

package runtimeconfig

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

// The `VariableState` describes the last known state of the variable and is
// used during a `variables().watch` call to distinguish the state of the
// variable.
type VariableState int32

const (
	// Default variable state.
	VariableState_VARIABLE_STATE_UNSPECIFIED VariableState = 0
	// The variable was updated, while `variables().watch` was executing.
	VariableState_UPDATED VariableState = 1
	// The variable was deleted, while `variables().watch` was executing.
	VariableState_DELETED VariableState = 2
)

var VariableState_name = map[int32]string{
	0: "VARIABLE_STATE_UNSPECIFIED",
	1: "UPDATED",
	2: "DELETED",
}

var VariableState_value = map[string]int32{
	"VARIABLE_STATE_UNSPECIFIED": 0,
	"UPDATED":                    1,
	"DELETED":                    2,
}

func (x VariableState) String() string {
	return proto.EnumName(VariableState_name, int32(x))
}

func (VariableState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7913f3704a8d250c, []int{0}
}

// A RuntimeConfig resource is the primary resource in the Cloud RuntimeConfig
// service. A RuntimeConfig resource consists of metadata and a hierarchy of
// variables.
type RuntimeConfig struct {
	// The resource name of a runtime config. The name must have the format:
	//
	//     projects/[PROJECT_ID]/configs/[CONFIG_NAME]
	//
	// The `[PROJECT_ID]` must be a valid project ID, and `[CONFIG_NAME]` is an
	// arbitrary name that matches RFC 1035 segment specification. The length of
	// `[CONFIG_NAME]` must be less than 64 bytes.
	//
	// You pick the RuntimeConfig resource name, but the server will validate that
	// the name adheres to this format. After you create the resource, you cannot
	// change the resource's name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// An optional description of the RuntimeConfig object.
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RuntimeConfig) Reset()         { *m = RuntimeConfig{} }
func (m *RuntimeConfig) String() string { return proto.CompactTextString(m) }
func (*RuntimeConfig) ProtoMessage()    {}
func (*RuntimeConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_7913f3704a8d250c, []int{0}
}

func (m *RuntimeConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RuntimeConfig.Unmarshal(m, b)
}
func (m *RuntimeConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RuntimeConfig.Marshal(b, m, deterministic)
}
func (m *RuntimeConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RuntimeConfig.Merge(m, src)
}
func (m *RuntimeConfig) XXX_Size() int {
	return xxx_messageInfo_RuntimeConfig.Size(m)
}
func (m *RuntimeConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RuntimeConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RuntimeConfig proto.InternalMessageInfo

func (m *RuntimeConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RuntimeConfig) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// Describes a single variable within a RuntimeConfig resource.
// The name denotes the hierarchical variable name. For example,
// `ports/serving_port` is a valid variable name. The variable value is an
// opaque string and only leaf variables can have values (that is, variables
// that do not have any child variables).
type Variable struct {
	// The name of the variable resource, in the format:
	//
	//     projects/[PROJECT_ID]/configs/[CONFIG_NAME]/variables/[VARIABLE_NAME]
	//
	// The `[PROJECT_ID]` must be a valid project ID, `[CONFIG_NAME]` must be a
	// valid RuntimeConfig reource and `[VARIABLE_NAME]` follows Unix file system
	// file path naming.
	//
	// The `[VARIABLE_NAME]` can contain ASCII letters, numbers, slashes and
	// dashes. Slashes are used as path element separators and are not part of the
	// `[VARIABLE_NAME]` itself, so `[VARIABLE_NAME]` must contain at least one
	// non-slash character. Multiple slashes are coalesced into single slash
	// character. Each path segment should follow RFC 1035 segment specification.
	// The length of a `[VARIABLE_NAME]` must be less than 256 bytes.
	//
	// Once you create a variable, you cannot change the variable name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The value of the variable. It can be either a binary or a string
	// value. You must specify one of either `value` or `text`. Specifying both
	// will cause the server to return an error.
	//
	// Types that are valid to be assigned to Contents:
	//	*Variable_Value
	//	*Variable_Text
	Contents isVariable_Contents `protobuf_oneof:"contents"`
	// [Output Only] The time of the last variable update.
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// [Ouput only] The current state of the variable. The variable state
	// indicates the outcome of the `variables().watch` call and is visible
	// through the `get` and `list` calls.
	State                VariableState `protobuf:"varint,4,opt,name=state,proto3,enum=google.cloud.runtimeconfig.v1beta1.VariableState" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Variable) Reset()         { *m = Variable{} }
func (m *Variable) String() string { return proto.CompactTextString(m) }
func (*Variable) ProtoMessage()    {}
func (*Variable) Descriptor() ([]byte, []int) {
	return fileDescriptor_7913f3704a8d250c, []int{1}
}

func (m *Variable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Variable.Unmarshal(m, b)
}
func (m *Variable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Variable.Marshal(b, m, deterministic)
}
func (m *Variable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Variable.Merge(m, src)
}
func (m *Variable) XXX_Size() int {
	return xxx_messageInfo_Variable.Size(m)
}
func (m *Variable) XXX_DiscardUnknown() {
	xxx_messageInfo_Variable.DiscardUnknown(m)
}

var xxx_messageInfo_Variable proto.InternalMessageInfo

func (m *Variable) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isVariable_Contents interface {
	isVariable_Contents()
}

type Variable_Value struct {
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3,oneof"`
}

type Variable_Text struct {
	Text string `protobuf:"bytes,5,opt,name=text,proto3,oneof"`
}

func (*Variable_Value) isVariable_Contents() {}

func (*Variable_Text) isVariable_Contents() {}

func (m *Variable) GetContents() isVariable_Contents {
	if m != nil {
		return m.Contents
	}
	return nil
}

func (m *Variable) GetValue() []byte {
	if x, ok := m.GetContents().(*Variable_Value); ok {
		return x.Value
	}
	return nil
}

func (m *Variable) GetText() string {
	if x, ok := m.GetContents().(*Variable_Text); ok {
		return x.Text
	}
	return ""
}

func (m *Variable) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func (m *Variable) GetState() VariableState {
	if m != nil {
		return m.State
	}
	return VariableState_VARIABLE_STATE_UNSPECIFIED
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Variable) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Variable_Value)(nil),
		(*Variable_Text)(nil),
	}
}

// The condition that a Waiter resource is waiting for.
type EndCondition struct {
	// The condition oneof holds the available condition types for this
	// EndCondition. Currently, the only available type is Cardinality.
	//
	// Types that are valid to be assigned to Condition:
	//	*EndCondition_Cardinality_
	Condition            isEndCondition_Condition `protobuf_oneof:"condition"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *EndCondition) Reset()         { *m = EndCondition{} }
func (m *EndCondition) String() string { return proto.CompactTextString(m) }
func (*EndCondition) ProtoMessage()    {}
func (*EndCondition) Descriptor() ([]byte, []int) {
	return fileDescriptor_7913f3704a8d250c, []int{2}
}

func (m *EndCondition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndCondition.Unmarshal(m, b)
}
func (m *EndCondition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndCondition.Marshal(b, m, deterministic)
}
func (m *EndCondition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndCondition.Merge(m, src)
}
func (m *EndCondition) XXX_Size() int {
	return xxx_messageInfo_EndCondition.Size(m)
}
func (m *EndCondition) XXX_DiscardUnknown() {
	xxx_messageInfo_EndCondition.DiscardUnknown(m)
}

var xxx_messageInfo_EndCondition proto.InternalMessageInfo

type isEndCondition_Condition interface {
	isEndCondition_Condition()
}

type EndCondition_Cardinality_ struct {
	Cardinality *EndCondition_Cardinality `protobuf:"bytes,1,opt,name=cardinality,proto3,oneof"`
}

func (*EndCondition_Cardinality_) isEndCondition_Condition() {}

func (m *EndCondition) GetCondition() isEndCondition_Condition {
	if m != nil {
		return m.Condition
	}
	return nil
}

func (m *EndCondition) GetCardinality() *EndCondition_Cardinality {
	if x, ok := m.GetCondition().(*EndCondition_Cardinality_); ok {
		return x.Cardinality
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*EndCondition) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*EndCondition_Cardinality_)(nil),
	}
}

// A Cardinality condition for the Waiter resource. A cardinality condition is
// met when the number of variables under a specified path prefix reaches a
// predefined number. For example, if you set a Cardinality condition where
// the `path` is set to `/foo` and the number of paths is set to 2, the
// following variables would meet the condition in a RuntimeConfig resource:
//
// + `/foo/variable1 = "value1"`
// + `/foo/variable2 = "value2"`
// + `/bar/variable3 = "value3"`
//
// It would not would not satisify the same condition with the `number` set to
// 3, however, because there is only 2 paths that start with `/foo`.
// Cardinality conditions are recursive; all subtrees under the specific
// path prefix are counted.
type EndCondition_Cardinality struct {
	// The root of the variable subtree to monitor. For example, `/foo`.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// The number variables under the `path` that must exist to meet this
	// condition. Defaults to 1 if not specified.
	Number               int32    `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EndCondition_Cardinality) Reset()         { *m = EndCondition_Cardinality{} }
func (m *EndCondition_Cardinality) String() string { return proto.CompactTextString(m) }
func (*EndCondition_Cardinality) ProtoMessage()    {}
func (*EndCondition_Cardinality) Descriptor() ([]byte, []int) {
	return fileDescriptor_7913f3704a8d250c, []int{2, 0}
}

func (m *EndCondition_Cardinality) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndCondition_Cardinality.Unmarshal(m, b)
}
func (m *EndCondition_Cardinality) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndCondition_Cardinality.Marshal(b, m, deterministic)
}
func (m *EndCondition_Cardinality) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndCondition_Cardinality.Merge(m, src)
}
func (m *EndCondition_Cardinality) XXX_Size() int {
	return xxx_messageInfo_EndCondition_Cardinality.Size(m)
}
func (m *EndCondition_Cardinality) XXX_DiscardUnknown() {
	xxx_messageInfo_EndCondition_Cardinality.DiscardUnknown(m)
}

var xxx_messageInfo_EndCondition_Cardinality proto.InternalMessageInfo

func (m *EndCondition_Cardinality) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *EndCondition_Cardinality) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

// A Waiter resource waits for some end condition within a RuntimeConfig
// resource to be met before it returns. For example, assume you have a
// distributed system where each node writes to a Variable resource indidicating
// the node's readiness as part of the startup process.
//
// You then configure a Waiter resource with the success condition set to wait
// until some number of nodes have checked in. Afterwards, your application
// runs some arbitrary code after the condition has been met and the waiter
// returns successfully.
//
// Once created, a Waiter resource is immutable.
//
// To learn more about using waiters, read the
// [Creating a
// Waiter](/deployment-manager/runtime-configurator/creating-a-waiter)
// documentation.
type Waiter struct {
	// The name of the Waiter resource, in the format:
	//
	//     projects/[PROJECT_ID]/configs/[CONFIG_NAME]/waiters/[WAITER_NAME]
	//
	// The `[PROJECT_ID]` must be a valid Google Cloud project ID,
	// the `[CONFIG_NAME]` must be a valid RuntimeConfig resource, the
	// `[WAITER_NAME]` must match RFC 1035 segment specification, and the length
	// of `[WAITER_NAME]` must be less than 64 bytes.
	//
	// After you create a Waiter resource, you cannot change the resource name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// [Required] Specifies the timeout of the waiter in seconds, beginning from
	// the instant that `waiters().create` method is called. If this time elapses
	// before the success or failure conditions are met, the waiter fails and sets
	// the `error` code to `DEADLINE_EXCEEDED`.
	Timeout *duration.Duration `protobuf:"bytes,2,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// [Optional] The failure condition of this waiter. If this condition is met,
	// `done` will be set to `true` and the `error` code will be set to `ABORTED`.
	// The failure condition takes precedence over the success condition. If both
	// conditions are met, a failure will be indicated. This value is optional; if
	// no failure condition is set, the only failure scenario will be a timeout.
	Failure *EndCondition `protobuf:"bytes,3,opt,name=failure,proto3" json:"failure,omitempty"`
	// [Required] The success condition. If this condition is met, `done` will be
	// set to `true` and the `error` value will remain unset. The failure
	// condition takes precedence over the success condition. If both conditions
	// are met, a failure will be indicated.
	Success *EndCondition `protobuf:"bytes,4,opt,name=success,proto3" json:"success,omitempty"`
	// [Output Only] The instant at which this Waiter resource was created. Adding
	// the value of `timeout` to this instant yields the timeout deadline for the
	// waiter.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// [Output Only] If the value is `false`, it means the waiter is still waiting
	// for one of its conditions to be met.
	//
	// If true, the waiter has finished. If the waiter finished due to a timeout
	// or failure, `error` will be set.
	Done bool `protobuf:"varint,6,opt,name=done,proto3" json:"done,omitempty"`
	// [Output Only] If the waiter ended due to a failure or timeout, this value
	// will be set.
	Error                *status.Status `protobuf:"bytes,7,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Waiter) Reset()         { *m = Waiter{} }
func (m *Waiter) String() string { return proto.CompactTextString(m) }
func (*Waiter) ProtoMessage()    {}
func (*Waiter) Descriptor() ([]byte, []int) {
	return fileDescriptor_7913f3704a8d250c, []int{3}
}

func (m *Waiter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Waiter.Unmarshal(m, b)
}
func (m *Waiter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Waiter.Marshal(b, m, deterministic)
}
func (m *Waiter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Waiter.Merge(m, src)
}
func (m *Waiter) XXX_Size() int {
	return xxx_messageInfo_Waiter.Size(m)
}
func (m *Waiter) XXX_DiscardUnknown() {
	xxx_messageInfo_Waiter.DiscardUnknown(m)
}

var xxx_messageInfo_Waiter proto.InternalMessageInfo

func (m *Waiter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Waiter) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *Waiter) GetFailure() *EndCondition {
	if m != nil {
		return m.Failure
	}
	return nil
}

func (m *Waiter) GetSuccess() *EndCondition {
	if m != nil {
		return m.Success
	}
	return nil
}

func (m *Waiter) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Waiter) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

func (m *Waiter) GetError() *status.Status {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterEnum("google.cloud.runtimeconfig.v1beta1.VariableState", VariableState_name, VariableState_value)
	proto.RegisterType((*RuntimeConfig)(nil), "google.cloud.runtimeconfig.v1beta1.RuntimeConfig")
	proto.RegisterType((*Variable)(nil), "google.cloud.runtimeconfig.v1beta1.Variable")
	proto.RegisterType((*EndCondition)(nil), "google.cloud.runtimeconfig.v1beta1.EndCondition")
	proto.RegisterType((*EndCondition_Cardinality)(nil), "google.cloud.runtimeconfig.v1beta1.EndCondition.Cardinality")
	proto.RegisterType((*Waiter)(nil), "google.cloud.runtimeconfig.v1beta1.Waiter")
}

func init() {
	proto.RegisterFile("google/cloud/runtimeconfig/v1beta1/resources.proto", fileDescriptor_7913f3704a8d250c)
}

var fileDescriptor_7913f3704a8d250c = []byte{
	// 628 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0xc7, 0x9b, 0xd2, 0x8f, 0xed, 0x64, 0x43, 0x93, 0x85, 0x46, 0xa8, 0xd0, 0xa8, 0x7a, 0x81,
	0x2a, 0x2e, 0x12, 0xda, 0x5d, 0xa1, 0x71, 0xd3, 0x8f, 0xb0, 0x15, 0x4d, 0x30, 0xa5, 0x5d, 0x91,
	0xd0, 0xa4, 0xe1, 0x3a, 0x6e, 0x88, 0x94, 0xda, 0x91, 0xe3, 0x4c, 0xf0, 0x4a, 0x3c, 0x01, 0x2f,
	0xc0, 0x0d, 0x0f, 0xc1, 0x15, 0x0f, 0x82, 0xec, 0x38, 0xd0, 0xc2, 0xc4, 0x06, 0x77, 0x3e, 0x3e,
	0xff, 0xf3, 0x3b, 0x1f, 0x3e, 0x09, 0xf4, 0x23, 0xce, 0xa3, 0x84, 0x7a, 0x24, 0xe1, 0x79, 0xe8,
	0x89, 0x9c, 0xc9, 0x78, 0x45, 0x09, 0x67, 0xcb, 0x38, 0xf2, 0xae, 0x7a, 0x0b, 0x2a, 0x71, 0xcf,
	0x13, 0x34, 0xe3, 0xb9, 0x20, 0x34, 0x73, 0x53, 0xc1, 0x25, 0x47, 0x9d, 0x22, 0xc6, 0xd5, 0x31,
	0xee, 0x46, 0x8c, 0x6b, 0x62, 0x5a, 0x0f, 0x0d, 0x17, 0xa7, 0xb1, 0x87, 0x19, 0xe3, 0x12, 0xcb,
	0x98, 0x33, 0x43, 0x68, 0x1d, 0x18, 0xaf, 0xb6, 0x16, 0xf9, 0xd2, 0x0b, 0x73, 0xa1, 0x05, 0xc6,
	0xff, 0xe8, 0x77, 0xbf, 0xca, 0x90, 0x49, 0xbc, 0x4a, 0x8d, 0xe0, 0xbe, 0x11, 0x88, 0x94, 0x78,
	0x99, 0xc4, 0x32, 0x37, 0xe4, 0x8e, 0x0f, 0xbb, 0x41, 0x51, 0xd0, 0x48, 0x17, 0x84, 0x10, 0xd4,
	0x18, 0x5e, 0x51, 0xc7, 0x6a, 0x5b, 0xdd, 0xed, 0x40, 0x9f, 0x51, 0x1b, 0xec, 0x90, 0x66, 0x44,
	0xc4, 0xa9, 0xca, 0xe9, 0x54, 0xb5, 0x6b, 0xfd, 0xaa, 0xf3, 0xcd, 0x82, 0xad, 0x39, 0x16, 0x31,
	0x5e, 0x24, 0xf4, 0x5a, 0xc4, 0x3e, 0xd4, 0xaf, 0x70, 0x92, 0x53, 0x1d, 0xbc, 0x73, 0x52, 0x09,
	0x0a, 0x13, 0xdd, 0x83, 0x9a, 0xa4, 0x1f, 0xa4, 0x53, 0x57, 0xda, 0x93, 0x4a, 0xa0, 0x2d, 0x74,
	0x04, 0x76, 0x9e, 0x86, 0x58, 0xd2, 0x4b, 0x55, 0x99, 0x73, 0xa7, 0x6d, 0x75, 0xed, 0x7e, 0xcb,
	0x35, 0x73, 0x2c, 0xbb, 0x74, 0x67, 0x65, 0x97, 0x01, 0x14, 0x72, 0x75, 0x81, 0x8e, 0xa1, 0xae,
	0x5a, 0xa4, 0x4e, 0xad, 0x6d, 0x75, 0xef, 0xf6, 0x7b, 0xee, 0xcd, 0xe3, 0x77, 0xcb, 0xda, 0xa7,
	0x2a, 0x30, 0x28, 0xe2, 0x87, 0x00, 0x5b, 0x84, 0x33, 0x49, 0x99, 0xcc, 0x3a, 0x9f, 0x2d, 0xd8,
	0xf1, 0x59, 0x38, 0xe2, 0x2c, 0x8c, 0x55, 0xc7, 0xe8, 0x1d, 0xd8, 0x04, 0x8b, 0x30, 0x66, 0x38,
	0x89, 0xe5, 0x47, 0xdd, 0xab, 0xdd, 0x7f, 0x7e, 0x9b, 0x5c, 0xeb, 0x18, 0x77, 0xf4, 0x8b, 0x71,
	0x52, 0x09, 0xd6, 0x91, 0xad, 0x67, 0x60, 0xaf, 0x79, 0xd5, 0x54, 0x53, 0x2c, 0xdf, 0x97, 0x53,
	0x55, 0x67, 0xb4, 0x0f, 0x0d, 0x96, 0xaf, 0x16, 0x54, 0xe8, 0xb1, 0xd6, 0x03, 0x63, 0x0d, 0x6d,
	0xd8, 0x26, 0x65, 0x8a, 0xce, 0xf7, 0x2a, 0x34, 0xde, 0xe0, 0x58, 0x52, 0x71, 0xed, 0xcb, 0x1c,
	0x42, 0x53, 0x15, 0xc9, 0x73, 0xa9, 0x21, 0x76, 0xff, 0xc1, 0x1f, 0x73, 0x1e, 0x9b, 0x6d, 0x0b,
	0x4a, 0x25, 0x7a, 0x09, 0xcd, 0x25, 0x8e, 0x93, 0x5c, 0x94, 0x8f, 0xf3, 0xf4, 0x5f, 0x3b, 0x0f,
	0x4a, 0x80, 0x62, 0x65, 0x39, 0x21, 0x34, 0xcb, 0xf4, 0x8b, 0xfd, 0x17, 0xcb, 0x00, 0xd4, 0xe2,
	0x10, 0x41, 0x7f, 0x2e, 0x4e, 0xfd, 0xe6, 0xc5, 0x29, 0xe4, 0x7a, 0x71, 0x10, 0xd4, 0x42, 0xce,
	0xa8, 0xd3, 0x68, 0x5b, 0xdd, 0xad, 0x40, 0x9f, 0x51, 0x17, 0xea, 0x54, 0x08, 0x2e, 0x9c, 0xa6,
	0x46, 0xa1, 0x12, 0x25, 0x52, 0xe2, 0x4e, 0xf5, 0x87, 0x14, 0x14, 0x82, 0x27, 0x13, 0xd8, 0xdd,
	0xd8, 0x22, 0x74, 0x00, 0xad, 0xf9, 0x20, 0x98, 0x0c, 0x86, 0xa7, 0xfe, 0xe5, 0x74, 0x36, 0x98,
	0xf9, 0x97, 0xe7, 0xaf, 0xa6, 0x67, 0xfe, 0x68, 0xf2, 0x62, 0xe2, 0x8f, 0xf7, 0x2a, 0xc8, 0x86,
	0xe6, 0xf9, 0xd9, 0x78, 0x30, 0xf3, 0xc7, 0x7b, 0x96, 0x32, 0xc6, 0xfe, 0xa9, 0xaf, 0x8c, 0xea,
	0xf0, 0x8b, 0x05, 0x8f, 0x09, 0x5f, 0xdd, 0x62, 0x0c, 0x67, 0xd6, 0xdb, 0xd7, 0x46, 0x15, 0xf1,
	0x04, 0xb3, 0xc8, 0xe5, 0x22, 0xf2, 0x22, 0xca, 0x74, 0xab, 0x5e, 0xe1, 0xc2, 0x69, 0x9c, 0xfd,
	0xed, 0x87, 0x75, 0xb4, 0x71, 0xfb, 0xa9, 0xda, 0x39, 0x2e, 0x88, 0x23, 0x9d, 0x77, 0xe3, 0xf7,
	0xe0, 0xce, 0x7b, 0x43, 0x15, 0xf2, 0xb5, 0x14, 0x5d, 0x68, 0xd1, 0xc5, 0x86, 0xe8, 0x62, 0x5e,
	0x70, 0x17, 0x0d, 0x5d, 0xc5, 0xe1, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1a, 0xc9, 0x60, 0x90,
	0x35, 0x05, 0x00, 0x00,
}
