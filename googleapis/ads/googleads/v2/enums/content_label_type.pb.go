// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/content_label_type.proto

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

// Enum listing the content label types supported by ContentLabel criterion.
type ContentLabelTypeEnum_ContentLabelType int32

const (
	// Not specified.
	ContentLabelTypeEnum_UNSPECIFIED ContentLabelTypeEnum_ContentLabelType = 0
	// Used for return value only. Represents value unknown in this version.
	ContentLabelTypeEnum_UNKNOWN ContentLabelTypeEnum_ContentLabelType = 1
	// Sexually suggestive content.
	ContentLabelTypeEnum_SEXUALLY_SUGGESTIVE ContentLabelTypeEnum_ContentLabelType = 2
	// Below the fold placement.
	ContentLabelTypeEnum_BELOW_THE_FOLD ContentLabelTypeEnum_ContentLabelType = 3
	// Parked domain.
	ContentLabelTypeEnum_PARKED_DOMAIN ContentLabelTypeEnum_ContentLabelType = 4
	// Game.
	ContentLabelTypeEnum_GAME ContentLabelTypeEnum_ContentLabelType = 5
	// Juvenile, gross & bizarre content.
	ContentLabelTypeEnum_JUVENILE ContentLabelTypeEnum_ContentLabelType = 6
	// Profanity & rough language.
	ContentLabelTypeEnum_PROFANITY ContentLabelTypeEnum_ContentLabelType = 7
	// Death & tragedy.
	ContentLabelTypeEnum_TRAGEDY ContentLabelTypeEnum_ContentLabelType = 8
	// Video.
	ContentLabelTypeEnum_VIDEO ContentLabelTypeEnum_ContentLabelType = 9
	// Content rating: G.
	ContentLabelTypeEnum_VIDEO_RATING_DV_G ContentLabelTypeEnum_ContentLabelType = 10
	// Content rating: PG.
	ContentLabelTypeEnum_VIDEO_RATING_DV_PG ContentLabelTypeEnum_ContentLabelType = 11
	// Content rating: T.
	ContentLabelTypeEnum_VIDEO_RATING_DV_T ContentLabelTypeEnum_ContentLabelType = 12
	// Content rating: MA.
	ContentLabelTypeEnum_VIDEO_RATING_DV_MA ContentLabelTypeEnum_ContentLabelType = 13
	// Content rating: not yet rated.
	ContentLabelTypeEnum_VIDEO_NOT_YET_RATED ContentLabelTypeEnum_ContentLabelType = 14
	// Embedded video.
	ContentLabelTypeEnum_EMBEDDED_VIDEO ContentLabelTypeEnum_ContentLabelType = 15
	// Live streaming video.
	ContentLabelTypeEnum_LIVE_STREAMING_VIDEO ContentLabelTypeEnum_ContentLabelType = 16
	// Sensitive social issues.
	ContentLabelTypeEnum_SOCIAL_ISSUES ContentLabelTypeEnum_ContentLabelType = 17
)

var ContentLabelTypeEnum_ContentLabelType_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "SEXUALLY_SUGGESTIVE",
	3:  "BELOW_THE_FOLD",
	4:  "PARKED_DOMAIN",
	5:  "GAME",
	6:  "JUVENILE",
	7:  "PROFANITY",
	8:  "TRAGEDY",
	9:  "VIDEO",
	10: "VIDEO_RATING_DV_G",
	11: "VIDEO_RATING_DV_PG",
	12: "VIDEO_RATING_DV_T",
	13: "VIDEO_RATING_DV_MA",
	14: "VIDEO_NOT_YET_RATED",
	15: "EMBEDDED_VIDEO",
	16: "LIVE_STREAMING_VIDEO",
	17: "SOCIAL_ISSUES",
}

var ContentLabelTypeEnum_ContentLabelType_value = map[string]int32{
	"UNSPECIFIED":          0,
	"UNKNOWN":              1,
	"SEXUALLY_SUGGESTIVE":  2,
	"BELOW_THE_FOLD":       3,
	"PARKED_DOMAIN":        4,
	"GAME":                 5,
	"JUVENILE":             6,
	"PROFANITY":            7,
	"TRAGEDY":              8,
	"VIDEO":                9,
	"VIDEO_RATING_DV_G":    10,
	"VIDEO_RATING_DV_PG":   11,
	"VIDEO_RATING_DV_T":    12,
	"VIDEO_RATING_DV_MA":   13,
	"VIDEO_NOT_YET_RATED":  14,
	"EMBEDDED_VIDEO":       15,
	"LIVE_STREAMING_VIDEO": 16,
	"SOCIAL_ISSUES":        17,
}

func (x ContentLabelTypeEnum_ContentLabelType) String() string {
	return proto.EnumName(ContentLabelTypeEnum_ContentLabelType_name, int32(x))
}

func (ContentLabelTypeEnum_ContentLabelType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bc9ccf58e1875500, []int{0, 0}
}

// Container for enum describing content label types in ContentLabel.
type ContentLabelTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContentLabelTypeEnum) Reset()         { *m = ContentLabelTypeEnum{} }
func (m *ContentLabelTypeEnum) String() string { return proto.CompactTextString(m) }
func (*ContentLabelTypeEnum) ProtoMessage()    {}
func (*ContentLabelTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc9ccf58e1875500, []int{0}
}

func (m *ContentLabelTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContentLabelTypeEnum.Unmarshal(m, b)
}
func (m *ContentLabelTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContentLabelTypeEnum.Marshal(b, m, deterministic)
}
func (m *ContentLabelTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentLabelTypeEnum.Merge(m, src)
}
func (m *ContentLabelTypeEnum) XXX_Size() int {
	return xxx_messageInfo_ContentLabelTypeEnum.Size(m)
}
func (m *ContentLabelTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentLabelTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ContentLabelTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.ContentLabelTypeEnum_ContentLabelType", ContentLabelTypeEnum_ContentLabelType_name, ContentLabelTypeEnum_ContentLabelType_value)
	proto.RegisterType((*ContentLabelTypeEnum)(nil), "google.ads.googleads.v2.enums.ContentLabelTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/content_label_type.proto", fileDescriptor_bc9ccf58e1875500)
}

var fileDescriptor_bc9ccf58e1875500 = []byte{
	// 486 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xcb, 0x6a, 0xdb, 0x40,
	0x14, 0xad, 0xed, 0x3c, 0xec, 0x71, 0x9c, 0x8c, 0xa7, 0x49, 0x5b, 0x4a, 0xb3, 0x48, 0x3e, 0x40,
	0x02, 0x17, 0xba, 0x50, 0x57, 0x63, 0xcf, 0xb5, 0x3a, 0x8d, 0x2c, 0x09, 0xeb, 0x91, 0xba, 0x18,
	0x06, 0x25, 0x16, 0xc2, 0x60, 0x6b, 0x44, 0xa4, 0x04, 0xb2, 0xeb, 0xb7, 0x74, 0xd9, 0x4f, 0xe9,
	0x8f, 0x14, 0xf2, 0x15, 0x65, 0x34, 0xb5, 0x17, 0x26, 0xcd, 0x46, 0x1c, 0xee, 0x79, 0x68, 0x2e,
	0xe7, 0xa2, 0x4f, 0x99, 0x94, 0xd9, 0x2a, 0x35, 0x93, 0x45, 0x69, 0x6a, 0xa8, 0xd0, 0xc3, 0xc0,
	0x4c, 0xf3, 0xfb, 0x75, 0x69, 0xde, 0xca, 0xbc, 0x4a, 0xf3, 0x4a, 0xac, 0x92, 0x9b, 0x74, 0x25,
	0xaa, 0xc7, 0x22, 0x35, 0x8a, 0x3b, 0x59, 0x49, 0x72, 0xae, 0xc5, 0x46, 0xb2, 0x28, 0x8d, 0xad,
	0xcf, 0x78, 0x18, 0x18, 0xb5, 0xef, 0xfd, 0x87, 0x4d, 0x6c, 0xb1, 0x34, 0x93, 0x3c, 0x97, 0x55,
	0x52, 0x2d, 0x65, 0x5e, 0x6a, 0xf3, 0xe5, 0x8f, 0x16, 0x3a, 0x1d, 0xe9, 0x64, 0x47, 0x05, 0x87,
	0x8f, 0x45, 0x0a, 0xf9, 0xfd, 0xfa, 0xf2, 0x4f, 0x13, 0xe1, 0x5d, 0x82, 0x9c, 0xa0, 0x6e, 0xe4,
	0x06, 0x3e, 0x8c, 0xf8, 0x98, 0x03, 0xc3, 0xaf, 0x48, 0x17, 0x1d, 0x46, 0xee, 0x95, 0xeb, 0x5d,
	0xbb, 0xb8, 0x41, 0xde, 0xa2, 0xd7, 0x01, 0x7c, 0x8b, 0xa8, 0xe3, 0xcc, 0x44, 0x10, 0xd9, 0x36,
	0x04, 0x21, 0x8f, 0x01, 0x37, 0x09, 0x41, 0xc7, 0x43, 0x70, 0xbc, 0x6b, 0x11, 0x7e, 0x01, 0x31,
	0xf6, 0x1c, 0x86, 0x5b, 0xa4, 0x8f, 0x7a, 0x3e, 0x9d, 0x5e, 0x01, 0x13, 0xcc, 0x9b, 0x50, 0xee,
	0xe2, 0x3d, 0xd2, 0x46, 0x7b, 0x36, 0x9d, 0x00, 0xde, 0x27, 0x47, 0xa8, 0xfd, 0x35, 0x8a, 0xc1,
	0xe5, 0x0e, 0xe0, 0x03, 0xd2, 0x43, 0x1d, 0x7f, 0xea, 0x8d, 0xa9, 0xcb, 0xc3, 0x19, 0x3e, 0x54,
	0xff, 0x0c, 0xa7, 0xd4, 0x06, 0x36, 0xc3, 0x6d, 0xd2, 0x41, 0xfb, 0x31, 0x67, 0xe0, 0xe1, 0x0e,
	0x39, 0x43, 0xfd, 0x1a, 0x8a, 0x29, 0x0d, 0xb9, 0x6b, 0x0b, 0x16, 0x0b, 0x1b, 0x23, 0xf2, 0x06,
	0x91, 0xdd, 0xb1, 0x6f, 0xe3, 0xee, 0x73, 0xf2, 0x10, 0x1f, 0x3d, 0x27, 0x9f, 0x50, 0xdc, 0x53,
	0xcb, 0xe9, 0xb9, 0xeb, 0x85, 0x62, 0x06, 0xa1, 0xe2, 0x81, 0xe1, 0x63, 0xb5, 0x1c, 0x4c, 0x86,
	0xc0, 0x18, 0x30, 0xa1, 0x9f, 0x72, 0x42, 0xde, 0xa1, 0x53, 0x87, 0xc7, 0x20, 0x82, 0x70, 0x0a,
	0x74, 0xa2, 0x62, 0x34, 0x83, 0xd5, 0xda, 0x81, 0x37, 0xe2, 0xd4, 0x11, 0x3c, 0x08, 0x22, 0x08,
	0x70, 0x7f, 0xf8, 0xd4, 0x40, 0x17, 0xb7, 0x72, 0x6d, 0xbc, 0x58, 0xe3, 0xf0, 0x6c, 0xb7, 0x0c,
	0x5f, 0xf5, 0xe7, 0x37, 0xbe, 0x0f, 0xff, 0xf9, 0x32, 0xb9, 0x4a, 0xf2, 0xcc, 0x90, 0x77, 0x99,
	0x99, 0xa5, 0x79, 0xdd, 0xee, 0xe6, 0x8c, 0x8a, 0x65, 0xf9, 0x9f, 0xab, 0xfa, 0x5c, 0x7f, 0x7f,
	0x36, 0x5b, 0x36, 0xa5, 0xbf, 0x9a, 0xe7, 0xb6, 0x8e, 0xa2, 0x8b, 0xd2, 0xd0, 0x50, 0xa1, 0x78,
	0x60, 0xa8, 0x8b, 0x28, 0x7f, 0x6f, 0xf8, 0x39, 0x5d, 0x94, 0xf3, 0x2d, 0x3f, 0x8f, 0x07, 0xf3,
	0x9a, 0x7f, 0x6a, 0x5e, 0xe8, 0xa1, 0x65, 0xd1, 0x45, 0x69, 0x59, 0x5b, 0x85, 0x65, 0xc5, 0x03,
	0xcb, 0xaa, 0x35, 0x37, 0x07, 0xf5, 0xc3, 0x3e, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xcc, 0xe2,
	0x93, 0x45, 0xed, 0x02, 0x00, 0x00,
}
