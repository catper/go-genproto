// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/distance_bucket.proto

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

// The distance bucket for a user’s distance from an advertiser’s location
// extension.
type DistanceBucketEnum_DistanceBucket int32

const (
	// Not specified.
	DistanceBucketEnum_UNSPECIFIED DistanceBucketEnum_DistanceBucket = 0
	// Used for return value only. Represents value unknown in this version.
	DistanceBucketEnum_UNKNOWN DistanceBucketEnum_DistanceBucket = 1
	// User was within 700m of the location.
	DistanceBucketEnum_WITHIN_700M DistanceBucketEnum_DistanceBucket = 2
	// User was within 1KM of the location.
	DistanceBucketEnum_WITHIN_1KM DistanceBucketEnum_DistanceBucket = 3
	// User was within 5KM of the location.
	DistanceBucketEnum_WITHIN_5KM DistanceBucketEnum_DistanceBucket = 4
	// User was within 10KM of the location.
	DistanceBucketEnum_WITHIN_10KM DistanceBucketEnum_DistanceBucket = 5
	// User was within 15KM of the location.
	DistanceBucketEnum_WITHIN_15KM DistanceBucketEnum_DistanceBucket = 6
	// User was within 20KM of the location.
	DistanceBucketEnum_WITHIN_20KM DistanceBucketEnum_DistanceBucket = 7
	// User was within 25KM of the location.
	DistanceBucketEnum_WITHIN_25KM DistanceBucketEnum_DistanceBucket = 8
	// User was within 30KM of the location.
	DistanceBucketEnum_WITHIN_30KM DistanceBucketEnum_DistanceBucket = 9
	// User was within 35KM of the location.
	DistanceBucketEnum_WITHIN_35KM DistanceBucketEnum_DistanceBucket = 10
	// User was within 40KM of the location.
	DistanceBucketEnum_WITHIN_40KM DistanceBucketEnum_DistanceBucket = 11
	// User was within 45KM of the location.
	DistanceBucketEnum_WITHIN_45KM DistanceBucketEnum_DistanceBucket = 12
	// User was within 50KM of the location.
	DistanceBucketEnum_WITHIN_50KM DistanceBucketEnum_DistanceBucket = 13
	// User was within 55KM of the location.
	DistanceBucketEnum_WITHIN_55KM DistanceBucketEnum_DistanceBucket = 14
	// User was within 60KM of the location.
	DistanceBucketEnum_WITHIN_60KM DistanceBucketEnum_DistanceBucket = 15
	// User was within 65KM of the location.
	DistanceBucketEnum_WITHIN_65KM DistanceBucketEnum_DistanceBucket = 16
	// User was beyond 65KM of the location.
	DistanceBucketEnum_BEYOND_65KM DistanceBucketEnum_DistanceBucket = 17
	// User was within 0.7 miles of the location.
	DistanceBucketEnum_WITHIN_0_7MILES DistanceBucketEnum_DistanceBucket = 18
	// User was within 1 mile of the location.
	DistanceBucketEnum_WITHIN_1MILE DistanceBucketEnum_DistanceBucket = 19
	// User was within 5 miles of the location.
	DistanceBucketEnum_WITHIN_5MILES DistanceBucketEnum_DistanceBucket = 20
	// User was within 10 miles of the location.
	DistanceBucketEnum_WITHIN_10MILES DistanceBucketEnum_DistanceBucket = 21
	// User was within 15 miles of the location.
	DistanceBucketEnum_WITHIN_15MILES DistanceBucketEnum_DistanceBucket = 22
	// User was within 20 miles of the location.
	DistanceBucketEnum_WITHIN_20MILES DistanceBucketEnum_DistanceBucket = 23
	// User was within 25 miles of the location.
	DistanceBucketEnum_WITHIN_25MILES DistanceBucketEnum_DistanceBucket = 24
	// User was within 30 miles of the location.
	DistanceBucketEnum_WITHIN_30MILES DistanceBucketEnum_DistanceBucket = 25
	// User was within 35 miles of the location.
	DistanceBucketEnum_WITHIN_35MILES DistanceBucketEnum_DistanceBucket = 26
	// User was within 40 miles of the location.
	DistanceBucketEnum_WITHIN_40MILES DistanceBucketEnum_DistanceBucket = 27
	// User was beyond 40 miles of the location.
	DistanceBucketEnum_BEYOND_40MILES DistanceBucketEnum_DistanceBucket = 28
)

var DistanceBucketEnum_DistanceBucket_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "WITHIN_700M",
	3:  "WITHIN_1KM",
	4:  "WITHIN_5KM",
	5:  "WITHIN_10KM",
	6:  "WITHIN_15KM",
	7:  "WITHIN_20KM",
	8:  "WITHIN_25KM",
	9:  "WITHIN_30KM",
	10: "WITHIN_35KM",
	11: "WITHIN_40KM",
	12: "WITHIN_45KM",
	13: "WITHIN_50KM",
	14: "WITHIN_55KM",
	15: "WITHIN_60KM",
	16: "WITHIN_65KM",
	17: "BEYOND_65KM",
	18: "WITHIN_0_7MILES",
	19: "WITHIN_1MILE",
	20: "WITHIN_5MILES",
	21: "WITHIN_10MILES",
	22: "WITHIN_15MILES",
	23: "WITHIN_20MILES",
	24: "WITHIN_25MILES",
	25: "WITHIN_30MILES",
	26: "WITHIN_35MILES",
	27: "WITHIN_40MILES",
	28: "BEYOND_40MILES",
}

var DistanceBucketEnum_DistanceBucket_value = map[string]int32{
	"UNSPECIFIED":     0,
	"UNKNOWN":         1,
	"WITHIN_700M":     2,
	"WITHIN_1KM":      3,
	"WITHIN_5KM":      4,
	"WITHIN_10KM":     5,
	"WITHIN_15KM":     6,
	"WITHIN_20KM":     7,
	"WITHIN_25KM":     8,
	"WITHIN_30KM":     9,
	"WITHIN_35KM":     10,
	"WITHIN_40KM":     11,
	"WITHIN_45KM":     12,
	"WITHIN_50KM":     13,
	"WITHIN_55KM":     14,
	"WITHIN_60KM":     15,
	"WITHIN_65KM":     16,
	"BEYOND_65KM":     17,
	"WITHIN_0_7MILES": 18,
	"WITHIN_1MILE":    19,
	"WITHIN_5MILES":   20,
	"WITHIN_10MILES":  21,
	"WITHIN_15MILES":  22,
	"WITHIN_20MILES":  23,
	"WITHIN_25MILES":  24,
	"WITHIN_30MILES":  25,
	"WITHIN_35MILES":  26,
	"WITHIN_40MILES":  27,
	"BEYOND_40MILES":  28,
}

func (x DistanceBucketEnum_DistanceBucket) String() string {
	return proto.EnumName(DistanceBucketEnum_DistanceBucket_name, int32(x))
}

func (DistanceBucketEnum_DistanceBucket) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4666d419e0011116, []int{0, 0}
}

// Container for distance buckets of a user’s distance from an advertiser’s
// location extension.
type DistanceBucketEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DistanceBucketEnum) Reset()         { *m = DistanceBucketEnum{} }
func (m *DistanceBucketEnum) String() string { return proto.CompactTextString(m) }
func (*DistanceBucketEnum) ProtoMessage()    {}
func (*DistanceBucketEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_4666d419e0011116, []int{0}
}

func (m *DistanceBucketEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DistanceBucketEnum.Unmarshal(m, b)
}
func (m *DistanceBucketEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DistanceBucketEnum.Marshal(b, m, deterministic)
}
func (m *DistanceBucketEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistanceBucketEnum.Merge(m, src)
}
func (m *DistanceBucketEnum) XXX_Size() int {
	return xxx_messageInfo_DistanceBucketEnum.Size(m)
}
func (m *DistanceBucketEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_DistanceBucketEnum.DiscardUnknown(m)
}

var xxx_messageInfo_DistanceBucketEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.DistanceBucketEnum_DistanceBucket", DistanceBucketEnum_DistanceBucket_name, DistanceBucketEnum_DistanceBucket_value)
	proto.RegisterType((*DistanceBucketEnum)(nil), "google.ads.googleads.v3.enums.DistanceBucketEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/distance_bucket.proto", fileDescriptor_4666d419e0011116)
}

var fileDescriptor_4666d419e0011116 = []byte{
	// 456 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xdf, 0x4e, 0xdb, 0x30,
	0x14, 0x87, 0xd7, 0xd2, 0xc1, 0xe6, 0x42, 0x6b, 0xdc, 0xfd, 0x65, 0x70, 0x01, 0x0f, 0xe0, 0x14,
	0x0c, 0x43, 0xca, 0xae, 0x92, 0x35, 0x63, 0x51, 0x97, 0x50, 0x89, 0x51, 0xb4, 0xa9, 0x52, 0x15,
	0x9a, 0x28, 0xaa, 0x46, 0xed, 0x0a, 0xa7, 0x3c, 0xca, 0x1e, 0x60, 0x97, 0x93, 0xf6, 0x22, 0x7b,
	0x92, 0x69, 0x4f, 0x31, 0x9d, 0xd8, 0x89, 0x0e, 0x17, 0xe3, 0x26, 0x3a, 0xf9, 0xce, 0x77, 0x4e,
	0x62, 0xeb, 0x47, 0x44, 0xae, 0x54, 0x7e, 0x93, 0x39, 0x49, 0xaa, 0x1d, 0x53, 0x42, 0x75, 0x27,
	0x9c, 0x4c, 0xae, 0x16, 0xda, 0x49, 0xe7, 0xba, 0x48, 0xe4, 0x2c, 0x9b, 0x5e, 0xaf, 0x66, 0xdf,
	0xb2, 0x82, 0x2f, 0x6f, 0x55, 0xa1, 0xd8, 0x9e, 0x31, 0x79, 0x92, 0x6a, 0x5e, 0x0f, 0xf1, 0x3b,
	0xc1, 0xcb, 0xa1, 0x9d, 0xdd, 0x6a, 0xe7, 0x72, 0xee, 0x24, 0x52, 0xaa, 0x22, 0x29, 0xe6, 0x4a,
	0x6a, 0x33, 0x7c, 0xf0, 0xab, 0x45, 0xd8, 0xc0, 0xae, 0xf5, 0xcb, 0xad, 0x81, 0x5c, 0x2d, 0x0e,
	0xbe, 0xb7, 0x48, 0xe7, 0x3e, 0x66, 0x5d, 0xd2, 0xbe, 0x8c, 0x2f, 0x46, 0xc1, 0xfb, 0xf0, 0x43,
	0x18, 0x0c, 0xe8, 0x23, 0xd6, 0x26, 0x1b, 0x97, 0xf1, 0x30, 0x3e, 0xbf, 0x8a, 0x69, 0x03, 0xba,
	0x57, 0xe1, 0xe7, 0x8f, 0x61, 0x3c, 0x3d, 0xed, 0xf7, 0x23, 0xda, 0x64, 0x1d, 0x42, 0x2c, 0x38,
	0x1c, 0x46, 0x74, 0x0d, 0xbd, 0x9f, 0x0c, 0x23, 0xda, 0x42, 0x03, 0x87, 0xfd, 0x61, 0x44, 0x1f,
	0x63, 0x00, 0xc6, 0x3a, 0x02, 0x47, 0x60, 0x6c, 0x60, 0x00, 0xc6, 0x13, 0x04, 0x04, 0x18, 0x4f,
	0x31, 0x00, 0x83, 0x20, 0x70, 0x0c, 0x46, 0x1b, 0x03, 0x30, 0x36, 0x11, 0x38, 0x01, 0x63, 0x0b,
	0x03, 0x30, 0x3a, 0x08, 0xbc, 0x05, 0xa3, 0x8b, 0x01, 0x18, 0x14, 0x80, 0x1f, 0x7c, 0x39, 0x8f,
	0x07, 0x06, 0x6c, 0xb3, 0x1e, 0xe9, 0x5a, 0xa3, 0x3f, 0x3d, 0x8d, 0xc2, 0x4f, 0xc1, 0x05, 0x65,
	0x8c, 0x92, 0xcd, 0xea, 0x80, 0x80, 0x68, 0x8f, 0x6d, 0x93, 0xad, 0xea, 0x53, 0x46, 0x7a, 0xc6,
	0x18, 0xe9, 0xd4, 0xd7, 0x62, 0xd8, 0x73, 0xcc, 0xac, 0xf7, 0x02, 0xb1, 0x23, 0xeb, 0xbd, 0xc4,
	0xcc, 0x7a, 0xaf, 0x10, 0x13, 0xd6, 0x7b, 0x8d, 0x99, 0xf5, 0x76, 0x10, 0x3b, 0xb6, 0xde, 0x1b,
	0x60, 0xf6, 0x58, 0x15, 0xdb, 0xf5, 0xff, 0x34, 0xc8, 0xfe, 0x4c, 0x2d, 0xf8, 0x83, 0x99, 0xf3,
	0x7b, 0xf7, 0xb3, 0x33, 0x82, 0xa8, 0x8d, 0x1a, 0x5f, 0x7d, 0x3b, 0x95, 0xab, 0x9b, 0x44, 0xe6,
	0x5c, 0xdd, 0xe6, 0x4e, 0x9e, 0xc9, 0x32, 0x88, 0x55, 0xdc, 0x97, 0x73, 0xfd, 0x9f, 0xf4, 0xbf,
	0x2b, 0x9f, 0x3f, 0x9a, 0x6b, 0x67, 0x9e, 0xf7, 0xb3, 0xb9, 0x77, 0x66, 0x56, 0x79, 0xa9, 0xe6,
	0xa6, 0x84, 0x6a, 0x2c, 0x38, 0xc4, 0x57, 0xff, 0xae, 0xfa, 0x13, 0x2f, 0xd5, 0x93, 0xba, 0x3f,
	0x19, 0x8b, 0x49, 0xd9, 0xff, 0xdb, 0xdc, 0x37, 0xd0, 0x75, 0xbd, 0x54, 0xbb, 0x6e, 0x6d, 0xb8,
	0xee, 0x58, 0xb8, 0x6e, 0xe9, 0x5c, 0xaf, 0x97, 0x3f, 0x26, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff,
	0x98, 0xa9, 0x3d, 0xe9, 0x95, 0x03, 0x00, 0x00,
}
