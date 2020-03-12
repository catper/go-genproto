// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/distance_view.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
	wrappers "github.com/catper/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v2/enums"
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

// A distance view with metrics aggregated by the user's distance from an
// advertiser's location extensions. Each DistanceBucket includes all
// impressions that fall within its distance and a single impression will
// contribute to the metrics for all DistanceBuckets that include the user's
// distance.
type DistanceView struct {
	// Immutable. The resource name of the distance view.
	// Distance view resource names have the form:
	//
	// `customers/{customer_id}/distanceViews/1~{distance_bucket}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. Grouping of user distance from location extensions.
	DistanceBucket enums.DistanceBucketEnum_DistanceBucket `protobuf:"varint,2,opt,name=distance_bucket,json=distanceBucket,proto3,enum=google.ads.googleads.v2.enums.DistanceBucketEnum_DistanceBucket" json:"distance_bucket,omitempty"`
	// Output only. True if the DistanceBucket is using the metric system, false otherwise.
	MetricSystem         *wrappers.BoolValue `protobuf:"bytes,3,opt,name=metric_system,json=metricSystem,proto3" json:"metric_system,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *DistanceView) Reset()         { *m = DistanceView{} }
func (m *DistanceView) String() string { return proto.CompactTextString(m) }
func (*DistanceView) ProtoMessage()    {}
func (*DistanceView) Descriptor() ([]byte, []int) {
	return fileDescriptor_3430a19803c9c189, []int{0}
}

func (m *DistanceView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DistanceView.Unmarshal(m, b)
}
func (m *DistanceView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DistanceView.Marshal(b, m, deterministic)
}
func (m *DistanceView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistanceView.Merge(m, src)
}
func (m *DistanceView) XXX_Size() int {
	return xxx_messageInfo_DistanceView.Size(m)
}
func (m *DistanceView) XXX_DiscardUnknown() {
	xxx_messageInfo_DistanceView.DiscardUnknown(m)
}

var xxx_messageInfo_DistanceView proto.InternalMessageInfo

func (m *DistanceView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *DistanceView) GetDistanceBucket() enums.DistanceBucketEnum_DistanceBucket {
	if m != nil {
		return m.DistanceBucket
	}
	return enums.DistanceBucketEnum_UNSPECIFIED
}

func (m *DistanceView) GetMetricSystem() *wrappers.BoolValue {
	if m != nil {
		return m.MetricSystem
	}
	return nil
}

func init() {
	proto.RegisterType((*DistanceView)(nil), "google.ads.googleads.v2.resources.DistanceView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/distance_view.proto", fileDescriptor_3430a19803c9c189)
}

var fileDescriptor_3430a19803c9c189 = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x6b, 0xd4, 0x40,
	0x18, 0x26, 0x59, 0x14, 0x8c, 0xdb, 0x8a, 0x39, 0xad, 0x8b, 0xe8, 0x56, 0x58, 0xdc, 0x8b, 0x33,
	0x90, 0xe2, 0x65, 0xbc, 0x38, 0xa1, 0x52, 0xf0, 0x20, 0x65, 0x85, 0x1c, 0x74, 0x71, 0x99, 0x4d,
	0xde, 0xc6, 0xc1, 0xcc, 0x4c, 0x98, 0x49, 0x76, 0x91, 0xd2, 0x3f, 0xe3, 0xd1, 0x9f, 0xe2, 0xaf,
	0xe8, 0xb9, 0xbf, 0x40, 0x3c, 0x88, 0xe4, 0x63, 0xa6, 0x69, 0xa1, 0xea, 0xed, 0xc9, 0x3c, 0x1f,
	0xef, 0x47, 0xde, 0xe0, 0x65, 0xae, 0x54, 0x5e, 0x00, 0x66, 0x99, 0xc1, 0x1d, 0x6c, 0xd0, 0x36,
	0xc2, 0x1a, 0x8c, 0xaa, 0x75, 0x0a, 0x06, 0x67, 0xdc, 0x54, 0x4c, 0xa6, 0xb0, 0xde, 0x72, 0xd8,
	0xa1, 0x52, 0xab, 0x4a, 0x85, 0x07, 0x9d, 0x16, 0xb1, 0xcc, 0x20, 0x67, 0x43, 0xdb, 0x08, 0x39,
	0xdb, 0xf4, 0xf0, 0xb6, 0x64, 0x90, 0xb5, 0x18, 0xa4, 0x6e, 0xea, 0xf4, 0x0b, 0x54, 0x5d, 0xee,
	0xf4, 0xa9, 0x35, 0x95, 0x1c, 0x9f, 0x72, 0x28, 0xb2, 0xf5, 0x06, 0x3e, 0xb3, 0x2d, 0x57, 0xba,
	0x17, 0x3c, 0x1a, 0x08, 0x6c, 0xad, 0x9e, 0x7a, 0xd2, 0x53, 0xed, 0xd7, 0xa6, 0x3e, 0xc5, 0x3b,
	0xcd, 0xca, 0x12, 0xb4, 0xe9, 0xf9, 0xc7, 0x03, 0x2b, 0x93, 0x52, 0x55, 0xac, 0xe2, 0x4a, 0xf6,
	0xec, 0xb3, 0x9f, 0x7e, 0x30, 0x3e, 0xea, 0x7b, 0x4a, 0x38, 0xec, 0xc2, 0x65, 0xb0, 0x67, 0x0b,
	0xac, 0x25, 0x13, 0x30, 0xf1, 0x66, 0xde, 0xe2, 0x5e, 0xfc, 0xe2, 0x82, 0xde, 0xf9, 0x45, 0x9f,
	0x07, 0xf3, 0xab, 0xb1, 0x7b, 0x54, 0x72, 0x83, 0x52, 0x25, 0xf0, 0x30, 0x65, 0x39, 0xb6, 0x19,
	0xef, 0x98, 0x80, 0x50, 0x06, 0x0f, 0x6e, 0xcc, 0x3d, 0xf1, 0x67, 0xde, 0x62, 0x3f, 0x7a, 0x8d,
	0x6e, 0x5b, 0x68, 0xbb, 0x2d, 0x64, 0x33, 0xe3, 0xd6, 0xf4, 0x46, 0xd6, 0xe2, 0xc6, 0x53, 0x3c,
	0xba, 0xa0, 0xa3, 0xe5, 0x7e, 0x76, 0xed, 0x31, 0x3c, 0x0a, 0xf6, 0x04, 0x54, 0x9a, 0xa7, 0x6b,
	0xf3, 0xd5, 0x54, 0x20, 0x26, 0xa3, 0x99, 0xb7, 0xb8, 0x1f, 0x4d, 0x6d, 0x35, 0xbb, 0x2a, 0x14,
	0x2b, 0x55, 0x24, 0xac, 0xa8, 0xa1, 0xcb, 0x19, 0x77, 0xae, 0xf7, 0xad, 0x89, 0x7c, 0xba, 0xa4,
	0x1f, 0xff, 0x73, 0xde, 0x30, 0x4a, 0x6b, 0x53, 0x29, 0x01, 0xda, 0xe0, 0x33, 0x0b, 0xcf, 0xdd,
	0xcf, 0x6e, 0x24, 0x06, 0x9f, 0x5d, 0xbb, 0xa8, 0xf3, 0xf8, 0xb7, 0x17, 0xcc, 0x53, 0x25, 0xd0,
	0x3f, 0x6f, 0x2a, 0x7e, 0x38, 0xac, 0x75, 0xd2, 0x34, 0x7f, 0xe2, 0x7d, 0x78, 0xdb, 0xfb, 0x72,
	0x55, 0x30, 0x99, 0x23, 0xa5, 0x73, 0x9c, 0x83, 0x6c, 0x47, 0xc3, 0x57, 0xad, 0xfe, 0xe5, 0xc2,
	0x5f, 0x39, 0xf4, 0xcd, 0x1f, 0x1d, 0x53, 0xfa, 0xdd, 0x3f, 0x38, 0xee, 0x22, 0x69, 0x66, 0x50,
	0x07, 0x1b, 0x94, 0x44, 0x68, 0x69, 0x95, 0x3f, 0xac, 0x66, 0x45, 0x33, 0xb3, 0x72, 0x9a, 0x55,
	0x12, 0xad, 0x9c, 0xe6, 0xd2, 0x9f, 0x77, 0x04, 0x21, 0x34, 0x33, 0x84, 0x38, 0x15, 0x21, 0x49,
	0x44, 0x88, 0xd3, 0x6d, 0xee, 0xb6, 0xcd, 0x1e, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x40, 0xd0,
	0xa8, 0x55, 0x8d, 0x03, 0x00, 0x00,
}
