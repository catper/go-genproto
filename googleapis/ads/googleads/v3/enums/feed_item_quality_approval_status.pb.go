// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/feed_item_quality_approval_status.proto

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

// The possible quality evaluation approval statuses of a feed item.
type FeedItemQualityApprovalStatusEnum_FeedItemQualityApprovalStatus int32

const (
	// No value has been specified.
	FeedItemQualityApprovalStatusEnum_UNSPECIFIED FeedItemQualityApprovalStatusEnum_FeedItemQualityApprovalStatus = 0
	// Used for return value only. Represents value unknown in this version.
	FeedItemQualityApprovalStatusEnum_UNKNOWN FeedItemQualityApprovalStatusEnum_FeedItemQualityApprovalStatus = 1
	// Meets all quality expectations.
	FeedItemQualityApprovalStatusEnum_APPROVED FeedItemQualityApprovalStatusEnum_FeedItemQualityApprovalStatus = 2
	// Does not meet some quality expectations. The specific reason is found in
	// the quality_disapproval_reasons field.
	FeedItemQualityApprovalStatusEnum_DISAPPROVED FeedItemQualityApprovalStatusEnum_FeedItemQualityApprovalStatus = 3
)

var FeedItemQualityApprovalStatusEnum_FeedItemQualityApprovalStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "APPROVED",
	3: "DISAPPROVED",
}

var FeedItemQualityApprovalStatusEnum_FeedItemQualityApprovalStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"APPROVED":    2,
	"DISAPPROVED": 3,
}

func (x FeedItemQualityApprovalStatusEnum_FeedItemQualityApprovalStatus) String() string {
	return proto.EnumName(FeedItemQualityApprovalStatusEnum_FeedItemQualityApprovalStatus_name, int32(x))
}

func (FeedItemQualityApprovalStatusEnum_FeedItemQualityApprovalStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fc4171234ccf44c4, []int{0, 0}
}

// Container for enum describing possible quality evaluation approval statuses
// of a feed item.
type FeedItemQualityApprovalStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeedItemQualityApprovalStatusEnum) Reset()         { *m = FeedItemQualityApprovalStatusEnum{} }
func (m *FeedItemQualityApprovalStatusEnum) String() string { return proto.CompactTextString(m) }
func (*FeedItemQualityApprovalStatusEnum) ProtoMessage()    {}
func (*FeedItemQualityApprovalStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc4171234ccf44c4, []int{0}
}

func (m *FeedItemQualityApprovalStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedItemQualityApprovalStatusEnum.Unmarshal(m, b)
}
func (m *FeedItemQualityApprovalStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedItemQualityApprovalStatusEnum.Marshal(b, m, deterministic)
}
func (m *FeedItemQualityApprovalStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedItemQualityApprovalStatusEnum.Merge(m, src)
}
func (m *FeedItemQualityApprovalStatusEnum) XXX_Size() int {
	return xxx_messageInfo_FeedItemQualityApprovalStatusEnum.Size(m)
}
func (m *FeedItemQualityApprovalStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedItemQualityApprovalStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_FeedItemQualityApprovalStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.FeedItemQualityApprovalStatusEnum_FeedItemQualityApprovalStatus", FeedItemQualityApprovalStatusEnum_FeedItemQualityApprovalStatus_name, FeedItemQualityApprovalStatusEnum_FeedItemQualityApprovalStatus_value)
	proto.RegisterType((*FeedItemQualityApprovalStatusEnum)(nil), "google.ads.googleads.v3.enums.FeedItemQualityApprovalStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/feed_item_quality_approval_status.proto", fileDescriptor_fc4171234ccf44c4)
}

var fileDescriptor_fc4171234ccf44c4 = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0xb5, 0x1d, 0xa8, 0x64, 0x82, 0xa3, 0x47, 0x71, 0x87, 0xed, 0x07, 0xa4, 0x87, 0xdc, 0xe2,
	0x29, 0x73, 0xdd, 0x28, 0x42, 0x57, 0x1d, 0xab, 0x20, 0x85, 0x12, 0x4d, 0x0c, 0x85, 0x36, 0xa9,
	0x4b, 0x3a, 0xf0, 0xe8, 0x5f, 0xf1, 0xe8, 0x4f, 0xf1, 0xa7, 0xf8, 0x07, 0xbc, 0x4a, 0x93, 0xad,
	0x37, 0x77, 0x09, 0x8f, 0x7c, 0xef, 0x7b, 0xef, 0x7b, 0x0f, 0x44, 0x42, 0x29, 0x51, 0xf1, 0x90,
	0x32, 0x1d, 0x3a, 0xd8, 0xa1, 0x1d, 0x0a, 0xb9, 0x6c, 0x6b, 0x1d, 0xbe, 0x72, 0xce, 0x8a, 0xd2,
	0xf0, 0xba, 0x78, 0x6b, 0x69, 0x55, 0x9a, 0xf7, 0x82, 0x36, 0xcd, 0x56, 0xed, 0x68, 0x55, 0x68,
	0x43, 0x4d, 0xab, 0x61, 0xb3, 0x55, 0x46, 0x05, 0x63, 0xb7, 0x0b, 0x29, 0xd3, 0xb0, 0x97, 0x81,
	0x3b, 0x04, 0xad, 0xcc, 0xd5, 0xf5, 0xc1, 0xa5, 0x29, 0x43, 0x2a, 0xa5, 0x32, 0xd4, 0x94, 0x4a,
	0xee, 0x97, 0xa7, 0x1f, 0x1e, 0x98, 0x2c, 0x38, 0x67, 0xb1, 0xe1, 0xf5, 0xbd, 0xb3, 0x21, 0x7b,
	0x97, 0xb5, 0x35, 0x89, 0x64, 0x5b, 0x4f, 0x73, 0x30, 0x3e, 0x4a, 0x0a, 0x2e, 0xc1, 0x70, 0x93,
	0xac, 0xd3, 0xe8, 0x36, 0x5e, 0xc4, 0xd1, 0x7c, 0x74, 0x12, 0x0c, 0xc1, 0xd9, 0x26, 0xb9, 0x4b,
	0x56, 0x8f, 0xc9, 0xc8, 0x0b, 0x2e, 0xc0, 0x39, 0x49, 0xd3, 0x87, 0x55, 0x16, 0xcd, 0x47, 0x7e,
	0xc7, 0x9d, 0xc7, 0xeb, 0xfe, 0x63, 0x30, 0xfb, 0xf5, 0xc0, 0xe4, 0x45, 0xd5, 0xf0, 0x68, 0x8e,
	0xd9, 0xf4, 0xe8, 0x05, 0x69, 0x97, 0x26, 0xf5, 0x9e, 0x66, 0x7b, 0x11, 0xa1, 0x2a, 0x2a, 0x05,
	0x54, 0x5b, 0x11, 0x0a, 0x2e, 0x6d, 0xd6, 0x43, 0xc7, 0x4d, 0xa9, 0xff, 0xa9, 0xfc, 0xc6, 0xbe,
	0x9f, 0xfe, 0x60, 0x49, 0xc8, 0x97, 0x3f, 0x5e, 0x3a, 0x29, 0xc2, 0x34, 0x74, 0xb0, 0x43, 0x19,
	0x82, 0x5d, 0x25, 0xfa, 0xfb, 0x30, 0xcf, 0x09, 0xd3, 0x79, 0x3f, 0xcf, 0x33, 0x94, 0xdb, 0xf9,
	0x8f, 0x3f, 0x71, 0x9f, 0x18, 0x13, 0xa6, 0x31, 0xee, 0x19, 0x18, 0x67, 0x08, 0x63, 0xcb, 0x79,
	0x3e, 0xb5, 0x87, 0xa1, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x65, 0x71, 0x05, 0xda, 0x0a, 0x02,
	0x00, 0x00,
}
