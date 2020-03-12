// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/vision/v1p4beta1/web_detection.proto

package vision

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

// Relevant information for the image from the Internet.
type WebDetection struct {
	// Deduced entities from similar images on the Internet.
	WebEntities []*WebDetection_WebEntity `protobuf:"bytes,1,rep,name=web_entities,json=webEntities,proto3" json:"web_entities,omitempty"`
	// Fully matching images from the Internet.
	// Can include resized copies of the query image.
	FullMatchingImages []*WebDetection_WebImage `protobuf:"bytes,2,rep,name=full_matching_images,json=fullMatchingImages,proto3" json:"full_matching_images,omitempty"`
	// Partial matching images from the Internet.
	// Those images are similar enough to share some key-point features. For
	// example an original image will likely have partial matching for its crops.
	PartialMatchingImages []*WebDetection_WebImage `protobuf:"bytes,3,rep,name=partial_matching_images,json=partialMatchingImages,proto3" json:"partial_matching_images,omitempty"`
	// Web pages containing the matching images from the Internet.
	PagesWithMatchingImages []*WebDetection_WebPage `protobuf:"bytes,4,rep,name=pages_with_matching_images,json=pagesWithMatchingImages,proto3" json:"pages_with_matching_images,omitempty"`
	// The visually similar image results.
	VisuallySimilarImages []*WebDetection_WebImage `protobuf:"bytes,6,rep,name=visually_similar_images,json=visuallySimilarImages,proto3" json:"visually_similar_images,omitempty"`
	// The service's best guess as to the topic of the request image.
	// Inferred from similar images on the open web.
	BestGuessLabels      []*WebDetection_WebLabel `protobuf:"bytes,8,rep,name=best_guess_labels,json=bestGuessLabels,proto3" json:"best_guess_labels,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *WebDetection) Reset()         { *m = WebDetection{} }
func (m *WebDetection) String() string { return proto.CompactTextString(m) }
func (*WebDetection) ProtoMessage()    {}
func (*WebDetection) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3e7265af1b70904, []int{0}
}

func (m *WebDetection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WebDetection.Unmarshal(m, b)
}
func (m *WebDetection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WebDetection.Marshal(b, m, deterministic)
}
func (m *WebDetection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebDetection.Merge(m, src)
}
func (m *WebDetection) XXX_Size() int {
	return xxx_messageInfo_WebDetection.Size(m)
}
func (m *WebDetection) XXX_DiscardUnknown() {
	xxx_messageInfo_WebDetection.DiscardUnknown(m)
}

var xxx_messageInfo_WebDetection proto.InternalMessageInfo

func (m *WebDetection) GetWebEntities() []*WebDetection_WebEntity {
	if m != nil {
		return m.WebEntities
	}
	return nil
}

func (m *WebDetection) GetFullMatchingImages() []*WebDetection_WebImage {
	if m != nil {
		return m.FullMatchingImages
	}
	return nil
}

func (m *WebDetection) GetPartialMatchingImages() []*WebDetection_WebImage {
	if m != nil {
		return m.PartialMatchingImages
	}
	return nil
}

func (m *WebDetection) GetPagesWithMatchingImages() []*WebDetection_WebPage {
	if m != nil {
		return m.PagesWithMatchingImages
	}
	return nil
}

func (m *WebDetection) GetVisuallySimilarImages() []*WebDetection_WebImage {
	if m != nil {
		return m.VisuallySimilarImages
	}
	return nil
}

func (m *WebDetection) GetBestGuessLabels() []*WebDetection_WebLabel {
	if m != nil {
		return m.BestGuessLabels
	}
	return nil
}

// Entity deduced from similar images on the Internet.
type WebDetection_WebEntity struct {
	// Opaque entity ID.
	EntityId string `protobuf:"bytes,1,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	// Overall relevancy score for the entity.
	// Not normalized and not comparable across different image queries.
	Score float32 `protobuf:"fixed32,2,opt,name=score,proto3" json:"score,omitempty"`
	// Canonical description of the entity, in English.
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WebDetection_WebEntity) Reset()         { *m = WebDetection_WebEntity{} }
func (m *WebDetection_WebEntity) String() string { return proto.CompactTextString(m) }
func (*WebDetection_WebEntity) ProtoMessage()    {}
func (*WebDetection_WebEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3e7265af1b70904, []int{0, 0}
}

func (m *WebDetection_WebEntity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WebDetection_WebEntity.Unmarshal(m, b)
}
func (m *WebDetection_WebEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WebDetection_WebEntity.Marshal(b, m, deterministic)
}
func (m *WebDetection_WebEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebDetection_WebEntity.Merge(m, src)
}
func (m *WebDetection_WebEntity) XXX_Size() int {
	return xxx_messageInfo_WebDetection_WebEntity.Size(m)
}
func (m *WebDetection_WebEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_WebDetection_WebEntity.DiscardUnknown(m)
}

var xxx_messageInfo_WebDetection_WebEntity proto.InternalMessageInfo

func (m *WebDetection_WebEntity) GetEntityId() string {
	if m != nil {
		return m.EntityId
	}
	return ""
}

func (m *WebDetection_WebEntity) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *WebDetection_WebEntity) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// Metadata for online images.
type WebDetection_WebImage struct {
	// The result image URL.
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// (Deprecated) Overall relevancy score for the image.
	Score                float32  `protobuf:"fixed32,2,opt,name=score,proto3" json:"score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WebDetection_WebImage) Reset()         { *m = WebDetection_WebImage{} }
func (m *WebDetection_WebImage) String() string { return proto.CompactTextString(m) }
func (*WebDetection_WebImage) ProtoMessage()    {}
func (*WebDetection_WebImage) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3e7265af1b70904, []int{0, 1}
}

func (m *WebDetection_WebImage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WebDetection_WebImage.Unmarshal(m, b)
}
func (m *WebDetection_WebImage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WebDetection_WebImage.Marshal(b, m, deterministic)
}
func (m *WebDetection_WebImage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebDetection_WebImage.Merge(m, src)
}
func (m *WebDetection_WebImage) XXX_Size() int {
	return xxx_messageInfo_WebDetection_WebImage.Size(m)
}
func (m *WebDetection_WebImage) XXX_DiscardUnknown() {
	xxx_messageInfo_WebDetection_WebImage.DiscardUnknown(m)
}

var xxx_messageInfo_WebDetection_WebImage proto.InternalMessageInfo

func (m *WebDetection_WebImage) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *WebDetection_WebImage) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

// Metadata for web pages.
type WebDetection_WebPage struct {
	// The result web page URL.
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// (Deprecated) Overall relevancy score for the web page.
	Score float32 `protobuf:"fixed32,2,opt,name=score,proto3" json:"score,omitempty"`
	// Title for the web page, may contain HTML markups.
	PageTitle string `protobuf:"bytes,3,opt,name=page_title,json=pageTitle,proto3" json:"page_title,omitempty"`
	// Fully matching images on the page.
	// Can include resized copies of the query image.
	FullMatchingImages []*WebDetection_WebImage `protobuf:"bytes,4,rep,name=full_matching_images,json=fullMatchingImages,proto3" json:"full_matching_images,omitempty"`
	// Partial matching images on the page.
	// Those images are similar enough to share some key-point features. For
	// example an original image will likely have partial matching for its
	// crops.
	PartialMatchingImages []*WebDetection_WebImage `protobuf:"bytes,5,rep,name=partial_matching_images,json=partialMatchingImages,proto3" json:"partial_matching_images,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                 `json:"-"`
	XXX_unrecognized      []byte                   `json:"-"`
	XXX_sizecache         int32                    `json:"-"`
}

func (m *WebDetection_WebPage) Reset()         { *m = WebDetection_WebPage{} }
func (m *WebDetection_WebPage) String() string { return proto.CompactTextString(m) }
func (*WebDetection_WebPage) ProtoMessage()    {}
func (*WebDetection_WebPage) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3e7265af1b70904, []int{0, 2}
}

func (m *WebDetection_WebPage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WebDetection_WebPage.Unmarshal(m, b)
}
func (m *WebDetection_WebPage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WebDetection_WebPage.Marshal(b, m, deterministic)
}
func (m *WebDetection_WebPage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebDetection_WebPage.Merge(m, src)
}
func (m *WebDetection_WebPage) XXX_Size() int {
	return xxx_messageInfo_WebDetection_WebPage.Size(m)
}
func (m *WebDetection_WebPage) XXX_DiscardUnknown() {
	xxx_messageInfo_WebDetection_WebPage.DiscardUnknown(m)
}

var xxx_messageInfo_WebDetection_WebPage proto.InternalMessageInfo

func (m *WebDetection_WebPage) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *WebDetection_WebPage) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *WebDetection_WebPage) GetPageTitle() string {
	if m != nil {
		return m.PageTitle
	}
	return ""
}

func (m *WebDetection_WebPage) GetFullMatchingImages() []*WebDetection_WebImage {
	if m != nil {
		return m.FullMatchingImages
	}
	return nil
}

func (m *WebDetection_WebPage) GetPartialMatchingImages() []*WebDetection_WebImage {
	if m != nil {
		return m.PartialMatchingImages
	}
	return nil
}

// Label to provide extra metadata for the web detection.
type WebDetection_WebLabel struct {
	// Label for extra metadata.
	Label string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	// The BCP-47 language code for `label`, such as "en-US" or "sr-Latn".
	// For more information, see
	// http://www.unicode.org/reports/tr35/#Unicode_locale_identifier.
	LanguageCode         string   `protobuf:"bytes,2,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WebDetection_WebLabel) Reset()         { *m = WebDetection_WebLabel{} }
func (m *WebDetection_WebLabel) String() string { return proto.CompactTextString(m) }
func (*WebDetection_WebLabel) ProtoMessage()    {}
func (*WebDetection_WebLabel) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3e7265af1b70904, []int{0, 3}
}

func (m *WebDetection_WebLabel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WebDetection_WebLabel.Unmarshal(m, b)
}
func (m *WebDetection_WebLabel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WebDetection_WebLabel.Marshal(b, m, deterministic)
}
func (m *WebDetection_WebLabel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebDetection_WebLabel.Merge(m, src)
}
func (m *WebDetection_WebLabel) XXX_Size() int {
	return xxx_messageInfo_WebDetection_WebLabel.Size(m)
}
func (m *WebDetection_WebLabel) XXX_DiscardUnknown() {
	xxx_messageInfo_WebDetection_WebLabel.DiscardUnknown(m)
}

var xxx_messageInfo_WebDetection_WebLabel proto.InternalMessageInfo

func (m *WebDetection_WebLabel) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *WebDetection_WebLabel) GetLanguageCode() string {
	if m != nil {
		return m.LanguageCode
	}
	return ""
}

func init() {
	proto.RegisterType((*WebDetection)(nil), "google.cloud.vision.v1p4beta1.WebDetection")
	proto.RegisterType((*WebDetection_WebEntity)(nil), "google.cloud.vision.v1p4beta1.WebDetection.WebEntity")
	proto.RegisterType((*WebDetection_WebImage)(nil), "google.cloud.vision.v1p4beta1.WebDetection.WebImage")
	proto.RegisterType((*WebDetection_WebPage)(nil), "google.cloud.vision.v1p4beta1.WebDetection.WebPage")
	proto.RegisterType((*WebDetection_WebLabel)(nil), "google.cloud.vision.v1p4beta1.WebDetection.WebLabel")
}

func init() {
	proto.RegisterFile("google/cloud/vision/v1p4beta1/web_detection.proto", fileDescriptor_f3e7265af1b70904)
}

var fileDescriptor_f3e7265af1b70904 = []byte{
	// 519 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x95, 0xb4, 0x1b, 0x8d, 0x5b, 0x04, 0xb3, 0x86, 0x16, 0x05, 0x26, 0x15, 0xb8, 0xf4,
	0x94, 0xa8, 0xdb, 0x38, 0x71, 0x5b, 0x99, 0xa6, 0x49, 0x80, 0xaa, 0x80, 0x18, 0xe2, 0x92, 0x39,
	0x89, 0x97, 0x5a, 0x72, 0xe3, 0x28, 0x76, 0x5a, 0xf5, 0x0b, 0xc0, 0xf7, 0xe0, 0xc0, 0x67, 0xe4,
	0x88, 0x5e, 0xdb, 0x99, 0x2a, 0xca, 0x26, 0xc6, 0xd0, 0x6e, 0x7e, 0x1f, 0xeb, 0x79, 0x7e, 0xf6,
	0xeb, 0x3f, 0x68, 0x5c, 0x08, 0x51, 0x70, 0x1a, 0x65, 0x5c, 0x34, 0x79, 0xb4, 0x60, 0x92, 0x89,
	0x32, 0x5a, 0x8c, 0xab, 0xa3, 0x94, 0x2a, 0x32, 0x8e, 0x96, 0x34, 0x4d, 0x72, 0xaa, 0x68, 0xa6,
	0x98, 0x28, 0xc3, 0xaa, 0x16, 0x4a, 0xe0, 0x7d, 0x63, 0x09, 0xb5, 0x25, 0x34, 0x96, 0xf0, 0xca,
	0x12, 0x3c, 0xb3, 0x89, 0xa4, 0x62, 0x11, 0x29, 0x4b, 0xa1, 0x08, 0x78, 0xa5, 0x31, 0xbf, 0xf8,
	0xea, 0xa1, 0xc1, 0x39, 0x4d, 0xdf, 0xb4, 0x99, 0xf8, 0x33, 0x1a, 0x00, 0x84, 0x96, 0x8a, 0x29,
	0x46, 0xa5, 0xef, 0x0c, 0x3b, 0xa3, 0xfe, 0xc1, 0xab, 0xf0, 0x46, 0x48, 0xb8, 0x1e, 0x01, 0xc5,
	0x09, 0xd8, 0x57, 0x71, 0x7f, 0x69, 0x87, 0x8c, 0x4a, 0x7c, 0x89, 0x76, 0x2f, 0x1b, 0xce, 0x93,
	0x39, 0x51, 0xd9, 0x8c, 0x95, 0x45, 0xc2, 0xe6, 0xa4, 0xa0, 0xd2, 0x77, 0x35, 0xe1, 0xe8, 0x96,
	0x84, 0x33, 0x30, 0xc7, 0x18, 0x12, 0xdf, 0xd9, 0x40, 0x2d, 0x49, 0xcc, 0xd1, 0x5e, 0x45, 0x6a,
	0xc5, 0xc8, 0x26, 0xaa, 0x73, 0x07, 0xd4, 0x13, 0x1b, 0xfa, 0x1b, 0xad, 0x42, 0x41, 0x05, 0x83,
	0x64, 0xc9, 0xd4, 0x6c, 0x03, 0xd8, 0xd5, 0xc0, 0xc3, 0x5b, 0x02, 0xa7, 0xc0, 0xdb, 0xd3, 0xb1,
	0xe7, 0x4c, 0xcd, 0x36, 0xf7, 0xb7, 0x60, 0xb2, 0x21, 0x9c, 0xaf, 0x12, 0xc9, 0xe6, 0x8c, 0x93,
	0xba, 0xc5, 0x6d, 0xdf, 0x65, 0x7f, 0x6d, 0xe8, 0x07, 0x93, 0x69, 0x69, 0x17, 0x68, 0x27, 0xa5,
	0x52, 0x25, 0x45, 0x43, 0xa5, 0x4c, 0x38, 0x49, 0x29, 0x97, 0x7e, 0xef, 0x9f, 0x38, 0x6f, 0xc1,
	0x1c, 0x3f, 0x82, 0xb8, 0x53, 0x48, 0xd3, 0xb5, 0x0c, 0x2e, 0x90, 0x77, 0x75, 0x63, 0xf0, 0x53,
	0xe4, 0xe9, 0xab, 0xb7, 0x4a, 0x58, 0xee, 0x3b, 0x43, 0x67, 0xe4, 0xc5, 0x3d, 0x23, 0x9c, 0xe5,
	0x78, 0x17, 0x6d, 0xc9, 0x4c, 0xd4, 0xd4, 0x77, 0x87, 0xce, 0xc8, 0x8d, 0x4d, 0x81, 0x87, 0xa8,
	0x9f, 0x53, 0x99, 0xd5, 0xac, 0x02, 0x90, 0xdf, 0xd1, 0xa6, 0x75, 0x29, 0x38, 0x40, 0xbd, 0x76,
	0x9b, 0xf8, 0x31, 0xea, 0x34, 0x35, 0xb7, 0xd1, 0x30, 0xfc, 0x73, 0x6a, 0xf0, 0xc3, 0x45, 0x0f,
	0xec, 0x51, 0xfc, 0xad, 0x07, 0xef, 0x23, 0x04, 0x87, 0x96, 0x28, 0xa6, 0x38, 0xb5, 0x0b, 0xf1,
	0x40, 0xf9, 0x08, 0xc2, 0xb5, 0x0f, 0xa0, 0x7b, 0x7f, 0x0f, 0x60, 0xeb, 0xbf, 0x3f, 0x80, 0xe0,
	0x44, 0x37, 0x57, 0x9f, 0x25, 0xb4, 0x45, 0xdf, 0x10, 0xdb, 0x2a, 0x53, 0xe0, 0x97, 0xe8, 0x21,
	0x27, 0x65, 0xd1, 0x40, 0x6b, 0x32, 0x91, 0x9b, 0xa6, 0x79, 0xf1, 0xa0, 0x15, 0x27, 0x22, 0xa7,
	0xc7, 0xdf, 0x1c, 0xf4, 0x3c, 0x13, 0xf3, 0x9b, 0x57, 0x76, 0xbc, 0xb3, 0xbe, 0xb4, 0x29, 0xfc,
	0x60, 0x53, 0xe7, 0xcb, 0xc4, 0x7a, 0x0a, 0x01, 0x89, 0xa1, 0xa8, 0x8b, 0xa8, 0xa0, 0xa5, 0xfe,
	0xdf, 0x22, 0x33, 0x45, 0x2a, 0x26, 0xaf, 0xf9, 0x52, 0x5f, 0x1b, 0xe1, 0xa7, 0xe3, 0x7c, 0x77,
	0xbb, 0xa7, 0x93, 0x4f, 0xef, 0xd3, 0x6d, 0xed, 0x3c, 0xfc, 0x15, 0x00, 0x00, 0xff, 0xff, 0x46,
	0x6d, 0x5f, 0x2f, 0x8b, 0x05, 0x00, 0x00,
}
