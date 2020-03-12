// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/ad_group_ad_status.proto

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
	return fileDescriptor_39943830242fa845, []int{0, 0}
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
	return fileDescriptor_39943830242fa845, []int{0}
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
	proto.RegisterEnum("google.ads.googleads.v2.enums.AdGroupAdStatusEnum_AdGroupAdStatus", AdGroupAdStatusEnum_AdGroupAdStatus_name, AdGroupAdStatusEnum_AdGroupAdStatus_value)
	proto.RegisterType((*AdGroupAdStatusEnum)(nil), "google.ads.googleads.v2.enums.AdGroupAdStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/ad_group_ad_status.proto", fileDescriptor_39943830242fa845)
}

var fileDescriptor_39943830242fa845 = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xcf, 0x4a, 0xc3, 0x30,
	0x18, 0x77, 0x9d, 0x4c, 0xc8, 0x0e, 0x2b, 0xd5, 0x93, 0xb8, 0xc3, 0xf6, 0x00, 0x09, 0x54, 0xf0,
	0x10, 0x4f, 0xa9, 0x8d, 0x63, 0xa8, 0x5d, 0x71, 0xb4, 0x82, 0x14, 0x46, 0x34, 0x23, 0x0c, 0xda,
	0xa4, 0x34, 0xed, 0x1e, 0xc8, 0xa3, 0x8f, 0xe2, 0x9b, 0xe8, 0x53, 0x48, 0xd2, 0xb5, 0x87, 0x81,
	0x5e, 0xca, 0xaf, 0xdf, 0xef, 0x4f, 0x7e, 0xdf, 0x07, 0x6e, 0x84, 0x52, 0x22, 0xdf, 0x22, 0xc6,
	0x35, 0x6a, 0xa1, 0x41, 0x7b, 0x1f, 0x6d, 0x65, 0x53, 0x68, 0xc4, 0xf8, 0x46, 0x54, 0xaa, 0x29,
	0x37, 0x8c, 0x6f, 0x74, 0xcd, 0xea, 0x46, 0xc3, 0xb2, 0x52, 0xb5, 0xf2, 0xa6, 0xad, 0x18, 0x32,
	0xae, 0x61, 0xef, 0x83, 0x7b, 0x1f, 0x5a, 0xdf, 0xe5, 0x55, 0x17, 0x5b, 0xee, 0x10, 0x93, 0x52,
	0xd5, 0xac, 0xde, 0x29, 0x79, 0x30, 0xcf, 0x73, 0x70, 0x4e, 0xf8, 0xc2, 0xe4, 0x12, 0xbe, 0xb6,
	0xa9, 0x54, 0x36, 0xc5, 0x3c, 0x01, 0x93, 0xa3, 0xb1, 0x37, 0x01, 0xe3, 0x24, 0x5a, 0xc7, 0xf4,
	0x6e, 0x79, 0xbf, 0xa4, 0xa1, 0x7b, 0xe2, 0x8d, 0xc1, 0x59, 0x12, 0x3d, 0x44, 0xab, 0x97, 0xc8,
	0x1d, 0x98, 0x1f, 0x1a, 0x91, 0xe0, 0x91, 0x86, 0xae, 0xe3, 0x01, 0x30, 0x8a, 0x49, 0xb2, 0xa6,
	0xa1, 0x3b, 0x34, 0xc4, 0x33, 0x7d, 0x5a, 0xa5, 0x34, 0x74, 0x4f, 0x83, 0xef, 0x01, 0x98, 0xbd,
	0xab, 0x02, 0xfe, 0xdb, 0x38, 0xb8, 0x38, 0x7a, 0x3a, 0x36, 0x4d, 0xe3, 0xc1, 0x6b, 0x70, 0xb0,
	0x09, 0x95, 0x33, 0x29, 0xa0, 0xaa, 0x04, 0x12, 0x5b, 0x69, 0xf7, 0xe8, 0x0e, 0x56, 0xee, 0xf4,
	0x1f, 0xf7, 0xbb, 0xb5, 0xdf, 0x0f, 0x67, 0xb8, 0x20, 0xe4, 0xd3, 0x99, 0x2e, 0xda, 0x28, 0xc2,
	0x35, 0x6c, 0xa1, 0x41, 0xa9, 0x0f, 0xcd, 0xf6, 0xfa, 0xab, 0xe3, 0x33, 0xc2, 0x75, 0xd6, 0xf3,
	0x59, 0xea, 0x67, 0x96, 0xff, 0x71, 0x66, 0xed, 0x10, 0x63, 0xc2, 0x35, 0xc6, 0xbd, 0x02, 0xe3,
	0xd4, 0xc7, 0xd8, 0x6a, 0xde, 0x46, 0xb6, 0xd8, 0xf5, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7b,
	0x88, 0x35, 0xbf, 0xd7, 0x01, 0x00, 0x00,
}
