// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/retry/omit_host_metadata/v2/omit_host_metadata_config.proto

package envoy_config_retry_omit_host_metadata_v2

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/datawire/ambassador/pkg/api/envoy/api/v2/core"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type OmitHostMetadataConfig struct {
	MetadataMatch        *core.Metadata `protobuf:"bytes,1,opt,name=metadata_match,json=metadataMatch,proto3" json:"metadata_match,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *OmitHostMetadataConfig) Reset()         { *m = OmitHostMetadataConfig{} }
func (m *OmitHostMetadataConfig) String() string { return proto.CompactTextString(m) }
func (*OmitHostMetadataConfig) ProtoMessage()    {}
func (*OmitHostMetadataConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_94b8638ad6beee65, []int{0}
}

func (m *OmitHostMetadataConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OmitHostMetadataConfig.Unmarshal(m, b)
}
func (m *OmitHostMetadataConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OmitHostMetadataConfig.Marshal(b, m, deterministic)
}
func (m *OmitHostMetadataConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OmitHostMetadataConfig.Merge(m, src)
}
func (m *OmitHostMetadataConfig) XXX_Size() int {
	return xxx_messageInfo_OmitHostMetadataConfig.Size(m)
}
func (m *OmitHostMetadataConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_OmitHostMetadataConfig.DiscardUnknown(m)
}

var xxx_messageInfo_OmitHostMetadataConfig proto.InternalMessageInfo

func (m *OmitHostMetadataConfig) GetMetadataMatch() *core.Metadata {
	if m != nil {
		return m.MetadataMatch
	}
	return nil
}

func init() {
	proto.RegisterType((*OmitHostMetadataConfig)(nil), "envoy.config.retry.omit_host_metadata.v2.OmitHostMetadataConfig")
}

func init() {
	proto.RegisterFile("envoy/config/retry/omit_host_metadata/v2/omit_host_metadata_config.proto", fileDescriptor_94b8638ad6beee65)
}

var fileDescriptor_94b8638ad6beee65 = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x31, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0x65, 0x24, 0x3a, 0x04, 0x81, 0x50, 0x06, 0x84, 0x5a, 0x40, 0x88, 0xa9, 0x93, 0x2d,
	0x12, 0x89, 0x03, 0x84, 0xa5, 0x4b, 0x45, 0xc5, 0x8c, 0x14, 0xbd, 0xa6, 0xa6, 0xf5, 0x60, 0x3f,
	0xcb, 0x7e, 0x8d, 0x9a, 0x8d, 0x1b, 0xb0, 0xf6, 0x2c, 0x9c, 0x80, 0x95, 0xab, 0x70, 0x00, 0x84,
	0x1c, 0x27, 0x2c, 0xcd, 0xc0, 0xea, 0xa7, 0xef, 0xff, 0xff, 0xcf, 0xc9, 0x4c, 0x9a, 0x1a, 0x1b,
	0x51, 0xa1, 0x79, 0x55, 0x6b, 0xe1, 0x24, 0xb9, 0x46, 0xa0, 0x56, 0x54, 0x6e, 0xd0, 0x53, 0xa9,
	0x25, 0xc1, 0x0a, 0x08, 0x44, 0x9d, 0x0d, 0xbc, 0x96, 0x91, 0xe2, 0xd6, 0x21, 0x61, 0x3a, 0x6d,
	0x93, 0x78, 0xf7, 0xd6, 0x26, 0xf1, 0x43, 0x86, 0xd7, 0xd9, 0xf8, 0x2a, 0x76, 0x82, 0x55, 0x21,
	0xb7, 0x42, 0x27, 0xc5, 0x12, 0xbc, 0x8c, 0x39, 0xe3, 0x9b, 0xed, 0xca, 0x82, 0x00, 0x63, 0x90,
	0x80, 0x14, 0x1a, 0x2f, 0xb4, 0x5a, 0x3b, 0xa0, 0xfe, 0x7e, 0x7d, 0x70, 0xf7, 0x04, 0xb4, 0xf5,
	0xf1, 0x7c, 0xf7, 0x92, 0x5c, 0x3c, 0x69, 0x45, 0x33, 0xf4, 0x34, 0xef, 0x3a, 0x1f, 0xdb, 0x49,
	0x69, 0x91, 0x9c, 0xfd, 0x2d, 0xd7, 0x40, 0xd5, 0xe6, 0x92, 0xdd, 0xb2, 0xe9, 0x49, 0x36, 0xe1,
	0x71, 0x39, 0x58, 0xc5, 0xeb, 0x8c, 0x87, 0x3d, 0xbc, 0x47, 0x9f, 0x4f, 0x7b, 0x64, 0x1e, 0x88,
	0x62, 0xcf, 0xbe, 0xf7, 0x3f, 0xef, 0xc7, 0x79, 0x7a, 0x1f, 0x19, 0xb9, 0x23, 0x69, 0x7c, 0x58,
	0xd1, 0x19, 0x07, 0xd9, 0x41, 0xed, 0xfc, 0xe3, 0xed, 0xf3, 0x6b, 0x74, 0x74, 0xce, 0x92, 0x07,
	0x85, 0xb1, 0xd1, 0x3a, 0xdc, 0x35, 0xfc, 0xbf, 0xdf, 0x56, 0x4c, 0x86, 0xbd, 0x16, 0x41, 0x7b,
	0xc1, 0x96, 0xa3, 0xd6, 0x3f, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x3f, 0xbf, 0x66, 0xd2,
	0x01, 0x00, 0x00,
}