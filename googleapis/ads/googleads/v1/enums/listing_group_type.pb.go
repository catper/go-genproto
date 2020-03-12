// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/listing_group_type.proto

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

// The type of the listing group.
type ListingGroupTypeEnum_ListingGroupType int32

const (
	// Not specified.
	ListingGroupTypeEnum_UNSPECIFIED ListingGroupTypeEnum_ListingGroupType = 0
	// Used for return value only. Represents value unknown in this version.
	ListingGroupTypeEnum_UNKNOWN ListingGroupTypeEnum_ListingGroupType = 1
	// Subdivision of products along some listing dimension. These nodes
	// are not used by serving to target listing entries, but is purely
	// to define the structure of the tree.
	ListingGroupTypeEnum_SUBDIVISION ListingGroupTypeEnum_ListingGroupType = 2
	// Listing group unit that defines a bid.
	ListingGroupTypeEnum_UNIT ListingGroupTypeEnum_ListingGroupType = 3
)

var ListingGroupTypeEnum_ListingGroupType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "SUBDIVISION",
	3: "UNIT",
}

var ListingGroupTypeEnum_ListingGroupType_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"SUBDIVISION": 2,
	"UNIT":        3,
}

func (x ListingGroupTypeEnum_ListingGroupType) String() string {
	return proto.EnumName(ListingGroupTypeEnum_ListingGroupType_name, int32(x))
}

func (ListingGroupTypeEnum_ListingGroupType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fbd98b8004bba2ed, []int{0, 0}
}

// Container for enum describing the type of the listing group.
type ListingGroupTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListingGroupTypeEnum) Reset()         { *m = ListingGroupTypeEnum{} }
func (m *ListingGroupTypeEnum) String() string { return proto.CompactTextString(m) }
func (*ListingGroupTypeEnum) ProtoMessage()    {}
func (*ListingGroupTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbd98b8004bba2ed, []int{0}
}

func (m *ListingGroupTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListingGroupTypeEnum.Unmarshal(m, b)
}
func (m *ListingGroupTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListingGroupTypeEnum.Marshal(b, m, deterministic)
}
func (m *ListingGroupTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListingGroupTypeEnum.Merge(m, src)
}
func (m *ListingGroupTypeEnum) XXX_Size() int {
	return xxx_messageInfo_ListingGroupTypeEnum.Size(m)
}
func (m *ListingGroupTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ListingGroupTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ListingGroupTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.ListingGroupTypeEnum_ListingGroupType", ListingGroupTypeEnum_ListingGroupType_name, ListingGroupTypeEnum_ListingGroupType_value)
	proto.RegisterType((*ListingGroupTypeEnum)(nil), "google.ads.googleads.v1.enums.ListingGroupTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/listing_group_type.proto", fileDescriptor_fbd98b8004bba2ed)
}

var fileDescriptor_fbd98b8004bba2ed = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xcf, 0x4a, 0xc3, 0x30,
	0x18, 0x77, 0x9d, 0xa8, 0x64, 0x07, 0x4b, 0xd1, 0x8b, 0xb8, 0xc3, 0xf6, 0x00, 0x09, 0x45, 0xf0,
	0x10, 0x4f, 0xad, 0x9b, 0x23, 0x4c, 0xb2, 0xc1, 0xd6, 0x0a, 0x52, 0x18, 0x75, 0x2d, 0xa1, 0xb0,
	0x26, 0xa1, 0x69, 0x07, 0x7b, 0x1d, 0x8f, 0x3e, 0x8a, 0x8f, 0xb2, 0xa7, 0x90, 0x24, 0xb6, 0x87,
	0x81, 0x5e, 0xc2, 0x8f, 0xef, 0xf7, 0x27, 0xbf, 0xef, 0x03, 0x8f, 0x4c, 0x08, 0xb6, 0xcb, 0x51,
	0x9a, 0x29, 0x64, 0xa1, 0x46, 0x7b, 0x1f, 0xe5, 0xbc, 0x29, 0x15, 0xda, 0x15, 0xaa, 0x2e, 0x38,
	0xdb, 0xb0, 0x4a, 0x34, 0x72, 0x53, 0x1f, 0x64, 0x0e, 0x65, 0x25, 0x6a, 0xe1, 0x0d, 0xad, 0x18,
	0xa6, 0x99, 0x82, 0x9d, 0x0f, 0xee, 0x7d, 0x68, 0x7c, 0x77, 0xf7, 0x6d, 0xac, 0x2c, 0x50, 0xca,
	0xb9, 0xa8, 0xd3, 0xba, 0x10, 0x5c, 0x59, 0xf3, 0x78, 0x0b, 0x6e, 0x5e, 0x6d, 0xf0, 0x4c, 0xe7,
	0xae, 0x0f, 0x32, 0x9f, 0xf2, 0xa6, 0x1c, 0xcf, 0x81, 0x7b, 0x3a, 0xf7, 0xae, 0xc1, 0x20, 0xa2,
	0xab, 0xe5, 0xf4, 0x99, 0xbc, 0x90, 0xe9, 0xc4, 0x3d, 0xf3, 0x06, 0xe0, 0x32, 0xa2, 0x73, 0xba,
	0x78, 0xa3, 0x6e, 0x4f, 0xb3, 0xab, 0x28, 0x9c, 0x90, 0x98, 0xac, 0xc8, 0x82, 0xba, 0x8e, 0x77,
	0x05, 0xce, 0x23, 0x4a, 0xd6, 0x6e, 0x3f, 0x3c, 0xf6, 0xc0, 0x68, 0x2b, 0x4a, 0xf8, 0x6f, 0xd1,
	0xf0, 0xf6, 0xf4, 0xc3, 0xa5, 0x6e, 0xb8, 0xec, 0xbd, 0x87, 0xbf, 0x3e, 0x26, 0x76, 0x29, 0x67,
	0x50, 0x54, 0x0c, 0xb1, 0x9c, 0x9b, 0xfe, 0xed, 0xa1, 0x64, 0xa1, 0xfe, 0xb8, 0xdb, 0x93, 0x79,
	0x3f, 0x9d, 0xfe, 0x2c, 0x08, 0xbe, 0x9c, 0xe1, 0xcc, 0x46, 0x05, 0x99, 0x82, 0x16, 0x6a, 0x14,
	0xfb, 0x50, 0x2f, 0xad, 0xbe, 0x5b, 0x3e, 0x09, 0x32, 0x95, 0x74, 0x7c, 0x12, 0xfb, 0x89, 0xe1,
	0x8f, 0xce, 0xc8, 0x0e, 0x31, 0x0e, 0x32, 0x85, 0x71, 0xa7, 0xc0, 0x38, 0xf6, 0x31, 0x36, 0x9a,
	0x8f, 0x0b, 0x53, 0xec, 0xe1, 0x27, 0x00, 0x00, 0xff, 0xff, 0x5c, 0xbf, 0x3f, 0x92, 0xcf, 0x01,
	0x00, 0x00,
}
