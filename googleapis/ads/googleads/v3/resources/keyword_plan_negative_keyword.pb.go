// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/keyword_plan_negative_keyword.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
	wrappers "github.com/catper/protobuf/ptypes/wrappers"
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

// A Keyword Plan negative keyword.
// Max number of keyword plan negative keywords per plan: 1000.
type KeywordPlanNegativeKeyword struct {
	// The resource name of the Keyword Plan negative keyword.
	// KeywordPlanNegativeKeyword resource names have the form:
	//
	// `customers/{customer_id}/keywordPlanNegativeKeywords/{kp_negative_keyword_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The Keyword Plan campaign to which this negative keyword belongs.
	KeywordPlanCampaign *wrappers.StringValue `protobuf:"bytes,2,opt,name=keyword_plan_campaign,json=keywordPlanCampaign,proto3" json:"keyword_plan_campaign,omitempty"`
	// The ID of the Keyword Plan negative keyword.
	Id *wrappers.Int64Value `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// The keyword text.
	Text *wrappers.StringValue `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
	// The keyword match type.
	MatchType            enums.KeywordMatchTypeEnum_KeywordMatchType `protobuf:"varint,5,opt,name=match_type,json=matchType,proto3,enum=google.ads.googleads.v3.enums.KeywordMatchTypeEnum_KeywordMatchType" json:"match_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                    `json:"-"`
	XXX_unrecognized     []byte                                      `json:"-"`
	XXX_sizecache        int32                                       `json:"-"`
}

func (m *KeywordPlanNegativeKeyword) Reset()         { *m = KeywordPlanNegativeKeyword{} }
func (m *KeywordPlanNegativeKeyword) String() string { return proto.CompactTextString(m) }
func (*KeywordPlanNegativeKeyword) ProtoMessage()    {}
func (*KeywordPlanNegativeKeyword) Descriptor() ([]byte, []int) {
	return fileDescriptor_385c61d2484cb50d, []int{0}
}

func (m *KeywordPlanNegativeKeyword) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordPlanNegativeKeyword.Unmarshal(m, b)
}
func (m *KeywordPlanNegativeKeyword) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordPlanNegativeKeyword.Marshal(b, m, deterministic)
}
func (m *KeywordPlanNegativeKeyword) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordPlanNegativeKeyword.Merge(m, src)
}
func (m *KeywordPlanNegativeKeyword) XXX_Size() int {
	return xxx_messageInfo_KeywordPlanNegativeKeyword.Size(m)
}
func (m *KeywordPlanNegativeKeyword) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordPlanNegativeKeyword.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordPlanNegativeKeyword proto.InternalMessageInfo

func (m *KeywordPlanNegativeKeyword) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *KeywordPlanNegativeKeyword) GetKeywordPlanCampaign() *wrappers.StringValue {
	if m != nil {
		return m.KeywordPlanCampaign
	}
	return nil
}

func (m *KeywordPlanNegativeKeyword) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *KeywordPlanNegativeKeyword) GetText() *wrappers.StringValue {
	if m != nil {
		return m.Text
	}
	return nil
}

func (m *KeywordPlanNegativeKeyword) GetMatchType() enums.KeywordMatchTypeEnum_KeywordMatchType {
	if m != nil {
		return m.MatchType
	}
	return enums.KeywordMatchTypeEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*KeywordPlanNegativeKeyword)(nil), "google.ads.googleads.v3.resources.KeywordPlanNegativeKeyword")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/keyword_plan_negative_keyword.proto", fileDescriptor_385c61d2484cb50d)
}

var fileDescriptor_385c61d2484cb50d = []byte{
	// 483 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4d, 0x6b, 0xd4, 0x40,
	0x18, 0x26, 0xd9, 0x2a, 0x74, 0xfc, 0x38, 0x44, 0x84, 0xb8, 0x16, 0xdd, 0x2a, 0x85, 0x05, 0x61,
	0x46, 0x1a, 0xe9, 0x61, 0x3c, 0xa5, 0x5a, 0x8a, 0x8a, 0x25, 0xac, 0xb2, 0x07, 0x59, 0x08, 0xd3,
	0x64, 0x1c, 0xc3, 0x66, 0x3e, 0x98, 0x99, 0x6c, 0x5d, 0x4a, 0xcf, 0x1e, 0xf4, 0x57, 0x78, 0xf4,
	0xa7, 0xf8, 0x3f, 0xbc, 0xf4, 0x57, 0xc8, 0x26, 0x93, 0x44, 0x29, 0xbb, 0x7a, 0x7b, 0xf2, 0xbe,
	0xcf, 0xf3, 0x3e, 0xef, 0x3c, 0x33, 0x01, 0x47, 0x4c, 0x4a, 0x56, 0x52, 0x44, 0x72, 0x83, 0x1a,
	0xb8, 0x42, 0x8b, 0x08, 0x69, 0x6a, 0x64, 0xa5, 0x33, 0x6a, 0xd0, 0x9c, 0x2e, 0xcf, 0xa4, 0xce,
	0x53, 0x55, 0x12, 0x91, 0x0a, 0xca, 0x88, 0x2d, 0x16, 0x34, 0x75, 0x55, 0xa8, 0xb4, 0xb4, 0x32,
	0xd8, 0x6d, 0xb4, 0x90, 0xe4, 0x06, 0x76, 0x63, 0xe0, 0x22, 0x82, 0xdd, 0x98, 0xe1, 0xc1, 0x3a,
	0x27, 0x2a, 0x2a, 0xde, 0xbb, 0x70, 0x62, 0xb3, 0x4f, 0xa9, 0x5d, 0x2a, 0xda, 0x8c, 0x1e, 0xde,
	0x6b, 0x75, 0xaa, 0xe8, 0x96, 0x72, 0xad, 0x07, 0xae, 0x55, 0x7f, 0x9d, 0x56, 0x1f, 0xd1, 0x99,
	0x26, 0x4a, 0x51, 0x6d, 0x5c, 0x7f, 0xe7, 0x0f, 0x29, 0x11, 0x42, 0x5a, 0x62, 0x0b, 0x29, 0x5c,
	0xf7, 0xd1, 0xaf, 0x01, 0x18, 0xbe, 0x69, 0x5c, 0x93, 0x92, 0x88, 0x13, 0x77, 0x32, 0x57, 0x0a,
	0x1e, 0x83, 0x5b, 0xad, 0x5d, 0x2a, 0x08, 0xa7, 0xa1, 0x37, 0xf2, 0xc6, 0xdb, 0x93, 0x9b, 0x6d,
	0xf1, 0x84, 0x70, 0x1a, 0x24, 0xe0, 0xee, 0x5f, 0xf1, 0x64, 0x84, 0x2b, 0x52, 0x30, 0x11, 0xfa,
	0x23, 0x6f, 0x7c, 0x63, 0x7f, 0xc7, 0x85, 0x01, 0xdb, 0x0d, 0xe1, 0x3b, 0xab, 0x0b, 0xc1, 0xa6,
	0xa4, 0xac, 0xe8, 0xe4, 0xce, 0xbc, 0x77, 0x7f, 0xe1, 0x84, 0xc1, 0x13, 0xe0, 0x17, 0x79, 0x38,
	0xa8, 0xe5, 0xf7, 0xaf, 0xc8, 0x5f, 0x09, 0x7b, 0xf0, 0xac, 0x51, 0xfb, 0x45, 0x1e, 0x3c, 0x05,
	0x5b, 0x96, 0x7e, 0xb6, 0xe1, 0xd6, 0x7f, 0xb8, 0xd5, 0xcc, 0x20, 0x03, 0xa0, 0x4f, 0x38, 0xbc,
	0x36, 0xf2, 0xc6, 0xb7, 0xf7, 0x5f, 0xc2, 0x75, 0xb7, 0x57, 0x5f, 0x0d, 0x74, 0x89, 0xbc, 0x5d,
	0xe9, 0xde, 0x2f, 0x15, 0x3d, 0x12, 0x15, 0xbf, 0x52, 0x9c, 0x6c, 0xf3, 0x16, 0xe2, 0xaf, 0xde,
	0x65, 0xfc, 0xc5, 0x03, 0x51, 0x3f, 0xcb, 0x21, 0x55, 0x18, 0x98, 0x49, 0x8e, 0x36, 0xc4, 0x9e,
	0x64, 0x95, 0xb1, 0x92, 0x53, 0x6d, 0xd0, 0x79, 0x0b, 0x2f, 0xd0, 0x7c, 0xad, 0xc0, 0xa0, 0xf3,
	0x8d, 0x0f, 0xf4, 0xe2, 0xf0, 0x9b, 0x0f, 0xf6, 0x32, 0xc9, 0xe1, 0x3f, 0x9f, 0xe8, 0xe1, 0xc3,
	0xf5, 0x7b, 0x25, 0xab, 0x48, 0x13, 0xef, 0xc3, 0x6b, 0x37, 0x85, 0xc9, 0x92, 0x08, 0x06, 0xa5,
	0x66, 0x88, 0x51, 0x51, 0x07, 0x8e, 0xfa, 0x43, 0x6e, 0xf8, 0x9d, 0x9e, 0x77, 0xe8, 0xbb, 0x3f,
	0x38, 0x8e, 0xe3, 0x1f, 0xfe, 0xee, 0x71, 0x33, 0x32, 0xce, 0x0d, 0x6c, 0xe0, 0x0a, 0x4d, 0x23,
	0x38, 0x69, 0x99, 0x3f, 0x5b, 0xce, 0x2c, 0xce, 0xcd, 0xac, 0xe3, 0xcc, 0xa6, 0xd1, 0xac, 0xe3,
	0x5c, 0xfa, 0x7b, 0x4d, 0x03, 0xe3, 0x38, 0x37, 0x18, 0x77, 0x2c, 0x8c, 0xa7, 0x11, 0xc6, 0x1d,
	0xef, 0xf4, 0x7a, 0xbd, 0x6c, 0xf4, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x61, 0xe0, 0x73, 0xc2, 0xfa,
	0x03, 0x00, 0x00,
}
