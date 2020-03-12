// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/user_interest.proto

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

// A user interest: a particular interest-based vertical to be targeted.
type UserInterest struct {
	// The resource name of the user interest.
	// User interest resource names have the form:
	//
	// `customers/{customer_id}/userInterests/{user_interest_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Taxonomy type of the user interest.
	TaxonomyType enums.UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType `protobuf:"varint,2,opt,name=taxonomy_type,json=taxonomyType,proto3,enum=google.ads.googleads.v3.enums.UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType" json:"taxonomy_type,omitempty"`
	// The ID of the user interest.
	UserInterestId *wrappers.Int64Value `protobuf:"bytes,3,opt,name=user_interest_id,json=userInterestId,proto3" json:"user_interest_id,omitempty"`
	// The name of the user interest.
	Name *wrappers.StringValue `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// The parent of the user interest.
	UserInterestParent *wrappers.StringValue `protobuf:"bytes,5,opt,name=user_interest_parent,json=userInterestParent,proto3" json:"user_interest_parent,omitempty"`
	// True if the user interest is launched to all channels and locales.
	LaunchedToAll *wrappers.BoolValue `protobuf:"bytes,6,opt,name=launched_to_all,json=launchedToAll,proto3" json:"launched_to_all,omitempty"`
	// Availability information of the user interest.
	Availabilities       []*common.CriterionCategoryAvailability `protobuf:"bytes,7,rep,name=availabilities,proto3" json:"availabilities,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *UserInterest) Reset()         { *m = UserInterest{} }
func (m *UserInterest) String() string { return proto.CompactTextString(m) }
func (*UserInterest) ProtoMessage()    {}
func (*UserInterest) Descriptor() ([]byte, []int) {
	return fileDescriptor_853f2530b9dfe755, []int{0}
}

func (m *UserInterest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInterest.Unmarshal(m, b)
}
func (m *UserInterest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInterest.Marshal(b, m, deterministic)
}
func (m *UserInterest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInterest.Merge(m, src)
}
func (m *UserInterest) XXX_Size() int {
	return xxx_messageInfo_UserInterest.Size(m)
}
func (m *UserInterest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInterest.DiscardUnknown(m)
}

var xxx_messageInfo_UserInterest proto.InternalMessageInfo

func (m *UserInterest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *UserInterest) GetTaxonomyType() enums.UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType {
	if m != nil {
		return m.TaxonomyType
	}
	return enums.UserInterestTaxonomyTypeEnum_UNSPECIFIED
}

func (m *UserInterest) GetUserInterestId() *wrappers.Int64Value {
	if m != nil {
		return m.UserInterestId
	}
	return nil
}

func (m *UserInterest) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *UserInterest) GetUserInterestParent() *wrappers.StringValue {
	if m != nil {
		return m.UserInterestParent
	}
	return nil
}

func (m *UserInterest) GetLaunchedToAll() *wrappers.BoolValue {
	if m != nil {
		return m.LaunchedToAll
	}
	return nil
}

func (m *UserInterest) GetAvailabilities() []*common.CriterionCategoryAvailability {
	if m != nil {
		return m.Availabilities
	}
	return nil
}

func init() {
	proto.RegisterType((*UserInterest)(nil), "google.ads.googleads.v3.resources.UserInterest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/user_interest.proto", fileDescriptor_853f2530b9dfe755)
}

var fileDescriptor_853f2530b9dfe755 = []byte{
	// 564 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0xd3, 0x1f, 0x84, 0xdb, 0x06, 0xb0, 0x38, 0x98, 0x50, 0xa1, 0x14, 0x14, 0x29, 0xa7,
	0x5d, 0x14, 0x03, 0x07, 0x23, 0x84, 0x9c, 0x52, 0x55, 0xe1, 0x50, 0x45, 0x21, 0xe4, 0x00, 0x11,
	0xd6, 0xc6, 0x1e, 0xcc, 0x4a, 0xeb, 0x5d, 0x6b, 0x77, 0x1d, 0xb0, 0xaa, 0x5e, 0x78, 0x14, 0x8e,
	0x3c, 0x0a, 0x8f, 0xd2, 0x87, 0x40, 0x28, 0xfe, 0xab, 0xd3, 0x2a, 0xc0, 0xed, 0xdb, 0xcc, 0xf7,
	0x33, 0x33, 0x99, 0xc4, 0x7c, 0x1e, 0x09, 0x11, 0x31, 0xc0, 0x24, 0x54, 0xb8, 0x80, 0x2b, 0xb4,
	0x74, 0xb0, 0x04, 0x25, 0x52, 0x19, 0x80, 0xc2, 0xa9, 0x02, 0xe9, 0x53, 0xae, 0x41, 0x82, 0xd2,
	0x28, 0x91, 0x42, 0x0b, 0xeb, 0xa8, 0xe0, 0x22, 0x12, 0x2a, 0x54, 0xcb, 0xd0, 0xd2, 0x41, 0xb5,
	0xac, 0xf3, 0x66, 0x93, 0x73, 0x20, 0xe2, 0x58, 0x70, 0x1c, 0x48, 0xaa, 0x41, 0x52, 0xc1, 0xfd,
	0x80, 0x68, 0x88, 0x84, 0xcc, 0x7c, 0xb2, 0x24, 0x94, 0x91, 0x05, 0x65, 0x54, 0x67, 0x45, 0x50,
	0xe7, 0xf5, 0x26, 0x17, 0xe0, 0x69, 0x7c, 0xad, 0x37, 0x5f, 0x93, 0x6f, 0x82, 0x8b, 0x38, 0xf3,
	0x75, 0x96, 0x40, 0x69, 0xf0, 0xa0, 0x32, 0x48, 0x68, 0x3d, 0x53, 0x59, 0x7a, 0x54, 0x96, 0xf2,
	0xd7, 0x22, 0xfd, 0x8c, 0xbf, 0x4a, 0x92, 0x24, 0x20, 0x55, 0x59, 0x3f, 0x6c, 0x48, 0x09, 0xe7,
	0x42, 0x13, 0x4d, 0x05, 0x2f, 0xab, 0x8f, 0xbf, 0xef, 0x98, 0xfb, 0xef, 0x15, 0xc8, 0x51, 0x99,
	0x6e, 0x3d, 0x31, 0x0f, 0xaa, 0x00, 0x9f, 0x93, 0x18, 0x6c, 0xa3, 0x6b, 0xf4, 0x6f, 0x4f, 0xf6,
	0xab, 0x0f, 0xcf, 0x48, 0x0c, 0x56, 0x66, 0x1e, 0xac, 0x75, 0x69, 0xb7, 0xba, 0x46, 0xbf, 0x3d,
	0x98, 0xa2, 0x4d, 0x0b, 0xcd, 0xe7, 0x44, 0xcd, 0xa0, 0x69, 0xa9, 0x9f, 0x66, 0x09, 0x9c, 0xf0,
	0x34, 0xde, 0x58, 0x9c, 0xec, 0xeb, 0xc6, 0xcb, 0x3a, 0x31, 0xef, 0xae, 0xaf, 0x8b, 0x86, 0xf6,
	0x56, 0xd7, 0xe8, 0xef, 0x0d, 0x1e, 0x56, 0xe9, 0xd5, 0x26, 0xd0, 0x88, 0xeb, 0x17, 0xcf, 0x66,
	0x84, 0xa5, 0x30, 0x69, 0xa7, 0x0d, 0xfb, 0x51, 0x68, 0x3d, 0x35, 0xb7, 0xf3, 0xe9, 0xb6, 0x73,
	0xe9, 0xe1, 0x0d, 0xe9, 0x3b, 0x2d, 0x29, 0x8f, 0x0a, 0x6d, 0xce, 0xb4, 0xce, 0xcc, 0xfb, 0xeb,
	0xc1, 0x09, 0x91, 0xc0, 0xb5, 0xbd, 0xf3, 0x1f, 0x0e, 0x56, 0x33, 0x7d, 0x9c, 0xeb, 0xac, 0xa1,
	0x79, 0x87, 0x91, 0x94, 0x07, 0x5f, 0x20, 0xf4, 0xb5, 0xf0, 0x09, 0x63, 0xf6, 0x6e, 0x6e, 0xd5,
	0xb9, 0x61, 0x35, 0x14, 0x82, 0x15, 0x46, 0x07, 0x95, 0x64, 0x2a, 0x3c, 0xc6, 0x2c, 0x30, 0xdb,
	0x8d, 0x6b, 0xa3, 0xa0, 0xec, 0x5b, 0xdd, 0xad, 0xfe, 0xde, 0xe0, 0xd5, 0xc6, 0x2f, 0xa2, 0x38,
	0x5b, 0x74, 0x5c, 0x9d, 0xed, 0x71, 0x79, 0xb5, 0x5e, 0xe3, 0x68, 0x27, 0xd7, 0x4c, 0xdd, 0x4f,
	0x97, 0xde, 0x47, 0xb3, 0x77, 0xe5, 0x53, 0xa2, 0x84, 0xaa, 0x95, 0x1f, 0x5e, 0xbb, 0x9f, 0x41,
	0x90, 0x2a, 0x2d, 0x62, 0x90, 0x0a, 0x9f, 0x57, 0xf0, 0x02, 0x37, 0xe7, 0x57, 0xf8, 0x7c, 0x6d,
	0x91, 0x17, 0xc3, 0xdf, 0x86, 0xd9, 0x0b, 0x44, 0x8c, 0xfe, 0xf9, 0x73, 0x1c, 0xde, 0x6b, 0x66,
	0x8d, 0x57, 0xfb, 0x19, 0x1b, 0x1f, 0xde, 0x96, 0xba, 0x48, 0x30, 0xc2, 0x23, 0x24, 0x64, 0x84,
	0x23, 0xe0, 0xf9, 0xf6, 0xf0, 0x55, 0xab, 0x7f, 0xf9, 0x73, 0x78, 0x59, 0xa3, 0x1f, 0xad, 0xad,
	0x53, 0xcf, 0xfb, 0xd9, 0x3a, 0x3a, 0x2d, 0x2c, 0xbd, 0x50, 0xa1, 0x02, 0xae, 0xd0, 0xcc, 0x41,
	0x93, 0x8a, 0xf9, 0xab, 0xe2, 0xcc, 0xbd, 0x50, 0xcd, 0x6b, 0xce, 0x7c, 0xe6, 0xcc, 0x6b, 0xce,
	0x65, 0xab, 0x57, 0x14, 0x5c, 0xd7, 0x0b, 0x95, 0xeb, 0xd6, 0x2c, 0xd7, 0x9d, 0x39, 0xae, 0x5b,
	0xf3, 0x16, 0xbb, 0x79, 0xb3, 0xce, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x00, 0x06, 0xa2,
	0xc8, 0x04, 0x00, 0x00,
}
