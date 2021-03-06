// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/extension_feed_item.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
	wrappers "github.com/catper/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v1/common"
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

// An extension feed item.
type ExtensionFeedItem struct {
	// Immutable. The resource name of the extension feed item.
	// Extension feed item resource names have the form:
	//
	// `customers/{customer_id}/extensionFeedItems/{feed_item_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The extension type of the extension feed item.
	// This field is read-only.
	ExtensionType enums.ExtensionTypeEnum_ExtensionType `protobuf:"varint,13,opt,name=extension_type,json=extensionType,proto3,enum=google.ads.googleads.v1.enums.ExtensionTypeEnum_ExtensionType" json:"extension_type,omitempty"`
	// Start time in which this feed item is effective and can begin serving. The
	// time is in the customer's time zone.
	// The format is "YYYY-MM-DD HH:MM:SS".
	// Examples: "2018-03-05 09:15:00" or "2018-02-01 14:34:30"
	StartDateTime *wrappers.StringValue `protobuf:"bytes,5,opt,name=start_date_time,json=startDateTime,proto3" json:"start_date_time,omitempty"`
	// End time in which this feed item is no longer effective and will stop
	// serving. The time is in the customer's time zone.
	// The format is "YYYY-MM-DD HH:MM:SS".
	// Examples: "2018-03-05 09:15:00" or "2018-02-01 14:34:30"
	EndDateTime *wrappers.StringValue `protobuf:"bytes,6,opt,name=end_date_time,json=endDateTime,proto3" json:"end_date_time,omitempty"`
	// List of non-overlapping schedules specifying all time intervals
	// for which the feed item may serve. There can be a maximum of 6 schedules
	// per day.
	AdSchedules []*common.AdScheduleInfo `protobuf:"bytes,16,rep,name=ad_schedules,json=adSchedules,proto3" json:"ad_schedules,omitempty"`
	// The targeted device.
	Device enums.FeedItemTargetDeviceEnum_FeedItemTargetDevice `protobuf:"varint,17,opt,name=device,proto3,enum=google.ads.googleads.v1.enums.FeedItemTargetDeviceEnum_FeedItemTargetDevice" json:"device,omitempty"`
	// The targeted geo target constant.
	TargetedGeoTargetConstant *wrappers.StringValue `protobuf:"bytes,20,opt,name=targeted_geo_target_constant,json=targetedGeoTargetConstant,proto3" json:"targeted_geo_target_constant,omitempty"`
	// Output only. Status of the feed item.
	// This field is read-only.
	Status enums.FeedItemStatusEnum_FeedItemStatus `protobuf:"varint,4,opt,name=status,proto3,enum=google.ads.googleads.v1.enums.FeedItemStatusEnum_FeedItemStatus" json:"status,omitempty"`
	// Extension type.
	//
	// Types that are valid to be assigned to Extension:
	//	*ExtensionFeedItem_SitelinkFeedItem
	//	*ExtensionFeedItem_StructuredSnippetFeedItem
	//	*ExtensionFeedItem_AppFeedItem
	//	*ExtensionFeedItem_CallFeedItem
	//	*ExtensionFeedItem_CalloutFeedItem
	//	*ExtensionFeedItem_TextMessageFeedItem
	//	*ExtensionFeedItem_PriceFeedItem
	//	*ExtensionFeedItem_PromotionFeedItem
	//	*ExtensionFeedItem_LocationFeedItem
	//	*ExtensionFeedItem_AffiliateLocationFeedItem
	Extension isExtensionFeedItem_Extension `protobuf_oneof:"extension"`
	// Targeting at either the campaign or ad group level. Feed items that target
	// a campaign or ad group will only serve with that resource.
	//
	// Types that are valid to be assigned to ServingResourceTargeting:
	//	*ExtensionFeedItem_TargetedCampaign
	//	*ExtensionFeedItem_TargetedAdGroup
	ServingResourceTargeting isExtensionFeedItem_ServingResourceTargeting `protobuf_oneof:"serving_resource_targeting"`
	XXX_NoUnkeyedLiteral     struct{}                                     `json:"-"`
	XXX_unrecognized         []byte                                       `json:"-"`
	XXX_sizecache            int32                                        `json:"-"`
}

func (m *ExtensionFeedItem) Reset()         { *m = ExtensionFeedItem{} }
func (m *ExtensionFeedItem) String() string { return proto.CompactTextString(m) }
func (*ExtensionFeedItem) ProtoMessage()    {}
func (*ExtensionFeedItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_800f8a839b1b4536, []int{0}
}

func (m *ExtensionFeedItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtensionFeedItem.Unmarshal(m, b)
}
func (m *ExtensionFeedItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtensionFeedItem.Marshal(b, m, deterministic)
}
func (m *ExtensionFeedItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtensionFeedItem.Merge(m, src)
}
func (m *ExtensionFeedItem) XXX_Size() int {
	return xxx_messageInfo_ExtensionFeedItem.Size(m)
}
func (m *ExtensionFeedItem) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtensionFeedItem.DiscardUnknown(m)
}

var xxx_messageInfo_ExtensionFeedItem proto.InternalMessageInfo

func (m *ExtensionFeedItem) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *ExtensionFeedItem) GetExtensionType() enums.ExtensionTypeEnum_ExtensionType {
	if m != nil {
		return m.ExtensionType
	}
	return enums.ExtensionTypeEnum_UNSPECIFIED
}

func (m *ExtensionFeedItem) GetStartDateTime() *wrappers.StringValue {
	if m != nil {
		return m.StartDateTime
	}
	return nil
}

func (m *ExtensionFeedItem) GetEndDateTime() *wrappers.StringValue {
	if m != nil {
		return m.EndDateTime
	}
	return nil
}

func (m *ExtensionFeedItem) GetAdSchedules() []*common.AdScheduleInfo {
	if m != nil {
		return m.AdSchedules
	}
	return nil
}

func (m *ExtensionFeedItem) GetDevice() enums.FeedItemTargetDeviceEnum_FeedItemTargetDevice {
	if m != nil {
		return m.Device
	}
	return enums.FeedItemTargetDeviceEnum_UNSPECIFIED
}

func (m *ExtensionFeedItem) GetTargetedGeoTargetConstant() *wrappers.StringValue {
	if m != nil {
		return m.TargetedGeoTargetConstant
	}
	return nil
}

func (m *ExtensionFeedItem) GetStatus() enums.FeedItemStatusEnum_FeedItemStatus {
	if m != nil {
		return m.Status
	}
	return enums.FeedItemStatusEnum_UNSPECIFIED
}

type isExtensionFeedItem_Extension interface {
	isExtensionFeedItem_Extension()
}

type ExtensionFeedItem_SitelinkFeedItem struct {
	SitelinkFeedItem *common.SitelinkFeedItem `protobuf:"bytes,2,opt,name=sitelink_feed_item,json=sitelinkFeedItem,proto3,oneof"`
}

type ExtensionFeedItem_StructuredSnippetFeedItem struct {
	StructuredSnippetFeedItem *common.StructuredSnippetFeedItem `protobuf:"bytes,3,opt,name=structured_snippet_feed_item,json=structuredSnippetFeedItem,proto3,oneof"`
}

type ExtensionFeedItem_AppFeedItem struct {
	AppFeedItem *common.AppFeedItem `protobuf:"bytes,7,opt,name=app_feed_item,json=appFeedItem,proto3,oneof"`
}

type ExtensionFeedItem_CallFeedItem struct {
	CallFeedItem *common.CallFeedItem `protobuf:"bytes,8,opt,name=call_feed_item,json=callFeedItem,proto3,oneof"`
}

type ExtensionFeedItem_CalloutFeedItem struct {
	CalloutFeedItem *common.CalloutFeedItem `protobuf:"bytes,9,opt,name=callout_feed_item,json=calloutFeedItem,proto3,oneof"`
}

type ExtensionFeedItem_TextMessageFeedItem struct {
	TextMessageFeedItem *common.TextMessageFeedItem `protobuf:"bytes,10,opt,name=text_message_feed_item,json=textMessageFeedItem,proto3,oneof"`
}

type ExtensionFeedItem_PriceFeedItem struct {
	PriceFeedItem *common.PriceFeedItem `protobuf:"bytes,11,opt,name=price_feed_item,json=priceFeedItem,proto3,oneof"`
}

type ExtensionFeedItem_PromotionFeedItem struct {
	PromotionFeedItem *common.PromotionFeedItem `protobuf:"bytes,12,opt,name=promotion_feed_item,json=promotionFeedItem,proto3,oneof"`
}

type ExtensionFeedItem_LocationFeedItem struct {
	LocationFeedItem *common.LocationFeedItem `protobuf:"bytes,14,opt,name=location_feed_item,json=locationFeedItem,proto3,oneof"`
}

type ExtensionFeedItem_AffiliateLocationFeedItem struct {
	AffiliateLocationFeedItem *common.AffiliateLocationFeedItem `protobuf:"bytes,15,opt,name=affiliate_location_feed_item,json=affiliateLocationFeedItem,proto3,oneof"`
}

func (*ExtensionFeedItem_SitelinkFeedItem) isExtensionFeedItem_Extension() {}

func (*ExtensionFeedItem_StructuredSnippetFeedItem) isExtensionFeedItem_Extension() {}

func (*ExtensionFeedItem_AppFeedItem) isExtensionFeedItem_Extension() {}

func (*ExtensionFeedItem_CallFeedItem) isExtensionFeedItem_Extension() {}

func (*ExtensionFeedItem_CalloutFeedItem) isExtensionFeedItem_Extension() {}

func (*ExtensionFeedItem_TextMessageFeedItem) isExtensionFeedItem_Extension() {}

func (*ExtensionFeedItem_PriceFeedItem) isExtensionFeedItem_Extension() {}

func (*ExtensionFeedItem_PromotionFeedItem) isExtensionFeedItem_Extension() {}

func (*ExtensionFeedItem_LocationFeedItem) isExtensionFeedItem_Extension() {}

func (*ExtensionFeedItem_AffiliateLocationFeedItem) isExtensionFeedItem_Extension() {}

func (m *ExtensionFeedItem) GetExtension() isExtensionFeedItem_Extension {
	if m != nil {
		return m.Extension
	}
	return nil
}

func (m *ExtensionFeedItem) GetSitelinkFeedItem() *common.SitelinkFeedItem {
	if x, ok := m.GetExtension().(*ExtensionFeedItem_SitelinkFeedItem); ok {
		return x.SitelinkFeedItem
	}
	return nil
}

func (m *ExtensionFeedItem) GetStructuredSnippetFeedItem() *common.StructuredSnippetFeedItem {
	if x, ok := m.GetExtension().(*ExtensionFeedItem_StructuredSnippetFeedItem); ok {
		return x.StructuredSnippetFeedItem
	}
	return nil
}

func (m *ExtensionFeedItem) GetAppFeedItem() *common.AppFeedItem {
	if x, ok := m.GetExtension().(*ExtensionFeedItem_AppFeedItem); ok {
		return x.AppFeedItem
	}
	return nil
}

func (m *ExtensionFeedItem) GetCallFeedItem() *common.CallFeedItem {
	if x, ok := m.GetExtension().(*ExtensionFeedItem_CallFeedItem); ok {
		return x.CallFeedItem
	}
	return nil
}

func (m *ExtensionFeedItem) GetCalloutFeedItem() *common.CalloutFeedItem {
	if x, ok := m.GetExtension().(*ExtensionFeedItem_CalloutFeedItem); ok {
		return x.CalloutFeedItem
	}
	return nil
}

func (m *ExtensionFeedItem) GetTextMessageFeedItem() *common.TextMessageFeedItem {
	if x, ok := m.GetExtension().(*ExtensionFeedItem_TextMessageFeedItem); ok {
		return x.TextMessageFeedItem
	}
	return nil
}

func (m *ExtensionFeedItem) GetPriceFeedItem() *common.PriceFeedItem {
	if x, ok := m.GetExtension().(*ExtensionFeedItem_PriceFeedItem); ok {
		return x.PriceFeedItem
	}
	return nil
}

func (m *ExtensionFeedItem) GetPromotionFeedItem() *common.PromotionFeedItem {
	if x, ok := m.GetExtension().(*ExtensionFeedItem_PromotionFeedItem); ok {
		return x.PromotionFeedItem
	}
	return nil
}

func (m *ExtensionFeedItem) GetLocationFeedItem() *common.LocationFeedItem {
	if x, ok := m.GetExtension().(*ExtensionFeedItem_LocationFeedItem); ok {
		return x.LocationFeedItem
	}
	return nil
}

func (m *ExtensionFeedItem) GetAffiliateLocationFeedItem() *common.AffiliateLocationFeedItem {
	if x, ok := m.GetExtension().(*ExtensionFeedItem_AffiliateLocationFeedItem); ok {
		return x.AffiliateLocationFeedItem
	}
	return nil
}

type isExtensionFeedItem_ServingResourceTargeting interface {
	isExtensionFeedItem_ServingResourceTargeting()
}

type ExtensionFeedItem_TargetedCampaign struct {
	TargetedCampaign *wrappers.StringValue `protobuf:"bytes,18,opt,name=targeted_campaign,json=targetedCampaign,proto3,oneof"`
}

type ExtensionFeedItem_TargetedAdGroup struct {
	TargetedAdGroup *wrappers.StringValue `protobuf:"bytes,19,opt,name=targeted_ad_group,json=targetedAdGroup,proto3,oneof"`
}

func (*ExtensionFeedItem_TargetedCampaign) isExtensionFeedItem_ServingResourceTargeting() {}

func (*ExtensionFeedItem_TargetedAdGroup) isExtensionFeedItem_ServingResourceTargeting() {}

func (m *ExtensionFeedItem) GetServingResourceTargeting() isExtensionFeedItem_ServingResourceTargeting {
	if m != nil {
		return m.ServingResourceTargeting
	}
	return nil
}

func (m *ExtensionFeedItem) GetTargetedCampaign() *wrappers.StringValue {
	if x, ok := m.GetServingResourceTargeting().(*ExtensionFeedItem_TargetedCampaign); ok {
		return x.TargetedCampaign
	}
	return nil
}

func (m *ExtensionFeedItem) GetTargetedAdGroup() *wrappers.StringValue {
	if x, ok := m.GetServingResourceTargeting().(*ExtensionFeedItem_TargetedAdGroup); ok {
		return x.TargetedAdGroup
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ExtensionFeedItem) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ExtensionFeedItem_SitelinkFeedItem)(nil),
		(*ExtensionFeedItem_StructuredSnippetFeedItem)(nil),
		(*ExtensionFeedItem_AppFeedItem)(nil),
		(*ExtensionFeedItem_CallFeedItem)(nil),
		(*ExtensionFeedItem_CalloutFeedItem)(nil),
		(*ExtensionFeedItem_TextMessageFeedItem)(nil),
		(*ExtensionFeedItem_PriceFeedItem)(nil),
		(*ExtensionFeedItem_PromotionFeedItem)(nil),
		(*ExtensionFeedItem_LocationFeedItem)(nil),
		(*ExtensionFeedItem_AffiliateLocationFeedItem)(nil),
		(*ExtensionFeedItem_TargetedCampaign)(nil),
		(*ExtensionFeedItem_TargetedAdGroup)(nil),
	}
}

func init() {
	proto.RegisterType((*ExtensionFeedItem)(nil), "google.ads.googleads.v1.resources.ExtensionFeedItem")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/extension_feed_item.proto", fileDescriptor_800f8a839b1b4536)
}

var fileDescriptor_800f8a839b1b4536 = []byte{
	// 1019 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0xdd, 0x6e, 0xdc, 0x44,
	0x14, 0x66, 0x13, 0x1a, 0xc8, 0x6c, 0x36, 0xc9, 0x4e, 0x50, 0xe5, 0x44, 0x2b, 0x48, 0x8a, 0x82,
	0x22, 0x68, 0x6d, 0x36, 0xe5, 0x86, 0x8d, 0x40, 0xf5, 0x26, 0x25, 0x29, 0x2a, 0x28, 0x75, 0xa2,
	0x54, 0x42, 0xa9, 0xcc, 0xc4, 0x3e, 0xeb, 0x4e, 0x6b, 0xcf, 0x18, 0xcf, 0x38, 0xa4, 0x2a, 0xbd,
	0x40, 0x5c, 0x71, 0xc3, 0x43, 0x70, 0xc9, 0xa3, 0xf0, 0x14, 0xbd, 0xee, 0x23, 0xe4, 0x0a, 0xed,
	0xd8, 0xe3, 0x9d, 0xfc, 0x6c, 0xbc, 0xdc, 0xcd, 0x9c, 0x39, 0xdf, 0xf7, 0xcd, 0xf9, 0x7c, 0x66,
	0x3c, 0x68, 0x2b, 0xe2, 0x3c, 0x8a, 0xc1, 0x21, 0xa1, 0x70, 0x8a, 0xe1, 0x70, 0x74, 0xda, 0x75,
	0x32, 0x10, 0x3c, 0xcf, 0x02, 0x10, 0x0e, 0x9c, 0x49, 0x60, 0x82, 0x72, 0xe6, 0x0f, 0x00, 0x42,
	0x9f, 0x4a, 0x48, 0xec, 0x34, 0xe3, 0x92, 0xe3, 0xb5, 0x02, 0x61, 0x93, 0x50, 0xd8, 0x15, 0xd8,
	0x3e, 0xed, 0xda, 0x15, 0x78, 0xe5, 0xde, 0x38, 0xfe, 0x80, 0x27, 0x09, 0x67, 0x4e, 0x90, 0x51,
	0x09, 0x19, 0x25, 0x05, 0xe3, 0x8a, 0x53, 0x93, 0x5e, 0xed, 0x45, 0x94, 0x80, 0xcd, 0x71, 0x00,
	0x60, 0x79, 0x62, 0xee, 0x5d, 0xbe, 0x4a, 0xa1, 0xc4, 0x7c, 0x75, 0x33, 0xa6, 0xaa, 0xd2, 0x17,
	0x92, 0xc8, 0x5c, 0x2b, 0x6d, 0x4d, 0x8a, 0x92, 0x24, 0x8b, 0x40, 0xfa, 0x21, 0x9c, 0xd2, 0x40,
	0x4b, 0x7e, 0xa2, 0xc1, 0x29, 0x75, 0x06, 0x14, 0xe2, 0xd0, 0x3f, 0x81, 0xe7, 0xe4, 0x94, 0xf2,
	0xac, 0x4c, 0x58, 0x36, 0x12, 0xb4, 0x7b, 0xe5, 0xd2, 0xc7, 0xe5, 0x92, 0x9a, 0x9d, 0xe4, 0x03,
	0xe7, 0xd7, 0x8c, 0xa4, 0x29, 0x64, 0x7a, 0x63, 0x1d, 0x03, 0x4a, 0x18, 0xe3, 0x92, 0xc8, 0x91,
	0x41, 0x77, 0xfe, 0x6c, 0xa3, 0xf6, 0x43, 0xed, 0xc2, 0x77, 0x00, 0xe1, 0x23, 0x09, 0x09, 0x7e,
	0x8a, 0x5a, 0x5a, 0xc5, 0x67, 0x24, 0x01, 0xab, 0xb1, 0xda, 0xd8, 0x98, 0xed, 0x6f, 0xbe, 0x75,
	0x6f, 0x9d, 0xbb, 0x77, 0xd1, 0xe7, 0xa3, 0xaf, 0x59, 0x8e, 0x52, 0x2a, 0xec, 0x80, 0x27, 0xce,
	0x15, 0x2a, 0x6f, 0x4e, 0x13, 0xfd, 0x48, 0x12, 0xc0, 0x2f, 0xd0, 0xfc, 0x45, 0xcf, 0xad, 0xd6,
	0x6a, 0x63, 0x63, 0x7e, 0xf3, 0x5b, 0x7b, 0x5c, 0xaf, 0x28, 0xfb, 0xec, 0x8a, 0xf7, 0xf0, 0x55,
	0x0a, 0x0f, 0x59, 0x9e, 0x5c, 0x8c, 0xf4, 0xa7, 0xdf, 0xba, 0xd3, 0x5e, 0x0b, 0xcc, 0x18, 0xde,
	0x41, 0x0b, 0x42, 0x92, 0x4c, 0xfa, 0x21, 0x91, 0xe0, 0x4b, 0x9a, 0x80, 0x75, 0x6b, 0xb5, 0xb1,
	0xd1, 0xdc, 0xec, 0x68, 0x31, 0x6d, 0x99, 0x7d, 0x20, 0x33, 0xca, 0xa2, 0x23, 0x12, 0xe7, 0xe0,
	0xb5, 0x14, 0x68, 0x87, 0x48, 0x38, 0xa4, 0x09, 0xe0, 0x07, 0xa8, 0x05, 0x2c, 0x34, 0x38, 0x66,
	0x26, 0xe0, 0x68, 0x02, 0x0b, 0x2b, 0x86, 0x27, 0x68, 0x8e, 0x84, 0xbe, 0x08, 0x9e, 0x43, 0x98,
	0xc7, 0x20, 0xac, 0xc5, 0xd5, 0xe9, 0x8d, 0xe6, 0xa6, 0x3d, 0xb6, 0xe2, 0xa2, 0x97, 0x6d, 0x37,
	0x3c, 0x28, 0x21, 0x8f, 0xd8, 0x80, 0x7b, 0x4d, 0x52, 0xcd, 0x05, 0x0e, 0xd1, 0x4c, 0xd1, 0x3f,
	0x56, 0x5b, 0xd9, 0xf7, 0xb8, 0xc6, 0x3e, 0xfd, 0x35, 0x0e, 0x55, 0xef, 0xed, 0x28, 0xa8, 0x72,
	0xf1, 0xba, 0x05, 0xaf, 0xe4, 0xc6, 0x7f, 0x35, 0x50, 0xa7, 0xe8, 0x56, 0x08, 0xfd, 0x08, 0xb8,
	0x6e, 0xdd, 0x80, 0x33, 0x21, 0x09, 0x93, 0xd6, 0x47, 0xf5, 0x56, 0xf4, 0x9d, 0x9b, 0x1b, 0x66,
	0x17, 0x78, 0xa1, 0xbd, 0x5d, 0x92, 0x7a, 0xcb, 0x5a, 0xf2, 0xca, 0x12, 0x7e, 0x86, 0x66, 0x8a,
	0x33, 0x67, 0xbd, 0xaf, 0xca, 0x7e, 0x30, 0x61, 0xd9, 0x07, 0x0a, 0x74, 0xa1, 0xe0, 0x22, 0x54,
	0xf4, 0x4d, 0x49, 0x8a, 0x7f, 0x46, 0x58, 0x50, 0x09, 0x31, 0x65, 0x2f, 0x47, 0x77, 0x99, 0x35,
	0xa5, 0x8a, 0xfc, 0xb2, 0xee, 0x73, 0x1d, 0x94, 0x48, 0x2d, 0xb0, 0xf7, 0x9e, 0xb7, 0x28, 0x2e,
	0xc5, 0xf0, 0x6f, 0xa8, 0x23, 0x64, 0x96, 0x07, 0x32, 0xcf, 0x20, 0xf4, 0x05, 0xa3, 0x69, 0x0a,
	0xd2, 0xd0, 0x9a, 0x56, 0x5a, 0x5f, 0xd7, 0x6a, 0x55, 0x1c, 0x07, 0x05, 0x85, 0x21, 0xba, 0x2c,
	0xc6, 0x2d, 0xe2, 0x27, 0xa8, 0x45, 0xd2, 0xd4, 0x90, 0xfb, 0x40, 0xc9, 0x7d, 0x51, 0xdb, 0x89,
	0x69, 0x6a, 0x08, 0x34, 0xc9, 0x68, 0x8a, 0x0f, 0xd1, 0x7c, 0x40, 0xe2, 0xd8, 0xe0, 0xfc, 0x50,
	0x71, 0xde, 0xad, 0xe3, 0xdc, 0x26, 0x71, 0x6c, 0x90, 0xce, 0x05, 0xc6, 0x1c, 0x3f, 0x43, 0xed,
	0xe1, 0x9c, 0xe7, 0xa6, 0x37, 0xb3, 0x8a, 0xd8, 0x99, 0x84, 0x98, 0xe7, 0xa6, 0x23, 0x0b, 0xc1,
	0xc5, 0x10, 0x7e, 0x81, 0x6e, 0x4b, 0x38, 0x93, 0x7e, 0x02, 0x42, 0x90, 0x08, 0x0c, 0x0d, 0xa4,
	0x34, 0xee, 0xd7, 0x69, 0x1c, 0xc2, 0x99, 0xfc, 0xa1, 0x00, 0x1b, 0x3a, 0x4b, 0xf2, 0x6a, 0x18,
	0x3f, 0x45, 0x0b, 0x69, 0x46, 0x03, 0x53, 0xa4, 0xa9, 0x44, 0xee, 0xd5, 0x89, 0xec, 0x0f, 0x61,
	0x06, 0x7d, 0x2b, 0x35, 0x03, 0x38, 0x40, 0x4b, 0x69, 0xc6, 0x13, 0x2e, 0x2f, 0xfc, 0x79, 0xad,
	0x39, 0x45, 0xde, 0xad, 0x27, 0x2f, 0xa1, 0x86, 0x40, 0x3b, 0xbd, 0x1c, 0xc4, 0x80, 0x70, 0xcc,
	0x03, 0x72, 0x49, 0x63, 0x7e, 0xb2, 0x13, 0xf1, 0xb8, 0x44, 0x6a, 0x36, 0x75, 0xd8, 0x86, 0xc7,
	0x22, 0xbe, 0xb4, 0x80, 0x7f, 0x6f, 0xa0, 0x0e, 0x19, 0x0c, 0x68, 0x4c, 0x87, 0xb7, 0xec, 0x35,
	0x8a, 0x0b, 0x93, 0x9d, 0x0b, 0x57, 0x73, 0x8c, 0x93, 0x5e, 0x26, 0xe3, 0x32, 0xb0, 0x40, 0xed,
	0xea, 0xae, 0x0b, 0x48, 0x92, 0x12, 0x1a, 0x31, 0x0b, 0x4f, 0x70, 0xc1, 0x7d, 0x76, 0xee, 0x7e,
	0x8a, 0xd6, 0xc6, 0x5e, 0x70, 0xdb, 0x25, 0xd7, 0x5e, 0xc3, 0x5b, 0xd4, 0x02, 0x3a, 0x86, 0x7f,
	0x31, 0x44, 0x49, 0xe8, 0x47, 0x19, 0xcf, 0x53, 0x6b, 0x69, 0x02, 0xd1, 0xf5, 0x73, 0xf7, 0x0e,
	0x5a, 0x1d, 0x2b, 0xea, 0x86, 0xbb, 0x43, 0xaa, 0xbd, 0x86, 0xb7, 0xa0, 0xf9, 0xcb, 0x50, 0x8f,
	0xbd, 0x73, 0x5f, 0xfe, 0x9f, 0x1f, 0x38, 0xfe, 0x26, 0xc8, 0x85, 0xe4, 0x09, 0x64, 0xc2, 0x79,
	0xad, 0x87, 0x6f, 0x46, 0x2f, 0x27, 0x9d, 0x27, 0x9c, 0xd7, 0xd7, 0xbc, 0x04, 0xdf, 0xf4, 0x9b,
	0x68, 0xb6, 0x8a, 0xf7, 0x3b, 0x68, 0x45, 0x40, 0x76, 0x4a, 0x59, 0xe4, 0x57, 0xef, 0x8b, 0x62,
	0x83, 0x94, 0x45, 0xfd, 0x3f, 0xa6, 0xd0, 0x7a, 0xc0, 0x13, 0xbb, 0xf6, 0xd9, 0xd8, 0xbf, 0x7d,
	0x65, 0x9b, 0xfb, 0x43, 0x97, 0xf6, 0x1b, 0x3f, 0x7d, 0x5f, 0x82, 0x23, 0x1e, 0x13, 0x16, 0xd9,
	0x3c, 0x8b, 0x9c, 0x08, 0x98, 0xf2, 0xd0, 0x19, 0x95, 0x7a, 0xc3, 0x7b, 0x76, 0xab, 0x1a, 0xfd,
	0x3d, 0x35, 0xbd, 0xeb, 0xba, 0xff, 0x4c, 0xad, 0xed, 0x16, 0x94, 0x6e, 0x28, 0xec, 0x62, 0x38,
	0x1c, 0x1d, 0x75, 0x6d, 0x4f, 0x67, 0xfe, 0xab, 0x73, 0x8e, 0xdd, 0x50, 0x1c, 0x57, 0x39, 0xc7,
	0x47, 0xdd, 0xe3, 0x2a, 0xe7, 0xdd, 0xd4, 0x7a, 0xb1, 0xd0, 0xeb, 0xb9, 0xa1, 0xe8, 0xf5, 0xaa,
	0xac, 0x5e, 0xef, 0xa8, 0xdb, 0xeb, 0x55, 0x79, 0x27, 0x33, 0x6a, 0xb3, 0xf7, 0xff, 0x0b, 0x00,
	0x00, 0xff, 0xff, 0x5c, 0xb3, 0xdc, 0x77, 0x7b, 0x0b, 0x00, 0x00,
}
