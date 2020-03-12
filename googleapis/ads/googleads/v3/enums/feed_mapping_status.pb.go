// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/feed_mapping_status.proto

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

// Possible statuses of a feed mapping.
type FeedMappingStatusEnum_FeedMappingStatus int32

const (
	// Not specified.
	FeedMappingStatusEnum_UNSPECIFIED FeedMappingStatusEnum_FeedMappingStatus = 0
	// Used for return value only. Represents value unknown in this version.
	FeedMappingStatusEnum_UNKNOWN FeedMappingStatusEnum_FeedMappingStatus = 1
	// Feed mapping is enabled.
	FeedMappingStatusEnum_ENABLED FeedMappingStatusEnum_FeedMappingStatus = 2
	// Feed mapping has been removed.
	FeedMappingStatusEnum_REMOVED FeedMappingStatusEnum_FeedMappingStatus = 3
)

var FeedMappingStatusEnum_FeedMappingStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ENABLED",
	3: "REMOVED",
}

var FeedMappingStatusEnum_FeedMappingStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ENABLED":     2,
	"REMOVED":     3,
}

func (x FeedMappingStatusEnum_FeedMappingStatus) String() string {
	return proto.EnumName(FeedMappingStatusEnum_FeedMappingStatus_name, int32(x))
}

func (FeedMappingStatusEnum_FeedMappingStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ce0d893a1df7704e, []int{0, 0}
}

// Container for enum describing possible statuses of a feed mapping.
type FeedMappingStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeedMappingStatusEnum) Reset()         { *m = FeedMappingStatusEnum{} }
func (m *FeedMappingStatusEnum) String() string { return proto.CompactTextString(m) }
func (*FeedMappingStatusEnum) ProtoMessage()    {}
func (*FeedMappingStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce0d893a1df7704e, []int{0}
}

func (m *FeedMappingStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedMappingStatusEnum.Unmarshal(m, b)
}
func (m *FeedMappingStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedMappingStatusEnum.Marshal(b, m, deterministic)
}
func (m *FeedMappingStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedMappingStatusEnum.Merge(m, src)
}
func (m *FeedMappingStatusEnum) XXX_Size() int {
	return xxx_messageInfo_FeedMappingStatusEnum.Size(m)
}
func (m *FeedMappingStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedMappingStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_FeedMappingStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.FeedMappingStatusEnum_FeedMappingStatus", FeedMappingStatusEnum_FeedMappingStatus_name, FeedMappingStatusEnum_FeedMappingStatus_value)
	proto.RegisterType((*FeedMappingStatusEnum)(nil), "google.ads.googleads.v3.enums.FeedMappingStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/feed_mapping_status.proto", fileDescriptor_ce0d893a1df7704e)
}

var fileDescriptor_ce0d893a1df7704e = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xcd, 0x4a, 0x33, 0x31,
	0x14, 0xfd, 0x3a, 0x85, 0x4f, 0x48, 0x17, 0x8e, 0x05, 0x5d, 0x88, 0x5d, 0xb4, 0x0f, 0x90, 0x2c,
	0xb2, 0x10, 0xe2, 0x2a, 0x63, 0xd3, 0x52, 0x6a, 0xa7, 0xc5, 0xd2, 0x11, 0x64, 0xa0, 0x44, 0x13,
	0xc3, 0x40, 0x27, 0x19, 0x9a, 0x69, 0x1f, 0xc8, 0xa5, 0x8f, 0xe2, 0xa3, 0xe8, 0x4b, 0x48, 0x92,
	0x76, 0x36, 0x45, 0x37, 0xc3, 0xb9, 0x73, 0x7e, 0x72, 0xee, 0x05, 0xb7, 0xca, 0x18, 0xb5, 0x91,
	0x88, 0x0b, 0x8b, 0x02, 0x74, 0x68, 0x8f, 0x91, 0xd4, 0xbb, 0xd2, 0xa2, 0x37, 0x29, 0xc5, 0xba,
	0xe4, 0x55, 0x55, 0x68, 0xb5, 0xb6, 0x35, 0xaf, 0x77, 0x16, 0x56, 0x5b, 0x53, 0x9b, 0x6e, 0x2f,
	0xa8, 0x21, 0x17, 0x16, 0x36, 0x46, 0xb8, 0xc7, 0xd0, 0x1b, 0xaf, 0x6f, 0x8e, 0xb9, 0x55, 0x81,
	0xb8, 0xd6, 0xa6, 0xe6, 0x75, 0x61, 0xf4, 0xc1, 0x3c, 0x10, 0xe0, 0x72, 0x24, 0xa5, 0x98, 0x85,
	0xe0, 0xa5, 0xcf, 0x65, 0x7a, 0x57, 0x0e, 0xa6, 0xe0, 0xe2, 0x84, 0xe8, 0x9e, 0x83, 0xce, 0x2a,
	0x5d, 0x2e, 0xd8, 0xfd, 0x64, 0x34, 0x61, 0xc3, 0xf8, 0x5f, 0xb7, 0x03, 0xce, 0x56, 0xe9, 0x34,
	0x9d, 0x3f, 0xa5, 0x71, 0xcb, 0x0d, 0x2c, 0xa5, 0xc9, 0x03, 0x1b, 0xc6, 0x91, 0x1b, 0x1e, 0xd9,
	0x6c, 0x9e, 0xb1, 0x61, 0xdc, 0x4e, 0xbe, 0x5b, 0xa0, 0xff, 0x6a, 0x4a, 0xf8, 0x67, 0xd3, 0xe4,
	0xea, 0xe4, 0xc1, 0x85, 0xeb, 0xb8, 0x68, 0x3d, 0x27, 0x07, 0xa3, 0x32, 0x1b, 0xae, 0x15, 0x34,
	0x5b, 0x85, 0x94, 0xd4, 0x7e, 0x83, 0xe3, 0xad, 0xaa, 0xc2, 0xfe, 0x72, 0xba, 0x3b, 0xff, 0x7d,
	0x8f, 0xda, 0x63, 0x4a, 0x3f, 0xa2, 0xde, 0x38, 0x44, 0x51, 0x61, 0x61, 0x80, 0x0e, 0x65, 0x18,
	0xba, 0xad, 0xed, 0xe7, 0x91, 0xcf, 0xa9, 0xb0, 0x79, 0xc3, 0xe7, 0x19, 0xce, 0x3d, 0xff, 0x15,
	0xf5, 0xc3, 0x4f, 0x42, 0xa8, 0xb0, 0x84, 0x34, 0x0a, 0x42, 0x32, 0x4c, 0x88, 0xd7, 0xbc, 0xfc,
	0xf7, 0xc5, 0xf0, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x77, 0xe5, 0x4c, 0x93, 0xd2, 0x01, 0x00,
	0x00,
}
