// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/google_ads_field.proto

package resources

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

// A field or resource (artifact) used by GoogleAdsService.
type GoogleAdsField struct {
	// Immutable. The resource name of the artifact.
	// Artifact resource names have the form:
	//
	// `googleAdsFields/{name}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The name of the artifact.
	Name *wrappers.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. The category of the artifact.
	Category enums.GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory `protobuf:"varint,3,opt,name=category,proto3,enum=google.ads.googleads.v2.enums.GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory" json:"category,omitempty"`
	// Output only. Whether the artifact can be used in a SELECT clause in search
	// queries.
	Selectable *wrappers.BoolValue `protobuf:"bytes,4,opt,name=selectable,proto3" json:"selectable,omitempty"`
	// Output only. Whether the artifact can be used in a WHERE clause in search
	// queries.
	Filterable *wrappers.BoolValue `protobuf:"bytes,5,opt,name=filterable,proto3" json:"filterable,omitempty"`
	// Output only. Whether the artifact can be used in a ORDER BY clause in search
	// queries.
	Sortable *wrappers.BoolValue `protobuf:"bytes,6,opt,name=sortable,proto3" json:"sortable,omitempty"`
	// Output only. The names of all resources, segments, and metrics that are selectable with
	// the described artifact.
	SelectableWith []*wrappers.StringValue `protobuf:"bytes,7,rep,name=selectable_with,json=selectableWith,proto3" json:"selectable_with,omitempty"`
	// Output only. The names of all resources that are selectable with the described
	// artifact. Fields from these resources do not segment metrics when included
	// in search queries.
	//
	// This field is only set for artifacts whose category is RESOURCE.
	AttributeResources []*wrappers.StringValue `protobuf:"bytes,8,rep,name=attribute_resources,json=attributeResources,proto3" json:"attribute_resources,omitempty"`
	// Output only. At and beyond version V1 this field lists the names of all metrics that are
	// selectable with the described artifact when it is used in the FROM clause.
	// It is only set for artifacts whose category is RESOURCE.
	//
	// Before version V1 this field lists the names of all metrics that are
	// selectable with the described artifact. It is only set for artifacts whose
	// category is either RESOURCE or SEGMENT
	Metrics []*wrappers.StringValue `protobuf:"bytes,9,rep,name=metrics,proto3" json:"metrics,omitempty"`
	// Output only. At and beyond version V1 this field lists the names of all artifacts,
	// whether a segment or another resource, that segment metrics when included
	// in search queries and when the described artifact is used in the FROM
	// clause. It is only set for artifacts whose category is RESOURCE.
	//
	// Before version V1 this field lists the names of all artifacts, whether a
	// segment or another resource, that segment metrics when included in search
	// queries. It is only set for artifacts of category RESOURCE, SEGMENT or
	// METRIC.
	Segments []*wrappers.StringValue `protobuf:"bytes,10,rep,name=segments,proto3" json:"segments,omitempty"`
	// Output only. Values the artifact can assume if it is a field of type ENUM.
	//
	// This field is only set for artifacts of category SEGMENT or ATTRIBUTE.
	EnumValues []*wrappers.StringValue `protobuf:"bytes,11,rep,name=enum_values,json=enumValues,proto3" json:"enum_values,omitempty"`
	// Output only. This field determines the operators that can be used with the artifact
	// in WHERE clauses.
	DataType enums.GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType `protobuf:"varint,12,opt,name=data_type,json=dataType,proto3,enum=google.ads.googleads.v2.enums.GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType" json:"data_type,omitempty"`
	// Output only. The URL of proto describing the artifact's data type.
	TypeUrl *wrappers.StringValue `protobuf:"bytes,13,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	// Output only. Whether the field artifact is repeated.
	IsRepeated           *wrappers.BoolValue `protobuf:"bytes,14,opt,name=is_repeated,json=isRepeated,proto3" json:"is_repeated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *GoogleAdsField) Reset()         { *m = GoogleAdsField{} }
func (m *GoogleAdsField) String() string { return proto.CompactTextString(m) }
func (*GoogleAdsField) ProtoMessage()    {}
func (*GoogleAdsField) Descriptor() ([]byte, []int) {
	return fileDescriptor_e51b907878502641, []int{0}
}

func (m *GoogleAdsField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoogleAdsField.Unmarshal(m, b)
}
func (m *GoogleAdsField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoogleAdsField.Marshal(b, m, deterministic)
}
func (m *GoogleAdsField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoogleAdsField.Merge(m, src)
}
func (m *GoogleAdsField) XXX_Size() int {
	return xxx_messageInfo_GoogleAdsField.Size(m)
}
func (m *GoogleAdsField) XXX_DiscardUnknown() {
	xxx_messageInfo_GoogleAdsField.DiscardUnknown(m)
}

var xxx_messageInfo_GoogleAdsField proto.InternalMessageInfo

func (m *GoogleAdsField) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *GoogleAdsField) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *GoogleAdsField) GetCategory() enums.GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory {
	if m != nil {
		return m.Category
	}
	return enums.GoogleAdsFieldCategoryEnum_UNSPECIFIED
}

func (m *GoogleAdsField) GetSelectable() *wrappers.BoolValue {
	if m != nil {
		return m.Selectable
	}
	return nil
}

func (m *GoogleAdsField) GetFilterable() *wrappers.BoolValue {
	if m != nil {
		return m.Filterable
	}
	return nil
}

func (m *GoogleAdsField) GetSortable() *wrappers.BoolValue {
	if m != nil {
		return m.Sortable
	}
	return nil
}

func (m *GoogleAdsField) GetSelectableWith() []*wrappers.StringValue {
	if m != nil {
		return m.SelectableWith
	}
	return nil
}

func (m *GoogleAdsField) GetAttributeResources() []*wrappers.StringValue {
	if m != nil {
		return m.AttributeResources
	}
	return nil
}

func (m *GoogleAdsField) GetMetrics() []*wrappers.StringValue {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *GoogleAdsField) GetSegments() []*wrappers.StringValue {
	if m != nil {
		return m.Segments
	}
	return nil
}

func (m *GoogleAdsField) GetEnumValues() []*wrappers.StringValue {
	if m != nil {
		return m.EnumValues
	}
	return nil
}

func (m *GoogleAdsField) GetDataType() enums.GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType {
	if m != nil {
		return m.DataType
	}
	return enums.GoogleAdsFieldDataTypeEnum_UNSPECIFIED
}

func (m *GoogleAdsField) GetTypeUrl() *wrappers.StringValue {
	if m != nil {
		return m.TypeUrl
	}
	return nil
}

func (m *GoogleAdsField) GetIsRepeated() *wrappers.BoolValue {
	if m != nil {
		return m.IsRepeated
	}
	return nil
}

func init() {
	proto.RegisterType((*GoogleAdsField)(nil), "google.ads.googleads.v2.resources.GoogleAdsField")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/google_ads_field.proto", fileDescriptor_e51b907878502641)
}

var fileDescriptor_e51b907878502641 = []byte{
	// 652 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x95, 0x5d, 0x6f, 0xd3, 0x3e,
	0x14, 0xc6, 0xd5, 0x76, 0x2f, 0x9d, 0xbb, 0xf5, 0x2f, 0x79, 0x37, 0xf9, 0x4f, 0x13, 0x74, 0x93,
	0x26, 0xc6, 0x8d, 0x23, 0x15, 0x21, 0xa1, 0x8c, 0x01, 0x29, 0x2f, 0x93, 0x26, 0x81, 0xaa, 0x32,
	0x8a, 0x84, 0x2a, 0x45, 0x6e, 0x73, 0x96, 0x59, 0x4a, 0xe2, 0x60, 0x3b, 0x9d, 0x26, 0x84, 0xc4,
	0x67, 0xe1, 0x92, 0x8f, 0xc2, 0xa7, 0xd8, 0xf5, 0xee, 0xb9, 0xe1, 0x0a, 0x25, 0x71, 0xdc, 0x96,
	0x31, 0x08, 0xdc, 0xd9, 0x39, 0xcf, 0xef, 0x39, 0x4f, 0x9d, 0xe3, 0x14, 0x3d, 0x08, 0x38, 0x0f,
	0x42, 0xb0, 0xa9, 0x2f, 0xed, 0x62, 0x99, 0xad, 0xa6, 0x5d, 0x5b, 0x80, 0xe4, 0xa9, 0x98, 0x40,
	0xf9, 0xd8, 0xa3, 0xbe, 0xf4, 0x4e, 0x19, 0x84, 0x3e, 0x49, 0x04, 0x57, 0x1c, 0xef, 0x14, 0xcf,
	0x09, 0xf5, 0x25, 0x31, 0x24, 0x99, 0x76, 0x89, 0x21, 0xb7, 0x0e, 0x6f, 0x32, 0x87, 0x38, 0x8d,
	0xae, 0x1b, 0x7b, 0x13, 0xaa, 0x20, 0xe0, 0xe2, 0xa2, 0xe8, 0xb0, 0xf5, 0xe8, 0x2f, 0x71, 0x9f,
	0x2a, 0xea, 0xa9, 0x8b, 0x04, 0x34, 0x7f, 0xbb, 0xe4, 0x13, 0x66, 0x17, 0x8a, 0x31, 0x9c, 0xd1,
	0x29, 0xe3, 0x42, 0x0b, 0xfe, 0x9f, 0x13, 0x94, 0xa9, 0x75, 0xe9, 0x96, 0x2e, 0xe5, 0xbb, 0x71,
	0x7a, 0x6a, 0x9f, 0x0b, 0x9a, 0x24, 0x20, 0xa4, 0xae, 0x6f, 0xcf, 0xa1, 0x34, 0x8e, 0xb9, 0xa2,
	0x8a, 0xf1, 0x58, 0x57, 0x77, 0xbf, 0x35, 0x51, 0xfb, 0x28, 0x17, 0xb8, 0xbe, 0x7c, 0x91, 0xb5,
	0xc6, 0x27, 0x68, 0xa3, 0x6c, 0xe1, 0xc5, 0x34, 0x02, 0xab, 0xd6, 0xa9, 0xed, 0xaf, 0xf5, 0xec,
	0x4b, 0x77, 0xf9, 0xbb, 0x7b, 0x17, 0xdd, 0x99, 0x1d, 0xa1, 0x5e, 0x25, 0x4c, 0x92, 0x09, 0x8f,
	0xec, 0x45, 0x9f, 0xc1, 0x7a, 0xe9, 0xf2, 0x8a, 0x46, 0x80, 0xef, 0xa3, 0xa5, 0xdc, 0xac, 0xde,
	0xa9, 0xed, 0xb7, 0xba, 0xdb, 0x9a, 0x25, 0x65, 0x6a, 0xf2, 0x5a, 0x09, 0x16, 0x07, 0x43, 0x1a,
	0xa6, 0xd0, 0x6b, 0x5c, 0xba, 0x8d, 0x41, 0x2e, 0xc7, 0xef, 0x51, 0xb3, 0x3c, 0x6b, 0xab, 0xd1,
	0xa9, 0xed, 0xb7, 0xbb, 0x7d, 0x72, 0xd3, 0xeb, 0xcc, 0x0f, 0x9b, 0x2c, 0xa6, 0x78, 0xaa, 0xe1,
	0xe7, 0x71, 0x1a, 0xdd, 0x50, 0x2a, 0xda, 0x99, 0x36, 0xf8, 0x31, 0x42, 0x12, 0x42, 0x98, 0x28,
	0x3a, 0x0e, 0xc1, 0x5a, 0xca, 0xf3, 0x6e, 0x5d, 0xcb, 0xdb, 0xe3, 0x3c, 0x9c, 0x4b, 0x3b, 0x87,
	0x64, 0x06, 0xa7, 0x2c, 0x54, 0x20, 0x72, 0x83, 0xe5, 0x8a, 0x06, 0x33, 0x04, 0x1f, 0xa0, 0xa6,
	0xe4, 0xa2, 0xe8, 0xbf, 0x52, 0x0d, 0x37, 0x00, 0x3e, 0x46, 0xff, 0xcd, 0xb2, 0x78, 0xe7, 0x4c,
	0x9d, 0x59, 0xab, 0x9d, 0x46, 0xb5, 0x33, 0x6f, 0xcf, 0xc8, 0xb7, 0x4c, 0x9d, 0xe1, 0x01, 0xda,
	0xa4, 0x4a, 0x09, 0x36, 0x4e, 0x15, 0x78, 0xe6, 0xb6, 0x58, 0xcd, 0xaa, 0x7e, 0xd8, 0xd0, 0x83,
	0x12, 0xc6, 0x07, 0x68, 0x35, 0x02, 0x25, 0xd8, 0x44, 0x5a, 0x6b, 0x55, 0x7d, 0x4a, 0x02, 0x1f,
	0xa2, 0xa6, 0x84, 0x20, 0x82, 0x58, 0x49, 0x0b, 0x55, 0xa5, 0x0d, 0x82, 0x7b, 0xa8, 0x95, 0x0d,
	0x89, 0x37, 0xcd, 0x8a, 0xd2, 0x6a, 0x55, 0x75, 0x40, 0x19, 0x95, 0xef, 0x25, 0x16, 0x68, 0xcd,
	0x5c, 0x5f, 0x6b, 0xfd, 0x1f, 0x46, 0xf2, 0x19, 0x55, 0xf4, 0xe4, 0x22, 0x81, 0x5f, 0x8c, 0x64,
	0x59, 0xd2, 0xb9, 0x7d, 0xbd, 0xc5, 0x0f, 0x51, 0x33, 0x6b, 0xe7, 0xa5, 0x22, 0xb4, 0x36, 0xaa,
	0x5e, 0xa0, 0xd5, 0x0c, 0x79, 0x23, 0x42, 0xfc, 0x04, 0xb5, 0x98, 0xf4, 0x04, 0x24, 0x40, 0x15,
	0xf8, 0x56, 0xbb, 0xe2, 0x40, 0x32, 0x39, 0xd0, 0x88, 0xd3, 0xbf, 0x72, 0x5f, 0x56, 0xbe, 0xf8,
	0x78, 0x37, 0x58, 0xd8, 0x4b, 0xfb, 0xc3, 0xcf, 0x1f, 0xc0, 0x8f, 0xbd, 0x4f, 0x75, 0xb4, 0x37,
	0xe1, 0x11, 0xf9, 0xe3, 0xa7, 0xb9, 0xb7, 0xb9, 0xe8, 0xde, 0xcf, 0xe2, 0xf6, 0x6b, 0xef, 0x8e,
	0x35, 0x19, 0xf0, 0x90, 0xc6, 0x01, 0xe1, 0x22, 0xb0, 0x03, 0x88, 0xf3, 0x1f, 0x63, 0xcf, 0xe2,
	0xfd, 0xe6, 0xdf, 0xe2, 0xc0, 0xac, 0x3e, 0xd7, 0x1b, 0x47, 0xae, 0xfb, 0xa5, 0xbe, 0x53, 0x74,
	0x22, 0xae, 0x3f, 0xf7, 0xca, 0xc8, 0xb0, 0x4b, 0xcc, 0xf0, 0x7e, 0x2d, 0x35, 0x23, 0xd7, 0x97,
	0x23, 0xa3, 0x19, 0x0d, 0xbb, 0x23, 0xa3, 0xb9, 0xaa, 0xef, 0x15, 0x05, 0xc7, 0x71, 0x7d, 0xe9,
	0x38, 0x46, 0xe5, 0x38, 0xc3, 0xae, 0xe3, 0x18, 0xdd, 0x78, 0x25, 0x0f, 0x7b, 0xef, 0x47, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x3f, 0x9b, 0xac, 0xb9, 0xd9, 0x06, 0x00, 0x00,
}
