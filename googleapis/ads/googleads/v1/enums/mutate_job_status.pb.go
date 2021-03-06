// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/mutate_job_status.proto

package enums

import (
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
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

// The mutate job statuses.
type MutateJobStatusEnum_MutateJobStatus int32

const (
	// Not specified.
	MutateJobStatusEnum_UNSPECIFIED MutateJobStatusEnum_MutateJobStatus = 0
	// Used for return value only. Represents value unknown in this version.
	MutateJobStatusEnum_UNKNOWN MutateJobStatusEnum_MutateJobStatus = 1
	// The job is not currently running.
	MutateJobStatusEnum_PENDING MutateJobStatusEnum_MutateJobStatus = 2
	// The job is running.
	MutateJobStatusEnum_RUNNING MutateJobStatusEnum_MutateJobStatus = 3
	// The job is done.
	MutateJobStatusEnum_DONE MutateJobStatusEnum_MutateJobStatus = 4
)

var MutateJobStatusEnum_MutateJobStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "PENDING",
	3: "RUNNING",
	4: "DONE",
}

var MutateJobStatusEnum_MutateJobStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"PENDING":     2,
	"RUNNING":     3,
	"DONE":        4,
}

func (x MutateJobStatusEnum_MutateJobStatus) String() string {
	return proto.EnumName(MutateJobStatusEnum_MutateJobStatus_name, int32(x))
}

func (MutateJobStatusEnum_MutateJobStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_48c0d38bc796ede0, []int{0, 0}
}

// Container for enum describing possible mutate job statuses.
type MutateJobStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateJobStatusEnum) Reset()         { *m = MutateJobStatusEnum{} }
func (m *MutateJobStatusEnum) String() string { return proto.CompactTextString(m) }
func (*MutateJobStatusEnum) ProtoMessage()    {}
func (*MutateJobStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_48c0d38bc796ede0, []int{0}
}

func (m *MutateJobStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateJobStatusEnum.Unmarshal(m, b)
}
func (m *MutateJobStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateJobStatusEnum.Marshal(b, m, deterministic)
}
func (m *MutateJobStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateJobStatusEnum.Merge(m, src)
}
func (m *MutateJobStatusEnum) XXX_Size() int {
	return xxx_messageInfo_MutateJobStatusEnum.Size(m)
}
func (m *MutateJobStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateJobStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_MutateJobStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.MutateJobStatusEnum_MutateJobStatus", MutateJobStatusEnum_MutateJobStatus_name, MutateJobStatusEnum_MutateJobStatus_value)
	proto.RegisterType((*MutateJobStatusEnum)(nil), "google.ads.googleads.v1.enums.MutateJobStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/mutate_job_status.proto", fileDescriptor_48c0d38bc796ede0)
}

var fileDescriptor_48c0d38bc796ede0 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xd1, 0x4a, 0xc3, 0x30,
	0x14, 0x75, 0xdd, 0x50, 0xc9, 0x1e, 0x56, 0xaa, 0x4f, 0xe2, 0x1e, 0xb6, 0x0f, 0x48, 0x28, 0xe2,
	0x4b, 0x7c, 0xea, 0x5c, 0x1d, 0x53, 0xcc, 0x8a, 0xa5, 0x15, 0xa4, 0x30, 0x52, 0x5b, 0x42, 0xc7,
	0x9a, 0x94, 0x25, 0xdd, 0x07, 0xf9, 0xe8, 0xa7, 0xf8, 0x27, 0xfa, 0x15, 0x92, 0x74, 0xed, 0xc3,
	0x40, 0x5f, 0xc2, 0xb9, 0xf7, 0x9e, 0x73, 0x72, 0xee, 0x05, 0xb7, 0x4c, 0x08, 0xb6, 0xcd, 0x11,
	0xcd, 0x24, 0x6a, 0xa0, 0x46, 0x7b, 0x17, 0xe5, 0xbc, 0x2e, 0x25, 0x2a, 0x6b, 0x45, 0x55, 0xbe,
	0xde, 0x88, 0x74, 0x2d, 0x15, 0x55, 0xb5, 0x84, 0xd5, 0x4e, 0x28, 0xe1, 0x8c, 0x1b, 0x2e, 0xa4,
	0x99, 0x84, 0x9d, 0x0c, 0xee, 0x5d, 0x68, 0x64, 0x57, 0xd7, 0xad, 0x6b, 0x55, 0x20, 0xca, 0xb9,
	0x50, 0x54, 0x15, 0x82, 0x1f, 0xc4, 0xd3, 0x0d, 0xb8, 0x78, 0x36, 0xbe, 0x8f, 0x22, 0x0d, 0x8d,
	0xab, 0xcf, 0xeb, 0x72, 0x1a, 0x82, 0xd1, 0x51, 0xdb, 0x19, 0x81, 0x61, 0x44, 0xc2, 0xc0, 0xbf,
	0x5f, 0x3e, 0x2c, 0xfd, 0xb9, 0x7d, 0xe2, 0x0c, 0xc1, 0x59, 0x44, 0x9e, 0xc8, 0xea, 0x95, 0xd8,
	0x3d, 0x5d, 0x04, 0x3e, 0x99, 0x2f, 0xc9, 0xc2, 0xb6, 0x74, 0xf1, 0x12, 0x11, 0xa2, 0x8b, 0xbe,
	0x73, 0x0e, 0x06, 0xf3, 0x15, 0xf1, 0xed, 0xc1, 0xec, 0xbb, 0x07, 0x26, 0xef, 0xa2, 0x84, 0xff,
	0xe6, 0x9d, 0x5d, 0x1e, 0x7d, 0x1c, 0xe8, 0x9c, 0x41, 0xef, 0x6d, 0x76, 0x90, 0x31, 0xb1, 0xa5,
	0x9c, 0x41, 0xb1, 0x63, 0x88, 0xe5, 0xdc, 0x6c, 0xd1, 0x5e, 0xab, 0x2a, 0xe4, 0x1f, 0xc7, 0xbb,
	0x33, 0xef, 0x87, 0xd5, 0x5f, 0x78, 0xde, 0xa7, 0x35, 0x5e, 0x34, 0x56, 0x5e, 0x26, 0x61, 0x03,
	0x35, 0x8a, 0x5d, 0xa8, 0x77, 0x97, 0x5f, 0xed, 0x3c, 0xf1, 0x32, 0x99, 0x74, 0xf3, 0x24, 0x76,
	0x13, 0x33, 0xff, 0xb1, 0x26, 0x4d, 0x13, 0x63, 0x2f, 0x93, 0x18, 0x77, 0x0c, 0x8c, 0x63, 0x17,
	0x63, 0xc3, 0x49, 0x4f, 0x4d, 0xb0, 0x9b, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x36, 0x88,
	0xfe, 0xd4, 0x01, 0x00, 0x00,
}
