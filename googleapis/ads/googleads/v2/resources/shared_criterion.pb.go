// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/shared_criterion.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
	wrappers "github.com/catper/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v2/common"
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

// A criterion belonging to a shared set.
type SharedCriterion struct {
	// Immutable. The resource name of the shared criterion.
	// Shared set resource names have the form:
	//
	// `customers/{customer_id}/sharedCriteria/{shared_set_id}~{criterion_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Immutable. The shared set to which the shared criterion belongs.
	SharedSet *wrappers.StringValue `protobuf:"bytes,2,opt,name=shared_set,json=sharedSet,proto3" json:"shared_set,omitempty"`
	// Output only. The ID of the criterion.
	//
	// This field is ignored for mutates.
	CriterionId *wrappers.Int64Value `protobuf:"bytes,26,opt,name=criterion_id,json=criterionId,proto3" json:"criterion_id,omitempty"`
	// Output only. The type of the criterion.
	Type enums.CriterionTypeEnum_CriterionType `protobuf:"varint,4,opt,name=type,proto3,enum=google.ads.googleads.v2.enums.CriterionTypeEnum_CriterionType" json:"type,omitempty"`
	// The criterion.
	//
	// Exactly one must be set.
	//
	// Types that are valid to be assigned to Criterion:
	//	*SharedCriterion_Keyword
	//	*SharedCriterion_YoutubeVideo
	//	*SharedCriterion_YoutubeChannel
	//	*SharedCriterion_Placement
	//	*SharedCriterion_MobileAppCategory
	//	*SharedCriterion_MobileApplication
	Criterion            isSharedCriterion_Criterion `protobuf_oneof:"criterion"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *SharedCriterion) Reset()         { *m = SharedCriterion{} }
func (m *SharedCriterion) String() string { return proto.CompactTextString(m) }
func (*SharedCriterion) ProtoMessage()    {}
func (*SharedCriterion) Descriptor() ([]byte, []int) {
	return fileDescriptor_1041cb67d99aea67, []int{0}
}

func (m *SharedCriterion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SharedCriterion.Unmarshal(m, b)
}
func (m *SharedCriterion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SharedCriterion.Marshal(b, m, deterministic)
}
func (m *SharedCriterion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SharedCriterion.Merge(m, src)
}
func (m *SharedCriterion) XXX_Size() int {
	return xxx_messageInfo_SharedCriterion.Size(m)
}
func (m *SharedCriterion) XXX_DiscardUnknown() {
	xxx_messageInfo_SharedCriterion.DiscardUnknown(m)
}

var xxx_messageInfo_SharedCriterion proto.InternalMessageInfo

func (m *SharedCriterion) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *SharedCriterion) GetSharedSet() *wrappers.StringValue {
	if m != nil {
		return m.SharedSet
	}
	return nil
}

func (m *SharedCriterion) GetCriterionId() *wrappers.Int64Value {
	if m != nil {
		return m.CriterionId
	}
	return nil
}

func (m *SharedCriterion) GetType() enums.CriterionTypeEnum_CriterionType {
	if m != nil {
		return m.Type
	}
	return enums.CriterionTypeEnum_UNSPECIFIED
}

type isSharedCriterion_Criterion interface {
	isSharedCriterion_Criterion()
}

type SharedCriterion_Keyword struct {
	Keyword *common.KeywordInfo `protobuf:"bytes,3,opt,name=keyword,proto3,oneof"`
}

type SharedCriterion_YoutubeVideo struct {
	YoutubeVideo *common.YouTubeVideoInfo `protobuf:"bytes,5,opt,name=youtube_video,json=youtubeVideo,proto3,oneof"`
}

type SharedCriterion_YoutubeChannel struct {
	YoutubeChannel *common.YouTubeChannelInfo `protobuf:"bytes,6,opt,name=youtube_channel,json=youtubeChannel,proto3,oneof"`
}

type SharedCriterion_Placement struct {
	Placement *common.PlacementInfo `protobuf:"bytes,7,opt,name=placement,proto3,oneof"`
}

type SharedCriterion_MobileAppCategory struct {
	MobileAppCategory *common.MobileAppCategoryInfo `protobuf:"bytes,8,opt,name=mobile_app_category,json=mobileAppCategory,proto3,oneof"`
}

type SharedCriterion_MobileApplication struct {
	MobileApplication *common.MobileApplicationInfo `protobuf:"bytes,9,opt,name=mobile_application,json=mobileApplication,proto3,oneof"`
}

func (*SharedCriterion_Keyword) isSharedCriterion_Criterion() {}

func (*SharedCriterion_YoutubeVideo) isSharedCriterion_Criterion() {}

func (*SharedCriterion_YoutubeChannel) isSharedCriterion_Criterion() {}

func (*SharedCriterion_Placement) isSharedCriterion_Criterion() {}

func (*SharedCriterion_MobileAppCategory) isSharedCriterion_Criterion() {}

func (*SharedCriterion_MobileApplication) isSharedCriterion_Criterion() {}

func (m *SharedCriterion) GetCriterion() isSharedCriterion_Criterion {
	if m != nil {
		return m.Criterion
	}
	return nil
}

func (m *SharedCriterion) GetKeyword() *common.KeywordInfo {
	if x, ok := m.GetCriterion().(*SharedCriterion_Keyword); ok {
		return x.Keyword
	}
	return nil
}

func (m *SharedCriterion) GetYoutubeVideo() *common.YouTubeVideoInfo {
	if x, ok := m.GetCriterion().(*SharedCriterion_YoutubeVideo); ok {
		return x.YoutubeVideo
	}
	return nil
}

func (m *SharedCriterion) GetYoutubeChannel() *common.YouTubeChannelInfo {
	if x, ok := m.GetCriterion().(*SharedCriterion_YoutubeChannel); ok {
		return x.YoutubeChannel
	}
	return nil
}

func (m *SharedCriterion) GetPlacement() *common.PlacementInfo {
	if x, ok := m.GetCriterion().(*SharedCriterion_Placement); ok {
		return x.Placement
	}
	return nil
}

func (m *SharedCriterion) GetMobileAppCategory() *common.MobileAppCategoryInfo {
	if x, ok := m.GetCriterion().(*SharedCriterion_MobileAppCategory); ok {
		return x.MobileAppCategory
	}
	return nil
}

func (m *SharedCriterion) GetMobileApplication() *common.MobileApplicationInfo {
	if x, ok := m.GetCriterion().(*SharedCriterion_MobileApplication); ok {
		return x.MobileApplication
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SharedCriterion) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SharedCriterion_Keyword)(nil),
		(*SharedCriterion_YoutubeVideo)(nil),
		(*SharedCriterion_YoutubeChannel)(nil),
		(*SharedCriterion_Placement)(nil),
		(*SharedCriterion_MobileAppCategory)(nil),
		(*SharedCriterion_MobileApplication)(nil),
	}
}

func init() {
	proto.RegisterType((*SharedCriterion)(nil), "google.ads.googleads.v2.resources.SharedCriterion")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/shared_criterion.proto", fileDescriptor_1041cb67d99aea67)
}

var fileDescriptor_1041cb67d99aea67 = []byte{
	// 687 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xdb, 0x6a, 0xdb, 0x4a,
	0x14, 0x86, 0xb7, 0xed, 0x9c, 0x3c, 0x39, 0xb1, 0xb5, 0xf7, 0x85, 0x76, 0x76, 0x68, 0x93, 0xd0,
	0x80, 0x49, 0xc9, 0x28, 0xa8, 0x6d, 0x28, 0x2a, 0x14, 0x64, 0x53, 0xd2, 0xf4, 0x44, 0x70, 0x52,
	0x97, 0x16, 0x83, 0x18, 0x49, 0x2b, 0x8a, 0x88, 0x34, 0x23, 0x34, 0x92, 0x83, 0x09, 0xb9, 0xe8,
	0xab, 0xe4, 0xb2, 0x8f, 0xd2, 0xa7, 0xc8, 0x75, 0x1e, 0xa1, 0x57, 0xc5, 0xa3, 0x19, 0xd9, 0x75,
	0xeb, 0x3a, 0xbd, 0x5b, 0xd2, 0xfa, 0xff, 0xef, 0x5f, 0x9a, 0xd1, 0x0c, 0x7a, 0x1a, 0x30, 0x16,
	0x44, 0x60, 0x10, 0x9f, 0x1b, 0x45, 0x39, 0xa8, 0x7a, 0xa6, 0x91, 0x02, 0x67, 0x79, 0xea, 0x01,
	0x37, 0xf8, 0x19, 0x49, 0xc1, 0x77, 0xbc, 0x34, 0xcc, 0x20, 0x0d, 0x19, 0xc5, 0x49, 0xca, 0x32,
	0xa6, 0x6d, 0x16, 0x72, 0x4c, 0x7c, 0x8e, 0x4b, 0x27, 0xee, 0x99, 0xb8, 0x74, 0xae, 0xed, 0x4e,
	0x82, 0x7b, 0x2c, 0x8e, 0x19, 0x35, 0x24, 0x92, 0x14, 0xc4, 0x35, 0x73, 0x92, 0x1c, 0x68, 0x1e,
	0x73, 0xa3, 0x1c, 0xc0, 0xc9, 0xfa, 0x09, 0x48, 0xcf, 0x7d, 0xe5, 0x49, 0x42, 0xe3, 0x34, 0x84,
	0xc8, 0x77, 0x5c, 0x38, 0x23, 0xbd, 0x90, 0xa5, 0x52, 0xf0, 0xdf, 0x88, 0x40, 0x4d, 0x26, 0x5b,
	0xf7, 0x64, 0x4b, 0x3c, 0xb9, 0xf9, 0xa9, 0x71, 0x91, 0x92, 0x24, 0x81, 0x94, 0xcb, 0xfe, 0xfa,
	0x88, 0x95, 0x50, 0xca, 0x32, 0x92, 0x85, 0x8c, 0xca, 0xee, 0xd6, 0xf5, 0x02, 0x5a, 0x3d, 0x16,
	0x4b, 0xd3, 0x52, 0x83, 0x69, 0xef, 0xd1, 0xb2, 0xca, 0x70, 0x28, 0x89, 0x41, 0xaf, 0x6c, 0x54,
	0x1a, 0xf5, 0xe6, 0xde, 0x8d, 0x3d, 0xfb, 0xcd, 0xde, 0x41, 0x8d, 0xe1, 0x3a, 0xc9, 0x2a, 0x09,
	0x39, 0xf6, 0x58, 0x6c, 0x8c, 0x81, 0xda, 0x4b, 0x0a, 0xf3, 0x8e, 0xc4, 0xa0, 0x05, 0x08, 0xc9,
	0x4d, 0xe0, 0x90, 0xe9, 0xd5, 0x8d, 0x4a, 0x63, 0xd1, 0x5c, 0x97, 0x08, 0xac, 0xa6, 0xc7, 0xc7,
	0x59, 0x1a, 0xd2, 0xa0, 0x43, 0xa2, 0x1c, 0x9a, 0x3b, 0x22, 0xf1, 0x01, 0xda, 0x9a, 0x92, 0x78,
	0x0c, 0x59, 0xbb, 0xce, 0x55, 0xa9, 0xb5, 0xd0, 0xd2, 0x70, 0x95, 0x43, 0x5f, 0x5f, 0x13, 0x51,
	0xff, 0xff, 0x14, 0x75, 0x48, 0xb3, 0xfd, 0xc7, 0x45, 0x52, 0xed, 0xc6, 0xae, 0xb5, 0x17, 0x4b,
	0xd7, 0xa1, 0xaf, 0x7d, 0x40, 0x33, 0x83, 0x0d, 0xd2, 0x67, 0x36, 0x2a, 0x8d, 0x15, 0xf3, 0x39,
	0x9e, 0xf4, 0x9f, 0x88, 0x5d, 0xc5, 0xe5, 0x37, 0x9f, 0xf4, 0x13, 0x78, 0x41, 0xf3, 0xf8, 0xc7,
	0x37, 0x05, 0x5f, 0x00, 0xb5, 0x37, 0x68, 0xfe, 0x1c, 0xfa, 0x17, 0x2c, 0xf5, 0xf5, 0x9a, 0x18,
	0xec, 0xe1, 0x44, 0x76, 0xf1, 0x83, 0xe1, 0xd7, 0x85, 0xfc, 0x90, 0x9e, 0xb2, 0x01, 0x68, 0xf6,
	0xe5, 0x5f, 0x6d, 0x85, 0xd0, 0xba, 0x68, 0xb9, 0xcf, 0xf2, 0x2c, 0x77, 0xc1, 0xe9, 0x85, 0x3e,
	0x30, 0x7d, 0x56, 0x30, 0xf7, 0xa6, 0x31, 0x3f, 0xb2, 0xfc, 0x24, 0x77, 0xa1, 0x33, 0xf0, 0x8c,
	0x82, 0x97, 0x24, 0x4d, 0x34, 0x34, 0x17, 0xad, 0x2a, 0xba, 0x77, 0x46, 0x28, 0x85, 0x48, 0x9f,
	0x13, 0x7c, 0xf3, 0x8e, 0xfc, 0x56, 0xe1, 0x1a, 0x4d, 0x58, 0x91, 0x44, 0xd9, 0xd2, 0xda, 0xa8,
	0x9e, 0x44, 0xc4, 0x83, 0x18, 0x68, 0xa6, 0xcf, 0x0b, 0xfa, 0xee, 0x34, 0xfa, 0x91, 0x32, 0x8c,
	0x82, 0x87, 0x18, 0x2d, 0x42, 0xff, 0xc4, 0xcc, 0x0d, 0x23, 0x70, 0x48, 0x92, 0x38, 0x1e, 0xc9,
	0x20, 0x60, 0x69, 0x5f, 0x5f, 0x10, 0xf4, 0x27, 0xd3, 0xe8, 0x6f, 0x85, 0xd5, 0x4e, 0x92, 0x96,
	0x34, 0x8e, 0xa6, 0xfc, 0x1d, 0x8f, 0x77, 0xb5, 0x73, 0xa4, 0x0d, 0xd3, 0xa2, 0xd0, 0x13, 0x07,
	0x4c, 0xaf, 0xff, 0x61, 0x98, 0x32, 0xfe, 0x3a, 0x4c, 0x75, 0x2d, 0xb8, 0xb5, 0xdd, 0xbb, 0x1f,
	0x41, 0x6d, 0xdf, 0xcb, 0x79, 0xc6, 0x62, 0x48, 0xb9, 0x71, 0xa9, 0xca, 0x2b, 0x79, 0x19, 0x4a,
	0x15, 0x31, 0x2e, 0xc7, 0x2f, 0xc7, 0xab, 0xe6, 0x22, 0xaa, 0x97, 0x4f, 0xcd, 0xcf, 0x55, 0xb4,
	0xed, 0xb1, 0x18, 0x4f, 0xbd, 0x2b, 0x9b, 0xff, 0x8e, 0xe5, 0x1f, 0x0d, 0xce, 0xda, 0x51, 0xe5,
	0xd3, 0x2b, 0x69, 0x0d, 0x58, 0x44, 0x68, 0x80, 0x59, 0x1a, 0x18, 0x01, 0x50, 0x71, 0x12, 0x8d,
	0xe1, 0x17, 0xfc, 0xe6, 0xfe, 0x7e, 0x56, 0x56, 0xd7, 0xd5, 0xda, 0x81, 0x6d, 0x7f, 0xa9, 0x6e,
	0x1e, 0x14, 0x48, 0xdb, 0xe7, 0xb8, 0x28, 0x07, 0x55, 0xc7, 0xc4, 0x6d, 0xa5, 0xfc, 0xaa, 0x34,
	0x5d, 0xdb, 0xe7, 0xdd, 0x52, 0xd3, 0xed, 0x98, 0xdd, 0x52, 0x73, 0x5b, 0xdd, 0x2e, 0x1a, 0x96,
	0x65, 0xfb, 0xdc, 0xb2, 0x4a, 0x95, 0x65, 0x75, 0x4c, 0xcb, 0x2a, 0x75, 0xee, 0x9c, 0x18, 0xf6,
	0xd1, 0xf7, 0x00, 0x00, 0x00, 0xff, 0xff, 0xac, 0xb6, 0x7d, 0xcc, 0x6b, 0x06, 0x00, 0x00,
}
