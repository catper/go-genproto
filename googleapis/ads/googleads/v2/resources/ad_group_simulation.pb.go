// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/ad_group_simulation.proto

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
	// Immutable. The resource name of the ad group simulation.
	// Ad group simulation resource names have the form:
	//
	// `customers/{customer_id}/adGroupSimulations/{ad_group_id}~{type}~{modification_method}~{start_date}~{end_date}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. Ad group id of the simulation.
	AdGroupId *wrappers.Int64Value `protobuf:"bytes,2,opt,name=ad_group_id,json=adGroupId,proto3" json:"ad_group_id,omitempty"`
	// Output only. The field that the simulation modifies.
	Type enums.SimulationTypeEnum_SimulationType `protobuf:"varint,3,opt,name=type,proto3,enum=google.ads.googleads.v2.enums.SimulationTypeEnum_SimulationType" json:"type,omitempty"`
	// Output only. How the simulation modifies the field.
	ModificationMethod enums.SimulationModificationMethodEnum_SimulationModificationMethod `protobuf:"varint,4,opt,name=modification_method,json=modificationMethod,proto3,enum=google.ads.googleads.v2.enums.SimulationModificationMethodEnum_SimulationModificationMethod" json:"modification_method,omitempty"`
	// Output only. First day on which the simulation is based, in YYYY-MM-DD format.
	StartDate *wrappers.StringValue `protobuf:"bytes,5,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	// Output only. Last day on which the simulation is based, in YYYY-MM-DD format
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
	return fileDescriptor_b65db7c3498c5238, []int{0}
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
	proto.RegisterType((*AdGroupSimulation)(nil), "google.ads.googleads.v2.resources.AdGroupSimulation")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/ad_group_simulation.proto", fileDescriptor_b65db7c3498c5238)
}

var fileDescriptor_b65db7c3498c5238 = []byte{
	// 660 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x6a, 0x13, 0x41,
	0x14, 0x36, 0x49, 0x7f, 0xa7, 0x55, 0x71, 0x45, 0x89, 0xb5, 0x68, 0x2b, 0x14, 0x8a, 0xc8, 0x0c,
	0x6c, 0x45, 0x70, 0x5b, 0xa1, 0xbb, 0x55, 0x6a, 0xc5, 0x4a, 0x49, 0x4b, 0x44, 0x09, 0x2c, 0x93,
	0x9d, 0xe9, 0x76, 0x30, 0x3b, 0x33, 0xec, 0xcc, 0xae, 0x94, 0xd2, 0x2b, 0xdf, 0xc0, 0x47, 0xf0,
	0xd2, 0x47, 0xf1, 0x29, 0x7a, 0xdd, 0x4b, 0x2f, 0xbd, 0x92, 0xcc, 0xfe, 0x24, 0x26, 0x4d, 0x13,
	0xef, 0xce, 0xee, 0xf9, 0x7e, 0xce, 0xf9, 0x76, 0x76, 0xc0, 0x66, 0x28, 0x44, 0xd8, 0xa1, 0x08,
	0x13, 0x85, 0xb2, 0xb2, 0x5b, 0xa5, 0x36, 0x8a, 0xa9, 0x12, 0x49, 0x1c, 0x50, 0x85, 0x30, 0xf1,
	0xc3, 0x58, 0x24, 0xd2, 0x57, 0x2c, 0x4a, 0x3a, 0x58, 0x33, 0xc1, 0xa1, 0x8c, 0x85, 0x16, 0xd6,
	0x6a, 0xc6, 0x80, 0x98, 0x28, 0x58, 0x92, 0x61, 0x6a, 0xc3, 0x92, 0xbc, 0x84, 0x46, 0xe9, 0x07,
	0x22, 0x8a, 0x04, 0x47, 0x83, 0x9a, 0x4b, 0xde, 0x28, 0x02, 0xe5, 0x49, 0xa4, 0xfa, 0xf0, 0x7e,
	0x24, 0x08, 0x3b, 0x66, 0x41, 0xfe, 0x40, 0xf5, 0x89, 0x20, 0xb9, 0xc6, 0xc6, 0xc4, 0x1a, 0xfa,
	0x54, 0xd2, 0x9c, 0xf4, 0xb8, 0x20, 0x49, 0x86, 0x8e, 0x19, 0xed, 0x10, 0xbf, 0x4d, 0x4f, 0x70,
	0xca, 0x44, 0x9c, 0x03, 0x1e, 0xf4, 0x01, 0x8a, 0x05, 0xf3, 0xd6, 0xa3, 0xbc, 0x65, 0x9e, 0xda,
	0xc9, 0x31, 0xfa, 0x1a, 0x63, 0x29, 0x69, 0xac, 0xf2, 0xfe, 0x72, 0x1f, 0x15, 0x73, 0x2e, 0xb4,
	0x71, 0xcf, 0xbb, 0x4f, 0x7e, 0xcf, 0x82, 0x3b, 0x2e, 0xd9, 0xed, 0x66, 0x7c, 0x58, 0x8e, 0x66,
	0x7d, 0x04, 0x37, 0x0b, 0x17, 0x9f, 0xe3, 0x88, 0xd6, 0x2b, 0x2b, 0x95, 0xf5, 0x79, 0xcf, 0xbe,
	0x70, 0xa7, 0xff, 0xb8, 0xcf, 0xc0, 0xd3, 0x5e, 0xe0, 0x79, 0x25, 0x99, 0x82, 0x81, 0x88, 0xd0,
	0x90, 0x54, 0x63, 0xb1, 0x10, 0xfa, 0x80, 0x23, 0x6a, 0x6d, 0x83, 0x85, 0xf2, 0x93, 0x32, 0x52,
	0xaf, 0xae, 0x54, 0xd6, 0x17, 0xec, 0x87, 0xb9, 0x0a, 0x2c, 0x56, 0x80, 0x7b, 0x5c, 0xbf, 0x78,
	0xde, 0xc4, 0x9d, 0x84, 0x7a, 0xb5, 0x0b, 0xb7, 0xd6, 0x98, 0xc7, 0x99, 0xee, 0x1e, 0xb1, 0x3e,
	0x81, 0xa9, 0x6e, 0x70, 0xf5, 0xda, 0x4a, 0x65, 0xfd, 0x96, 0xbd, 0x0d, 0x47, 0x1d, 0x03, 0x13,
	0x37, 0xec, 0x0d, 0x72, 0x74, 0x2a, 0xe9, 0x1b, 0x9e, 0x44, 0x03, 0xaf, 0x32, 0x7d, 0x23, 0x69,
	0x7d, 0xaf, 0x80, 0xbb, 0x57, 0x7c, 0xd8, 0xfa, 0x94, 0xb1, 0x6a, 0x4d, 0x6c, 0xb5, 0xdf, 0xa7,
	0xb1, 0x6f, 0x24, 0x06, 0x8c, 0x87, 0x01, 0xd9, 0x18, 0x56, 0x34, 0xd4, 0xb0, 0xb6, 0x01, 0x50,
	0x1a, 0xc7, 0xda, 0x27, 0x58, 0xd3, 0xfa, 0xb4, 0x09, 0x6c, 0x79, 0x28, 0xb0, 0x43, 0x1d, 0x33,
	0x1e, 0xf6, 0x27, 0x66, 0x48, 0xaf, 0xb1, 0xa6, 0xd6, 0x16, 0x98, 0xa3, 0x9c, 0x64, 0xfc, 0x99,
	0x49, 0xf9, 0xb3, 0x94, 0x13, 0xc3, 0x8e, 0x80, 0x15, 0xc8, 0xc0, 0x6f, 0x33, 0xe2, 0x4b, 0xc1,
	0xb8, 0xf6, 0x3b, 0x4c, 0xe9, 0xfa, 0x9c, 0xd1, 0x79, 0x39, 0x32, 0x92, 0xec, 0x0f, 0x83, 0x3b,
	0x32, 0xf0, 0x18, 0xe9, 0x2d, 0x7e, 0xd0, 0x55, 0x78, 0xcf, 0x94, 0x36, 0x26, 0x6f, 0x6f, 0x34,
	0x6e, 0x07, 0x06, 0x51, 0xbe, 0xcf, 0xec, 0xd2, 0x41, 0x3b, 0x30, 0xa9, 0x5d, 0x3a, 0xd6, 0x2e,
	0xfd, 0xc7, 0x2e, 0x05, 0xf7, 0x34, 0x8e, 0x43, 0xaa, 0xfd, 0x40, 0xe2, 0x7e, 0xc7, 0x79, 0xe3,
	0xb8, 0x35, 0xce, 0xf1, 0xc8, 0x90, 0x77, 0x24, 0xbe, 0xc6, 0xd4, 0xd2, 0x05, 0xa8, 0x6c, 0x39,
	0xfc, 0xd2, 0xfd, 0xf2, 0x3f, 0xbf, 0x91, 0xf5, 0x2a, 0x48, 0x94, 0x16, 0x11, 0x8d, 0x15, 0x3a,
	0x2b, 0xca, 0x73, 0x84, 0x07, 0x71, 0x0a, 0x9d, 0x5d, 0x71, 0x65, 0x9e, 0x7b, 0x8b, 0x00, 0xf4,
	0x96, 0xf3, 0xbe, 0x55, 0xc1, 0x5a, 0x20, 0x22, 0x38, 0xf6, 0x0a, 0xf5, 0xee, 0x0f, 0x4d, 0x72,
	0xd0, 0x3d, 0x32, 0x07, 0x95, 0xcf, 0xef, 0x72, 0x72, 0x28, 0x3a, 0x98, 0x87, 0x50, 0xc4, 0x21,
	0x0a, 0x29, 0x37, 0x07, 0x0a, 0xf5, 0xb6, 0xb9, 0xe6, 0x6e, 0xdf, 0x2c, 0xab, 0x1f, 0xd5, 0xda,
	0xae, 0xeb, 0xfe, 0xac, 0xae, 0xee, 0x66, 0x92, 0x2e, 0x51, 0x30, 0x2b, 0xbb, 0x55, 0xd3, 0x86,
	0x8d, 0x02, 0xf9, 0xab, 0xc0, 0xb4, 0x5c, 0xa2, 0x5a, 0x25, 0xa6, 0xd5, 0xb4, 0x5b, 0x25, 0xe6,
	0xb2, 0xba, 0x96, 0x35, 0x1c, 0xc7, 0x25, 0xca, 0x71, 0x4a, 0x94, 0xe3, 0x34, 0x6d, 0xc7, 0x29,
	0x71, 0xed, 0x19, 0x33, 0xec, 0xc6, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x00, 0x9b, 0x91, 0x55,
	0x87, 0x06, 0x00, 0x00,
}
