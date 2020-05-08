// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/extensions/filters/http/rbac/v3/rbac.proto

package envoy_extensions_filters_http_rbac_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/datawire/ambassador/pkg/api/envoy/config/rbac/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// RBAC filter config.
type RBAC struct {
	// Specify the RBAC rules to be applied globally.
	// If absent, no enforcing RBAC policy will be applied.
	Rules *v3.RBAC `protobuf:"bytes,1,opt,name=rules,proto3" json:"rules,omitempty"`
	// Shadow rules are not enforced by the filter (i.e., returning a 403)
	// but will emit stats and logs and can be used for rule testing.
	// If absent, no shadow RBAC policy will be applied.
	ShadowRules          *v3.RBAC `protobuf:"bytes,2,opt,name=shadow_rules,json=shadowRules,proto3" json:"shadow_rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RBAC) Reset()         { *m = RBAC{} }
func (m *RBAC) String() string { return proto.CompactTextString(m) }
func (*RBAC) ProtoMessage()    {}
func (*RBAC) Descriptor() ([]byte, []int) {
	return fileDescriptor_b77e76eac62eed05, []int{0}
}
func (m *RBAC) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RBAC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RBAC.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RBAC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RBAC.Merge(m, src)
}
func (m *RBAC) XXX_Size() int {
	return m.Size()
}
func (m *RBAC) XXX_DiscardUnknown() {
	xxx_messageInfo_RBAC.DiscardUnknown(m)
}

var xxx_messageInfo_RBAC proto.InternalMessageInfo

func (m *RBAC) GetRules() *v3.RBAC {
	if m != nil {
		return m.Rules
	}
	return nil
}

func (m *RBAC) GetShadowRules() *v3.RBAC {
	if m != nil {
		return m.ShadowRules
	}
	return nil
}

type RBACPerRoute struct {
	// Override the global configuration of the filter with this new config.
	// If absent, the global RBAC policy will be disabled for this route.
	Rbac                 *RBAC    `protobuf:"bytes,2,opt,name=rbac,proto3" json:"rbac,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RBACPerRoute) Reset()         { *m = RBACPerRoute{} }
func (m *RBACPerRoute) String() string { return proto.CompactTextString(m) }
func (*RBACPerRoute) ProtoMessage()    {}
func (*RBACPerRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_b77e76eac62eed05, []int{1}
}
func (m *RBACPerRoute) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RBACPerRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RBACPerRoute.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RBACPerRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RBACPerRoute.Merge(m, src)
}
func (m *RBACPerRoute) XXX_Size() int {
	return m.Size()
}
func (m *RBACPerRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_RBACPerRoute.DiscardUnknown(m)
}

var xxx_messageInfo_RBACPerRoute proto.InternalMessageInfo

func (m *RBACPerRoute) GetRbac() *RBAC {
	if m != nil {
		return m.Rbac
	}
	return nil
}

func init() {
	proto.RegisterType((*RBAC)(nil), "envoy.extensions.filters.http.rbac.v3.RBAC")
	proto.RegisterType((*RBACPerRoute)(nil), "envoy.extensions.filters.http.rbac.v3.RBACPerRoute")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/http/rbac/v3/rbac.proto", fileDescriptor_b77e76eac62eed05)
}

var fileDescriptor_b77e76eac62eed05 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xb1, 0x4a, 0x03, 0x31,
	0x18, 0xc7, 0x49, 0xa9, 0x45, 0xd3, 0x0e, 0xe5, 0x16, 0xa5, 0xe0, 0xa9, 0x85, 0xa2, 0xa0, 0x26,
	0xa5, 0xe7, 0x54, 0x10, 0xb1, 0x6e, 0x4e, 0xe5, 0x06, 0x57, 0x49, 0x7b, 0x69, 0x1b, 0x38, 0x92,
	0x23, 0xc9, 0x9d, 0xed, 0xe6, 0x28, 0xf8, 0x06, 0xee, 0xbe, 0x84, 0xbb, 0xd0, 0x51, 0xdf, 0x40,
	0xee, 0x49, 0x24, 0xc9, 0x1d, 0x2a, 0x0a, 0xde, 0x94, 0xc0, 0xf7, 0xfb, 0xe5, 0xff, 0xe7, 0x0b,
	0xec, 0x53, 0x9e, 0x89, 0x15, 0xa6, 0x4b, 0x4d, 0xb9, 0x62, 0x82, 0x2b, 0x3c, 0x63, 0xb1, 0xa6,
	0x52, 0xe1, 0x85, 0xd6, 0x09, 0x96, 0x13, 0x32, 0xc5, 0x59, 0x60, 0x4f, 0x94, 0x48, 0xa1, 0x85,
	0xd7, 0xb3, 0x06, 0xfa, 0x32, 0x50, 0x61, 0x20, 0x63, 0x20, 0x4b, 0x66, 0x41, 0x67, 0xcf, 0x3d,
	0x3c, 0x15, 0x7c, 0xc6, 0xe6, 0x7f, 0xbc, 0xd3, 0xd9, 0x4d, 0xa3, 0x84, 0x60, 0xc2, 0xb9, 0xd0,
	0x44, 0xdb, 0x64, 0xa5, 0x89, 0x4e, 0x55, 0x31, 0x3e, 0xf8, 0x35, 0xce, 0xa8, 0x34, 0x79, 0x8c,
	0xcf, 0x0b, 0x64, 0x3b, 0x23, 0x31, 0x8b, 0x88, 0xa6, 0xb8, 0xbc, 0xb8, 0x41, 0xf7, 0x19, 0xc0,
	0x7a, 0x38, 0xba, 0xbc, 0xf2, 0xfa, 0x70, 0x43, 0xa6, 0x31, 0x55, 0x3b, 0x60, 0x1f, 0x1c, 0x35,
	0x07, 0x1d, 0xe4, 0xba, 0xbb, 0x52, 0x65, 0x55, 0x64, 0xd0, 0xd0, 0x81, 0xde, 0x39, 0x6c, 0xa9,
	0x05, 0x89, 0xc4, 0xdd, 0xad, 0x13, 0x6b, 0xff, 0x8a, 0x4d, 0xc7, 0x87, 0x06, 0x1f, 0x9e, 0x3c,
	0xbd, 0x3e, 0xf8, 0x87, 0xb0, 0xf7, 0x03, 0x77, 0xfb, 0xf9, 0xbe, 0x9e, 0x81, 0x55, 0xbb, 0x8f,
	0x00, 0xb6, 0xcc, 0x65, 0x4c, 0x65, 0x28, 0x52, 0x4d, 0xbd, 0x0b, 0x58, 0x37, 0x40, 0x91, 0x7a,
	0x8c, 0x2a, 0xad, 0xda, 0xd5, 0xb0, 0xe2, 0xf0, 0xcc, 0xe4, 0x63, 0x78, 0x5a, 0x29, 0xbf, 0x8c,
	0xbd, 0xae, 0x6f, 0x82, 0x76, 0x6d, 0x74, 0xb3, 0xce, 0x7d, 0xf0, 0x96, 0xfb, 0xe0, 0x23, 0xf7,
	0xc1, 0xcb, 0xfd, 0xfa, 0xbd, 0x51, 0x6b, 0x03, 0x18, 0x30, 0xe1, 0x6a, 0x24, 0x52, 0x2c, 0x57,
	0xd5, 0x1a, 0x8d, 0xb6, 0xc2, 0x09, 0x99, 0x8e, 0xcd, 0x5f, 0x8c, 0xc1, 0xa4, 0x61, 0x3f, 0x25,
	0xf8, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xf7, 0xa5, 0xcf, 0x1f, 0x6b, 0x02, 0x00, 0x00,
}

func (m *RBAC) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RBAC) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RBAC) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.ShadowRules != nil {
		{
			size, err := m.ShadowRules.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRbac(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Rules != nil {
		{
			size, err := m.Rules.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRbac(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RBACPerRoute) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RBACPerRoute) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RBACPerRoute) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Rbac != nil {
		{
			size, err := m.Rbac.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRbac(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func encodeVarintRbac(dAtA []byte, offset int, v uint64) int {
	offset -= sovRbac(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RBAC) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Rules != nil {
		l = m.Rules.Size()
		n += 1 + l + sovRbac(uint64(l))
	}
	if m.ShadowRules != nil {
		l = m.ShadowRules.Size()
		n += 1 + l + sovRbac(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RBACPerRoute) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Rbac != nil {
		l = m.Rbac.Size()
		n += 1 + l + sovRbac(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovRbac(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRbac(x uint64) (n int) {
	return sovRbac(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RBAC) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRbac
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RBAC: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RBAC: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rules", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRbac
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRbac
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRbac
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Rules == nil {
				m.Rules = &v3.RBAC{}
			}
			if err := m.Rules.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShadowRules", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRbac
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRbac
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRbac
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ShadowRules == nil {
				m.ShadowRules = &v3.RBAC{}
			}
			if err := m.ShadowRules.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRbac(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRbac
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRbac
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RBACPerRoute) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRbac
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RBACPerRoute: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RBACPerRoute: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rbac", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRbac
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRbac
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRbac
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Rbac == nil {
				m.Rbac = &RBAC{}
			}
			if err := m.Rbac.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRbac(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRbac
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRbac
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRbac(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRbac
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRbac
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRbac
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthRbac
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRbac
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRbac
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRbac        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRbac          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRbac = fmt.Errorf("proto: unexpected end of group")
)