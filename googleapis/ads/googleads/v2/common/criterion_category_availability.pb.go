// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/common/criterion_category_availability.proto

package common

import (
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
	wrappers "github.com/catper/protobuf/ptypes/wrappers"
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

// Information of category availability, per advertising channel.
type CriterionCategoryAvailability struct {
	// Channel types and subtypes that are available to the category.
	Channel *CriterionCategoryChannelAvailability `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	// Locales that are available to the category for the channel.
	Locale               []*CriterionCategoryLocaleAvailability `protobuf:"bytes,2,rep,name=locale,proto3" json:"locale,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *CriterionCategoryAvailability) Reset()         { *m = CriterionCategoryAvailability{} }
func (m *CriterionCategoryAvailability) String() string { return proto.CompactTextString(m) }
func (*CriterionCategoryAvailability) ProtoMessage()    {}
func (*CriterionCategoryAvailability) Descriptor() ([]byte, []int) {
	return fileDescriptor_670ab9bba1a25e93, []int{0}
}

func (m *CriterionCategoryAvailability) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CriterionCategoryAvailability.Unmarshal(m, b)
}
func (m *CriterionCategoryAvailability) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CriterionCategoryAvailability.Marshal(b, m, deterministic)
}
func (m *CriterionCategoryAvailability) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CriterionCategoryAvailability.Merge(m, src)
}
func (m *CriterionCategoryAvailability) XXX_Size() int {
	return xxx_messageInfo_CriterionCategoryAvailability.Size(m)
}
func (m *CriterionCategoryAvailability) XXX_DiscardUnknown() {
	xxx_messageInfo_CriterionCategoryAvailability.DiscardUnknown(m)
}

var xxx_messageInfo_CriterionCategoryAvailability proto.InternalMessageInfo

func (m *CriterionCategoryAvailability) GetChannel() *CriterionCategoryChannelAvailability {
	if m != nil {
		return m.Channel
	}
	return nil
}

func (m *CriterionCategoryAvailability) GetLocale() []*CriterionCategoryLocaleAvailability {
	if m != nil {
		return m.Locale
	}
	return nil
}

// Information of advertising channel type and subtypes a category is available
// in.
type CriterionCategoryChannelAvailability struct {
	// Format of the channel availability. Can be ALL_CHANNELS (the rest of the
	// fields will not be set), CHANNEL_TYPE (only advertising_channel_type type
	// will be set, the category is available to all sub types under it) or
	// CHANNEL_TYPE_AND_SUBTYPES (advertising_channel_type,
	// advertising_channel_sub_type, and include_default_channel_sub_type will all
	// be set).
	AvailabilityMode enums.CriterionCategoryChannelAvailabilityModeEnum_CriterionCategoryChannelAvailabilityMode `protobuf:"varint,1,opt,name=availability_mode,json=availabilityMode,proto3,enum=google.ads.googleads.v2.enums.CriterionCategoryChannelAvailabilityModeEnum_CriterionCategoryChannelAvailabilityMode" json:"availability_mode,omitempty"`
	// Channel type the category is available to.
	AdvertisingChannelType enums.AdvertisingChannelTypeEnum_AdvertisingChannelType `protobuf:"varint,2,opt,name=advertising_channel_type,json=advertisingChannelType,proto3,enum=google.ads.googleads.v2.enums.AdvertisingChannelTypeEnum_AdvertisingChannelType" json:"advertising_channel_type,omitempty"`
	// Channel subtypes under the channel type the category is available to.
	AdvertisingChannelSubType []enums.AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType `protobuf:"varint,3,rep,packed,name=advertising_channel_sub_type,json=advertisingChannelSubType,proto3,enum=google.ads.googleads.v2.enums.AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType" json:"advertising_channel_sub_type,omitempty"`
	// Whether default channel sub type is included. For example,
	// advertising_channel_type being DISPLAY and include_default_channel_sub_type
	// being false means that the default display campaign where channel sub type
	// is not set is not included in this availability configuration.
	IncludeDefaultChannelSubType *wrappers.BoolValue `protobuf:"bytes,4,opt,name=include_default_channel_sub_type,json=includeDefaultChannelSubType,proto3" json:"include_default_channel_sub_type,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}            `json:"-"`
	XXX_unrecognized             []byte              `json:"-"`
	XXX_sizecache                int32               `json:"-"`
}

func (m *CriterionCategoryChannelAvailability) Reset()         { *m = CriterionCategoryChannelAvailability{} }
func (m *CriterionCategoryChannelAvailability) String() string { return proto.CompactTextString(m) }
func (*CriterionCategoryChannelAvailability) ProtoMessage()    {}
func (*CriterionCategoryChannelAvailability) Descriptor() ([]byte, []int) {
	return fileDescriptor_670ab9bba1a25e93, []int{1}
}

func (m *CriterionCategoryChannelAvailability) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CriterionCategoryChannelAvailability.Unmarshal(m, b)
}
func (m *CriterionCategoryChannelAvailability) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CriterionCategoryChannelAvailability.Marshal(b, m, deterministic)
}
func (m *CriterionCategoryChannelAvailability) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CriterionCategoryChannelAvailability.Merge(m, src)
}
func (m *CriterionCategoryChannelAvailability) XXX_Size() int {
	return xxx_messageInfo_CriterionCategoryChannelAvailability.Size(m)
}
func (m *CriterionCategoryChannelAvailability) XXX_DiscardUnknown() {
	xxx_messageInfo_CriterionCategoryChannelAvailability.DiscardUnknown(m)
}

var xxx_messageInfo_CriterionCategoryChannelAvailability proto.InternalMessageInfo

func (m *CriterionCategoryChannelAvailability) GetAvailabilityMode() enums.CriterionCategoryChannelAvailabilityModeEnum_CriterionCategoryChannelAvailabilityMode {
	if m != nil {
		return m.AvailabilityMode
	}
	return enums.CriterionCategoryChannelAvailabilityModeEnum_UNSPECIFIED
}

func (m *CriterionCategoryChannelAvailability) GetAdvertisingChannelType() enums.AdvertisingChannelTypeEnum_AdvertisingChannelType {
	if m != nil {
		return m.AdvertisingChannelType
	}
	return enums.AdvertisingChannelTypeEnum_UNSPECIFIED
}

func (m *CriterionCategoryChannelAvailability) GetAdvertisingChannelSubType() []enums.AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType {
	if m != nil {
		return m.AdvertisingChannelSubType
	}
	return nil
}

func (m *CriterionCategoryChannelAvailability) GetIncludeDefaultChannelSubType() *wrappers.BoolValue {
	if m != nil {
		return m.IncludeDefaultChannelSubType
	}
	return nil
}

// Information about which locales a category is available in.
type CriterionCategoryLocaleAvailability struct {
	// Format of the locale availability. Can be LAUNCHED_TO_ALL (both country and
	// language will be empty), COUNTRY (only country will be set), LANGUAGE (only
	// language wil be set), COUNTRY_AND_LANGUAGE (both country and language will
	// be set).
	AvailabilityMode enums.CriterionCategoryLocaleAvailabilityModeEnum_CriterionCategoryLocaleAvailabilityMode `protobuf:"varint,1,opt,name=availability_mode,json=availabilityMode,proto3,enum=google.ads.googleads.v2.enums.CriterionCategoryLocaleAvailabilityModeEnum_CriterionCategoryLocaleAvailabilityMode" json:"availability_mode,omitempty"`
	// Code of the country.
	CountryCode *wrappers.StringValue `protobuf:"bytes,2,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	// Code of the language.
	LanguageCode         *wrappers.StringValue `protobuf:"bytes,3,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CriterionCategoryLocaleAvailability) Reset()         { *m = CriterionCategoryLocaleAvailability{} }
func (m *CriterionCategoryLocaleAvailability) String() string { return proto.CompactTextString(m) }
func (*CriterionCategoryLocaleAvailability) ProtoMessage()    {}
func (*CriterionCategoryLocaleAvailability) Descriptor() ([]byte, []int) {
	return fileDescriptor_670ab9bba1a25e93, []int{2}
}

func (m *CriterionCategoryLocaleAvailability) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CriterionCategoryLocaleAvailability.Unmarshal(m, b)
}
func (m *CriterionCategoryLocaleAvailability) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CriterionCategoryLocaleAvailability.Marshal(b, m, deterministic)
}
func (m *CriterionCategoryLocaleAvailability) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CriterionCategoryLocaleAvailability.Merge(m, src)
}
func (m *CriterionCategoryLocaleAvailability) XXX_Size() int {
	return xxx_messageInfo_CriterionCategoryLocaleAvailability.Size(m)
}
func (m *CriterionCategoryLocaleAvailability) XXX_DiscardUnknown() {
	xxx_messageInfo_CriterionCategoryLocaleAvailability.DiscardUnknown(m)
}

var xxx_messageInfo_CriterionCategoryLocaleAvailability proto.InternalMessageInfo

func (m *CriterionCategoryLocaleAvailability) GetAvailabilityMode() enums.CriterionCategoryLocaleAvailabilityModeEnum_CriterionCategoryLocaleAvailabilityMode {
	if m != nil {
		return m.AvailabilityMode
	}
	return enums.CriterionCategoryLocaleAvailabilityModeEnum_UNSPECIFIED
}

func (m *CriterionCategoryLocaleAvailability) GetCountryCode() *wrappers.StringValue {
	if m != nil {
		return m.CountryCode
	}
	return nil
}

func (m *CriterionCategoryLocaleAvailability) GetLanguageCode() *wrappers.StringValue {
	if m != nil {
		return m.LanguageCode
	}
	return nil
}

func init() {
	proto.RegisterType((*CriterionCategoryAvailability)(nil), "google.ads.googleads.v2.common.CriterionCategoryAvailability")
	proto.RegisterType((*CriterionCategoryChannelAvailability)(nil), "google.ads.googleads.v2.common.CriterionCategoryChannelAvailability")
	proto.RegisterType((*CriterionCategoryLocaleAvailability)(nil), "google.ads.googleads.v2.common.CriterionCategoryLocaleAvailability")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/common/criterion_category_availability.proto", fileDescriptor_670ab9bba1a25e93)
}

var fileDescriptor_670ab9bba1a25e93 = []byte{
	// 608 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x95, 0x41, 0x8e, 0xd3, 0x30,
	0x14, 0x86, 0x95, 0x74, 0x34, 0x48, 0xee, 0x30, 0x82, 0x2c, 0x50, 0xa9, 0xca, 0xa8, 0x0a, 0x2c,
	0xba, 0x72, 0xa4, 0xb0, 0x0b, 0x48, 0x90, 0xa6, 0x68, 0x36, 0x20, 0xaa, 0x0e, 0xea, 0x02, 0x2a,
	0x22, 0x27, 0xf1, 0x04, 0x4b, 0x8e, 0x1d, 0x25, 0x4e, 0x51, 0xaf, 0xc0, 0x05, 0x58, 0x22, 0x58,
	0x72, 0x14, 0x58, 0xb3, 0xe3, 0x04, 0xdc, 0x01, 0x09, 0x25, 0x76, 0x4a, 0x67, 0xda, 0x64, 0xda,
	0x55, 0xdd, 0xf8, 0xbd, 0xef, 0xff, 0xdf, 0xcb, 0xb3, 0x03, 0x26, 0x31, 0xe7, 0x31, 0xc5, 0x16,
	0x8a, 0x72, 0x4b, 0x2e, 0xcb, 0xd5, 0xd2, 0xb6, 0x42, 0x9e, 0x24, 0x9c, 0x59, 0x61, 0x46, 0x04,
	0xce, 0x08, 0x67, 0x7e, 0x88, 0x04, 0x8e, 0x79, 0xb6, 0xf2, 0xd1, 0x12, 0x11, 0x8a, 0x02, 0x42,
	0x89, 0x58, 0xc1, 0x34, 0xe3, 0x82, 0x1b, 0x67, 0x32, 0x15, 0xa2, 0x28, 0x87, 0x6b, 0x0a, 0x5c,
	0xda, 0x50, 0x52, 0xfa, 0xcf, 0x9b, 0x54, 0x30, 0x2b, 0x92, 0xdc, 0x42, 0xd1, 0x12, 0x67, 0x82,
	0xe4, 0x84, 0xc5, 0x7e, 0xf8, 0x01, 0x31, 0x86, 0xa9, 0x9f, 0x17, 0x81, 0x2f, 0x56, 0x29, 0x96,
	0x0a, 0xfd, 0xa7, 0x87, 0x13, 0x36, 0xb2, 0xa7, 0xed, 0xd9, 0x3b, 0x8a, 0xac, 0x21, 0x9b, 0xc5,
	0xfa, 0x09, 0x8f, 0x6a, 0xe2, 0xeb, 0x83, 0x89, 0x94, 0x87, 0x88, 0xe2, 0x46, 0xa0, 0x6a, 0xa1,
	0x55, 0xfd, 0x0b, 0x8a, 0x4b, 0xeb, 0x63, 0x86, 0xd2, 0x14, 0x67, 0xb9, 0xda, 0x1f, 0xd4, 0x82,
	0x29, 0xb1, 0x10, 0x63, 0x5c, 0x20, 0x41, 0x38, 0x53, 0xbb, 0xe6, 0x6f, 0x0d, 0x3c, 0xf0, 0x6a,
	0x4d, 0x4f, 0x49, 0xba, 0x1b, 0x52, 0xc6, 0x7b, 0x70, 0x4b, 0xd5, 0xd4, 0xd3, 0x86, 0xda, 0xa8,
	0x6b, 0x4f, 0x60, 0xfb, 0x4b, 0x83, 0x5b, 0x3c, 0x4f, 0xe6, 0x6f, 0x62, 0x67, 0x35, 0xd4, 0x78,
	0x07, 0x8e, 0x65, 0x85, 0x3d, 0x7d, 0xd8, 0x19, 0x75, 0x6d, 0xef, 0x60, 0xfc, 0xcb, 0x2a, 0xfd,
	0x0a, 0x5d, 0x21, 0xcd, 0x5f, 0x47, 0xe0, 0xd1, 0x3e, 0x76, 0x8c, 0xaf, 0x1a, 0xb8, 0xbb, 0xd5,
	0xe1, 0xaa, 0xe0, 0x53, 0x5b, 0x34, 0x3a, 0xaa, 0xde, 0xd9, 0x5e, 0xf5, 0xbe, 0xe2, 0x11, 0x7e,
	0xc1, 0x8a, 0x64, 0xef, 0xe0, 0xd9, 0x1d, 0x74, 0xed, 0x89, 0xf1, 0x49, 0x03, 0xbd, 0xa6, 0x79,
	0xed, 0xe9, 0x95, 0xd5, 0xe9, 0x0d, 0x56, 0xdd, 0xff, 0xe9, 0x4a, 0xf7, 0xcd, 0x2a, 0x95, 0xc6,
	0x76, 0x6f, 0xcd, 0xee, 0xa1, 0x9d, 0xcf, 0x8d, 0xcf, 0x1a, 0x18, 0xb4, 0x1d, 0xbf, 0x5e, 0x67,
	0xd8, 0x19, 0x9d, 0xda, 0xf3, 0x83, 0x0d, 0x5d, 0x14, 0x41, 0x8b, 0x27, 0xb5, 0x3b, 0xbb, 0x8f,
	0x9a, 0xb6, 0x8c, 0x00, 0x0c, 0x09, 0x0b, 0x69, 0x11, 0x61, 0x3f, 0xc2, 0x97, 0xa8, 0xa0, 0x62,
	0xdb, 0xdc, 0x51, 0x35, 0xc9, 0xfd, 0xda, 0x5c, 0x7d, 0x76, 0xe0, 0x98, 0x73, 0x3a, 0x47, 0xb4,
	0xc0, 0xb3, 0x81, 0x62, 0x4c, 0x24, 0xe2, 0xaa, 0x86, 0xf9, 0x53, 0x07, 0x0f, 0xf7, 0x98, 0x43,
	0xe3, 0x4b, 0xcb, 0x58, 0x65, 0x87, 0x8e, 0xd5, 0x36, 0xbf, 0x79, 0xaa, 0x76, 0xc7, 0xee, 0x18,
	0xaa, 0x67, 0xe0, 0x24, 0xe4, 0x05, 0x13, 0xe5, 0xd5, 0x55, 0x7a, 0xd3, 0xab, 0xce, 0x0c, 0xb6,
	0x3a, 0x73, 0x21, 0x32, 0xc2, 0x62, 0xd9, 0x9b, 0xae, 0xca, 0xf0, 0x4a, 0x80, 0x0b, 0x6e, 0x53,
	0xc4, 0xe2, 0x02, 0xc5, 0x58, 0x12, 0x3a, 0x7b, 0x10, 0x4e, 0xea, 0x94, 0x12, 0x31, 0xfe, 0xab,
	0x01, 0x33, 0xe4, 0xc9, 0x0d, 0x07, 0x7f, 0x6c, 0xb6, 0x5e, 0x54, 0xd3, 0x52, 0x67, 0xaa, 0xbd,
	0x55, 0x1f, 0x26, 0x18, 0xf3, 0x52, 0x03, 0xf2, 0x2c, 0xb6, 0x62, 0xcc, 0x2a, 0x17, 0xf5, 0x85,
	0x9b, 0x92, 0xbc, 0xe9, 0xbb, 0xf5, 0x44, 0xfe, 0x7c, 0xd3, 0x3b, 0xe7, 0xae, 0xfb, 0x5d, 0x3f,
	0x3b, 0x97, 0x30, 0x37, 0xca, 0xa1, 0x5c, 0x96, 0xab, 0xb9, 0x0d, 0xbd, 0x2a, 0xec, 0x47, 0x1d,
	0xb0, 0x70, 0xa3, 0x7c, 0xb1, 0x0e, 0x58, 0xcc, 0xed, 0x85, 0x0c, 0xf8, 0xa3, 0x9b, 0xf2, 0xa9,
	0xe3, 0xb8, 0x51, 0xee, 0x38, 0xeb, 0x10, 0xc7, 0x99, 0xdb, 0x8e, 0x23, 0x83, 0x82, 0xe3, 0xca,
	0xdd, 0xe3, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x99, 0xae, 0x1d, 0x9e, 0x54, 0x07, 0x00, 0x00,
}
