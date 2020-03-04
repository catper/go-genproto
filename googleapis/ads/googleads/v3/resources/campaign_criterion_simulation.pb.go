// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/campaign_criterion_simulation.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// A campaign criterion simulation. Supported combinations of advertising
// channel type, criterion ids, simulation type and simulation modification
// method is detailed below respectively.
//
// 1. SEARCH - 30000,30001,30002 - BID_MODIFIER - UNIFORM
// 2. SHOPPING - 30000,30001,30002 - BID_MODIFIER - UNIFORM
// 3. DISPLAY - 30001 - BID_MODIFIER - UNIFORM
type CampaignCriterionSimulation struct {
	// The resource name of the campaign criterion simulation.
	// Campaign criterion simulation resource names have the form:
	//
	// `customers/{customer_id}/campaignCriterionSimulations/{campaign_id}~{criterion_id}~{type}~{modification_method}~{start_date}~{end_date}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Campaign ID of the simulation.
	CampaignId *wrappers.Int64Value `protobuf:"bytes,2,opt,name=campaign_id,json=campaignId,proto3" json:"campaign_id,omitempty"`
	// Criterion ID of the simulation.
	CriterionId *wrappers.Int64Value `protobuf:"bytes,3,opt,name=criterion_id,json=criterionId,proto3" json:"criterion_id,omitempty"`
	// The field that the simulation modifies.
	Type enums.SimulationTypeEnum_SimulationType `protobuf:"varint,4,opt,name=type,proto3,enum=google.ads.googleads.v3.enums.SimulationTypeEnum_SimulationType" json:"type,omitempty"`
	// How the simulation modifies the field.
	ModificationMethod enums.SimulationModificationMethodEnum_SimulationModificationMethod `protobuf:"varint,5,opt,name=modification_method,json=modificationMethod,proto3,enum=google.ads.googleads.v3.enums.SimulationModificationMethodEnum_SimulationModificationMethod" json:"modification_method,omitempty"`
	// First day on which the simulation is based, in YYYY-MM-DD format.
	StartDate *wrappers.StringValue `protobuf:"bytes,6,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	// Last day on which the simulation is based, in YYYY-MM-DD format.
	EndDate *wrappers.StringValue `protobuf:"bytes,7,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	// List of simulation points.
	//
	// Types that are valid to be assigned to PointList:
	//	*CampaignCriterionSimulation_BidModifierPointList
	PointList            isCampaignCriterionSimulation_PointList `protobuf_oneof:"point_list"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *CampaignCriterionSimulation) Reset()         { *m = CampaignCriterionSimulation{} }
func (m *CampaignCriterionSimulation) String() string { return proto.CompactTextString(m) }
func (*CampaignCriterionSimulation) ProtoMessage()    {}
func (*CampaignCriterionSimulation) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4b0653a9a2bb524, []int{0}
}

func (m *CampaignCriterionSimulation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignCriterionSimulation.Unmarshal(m, b)
}
func (m *CampaignCriterionSimulation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignCriterionSimulation.Marshal(b, m, deterministic)
}
func (m *CampaignCriterionSimulation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignCriterionSimulation.Merge(m, src)
}
func (m *CampaignCriterionSimulation) XXX_Size() int {
	return xxx_messageInfo_CampaignCriterionSimulation.Size(m)
}
func (m *CampaignCriterionSimulation) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignCriterionSimulation.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignCriterionSimulation proto.InternalMessageInfo

func (m *CampaignCriterionSimulation) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CampaignCriterionSimulation) GetCampaignId() *wrappers.Int64Value {
	if m != nil {
		return m.CampaignId
	}
	return nil
}

func (m *CampaignCriterionSimulation) GetCriterionId() *wrappers.Int64Value {
	if m != nil {
		return m.CriterionId
	}
	return nil
}

func (m *CampaignCriterionSimulation) GetType() enums.SimulationTypeEnum_SimulationType {
	if m != nil {
		return m.Type
	}
	return enums.SimulationTypeEnum_UNSPECIFIED
}

func (m *CampaignCriterionSimulation) GetModificationMethod() enums.SimulationModificationMethodEnum_SimulationModificationMethod {
	if m != nil {
		return m.ModificationMethod
	}
	return enums.SimulationModificationMethodEnum_UNSPECIFIED
}

func (m *CampaignCriterionSimulation) GetStartDate() *wrappers.StringValue {
	if m != nil {
		return m.StartDate
	}
	return nil
}

func (m *CampaignCriterionSimulation) GetEndDate() *wrappers.StringValue {
	if m != nil {
		return m.EndDate
	}
	return nil
}

type isCampaignCriterionSimulation_PointList interface {
	isCampaignCriterionSimulation_PointList()
}

type CampaignCriterionSimulation_BidModifierPointList struct {
	BidModifierPointList *common.BidModifierSimulationPointList `protobuf:"bytes,8,opt,name=bid_modifier_point_list,json=bidModifierPointList,proto3,oneof"`
}

func (*CampaignCriterionSimulation_BidModifierPointList) isCampaignCriterionSimulation_PointList() {}

func (m *CampaignCriterionSimulation) GetPointList() isCampaignCriterionSimulation_PointList {
	if m != nil {
		return m.PointList
	}
	return nil
}

func (m *CampaignCriterionSimulation) GetBidModifierPointList() *common.BidModifierSimulationPointList {
	if x, ok := m.GetPointList().(*CampaignCriterionSimulation_BidModifierPointList); ok {
		return x.BidModifierPointList
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CampaignCriterionSimulation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CampaignCriterionSimulation_BidModifierPointList)(nil),
	}
}

func init() {
	proto.RegisterType((*CampaignCriterionSimulation)(nil), "google.ads.googleads.v3.resources.CampaignCriterionSimulation")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/campaign_criterion_simulation.proto", fileDescriptor_d4b0653a9a2bb524)
}

var fileDescriptor_d4b0653a9a2bb524 = []byte{
	// 605 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdd, 0x4e, 0x14, 0x31,
	0x14, 0x76, 0x16, 0xe4, 0xa7, 0xa0, 0x17, 0xa3, 0x89, 0x23, 0x10, 0xb3, 0x68, 0x48, 0xb8, 0xea,
	0x24, 0x0c, 0xd1, 0x64, 0x30, 0xc4, 0x5d, 0x24, 0x88, 0x11, 0x83, 0x0b, 0xd9, 0x0b, 0xb3, 0xc9,
	0xa4, 0x3b, 0x2d, 0x63, 0x93, 0x6d, 0x3b, 0x69, 0x3b, 0x10, 0x62, 0x78, 0x00, 0x2f, 0x78, 0x09,
	0x2f, 0xbd, 0xf5, 0x2d, 0x7c, 0x14, 0x9e, 0xc2, 0x4c, 0xa7, 0xd3, 0x21, 0xc0, 0x2e, 0xdc, 0x9d,
	0xe9, 0xf9, 0xbe, 0xef, 0xf4, 0x7c, 0xa7, 0x73, 0xc0, 0x6e, 0x26, 0x44, 0x36, 0x22, 0x21, 0xc2,
	0x2a, 0xac, 0xc2, 0x32, 0x3a, 0x8d, 0x42, 0x49, 0x94, 0x28, 0x64, 0x4a, 0x54, 0x98, 0x22, 0x96,
	0x23, 0x9a, 0xf1, 0x24, 0x95, 0x54, 0x13, 0x49, 0x05, 0x4f, 0x14, 0x65, 0xc5, 0x08, 0x69, 0x2a,
	0x38, 0xcc, 0xa5, 0xd0, 0xc2, 0x5f, 0xad, 0xb8, 0x10, 0x61, 0x05, 0x9d, 0x0c, 0x3c, 0x8d, 0xa0,
	0x93, 0x59, 0x0a, 0xc7, 0x55, 0x4a, 0x05, 0x63, 0x82, 0x87, 0x37, 0x35, 0x97, 0xba, 0xe3, 0x08,
	0x84, 0x17, 0x4c, 0x5d, 0xc3, 0x27, 0x4c, 0x60, 0x7a, 0x42, 0x53, 0xfb, 0x41, 0xf4, 0x0f, 0x81,
	0xad, 0x46, 0xf4, 0x60, 0x0d, 0x7d, 0x9e, 0x13, 0x4b, 0x7a, 0x59, 0x93, 0x72, 0xea, 0x6c, 0xb0,
	0xa9, 0x57, 0x36, 0x65, 0xbe, 0x86, 0xc5, 0x49, 0x78, 0x26, 0x51, 0x9e, 0x13, 0xa9, 0x6c, 0x7e,
	0xe5, 0x1a, 0x15, 0x71, 0x2e, 0xb4, 0x11, 0xb7, 0xd9, 0xd7, 0x7f, 0x67, 0xc0, 0xf2, 0x8e, 0x75,
	0x73, 0xa7, 0x36, 0xf3, 0xc8, 0xdd, 0xc1, 0x7f, 0x03, 0x9e, 0xd4, 0xf5, 0x12, 0x8e, 0x18, 0x09,
	0xbc, 0xb6, 0xb7, 0x3e, 0xdf, 0x5b, 0xac, 0x0f, 0xbf, 0x22, 0x46, 0xfc, 0xf7, 0x60, 0xc1, 0x4d,
	0x84, 0xe2, 0xa0, 0xd5, 0xf6, 0xd6, 0x17, 0x36, 0x96, 0xad, 0xeb, 0xb0, 0xbe, 0x18, 0xdc, 0xe7,
	0xfa, 0xed, 0x66, 0x1f, 0x8d, 0x0a, 0xd2, 0x03, 0x35, 0x7e, 0x1f, 0xfb, 0xdb, 0x60, 0xb1, 0x19,
	0x23, 0xc5, 0xc1, 0xd4, 0xfd, 0xf4, 0x05, 0x47, 0xd8, 0xc7, 0xfe, 0x31, 0x98, 0x2e, 0x9d, 0x0a,
	0xa6, 0xdb, 0xde, 0xfa, 0xd3, 0x8d, 0x0f, 0x70, 0xdc, 0xdc, 0x8d, 0xbf, 0xb0, 0xe9, 0xed, 0xf8,
	0x3c, 0x27, 0xbb, 0xbc, 0x60, 0x37, 0x8e, 0x7a, 0x46, 0xcd, 0xbf, 0xf4, 0xc0, 0xb3, 0x3b, 0x86,
	0x18, 0x3c, 0x36, 0x55, 0x06, 0x0f, 0xae, 0x72, 0x70, 0x4d, 0xe3, 0xc0, 0x48, 0xdc, 0xa8, 0x79,
	0x1b, 0xd0, 0xf3, 0xd9, 0xad, 0x33, 0x7f, 0x0b, 0x00, 0xa5, 0x91, 0xd4, 0x09, 0x46, 0x9a, 0x04,
	0x33, 0xc6, 0xa3, 0x95, 0x5b, 0x1e, 0x1d, 0x69, 0x49, 0x79, 0x56, 0x99, 0x34, 0x6f, 0xf0, 0x1f,
	0x91, 0x26, 0xfe, 0x3b, 0x30, 0x47, 0x38, 0xae, 0xa8, 0xb3, 0x0f, 0xa0, 0xce, 0x12, 0x8e, 0x0d,
	0xf1, 0x0c, 0xbc, 0x18, 0x52, 0x6c, 0x5f, 0x33, 0x91, 0x49, 0x2e, 0x28, 0xd7, 0xc9, 0x88, 0x2a,
	0x1d, 0xcc, 0x19, 0x9d, 0xed, 0xb1, 0x46, 0x54, 0xff, 0x10, 0xec, 0x52, 0x7c, 0x60, 0xd9, 0x4d,
	0xcf, 0x87, 0xa5, 0xcc, 0x17, 0xaa, 0xf4, 0xa7, 0x47, 0xbd, 0xe7, 0xc3, 0x06, 0xe1, 0xce, 0xe3,
	0x4b, 0xef, 0xaa, 0xf3, 0xcb, 0x03, 0x9b, 0x8d, 0xa4, 0x8d, 0x72, 0xaa, 0x4a, 0xe9, 0x70, 0xd2,
	0xa3, 0xfd, 0x96, 0x16, 0x4a, 0x0b, 0x46, 0xa4, 0x0a, 0x7f, 0xd6, 0xe1, 0x85, 0x5b, 0x1a, 0x77,
	0x30, 0x4a, 0xdc, 0xa4, 0x95, 0x72, 0xd1, 0x5d, 0x04, 0xa0, 0xe9, 0xbd, 0x7b, 0xd9, 0x02, 0x6b,
	0xa9, 0x60, 0xf0, 0xde, 0x15, 0xd3, 0x6d, 0x4f, 0xb8, 0xe7, 0x61, 0x69, 0xfd, 0xa1, 0xf7, 0xfd,
	0xb3, 0x95, 0xc9, 0xc4, 0x08, 0xf1, 0x0c, 0x0a, 0x99, 0x85, 0x19, 0xe1, 0x66, 0x30, 0x61, 0xd3,
	0xf5, 0x84, 0x7d, 0xb8, 0xe5, 0xa2, 0xdf, 0xad, 0xa9, 0xbd, 0x4e, 0xe7, 0x4f, 0x6b, 0x75, 0xaf,
	0x92, 0xec, 0x60, 0x05, 0xab, 0xb0, 0x8c, 0xfa, 0x11, 0xec, 0xd5, 0xc8, 0x7f, 0x35, 0x66, 0xd0,
	0xc1, 0x6a, 0xe0, 0x30, 0x83, 0x7e, 0x34, 0x70, 0x98, 0xab, 0xd6, 0x5a, 0x95, 0x88, 0xe3, 0x0e,
	0x56, 0x71, 0xec, 0x50, 0x71, 0xdc, 0x8f, 0xe2, 0xd8, 0xe1, 0x86, 0x33, 0xe6, 0xb2, 0xd1, 0xff,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x36, 0xe3, 0x8f, 0xbb, 0x05, 0x00, 0x00,
}