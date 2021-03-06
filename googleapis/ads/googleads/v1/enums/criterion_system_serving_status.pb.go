// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/criterion_system_serving_status.proto

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

// Enumerates criterion system serving statuses.
type CriterionSystemServingStatusEnum_CriterionSystemServingStatus int32

const (
	// Not specified.
	CriterionSystemServingStatusEnum_UNSPECIFIED CriterionSystemServingStatusEnum_CriterionSystemServingStatus = 0
	// The value is unknown in this version.
	CriterionSystemServingStatusEnum_UNKNOWN CriterionSystemServingStatusEnum_CriterionSystemServingStatus = 1
	// Eligible.
	CriterionSystemServingStatusEnum_ELIGIBLE CriterionSystemServingStatusEnum_CriterionSystemServingStatus = 2
	// Low search volume.
	CriterionSystemServingStatusEnum_RARELY_SERVED CriterionSystemServingStatusEnum_CriterionSystemServingStatus = 3
)

var CriterionSystemServingStatusEnum_CriterionSystemServingStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ELIGIBLE",
	3: "RARELY_SERVED",
}

var CriterionSystemServingStatusEnum_CriterionSystemServingStatus_value = map[string]int32{
	"UNSPECIFIED":   0,
	"UNKNOWN":       1,
	"ELIGIBLE":      2,
	"RARELY_SERVED": 3,
}

func (x CriterionSystemServingStatusEnum_CriterionSystemServingStatus) String() string {
	return proto.EnumName(CriterionSystemServingStatusEnum_CriterionSystemServingStatus_name, int32(x))
}

func (CriterionSystemServingStatusEnum_CriterionSystemServingStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b78d45305d2292b6, []int{0, 0}
}

// Container for enum describing possible criterion system serving statuses.
type CriterionSystemServingStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CriterionSystemServingStatusEnum) Reset()         { *m = CriterionSystemServingStatusEnum{} }
func (m *CriterionSystemServingStatusEnum) String() string { return proto.CompactTextString(m) }
func (*CriterionSystemServingStatusEnum) ProtoMessage()    {}
func (*CriterionSystemServingStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_b78d45305d2292b6, []int{0}
}

func (m *CriterionSystemServingStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CriterionSystemServingStatusEnum.Unmarshal(m, b)
}
func (m *CriterionSystemServingStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CriterionSystemServingStatusEnum.Marshal(b, m, deterministic)
}
func (m *CriterionSystemServingStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CriterionSystemServingStatusEnum.Merge(m, src)
}
func (m *CriterionSystemServingStatusEnum) XXX_Size() int {
	return xxx_messageInfo_CriterionSystemServingStatusEnum.Size(m)
}
func (m *CriterionSystemServingStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CriterionSystemServingStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CriterionSystemServingStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.CriterionSystemServingStatusEnum_CriterionSystemServingStatus", CriterionSystemServingStatusEnum_CriterionSystemServingStatus_name, CriterionSystemServingStatusEnum_CriterionSystemServingStatus_value)
	proto.RegisterType((*CriterionSystemServingStatusEnum)(nil), "google.ads.googleads.v1.enums.CriterionSystemServingStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/criterion_system_serving_status.proto", fileDescriptor_b78d45305d2292b6)
}

var fileDescriptor_b78d45305d2292b6 = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xc1, 0x4a, 0xf3, 0x40,
	0x18, 0xfc, 0x93, 0xc2, 0xaf, 0x6c, 0x15, 0x63, 0x8e, 0xd2, 0x82, 0xed, 0x03, 0x6c, 0x08, 0xde,
	0xd6, 0x53, 0xd2, 0xae, 0x25, 0x58, 0x62, 0x69, 0x68, 0x44, 0x89, 0x84, 0xd8, 0x84, 0x25, 0xd0,
	0xec, 0x96, 0x7c, 0xdb, 0x82, 0x47, 0x5f, 0xc5, 0xa3, 0x8f, 0xe2, 0xa3, 0xf8, 0x00, 0x9e, 0x25,
	0xbb, 0x4d, 0x6f, 0xe6, 0xb2, 0x0c, 0xfb, 0xcd, 0x37, 0xf3, 0xcd, 0xa0, 0x09, 0x13, 0x82, 0x6d,
	0x0a, 0x27, 0xcb, 0xc1, 0xd1, 0xb0, 0x41, 0x7b, 0xd7, 0x29, 0xf8, 0xae, 0x02, 0x67, 0x5d, 0x97,
	0xb2, 0xa8, 0x4b, 0xc1, 0x53, 0x78, 0x03, 0x59, 0x54, 0x29, 0x14, 0xf5, 0xbe, 0xe4, 0x2c, 0x05,
	0x99, 0xc9, 0x1d, 0xe0, 0x6d, 0x2d, 0xa4, 0xb0, 0x87, 0x7a, 0x13, 0x67, 0x39, 0xe0, 0xa3, 0x08,
	0xde, 0xbb, 0x58, 0x89, 0x5c, 0x0d, 0x5a, 0x8f, 0x6d, 0xe9, 0x64, 0x9c, 0x0b, 0x99, 0xc9, 0x52,
	0xf0, 0xc3, 0xf2, 0xf8, 0xdd, 0x40, 0xd7, 0x93, 0xd6, 0x26, 0x52, 0x2e, 0x91, 0x36, 0x89, 0x94,
	0x07, 0xe5, 0xbb, 0x6a, 0xfc, 0x82, 0x06, 0x5d, 0x1c, 0xfb, 0x02, 0xf5, 0x57, 0x61, 0xb4, 0xa0,
	0x93, 0xe0, 0x2e, 0xa0, 0x53, 0xeb, 0x9f, 0xdd, 0x47, 0x27, 0xab, 0xf0, 0x3e, 0x7c, 0x78, 0x0c,
	0x2d, 0xc3, 0x3e, 0x43, 0xa7, 0x74, 0x1e, 0xcc, 0x02, 0x7f, 0x4e, 0x2d, 0xd3, 0xbe, 0x44, 0xe7,
	0x4b, 0x6f, 0x49, 0xe7, 0x4f, 0x69, 0x44, 0x97, 0x31, 0x9d, 0x5a, 0x3d, 0xff, 0xc7, 0x40, 0xa3,
	0xb5, 0xa8, 0x70, 0x67, 0x0e, 0x7f, 0xd4, 0x75, 0xc2, 0xa2, 0x09, 0xb3, 0x30, 0x9e, 0xfd, 0x83,
	0x06, 0x13, 0x9b, 0x8c, 0x33, 0x2c, 0x6a, 0xe6, 0xb0, 0x82, 0xab, 0xa8, 0x6d, 0xc1, 0xdb, 0x12,
	0xfe, 0xe8, 0xfb, 0x56, 0xbd, 0x1f, 0x66, 0x6f, 0xe6, 0x79, 0x9f, 0xe6, 0x70, 0xa6, 0xa5, 0xbc,
	0x1c, 0xb0, 0x86, 0x0d, 0x8a, 0x5d, 0xdc, 0x54, 0x02, 0x5f, 0xed, 0x3c, 0xf1, 0x72, 0x48, 0x8e,
	0xf3, 0x24, 0x76, 0x13, 0x35, 0xff, 0x36, 0x47, 0xfa, 0x93, 0x10, 0x2f, 0x07, 0x42, 0x8e, 0x0c,
	0x42, 0x62, 0x97, 0x10, 0xc5, 0x79, 0xfd, 0xaf, 0x0e, 0xbb, 0xf9, 0x0d, 0x00, 0x00, 0xff, 0xff,
	0x2c, 0x98, 0x27, 0x97, 0x07, 0x02, 0x00, 0x00,
}
