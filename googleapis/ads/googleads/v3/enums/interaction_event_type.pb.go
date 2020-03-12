// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/interaction_event_type.proto

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
	return fileDescriptor_d3dbb1c76f569f9a, []int{0, 0}
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
	return fileDescriptor_d3dbb1c76f569f9a, []int{0}
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
	proto.RegisterEnum("google.ads.googleads.v3.enums.InteractionEventTypeEnum_InteractionEventType", InteractionEventTypeEnum_InteractionEventType_name, InteractionEventTypeEnum_InteractionEventType_value)
	proto.RegisterType((*InteractionEventTypeEnum)(nil), "google.ads.googleads.v3.enums.InteractionEventTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/interaction_event_type.proto", fileDescriptor_d3dbb1c76f569f9a)
}

var fileDescriptor_d3dbb1c76f569f9a = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xdd, 0x4e, 0xc2, 0x30,
	0x18, 0x75, 0x03, 0xfc, 0x29, 0x89, 0x2e, 0x8b, 0x17, 0x6a, 0xe4, 0x02, 0x1e, 0xa0, 0xbb, 0xd8,
	0x5d, 0xbd, 0x1a, 0x50, 0x97, 0x05, 0x2d, 0x24, 0xc2, 0x48, 0xcc, 0x12, 0x32, 0x59, 0xd3, 0x2c,
	0x81, 0x76, 0xa1, 0x85, 0x84, 0x07, 0xf0, 0x45, 0xbc, 0xf4, 0x51, 0x7c, 0x14, 0x2f, 0x7c, 0x06,
	0xd3, 0x56, 0xf0, 0x06, 0xbd, 0x69, 0x4e, 0xbf, 0xf3, 0x9d, 0x93, 0xef, 0x1c, 0x80, 0x98, 0x10,
	0x6c, 0x41, 0x83, 0xbc, 0x90, 0x81, 0x85, 0x1a, 0x6d, 0xc2, 0x80, 0xf2, 0xf5, 0x52, 0x06, 0x25,
	0x57, 0x74, 0x95, 0xcf, 0x55, 0x29, 0xf8, 0x8c, 0x6e, 0x28, 0x57, 0x33, 0xb5, 0xad, 0x28, 0xac,
	0x56, 0x42, 0x09, 0xbf, 0x65, 0x05, 0x30, 0x2f, 0x24, 0xdc, 0x6b, 0xe1, 0x26, 0x84, 0x46, 0x7b,
	0x73, 0xbb, 0xb3, 0xae, 0xca, 0x20, 0xe7, 0x5c, 0xa8, 0x5c, 0xdb, 0x48, 0x2b, 0xee, 0xbc, 0x3a,
	0xe0, 0x2a, 0xf9, 0x75, 0xc7, 0xda, 0x7c, 0xbc, 0xad, 0x28, 0xe6, 0xeb, 0x65, 0xa7, 0x04, 0x97,
	0x87, 0x38, 0xff, 0x02, 0x34, 0x27, 0xe4, 0x69, 0x84, 0x7b, 0xc9, 0x7d, 0x82, 0xfb, 0xde, 0x91,
	0xdf, 0x04, 0x27, 0x13, 0x32, 0x20, 0xc3, 0x29, 0xf1, 0x1c, 0xff, 0x0c, 0x34, 0x7a, 0x0f, 0x49,
	0x6f, 0xe0, 0xb9, 0xfe, 0x39, 0x00, 0x98, 0xc4, 0x51, 0x8c, 0x1f, 0x31, 0x19, 0x7b, 0x35, 0xfd,
	0x4f, 0x93, 0x3e, 0x1e, 0xce, 0xd2, 0x04, 0x4f, 0xbd, 0xba, 0x7f, 0x0a, 0xea, 0x64, 0x48, 0xb0,
	0xd7, 0xe8, 0x7e, 0x39, 0xa0, 0x3d, 0x17, 0x4b, 0xf8, 0x6f, 0x96, 0xee, 0xf5, 0xa1, 0x73, 0x46,
	0x3a, 0xc8, 0xc8, 0x79, 0xee, 0xfe, 0x68, 0x99, 0x58, 0xe4, 0x9c, 0x41, 0xb1, 0x62, 0x01, 0xa3,
	0xdc, 0xc4, 0xdc, 0x75, 0x5a, 0x95, 0xf2, 0x8f, 0x8a, 0xef, 0xcc, 0xfb, 0xe6, 0xd6, 0xe2, 0x28,
	0x7a, 0x77, 0x5b, 0xb1, 0xb5, 0x8a, 0x0a, 0x09, 0x2d, 0xd4, 0x28, 0x0d, 0xa1, 0xae, 0x45, 0x7e,
	0xec, 0xf8, 0x2c, 0x2a, 0x64, 0xb6, 0xe7, 0xb3, 0x34, 0xcc, 0x0c, 0xff, 0xe9, 0xb6, 0xed, 0x10,
	0xa1, 0xa8, 0x90, 0x08, 0xed, 0x37, 0x10, 0x4a, 0x43, 0x84, 0xcc, 0xce, 0xcb, 0xb1, 0x39, 0x2c,
	0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x28, 0x56, 0xeb, 0x66, 0xfa, 0x01, 0x00, 0x00,
}
