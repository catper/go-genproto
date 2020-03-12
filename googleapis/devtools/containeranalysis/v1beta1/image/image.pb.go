// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/devtools/containeranalysis/v1beta1/image/image.proto

package image

import (
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
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

// Instructions from Dockerfile.
type Layer_Directive int32

const (
	// Default value for unsupported/missing directive.
	Layer_DIRECTIVE_UNSPECIFIED Layer_Directive = 0
	// https://docs.docker.com/engine/reference/builder/
	Layer_MAINTAINER Layer_Directive = 1
	// https://docs.docker.com/engine/reference/builder/
	Layer_RUN Layer_Directive = 2
	// https://docs.docker.com/engine/reference/builder/
	Layer_CMD Layer_Directive = 3
	// https://docs.docker.com/engine/reference/builder/
	Layer_LABEL Layer_Directive = 4
	// https://docs.docker.com/engine/reference/builder/
	Layer_EXPOSE Layer_Directive = 5
	// https://docs.docker.com/engine/reference/builder/
	Layer_ENV Layer_Directive = 6
	// https://docs.docker.com/engine/reference/builder/
	Layer_ADD Layer_Directive = 7
	// https://docs.docker.com/engine/reference/builder/
	Layer_COPY Layer_Directive = 8
	// https://docs.docker.com/engine/reference/builder/
	Layer_ENTRYPOINT Layer_Directive = 9
	// https://docs.docker.com/engine/reference/builder/
	Layer_VOLUME Layer_Directive = 10
	// https://docs.docker.com/engine/reference/builder/
	Layer_USER Layer_Directive = 11
	// https://docs.docker.com/engine/reference/builder/
	Layer_WORKDIR Layer_Directive = 12
	// https://docs.docker.com/engine/reference/builder/
	Layer_ARG Layer_Directive = 13
	// https://docs.docker.com/engine/reference/builder/
	Layer_ONBUILD Layer_Directive = 14
	// https://docs.docker.com/engine/reference/builder/
	Layer_STOPSIGNAL Layer_Directive = 15
	// https://docs.docker.com/engine/reference/builder/
	Layer_HEALTHCHECK Layer_Directive = 16
	// https://docs.docker.com/engine/reference/builder/
	Layer_SHELL Layer_Directive = 17
)

var Layer_Directive_name = map[int32]string{
	0:  "DIRECTIVE_UNSPECIFIED",
	1:  "MAINTAINER",
	2:  "RUN",
	3:  "CMD",
	4:  "LABEL",
	5:  "EXPOSE",
	6:  "ENV",
	7:  "ADD",
	8:  "COPY",
	9:  "ENTRYPOINT",
	10: "VOLUME",
	11: "USER",
	12: "WORKDIR",
	13: "ARG",
	14: "ONBUILD",
	15: "STOPSIGNAL",
	16: "HEALTHCHECK",
	17: "SHELL",
}

var Layer_Directive_value = map[string]int32{
	"DIRECTIVE_UNSPECIFIED": 0,
	"MAINTAINER":            1,
	"RUN":                   2,
	"CMD":                   3,
	"LABEL":                 4,
	"EXPOSE":                5,
	"ENV":                   6,
	"ADD":                   7,
	"COPY":                  8,
	"ENTRYPOINT":            9,
	"VOLUME":                10,
	"USER":                  11,
	"WORKDIR":               12,
	"ARG":                   13,
	"ONBUILD":               14,
	"STOPSIGNAL":            15,
	"HEALTHCHECK":           16,
	"SHELL":                 17,
}

func (x Layer_Directive) String() string {
	return proto.EnumName(Layer_Directive_name, int32(x))
}

func (Layer_Directive) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ab0486c83d3ae0f1, []int{0, 0}
}

// Layer holds metadata specific to a layer of a Docker image.
type Layer struct {
	// Required. The recovered Dockerfile directive used to construct this layer.
	Directive Layer_Directive `protobuf:"varint,1,opt,name=directive,proto3,enum=grafeas.v1beta1.image.Layer_Directive" json:"directive,omitempty"`
	// The recovered arguments to the Dockerfile directive.
	Arguments            string   `protobuf:"bytes,2,opt,name=arguments,proto3" json:"arguments,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Layer) Reset()         { *m = Layer{} }
func (m *Layer) String() string { return proto.CompactTextString(m) }
func (*Layer) ProtoMessage()    {}
func (*Layer) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab0486c83d3ae0f1, []int{0}
}

func (m *Layer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Layer.Unmarshal(m, b)
}
func (m *Layer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Layer.Marshal(b, m, deterministic)
}
func (m *Layer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Layer.Merge(m, src)
}
func (m *Layer) XXX_Size() int {
	return xxx_messageInfo_Layer.Size(m)
}
func (m *Layer) XXX_DiscardUnknown() {
	xxx_messageInfo_Layer.DiscardUnknown(m)
}

var xxx_messageInfo_Layer proto.InternalMessageInfo

func (m *Layer) GetDirective() Layer_Directive {
	if m != nil {
		return m.Directive
	}
	return Layer_DIRECTIVE_UNSPECIFIED
}

func (m *Layer) GetArguments() string {
	if m != nil {
		return m.Arguments
	}
	return ""
}

// A set of properties that uniquely identify a given Docker image.
type Fingerprint struct {
	// Required. The layer ID of the final layer in the Docker image's v1
	// representation.
	V1Name string `protobuf:"bytes,1,opt,name=v1_name,json=v1Name,proto3" json:"v1_name,omitempty"`
	// Required. The ordered list of v2 blobs that represent a given image.
	V2Blob []string `protobuf:"bytes,2,rep,name=v2_blob,json=v2Blob,proto3" json:"v2_blob,omitempty"`
	// Output only. The name of the image's v2 blobs computed via:
	//   [bottom] := v2_blob[bottom]
	//   [N] := sha256(v2_blob[N] + " " + v2_name[N+1])
	// Only the name of the final blob is kept.
	V2Name               string   `protobuf:"bytes,3,opt,name=v2_name,json=v2Name,proto3" json:"v2_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Fingerprint) Reset()         { *m = Fingerprint{} }
func (m *Fingerprint) String() string { return proto.CompactTextString(m) }
func (*Fingerprint) ProtoMessage()    {}
func (*Fingerprint) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab0486c83d3ae0f1, []int{1}
}

func (m *Fingerprint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Fingerprint.Unmarshal(m, b)
}
func (m *Fingerprint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Fingerprint.Marshal(b, m, deterministic)
}
func (m *Fingerprint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fingerprint.Merge(m, src)
}
func (m *Fingerprint) XXX_Size() int {
	return xxx_messageInfo_Fingerprint.Size(m)
}
func (m *Fingerprint) XXX_DiscardUnknown() {
	xxx_messageInfo_Fingerprint.DiscardUnknown(m)
}

var xxx_messageInfo_Fingerprint proto.InternalMessageInfo

func (m *Fingerprint) GetV1Name() string {
	if m != nil {
		return m.V1Name
	}
	return ""
}

func (m *Fingerprint) GetV2Blob() []string {
	if m != nil {
		return m.V2Blob
	}
	return nil
}

func (m *Fingerprint) GetV2Name() string {
	if m != nil {
		return m.V2Name
	}
	return ""
}

// Basis describes the base image portion (Note) of the DockerImage
// relationship. Linked occurrences are derived from this or an
// equivalent image via:
//   FROM <Basis.resource_url>
// Or an equivalent reference, e.g. a tag of the resource_url.
type Basis struct {
	// Required. Immutable. The resource_url for the resource representing the
	// basis of associated occurrence images.
	ResourceUrl string `protobuf:"bytes,1,opt,name=resource_url,json=resourceUrl,proto3" json:"resource_url,omitempty"`
	// Required. Immutable. The fingerprint of the base image.
	Fingerprint          *Fingerprint `protobuf:"bytes,2,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Basis) Reset()         { *m = Basis{} }
func (m *Basis) String() string { return proto.CompactTextString(m) }
func (*Basis) ProtoMessage()    {}
func (*Basis) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab0486c83d3ae0f1, []int{2}
}

func (m *Basis) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Basis.Unmarshal(m, b)
}
func (m *Basis) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Basis.Marshal(b, m, deterministic)
}
func (m *Basis) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Basis.Merge(m, src)
}
func (m *Basis) XXX_Size() int {
	return xxx_messageInfo_Basis.Size(m)
}
func (m *Basis) XXX_DiscardUnknown() {
	xxx_messageInfo_Basis.DiscardUnknown(m)
}

var xxx_messageInfo_Basis proto.InternalMessageInfo

func (m *Basis) GetResourceUrl() string {
	if m != nil {
		return m.ResourceUrl
	}
	return ""
}

func (m *Basis) GetFingerprint() *Fingerprint {
	if m != nil {
		return m.Fingerprint
	}
	return nil
}

// Details of an image occurrence.
type Details struct {
	// Required. Immutable. The child image derived from the base image.
	DerivedImage         *Derived `protobuf:"bytes,1,opt,name=derived_image,json=derivedImage,proto3" json:"derived_image,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Details) Reset()         { *m = Details{} }
func (m *Details) String() string { return proto.CompactTextString(m) }
func (*Details) ProtoMessage()    {}
func (*Details) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab0486c83d3ae0f1, []int{3}
}

func (m *Details) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Details.Unmarshal(m, b)
}
func (m *Details) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Details.Marshal(b, m, deterministic)
}
func (m *Details) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Details.Merge(m, src)
}
func (m *Details) XXX_Size() int {
	return xxx_messageInfo_Details.Size(m)
}
func (m *Details) XXX_DiscardUnknown() {
	xxx_messageInfo_Details.DiscardUnknown(m)
}

var xxx_messageInfo_Details proto.InternalMessageInfo

func (m *Details) GetDerivedImage() *Derived {
	if m != nil {
		return m.DerivedImage
	}
	return nil
}

// Derived describes the derived image portion (Occurrence) of the DockerImage
// relationship. This image would be produced from a Dockerfile with FROM
// <DockerImage.Basis in attached Note>.
type Derived struct {
	// Required. The fingerprint of the derived image.
	Fingerprint *Fingerprint `protobuf:"bytes,1,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
	// Output only. The number of layers by which this image differs from the
	// associated image basis.
	Distance int32 `protobuf:"varint,2,opt,name=distance,proto3" json:"distance,omitempty"`
	// This contains layer-specific metadata, if populated it has length
	// "distance" and is ordered with [distance] being the layer immediately
	// following the base image and [1] being the final layer.
	LayerInfo []*Layer `protobuf:"bytes,3,rep,name=layer_info,json=layerInfo,proto3" json:"layer_info,omitempty"`
	// Output only. This contains the base image URL for the derived image
	// occurrence.
	BaseResourceUrl      string   `protobuf:"bytes,4,opt,name=base_resource_url,json=baseResourceUrl,proto3" json:"base_resource_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Derived) Reset()         { *m = Derived{} }
func (m *Derived) String() string { return proto.CompactTextString(m) }
func (*Derived) ProtoMessage()    {}
func (*Derived) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab0486c83d3ae0f1, []int{4}
}

func (m *Derived) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Derived.Unmarshal(m, b)
}
func (m *Derived) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Derived.Marshal(b, m, deterministic)
}
func (m *Derived) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Derived.Merge(m, src)
}
func (m *Derived) XXX_Size() int {
	return xxx_messageInfo_Derived.Size(m)
}
func (m *Derived) XXX_DiscardUnknown() {
	xxx_messageInfo_Derived.DiscardUnknown(m)
}

var xxx_messageInfo_Derived proto.InternalMessageInfo

func (m *Derived) GetFingerprint() *Fingerprint {
	if m != nil {
		return m.Fingerprint
	}
	return nil
}

func (m *Derived) GetDistance() int32 {
	if m != nil {
		return m.Distance
	}
	return 0
}

func (m *Derived) GetLayerInfo() []*Layer {
	if m != nil {
		return m.LayerInfo
	}
	return nil
}

func (m *Derived) GetBaseResourceUrl() string {
	if m != nil {
		return m.BaseResourceUrl
	}
	return ""
}

func init() {
	proto.RegisterEnum("grafeas.v1beta1.image.Layer_Directive", Layer_Directive_name, Layer_Directive_value)
	proto.RegisterType((*Layer)(nil), "grafeas.v1beta1.image.Layer")
	proto.RegisterType((*Fingerprint)(nil), "grafeas.v1beta1.image.Fingerprint")
	proto.RegisterType((*Basis)(nil), "grafeas.v1beta1.image.Basis")
	proto.RegisterType((*Details)(nil), "grafeas.v1beta1.image.Details")
	proto.RegisterType((*Derived)(nil), "grafeas.v1beta1.image.Derived")
}

func init() {
	proto.RegisterFile("google/devtools/containeranalysis/v1beta1/image/image.proto", fileDescriptor_ab0486c83d3ae0f1)
}

var fileDescriptor_ab0486c83d3ae0f1 = []byte{
	// 613 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcf, 0x6e, 0xda, 0x40,
	0x10, 0xc6, 0x6b, 0x08, 0x10, 0x8f, 0xf3, 0x67, 0xb3, 0x52, 0x54, 0x5a, 0x45, 0x15, 0xe5, 0x50,
	0x45, 0x3d, 0x18, 0x41, 0x8f, 0x39, 0x01, 0xde, 0x04, 0x2b, 0x8e, 0x8d, 0x16, 0x93, 0x26, 0xbd,
	0x58, 0x0b, 0x2c, 0xd6, 0x4a, 0xc6, 0x8b, 0xd6, 0x0e, 0x6a, 0x5e, 0xa7, 0xcf, 0xd4, 0x73, 0x9f,
	0xa1, 0x8f, 0x50, 0x79, 0x21, 0x21, 0xad, 0x92, 0xaa, 0xea, 0x65, 0xb5, 0xcc, 0x37, 0xbf, 0x8f,
	0xf1, 0xec, 0x0c, 0x9c, 0xc5, 0x52, 0xc6, 0x09, 0x6f, 0xcd, 0xf8, 0x2a, 0x97, 0x32, 0xc9, 0x5a,
	0x53, 0x99, 0xe6, 0x4c, 0xa4, 0x5c, 0xb1, 0x94, 0x25, 0xf7, 0x99, 0xc8, 0x5a, 0xab, 0xf6, 0x84,
	0xe7, 0xac, 0xdd, 0x12, 0x0b, 0x16, 0xf3, 0xf5, 0x69, 0x2f, 0x95, 0xcc, 0x25, 0x3e, 0x8e, 0x15,
	0x9b, 0x73, 0x96, 0xd9, 0x9b, 0x14, 0x5b, 0x8b, 0xcd, 0x1f, 0x25, 0xa8, 0x78, 0xec, 0x9e, 0x2b,
	0xec, 0x80, 0x39, 0x13, 0x8a, 0x4f, 0x73, 0xb1, 0xe2, 0x75, 0xa3, 0x61, 0x9c, 0x1e, 0x74, 0x3e,
	0xd8, 0xcf, 0x42, 0xb6, 0x06, 0x6c, 0xe7, 0x21, 0x9b, 0x6e, 0x41, 0x7c, 0x02, 0x26, 0x53, 0xf1,
	0xdd, 0x82, 0xa7, 0x79, 0x56, 0x2f, 0x35, 0x8c, 0x53, 0x93, 0x6e, 0x03, 0xcd, 0x9f, 0x06, 0x98,
	0x8f, 0x18, 0x7e, 0x03, 0xc7, 0x8e, 0x4b, 0x49, 0x3f, 0x74, 0xaf, 0x49, 0x34, 0xf6, 0x47, 0x43,
	0xd2, 0x77, 0xcf, 0x5d, 0xe2, 0xa0, 0x57, 0xf8, 0x00, 0xe0, 0xaa, 0xeb, 0xfa, 0x61, 0xd7, 0xf5,
	0x09, 0x45, 0x06, 0xae, 0x41, 0x99, 0x8e, 0x7d, 0x54, 0x2a, 0x2e, 0xfd, 0x2b, 0x07, 0x95, 0xb1,
	0x09, 0x15, 0xaf, 0xdb, 0x23, 0x1e, 0xda, 0xc1, 0x00, 0x55, 0x72, 0x33, 0x0c, 0x46, 0x04, 0x55,
	0x0a, 0x9d, 0xf8, 0xd7, 0xa8, 0x5a, 0x5c, 0xba, 0x8e, 0x83, 0x6a, 0x78, 0x17, 0x76, 0xfa, 0xc1,
	0xf0, 0x16, 0xed, 0x16, 0xa6, 0xc4, 0x0f, 0xe9, 0xed, 0x30, 0x70, 0xfd, 0x10, 0x99, 0x05, 0x77,
	0x1d, 0x78, 0xe3, 0x2b, 0x82, 0xa0, 0xc8, 0x1a, 0x8f, 0x08, 0x45, 0x16, 0xb6, 0xa0, 0xf6, 0x39,
	0xa0, 0x97, 0x8e, 0x4b, 0xd1, 0x9e, 0x76, 0xa1, 0x17, 0x68, 0xbf, 0x88, 0x06, 0x7e, 0x6f, 0xec,
	0x7a, 0x0e, 0x3a, 0x28, 0x8c, 0x46, 0x61, 0x30, 0x1c, 0xb9, 0x17, 0x7e, 0xd7, 0x43, 0x87, 0xf8,
	0x10, 0xac, 0x01, 0xe9, 0x7a, 0xe1, 0xa0, 0x3f, 0x20, 0xfd, 0x4b, 0x84, 0x8a, 0xe2, 0x46, 0x03,
	0xe2, 0x79, 0xe8, 0xa8, 0x79, 0x03, 0xd6, 0xb9, 0x48, 0x63, 0xae, 0x96, 0x4a, 0xa4, 0x39, 0x7e,
	0x0d, 0xb5, 0x55, 0x3b, 0x4a, 0xd9, 0x62, 0xdd, 0x63, 0x93, 0x56, 0x57, 0x6d, 0x9f, 0x2d, 0xb8,
	0x16, 0x3a, 0xd1, 0x24, 0x91, 0x93, 0x7a, 0xa9, 0x51, 0xd6, 0x42, 0xa7, 0x97, 0xc8, 0xc9, 0x46,
	0xd0, 0x44, 0x79, 0x43, 0x74, 0x0a, 0xa2, 0xb9, 0x84, 0x4a, 0x8f, 0x65, 0x22, 0xc3, 0xef, 0x61,
	0x4f, 0xf1, 0x4c, 0xde, 0xa9, 0x29, 0x8f, 0xee, 0x54, 0xb2, 0x31, 0xb6, 0x1e, 0x62, 0x63, 0x95,
	0x60, 0x07, 0xac, 0xf9, 0xb6, 0x0a, 0xfd, 0x30, 0x56, 0xa7, 0xf9, 0xc2, 0xf3, 0x3e, 0xa9, 0x97,
	0x3e, 0xc5, 0x9a, 0x3e, 0xd4, 0x1c, 0x9e, 0x33, 0x91, 0x64, 0xb8, 0x0f, 0xfb, 0x33, 0xae, 0xc4,
	0x8a, 0xcf, 0x22, 0x0d, 0xe9, 0x3f, 0xb5, 0x3a, 0xef, 0x5e, 0xb0, 0x74, 0xd6, 0xb9, 0x74, 0x6f,
	0x03, 0xb9, 0x7a, 0xf8, 0xbe, 0x1b, 0x85, 0xa1, 0x0e, 0xfc, 0x59, 0xa1, 0xf1, 0x5f, 0x15, 0xe2,
	0xb7, 0xb0, 0x3b, 0x13, 0x59, 0xce, 0xd2, 0x29, 0xd7, 0x1f, 0x59, 0xa1, 0x8f, 0xbf, 0xf1, 0x19,
	0x40, 0x52, 0x0c, 0x6e, 0x24, 0xd2, 0xb9, 0xac, 0x97, 0x1b, 0xe5, 0x53, 0xab, 0x73, 0xf2, 0xb7,
	0x09, 0xa7, 0xa6, 0xce, 0x77, 0xd3, 0xb9, 0xc4, 0x1f, 0xe1, 0x68, 0xc2, 0x32, 0x1e, 0xfd, 0xd6,
	0xe8, 0x1d, 0xdd, 0xe8, 0xc3, 0x42, 0xa0, 0xdb, 0x66, 0xf7, 0xbe, 0x42, 0x5d, 0xc8, 0xe7, 0x8d,
	0x87, 0xc6, 0x97, 0x70, 0xbd, 0xc5, 0x76, 0x2c, 0x13, 0x96, 0xc6, 0xb6, 0x54, 0x71, 0x2b, 0xe6,
	0xa9, 0x5e, 0xd2, 0xd6, 0x5a, 0x62, 0x4b, 0x91, 0xfd, 0xf3, 0x92, 0x9f, 0xe9, 0xf3, 0x5b, 0xa9,
	0x7c, 0x41, 0xbb, 0x93, 0xaa, 0xb6, 0xf9, 0xf4, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x06, 0xe8,
	0x32, 0x2a, 0x04, 0x00, 0x00,
}
