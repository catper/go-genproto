// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/maps/playablelocations/v3/playablelocations.proto

package playablelocations

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
	duration "github.com/catper/protobuf/ptypes/duration"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	sample "google.golang.org/genproto/googleapis/maps/playablelocations/v3/sample"
	unity "google.golang.org/genproto/googleapis/maps/unity"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

//
// Life of a query:
//
// - When a game starts in a new location, your game server issues a
// [SamplePlayableLocations][google.maps.playablelocations.v3.PlayableLocations.SamplePlayableLocations]
// request. The request specifies the S2 cell, and contains one or more
// "criteria" for filtering:
//
// - Criterion 0: i locations for long-lived bases, or level 0 monsters, or...
// - Criterion 1: j locations for short-lived bases, or level 1 monsters, ...
// - Criterion 2: k locations for random objects.
// - etc (up to 5 criterion may be specified).
//
// `PlayableLocationList` will then contain mutually
// exclusive lists of `PlayableLocation` objects that satisfy each of
// the criteria. Think of it as a collection of real-world locations that you
// can then associate with your game state.
//
// Note: These points are impermanent in nature. E.g, parks can close, and
// places can be removed.
//
// The response specifies how long you can expect the playable locations to
// last. Once they expire, you should query the `samplePlayableLocations` API
// again to get a fresh view of the real world.
type SamplePlayableLocationsRequest struct {
	// Required. Specifies the area to search within for playable locations.
	AreaFilter *sample.AreaFilter `protobuf:"bytes,1,opt,name=area_filter,json=areaFilter,proto3" json:"area_filter,omitempty"`
	// Required. Specifies one or more (up to 5) criteria for filtering the
	// returned playable locations.
	Criteria             []*sample.Criterion `protobuf:"bytes,2,rep,name=criteria,proto3" json:"criteria,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *SamplePlayableLocationsRequest) Reset()         { *m = SamplePlayableLocationsRequest{} }
func (m *SamplePlayableLocationsRequest) String() string { return proto.CompactTextString(m) }
func (*SamplePlayableLocationsRequest) ProtoMessage()    {}
func (*SamplePlayableLocationsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0280e761c74de946, []int{0}
}

func (m *SamplePlayableLocationsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SamplePlayableLocationsRequest.Unmarshal(m, b)
}
func (m *SamplePlayableLocationsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SamplePlayableLocationsRequest.Marshal(b, m, deterministic)
}
func (m *SamplePlayableLocationsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SamplePlayableLocationsRequest.Merge(m, src)
}
func (m *SamplePlayableLocationsRequest) XXX_Size() int {
	return xxx_messageInfo_SamplePlayableLocationsRequest.Size(m)
}
func (m *SamplePlayableLocationsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SamplePlayableLocationsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SamplePlayableLocationsRequest proto.InternalMessageInfo

func (m *SamplePlayableLocationsRequest) GetAreaFilter() *sample.AreaFilter {
	if m != nil {
		return m.AreaFilter
	}
	return nil
}

func (m *SamplePlayableLocationsRequest) GetCriteria() []*sample.Criterion {
	if m != nil {
		return m.Criteria
	}
	return nil
}

//
// Response for the
// [SamplePlayableLocations][google.maps.playablelocations.v3.PlayableLocations.SamplePlayableLocations]
// method.
type SamplePlayableLocationsResponse struct {
	// Each PlayableLocation object corresponds to a game_object_type specified
	// in the request.
	LocationsPerGameObjectType map[int32]*sample.PlayableLocationList `protobuf:"bytes,1,rep,name=locations_per_game_object_type,json=locationsPerGameObjectType,proto3" json:"locations_per_game_object_type,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Required. Specifies the "time-to-live" for the set of playable locations. You can use
	// this value to determine how long to cache the set of playable locations.
	// After this length of time, your back-end game server should issue a new
	// [SamplePlayableLocations][google.maps.playablelocations.v3.PlayableLocations.SamplePlayableLocations]
	// request to get a fresh set of playable locations (because for example, they
	// might have been removed, a park might have closed for the day, a
	// business might have closed permanently).
	Ttl                  *duration.Duration `protobuf:"bytes,9,opt,name=ttl,proto3" json:"ttl,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *SamplePlayableLocationsResponse) Reset()         { *m = SamplePlayableLocationsResponse{} }
func (m *SamplePlayableLocationsResponse) String() string { return proto.CompactTextString(m) }
func (*SamplePlayableLocationsResponse) ProtoMessage()    {}
func (*SamplePlayableLocationsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0280e761c74de946, []int{1}
}

func (m *SamplePlayableLocationsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SamplePlayableLocationsResponse.Unmarshal(m, b)
}
func (m *SamplePlayableLocationsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SamplePlayableLocationsResponse.Marshal(b, m, deterministic)
}
func (m *SamplePlayableLocationsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SamplePlayableLocationsResponse.Merge(m, src)
}
func (m *SamplePlayableLocationsResponse) XXX_Size() int {
	return xxx_messageInfo_SamplePlayableLocationsResponse.Size(m)
}
func (m *SamplePlayableLocationsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SamplePlayableLocationsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SamplePlayableLocationsResponse proto.InternalMessageInfo

func (m *SamplePlayableLocationsResponse) GetLocationsPerGameObjectType() map[int32]*sample.PlayableLocationList {
	if m != nil {
		return m.LocationsPerGameObjectType
	}
	return nil
}

func (m *SamplePlayableLocationsResponse) GetTtl() *duration.Duration {
	if m != nil {
		return m.Ttl
	}
	return nil
}

// A request for logging your player's bad location reports.
type LogPlayerReportsRequest struct {
	// Required. Player reports. The maximum number of player reports that you can log at
	// once is 50.
	PlayerReports []*PlayerReport `protobuf:"bytes,1,rep,name=player_reports,json=playerReports,proto3" json:"player_reports,omitempty"`
	// Required. A string that uniquely identifies the log player reports request. This
	// allows you to detect duplicate requests. We recommend that you use UUIDs
	// for this value. The value must not exceed 50 characters.
	//
	// You should reuse the `request_id` only when retrying a request in the case
	// of a failure. In that case, the request must be identical to the one that
	// failed.
	RequestId string `protobuf:"bytes,2,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// Required. Information about the client device (for example, device model and
	// operating system).
	ClientInfo           *unity.ClientInfo `protobuf:"bytes,3,opt,name=client_info,json=clientInfo,proto3" json:"client_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *LogPlayerReportsRequest) Reset()         { *m = LogPlayerReportsRequest{} }
func (m *LogPlayerReportsRequest) String() string { return proto.CompactTextString(m) }
func (*LogPlayerReportsRequest) ProtoMessage()    {}
func (*LogPlayerReportsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0280e761c74de946, []int{2}
}

func (m *LogPlayerReportsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogPlayerReportsRequest.Unmarshal(m, b)
}
func (m *LogPlayerReportsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogPlayerReportsRequest.Marshal(b, m, deterministic)
}
func (m *LogPlayerReportsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogPlayerReportsRequest.Merge(m, src)
}
func (m *LogPlayerReportsRequest) XXX_Size() int {
	return xxx_messageInfo_LogPlayerReportsRequest.Size(m)
}
func (m *LogPlayerReportsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogPlayerReportsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogPlayerReportsRequest proto.InternalMessageInfo

func (m *LogPlayerReportsRequest) GetPlayerReports() []*PlayerReport {
	if m != nil {
		return m.PlayerReports
	}
	return nil
}

func (m *LogPlayerReportsRequest) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *LogPlayerReportsRequest) GetClientInfo() *unity.ClientInfo {
	if m != nil {
		return m.ClientInfo
	}
	return nil
}

// A response for the [LogPlayerReports][google.maps.playablelocations.v3.PlayableLocations.LogPlayerReports]
// method.
//
// This method returns no data upon success.
type LogPlayerReportsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogPlayerReportsResponse) Reset()         { *m = LogPlayerReportsResponse{} }
func (m *LogPlayerReportsResponse) String() string { return proto.CompactTextString(m) }
func (*LogPlayerReportsResponse) ProtoMessage()    {}
func (*LogPlayerReportsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0280e761c74de946, []int{3}
}

func (m *LogPlayerReportsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogPlayerReportsResponse.Unmarshal(m, b)
}
func (m *LogPlayerReportsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogPlayerReportsResponse.Marshal(b, m, deterministic)
}
func (m *LogPlayerReportsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogPlayerReportsResponse.Merge(m, src)
}
func (m *LogPlayerReportsResponse) XXX_Size() int {
	return xxx_messageInfo_LogPlayerReportsResponse.Size(m)
}
func (m *LogPlayerReportsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LogPlayerReportsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LogPlayerReportsResponse proto.InternalMessageInfo

// A request for logging impressions.
type LogImpressionsRequest struct {
	// Required. Impression event details. The maximum number of impression reports that you
	// can log at once is 50.
	Impressions []*Impression `protobuf:"bytes,1,rep,name=impressions,proto3" json:"impressions,omitempty"`
	// Required. A string that uniquely identifies the log impressions request. This allows
	// you to detect duplicate requests. We recommend that you use UUIDs for this
	// value. The value must not exceed 50 characters.
	//
	// You should reuse the `request_id` only when retrying a request in case of
	// failure. In this case, the request must be identical to the one that
	// failed.
	RequestId string `protobuf:"bytes,2,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// Required. Information about the client device. For example, device model and
	// operating system.
	ClientInfo           *unity.ClientInfo `protobuf:"bytes,3,opt,name=client_info,json=clientInfo,proto3" json:"client_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *LogImpressionsRequest) Reset()         { *m = LogImpressionsRequest{} }
func (m *LogImpressionsRequest) String() string { return proto.CompactTextString(m) }
func (*LogImpressionsRequest) ProtoMessage()    {}
func (*LogImpressionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0280e761c74de946, []int{4}
}

func (m *LogImpressionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogImpressionsRequest.Unmarshal(m, b)
}
func (m *LogImpressionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogImpressionsRequest.Marshal(b, m, deterministic)
}
func (m *LogImpressionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogImpressionsRequest.Merge(m, src)
}
func (m *LogImpressionsRequest) XXX_Size() int {
	return xxx_messageInfo_LogImpressionsRequest.Size(m)
}
func (m *LogImpressionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogImpressionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogImpressionsRequest proto.InternalMessageInfo

func (m *LogImpressionsRequest) GetImpressions() []*Impression {
	if m != nil {
		return m.Impressions
	}
	return nil
}

func (m *LogImpressionsRequest) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *LogImpressionsRequest) GetClientInfo() *unity.ClientInfo {
	if m != nil {
		return m.ClientInfo
	}
	return nil
}

// A response for the [LogImpressions][google.maps.playablelocations.v3.PlayableLocations.LogImpressions] method.
// This method returns no data upon success.
type LogImpressionsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogImpressionsResponse) Reset()         { *m = LogImpressionsResponse{} }
func (m *LogImpressionsResponse) String() string { return proto.CompactTextString(m) }
func (*LogImpressionsResponse) ProtoMessage()    {}
func (*LogImpressionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0280e761c74de946, []int{5}
}

func (m *LogImpressionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogImpressionsResponse.Unmarshal(m, b)
}
func (m *LogImpressionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogImpressionsResponse.Marshal(b, m, deterministic)
}
func (m *LogImpressionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogImpressionsResponse.Merge(m, src)
}
func (m *LogImpressionsResponse) XXX_Size() int {
	return xxx_messageInfo_LogImpressionsResponse.Size(m)
}
func (m *LogImpressionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LogImpressionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LogImpressionsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SamplePlayableLocationsRequest)(nil), "google.maps.playablelocations.v3.SamplePlayableLocationsRequest")
	proto.RegisterType((*SamplePlayableLocationsResponse)(nil), "google.maps.playablelocations.v3.SamplePlayableLocationsResponse")
	proto.RegisterMapType((map[int32]*sample.PlayableLocationList)(nil), "google.maps.playablelocations.v3.SamplePlayableLocationsResponse.LocationsPerGameObjectTypeEntry")
	proto.RegisterType((*LogPlayerReportsRequest)(nil), "google.maps.playablelocations.v3.LogPlayerReportsRequest")
	proto.RegisterType((*LogPlayerReportsResponse)(nil), "google.maps.playablelocations.v3.LogPlayerReportsResponse")
	proto.RegisterType((*LogImpressionsRequest)(nil), "google.maps.playablelocations.v3.LogImpressionsRequest")
	proto.RegisterType((*LogImpressionsResponse)(nil), "google.maps.playablelocations.v3.LogImpressionsResponse")
}

func init() {
	proto.RegisterFile("google/maps/playablelocations/v3/playablelocations.proto", fileDescriptor_0280e761c74de946)
}

var fileDescriptor_0280e761c74de946 = []byte{
	// 748 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0x4f, 0x6b, 0x13, 0x41,
	0x14, 0x67, 0x92, 0x56, 0xec, 0x04, 0x4b, 0x1d, 0xb4, 0xd9, 0xae, 0xb6, 0x09, 0xab, 0x48, 0xa9,
	0xb2, 0x2b, 0xc9, 0xa1, 0x35, 0x22, 0x98, 0x56, 0x2d, 0x85, 0x14, 0x63, 0x22, 0x82, 0xbd, 0x2c,
	0x93, 0xcd, 0x64, 0x5d, 0xdd, 0xec, 0xac, 0xb3, 0x93, 0x40, 0xae, 0x82, 0x37, 0x6f, 0x7e, 0x83,
	0x82, 0x97, 0x7e, 0x07, 0xaf, 0x1e, 0x3c, 0x09, 0xc5, 0xab, 0x78, 0xf0, 0x83, 0xc8, 0xce, 0x4c,
	0xd2, 0x6d, 0xd6, 0xb0, 0xd5, 0x82, 0xb7, 0xf6, 0xcd, 0xfb, 0xfd, 0x9b, 0x37, 0x6f, 0x03, 0xb7,
	0x5c, 0x4a, 0x5d, 0x9f, 0x58, 0x7d, 0x1c, 0x46, 0x56, 0xe8, 0xe3, 0x11, 0xee, 0xf8, 0xc4, 0xa7,
	0x0e, 0xe6, 0x1e, 0x0d, 0x22, 0x6b, 0x58, 0x4d, 0x17, 0xcd, 0x90, 0x51, 0x4e, 0x51, 0x59, 0x22,
	0xcd, 0x18, 0x69, 0xa6, 0x9b, 0x86, 0x55, 0xfd, 0xba, 0xe2, 0xc6, 0xa1, 0x67, 0xe1, 0x20, 0xa0,
	0x3c, 0x89, 0xd7, 0x4b, 0x89, 0xd3, 0x9e, 0x47, 0xfc, 0xae, 0xdd, 0x21, 0xaf, 0xf0, 0xd0, 0xa3,
	0x4c, 0x35, 0xdc, 0xcd, 0xb4, 0xc6, 0x48, 0x44, 0x07, 0xcc, 0x21, 0x63, 0xca, 0xcd, 0x4c, 0x44,
	0x84, 0xfb, 0xa1, 0x4f, 0x52, 0x40, 0x23, 0x09, 0x1c, 0x04, 0x1e, 0x1f, 0x59, 0x8e, 0xef, 0x91,
	0x80, 0x7b, 0x41, 0x8f, 0xaa, 0x9e, 0x35, 0xd5, 0x23, 0xfe, 0xeb, 0x0c, 0x7a, 0x56, 0x77, 0xc0,
	0x04, 0xb1, 0x3a, 0x2f, 0x26, 0xf2, 0x48, 0xb0, 0x3c, 0x30, 0xbe, 0x01, 0xb8, 0xd6, 0x16, 0xba,
	0x4d, 0x65, 0xa9, 0x31, 0xb6, 0xd4, 0x22, 0x6f, 0x07, 0x24, 0xe2, 0xe8, 0x25, 0x2c, 0x60, 0x46,
	0xb0, 0xdd, 0xf3, 0x7c, 0x4e, 0x98, 0x06, 0xca, 0x60, 0xbd, 0x50, 0xa9, 0x9a, 0x59, 0x37, 0x6c,
	0xca, 0x38, 0x66, 0x9d, 0x11, 0xfc, 0x44, 0x40, 0xb7, 0xf3, 0x3f, 0xeb, 0xb9, 0x16, 0xc4, 0x93,
	0x02, 0x6a, 0xc3, 0x8b, 0x0e, 0xf3, 0x38, 0x61, 0x1e, 0xd6, 0x72, 0xe5, 0xfc, 0x7a, 0xa1, 0x52,
	0x39, 0x33, 0xef, 0x8e, 0x04, 0xd2, 0x40, 0xd2, 0x4e, 0x88, 0x8c, 0xf7, 0x79, 0x58, 0x9a, 0x19,
	0x29, 0x0a, 0x69, 0x10, 0x11, 0xf4, 0x09, 0xc0, 0xb5, 0x09, 0xa9, 0x1d, 0x12, 0x66, 0xbb, 0xb8,
	0x4f, 0x6c, 0xda, 0x79, 0x4d, 0x1c, 0x6e, 0xf3, 0x51, 0x48, 0x34, 0x20, 0xfc, 0xe0, 0x6c, 0x3f,
	0x19, 0x5a, 0xe6, 0xa4, 0xd2, 0x24, 0x6c, 0x17, 0xf7, 0xc9, 0x53, 0x21, 0xf2, 0x7c, 0x14, 0x92,
	0xc7, 0x01, 0x67, 0xa3, 0x96, 0xee, 0xcf, 0x6c, 0x40, 0xb7, 0x61, 0x9e, 0x73, 0x5f, 0x5b, 0x10,
	0x77, 0xbe, 0x32, 0xf6, 0x32, 0x9e, 0xb2, 0xf9, 0x48, 0x4d, 0xb9, 0x15, 0x77, 0xe9, 0x1f, 0x00,
	0x2c, 0x65, 0x88, 0xa1, 0x25, 0x98, 0x7f, 0x43, 0x46, 0x62, 0x88, 0xf3, 0xad, 0xf8, 0x4f, 0xd4,
	0x86, 0xf3, 0x43, 0xec, 0x0f, 0x88, 0x96, 0x13, 0x22, 0x0f, 0xce, 0x3c, 0x80, 0xe9, 0xc4, 0x0d,
	0x2f, 0xe2, 0x2d, 0xc9, 0x55, 0xcb, 0x6d, 0x01, 0xe3, 0x07, 0x80, 0xc5, 0x06, 0x75, 0xe3, 0x36,
	0xc2, 0x5a, 0x24, 0xa4, 0x8c, 0x4f, 0xde, 0xd4, 0x01, 0x5c, 0x0c, 0x45, 0xdd, 0x66, 0xf2, 0x40,
	0x5d, 0xb7, 0x99, 0xad, 0x9e, 0xe4, 0x93, 0xa3, 0xbf, 0x14, 0x26, 0x25, 0x90, 0x01, 0x21, 0x93,
	0x32, 0xb6, 0xd7, 0x15, 0xa9, 0x16, 0x64, 0xdf, 0x82, 0x2a, 0xef, 0x75, 0xd1, 0x0e, 0x2c, 0xc8,
	0x35, 0xb0, 0xe3, 0x25, 0xd2, 0xf2, 0x22, 0xfa, 0xea, 0x29, 0x71, 0xb1, 0x69, 0xe6, 0x8e, 0xe8,
	0xda, 0x0b, 0x7a, 0x54, 0xbd, 0x5e, 0x67, 0x52, 0x30, 0x74, 0xa8, 0xa5, 0xf3, 0xc9, 0xa1, 0x1b,
	0xdf, 0x01, 0xbc, 0xda, 0xa0, 0xee, 0x5e, 0x3f, 0x64, 0x24, 0x8a, 0x92, 0xeb, 0xd4, 0x86, 0x05,
	0xef, 0xa4, 0xaa, 0x72, 0xdf, 0xc9, 0xce, 0x7d, 0x42, 0x25, 0x9d, 0x24, 0x59, 0xfe, 0x5f, 0x66,
	0x0d, 0x2e, 0x4f, 0xc7, 0x92, 0x89, 0x2b, 0xc7, 0x73, 0xf0, 0x72, 0x6a, 0x09, 0xd0, 0x17, 0x00,
	0x8b, 0x33, 0x16, 0x04, 0x3d, 0x3c, 0xc7, 0x6e, 0x89, 0x50, 0x7a, 0xfd, 0xdc, 0xdb, 0x69, 0xdc,
	0x7a, 0x77, 0xfc, 0xeb, 0x63, 0xae, 0x6c, 0x5c, 0xb3, 0x86, 0xd5, 0x5a, 0xf4, 0xe7, 0xe6, 0x1a,
	0xd8, 0x40, 0x47, 0x00, 0x2e, 0x4d, 0x4f, 0x1b, 0xdd, 0xcb, 0xd6, 0x9f, 0xb1, 0x01, 0x7a, 0xed,
	0x5f, 0xa0, 0xca, 0x73, 0x49, 0x78, 0x5e, 0x31, 0xae, 0xc4, 0x9e, 0xfd, 0xa9, 0xae, 0xd8, 0xec,
	0x21, 0x80, 0x8b, 0xa7, 0xc7, 0x84, 0x36, 0xcf, 0xa4, 0x97, 0x7e, 0xaf, 0xfa, 0xd6, 0xdf, 0x03,
	0x95, 0xcd, 0x55, 0x61, 0xb3, 0x68, 0x20, 0x65, 0x33, 0xd1, 0x53, 0x03, 0x1b, 0xfa, 0x8d, 0xaf,
	0xf5, 0x72, 0x9a, 0x4e, 0x8a, 0xe1, 0xd0, 0x8b, 0x4c, 0x87, 0xf6, 0xb7, 0x3f, 0x03, 0x78, 0xd3,
	0xa1, 0xfd, 0x4c, 0x0f, 0xdb, 0xcb, 0xa9, 0xa9, 0x35, 0xe3, 0xaf, 0x64, 0x13, 0x1c, 0x3c, 0x53,
	0x58, 0x97, 0xfa, 0x38, 0x70, 0x4d, 0xca, 0x5c, 0xcb, 0x25, 0x81, 0xf8, 0x86, 0x5a, 0x27, 0x6a,
	0xb3, 0x7f, 0x97, 0xef, 0xa7, 0x8a, 0x87, 0xb9, 0xb9, 0xdd, 0xfd, 0x66, 0xe3, 0x28, 0x57, 0xde,
	0x95, 0xd4, 0xfb, 0xb1, 0xad, 0x94, 0x01, 0xf3, 0x45, 0xb5, 0x73, 0x41, 0xe8, 0x54, 0x7f, 0x07,
	0x00, 0x00, 0xff, 0xff, 0xed, 0xf1, 0x01, 0x5c, 0xcb, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PlayableLocationsClient is the client API for PlayableLocations service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PlayableLocationsClient interface {
	// Returns a set of playable locations that lie within a specified area,
	// that satisfy optional filter criteria.
	//
	// Note: Identical `SamplePlayableLocations` requests can return different
	// results as the state of the world changes over time.
	SamplePlayableLocations(ctx context.Context, in *SamplePlayableLocationsRequest, opts ...grpc.CallOption) (*SamplePlayableLocationsResponse, error)
	// Logs bad playable location reports submitted by players.
	//
	// Reports are not partially saved; either all reports are saved and this
	// request succeeds, or no reports are saved, and this request fails.
	LogPlayerReports(ctx context.Context, in *LogPlayerReportsRequest, opts ...grpc.CallOption) (*LogPlayerReportsResponse, error)
	// Logs new events when playable locations are displayed, and when they are
	// interacted with.
	//
	// Impressions are not partially saved; either all impressions are saved and
	// this request succeeds, or no impressions are saved, and this request fails.
	LogImpressions(ctx context.Context, in *LogImpressionsRequest, opts ...grpc.CallOption) (*LogImpressionsResponse, error)
}

type playableLocationsClient struct {
	cc grpc.ClientConnInterface
}

func NewPlayableLocationsClient(cc grpc.ClientConnInterface) PlayableLocationsClient {
	return &playableLocationsClient{cc}
}

func (c *playableLocationsClient) SamplePlayableLocations(ctx context.Context, in *SamplePlayableLocationsRequest, opts ...grpc.CallOption) (*SamplePlayableLocationsResponse, error) {
	out := new(SamplePlayableLocationsResponse)
	err := c.cc.Invoke(ctx, "/google.maps.playablelocations.v3.PlayableLocations/SamplePlayableLocations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playableLocationsClient) LogPlayerReports(ctx context.Context, in *LogPlayerReportsRequest, opts ...grpc.CallOption) (*LogPlayerReportsResponse, error) {
	out := new(LogPlayerReportsResponse)
	err := c.cc.Invoke(ctx, "/google.maps.playablelocations.v3.PlayableLocations/LogPlayerReports", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playableLocationsClient) LogImpressions(ctx context.Context, in *LogImpressionsRequest, opts ...grpc.CallOption) (*LogImpressionsResponse, error) {
	out := new(LogImpressionsResponse)
	err := c.cc.Invoke(ctx, "/google.maps.playablelocations.v3.PlayableLocations/LogImpressions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlayableLocationsServer is the server API for PlayableLocations service.
type PlayableLocationsServer interface {
	// Returns a set of playable locations that lie within a specified area,
	// that satisfy optional filter criteria.
	//
	// Note: Identical `SamplePlayableLocations` requests can return different
	// results as the state of the world changes over time.
	SamplePlayableLocations(context.Context, *SamplePlayableLocationsRequest) (*SamplePlayableLocationsResponse, error)
	// Logs bad playable location reports submitted by players.
	//
	// Reports are not partially saved; either all reports are saved and this
	// request succeeds, or no reports are saved, and this request fails.
	LogPlayerReports(context.Context, *LogPlayerReportsRequest) (*LogPlayerReportsResponse, error)
	// Logs new events when playable locations are displayed, and when they are
	// interacted with.
	//
	// Impressions are not partially saved; either all impressions are saved and
	// this request succeeds, or no impressions are saved, and this request fails.
	LogImpressions(context.Context, *LogImpressionsRequest) (*LogImpressionsResponse, error)
}

// UnimplementedPlayableLocationsServer can be embedded to have forward compatible implementations.
type UnimplementedPlayableLocationsServer struct {
}

func (*UnimplementedPlayableLocationsServer) SamplePlayableLocations(ctx context.Context, req *SamplePlayableLocationsRequest) (*SamplePlayableLocationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SamplePlayableLocations not implemented")
}
func (*UnimplementedPlayableLocationsServer) LogPlayerReports(ctx context.Context, req *LogPlayerReportsRequest) (*LogPlayerReportsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogPlayerReports not implemented")
}
func (*UnimplementedPlayableLocationsServer) LogImpressions(ctx context.Context, req *LogImpressionsRequest) (*LogImpressionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogImpressions not implemented")
}

func RegisterPlayableLocationsServer(s *grpc.Server, srv PlayableLocationsServer) {
	s.RegisterService(&_PlayableLocations_serviceDesc, srv)
}

func _PlayableLocations_SamplePlayableLocations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SamplePlayableLocationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayableLocationsServer).SamplePlayableLocations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.maps.playablelocations.v3.PlayableLocations/SamplePlayableLocations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayableLocationsServer).SamplePlayableLocations(ctx, req.(*SamplePlayableLocationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayableLocations_LogPlayerReports_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogPlayerReportsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayableLocationsServer).LogPlayerReports(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.maps.playablelocations.v3.PlayableLocations/LogPlayerReports",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayableLocationsServer).LogPlayerReports(ctx, req.(*LogPlayerReportsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayableLocations_LogImpressions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogImpressionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayableLocationsServer).LogImpressions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.maps.playablelocations.v3.PlayableLocations/LogImpressions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayableLocationsServer).LogImpressions(ctx, req.(*LogImpressionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PlayableLocations_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.maps.playablelocations.v3.PlayableLocations",
	HandlerType: (*PlayableLocationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SamplePlayableLocations",
			Handler:    _PlayableLocations_SamplePlayableLocations_Handler,
		},
		{
			MethodName: "LogPlayerReports",
			Handler:    _PlayableLocations_LogPlayerReports_Handler,
		},
		{
			MethodName: "LogImpressions",
			Handler:    _PlayableLocations_LogImpressions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/maps/playablelocations/v3/playablelocations.proto",
}
