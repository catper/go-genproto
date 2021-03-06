// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/feed_item.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
	wrappers "github.com/catper/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v3/common"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v3/enums"
	errors "google.golang.org/genproto/googleapis/ads/googleads/v3/errors"
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

// A feed item.
type FeedItem struct {
	// The resource name of the feed item.
	// Feed item resource names have the form:
	//
	// `customers/{customer_id}/feedItems/{feed_id}~{feed_item_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The feed to which this feed item belongs.
	Feed *wrappers.StringValue `protobuf:"bytes,2,opt,name=feed,proto3" json:"feed,omitempty"`
	// The ID of this feed item.
	Id *wrappers.Int64Value `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// Start time in which this feed item is effective and can begin serving. The
	// time is in the customer's time zone.
	// The format is "YYYY-MM-DD HH:MM:SS".
	// Examples: "2018-03-05 09:15:00" or "2018-02-01 14:34:30"
	StartDateTime *wrappers.StringValue `protobuf:"bytes,4,opt,name=start_date_time,json=startDateTime,proto3" json:"start_date_time,omitempty"`
	// End time in which this feed item is no longer effective and will stop
	// serving. The time is in the customer's time zone.
	// The format is "YYYY-MM-DD HH:MM:SS".
	// Examples: "2018-03-05 09:15:00" or "2018-02-01 14:34:30"
	EndDateTime *wrappers.StringValue `protobuf:"bytes,5,opt,name=end_date_time,json=endDateTime,proto3" json:"end_date_time,omitempty"`
	// The feed item's attribute values.
	AttributeValues []*FeedItemAttributeValue `protobuf:"bytes,6,rep,name=attribute_values,json=attributeValues,proto3" json:"attribute_values,omitempty"`
	// Geo targeting restriction specifies the type of location that can be used
	// for targeting.
	GeoTargetingRestriction enums.GeoTargetingRestrictionEnum_GeoTargetingRestriction `protobuf:"varint,7,opt,name=geo_targeting_restriction,json=geoTargetingRestriction,proto3,enum=google.ads.googleads.v3.enums.GeoTargetingRestrictionEnum_GeoTargetingRestriction" json:"geo_targeting_restriction,omitempty"`
	// The list of mappings used to substitute custom parameter tags in a
	// `tracking_url_template`, `final_urls`, or `mobile_final_urls`.
	UrlCustomParameters []*common.CustomParameter `protobuf:"bytes,8,rep,name=url_custom_parameters,json=urlCustomParameters,proto3" json:"url_custom_parameters,omitempty"`
	// Status of the feed item.
	// This field is read-only.
	Status enums.FeedItemStatusEnum_FeedItemStatus `protobuf:"varint,9,opt,name=status,proto3,enum=google.ads.googleads.v3.enums.FeedItemStatusEnum_FeedItemStatus" json:"status,omitempty"`
	// List of info about a feed item's validation and approval state for active
	// feed mappings. There will be an entry in the list for each type of feed
	// mapping associated with the feed, e.g. a feed with a sitelink and a call
	// feed mapping would cause every feed item associated with that feed to have
	// an entry in this list for both sitelink and call.
	// This field is read-only.
	PolicyInfos          []*FeedItemPlaceholderPolicyInfo `protobuf:"bytes,10,rep,name=policy_infos,json=policyInfos,proto3" json:"policy_infos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *FeedItem) Reset()         { *m = FeedItem{} }
func (m *FeedItem) String() string { return proto.CompactTextString(m) }
func (*FeedItem) ProtoMessage()    {}
func (*FeedItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa81ffed50546ca6, []int{0}
}

func (m *FeedItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedItem.Unmarshal(m, b)
}
func (m *FeedItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedItem.Marshal(b, m, deterministic)
}
func (m *FeedItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedItem.Merge(m, src)
}
func (m *FeedItem) XXX_Size() int {
	return xxx_messageInfo_FeedItem.Size(m)
}
func (m *FeedItem) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedItem.DiscardUnknown(m)
}

var xxx_messageInfo_FeedItem proto.InternalMessageInfo

func (m *FeedItem) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *FeedItem) GetFeed() *wrappers.StringValue {
	if m != nil {
		return m.Feed
	}
	return nil
}

func (m *FeedItem) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *FeedItem) GetStartDateTime() *wrappers.StringValue {
	if m != nil {
		return m.StartDateTime
	}
	return nil
}

func (m *FeedItem) GetEndDateTime() *wrappers.StringValue {
	if m != nil {
		return m.EndDateTime
	}
	return nil
}

func (m *FeedItem) GetAttributeValues() []*FeedItemAttributeValue {
	if m != nil {
		return m.AttributeValues
	}
	return nil
}

func (m *FeedItem) GetGeoTargetingRestriction() enums.GeoTargetingRestrictionEnum_GeoTargetingRestriction {
	if m != nil {
		return m.GeoTargetingRestriction
	}
	return enums.GeoTargetingRestrictionEnum_UNSPECIFIED
}

func (m *FeedItem) GetUrlCustomParameters() []*common.CustomParameter {
	if m != nil {
		return m.UrlCustomParameters
	}
	return nil
}

func (m *FeedItem) GetStatus() enums.FeedItemStatusEnum_FeedItemStatus {
	if m != nil {
		return m.Status
	}
	return enums.FeedItemStatusEnum_UNSPECIFIED
}

func (m *FeedItem) GetPolicyInfos() []*FeedItemPlaceholderPolicyInfo {
	if m != nil {
		return m.PolicyInfos
	}
	return nil
}

// A feed item attribute value.
type FeedItemAttributeValue struct {
	// Id of the feed attribute for which the value is associated with.
	FeedAttributeId *wrappers.Int64Value `protobuf:"bytes,1,opt,name=feed_attribute_id,json=feedAttributeId,proto3" json:"feed_attribute_id,omitempty"`
	// Int64 value. Should be set if feed_attribute_id refers to a feed attribute
	// of type INT64.
	IntegerValue *wrappers.Int64Value `protobuf:"bytes,2,opt,name=integer_value,json=integerValue,proto3" json:"integer_value,omitempty"`
	// Bool value. Should be set if feed_attribute_id refers to a feed attribute
	// of type BOOLEAN.
	BooleanValue *wrappers.BoolValue `protobuf:"bytes,3,opt,name=boolean_value,json=booleanValue,proto3" json:"boolean_value,omitempty"`
	// String value. Should be set if feed_attribute_id refers to a feed attribute
	// of type STRING, URL or DATE_TIME.
	// For STRING the maximum length is 1500 characters. For URL the maximum
	// length is 2076 characters. For DATE_TIME the string must be in the format
	// "YYYYMMDD HHMMSS".
	StringValue *wrappers.StringValue `protobuf:"bytes,4,opt,name=string_value,json=stringValue,proto3" json:"string_value,omitempty"`
	// Double value. Should be set if feed_attribute_id refers to a feed attribute
	// of type DOUBLE.
	DoubleValue *wrappers.DoubleValue `protobuf:"bytes,5,opt,name=double_value,json=doubleValue,proto3" json:"double_value,omitempty"`
	// Price value. Should be set if feed_attribute_id refers to a feed attribute
	// of type PRICE.
	PriceValue *common.Money `protobuf:"bytes,6,opt,name=price_value,json=priceValue,proto3" json:"price_value,omitempty"`
	// Repeated int64 value. Should be set if feed_attribute_id refers to a feed
	// attribute of type INT64_LIST.
	IntegerValues []*wrappers.Int64Value `protobuf:"bytes,7,rep,name=integer_values,json=integerValues,proto3" json:"integer_values,omitempty"`
	// Repeated bool value. Should be set if feed_attribute_id refers to a feed
	// attribute of type BOOLEAN_LIST.
	BooleanValues []*wrappers.BoolValue `protobuf:"bytes,8,rep,name=boolean_values,json=booleanValues,proto3" json:"boolean_values,omitempty"`
	// Repeated string value. Should be set if feed_attribute_id refers to a feed
	// attribute of type STRING_LIST, URL_LIST or DATE_TIME_LIST.
	// For STRING_LIST and URL_LIST the total size of the list in bytes may not
	// exceed 3000. For DATE_TIME_LIST the number of elements may not exceed 200.
	//
	// For STRING_LIST the maximum length of each string element is 1500
	// characters. For URL_LIST the maximum length is 2076 characters. For
	// DATE_TIME the format of the string must be the same as start and end time
	// for the feed item.
	StringValues []*wrappers.StringValue `protobuf:"bytes,9,rep,name=string_values,json=stringValues,proto3" json:"string_values,omitempty"`
	// Repeated double value. Should be set if feed_attribute_id refers to a feed
	// attribute of type DOUBLE_LIST.
	DoubleValues         []*wrappers.DoubleValue `protobuf:"bytes,10,rep,name=double_values,json=doubleValues,proto3" json:"double_values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *FeedItemAttributeValue) Reset()         { *m = FeedItemAttributeValue{} }
func (m *FeedItemAttributeValue) String() string { return proto.CompactTextString(m) }
func (*FeedItemAttributeValue) ProtoMessage()    {}
func (*FeedItemAttributeValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa81ffed50546ca6, []int{1}
}

func (m *FeedItemAttributeValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedItemAttributeValue.Unmarshal(m, b)
}
func (m *FeedItemAttributeValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedItemAttributeValue.Marshal(b, m, deterministic)
}
func (m *FeedItemAttributeValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedItemAttributeValue.Merge(m, src)
}
func (m *FeedItemAttributeValue) XXX_Size() int {
	return xxx_messageInfo_FeedItemAttributeValue.Size(m)
}
func (m *FeedItemAttributeValue) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedItemAttributeValue.DiscardUnknown(m)
}

var xxx_messageInfo_FeedItemAttributeValue proto.InternalMessageInfo

func (m *FeedItemAttributeValue) GetFeedAttributeId() *wrappers.Int64Value {
	if m != nil {
		return m.FeedAttributeId
	}
	return nil
}

func (m *FeedItemAttributeValue) GetIntegerValue() *wrappers.Int64Value {
	if m != nil {
		return m.IntegerValue
	}
	return nil
}

func (m *FeedItemAttributeValue) GetBooleanValue() *wrappers.BoolValue {
	if m != nil {
		return m.BooleanValue
	}
	return nil
}

func (m *FeedItemAttributeValue) GetStringValue() *wrappers.StringValue {
	if m != nil {
		return m.StringValue
	}
	return nil
}

func (m *FeedItemAttributeValue) GetDoubleValue() *wrappers.DoubleValue {
	if m != nil {
		return m.DoubleValue
	}
	return nil
}

func (m *FeedItemAttributeValue) GetPriceValue() *common.Money {
	if m != nil {
		return m.PriceValue
	}
	return nil
}

func (m *FeedItemAttributeValue) GetIntegerValues() []*wrappers.Int64Value {
	if m != nil {
		return m.IntegerValues
	}
	return nil
}

func (m *FeedItemAttributeValue) GetBooleanValues() []*wrappers.BoolValue {
	if m != nil {
		return m.BooleanValues
	}
	return nil
}

func (m *FeedItemAttributeValue) GetStringValues() []*wrappers.StringValue {
	if m != nil {
		return m.StringValues
	}
	return nil
}

func (m *FeedItemAttributeValue) GetDoubleValues() []*wrappers.DoubleValue {
	if m != nil {
		return m.DoubleValues
	}
	return nil
}

// Policy, validation, and quality approval info for a feed item for the
// specified placeholder type.
type FeedItemPlaceholderPolicyInfo struct {
	// The placeholder type.
	PlaceholderTypeEnum enums.PlaceholderTypeEnum_PlaceholderType `protobuf:"varint,10,opt,name=placeholder_type_enum,json=placeholderTypeEnum,proto3,enum=google.ads.googleads.v3.enums.PlaceholderTypeEnum_PlaceholderType" json:"placeholder_type_enum,omitempty"`
	// The FeedMapping that contains the placeholder type.
	FeedMappingResourceName *wrappers.StringValue `protobuf:"bytes,2,opt,name=feed_mapping_resource_name,json=feedMappingResourceName,proto3" json:"feed_mapping_resource_name,omitempty"`
	// Where the placeholder type is in the review process.
	ReviewStatus enums.PolicyReviewStatusEnum_PolicyReviewStatus `protobuf:"varint,3,opt,name=review_status,json=reviewStatus,proto3,enum=google.ads.googleads.v3.enums.PolicyReviewStatusEnum_PolicyReviewStatus" json:"review_status,omitempty"`
	// The overall approval status of the placeholder type, calculated based on
	// the status of its individual policy topic entries.
	ApprovalStatus enums.PolicyApprovalStatusEnum_PolicyApprovalStatus `protobuf:"varint,4,opt,name=approval_status,json=approvalStatus,proto3,enum=google.ads.googleads.v3.enums.PolicyApprovalStatusEnum_PolicyApprovalStatus" json:"approval_status,omitempty"`
	// The list of policy findings for the placeholder type.
	PolicyTopicEntries []*common.PolicyTopicEntry `protobuf:"bytes,5,rep,name=policy_topic_entries,json=policyTopicEntries,proto3" json:"policy_topic_entries,omitempty"`
	// The validation status of the palceholder type.
	ValidationStatus enums.FeedItemValidationStatusEnum_FeedItemValidationStatus `protobuf:"varint,6,opt,name=validation_status,json=validationStatus,proto3,enum=google.ads.googleads.v3.enums.FeedItemValidationStatusEnum_FeedItemValidationStatus" json:"validation_status,omitempty"`
	// List of placeholder type validation errors.
	ValidationErrors []*FeedItemValidationError `protobuf:"bytes,7,rep,name=validation_errors,json=validationErrors,proto3" json:"validation_errors,omitempty"`
	// Placeholder type quality evaluation approval status.
	QualityApprovalStatus enums.FeedItemQualityApprovalStatusEnum_FeedItemQualityApprovalStatus `protobuf:"varint,8,opt,name=quality_approval_status,json=qualityApprovalStatus,proto3,enum=google.ads.googleads.v3.enums.FeedItemQualityApprovalStatusEnum_FeedItemQualityApprovalStatus" json:"quality_approval_status,omitempty"`
	// List of placeholder type quality evaluation disapproval reasons.
	QualityDisapprovalReasons []enums.FeedItemQualityDisapprovalReasonEnum_FeedItemQualityDisapprovalReason `protobuf:"varint,9,rep,packed,name=quality_disapproval_reasons,json=qualityDisapprovalReasons,proto3,enum=google.ads.googleads.v3.enums.FeedItemQualityDisapprovalReasonEnum_FeedItemQualityDisapprovalReason" json:"quality_disapproval_reasons,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}                                                                      `json:"-"`
	XXX_unrecognized          []byte                                                                        `json:"-"`
	XXX_sizecache             int32                                                                         `json:"-"`
}

func (m *FeedItemPlaceholderPolicyInfo) Reset()         { *m = FeedItemPlaceholderPolicyInfo{} }
func (m *FeedItemPlaceholderPolicyInfo) String() string { return proto.CompactTextString(m) }
func (*FeedItemPlaceholderPolicyInfo) ProtoMessage()    {}
func (*FeedItemPlaceholderPolicyInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa81ffed50546ca6, []int{2}
}

func (m *FeedItemPlaceholderPolicyInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedItemPlaceholderPolicyInfo.Unmarshal(m, b)
}
func (m *FeedItemPlaceholderPolicyInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedItemPlaceholderPolicyInfo.Marshal(b, m, deterministic)
}
func (m *FeedItemPlaceholderPolicyInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedItemPlaceholderPolicyInfo.Merge(m, src)
}
func (m *FeedItemPlaceholderPolicyInfo) XXX_Size() int {
	return xxx_messageInfo_FeedItemPlaceholderPolicyInfo.Size(m)
}
func (m *FeedItemPlaceholderPolicyInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedItemPlaceholderPolicyInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FeedItemPlaceholderPolicyInfo proto.InternalMessageInfo

func (m *FeedItemPlaceholderPolicyInfo) GetPlaceholderTypeEnum() enums.PlaceholderTypeEnum_PlaceholderType {
	if m != nil {
		return m.PlaceholderTypeEnum
	}
	return enums.PlaceholderTypeEnum_UNSPECIFIED
}

func (m *FeedItemPlaceholderPolicyInfo) GetFeedMappingResourceName() *wrappers.StringValue {
	if m != nil {
		return m.FeedMappingResourceName
	}
	return nil
}

func (m *FeedItemPlaceholderPolicyInfo) GetReviewStatus() enums.PolicyReviewStatusEnum_PolicyReviewStatus {
	if m != nil {
		return m.ReviewStatus
	}
	return enums.PolicyReviewStatusEnum_UNSPECIFIED
}

func (m *FeedItemPlaceholderPolicyInfo) GetApprovalStatus() enums.PolicyApprovalStatusEnum_PolicyApprovalStatus {
	if m != nil {
		return m.ApprovalStatus
	}
	return enums.PolicyApprovalStatusEnum_UNSPECIFIED
}

func (m *FeedItemPlaceholderPolicyInfo) GetPolicyTopicEntries() []*common.PolicyTopicEntry {
	if m != nil {
		return m.PolicyTopicEntries
	}
	return nil
}

func (m *FeedItemPlaceholderPolicyInfo) GetValidationStatus() enums.FeedItemValidationStatusEnum_FeedItemValidationStatus {
	if m != nil {
		return m.ValidationStatus
	}
	return enums.FeedItemValidationStatusEnum_UNSPECIFIED
}

func (m *FeedItemPlaceholderPolicyInfo) GetValidationErrors() []*FeedItemValidationError {
	if m != nil {
		return m.ValidationErrors
	}
	return nil
}

func (m *FeedItemPlaceholderPolicyInfo) GetQualityApprovalStatus() enums.FeedItemQualityApprovalStatusEnum_FeedItemQualityApprovalStatus {
	if m != nil {
		return m.QualityApprovalStatus
	}
	return enums.FeedItemQualityApprovalStatusEnum_UNSPECIFIED
}

func (m *FeedItemPlaceholderPolicyInfo) GetQualityDisapprovalReasons() []enums.FeedItemQualityDisapprovalReasonEnum_FeedItemQualityDisapprovalReason {
	if m != nil {
		return m.QualityDisapprovalReasons
	}
	return nil
}

// Stores a validation error and the set of offending feed attributes which
// together are responsible for causing a feed item validation error.
type FeedItemValidationError struct {
	// Error code indicating what validation error was triggered. The description
	// of the error can be found in the 'description' field.
	ValidationError errors.FeedItemValidationErrorEnum_FeedItemValidationError `protobuf:"varint,1,opt,name=validation_error,json=validationError,proto3,enum=google.ads.googleads.v3.errors.FeedItemValidationErrorEnum_FeedItemValidationError" json:"validation_error,omitempty"`
	// The description of the validation error.
	Description *wrappers.StringValue `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Set of feed attributes in the feed item flagged during validation. If
	// empty, no specific feed attributes can be associated with the error
	// (e.g. error across the entire feed item).
	FeedAttributeIds []*wrappers.Int64Value `protobuf:"bytes,3,rep,name=feed_attribute_ids,json=feedAttributeIds,proto3" json:"feed_attribute_ids,omitempty"`
	// Any extra information related to this error which is not captured by
	// validation_error and feed_attribute_id (e.g. placeholder field IDs when
	// feed_attribute_id is not mapped). Note that extra_info is not localized.
	ExtraInfo            *wrappers.StringValue `protobuf:"bytes,5,opt,name=extra_info,json=extraInfo,proto3" json:"extra_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *FeedItemValidationError) Reset()         { *m = FeedItemValidationError{} }
func (m *FeedItemValidationError) String() string { return proto.CompactTextString(m) }
func (*FeedItemValidationError) ProtoMessage()    {}
func (*FeedItemValidationError) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa81ffed50546ca6, []int{3}
}

func (m *FeedItemValidationError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedItemValidationError.Unmarshal(m, b)
}
func (m *FeedItemValidationError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedItemValidationError.Marshal(b, m, deterministic)
}
func (m *FeedItemValidationError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedItemValidationError.Merge(m, src)
}
func (m *FeedItemValidationError) XXX_Size() int {
	return xxx_messageInfo_FeedItemValidationError.Size(m)
}
func (m *FeedItemValidationError) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedItemValidationError.DiscardUnknown(m)
}

var xxx_messageInfo_FeedItemValidationError proto.InternalMessageInfo

func (m *FeedItemValidationError) GetValidationError() errors.FeedItemValidationErrorEnum_FeedItemValidationError {
	if m != nil {
		return m.ValidationError
	}
	return errors.FeedItemValidationErrorEnum_UNSPECIFIED
}

func (m *FeedItemValidationError) GetDescription() *wrappers.StringValue {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *FeedItemValidationError) GetFeedAttributeIds() []*wrappers.Int64Value {
	if m != nil {
		return m.FeedAttributeIds
	}
	return nil
}

func (m *FeedItemValidationError) GetExtraInfo() *wrappers.StringValue {
	if m != nil {
		return m.ExtraInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*FeedItem)(nil), "google.ads.googleads.v3.resources.FeedItem")
	proto.RegisterType((*FeedItemAttributeValue)(nil), "google.ads.googleads.v3.resources.FeedItemAttributeValue")
	proto.RegisterType((*FeedItemPlaceholderPolicyInfo)(nil), "google.ads.googleads.v3.resources.FeedItemPlaceholderPolicyInfo")
	proto.RegisterType((*FeedItemValidationError)(nil), "google.ads.googleads.v3.resources.FeedItemValidationError")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/feed_item.proto", fileDescriptor_aa81ffed50546ca6)
}

var fileDescriptor_aa81ffed50546ca6 = []byte{
	// 1275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x57, 0xdf, 0x6e, 0xdc, 0xc4,
	0x17, 0xd6, 0x6e, 0xda, 0x34, 0x9d, 0xfd, 0x93, 0x74, 0xfa, 0xeb, 0x2f, 0x6e, 0x5a, 0x50, 0x1a,
	0x54, 0x29, 0xa2, 0x92, 0x1d, 0x76, 0x0b, 0x82, 0xad, 0xa0, 0xd9, 0x90, 0x34, 0x0d, 0xa2, 0x10,
	0x9c, 0x28, 0x02, 0x14, 0x61, 0xcd, 0xda, 0x13, 0x33, 0x92, 0xed, 0x71, 0x66, 0xc6, 0x5b, 0x56,
	0x55, 0x91, 0xb8, 0xe0, 0x96, 0x4b, 0x1e, 0x00, 0x71, 0xc5, 0x2b, 0xf0, 0x06, 0xbc, 0x05, 0x97,
	0xf4, 0x11, 0xb8, 0x42, 0x1e, 0x8f, 0xbd, 0xde, 0x3f, 0x5e, 0x3b, 0x77, 0xe3, 0x99, 0xf3, 0x7d,
	0xe7, 0xcc, 0x99, 0xf9, 0xce, 0x19, 0x83, 0xf7, 0x5c, 0x4a, 0x5d, 0x0f, 0x1b, 0xc8, 0xe1, 0x46,
	0x32, 0x8c, 0x47, 0xc3, 0xae, 0xc1, 0x30, 0xa7, 0x11, 0xb3, 0x31, 0x37, 0x2e, 0x30, 0x76, 0x2c,
	0x22, 0xb0, 0xaf, 0x87, 0x8c, 0x0a, 0x0a, 0x1f, 0x24, 0x76, 0x3a, 0x72, 0xb8, 0x9e, 0x41, 0xf4,
	0x61, 0x57, 0xcf, 0x20, 0x1b, 0xef, 0x17, 0xb1, 0xda, 0xd4, 0xf7, 0x69, 0x60, 0xd8, 0x11, 0x17,
	0xd4, 0xb7, 0x42, 0xc4, 0x90, 0x8f, 0x05, 0x66, 0x09, 0xf3, 0xc6, 0x4e, 0x09, 0x4c, 0x46, 0x92,
	0x8c, 0x15, 0xe2, 0x51, 0x09, 0x22, 0xa4, 0x1e, 0xb1, 0x47, 0xca, 0xf8, 0xa0, 0xc8, 0x18, 0x07,
	0x91, 0x9f, 0xdb, 0xa7, 0x75, 0x19, 0x21, 0x8f, 0x88, 0x91, 0x85, 0xc2, 0x90, 0xd1, 0x21, 0xf2,
	0x2c, 0x2e, 0x90, 0x88, 0xb8, 0xa2, 0x79, 0x7e, 0x55, 0x1a, 0x87, 0xf0, 0x8c, 0x89, 0x61, 0xc4,
	0xb3, 0xe8, 0x1f, 0x57, 0x65, 0x9a, 0xf0, 0xff, 0xb4, 0x2a, 0x6a, 0x88, 0x3c, 0xe2, 0x20, 0x41,
	0x68, 0x30, 0x49, 0xf0, 0xf1, 0x62, 0x02, 0x17, 0x53, 0x4b, 0x20, 0xe6, 0x62, 0x41, 0x02, 0xd7,
	0x62, 0x98, 0x0b, 0x46, 0xec, 0x98, 0xa5, 0x5a, 0xd4, 0xa1, 0x87, 0x6c, 0xfc, 0x3d, 0xf5, 0x1c,
	0xcc, 0x2c, 0x31, 0x0a, 0xb1, 0x42, 0xf5, 0x4a, 0x50, 0xf2, 0xa0, 0x0a, 0x32, 0xfe, 0x61, 0x25,
	0x2c, 0xc3, 0x43, 0x82, 0x5f, 0x56, 0xcd, 0x15, 0x63, 0x94, 0x15, 0x24, 0x4b, 0xae, 0x29, 0x82,
	0xbb, 0x29, 0x41, 0x48, 0x32, 0x49, 0xa8, 0xa5, 0xb7, 0xd5, 0x92, 0xfc, 0x1a, 0x44, 0x17, 0xc6,
	0x4b, 0x86, 0xc2, 0x10, 0xb3, 0xd4, 0xf7, 0xfd, 0x1c, 0x14, 0x05, 0x01, 0x15, 0x92, 0x5d, 0xad,
	0x6e, 0xfd, 0x7c, 0x03, 0xac, 0x3c, 0xc3, 0xd8, 0x39, 0x12, 0xd8, 0x87, 0xef, 0x80, 0x56, 0x4a,
	0x6e, 0x05, 0xc8, 0xc7, 0x5a, 0x6d, 0xb3, 0xb6, 0x7d, 0xd3, 0x6c, 0xa6, 0x93, 0x5f, 0x20, 0x1f,
	0xc3, 0x1d, 0x70, 0x2d, 0x0e, 0x57, 0xab, 0x6f, 0xd6, 0xb6, 0x1b, 0x9d, 0xfb, 0x4a, 0x7b, 0x7a,
	0xea, 0x5e, 0x3f, 0x11, 0x8c, 0x04, 0xee, 0x19, 0xf2, 0x22, 0x6c, 0x4a, 0x4b, 0xf8, 0x08, 0xd4,
	0x89, 0xa3, 0x2d, 0x49, 0xfb, 0x7b, 0x33, 0xf6, 0x47, 0x81, 0xf8, 0xe0, 0x71, 0x62, 0x5e, 0x27,
	0x0e, 0xdc, 0x07, 0xab, 0x5c, 0x20, 0x26, 0x2c, 0x07, 0x09, 0x6c, 0x09, 0xe2, 0x63, 0xed, 0x5a,
	0x05, 0x4f, 0x2d, 0x09, 0xda, 0x47, 0x02, 0x9f, 0x12, 0x1f, 0xc3, 0x5d, 0xd0, 0xc2, 0x81, 0x93,
	0xe3, 0xb8, 0x5e, 0x81, 0xa3, 0x81, 0x03, 0x27, 0x63, 0x70, 0xc0, 0x1a, 0x12, 0x82, 0x91, 0x41,
	0x24, 0x70, 0x7c, 0x2a, 0x11, 0xe6, 0xda, 0xf2, 0xe6, 0xd2, 0x76, 0xa3, 0xf3, 0x91, 0x5e, 0x5a,
	0x79, 0xf4, 0x34, 0xa5, 0xfd, 0x94, 0x22, 0xf1, 0xb0, 0x8a, 0x26, 0xbe, 0x39, 0xfc, 0xa5, 0x06,
	0xee, 0x16, 0x5e, 0x74, 0xed, 0xc6, 0x66, 0x6d, 0xbb, 0xdd, 0x31, 0x0b, 0xfd, 0xc9, 0x7b, 0xa7,
	0x1f, 0x62, 0x7a, 0x9a, 0xc2, 0xcd, 0x31, 0xfa, 0x20, 0x88, 0xfc, 0xa2, 0x35, 0x73, 0xdd, 0x9d,
	0xbf, 0x00, 0x6d, 0x70, 0x27, 0x62, 0x9e, 0x35, 0x5d, 0x19, 0xb9, 0xb6, 0x22, 0xf7, 0x6e, 0x14,
	0xc6, 0xa2, 0xea, 0xe1, 0xa7, 0x12, 0x78, 0x9c, 0xe2, 0xcc, 0xdb, 0x11, 0xf3, 0xa6, 0xe6, 0x38,
	0xfc, 0x1a, 0x2c, 0x27, 0xf2, 0xd0, 0x6e, 0xca, 0x1d, 0xee, 0x96, 0xec, 0x30, 0xcd, 0xe6, 0x89,
	0x04, 0xc9, 0x8d, 0x4d, 0x4e, 0x99, 0x8a, 0x0f, 0xda, 0xa0, 0xa9, 0x64, 0x48, 0x82, 0x0b, 0xca,
	0x35, 0x20, 0xa3, 0xde, 0xbd, 0xc2, 0x89, 0x1d, 0x8f, 0xeb, 0xc6, 0xb1, 0x64, 0x3a, 0x0a, 0x2e,
	0xa8, 0xd9, 0x08, 0xb3, 0x31, 0xef, 0x99, 0x6f, 0xfa, 0x5f, 0x82, 0x07, 0x63, 0x1e, 0x35, 0x0a,
	0x09, 0x8f, 0xb3, 0x60, 0x64, 0x72, 0x7a, 0x37, 0xc9, 0x22, 0x66, 0xdc, 0x78, 0x95, 0x0e, 0x5f,
	0x4b, 0xad, 0xc7, 0xcb, 0xdc, 0x78, 0x95, 0xc9, 0xfe, 0xf5, 0xd6, 0x9f, 0xd7, 0xc1, 0xff, 0xe7,
	0x5f, 0x1a, 0x78, 0x08, 0x6e, 0x49, 0xc3, 0xf1, 0x75, 0x24, 0x8e, 0x54, 0x66, 0x89, 0x9a, 0x56,
	0x63, 0x54, 0xc6, 0x75, 0xe4, 0xc4, 0xa2, 0x20, 0x81, 0xc0, 0x2e, 0x66, 0xc9, 0x85, 0x56, 0x12,
	0x5e, 0x48, 0xd2, 0x54, 0x88, 0x24, 0x94, 0xa7, 0xa0, 0x35, 0xa0, 0xd4, 0xc3, 0x28, 0x50, 0x0c,
	0x89, 0xa8, 0x37, 0x66, 0x18, 0xf6, 0x28, 0xf5, 0x14, 0x81, 0x02, 0xa4, 0x04, 0x4d, 0x2e, 0x15,
	0xa7, 0xf0, 0x55, 0xa4, 0xdd, 0xe0, 0xe3, 0x8f, 0x98, 0xc0, 0xa1, 0xd1, 0xc0, 0x53, 0x9a, 0x2c,
	0xd4, 0xf5, 0xbe, 0x34, 0x52, 0x04, 0xce, 0xf8, 0x03, 0x3e, 0x03, 0x8d, 0x90, 0x11, 0x3b, 0xc5,
	0x2f, 0x4b, 0xfc, 0xc3, 0xb2, 0x6b, 0xfd, 0x82, 0x06, 0x78, 0x64, 0x02, 0x89, 0x4c, 0x78, 0xf6,
	0x40, 0x7b, 0x22, 0x99, 0x5c, 0xbb, 0x21, 0xef, 0xda, 0xc2, 0x6c, 0xb6, 0xf2, 0xd9, 0xe4, 0xb0,
	0x0f, 0xda, 0x13, 0xe9, 0x4c, 0x55, 0xb6, 0x28, 0x9f, 0xad, 0x7c, 0x3e, 0x63, 0x8a, 0x56, 0x3e,
	0xa1, 0xb1, 0xa2, 0x96, 0x4a, 0x33, 0xda, 0xcc, 0x65, 0x54, 0x52, 0xe4, 0x53, 0x9a, 0x8a, 0x66,
	0x71, 0x4e, 0x9b, 0xb9, 0x9c, 0xf2, 0xad, 0x7f, 0x56, 0xc0, 0x5b, 0x0b, 0x05, 0x04, 0x87, 0xe0,
	0xce, 0x74, 0x47, 0xb6, 0x62, 0x71, 0x6b, 0x40, 0x56, 0x80, 0xbd, 0x92, 0x0a, 0x90, 0x23, 0x3d,
	0x1d, 0x85, 0x58, 0x96, 0x80, 0xa9, 0x39, 0xf3, 0x76, 0x38, 0x6b, 0x04, 0xbf, 0x01, 0x1b, 0x52,
	0x3c, 0x3e, 0x0a, 0x43, 0x55, 0x5e, 0x73, 0xfd, 0xad, 0x4a, 0x0f, 0x5b, 0x8f, 0xf1, 0x2f, 0x12,
	0xb8, 0x99, 0x6f, 0x84, 0x7e, 0xdc, 0x2d, 0x73, 0xbd, 0x5e, 0x8a, 0xa1, 0xdd, 0x79, 0x5e, 0xb6,
	0x15, 0x99, 0x14, 0x53, 0x22, 0x73, 0x05, 0x6d, 0x76, 0x3a, 0xee, 0xbb, 0xe3, 0x2f, 0x18, 0x81,
	0xd5, 0xa9, 0x67, 0x89, 0x54, 0x4f, 0xbb, 0xf3, 0x79, 0x25, 0x87, 0x7d, 0x85, 0x9d, 0x71, 0x39,
	0xb9, 0x60, 0xb6, 0xd1, 0xc4, 0x37, 0x1c, 0x80, 0xff, 0xa9, 0x8a, 0x2a, 0x68, 0x48, 0x6c, 0x0b,
	0x07, 0x82, 0x11, 0xcc, 0xb5, 0xeb, 0xf2, 0x92, 0xec, 0x94, 0x09, 0x27, 0xf1, 0x71, 0x1a, 0x43,
	0x0f, 0x02, 0xc1, 0x46, 0x26, 0x0c, 0x27, 0x67, 0x08, 0xe6, 0xf0, 0xa7, 0x1a, 0xb8, 0x35, 0xf3,
	0x4a, 0x94, 0xd2, 0x6c, 0x77, 0x4e, 0x2b, 0xf6, 0x86, 0xb3, 0x0c, 0x3f, 0xa7, 0x4b, 0x4c, 0x2f,
	0x9a, 0x6b, 0xc3, 0xa9, 0x19, 0xe8, 0x4e, 0x84, 0x90, 0xbc, 0xcb, 0x94, 0xa4, 0x7b, 0x57, 0x68,
	0x1f, 0x63, 0x4f, 0x07, 0x31, 0x45, 0xde, 0x91, 0x9c, 0xe0, 0xf0, 0xd7, 0x1a, 0x58, 0x2f, 0x78,
	0xd9, 0x6b, 0x2b, 0x72, 0xcb, 0xdf, 0x55, 0xdc, 0xf2, 0x57, 0x09, 0xcb, 0x9c, 0x93, 0x5d, 0x68,
	0x61, 0xde, 0xb9, 0x9c, 0x37, 0x0d, 0x7f, 0xaf, 0x81, 0x7b, 0xc5, 0xff, 0x0a, 0x49, 0x65, 0x69,
	0x77, 0x9c, 0xab, 0x05, 0xb7, 0x3f, 0x26, 0x32, 0x25, 0xcf, 0xbc, 0xf8, 0x66, 0x8c, 0xcc, 0xbb,
	0x97, 0x05, 0x2b, 0x7c, 0xeb, 0xef, 0x3a, 0x58, 0x2f, 0xc8, 0x36, 0xfc, 0x11, 0xac, 0x4d, 0x1f,
	0xa2, 0xec, 0x94, 0xed, 0xce, 0x49, 0x71, 0xd8, 0xf2, 0x58, 0x8a, 0x0e, 0xb0, 0xe0, 0x1a, 0x25,
	0x87, 0xbb, 0x3a, 0x75, 0xb8, 0xf0, 0x13, 0xd0, 0x70, 0x30, 0xb7, 0x19, 0x09, 0xe5, 0xfb, 0xad,
	0x4a, 0x79, 0xc9, 0x03, 0xe0, 0x11, 0x80, 0x33, 0xad, 0x3e, 0xae, 0x2b, 0xa5, 0x8d, 0x65, 0x6d,
	0xaa, 0xd7, 0x73, 0xf8, 0x04, 0x00, 0xfc, 0x83, 0x60, 0x48, 0x3e, 0x84, 0x2a, 0x3d, 0x7f, 0x6f,
	0x4a, 0xfb, 0xb8, 0x5a, 0xef, 0xfd, 0x5b, 0x03, 0x0f, 0x6d, 0xea, 0x97, 0xdf, 0xfb, 0xbd, 0x56,
	0x56, 0xf6, 0x63, 0xca, 0xe3, 0xda, 0xb7, 0x9f, 0x29, 0x8c, 0x4b, 0x3d, 0x14, 0xb8, 0x3a, 0x65,
	0xae, 0xe1, 0xe2, 0x40, 0x3a, 0x34, 0xc6, 0x0f, 0xa5, 0x05, 0x3f, 0xfa, 0x4f, 0xb2, 0xd1, 0x6f,
	0xf5, 0xa5, 0xc3, 0x7e, 0xff, 0x8f, 0xfa, 0x83, 0xc3, 0x84, 0xb2, 0xef, 0x70, 0x3d, 0x19, 0xc6,
	0xa3, 0xb3, 0xae, 0x9e, 0x96, 0x63, 0xfe, 0x57, 0x6a, 0x73, 0xde, 0x77, 0xf8, 0x79, 0x66, 0x73,
	0x7e, 0xd6, 0x3d, 0xcf, 0x6c, 0xde, 0xd4, 0x1f, 0x26, 0x0b, 0xbd, 0x5e, 0xdf, 0xe1, 0xbd, 0x5e,
	0x66, 0xd5, 0xeb, 0x9d, 0x75, 0x7b, 0xbd, 0xcc, 0x6e, 0xb0, 0x2c, 0x83, 0xed, 0xfe, 0x17, 0x00,
	0x00, 0xff, 0xff, 0x2c, 0x38, 0x2b, 0x40, 0x94, 0x10, 0x00, 0x00,
}
