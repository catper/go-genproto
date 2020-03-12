// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/mutate_job_status.proto

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
	return fileDescriptor_52c7990b70abb227, []int{0, 0}
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
	return fileDescriptor_52c7990b70abb227, []int{0}
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
	proto.RegisterEnum("google.ads.googleads.v3.enums.MutateJobStatusEnum_MutateJobStatus", MutateJobStatusEnum_MutateJobStatus_name, MutateJobStatusEnum_MutateJobStatus_value)
	proto.RegisterType((*MutateJobStatusEnum)(nil), "google.ads.googleads.v3.enums.MutateJobStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/mutate_job_status.proto", fileDescriptor_52c7990b70abb227)
}

var fileDescriptor_52c7990b70abb227 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0x41, 0x4e, 0xeb, 0x30,
	0x10, 0xfd, 0x4d, 0xab, 0x0f, 0x72, 0x17, 0x8d, 0x02, 0x2b, 0x44, 0x17, 0xed, 0x01, 0xec, 0x85,
	0xc5, 0xc6, 0xac, 0x52, 0x1a, 0xaa, 0x82, 0x70, 0x23, 0xaa, 0x04, 0x09, 0x45, 0xaa, 0x1c, 0x12,
	0x59, 0xa9, 0x1a, 0x3b, 0xaa, 0x9d, 0x1e, 0x88, 0x25, 0x47, 0xe1, 0x26, 0x70, 0x0a, 0x64, 0xa7,
	0xc9, 0xa2, 0x12, 0x6c, 0xac, 0x37, 0x33, 0xef, 0x3d, 0xbf, 0x19, 0x70, 0xc3, 0xa5, 0xe4, 0xbb,
	0x1c, 0xb1, 0x4c, 0xa1, 0x06, 0x1a, 0x74, 0xc0, 0x28, 0x17, 0x75, 0xa9, 0x50, 0x59, 0x6b, 0xa6,
	0xf3, 0xcd, 0x56, 0xa6, 0x1b, 0xa5, 0x99, 0xae, 0x15, 0xac, 0xf6, 0x52, 0x4b, 0x6f, 0xdc, 0x70,
	0x21, 0xcb, 0x14, 0xec, 0x64, 0xf0, 0x80, 0xa1, 0x95, 0x5d, 0x5d, 0xb7, 0xae, 0x55, 0x81, 0x98,
	0x10, 0x52, 0x33, 0x5d, 0x48, 0x71, 0x14, 0x4f, 0xb7, 0xe0, 0xe2, 0xc9, 0xfa, 0x3e, 0xc8, 0x74,
	0x6d, 0x5d, 0x03, 0x51, 0x97, 0xd3, 0x35, 0x18, 0x9d, 0xb4, 0xbd, 0x11, 0x18, 0x46, 0x74, 0x1d,
	0x06, 0x77, 0xcb, 0xfb, 0x65, 0x30, 0x77, 0xff, 0x79, 0x43, 0x70, 0x16, 0xd1, 0x47, 0xba, 0x7a,
	0xa1, 0x6e, 0xcf, 0x14, 0x61, 0x40, 0xe7, 0x4b, 0xba, 0x70, 0x1d, 0x53, 0x3c, 0x47, 0x94, 0x9a,
	0xa2, 0xef, 0x9d, 0x83, 0xc1, 0x7c, 0x45, 0x03, 0x77, 0x30, 0xfb, 0xea, 0x81, 0xc9, 0x9b, 0x2c,
	0xe1, 0x9f, 0x79, 0x67, 0x97, 0x27, 0x1f, 0x87, 0x26, 0x67, 0xd8, 0x7b, 0x9d, 0x1d, 0x65, 0x5c,
	0xee, 0x98, 0xe0, 0x50, 0xee, 0x39, 0xe2, 0xb9, 0xb0, 0x5b, 0xb4, 0xd7, 0xaa, 0x0a, 0xf5, 0xcb,
	0xf1, 0x6e, 0xed, 0xfb, 0xee, 0xf4, 0x17, 0xbe, 0xff, 0xe1, 0x8c, 0x17, 0x8d, 0x95, 0x9f, 0x29,
	0xd8, 0x40, 0x83, 0x62, 0x0c, 0xcd, 0xee, 0xea, 0xb3, 0x9d, 0x27, 0x7e, 0xa6, 0x92, 0x6e, 0x9e,
	0xc4, 0x38, 0xb1, 0xf3, 0x6f, 0x67, 0xd2, 0x34, 0x09, 0xf1, 0x33, 0x45, 0x48, 0xc7, 0x20, 0x24,
	0xc6, 0x84, 0x58, 0x4e, 0xfa, 0xdf, 0x06, 0xc3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7d, 0xc6,
	0xee, 0xec, 0xd4, 0x01, 0x00, 0x00,
}
