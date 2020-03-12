// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/hotel_price_bucket.proto

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

// Enum describing possible hotel price buckets.
type HotelPriceBucketEnum_HotelPriceBucket int32

const (
	// Not specified.
	HotelPriceBucketEnum_UNSPECIFIED HotelPriceBucketEnum_HotelPriceBucket = 0
	// The value is unknown in this version.
	HotelPriceBucketEnum_UNKNOWN HotelPriceBucketEnum_HotelPriceBucket = 1
	// Tied for lowest price. Partner is within a small variance of the lowest
	// price.
	HotelPriceBucketEnum_LOWEST_TIED HotelPriceBucketEnum_HotelPriceBucket = 3
	// Not lowest price. Partner is not within a small variance of the lowest
	// price.
	HotelPriceBucketEnum_NOT_LOWEST HotelPriceBucketEnum_HotelPriceBucket = 4
)

var HotelPriceBucketEnum_HotelPriceBucket_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	3: "LOWEST_TIED",
	4: "NOT_LOWEST",
}

var HotelPriceBucketEnum_HotelPriceBucket_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"LOWEST_TIED": 3,
	"NOT_LOWEST":  4,
}

func (x HotelPriceBucketEnum_HotelPriceBucket) String() string {
	return proto.EnumName(HotelPriceBucketEnum_HotelPriceBucket_name, int32(x))
}

func (HotelPriceBucketEnum_HotelPriceBucket) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bd2118b8b82d2ea6, []int{0, 0}
}

// Container for enum describing hotel price bucket for a hotel itinerary.
type HotelPriceBucketEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HotelPriceBucketEnum) Reset()         { *m = HotelPriceBucketEnum{} }
func (m *HotelPriceBucketEnum) String() string { return proto.CompactTextString(m) }
func (*HotelPriceBucketEnum) ProtoMessage()    {}
func (*HotelPriceBucketEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd2118b8b82d2ea6, []int{0}
}

func (m *HotelPriceBucketEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HotelPriceBucketEnum.Unmarshal(m, b)
}
func (m *HotelPriceBucketEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HotelPriceBucketEnum.Marshal(b, m, deterministic)
}
func (m *HotelPriceBucketEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HotelPriceBucketEnum.Merge(m, src)
}
func (m *HotelPriceBucketEnum) XXX_Size() int {
	return xxx_messageInfo_HotelPriceBucketEnum.Size(m)
}
func (m *HotelPriceBucketEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_HotelPriceBucketEnum.DiscardUnknown(m)
}

var xxx_messageInfo_HotelPriceBucketEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.HotelPriceBucketEnum_HotelPriceBucket", HotelPriceBucketEnum_HotelPriceBucket_name, HotelPriceBucketEnum_HotelPriceBucket_value)
	proto.RegisterType((*HotelPriceBucketEnum)(nil), "google.ads.googleads.v2.enums.HotelPriceBucketEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/hotel_price_bucket.proto", fileDescriptor_bd2118b8b82d2ea6)
}

var fileDescriptor_bd2118b8b82d2ea6 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xcf, 0x4a, 0xc3, 0x30,
	0x18, 0x77, 0x9b, 0x28, 0x64, 0xa0, 0xa5, 0xe8, 0x45, 0xdc, 0x61, 0x7b, 0x80, 0x04, 0x2a, 0x78,
	0x88, 0xa7, 0x56, 0xeb, 0x1c, 0x4a, 0x57, 0x59, 0xd7, 0x81, 0x14, 0x4a, 0xd7, 0x86, 0x58, 0x6c,
	0x93, 0xd2, 0xb4, 0x7b, 0x20, 0x8f, 0x3e, 0x8a, 0x8f, 0xb2, 0xa7, 0x90, 0x24, 0xb6, 0x87, 0x81,
	0x5e, 0xc2, 0x8f, 0xef, 0xf7, 0x27, 0xbf, 0xef, 0x03, 0xb7, 0x94, 0x73, 0x5a, 0x10, 0x94, 0x64,
	0x02, 0x69, 0x28, 0xd1, 0xce, 0x42, 0x84, 0xb5, 0xa5, 0x40, 0xef, 0xbc, 0x21, 0x45, 0x5c, 0xd5,
	0x79, 0x4a, 0xe2, 0x6d, 0x9b, 0x7e, 0x90, 0x06, 0x56, 0x35, 0x6f, 0xb8, 0x39, 0xd1, 0x62, 0x98,
	0x64, 0x02, 0xf6, 0x3e, 0xb8, 0xb3, 0xa0, 0xf2, 0x5d, 0x5d, 0x77, 0xb1, 0x55, 0x8e, 0x12, 0xc6,
	0x78, 0x93, 0x34, 0x39, 0x67, 0x42, 0x9b, 0x67, 0x39, 0xb8, 0x78, 0x92, 0xc1, 0xbe, 0xcc, 0x75,
	0x54, 0xac, 0xcb, 0xda, 0x72, 0xf6, 0x0a, 0x8c, 0xc3, 0xb9, 0x79, 0x0e, 0xc6, 0x6b, 0x6f, 0xe5,
	0xbb, 0xf7, 0x8b, 0xc7, 0x85, 0xfb, 0x60, 0x1c, 0x99, 0x63, 0x70, 0xba, 0xf6, 0x9e, 0xbd, 0xe5,
	0xc6, 0x33, 0x06, 0x92, 0x7d, 0x59, 0x6e, 0xdc, 0x55, 0x10, 0x07, 0x92, 0x1d, 0x99, 0x67, 0x00,
	0x78, 0xcb, 0x20, 0xd6, 0x43, 0xe3, 0xd8, 0xd9, 0x0f, 0xc0, 0x34, 0xe5, 0x25, 0xfc, 0xb7, 0xae,
	0x73, 0x79, 0xf8, 0xad, 0x2f, 0x7b, 0xfa, 0x83, 0x37, 0xe7, 0xd7, 0x47, 0x79, 0x91, 0x30, 0x0a,
	0x79, 0x4d, 0x11, 0x25, 0x4c, 0x6d, 0xd1, 0x9d, 0xab, 0xca, 0xc5, 0x1f, 0xd7, 0xbb, 0x53, 0xef,
	0xe7, 0x70, 0x34, 0xb7, 0xed, 0xaf, 0xe1, 0x64, 0xae, 0xa3, 0xec, 0x4c, 0x40, 0x0d, 0x25, 0x0a,
	0x2d, 0x28, 0x57, 0x17, 0xdf, 0x1d, 0x1f, 0xd9, 0x99, 0x88, 0x7a, 0x3e, 0x0a, 0xad, 0x48, 0xf1,
	0xfb, 0xe1, 0x54, 0x0f, 0x31, 0xb6, 0x33, 0x81, 0x71, 0xaf, 0xc0, 0x38, 0xb4, 0x30, 0x56, 0x9a,
	0xed, 0x89, 0x2a, 0x76, 0xf3, 0x13, 0x00, 0x00, 0xff, 0xff, 0x33, 0x73, 0x88, 0x9b, 0xd5, 0x01,
	0x00, 0x00,
}
