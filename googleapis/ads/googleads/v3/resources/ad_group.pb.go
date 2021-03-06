// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/ad_group.proto

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

// An ad group.
type AdGroup struct {
	// The resource name of the ad group.
	// Ad group resource names have the form:
	//
	// `customers/{customer_id}/adGroups/{ad_group_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the ad group.
	Id *wrappers.Int64Value `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the ad group.
	//
	// This field is required and should not be empty when creating new ad
	// groups.
	//
	// It must contain fewer than 255 UTF-8 full-width characters.
	//
	// It must not contain any null (code point 0x0), NL line feed
	// (code point 0xA) or carriage return (code point 0xD) characters.
	Name *wrappers.StringValue `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// The status of the ad group.
	Status enums.AdGroupStatusEnum_AdGroupStatus `protobuf:"varint,5,opt,name=status,proto3,enum=google.ads.googleads.v3.enums.AdGroupStatusEnum_AdGroupStatus" json:"status,omitempty"`
	// The type of the ad group.
	Type enums.AdGroupTypeEnum_AdGroupType `protobuf:"varint,12,opt,name=type,proto3,enum=google.ads.googleads.v3.enums.AdGroupTypeEnum_AdGroupType" json:"type,omitempty"`
	// The ad rotation mode of the ad group.
	AdRotationMode enums.AdGroupAdRotationModeEnum_AdGroupAdRotationMode `protobuf:"varint,22,opt,name=ad_rotation_mode,json=adRotationMode,proto3,enum=google.ads.googleads.v3.enums.AdGroupAdRotationModeEnum_AdGroupAdRotationMode" json:"ad_rotation_mode,omitempty"`
	// For draft or experiment ad groups, this field is the resource name of the
	// base ad group from which this ad group was created. If a draft or
	// experiment ad group does not have a base ad group, then this field is null.
	//
	// For base ad groups, this field equals the ad group resource name.
	//
	// This field is read-only.
	BaseAdGroup *wrappers.StringValue `protobuf:"bytes,18,opt,name=base_ad_group,json=baseAdGroup,proto3" json:"base_ad_group,omitempty"`
	// The URL template for constructing a tracking URL.
	TrackingUrlTemplate *wrappers.StringValue `protobuf:"bytes,13,opt,name=tracking_url_template,json=trackingUrlTemplate,proto3" json:"tracking_url_template,omitempty"`
	// The list of mappings used to substitute custom parameter tags in a
	// `tracking_url_template`, `final_urls`, or `mobile_final_urls`.
	UrlCustomParameters []*common.CustomParameter `protobuf:"bytes,6,rep,name=url_custom_parameters,json=urlCustomParameters,proto3" json:"url_custom_parameters,omitempty"`
	// The campaign to which the ad group belongs.
	Campaign *wrappers.StringValue `protobuf:"bytes,10,opt,name=campaign,proto3" json:"campaign,omitempty"`
	// The maximum CPC (cost-per-click) bid.
	CpcBidMicros *wrappers.Int64Value `protobuf:"bytes,14,opt,name=cpc_bid_micros,json=cpcBidMicros,proto3" json:"cpc_bid_micros,omitempty"`
	// The maximum CPM (cost-per-thousand viewable impressions) bid.
	CpmBidMicros *wrappers.Int64Value `protobuf:"bytes,15,opt,name=cpm_bid_micros,json=cpmBidMicros,proto3" json:"cpm_bid_micros,omitempty"`
	// The target CPA (cost-per-acquisition).
	TargetCpaMicros *wrappers.Int64Value `protobuf:"bytes,27,opt,name=target_cpa_micros,json=targetCpaMicros,proto3" json:"target_cpa_micros,omitempty"`
	// The CPV (cost-per-view) bid.
	CpvBidMicros *wrappers.Int64Value `protobuf:"bytes,17,opt,name=cpv_bid_micros,json=cpvBidMicros,proto3" json:"cpv_bid_micros,omitempty"`
	// Average amount in micros that the advertiser is willing to pay for every
	// thousand times the ad is shown.
	TargetCpmMicros *wrappers.Int64Value `protobuf:"bytes,26,opt,name=target_cpm_micros,json=targetCpmMicros,proto3" json:"target_cpm_micros,omitempty"`
	// The target ROAS (return-on-ad-spend) override. If the ad group's campaign
	// bidding strategy is a standard Target ROAS strategy, then this field
	// overrides the target ROAS specified in the campaign's bidding strategy.
	// Otherwise, this value is ignored.
	TargetRoas *wrappers.DoubleValue `protobuf:"bytes,30,opt,name=target_roas,json=targetRoas,proto3" json:"target_roas,omitempty"`
	// The percent cpc bid amount, expressed as a fraction of the advertised price
	// for some good or service. The valid range for the fraction is [0,1) and the
	// value stored here is 1,000,000 * [fraction].
	PercentCpcBidMicros *wrappers.Int64Value `protobuf:"bytes,20,opt,name=percent_cpc_bid_micros,json=percentCpcBidMicros,proto3" json:"percent_cpc_bid_micros,omitempty"`
	// Settings for the Display Campaign Optimizer, initially termed "Explorer".
	ExplorerAutoOptimizerSetting *common.ExplorerAutoOptimizerSetting `protobuf:"bytes,21,opt,name=explorer_auto_optimizer_setting,json=explorerAutoOptimizerSetting,proto3" json:"explorer_auto_optimizer_setting,omitempty"`
	// Allows advertisers to specify a targeting dimension on which to place
	// absolute bids. This is only applicable for campaigns that target only the
	// display network and not search.
	DisplayCustomBidDimension enums.TargetingDimensionEnum_TargetingDimension `protobuf:"varint,23,opt,name=display_custom_bid_dimension,json=displayCustomBidDimension,proto3,enum=google.ads.googleads.v3.enums.TargetingDimensionEnum_TargetingDimension" json:"display_custom_bid_dimension,omitempty"`
	// URL template for appending params to Final URL.
	FinalUrlSuffix *wrappers.StringValue `protobuf:"bytes,24,opt,name=final_url_suffix,json=finalUrlSuffix,proto3" json:"final_url_suffix,omitempty"`
	// Setting for targeting related features.
	TargetingSetting *common.TargetingSetting `protobuf:"bytes,25,opt,name=targeting_setting,json=targetingSetting,proto3" json:"targeting_setting,omitempty"`
	// The effective target CPA (cost-per-acquisition).
	// This field is read-only.
	EffectiveTargetCpaMicros *wrappers.Int64Value `protobuf:"bytes,28,opt,name=effective_target_cpa_micros,json=effectiveTargetCpaMicros,proto3" json:"effective_target_cpa_micros,omitempty"`
	// Source of the effective target CPA.
	// This field is read-only.
	EffectiveTargetCpaSource enums.BiddingSourceEnum_BiddingSource `protobuf:"varint,29,opt,name=effective_target_cpa_source,json=effectiveTargetCpaSource,proto3,enum=google.ads.googleads.v3.enums.BiddingSourceEnum_BiddingSource" json:"effective_target_cpa_source,omitempty"`
	// The effective target ROAS (return-on-ad-spend).
	// This field is read-only.
	EffectiveTargetRoas *wrappers.DoubleValue `protobuf:"bytes,31,opt,name=effective_target_roas,json=effectiveTargetRoas,proto3" json:"effective_target_roas,omitempty"`
	// Source of the effective target ROAS.
	// This field is read-only.
	EffectiveTargetRoasSource enums.BiddingSourceEnum_BiddingSource `protobuf:"varint,32,opt,name=effective_target_roas_source,json=effectiveTargetRoasSource,proto3,enum=google.ads.googleads.v3.enums.BiddingSourceEnum_BiddingSource" json:"effective_target_roas_source,omitempty"`
	// The resource names of labels attached to this ad group.
	Labels               []*wrappers.StringValue `protobuf:"bytes,33,rep,name=labels,proto3" json:"labels,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *AdGroup) Reset()         { *m = AdGroup{} }
func (m *AdGroup) String() string { return proto.CompactTextString(m) }
func (*AdGroup) ProtoMessage()    {}
func (*AdGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_3478586bae6468dc, []int{0}
}

func (m *AdGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroup.Unmarshal(m, b)
}
func (m *AdGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroup.Marshal(b, m, deterministic)
}
func (m *AdGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroup.Merge(m, src)
}
func (m *AdGroup) XXX_Size() int {
	return xxx_messageInfo_AdGroup.Size(m)
}
func (m *AdGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroup.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroup proto.InternalMessageInfo

func (m *AdGroup) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *AdGroup) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *AdGroup) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *AdGroup) GetStatus() enums.AdGroupStatusEnum_AdGroupStatus {
	if m != nil {
		return m.Status
	}
	return enums.AdGroupStatusEnum_UNSPECIFIED
}

func (m *AdGroup) GetType() enums.AdGroupTypeEnum_AdGroupType {
	if m != nil {
		return m.Type
	}
	return enums.AdGroupTypeEnum_UNSPECIFIED
}

func (m *AdGroup) GetAdRotationMode() enums.AdGroupAdRotationModeEnum_AdGroupAdRotationMode {
	if m != nil {
		return m.AdRotationMode
	}
	return enums.AdGroupAdRotationModeEnum_UNSPECIFIED
}

func (m *AdGroup) GetBaseAdGroup() *wrappers.StringValue {
	if m != nil {
		return m.BaseAdGroup
	}
	return nil
}

func (m *AdGroup) GetTrackingUrlTemplate() *wrappers.StringValue {
	if m != nil {
		return m.TrackingUrlTemplate
	}
	return nil
}

func (m *AdGroup) GetUrlCustomParameters() []*common.CustomParameter {
	if m != nil {
		return m.UrlCustomParameters
	}
	return nil
}

func (m *AdGroup) GetCampaign() *wrappers.StringValue {
	if m != nil {
		return m.Campaign
	}
	return nil
}

func (m *AdGroup) GetCpcBidMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CpcBidMicros
	}
	return nil
}

func (m *AdGroup) GetCpmBidMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CpmBidMicros
	}
	return nil
}

func (m *AdGroup) GetTargetCpaMicros() *wrappers.Int64Value {
	if m != nil {
		return m.TargetCpaMicros
	}
	return nil
}

func (m *AdGroup) GetCpvBidMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CpvBidMicros
	}
	return nil
}

func (m *AdGroup) GetTargetCpmMicros() *wrappers.Int64Value {
	if m != nil {
		return m.TargetCpmMicros
	}
	return nil
}

func (m *AdGroup) GetTargetRoas() *wrappers.DoubleValue {
	if m != nil {
		return m.TargetRoas
	}
	return nil
}

func (m *AdGroup) GetPercentCpcBidMicros() *wrappers.Int64Value {
	if m != nil {
		return m.PercentCpcBidMicros
	}
	return nil
}

func (m *AdGroup) GetExplorerAutoOptimizerSetting() *common.ExplorerAutoOptimizerSetting {
	if m != nil {
		return m.ExplorerAutoOptimizerSetting
	}
	return nil
}

func (m *AdGroup) GetDisplayCustomBidDimension() enums.TargetingDimensionEnum_TargetingDimension {
	if m != nil {
		return m.DisplayCustomBidDimension
	}
	return enums.TargetingDimensionEnum_UNSPECIFIED
}

func (m *AdGroup) GetFinalUrlSuffix() *wrappers.StringValue {
	if m != nil {
		return m.FinalUrlSuffix
	}
	return nil
}

func (m *AdGroup) GetTargetingSetting() *common.TargetingSetting {
	if m != nil {
		return m.TargetingSetting
	}
	return nil
}

func (m *AdGroup) GetEffectiveTargetCpaMicros() *wrappers.Int64Value {
	if m != nil {
		return m.EffectiveTargetCpaMicros
	}
	return nil
}

func (m *AdGroup) GetEffectiveTargetCpaSource() enums.BiddingSourceEnum_BiddingSource {
	if m != nil {
		return m.EffectiveTargetCpaSource
	}
	return enums.BiddingSourceEnum_UNSPECIFIED
}

func (m *AdGroup) GetEffectiveTargetRoas() *wrappers.DoubleValue {
	if m != nil {
		return m.EffectiveTargetRoas
	}
	return nil
}

func (m *AdGroup) GetEffectiveTargetRoasSource() enums.BiddingSourceEnum_BiddingSource {
	if m != nil {
		return m.EffectiveTargetRoasSource
	}
	return enums.BiddingSourceEnum_UNSPECIFIED
}

func (m *AdGroup) GetLabels() []*wrappers.StringValue {
	if m != nil {
		return m.Labels
	}
	return nil
}

func init() {
	proto.RegisterType((*AdGroup)(nil), "google.ads.googleads.v3.resources.AdGroup")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/ad_group.proto", fileDescriptor_3478586bae6468dc)
}

var fileDescriptor_3478586bae6468dc = []byte{
	// 1026 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0xdd, 0x6e, 0xe4, 0x34,
	0x14, 0xc7, 0x35, 0x6d, 0x29, 0xe0, 0x7e, 0x6c, 0x9b, 0xd2, 0x25, 0x6d, 0x87, 0xdd, 0x29, 0x68,
	0xa5, 0x4a, 0x48, 0x49, 0xe9, 0x2c, 0x0b, 0x0a, 0x2c, 0x22, 0x6d, 0x97, 0x02, 0x62, 0xbb, 0x55,
	0xda, 0xed, 0xc5, 0xaa, 0x28, 0xf2, 0xc4, 0x9e, 0xc8, 0x22, 0x89, 0x23, 0xdb, 0x29, 0x2d, 0xab,
	0x8a, 0x0b, 0xae, 0x78, 0x0d, 0x2e, 0x79, 0x06, 0x9e, 0x80, 0x47, 0xd9, 0x37, 0xe0, 0x0e, 0xc5,
	0x1f, 0x99, 0x8f, 0x76, 0x26, 0x41, 0xe2, 0xce, 0xf1, 0x39, 0xff, 0x9f, 0x8f, 0xed, 0x73, 0x7c,
	0x02, 0x76, 0x63, 0x4a, 0xe3, 0x04, 0xbb, 0x10, 0x71, 0x57, 0x0d, 0xcb, 0xd1, 0x65, 0xd7, 0x65,
	0x98, 0xd3, 0x82, 0x45, 0x98, 0xbb, 0x10, 0x85, 0x31, 0xa3, 0x45, 0xee, 0xe4, 0x8c, 0x0a, 0x6a,
	0x6d, 0x2b, 0x37, 0x07, 0x22, 0xee, 0x54, 0x0a, 0xe7, 0xb2, 0xeb, 0x54, 0x8a, 0xcd, 0x4f, 0x27,
	0x41, 0x23, 0x9a, 0xa6, 0x34, 0x73, 0xa3, 0x82, 0x0b, 0x9a, 0x86, 0x39, 0x64, 0x30, 0xc5, 0x02,
	0x33, 0x45, 0xde, 0x3c, 0xac, 0x91, 0xe1, 0xab, 0x3c, 0xa1, 0x0c, 0xb3, 0x10, 0x16, 0x82, 0x86,
	0x34, 0x17, 0x24, 0x25, 0xbf, 0x60, 0x16, 0x72, 0x2c, 0x04, 0xc9, 0x62, 0x4d, 0x79, 0x52, 0x43,
	0x11, 0x90, 0xc5, 0xb8, 0xf4, 0x1f, 0xd3, 0x3d, 0x9d, 0xa4, 0xc3, 0x59, 0x91, 0x0e, 0x4e, 0x21,
	0x84, 0x28, 0x64, 0x54, 0x40, 0x41, 0x68, 0x16, 0xa6, 0x14, 0x61, 0x2d, 0xef, 0x36, 0x94, 0x73,
	0x01, 0x45, 0xc1, 0xb5, 0xe8, 0x93, 0x86, 0x22, 0x71, 0x9d, 0x9b, 0x75, 0xf6, 0xa6, 0x4b, 0x7a,
	0x04, 0x21, 0xb9, 0x37, 0x79, 0x15, 0x5a, 0xf3, 0xd9, 0x74, 0xcd, 0xe0, 0x44, 0x10, 0x49, 0x71,
	0xc6, 0x09, 0xcd, 0xb4, 0x70, 0xc3, 0x08, 0x73, 0x52, 0x25, 0x84, 0x36, 0x3d, 0xd0, 0x26, 0xf9,
	0xd5, 0x2b, 0xfa, 0xee, 0xcf, 0x0c, 0xe6, 0x39, 0x66, 0x66, 0x6b, 0xed, 0x21, 0x29, 0xcc, 0x32,
	0x7d, 0x62, 0xda, 0xfa, 0xe1, 0x5f, 0xab, 0xe0, 0x6d, 0x1f, 0x1d, 0x95, 0x9b, 0xb3, 0x3e, 0x02,
	0x4b, 0x86, 0x1d, 0x66, 0x30, 0xc5, 0x76, 0xab, 0xd3, 0xda, 0x79, 0x37, 0x58, 0x34, 0x93, 0xc7,
	0x30, 0xc5, 0xd6, 0xc7, 0x60, 0x86, 0x20, 0x7b, 0xb6, 0xd3, 0xda, 0x59, 0xd8, 0xdb, 0xd2, 0x79,
	0xe7, 0x98, 0xb5, 0x9d, 0xef, 0x32, 0xf1, 0xe4, 0xf1, 0x39, 0x4c, 0x0a, 0x1c, 0xcc, 0x10, 0x64,
	0xed, 0x82, 0x39, 0x09, 0x9a, 0x93, 0xee, 0xed, 0x5b, 0xee, 0xa7, 0x82, 0x91, 0x2c, 0x56, 0xfe,
	0xd2, 0xd3, 0x3a, 0x07, 0xf3, 0xea, 0x62, 0xec, 0xb7, 0x3a, 0xad, 0x9d, 0xe5, 0xbd, 0xaf, 0x9c,
	0x49, 0x59, 0x2e, 0x8f, 0xcc, 0xd1, 0xb1, 0x9f, 0x4a, 0xcd, 0xb3, 0xac, 0x48, 0x47, 0x67, 0x02,
	0x4d, 0xb3, 0x8e, 0xc1, 0x5c, 0x79, 0x77, 0xf6, 0xa2, 0xa4, 0x7a, 0xcd, 0xa8, 0x67, 0xd7, 0x39,
	0x1e, 0x66, 0x96, 0xdf, 0x81, 0xe4, 0x58, 0x57, 0x60, 0x65, 0x3c, 0xff, 0xec, 0xfb, 0x92, 0x7d,
	0xdc, 0x8c, 0xed, 0xa3, 0x40, 0x8b, 0x9f, 0x53, 0x34, 0xb2, 0xca, 0xa8, 0x25, 0x58, 0x86, 0x23,
	0xdf, 0xd6, 0xd7, 0x60, 0xa9, 0x07, 0x39, 0x0e, 0x4d, 0x4e, 0xda, 0x56, 0x83, 0xc3, 0x5d, 0x28,
	0x25, 0xe6, 0x9e, 0x4f, 0xc0, 0xba, 0x60, 0x30, 0xfa, 0xa9, 0x4c, 0xb4, 0x82, 0x25, 0xa1, 0xc0,
	0x69, 0x9e, 0x40, 0x81, 0xed, 0xa5, 0x06, 0xa4, 0x35, 0x23, 0x7d, 0xc9, 0x92, 0x33, 0x2d, 0xb4,
	0x22, 0xb0, 0x5e, 0x82, 0xc6, 0x9f, 0x13, 0x6e, 0xcf, 0x77, 0x66, 0x77, 0x16, 0xf6, 0xdc, 0x89,
	0x47, 0xa2, 0x9e, 0x02, 0xe7, 0x40, 0x0a, 0x4f, 0x8c, 0x2e, 0x58, 0x2b, 0x58, 0x32, 0x36, 0xc7,
	0xad, 0xcf, 0xc1, 0x3b, 0x11, 0x4c, 0x73, 0x48, 0xe2, 0xcc, 0x06, 0x0d, 0x22, 0xad, 0xbc, 0x2d,
	0x1f, 0x2c, 0x47, 0x79, 0x14, 0xf6, 0x08, 0x0a, 0x53, 0x12, 0x31, 0xca, 0xed, 0xe5, 0xfa, 0xfc,
	0x5d, 0x8c, 0xf2, 0x68, 0x9f, 0xa0, 0xe7, 0x52, 0xa0, 0x10, 0xe9, 0x30, 0xe2, 0x5e, 0x23, 0x44,
	0x3a, 0x40, 0x1c, 0x81, 0x55, 0x55, 0xe0, 0x61, 0x94, 0x43, 0x43, 0xd9, 0xaa, 0xa7, 0xdc, 0x53,
	0xaa, 0x83, 0x1c, 0x0e, 0xc7, 0x72, 0x39, 0x1c, 0xcb, 0x6a, 0xa3, 0x58, 0x2e, 0xef, 0x8e, 0x25,
	0x35, 0x94, 0xcd, 0xff, 0x10, 0x4b, 0xaa, 0x41, 0x4f, 0xc1, 0x82, 0x06, 0x31, 0x0a, 0xb9, 0xfd,
	0x60, 0xc2, 0xbd, 0x1c, 0xd2, 0xa2, 0x97, 0x60, 0xc5, 0x00, 0x4a, 0x10, 0x50, 0xc8, 0xad, 0x13,
	0x70, 0x3f, 0xc7, 0x2c, 0xc2, 0x59, 0x19, 0xc8, 0xc8, 0x0d, 0xbd, 0x57, 0x1f, 0xcc, 0x9a, 0x96,
	0x1e, 0x0c, 0x5f, 0xd4, 0x6f, 0x2d, 0xf0, 0xb0, 0xa6, 0x3f, 0xd9, 0xeb, 0x92, 0xfd, 0x65, 0x5d,
	0x56, 0x3e, 0xd3, 0x18, 0xbf, 0x10, 0xf4, 0x85, 0x81, 0x9c, 0x2a, 0x46, 0xd0, 0xc6, 0x53, 0xac,
	0xd6, 0xef, 0x2d, 0xd0, 0x46, 0x84, 0xe7, 0x09, 0xbc, 0x36, 0x55, 0x51, 0xee, 0xad, 0x7a, 0xd6,
	0xed, 0xf7, 0xe5, 0x5b, 0xf1, 0x6d, 0xcd, 0x5b, 0x71, 0x66, 0x1a, 0xc2, 0xa1, 0x11, 0xca, 0x87,
	0xe2, 0xf6, 0x74, 0xb0, 0xa1, 0x57, 0x53, 0x55, 0xb3, 0x4f, 0x50, 0x65, 0xb2, 0xbe, 0x01, 0x2b,
	0x7d, 0x92, 0xc1, 0x44, 0xd6, 0x3a, 0x2f, 0xfa, 0x7d, 0x72, 0x65, 0xdb, 0x0d, 0xea, 0x67, 0x59,
	0xaa, 0x5e, 0xb2, 0xe4, 0x54, 0x6a, 0xac, 0x1f, 0x4d, 0xce, 0x0c, 0xb5, 0x6c, 0x7b, 0x43, 0x82,
	0x76, 0xeb, 0x8e, 0xb2, 0x8a, 0xd8, 0x1c, 0xdf, 0x8a, 0x18, 0x9b, 0xb1, 0x5e, 0x81, 0x2d, 0xdc,
	0xef, 0xe3, 0x48, 0x90, 0x4b, 0x1c, 0xde, 0x2e, 0x94, 0x76, 0x7d, 0x3e, 0xd8, 0x95, 0xfe, 0x6c,
	0xac, 0x62, 0x6e, 0x26, 0xb0, 0x55, 0x5f, 0xb3, 0x3f, 0x68, 0xd4, 0x6a, 0xf6, 0x55, 0x47, 0x3f,
	0x95, 0x1a, 0x79, 0x0f, 0x23, 0x33, 0x77, 0x2d, 0xaf, 0x2c, 0xe5, 0x83, 0x7b, 0x6b, 0x79, 0x59,
	0x2e, 0x0f, 0x1b, 0x94, 0xcb, 0xda, 0x18, 0x56, 0xd6, 0xcd, 0xaf, 0xa0, 0x7d, 0x27, 0xd1, 0xec,
	0xa8, 0xf3, 0xbf, 0xec, 0x68, 0xe3, 0x8e, 0xa5, 0xf5, 0x96, 0x1e, 0x83, 0xf9, 0x04, 0xf6, 0x70,
	0xc2, 0xed, 0x6d, 0xf9, 0xc4, 0x4f, 0x4f, 0x25, 0xed, 0xeb, 0xbd, 0x78, 0xe3, 0xff, 0x00, 0x3a,
	0x83, 0x48, 0xf4, 0x28, 0x27, 0xbc, 0x4c, 0x14, 0xd7, 0x34, 0xa8, 0x1d, 0x55, 0x32, 0x98, 0x71,
	0xf7, 0xb5, 0x19, 0xde, 0xb8, 0x50, 0x59, 0xb9, 0xfb, 0xda, 0xb4, 0xbe, 0x9b, 0xfd, 0x7f, 0x5a,
	0xe0, 0x51, 0x44, 0x53, 0xa7, 0xf6, 0x57, 0x78, 0x7f, 0x51, 0xc3, 0x4f, 0xca, 0xf8, 0x4e, 0x5a,
	0xaf, 0xbe, 0xd7, 0x92, 0x98, 0x26, 0x30, 0x8b, 0x1d, 0xca, 0x62, 0x37, 0xc6, 0x99, 0x8c, 0xde,
	0x1d, 0x84, 0x35, 0xe5, 0x77, 0xfc, 0x8b, 0x6a, 0xf4, 0xc7, 0xcc, 0xec, 0x91, 0xef, 0xff, 0x39,
	0xb3, 0x7d, 0xa4, 0x90, 0x3e, 0xe2, 0x8e, 0x1a, 0x96, 0xa3, 0xf3, 0xae, 0x13, 0x18, 0xcf, 0xbf,
	0x8d, 0xcf, 0x85, 0x8f, 0xf8, 0x45, 0xe5, 0x73, 0x71, 0xde, 0xbd, 0xa8, 0x7c, 0xde, 0xcc, 0x3c,
	0x52, 0x06, 0xcf, 0xf3, 0x11, 0xf7, 0xbc, 0xca, 0xcb, 0xf3, 0xce, 0xbb, 0x9e, 0x57, 0xf9, 0xf5,
	0xe6, 0x65, 0xb0, 0xdd, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x0f, 0xe9, 0xf5, 0xae, 0x3a, 0x0c,
	0x00, 0x00,
}
