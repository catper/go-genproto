// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/devtools/containeranalysis/v1beta1/common/common.proto

package common

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

// Kind represents the kinds of notes supported.
type NoteKind int32

const (
	// Unknown.
	NoteKind_NOTE_KIND_UNSPECIFIED NoteKind = 0
	// The note and occurrence represent a package vulnerability.
	NoteKind_VULNERABILITY NoteKind = 1
	// The note and occurrence assert build provenance.
	NoteKind_BUILD NoteKind = 2
	// This represents an image basis relationship.
	NoteKind_IMAGE NoteKind = 3
	// This represents a package installed via a package manager.
	NoteKind_PACKAGE NoteKind = 4
	// The note and occurrence track deployment events.
	NoteKind_DEPLOYMENT NoteKind = 5
	// The note and occurrence track the initial discovery status of a resource.
	NoteKind_DISCOVERY NoteKind = 6
	// This represents a logical "role" that can attest to artifacts.
	NoteKind_ATTESTATION NoteKind = 7
)

var NoteKind_name = map[int32]string{
	0: "NOTE_KIND_UNSPECIFIED",
	1: "VULNERABILITY",
	2: "BUILD",
	3: "IMAGE",
	4: "PACKAGE",
	5: "DEPLOYMENT",
	6: "DISCOVERY",
	7: "ATTESTATION",
}

var NoteKind_value = map[string]int32{
	"NOTE_KIND_UNSPECIFIED": 0,
	"VULNERABILITY":         1,
	"BUILD":                 2,
	"IMAGE":                 3,
	"PACKAGE":               4,
	"DEPLOYMENT":            5,
	"DISCOVERY":             6,
	"ATTESTATION":           7,
}

func (x NoteKind) String() string {
	return proto.EnumName(NoteKind_name, int32(x))
}

func (NoteKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_baf5189c00300310, []int{0}
}

// Metadata for any related URL information.
type RelatedUrl struct {
	// Specific URL associated with the resource.
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// Label to describe usage of the URL.
	Label                string   `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RelatedUrl) Reset()         { *m = RelatedUrl{} }
func (m *RelatedUrl) String() string { return proto.CompactTextString(m) }
func (*RelatedUrl) ProtoMessage()    {}
func (*RelatedUrl) Descriptor() ([]byte, []int) {
	return fileDescriptor_baf5189c00300310, []int{0}
}

func (m *RelatedUrl) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RelatedUrl.Unmarshal(m, b)
}
func (m *RelatedUrl) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RelatedUrl.Marshal(b, m, deterministic)
}
func (m *RelatedUrl) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RelatedUrl.Merge(m, src)
}
func (m *RelatedUrl) XXX_Size() int {
	return xxx_messageInfo_RelatedUrl.Size(m)
}
func (m *RelatedUrl) XXX_DiscardUnknown() {
	xxx_messageInfo_RelatedUrl.DiscardUnknown(m)
}

var xxx_messageInfo_RelatedUrl proto.InternalMessageInfo

func (m *RelatedUrl) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *RelatedUrl) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

// Verifiers (e.g. Kritis implementations) MUST verify signatures
// with respect to the trust anchors defined in policy (e.g. a Kritis policy).
// Typically this means that the verifier has been configured with a map from
// `public_key_id` to public key material (and any required parameters, e.g.
// signing algorithm).
//
// In particular, verification implementations MUST NOT treat the signature
// `public_key_id` as anything more than a key lookup hint. The `public_key_id`
// DOES NOT validate or authenticate a public key; it only provides a mechanism
// for quickly selecting a public key ALREADY CONFIGURED on the verifier through
// a trusted channel. Verification implementations MUST reject signatures in any
// of the following circumstances:
//   * The `public_key_id` is not recognized by the verifier.
//   * The public key that `public_key_id` refers to does not verify the
//     signature with respect to the payload.
//
// The `signature` contents SHOULD NOT be "attached" (where the payload is
// included with the serialized `signature` bytes). Verifiers MUST ignore any
// "attached" payload and only verify signatures with respect to explicitly
// provided payload (e.g. a `payload` field on the proto message that holds
// this Signature, or the canonical serialization of the proto message that
// holds this signature).
type Signature struct {
	// The content of the signature, an opaque bytestring.
	// The payload that this signature verifies MUST be unambiguously provided
	// with the Signature during verification. A wrapper message might provide
	// the payload explicitly. Alternatively, a message might have a canonical
	// serialization that can always be unambiguously computed to derive the
	// payload.
	Signature []byte `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	// The identifier for the public key that verifies this signature.
	//   * The `public_key_id` is required.
	//   * The `public_key_id` MUST be an RFC3986 conformant URI.
	//   * When possible, the `public_key_id` SHOULD be an immutable reference,
	//     such as a cryptographic digest.
	//
	// Examples of valid `public_key_id`s:
	//
	// OpenPGP V4 public key fingerprint:
	//   * "openpgp4fpr:74FAF3B861BDA0870C7B6DEF607E48D2A663AEEA"
	// See https://www.iana.org/assignments/uri-schemes/prov/openpgp4fpr for more
	// details on this scheme.
	//
	// RFC6920 digest-named SubjectPublicKeyInfo (digest of the DER
	// serialization):
	//   * "ni:///sha-256;cD9o9Cq6LG3jD0iKXqEi_vdjJGecm_iXkbqVoScViaU"
	//   * "nih:///sha-256;703f68f42aba2c6de30f488a5ea122fef76324679c9bf89791ba95a1271589a5"
	PublicKeyId          string   `protobuf:"bytes,2,opt,name=public_key_id,json=publicKeyId,proto3" json:"public_key_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Signature) Reset()         { *m = Signature{} }
func (m *Signature) String() string { return proto.CompactTextString(m) }
func (*Signature) ProtoMessage()    {}
func (*Signature) Descriptor() ([]byte, []int) {
	return fileDescriptor_baf5189c00300310, []int{1}
}

func (m *Signature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Signature.Unmarshal(m, b)
}
func (m *Signature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Signature.Marshal(b, m, deterministic)
}
func (m *Signature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Signature.Merge(m, src)
}
func (m *Signature) XXX_Size() int {
	return xxx_messageInfo_Signature.Size(m)
}
func (m *Signature) XXX_DiscardUnknown() {
	xxx_messageInfo_Signature.DiscardUnknown(m)
}

var xxx_messageInfo_Signature proto.InternalMessageInfo

func (m *Signature) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Signature) GetPublicKeyId() string {
	if m != nil {
		return m.PublicKeyId
	}
	return ""
}

func init() {
	proto.RegisterEnum("grafeas.v1beta1.NoteKind", NoteKind_name, NoteKind_value)
	proto.RegisterType((*RelatedUrl)(nil), "grafeas.v1beta1.RelatedUrl")
	proto.RegisterType((*Signature)(nil), "grafeas.v1beta1.Signature")
}

func init() {
	proto.RegisterFile("google/devtools/containeranalysis/v1beta1/common/common.proto", fileDescriptor_baf5189c00300310)
}

var fileDescriptor_baf5189c00300310 = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0xcd, 0xaa, 0xd3, 0x40,
	0x18, 0x35, 0xad, 0xbd, 0xd7, 0x7c, 0xb5, 0xde, 0x71, 0x50, 0xb8, 0x17, 0x5c, 0x48, 0x57, 0xe2,
	0x22, 0xe1, 0xa2, 0x3b, 0x71, 0x91, 0x36, 0x63, 0x19, 0x92, 0x26, 0x21, 0x3f, 0x85, 0xba, 0x09,
	0x93, 0x66, 0x1c, 0x06, 0xa7, 0x33, 0x25, 0x49, 0x0b, 0xc5, 0x47, 0xf0, 0x2d, 0x7c, 0x52, 0x69,
	0x52, 0x11, 0x5c, 0xb9, 0x3a, 0x3f, 0x33, 0xdf, 0x59, 0x9c, 0x03, 0x9f, 0x85, 0x31, 0x42, 0x71,
	0xb7, 0xe6, 0xa7, 0xce, 0x18, 0xd5, 0xba, 0x3b, 0xa3, 0x3b, 0x26, 0x35, 0x6f, 0x98, 0x66, 0xea,
	0xdc, 0xca, 0xd6, 0x3d, 0x3d, 0x56, 0xbc, 0x63, 0x8f, 0xee, 0xce, 0xec, 0xf7, 0x46, 0x5f, 0xc1,
	0x39, 0x34, 0xa6, 0x33, 0xf8, 0x4e, 0x34, 0xec, 0x1b, 0x67, 0xad, 0x73, 0xfd, 0x34, 0xff, 0x08,
	0x90, 0x72, 0xc5, 0x3a, 0x5e, 0x17, 0x8d, 0xc2, 0x08, 0xc6, 0xc7, 0x46, 0xdd, 0x5b, 0x6f, 0xad,
	0x77, 0x76, 0x7a, 0xa1, 0xf8, 0x15, 0x4c, 0x14, 0xab, 0xb8, 0xba, 0x1f, 0xf5, 0xde, 0x20, 0xe6,
	0x6b, 0xb0, 0x33, 0x29, 0x34, 0xeb, 0x8e, 0x0d, 0xc7, 0x6f, 0xc0, 0x6e, 0xff, 0x88, 0xfe, 0xf4,
	0x79, 0xfa, 0xd7, 0xc0, 0x73, 0x98, 0x1d, 0x8e, 0x95, 0x92, 0xbb, 0xf2, 0x3b, 0x3f, 0x97, 0xb2,
	0xbe, 0x06, 0x4d, 0x07, 0x33, 0xe0, 0x67, 0x5a, 0xbf, 0xff, 0x69, 0xc1, 0xb3, 0xc8, 0x74, 0x3c,
	0x90, 0xba, 0xc6, 0x0f, 0xf0, 0x3a, 0x8a, 0x73, 0x52, 0x06, 0x34, 0xf2, 0xcb, 0x22, 0xca, 0x12,
	0xb2, 0xa4, 0x5f, 0x28, 0xf1, 0xd1, 0x13, 0xfc, 0x12, 0x66, 0x9b, 0x22, 0x8c, 0x48, 0xea, 0x2d,
	0x68, 0x48, 0xf3, 0x2d, 0xb2, 0xb0, 0x0d, 0x93, 0x45, 0x41, 0x43, 0x1f, 0x8d, 0x2e, 0x94, 0xae,
	0xbd, 0x15, 0x41, 0x63, 0x3c, 0x85, 0xdb, 0xc4, 0x5b, 0x06, 0x17, 0xf1, 0x14, 0xbf, 0x00, 0xf0,
	0x49, 0x12, 0xc6, 0xdb, 0x35, 0x89, 0x72, 0x34, 0xc1, 0x33, 0xb0, 0x7d, 0x9a, 0x2d, 0xe3, 0x0d,
	0x49, 0xb7, 0xe8, 0x06, 0xdf, 0xc1, 0xd4, 0xcb, 0x73, 0x92, 0xe5, 0x5e, 0x4e, 0xe3, 0x08, 0xdd,
	0x2e, 0x7e, 0xc0, 0x83, 0x34, 0xce, 0x3f, 0x45, 0x39, 0x43, 0x8d, 0x89, 0xf5, 0x75, 0x33, 0x2c,
	0xe0, 0x08, 0xa3, 0x98, 0x16, 0x8e, 0x69, 0x84, 0x2b, 0xb8, 0xee, 0xeb, 0x75, 0x87, 0x27, 0x76,
	0x90, 0xed, 0xff, 0x0f, 0xf4, 0x69, 0x80, 0x5f, 0xa3, 0xf1, 0x2a, 0xf5, 0xaa, 0x9b, 0x3e, 0xe8,
	0xc3, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x02, 0xf2, 0x6d, 0xd4, 0xe8, 0x01, 0x00, 0x00,
}
