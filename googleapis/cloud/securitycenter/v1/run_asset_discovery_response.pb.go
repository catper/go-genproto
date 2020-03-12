// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/securitycenter/v1/run_asset_discovery_response.proto

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
	return fileDescriptor_421ea6ed240ae1f7, []int{0, 0}
}

// Response of asset discovery run
type RunAssetDiscoveryResponse struct {
	// The state of an asset discovery run.
	State RunAssetDiscoveryResponse_State `protobuf:"varint,1,opt,name=state,proto3,enum=google.cloud.securitycenter.v1.RunAssetDiscoveryResponse_State" json:"state,omitempty"`
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
	return fileDescriptor_421ea6ed240ae1f7, []int{0}
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
	proto.RegisterEnum("google.cloud.securitycenter.v1.RunAssetDiscoveryResponse_State", RunAssetDiscoveryResponse_State_name, RunAssetDiscoveryResponse_State_value)
	proto.RegisterType((*RunAssetDiscoveryResponse)(nil), "google.cloud.securitycenter.v1.RunAssetDiscoveryResponse")
}

func init() {
	proto.RegisterFile("google/cloud/securitycenter/v1/run_asset_discovery_response.proto", fileDescriptor_421ea6ed240ae1f7)
}

var fileDescriptor_421ea6ed240ae1f7 = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x8a, 0x9b, 0x40,
	0x18, 0x80, 0xab, 0x25, 0xa5, 0x9d, 0xd2, 0x90, 0x0a, 0x85, 0x24, 0x94, 0x90, 0xe6, 0x94, 0xd3,
	0x0c, 0xa6, 0xf4, 0x62, 0x0f, 0xc5, 0xea, 0xb4, 0x04, 0x92, 0x54, 0xd4, 0xe4, 0x50, 0x02, 0x62,
	0xcc, 0x54, 0x84, 0x74, 0x46, 0x66, 0xc6, 0x40, 0x5e, 0x69, 0x1f, 0x65, 0x1f, 0x61, 0x1f, 0x61,
	0x9f, 0x60, 0x8f, 0x8b, 0x33, 0xba, 0xe0, 0x21, 0xbb, 0xc7, 0xf1, 0xff, 0xe6, 0xfb, 0x3f, 0x19,
	0xe0, 0xe6, 0x8c, 0xe5, 0x27, 0x82, 0xb2, 0x13, 0xab, 0x8e, 0x48, 0x90, 0xac, 0xe2, 0x85, 0xbc,
	0x64, 0x84, 0x4a, 0xc2, 0xd1, 0xd9, 0x46, 0xbc, 0xa2, 0x49, 0x2a, 0x04, 0x91, 0xc9, 0xb1, 0x10,
	0x19, 0x3b, 0x13, 0x7e, 0x49, 0x38, 0x11, 0x25, 0xa3, 0x82, 0xc0, 0x92, 0x33, 0xc9, 0xac, 0x89,
	0x56, 0x40, 0xa5, 0x80, 0x5d, 0x05, 0x3c, 0xdb, 0xe3, 0xcf, 0xcd, 0x8a, 0xb4, 0x2c, 0x50, 0x4a,
	0x29, 0x93, 0xa9, 0x2c, 0x18, 0x15, 0xfa, 0xf6, 0xb8, 0xb9, 0x8d, 0xd4, 0xe9, 0x50, 0xfd, 0x43,
	0xc7, 0x8a, 0x2b, 0x40, 0xcf, 0x67, 0x0f, 0x06, 0x18, 0x85, 0x15, 0x75, 0xeb, 0x06, 0xbf, 0x4d,
	0x08, 0x9b, 0x02, 0x6b, 0x0b, 0x7a, 0x42, 0xa6, 0x92, 0x0c, 0x8d, 0xa9, 0x31, 0xef, 0x2f, 0x7e,
	0xc0, 0xe7, 0x5b, 0xe0, 0x55, 0x13, 0x8c, 0x6a, 0x4d, 0xa8, 0x6d, 0xd6, 0x37, 0xf0, 0xb6, 0xcd,
	0x18, 0x9a, 0x53, 0x63, 0xfe, 0x7e, 0x31, 0x6a, 0xcd, 0x6d, 0x27, 0xf4, 0x1b, 0x20, 0x7c, 0x42,
	0x67, 0x6b, 0xd0, 0x53, 0x1a, 0xeb, 0x13, 0xf8, 0x18, 0xc5, 0x6e, 0x8c, 0x93, 0xed, 0x26, 0x0a,
	0xb0, 0xb7, 0xfc, 0xb5, 0xc4, 0xfe, 0xe0, 0x95, 0xf5, 0x01, 0xbc, 0xf3, 0xfe, 0xac, 0x83, 0x15,
	0x8e, 0xb1, 0x3f, 0x30, 0xac, 0x3e, 0x00, 0xd1, 0x36, 0xc0, 0x61, 0x84, 0x7d, 0xec, 0x0f, 0xcc,
	0xfa, 0x1c, 0xe3, 0x70, 0xbd, 0xdc, 0xb8, 0xf5, 0xfc, 0xf5, 0xcf, 0x3b, 0x03, 0xcc, 0x32, 0xf6,
	0xff, 0x85, 0x7f, 0x0a, 0x8c, 0xbf, 0xab, 0x86, 0xc8, 0xd9, 0x29, 0xa5, 0x39, 0x64, 0x3c, 0x47,
	0x39, 0xa1, 0xaa, 0x14, 0xe9, 0x51, 0x5a, 0x16, 0xe2, 0xda, 0x1b, 0x7f, 0xef, 0x7e, 0xb9, 0x31,
	0x27, 0xbf, 0xb5, 0xce, 0x53, 0x0b, 0xa3, 0x66, 0xea, 0xe9, 0x85, 0x3b, 0xfb, 0xb6, 0x05, 0xf6,
	0x0a, 0xd8, 0x77, 0x81, 0xfd, 0xce, 0xbe, 0x37, 0xbf, 0x68, 0xc0, 0x71, 0x14, 0xe1, 0x38, 0x5d,
	0xc4, 0x71, 0x76, 0xf6, 0xe1, 0x8d, 0xca, 0xfb, 0xfa, 0x18, 0x00, 0x00, 0xff, 0xff, 0x8e, 0xd3,
	0x2e, 0xc5, 0x81, 0x02, 0x00, 0x00,
}
