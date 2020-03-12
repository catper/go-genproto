// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/ad_group_simulation.proto

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

// An ad group simulation. Supported combinations of advertising
// channel type, simulation type and simulation modification method is
// detailed below respectively.
//
// 1. SEARCH - CPC_BID - DEFAULT
// 2. SEARCH - CPC_BID - UNIFORM
// 3. SEARCH - TARGET_CPA - UNIFORM
// 4. DISPLAY - CPC_BID - DEFAULT
// 5. DISPLAY - CPC_BID - UNIFORM
// 6. DISPLAY - TARGET_CPA - UNIFORM
// 7. VIDEO - CPV_BID - DEFAULT
// 8. VIDEO - CPV_BID - UNIFORM
type AdGroupSimulation struct {
	// The resource name of the ad group simulation.
	// Ad group simulation resource names have the form:
	//
	// `customers/{customer_id}/adGroupSimulations/{ad_group_id}~{type}~{modification_method}~{start_date}~{end_date}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Ad group id of the simulation.
	AdGroupId *wrappers.Int64Value `protobuf:"bytes,2,opt,name=ad_group_id,json=adGroupId,proto3" json:"ad_group_id,omitempty"`
	// The field that the simulation modifies.
	Type enums.SimulationTypeEnum_SimulationType `protobuf:"varint,3,opt,name=type,proto3,enum=google.ads.googleads.v3.enums.SimulationTypeEnum_SimulationType" json:"type,omitempty"`
	// How the simulation modifies the field.
	ModificationMethod enums.SimulationModificationMethodEnum_SimulationModificationMethod `protobuf:"varint,4,opt,name=modification_method,json=modificationMethod,proto3,enum=google.ads.googleads.v3.enums.SimulationModificationMethodEnum_SimulationModificationMethod" json:"modification_method,omitempty"`
	// First day on which the simulation is based, in YYYY-MM-DD format.
	StartDate *wrappers.StringValue `protobuf:"bytes,5,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	// Last day on which the simulation is based, in YYYY-MM-DD format
	EndDate *wrappers.StringValue `protobuf:"bytes,6,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	// List of simulation points.
	//
	// Types that are valid to be assigned to PointList:
	//	*AdGroupSimulation_CpcBidPointList
	//	*AdGroupSimulation_CpvBidPointList
	//	*AdGroupSimulation_TargetCpaPointList
	PointList            isAdGroupSimulation_PointList `protobuf_oneof:"point_list"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *AdGroupSimulation) Reset()         { *m = AdGroupSimulation{} }
func (m *AdGroupSimulation) String() string { return proto.CompactTextString(m) }
func (*AdGroupSimulation) ProtoMessage()    {}
func (*AdGroupSimulation) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfe8b62d7d533dcc, []int{0}
}

func (m *AdGroupSimulation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupSimulation.Unmarshal(m, b)
}
func (m *AdGroupSimulation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupSimulation.Marshal(b, m, deterministic)
}
func (m *AdGroupSimulation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupSimulation.Merge(m, src)
}
func (m *AdGroupSimulation) XXX_Size() int {
	return xxx_messageInfo_AdGroupSimulation.Size(m)
}
func (m *AdGroupSimulation) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupSimulation.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupSimulation proto.InternalMessageInfo

func (m *AdGroupSimulation) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *AdGroupSimulation) GetAdGroupId() *wrappers.Int64Value {
	if m != nil {
		return m.AdGroupId
	}
	return nil
}

func (m *AdGroupSimulation) GetType() enums.SimulationTypeEnum_SimulationType {
	if m != nil {
		return m.Type
	}
	return enums.SimulationTypeEnum_UNSPECIFIED
}

func (m *AdGroupSimulation) GetModificationMethod() enums.SimulationModificationMethodEnum_SimulationModificationMethod {
	if m != nil {
		return m.ModificationMethod
	}
	return enums.SimulationModificationMethodEnum_UNSPECIFIED
}

func (m *AdGroupSimulation) GetStartDate() *wrappers.StringValue {
	if m != nil {
		return m.StartDate
	}
	return nil
}

func (m *AdGroupSimulation) GetEndDate() *wrappers.StringValue {
	if m != nil {
		return m.EndDate
	}
	return nil
}

type isAdGroupSimulation_PointList interface {
	isAdGroupSimulation_PointList()
}

type AdGroupSimulation_CpcBidPointList struct {
	CpcBidPointList *common.CpcBidSimulationPointList `protobuf:"bytes,8,opt,name=cpc_bid_point_list,json=cpcBidPointList,proto3,oneof"`
}

type AdGroupSimulation_CpvBidPointList struct {
	CpvBidPointList *common.CpvBidSimulationPointList `protobuf:"bytes,10,opt,name=cpv_bid_point_list,json=cpvBidPointList,proto3,oneof"`
}

type AdGroupSimulation_TargetCpaPointList struct {
	TargetCpaPointList *common.TargetCpaSimulationPointList `protobuf:"bytes,9,opt,name=target_cpa_point_list,json=targetCpaPointList,proto3,oneof"`
}

func (*AdGroupSimulation_CpcBidPointList) isAdGroupSimulation_PointList() {}

func (*AdGroupSimulation_CpvBidPointList) isAdGroupSimulation_PointList() {}

func (*AdGroupSimulation_TargetCpaPointList) isAdGroupSimulation_PointList() {}

func (m *AdGroupSimulation) GetPointList() isAdGroupSimulation_PointList {
	if m != nil {
		return m.PointList
	}
	return nil
}

func (m *AdGroupSimulation) GetCpcBidPointList() *common.CpcBidSimulationPointList {
	if x, ok := m.GetPointList().(*AdGroupSimulation_CpcBidPointList); ok {
		return x.CpcBidPointList
	}
	return nil
}

func (m *AdGroupSimulation) GetCpvBidPointList() *common.CpvBidSimulationPointList {
	if x, ok := m.GetPointList().(*AdGroupSimulation_CpvBidPointList); ok {
		return x.CpvBidPointList
	}
	return nil
}

func (m *AdGroupSimulation) GetTargetCpaPointList() *common.TargetCpaSimulationPointList {
	if x, ok := m.GetPointList().(*AdGroupSimulation_TargetCpaPointList); ok {
		return x.TargetCpaPointList
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AdGroupSimulation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AdGroupSimulation_CpcBidPointList)(nil),
		(*AdGroupSimulation_CpvBidPointList)(nil),
		(*AdGroupSimulation_TargetCpaPointList)(nil),
	}
}

func init() {
	proto.RegisterType((*AdGroupSimulation)(nil), "google.ads.googleads.v3.resources.AdGroupSimulation")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/ad_group_simulation.proto", fileDescriptor_bfe8b62d7d533dcc)
}

var fileDescriptor_bfe8b62d7d533dcc = []byte{
	// 620 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0xa6, 0xdd, 0xd8, 0x8f, 0x37, 0x40, 0x04, 0x81, 0xc2, 0x98, 0xd0, 0x06, 0x9a, 0x34, 0x71,
	0x61, 0x4b, 0x0b, 0x02, 0x91, 0x82, 0x44, 0x3a, 0xd0, 0x18, 0x62, 0x68, 0xca, 0xa6, 0x5e, 0xa0,
	0x4a, 0x91, 0x17, 0x7b, 0x99, 0x45, 0x63, 0x9b, 0xd8, 0x29, 0x9a, 0xa6, 0x5d, 0x71, 0xcd, 0x2d,
	0x0f, 0xc0, 0x25, 0x8f, 0xc2, 0xa3, 0xec, 0x29, 0x50, 0x9c, 0xc4, 0x2d, 0xed, 0xba, 0xf5, 0xee,
	0xe4, 0x9c, 0xef, 0xc7, 0xe7, 0xab, 0x6b, 0xd0, 0x4a, 0x84, 0x48, 0x7a, 0x14, 0x61, 0xa2, 0x50,
	0x59, 0x16, 0x55, 0xdf, 0x43, 0x19, 0x55, 0x22, 0xcf, 0x62, 0xaa, 0x10, 0x26, 0x51, 0x92, 0x89,
	0x5c, 0x46, 0x8a, 0xa5, 0x79, 0x0f, 0x6b, 0x26, 0x38, 0x94, 0x99, 0xd0, 0xc2, 0x59, 0x2f, 0x19,
	0x10, 0x13, 0x05, 0x2d, 0x19, 0xf6, 0x3d, 0x68, 0xc9, 0x2b, 0x68, 0x92, 0x7e, 0x2c, 0xd2, 0x54,
	0x70, 0x34, 0xaa, 0xb9, 0xd2, 0x9e, 0x44, 0xa0, 0x3c, 0x4f, 0xd5, 0x10, 0x3e, 0x4a, 0x05, 0x61,
	0xc7, 0x2c, 0xae, 0x3e, 0xa8, 0x3e, 0x11, 0xa4, 0xd2, 0xf0, 0xa6, 0xd6, 0xd0, 0xa7, 0x92, 0x56,
	0xa4, 0x87, 0x35, 0x49, 0x32, 0xbb, 0x7c, 0x35, 0x7a, 0x5c, 0x8d, 0xcc, 0xd7, 0x51, 0x7e, 0x8c,
	0xbe, 0x67, 0x58, 0x4a, 0x9a, 0xa9, 0x6a, 0xbe, 0x3a, 0x44, 0xc5, 0x9c, 0x0b, 0x6d, 0xc4, 0xab,
	0xe9, 0x93, 0x5f, 0xf3, 0xe0, 0x6e, 0x40, 0x76, 0x8a, 0x08, 0x0f, 0xac, 0xb3, 0xf3, 0x14, 0xdc,
	0xaa, 0x5d, 0x22, 0x8e, 0x53, 0xea, 0x36, 0xd6, 0x1a, 0x9b, 0x8b, 0xe1, 0x72, 0xdd, 0xfc, 0x8c,
	0x53, 0xea, 0xb4, 0xc0, 0x92, 0x4d, 0x9f, 0x11, 0xb7, 0xb9, 0xd6, 0xd8, 0x5c, 0xda, 0x7a, 0x54,
	0x65, 0x0d, 0xeb, 0xe3, 0xc0, 0x5d, 0xae, 0x5f, 0x3c, 0xef, 0xe0, 0x5e, 0x4e, 0xc3, 0x45, 0x5c,
	0x3a, 0xed, 0x12, 0xe7, 0x10, 0xcc, 0x16, 0xeb, 0xb9, 0x33, 0x6b, 0x8d, 0xcd, 0xdb, 0x5b, 0x6f,
	0xe1, 0xa4, 0x1f, 0xcb, 0x84, 0x02, 0x07, 0x47, 0x3b, 0x3c, 0x95, 0xf4, 0x3d, 0xcf, 0xd3, 0x91,
	0x56, 0x68, 0xd4, 0x9c, 0x9f, 0x0d, 0x70, 0xef, 0x92, 0xe4, 0xdd, 0x59, 0xe3, 0xd2, 0x9d, 0xda,
	0x65, 0x6f, 0x48, 0x63, 0xcf, 0x48, 0x8c, 0x78, 0x8e, 0x03, 0x42, 0x27, 0x1d, 0xeb, 0x39, 0x2d,
	0x00, 0x94, 0xc6, 0x99, 0x8e, 0x08, 0xd6, 0xd4, 0xbd, 0x69, 0x12, 0x5a, 0x1d, 0x4b, 0xe8, 0x40,
	0x67, 0x8c, 0x27, 0x55, 0x44, 0x06, 0xff, 0x0e, 0x6b, 0xea, 0xbc, 0x04, 0x0b, 0x94, 0x93, 0x92,
	0x3a, 0x37, 0x05, 0x75, 0x9e, 0x72, 0x62, 0x88, 0x27, 0xc0, 0x89, 0x65, 0x1c, 0x1d, 0x31, 0x12,
	0x49, 0xc1, 0xb8, 0x8e, 0x7a, 0x4c, 0x69, 0x77, 0xc1, 0x48, 0xbc, 0x9a, 0x98, 0x41, 0x79, 0xe7,
	0xe1, 0xb6, 0x8c, 0xdb, 0x8c, 0x0c, 0x36, 0xdd, 0x2f, 0x14, 0x3e, 0x31, 0xa5, 0x3f, 0xdc, 0x08,
	0xef, 0xc4, 0x66, 0x68, 0x5b, 0xa5, 0x53, 0x7f, 0xd4, 0x09, 0x4c, 0xeb, 0xd4, 0xbf, 0xca, 0xa9,
	0xff, 0x9f, 0xd3, 0x37, 0x70, 0x5f, 0xe3, 0x2c, 0xa1, 0x3a, 0x8a, 0x25, 0x1e, 0x36, 0x5b, 0x34,
	0x66, 0xaf, 0xaf, 0x33, 0x3b, 0x34, 0xe4, 0x6d, 0x89, 0x2f, 0xf7, 0x73, 0x74, 0x3d, 0xb7, 0x5d,
	0x9f, 0x5f, 0x04, 0x5f, 0xc1, 0xb3, 0x81, 0x58, 0x55, 0x49, 0xa6, 0x0a, 0x51, 0x34, 0xfe, 0xaf,
	0x79, 0x13, 0xe7, 0x4a, 0x8b, 0x94, 0x66, 0x0a, 0x9d, 0xd5, 0xe5, 0x39, 0xc2, 0xa3, 0x38, 0x85,
	0xce, 0x2e, 0x79, 0xb5, 0xce, 0xdb, 0xcb, 0x00, 0x0c, 0xf6, 0x6a, 0xff, 0x68, 0x82, 0x8d, 0x58,
	0xa4, 0xf0, 0xda, 0x57, 0xac, 0xfd, 0x60, 0xec, 0x24, 0xfb, 0xc5, 0xf5, 0xd8, 0x6f, 0x7c, 0xf9,
	0x58, 0x91, 0x13, 0xd1, 0xc3, 0x3c, 0x81, 0x22, 0x4b, 0x50, 0x42, 0xb9, 0xb9, 0x3c, 0x68, 0xb0,
	0xcd, 0x15, 0xcf, 0x6b, 0xcb, 0x56, 0xbf, 0x9b, 0x33, 0x3b, 0x41, 0xf0, 0xa7, 0xb9, 0xbe, 0x53,
	0x4a, 0x06, 0x44, 0xc1, 0xb2, 0x2c, 0xaa, 0x8e, 0x07, 0xc3, 0x1a, 0xf9, 0xb7, 0xc6, 0x74, 0x03,
	0xa2, 0xba, 0x16, 0xd3, 0xed, 0x78, 0x5d, 0x8b, 0xb9, 0x68, 0x6e, 0x94, 0x03, 0xdf, 0x0f, 0x88,
	0xf2, 0x7d, 0x8b, 0xf2, 0xfd, 0x8e, 0xe7, 0xfb, 0x16, 0x77, 0x34, 0x67, 0x0e, 0xeb, 0xfd, 0x0b,
	0x00, 0x00, 0xff, 0xff, 0x6a, 0x12, 0x59, 0x22, 0x0a, 0x06, 0x00, 0x00,
}
