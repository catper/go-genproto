// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/conversion_action_counting_type.proto

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

// Indicates how conversions for this action will be counted. For more
// information, see https://support.google.com/google-ads/answer/3438531.
type ConversionActionCountingTypeEnum_ConversionActionCountingType int32

const (
	// Not specified.
	ConversionActionCountingTypeEnum_UNSPECIFIED ConversionActionCountingTypeEnum_ConversionActionCountingType = 0
	// Used for return value only. Represents value unknown in this version.
	ConversionActionCountingTypeEnum_UNKNOWN ConversionActionCountingTypeEnum_ConversionActionCountingType = 1
	// Count only one conversion per click.
	ConversionActionCountingTypeEnum_ONE_PER_CLICK ConversionActionCountingTypeEnum_ConversionActionCountingType = 2
	// Count all conversions per click.
	ConversionActionCountingTypeEnum_MANY_PER_CLICK ConversionActionCountingTypeEnum_ConversionActionCountingType = 3
)

var ConversionActionCountingTypeEnum_ConversionActionCountingType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ONE_PER_CLICK",
	3: "MANY_PER_CLICK",
}

var ConversionActionCountingTypeEnum_ConversionActionCountingType_value = map[string]int32{
	"UNSPECIFIED":    0,
	"UNKNOWN":        1,
	"ONE_PER_CLICK":  2,
	"MANY_PER_CLICK": 3,
}

func (x ConversionActionCountingTypeEnum_ConversionActionCountingType) String() string {
	return proto.EnumName(ConversionActionCountingTypeEnum_ConversionActionCountingType_name, int32(x))
}

func (ConversionActionCountingTypeEnum_ConversionActionCountingType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3b999761ccfaadcc, []int{0, 0}
}

// Container for enum describing the conversion deduplication mode for
// conversion optimizer.
type ConversionActionCountingTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConversionActionCountingTypeEnum) Reset()         { *m = ConversionActionCountingTypeEnum{} }
func (m *ConversionActionCountingTypeEnum) String() string { return proto.CompactTextString(m) }
func (*ConversionActionCountingTypeEnum) ProtoMessage()    {}
func (*ConversionActionCountingTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b999761ccfaadcc, []int{0}
}

func (m *ConversionActionCountingTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConversionActionCountingTypeEnum.Unmarshal(m, b)
}
func (m *ConversionActionCountingTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConversionActionCountingTypeEnum.Marshal(b, m, deterministic)
}
func (m *ConversionActionCountingTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConversionActionCountingTypeEnum.Merge(m, src)
}
func (m *ConversionActionCountingTypeEnum) XXX_Size() int {
	return xxx_messageInfo_ConversionActionCountingTypeEnum.Size(m)
}
func (m *ConversionActionCountingTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ConversionActionCountingTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ConversionActionCountingTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.ConversionActionCountingTypeEnum_ConversionActionCountingType", ConversionActionCountingTypeEnum_ConversionActionCountingType_name, ConversionActionCountingTypeEnum_ConversionActionCountingType_value)
	proto.RegisterType((*ConversionActionCountingTypeEnum)(nil), "google.ads.googleads.v3.enums.ConversionActionCountingTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/conversion_action_counting_type.proto", fileDescriptor_3b999761ccfaadcc)
}

var fileDescriptor_3b999761ccfaadcc = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xcf, 0x4a, 0xc3, 0x30,
	0x1c, 0xb6, 0x1d, 0x28, 0x64, 0xa8, 0xb3, 0x47, 0xd9, 0xc0, 0xed, 0x01, 0xd2, 0x43, 0x6f, 0xf1,
	0x94, 0xc5, 0x3a, 0xc6, 0x34, 0x2b, 0xea, 0x26, 0x4a, 0xa1, 0xc4, 0x36, 0x84, 0xc2, 0x96, 0x94,
	0xa5, 0x1b, 0xec, 0x09, 0x7c, 0x0f, 0x8f, 0x3e, 0x8a, 0x8f, 0xe2, 0x03, 0x78, 0x96, 0x24, 0x5b,
	0xf1, 0xe2, 0x2e, 0xc9, 0xc7, 0xef, 0xcf, 0xf7, 0xfd, 0xbe, 0x0f, 0x10, 0xa1, 0x94, 0x58, 0xf0,
	0x90, 0x15, 0x3a, 0x74, 0xd0, 0xa0, 0x4d, 0x14, 0x72, 0xb9, 0x5e, 0xea, 0x30, 0x57, 0x72, 0xc3,
	0x57, 0xba, 0x54, 0x32, 0x63, 0x79, 0x6d, 0xbe, 0x5c, 0xad, 0x65, 0x5d, 0x4a, 0x91, 0xd5, 0xdb,
	0x8a, 0xc3, 0x6a, 0xa5, 0x6a, 0x15, 0xf4, 0xdc, 0x26, 0x64, 0x85, 0x86, 0x0d, 0x09, 0xdc, 0x44,
	0xd0, 0x92, 0x5c, 0x76, 0xf7, 0x1a, 0x55, 0x19, 0x32, 0x29, 0x55, 0xcd, 0x0c, 0x93, 0x76, 0xcb,
	0x83, 0x77, 0x0f, 0x5c, 0x91, 0x46, 0x06, 0x5b, 0x15, 0xb2, 0x13, 0x79, 0xda, 0x56, 0x3c, 0x96,
	0xeb, 0xe5, 0x20, 0x07, 0xdd, 0x43, 0x33, 0xc1, 0x39, 0x68, 0xcf, 0xe8, 0x63, 0x12, 0x93, 0xf1,
	0xed, 0x38, 0xbe, 0xe9, 0x1c, 0x05, 0x6d, 0x70, 0x32, 0xa3, 0x13, 0x3a, 0x7d, 0xa6, 0x1d, 0x2f,
	0xb8, 0x00, 0xa7, 0x53, 0x1a, 0x67, 0x49, 0xfc, 0x90, 0x91, 0xbb, 0x31, 0x99, 0x74, 0xfc, 0x20,
	0x00, 0x67, 0xf7, 0x98, 0xbe, 0xfc, 0xa9, 0xb5, 0x86, 0x3f, 0x1e, 0xe8, 0xe7, 0x6a, 0x09, 0x0f,
	0xba, 0x19, 0xf6, 0x0f, 0x1d, 0x92, 0x18, 0x4b, 0x89, 0xf7, 0x3a, 0xdc, 0x71, 0x08, 0xb5, 0x60,
	0x52, 0x40, 0xb5, 0x12, 0xa1, 0xe0, 0xd2, 0x1a, 0xde, 0xc7, 0x5c, 0x95, 0xfa, 0x9f, 0xd4, 0xaf,
	0xed, 0xfb, 0xe1, 0xb7, 0x46, 0x18, 0x7f, 0xfa, 0xbd, 0x91, 0xa3, 0xc2, 0x85, 0x86, 0x0e, 0x1a,
	0x34, 0x8f, 0xa0, 0x09, 0x46, 0x7f, 0xed, 0xfb, 0x29, 0x2e, 0x74, 0xda, 0xf4, 0xd3, 0x79, 0x94,
	0xda, 0xfe, 0xb7, 0xdf, 0x77, 0x45, 0x84, 0x70, 0xa1, 0x11, 0x6a, 0x26, 0x10, 0x9a, 0x47, 0x08,
	0xd9, 0x99, 0xb7, 0x63, 0x7b, 0x58, 0xf4, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x79, 0x95, 0x05, 0x66,
	0x0d, 0x02, 0x00, 0x00,
}
