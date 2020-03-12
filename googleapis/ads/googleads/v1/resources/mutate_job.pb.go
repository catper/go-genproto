// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/mutate_job.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
	wrappers "github.com/catper/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
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

// A list of mutates being processed asynchronously. The mutates are uploaded
// by the user. The mutates themselves aren't readable and the results of the
// job can only be read using MutateJobService.ListMutateJobResults.
type MutateJob struct {
	// Immutable. The resource name of the mutate job.
	// Mutate job resource names have the form:
	//
	// `customers/{customer_id}/mutateJobs/{mutate_job_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. ID of this mutate job.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Output only. The next sequence token to use when adding operations. Only set when the
	// mutate job status is PENDING.
	NextAddSequenceToken *wrappers.StringValue `protobuf:"bytes,3,opt,name=next_add_sequence_token,json=nextAddSequenceToken,proto3" json:"next_add_sequence_token,omitempty"`
	// Output only. Contains additional information about this mutate job.
	Metadata *MutateJob_MutateJobMetadata `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Output only. Status of this mutate job.
	Status enums.MutateJobStatusEnum_MutateJobStatus `protobuf:"varint,5,opt,name=status,proto3,enum=google.ads.googleads.v1.enums.MutateJobStatusEnum_MutateJobStatus" json:"status,omitempty"`
	// Output only. The resource name of the long-running operation that can be used to poll
	// for completion. Only set when the mutate job status is RUNNING or DONE.
	LongRunningOperation *wrappers.StringValue `protobuf:"bytes,6,opt,name=long_running_operation,json=longRunningOperation,proto3" json:"long_running_operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MutateJob) Reset()         { *m = MutateJob{} }
func (m *MutateJob) String() string { return proto.CompactTextString(m) }
func (*MutateJob) ProtoMessage()    {}
func (*MutateJob) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5e8cce5a5c605aa, []int{0}
}

func (m *MutateJob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateJob.Unmarshal(m, b)
}
func (m *MutateJob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateJob.Marshal(b, m, deterministic)
}
func (m *MutateJob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateJob.Merge(m, src)
}
func (m *MutateJob) XXX_Size() int {
	return xxx_messageInfo_MutateJob.Size(m)
}
func (m *MutateJob) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateJob.DiscardUnknown(m)
}

var xxx_messageInfo_MutateJob proto.InternalMessageInfo

func (m *MutateJob) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *MutateJob) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *MutateJob) GetNextAddSequenceToken() *wrappers.StringValue {
	if m != nil {
		return m.NextAddSequenceToken
	}
	return nil
}

func (m *MutateJob) GetMetadata() *MutateJob_MutateJobMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *MutateJob) GetStatus() enums.MutateJobStatusEnum_MutateJobStatus {
	if m != nil {
		return m.Status
	}
	return enums.MutateJobStatusEnum_UNSPECIFIED
}

func (m *MutateJob) GetLongRunningOperation() *wrappers.StringValue {
	if m != nil {
		return m.LongRunningOperation
	}
	return nil
}

// Additional information about the mutate job. This message is also used as
// metadata returned in mutate job Long Running Operations.
type MutateJob_MutateJobMetadata struct {
	// Output only. The time when this mutate job was created.
	// Formatted as yyyy-mm-dd hh:mm:ss. Example: "2018-03-05 09:15:00"
	CreationDateTime *wrappers.StringValue `protobuf:"bytes,1,opt,name=creation_date_time,json=creationDateTime,proto3" json:"creation_date_time,omitempty"`
	// Output only. The time when this mutate job was completed.
	// Formatted as yyyy-MM-dd HH:mm:ss. Example: "2018-03-05 09:16:00"
	CompletionDateTime *wrappers.StringValue `protobuf:"bytes,2,opt,name=completion_date_time,json=completionDateTime,proto3" json:"completion_date_time,omitempty"`
	// Output only. The fraction (between 0.0 and 1.0) of mutates that have been processed.
	// This is empty if the job hasn't started running yet.
	EstimatedCompletionRatio *wrappers.DoubleValue `protobuf:"bytes,3,opt,name=estimated_completion_ratio,json=estimatedCompletionRatio,proto3" json:"estimated_completion_ratio,omitempty"`
	// Output only. The number of mutate operations in the mutate job.
	OperationCount *wrappers.Int64Value `protobuf:"bytes,4,opt,name=operation_count,json=operationCount,proto3" json:"operation_count,omitempty"`
	// Output only. The number of mutate operations executed by the mutate job.
	// Present only if the job has started running.
	ExecutedOperationCount *wrappers.Int64Value `protobuf:"bytes,5,opt,name=executed_operation_count,json=executedOperationCount,proto3" json:"executed_operation_count,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}             `json:"-"`
	XXX_unrecognized       []byte               `json:"-"`
	XXX_sizecache          int32                `json:"-"`
}

func (m *MutateJob_MutateJobMetadata) Reset()         { *m = MutateJob_MutateJobMetadata{} }
func (m *MutateJob_MutateJobMetadata) String() string { return proto.CompactTextString(m) }
func (*MutateJob_MutateJobMetadata) ProtoMessage()    {}
func (*MutateJob_MutateJobMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5e8cce5a5c605aa, []int{0, 0}
}

func (m *MutateJob_MutateJobMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateJob_MutateJobMetadata.Unmarshal(m, b)
}
func (m *MutateJob_MutateJobMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateJob_MutateJobMetadata.Marshal(b, m, deterministic)
}
func (m *MutateJob_MutateJobMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateJob_MutateJobMetadata.Merge(m, src)
}
func (m *MutateJob_MutateJobMetadata) XXX_Size() int {
	return xxx_messageInfo_MutateJob_MutateJobMetadata.Size(m)
}
func (m *MutateJob_MutateJobMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateJob_MutateJobMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_MutateJob_MutateJobMetadata proto.InternalMessageInfo

func (m *MutateJob_MutateJobMetadata) GetCreationDateTime() *wrappers.StringValue {
	if m != nil {
		return m.CreationDateTime
	}
	return nil
}

func (m *MutateJob_MutateJobMetadata) GetCompletionDateTime() *wrappers.StringValue {
	if m != nil {
		return m.CompletionDateTime
	}
	return nil
}

func (m *MutateJob_MutateJobMetadata) GetEstimatedCompletionRatio() *wrappers.DoubleValue {
	if m != nil {
		return m.EstimatedCompletionRatio
	}
	return nil
}

func (m *MutateJob_MutateJobMetadata) GetOperationCount() *wrappers.Int64Value {
	if m != nil {
		return m.OperationCount
	}
	return nil
}

func (m *MutateJob_MutateJobMetadata) GetExecutedOperationCount() *wrappers.Int64Value {
	if m != nil {
		return m.ExecutedOperationCount
	}
	return nil
}

func init() {
	proto.RegisterType((*MutateJob)(nil), "google.ads.googleads.v1.resources.MutateJob")
	proto.RegisterType((*MutateJob_MutateJobMetadata)(nil), "google.ads.googleads.v1.resources.MutateJob.MutateJobMetadata")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/mutate_job.proto", fileDescriptor_a5e8cce5a5c605aa)
}

var fileDescriptor_a5e8cce5a5c605aa = []byte{
	// 670 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0x5d, 0x4f, 0xd4, 0x4c,
	0x14, 0xc7, 0xb3, 0xdd, 0x07, 0x02, 0xf3, 0x3c, 0x0f, 0xea, 0x84, 0x60, 0x5d, 0x89, 0x02, 0x91,
	0x84, 0x18, 0xd3, 0xba, 0xf8, 0x72, 0x51, 0x13, 0x93, 0x2e, 0x18, 0x02, 0x09, 0x2e, 0xd9, 0xc5,
	0xd5, 0x90, 0x4d, 0x26, 0xd3, 0xce, 0xa1, 0x56, 0xdb, 0x99, 0xda, 0x99, 0x22, 0x09, 0xe1, 0xab,
	0x78, 0xe1, 0xa5, 0xdf, 0xc2, 0x5b, 0x3f, 0x05, 0xd7, 0x7c, 0x04, 0xbd, 0x31, 0x7d, 0x9b, 0x5d,
	0x21, 0xc8, 0xde, 0x9d, 0xdd, 0xf3, 0xff, 0xff, 0xce, 0x39, 0x9d, 0x33, 0x83, 0xd6, 0x03, 0x21,
	0x82, 0x08, 0x6c, 0xca, 0xa4, 0x5d, 0x86, 0x79, 0x74, 0xd4, 0xb6, 0x53, 0x90, 0x22, 0x4b, 0x7d,
	0x90, 0x76, 0x9c, 0x29, 0xaa, 0x80, 0x7c, 0x10, 0x9e, 0x95, 0xa4, 0x42, 0x09, 0xbc, 0x5c, 0x0a,
	0x2d, 0xca, 0xa4, 0xa5, 0x3d, 0xd6, 0x51, 0xdb, 0xd2, 0x9e, 0xd6, 0xb3, 0xab, 0xb0, 0xc0, 0xb3,
	0x78, 0x1c, 0x49, 0xa4, 0xa2, 0x2a, 0x93, 0x25, 0xb9, 0x75, 0xbf, 0xb6, 0x25, 0xa1, 0x7d, 0x18,
	0x42, 0xc4, 0x88, 0x07, 0xef, 0xe9, 0x51, 0x28, 0xd2, 0x4a, 0x70, 0x67, 0x4c, 0x50, 0x57, 0xab,
	0x52, 0xf7, 0xaa, 0x54, 0xf1, 0xcb, 0xcb, 0x0e, 0xed, 0xcf, 0x29, 0x4d, 0x12, 0x48, 0x6b, 0xf6,
	0xe2, 0x98, 0x95, 0x72, 0x2e, 0x14, 0x55, 0xa1, 0xe0, 0x55, 0x76, 0xe5, 0xcb, 0x0c, 0x9a, 0xdd,
	0x2d, 0xba, 0xda, 0x11, 0x1e, 0xee, 0xa2, 0xff, 0x6b, 0x3a, 0xe1, 0x34, 0x06, 0xb3, 0xb1, 0xd4,
	0x58, 0x9b, 0xed, 0x3c, 0x3c, 0x73, 0xa7, 0x7e, 0xba, 0x0f, 0xd0, 0xca, 0x68, 0xea, 0x2a, 0x4a,
	0x42, 0x69, 0xf9, 0x22, 0xb6, 0x35, 0xa2, 0xf7, 0x5f, 0x0d, 0x78, 0x4d, 0x63, 0xc0, 0x8f, 0x91,
	0x11, 0x32, 0xd3, 0x58, 0x6a, 0xac, 0xfd, 0xbb, 0x7e, 0xb7, 0x32, 0x59, 0x75, 0xa7, 0xd6, 0x36,
	0x57, 0xcf, 0x9f, 0x0e, 0x68, 0x94, 0x41, 0xa7, 0x79, 0xe6, 0x36, 0x7b, 0x46, 0xc8, 0xf0, 0x3b,
	0x74, 0x9b, 0xc3, 0xb1, 0x22, 0x94, 0x31, 0x22, 0xe1, 0x53, 0x06, 0xdc, 0x07, 0xa2, 0xc4, 0x47,
	0xe0, 0x66, 0xb3, 0xc0, 0x2c, 0x5e, 0xc2, 0xf4, 0x55, 0x1a, 0xf2, 0x60, 0x8c, 0x33, 0x9f, 0x13,
	0x5c, 0xc6, 0xfa, 0x95, 0x7f, 0x3f, 0xb7, 0x63, 0x82, 0x66, 0x62, 0x50, 0x94, 0x51, 0x45, 0xcd,
	0x7f, 0x0a, 0xd4, 0x4b, 0xeb, 0xda, 0x13, 0xb5, 0xf4, 0x64, 0xa3, 0x68, 0xb7, 0xa2, 0x94, 0xc5,
	0x34, 0x14, 0x13, 0x34, 0x5d, 0x9e, 0xaa, 0x39, 0xb5, 0xd4, 0x58, 0x9b, 0x5b, 0xef, 0x5c, 0x89,
	0x2f, 0xb6, 0x61, 0x04, 0xec, 0x17, 0xae, 0x57, 0x3c, 0x8b, 0x2f, 0xfe, 0x57, 0x96, 0xa8, 0xb0,
	0xf8, 0x2d, 0x5a, 0x88, 0x04, 0x0f, 0x48, 0x9a, 0x71, 0x1e, 0xf2, 0x80, 0x88, 0x04, 0xd2, 0xe2,
	0x34, 0xcd, 0xe9, 0x89, 0x3f, 0x4d, 0x0e, 0xe8, 0x95, 0xfe, 0x6e, 0x6d, 0x6f, 0x7d, 0x6f, 0xa2,
	0x5b, 0x97, 0xc6, 0xc3, 0x5d, 0x84, 0xfd, 0x14, 0x0a, 0x05, 0x61, 0xf9, 0xde, 0xaa, 0xb0, 0x5a,
	0x89, 0x89, 0x4a, 0xdd, 0xac, 0xcd, 0x9b, 0x54, 0xc1, 0x7e, 0x18, 0x03, 0xee, 0xa3, 0x79, 0x5f,
	0xc4, 0x49, 0x04, 0x17, 0x90, 0xc6, 0xa4, 0x48, 0x3c, 0xb2, 0x6b, 0x28, 0x41, 0x2d, 0x90, 0x2a,
	0x8c, 0xa9, 0x02, 0x46, 0xc6, 0xf0, 0xc5, 0x68, 0x57, 0xee, 0xcc, 0xa6, 0xc8, 0xbc, 0x08, 0xc6,
	0xd0, 0xa6, 0x86, 0x6c, 0x68, 0x46, 0x2f, 0x47, 0xe0, 0x6d, 0x74, 0x43, 0x7f, 0x68, 0xe2, 0x8b,
	0x8c, 0xab, 0x6a, 0x7d, 0xae, 0x5f, 0xe8, 0x39, 0x6d, 0xdc, 0xc8, 0x7d, 0xf8, 0x00, 0x99, 0x70,
	0x0c, 0x7e, 0x96, 0xb7, 0x7a, 0x91, 0x39, 0x35, 0x21, 0x73, 0xa1, 0x26, 0x74, 0xff, 0x60, 0x3b,
	0x6f, 0xce, 0xdd, 0xde, 0x24, 0x37, 0x14, 0x3f, 0xf2, 0x33, 0xa9, 0x44, 0x0c, 0xa9, 0xb4, 0x4f,
	0xea, 0xf0, 0xb4, 0x7a, 0x9a, 0x76, 0x84, 0x27, 0xed, 0x93, 0xd1, 0x33, 0x75, 0xda, 0xf9, 0xd5,
	0x40, 0xab, 0xbe, 0x88, 0xaf, 0xbf, 0x29, 0x9d, 0x39, 0x5d, 0x62, 0x2f, 0xef, 0x7d, 0xaf, 0x71,
	0xb0, 0x53, 0x99, 0x02, 0x11, 0x51, 0x1e, 0x58, 0x22, 0x0d, 0xec, 0x00, 0x78, 0x31, 0x99, 0x3d,
	0x6a, 0xef, 0x2f, 0x6f, 0xf0, 0x0b, 0x1d, 0x7d, 0x35, 0x9a, 0x5b, 0xae, 0xfb, 0xcd, 0x58, 0xde,
	0x2a, 0x91, 0x2e, 0x93, 0x56, 0x19, 0xe6, 0xd1, 0xa0, 0x6d, 0xf5, 0x6a, 0xe5, 0x8f, 0x5a, 0x33,
	0x74, 0x99, 0x1c, 0x6a, 0xcd, 0x70, 0xd0, 0x1e, 0x6a, 0xcd, 0xb9, 0xb1, 0x5a, 0x26, 0x1c, 0xc7,
	0x65, 0xd2, 0x71, 0xb4, 0xca, 0x71, 0x06, 0x6d, 0xc7, 0xd1, 0x3a, 0x6f, 0xba, 0x68, 0xf6, 0xc9,
	0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9b, 0x6a, 0xd1, 0xdd, 0x2f, 0x06, 0x00, 0x00,
}
