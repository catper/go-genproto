// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/servicedirectory/v1beta1/namespace.proto

package servicedirectory

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

// A container for [services][google.cloud.servicedirectory.v1beta1.Service].
// Namespaces allow administrators to group services together and define
// permissions for a collection of services.
type Namespace struct {
	// Immutable. The resource name for the namespace in the format
	// 'projects/*/locations/*/namespaces/*'.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optional. Resource labels associated with this Namespace.
	// No more than 64 user labels can be associated with a given resource.  Label
	// keys and values can be no longer than 63 characters.
	Labels               map[string]string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Namespace) Reset()         { *m = Namespace{} }
func (m *Namespace) String() string { return proto.CompactTextString(m) }
func (*Namespace) ProtoMessage()    {}
func (*Namespace) Descriptor() ([]byte, []int) {
	return fileDescriptor_aedd900d4e39df62, []int{0}
}

func (m *Namespace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Namespace.Unmarshal(m, b)
}
func (m *Namespace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Namespace.Marshal(b, m, deterministic)
}
func (m *Namespace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Namespace.Merge(m, src)
}
func (m *Namespace) XXX_Size() int {
	return xxx_messageInfo_Namespace.Size(m)
}
func (m *Namespace) XXX_DiscardUnknown() {
	xxx_messageInfo_Namespace.DiscardUnknown(m)
}

var xxx_messageInfo_Namespace proto.InternalMessageInfo

func (m *Namespace) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Namespace) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func init() {
	proto.RegisterType((*Namespace)(nil), "google.cloud.servicedirectory.v1beta1.Namespace")
	proto.RegisterMapType((map[string]string)(nil), "google.cloud.servicedirectory.v1beta1.Namespace.LabelsEntry")
}

func init() {
	proto.RegisterFile("google/cloud/servicedirectory/v1beta1/namespace.proto", fileDescriptor_aedd900d4e39df62)
}

var fileDescriptor_aedd900d4e39df62 = []byte{
	// 405 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xd1, 0xca, 0xd3, 0x30,
	0x14, 0xc7, 0x69, 0xe7, 0xf7, 0xc1, 0x97, 0x81, 0x48, 0x11, 0x9c, 0x45, 0x70, 0x08, 0x83, 0xed,
	0x26, 0x61, 0x8a, 0xa0, 0x51, 0x84, 0x55, 0xc5, 0x1b, 0x91, 0x31, 0x71, 0xa0, 0x0c, 0x24, 0xcd,
	0x62, 0xad, 0xcb, 0x72, 0x4a, 0xd2, 0x15, 0xc6, 0xd8, 0x4b, 0xf8, 0x18, 0x7b, 0x14, 0x9f, 0xc2,
	0xeb, 0x3d, 0x81, 0x97, 0xd2, 0xa4, 0xed, 0x46, 0xbd, 0x70, 0x77, 0xa7, 0x39, 0xff, 0xff, 0xff,
	0xfc, 0x92, 0x53, 0xf4, 0x34, 0x01, 0x48, 0xa4, 0x20, 0x5c, 0xc2, 0x66, 0x49, 0x8c, 0xd0, 0x45,
	0xca, 0xc5, 0x32, 0xd5, 0x82, 0xe7, 0xa0, 0xb7, 0xa4, 0x18, 0xc7, 0x22, 0x67, 0x63, 0xa2, 0xd8,
	0x5a, 0x98, 0x8c, 0x71, 0x81, 0x33, 0x0d, 0x39, 0x04, 0x03, 0x67, 0xc3, 0xd6, 0x86, 0xdb, 0x36,
	0x5c, 0xd9, 0xc2, 0x87, 0x55, 0x3a, 0xcb, 0x52, 0xf2, 0x2d, 0x15, 0x72, 0xf9, 0x35, 0x16, 0xdf,
	0x59, 0x91, 0x82, 0x76, 0x39, 0xe1, 0xfd, 0x33, 0x81, 0x16, 0x06, 0x36, 0xba, 0x1e, 0x11, 0x3e,
	0x38, 0x6b, 0x31, 0xa5, 0x20, 0x67, 0x79, 0x0a, 0xca, 0xb8, 0xee, 0xa3, 0x83, 0x8f, 0x6e, 0x3e,
	0xd4, 0x50, 0xc1, 0x3d, 0x74, 0xab, 0x24, 0xec, 0x79, 0x7d, 0x6f, 0x78, 0x13, 0x75, 0x7e, 0x4f,
	0xae, 0x66, 0xf6, 0x20, 0xf8, 0x8c, 0xae, 0x25, 0x8b, 0x85, 0x34, 0x3d, 0xbf, 0xdf, 0x19, 0x76,
	0x1f, 0xbf, 0xc4, 0x17, 0x81, 0xe3, 0x26, 0x1a, 0xbf, 0xb7, 0xf6, 0xb7, 0x2a, 0xd7, 0xdb, 0x32,
	0xd8, 0x9b, 0x55, 0x81, 0xe1, 0x73, 0xd4, 0x3d, 0xeb, 0x05, 0x77, 0x50, 0x67, 0x25, 0xb6, 0x8e,
	0x60, 0x56, 0x96, 0xc1, 0x5d, 0x74, 0x55, 0x30, 0xb9, 0x11, 0x3d, 0xdf, 0x9e, 0xb9, 0x0f, 0xea,
	0x3f, 0xf3, 0xa8, 0x3a, 0x4e, 0x56, 0x68, 0xf4, 0xcf, 0x70, 0x87, 0xc6, 0xb2, 0xd4, 0x60, 0x0e,
	0x6b, 0x72, 0xba, 0xde, 0xab, 0x4c, 0xc3, 0x0f, 0xc1, 0x73, 0x43, 0x76, 0x55, 0xb5, 0x27, 0x12,
	0xb8, 0x7b, 0x10, 0xb2, 0xab, 0xcb, 0xfd, 0x69, 0x4d, 0x86, 0xec, 0x9a, 0x7a, 0x1f, 0xfd, 0xf4,
	0xd1, 0x88, 0xc3, 0xfa, 0xb2, 0xbb, 0x47, 0xb7, 0x9b, 0xc1, 0xd3, 0xf2, 0xa9, 0xa7, 0xde, 0x97,
	0x4f, 0x95, 0x31, 0x01, 0xc9, 0x54, 0x82, 0x41, 0x27, 0x24, 0x11, 0xca, 0x2e, 0x82, 0x9c, 0xa0,
	0xff, 0xf3, 0x0f, 0xbd, 0x68, 0x37, 0xfe, 0x78, 0xde, 0xc1, 0x1f, 0xbc, 0x73, 0xd1, 0xaf, 0x2d,
	0xd3, 0x47, 0x27, 0x79, 0xd3, 0x30, 0xcd, 0xc7, 0x51, 0xe9, 0xfd, 0x55, 0xeb, 0x16, 0x56, 0xb7,
	0x68, 0xeb, 0x16, 0x73, 0x37, 0xe3, 0xe8, 0x0f, 0x9d, 0x8e, 0x52, 0x2b, 0xa4, 0xb4, 0xad, 0xa4,
	0xb4, 0x92, 0xc6, 0xd7, 0x96, 0xff, 0xc9, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd6, 0xa6, 0x6d,
	0x29, 0x02, 0x03, 0x00, 0x00,
}
