// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/label.proto

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

// A label.
type Label struct {
	// Immutable. Name of the resource.
	// Label resource names have the form:
	// `customers/{customer_id}/labels/{label_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. Id of the label. Read only.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the label.
	//
	// This field is required and should not be empty when creating a new label.
	//
	// The length of this string should be between 1 and 80, inclusive.
	Name *wrappers.StringValue `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. Status of the label. Read only.
	Status enums.LabelStatusEnum_LabelStatus `protobuf:"varint,4,opt,name=status,proto3,enum=google.ads.googleads.v2.enums.LabelStatusEnum_LabelStatus" json:"status,omitempty"`
	// A type of label displaying text on a colored background.
	TextLabel            *common.TextLabel `protobuf:"bytes,5,opt,name=text_label,json=textLabel,proto3" json:"text_label,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Label) Reset()         { *m = Label{} }
func (m *Label) String() string { return proto.CompactTextString(m) }
func (*Label) ProtoMessage()    {}
func (*Label) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff08e2c812fe9893, []int{0}
}

func (m *Label) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Label.Unmarshal(m, b)
}
func (m *Label) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Label.Marshal(b, m, deterministic)
}
func (m *Label) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Label.Merge(m, src)
}
func (m *Label) XXX_Size() int {
	return xxx_messageInfo_Label.Size(m)
}
func (m *Label) XXX_DiscardUnknown() {
	xxx_messageInfo_Label.DiscardUnknown(m)
}

var xxx_messageInfo_Label proto.InternalMessageInfo

func (m *Label) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *Label) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Label) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Label) GetStatus() enums.LabelStatusEnum_LabelStatus {
	if m != nil {
		return m.Status
	}
	return enums.LabelStatusEnum_UNSPECIFIED
}

func (m *Label) GetTextLabel() *common.TextLabel {
	if m != nil {
		return m.TextLabel
	}
	return nil
}

func init() {
	proto.RegisterType((*Label)(nil), "google.ads.googleads.v2.resources.Label")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/label.proto", fileDescriptor_ff08e2c812fe9893)
}

var fileDescriptor_ff08e2c812fe9893 = []byte{
	// 489 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xcd, 0x6a, 0x1b, 0x31,
	0x18, 0x64, 0xd7, 0x71, 0x20, 0xea, 0xcf, 0x61, 0x4f, 0x6e, 0x1a, 0x52, 0xa7, 0x25, 0xc5, 0x3d,
	0x54, 0x32, 0xdb, 0xd2, 0x83, 0x7a, 0x92, 0x21, 0x24, 0xfd, 0xa1, 0x04, 0xa7, 0xb8, 0x50, 0x0c,
	0x46, 0xde, 0x55, 0xb6, 0x0b, 0x2b, 0x69, 0x91, 0xb4, 0x6e, 0x20, 0xe4, 0x65, 0x0a, 0xbd, 0xf4,
	0x51, 0xfa, 0x14, 0x39, 0xe7, 0x05, 0x0a, 0x3d, 0x15, 0xeb, 0x67, 0x6d, 0x28, 0x4e, 0x4e, 0x3b,
	0xe2, 0x9b, 0x19, 0x0d, 0xa3, 0x6f, 0xc1, 0xcb, 0x42, 0xca, 0xa2, 0x62, 0x88, 0xe6, 0x1a, 0x39,
	0xb8, 0x44, 0x8b, 0x14, 0x29, 0xa6, 0x65, 0xa3, 0x32, 0xa6, 0x51, 0x45, 0xe7, 0xac, 0x82, 0xb5,
	0x92, 0x46, 0x26, 0x07, 0x8e, 0x03, 0x69, 0xae, 0x61, 0x4b, 0x87, 0x8b, 0x14, 0xb6, 0xf4, 0x5d,
	0xb4, 0xc9, 0x31, 0x93, 0x9c, 0x4b, 0x81, 0x0c, 0xbb, 0x30, 0xb3, 0x35, 0xcf, 0xdd, 0xe1, 0x26,
	0x01, 0x13, 0x0d, 0xf7, 0xd7, 0xcf, 0xb4, 0xa1, 0xa6, 0xd1, 0x5e, 0xf1, 0x24, 0x28, 0xea, 0x12,
	0x9d, 0x97, 0xac, 0xca, 0x67, 0x73, 0xf6, 0x8d, 0x2e, 0x4a, 0xa9, 0x3c, 0xe1, 0xd1, 0x1a, 0x21,
	0x24, 0xf3, 0xa3, 0x7d, 0x3f, 0xb2, 0xa7, 0x79, 0x73, 0x8e, 0xbe, 0x2b, 0x5a, 0xd7, 0x4c, 0x05,
	0xef, 0xbd, 0x35, 0x29, 0x15, 0x42, 0x1a, 0x6a, 0x4a, 0x29, 0xfc, 0xf4, 0xe9, 0xcf, 0x0e, 0xe8,
	0x7e, 0x5c, 0x06, 0x4a, 0x3e, 0x80, 0x07, 0xc1, 0x79, 0x26, 0x28, 0x67, 0xbd, 0xa8, 0x1f, 0x0d,
	0x76, 0x46, 0xcf, 0xaf, 0x49, 0xf7, 0x2f, 0xe9, 0x83, 0xfd, 0x55, 0x3b, 0x1e, 0xd5, 0xa5, 0x86,
	0x99, 0xe4, 0xc8, 0xca, 0xc7, 0xf7, 0x83, 0xf8, 0x13, 0xe5, 0x2c, 0x19, 0x82, 0xb8, 0xcc, 0x7b,
	0x71, 0x3f, 0x1a, 0xdc, 0x4b, 0x1f, 0x7b, 0x01, 0x0c, 0x09, 0xe1, 0x3b, 0x61, 0xde, 0xbc, 0x9e,
	0xd0, 0xaa, 0x61, 0xa3, 0xce, 0x35, 0xe9, 0x8c, 0xe3, 0x32, 0x4f, 0x86, 0x60, 0xcb, 0xde, 0xda,
	0xb1, 0x9a, 0xbd, 0xff, 0x34, 0x67, 0x46, 0x95, 0xa2, 0xb0, 0xa2, 0xb1, 0x65, 0x26, 0x5f, 0xc0,
	0xb6, 0x2b, 0xb1, 0xb7, 0xd5, 0x8f, 0x06, 0x0f, 0x53, 0x0c, 0x37, 0xbd, 0xa5, 0xed, 0x1d, 0xda,
	0x9c, 0x67, 0x56, 0x71, 0x24, 0x1a, 0xbe, 0x7e, 0x76, 0x31, 0xbc, 0x5d, 0x72, 0x02, 0xc0, 0xea,
	0x4d, 0x7b, 0x5d, 0x1b, 0xe8, 0xc5, 0x46, 0x73, 0xb7, 0x05, 0xf0, 0x33, 0xbb, 0x30, 0xae, 0x89,
	0x1d, 0x13, 0x20, 0x3e, 0xb9, 0x21, 0x47, 0x77, 0x35, 0x97, 0x3c, 0xcb, 0x1a, 0x6d, 0x24, 0x67,
	0x4a, 0xa3, 0xcb, 0x00, 0xaf, 0xdc, 0x96, 0x68, 0x74, 0x69, 0xbf, 0x57, 0xa3, 0x3f, 0x11, 0x38,
	0xcc, 0x24, 0x87, 0x77, 0xae, 0xeb, 0x08, 0x58, 0xd7, 0xd3, 0x65, 0x6f, 0xa7, 0xd1, 0xd7, 0xf7,
	0x5e, 0x50, 0xc8, 0x8a, 0x8a, 0x02, 0x4a, 0x55, 0xa0, 0x82, 0x09, 0xdb, 0x2a, 0x5a, 0xa5, 0xb9,
	0xe5, 0x6f, 0x79, 0xdb, 0xa2, 0x1f, 0x71, 0xe7, 0x98, 0x90, 0x5f, 0xf1, 0xc1, 0xb1, 0xb3, 0x24,
	0xb9, 0x86, 0x0e, 0x2e, 0xd1, 0x24, 0x85, 0xe3, 0xc0, 0xfc, 0x1d, 0x38, 0x53, 0x92, 0xeb, 0x69,
	0xcb, 0x99, 0x4e, 0xd2, 0x69, 0xcb, 0xb9, 0x89, 0x0f, 0xdd, 0x00, 0x63, 0x92, 0x6b, 0x8c, 0x5b,
	0x16, 0xc6, 0x93, 0x14, 0xe3, 0x96, 0x37, 0xdf, 0xb6, 0x61, 0x5f, 0xfd, 0x0b, 0x00, 0x00, 0xff,
	0xff, 0x9f, 0xf1, 0x42, 0xad, 0xd9, 0x03, 0x00, 0x00,
}
