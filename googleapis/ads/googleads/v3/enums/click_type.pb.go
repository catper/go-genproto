// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/click_type.proto

package enums

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

// Enumerates Google Ads click types.
type ClickTypeEnum_ClickType int32

const (
	// Not specified.
	ClickTypeEnum_UNSPECIFIED ClickTypeEnum_ClickType = 0
	// The value is unknown in this version.
	ClickTypeEnum_UNKNOWN ClickTypeEnum_ClickType = 1
	// App engagement ad deep link.
	ClickTypeEnum_APP_DEEPLINK ClickTypeEnum_ClickType = 2
	// Breadcrumbs.
	ClickTypeEnum_BREADCRUMBS ClickTypeEnum_ClickType = 3
	// Broadband Plan.
	ClickTypeEnum_BROADBAND_PLAN ClickTypeEnum_ClickType = 4
	// Manually dialed phone calls.
	ClickTypeEnum_CALL_TRACKING ClickTypeEnum_ClickType = 5
	// Phone calls.
	ClickTypeEnum_CALLS ClickTypeEnum_ClickType = 6
	// Click on engagement ad.
	ClickTypeEnum_CLICK_ON_ENGAGEMENT_AD ClickTypeEnum_ClickType = 7
	// Driving direction.
	ClickTypeEnum_GET_DIRECTIONS ClickTypeEnum_ClickType = 8
	// Get location details.
	ClickTypeEnum_LOCATION_EXPANSION ClickTypeEnum_ClickType = 9
	// Call.
	ClickTypeEnum_LOCATION_FORMAT_CALL ClickTypeEnum_ClickType = 10
	// Directions.
	ClickTypeEnum_LOCATION_FORMAT_DIRECTIONS ClickTypeEnum_ClickType = 11
	// Image(s).
	ClickTypeEnum_LOCATION_FORMAT_IMAGE ClickTypeEnum_ClickType = 12
	// Go to landing page.
	ClickTypeEnum_LOCATION_FORMAT_LANDING_PAGE ClickTypeEnum_ClickType = 13
	// Map.
	ClickTypeEnum_LOCATION_FORMAT_MAP ClickTypeEnum_ClickType = 14
	// Go to store info.
	ClickTypeEnum_LOCATION_FORMAT_STORE_INFO ClickTypeEnum_ClickType = 15
	// Text.
	ClickTypeEnum_LOCATION_FORMAT_TEXT ClickTypeEnum_ClickType = 16
	// Mobile phone calls.
	ClickTypeEnum_MOBILE_CALL_TRACKING ClickTypeEnum_ClickType = 17
	// Print offer.
	ClickTypeEnum_OFFER_PRINTS ClickTypeEnum_ClickType = 18
	// Other.
	ClickTypeEnum_OTHER ClickTypeEnum_ClickType = 19
	// Product plusbox offer.
	ClickTypeEnum_PRODUCT_EXTENSION_CLICKS ClickTypeEnum_ClickType = 20
	// Shopping - Product - Online.
	ClickTypeEnum_PRODUCT_LISTING_AD_CLICKS ClickTypeEnum_ClickType = 21
	// Sitelink.
	ClickTypeEnum_SITELINKS ClickTypeEnum_ClickType = 22
	// Show nearby locations.
	ClickTypeEnum_STORE_LOCATOR ClickTypeEnum_ClickType = 23
	// Headline.
	ClickTypeEnum_URL_CLICKS ClickTypeEnum_ClickType = 25
	// App store.
	ClickTypeEnum_VIDEO_APP_STORE_CLICKS ClickTypeEnum_ClickType = 26
	// Call-to-Action overlay.
	ClickTypeEnum_VIDEO_CALL_TO_ACTION_CLICKS ClickTypeEnum_ClickType = 27
	// Cards.
	ClickTypeEnum_VIDEO_CARD_ACTION_HEADLINE_CLICKS ClickTypeEnum_ClickType = 28
	// End cap.
	ClickTypeEnum_VIDEO_END_CAP_CLICKS ClickTypeEnum_ClickType = 29
	// Website.
	ClickTypeEnum_VIDEO_WEBSITE_CLICKS ClickTypeEnum_ClickType = 30
	// Visual Sitelinks.
	ClickTypeEnum_VISUAL_SITELINKS ClickTypeEnum_ClickType = 31
	// Wireless Plan.
	ClickTypeEnum_WIRELESS_PLAN ClickTypeEnum_ClickType = 32
	// Shopping - Product - Local.
	ClickTypeEnum_PRODUCT_LISTING_AD_LOCAL ClickTypeEnum_ClickType = 33
	// Shopping - Product - MultiChannel Local.
	ClickTypeEnum_PRODUCT_LISTING_AD_MULTICHANNEL_LOCAL ClickTypeEnum_ClickType = 34
	// Shopping - Product - MultiChannel Online.
	ClickTypeEnum_PRODUCT_LISTING_AD_MULTICHANNEL_ONLINE ClickTypeEnum_ClickType = 35
	// Shopping - Product - Coupon.
	ClickTypeEnum_PRODUCT_LISTING_ADS_COUPON ClickTypeEnum_ClickType = 36
	// Shopping - Product - Sell on Google.
	ClickTypeEnum_PRODUCT_LISTING_AD_TRANSACTABLE ClickTypeEnum_ClickType = 37
	// Shopping - Product - App engagement ad deep link.
	ClickTypeEnum_PRODUCT_AD_APP_DEEPLINK ClickTypeEnum_ClickType = 38
	// Shopping - Showcase - Category.
	ClickTypeEnum_SHOWCASE_AD_CATEGORY_LINK ClickTypeEnum_ClickType = 39
	// Shopping - Showcase - Local storefront.
	ClickTypeEnum_SHOWCASE_AD_LOCAL_STOREFRONT_LINK ClickTypeEnum_ClickType = 40
	// Shopping - Showcase - Online product.
	ClickTypeEnum_SHOWCASE_AD_ONLINE_PRODUCT_LINK ClickTypeEnum_ClickType = 42
	// Shopping - Showcase - Local product.
	ClickTypeEnum_SHOWCASE_AD_LOCAL_PRODUCT_LINK ClickTypeEnum_ClickType = 43
	// Promotion Extension.
	ClickTypeEnum_PROMOTION_EXTENSION ClickTypeEnum_ClickType = 44
	// Ad Headline.
	ClickTypeEnum_SWIPEABLE_GALLERY_AD_HEADLINE ClickTypeEnum_ClickType = 45
	// Swipes.
	ClickTypeEnum_SWIPEABLE_GALLERY_AD_SWIPES ClickTypeEnum_ClickType = 46
	// See More.
	ClickTypeEnum_SWIPEABLE_GALLERY_AD_SEE_MORE ClickTypeEnum_ClickType = 47
	// Sitelink 1.
	ClickTypeEnum_SWIPEABLE_GALLERY_AD_SITELINK_ONE ClickTypeEnum_ClickType = 48
	// Sitelink 2.
	ClickTypeEnum_SWIPEABLE_GALLERY_AD_SITELINK_TWO ClickTypeEnum_ClickType = 49
	// Sitelink 3.
	ClickTypeEnum_SWIPEABLE_GALLERY_AD_SITELINK_THREE ClickTypeEnum_ClickType = 50
	// Sitelink 4.
	ClickTypeEnum_SWIPEABLE_GALLERY_AD_SITELINK_FOUR ClickTypeEnum_ClickType = 51
	// Sitelink 5.
	ClickTypeEnum_SWIPEABLE_GALLERY_AD_SITELINK_FIVE ClickTypeEnum_ClickType = 52
	// Hotel price.
	ClickTypeEnum_HOTEL_PRICE ClickTypeEnum_ClickType = 53
	// Price Extension.
	ClickTypeEnum_PRICE_EXTENSION ClickTypeEnum_ClickType = 54
	// Book on Google hotel room selection.
	ClickTypeEnum_HOTEL_BOOK_ON_GOOGLE_ROOM_SELECTION ClickTypeEnum_ClickType = 55
	// Shopping - Comparison Listing.
	ClickTypeEnum_SHOPPING_COMPARISON_LISTING ClickTypeEnum_ClickType = 56
)

var ClickTypeEnum_ClickType_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "APP_DEEPLINK",
	3:  "BREADCRUMBS",
	4:  "BROADBAND_PLAN",
	5:  "CALL_TRACKING",
	6:  "CALLS",
	7:  "CLICK_ON_ENGAGEMENT_AD",
	8:  "GET_DIRECTIONS",
	9:  "LOCATION_EXPANSION",
	10: "LOCATION_FORMAT_CALL",
	11: "LOCATION_FORMAT_DIRECTIONS",
	12: "LOCATION_FORMAT_IMAGE",
	13: "LOCATION_FORMAT_LANDING_PAGE",
	14: "LOCATION_FORMAT_MAP",
	15: "LOCATION_FORMAT_STORE_INFO",
	16: "LOCATION_FORMAT_TEXT",
	17: "MOBILE_CALL_TRACKING",
	18: "OFFER_PRINTS",
	19: "OTHER",
	20: "PRODUCT_EXTENSION_CLICKS",
	21: "PRODUCT_LISTING_AD_CLICKS",
	22: "SITELINKS",
	23: "STORE_LOCATOR",
	25: "URL_CLICKS",
	26: "VIDEO_APP_STORE_CLICKS",
	27: "VIDEO_CALL_TO_ACTION_CLICKS",
	28: "VIDEO_CARD_ACTION_HEADLINE_CLICKS",
	29: "VIDEO_END_CAP_CLICKS",
	30: "VIDEO_WEBSITE_CLICKS",
	31: "VISUAL_SITELINKS",
	32: "WIRELESS_PLAN",
	33: "PRODUCT_LISTING_AD_LOCAL",
	34: "PRODUCT_LISTING_AD_MULTICHANNEL_LOCAL",
	35: "PRODUCT_LISTING_AD_MULTICHANNEL_ONLINE",
	36: "PRODUCT_LISTING_ADS_COUPON",
	37: "PRODUCT_LISTING_AD_TRANSACTABLE",
	38: "PRODUCT_AD_APP_DEEPLINK",
	39: "SHOWCASE_AD_CATEGORY_LINK",
	40: "SHOWCASE_AD_LOCAL_STOREFRONT_LINK",
	42: "SHOWCASE_AD_ONLINE_PRODUCT_LINK",
	43: "SHOWCASE_AD_LOCAL_PRODUCT_LINK",
	44: "PROMOTION_EXTENSION",
	45: "SWIPEABLE_GALLERY_AD_HEADLINE",
	46: "SWIPEABLE_GALLERY_AD_SWIPES",
	47: "SWIPEABLE_GALLERY_AD_SEE_MORE",
	48: "SWIPEABLE_GALLERY_AD_SITELINK_ONE",
	49: "SWIPEABLE_GALLERY_AD_SITELINK_TWO",
	50: "SWIPEABLE_GALLERY_AD_SITELINK_THREE",
	51: "SWIPEABLE_GALLERY_AD_SITELINK_FOUR",
	52: "SWIPEABLE_GALLERY_AD_SITELINK_FIVE",
	53: "HOTEL_PRICE",
	54: "PRICE_EXTENSION",
	55: "HOTEL_BOOK_ON_GOOGLE_ROOM_SELECTION",
	56: "SHOPPING_COMPARISON_LISTING",
}

var ClickTypeEnum_ClickType_value = map[string]int32{
	"UNSPECIFIED":                            0,
	"UNKNOWN":                                1,
	"APP_DEEPLINK":                           2,
	"BREADCRUMBS":                            3,
	"BROADBAND_PLAN":                         4,
	"CALL_TRACKING":                          5,
	"CALLS":                                  6,
	"CLICK_ON_ENGAGEMENT_AD":                 7,
	"GET_DIRECTIONS":                         8,
	"LOCATION_EXPANSION":                     9,
	"LOCATION_FORMAT_CALL":                   10,
	"LOCATION_FORMAT_DIRECTIONS":             11,
	"LOCATION_FORMAT_IMAGE":                  12,
	"LOCATION_FORMAT_LANDING_PAGE":           13,
	"LOCATION_FORMAT_MAP":                    14,
	"LOCATION_FORMAT_STORE_INFO":             15,
	"LOCATION_FORMAT_TEXT":                   16,
	"MOBILE_CALL_TRACKING":                   17,
	"OFFER_PRINTS":                           18,
	"OTHER":                                  19,
	"PRODUCT_EXTENSION_CLICKS":               20,
	"PRODUCT_LISTING_AD_CLICKS":              21,
	"SITELINKS":                              22,
	"STORE_LOCATOR":                          23,
	"URL_CLICKS":                             25,
	"VIDEO_APP_STORE_CLICKS":                 26,
	"VIDEO_CALL_TO_ACTION_CLICKS":            27,
	"VIDEO_CARD_ACTION_HEADLINE_CLICKS":      28,
	"VIDEO_END_CAP_CLICKS":                   29,
	"VIDEO_WEBSITE_CLICKS":                   30,
	"VISUAL_SITELINKS":                       31,
	"WIRELESS_PLAN":                          32,
	"PRODUCT_LISTING_AD_LOCAL":               33,
	"PRODUCT_LISTING_AD_MULTICHANNEL_LOCAL":  34,
	"PRODUCT_LISTING_AD_MULTICHANNEL_ONLINE": 35,
	"PRODUCT_LISTING_ADS_COUPON":             36,
	"PRODUCT_LISTING_AD_TRANSACTABLE":        37,
	"PRODUCT_AD_APP_DEEPLINK":                38,
	"SHOWCASE_AD_CATEGORY_LINK":              39,
	"SHOWCASE_AD_LOCAL_STOREFRONT_LINK":      40,
	"SHOWCASE_AD_ONLINE_PRODUCT_LINK":        42,
	"SHOWCASE_AD_LOCAL_PRODUCT_LINK":         43,
	"PROMOTION_EXTENSION":                    44,
	"SWIPEABLE_GALLERY_AD_HEADLINE":          45,
	"SWIPEABLE_GALLERY_AD_SWIPES":            46,
	"SWIPEABLE_GALLERY_AD_SEE_MORE":          47,
	"SWIPEABLE_GALLERY_AD_SITELINK_ONE":      48,
	"SWIPEABLE_GALLERY_AD_SITELINK_TWO":      49,
	"SWIPEABLE_GALLERY_AD_SITELINK_THREE":    50,
	"SWIPEABLE_GALLERY_AD_SITELINK_FOUR":     51,
	"SWIPEABLE_GALLERY_AD_SITELINK_FIVE":     52,
	"HOTEL_PRICE":                            53,
	"PRICE_EXTENSION":                        54,
	"HOTEL_BOOK_ON_GOOGLE_ROOM_SELECTION":    55,
	"SHOPPING_COMPARISON_LISTING":            56,
}

func (x ClickTypeEnum_ClickType) String() string {
	return proto.EnumName(ClickTypeEnum_ClickType_name, int32(x))
}

func (ClickTypeEnum_ClickType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ca334f107a4ef54e, []int{0, 0}
}

// Container for enumeration of Google Ads click types.
type ClickTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClickTypeEnum) Reset()         { *m = ClickTypeEnum{} }
func (m *ClickTypeEnum) String() string { return proto.CompactTextString(m) }
func (*ClickTypeEnum) ProtoMessage()    {}
func (*ClickTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca334f107a4ef54e, []int{0}
}

func (m *ClickTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClickTypeEnum.Unmarshal(m, b)
}
func (m *ClickTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClickTypeEnum.Marshal(b, m, deterministic)
}
func (m *ClickTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClickTypeEnum.Merge(m, src)
}
func (m *ClickTypeEnum) XXX_Size() int {
	return xxx_messageInfo_ClickTypeEnum.Size(m)
}
func (m *ClickTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ClickTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ClickTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.ClickTypeEnum_ClickType", ClickTypeEnum_ClickType_name, ClickTypeEnum_ClickType_value)
	proto.RegisterType((*ClickTypeEnum)(nil), "google.ads.googleads.v3.enums.ClickTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/click_type.proto", fileDescriptor_ca334f107a4ef54e)
}

var fileDescriptor_ca334f107a4ef54e = []byte{
	// 947 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x95, 0xdd, 0x52, 0x1b, 0x37,
	0x14, 0xc7, 0x1b, 0xd2, 0x24, 0x45, 0x7c, 0x9d, 0x08, 0x02, 0xe1, 0x3b, 0x40, 0x21, 0x6d, 0xda,
	0xae, 0xdb, 0xba, 0x5f, 0xe3, 0x5e, 0x69, 0x77, 0x8f, 0xd7, 0x1a, 0xb4, 0xd2, 0x8e, 0x24, 0xdb,
	0xa4, 0xc3, 0x8c, 0x86, 0x02, 0xe3, 0x61, 0x0a, 0x36, 0x13, 0x93, 0xcc, 0xe4, 0x3d, 0xfa, 0x04,
	0xbd, 0xe8, 0x45, 0x1f, 0xa5, 0xaf, 0xd1, 0xbb, 0x3e, 0x45, 0x47, 0xbb, 0x5e, 0x63, 0x88, 0x13,
	0x7a, 0xe3, 0x91, 0xf5, 0xff, 0xe9, 0xec, 0xf9, 0x92, 0x0e, 0x09, 0x3a, 0xbd, 0x5e, 0xe7, 0xfc,
	0xb4, 0x72, 0x74, 0xd2, 0xaf, 0x14, 0x4b, 0xbf, 0x7a, 0x53, 0xad, 0x9c, 0x76, 0x5f, 0x5f, 0xf4,
	0x2b, 0xc7, 0xe7, 0x67, 0xc7, 0xbf, 0xb9, 0xab, 0xb7, 0x97, 0xa7, 0xc1, 0xe5, 0xab, 0xde, 0x55,
	0x8f, 0xae, 0x17, 0x50, 0x70, 0x74, 0xd2, 0x0f, 0x86, 0x7c, 0xf0, 0xa6, 0x1a, 0xe4, 0xfc, 0xca,
	0x5a, 0x69, 0xee, 0xf2, 0xac, 0x72, 0xd4, 0xed, 0xf6, 0xae, 0x8e, 0xae, 0xce, 0x7a, 0xdd, 0x7e,
	0x71, 0x78, 0xfb, 0xcf, 0x69, 0x32, 0x13, 0x79, 0x8b, 0xf6, 0xed, 0xe5, 0x29, 0x76, 0x5f, 0x5f,
	0x6c, 0xff, 0x3e, 0x4d, 0x26, 0x87, 0x3b, 0x74, 0x8e, 0x4c, 0x35, 0xa5, 0xc9, 0x30, 0xe2, 0x75,
	0x8e, 0x31, 0x7c, 0x44, 0xa7, 0xc8, 0xa3, 0xa6, 0xdc, 0x97, 0xaa, 0x2d, 0xe1, 0x1e, 0x05, 0x32,
	0xcd, 0xb2, 0xcc, 0xc5, 0x88, 0x99, 0xe0, 0x72, 0x1f, 0x26, 0x3c, 0x1f, 0x6a, 0x64, 0x71, 0xa4,
	0x9b, 0x69, 0x68, 0xe0, 0x3e, 0xa5, 0x64, 0x36, 0xd4, 0x8a, 0xc5, 0x21, 0x93, 0xb1, 0xcb, 0x04,
	0x93, 0xf0, 0x31, 0x7d, 0x4c, 0x66, 0x22, 0x26, 0x84, 0xb3, 0x9a, 0x45, 0xfb, 0x5c, 0x26, 0xf0,
	0x80, 0x4e, 0x92, 0x07, 0x7e, 0xcb, 0xc0, 0x43, 0xba, 0x42, 0x16, 0x23, 0xc1, 0xa3, 0x7d, 0xa7,
	0xa4, 0x43, 0x99, 0xb0, 0x04, 0x53, 0x94, 0xd6, 0xb1, 0x18, 0x1e, 0x79, 0x6b, 0x09, 0x5a, 0x17,
	0x73, 0x8d, 0x91, 0xe5, 0x4a, 0x1a, 0xf8, 0x84, 0x2e, 0x12, 0x2a, 0x54, 0xc4, 0xfc, 0x5f, 0x87,
	0x07, 0x19, 0x93, 0x86, 0x2b, 0x09, 0x93, 0xf4, 0x29, 0x59, 0x18, 0xee, 0xd7, 0x95, 0x4e, 0x99,
	0x75, 0xfe, 0x13, 0x40, 0xe8, 0x06, 0x59, 0xb9, 0xad, 0x8c, 0x58, 0x9c, 0xa2, 0xcb, 0xe4, 0xc9,
	0x6d, 0x9d, 0xa7, 0x2c, 0x41, 0x98, 0xa6, 0xcf, 0xc8, 0xda, 0x6d, 0x49, 0x30, 0x19, 0x73, 0x99,
	0xb8, 0xcc, 0x13, 0x33, 0x74, 0x89, 0xcc, 0xdf, 0x26, 0x52, 0x96, 0xc1, 0xec, 0xb8, 0xaf, 0x1a,
	0xab, 0x34, 0x3a, 0x2e, 0xeb, 0x0a, 0xe6, 0xc6, 0xf9, 0x6b, 0xf1, 0xc0, 0x02, 0x78, 0x25, 0x55,
	0x21, 0x17, 0xe8, 0x6e, 0xa6, 0xed, 0xb1, 0x2f, 0x80, 0xaa, 0xd7, 0x51, 0xbb, 0x4c, 0x73, 0x69,
	0x0d, 0x50, 0x9f, 0x48, 0x65, 0x1b, 0xa8, 0x61, 0x9e, 0xae, 0x91, 0xa7, 0x99, 0x56, 0x71, 0x33,
	0xb2, 0x0e, 0x0f, 0x2c, 0xe6, 0x79, 0x71, 0x79, 0x6a, 0x0d, 0x2c, 0xd0, 0x75, 0xb2, 0x5c, 0xaa,
	0x82, 0x1b, 0xeb, 0x23, 0x60, 0x71, 0x29, 0x3f, 0xa1, 0x33, 0x64, 0xd2, 0x70, 0x8b, 0xbe, 0xac,
	0x06, 0x16, 0x7d, 0xc9, 0x0a, 0x67, 0x73, 0x17, 0x95, 0x86, 0x25, 0x3a, 0x4b, 0x48, 0x53, 0x8b,
	0xf2, 0xc4, 0xb2, 0xaf, 0x5b, 0x8b, 0xc7, 0xa8, 0x9c, 0x6f, 0x89, 0x02, 0x1e, 0x68, 0x2b, 0x74,
	0x93, 0xac, 0x16, 0x5a, 0x11, 0x80, 0x72, 0x2c, 0x4f, 0x76, 0x09, 0xac, 0xd2, 0x5d, 0xb2, 0x55,
	0x02, 0x3a, 0x2e, 0xd5, 0x06, 0xb2, 0x58, 0x70, 0x39, 0xb4, 0xb3, 0xe6, 0x33, 0x51, 0x60, 0x28,
	0x63, 0x17, 0xb1, 0xac, 0x54, 0xd6, 0xaf, 0x95, 0x36, 0x86, 0xde, 0xf1, 0x52, 0xd9, 0xa0, 0x0b,
	0x04, 0x5a, 0xdc, 0x34, 0x99, 0x70, 0xd7, 0x01, 0x6d, 0xfa, 0x80, 0xda, 0x5c, 0xa3, 0x40, 0x63,
	0x8a, 0xb6, 0x7c, 0x36, 0x9a, 0xaf, 0x91, 0x8c, 0xf8, 0x80, 0x05, 0x6c, 0xd1, 0xcf, 0xc9, 0xee,
	0x18, 0x35, 0x6d, 0x0a, 0xcb, 0xa3, 0x06, 0x93, 0x12, 0xc5, 0x00, 0xdd, 0xa6, 0x2f, 0xc8, 0xde,
	0x5d, 0xa8, 0x92, 0x3e, 0x30, 0xd8, 0xf1, 0x5d, 0xf1, 0x2e, 0x6b, 0x5c, 0xa4, 0x9a, 0x99, 0x92,
	0xf0, 0x29, 0xdd, 0x21, 0x9b, 0x63, 0x6c, 0x59, 0xcd, 0xa4, 0x61, 0x91, 0x65, 0xa1, 0x40, 0xd8,
	0xa5, 0xab, 0x64, 0xa9, 0x84, 0x58, 0xec, 0x6e, 0x5c, 0xc9, 0x3d, 0x5f, 0x68, 0xd3, 0x50, 0xed,
	0x88, 0x19, 0xcc, 0x2b, 0xcc, 0x2c, 0x26, 0x4a, 0xbf, 0x74, 0xb9, 0xfc, 0xdc, 0x67, 0x7e, 0x54,
	0xce, 0x63, 0x28, 0xca, 0x57, 0xd7, 0x4a, 0xda, 0x02, 0xfb, 0xcc, 0xfb, 0x31, 0x8a, 0x15, 0xfe,
	0xbb, 0x6b, 0xd7, 0xe4, 0x3e, 0xbc, 0xa0, 0xdb, 0x64, 0xe3, 0x5d, 0x5b, 0x37, 0x98, 0x2f, 0xfc,
	0xfd, 0xc8, 0xb4, 0x4a, 0xd5, 0xe0, 0xbe, 0x0e, 0xfa, 0x12, 0xbe, 0xa4, 0x5b, 0x64, 0xdd, 0xb4,
	0x79, 0x86, 0x3e, 0x26, 0x97, 0x30, 0x21, 0x50, 0xbf, 0xf4, 0x56, 0xca, 0x2e, 0x80, 0xaf, 0x7c,
	0x1b, 0x8d, 0x45, 0xf2, 0x4d, 0x03, 0xc1, 0x7b, 0x6d, 0x18, 0x44, 0x97, 0x2a, 0x8d, 0x50, 0xc9,
	0xe3, 0x1d, 0x8b, 0x0c, 0x9a, 0xc3, 0x29, 0x89, 0xf0, 0xf5, 0xdd, 0x98, 0x6d, 0x2b, 0xf8, 0x86,
	0x3e, 0x27, 0x3b, 0x77, 0x60, 0x0d, 0x8d, 0x08, 0xdf, 0xd2, 0x3d, 0xb2, 0xfd, 0x61, 0xb0, 0xae,
	0x9a, 0x1a, 0xaa, 0xff, 0x83, 0xe3, 0x2d, 0x84, 0xef, 0xfc, 0x43, 0xdb, 0x50, 0x16, 0x7d, 0x7a,
	0x79, 0x84, 0xf0, 0x3d, 0x9d, 0x27, 0x73, 0xf9, 0x72, 0x24, 0xa7, 0x3f, 0x78, 0xf7, 0x0a, 0x2a,
	0x54, 0x2a, 0x7f, 0x50, 0x13, 0xa5, 0x12, 0x81, 0x4e, 0x2b, 0x95, 0x3a, 0x83, 0xa2, 0x78, 0xf3,
	0xe0, 0xc7, 0x3c, 0xb3, 0x0d, 0x95, 0x65, 0xbe, 0xbf, 0x22, 0x95, 0x66, 0x4c, 0x73, 0xa3, 0x64,
	0xd9, 0x72, 0xf0, 0x53, 0xf8, 0xcf, 0x3d, 0xb2, 0x75, 0xdc, 0xbb, 0x08, 0x3e, 0x38, 0x6c, 0xc2,
	0xd9, 0xe1, 0xe4, 0xc8, 0xfc, 0x78, 0xc9, 0xee, 0xfd, 0x12, 0x0e, 0x0e, 0x74, 0x7a, 0xe7, 0x47,
	0xdd, 0x4e, 0xd0, 0x7b, 0xd5, 0xa9, 0x74, 0x4e, 0xbb, 0xf9, 0xf0, 0x29, 0xa7, 0xdb, 0xe5, 0x59,
	0xff, 0x3d, 0xc3, 0xee, 0xe7, 0xfc, 0xf7, 0x8f, 0x89, 0xfb, 0x09, 0x63, 0x7f, 0x4d, 0xac, 0x27,
	0x85, 0x29, 0x76, 0xd2, 0x0f, 0x8a, 0xa5, 0x5f, 0xb5, 0xaa, 0x81, 0x9f, 0x5b, 0xfd, 0xbf, 0x4b,
	0xfd, 0x90, 0x9d, 0xf4, 0x0f, 0x87, 0xfa, 0x61, 0xab, 0x7a, 0x98, 0xeb, 0xff, 0x4e, 0x6c, 0x15,
	0x9b, 0xb5, 0x1a, 0x3b, 0xe9, 0xd7, 0x6a, 0x43, 0xa2, 0x56, 0x6b, 0x55, 0x6b, 0xb5, 0x9c, 0xf9,
	0xf5, 0x61, 0xee, 0x58, 0xf5, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x11, 0x88, 0xfc, 0xd8, 0x84,
	0x07, 0x00, 0x00,
}
