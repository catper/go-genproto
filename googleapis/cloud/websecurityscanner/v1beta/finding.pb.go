// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/websecurityscanner/v1beta/finding.proto

package websecurityscanner

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

// A Finding resource represents a vulnerability instance identified during a
// ScanRun.
type Finding struct {
	// The resource name of the Finding. The name follows the format of
	// 'projects/{projectId}/scanConfigs/{scanConfigId}/scanruns/{scanRunId}/findings/{findingId}'.
	// The finding IDs are generated by the system.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The type of the Finding.
	// Detailed and up-to-date information on findings can be found here:
	// https://cloud.google.com/security-scanner/docs/scan-result-details
	FindingType string `protobuf:"bytes,2,opt,name=finding_type,json=findingType,proto3" json:"finding_type,omitempty"`
	// The http method of the request that triggered the vulnerability, in
	// uppercase.
	HttpMethod string `protobuf:"bytes,3,opt,name=http_method,json=httpMethod,proto3" json:"http_method,omitempty"`
	// The URL produced by the server-side fuzzer and used in the request that
	// triggered the vulnerability.
	FuzzedUrl string `protobuf:"bytes,4,opt,name=fuzzed_url,json=fuzzedUrl,proto3" json:"fuzzed_url,omitempty"`
	// The body of the request that triggered the vulnerability.
	Body string `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	// The description of the vulnerability.
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	// The URL containing human-readable payload that user can leverage to
	// reproduce the vulnerability.
	ReproductionUrl string `protobuf:"bytes,7,opt,name=reproduction_url,json=reproductionUrl,proto3" json:"reproduction_url,omitempty"`
	// If the vulnerability was originated from nested IFrame, the immediate
	// parent IFrame is reported.
	FrameUrl string `protobuf:"bytes,8,opt,name=frame_url,json=frameUrl,proto3" json:"frame_url,omitempty"`
	// The URL where the browser lands when the vulnerability is detected.
	FinalUrl string `protobuf:"bytes,9,opt,name=final_url,json=finalUrl,proto3" json:"final_url,omitempty"`
	// The tracking ID uniquely identifies a vulnerability instance across
	// multiple ScanRuns.
	TrackingId string `protobuf:"bytes,10,opt,name=tracking_id,json=trackingId,proto3" json:"tracking_id,omitempty"`
	// An addon containing information reported for a vulnerability with an HTML
	// form, if any.
	Form *Form `protobuf:"bytes,16,opt,name=form,proto3" json:"form,omitempty"`
	// An addon containing information about outdated libraries.
	OutdatedLibrary *OutdatedLibrary `protobuf:"bytes,11,opt,name=outdated_library,json=outdatedLibrary,proto3" json:"outdated_library,omitempty"`
	// An addon containing detailed information regarding any resource causing the
	// vulnerability such as JavaScript sources, image, audio files, etc.
	ViolatingResource *ViolatingResource `protobuf:"bytes,12,opt,name=violating_resource,json=violatingResource,proto3" json:"violating_resource,omitempty"`
	// An addon containing information about vulnerable or missing HTTP headers.
	VulnerableHeaders *VulnerableHeaders `protobuf:"bytes,15,opt,name=vulnerable_headers,json=vulnerableHeaders,proto3" json:"vulnerable_headers,omitempty"`
	// An addon containing information about request parameters which were found
	// to be vulnerable.
	VulnerableParameters *VulnerableParameters `protobuf:"bytes,13,opt,name=vulnerable_parameters,json=vulnerableParameters,proto3" json:"vulnerable_parameters,omitempty"`
	// An addon containing information reported for an XSS, if any.
	Xss                  *Xss     `protobuf:"bytes,14,opt,name=xss,proto3" json:"xss,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Finding) Reset()         { *m = Finding{} }
func (m *Finding) String() string { return proto.CompactTextString(m) }
func (*Finding) ProtoMessage()    {}
func (*Finding) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1082375f5efc351, []int{0}
}

func (m *Finding) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Finding.Unmarshal(m, b)
}
func (m *Finding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Finding.Marshal(b, m, deterministic)
}
func (m *Finding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Finding.Merge(m, src)
}
func (m *Finding) XXX_Size() int {
	return xxx_messageInfo_Finding.Size(m)
}
func (m *Finding) XXX_DiscardUnknown() {
	xxx_messageInfo_Finding.DiscardUnknown(m)
}

var xxx_messageInfo_Finding proto.InternalMessageInfo

func (m *Finding) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Finding) GetFindingType() string {
	if m != nil {
		return m.FindingType
	}
	return ""
}

func (m *Finding) GetHttpMethod() string {
	if m != nil {
		return m.HttpMethod
	}
	return ""
}

func (m *Finding) GetFuzzedUrl() string {
	if m != nil {
		return m.FuzzedUrl
	}
	return ""
}

func (m *Finding) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Finding) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Finding) GetReproductionUrl() string {
	if m != nil {
		return m.ReproductionUrl
	}
	return ""
}

func (m *Finding) GetFrameUrl() string {
	if m != nil {
		return m.FrameUrl
	}
	return ""
}

func (m *Finding) GetFinalUrl() string {
	if m != nil {
		return m.FinalUrl
	}
	return ""
}

func (m *Finding) GetTrackingId() string {
	if m != nil {
		return m.TrackingId
	}
	return ""
}

func (m *Finding) GetForm() *Form {
	if m != nil {
		return m.Form
	}
	return nil
}

func (m *Finding) GetOutdatedLibrary() *OutdatedLibrary {
	if m != nil {
		return m.OutdatedLibrary
	}
	return nil
}

func (m *Finding) GetViolatingResource() *ViolatingResource {
	if m != nil {
		return m.ViolatingResource
	}
	return nil
}

func (m *Finding) GetVulnerableHeaders() *VulnerableHeaders {
	if m != nil {
		return m.VulnerableHeaders
	}
	return nil
}

func (m *Finding) GetVulnerableParameters() *VulnerableParameters {
	if m != nil {
		return m.VulnerableParameters
	}
	return nil
}

func (m *Finding) GetXss() *Xss {
	if m != nil {
		return m.Xss
	}
	return nil
}

func init() {
	proto.RegisterType((*Finding)(nil), "google.cloud.websecurityscanner.v1beta.Finding")
}

func init() {
	proto.RegisterFile("google/cloud/websecurityscanner/v1beta/finding.proto", fileDescriptor_a1082375f5efc351)
}

var fileDescriptor_a1082375f5efc351 = []byte{
	// 605 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x5f, 0x4f, 0xd5, 0x3e,
	0x1c, 0xc6, 0x33, 0xe0, 0x07, 0x9c, 0x8e, 0x9f, 0x40, 0xa3, 0x49, 0xc5, 0x18, 0xd0, 0x0b, 0x02,
	0x6a, 0xd6, 0xf8, 0x27, 0x31, 0xa2, 0x26, 0x0a, 0x09, 0x6a, 0xa2, 0x91, 0x0c, 0x45, 0x62, 0x48,
	0x96, 0x6e, 0xed, 0xd9, 0xa9, 0x6e, 0xed, 0x6c, 0xbb, 0xa3, 0x07, 0x42, 0xe2, 0x85, 0xef, 0xcb,
	0x7b, 0x5f, 0x8a, 0x37, 0xbe, 0x05, 0xd3, 0x76, 0xc3, 0x23, 0x90, 0x78, 0xce, 0x5d, 0xfb, 0x3c,
	0xdf, 0xe7, 0xf9, 0xac, 0xd9, 0x56, 0x70, 0x2f, 0x97, 0x32, 0x2f, 0x18, 0xce, 0x0a, 0x59, 0x53,
	0xfc, 0x99, 0xa5, 0x9a, 0x65, 0xb5, 0xe2, 0x66, 0xa0, 0x33, 0x22, 0x04, 0x53, 0xb8, 0x7f, 0x3b,
	0x65, 0x86, 0xe0, 0x2e, 0x17, 0x94, 0x8b, 0x3c, 0xaa, 0x94, 0x34, 0x12, 0xae, 0xfa, 0x54, 0xe4,
	0x52, 0xd1, 0xd9, 0x54, 0xe4, 0x53, 0x4b, 0x97, 0x9b, 0x76, 0x52, 0x71, 0xac, 0x98, 0x96, 0xb5,
	0xca, 0x98, 0xaf, 0x58, 0xda, 0x18, 0x0f, 0x9c, 0x10, 0x4a, 0xa5, 0xf0, 0xd9, 0xeb, 0xbf, 0x66,
	0xc0, 0xcc, 0xb6, 0xd7, 0x21, 0x04, 0x53, 0x82, 0x94, 0x0c, 0x05, 0x2b, 0xc1, 0x5a, 0x27, 0x76,
	0x6b, 0x78, 0x0d, 0xcc, 0xb5, 0x31, 0x33, 0xa8, 0x18, 0x9a, 0x70, 0x5e, 0xd8, 0x68, 0x6f, 0x06,
	0x15, 0x83, 0xcb, 0x20, 0xec, 0x19, 0x53, 0x25, 0x25, 0x33, 0x3d, 0x49, 0xd1, 0xa4, 0x9b, 0x00,
	0x56, 0x7a, 0xe5, 0x14, 0x78, 0x15, 0x80, 0x6e, 0x7d, 0x78, 0xc8, 0x68, 0x52, 0xab, 0x02, 0x4d,
	0x39, 0xbf, 0xe3, 0x95, 0xb7, 0xaa, 0xb0, 0xd8, 0x54, 0xd2, 0x01, 0xfa, 0xcf, 0x63, 0xed, 0x1a,
	0xae, 0x80, 0x90, 0x32, 0x9d, 0x29, 0x5e, 0x19, 0x2e, 0x05, 0x9a, 0xf6, 0xd4, 0x21, 0x09, 0xae,
	0x83, 0x05, 0xc5, 0x2a, 0x25, 0x69, 0x9d, 0xd9, 0xbd, 0xab, 0x9e, 0x71, 0x63, 0xf3, 0xc3, 0xba,
	0x05, 0x5c, 0x01, 0x9d, 0xae, 0x22, 0x25, 0x73, 0x33, 0xb3, 0x6e, 0x66, 0xd6, 0x09, 0xad, 0xc9,
	0x05, 0x29, 0x9c, 0xd9, 0x69, 0x4c, 0x2b, 0x58, 0x73, 0x19, 0x84, 0x46, 0x91, 0xec, 0xa3, 0x3d,
	0x3e, 0xa7, 0x08, 0xf8, 0xa3, 0xb5, 0xd2, 0x0b, 0x0a, 0x9f, 0x80, 0xa9, 0xae, 0x54, 0x25, 0x5a,
	0x58, 0x09, 0xd6, 0xc2, 0x3b, 0xb7, 0xa2, 0xd1, 0x5e, 0x66, 0xb4, 0x2d, 0x55, 0x19, 0xbb, 0x24,
	0x4c, 0xc1, 0x82, 0xac, 0x0d, 0x25, 0x86, 0xd1, 0xa4, 0xe0, 0xa9, 0x22, 0x6a, 0x80, 0x42, 0xd7,
	0x76, 0x7f, 0xd4, 0xb6, 0xd7, 0x4d, 0xfe, 0xa5, 0x8f, 0xc7, 0xf3, 0xf2, 0x6f, 0x01, 0xf6, 0x00,
	0xec, 0x73, 0x59, 0x10, 0x63, 0xcf, 0xd1, 0x7e, 0x3c, 0x68, 0xce, 0x51, 0x1e, 0x8c, 0x4a, 0xd9,
	0x6b, 0x1b, 0xe2, 0xa6, 0x20, 0x5e, 0xec, 0x9f, 0x96, 0x1c, 0xa9, 0x2e, 0x04, 0x53, 0x24, 0x2d,
	0x58, 0xd2, 0x63, 0x84, 0x32, 0xa5, 0xd1, 0xfc, 0x98, 0xa4, 0x93, 0x86, 0xe7, 0xbe, 0x20, 0x5e,
	0xec, 0x9f, 0x96, 0xe0, 0x27, 0x70, 0x69, 0x88, 0x54, 0x11, 0xfb, 0x3e, 0x8d, 0x85, 0xfd, 0xef,
	0x60, 0x8f, 0xc6, 0x87, 0xed, 0x9c, 0x74, 0xc4, 0x17, 0xfb, 0xe7, 0xa8, 0xf0, 0x31, 0x98, 0xfc,
	0xa2, 0x35, 0xba, 0xe0, 0x00, 0x37, 0x47, 0x05, 0xec, 0x6b, 0x1d, 0xdb, 0xdc, 0xc6, 0xb7, 0xe0,
	0xe7, 0xd3, 0xaf, 0x01, 0x58, 0x3f, 0x67, 0xd4, 0x57, 0x91, 0x8a, 0xeb, 0x28, 0x93, 0x25, 0x6e,
	0x7f, 0xc8, 0xdd, 0x4a, 0xc9, 0x0f, 0x2c, 0x33, 0x1a, 0x1f, 0x35, 0xab, 0x63, 0x6c, 0x23, 0x5b,
	0x52, 0x74, 0x79, 0xae, 0xf1, 0x91, 0xdd, 0x24, 0x99, 0xdb, 0x79, 0x27, 0xae, 0x45, 0x2b, 0xab,
	0x5a, 0x1c, 0xb7, 0xbf, 0xbc, 0xc6, 0x47, 0xcd, 0xea, 0x78, 0xf3, 0x7b, 0x00, 0x6e, 0x64, 0xb2,
	0x1c, 0xf1, 0xf1, 0x37, 0xe7, 0x9a, 0x87, 0xd9, 0xb1, 0xd7, 0xc5, 0x4e, 0xf0, 0x7e, 0xbf, 0xc9,
	0xe5, 0xb2, 0x20, 0x22, 0x8f, 0xa4, 0xca, 0x71, 0xce, 0x84, 0xbb, 0x4c, 0xf0, 0x9f, 0x63, 0xfc,
	0xeb, 0x2e, 0x7a, 0x78, 0xd6, 0xf9, 0x31, 0xb1, 0xfa, 0xcc, 0xe5, 0x0f, 0xb6, 0x6c, 0xf6, 0xe0,
	0x1d, 0x4b, 0x77, 0x9b, 0x89, 0x5d, 0x3f, 0x71, 0xb0, 0xe7, 0xb2, 0xe9, 0xb4, 0xa3, 0xdd, 0xfd,
	0x1d, 0x00, 0x00, 0xff, 0xff, 0x34, 0x35, 0xea, 0x27, 0x71, 0x05, 0x00, 0x00,
}
