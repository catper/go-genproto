// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/proximity_radius_units.proto

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

// The unit of radius distance in proximity (e.g. MILES)
type ProximityRadiusUnitsEnum_ProximityRadiusUnits int32

const (
	// Not specified.
	ProximityRadiusUnitsEnum_UNSPECIFIED ProximityRadiusUnitsEnum_ProximityRadiusUnits = 0
	// Used for return value only. Represents value unknown in this version.
	ProximityRadiusUnitsEnum_UNKNOWN ProximityRadiusUnitsEnum_ProximityRadiusUnits = 1
	// Miles
	ProximityRadiusUnitsEnum_MILES ProximityRadiusUnitsEnum_ProximityRadiusUnits = 2
	// Kilometers
	ProximityRadiusUnitsEnum_KILOMETERS ProximityRadiusUnitsEnum_ProximityRadiusUnits = 3
)

var ProximityRadiusUnitsEnum_ProximityRadiusUnits_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "MILES",
	3: "KILOMETERS",
}

var ProximityRadiusUnitsEnum_ProximityRadiusUnits_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"MILES":       2,
	"KILOMETERS":  3,
}

func (x ProximityRadiusUnitsEnum_ProximityRadiusUnits) String() string {
	return proto.EnumName(ProximityRadiusUnitsEnum_ProximityRadiusUnits_name, int32(x))
}

func (ProximityRadiusUnitsEnum_ProximityRadiusUnits) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a987958cbc0f39f4, []int{0, 0}
}

// Container for enum describing unit of radius in proximity.
type ProximityRadiusUnitsEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProximityRadiusUnitsEnum) Reset()         { *m = ProximityRadiusUnitsEnum{} }
func (m *ProximityRadiusUnitsEnum) String() string { return proto.CompactTextString(m) }
func (*ProximityRadiusUnitsEnum) ProtoMessage()    {}
func (*ProximityRadiusUnitsEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_a987958cbc0f39f4, []int{0}
}

func (m *ProximityRadiusUnitsEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProximityRadiusUnitsEnum.Unmarshal(m, b)
}
func (m *ProximityRadiusUnitsEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProximityRadiusUnitsEnum.Marshal(b, m, deterministic)
}
func (m *ProximityRadiusUnitsEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProximityRadiusUnitsEnum.Merge(m, src)
}
func (m *ProximityRadiusUnitsEnum) XXX_Size() int {
	return xxx_messageInfo_ProximityRadiusUnitsEnum.Size(m)
}
func (m *ProximityRadiusUnitsEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ProximityRadiusUnitsEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ProximityRadiusUnitsEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.ProximityRadiusUnitsEnum_ProximityRadiusUnits", ProximityRadiusUnitsEnum_ProximityRadiusUnits_name, ProximityRadiusUnitsEnum_ProximityRadiusUnits_value)
	proto.RegisterType((*ProximityRadiusUnitsEnum)(nil), "google.ads.googleads.v3.enums.ProximityRadiusUnitsEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/proximity_radius_units.proto", fileDescriptor_a987958cbc0f39f4)
}

var fileDescriptor_a987958cbc0f39f4 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0xc3, 0x30,
	0x1c, 0xc6, 0x5d, 0x87, 0x8a, 0x19, 0x68, 0x29, 0x1e, 0x54, 0xdc, 0x61, 0x7b, 0x80, 0xe4, 0x90,
	0x5b, 0x3c, 0x75, 0x1a, 0x47, 0xd9, 0xd6, 0x95, 0xcd, 0x4d, 0x90, 0xc2, 0x88, 0x66, 0x84, 0xe0,
	0x9a, 0x94, 0xa6, 0x1d, 0xfa, 0x3a, 0x1e, 0x7d, 0x14, 0x1f, 0xc5, 0x83, 0xcf, 0x20, 0x4d, 0xd7,
	0x9e, 0xa6, 0x97, 0xf2, 0xd1, 0xef, 0xff, 0xfd, 0xf2, 0xfd, 0xff, 0x80, 0x08, 0xad, 0xc5, 0x66,
	0x8d, 0x18, 0x37, 0xa8, 0x92, 0xa5, 0xda, 0x62, 0xb4, 0x56, 0x45, 0x62, 0x50, 0x9a, 0xe9, 0x37,
	0x99, 0xc8, 0xfc, 0x7d, 0x95, 0x31, 0x2e, 0x0b, 0xb3, 0x2a, 0x94, 0xcc, 0x0d, 0x4c, 0x33, 0x9d,
	0x6b, 0xaf, 0x5b, 0x05, 0x20, 0xe3, 0x06, 0x36, 0x59, 0xb8, 0xc5, 0xd0, 0x66, 0xaf, 0xae, 0x6b,
	0x74, 0x2a, 0x11, 0x53, 0x4a, 0xe7, 0x2c, 0x97, 0x5a, 0xed, 0xc2, 0xfd, 0x57, 0x70, 0x11, 0xd5,
	0xf0, 0x99, 0x65, 0x2f, 0x4a, 0x34, 0x55, 0x45, 0xd2, 0x9f, 0x82, 0xf3, 0x7d, 0x9e, 0x77, 0x06,
	0x3a, 0x8b, 0x70, 0x1e, 0xd1, 0xdb, 0xe0, 0x3e, 0xa0, 0x77, 0xee, 0x81, 0xd7, 0x01, 0xc7, 0x8b,
	0x70, 0x14, 0x4e, 0x1f, 0x43, 0xb7, 0xe5, 0x9d, 0x80, 0xc3, 0x49, 0x30, 0xa6, 0x73, 0xd7, 0xf1,
	0x4e, 0x01, 0x18, 0x05, 0xe3, 0xe9, 0x84, 0x3e, 0xd0, 0xd9, 0xdc, 0x6d, 0x0f, 0x7e, 0x5a, 0xa0,
	0xf7, 0xa2, 0x13, 0xf8, 0x6f, 0xe1, 0xc1, 0xe5, 0xbe, 0x47, 0xa3, 0xb2, 0x6d, 0xd4, 0x7a, 0x1a,
	0xec, 0xb2, 0x42, 0x6f, 0x98, 0x12, 0x50, 0x67, 0x02, 0x89, 0xb5, 0xb2, 0xbb, 0xd4, 0x87, 0x4b,
	0xa5, 0xf9, 0xe3, 0x8e, 0x37, 0xf6, 0xfb, 0xe1, 0xb4, 0x87, 0xbe, 0xff, 0xe9, 0x74, 0x87, 0x15,
	0xca, 0xe7, 0x06, 0x56, 0xb2, 0x54, 0x4b, 0x0c, 0xcb, 0xe5, 0xcd, 0x57, 0xed, 0xc7, 0x3e, 0x37,
	0x71, 0xe3, 0xc7, 0x4b, 0x1c, 0x5b, 0xff, 0xdb, 0xe9, 0x55, 0x3f, 0x09, 0xf1, 0xb9, 0x21, 0xa4,
	0x99, 0x20, 0x64, 0x89, 0x09, 0xb1, 0x33, 0xcf, 0x47, 0xb6, 0x18, 0xfe, 0x0d, 0x00, 0x00, 0xff,
	0xff, 0xcc, 0x1e, 0x70, 0x1d, 0xdf, 0x01, 0x00, 0x00,
}
