// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/content_label_type.proto

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
	return fileDescriptor_b519dcc6be978913, []int{0, 0}
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
	return fileDescriptor_b519dcc6be978913, []int{0}
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
	proto.RegisterEnum("google.ads.googleads.v3.enums.ContentLabelTypeEnum_ContentLabelType", ContentLabelTypeEnum_ContentLabelType_name, ContentLabelTypeEnum_ContentLabelType_value)
	proto.RegisterType((*ContentLabelTypeEnum)(nil), "google.ads.googleads.v3.enums.ContentLabelTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/content_label_type.proto", fileDescriptor_b519dcc6be978913)
}

var fileDescriptor_b519dcc6be978913 = []byte{
	// 479 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x4d, 0x6e, 0xd3, 0x40,
	0x14, 0x26, 0x2e, 0xb4, 0xcd, 0xa4, 0x69, 0x27, 0x43, 0x0b, 0x08, 0xd1, 0x45, 0x7b, 0x00, 0x7b,
	0x61, 0x89, 0x85, 0x59, 0x4d, 0x32, 0x2f, 0x66, 0xa8, 0x63, 0x5b, 0xf1, 0x4f, 0x09, 0x8a, 0x34,
	0x72, 0x1b, 0xcb, 0x8a, 0x94, 0x78, 0xac, 0xda, 0xad, 0xd4, 0xeb, 0xb0, 0xe4, 0x28, 0xdc, 0x81,
	0x0b, 0xf4, 0x00, 0xac, 0xd1, 0x78, 0x48, 0x16, 0x51, 0x61, 0x63, 0x7d, 0x7a, 0xdf, 0x8f, 0xe7,
	0xe9, 0x7b, 0xe8, 0x63, 0x21, 0x65, 0xb1, 0xca, 0xad, 0x6c, 0x51, 0x5b, 0x1a, 0x2a, 0xf4, 0x60,
	0x5b, 0x79, 0x79, 0xbf, 0xae, 0xad, 0x5b, 0x59, 0x36, 0x79, 0xd9, 0x88, 0x55, 0x76, 0x93, 0xaf,
	0x44, 0xf3, 0x58, 0xe5, 0x66, 0x75, 0x27, 0x1b, 0x49, 0xce, 0xb5, 0xd8, 0xcc, 0x16, 0xb5, 0xb9,
	0xf5, 0x99, 0x0f, 0xb6, 0xd9, 0xfa, 0xde, 0x7f, 0xd8, 0xc4, 0x56, 0x4b, 0x2b, 0x2b, 0x4b, 0xd9,
	0x64, 0xcd, 0x52, 0x96, 0xb5, 0x36, 0x5f, 0xfe, 0x36, 0xd0, 0xe9, 0x48, 0x27, 0x7b, 0x2a, 0x38,
	0x7e, 0xac, 0x72, 0x28, 0xef, 0xd7, 0x97, 0xbf, 0x0c, 0x84, 0x77, 0x09, 0x72, 0x82, 0x7a, 0x89,
	0x1f, 0x85, 0x30, 0xe2, 0x63, 0x0e, 0x0c, 0xbf, 0x20, 0x3d, 0x74, 0x90, 0xf8, 0x57, 0x7e, 0x70,
	0xed, 0xe3, 0x0e, 0x79, 0x8b, 0x5e, 0x47, 0xf0, 0x35, 0xa1, 0x9e, 0x37, 0x13, 0x51, 0xe2, 0xba,
	0x10, 0xc5, 0x3c, 0x05, 0x6c, 0x10, 0x82, 0x8e, 0x87, 0xe0, 0x05, 0xd7, 0x22, 0xfe, 0x0c, 0x62,
	0x1c, 0x78, 0x0c, 0xef, 0x91, 0x01, 0xea, 0x87, 0x74, 0x7a, 0x05, 0x4c, 0xb0, 0x60, 0x42, 0xb9,
	0x8f, 0x5f, 0x92, 0x23, 0x74, 0xf8, 0x25, 0x49, 0xc1, 0xe7, 0x1e, 0xe0, 0x7d, 0xd2, 0x47, 0xdd,
	0x70, 0x1a, 0x8c, 0xa9, 0xcf, 0xe3, 0x19, 0x3e, 0x50, 0x7f, 0x8a, 0xa7, 0xd4, 0x05, 0x36, 0xc3,
	0x87, 0xa4, 0x8b, 0x5e, 0xa5, 0x9c, 0x41, 0x80, 0xbb, 0xe4, 0x0c, 0x0d, 0x5a, 0x28, 0xa6, 0x34,
	0xe6, 0xbe, 0x2b, 0x58, 0x2a, 0x5c, 0x8c, 0xc8, 0x1b, 0x44, 0x76, 0xc7, 0xa1, 0x8b, 0x7b, 0xcf,
	0xc9, 0x63, 0x7c, 0xf4, 0x9c, 0x7c, 0x42, 0x71, 0x5f, 0xad, 0xa4, 0xe7, 0x7e, 0x10, 0x8b, 0x19,
	0xc4, 0x8a, 0x07, 0x86, 0x8f, 0xd5, 0x4a, 0x30, 0x19, 0x02, 0x63, 0xc0, 0x84, 0x7e, 0xca, 0x09,
	0x79, 0x87, 0x4e, 0x3d, 0x9e, 0x82, 0x88, 0xe2, 0x29, 0xd0, 0x89, 0x8a, 0xd1, 0x0c, 0x56, 0xcb,
	0x46, 0xc1, 0x88, 0x53, 0x4f, 0xf0, 0x28, 0x4a, 0x20, 0xc2, 0x83, 0xe1, 0x53, 0x07, 0x5d, 0xdc,
	0xca, 0xb5, 0xf9, 0xdf, 0xf2, 0x86, 0x67, 0xbb, 0x15, 0x84, 0xaa, 0xb5, 0xb0, 0xf3, 0x6d, 0xf8,
	0xd7, 0x57, 0xc8, 0x55, 0x56, 0x16, 0xa6, 0xbc, 0x2b, 0xac, 0x22, 0x2f, 0xdb, 0x4e, 0x37, 0xc7,
	0x53, 0x2d, 0xeb, 0x7f, 0xdc, 0xd2, 0xa7, 0xf6, 0xfb, 0xdd, 0xd8, 0x73, 0x29, 0xfd, 0x61, 0x9c,
	0xbb, 0x3a, 0x8a, 0x2e, 0x6a, 0x53, 0x43, 0x85, 0x52, 0xdb, 0x54, 0x77, 0x50, 0xff, 0xdc, 0xf0,
	0x73, 0xba, 0xa8, 0xe7, 0x5b, 0x7e, 0x9e, 0xda, 0xf3, 0x96, 0x7f, 0x32, 0x2e, 0xf4, 0xd0, 0x71,
	0xe8, 0xa2, 0x76, 0x9c, 0xad, 0xc2, 0x71, 0x52, 0xdb, 0x71, 0x5a, 0xcd, 0xcd, 0x7e, 0xfb, 0x30,
	0xfb, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdc, 0x73, 0x79, 0x58, 0xe3, 0x02, 0x00, 0x00,
}
