// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/geographic_view.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
	wrappers "github.com/catper/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
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

// A geographic view.
//
// Geographic View includes all metrics aggregated at the country level,
// one row per country. It reports metrics at either actual physical location of
// the user or an area of interest. If other segment fields are used, you may
// get more than one row per country.
type GeographicView struct {
	// Immutable. The resource name of the geographic view.
	// Geographic view resource names have the form:
	//
	// `customers/{customer_id}/geographicViews/{country_criterion_id}~{location_type}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. CriterionId for the geo target for a country.
	CountryGeoTargetConstant *wrappers.StringValue `protobuf:"bytes,2,opt,name=country_geo_target_constant,json=countryGeoTargetConstant,proto3" json:"country_geo_target_constant,omitempty"`
	// Output only. Type of the geo targeting of the campaign.
	LocationType         enums.GeoTargetingTypeEnum_GeoTargetingType `protobuf:"varint,3,opt,name=location_type,json=locationType,proto3,enum=google.ads.googleads.v1.enums.GeoTargetingTypeEnum_GeoTargetingType" json:"location_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                    `json:"-"`
	XXX_unrecognized     []byte                                      `json:"-"`
	XXX_sizecache        int32                                       `json:"-"`
}

func (m *GeographicView) Reset()         { *m = GeographicView{} }
func (m *GeographicView) String() string { return proto.CompactTextString(m) }
func (*GeographicView) ProtoMessage()    {}
func (*GeographicView) Descriptor() ([]byte, []int) {
	return fileDescriptor_45a03e49a5e672c5, []int{0}
}

func (m *GeographicView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GeographicView.Unmarshal(m, b)
}
func (m *GeographicView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GeographicView.Marshal(b, m, deterministic)
}
func (m *GeographicView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GeographicView.Merge(m, src)
}
func (m *GeographicView) XXX_Size() int {
	return xxx_messageInfo_GeographicView.Size(m)
}
func (m *GeographicView) XXX_DiscardUnknown() {
	xxx_messageInfo_GeographicView.DiscardUnknown(m)
}

var xxx_messageInfo_GeographicView proto.InternalMessageInfo

func (m *GeographicView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *GeographicView) GetCountryGeoTargetConstant() *wrappers.StringValue {
	if m != nil {
		return m.CountryGeoTargetConstant
	}
	return nil
}

func (m *GeographicView) GetLocationType() enums.GeoTargetingTypeEnum_GeoTargetingType {
	if m != nil {
		return m.LocationType
	}
	return enums.GeoTargetingTypeEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*GeographicView)(nil), "google.ads.googleads.v1.resources.GeographicView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/geographic_view.proto", fileDescriptor_45a03e49a5e672c5)
}

var fileDescriptor_45a03e49a5e672c5 = []byte{
	// 498 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcf, 0x6a, 0xd4, 0x40,
	0x1c, 0x26, 0x1b, 0x14, 0x8c, 0x6d, 0x0f, 0xf1, 0xb2, 0xd6, 0xa2, 0x5b, 0xa1, 0xb8, 0x8a, 0xcc,
	0x90, 0x15, 0x2a, 0xc4, 0xd3, 0x44, 0x65, 0xc1, 0x83, 0x94, 0x75, 0xc9, 0x41, 0x16, 0xc2, 0x6c,
	0x32, 0x9d, 0x0e, 0x24, 0x33, 0x61, 0x66, 0xb2, 0xcb, 0x52, 0x0a, 0xbe, 0x81, 0xef, 0xe0, 0xd1,
	0x47, 0xf1, 0xe0, 0x33, 0xec, 0xb9, 0x8f, 0xd0, 0x93, 0x24, 0x99, 0xcc, 0x6e, 0x2d, 0xad, 0xde,
	0xbe, 0xcc, 0xf7, 0xfd, 0xbe, 0xdf, 0xdf, 0x78, 0x6f, 0xa9, 0x10, 0x34, 0x27, 0x10, 0x67, 0x0a,
	0xb6, 0xb0, 0x46, 0x8b, 0x00, 0x4a, 0xa2, 0x44, 0x25, 0x53, 0xa2, 0x20, 0x25, 0x82, 0x4a, 0x5c,
	0x9e, 0xb1, 0x34, 0x59, 0x30, 0xb2, 0x04, 0xa5, 0x14, 0x5a, 0xf8, 0x87, 0xad, 0x1a, 0xe0, 0x4c,
	0x01, 0x1b, 0x08, 0x16, 0x01, 0xb0, 0x81, 0xfb, 0xc7, 0xb7, 0x79, 0x13, 0x5e, 0x15, 0x8d, 0x6f,
	0xa2, 0xb1, 0xa4, 0x44, 0x33, 0x4e, 0x13, 0xbd, 0x2a, 0x49, 0x6b, 0xbd, 0xff, 0xac, 0x8b, 0x2b,
	0x19, 0x3c, 0x65, 0x24, 0xcf, 0x92, 0x39, 0x39, 0xc3, 0x0b, 0x26, 0xa4, 0x11, 0x3c, 0xde, 0x12,
	0x74, 0xe9, 0x0c, 0xf5, 0xd4, 0x50, 0xcd, 0xd7, 0xbc, 0x3a, 0x85, 0x4b, 0x89, 0xcb, 0x92, 0x48,
	0x65, 0xf8, 0x83, 0xad, 0x50, 0xcc, 0xb9, 0xd0, 0x58, 0x33, 0xc1, 0x0d, 0xfb, 0xfc, 0xb7, 0xeb,
	0xed, 0x8d, 0x6d, 0xbb, 0x31, 0x23, 0x4b, 0x7f, 0xea, 0xed, 0x76, 0x29, 0x12, 0x8e, 0x0b, 0xd2,
	0x77, 0x06, 0xce, 0xf0, 0x41, 0x04, 0xd7, 0xe8, 0xde, 0x15, 0x7a, 0xe9, 0xbd, 0xd8, 0xf4, 0x6e,
	0x50, 0xc9, 0x14, 0x48, 0x45, 0x01, 0xaf, 0xfb, 0x4c, 0x76, 0x3a, 0x97, 0xcf, 0xb8, 0x20, 0xfe,
	0x77, 0xc7, 0x7b, 0x92, 0x8a, 0x8a, 0x6b, 0xb9, 0x4a, 0x36, 0x73, 0x48, 0x52, 0xc1, 0x95, 0xc6,
	0x5c, 0xf7, 0x7b, 0x03, 0x67, 0xf8, 0x70, 0x74, 0x60, 0x3c, 0x41, 0xd7, 0x0d, 0xf8, 0xa2, 0x25,
	0xe3, 0x34, 0xc6, 0x79, 0x45, 0xa2, 0xd1, 0x1a, 0xb9, 0x57, 0xe8, 0xb5, 0xf7, 0xea, 0xae, 0x12,
	0xa6, 0x8d, 0xf1, 0x7b, 0xe3, 0x3b, 0xe9, 0x9b, 0xa4, 0x37, 0x18, 0x9f, 0x7b, 0xbb, 0xb9, 0x48,
	0x9b, 0x69, 0x34, 0xbb, 0xe8, 0xbb, 0x03, 0x67, 0xb8, 0x37, 0xfa, 0x00, 0x6e, 0xdb, 0x73, 0xb3,
	0x44, 0x60, 0x8d, 0x18, 0xa7, 0xd3, 0x55, 0x49, 0x3e, 0xf2, 0xaa, 0xb8, 0xf1, 0x18, 0xb9, 0x6b,
	0xe4, 0x4e, 0x76, 0x3a, 0xff, 0xfa, 0x29, 0xcc, 0x2e, 0x11, 0xfe, 0xef, 0xe9, 0xf9, 0xc7, 0x69,
	0xa5, 0xb4, 0x28, 0x88, 0x54, 0xf0, 0xbc, 0x83, 0x17, 0x5b, 0x97, 0x59, 0x8b, 0x14, 0x3c, 0xff,
	0xeb, 0x54, 0x2f, 0xa2, 0x6f, 0x3d, 0xef, 0x28, 0x15, 0x05, 0xf8, 0xe7, 0xb1, 0x46, 0x8f, 0xae,
	0x67, 0x3c, 0xa9, 0x27, 0x7e, 0xe2, 0x7c, 0xfd, 0x64, 0x22, 0xa9, 0xc8, 0x31, 0xa7, 0x40, 0x48,
	0x0a, 0x29, 0xe1, 0xcd, 0x3e, 0xe0, 0xa6, 0xe4, 0x3b, 0x7e, 0x9f, 0x77, 0x16, 0xfd, 0xe8, 0xb9,
	0x63, 0x84, 0x7e, 0xf6, 0x0e, 0xc7, 0xad, 0x25, 0xca, 0x14, 0x68, 0x61, 0x8d, 0xe2, 0x00, 0x4c,
	0x3a, 0xe5, 0xaf, 0x4e, 0x33, 0x43, 0x99, 0x9a, 0x59, 0xcd, 0x2c, 0x0e, 0x66, 0x56, 0x73, 0xd9,
	0x3b, 0x6a, 0x89, 0x30, 0x44, 0x99, 0x0a, 0x43, 0xab, 0x0a, 0xc3, 0x38, 0x08, 0x43, 0xab, 0x9b,
	0xdf, 0x6f, 0x8a, 0x7d, 0xf3, 0x27, 0x00, 0x00, 0xff, 0xff, 0x59, 0xc5, 0xae, 0xf0, 0xea, 0x03,
	0x00, 0x00,
}
