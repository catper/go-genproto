// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/ad_group_ad_status.proto

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

// The possible statuses of an AdGroupAd.
type AdGroupAdStatusEnum_AdGroupAdStatus int32

const (
	// No value has been specified.
	AdGroupAdStatusEnum_UNSPECIFIED AdGroupAdStatusEnum_AdGroupAdStatus = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	AdGroupAdStatusEnum_UNKNOWN AdGroupAdStatusEnum_AdGroupAdStatus = 1
	// The ad group ad is enabled.
	AdGroupAdStatusEnum_ENABLED AdGroupAdStatusEnum_AdGroupAdStatus = 2
	// The ad group ad is paused.
	AdGroupAdStatusEnum_PAUSED AdGroupAdStatusEnum_AdGroupAdStatus = 3
	// The ad group ad is removed.
	AdGroupAdStatusEnum_REMOVED AdGroupAdStatusEnum_AdGroupAdStatus = 4
)

var AdGroupAdStatusEnum_AdGroupAdStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ENABLED",
	3: "PAUSED",
	4: "REMOVED",
}

var AdGroupAdStatusEnum_AdGroupAdStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ENABLED":     2,
	"PAUSED":      3,
	"REMOVED":     4,
}

func (x AdGroupAdStatusEnum_AdGroupAdStatus) String() string {
	return proto.EnumName(AdGroupAdStatusEnum_AdGroupAdStatus_name, int32(x))
}

func (AdGroupAdStatusEnum_AdGroupAdStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_469b781fb0018c61, []int{0, 0}
}

// Container for enum describing possible statuses of an AdGroupAd.
type AdGroupAdStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdGroupAdStatusEnum) Reset()         { *m = AdGroupAdStatusEnum{} }
func (m *AdGroupAdStatusEnum) String() string { return proto.CompactTextString(m) }
func (*AdGroupAdStatusEnum) ProtoMessage()    {}
func (*AdGroupAdStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_469b781fb0018c61, []int{0}
}

func (m *AdGroupAdStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupAdStatusEnum.Unmarshal(m, b)
}
func (m *AdGroupAdStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupAdStatusEnum.Marshal(b, m, deterministic)
}
func (m *AdGroupAdStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupAdStatusEnum.Merge(m, src)
}
func (m *AdGroupAdStatusEnum) XXX_Size() int {
	return xxx_messageInfo_AdGroupAdStatusEnum.Size(m)
}
func (m *AdGroupAdStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupAdStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupAdStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.AdGroupAdStatusEnum_AdGroupAdStatus", AdGroupAdStatusEnum_AdGroupAdStatus_name, AdGroupAdStatusEnum_AdGroupAdStatus_value)
	proto.RegisterType((*AdGroupAdStatusEnum)(nil), "google.ads.googleads.v3.enums.AdGroupAdStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/ad_group_ad_status.proto", fileDescriptor_469b781fb0018c61)
}

var fileDescriptor_469b781fb0018c61 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0x4f, 0x4b, 0xfb, 0x30,
	0x18, 0xfe, 0xad, 0xfb, 0x31, 0x21, 0x3b, 0xac, 0x54, 0x4f, 0xe2, 0x0e, 0xdb, 0x07, 0x48, 0x0e,
	0x01, 0x0f, 0xf1, 0x94, 0xda, 0x38, 0x86, 0xda, 0x15, 0x47, 0x2b, 0x48, 0x61, 0x44, 0x33, 0xc2,
	0x60, 0x4d, 0x4a, 0xd3, 0xee, 0x03, 0x79, 0xf4, 0xa3, 0xf8, 0x4d, 0xf4, 0x53, 0x48, 0xd2, 0xb5,
	0x87, 0x81, 0x5e, 0xca, 0xd3, 0xf7, 0xf9, 0x93, 0xe7, 0x7d, 0xc1, 0xb5, 0xd4, 0x5a, 0xee, 0xb7,
	0x88, 0x0b, 0x83, 0x5a, 0x68, 0xd1, 0x01, 0xa3, 0xad, 0x6a, 0x0a, 0x83, 0xb8, 0xd8, 0xc8, 0x4a,
	0x37, 0xe5, 0x86, 0x8b, 0x8d, 0xa9, 0x79, 0xdd, 0x18, 0x58, 0x56, 0xba, 0xd6, 0xc1, 0xb4, 0x15,
	0x43, 0x2e, 0x0c, 0xec, 0x7d, 0xf0, 0x80, 0xa1, 0xf3, 0x5d, 0x5e, 0x75, 0xb1, 0xe5, 0x0e, 0x71,
	0xa5, 0x74, 0xcd, 0xeb, 0x9d, 0x56, 0x47, 0xf3, 0x7c, 0x0f, 0xce, 0xa9, 0x58, 0xd8, 0x5c, 0x2a,
	0xd6, 0x2e, 0x95, 0xa9, 0xa6, 0x98, 0xa7, 0x60, 0x72, 0x32, 0x0e, 0x26, 0x60, 0x9c, 0xc6, 0xeb,
	0x84, 0xdd, 0x2e, 0xef, 0x96, 0x2c, 0xf2, 0xff, 0x05, 0x63, 0x70, 0x96, 0xc6, 0xf7, 0xf1, 0xea,
	0x39, 0xf6, 0x07, 0xf6, 0x87, 0xc5, 0x34, 0x7c, 0x60, 0x91, 0xef, 0x05, 0x00, 0x8c, 0x12, 0x9a,
	0xae, 0x59, 0xe4, 0x0f, 0x2d, 0xf1, 0xc4, 0x1e, 0x57, 0x19, 0x8b, 0xfc, 0xff, 0xe1, 0xd7, 0x00,
	0xcc, 0xde, 0x74, 0x01, 0xff, 0x6c, 0x1c, 0x5e, 0x9c, 0x3c, 0x9d, 0xd8, 0xa6, 0xc9, 0xe0, 0x25,
	0x3c, 0xda, 0xa4, 0xde, 0x73, 0x25, 0xa1, 0xae, 0x24, 0x92, 0x5b, 0xe5, 0xf6, 0xe8, 0x0e, 0x56,
	0xee, 0xcc, 0x2f, 0xf7, 0xbb, 0x71, 0xdf, 0x77, 0x6f, 0xb8, 0xa0, 0xf4, 0xc3, 0x9b, 0x2e, 0xda,
	0x28, 0x2a, 0x0c, 0x6c, 0xa1, 0x45, 0x19, 0x86, 0x76, 0x7b, 0xf3, 0xd9, 0xf1, 0x39, 0x15, 0x26,
	0xef, 0xf9, 0x3c, 0xc3, 0xb9, 0xe3, 0xbf, 0xbd, 0x59, 0x3b, 0x24, 0x84, 0x0a, 0x43, 0x48, 0xaf,
	0x20, 0x24, 0xc3, 0x84, 0x38, 0xcd, 0xeb, 0xc8, 0x15, 0xc3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x9e, 0xcd, 0x77, 0xa8, 0xd7, 0x01, 0x00, 0x00,
}
