// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/dataproc/v1/operations.proto

package dataproc

import (
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
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

// The operation state.
type ClusterOperationStatus_State int32

const (
	// Unused.
	ClusterOperationStatus_UNKNOWN ClusterOperationStatus_State = 0
	// The operation has been created.
	ClusterOperationStatus_PENDING ClusterOperationStatus_State = 1
	// The operation is running.
	ClusterOperationStatus_RUNNING ClusterOperationStatus_State = 2
	// The operation is done; either cancelled or completed.
	ClusterOperationStatus_DONE ClusterOperationStatus_State = 3
)

var ClusterOperationStatus_State_name = map[int32]string{
	0: "UNKNOWN",
	1: "PENDING",
	2: "RUNNING",
	3: "DONE",
}

var ClusterOperationStatus_State_value = map[string]int32{
	"UNKNOWN": 0,
	"PENDING": 1,
	"RUNNING": 2,
	"DONE":    3,
}

func (x ClusterOperationStatus_State) String() string {
	return proto.EnumName(ClusterOperationStatus_State_name, int32(x))
}

func (ClusterOperationStatus_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_78b4f56947714756, []int{0, 0}
}

// The status of the operation.
type ClusterOperationStatus struct {
	// Output only. A message containing the operation state.
	State ClusterOperationStatus_State `protobuf:"varint,1,opt,name=state,proto3,enum=google.cloud.dataproc.v1.ClusterOperationStatus_State" json:"state,omitempty"`
	// Output only. A message containing the detailed operation state.
	InnerState string `protobuf:"bytes,2,opt,name=inner_state,json=innerState,proto3" json:"inner_state,omitempty"`
	// Output only. A message containing any operation metadata details.
	Details string `protobuf:"bytes,3,opt,name=details,proto3" json:"details,omitempty"`
	// Output only. The time this state was entered.
	StateStartTime       *timestamp.Timestamp `protobuf:"bytes,4,opt,name=state_start_time,json=stateStartTime,proto3" json:"state_start_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ClusterOperationStatus) Reset()         { *m = ClusterOperationStatus{} }
func (m *ClusterOperationStatus) String() string { return proto.CompactTextString(m) }
func (*ClusterOperationStatus) ProtoMessage()    {}
func (*ClusterOperationStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_78b4f56947714756, []int{0}
}

func (m *ClusterOperationStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterOperationStatus.Unmarshal(m, b)
}
func (m *ClusterOperationStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterOperationStatus.Marshal(b, m, deterministic)
}
func (m *ClusterOperationStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterOperationStatus.Merge(m, src)
}
func (m *ClusterOperationStatus) XXX_Size() int {
	return xxx_messageInfo_ClusterOperationStatus.Size(m)
}
func (m *ClusterOperationStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterOperationStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterOperationStatus proto.InternalMessageInfo

func (m *ClusterOperationStatus) GetState() ClusterOperationStatus_State {
	if m != nil {
		return m.State
	}
	return ClusterOperationStatus_UNKNOWN
}

func (m *ClusterOperationStatus) GetInnerState() string {
	if m != nil {
		return m.InnerState
	}
	return ""
}

func (m *ClusterOperationStatus) GetDetails() string {
	if m != nil {
		return m.Details
	}
	return ""
}

func (m *ClusterOperationStatus) GetStateStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StateStartTime
	}
	return nil
}

// Metadata describing the operation.
type ClusterOperationMetadata struct {
	// Output only. Name of the cluster for the operation.
	ClusterName string `protobuf:"bytes,7,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	// Output only. Cluster UUID for the operation.
	ClusterUuid string `protobuf:"bytes,8,opt,name=cluster_uuid,json=clusterUuid,proto3" json:"cluster_uuid,omitempty"`
	// Output only. Current operation status.
	Status *ClusterOperationStatus `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
	// Output only. The previous operation status.
	StatusHistory []*ClusterOperationStatus `protobuf:"bytes,10,rep,name=status_history,json=statusHistory,proto3" json:"status_history,omitempty"`
	// Output only. The operation type.
	OperationType string `protobuf:"bytes,11,opt,name=operation_type,json=operationType,proto3" json:"operation_type,omitempty"`
	// Output only. Short description of operation.
	Description string `protobuf:"bytes,12,opt,name=description,proto3" json:"description,omitempty"`
	// Output only. Labels associated with the operation
	Labels map[string]string `protobuf:"bytes,13,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Output only. Errors encountered during operation execution.
	Warnings             []string `protobuf:"bytes,14,rep,name=warnings,proto3" json:"warnings,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterOperationMetadata) Reset()         { *m = ClusterOperationMetadata{} }
func (m *ClusterOperationMetadata) String() string { return proto.CompactTextString(m) }
func (*ClusterOperationMetadata) ProtoMessage()    {}
func (*ClusterOperationMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_78b4f56947714756, []int{1}
}

func (m *ClusterOperationMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterOperationMetadata.Unmarshal(m, b)
}
func (m *ClusterOperationMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterOperationMetadata.Marshal(b, m, deterministic)
}
func (m *ClusterOperationMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterOperationMetadata.Merge(m, src)
}
func (m *ClusterOperationMetadata) XXX_Size() int {
	return xxx_messageInfo_ClusterOperationMetadata.Size(m)
}
func (m *ClusterOperationMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterOperationMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterOperationMetadata proto.InternalMessageInfo

func (m *ClusterOperationMetadata) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *ClusterOperationMetadata) GetClusterUuid() string {
	if m != nil {
		return m.ClusterUuid
	}
	return ""
}

func (m *ClusterOperationMetadata) GetStatus() *ClusterOperationStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ClusterOperationMetadata) GetStatusHistory() []*ClusterOperationStatus {
	if m != nil {
		return m.StatusHistory
	}
	return nil
}

func (m *ClusterOperationMetadata) GetOperationType() string {
	if m != nil {
		return m.OperationType
	}
	return ""
}

func (m *ClusterOperationMetadata) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ClusterOperationMetadata) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *ClusterOperationMetadata) GetWarnings() []string {
	if m != nil {
		return m.Warnings
	}
	return nil
}

func init() {
	proto.RegisterEnum("google.cloud.dataproc.v1.ClusterOperationStatus_State", ClusterOperationStatus_State_name, ClusterOperationStatus_State_value)
	proto.RegisterType((*ClusterOperationStatus)(nil), "google.cloud.dataproc.v1.ClusterOperationStatus")
	proto.RegisterType((*ClusterOperationMetadata)(nil), "google.cloud.dataproc.v1.ClusterOperationMetadata")
	proto.RegisterMapType((map[string]string)(nil), "google.cloud.dataproc.v1.ClusterOperationMetadata.LabelsEntry")
}

func init() {
	proto.RegisterFile("google/cloud/dataproc/v1/operations.proto", fileDescriptor_78b4f56947714756)
}

var fileDescriptor_78b4f56947714756 = []byte{
	// 563 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x7f, 0x6b, 0xd3, 0x40,
	0x18, 0xc7, 0x4d, 0xb3, 0xfe, 0xba, 0xac, 0xb5, 0x1c, 0x22, 0x47, 0x99, 0xac, 0x14, 0x95, 0xea,
	0x1f, 0x89, 0x9b, 0x20, 0x53, 0x41, 0xb4, 0xae, 0xe8, 0xd0, 0xa5, 0xb5, 0x5b, 0x11, 0x54, 0x08,
	0xd7, 0xe6, 0x96, 0x1d, 0xa6, 0xb9, 0x70, 0x77, 0xa9, 0xf4, 0xf5, 0xf8, 0x1a, 0x7c, 0x3f, 0xbe,
	0x14, 0xb9, 0xbb, 0xa4, 0x0d, 0xce, 0x81, 0xfa, 0x57, 0xef, 0x79, 0x9e, 0xef, 0xf3, 0xc9, 0xf7,
	0x79, 0x72, 0x0d, 0x78, 0x10, 0x31, 0x16, 0xc5, 0xc4, 0x5b, 0xc4, 0x2c, 0x0b, 0xbd, 0x10, 0x4b,
	0x9c, 0x72, 0xb6, 0xf0, 0x56, 0x07, 0x1e, 0x4b, 0x09, 0xc7, 0x92, 0xb2, 0x44, 0xb8, 0x29, 0x67,
	0x92, 0x41, 0x64, 0xa4, 0xae, 0x96, 0xba, 0x85, 0xd4, 0x5d, 0x1d, 0x74, 0xf7, 0x73, 0x08, 0x4e,
	0xa9, 0x77, 0x41, 0x49, 0x1c, 0x06, 0x73, 0x72, 0x89, 0x57, 0x94, 0x71, 0xd3, 0xba, 0x11, 0xe8,
	0x68, 0x9e, 0x5d, 0x78, 0x92, 0x2e, 0x89, 0x90, 0x78, 0x99, 0xe6, 0x82, 0xbd, 0x12, 0x01, 0x27,
	0x09, 0x93, 0xe5, 0x27, 0xf7, 0x7f, 0x54, 0xc0, 0xed, 0xd7, 0x71, 0x26, 0x24, 0xe1, 0xe3, 0xc2,
	0xd5, 0x99, 0xc4, 0x32, 0x13, 0xf0, 0x03, 0xa8, 0x0a, 0x89, 0x25, 0x41, 0x56, 0xcf, 0x1a, 0xb4,
	0x0f, 0x9f, 0xb8, 0xd7, 0x99, 0x74, 0xff, 0x0c, 0x70, 0xd5, 0x0f, 0x19, 0xda, 0x3f, 0x5f, 0xd9,
	0x53, 0x43, 0x82, 0x77, 0x81, 0x43, 0x93, 0x84, 0xf0, 0xc0, 0x80, 0x2b, 0x3d, 0x6b, 0xd0, 0x34,
	0x02, 0xa0, 0xf3, 0xba, 0x03, 0xde, 0x01, 0xf5, 0x90, 0x48, 0x4c, 0x63, 0x81, 0xec, 0xad, 0xa2,
	0xc8, 0xc1, 0x13, 0xd0, 0xd1, 0xed, 0x0a, 0xc2, 0x65, 0xa0, 0xe6, 0x45, 0x3b, 0x3d, 0x6b, 0xe0,
	0x1c, 0x76, 0x0b, 0x8b, 0xc5, 0x32, 0xdc, 0xf3, 0x62, 0x19, 0x86, 0xd1, 0xd6, 0x8d, 0x67, 0xaa,
	0x4f, 0x55, 0xfa, 0x47, 0xa0, 0x6a, 0x1e, 0xe9, 0x80, 0xfa, 0xcc, 0x7f, 0xe7, 0x8f, 0x3f, 0xfa,
	0x9d, 0x1b, 0x2a, 0x98, 0x8c, 0xfc, 0xe3, 0x13, 0xff, 0x4d, 0xc7, 0x52, 0xc1, 0x74, 0xe6, 0xfb,
	0x2a, 0xa8, 0xc0, 0x06, 0xd8, 0x39, 0x1e, 0xfb, 0xa3, 0x8e, 0xdd, 0xff, 0xbe, 0x03, 0xd0, 0xef,
	0x63, 0x9f, 0x12, 0x89, 0xd5, 0x5a, 0xe0, 0x7d, 0xb0, 0xbb, 0x30, 0xb5, 0x20, 0xc1, 0x4b, 0x82,
	0xea, 0xdb, 0x29, 0x9c, 0xbc, 0xe0, 0xe3, 0x25, 0x29, 0xeb, 0xb2, 0x8c, 0x86, 0xa8, 0x71, 0x55,
	0x37, 0xcb, 0x68, 0x08, 0x4f, 0x41, 0x4d, 0xe8, 0x95, 0xa2, 0xa6, 0x9e, 0xf3, 0xd1, 0xbf, 0xbe,
	0x0a, 0xc3, 0xcc, 0x21, 0xf0, 0x0b, 0x68, 0x9b, 0x53, 0x70, 0x49, 0x85, 0x64, 0x7c, 0x8d, 0x40,
	0xcf, 0xfe, 0x7f, 0x6c, 0xcb, 0xc0, 0xde, 0x1a, 0x16, 0x7c, 0x08, 0xda, 0x9b, 0xfb, 0x1d, 0xc8,
	0x75, 0x4a, 0x90, 0xb3, 0x1d, 0xab, 0xb5, 0x29, 0x9d, 0xaf, 0x53, 0x02, 0xef, 0x01, 0x27, 0x24,
	0x62, 0xc1, 0x69, 0xaa, 0x52, 0x68, 0xb7, 0x34, 0x7f, 0x29, 0x0f, 0x3f, 0x83, 0x5a, 0x8c, 0xe7,
	0x24, 0x16, 0xa8, 0xa5, 0x8d, 0xbe, 0xf8, 0x7b, 0xa3, 0xc5, 0x3b, 0x71, 0xdf, 0x6b, 0xc0, 0x28,
	0x91, 0x7c, 0x9d, 0x6f, 0xc3, 0x20, 0xe1, 0x3e, 0x68, 0x7c, 0xc3, 0x3c, 0xa1, 0x49, 0x24, 0x50,
	0xbb, 0x67, 0x17, 0x06, 0x36, 0xc9, 0xee, 0x53, 0xe0, 0x94, 0x9a, 0x61, 0x07, 0xd8, 0x5f, 0xc9,
	0x5a, 0xff, 0x29, 0x9a, 0x53, 0x75, 0x84, 0xb7, 0x40, 0x75, 0x85, 0xe3, 0x2c, 0xbf, 0xcf, 0x53,
	0x13, 0x3c, 0xab, 0x1c, 0x59, 0x43, 0x01, 0xf6, 0x16, 0x6c, 0x79, 0xad, 0xdb, 0xe1, 0xcd, 0x8d,
	0x4f, 0x31, 0x51, 0x57, 0x76, 0x62, 0x7d, 0x7a, 0x99, 0x8b, 0x23, 0x16, 0xe3, 0x24, 0x72, 0x19,
	0x8f, 0xbc, 0x88, 0x24, 0xfa, 0x42, 0x7b, 0xa6, 0x84, 0x53, 0x2a, 0xae, 0x7e, 0x54, 0x9e, 0x17,
	0xe7, 0x79, 0x4d, 0x8b, 0x1f, 0xff, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x9d, 0xc4, 0x04, 0x47, 0x80,
	0x04, 0x00, 0x00,
}
