// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/location_group_radius_units.proto

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

// The unit of radius distance in location group (e.g. MILES)
type LocationGroupRadiusUnitsEnum_LocationGroupRadiusUnits int32

const (
	// Not specified.
	LocationGroupRadiusUnitsEnum_UNSPECIFIED LocationGroupRadiusUnitsEnum_LocationGroupRadiusUnits = 0
	// Used for return value only. Represents value unknown in this version.
	LocationGroupRadiusUnitsEnum_UNKNOWN LocationGroupRadiusUnitsEnum_LocationGroupRadiusUnits = 1
	// Meters
	LocationGroupRadiusUnitsEnum_METERS LocationGroupRadiusUnitsEnum_LocationGroupRadiusUnits = 2
	// Miles
	LocationGroupRadiusUnitsEnum_MILES LocationGroupRadiusUnitsEnum_LocationGroupRadiusUnits = 3
)

var LocationGroupRadiusUnitsEnum_LocationGroupRadiusUnits_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "METERS",
	3: "MILES",
}

var LocationGroupRadiusUnitsEnum_LocationGroupRadiusUnits_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"METERS":      2,
	"MILES":       3,
}

func (x LocationGroupRadiusUnitsEnum_LocationGroupRadiusUnits) String() string {
	return proto.EnumName(LocationGroupRadiusUnitsEnum_LocationGroupRadiusUnits_name, int32(x))
}

func (LocationGroupRadiusUnitsEnum_LocationGroupRadiusUnits) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4d3468bd973fd435, []int{0, 0}
}

// Container for enum describing unit of radius in location group.
type LocationGroupRadiusUnitsEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocationGroupRadiusUnitsEnum) Reset()         { *m = LocationGroupRadiusUnitsEnum{} }
func (m *LocationGroupRadiusUnitsEnum) String() string { return proto.CompactTextString(m) }
func (*LocationGroupRadiusUnitsEnum) ProtoMessage()    {}
func (*LocationGroupRadiusUnitsEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d3468bd973fd435, []int{0}
}

func (m *LocationGroupRadiusUnitsEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocationGroupRadiusUnitsEnum.Unmarshal(m, b)
}
func (m *LocationGroupRadiusUnitsEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocationGroupRadiusUnitsEnum.Marshal(b, m, deterministic)
}
func (m *LocationGroupRadiusUnitsEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocationGroupRadiusUnitsEnum.Merge(m, src)
}
func (m *LocationGroupRadiusUnitsEnum) XXX_Size() int {
	return xxx_messageInfo_LocationGroupRadiusUnitsEnum.Size(m)
}
func (m *LocationGroupRadiusUnitsEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_LocationGroupRadiusUnitsEnum.DiscardUnknown(m)
}

var xxx_messageInfo_LocationGroupRadiusUnitsEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.LocationGroupRadiusUnitsEnum_LocationGroupRadiusUnits", LocationGroupRadiusUnitsEnum_LocationGroupRadiusUnits_name, LocationGroupRadiusUnitsEnum_LocationGroupRadiusUnits_value)
	proto.RegisterType((*LocationGroupRadiusUnitsEnum)(nil), "google.ads.googleads.v1.enums.LocationGroupRadiusUnitsEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/location_group_radius_units.proto", fileDescriptor_4d3468bd973fd435)
}

var fileDescriptor_4d3468bd973fd435 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xcf, 0x6a, 0xf2, 0x40,
	0x14, 0xc5, 0x3f, 0x23, 0x9f, 0xa5, 0xe3, 0xa2, 0x21, 0xab, 0x52, 0x74, 0xa1, 0x0f, 0x30, 0x43,
	0xe8, 0x6e, 0xba, 0x28, 0xb1, 0x9d, 0x8a, 0x54, 0xa3, 0x68, 0xb5, 0x50, 0x02, 0x32, 0x75, 0x64,
	0x08, 0x98, 0xb9, 0x21, 0x93, 0xf8, 0x40, 0x5d, 0xf6, 0x51, 0xfa, 0x28, 0x5d, 0xf6, 0x09, 0xca,
	0x4c, 0xfe, 0xec, 0xd2, 0x4d, 0x38, 0xe4, 0xdc, 0xfb, 0xbb, 0x67, 0x0e, 0xba, 0x97, 0x00, 0xf2,
	0x74, 0x24, 0x5c, 0x68, 0x52, 0x4a, 0xa3, 0xce, 0x3e, 0x39, 0xaa, 0x22, 0xd1, 0xe4, 0x04, 0x07,
	0x9e, 0xc7, 0xa0, 0xf6, 0x32, 0x83, 0x22, 0xdd, 0x67, 0x5c, 0xc4, 0x85, 0xde, 0x17, 0x2a, 0xce,
	0x35, 0x4e, 0x33, 0xc8, 0xc1, 0x1b, 0x96, 0x5b, 0x98, 0x0b, 0x8d, 0x1b, 0x00, 0x3e, 0xfb, 0xd8,
	0x02, 0x6e, 0x06, 0x35, 0x3f, 0x8d, 0x09, 0x57, 0x0a, 0x72, 0x8b, 0xab, 0x96, 0xc7, 0x80, 0x06,
	0xf3, 0xea, 0xc2, 0xd4, 0x1c, 0x58, 0x5b, 0xfe, 0xd6, 0xe0, 0x99, 0x2a, 0x92, 0xf1, 0x12, 0x5d,
	0xb7, 0xf9, 0xde, 0x15, 0xea, 0x6f, 0xc3, 0xcd, 0x8a, 0x3d, 0xcc, 0x9e, 0x66, 0xec, 0xd1, 0xfd,
	0xe7, 0xf5, 0xd1, 0xc5, 0x36, 0x7c, 0x0e, 0x97, 0xaf, 0xa1, 0xdb, 0xf1, 0x10, 0xea, 0x2d, 0xd8,
	0x0b, 0x5b, 0x6f, 0x5c, 0xc7, 0xbb, 0x44, 0xff, 0x17, 0xb3, 0x39, 0xdb, 0xb8, 0xdd, 0xc9, 0x4f,
	0x07, 0x8d, 0x0e, 0x90, 0xe0, 0x3f, 0x43, 0x4f, 0x86, 0x6d, 0x47, 0x57, 0x26, 0xf5, 0xaa, 0xf3,
	0x36, 0xa9, 0xf6, 0x25, 0x9c, 0xb8, 0x92, 0x18, 0x32, 0x49, 0xe4, 0x51, 0xd9, 0x37, 0xd5, 0x2d,
	0xa6, 0xb1, 0x6e, 0x29, 0xf5, 0xce, 0x7e, 0x3f, 0x9c, 0xee, 0x34, 0x08, 0x3e, 0x9d, 0xe1, 0xb4,
	0x44, 0x05, 0x42, 0xe3, 0x52, 0x1a, 0xb5, 0xf3, 0xb1, 0x29, 0x40, 0x7f, 0xd5, 0x7e, 0x14, 0x08,
	0x1d, 0x35, 0x7e, 0xb4, 0xf3, 0x23, 0xeb, 0x7f, 0x3b, 0xa3, 0xf2, 0x27, 0xa5, 0x81, 0xd0, 0x94,
	0x36, 0x13, 0x94, 0xee, 0x7c, 0x4a, 0xed, 0xcc, 0x7b, 0xcf, 0x06, 0xbb, 0xfd, 0x0d, 0x00, 0x00,
	0xff, 0xff, 0x92, 0x60, 0x22, 0x8a, 0xec, 0x01, 0x00, 0x00,
}
