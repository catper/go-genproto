// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/extension_type.proto

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

// Possible data types for an extension in an extension setting.
type ExtensionTypeEnum_ExtensionType int32

const (
	// Not specified.
	ExtensionTypeEnum_UNSPECIFIED ExtensionTypeEnum_ExtensionType = 0
	// Used for return value only. Represents value unknown in this version.
	ExtensionTypeEnum_UNKNOWN ExtensionTypeEnum_ExtensionType = 1
	// None.
	ExtensionTypeEnum_NONE ExtensionTypeEnum_ExtensionType = 2
	// App.
	ExtensionTypeEnum_APP ExtensionTypeEnum_ExtensionType = 3
	// Call.
	ExtensionTypeEnum_CALL ExtensionTypeEnum_ExtensionType = 4
	// Callout.
	ExtensionTypeEnum_CALLOUT ExtensionTypeEnum_ExtensionType = 5
	// Message.
	ExtensionTypeEnum_MESSAGE ExtensionTypeEnum_ExtensionType = 6
	// Price.
	ExtensionTypeEnum_PRICE ExtensionTypeEnum_ExtensionType = 7
	// Promotion.
	ExtensionTypeEnum_PROMOTION ExtensionTypeEnum_ExtensionType = 8
	// Sitelink.
	ExtensionTypeEnum_SITELINK ExtensionTypeEnum_ExtensionType = 10
	// Structured snippet.
	ExtensionTypeEnum_STRUCTURED_SNIPPET ExtensionTypeEnum_ExtensionType = 11
	// Location.
	ExtensionTypeEnum_LOCATION ExtensionTypeEnum_ExtensionType = 12
	// Affiliate location.
	ExtensionTypeEnum_AFFILIATE_LOCATION ExtensionTypeEnum_ExtensionType = 13
	// Hotel callout
	ExtensionTypeEnum_HOTEL_CALLOUT ExtensionTypeEnum_ExtensionType = 15
)

var ExtensionTypeEnum_ExtensionType_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "NONE",
	3:  "APP",
	4:  "CALL",
	5:  "CALLOUT",
	6:  "MESSAGE",
	7:  "PRICE",
	8:  "PROMOTION",
	10: "SITELINK",
	11: "STRUCTURED_SNIPPET",
	12: "LOCATION",
	13: "AFFILIATE_LOCATION",
	15: "HOTEL_CALLOUT",
}

var ExtensionTypeEnum_ExtensionType_value = map[string]int32{
	"UNSPECIFIED":        0,
	"UNKNOWN":            1,
	"NONE":               2,
	"APP":                3,
	"CALL":               4,
	"CALLOUT":            5,
	"MESSAGE":            6,
	"PRICE":              7,
	"PROMOTION":          8,
	"SITELINK":           10,
	"STRUCTURED_SNIPPET": 11,
	"LOCATION":           12,
	"AFFILIATE_LOCATION": 13,
	"HOTEL_CALLOUT":      15,
}

func (x ExtensionTypeEnum_ExtensionType) String() string {
	return proto.EnumName(ExtensionTypeEnum_ExtensionType_name, int32(x))
}

func (ExtensionTypeEnum_ExtensionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_db17f887bbab658a, []int{0, 0}
}

// Container for enum describing possible data types for an extension in an
// extension setting.
type ExtensionTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExtensionTypeEnum) Reset()         { *m = ExtensionTypeEnum{} }
func (m *ExtensionTypeEnum) String() string { return proto.CompactTextString(m) }
func (*ExtensionTypeEnum) ProtoMessage()    {}
func (*ExtensionTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_db17f887bbab658a, []int{0}
}

func (m *ExtensionTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtensionTypeEnum.Unmarshal(m, b)
}
func (m *ExtensionTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtensionTypeEnum.Marshal(b, m, deterministic)
}
func (m *ExtensionTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtensionTypeEnum.Merge(m, src)
}
func (m *ExtensionTypeEnum) XXX_Size() int {
	return xxx_messageInfo_ExtensionTypeEnum.Size(m)
}
func (m *ExtensionTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtensionTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ExtensionTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.ExtensionTypeEnum_ExtensionType", ExtensionTypeEnum_ExtensionType_name, ExtensionTypeEnum_ExtensionType_value)
	proto.RegisterType((*ExtensionTypeEnum)(nil), "google.ads.googleads.v3.enums.ExtensionTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/extension_type.proto", fileDescriptor_db17f887bbab658a)
}

var fileDescriptor_db17f887bbab658a = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0x41, 0x6b, 0xdb, 0x30,
	0x18, 0x9d, 0x93, 0xb6, 0x49, 0x95, 0x86, 0xaa, 0x3a, 0xec, 0x30, 0xd6, 0x43, 0xfb, 0x03, 0x64,
	0x98, 0x6f, 0xda, 0x49, 0x71, 0x95, 0x4c, 0xd4, 0x95, 0x44, 0x2c, 0x67, 0x30, 0x02, 0xc1, 0x9b,
	0x8d, 0x09, 0x34, 0x92, 0xa9, 0xdc, 0xb2, 0xfe, 0x9d, 0x1d, 0xf7, 0x53, 0xf6, 0x1f, 0x76, 0x1e,
	0xec, 0xba, 0x3f, 0x30, 0x64, 0xcf, 0x81, 0x1c, 0xd6, 0x8b, 0x78, 0x7a, 0xdf, 0x7b, 0x1f, 0x1f,
	0xef, 0x81, 0x77, 0x95, 0xb5, 0xd5, 0x7d, 0x19, 0xe6, 0x85, 0x0b, 0x3b, 0xe8, 0xd1, 0x53, 0x14,
	0x96, 0xe6, 0x71, 0xe7, 0xc2, 0xf2, 0x6b, 0x53, 0x1a, 0xb7, 0xb5, 0x66, 0xd3, 0x3c, 0xd7, 0x25,
	0xae, 0x1f, 0x6c, 0x63, 0xd1, 0x65, 0x27, 0xc4, 0x79, 0xe1, 0xf0, 0xde, 0x83, 0x9f, 0x22, 0xdc,
	0x7a, 0xde, 0xbc, 0xed, 0x57, 0xd6, 0xdb, 0x30, 0x37, 0xc6, 0x36, 0x79, 0xb3, 0xb5, 0xc6, 0x75,
	0xe6, 0xeb, 0x3f, 0x01, 0xb8, 0x60, 0xfd, 0x56, 0xfd, 0x5c, 0x97, 0xcc, 0x3c, 0xee, 0xae, 0x7f,
	0x06, 0x60, 0x7a, 0xc0, 0xa2, 0x73, 0x30, 0xc9, 0x44, 0xaa, 0x58, 0xcc, 0xe7, 0x9c, 0xdd, 0xc0,
	0x57, 0x68, 0x02, 0x46, 0x99, 0xb8, 0x15, 0xf2, 0xa3, 0x80, 0x01, 0x1a, 0x83, 0x23, 0x21, 0x05,
	0x83, 0x03, 0x34, 0x02, 0x43, 0xaa, 0x14, 0x1c, 0x7a, 0x2a, 0xa6, 0x49, 0x02, 0x8f, 0xbc, 0xd2,
	0x23, 0x99, 0x69, 0x78, 0xec, 0x3f, 0x77, 0x2c, 0x4d, 0xe9, 0x82, 0xc1, 0x13, 0x74, 0x0a, 0x8e,
	0xd5, 0x92, 0xc7, 0x0c, 0x8e, 0xd0, 0x14, 0x9c, 0xaa, 0xa5, 0xbc, 0x93, 0x9a, 0x4b, 0x01, 0xc7,
	0xe8, 0x0c, 0x8c, 0x53, 0xae, 0x59, 0xc2, 0xc5, 0x2d, 0x04, 0xe8, 0x35, 0x40, 0xa9, 0x5e, 0x66,
	0xb1, 0xce, 0x96, 0xec, 0x66, 0x93, 0x0a, 0xae, 0x14, 0xd3, 0x70, 0xe2, 0x55, 0x89, 0x8c, 0x69,
	0xeb, 0x39, 0xf3, 0x2a, 0x3a, 0x9f, 0xf3, 0x84, 0x53, 0xcd, 0x36, 0x7b, 0x7e, 0x8a, 0x2e, 0xc0,
	0xf4, 0x83, 0xd4, 0x2c, 0xd9, 0xf4, 0x57, 0x9c, 0xcf, 0x7e, 0x05, 0xe0, 0xea, 0x8b, 0xdd, 0xe1,
	0x17, 0x93, 0x9b, 0xa1, 0x83, 0x08, 0x94, 0xcf, 0x4b, 0x05, 0x9f, 0x66, 0xff, 0x4c, 0x95, 0xbd,
	0xcf, 0x4d, 0x85, 0xed, 0x43, 0x15, 0x56, 0xa5, 0x69, 0xd3, 0xec, 0x2b, 0xab, 0xb7, 0xee, 0x3f,
	0x0d, 0xbe, 0x6f, 0xdf, 0x6f, 0x83, 0xe1, 0x82, 0xd2, 0xef, 0x83, 0xcb, 0x45, 0xb7, 0x8a, 0x16,
	0x0e, 0x77, 0xd0, 0xa3, 0x55, 0x84, 0x7d, 0x09, 0xee, 0x47, 0x3f, 0x5f, 0xd3, 0xc2, 0xad, 0xf7,
	0xf3, 0xf5, 0x2a, 0x5a, 0xb7, 0xf3, 0xdf, 0x83, 0xab, 0x8e, 0x24, 0x84, 0x16, 0x8e, 0x90, 0xbd,
	0x82, 0x90, 0x55, 0x44, 0x48, 0xab, 0xf9, 0x7c, 0xd2, 0x1e, 0x16, 0xfd, 0x0d, 0x00, 0x00, 0xff,
	0xff, 0x09, 0x27, 0xaa, 0xa1, 0x59, 0x02, 0x00, 0x00,
}
