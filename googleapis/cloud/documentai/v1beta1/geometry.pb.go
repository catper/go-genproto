// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/documentai/v1beta1/geometry.proto

package documentai

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

// A vertex represents a 2D point in the image.
// NOTE: the vertex coordinates are in the same scale as the original image.
type Vertex struct {
	// X coordinate.
	X int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	// Y coordinate.
	Y                    int32    `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vertex) Reset()         { *m = Vertex{} }
func (m *Vertex) String() string { return proto.CompactTextString(m) }
func (*Vertex) ProtoMessage()    {}
func (*Vertex) Descriptor() ([]byte, []int) {
	return fileDescriptor_fcd2e1229b566cf9, []int{0}
}

func (m *Vertex) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vertex.Unmarshal(m, b)
}
func (m *Vertex) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vertex.Marshal(b, m, deterministic)
}
func (m *Vertex) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vertex.Merge(m, src)
}
func (m *Vertex) XXX_Size() int {
	return xxx_messageInfo_Vertex.Size(m)
}
func (m *Vertex) XXX_DiscardUnknown() {
	xxx_messageInfo_Vertex.DiscardUnknown(m)
}

var xxx_messageInfo_Vertex proto.InternalMessageInfo

func (m *Vertex) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Vertex) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

// A vertex represents a 2D point in the image.
// NOTE: the normalized vertex coordinates are relative to the original image
// and range from 0 to 1.
type NormalizedVertex struct {
	// X coordinate.
	X float32 `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	// Y coordinate.
	Y                    float32  `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NormalizedVertex) Reset()         { *m = NormalizedVertex{} }
func (m *NormalizedVertex) String() string { return proto.CompactTextString(m) }
func (*NormalizedVertex) ProtoMessage()    {}
func (*NormalizedVertex) Descriptor() ([]byte, []int) {
	return fileDescriptor_fcd2e1229b566cf9, []int{1}
}

func (m *NormalizedVertex) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NormalizedVertex.Unmarshal(m, b)
}
func (m *NormalizedVertex) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NormalizedVertex.Marshal(b, m, deterministic)
}
func (m *NormalizedVertex) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NormalizedVertex.Merge(m, src)
}
func (m *NormalizedVertex) XXX_Size() int {
	return xxx_messageInfo_NormalizedVertex.Size(m)
}
func (m *NormalizedVertex) XXX_DiscardUnknown() {
	xxx_messageInfo_NormalizedVertex.DiscardUnknown(m)
}

var xxx_messageInfo_NormalizedVertex proto.InternalMessageInfo

func (m *NormalizedVertex) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *NormalizedVertex) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

// A bounding polygon for the detected image annotation.
type BoundingPoly struct {
	// The bounding polygon vertices.
	Vertices []*Vertex `protobuf:"bytes,1,rep,name=vertices,proto3" json:"vertices,omitempty"`
	// The bounding polygon normalized vertices.
	NormalizedVertices   []*NormalizedVertex `protobuf:"bytes,2,rep,name=normalized_vertices,json=normalizedVertices,proto3" json:"normalized_vertices,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *BoundingPoly) Reset()         { *m = BoundingPoly{} }
func (m *BoundingPoly) String() string { return proto.CompactTextString(m) }
func (*BoundingPoly) ProtoMessage()    {}
func (*BoundingPoly) Descriptor() ([]byte, []int) {
	return fileDescriptor_fcd2e1229b566cf9, []int{2}
}

func (m *BoundingPoly) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoundingPoly.Unmarshal(m, b)
}
func (m *BoundingPoly) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoundingPoly.Marshal(b, m, deterministic)
}
func (m *BoundingPoly) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoundingPoly.Merge(m, src)
}
func (m *BoundingPoly) XXX_Size() int {
	return xxx_messageInfo_BoundingPoly.Size(m)
}
func (m *BoundingPoly) XXX_DiscardUnknown() {
	xxx_messageInfo_BoundingPoly.DiscardUnknown(m)
}

var xxx_messageInfo_BoundingPoly proto.InternalMessageInfo

func (m *BoundingPoly) GetVertices() []*Vertex {
	if m != nil {
		return m.Vertices
	}
	return nil
}

func (m *BoundingPoly) GetNormalizedVertices() []*NormalizedVertex {
	if m != nil {
		return m.NormalizedVertices
	}
	return nil
}

func init() {
	proto.RegisterType((*Vertex)(nil), "google.cloud.documentai.v1beta1.Vertex")
	proto.RegisterType((*NormalizedVertex)(nil), "google.cloud.documentai.v1beta1.NormalizedVertex")
	proto.RegisterType((*BoundingPoly)(nil), "google.cloud.documentai.v1beta1.BoundingPoly")
}

func init() {
	proto.RegisterFile("google/cloud/documentai/v1beta1/geometry.proto", fileDescriptor_fcd2e1229b566cf9)
}

var fileDescriptor_fcd2e1229b566cf9 = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xbf, 0x4a, 0xc4, 0x40,
	0x10, 0x87, 0xd9, 0x88, 0x87, 0xac, 0x27, 0x48, 0x6c, 0x82, 0x08, 0x1e, 0xa7, 0xe0, 0x55, 0xbb,
	0x44, 0x4b, 0xbb, 0x58, 0x88, 0x8d, 0x84, 0x14, 0x16, 0x36, 0xb2, 0x49, 0x86, 0x65, 0x21, 0x99,
	0x09, 0x9b, 0xcd, 0x71, 0xb1, 0xf3, 0xa9, 0x7c, 0x3d, 0xc9, 0x1f, 0x72, 0xdc, 0x81, 0xa4, 0x9c,
	0xc9, 0xf7, 0xfb, 0x32, 0x3b, 0xc3, 0x85, 0x26, 0xd2, 0x05, 0xc8, 0xac, 0xa0, 0x26, 0x97, 0x39,
	0x65, 0x4d, 0x09, 0xe8, 0x94, 0x91, 0xdb, 0x30, 0x05, 0xa7, 0x42, 0xa9, 0x81, 0x4a, 0x70, 0xb6,
	0x15, 0x95, 0x25, 0x47, 0xfe, 0xed, 0xc0, 0x8b, 0x9e, 0x17, 0x7b, 0x5e, 0x8c, 0xfc, 0xf5, 0xcd,
	0x28, 0x54, 0x95, 0x91, 0x0a, 0x91, 0x9c, 0x72, 0x86, 0xb0, 0x1e, 0xe2, 0xeb, 0x7b, 0xbe, 0xf8,
	0x00, 0xeb, 0x60, 0xe7, 0x2f, 0x39, 0xdb, 0x05, 0x6c, 0xc5, 0x36, 0xa7, 0x09, 0xeb, 0xab, 0x36,
	0xf0, 0x86, 0xaa, 0x5d, 0x0b, 0x7e, 0xf9, 0x4e, 0xb6, 0x54, 0x85, 0xf9, 0x86, 0xfc, 0x98, 0xf7,
	0x0e, 0x78, 0xaf, 0xe3, 0x7f, 0x19, 0x5f, 0x46, 0xd4, 0x60, 0x6e, 0x50, 0xc7, 0x54, 0xb4, 0xfe,
	0x0b, 0x3f, 0xdb, 0x82, 0x75, 0x26, 0x83, 0x3a, 0x60, 0xab, 0x93, 0xcd, 0xf9, 0xe3, 0x83, 0x98,
	0x19, 0x5c, 0x0c, 0xff, 0x49, 0xa6, 0xa0, 0x9f, 0xf2, 0x2b, 0x9c, 0xa6, 0xf8, 0x9a, 0x7c, 0x5e,
	0xef, 0x0b, 0x67, 0x7d, 0xc7, 0x2f, 0x48, 0x7c, 0x3c, 0xe8, 0x74, 0xb2, 0xe8, 0x87, 0xf1, 0xbb,
	0x8c, 0xca, 0x39, 0x59, 0x74, 0xf1, 0x3a, 0x9e, 0x21, 0xee, 0xd6, 0x18, 0xb3, 0xcf, 0xb7, 0x31,
	0xa1, 0xa9, 0x50, 0xa8, 0x05, 0x59, 0x2d, 0x35, 0x60, 0xbf, 0x64, 0x39, 0x7c, 0x52, 0x95, 0xa9,
	0xff, 0x3d, 0xeb, 0xf3, 0xbe, 0x95, 0x2e, 0xfa, 0xd4, 0xd3, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x90, 0x79, 0xe3, 0xf1, 0x0b, 0x02, 0x00, 0x00,
}
