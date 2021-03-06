// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/shared_criterion.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v2/common"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v2/enums"
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

// A criterion belonging to a shared set.
type SharedCriterion struct {
	// The resource name of the shared criterion.
	// Shared set resource names have the form:
	//
	// `customers/{customer_id}/sharedCriteria/{shared_set_id}~{criterion_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The shared set to which the shared criterion belongs.
	SharedSet *wrappers.StringValue `protobuf:"bytes,2,opt,name=shared_set,json=sharedSet,proto3" json:"shared_set,omitempty"`
	// The ID of the criterion.
	//
	// This field is ignored for mutates.
	CriterionId *wrappers.Int64Value `protobuf:"bytes,26,opt,name=criterion_id,json=criterionId,proto3" json:"criterion_id,omitempty"`
	// The type of the criterion.
	Type enums.CriterionTypeEnum_CriterionType `protobuf:"varint,4,opt,name=type,proto3,enum=google.ads.googleads.v2.enums.CriterionTypeEnum_CriterionType" json:"type,omitempty"`
	// The criterion.
	//
	// Exactly one must be set.
	//
	// Types that are valid to be assigned to Criterion:
	//	*SharedCriterion_Keyword
	//	*SharedCriterion_YoutubeVideo
	//	*SharedCriterion_YoutubeChannel
	//	*SharedCriterion_Placement
	//	*SharedCriterion_MobileAppCategory
	//	*SharedCriterion_MobileApplication
	Criterion            isSharedCriterion_Criterion `protobuf_oneof:"criterion"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *SharedCriterion) Reset()         { *m = SharedCriterion{} }
func (m *SharedCriterion) String() string { return proto.CompactTextString(m) }
func (*SharedCriterion) ProtoMessage()    {}
func (*SharedCriterion) Descriptor() ([]byte, []int) {
	return fileDescriptor_1041cb67d99aea67, []int{0}
}

func (m *SharedCriterion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SharedCriterion.Unmarshal(m, b)
}
func (m *SharedCriterion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SharedCriterion.Marshal(b, m, deterministic)
}
func (m *SharedCriterion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SharedCriterion.Merge(m, src)
}
func (m *SharedCriterion) XXX_Size() int {
	return xxx_messageInfo_SharedCriterion.Size(m)
}
func (m *SharedCriterion) XXX_DiscardUnknown() {
	xxx_messageInfo_SharedCriterion.DiscardUnknown(m)
}

var xxx_messageInfo_SharedCriterion proto.InternalMessageInfo

func (m *SharedCriterion) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *SharedCriterion) GetSharedSet() *wrappers.StringValue {
	if m != nil {
		return m.SharedSet
	}
	return nil
}

func (m *SharedCriterion) GetCriterionId() *wrappers.Int64Value {
	if m != nil {
		return m.CriterionId
	}
	return nil
}

func (m *SharedCriterion) GetType() enums.CriterionTypeEnum_CriterionType {
	if m != nil {
		return m.Type
	}
	return enums.CriterionTypeEnum_UNSPECIFIED
}

type isSharedCriterion_Criterion interface {
	isSharedCriterion_Criterion()
}

type SharedCriterion_Keyword struct {
	Keyword *common.KeywordInfo `protobuf:"bytes,3,opt,name=keyword,proto3,oneof"`
}

type SharedCriterion_YoutubeVideo struct {
	YoutubeVideo *common.YouTubeVideoInfo `protobuf:"bytes,5,opt,name=youtube_video,json=youtubeVideo,proto3,oneof"`
}

type SharedCriterion_YoutubeChannel struct {
	YoutubeChannel *common.YouTubeChannelInfo `protobuf:"bytes,6,opt,name=youtube_channel,json=youtubeChannel,proto3,oneof"`
}

type SharedCriterion_Placement struct {
	Placement *common.PlacementInfo `protobuf:"bytes,7,opt,name=placement,proto3,oneof"`
}

type SharedCriterion_MobileAppCategory struct {
	MobileAppCategory *common.MobileAppCategoryInfo `protobuf:"bytes,8,opt,name=mobile_app_category,json=mobileAppCategory,proto3,oneof"`
}

type SharedCriterion_MobileApplication struct {
	MobileApplication *common.MobileApplicationInfo `protobuf:"bytes,9,opt,name=mobile_application,json=mobileApplication,proto3,oneof"`
}

func (*SharedCriterion_Keyword) isSharedCriterion_Criterion() {}

func (*SharedCriterion_YoutubeVideo) isSharedCriterion_Criterion() {}

func (*SharedCriterion_YoutubeChannel) isSharedCriterion_Criterion() {}

func (*SharedCriterion_Placement) isSharedCriterion_Criterion() {}

func (*SharedCriterion_MobileAppCategory) isSharedCriterion_Criterion() {}

func (*SharedCriterion_MobileApplication) isSharedCriterion_Criterion() {}

func (m *SharedCriterion) GetCriterion() isSharedCriterion_Criterion {
	if m != nil {
		return m.Criterion
	}
	return nil
}

func (m *SharedCriterion) GetKeyword() *common.KeywordInfo {
	if x, ok := m.GetCriterion().(*SharedCriterion_Keyword); ok {
		return x.Keyword
	}
	return nil
}

func (m *SharedCriterion) GetYoutubeVideo() *common.YouTubeVideoInfo {
	if x, ok := m.GetCriterion().(*SharedCriterion_YoutubeVideo); ok {
		return x.YoutubeVideo
	}
	return nil
}

func (m *SharedCriterion) GetYoutubeChannel() *common.YouTubeChannelInfo {
	if x, ok := m.GetCriterion().(*SharedCriterion_YoutubeChannel); ok {
		return x.YoutubeChannel
	}
	return nil
}

func (m *SharedCriterion) GetPlacement() *common.PlacementInfo {
	if x, ok := m.GetCriterion().(*SharedCriterion_Placement); ok {
		return x.Placement
	}
	return nil
}

func (m *SharedCriterion) GetMobileAppCategory() *common.MobileAppCategoryInfo {
	if x, ok := m.GetCriterion().(*SharedCriterion_MobileAppCategory); ok {
		return x.MobileAppCategory
	}
	return nil
}

func (m *SharedCriterion) GetMobileApplication() *common.MobileApplicationInfo {
	if x, ok := m.GetCriterion().(*SharedCriterion_MobileApplication); ok {
		return x.MobileApplication
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SharedCriterion) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SharedCriterion_Keyword)(nil),
		(*SharedCriterion_YoutubeVideo)(nil),
		(*SharedCriterion_YoutubeChannel)(nil),
		(*SharedCriterion_Placement)(nil),
		(*SharedCriterion_MobileAppCategory)(nil),
		(*SharedCriterion_MobileApplication)(nil),
	}
}

func init() {
	proto.RegisterType((*SharedCriterion)(nil), "google.ads.googleads.v2.resources.SharedCriterion")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/shared_criterion.proto", fileDescriptor_1041cb67d99aea67)
}

var fileDescriptor_1041cb67d99aea67 = []byte{
	// 591 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xdb, 0x6a, 0xd4, 0x40,
	0x18, 0xc7, 0xdd, 0x6d, 0x6d, 0xdd, 0xe9, 0x09, 0x47, 0x2f, 0x42, 0x2d, 0xd2, 0x2a, 0x85, 0x82,
	0x74, 0x22, 0xf1, 0x80, 0xa4, 0x50, 0x48, 0x8b, 0xd4, 0x2a, 0x95, 0x92, 0x96, 0x15, 0x65, 0x25,
	0xcc, 0x26, 0x5f, 0xd3, 0x60, 0x32, 0x33, 0xcc, 0x4c, 0x5a, 0xf6, 0xd2, 0x57, 0xf1, 0xd2, 0x7b,
	0x5f, 0xc2, 0x47, 0xf1, 0x29, 0x64, 0x27, 0x33, 0x29, 0xb6, 0xac, 0xab, 0x77, 0x5f, 0xe6, 0xfb,
	0xff, 0x7e, 0x73, 0xc8, 0x24, 0xe8, 0x55, 0xce, 0x79, 0x5e, 0x82, 0x4f, 0x33, 0xe5, 0x37, 0xe5,
	0xb8, 0xba, 0x08, 0x7c, 0x09, 0x8a, 0xd7, 0x32, 0x05, 0xe5, 0xab, 0x73, 0x2a, 0x21, 0x4b, 0x52,
	0x59, 0x68, 0x90, 0x05, 0x67, 0x44, 0x48, 0xae, 0x39, 0xde, 0x68, 0xe2, 0x84, 0x66, 0x8a, 0xb4,
	0x24, 0xb9, 0x08, 0x48, 0x4b, 0xae, 0x6e, 0x4f, 0x92, 0xa7, 0xbc, 0xaa, 0x38, 0xf3, 0xad, 0x92,
	0x36, 0xc6, 0xd5, 0x60, 0x52, 0x1c, 0x58, 0x5d, 0x29, 0xbf, 0x5d, 0x40, 0xa2, 0x47, 0x02, 0x2c,
	0xf3, 0xd0, 0x32, 0xe6, 0x69, 0x58, 0x9f, 0xf9, 0x97, 0x92, 0x0a, 0x01, 0x52, 0xd9, 0xfe, 0x9a,
	0x73, 0x8a, 0xc2, 0xa7, 0x8c, 0x71, 0x4d, 0x75, 0xc1, 0x99, 0xed, 0x3e, 0xfa, 0x31, 0x87, 0x56,
	0x4e, 0xcc, 0xf6, 0xf6, 0x9d, 0x1c, 0x3f, 0x46, 0x4b, 0x6e, 0x07, 0x09, 0xa3, 0x15, 0x78, 0x9d,
	0xf5, 0xce, 0x56, 0x2f, 0x5e, 0x74, 0x83, 0xef, 0x69, 0x05, 0x78, 0x07, 0x21, 0x7b, 0x2c, 0x0a,
	0xb4, 0xd7, 0x5d, 0xef, 0x6c, 0x2d, 0x04, 0x6b, 0xf6, 0x18, 0x88, 0x5b, 0x0b, 0x39, 0xd1, 0xb2,
	0x60, 0x79, 0x9f, 0x96, 0x35, 0xc4, 0xbd, 0x26, 0x7f, 0x02, 0x1a, 0xef, 0xa2, 0xc5, 0xab, 0xbd,
	0x14, 0x99, 0xb7, 0x6a, 0xf0, 0x07, 0x37, 0xf0, 0x43, 0xa6, 0x5f, 0x3e, 0x6f, 0xe8, 0x85, 0x16,
	0x38, 0xcc, 0x70, 0x8c, 0x66, 0xc7, 0x27, 0xe0, 0xcd, 0xae, 0x77, 0xb6, 0x96, 0x83, 0x5d, 0x32,
	0xe9, 0x45, 0x98, 0x63, 0x23, 0xed, 0xce, 0x4e, 0x47, 0x02, 0x5e, 0xb3, 0xba, 0xfa, 0x73, 0x24,
	0x36, 0x2e, 0x7c, 0x80, 0xe6, 0xbf, 0xc0, 0xe8, 0x92, 0xcb, 0xcc, 0x9b, 0x31, 0xcb, 0x79, 0x32,
	0x51, 0xdb, 0xbc, 0x3c, 0xf2, 0xae, 0x89, 0x1f, 0xb2, 0x33, 0xfe, 0xe6, 0x56, 0xec, 0x68, 0xfc,
	0x01, 0x2d, 0x8d, 0x78, 0xad, 0xeb, 0x21, 0x24, 0x17, 0x45, 0x06, 0xdc, 0xbb, 0x6d, 0x74, 0x4f,
	0xa7, 0xe9, 0x3e, 0xf2, 0xfa, 0xb4, 0x1e, 0x42, 0x7f, 0xcc, 0x58, 0xe7, 0xa2, 0x15, 0x99, 0x31,
	0xfc, 0x19, 0xad, 0x38, 0x71, 0x7a, 0x4e, 0x19, 0x83, 0xd2, 0x9b, 0x33, 0xea, 0xe0, 0x1f, 0xd5,
	0xfb, 0x0d, 0x65, 0xe5, 0xcb, 0x56, 0x66, 0x47, 0xf1, 0x11, 0xea, 0x89, 0x92, 0xa6, 0x50, 0x01,
	0xd3, 0xde, 0xbc, 0x11, 0x6f, 0x4f, 0x13, 0x1f, 0x3b, 0xc0, 0x3a, 0xaf, 0x0c, 0x38, 0x47, 0xf7,
	0x2a, 0x3e, 0x2c, 0x4a, 0x48, 0xa8, 0x10, 0x49, 0x4a, 0x35, 0xe4, 0x5c, 0x8e, 0xbc, 0x3b, 0x46,
	0xfc, 0x62, 0x9a, 0xf8, 0xc8, 0xa0, 0x91, 0x10, 0xfb, 0x16, 0xb4, 0x13, 0xdc, 0xad, 0xae, 0x37,
	0xf0, 0x19, 0xc2, 0x57, 0x13, 0x95, 0x45, 0x6a, 0xee, 0xb7, 0xd7, 0xfb, 0xcf, 0x79, 0x1c, 0x78,
	0x63, 0x1e, 0xd7, 0xd8, 0x5b, 0x40, 0xbd, 0xf6, 0x0e, 0xee, 0x7d, 0xed, 0xa2, 0xcd, 0x94, 0x57,
	0x64, 0xea, 0x2f, 0x60, 0xef, 0xfe, 0xb5, 0xcf, 0xeb, 0x78, 0x7c, 0xb9, 0x8f, 0x3b, 0x9f, 0xde,
	0x5a, 0x34, 0xe7, 0x25, 0x65, 0x39, 0xe1, 0x32, 0xf7, 0x73, 0x60, 0xe6, 0xea, 0xbb, 0x6f, 0x5f,
	0x14, 0xea, 0x2f, 0xbf, 0xa5, 0x9d, 0xb6, 0xfa, 0xd6, 0x9d, 0x39, 0x88, 0xa2, 0xef, 0xdd, 0x8d,
	0x83, 0x46, 0x19, 0x65, 0x8a, 0x34, 0xe5, 0xb8, 0xea, 0x07, 0x24, 0x76, 0xc9, 0x9f, 0x2e, 0x33,
	0x88, 0x32, 0x35, 0x68, 0x33, 0x83, 0x7e, 0x30, 0x68, 0x33, 0xbf, 0xba, 0x9b, 0x4d, 0x23, 0x0c,
	0xa3, 0x4c, 0x85, 0x61, 0x9b, 0x0a, 0xc3, 0x7e, 0x10, 0x86, 0x6d, 0x6e, 0x38, 0x67, 0x16, 0xfb,
	0xec, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf1, 0x88, 0x96, 0x3b, 0x42, 0x05, 0x00, 0x00,
}
