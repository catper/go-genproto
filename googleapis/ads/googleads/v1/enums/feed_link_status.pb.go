// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/feed_link_status.proto

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

// Possible statuses of a feed link.
type FeedLinkStatusEnum_FeedLinkStatus int32

const (
	// Not specified.
	FeedLinkStatusEnum_UNSPECIFIED FeedLinkStatusEnum_FeedLinkStatus = 0
	// Used for return value only. Represents value unknown in this version.
	FeedLinkStatusEnum_UNKNOWN FeedLinkStatusEnum_FeedLinkStatus = 1
	// Feed link is enabled.
	FeedLinkStatusEnum_ENABLED FeedLinkStatusEnum_FeedLinkStatus = 2
	// Feed link has been removed.
	FeedLinkStatusEnum_REMOVED FeedLinkStatusEnum_FeedLinkStatus = 3
)

var FeedLinkStatusEnum_FeedLinkStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ENABLED",
	3: "REMOVED",
}

var FeedLinkStatusEnum_FeedLinkStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ENABLED":     2,
	"REMOVED":     3,
}

func (x FeedLinkStatusEnum_FeedLinkStatus) String() string {
	return proto.EnumName(FeedLinkStatusEnum_FeedLinkStatus_name, int32(x))
}

func (FeedLinkStatusEnum_FeedLinkStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_23c30109a5e5efe9, []int{0, 0}
}

// Container for an enum describing possible statuses of a feed link.
type FeedLinkStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeedLinkStatusEnum) Reset()         { *m = FeedLinkStatusEnum{} }
func (m *FeedLinkStatusEnum) String() string { return proto.CompactTextString(m) }
func (*FeedLinkStatusEnum) ProtoMessage()    {}
func (*FeedLinkStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_23c30109a5e5efe9, []int{0}
}

func (m *FeedLinkStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedLinkStatusEnum.Unmarshal(m, b)
}
func (m *FeedLinkStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedLinkStatusEnum.Marshal(b, m, deterministic)
}
func (m *FeedLinkStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedLinkStatusEnum.Merge(m, src)
}
func (m *FeedLinkStatusEnum) XXX_Size() int {
	return xxx_messageInfo_FeedLinkStatusEnum.Size(m)
}
func (m *FeedLinkStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedLinkStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_FeedLinkStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.FeedLinkStatusEnum_FeedLinkStatus", FeedLinkStatusEnum_FeedLinkStatus_name, FeedLinkStatusEnum_FeedLinkStatus_value)
	proto.RegisterType((*FeedLinkStatusEnum)(nil), "google.ads.googleads.v1.enums.FeedLinkStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/feed_link_status.proto", fileDescriptor_23c30109a5e5efe9)
}

var fileDescriptor_23c30109a5e5efe9 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xcf, 0x4a, 0xc3, 0x30,
	0x18, 0x77, 0x1d, 0x28, 0x64, 0xa0, 0xa5, 0xde, 0xc4, 0x1d, 0xb6, 0x07, 0x48, 0x28, 0x7a, 0x8a,
	0xa7, 0xd4, 0x65, 0x73, 0x38, 0xbb, 0xe1, 0x58, 0x05, 0x29, 0x8e, 0x68, 0x62, 0x28, 0xdb, 0x92,
	0xb2, 0xb4, 0x7b, 0x20, 0x8f, 0x3e, 0x8a, 0x4f, 0x22, 0x3e, 0x85, 0x24, 0x5d, 0x0b, 0x3b, 0xe8,
	0xa5, 0xfc, 0xbe, 0xef, 0xf7, 0xa7, 0xbf, 0x7c, 0xe0, 0x5a, 0x6a, 0x2d, 0xd7, 0x02, 0x31, 0x6e,
	0x50, 0x05, 0x2d, 0xda, 0x85, 0x48, 0xa8, 0x72, 0x63, 0xd0, 0xbb, 0x10, 0x7c, 0xb9, 0xce, 0xd4,
	0x6a, 0x69, 0x0a, 0x56, 0x94, 0x06, 0xe6, 0x5b, 0x5d, 0xe8, 0xa0, 0x5b, 0x49, 0x21, 0xe3, 0x06,
	0x36, 0x2e, 0xb8, 0x0b, 0xa1, 0x73, 0x5d, 0x5c, 0xd6, 0xa1, 0x79, 0x86, 0x98, 0x52, 0xba, 0x60,
	0x45, 0xa6, 0xd5, 0xde, 0xdc, 0x7f, 0x01, 0xc1, 0x50, 0x08, 0x3e, 0xc9, 0xd4, 0x6a, 0xee, 0x42,
	0xa9, 0x2a, 0x37, 0xfd, 0x3b, 0x70, 0x7a, 0xb8, 0x0d, 0xce, 0x40, 0x67, 0x11, 0xcf, 0x67, 0xf4,
	0x76, 0x3c, 0x1c, 0xd3, 0x81, 0x7f, 0x14, 0x74, 0xc0, 0xc9, 0x22, 0xbe, 0x8f, 0xa7, 0x4f, 0xb1,
	0xdf, 0xb2, 0x03, 0x8d, 0x49, 0x34, 0xa1, 0x03, 0xdf, 0xb3, 0xc3, 0x23, 0x7d, 0x98, 0x26, 0x74,
	0xe0, 0xb7, 0xa3, 0xef, 0x16, 0xe8, 0xbd, 0xe9, 0x0d, 0xfc, 0xb7, 0x63, 0x74, 0x7e, 0xf8, 0xb7,
	0x99, 0xad, 0x36, 0x6b, 0x3d, 0x47, 0x7b, 0x97, 0xd4, 0x6b, 0xa6, 0x24, 0xd4, 0x5b, 0x89, 0xa4,
	0x50, 0xae, 0x78, 0x7d, 0x9f, 0x3c, 0x33, 0x7f, 0x9c, 0xeb, 0xc6, 0x7d, 0x3f, 0xbc, 0xf6, 0x88,
	0x90, 0x4f, 0xaf, 0x3b, 0xaa, 0xa2, 0x08, 0x37, 0xb0, 0x82, 0x16, 0x25, 0x21, 0xb4, 0xef, 0x35,
	0x5f, 0x35, 0x9f, 0x12, 0x6e, 0xd2, 0x86, 0x4f, 0x93, 0x30, 0x75, 0xfc, 0x8f, 0xd7, 0xab, 0x96,
	0x18, 0x13, 0x6e, 0x30, 0x6e, 0x14, 0x18, 0x27, 0x21, 0xc6, 0x4e, 0xf3, 0x7a, 0xec, 0x8a, 0x5d,
	0xfd, 0x06, 0x00, 0x00, 0xff, 0xff, 0xe1, 0x0e, 0x39, 0x40, 0xc6, 0x01, 0x00, 0x00,
}
