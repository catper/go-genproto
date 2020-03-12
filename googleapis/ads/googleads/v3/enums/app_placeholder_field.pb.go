// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/app_placeholder_field.proto

package enums

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

// Possible values for App placeholder fields.
type AppPlaceholderFieldEnum_AppPlaceholderField int32

const (
	// Not specified.
	AppPlaceholderFieldEnum_UNSPECIFIED AppPlaceholderFieldEnum_AppPlaceholderField = 0
	// Used for return value only. Represents value unknown in this version.
	AppPlaceholderFieldEnum_UNKNOWN AppPlaceholderFieldEnum_AppPlaceholderField = 1
	// Data Type: INT64. The application store that the target application
	// belongs to. Valid values are: 1 = Apple iTunes Store; 2 = Google Play
	// Store.
	AppPlaceholderFieldEnum_STORE AppPlaceholderFieldEnum_AppPlaceholderField = 2
	// Data Type: STRING. The store-specific ID for the target application.
	AppPlaceholderFieldEnum_ID AppPlaceholderFieldEnum_AppPlaceholderField = 3
	// Data Type: STRING. The visible text displayed when the link is rendered
	// in an ad.
	AppPlaceholderFieldEnum_LINK_TEXT AppPlaceholderFieldEnum_AppPlaceholderField = 4
	// Data Type: STRING. The destination URL of the in-app link.
	AppPlaceholderFieldEnum_URL AppPlaceholderFieldEnum_AppPlaceholderField = 5
	// Data Type: URL_LIST. Final URLs for the in-app link when using Upgraded
	// URLs.
	AppPlaceholderFieldEnum_FINAL_URLS AppPlaceholderFieldEnum_AppPlaceholderField = 6
	// Data Type: URL_LIST. Final Mobile URLs for the in-app link when using
	// Upgraded URLs.
	AppPlaceholderFieldEnum_FINAL_MOBILE_URLS AppPlaceholderFieldEnum_AppPlaceholderField = 7
	// Data Type: URL. Tracking template for the in-app link when using Upgraded
	// URLs.
	AppPlaceholderFieldEnum_TRACKING_URL AppPlaceholderFieldEnum_AppPlaceholderField = 8
	// Data Type: STRING. Final URL suffix for the in-app link when using
	// parallel tracking.
	AppPlaceholderFieldEnum_FINAL_URL_SUFFIX AppPlaceholderFieldEnum_AppPlaceholderField = 9
)

var AppPlaceholderFieldEnum_AppPlaceholderField_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "STORE",
	3: "ID",
	4: "LINK_TEXT",
	5: "URL",
	6: "FINAL_URLS",
	7: "FINAL_MOBILE_URLS",
	8: "TRACKING_URL",
	9: "FINAL_URL_SUFFIX",
}

var AppPlaceholderFieldEnum_AppPlaceholderField_value = map[string]int32{
	"UNSPECIFIED":       0,
	"UNKNOWN":           1,
	"STORE":             2,
	"ID":                3,
	"LINK_TEXT":         4,
	"URL":               5,
	"FINAL_URLS":        6,
	"FINAL_MOBILE_URLS": 7,
	"TRACKING_URL":      8,
	"FINAL_URL_SUFFIX":  9,
}

func (x AppPlaceholderFieldEnum_AppPlaceholderField) String() string {
	return proto.EnumName(AppPlaceholderFieldEnum_AppPlaceholderField_name, int32(x))
}

func (AppPlaceholderFieldEnum_AppPlaceholderField) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_60a611ef024f5427, []int{0, 0}
}

// Values for App placeholder fields.
type AppPlaceholderFieldEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppPlaceholderFieldEnum) Reset()         { *m = AppPlaceholderFieldEnum{} }
func (m *AppPlaceholderFieldEnum) String() string { return proto.CompactTextString(m) }
func (*AppPlaceholderFieldEnum) ProtoMessage()    {}
func (*AppPlaceholderFieldEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_60a611ef024f5427, []int{0}
}

func (m *AppPlaceholderFieldEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppPlaceholderFieldEnum.Unmarshal(m, b)
}
func (m *AppPlaceholderFieldEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppPlaceholderFieldEnum.Marshal(b, m, deterministic)
}
func (m *AppPlaceholderFieldEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppPlaceholderFieldEnum.Merge(m, src)
}
func (m *AppPlaceholderFieldEnum) XXX_Size() int {
	return xxx_messageInfo_AppPlaceholderFieldEnum.Size(m)
}
func (m *AppPlaceholderFieldEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AppPlaceholderFieldEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AppPlaceholderFieldEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.AppPlaceholderFieldEnum_AppPlaceholderField", AppPlaceholderFieldEnum_AppPlaceholderField_name, AppPlaceholderFieldEnum_AppPlaceholderField_value)
	proto.RegisterType((*AppPlaceholderFieldEnum)(nil), "google.ads.googleads.v3.enums.AppPlaceholderFieldEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/app_placeholder_field.proto", fileDescriptor_60a611ef024f5427)
}

var fileDescriptor_60a611ef024f5427 = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xdd, 0xce, 0x93, 0x30,
	0x18, 0x16, 0xe6, 0xb7, 0xb9, 0x7e, 0xfe, 0xd4, 0xaa, 0xd1, 0x18, 0x77, 0xb0, 0x5d, 0x40, 0x39,
	0xe0, 0xc8, 0x7a, 0x54, 0x36, 0x58, 0x9a, 0x21, 0x23, 0xfc, 0xcc, 0xc5, 0x90, 0x10, 0x1c, 0x88,
	0x24, 0x8c, 0x36, 0xeb, 0xb6, 0x0b, 0xf2, 0xd0, 0xc4, 0x0b, 0xd1, 0x4b, 0x31, 0xf1, 0x1e, 0x0c,
	0xd4, 0xe1, 0xc9, 0xfc, 0x4e, 0x9a, 0xa7, 0xef, 0xf3, 0x93, 0xf6, 0x79, 0xc1, 0xdb, 0x92, 0xf3,
	0xb2, 0x2e, 0x8c, 0x2c, 0x97, 0x86, 0x82, 0x2d, 0x3a, 0x9b, 0x46, 0xd1, 0x9c, 0xf6, 0xd2, 0xc8,
	0x84, 0x48, 0x45, 0x9d, 0xed, 0x8a, 0x2f, 0xbc, 0xce, 0x8b, 0x43, 0xfa, 0xb9, 0x2a, 0xea, 0x1c,
	0x8b, 0x03, 0x3f, 0x72, 0x34, 0x51, 0x7a, 0x9c, 0xe5, 0x12, 0xf7, 0x56, 0x7c, 0x36, 0x71, 0x67,
	0x7d, 0xfd, 0xe6, 0x92, 0x2c, 0x2a, 0x23, 0x6b, 0x1a, 0x7e, 0xcc, 0x8e, 0x15, 0x6f, 0xa4, 0x32,
	0xcf, 0x7e, 0x68, 0xe0, 0x25, 0x15, 0xc2, 0xff, 0x97, 0xed, 0xb4, 0xd1, 0x76, 0x73, 0xda, 0xcf,
	0xbe, 0x6b, 0xe0, 0xd9, 0x15, 0x0e, 0x3d, 0x01, 0xb7, 0xb1, 0x17, 0xfa, 0xf6, 0x9c, 0x39, 0xcc,
	0x5e, 0xc0, 0x7b, 0xe8, 0x16, 0x8c, 0x62, 0x6f, 0xe5, 0xad, 0x3f, 0x78, 0x50, 0x43, 0x63, 0x70,
	0x13, 0x46, 0xeb, 0xc0, 0x86, 0x3a, 0x1a, 0x02, 0x9d, 0x2d, 0xe0, 0x00, 0x3d, 0x02, 0x63, 0x97,
	0x79, 0xab, 0x34, 0xb2, 0xb7, 0x11, 0xbc, 0x8f, 0x46, 0x60, 0x10, 0x07, 0x2e, 0xbc, 0x41, 0x8f,
	0x01, 0x70, 0x98, 0x47, 0xdd, 0x34, 0x0e, 0xdc, 0x10, 0x0e, 0xd1, 0x0b, 0xf0, 0x54, 0xdd, 0xdf,
	0xaf, 0x2d, 0xe6, 0xda, 0x6a, 0x3c, 0x42, 0x10, 0x3c, 0x8c, 0x02, 0x3a, 0x5f, 0x31, 0x6f, 0xd9,
	0x8e, 0xe0, 0x03, 0xf4, 0x1c, 0xc0, 0xde, 0x98, 0x86, 0xb1, 0xe3, 0xb0, 0x2d, 0x1c, 0x5b, 0xbf,
	0x35, 0x30, 0xdd, 0xf1, 0x3d, 0xbe, 0xb3, 0x0f, 0xeb, 0xd5, 0x95, 0x2f, 0xf9, 0x6d, 0x17, 0xbe,
	0xf6, 0xd1, 0xfa, 0x6b, 0x2d, 0x79, 0x9d, 0x35, 0x25, 0xe6, 0x87, 0xd2, 0x28, 0x8b, 0xa6, 0x6b,
	0xea, 0xb2, 0x15, 0x51, 0xc9, 0xff, 0x2c, 0xe9, 0x5d, 0x77, 0x7e, 0xd5, 0x07, 0x4b, 0x4a, 0xbf,
	0xe9, 0x93, 0xa5, 0x8a, 0xa2, 0xb9, 0xc4, 0x0a, 0xb6, 0x68, 0x63, 0xe2, 0xb6, 0x5a, 0xf9, 0xf3,
	0xc2, 0x27, 0x34, 0x97, 0x49, 0xcf, 0x27, 0x1b, 0x33, 0xe9, 0xf8, 0x5f, 0xfa, 0x54, 0x0d, 0x09,
	0xa1, 0xb9, 0x24, 0xa4, 0x57, 0x10, 0xb2, 0x31, 0x09, 0xe9, 0x34, 0x9f, 0x86, 0xdd, 0xc3, 0xcc,
	0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd1, 0xe8, 0xae, 0xdb, 0x3c, 0x02, 0x00, 0x00,
}
