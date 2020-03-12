// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/maps/roads/v1op/roads.proto

package roads

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
	wrappers "github.com/catper/protobuf/ptypes/wrappers"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	latlng "google.golang.org/genproto/googleapis/type/latlng"
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

// An enum representing the mode of travel used for snapping.
type TravelMode int32

const (
	TravelMode_TRAVEL_MODE_UNSPECIFIED TravelMode = 0
	TravelMode_DRIVING                 TravelMode = 1
	TravelMode_CYCLING                 TravelMode = 2
	TravelMode_WALKING                 TravelMode = 3
)

var TravelMode_name = map[int32]string{
	0: "TRAVEL_MODE_UNSPECIFIED",
	1: "DRIVING",
	2: "CYCLING",
	3: "WALKING",
}

var TravelMode_value = map[string]int32{
	"TRAVEL_MODE_UNSPECIFIED": 0,
	"DRIVING":                 1,
	"CYCLING":                 2,
	"WALKING":                 3,
}

func (x TravelMode) String() string {
	return proto.EnumName(TravelMode_name, int32(x))
}

func (TravelMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c8e70e24aafa772e, []int{0}
}

// A request to the SnapToRoads method, requesting that a sequence of points be
// snapped to road segments.
type SnapToRoadsRequest struct {
	// The path to be snapped as a series of lat, lng points. Specified as
	// a string of the format: lat,lng|lat,lng|...
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// Whether to interpolate the points to return full road geometry.
	Interpolate bool `protobuf:"varint,2,opt,name=interpolate,proto3" json:"interpolate,omitempty"`
	// The asset ID of the asset to which this path relates. This is used for
	// abuse detection purposes for clients with asset-based SKUs.
	AssetId string `protobuf:"bytes,3,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	// The type of travel being tracked. This will constrain the paths we snap to.
	TravelMode           TravelMode `protobuf:"varint,4,opt,name=travel_mode,json=travelMode,proto3,enum=google.maps.roads.v1op.TravelMode" json:"travel_mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SnapToRoadsRequest) Reset()         { *m = SnapToRoadsRequest{} }
func (m *SnapToRoadsRequest) String() string { return proto.CompactTextString(m) }
func (*SnapToRoadsRequest) ProtoMessage()    {}
func (*SnapToRoadsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8e70e24aafa772e, []int{0}
}

func (m *SnapToRoadsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnapToRoadsRequest.Unmarshal(m, b)
}
func (m *SnapToRoadsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnapToRoadsRequest.Marshal(b, m, deterministic)
}
func (m *SnapToRoadsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnapToRoadsRequest.Merge(m, src)
}
func (m *SnapToRoadsRequest) XXX_Size() int {
	return xxx_messageInfo_SnapToRoadsRequest.Size(m)
}
func (m *SnapToRoadsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SnapToRoadsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SnapToRoadsRequest proto.InternalMessageInfo

func (m *SnapToRoadsRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *SnapToRoadsRequest) GetInterpolate() bool {
	if m != nil {
		return m.Interpolate
	}
	return false
}

func (m *SnapToRoadsRequest) GetAssetId() string {
	if m != nil {
		return m.AssetId
	}
	return ""
}

func (m *SnapToRoadsRequest) GetTravelMode() TravelMode {
	if m != nil {
		return m.TravelMode
	}
	return TravelMode_TRAVEL_MODE_UNSPECIFIED
}

// A snapped point object, representing the result of snapping.
type SnappedPoint struct {
	// The lat,lng of the snapped location.
	Location *latlng.LatLng `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	// The index into the original path of the equivalent pre-snapped point.
	// This allows for identification of points which have been interpolated if
	// this index is missing.
	OriginalIndex *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=original_index,json=originalIndex,proto3" json:"original_index,omitempty"`
	// The place ID for this snapped location (road segment). These are the same
	// as are currently used by the Places API.
	PlaceId              string   `protobuf:"bytes,3,opt,name=place_id,json=placeId,proto3" json:"place_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SnappedPoint) Reset()         { *m = SnappedPoint{} }
func (m *SnappedPoint) String() string { return proto.CompactTextString(m) }
func (*SnappedPoint) ProtoMessage()    {}
func (*SnappedPoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8e70e24aafa772e, []int{1}
}

func (m *SnappedPoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnappedPoint.Unmarshal(m, b)
}
func (m *SnappedPoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnappedPoint.Marshal(b, m, deterministic)
}
func (m *SnappedPoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnappedPoint.Merge(m, src)
}
func (m *SnappedPoint) XXX_Size() int {
	return xxx_messageInfo_SnappedPoint.Size(m)
}
func (m *SnappedPoint) XXX_DiscardUnknown() {
	xxx_messageInfo_SnappedPoint.DiscardUnknown(m)
}

var xxx_messageInfo_SnappedPoint proto.InternalMessageInfo

func (m *SnappedPoint) GetLocation() *latlng.LatLng {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *SnappedPoint) GetOriginalIndex() *wrappers.UInt32Value {
	if m != nil {
		return m.OriginalIndex
	}
	return nil
}

func (m *SnappedPoint) GetPlaceId() string {
	if m != nil {
		return m.PlaceId
	}
	return ""
}

// The response from the SnapToRoads method, returning a sequence of snapped
// points.
type SnapToRoadsResponse struct {
	// A list of snapped points.
	SnappedPoints []*SnappedPoint `protobuf:"bytes,1,rep,name=snapped_points,json=snappedPoints,proto3" json:"snapped_points,omitempty"`
	// User-visible warning message, if any, which can be shown alongside a valid
	// result.
	WarningMessage       string   `protobuf:"bytes,2,opt,name=warning_message,json=warningMessage,proto3" json:"warning_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SnapToRoadsResponse) Reset()         { *m = SnapToRoadsResponse{} }
func (m *SnapToRoadsResponse) String() string { return proto.CompactTextString(m) }
func (*SnapToRoadsResponse) ProtoMessage()    {}
func (*SnapToRoadsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8e70e24aafa772e, []int{2}
}

func (m *SnapToRoadsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnapToRoadsResponse.Unmarshal(m, b)
}
func (m *SnapToRoadsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnapToRoadsResponse.Marshal(b, m, deterministic)
}
func (m *SnapToRoadsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnapToRoadsResponse.Merge(m, src)
}
func (m *SnapToRoadsResponse) XXX_Size() int {
	return xxx_messageInfo_SnapToRoadsResponse.Size(m)
}
func (m *SnapToRoadsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SnapToRoadsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SnapToRoadsResponse proto.InternalMessageInfo

func (m *SnapToRoadsResponse) GetSnappedPoints() []*SnappedPoint {
	if m != nil {
		return m.SnappedPoints
	}
	return nil
}

func (m *SnapToRoadsResponse) GetWarningMessage() string {
	if m != nil {
		return m.WarningMessage
	}
	return ""
}

// A request to the ListNearestRoads method, requesting that a sequence of
// points be snapped individually to the road segment that each is closest to.
type ListNearestRoadsRequest struct {
	// The points to be snapped as a series of lat, lng points. Specified as
	// a string of the format: lat,lng|lat,lng|...
	Points string `protobuf:"bytes,1,opt,name=points,proto3" json:"points,omitempty"`
	// The type of travel being tracked. This will constrain the roads we snap to.
	TravelMode           TravelMode `protobuf:"varint,2,opt,name=travel_mode,json=travelMode,proto3,enum=google.maps.roads.v1op.TravelMode" json:"travel_mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListNearestRoadsRequest) Reset()         { *m = ListNearestRoadsRequest{} }
func (m *ListNearestRoadsRequest) String() string { return proto.CompactTextString(m) }
func (*ListNearestRoadsRequest) ProtoMessage()    {}
func (*ListNearestRoadsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8e70e24aafa772e, []int{3}
}

func (m *ListNearestRoadsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListNearestRoadsRequest.Unmarshal(m, b)
}
func (m *ListNearestRoadsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListNearestRoadsRequest.Marshal(b, m, deterministic)
}
func (m *ListNearestRoadsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListNearestRoadsRequest.Merge(m, src)
}
func (m *ListNearestRoadsRequest) XXX_Size() int {
	return xxx_messageInfo_ListNearestRoadsRequest.Size(m)
}
func (m *ListNearestRoadsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListNearestRoadsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListNearestRoadsRequest proto.InternalMessageInfo

func (m *ListNearestRoadsRequest) GetPoints() string {
	if m != nil {
		return m.Points
	}
	return ""
}

func (m *ListNearestRoadsRequest) GetTravelMode() TravelMode {
	if m != nil {
		return m.TravelMode
	}
	return TravelMode_TRAVEL_MODE_UNSPECIFIED
}

// The response from the ListNearestRoads method, returning a list of snapped
// points.
type ListNearestRoadsResponse struct {
	// A list of snapped points.
	SnappedPoints        []*SnappedPoint `protobuf:"bytes,1,rep,name=snapped_points,json=snappedPoints,proto3" json:"snapped_points,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ListNearestRoadsResponse) Reset()         { *m = ListNearestRoadsResponse{} }
func (m *ListNearestRoadsResponse) String() string { return proto.CompactTextString(m) }
func (*ListNearestRoadsResponse) ProtoMessage()    {}
func (*ListNearestRoadsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8e70e24aafa772e, []int{4}
}

func (m *ListNearestRoadsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListNearestRoadsResponse.Unmarshal(m, b)
}
func (m *ListNearestRoadsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListNearestRoadsResponse.Marshal(b, m, deterministic)
}
func (m *ListNearestRoadsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListNearestRoadsResponse.Merge(m, src)
}
func (m *ListNearestRoadsResponse) XXX_Size() int {
	return xxx_messageInfo_ListNearestRoadsResponse.Size(m)
}
func (m *ListNearestRoadsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListNearestRoadsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListNearestRoadsResponse proto.InternalMessageInfo

func (m *ListNearestRoadsResponse) GetSnappedPoints() []*SnappedPoint {
	if m != nil {
		return m.SnappedPoints
	}
	return nil
}

func init() {
	proto.RegisterEnum("google.maps.roads.v1op.TravelMode", TravelMode_name, TravelMode_value)
	proto.RegisterType((*SnapToRoadsRequest)(nil), "google.maps.roads.v1op.SnapToRoadsRequest")
	proto.RegisterType((*SnappedPoint)(nil), "google.maps.roads.v1op.SnappedPoint")
	proto.RegisterType((*SnapToRoadsResponse)(nil), "google.maps.roads.v1op.SnapToRoadsResponse")
	proto.RegisterType((*ListNearestRoadsRequest)(nil), "google.maps.roads.v1op.ListNearestRoadsRequest")
	proto.RegisterType((*ListNearestRoadsResponse)(nil), "google.maps.roads.v1op.ListNearestRoadsResponse")
}

func init() {
	proto.RegisterFile("google/maps/roads/v1op/roads.proto", fileDescriptor_c8e70e24aafa772e)
}

var fileDescriptor_c8e70e24aafa772e = []byte{
	// 609 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x5f, 0x6f, 0xd3, 0x3e,
	0x14, 0x5d, 0xba, 0x69, 0x7f, 0x6e, 0xb6, 0xfe, 0x26, 0xef, 0xa7, 0x35, 0x94, 0x09, 0x55, 0x11,
	0x12, 0xd5, 0x90, 0x12, 0xe8, 0x1e, 0xf7, 0xb4, 0x75, 0x05, 0x45, 0xeb, 0x46, 0x95, 0xfd, 0x41,
	0xf0, 0x12, 0x79, 0x8d, 0xf1, 0x2c, 0xa5, 0xb6, 0x89, 0xbd, 0x0e, 0xbe, 0x02, 0x7c, 0x08, 0xde,
	0xf9, 0x28, 0x7c, 0x2a, 0x64, 0x27, 0xd9, 0xb2, 0x41, 0xd1, 0x84, 0xc4, 0x9b, 0x8f, 0xef, 0x39,
	0x37, 0xe7, 0x9e, 0x2b, 0x07, 0x7c, 0x2a, 0x04, 0xcd, 0x48, 0x38, 0xc1, 0x52, 0x85, 0xb9, 0xc0,
	0xa9, 0x0a, 0xa7, 0x2f, 0x85, 0x2c, 0x8e, 0x81, 0xcc, 0x85, 0x16, 0x68, 0xb3, 0xe0, 0x04, 0x86,
	0x13, 0x14, 0x05, 0xc3, 0x69, 0x6f, 0x95, 0x5a, 0x2c, 0x59, 0x88, 0x39, 0x17, 0x1a, 0x6b, 0x26,
	0x78, 0xa9, 0x6a, 0xb7, 0x6a, 0xd5, 0x71, 0xc6, 0x08, 0xd7, 0x65, 0xe1, 0x49, 0x59, 0xb0, 0xe8,
	0xe2, 0xea, 0x43, 0x78, 0x9d, 0x63, 0x29, 0x49, 0x5e, 0x09, 0xbd, 0xb2, 0xae, 0x3f, 0x4b, 0x12,
	0x66, 0x58, 0x67, 0x9c, 0x16, 0x15, 0xff, 0xbb, 0x03, 0xe8, 0x84, 0x63, 0x79, 0x2a, 0x62, 0xe3,
	0x22, 0x26, 0x1f, 0xaf, 0x88, 0xd2, 0x08, 0xc1, 0x82, 0xc4, 0xfa, 0xd2, 0x73, 0x3a, 0x4e, 0x77,
	0x25, 0xb6, 0x67, 0xd4, 0x01, 0x97, 0x71, 0x4d, 0x72, 0x29, 0x32, 0xac, 0x89, 0xd7, 0xe8, 0x38,
	0xdd, 0xe5, 0xb8, 0x7e, 0x85, 0x1e, 0xc1, 0x32, 0x56, 0x8a, 0xe8, 0x84, 0xa5, 0xde, 0xbc, 0x55,
	0x2e, 0x59, 0x1c, 0xa5, 0xa8, 0x0f, 0xae, 0xce, 0xf1, 0x94, 0x64, 0xc9, 0x44, 0xa4, 0xc4, 0x5b,
	0xe8, 0x38, 0xdd, 0x66, 0xcf, 0x0f, 0x7e, 0x1f, 0x43, 0x70, 0x6a, 0xa9, 0x47, 0x22, 0x25, 0x31,
	0xe8, 0x9b, 0xb3, 0xff, 0xcd, 0x81, 0x55, 0x63, 0x56, 0x92, 0x74, 0x24, 0x18, 0xd7, 0x28, 0x84,
	0xe5, 0x4c, 0x8c, 0x6d, 0x46, 0xd6, 0xaa, 0xdb, 0xdb, 0xa8, 0x5a, 0x9a, 0x51, 0x83, 0x21, 0xd6,
	0x43, 0x4e, 0xe3, 0x1b, 0x12, 0xea, 0x43, 0x53, 0xe4, 0x8c, 0x32, 0x8e, 0xb3, 0x84, 0xf1, 0x94,
	0x7c, 0xb2, 0x63, 0xb8, 0xbd, 0xad, 0x4a, 0x56, 0x25, 0x18, 0x9c, 0x45, 0x5c, 0xef, 0xf4, 0xce,
	0x71, 0x76, 0x45, 0xe2, 0xb5, 0x4a, 0x13, 0x19, 0x89, 0x19, 0x53, 0x66, 0x78, 0x4c, 0x6a, 0x63,
	0x5a, 0x1c, 0xa5, 0xfe, 0x57, 0x07, 0x36, 0xee, 0xc4, 0xa9, 0xa4, 0xe0, 0x8a, 0xa0, 0x43, 0x68,
	0xaa, 0xc2, 0x78, 0x22, 0x8d, 0x73, 0xe5, 0x39, 0x9d, 0xf9, 0xae, 0xdb, 0x7b, 0x3a, 0x2b, 0x81,
	0xfa, 0x98, 0xf1, 0x9a, 0xaa, 0x21, 0x85, 0x9e, 0xc1, 0x7f, 0xd7, 0x38, 0xe7, 0x8c, 0xd3, 0x64,
	0x42, 0x94, 0xc2, 0xb4, 0x58, 0xc6, 0x4a, 0xdc, 0x2c, 0xaf, 0x8f, 0x8a, 0x5b, 0x7f, 0x0a, 0xad,
	0x21, 0x53, 0xfa, 0x98, 0xe0, 0x9c, 0x28, 0x7d, 0x67, 0xc1, 0x9b, 0xb0, 0x78, 0x63, 0xc4, 0x48,
	0x4b, 0x74, 0x7f, 0x4f, 0x8d, 0xbf, 0xda, 0x13, 0x05, 0xef, 0xd7, 0xef, 0xfe, 0x83, 0x24, 0xb6,
	0x47, 0x00, 0xb7, 0x16, 0xd0, 0x63, 0x68, 0x9d, 0xc6, 0x7b, 0xe7, 0x83, 0x61, 0x72, 0xf4, 0xe6,
	0x60, 0x90, 0x9c, 0x1d, 0x9f, 0x8c, 0x06, 0xfd, 0xe8, 0x55, 0x34, 0x38, 0x58, 0x9f, 0x43, 0x2e,
	0x2c, 0x1d, 0xc4, 0xd1, 0x79, 0x74, 0xfc, 0x7a, 0xdd, 0x31, 0xa0, 0xff, 0xae, 0x3f, 0x34, 0xa0,
	0x61, 0xc0, 0xdb, 0xbd, 0xe1, 0xa1, 0x01, 0xf3, 0xbd, 0x2f, 0x0d, 0x58, 0xb5, 0x86, 0x4f, 0x48,
	0x3e, 0x65, 0x63, 0x82, 0x2e, 0xc1, 0xad, 0x2d, 0x14, 0x6d, 0xff, 0xc9, 0xe6, 0xdd, 0x47, 0xd4,
	0x7e, 0xfe, 0x20, 0x6e, 0x91, 0x8b, 0x3f, 0x87, 0xae, 0x61, 0xfd, 0x7e, 0x6a, 0x28, 0x9c, 0xd5,
	0x62, 0xc6, 0x5e, 0xdb, 0x2f, 0x1e, 0x2e, 0xa8, 0x3e, 0xdc, 0x6e, 0xfd, 0xd8, 0xfb, 0xbf, 0x60,
	0x16, 0x5a, 0x2c, 0x99, 0x0a, 0xc6, 0x62, 0xb2, 0x4f, 0xa1, 0x3d, 0x16, 0x93, 0x19, 0x1d, 0xf7,
	0xc1, 0xf6, 0x19, 0x99, 0x07, 0x33, 0x72, 0xde, 0xef, 0x96, 0x2c, 0x2a, 0x32, 0xcc, 0x69, 0x20,
	0x72, 0x1a, 0x52, 0xc2, 0xed, 0x73, 0x0a, 0x6f, 0xdb, 0xde, 0xff, 0x29, 0xee, 0xda, 0xe3, 0xc5,
	0xa2, 0xe5, 0xed, 0xfc, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xa8, 0xb6, 0x93, 0xe3, 0x3b, 0x05, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RoadsServiceClient is the client API for RoadsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RoadsServiceClient interface {
	// This method takes a sequence of latitude,longitude points and snaps them to
	// the most likely road segments. Optionally returns additional points giving
	// the full road geometry. Also returns a place ID for each snapped point.
	SnapToRoads(ctx context.Context, in *SnapToRoadsRequest, opts ...grpc.CallOption) (*SnapToRoadsResponse, error)
	// This method takes a list of latitude,longitude points and snaps them each
	// to their nearest road. Also returns a place ID for each snapped point.
	ListNearestRoads(ctx context.Context, in *ListNearestRoadsRequest, opts ...grpc.CallOption) (*ListNearestRoadsResponse, error)
}

type roadsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRoadsServiceClient(cc grpc.ClientConnInterface) RoadsServiceClient {
	return &roadsServiceClient{cc}
}

func (c *roadsServiceClient) SnapToRoads(ctx context.Context, in *SnapToRoadsRequest, opts ...grpc.CallOption) (*SnapToRoadsResponse, error) {
	out := new(SnapToRoadsResponse)
	err := c.cc.Invoke(ctx, "/google.maps.roads.v1op.RoadsService/SnapToRoads", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roadsServiceClient) ListNearestRoads(ctx context.Context, in *ListNearestRoadsRequest, opts ...grpc.CallOption) (*ListNearestRoadsResponse, error) {
	out := new(ListNearestRoadsResponse)
	err := c.cc.Invoke(ctx, "/google.maps.roads.v1op.RoadsService/ListNearestRoads", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoadsServiceServer is the server API for RoadsService service.
type RoadsServiceServer interface {
	// This method takes a sequence of latitude,longitude points and snaps them to
	// the most likely road segments. Optionally returns additional points giving
	// the full road geometry. Also returns a place ID for each snapped point.
	SnapToRoads(context.Context, *SnapToRoadsRequest) (*SnapToRoadsResponse, error)
	// This method takes a list of latitude,longitude points and snaps them each
	// to their nearest road. Also returns a place ID for each snapped point.
	ListNearestRoads(context.Context, *ListNearestRoadsRequest) (*ListNearestRoadsResponse, error)
}

// UnimplementedRoadsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRoadsServiceServer struct {
}

func (*UnimplementedRoadsServiceServer) SnapToRoads(ctx context.Context, req *SnapToRoadsRequest) (*SnapToRoadsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SnapToRoads not implemented")
}
func (*UnimplementedRoadsServiceServer) ListNearestRoads(ctx context.Context, req *ListNearestRoadsRequest) (*ListNearestRoadsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNearestRoads not implemented")
}

func RegisterRoadsServiceServer(s *grpc.Server, srv RoadsServiceServer) {
	s.RegisterService(&_RoadsService_serviceDesc, srv)
}

func _RoadsService_SnapToRoads_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SnapToRoadsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoadsServiceServer).SnapToRoads(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.maps.roads.v1op.RoadsService/SnapToRoads",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoadsServiceServer).SnapToRoads(ctx, req.(*SnapToRoadsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoadsService_ListNearestRoads_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNearestRoadsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoadsServiceServer).ListNearestRoads(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.maps.roads.v1op.RoadsService/ListNearestRoads",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoadsServiceServer).ListNearestRoads(ctx, req.(*ListNearestRoadsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RoadsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.maps.roads.v1op.RoadsService",
	HandlerType: (*RoadsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SnapToRoads",
			Handler:    _RoadsService_SnapToRoads_Handler,
		},
		{
			MethodName: "ListNearestRoads",
			Handler:    _RoadsService_ListNearestRoads_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/maps/roads/v1op/roads.proto",
}
