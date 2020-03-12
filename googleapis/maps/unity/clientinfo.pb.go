// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/maps/unity/clientinfo.proto

package unity

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

// Platform enum.
type ClientInfo_Platform int32

const (
	// Unspecified or unknown OS.
	ClientInfo_PLATFORM_UNSPECIFIED ClientInfo_Platform = 0
	// Development environment.
	ClientInfo_EDITOR ClientInfo_Platform = 1
	// macOS.
	ClientInfo_MAC_OS ClientInfo_Platform = 2
	// Windows.
	ClientInfo_WINDOWS ClientInfo_Platform = 3
	// Linux
	ClientInfo_LINUX ClientInfo_Platform = 4
	// Android
	ClientInfo_ANDROID ClientInfo_Platform = 5
	// iOS
	ClientInfo_IOS ClientInfo_Platform = 6
	// WebGL.
	ClientInfo_WEB_GL ClientInfo_Platform = 7
)

var ClientInfo_Platform_name = map[int32]string{
	0: "PLATFORM_UNSPECIFIED",
	1: "EDITOR",
	2: "MAC_OS",
	3: "WINDOWS",
	4: "LINUX",
	5: "ANDROID",
	6: "IOS",
	7: "WEB_GL",
}

var ClientInfo_Platform_value = map[string]int32{
	"PLATFORM_UNSPECIFIED": 0,
	"EDITOR":               1,
	"MAC_OS":               2,
	"WINDOWS":              3,
	"LINUX":                4,
	"ANDROID":              5,
	"IOS":                  6,
	"WEB_GL":               7,
}

func (x ClientInfo_Platform) String() string {
	return proto.EnumName(ClientInfo_Platform_name, int32(x))
}

func (ClientInfo_Platform) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f381a22e34055ff1, []int{0, 0}
}

// Client information.
type ClientInfo struct {
	// Application ID, such as the package name on Android and the bundle
	// identifier on iOS platforms.
	ApplicationId string `protobuf:"bytes,1,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
	// Application version number, such as "1.2.3". The exact format is
	// application-dependent.
	ApplicationVersion string `protobuf:"bytes,2,opt,name=application_version,json=applicationVersion,proto3" json:"application_version,omitempty"`
	// Platform where the application is running.
	Platform ClientInfo_Platform `protobuf:"varint,3,opt,name=platform,proto3,enum=google.maps.unity.ClientInfo_Platform" json:"platform,omitempty"`
	// Operating system name and version as reported by the OS. For example,
	// "Mac OS X 10.10.4". The exact format is platform-dependent.
	OperatingSystem string `protobuf:"bytes,4,opt,name=operating_system,json=operatingSystem,proto3" json:"operating_system,omitempty"`
	// API client name and version. For example, the SDK calling the API. The
	// exact format is up to the client.
	ApiClient string `protobuf:"bytes,5,opt,name=api_client,json=apiClient,proto3" json:"api_client,omitempty"`
	// Device model as reported by the device. The exact format is
	// platform-dependent.
	DeviceModel string `protobuf:"bytes,6,opt,name=device_model,json=deviceModel,proto3" json:"device_model,omitempty"`
	// Language code (in BCP-47 format) indicating the UI language of the client.
	// Examples are "en", "en-US" or "ja-Latn". For more information, see
	// http://www.unicode.org/reports/tr35/#Unicode_locale_identifier.
	LanguageCode string `protobuf:"bytes,7,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	// Build number/version of the operating system. e.g., the contents of
	// android.os.Build.ID in Android, or the contents of sysctl "kern.osversion"
	// in iOS.
	OperatingSystemBuild string   `protobuf:"bytes,8,opt,name=operating_system_build,json=operatingSystemBuild,proto3" json:"operating_system_build,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientInfo) Reset()         { *m = ClientInfo{} }
func (m *ClientInfo) String() string { return proto.CompactTextString(m) }
func (*ClientInfo) ProtoMessage()    {}
func (*ClientInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f381a22e34055ff1, []int{0}
}

func (m *ClientInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientInfo.Unmarshal(m, b)
}
func (m *ClientInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientInfo.Marshal(b, m, deterministic)
}
func (m *ClientInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientInfo.Merge(m, src)
}
func (m *ClientInfo) XXX_Size() int {
	return xxx_messageInfo_ClientInfo.Size(m)
}
func (m *ClientInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ClientInfo proto.InternalMessageInfo

func (m *ClientInfo) GetApplicationId() string {
	if m != nil {
		return m.ApplicationId
	}
	return ""
}

func (m *ClientInfo) GetApplicationVersion() string {
	if m != nil {
		return m.ApplicationVersion
	}
	return ""
}

func (m *ClientInfo) GetPlatform() ClientInfo_Platform {
	if m != nil {
		return m.Platform
	}
	return ClientInfo_PLATFORM_UNSPECIFIED
}

func (m *ClientInfo) GetOperatingSystem() string {
	if m != nil {
		return m.OperatingSystem
	}
	return ""
}

func (m *ClientInfo) GetApiClient() string {
	if m != nil {
		return m.ApiClient
	}
	return ""
}

func (m *ClientInfo) GetDeviceModel() string {
	if m != nil {
		return m.DeviceModel
	}
	return ""
}

func (m *ClientInfo) GetLanguageCode() string {
	if m != nil {
		return m.LanguageCode
	}
	return ""
}

func (m *ClientInfo) GetOperatingSystemBuild() string {
	if m != nil {
		return m.OperatingSystemBuild
	}
	return ""
}

func init() {
	proto.RegisterEnum("google.maps.unity.ClientInfo_Platform", ClientInfo_Platform_name, ClientInfo_Platform_value)
	proto.RegisterType((*ClientInfo)(nil), "google.maps.unity.ClientInfo")
}

func init() {
	proto.RegisterFile("google/maps/unity/clientinfo.proto", fileDescriptor_f381a22e34055ff1)
}

var fileDescriptor_f381a22e34055ff1 = []byte{
	// 447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xcd, 0x6a, 0xdb, 0x40,
	0x14, 0x85, 0x6b, 0x2b, 0xfe, 0xbb, 0xf9, 0x53, 0xa6, 0x69, 0xd1, 0xa6, 0x90, 0xba, 0xb4, 0xa4,
	0x1b, 0x09, 0xda, 0xd2, 0x4d, 0x57, 0x96, 0xe5, 0x18, 0x81, 0x65, 0x09, 0x29, 0xae, 0x4b, 0x37,
	0xc3, 0x44, 0x1a, 0x8b, 0x01, 0x69, 0x66, 0x90, 0x64, 0x43, 0xa0, 0x4f, 0xd3, 0x65, 0x9f, 0xac,
	0x8f, 0x11, 0x66, 0x94, 0x1f, 0x93, 0x6c, 0x84, 0x38, 0xe7, 0x3b, 0xdc, 0x33, 0x97, 0x0b, 0xe3,
	0x5c, 0x88, 0xbc, 0xa0, 0x4e, 0x49, 0x64, 0xed, 0x6c, 0x39, 0x6b, 0x6e, 0x9d, 0xb4, 0x60, 0x94,
	0x37, 0x8c, 0x6f, 0x84, 0x2d, 0x2b, 0xd1, 0x08, 0x74, 0xd6, 0x32, 0xb6, 0x62, 0x6c, 0xcd, 0x8c,
	0xff, 0x1b, 0x00, 0x53, 0xcd, 0xf9, 0x7c, 0x23, 0xd0, 0x47, 0x38, 0x21, 0x52, 0x16, 0x2c, 0x25,
	0x0d, 0x13, 0x1c, 0xb3, 0xcc, 0xea, 0x5c, 0x74, 0x2e, 0x47, 0xf1, 0xf1, 0x9e, 0xea, 0x67, 0xc8,
	0x81, 0xd7, 0xfb, 0xd8, 0x8e, 0x56, 0x35, 0x13, 0xdc, 0xea, 0x6a, 0x16, 0xed, 0x59, 0x3f, 0x5b,
	0x07, 0xb9, 0x30, 0x94, 0x05, 0x69, 0x36, 0xa2, 0x2a, 0x2d, 0xe3, 0xa2, 0x73, 0x79, 0xf2, 0xe5,
	0x93, 0xfd, 0xa2, 0x8c, 0xfd, 0x54, 0xc4, 0x8e, 0xee, 0xe9, 0xf8, 0x31, 0x87, 0x3e, 0x83, 0x29,
	0x24, 0xad, 0x48, 0xc3, 0x78, 0x8e, 0xeb, 0xdb, 0xba, 0xa1, 0xa5, 0x75, 0xa0, 0x27, 0x9e, 0x3e,
	0xea, 0x89, 0x96, 0xd1, 0x3b, 0x00, 0x22, 0x19, 0x6e, 0x17, 0x60, 0xf5, 0x34, 0x34, 0x22, 0x92,
	0xb5, 0x03, 0xd0, 0x7b, 0x38, 0xca, 0xe8, 0x8e, 0xa5, 0x14, 0x97, 0x22, 0xa3, 0x85, 0xd5, 0xd7,
	0xc0, 0x61, 0xab, 0x05, 0x4a, 0x42, 0x1f, 0xe0, 0xb8, 0x20, 0x3c, 0xdf, 0x92, 0x9c, 0xe2, 0x54,
	0x64, 0xd4, 0x1a, 0x68, 0xe6, 0xe8, 0x41, 0x9c, 0x8a, 0x8c, 0xa2, 0x6f, 0xf0, 0xf6, 0x79, 0x23,
	0x7c, 0xb3, 0x65, 0x45, 0x66, 0x0d, 0x35, 0x7d, 0xfe, 0xac, 0x97, 0xab, 0xbc, 0xf1, 0x0e, 0x86,
	0x0f, 0xaf, 0x43, 0x16, 0x9c, 0x47, 0x8b, 0xc9, 0xf5, 0x55, 0x18, 0x07, 0x78, 0xb5, 0x4c, 0xa2,
	0xd9, 0xd4, 0xbf, 0xf2, 0x67, 0x9e, 0xf9, 0x0a, 0x01, 0xf4, 0x67, 0x9e, 0x7f, 0x1d, 0xc6, 0x66,
	0x47, 0xfd, 0x07, 0x93, 0x29, 0x0e, 0x13, 0xb3, 0x8b, 0x0e, 0x61, 0xb0, 0xf6, 0x97, 0x5e, 0xb8,
	0x4e, 0x4c, 0x03, 0x8d, 0xa0, 0xb7, 0xf0, 0x97, 0xab, 0x5f, 0xe6, 0x81, 0xd2, 0x27, 0x4b, 0x2f,
	0x0e, 0x7d, 0xcf, 0xec, 0xa1, 0x01, 0x18, 0x7e, 0x98, 0x98, 0x7d, 0x95, 0x5c, 0xcf, 0x5c, 0x3c,
	0x5f, 0x98, 0x03, 0xf7, 0x0f, 0xbc, 0x49, 0x45, 0xf9, 0x72, 0xed, 0xee, 0xe9, 0xd3, 0xde, 0x23,
	0x75, 0x27, 0x51, 0xe7, 0xf7, 0xf7, 0x7b, 0x2a, 0x17, 0xea, 0xc1, 0xb6, 0xa8, 0x72, 0x27, 0xa7,
	0x5c, 0x5f, 0x91, 0xd3, 0x5a, 0x44, 0xb2, 0x7a, 0xef, 0xd8, 0x7e, 0xe8, 0xef, 0xdf, 0xae, 0x31,
	0x0f, 0x56, 0xff, 0xba, 0x67, 0xf3, 0x36, 0x1e, 0xa8, 0x21, 0x2b, 0xe5, 0xdc, 0xf4, 0x75, 0xf8,
	0xeb, 0x5d, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd9, 0x70, 0x1f, 0xe7, 0xa8, 0x02, 0x00, 0x00,
}
