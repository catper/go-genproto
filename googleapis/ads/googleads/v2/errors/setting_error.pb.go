// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/errors/setting_error.proto

package errors

import (
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
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

// Enum describing possible setting errors.
type SettingErrorEnum_SettingError int32

const (
	// Enum unspecified.
	SettingErrorEnum_UNSPECIFIED SettingErrorEnum_SettingError = 0
	// The received error code is not known in this version.
	SettingErrorEnum_UNKNOWN SettingErrorEnum_SettingError = 1
	// The campaign setting is not available for this Google Ads account.
	SettingErrorEnum_SETTING_TYPE_IS_NOT_AVAILABLE SettingErrorEnum_SettingError = 3
	// The setting is not compatible with the campaign.
	SettingErrorEnum_SETTING_TYPE_IS_NOT_COMPATIBLE_WITH_CAMPAIGN SettingErrorEnum_SettingError = 4
	// The supplied TargetingSetting contains an invalid CriterionTypeGroup. See
	// CriterionTypeGroup documentation for CriterionTypeGroups allowed
	// in Campaign or AdGroup TargetingSettings.
	SettingErrorEnum_TARGETING_SETTING_CONTAINS_INVALID_CRITERION_TYPE_GROUP SettingErrorEnum_SettingError = 5
	// TargetingSetting must not explicitly
	// set any of the Demographic CriterionTypeGroups (AGE_RANGE, GENDER,
	// PARENT, INCOME_RANGE) to false (it's okay to not set them at all, in
	// which case the system will set them to true automatically).
	SettingErrorEnum_TARGETING_SETTING_DEMOGRAPHIC_CRITERION_TYPE_GROUPS_MUST_BE_SET_TO_TARGET_ALL SettingErrorEnum_SettingError = 6
	// TargetingSetting cannot change any of
	// the Demographic CriterionTypeGroups (AGE_RANGE, GENDER, PARENT,
	// INCOME_RANGE) from true to false.
	SettingErrorEnum_TARGETING_SETTING_CANNOT_CHANGE_TARGET_ALL_TO_FALSE_FOR_DEMOGRAPHIC_CRITERION_TYPE_GROUP SettingErrorEnum_SettingError = 7
	// At least one feed id should be present.
	SettingErrorEnum_DYNAMIC_SEARCH_ADS_SETTING_AT_LEAST_ONE_FEED_ID_MUST_BE_PRESENT SettingErrorEnum_SettingError = 8
	// The supplied DynamicSearchAdsSetting contains an invalid domain name.
	SettingErrorEnum_DYNAMIC_SEARCH_ADS_SETTING_CONTAINS_INVALID_DOMAIN_NAME SettingErrorEnum_SettingError = 9
	// The supplied DynamicSearchAdsSetting contains a subdomain name.
	SettingErrorEnum_DYNAMIC_SEARCH_ADS_SETTING_CONTAINS_SUBDOMAIN_NAME SettingErrorEnum_SettingError = 10
	// The supplied DynamicSearchAdsSetting contains an invalid language code.
	SettingErrorEnum_DYNAMIC_SEARCH_ADS_SETTING_CONTAINS_INVALID_LANGUAGE_CODE SettingErrorEnum_SettingError = 11
	// TargetingSettings in search campaigns should not have
	// CriterionTypeGroup.PLACEMENT set to targetAll.
	SettingErrorEnum_TARGET_ALL_IS_NOT_ALLOWED_FOR_PLACEMENT_IN_SEARCH_CAMPAIGN SettingErrorEnum_SettingError = 12
	// Duplicate description in universal app setting description field.
	SettingErrorEnum_UNIVERSAL_APP_CAMPAIGN_SETTING_DUPLICATE_DESCRIPTION SettingErrorEnum_SettingError = 13
	// Description line width is too long in universal app setting description
	// field.
	SettingErrorEnum_UNIVERSAL_APP_CAMPAIGN_SETTING_DESCRIPTION_LINE_WIDTH_TOO_LONG SettingErrorEnum_SettingError = 14
	// Universal app setting appId field cannot be modified for COMPLETE
	// campaigns.
	SettingErrorEnum_UNIVERSAL_APP_CAMPAIGN_SETTING_APP_ID_CANNOT_BE_MODIFIED SettingErrorEnum_SettingError = 15
	// YoutubeVideoMediaIds in universal app setting cannot exceed size limit.
	SettingErrorEnum_TOO_MANY_YOUTUBE_MEDIA_IDS_IN_UNIVERSAL_APP_CAMPAIGN SettingErrorEnum_SettingError = 16
	// ImageMediaIds in universal app setting cannot exceed size limit.
	SettingErrorEnum_TOO_MANY_IMAGE_MEDIA_IDS_IN_UNIVERSAL_APP_CAMPAIGN SettingErrorEnum_SettingError = 17
	// Media is incompatible for universal app campaign.
	SettingErrorEnum_MEDIA_INCOMPATIBLE_FOR_UNIVERSAL_APP_CAMPAIGN SettingErrorEnum_SettingError = 18
	// Too many exclamation marks in universal app campaign ad text ideas.
	SettingErrorEnum_TOO_MANY_EXCLAMATION_MARKS SettingErrorEnum_SettingError = 19
)

var SettingErrorEnum_SettingError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	3:  "SETTING_TYPE_IS_NOT_AVAILABLE",
	4:  "SETTING_TYPE_IS_NOT_COMPATIBLE_WITH_CAMPAIGN",
	5:  "TARGETING_SETTING_CONTAINS_INVALID_CRITERION_TYPE_GROUP",
	6:  "TARGETING_SETTING_DEMOGRAPHIC_CRITERION_TYPE_GROUPS_MUST_BE_SET_TO_TARGET_ALL",
	7:  "TARGETING_SETTING_CANNOT_CHANGE_TARGET_ALL_TO_FALSE_FOR_DEMOGRAPHIC_CRITERION_TYPE_GROUP",
	8:  "DYNAMIC_SEARCH_ADS_SETTING_AT_LEAST_ONE_FEED_ID_MUST_BE_PRESENT",
	9:  "DYNAMIC_SEARCH_ADS_SETTING_CONTAINS_INVALID_DOMAIN_NAME",
	10: "DYNAMIC_SEARCH_ADS_SETTING_CONTAINS_SUBDOMAIN_NAME",
	11: "DYNAMIC_SEARCH_ADS_SETTING_CONTAINS_INVALID_LANGUAGE_CODE",
	12: "TARGET_ALL_IS_NOT_ALLOWED_FOR_PLACEMENT_IN_SEARCH_CAMPAIGN",
	13: "UNIVERSAL_APP_CAMPAIGN_SETTING_DUPLICATE_DESCRIPTION",
	14: "UNIVERSAL_APP_CAMPAIGN_SETTING_DESCRIPTION_LINE_WIDTH_TOO_LONG",
	15: "UNIVERSAL_APP_CAMPAIGN_SETTING_APP_ID_CANNOT_BE_MODIFIED",
	16: "TOO_MANY_YOUTUBE_MEDIA_IDS_IN_UNIVERSAL_APP_CAMPAIGN",
	17: "TOO_MANY_IMAGE_MEDIA_IDS_IN_UNIVERSAL_APP_CAMPAIGN",
	18: "MEDIA_INCOMPATIBLE_FOR_UNIVERSAL_APP_CAMPAIGN",
	19: "TOO_MANY_EXCLAMATION_MARKS",
}

var SettingErrorEnum_SettingError_value = map[string]int32{
	"UNSPECIFIED":                   0,
	"UNKNOWN":                       1,
	"SETTING_TYPE_IS_NOT_AVAILABLE": 3,
	"SETTING_TYPE_IS_NOT_COMPATIBLE_WITH_CAMPAIGN":                                             4,
	"TARGETING_SETTING_CONTAINS_INVALID_CRITERION_TYPE_GROUP":                                  5,
	"TARGETING_SETTING_DEMOGRAPHIC_CRITERION_TYPE_GROUPS_MUST_BE_SET_TO_TARGET_ALL":            6,
	"TARGETING_SETTING_CANNOT_CHANGE_TARGET_ALL_TO_FALSE_FOR_DEMOGRAPHIC_CRITERION_TYPE_GROUP": 7,
	"DYNAMIC_SEARCH_ADS_SETTING_AT_LEAST_ONE_FEED_ID_MUST_BE_PRESENT":                          8,
	"DYNAMIC_SEARCH_ADS_SETTING_CONTAINS_INVALID_DOMAIN_NAME":                                  9,
	"DYNAMIC_SEARCH_ADS_SETTING_CONTAINS_SUBDOMAIN_NAME":                                       10,
	"DYNAMIC_SEARCH_ADS_SETTING_CONTAINS_INVALID_LANGUAGE_CODE":                                11,
	"TARGET_ALL_IS_NOT_ALLOWED_FOR_PLACEMENT_IN_SEARCH_CAMPAIGN":                               12,
	"UNIVERSAL_APP_CAMPAIGN_SETTING_DUPLICATE_DESCRIPTION":                                     13,
	"UNIVERSAL_APP_CAMPAIGN_SETTING_DESCRIPTION_LINE_WIDTH_TOO_LONG":                           14,
	"UNIVERSAL_APP_CAMPAIGN_SETTING_APP_ID_CANNOT_BE_MODIFIED":                                 15,
	"TOO_MANY_YOUTUBE_MEDIA_IDS_IN_UNIVERSAL_APP_CAMPAIGN":                                     16,
	"TOO_MANY_IMAGE_MEDIA_IDS_IN_UNIVERSAL_APP_CAMPAIGN":                                       17,
	"MEDIA_INCOMPATIBLE_FOR_UNIVERSAL_APP_CAMPAIGN":                                            18,
	"TOO_MANY_EXCLAMATION_MARKS":                                                               19,
}

func (x SettingErrorEnum_SettingError) String() string {
	return proto.EnumName(SettingErrorEnum_SettingError_name, int32(x))
}

func (SettingErrorEnum_SettingError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6f7dbbb80c19998c, []int{0, 0}
}

// Container for enum describing possible setting errors.
type SettingErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SettingErrorEnum) Reset()         { *m = SettingErrorEnum{} }
func (m *SettingErrorEnum) String() string { return proto.CompactTextString(m) }
func (*SettingErrorEnum) ProtoMessage()    {}
func (*SettingErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f7dbbb80c19998c, []int{0}
}

func (m *SettingErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SettingErrorEnum.Unmarshal(m, b)
}
func (m *SettingErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SettingErrorEnum.Marshal(b, m, deterministic)
}
func (m *SettingErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SettingErrorEnum.Merge(m, src)
}
func (m *SettingErrorEnum) XXX_Size() int {
	return xxx_messageInfo_SettingErrorEnum.Size(m)
}
func (m *SettingErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_SettingErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_SettingErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.errors.SettingErrorEnum_SettingError", SettingErrorEnum_SettingError_name, SettingErrorEnum_SettingError_value)
	proto.RegisterType((*SettingErrorEnum)(nil), "google.ads.googleads.v2.errors.SettingErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/errors/setting_error.proto", fileDescriptor_6f7dbbb80c19998c)
}

var fileDescriptor_6f7dbbb80c19998c = []byte{
	// 695 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xd1, 0x8e, 0xdb, 0x44,
	0x14, 0x86, 0xd9, 0x2d, 0xb4, 0xcb, 0x6c, 0xa1, 0xd3, 0xe1, 0xae, 0x82, 0x95, 0xd8, 0x6b, 0x70,
	0x20, 0x20, 0x28, 0x2e, 0x14, 0x1d, 0x7b, 0xce, 0x3a, 0xa3, 0x8e, 0x67, 0x8c, 0x67, 0x9c, 0x6d,
	0x50, 0xa4, 0xa3, 0x40, 0xa2, 0x28, 0x52, 0x1b, 0xaf, 0xe2, 0xd0, 0x07, 0xe2, 0x0e, 0x1e, 0x82,
	0x07, 0xe0, 0x51, 0x10, 0x0f, 0x81, 0xc6, 0x6e, 0xdc, 0x48, 0x4d, 0xbb, 0xcb, 0x55, 0x26, 0xe3,
	0xff, 0xfb, 0xcf, 0xf8, 0xf7, 0x99, 0xc3, 0x86, 0xcb, 0xba, 0x5e, 0x3e, 0x5b, 0x0c, 0x66, 0xf3,
	0x66, 0xd0, 0x2d, 0xc3, 0xea, 0xc5, 0x70, 0xb0, 0xd8, 0x6c, 0xea, 0x4d, 0x33, 0x68, 0x16, 0xdb,
	0xed, 0x6a, 0xbd, 0xa4, 0xf6, 0x6f, 0x74, 0xb5, 0xa9, 0xb7, 0xb5, 0x38, 0xeb, 0x84, 0xd1, 0x6c,
	0xde, 0x44, 0x3d, 0x13, 0xbd, 0x18, 0x46, 0x1d, 0xf3, 0xe0, 0xe3, 0x9d, 0xe7, 0xd5, 0x6a, 0x30,
	0x5b, 0xaf, 0xeb, 0xed, 0x6c, 0xbb, 0xaa, 0xd7, 0x4d, 0x47, 0x9f, 0xff, 0x75, 0xc2, 0xb8, 0xeb,
	0x5c, 0x31, 0xe8, 0x71, 0xfd, 0xdb, 0xf3, 0xf3, 0x3f, 0x4e, 0xd8, 0xdd, 0xfd, 0x4d, 0x71, 0x8f,
	0x9d, 0x56, 0xc6, 0x15, 0x98, 0xaa, 0x0b, 0x85, 0x92, 0xbf, 0x23, 0x4e, 0xd9, 0x9d, 0xca, 0x3c,
	0x31, 0xf6, 0xd2, 0xf0, 0x23, 0xf1, 0x29, 0xfb, 0xc4, 0xa1, 0xf7, 0xca, 0x64, 0xe4, 0x27, 0x05,
	0x92, 0x72, 0x64, 0xac, 0x27, 0x18, 0x83, 0xd2, 0x90, 0x68, 0xe4, 0xb7, 0xc4, 0x17, 0xec, 0xb3,
	0x43, 0x92, 0xd4, 0xe6, 0x05, 0x78, 0x95, 0x68, 0xa4, 0x4b, 0xe5, 0x47, 0x94, 0x42, 0x5e, 0x80,
	0xca, 0x0c, 0x7f, 0x57, 0x3c, 0x62, 0xdf, 0x7a, 0x28, 0x33, 0x6c, 0x99, 0x1d, 0x9b, 0x5a, 0xe3,
	0x41, 0x19, 0x47, 0xca, 0x8c, 0x41, 0x2b, 0x49, 0x69, 0xa9, 0x3c, 0x96, 0xca, 0x9a, 0xce, 0x36,
	0x2b, 0x6d, 0x55, 0xf0, 0xf7, 0xc4, 0x4f, 0x2c, 0x7f, 0x1d, 0x96, 0x98, 0xdb, 0xac, 0x84, 0x62,
	0xa4, 0xd2, 0x83, 0x9c, 0xa3, 0xbc, 0x72, 0x9e, 0x12, 0x0c, 0x04, 0x79, 0x4b, 0x9d, 0x05, 0x81,
	0xd6, 0xfc, 0xb6, 0x98, 0xb2, 0xa7, 0x07, 0xce, 0x03, 0xa6, 0x7d, 0x8d, 0x11, 0x98, 0x0c, 0xf7,
	0xf4, 0x81, 0xbe, 0x00, 0xed, 0x90, 0x2e, 0x6c, 0x79, 0x6d, 0x61, 0x7e, 0x47, 0xa4, 0xec, 0x47,
	0x39, 0x31, 0x90, 0xab, 0x94, 0x1c, 0x42, 0x99, 0x8e, 0x08, 0xa4, 0xeb, 0xcb, 0x80, 0x27, 0x8d,
	0xe0, 0x3c, 0x59, 0x83, 0x74, 0x81, 0x28, 0x49, 0xc9, 0xfe, 0xb0, 0x45, 0x89, 0x0e, 0x8d, 0xe7,
	0x27, 0x21, 0xb2, 0xb7, 0x98, 0xbc, 0x96, 0x9d, 0xb4, 0x39, 0x28, 0x43, 0x06, 0x72, 0xe4, 0xef,
	0x8b, 0x6f, 0xd8, 0xf0, 0x26, 0xb0, 0xab, 0x92, 0x7d, 0x8e, 0x89, 0x1f, 0xd8, 0x77, 0xff, 0xa7,
	0xa8, 0x06, 0x93, 0x55, 0x90, 0x21, 0xa5, 0x56, 0x22, 0x3f, 0x15, 0x8f, 0x59, 0xbc, 0x17, 0xdb,
	0xae, 0x73, 0xb4, 0xb6, 0x97, 0x28, 0xdb, 0xf0, 0x0a, 0x0d, 0x29, 0xe6, 0x68, 0x3c, 0x29, 0xb3,
	0xab, 0xd0, 0xb7, 0xc9, 0x5d, 0xf1, 0x90, 0x7d, 0x5d, 0x19, 0x35, 0xc6, 0xd2, 0x81, 0x26, 0x28,
	0x8a, 0xfe, 0xd9, 0xab, 0xcf, 0x5e, 0x15, 0x5a, 0xa5, 0xe0, 0x91, 0x24, 0xba, 0xb4, 0x54, 0x85,
	0x57, 0xd6, 0xf0, 0x0f, 0x44, 0xc2, 0x1e, 0x5f, 0x47, 0xbe, 0xd2, 0x93, 0x56, 0x26, 0xf4, 0xa8,
	0xf4, 0x23, 0xf2, 0xd6, 0x92, 0xb6, 0x26, 0xe3, 0x1f, 0x8a, 0xef, 0xd9, 0xc3, 0x6b, 0x3c, 0xc2,
	0x66, 0xe8, 0xd3, 0xae, 0x51, 0x12, 0xa4, 0xdc, 0xca, 0xee, 0x12, 0xdd, 0x0b, 0x67, 0x0f, 0x5e,
	0x39, 0x98, 0x09, 0x4d, 0x6c, 0xe5, 0xab, 0xf0, 0x18, 0xa5, 0x02, 0x52, 0x32, 0x24, 0x46, 0x87,
	0xbd, 0x39, 0x0f, 0x1f, 0xab, 0x27, 0x55, 0x1e, 0xe2, 0xbc, 0x09, 0x77, 0x5f, 0x7c, 0xc9, 0x3e,
	0x7f, 0x29, 0x34, 0x7b, 0x57, 0x2f, 0xc4, 0xfc, 0x06, 0x44, 0x88, 0x33, 0xf6, 0xa0, 0x2f, 0x85,
	0x4f, 0x53, 0x0d, 0x39, 0xb4, 0x81, 0xe4, 0x50, 0x3e, 0x71, 0xfc, 0xa3, 0xe4, 0xdf, 0x23, 0x76,
	0xfe, 0x6b, 0xfd, 0x3c, 0x7a, 0xfb, 0x14, 0x4a, 0xee, 0xef, 0xcf, 0x93, 0x22, 0x8c, 0x9e, 0xe2,
	0xe8, 0x67, 0xf9, 0x12, 0x5a, 0xd6, 0xcf, 0x66, 0xeb, 0x65, 0x54, 0x6f, 0x96, 0x83, 0xe5, 0x62,
	0xdd, 0x0e, 0xa6, 0xdd, 0xf8, 0xbb, 0x5a, 0x35, 0x6f, 0x9a, 0x86, 0x8f, 0xba, 0x9f, 0xdf, 0x8f,
	0x6f, 0x65, 0x00, 0x7f, 0x1e, 0x9f, 0x65, 0x9d, 0x19, 0xcc, 0x9b, 0xa8, 0x5b, 0x86, 0xd5, 0x78,
	0x18, 0xb5, 0x25, 0x9b, 0xbf, 0x77, 0x82, 0x29, 0xcc, 0x9b, 0x69, 0x2f, 0x98, 0x8e, 0x87, 0xd3,
	0x4e, 0xf0, 0xcf, 0xf1, 0x79, 0xb7, 0x1b, 0xc7, 0x30, 0x6f, 0xe2, 0xb8, 0x97, 0xc4, 0xf1, 0x78,
	0x18, 0xc7, 0x9d, 0xe8, 0x97, 0xdb, 0xed, 0xe9, 0xbe, 0xfa, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x62,
	0x5f, 0x0a, 0xee, 0xaa, 0x05, 0x00, 0x00,
}
