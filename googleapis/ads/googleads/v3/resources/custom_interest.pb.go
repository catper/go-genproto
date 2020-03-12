// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/custom_interest.proto

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

// A custom interest. This is a list of users by interest.
type CustomInterest struct {
	// The resource name of the custom interest.
	// Custom interest resource names have the form:
	//
	// `customers/{customer_id}/customInterests/{custom_interest_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Id of the custom interest.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Status of this custom interest. Indicates whether the custom interest is
	// enabled or removed.
	Status enums.CustomInterestStatusEnum_CustomInterestStatus `protobuf:"varint,3,opt,name=status,proto3,enum=google.ads.googleads.v3.enums.CustomInterestStatusEnum_CustomInterestStatus" json:"status,omitempty"`
	// Name of the custom interest. It should be unique across the same custom
	// affinity audience.
	// This field is required for create operations.
	Name *wrappers.StringValue `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Type of the custom interest, CUSTOM_AFFINITY or CUSTOM_INTENT.
	// By default the type is set to CUSTOM_AFFINITY.
	Type enums.CustomInterestTypeEnum_CustomInterestType `protobuf:"varint,5,opt,name=type,proto3,enum=google.ads.googleads.v3.enums.CustomInterestTypeEnum_CustomInterestType" json:"type,omitempty"`
	// Description of this custom interest audience.
	Description *wrappers.StringValue `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	// List of custom interest members that this custom interest is composed of.
	// Members can be added during CustomInterest creation. If members are
	// presented in UPDATE operation, existing members will be overridden.
	Members              []*CustomInterestMember `protobuf:"bytes,7,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *CustomInterest) Reset()         { *m = CustomInterest{} }
func (m *CustomInterest) String() string { return proto.CompactTextString(m) }
func (*CustomInterest) ProtoMessage()    {}
func (*CustomInterest) Descriptor() ([]byte, []int) {
	return fileDescriptor_440649afc3ab8fa9, []int{0}
}

func (m *CustomInterest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomInterest.Unmarshal(m, b)
}
func (m *CustomInterest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomInterest.Marshal(b, m, deterministic)
}
func (m *CustomInterest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomInterest.Merge(m, src)
}
func (m *CustomInterest) XXX_Size() int {
	return xxx_messageInfo_CustomInterest.Size(m)
}
func (m *CustomInterest) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomInterest.DiscardUnknown(m)
}

var xxx_messageInfo_CustomInterest proto.InternalMessageInfo

func (m *CustomInterest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CustomInterest) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *CustomInterest) GetStatus() enums.CustomInterestStatusEnum_CustomInterestStatus {
	if m != nil {
		return m.Status
	}
	return enums.CustomInterestStatusEnum_UNSPECIFIED
}

func (m *CustomInterest) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *CustomInterest) GetType() enums.CustomInterestTypeEnum_CustomInterestType {
	if m != nil {
		return m.Type
	}
	return enums.CustomInterestTypeEnum_UNSPECIFIED
}

func (m *CustomInterest) GetDescription() *wrappers.StringValue {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *CustomInterest) GetMembers() []*CustomInterestMember {
	if m != nil {
		return m.Members
	}
	return nil
}

// A member of custom interest audience. A member can be a keyword or url.
// It is immutable, that is, it can only be created or removed but not changed.
type CustomInterestMember struct {
	// The type of custom interest member, KEYWORD or URL.
	MemberType enums.CustomInterestMemberTypeEnum_CustomInterestMemberType `protobuf:"varint,1,opt,name=member_type,json=memberType,proto3,enum=google.ads.googleads.v3.enums.CustomInterestMemberTypeEnum_CustomInterestMemberType" json:"member_type,omitempty"`
	// Keyword text when member_type is KEYWORD or URL string when
	// member_type is URL.
	Parameter            *wrappers.StringValue `protobuf:"bytes,2,opt,name=parameter,proto3" json:"parameter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CustomInterestMember) Reset()         { *m = CustomInterestMember{} }
func (m *CustomInterestMember) String() string { return proto.CompactTextString(m) }
func (*CustomInterestMember) ProtoMessage()    {}
func (*CustomInterestMember) Descriptor() ([]byte, []int) {
	return fileDescriptor_440649afc3ab8fa9, []int{1}
}

func (m *CustomInterestMember) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomInterestMember.Unmarshal(m, b)
}
func (m *CustomInterestMember) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomInterestMember.Marshal(b, m, deterministic)
}
func (m *CustomInterestMember) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomInterestMember.Merge(m, src)
}
func (m *CustomInterestMember) XXX_Size() int {
	return xxx_messageInfo_CustomInterestMember.Size(m)
}
func (m *CustomInterestMember) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomInterestMember.DiscardUnknown(m)
}

var xxx_messageInfo_CustomInterestMember proto.InternalMessageInfo

func (m *CustomInterestMember) GetMemberType() enums.CustomInterestMemberTypeEnum_CustomInterestMemberType {
	if m != nil {
		return m.MemberType
	}
	return enums.CustomInterestMemberTypeEnum_UNSPECIFIED
}

func (m *CustomInterestMember) GetParameter() *wrappers.StringValue {
	if m != nil {
		return m.Parameter
	}
	return nil
}

func init() {
	proto.RegisterType((*CustomInterest)(nil), "google.ads.googleads.v3.resources.CustomInterest")
	proto.RegisterType((*CustomInterestMember)(nil), "google.ads.googleads.v3.resources.CustomInterestMember")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/custom_interest.proto", fileDescriptor_440649afc3ab8fa9)
}

var fileDescriptor_440649afc3ab8fa9 = []byte{
	// 563 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xd1, 0x6a, 0x13, 0x4d,
	0x18, 0x65, 0x37, 0xf9, 0x53, 0x3a, 0xf9, 0xcd, 0xc5, 0xea, 0xc5, 0x5a, 0x8b, 0xa4, 0x95, 0x62,
	0x40, 0x98, 0x95, 0xac, 0xb4, 0xb2, 0x82, 0xb2, 0x15, 0xa9, 0x15, 0x95, 0xba, 0x2d, 0xb9, 0x90,
	0x40, 0x98, 0x64, 0xc7, 0x65, 0x21, 0x33, 0xb3, 0xcc, 0xcc, 0x56, 0x8a, 0x14, 0x7c, 0x16, 0x2f,
	0x7d, 0x13, 0xbd, 0xf3, 0x35, 0xfa, 0x14, 0xb2, 0x33, 0xb3, 0x93, 0xb4, 0x4d, 0x5a, 0x73, 0xf7,
	0x65, 0xbf, 0x73, 0xce, 0x77, 0x38, 0x7b, 0xb2, 0x60, 0x2f, 0x63, 0x2c, 0x9b, 0xe2, 0x00, 0xa5,
	0x22, 0xd0, 0x63, 0x35, 0x9d, 0x86, 0x01, 0xc7, 0x82, 0x95, 0x7c, 0x82, 0x45, 0x30, 0x29, 0x85,
	0x64, 0x64, 0x94, 0x53, 0x89, 0x39, 0x16, 0x12, 0x16, 0x9c, 0x49, 0xe6, 0x6d, 0x69, 0x34, 0x44,
	0xa9, 0x80, 0x96, 0x08, 0x4f, 0x43, 0x68, 0x89, 0x1b, 0xaf, 0x96, 0x69, 0x63, 0x5a, 0x92, 0x6b,
	0xba, 0x23, 0x82, 0xc9, 0x18, 0xf3, 0x91, 0x3c, 0x2b, 0xb0, 0xbe, 0xb1, 0x11, 0xad, 0x26, 0x20,
	0x24, 0x92, 0xa5, 0x30, 0xdc, 0xe7, 0xab, 0x71, 0xe7, 0xae, 0xde, 0xaf, 0x99, 0x45, 0x6e, 0x53,
	0x30, 0xab, 0x87, 0x66, 0xa5, 0x7e, 0x8d, 0xcb, 0x2f, 0xc1, 0x57, 0x8e, 0x8a, 0x02, 0xf3, 0xfa,
	0xe8, 0xe6, 0x1c, 0x15, 0x51, 0xca, 0x24, 0x92, 0x39, 0xa3, 0x66, 0xbb, 0xfd, 0xa7, 0x09, 0x3a,
	0xaf, 0xd5, 0xdd, 0x43, 0x73, 0xd6, 0x7b, 0x04, 0xee, 0xd4, 0x27, 0x46, 0x14, 0x11, 0xec, 0x3b,
	0x5d, 0xa7, 0xb7, 0x9e, 0xfc, 0x5f, 0x3f, 0xfc, 0x88, 0x08, 0xf6, 0x9e, 0x00, 0x37, 0x4f, 0x7d,
	0xb7, 0xeb, 0xf4, 0xda, 0xfd, 0x07, 0x26, 0x6c, 0x58, 0x5b, 0x80, 0x87, 0x54, 0xee, 0x3e, 0x1b,
	0xa0, 0x69, 0x89, 0x13, 0x37, 0x4f, 0xbd, 0x14, 0xb4, 0x74, 0x0e, 0x7e, 0xa3, 0xeb, 0xf4, 0x3a,
	0xfd, 0xf7, 0x70, 0xd9, 0x8b, 0x52, 0x41, 0xc0, 0xcb, 0x86, 0x8e, 0x15, 0xf5, 0x0d, 0x2d, 0xc9,
	0xc2, 0x45, 0x62, 0xb4, 0xbd, 0xa7, 0xa0, 0xa9, 0xec, 0x36, 0x95, 0xa9, 0xcd, 0x6b, 0xa6, 0x8e,
	0x25, 0xcf, 0x69, 0xa6, 0x5d, 0x29, 0xa4, 0x37, 0x04, 0xcd, 0x2a, 0x63, 0xff, 0x3f, 0xe5, 0xea,
	0xed, 0x4a, 0xae, 0x4e, 0xce, 0x0a, 0xbc, 0xc0, 0x53, 0xf5, 0x38, 0x51, 0xaa, 0xde, 0x4b, 0xd0,
	0x4e, 0xb1, 0x98, 0xf0, 0xbc, 0xa8, 0x02, 0xf7, 0x5b, 0xff, 0x60, 0x6b, 0x9e, 0xe0, 0x7d, 0x02,
	0x6b, 0xba, 0x7e, 0xc2, 0x5f, 0xeb, 0x36, 0x7a, 0xed, 0xfe, 0x1e, 0xbc, 0xb5, 0xdf, 0x57, 0xdc,
	0x7c, 0x50, 0xfc, 0xa4, 0xd6, 0x89, 0xd2, 0x8b, 0x18, 0x81, 0xc7, 0x33, 0xaa, 0x99, 0x8a, 0x5c,
	0xc0, 0x09, 0x23, 0xc1, 0x95, 0x22, 0xec, 0xea, 0x42, 0x62, 0x2e, 0x82, 0x6f, 0xf5, 0x78, 0x6e,
	0x5a, 0x5a, 0x83, 0xec, 0xca, 0xd6, 0xf6, 0x7c, 0xfb, 0x97, 0x03, 0xee, 0x2d, 0xf2, 0xe1, 0x95,
	0xa0, 0x3d, 0xf7, 0x87, 0x52, 0xbd, 0xea, 0xf4, 0x4f, 0x56, 0x8a, 0x5d, 0x2b, 0x2d, 0x09, 0x7f,
	0xb6, 0x4c, 0x00, 0xb1, 0xb3, 0x17, 0x81, 0xf5, 0x02, 0x71, 0x44, 0xb0, 0xc4, 0xdc, 0x54, 0xf6,
	0xe6, 0xd7, 0x30, 0x83, 0xef, 0x7f, 0x77, 0xc1, 0xce, 0x84, 0x91, 0xdb, 0x93, 0xdf, 0xbf, 0x7b,
	0xd9, 0xcb, 0x51, 0x25, 0x7c, 0xe4, 0x7c, 0x7e, 0x67, 0x98, 0x19, 0x9b, 0x22, 0x9a, 0x41, 0xc6,
	0xb3, 0x20, 0xc3, 0x54, 0x9d, 0x0d, 0x66, 0xf1, 0xdf, 0xf0, 0xad, 0x7b, 0x61, 0xa7, 0x1f, 0x6e,
	0xe3, 0x20, 0x8e, 0x7f, 0xba, 0x5b, 0x07, 0x5a, 0x32, 0x4e, 0x05, 0xd4, 0x63, 0x35, 0x0d, 0x42,
	0x98, 0xd4, 0xc8, 0xdf, 0x35, 0x66, 0x18, 0xa7, 0x62, 0x68, 0x31, 0xc3, 0x41, 0x38, 0xb4, 0x98,
	0x0b, 0x77, 0x47, 0x2f, 0xa2, 0x28, 0x4e, 0x45, 0x14, 0x59, 0x54, 0x14, 0x0d, 0xc2, 0x28, 0xb2,
	0xb8, 0x71, 0x4b, 0x99, 0x0d, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xcc, 0xef, 0xee, 0x77, 0x97,
	0x05, 0x00, 0x00,
}
