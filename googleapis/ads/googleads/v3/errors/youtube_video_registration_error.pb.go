// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/errors/youtube_video_registration_error.proto

package errors

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

// Enum describing YouTube video registration errors.
type YoutubeVideoRegistrationErrorEnum_YoutubeVideoRegistrationError int32

const (
	// Enum unspecified.
	YoutubeVideoRegistrationErrorEnum_UNSPECIFIED YoutubeVideoRegistrationErrorEnum_YoutubeVideoRegistrationError = 0
	// The received error code is not known in this version.
	YoutubeVideoRegistrationErrorEnum_UNKNOWN YoutubeVideoRegistrationErrorEnum_YoutubeVideoRegistrationError = 1
	// Video to be registered wasn't found.
	YoutubeVideoRegistrationErrorEnum_VIDEO_NOT_FOUND YoutubeVideoRegistrationErrorEnum_YoutubeVideoRegistrationError = 2
	// Video to be registered is not accessible (e.g. private).
	YoutubeVideoRegistrationErrorEnum_VIDEO_NOT_ACCESSIBLE YoutubeVideoRegistrationErrorEnum_YoutubeVideoRegistrationError = 3
	// Video to be registered is not eligible (e.g. mature content).
	YoutubeVideoRegistrationErrorEnum_VIDEO_NOT_ELIGIBLE YoutubeVideoRegistrationErrorEnum_YoutubeVideoRegistrationError = 4
)

var YoutubeVideoRegistrationErrorEnum_YoutubeVideoRegistrationError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "VIDEO_NOT_FOUND",
	3: "VIDEO_NOT_ACCESSIBLE",
	4: "VIDEO_NOT_ELIGIBLE",
}

var YoutubeVideoRegistrationErrorEnum_YoutubeVideoRegistrationError_value = map[string]int32{
	"UNSPECIFIED":          0,
	"UNKNOWN":              1,
	"VIDEO_NOT_FOUND":      2,
	"VIDEO_NOT_ACCESSIBLE": 3,
	"VIDEO_NOT_ELIGIBLE":   4,
}

func (x YoutubeVideoRegistrationErrorEnum_YoutubeVideoRegistrationError) String() string {
	return proto.EnumName(YoutubeVideoRegistrationErrorEnum_YoutubeVideoRegistrationError_name, int32(x))
}

func (YoutubeVideoRegistrationErrorEnum_YoutubeVideoRegistrationError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_37f72673454cf2d4, []int{0, 0}
}

// Container for enum describing YouTube video registration errors.
type YoutubeVideoRegistrationErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *YoutubeVideoRegistrationErrorEnum) Reset()         { *m = YoutubeVideoRegistrationErrorEnum{} }
func (m *YoutubeVideoRegistrationErrorEnum) String() string { return proto.CompactTextString(m) }
func (*YoutubeVideoRegistrationErrorEnum) ProtoMessage()    {}
func (*YoutubeVideoRegistrationErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_37f72673454cf2d4, []int{0}
}

func (m *YoutubeVideoRegistrationErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_YoutubeVideoRegistrationErrorEnum.Unmarshal(m, b)
}
func (m *YoutubeVideoRegistrationErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_YoutubeVideoRegistrationErrorEnum.Marshal(b, m, deterministic)
}
func (m *YoutubeVideoRegistrationErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_YoutubeVideoRegistrationErrorEnum.Merge(m, src)
}
func (m *YoutubeVideoRegistrationErrorEnum) XXX_Size() int {
	return xxx_messageInfo_YoutubeVideoRegistrationErrorEnum.Size(m)
}
func (m *YoutubeVideoRegistrationErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_YoutubeVideoRegistrationErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_YoutubeVideoRegistrationErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.errors.YoutubeVideoRegistrationErrorEnum_YoutubeVideoRegistrationError", YoutubeVideoRegistrationErrorEnum_YoutubeVideoRegistrationError_name, YoutubeVideoRegistrationErrorEnum_YoutubeVideoRegistrationError_value)
	proto.RegisterType((*YoutubeVideoRegistrationErrorEnum)(nil), "google.ads.googleads.v3.errors.YoutubeVideoRegistrationErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/errors/youtube_video_registration_error.proto", fileDescriptor_37f72673454cf2d4)
}

var fileDescriptor_37f72673454cf2d4 = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x4d, 0x2a, 0x0a, 0xdb, 0x43, 0xc3, 0x2a, 0x22, 0xa2, 0x05, 0xf3, 0x00, 0x9b, 0x43,
	0x6e, 0xeb, 0x29, 0x6d, 0xb6, 0x25, 0x58, 0x92, 0x62, 0x6d, 0x44, 0x09, 0x84, 0xd4, 0x84, 0x25,
	0xd0, 0x66, 0xca, 0x6e, 0x5a, 0xf0, 0xee, 0x93, 0x78, 0xec, 0xa3, 0xf8, 0x28, 0xbe, 0x83, 0x20,
	0xd9, 0xb5, 0xad, 0x17, 0x7b, 0xca, 0xcf, 0xe4, 0x9b, 0xff, 0x9f, 0xd9, 0x41, 0x8c, 0x03, 0xf0,
	0x79, 0xe1, 0x64, 0xb9, 0x74, 0xb4, 0x6c, 0xd4, 0xda, 0x75, 0x0a, 0x21, 0x40, 0x48, 0xe7, 0x0d,
	0x56, 0xf5, 0x6a, 0x56, 0xa4, 0xeb, 0x32, 0x2f, 0x20, 0x15, 0x05, 0x2f, 0x65, 0x2d, 0xb2, 0xba,
	0x84, 0x2a, 0x55, 0x04, 0x59, 0x0a, 0xa8, 0x01, 0x77, 0x75, 0x2f, 0xc9, 0x72, 0x49, 0x76, 0x36,
	0x64, 0xed, 0x12, 0x6d, 0x73, 0x75, 0xbd, 0x8d, 0x59, 0x96, 0x4e, 0x56, 0x55, 0x50, 0x2b, 0x0b,
	0xa9, 0xbb, 0xed, 0x8d, 0x81, 0x6e, 0x9f, 0x75, 0x50, 0xdc, 0xe4, 0x3c, 0xfc, 0x89, 0x61, 0x8d,
	0x01, 0xab, 0x56, 0x0b, 0xfb, 0xdd, 0x40, 0x37, 0x07, 0x29, 0xdc, 0x41, 0xed, 0x69, 0x38, 0x19,
	0xb3, 0x7e, 0x30, 0x08, 0x98, 0x6f, 0x1d, 0xe1, 0x36, 0x3a, 0x9d, 0x86, 0xf7, 0x61, 0xf4, 0x14,
	0x5a, 0x06, 0x3e, 0x43, 0x9d, 0x38, 0xf0, 0x59, 0x94, 0x86, 0xd1, 0x63, 0x3a, 0x88, 0xa6, 0xa1,
	0x6f, 0x99, 0xf8, 0x12, 0x9d, 0xef, 0x8b, 0x5e, 0xbf, 0xcf, 0x26, 0x93, 0xa0, 0x37, 0x62, 0x56,
	0x0b, 0x5f, 0x20, 0xbc, 0xff, 0xc3, 0x46, 0xc1, 0x50, 0xd5, 0x8f, 0x7b, 0xdf, 0x06, 0xb2, 0x5f,
	0x61, 0x41, 0x0e, 0x6f, 0xdc, 0xb3, 0x0f, 0x8e, 0x3a, 0x6e, 0xf6, 0x1e, 0x1b, 0x2f, 0xfe, 0xaf,
	0x0b, 0x87, 0x79, 0x56, 0x71, 0x02, 0x82, 0x3b, 0xbc, 0xa8, 0xd4, 0xab, 0x6c, 0xcf, 0xb1, 0x2c,
	0xe5, 0x7f, 0xd7, 0xb9, 0xd3, 0x9f, 0x0f, 0xb3, 0x35, 0xf4, 0xbc, 0x8d, 0xd9, 0x1d, 0x6a, 0x33,
	0x2f, 0x97, 0x44, 0xcb, 0x46, 0xc5, 0x2e, 0x51, 0x91, 0xf2, 0x73, 0x0b, 0x24, 0x5e, 0x2e, 0x93,
	0x1d, 0x90, 0xc4, 0x6e, 0xa2, 0x81, 0x2f, 0xd3, 0xd6, 0x55, 0x4a, 0xbd, 0x5c, 0x52, 0xba, 0x43,
	0x28, 0x8d, 0x5d, 0x4a, 0x35, 0x34, 0x3b, 0x51, 0xd3, 0xb9, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xf2, 0xe2, 0x8c, 0xd1, 0x3a, 0x02, 0x00, 0x00,
}
