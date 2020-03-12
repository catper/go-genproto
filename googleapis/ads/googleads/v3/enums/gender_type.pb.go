// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/gender_type.proto

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

// The type of demographic genders (e.g. female).
type GenderTypeEnum_GenderType int32

const (
	// Not specified.
	GenderTypeEnum_UNSPECIFIED GenderTypeEnum_GenderType = 0
	// Used for return value only. Represents value unknown in this version.
	GenderTypeEnum_UNKNOWN GenderTypeEnum_GenderType = 1
	// Male.
	GenderTypeEnum_MALE GenderTypeEnum_GenderType = 10
	// Female.
	GenderTypeEnum_FEMALE GenderTypeEnum_GenderType = 11
	// Undetermined gender.
	GenderTypeEnum_UNDETERMINED GenderTypeEnum_GenderType = 20
)

var GenderTypeEnum_GenderType_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	10: "MALE",
	11: "FEMALE",
	20: "UNDETERMINED",
}

var GenderTypeEnum_GenderType_value = map[string]int32{
	"UNSPECIFIED":  0,
	"UNKNOWN":      1,
	"MALE":         10,
	"FEMALE":       11,
	"UNDETERMINED": 20,
}

func (x GenderTypeEnum_GenderType) String() string {
	return proto.EnumName(GenderTypeEnum_GenderType_name, int32(x))
}

func (GenderTypeEnum_GenderType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_51e52a2335997197, []int{0, 0}
}

// Container for enum describing the type of demographic genders.
type GenderTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenderTypeEnum) Reset()         { *m = GenderTypeEnum{} }
func (m *GenderTypeEnum) String() string { return proto.CompactTextString(m) }
func (*GenderTypeEnum) ProtoMessage()    {}
func (*GenderTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_51e52a2335997197, []int{0}
}

func (m *GenderTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenderTypeEnum.Unmarshal(m, b)
}
func (m *GenderTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenderTypeEnum.Marshal(b, m, deterministic)
}
func (m *GenderTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenderTypeEnum.Merge(m, src)
}
func (m *GenderTypeEnum) XXX_Size() int {
	return xxx_messageInfo_GenderTypeEnum.Size(m)
}
func (m *GenderTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_GenderTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_GenderTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.GenderTypeEnum_GenderType", GenderTypeEnum_GenderType_name, GenderTypeEnum_GenderType_value)
	proto.RegisterType((*GenderTypeEnum)(nil), "google.ads.googleads.v3.enums.GenderTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/gender_type.proto", fileDescriptor_51e52a2335997197)
}

var fileDescriptor_51e52a2335997197 = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0x4f, 0x4f, 0xf2, 0x30,
	0x1c, 0x7e, 0xe1, 0x35, 0x68, 0x7e, 0x18, 0x59, 0x1a, 0x4f, 0x46, 0x0e, 0xf0, 0x01, 0xba, 0x43,
	0x6f, 0xf5, 0x54, 0xa4, 0x10, 0xa2, 0x54, 0x82, 0x80, 0x89, 0x59, 0x62, 0xa6, 0x6d, 0x1a, 0x12,
	0xd6, 0x2e, 0xeb, 0x20, 0xe1, 0xeb, 0x78, 0xf4, 0xa3, 0xf8, 0x39, 0x3c, 0xf9, 0x29, 0xcc, 0x3a,
	0x36, 0x4e, 0x7a, 0x59, 0x9e, 0xfc, 0x9e, 0x3f, 0x7b, 0xfa, 0x40, 0xa8, 0xad, 0xd5, 0x1b, 0x15,
	0xc6, 0xd2, 0x1d, 0x60, 0x81, 0x76, 0x24, 0x54, 0x66, 0x9b, 0xb8, 0x50, 0x2b, 0x23, 0x55, 0xf6,
	0x92, 0xef, 0x53, 0x85, 0xd3, 0xcc, 0xe6, 0x16, 0x75, 0x4b, 0x15, 0x8e, 0xa5, 0xc3, 0xb5, 0x01,
	0xef, 0x08, 0xf6, 0x86, 0xab, 0xeb, 0x2a, 0x2f, 0x5d, 0x87, 0xb1, 0x31, 0x36, 0x8f, 0xf3, 0xb5,
	0x35, 0xae, 0x34, 0xf7, 0x25, 0x5c, 0x8c, 0x7d, 0xe2, 0x62, 0x9f, 0x2a, 0x6e, 0xb6, 0x49, 0x7f,
	0x0e, 0x70, 0xbc, 0xa0, 0x0e, 0xb4, 0x97, 0xe2, 0x71, 0xc6, 0x6f, 0x27, 0xa3, 0x09, 0x1f, 0x06,
	0xff, 0x50, 0x1b, 0x4e, 0x97, 0xe2, 0x4e, 0x3c, 0x3c, 0x89, 0xa0, 0x81, 0xce, 0xe0, 0x64, 0xca,
	0xee, 0x79, 0x00, 0x08, 0xa0, 0x35, 0xe2, 0x1e, 0xb7, 0x51, 0x00, 0xe7, 0x4b, 0x31, 0xe4, 0x0b,
	0x3e, 0x9f, 0x4e, 0x04, 0x1f, 0x06, 0x97, 0x83, 0xaf, 0x06, 0xf4, 0xde, 0x6c, 0x82, 0xff, 0x6c,
	0x3a, 0xe8, 0x1c, 0xff, 0x3b, 0x2b, 0xca, 0xcd, 0x1a, 0xcf, 0x83, 0x83, 0x43, 0xdb, 0x4d, 0x6c,
	0x34, 0xb6, 0x99, 0x2e, 0x06, 0xf0, 0xd5, 0xab, 0x71, 0xd2, 0xb5, 0xfb, 0x65, 0xab, 0x1b, 0xff,
	0x7d, 0x6f, 0xfe, 0x1f, 0x33, 0xf6, 0xd1, 0xec, 0x8e, 0xcb, 0x28, 0x26, 0x1d, 0x2e, 0x61, 0x81,
	0x56, 0x04, 0x17, 0xaf, 0x76, 0x9f, 0x15, 0x1f, 0x31, 0xe9, 0xa2, 0x9a, 0x8f, 0x56, 0x24, 0xf2,
	0xfc, 0x77, 0xb3, 0x57, 0x1e, 0x29, 0x65, 0xd2, 0x51, 0x5a, 0x2b, 0x28, 0x5d, 0x11, 0x4a, 0xbd,
	0xe6, 0xb5, 0xe5, 0x8b, 0x91, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x97, 0xaf, 0x35, 0x59, 0xc3,
	0x01, 0x00, 0x00,
}
