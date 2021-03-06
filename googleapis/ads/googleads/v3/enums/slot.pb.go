// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/slot.proto

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

// Enumerates possible positions of the Ad.
type SlotEnum_Slot int32

const (
	// Not specified.
	SlotEnum_UNSPECIFIED SlotEnum_Slot = 0
	// The value is unknown in this version.
	SlotEnum_UNKNOWN SlotEnum_Slot = 1
	// Google search: Side.
	SlotEnum_SEARCH_SIDE SlotEnum_Slot = 2
	// Google search: Top.
	SlotEnum_SEARCH_TOP SlotEnum_Slot = 3
	// Google search: Other.
	SlotEnum_SEARCH_OTHER SlotEnum_Slot = 4
	// Google Display Network.
	SlotEnum_CONTENT SlotEnum_Slot = 5
	// Search partners: Top.
	SlotEnum_SEARCH_PARTNER_TOP SlotEnum_Slot = 6
	// Search partners: Other.
	SlotEnum_SEARCH_PARTNER_OTHER SlotEnum_Slot = 7
	// Cross-network.
	SlotEnum_MIXED SlotEnum_Slot = 8
)

var SlotEnum_Slot_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "SEARCH_SIDE",
	3: "SEARCH_TOP",
	4: "SEARCH_OTHER",
	5: "CONTENT",
	6: "SEARCH_PARTNER_TOP",
	7: "SEARCH_PARTNER_OTHER",
	8: "MIXED",
}

var SlotEnum_Slot_value = map[string]int32{
	"UNSPECIFIED":          0,
	"UNKNOWN":              1,
	"SEARCH_SIDE":          2,
	"SEARCH_TOP":           3,
	"SEARCH_OTHER":         4,
	"CONTENT":              5,
	"SEARCH_PARTNER_TOP":   6,
	"SEARCH_PARTNER_OTHER": 7,
	"MIXED":                8,
}

func (x SlotEnum_Slot) String() string {
	return proto.EnumName(SlotEnum_Slot_name, int32(x))
}

func (SlotEnum_Slot) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fef43849b89d590c, []int{0, 0}
}

// Container for enumeration of possible positions of the Ad.
type SlotEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SlotEnum) Reset()         { *m = SlotEnum{} }
func (m *SlotEnum) String() string { return proto.CompactTextString(m) }
func (*SlotEnum) ProtoMessage()    {}
func (*SlotEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_fef43849b89d590c, []int{0}
}

func (m *SlotEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SlotEnum.Unmarshal(m, b)
}
func (m *SlotEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SlotEnum.Marshal(b, m, deterministic)
}
func (m *SlotEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SlotEnum.Merge(m, src)
}
func (m *SlotEnum) XXX_Size() int {
	return xxx_messageInfo_SlotEnum.Size(m)
}
func (m *SlotEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_SlotEnum.DiscardUnknown(m)
}

var xxx_messageInfo_SlotEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.SlotEnum_Slot", SlotEnum_Slot_name, SlotEnum_Slot_value)
	proto.RegisterType((*SlotEnum)(nil), "google.ads.googleads.v3.enums.SlotEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/slot.proto", fileDescriptor_fef43849b89d590c)
}

var fileDescriptor_fef43849b89d590c = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xcf, 0x4a, 0xfb, 0x40,
	0x18, 0xfc, 0x25, 0xfd, 0xbf, 0xfd, 0xa1, 0xcb, 0x22, 0x22, 0x62, 0x0f, 0xed, 0xc9, 0xd3, 0xe6,
	0x90, 0xdb, 0x7a, 0x4a, 0xdb, 0xb5, 0x0d, 0x62, 0x12, 0x92, 0xb4, 0x8a, 0x04, 0x24, 0x9a, 0x12,
	0x0a, 0xe9, 0x6e, 0xe9, 0xa6, 0x7d, 0x1f, 0x3d, 0x7a, 0xf0, 0x41, 0x7c, 0x08, 0x1f, 0xc0, 0xa7,
	0x90, 0xdd, 0x4d, 0x7b, 0x10, 0xf4, 0xb2, 0xcc, 0xf7, 0xcd, 0xcc, 0xc7, 0xce, 0x80, 0xcb, 0x9c,
	0xf3, 0xbc, 0x58, 0x58, 0x69, 0x26, 0x2c, 0x0d, 0x25, 0xda, 0xd9, 0xd6, 0x82, 0x6d, 0x57, 0xc2,
	0x12, 0x05, 0x2f, 0xf1, 0x7a, 0xc3, 0x4b, 0x8e, 0x7a, 0x9a, 0xc6, 0x69, 0x26, 0xf0, 0x41, 0x89,
	0x77, 0x36, 0x56, 0xca, 0xf3, 0x8b, 0xfd, 0xa1, 0xf5, 0xd2, 0x4a, 0x19, 0xe3, 0x65, 0x5a, 0x2e,
	0x39, 0x13, 0xda, 0x3c, 0x78, 0x37, 0x40, 0x3b, 0x2a, 0x78, 0x49, 0xd9, 0x76, 0x35, 0x78, 0x31,
	0x40, 0x5d, 0x0e, 0xe8, 0x18, 0x74, 0x67, 0x5e, 0x14, 0xd0, 0x91, 0x7b, 0xed, 0xd2, 0x31, 0xfc,
	0x87, 0xba, 0xa0, 0x35, 0xf3, 0x6e, 0x3c, 0xff, 0xce, 0x83, 0x86, 0x64, 0x23, 0xea, 0x84, 0xa3,
	0xe9, 0x63, 0xe4, 0x8e, 0x29, 0x34, 0xd1, 0x11, 0x00, 0xd5, 0x22, 0xf6, 0x03, 0x58, 0x43, 0x10,
	0xfc, 0xaf, 0x66, 0x3f, 0x9e, 0xd2, 0x10, 0xd6, 0xa5, 0x7f, 0xe4, 0x7b, 0x31, 0xf5, 0x62, 0xd8,
	0x40, 0xa7, 0x00, 0x55, 0x74, 0xe0, 0x84, 0xb1, 0x47, 0x43, 0x65, 0x6b, 0xa2, 0x33, 0x70, 0xf2,
	0x63, 0xaf, 0xed, 0x2d, 0xd4, 0x01, 0x8d, 0x5b, 0xf7, 0x9e, 0x8e, 0x61, 0x7b, 0xf8, 0x69, 0x80,
	0xfe, 0x33, 0x5f, 0xe1, 0x3f, 0x43, 0x0f, 0x3b, 0x32, 0x46, 0x20, 0x13, 0x06, 0xc6, 0xc3, 0xb0,
	0xd2, 0xe6, 0xbc, 0x48, 0x59, 0x8e, 0xf9, 0x26, 0xb7, 0xf2, 0x05, 0x53, 0xf9, 0xf7, 0xd5, 0xae,
	0x97, 0xe2, 0x97, 0xa6, 0xaf, 0xd4, 0xfb, 0x6a, 0xd6, 0x26, 0x8e, 0xf3, 0x66, 0xf6, 0x26, 0xfa,
	0x94, 0x93, 0x09, 0xac, 0xa1, 0x44, 0x73, 0x1b, 0xcb, 0xfe, 0xc4, 0xc7, 0x9e, 0x4f, 0x9c, 0x4c,
	0x24, 0x07, 0x3e, 0x99, 0xdb, 0x89, 0xe2, 0xbf, 0xcc, 0xbe, 0x5e, 0x12, 0xe2, 0x64, 0x82, 0x90,
	0x83, 0x82, 0x90, 0xb9, 0x4d, 0x88, 0xd2, 0x3c, 0x35, 0xd5, 0xc7, 0xec, 0xef, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x11, 0xde, 0x43, 0x57, 0x01, 0x02, 0x00, 0x00,
}
