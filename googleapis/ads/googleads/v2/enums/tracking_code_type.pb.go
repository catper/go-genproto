// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/tracking_code_type.proto

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

// The type of the generated tag snippets for tracking conversions.
type TrackingCodeTypeEnum_TrackingCodeType int32

const (
	// Not specified.
	TrackingCodeTypeEnum_UNSPECIFIED TrackingCodeTypeEnum_TrackingCodeType = 0
	// Used for return value only. Represents value unknown in this version.
	TrackingCodeTypeEnum_UNKNOWN TrackingCodeTypeEnum_TrackingCodeType = 1
	// The snippet that is fired as a result of a website page loading.
	TrackingCodeTypeEnum_WEBPAGE TrackingCodeTypeEnum_TrackingCodeType = 2
	// The snippet contains a JavaScript function which fires the tag. This
	// function is typically called from an onClick handler added to a link or
	// button element on the page.
	TrackingCodeTypeEnum_WEBPAGE_ONCLICK TrackingCodeTypeEnum_TrackingCodeType = 3
	// For embedding on a mobile webpage. The snippet contains a JavaScript
	// function which fires the tag.
	TrackingCodeTypeEnum_CLICK_TO_CALL TrackingCodeTypeEnum_TrackingCodeType = 4
	// The snippet that is used to replace the phone number on your website with
	// a Google forwarding number for call tracking purposes.
	TrackingCodeTypeEnum_WEBSITE_CALL TrackingCodeTypeEnum_TrackingCodeType = 5
)

var TrackingCodeTypeEnum_TrackingCodeType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "WEBPAGE",
	3: "WEBPAGE_ONCLICK",
	4: "CLICK_TO_CALL",
	5: "WEBSITE_CALL",
}

var TrackingCodeTypeEnum_TrackingCodeType_value = map[string]int32{
	"UNSPECIFIED":     0,
	"UNKNOWN":         1,
	"WEBPAGE":         2,
	"WEBPAGE_ONCLICK": 3,
	"CLICK_TO_CALL":   4,
	"WEBSITE_CALL":    5,
}

func (x TrackingCodeTypeEnum_TrackingCodeType) String() string {
	return proto.EnumName(TrackingCodeTypeEnum_TrackingCodeType_name, int32(x))
}

func (TrackingCodeTypeEnum_TrackingCodeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bc3aa75f0cf51f59, []int{0, 0}
}

// Container for enum describing the type of the generated tag snippets for
// tracking conversions.
type TrackingCodeTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrackingCodeTypeEnum) Reset()         { *m = TrackingCodeTypeEnum{} }
func (m *TrackingCodeTypeEnum) String() string { return proto.CompactTextString(m) }
func (*TrackingCodeTypeEnum) ProtoMessage()    {}
func (*TrackingCodeTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc3aa75f0cf51f59, []int{0}
}

func (m *TrackingCodeTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrackingCodeTypeEnum.Unmarshal(m, b)
}
func (m *TrackingCodeTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrackingCodeTypeEnum.Marshal(b, m, deterministic)
}
func (m *TrackingCodeTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrackingCodeTypeEnum.Merge(m, src)
}
func (m *TrackingCodeTypeEnum) XXX_Size() int {
	return xxx_messageInfo_TrackingCodeTypeEnum.Size(m)
}
func (m *TrackingCodeTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_TrackingCodeTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_TrackingCodeTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.TrackingCodeTypeEnum_TrackingCodeType", TrackingCodeTypeEnum_TrackingCodeType_name, TrackingCodeTypeEnum_TrackingCodeType_value)
	proto.RegisterType((*TrackingCodeTypeEnum)(nil), "google.ads.googleads.v2.enums.TrackingCodeTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/tracking_code_type.proto", fileDescriptor_bc3aa75f0cf51f59)
}

var fileDescriptor_bc3aa75f0cf51f59 = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xd1, 0x4a, 0xf3, 0x30,
	0x18, 0xfd, 0xdb, 0xfd, 0x2a, 0x64, 0xca, 0x6a, 0xd5, 0x1b, 0x71, 0x17, 0xdb, 0x03, 0xa4, 0x50,
	0xc1, 0x8b, 0x78, 0xd5, 0xd6, 0x38, 0xca, 0x46, 0x57, 0x58, 0xb7, 0x81, 0x14, 0x4a, 0x5d, 0x42,
	0x28, 0x6e, 0x49, 0x59, 0xba, 0xc9, 0x9e, 0xc2, 0x77, 0xf0, 0xd2, 0x47, 0xf1, 0x51, 0xf6, 0x14,
	0xd2, 0x64, 0xdb, 0xc5, 0x40, 0x6f, 0xc2, 0xc9, 0xf7, 0x9d, 0x73, 0xf8, 0xce, 0x01, 0x0f, 0x4c,
	0x08, 0x36, 0xa7, 0x4e, 0x4e, 0xa4, 0xa3, 0x61, 0x8d, 0xd6, 0xae, 0x43, 0xf9, 0x6a, 0x21, 0x9d,
	0x6a, 0x99, 0xcf, 0xde, 0x0a, 0xce, 0xb2, 0x99, 0x20, 0x34, 0xab, 0x36, 0x25, 0x85, 0xe5, 0x52,
	0x54, 0xc2, 0x6e, 0x6b, 0x32, 0xcc, 0x89, 0x84, 0x07, 0x1d, 0x5c, 0xbb, 0x50, 0xe9, 0x6e, 0xef,
	0xf6, 0xb6, 0x65, 0xe1, 0xe4, 0x9c, 0x8b, 0x2a, 0xaf, 0x0a, 0xc1, 0xa5, 0x16, 0x77, 0x3f, 0x0c,
	0x70, 0x9d, 0xec, 0x9c, 0x03, 0x41, 0x68, 0xb2, 0x29, 0x29, 0xe6, 0xab, 0x45, 0xf7, 0x1d, 0x58,
	0xc7, 0x73, 0xbb, 0x05, 0x9a, 0xe3, 0x68, 0x14, 0xe3, 0x20, 0x7c, 0x0e, 0xf1, 0x93, 0xf5, 0xcf,
	0x6e, 0x82, 0xb3, 0x71, 0xd4, 0x8f, 0x86, 0xd3, 0xc8, 0x32, 0xea, 0xcf, 0x14, 0xfb, 0xb1, 0xd7,
	0xc3, 0x96, 0x69, 0x5f, 0x81, 0xd6, 0xee, 0x93, 0x0d, 0xa3, 0x60, 0x10, 0x06, 0x7d, 0xab, 0x61,
	0x5f, 0x82, 0x0b, 0x05, 0xb3, 0x64, 0x98, 0x05, 0xde, 0x60, 0x60, 0xfd, 0xb7, 0x2d, 0x70, 0x3e,
	0xc5, 0xfe, 0x28, 0x4c, 0xb0, 0x9e, 0x9c, 0xf8, 0x5b, 0x03, 0x74, 0x66, 0x62, 0x01, 0xff, 0x4c,
	0xe5, 0xdf, 0x1c, 0x1f, 0x17, 0xd7, 0x71, 0x62, 0xe3, 0xc5, 0xdf, 0xe9, 0x98, 0x98, 0xe7, 0x9c,
	0x41, 0xb1, 0x64, 0x0e, 0xa3, 0x5c, 0x85, 0xdd, 0xb7, 0x5a, 0x16, 0xf2, 0x97, 0x92, 0x1f, 0xd5,
	0xfb, 0x69, 0x36, 0x7a, 0x9e, 0xf7, 0x65, 0xb6, 0x7b, 0xda, 0xca, 0x23, 0x12, 0x6a, 0x58, 0xa3,
	0x89, 0x0b, 0xeb, 0x82, 0xe4, 0xf7, 0x7e, 0x9f, 0x7a, 0x44, 0xa6, 0x87, 0x7d, 0x3a, 0x71, 0x53,
	0xb5, 0xdf, 0x9a, 0x1d, 0x3d, 0x44, 0xc8, 0x23, 0x12, 0xa1, 0x03, 0x03, 0xa1, 0x89, 0x8b, 0x90,
	0xe2, 0xbc, 0x9e, 0xaa, 0xc3, 0xee, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf1, 0xf8, 0x40, 0x53,
	0xfc, 0x01, 0x00, 0x00,
}
