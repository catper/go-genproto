// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/detail_placement_view.proto

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

// A view with metrics aggregated by ad group and URL or YouTube video.
type DetailPlacementView struct {
	// Immutable. The resource name of the detail placement view.
	// Detail placement view resource names have the form:
	//
	// `customers/{customer_id}/detailPlacementViews/{ad_group_id}~{base64_placement}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The automatic placement string at detail level, e. g. website URL, mobile
	// application ID, or a YouTube video ID.
	Placement *wrappers.StringValue `protobuf:"bytes,2,opt,name=placement,proto3" json:"placement,omitempty"`
	// Output only. The display name is URL name for websites, YouTube video name for YouTube
	// videos, and translated mobile app name for mobile apps.
	DisplayName *wrappers.StringValue `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Output only. URL of the group placement, e.g. domain, link to the mobile application in
	// app store, or a YouTube channel URL.
	GroupPlacementTargetUrl *wrappers.StringValue `protobuf:"bytes,4,opt,name=group_placement_target_url,json=groupPlacementTargetUrl,proto3" json:"group_placement_target_url,omitempty"`
	// Output only. URL of the placement, e.g. website, link to the mobile application in app
	// store, or a YouTube video URL.
	TargetUrl *wrappers.StringValue `protobuf:"bytes,5,opt,name=target_url,json=targetUrl,proto3" json:"target_url,omitempty"`
	// Output only. Type of the placement, e.g. Website, YouTube Video, and Mobile Application.
	PlacementType        enums.PlacementTypeEnum_PlacementType `protobuf:"varint,6,opt,name=placement_type,json=placementType,proto3,enum=google.ads.googleads.v2.enums.PlacementTypeEnum_PlacementType" json:"placement_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *DetailPlacementView) Reset()         { *m = DetailPlacementView{} }
func (m *DetailPlacementView) String() string { return proto.CompactTextString(m) }
func (*DetailPlacementView) ProtoMessage()    {}
func (*DetailPlacementView) Descriptor() ([]byte, []int) {
	return fileDescriptor_64070670c65266a2, []int{0}
}

func (m *DetailPlacementView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetailPlacementView.Unmarshal(m, b)
}
func (m *DetailPlacementView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetailPlacementView.Marshal(b, m, deterministic)
}
func (m *DetailPlacementView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetailPlacementView.Merge(m, src)
}
func (m *DetailPlacementView) XXX_Size() int {
	return xxx_messageInfo_DetailPlacementView.Size(m)
}
func (m *DetailPlacementView) XXX_DiscardUnknown() {
	xxx_messageInfo_DetailPlacementView.DiscardUnknown(m)
}

var xxx_messageInfo_DetailPlacementView proto.InternalMessageInfo

func (m *DetailPlacementView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *DetailPlacementView) GetPlacement() *wrappers.StringValue {
	if m != nil {
		return m.Placement
	}
	return nil
}

func (m *DetailPlacementView) GetDisplayName() *wrappers.StringValue {
	if m != nil {
		return m.DisplayName
	}
	return nil
}

func (m *DetailPlacementView) GetGroupPlacementTargetUrl() *wrappers.StringValue {
	if m != nil {
		return m.GroupPlacementTargetUrl
	}
	return nil
}

func (m *DetailPlacementView) GetTargetUrl() *wrappers.StringValue {
	if m != nil {
		return m.TargetUrl
	}
	return nil
}

func (m *DetailPlacementView) GetPlacementType() enums.PlacementTypeEnum_PlacementType {
	if m != nil {
		return m.PlacementType
	}
	return enums.PlacementTypeEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*DetailPlacementView)(nil), "google.ads.googleads.v2.resources.DetailPlacementView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/detail_placement_view.proto", fileDescriptor_64070670c65266a2)
}

var fileDescriptor_64070670c65266a2 = []byte{
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcb, 0x6e, 0x13, 0x31,
	0x14, 0x55, 0x92, 0xb6, 0x52, 0xdd, 0xc7, 0x62, 0x58, 0x10, 0xaa, 0x0a, 0x52, 0xa4, 0x4a, 0x59,
	0x20, 0x5b, 0x1a, 0x58, 0x0d, 0xe2, 0xe1, 0xa8, 0xa8, 0x12, 0x0b, 0x14, 0x05, 0x88, 0x04, 0x8a,
	0x88, 0x9c, 0xcc, 0xed, 0x60, 0x34, 0x63, 0x1b, 0xdb, 0x93, 0x28, 0xaa, 0xba, 0xe4, 0x47, 0x58,
	0xf2, 0x29, 0x7c, 0x45, 0xd7, 0xfd, 0x04, 0xd8, 0xa0, 0xcc, 0x78, 0x9c, 0x44, 0x4a, 0x21, 0xec,
	0xce, 0xcc, 0x3d, 0xe7, 0xdc, 0x7b, 0x6c, 0x5f, 0xf4, 0x2c, 0x91, 0x32, 0x49, 0x81, 0xb0, 0xd8,
	0x90, 0x12, 0xce, 0xd1, 0x24, 0x24, 0x1a, 0x8c, 0xcc, 0xf5, 0x18, 0x0c, 0x89, 0xc1, 0x32, 0x9e,
	0x0e, 0x55, 0xca, 0xc6, 0x90, 0x81, 0xb0, 0xc3, 0x09, 0x87, 0x29, 0x56, 0x5a, 0x5a, 0x19, 0x9c,
	0x94, 0x1a, 0xcc, 0x62, 0x83, 0xbd, 0x1c, 0x4f, 0x42, 0xec, 0xe5, 0x47, 0xe1, 0x6d, 0x1d, 0x40,
	0xe4, 0x99, 0x21, 0x0b, 0x5b, 0x3b, 0x53, 0x50, 0xda, 0x1e, 0x3d, 0xa8, 0x34, 0x8a, 0x93, 0x0b,
	0x0e, 0x69, 0x3c, 0x1c, 0xc1, 0x67, 0x36, 0xe1, 0x52, 0x3b, 0xc2, 0xbd, 0x25, 0x42, 0xd5, 0xca,
	0x95, 0xee, 0xbb, 0x52, 0xf1, 0x35, 0xca, 0x2f, 0xc8, 0x54, 0x33, 0xa5, 0x40, 0x1b, 0x57, 0x3f,
	0x5e, 0x92, 0x32, 0x21, 0xa4, 0x65, 0x96, 0x4b, 0xe1, 0xaa, 0x0f, 0x7f, 0x6f, 0xa1, 0x3b, 0x67,
	0x45, 0xe0, 0x6e, 0x35, 0x58, 0x9f, 0xc3, 0x34, 0xf8, 0x80, 0x0e, 0xaa, 0x3e, 0x43, 0xc1, 0x32,
	0x68, 0xd6, 0x5a, 0xb5, 0xf6, 0x6e, 0xe7, 0xc9, 0x35, 0xdd, 0xfe, 0x45, 0x31, 0x7a, 0xb4, 0x08,
	0xef, 0x90, 0xe2, 0x06, 0x8f, 0x65, 0x46, 0xd6, 0x98, 0xf5, 0xf6, 0x2b, 0xab, 0x37, 0x2c, 0x83,
	0xe0, 0x05, 0xda, 0xf5, 0x87, 0xd0, 0xac, 0xb7, 0x6a, 0xed, 0xbd, 0xf0, 0xd8, 0xb9, 0xe0, 0x2a,
	0x04, 0x7e, 0x6b, 0x35, 0x17, 0x49, 0x9f, 0xa5, 0x39, 0x74, 0x1a, 0xd7, 0xb4, 0xd1, 0x5b, 0x68,
	0x82, 0x33, 0xb4, 0x1f, 0x73, 0xa3, 0x52, 0x36, 0x2b, 0x47, 0x6b, 0x6c, 0xea, 0xb1, 0xe7, 0x64,
	0xc5, 0x18, 0x9f, 0xd0, 0x51, 0xa2, 0x65, 0xae, 0x96, 0x2e, 0xda, 0x32, 0x9d, 0x80, 0x1d, 0xe6,
	0x3a, 0x6d, 0x6e, 0x6d, 0xea, 0x79, 0xb7, 0x30, 0xf1, 0x79, 0xdf, 0x15, 0x16, 0xef, 0x75, 0x1a,
	0xbc, 0x44, 0x68, 0xc9, 0x6f, 0x7b, 0xe3, 0x9c, 0xd6, 0x3b, 0x7c, 0x41, 0x87, 0xab, 0xaf, 0xa5,
	0xb9, 0xd3, 0xaa, 0xb5, 0x0f, 0xc3, 0xe7, 0xf8, 0xb6, 0x57, 0x58, 0x3c, 0x31, 0xbc, 0x18, 0x66,
	0xa6, 0xe0, 0x95, 0xc8, 0xb3, 0xd5, 0x3f, 0x65, 0x9f, 0x03, 0xb5, 0xfc, 0x2f, 0xb2, 0x37, 0xf4,
	0xeb, 0xff, 0xdd, 0x6a, 0x40, 0xc7, 0xb9, 0xb1, 0x32, 0x03, 0x6d, 0xc8, 0x65, 0x05, 0xaf, 0xdc,
	0xf6, 0xac, 0x30, 0x0d, 0xb9, 0x5c, 0xbb, 0x53, 0x57, 0x9d, 0x6f, 0x75, 0x74, 0x3a, 0x96, 0x19,
	0xfe, 0xe7, 0x56, 0x75, 0x9a, 0x6b, 0x26, 0xe8, 0xce, 0x0f, 0xb1, 0x5b, 0xfb, 0xf8, 0xda, 0xc9,
	0x13, 0x99, 0x32, 0x91, 0x60, 0xa9, 0x13, 0x92, 0x80, 0x28, 0x8e, 0x98, 0x2c, 0x72, 0xfc, 0x65,
	0xe5, 0x9f, 0x7a, 0xf4, 0xbd, 0xde, 0x38, 0xa7, 0xf4, 0x47, 0xfd, 0xe4, 0xbc, 0xb4, 0xa4, 0xb1,
	0xc1, 0x25, 0x9c, 0xa3, 0x7e, 0x88, 0x7b, 0x15, 0xf3, 0x67, 0xc5, 0x19, 0xd0, 0xd8, 0x0c, 0x3c,
	0x67, 0xd0, 0x0f, 0x07, 0x9e, 0x73, 0x53, 0x3f, 0x2d, 0x0b, 0x51, 0x44, 0x63, 0x13, 0x45, 0x9e,
	0x15, 0x45, 0xfd, 0x30, 0x8a, 0x3c, 0x6f, 0xb4, 0x53, 0x0c, 0xfb, 0xf8, 0x4f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x66, 0x08, 0xb7, 0xd4, 0x9e, 0x04, 0x00, 0x00,
}
