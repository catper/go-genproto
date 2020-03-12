// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/bidding_strategy.proto

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

// A bidding strategy.
type BiddingStrategy struct {
	// The resource name of the bidding strategy.
	// Bidding strategy resource names have the form:
	//
	// `customers/{customer_id}/biddingStrategies/{bidding_strategy_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the bidding strategy.
	Id *wrappers.Int64Value `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the bidding strategy.
	// All bidding strategies within an account must be named distinctly.
	//
	// The length of this string should be between 1 and 255, inclusive,
	// in UTF-8 bytes, (trimmed).
	Name *wrappers.StringValue `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// The status of the bidding strategy.
	//
	// This field is read-only.
	Status enums.BiddingStrategyStatusEnum_BiddingStrategyStatus `protobuf:"varint,15,opt,name=status,proto3,enum=google.ads.googleads.v3.enums.BiddingStrategyStatusEnum_BiddingStrategyStatus" json:"status,omitempty"`
	// The type of the bidding strategy.
	// Create a bidding strategy by setting the bidding scheme.
	//
	// This field is read-only.
	Type enums.BiddingStrategyTypeEnum_BiddingStrategyType `protobuf:"varint,5,opt,name=type,proto3,enum=google.ads.googleads.v3.enums.BiddingStrategyTypeEnum_BiddingStrategyType" json:"type,omitempty"`
	// The number of campaigns attached to this bidding strategy.
	//
	// This field is read-only.
	CampaignCount *wrappers.Int64Value `protobuf:"bytes,13,opt,name=campaign_count,json=campaignCount,proto3" json:"campaign_count,omitempty"`
	// The number of non-removed campaigns attached to this bidding strategy.
	//
	// This field is read-only.
	NonRemovedCampaignCount *wrappers.Int64Value `protobuf:"bytes,14,opt,name=non_removed_campaign_count,json=nonRemovedCampaignCount,proto3" json:"non_removed_campaign_count,omitempty"`
	// The bidding scheme.
	//
	// Only one can be set.
	//
	// Types that are valid to be assigned to Scheme:
	//	*BiddingStrategy_EnhancedCpc
	//	*BiddingStrategy_TargetCpa
	//	*BiddingStrategy_TargetImpressionShare
	//	*BiddingStrategy_TargetRoas
	//	*BiddingStrategy_TargetSpend
	Scheme               isBiddingStrategy_Scheme `protobuf_oneof:"scheme"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *BiddingStrategy) Reset()         { *m = BiddingStrategy{} }
func (m *BiddingStrategy) String() string { return proto.CompactTextString(m) }
func (*BiddingStrategy) ProtoMessage()    {}
func (*BiddingStrategy) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac3d020c2b1960dc, []int{0}
}

func (m *BiddingStrategy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BiddingStrategy.Unmarshal(m, b)
}
func (m *BiddingStrategy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BiddingStrategy.Marshal(b, m, deterministic)
}
func (m *BiddingStrategy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BiddingStrategy.Merge(m, src)
}
func (m *BiddingStrategy) XXX_Size() int {
	return xxx_messageInfo_BiddingStrategy.Size(m)
}
func (m *BiddingStrategy) XXX_DiscardUnknown() {
	xxx_messageInfo_BiddingStrategy.DiscardUnknown(m)
}

var xxx_messageInfo_BiddingStrategy proto.InternalMessageInfo

func (m *BiddingStrategy) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *BiddingStrategy) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *BiddingStrategy) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *BiddingStrategy) GetStatus() enums.BiddingStrategyStatusEnum_BiddingStrategyStatus {
	if m != nil {
		return m.Status
	}
	return enums.BiddingStrategyStatusEnum_UNSPECIFIED
}

func (m *BiddingStrategy) GetType() enums.BiddingStrategyTypeEnum_BiddingStrategyType {
	if m != nil {
		return m.Type
	}
	return enums.BiddingStrategyTypeEnum_UNSPECIFIED
}

func (m *BiddingStrategy) GetCampaignCount() *wrappers.Int64Value {
	if m != nil {
		return m.CampaignCount
	}
	return nil
}

func (m *BiddingStrategy) GetNonRemovedCampaignCount() *wrappers.Int64Value {
	if m != nil {
		return m.NonRemovedCampaignCount
	}
	return nil
}

type isBiddingStrategy_Scheme interface {
	isBiddingStrategy_Scheme()
}

type BiddingStrategy_EnhancedCpc struct {
	EnhancedCpc *common.EnhancedCpc `protobuf:"bytes,7,opt,name=enhanced_cpc,json=enhancedCpc,proto3,oneof"`
}

type BiddingStrategy_TargetCpa struct {
	TargetCpa *common.TargetCpa `protobuf:"bytes,9,opt,name=target_cpa,json=targetCpa,proto3,oneof"`
}

type BiddingStrategy_TargetImpressionShare struct {
	TargetImpressionShare *common.TargetImpressionShare `protobuf:"bytes,48,opt,name=target_impression_share,json=targetImpressionShare,proto3,oneof"`
}

type BiddingStrategy_TargetRoas struct {
	TargetRoas *common.TargetRoas `protobuf:"bytes,11,opt,name=target_roas,json=targetRoas,proto3,oneof"`
}

type BiddingStrategy_TargetSpend struct {
	TargetSpend *common.TargetSpend `protobuf:"bytes,12,opt,name=target_spend,json=targetSpend,proto3,oneof"`
}

func (*BiddingStrategy_EnhancedCpc) isBiddingStrategy_Scheme() {}

func (*BiddingStrategy_TargetCpa) isBiddingStrategy_Scheme() {}

func (*BiddingStrategy_TargetImpressionShare) isBiddingStrategy_Scheme() {}

func (*BiddingStrategy_TargetRoas) isBiddingStrategy_Scheme() {}

func (*BiddingStrategy_TargetSpend) isBiddingStrategy_Scheme() {}

func (m *BiddingStrategy) GetScheme() isBiddingStrategy_Scheme {
	if m != nil {
		return m.Scheme
	}
	return nil
}

func (m *BiddingStrategy) GetEnhancedCpc() *common.EnhancedCpc {
	if x, ok := m.GetScheme().(*BiddingStrategy_EnhancedCpc); ok {
		return x.EnhancedCpc
	}
	return nil
}

func (m *BiddingStrategy) GetTargetCpa() *common.TargetCpa {
	if x, ok := m.GetScheme().(*BiddingStrategy_TargetCpa); ok {
		return x.TargetCpa
	}
	return nil
}

func (m *BiddingStrategy) GetTargetImpressionShare() *common.TargetImpressionShare {
	if x, ok := m.GetScheme().(*BiddingStrategy_TargetImpressionShare); ok {
		return x.TargetImpressionShare
	}
	return nil
}

func (m *BiddingStrategy) GetTargetRoas() *common.TargetRoas {
	if x, ok := m.GetScheme().(*BiddingStrategy_TargetRoas); ok {
		return x.TargetRoas
	}
	return nil
}

func (m *BiddingStrategy) GetTargetSpend() *common.TargetSpend {
	if x, ok := m.GetScheme().(*BiddingStrategy_TargetSpend); ok {
		return x.TargetSpend
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*BiddingStrategy) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*BiddingStrategy_EnhancedCpc)(nil),
		(*BiddingStrategy_TargetCpa)(nil),
		(*BiddingStrategy_TargetImpressionShare)(nil),
		(*BiddingStrategy_TargetRoas)(nil),
		(*BiddingStrategy_TargetSpend)(nil),
	}
}

func init() {
	proto.RegisterType((*BiddingStrategy)(nil), "google.ads.googleads.v3.resources.BiddingStrategy")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/bidding_strategy.proto", fileDescriptor_ac3d020c2b1960dc)
}

var fileDescriptor_ac3d020c2b1960dc = []byte{
	// 676 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xdf, 0x4e, 0xdb, 0x30,
	0x14, 0xc6, 0xdb, 0xc2, 0x18, 0xb8, 0xfc, 0x91, 0xa2, 0x4d, 0x64, 0x0c, 0x4d, 0xb0, 0x09, 0xa9,
	0x1b, 0x93, 0x83, 0xe8, 0x36, 0x8d, 0x70, 0x95, 0x56, 0x88, 0x82, 0x34, 0x84, 0x52, 0x54, 0x4d,
	0x53, 0xb5, 0xc8, 0x24, 0x26, 0x8d, 0x44, 0x6c, 0xcb, 0x76, 0x98, 0x10, 0xe2, 0x62, 0x4f, 0x32,
	0x69, 0x97, 0x7b, 0x94, 0x3d, 0x0a, 0x4f, 0x31, 0xd9, 0x71, 0xd2, 0xad, 0xd0, 0x15, 0xee, 0x8e,
	0x73, 0xbe, 0xef, 0xf7, 0xd9, 0x27, 0x89, 0xc1, 0xc7, 0x98, 0xd2, 0xf8, 0x1c, 0x3b, 0x28, 0x12,
	0x4e, 0x5e, 0xaa, 0xea, 0xa2, 0xe9, 0x70, 0x2c, 0x68, 0xc6, 0x43, 0x2c, 0x9c, 0xd3, 0x24, 0x8a,
	0x12, 0x12, 0x07, 0x42, 0x72, 0x24, 0x71, 0x7c, 0x09, 0x19, 0xa7, 0x92, 0x5a, 0xeb, 0xb9, 0x1c,
	0xa2, 0x48, 0xc0, 0xd2, 0x09, 0x2f, 0x9a, 0xb0, 0x74, 0xae, 0xbc, 0x1d, 0x07, 0x0f, 0x69, 0x9a,
	0x52, 0x52, 0x90, 0x73, 0xe0, 0xca, 0xee, 0x38, 0x35, 0x26, 0x59, 0x7a, 0x7b, 0x1b, 0x81, 0x90,
	0x48, 0x66, 0xc2, 0x98, 0x77, 0x1e, 0x68, 0x96, 0x97, 0x0c, 0x1b, 0xeb, 0xb3, 0xc2, 0xca, 0x92,
	0xf2, 0xd4, 0xa6, 0xf5, 0xc2, 0xb4, 0xf4, 0xea, 0x34, 0x3b, 0x73, 0xbe, 0x71, 0xc4, 0x18, 0xe6,
	0x45, 0xea, 0xea, 0x5f, 0x56, 0x44, 0x08, 0x95, 0x48, 0x26, 0x94, 0x98, 0xee, 0xcb, 0x1f, 0xb3,
	0x60, 0xa9, 0x95, 0x07, 0x77, 0x4d, 0xae, 0xf5, 0x0a, 0x2c, 0x14, 0x19, 0x01, 0x41, 0x29, 0xb6,
	0xab, 0x6b, 0xd5, 0xc6, 0x9c, 0x3f, 0x5f, 0x3c, 0x3c, 0x42, 0x29, 0xb6, 0x36, 0x41, 0x2d, 0x89,
	0xec, 0xa9, 0xb5, 0x6a, 0xa3, 0xbe, 0xfd, 0xdc, 0x0c, 0x17, 0x16, 0x7b, 0x80, 0x07, 0x44, 0x7e,
	0x78, 0xd7, 0x43, 0xe7, 0x19, 0xf6, 0x6b, 0x49, 0x64, 0x6d, 0x81, 0x69, 0x0d, 0x9a, 0xd6, 0xf2,
	0xd5, 0x5b, 0xf2, 0xae, 0xe4, 0x09, 0x89, 0x73, 0xbd, 0x56, 0x5a, 0x67, 0x60, 0x26, 0x9f, 0x9d,
	0xbd, 0xb4, 0x56, 0x6d, 0x2c, 0x6e, 0x1f, 0xc1, 0x71, 0xaf, 0x52, 0x0f, 0x0f, 0x8e, 0x9c, 0xa1,
	0xab, 0xbd, 0x7b, 0x24, 0x4b, 0xef, 0xee, 0xf8, 0x86, 0x6e, 0x7d, 0x05, 0xd3, 0x6a, 0xcc, 0xf6,
	0x23, 0x9d, 0x72, 0xf8, 0xb0, 0x94, 0x93, 0x4b, 0x86, 0xef, 0xca, 0x50, 0xcf, 0x7d, 0xcd, 0xb5,
	0x5a, 0x60, 0x31, 0x44, 0x29, 0x43, 0x49, 0x4c, 0x82, 0x90, 0x66, 0x44, 0xda, 0x0b, 0x93, 0x47,
	0xb6, 0x50, 0x58, 0xda, 0xca, 0x61, 0x7d, 0x06, 0x2b, 0x84, 0x92, 0x80, 0xe3, 0x94, 0x5e, 0xe0,
	0x28, 0x18, 0xe1, 0x2d, 0x4e, 0xe6, 0x2d, 0x13, 0x4a, 0xfc, 0xdc, 0xdd, 0xfe, 0x87, 0x7c, 0x0c,
	0xe6, 0x31, 0x19, 0x20, 0x12, 0x2a, 0x2c, 0x0b, 0xed, 0xc7, 0x9a, 0xb5, 0x39, 0x76, 0x0a, 0xf9,
	0x3f, 0x01, 0xf7, 0x8c, 0xa7, 0xcd, 0xc2, 0x4e, 0xc5, 0xaf, 0xe3, 0xe1, 0xd2, 0x3a, 0x04, 0x40,
	0x22, 0x1e, 0x63, 0x19, 0x84, 0x0c, 0xd9, 0x73, 0x9a, 0xf7, 0x7a, 0x12, 0xef, 0x44, 0x3b, 0xda,
	0x0c, 0x75, 0x2a, 0xfe, 0x9c, 0x2c, 0x16, 0x16, 0x05, 0xcb, 0x86, 0x95, 0xa4, 0x8c, 0x63, 0x21,
	0x12, 0x4a, 0x02, 0x31, 0x40, 0x1c, 0xdb, 0x5b, 0x1a, 0xfc, 0xfe, 0x7e, 0xe0, 0x83, 0xd2, 0xdd,
	0x55, 0xe6, 0x4e, 0xc5, 0x7f, 0x2a, 0xef, 0x6a, 0x58, 0x9f, 0x40, 0xdd, 0x04, 0x72, 0x8a, 0x84,
	0x5d, 0xd7, 0x21, 0x6f, 0xee, 0x17, 0xe2, 0x53, 0x24, 0x3a, 0x15, 0xdf, 0x9c, 0x5e, 0xad, 0xd4,
	0x74, 0x0d, 0x4e, 0x30, 0x4c, 0x22, 0x7b, 0xfe, 0x7e, 0xd3, 0xcd, 0x79, 0x5d, 0x65, 0x51, 0xd3,
	0x95, 0xc3, 0xa5, 0x3b, 0xb8, 0xf1, 0x30, 0x68, 0x0c, 0x4d, 0xa6, 0x62, 0x89, 0x50, 0x66, 0x67,
	0xf4, 0x47, 0xde, 0x09, 0x33, 0x21, 0x69, 0x8a, 0xb9, 0x70, 0xae, 0x8a, 0xf2, 0xba, 0xb8, 0x67,
	0x8c, 0x2a, 0xc1, 0xc2, 0xb9, 0x1a, 0xbd, 0x7a, 0xae, 0x5b, 0xb3, 0x60, 0x46, 0x84, 0x03, 0x9c,
	0xe2, 0xd6, 0xf7, 0x1a, 0xd8, 0x08, 0x69, 0x0a, 0x27, 0x5e, 0xa5, 0xad, 0x27, 0x23, 0xf9, 0xc7,
	0xea, 0x53, 0x3c, 0xae, 0x7e, 0x39, 0x34, 0xd6, 0x98, 0x9e, 0x23, 0x12, 0x43, 0xca, 0x63, 0x27,
	0xc6, 0x44, 0x7f, 0xa8, 0xce, 0xf0, 0x04, 0xff, 0xb9, 0xde, 0x77, 0xcb, 0xea, 0x67, 0x6d, 0x6a,
	0xdf, 0xf3, 0x7e, 0xd5, 0xd6, 0xf7, 0x73, 0xa4, 0x17, 0x09, 0x98, 0x97, 0xaa, 0xea, 0x35, 0xa1,
	0x5f, 0x28, 0x7f, 0x17, 0x9a, 0xbe, 0x17, 0x89, 0x7e, 0xa9, 0xe9, 0xf7, 0x9a, 0xfd, 0x52, 0x73,
	0x53, 0xdb, 0xc8, 0x1b, 0xae, 0xeb, 0x45, 0xc2, 0x75, 0x4b, 0x95, 0xeb, 0xf6, 0x9a, 0xae, 0x5b,
	0xea, 0x4e, 0x67, 0xf4, 0x66, 0x9b, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x1b, 0x9e, 0x85, 0x4c,
	0x8a, 0x06, 0x00, 0x00,
}
