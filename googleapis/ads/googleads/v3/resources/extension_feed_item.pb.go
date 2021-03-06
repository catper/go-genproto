// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/extension_feed_item.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
	wrappers "github.com/catper/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v3/common"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v3/enums"
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
	// The resource name of the extension feed item.
	// Extension feed item resource names have the form:
	//
	// `customers/{customer_id}/extensionFeedItems/{feed_item_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of this feed item. Read-only.
	Id *wrappers.Int64Value `protobuf:"bytes,24,opt,name=id,proto3" json:"id,omitempty"`
	// The extension type of the extension feed item.
	// This field is read-only.
	ExtensionType enums.ExtensionTypeEnum_ExtensionType `protobuf:"varint,13,opt,name=extension_type,json=extensionType,proto3,enum=google.ads.googleads.v3.enums.ExtensionTypeEnum_ExtensionType" json:"extension_type,omitempty"`
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
	Device enums.FeedItemTargetDeviceEnum_FeedItemTargetDevice `protobuf:"varint,17,opt,name=device,proto3,enum=google.ads.googleads.v3.enums.FeedItemTargetDeviceEnum_FeedItemTargetDevice" json:"device,omitempty"`
	// The targeted geo target constant.
	TargetedGeoTargetConstant *wrappers.StringValue `protobuf:"bytes,20,opt,name=targeted_geo_target_constant,json=targetedGeoTargetConstant,proto3" json:"targeted_geo_target_constant,omitempty"`
	// The targeted keyword.
	TargetedKeyword *common.KeywordInfo `protobuf:"bytes,22,opt,name=targeted_keyword,json=targetedKeyword,proto3" json:"targeted_keyword,omitempty"`
	// Status of the feed item.
	// This field is read-only.
	Status enums.FeedItemStatusEnum_FeedItemStatus `protobuf:"varint,4,opt,name=status,proto3,enum=google.ads.googleads.v3.enums.FeedItemStatusEnum_FeedItemStatus" json:"status,omitempty"`
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
	//	*ExtensionFeedItem_HotelCalloutFeedItem
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
	return fileDescriptor_d54452736e95564f, []int{0}
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

func (m *ExtensionFeedItem) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
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

func (m *ExtensionFeedItem) GetTargetedKeyword() *common.KeywordInfo {
	if m != nil {
		return m.TargetedKeyword
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

type ExtensionFeedItem_HotelCalloutFeedItem struct {
	HotelCalloutFeedItem *common.HotelCalloutFeedItem `protobuf:"bytes,23,opt,name=hotel_callout_feed_item,json=hotelCalloutFeedItem,proto3,oneof"`
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

func (*ExtensionFeedItem_HotelCalloutFeedItem) isExtensionFeedItem_Extension() {}

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

func (m *ExtensionFeedItem) GetHotelCalloutFeedItem() *common.HotelCalloutFeedItem {
	if x, ok := m.GetExtension().(*ExtensionFeedItem_HotelCalloutFeedItem); ok {
		return x.HotelCalloutFeedItem
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
		(*ExtensionFeedItem_HotelCalloutFeedItem)(nil),
		(*ExtensionFeedItem_TargetedCampaign)(nil),
		(*ExtensionFeedItem_TargetedAdGroup)(nil),
	}
}

func init() {
	proto.RegisterType((*ExtensionFeedItem)(nil), "google.ads.googleads.v3.resources.ExtensionFeedItem")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/extension_feed_item.proto", fileDescriptor_d54452736e95564f)
}

var fileDescriptor_d54452736e95564f = []byte{
	// 1013 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0x61, 0x4f, 0xdc, 0x36,
	0x18, 0xde, 0x5d, 0x37, 0x36, 0x7c, 0x1c, 0x70, 0xa6, 0xa2, 0x81, 0xa1, 0x89, 0x6e, 0xaa, 0x84,
	0xd6, 0x35, 0xd9, 0x38, 0x34, 0x69, 0x87, 0x36, 0xf5, 0x80, 0x0e, 0x68, 0xbb, 0x89, 0x1e, 0x88,
	0x4d, 0x13, 0x55, 0xe6, 0xc6, 0x2f, 0xc1, 0x23, 0xb1, 0xa3, 0xd8, 0xa1, 0xa0, 0xae, 0x9f, 0xf6,
	0x4f, 0xf6, 0x71, 0x5f, 0xf7, 0x2f, 0xf6, 0x53, 0xfa, 0x2b, 0xa6, 0x38, 0x89, 0xcf, 0x70, 0x5c,
	0x73, 0xdf, 0xec, 0xd7, 0xef, 0xf3, 0x3c, 0xaf, 0x9f, 0xbc, 0x76, 0x8c, 0x36, 0x43, 0x21, 0xc2,
	0x08, 0x3c, 0x42, 0xa5, 0x57, 0x0c, 0xf3, 0xd1, 0x45, 0xd7, 0x4b, 0x41, 0x8a, 0x2c, 0x0d, 0x40,
	0x7a, 0x70, 0xa9, 0x80, 0x4b, 0x26, 0xb8, 0x7f, 0x0a, 0x40, 0x7d, 0xa6, 0x20, 0x76, 0x93, 0x54,
	0x28, 0x81, 0xef, 0x17, 0x08, 0x97, 0x50, 0xe9, 0x1a, 0xb0, 0x7b, 0xd1, 0x75, 0x0d, 0x78, 0xf9,
	0xd1, 0x38, 0xfe, 0x40, 0xc4, 0xb1, 0xe0, 0x5e, 0x90, 0x32, 0x05, 0x29, 0x23, 0x05, 0xe3, 0xb2,
	0x57, 0x93, 0x6e, 0x6a, 0x91, 0x25, 0x60, 0x7d, 0x1c, 0x00, 0x78, 0x16, 0xdb, 0xb5, 0xab, 0xab,
	0x04, 0x4a, 0xcc, 0xc6, 0xfb, 0x31, 0x66, 0x97, 0xbe, 0x54, 0x44, 0x65, 0x95, 0xd2, 0xe6, 0xa4,
	0x28, 0x45, 0xd2, 0x10, 0x94, 0x4f, 0xe1, 0x82, 0x05, 0x95, 0xe4, 0x52, 0x05, 0x4e, 0x98, 0x71,
	0xb6, 0x5c, 0xfa, 0xac, 0x5c, 0xd2, 0xb3, 0x57, 0xd9, 0xa9, 0xf7, 0x3a, 0x25, 0x49, 0x02, 0x69,
	0xa5, 0xbb, 0x62, 0x41, 0x09, 0xe7, 0x42, 0x11, 0x35, 0xdc, 0xff, 0xe7, 0xff, 0x76, 0x50, 0xe7,
	0x49, 0xb5, 0xc9, 0x1f, 0x01, 0xe8, 0xbe, 0x82, 0x18, 0x7f, 0x81, 0xda, 0x95, 0x8a, 0xcf, 0x49,
	0x0c, 0x4e, 0x63, 0xb5, 0xb1, 0x36, 0x3d, 0x98, 0xa9, 0x82, 0x3f, 0x93, 0x18, 0xf0, 0x43, 0xd4,
	0x64, 0xd4, 0x71, 0x56, 0x1b, 0x6b, 0xad, 0xf5, 0x4f, 0xcb, 0xef, 0xe7, 0x56, 0x55, 0xb8, 0xfb,
	0x5c, 0x7d, 0xbb, 0x71, 0x4c, 0xa2, 0x0c, 0x06, 0x4d, 0x46, 0x31, 0xa0, 0xd9, 0xeb, 0x5e, 0x3a,
	0xed, 0xd5, 0xc6, 0xda, 0xec, 0xfa, 0x0f, 0xee, 0xb8, 0x1e, 0xd0, 0xb6, 0xb8, 0xa6, 0xb6, 0xa3,
	0xab, 0x04, 0x9e, 0xf0, 0x2c, 0xbe, 0x1e, 0x19, 0xb4, 0xc1, 0x9e, 0xe2, 0x1d, 0x34, 0x27, 0x15,
	0x49, 0x95, 0x4f, 0x89, 0x02, 0x5f, 0xb1, 0x18, 0x9c, 0x8f, 0x74, 0x81, 0x2b, 0x23, 0x05, 0x1e,
	0xaa, 0x94, 0xf1, 0xb0, 0xa8, 0xb0, 0xad, 0x41, 0x3b, 0x44, 0xc1, 0x11, 0x8b, 0x01, 0x3f, 0x46,
	0x6d, 0xe0, 0xd4, 0xe2, 0x98, 0x9a, 0x80, 0xa3, 0x05, 0x9c, 0x1a, 0x86, 0x17, 0x68, 0x86, 0x50,
	0x5f, 0x06, 0x67, 0x40, 0xb3, 0x08, 0xa4, 0x33, 0xbf, 0x7a, 0x67, 0xad, 0xb5, 0xee, 0x8e, 0xdd,
	0x6c, 0xd1, 0x9e, 0x6e, 0x9f, 0x1e, 0x96, 0x90, 0x7d, 0x7e, 0x2a, 0x06, 0x2d, 0x62, 0xe6, 0x12,
	0x53, 0x34, 0x55, 0xb4, 0x84, 0xd3, 0xd1, 0xce, 0x3d, 0xaf, 0x71, 0xae, 0xfa, 0x98, 0x47, 0xba,
	0x9d, 0x76, 0x34, 0x54, 0x1b, 0x78, 0xdb, 0xc2, 0xa0, 0xe4, 0xc6, 0x2f, 0xd1, 0x4a, 0xd1, 0x7f,
	0x40, 0xfd, 0x10, 0x44, 0xd5, 0x8c, 0x81, 0xe0, 0x52, 0x11, 0xae, 0x9c, 0xbb, 0x13, 0x38, 0xb1,
	0x54, 0x31, 0xec, 0x82, 0x28, 0x44, 0xb6, 0x4b, 0x38, 0x3e, 0x46, 0xf3, 0x86, 0xfe, 0x1c, 0xae,
	0x5e, 0x8b, 0x94, 0x3a, 0x8b, 0x9a, 0xf2, 0x61, 0x9d, 0x37, 0xcf, 0x8a, 0x74, 0x6d, 0xcc, 0x5c,
	0x45, 0x52, 0x06, 0xf1, 0xaf, 0x68, 0xaa, 0x38, 0x6c, 0xce, 0x87, 0xda, 0x9c, 0xc7, 0x13, 0x9a,
	0x73, 0xa8, 0x41, 0xd7, 0x6c, 0x29, 0x42, 0x83, 0x92, 0x0f, 0xff, 0x8e, 0xb0, 0x64, 0x0a, 0x22,
	0xc6, 0xcf, 0x87, 0xf7, 0x97, 0xd3, 0xd4, 0x35, 0x7f, 0x5d, 0x57, 0xf3, 0x61, 0x89, 0xac, 0xb8,
	0xf7, 0x3e, 0x18, 0xcc, 0xcb, 0x1b, 0x31, 0xfc, 0x27, 0x5a, 0x91, 0x2a, 0xcd, 0x02, 0x95, 0xa5,
	0x40, 0x7d, 0xc9, 0x59, 0x92, 0x80, 0xb2, 0xb4, 0xee, 0x68, 0xad, 0xef, 0x6a, 0xb5, 0x0c, 0xc7,
	0x61, 0x41, 0x61, 0x89, 0x2e, 0xc9, 0x71, 0x8b, 0xf8, 0x05, 0x6a, 0x93, 0x24, 0xb1, 0xe4, 0x3e,
	0x9e, 0xec, 0x73, 0xf4, 0x93, 0xc4, 0x12, 0x68, 0x91, 0xe1, 0x14, 0x1f, 0xa1, 0xd9, 0x80, 0x44,
	0x91, 0xc5, 0xf9, 0x89, 0xe6, 0xfc, 0xaa, 0x8e, 0x73, 0x9b, 0x44, 0x91, 0x45, 0x3a, 0x13, 0x58,
	0x73, 0xfc, 0x12, 0x75, 0xf2, 0xb9, 0xc8, 0x6c, 0x6f, 0xa6, 0x35, 0xb1, 0x37, 0x09, 0xb1, 0xc8,
	0x6c, 0x47, 0xe6, 0x82, 0xeb, 0x21, 0xfc, 0x07, 0x5a, 0x54, 0x70, 0xa9, 0xfc, 0x18, 0xa4, 0x24,
	0x21, 0x58, 0x1a, 0x48, 0x6b, 0x74, 0xeb, 0x34, 0x8e, 0xe0, 0x52, 0xfd, 0x54, 0x80, 0x2d, 0x9d,
	0x05, 0x35, 0x1a, 0xc6, 0xbf, 0xa0, 0xb9, 0x24, 0x65, 0x81, 0x2d, 0xd2, 0xd2, 0x22, 0x8f, 0xea,
	0x44, 0x0e, 0x72, 0x98, 0x45, 0xdf, 0x4e, 0xec, 0x00, 0x0e, 0xd0, 0x42, 0x92, 0x8a, 0x58, 0xa8,
	0x6b, 0x7f, 0x5b, 0x67, 0x46, 0x93, 0x7f, 0x53, 0x4f, 0x5e, 0x42, 0x2d, 0x81, 0x4e, 0x72, 0x33,
	0x98, 0x9f, 0x88, 0x48, 0x04, 0xe4, 0x86, 0xc6, 0xec, 0x64, 0x27, 0xe2, 0x79, 0x89, 0xb4, 0x4f,
	0x44, 0x74, 0x23, 0x96, 0x9f, 0x08, 0x72, 0x7a, 0xca, 0x22, 0x96, 0x5f, 0xc0, 0xb7, 0x68, 0xcd,
	0x4d, 0x76, 0x22, 0xfa, 0x15, 0xc7, 0x2d, 0xa2, 0x4b, 0x64, 0xdc, 0x22, 0x8e, 0xd1, 0xbd, 0x33,
	0xa1, 0x20, 0xf2, 0x47, 0xdb, 0xed, 0x9e, 0x16, 0xde, 0xa8, 0x13, 0xde, 0xcb, 0xe1, 0xa3, 0x3d,
	0x77, 0xf7, 0xec, 0x96, 0x38, 0x7e, 0x86, 0x3a, 0xe6, 0x4a, 0x0c, 0x48, 0x9c, 0x10, 0x16, 0x72,
	0x07, 0xd7, 0x5f, 0xb3, 0x7b, 0x8d, 0x81, 0xb9, 0x4b, 0xb7, 0x4b, 0x1c, 0x7e, 0x6a, 0x91, 0x11,
	0xea, 0x87, 0xa9, 0xc8, 0x12, 0x67, 0x61, 0x22, 0x32, 0x73, 0xa7, 0xf6, 0xe9, 0x6e, 0x0e, 0xeb,
	0xf1, 0x77, 0xfd, 0x73, 0xf4, 0xe5, 0x70, 0x7f, 0xe5, 0x28, 0x61, 0x32, 0xdf, 0xa7, 0x37, 0xfa,
	0x6a, 0xf8, 0x3e, 0xc8, 0xa4, 0x12, 0x31, 0xa4, 0xd2, 0x7b, 0x53, 0x0d, 0xdf, 0x0e, 0x9f, 0x50,
	0x55, 0x9e, 0xf4, 0xde, 0xdc, 0xf2, 0x24, 0x7c, 0xbb, 0xd5, 0x42, 0xd3, 0x26, 0xbe, 0xb5, 0x82,
	0x96, 0x25, 0xa4, 0x17, 0x8c, 0x87, 0xbe, 0x79, 0x89, 0x14, 0x05, 0x32, 0x1e, 0x6e, 0xfd, 0xd5,
	0x44, 0x0f, 0x02, 0x11, 0xbb, 0xb5, 0xef, 0xc7, 0xad, 0xc5, 0x91, 0x32, 0x0f, 0xf2, 0xed, 0x1f,
	0x34, 0x7e, 0x7b, 0x5a, 0x82, 0x43, 0x11, 0x11, 0x1e, 0xba, 0x22, 0x0d, 0xbd, 0x10, 0xb8, 0x36,
	0xc7, 0x1b, 0x6e, 0xf5, 0x3d, 0x0f, 0xdb, 0x4d, 0x33, 0xfa, 0xbb, 0x79, 0x67, 0xb7, 0xdf, 0xff,
	0xa7, 0x79, 0x7f, 0xb7, 0xa0, 0xec, 0x53, 0xe9, 0x16, 0xc3, 0x7c, 0x74, 0xdc, 0x75, 0x07, 0x55,
	0xe6, 0x7f, 0x55, 0xce, 0x49, 0x9f, 0xca, 0x13, 0x93, 0x73, 0x72, 0xdc, 0x3d, 0x31, 0x39, 0xef,
	0x9a, 0x0f, 0x8a, 0x85, 0x5e, 0xaf, 0x4f, 0x65, 0xaf, 0x67, 0xb2, 0x7a, 0xbd, 0xe3, 0x6e, 0xaf,
	0x67, 0xf2, 0x5e, 0x4d, 0xe9, 0x62, 0xbb, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0x03, 0xd9, 0x3f,
	0x01, 0x84, 0x0b, 0x00, 0x00,
}
