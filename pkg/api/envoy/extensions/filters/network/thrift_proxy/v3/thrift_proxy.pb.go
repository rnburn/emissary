// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/network/thrift_proxy/v3/thrift_proxy.proto

package envoy_extensions_filters_network_thrift_proxy_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type TransportType int32

const (
	TransportType_AUTO_TRANSPORT TransportType = 0
	TransportType_FRAMED         TransportType = 1
	TransportType_UNFRAMED       TransportType = 2
	TransportType_HEADER         TransportType = 3
)

var TransportType_name = map[int32]string{
	0: "AUTO_TRANSPORT",
	1: "FRAMED",
	2: "UNFRAMED",
	3: "HEADER",
}

var TransportType_value = map[string]int32{
	"AUTO_TRANSPORT": 0,
	"FRAMED":         1,
	"UNFRAMED":       2,
	"HEADER":         3,
}

func (x TransportType) String() string {
	return proto.EnumName(TransportType_name, int32(x))
}

func (TransportType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a5befaf4ce0c1223, []int{0}
}

type ProtocolType int32

const (
	ProtocolType_AUTO_PROTOCOL ProtocolType = 0
	ProtocolType_BINARY        ProtocolType = 1
	ProtocolType_LAX_BINARY    ProtocolType = 2
	ProtocolType_COMPACT       ProtocolType = 3
	ProtocolType_TWITTER       ProtocolType = 4
)

var ProtocolType_name = map[int32]string{
	0: "AUTO_PROTOCOL",
	1: "BINARY",
	2: "LAX_BINARY",
	3: "COMPACT",
	4: "TWITTER",
}

var ProtocolType_value = map[string]int32{
	"AUTO_PROTOCOL": 0,
	"BINARY":        1,
	"LAX_BINARY":    2,
	"COMPACT":       3,
	"TWITTER":       4,
}

func (x ProtocolType) String() string {
	return proto.EnumName(ProtocolType_name, int32(x))
}

func (ProtocolType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a5befaf4ce0c1223, []int{1}
}

type ThriftProxy struct {
	Transport            TransportType       `protobuf:"varint,2,opt,name=transport,proto3,enum=envoy.extensions.filters.network.thrift_proxy.v3.TransportType" json:"transport,omitempty"`
	Protocol             ProtocolType        `protobuf:"varint,3,opt,name=protocol,proto3,enum=envoy.extensions.filters.network.thrift_proxy.v3.ProtocolType" json:"protocol,omitempty"`
	StatPrefix           string              `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	RouteConfig          *RouteConfiguration `protobuf:"bytes,4,opt,name=route_config,json=routeConfig,proto3" json:"route_config,omitempty"`
	ThriftFilters        []*ThriftFilter     `protobuf:"bytes,5,rep,name=thrift_filters,json=thriftFilters,proto3" json:"thrift_filters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ThriftProxy) Reset()         { *m = ThriftProxy{} }
func (m *ThriftProxy) String() string { return proto.CompactTextString(m) }
func (*ThriftProxy) ProtoMessage()    {}
func (*ThriftProxy) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5befaf4ce0c1223, []int{0}
}

func (m *ThriftProxy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThriftProxy.Unmarshal(m, b)
}
func (m *ThriftProxy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThriftProxy.Marshal(b, m, deterministic)
}
func (m *ThriftProxy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThriftProxy.Merge(m, src)
}
func (m *ThriftProxy) XXX_Size() int {
	return xxx_messageInfo_ThriftProxy.Size(m)
}
func (m *ThriftProxy) XXX_DiscardUnknown() {
	xxx_messageInfo_ThriftProxy.DiscardUnknown(m)
}

var xxx_messageInfo_ThriftProxy proto.InternalMessageInfo

func (m *ThriftProxy) GetTransport() TransportType {
	if m != nil {
		return m.Transport
	}
	return TransportType_AUTO_TRANSPORT
}

func (m *ThriftProxy) GetProtocol() ProtocolType {
	if m != nil {
		return m.Protocol
	}
	return ProtocolType_AUTO_PROTOCOL
}

func (m *ThriftProxy) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *ThriftProxy) GetRouteConfig() *RouteConfiguration {
	if m != nil {
		return m.RouteConfig
	}
	return nil
}

func (m *ThriftProxy) GetThriftFilters() []*ThriftFilter {
	if m != nil {
		return m.ThriftFilters
	}
	return nil
}

type ThriftFilter struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to ConfigType:
	//	*ThriftFilter_TypedConfig
	//	*ThriftFilter_HiddenEnvoyDeprecatedConfig
	ConfigType           isThriftFilter_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ThriftFilter) Reset()         { *m = ThriftFilter{} }
func (m *ThriftFilter) String() string { return proto.CompactTextString(m) }
func (*ThriftFilter) ProtoMessage()    {}
func (*ThriftFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5befaf4ce0c1223, []int{1}
}

func (m *ThriftFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThriftFilter.Unmarshal(m, b)
}
func (m *ThriftFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThriftFilter.Marshal(b, m, deterministic)
}
func (m *ThriftFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThriftFilter.Merge(m, src)
}
func (m *ThriftFilter) XXX_Size() int {
	return xxx_messageInfo_ThriftFilter.Size(m)
}
func (m *ThriftFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_ThriftFilter.DiscardUnknown(m)
}

var xxx_messageInfo_ThriftFilter proto.InternalMessageInfo

func (m *ThriftFilter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isThriftFilter_ConfigType interface {
	isThriftFilter_ConfigType()
}

type ThriftFilter_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,3,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

type ThriftFilter_HiddenEnvoyDeprecatedConfig struct {
	HiddenEnvoyDeprecatedConfig *_struct.Struct `protobuf:"bytes,2,opt,name=hidden_envoy_deprecated_config,json=hiddenEnvoyDeprecatedConfig,proto3,oneof"`
}

func (*ThriftFilter_TypedConfig) isThriftFilter_ConfigType() {}

func (*ThriftFilter_HiddenEnvoyDeprecatedConfig) isThriftFilter_ConfigType() {}

func (m *ThriftFilter) GetConfigType() isThriftFilter_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (m *ThriftFilter) GetTypedConfig() *any.Any {
	if x, ok := m.GetConfigType().(*ThriftFilter_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// Deprecated: Do not use.
func (m *ThriftFilter) GetHiddenEnvoyDeprecatedConfig() *_struct.Struct {
	if x, ok := m.GetConfigType().(*ThriftFilter_HiddenEnvoyDeprecatedConfig); ok {
		return x.HiddenEnvoyDeprecatedConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ThriftFilter) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ThriftFilter_TypedConfig)(nil),
		(*ThriftFilter_HiddenEnvoyDeprecatedConfig)(nil),
	}
}

type ThriftProtocolOptions struct {
	Transport            TransportType `protobuf:"varint,1,opt,name=transport,proto3,enum=envoy.extensions.filters.network.thrift_proxy.v3.TransportType" json:"transport,omitempty"`
	Protocol             ProtocolType  `protobuf:"varint,2,opt,name=protocol,proto3,enum=envoy.extensions.filters.network.thrift_proxy.v3.ProtocolType" json:"protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ThriftProtocolOptions) Reset()         { *m = ThriftProtocolOptions{} }
func (m *ThriftProtocolOptions) String() string { return proto.CompactTextString(m) }
func (*ThriftProtocolOptions) ProtoMessage()    {}
func (*ThriftProtocolOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5befaf4ce0c1223, []int{2}
}

func (m *ThriftProtocolOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThriftProtocolOptions.Unmarshal(m, b)
}
func (m *ThriftProtocolOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThriftProtocolOptions.Marshal(b, m, deterministic)
}
func (m *ThriftProtocolOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThriftProtocolOptions.Merge(m, src)
}
func (m *ThriftProtocolOptions) XXX_Size() int {
	return xxx_messageInfo_ThriftProtocolOptions.Size(m)
}
func (m *ThriftProtocolOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_ThriftProtocolOptions.DiscardUnknown(m)
}

var xxx_messageInfo_ThriftProtocolOptions proto.InternalMessageInfo

func (m *ThriftProtocolOptions) GetTransport() TransportType {
	if m != nil {
		return m.Transport
	}
	return TransportType_AUTO_TRANSPORT
}

func (m *ThriftProtocolOptions) GetProtocol() ProtocolType {
	if m != nil {
		return m.Protocol
	}
	return ProtocolType_AUTO_PROTOCOL
}

func init() {
	proto.RegisterEnum("envoy.extensions.filters.network.thrift_proxy.v3.TransportType", TransportType_name, TransportType_value)
	proto.RegisterEnum("envoy.extensions.filters.network.thrift_proxy.v3.ProtocolType", ProtocolType_name, ProtocolType_value)
	proto.RegisterType((*ThriftProxy)(nil), "envoy.extensions.filters.network.thrift_proxy.v3.ThriftProxy")
	proto.RegisterType((*ThriftFilter)(nil), "envoy.extensions.filters.network.thrift_proxy.v3.ThriftFilter")
	proto.RegisterType((*ThriftProtocolOptions)(nil), "envoy.extensions.filters.network.thrift_proxy.v3.ThriftProtocolOptions")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/network/thrift_proxy/v3/thrift_proxy.proto", fileDescriptor_a5befaf4ce0c1223)
}

var fileDescriptor_a5befaf4ce0c1223 = []byte{
	// 684 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x6b, 0x27, 0xfd, 0x37, 0x4e, 0x22, 0xb3, 0x02, 0x35, 0xb4, 0x50, 0x85, 0x9e, 0xa2,
	0x1e, 0x6c, 0x68, 0x4f, 0x54, 0xd0, 0x62, 0x27, 0x29, 0x8d, 0xd4, 0xc6, 0xd6, 0xd6, 0x55, 0x41,
	0x42, 0xb2, 0xdc, 0x78, 0x93, 0x5a, 0x04, 0xaf, 0xb5, 0xde, 0x84, 0xe6, 0x86, 0x38, 0xf1, 0x0c,
	0x88, 0x03, 0x4f, 0xc1, 0x81, 0x3b, 0x12, 0x57, 0x5e, 0xa7, 0x27, 0xe4, 0xb5, 0xd3, 0x24, 0xad,
	0x38, 0xa4, 0x08, 0x6e, 0x1e, 0xcf, 0xee, 0x6f, 0xbe, 0xf9, 0x76, 0x06, 0x6a, 0x24, 0x1c, 0xd0,
	0xa1, 0x4e, 0x2e, 0x38, 0x09, 0xe3, 0x80, 0x86, 0xb1, 0xde, 0x09, 0x7a, 0x9c, 0xb0, 0x58, 0x0f,
	0x09, 0x7f, 0x4f, 0xd9, 0x5b, 0x9d, 0x9f, 0xb3, 0xa0, 0xc3, 0xdd, 0x88, 0xd1, 0x8b, 0xa1, 0x3e,
	0xd8, 0x9e, 0x8a, 0xb5, 0x88, 0x51, 0x4e, 0xd1, 0x63, 0x01, 0xd1, 0xc6, 0x10, 0x2d, 0x83, 0x68,
	0x19, 0x44, 0x9b, 0xba, 0x34, 0xd8, 0x5e, 0x7d, 0x36, 0x73, 0x59, 0x46, 0xfb, 0x9c, 0xa4, 0xf5,
	0x56, 0xef, 0x77, 0x29, 0xed, 0xf6, 0x88, 0x2e, 0xa2, 0xb3, 0x7e, 0x47, 0xf7, 0xc2, 0x4c, 0xca,
	0xea, 0x83, 0xeb, 0xa9, 0x98, 0xb3, 0x7e, 0x9b, 0x67, 0xd9, 0x87, 0x7d, 0x3f, 0xf2, 0x74, 0x2f,
	0x0c, 0x29, 0xf7, 0xb8, 0x28, 0x1b, 0x73, 0x8f, 0xf7, 0xe3, 0x2c, 0xfd, 0xe8, 0x46, 0x7a, 0x40,
	0x58, 0x22, 0x2f, 0x08, 0xbb, 0xd9, 0x91, 0x95, 0x81, 0xd7, 0x0b, 0x7c, 0x8f, 0x13, 0x7d, 0xf4,
	0x91, 0x26, 0x36, 0xbe, 0xe4, 0x41, 0x71, 0x84, 0x66, 0x3b, 0x91, 0x8c, 0xba, 0xb0, 0xcc, 0x99,
	0x17, 0xc6, 0x11, 0x65, 0xbc, 0x2c, 0x57, 0xa4, 0x6a, 0x69, 0x6b, 0x4f, 0x9b, 0xd5, 0x27, 0xcd,
	0x19, 0x21, 0x9c, 0x61, 0x44, 0xcc, 0xa5, 0x4b, 0x73, 0xfe, 0xa3, 0x24, 0xab, 0x12, 0x1e, 0xb3,
	0x91, 0x0f, 0x4b, 0x42, 0x41, 0x9b, 0xf6, 0xca, 0x39, 0x51, 0x67, 0x77, 0xf6, 0x3a, 0x76, 0x46,
	0xb8, 0x56, 0xe6, 0x8a, 0x8c, 0xaa, 0xa0, 0x24, 0x56, 0xb9, 0x11, 0x23, 0x9d, 0xe0, 0xa2, 0x2c,
	0x55, 0xa4, 0xea, 0xb2, 0xb9, 0x78, 0x69, 0xe6, 0x99, 0x5c, 0x91, 0x30, 0x24, 0x39, 0x5b, 0xa4,
	0x50, 0x17, 0x0a, 0xe2, 0xad, 0xdc, 0x36, 0x0d, 0x3b, 0x41, 0xb7, 0x9c, 0xaf, 0x48, 0x55, 0x65,
	0xab, 0x3e, 0xbb, 0x26, 0x9c, 0x50, 0x6a, 0x02, 0xd2, 0x67, 0xe2, 0x3d, 0xb0, 0xc2, 0xc6, 0xff,
	0x10, 0x81, 0x52, 0x76, 0x25, 0x23, 0x95, 0xe7, 0x2b, 0xb9, 0xaa, 0x72, 0x9b, 0xf6, 0xd3, 0x87,
	0xdb, 0x17, 0xa7, 0x70, 0x91, 0x4f, 0x44, 0xf1, 0x4e, 0xfd, 0xf3, 0x8f, 0x4f, 0xeb, 0x7b, 0xf0,
	0x3c, 0x85, 0xa6, 0x4d, 0x65, 0xc0, 0x3f, 0xf0, 0xb6, 0xbc, 0x5e, 0x74, 0xee, 0x3d, 0xd1, 0x26,
	0xc6, 0x61, 0xe3, 0xab, 0x0c, 0x85, 0xc9, 0x2a, 0x68, 0x0d, 0xf2, 0xa1, 0xf7, 0x8e, 0x5c, 0x77,
	0x52, 0xfc, 0x44, 0x4f, 0xa1, 0xc0, 0x87, 0x11, 0xf1, 0x47, 0x1e, 0xe6, 0x84, 0x87, 0x77, 0xb5,
	0x74, 0xb8, 0xb5, 0xd1, 0x70, 0x6b, 0x46, 0x38, 0x3c, 0x98, 0xc3, 0x8a, 0x38, 0x9b, 0xb9, 0x72,
	0x06, 0xeb, 0xe7, 0x81, 0xef, 0x93, 0xd0, 0x15, 0x82, 0x5d, 0x9f, 0x44, 0x8c, 0xb4, 0x3d, 0x3e,
	0x86, 0xc9, 0x02, 0xb6, 0x72, 0x03, 0x76, 0x2c, 0x36, 0xc5, 0x94, 0xcb, 0xd2, 0xc1, 0x1c, 0x5e,
	0x4b, 0x21, 0x8d, 0x84, 0x51, 0xbf, 0x42, 0xa4, 0x35, 0x76, 0x1a, 0x89, 0x25, 0x2f, 0x60, 0xf7,
	0xb6, 0x96, 0xa4, 0x16, 0x98, 0x45, 0x50, 0xd2, 0xbb, 0x6e, 0xd2, 0xc0, 0xc6, 0x37, 0x19, 0xee,
	0x5d, 0x59, 0x26, 0xa6, 0xce, 0x8a, 0xc4, 0x16, 0x4e, 0xef, 0x92, 0xf4, 0x9f, 0x76, 0x49, 0xfe,
	0x57, 0xbb, 0xb4, 0xd3, 0x4a, 0xec, 0x6b, 0xc2, 0xcb, 0xbf, 0x98, 0xa8, 0x49, 0x7b, 0x36, 0x9b,
	0x50, 0x9c, 0xea, 0x0d, 0x21, 0x28, 0x19, 0x27, 0x8e, 0xe5, 0x3a, 0xd8, 0x68, 0x1d, 0xdb, 0x16,
	0x76, 0xd4, 0x39, 0x04, 0xb0, 0xb0, 0x8f, 0x8d, 0xa3, 0x46, 0x5d, 0x95, 0x50, 0x01, 0x96, 0x4e,
	0x5a, 0x59, 0x24, 0x27, 0x99, 0x83, 0x86, 0x51, 0x6f, 0x60, 0x35, 0xb7, 0x79, 0x0a, 0x85, 0x49,
	0xf9, 0xe8, 0x0e, 0x14, 0x05, 0xc9, 0xc6, 0x96, 0x63, 0xd5, 0xac, 0xc3, 0x14, 0x64, 0x36, 0x5b,
	0x06, 0x7e, 0xad, 0x4a, 0xa8, 0x04, 0x70, 0x68, 0xbc, 0x72, 0xb3, 0x58, 0x46, 0x0a, 0x2c, 0xd6,
	0xac, 0x23, 0xdb, 0xa8, 0x39, 0x6a, 0x2e, 0x09, 0x9c, 0xd3, 0xa6, 0xe3, 0x34, 0xb0, 0x9a, 0x37,
	0xdf, 0x7c, 0xff, 0xf0, 0xf3, 0xd7, 0x82, 0xac, 0xca, 0xb0, 0x1b, 0xd0, 0xd4, 0xd3, 0xb4, 0xb7,
	0x59, 0xed, 0x35, 0xd5, 0x89, 0xb5, 0x12, 0x5a, 0x6d, 0xe9, 0x6c, 0x41, 0x78, 0xbb, 0xfd, 0x3b,
	0x00, 0x00, 0xff, 0xff, 0xcd, 0xc1, 0x9c, 0xcf, 0xce, 0x06, 0x00, 0x00,
}