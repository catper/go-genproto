// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/interaction_event_type.proto

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

// Enum describing possible types of payable and free interactions.
type InteractionEventTypeEnum_InteractionEventType int32

const (
	// Not specified.
	InteractionEventTypeEnum_UNSPECIFIED InteractionEventTypeEnum_InteractionEventType = 0
	// Used for return value only. Represents value unknown in this version.
	InteractionEventTypeEnum_UNKNOWN InteractionEventTypeEnum_InteractionEventType = 1
	// Click to site. In most cases, this interaction navigates to an external
	// location, usually the advertiser's landing page. This is also the default
	// InteractionEventType for click events.
	InteractionEventTypeEnum_CLICK InteractionEventTypeEnum_InteractionEventType = 2
	// The user's expressed intent to engage with the ad in-place.
	InteractionEventTypeEnum_ENGAGEMENT InteractionEventTypeEnum_InteractionEventType = 3
	// User viewed a video ad.
	InteractionEventTypeEnum_VIDEO_VIEW InteractionEventTypeEnum_InteractionEventType = 4
	// The default InteractionEventType for ad conversion events.
	// This is used when an ad conversion row does NOT indicate
	// that the free interactions (i.e., the ad conversions)
	// should be 'promoted' and reported as part of the core metrics.
	// These are simply other (ad) conversions.
	InteractionEventTypeEnum_NONE InteractionEventTypeEnum_InteractionEventType = 5
)

var InteractionEventTypeEnum_InteractionEventType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "CLICK",
	3: "ENGAGEMENT",
	4: "VIDEO_VIEW",
	5: "NONE",
}

var InteractionEventTypeEnum_InteractionEventType_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"CLICK":       2,
	"ENGAGEMENT":  3,
	"VIDEO_VIEW":  4,
	"NONE":        5,
}

func (x InteractionEventTypeEnum_InteractionEventType) String() string {
	return proto.EnumName(InteractionEventTypeEnum_InteractionEventType_name, int32(x))
}

func (InteractionEventTypeEnum_InteractionEventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_68735a9e92f87cdf, []int{0, 0}
}

// Container for enum describing types of payable and free interactions.
type InteractionEventTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InteractionEventTypeEnum) Reset()         { *m = InteractionEventTypeEnum{} }
func (m *InteractionEventTypeEnum) String() string { return proto.CompactTextString(m) }
func (*InteractionEventTypeEnum) ProtoMessage()    {}
func (*InteractionEventTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_68735a9e92f87cdf, []int{0}
}

func (m *InteractionEventTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InteractionEventTypeEnum.Unmarshal(m, b)
}
func (m *InteractionEventTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InteractionEventTypeEnum.Marshal(b, m, deterministic)
}
func (m *InteractionEventTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InteractionEventTypeEnum.Merge(m, src)
}
func (m *InteractionEventTypeEnum) XXX_Size() int {
	return xxx_messageInfo_InteractionEventTypeEnum.Size(m)
}
func (m *InteractionEventTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_InteractionEventTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_InteractionEventTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.InteractionEventTypeEnum_InteractionEventType", InteractionEventTypeEnum_InteractionEventType_name, InteractionEventTypeEnum_InteractionEventType_value)
	proto.RegisterType((*InteractionEventTypeEnum)(nil), "google.ads.googleads.v2.enums.InteractionEventTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/interaction_event_type.proto", fileDescriptor_68735a9e92f87cdf)
}

var fileDescriptor_68735a9e92f87cdf = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xd1, 0x4e, 0xf2, 0x30,
	0x18, 0xfd, 0x37, 0xe0, 0x57, 0x4b, 0xa2, 0xcb, 0xe2, 0x85, 0x1a, 0xb9, 0x80, 0x07, 0xe8, 0x92,
	0x79, 0x57, 0xaf, 0x06, 0xd4, 0x65, 0x41, 0x0b, 0x89, 0x30, 0x12, 0xb3, 0x84, 0x4c, 0xd6, 0x34,
	0x4b, 0xa0, 0x5d, 0x68, 0x21, 0xe1, 0x01, 0x7c, 0x11, 0x2f, 0x7d, 0x14, 0x1f, 0xc5, 0x0b, 0x9f,
	0xc1, 0xb4, 0x15, 0xbc, 0x41, 0x6f, 0x9a, 0xd3, 0xef, 0x7c, 0xe7, 0xe4, 0x3b, 0x07, 0x20, 0x26,
	0x04, 0x5b, 0xd0, 0x20, 0x2f, 0x64, 0x60, 0xa1, 0x46, 0x9b, 0x30, 0xa0, 0x7c, 0xbd, 0x94, 0x41,
	0xc9, 0x15, 0x5d, 0xe5, 0x73, 0x55, 0x0a, 0x3e, 0xa3, 0x1b, 0xca, 0xd5, 0x4c, 0x6d, 0x2b, 0x0a,
	0xab, 0x95, 0x50, 0xc2, 0x6f, 0x59, 0x01, 0xcc, 0x0b, 0x09, 0xf7, 0x5a, 0xb8, 0x09, 0xa1, 0xd1,
	0x5e, 0x5d, 0xef, 0xac, 0xab, 0x32, 0xc8, 0x39, 0x17, 0x2a, 0xd7, 0x36, 0xd2, 0x8a, 0x3b, 0x2f,
	0x0e, 0xb8, 0x48, 0x7e, 0xdc, 0xb1, 0x36, 0x1f, 0x6f, 0x2b, 0x8a, 0xf9, 0x7a, 0xd9, 0x29, 0xc1,
	0xf9, 0x21, 0xce, 0x3f, 0x03, 0xcd, 0x09, 0x79, 0x1c, 0xe1, 0x5e, 0x72, 0x97, 0xe0, 0xbe, 0xf7,
	0xcf, 0x6f, 0x82, 0xa3, 0x09, 0x19, 0x90, 0xe1, 0x94, 0x78, 0x8e, 0x7f, 0x02, 0x1a, 0xbd, 0xfb,
	0xa4, 0x37, 0xf0, 0x5c, 0xff, 0x14, 0x00, 0x4c, 0xe2, 0x28, 0xc6, 0x0f, 0x98, 0x8c, 0xbd, 0x9a,
	0xfe, 0xa7, 0x49, 0x1f, 0x0f, 0x67, 0x69, 0x82, 0xa7, 0x5e, 0xdd, 0x3f, 0x06, 0x75, 0x32, 0x24,
	0xd8, 0x6b, 0x74, 0x3f, 0x1d, 0xd0, 0x9e, 0x8b, 0x25, 0xfc, 0x33, 0x4b, 0xf7, 0xf2, 0xd0, 0x39,
	0x23, 0x1d, 0x64, 0xe4, 0x3c, 0x75, 0xbf, 0xb5, 0x4c, 0x2c, 0x72, 0xce, 0xa0, 0x58, 0xb1, 0x80,
	0x51, 0x6e, 0x62, 0xee, 0x3a, 0xad, 0x4a, 0xf9, 0x4b, 0xc5, 0xb7, 0xe6, 0x7d, 0x75, 0x6b, 0x71,
	0x14, 0xbd, 0xb9, 0xad, 0xd8, 0x5a, 0x45, 0x85, 0x84, 0x16, 0x6a, 0x94, 0x86, 0x50, 0xd7, 0x22,
	0xdf, 0x77, 0x7c, 0x16, 0x15, 0x32, 0xdb, 0xf3, 0x59, 0x1a, 0x66, 0x86, 0xff, 0x70, 0xdb, 0x76,
	0x88, 0x50, 0x54, 0x48, 0x84, 0xf6, 0x1b, 0x08, 0xa5, 0x21, 0x42, 0x66, 0xe7, 0xf9, 0xbf, 0x39,
	0xec, 0xe6, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xf2, 0xc0, 0x2c, 0x70, 0xfa, 0x01, 0x00, 0x00,
}
