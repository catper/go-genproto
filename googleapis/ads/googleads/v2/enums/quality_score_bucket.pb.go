// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/quality_score_bucket.proto

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

// Enum listing the possible quality score buckets.
type QualityScoreBucketEnum_QualityScoreBucket int32

const (
	// Not specified.
	QualityScoreBucketEnum_UNSPECIFIED QualityScoreBucketEnum_QualityScoreBucket = 0
	// Used for return value only. Represents value unknown in this version.
	QualityScoreBucketEnum_UNKNOWN QualityScoreBucketEnum_QualityScoreBucket = 1
	// Quality of the creative is below average.
	QualityScoreBucketEnum_BELOW_AVERAGE QualityScoreBucketEnum_QualityScoreBucket = 2
	// Quality of the creative is average.
	QualityScoreBucketEnum_AVERAGE QualityScoreBucketEnum_QualityScoreBucket = 3
	// Quality of the creative is above average.
	QualityScoreBucketEnum_ABOVE_AVERAGE QualityScoreBucketEnum_QualityScoreBucket = 4
)

var QualityScoreBucketEnum_QualityScoreBucket_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "BELOW_AVERAGE",
	3: "AVERAGE",
	4: "ABOVE_AVERAGE",
}

var QualityScoreBucketEnum_QualityScoreBucket_value = map[string]int32{
	"UNSPECIFIED":   0,
	"UNKNOWN":       1,
	"BELOW_AVERAGE": 2,
	"AVERAGE":       3,
	"ABOVE_AVERAGE": 4,
}

func (x QualityScoreBucketEnum_QualityScoreBucket) String() string {
	return proto.EnumName(QualityScoreBucketEnum_QualityScoreBucket_name, int32(x))
}

func (QualityScoreBucketEnum_QualityScoreBucket) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f6336f760aaed82f, []int{0, 0}
}

// The relative performance compared to other advertisers.
type QualityScoreBucketEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QualityScoreBucketEnum) Reset()         { *m = QualityScoreBucketEnum{} }
func (m *QualityScoreBucketEnum) String() string { return proto.CompactTextString(m) }
func (*QualityScoreBucketEnum) ProtoMessage()    {}
func (*QualityScoreBucketEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6336f760aaed82f, []int{0}
}

func (m *QualityScoreBucketEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QualityScoreBucketEnum.Unmarshal(m, b)
}
func (m *QualityScoreBucketEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QualityScoreBucketEnum.Marshal(b, m, deterministic)
}
func (m *QualityScoreBucketEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QualityScoreBucketEnum.Merge(m, src)
}
func (m *QualityScoreBucketEnum) XXX_Size() int {
	return xxx_messageInfo_QualityScoreBucketEnum.Size(m)
}
func (m *QualityScoreBucketEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_QualityScoreBucketEnum.DiscardUnknown(m)
}

var xxx_messageInfo_QualityScoreBucketEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.QualityScoreBucketEnum_QualityScoreBucket", QualityScoreBucketEnum_QualityScoreBucket_name, QualityScoreBucketEnum_QualityScoreBucket_value)
	proto.RegisterType((*QualityScoreBucketEnum)(nil), "google.ads.googleads.v2.enums.QualityScoreBucketEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/quality_score_bucket.proto", fileDescriptor_f6336f760aaed82f)
}

var fileDescriptor_f6336f760aaed82f = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xd1, 0x4a, 0xfb, 0x30,
	0x18, 0xc5, 0xff, 0xeb, 0xfe, 0x28, 0x64, 0x88, 0xb3, 0x17, 0x0a, 0xe2, 0x2e, 0xb6, 0x07, 0x48,
	0xa0, 0xde, 0x48, 0xbc, 0x4a, 0x34, 0x8e, 0xa1, 0x74, 0xd3, 0xb1, 0x0e, 0xa4, 0x30, 0xb2, 0x35,
	0x84, 0xe2, 0x96, 0xcc, 0xa5, 0x1d, 0x78, 0xe5, 0xbb, 0x78, 0xe9, 0xa3, 0xf8, 0x28, 0xe2, 0x43,
	0x48, 0x12, 0xdb, 0x9b, 0xa1, 0x37, 0xe5, 0x34, 0xbf, 0xef, 0x1c, 0xbe, 0xef, 0x80, 0x0b, 0xa9,
	0xb5, 0x5c, 0x0a, 0xc4, 0x33, 0x83, 0xbc, 0xb4, 0x6a, 0x1b, 0x21, 0xa1, 0xca, 0x95, 0x41, 0xcf,
	0x25, 0x5f, 0xe6, 0xc5, 0xcb, 0xcc, 0x2c, 0xf4, 0x46, 0xcc, 0xe6, 0xe5, 0xe2, 0x49, 0x14, 0x70,
	0xbd, 0xd1, 0x85, 0x0e, 0x3b, 0x7e, 0x1c, 0xf2, 0xcc, 0xc0, 0xda, 0x09, 0xb7, 0x11, 0x74, 0xce,
	0xd3, 0xb3, 0x2a, 0x78, 0x9d, 0x23, 0xae, 0x94, 0x2e, 0x78, 0x91, 0x6b, 0x65, 0xbc, 0xb9, 0xf7,
	0x0a, 0x8e, 0xef, 0x7d, 0xf4, 0xd8, 0x26, 0x53, 0x17, 0xcc, 0x54, 0xb9, 0xea, 0x09, 0x10, 0xee,
	0x92, 0xf0, 0x10, 0xb4, 0x26, 0xf1, 0x78, 0xc4, 0xae, 0x06, 0x37, 0x03, 0x76, 0xdd, 0xfe, 0x17,
	0xb6, 0xc0, 0xfe, 0x24, 0xbe, 0x8d, 0x87, 0xd3, 0xb8, 0xdd, 0x08, 0x8f, 0xc0, 0x01, 0x65, 0x77,
	0xc3, 0xe9, 0x8c, 0x24, 0xec, 0x81, 0xf4, 0x59, 0x3b, 0xb0, 0xbc, 0xfa, 0x69, 0x5a, 0x4e, 0xe8,
	0x30, 0x61, 0x35, 0xff, 0x4f, 0xbf, 0x1a, 0xa0, 0xbb, 0xd0, 0x2b, 0xf8, 0xe7, 0x11, 0xf4, 0x64,
	0x77, 0x95, 0x91, 0xdd, 0x7f, 0xd4, 0x78, 0xa4, 0x3f, 0x4e, 0xa9, 0x97, 0x5c, 0x49, 0xa8, 0x37,
	0x12, 0x49, 0xa1, 0xdc, 0x75, 0x55, 0x91, 0xeb, 0xdc, 0xfc, 0xd2, 0xeb, 0xa5, 0xfb, 0xbe, 0x05,
	0xcd, 0x3e, 0x21, 0xef, 0x41, 0xa7, 0xef, 0xa3, 0x48, 0x66, 0xa0, 0x97, 0x56, 0x25, 0x11, 0xb4,
	0x85, 0x98, 0x8f, 0x8a, 0xa7, 0x24, 0x33, 0x69, 0xcd, 0xd3, 0x24, 0x4a, 0x1d, 0xff, 0x0c, 0xba,
	0xfe, 0x11, 0x63, 0x92, 0x19, 0x8c, 0xeb, 0x09, 0x8c, 0x93, 0x08, 0x63, 0x37, 0x33, 0xdf, 0x73,
	0x8b, 0x9d, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x4a, 0xf4, 0x2c, 0xef, 0x01, 0x00, 0x00,
}
