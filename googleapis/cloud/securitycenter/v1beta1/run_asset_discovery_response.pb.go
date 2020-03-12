// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/securitycenter/v1beta1/run_asset_discovery_response.proto

package securitycenter

import (
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
	duration "github.com/catper/protobuf/ptypes/duration"
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

// The state of an asset discovery run.
type RunAssetDiscoveryResponse_State int32

const (
	// Asset discovery run state was unspecified.
	RunAssetDiscoveryResponse_STATE_UNSPECIFIED RunAssetDiscoveryResponse_State = 0
	// Asset discovery run completed successfully.
	RunAssetDiscoveryResponse_COMPLETED RunAssetDiscoveryResponse_State = 1
	// Asset discovery run was cancelled with tasks still pending, as another
	// run for the same organization was started with a higher priority.
	RunAssetDiscoveryResponse_SUPERSEDED RunAssetDiscoveryResponse_State = 2
	// Asset discovery run was killed and terminated.
	RunAssetDiscoveryResponse_TERMINATED RunAssetDiscoveryResponse_State = 3
)

var RunAssetDiscoveryResponse_State_name = map[int32]string{
	0: "STATE_UNSPECIFIED",
	1: "COMPLETED",
	2: "SUPERSEDED",
	3: "TERMINATED",
}

var RunAssetDiscoveryResponse_State_value = map[string]int32{
	"STATE_UNSPECIFIED": 0,
	"COMPLETED":         1,
	"SUPERSEDED":        2,
	"TERMINATED":        3,
}

func (x RunAssetDiscoveryResponse_State) String() string {
	return proto.EnumName(RunAssetDiscoveryResponse_State_name, int32(x))
}

func (RunAssetDiscoveryResponse_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e341f63197e359b4, []int{0, 0}
}

// Response of asset discovery run
type RunAssetDiscoveryResponse struct {
	// The state of an asset discovery run.
	State RunAssetDiscoveryResponse_State `protobuf:"varint,1,opt,name=state,proto3,enum=google.cloud.securitycenter.v1beta1.RunAssetDiscoveryResponse_State" json:"state,omitempty"`
	// The duration between asset discovery run start and end
	Duration             *duration.Duration `protobuf:"bytes,2,opt,name=duration,proto3" json:"duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RunAssetDiscoveryResponse) Reset()         { *m = RunAssetDiscoveryResponse{} }
func (m *RunAssetDiscoveryResponse) String() string { return proto.CompactTextString(m) }
func (*RunAssetDiscoveryResponse) ProtoMessage()    {}
func (*RunAssetDiscoveryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e341f63197e359b4, []int{0}
}

func (m *RunAssetDiscoveryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunAssetDiscoveryResponse.Unmarshal(m, b)
}
func (m *RunAssetDiscoveryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunAssetDiscoveryResponse.Marshal(b, m, deterministic)
}
func (m *RunAssetDiscoveryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunAssetDiscoveryResponse.Merge(m, src)
}
func (m *RunAssetDiscoveryResponse) XXX_Size() int {
	return xxx_messageInfo_RunAssetDiscoveryResponse.Size(m)
}
func (m *RunAssetDiscoveryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RunAssetDiscoveryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RunAssetDiscoveryResponse proto.InternalMessageInfo

func (m *RunAssetDiscoveryResponse) GetState() RunAssetDiscoveryResponse_State {
	if m != nil {
		return m.State
	}
	return RunAssetDiscoveryResponse_STATE_UNSPECIFIED
}

func (m *RunAssetDiscoveryResponse) GetDuration() *duration.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func init() {
	proto.RegisterEnum("google.cloud.securitycenter.v1beta1.RunAssetDiscoveryResponse_State", RunAssetDiscoveryResponse_State_name, RunAssetDiscoveryResponse_State_value)
	proto.RegisterType((*RunAssetDiscoveryResponse)(nil), "google.cloud.securitycenter.v1beta1.RunAssetDiscoveryResponse")
}

func init() {
	proto.RegisterFile("google/cloud/securitycenter/v1beta1/run_asset_discovery_response.proto", fileDescriptor_e341f63197e359b4)
}

var fileDescriptor_e341f63197e359b4 = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcf, 0x4b, 0xc3, 0x30,
	0x18, 0x86, 0xed, 0x64, 0xa2, 0x11, 0xc7, 0x2c, 0x08, 0xdb, 0x10, 0x19, 0xf3, 0xe0, 0x4e, 0x09,
	0x9b, 0x78, 0xf2, 0x34, 0x97, 0x0c, 0x06, 0x6e, 0xd6, 0xb6, 0xbb, 0xec, 0x52, 0xb2, 0x2e, 0x96,
	0x42, 0xcd, 0x57, 0x92, 0x74, 0xb0, 0x8b, 0xff, 0xb9, 0x20, 0xfd, 0x25, 0xec, 0xa0, 0xec, 0x98,
	0xbc, 0xef, 0xf7, 0xe4, 0xf9, 0x08, 0x9a, 0x45, 0x00, 0x51, 0x22, 0x48, 0x98, 0x40, 0xb6, 0x25,
	0x5a, 0x84, 0x99, 0x8a, 0xcd, 0x3e, 0x14, 0xd2, 0x08, 0x45, 0x76, 0xa3, 0x8d, 0x30, 0x7c, 0x44,
	0x54, 0x26, 0x03, 0xae, 0xb5, 0x30, 0xc1, 0x36, 0xd6, 0x21, 0xec, 0x84, 0xda, 0x07, 0x4a, 0xe8,
	0x14, 0xa4, 0x16, 0x38, 0x55, 0x60, 0xc0, 0xbe, 0x2f, 0x39, 0xb8, 0xe0, 0xe0, 0x43, 0x0e, 0xae,
	0x38, 0xbd, 0xbb, 0xea, 0xb1, 0x62, 0x64, 0x93, 0x7d, 0x90, 0x6d, 0xa6, 0xb8, 0x89, 0x41, 0x96,
	0x90, 0xde, 0x6d, 0x95, 0xf3, 0x34, 0x26, 0x5c, 0x4a, 0x30, 0x45, 0xa8, 0xcb, 0x74, 0xf0, 0x6d,
	0xa1, 0xae, 0x9b, 0xc9, 0x49, 0x2e, 0x42, 0x6b, 0x0f, 0xb7, 0xd2, 0xb0, 0xd7, 0xa8, 0xa9, 0x0d,
	0x37, 0xa2, 0x63, 0xf5, 0xad, 0x61, 0x6b, 0x4c, 0xf1, 0x11, 0x42, 0xf8, 0x4f, 0x1c, 0xf6, 0x72,
	0x96, 0x5b, 0x22, 0xed, 0x27, 0x74, 0x5e, 0x9b, 0x76, 0x1a, 0x7d, 0x6b, 0x78, 0x39, 0xee, 0xd6,
	0xf8, 0x7a, 0x15, 0x4c, 0xab, 0x82, 0xfb, 0x5b, 0x1d, 0x2c, 0x50, 0xb3, 0xc0, 0xd8, 0x37, 0xe8,
	0xda, 0xf3, 0x27, 0x3e, 0x0b, 0x56, 0x4b, 0xcf, 0x61, 0xd3, 0xf9, 0x6c, 0xce, 0x68, 0xfb, 0xc4,
	0xbe, 0x42, 0x17, 0xd3, 0xb7, 0x85, 0xf3, 0xca, 0x7c, 0x46, 0xdb, 0x96, 0xdd, 0x42, 0xc8, 0x5b,
	0x39, 0xcc, 0xf5, 0x18, 0x65, 0xb4, 0xdd, 0xc8, 0xcf, 0x3e, 0x73, 0x17, 0xf3, 0xe5, 0x24, 0xcf,
	0x4f, 0x5f, 0xbe, 0xd0, 0x43, 0x08, 0x9f, 0xc7, 0xec, 0xe5, 0x58, 0xeb, 0xf7, 0xaa, 0x16, 0x41,
	0xc2, 0x65, 0x84, 0x41, 0x45, 0x24, 0x12, 0xb2, 0xb0, 0x25, 0x65, 0xc4, 0xd3, 0x58, 0xff, 0xfb,
	0xed, 0xcf, 0x87, 0xd7, 0x9b, 0xb3, 0x62, 0xfa, 0xf1, 0x27, 0x00, 0x00, 0xff, 0xff, 0x33, 0x4f,
	0x03, 0xeb, 0x33, 0x02, 0x00, 0x00,
}
