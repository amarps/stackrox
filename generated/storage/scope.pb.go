// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage/scope.proto

package storage

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Scope struct {
	Cluster              string       `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Namespace            string       `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Label                *Scope_Label `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Scope) Reset()         { *m = Scope{} }
func (m *Scope) String() string { return proto.CompactTextString(m) }
func (*Scope) ProtoMessage()    {}
func (*Scope) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c72a373e2c7b3ae, []int{0}
}
func (m *Scope) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Scope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Scope.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Scope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Scope.Merge(m, src)
}
func (m *Scope) XXX_Size() int {
	return m.Size()
}
func (m *Scope) XXX_DiscardUnknown() {
	xxx_messageInfo_Scope.DiscardUnknown(m)
}

var xxx_messageInfo_Scope proto.InternalMessageInfo

func (m *Scope) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *Scope) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *Scope) GetLabel() *Scope_Label {
	if m != nil {
		return m.Label
	}
	return nil
}

func (m *Scope) MessageClone() proto.Message {
	return m.Clone()
}
func (m *Scope) Clone() *Scope {
	if m == nil {
		return nil
	}
	cloned := new(Scope)
	*cloned = *m

	cloned.Label = m.Label.Clone()
	return cloned
}

type Scope_Label struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Scope_Label) Reset()         { *m = Scope_Label{} }
func (m *Scope_Label) String() string { return proto.CompactTextString(m) }
func (*Scope_Label) ProtoMessage()    {}
func (*Scope_Label) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c72a373e2c7b3ae, []int{0, 0}
}
func (m *Scope_Label) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Scope_Label) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Scope_Label.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Scope_Label) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Scope_Label.Merge(m, src)
}
func (m *Scope_Label) XXX_Size() int {
	return m.Size()
}
func (m *Scope_Label) XXX_DiscardUnknown() {
	xxx_messageInfo_Scope_Label.DiscardUnknown(m)
}

var xxx_messageInfo_Scope_Label proto.InternalMessageInfo

func (m *Scope_Label) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Scope_Label) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Scope_Label) MessageClone() proto.Message {
	return m.Clone()
}
func (m *Scope_Label) Clone() *Scope_Label {
	if m == nil {
		return nil
	}
	cloned := new(Scope_Label)
	*cloned = *m

	return cloned
}

func init() {
	proto.RegisterType((*Scope)(nil), "storage.Scope")
	proto.RegisterType((*Scope_Label)(nil), "storage.Scope.Label")
}

func init() { proto.RegisterFile("storage/scope.proto", fileDescriptor_7c72a373e2c7b3ae) }

var fileDescriptor_7c72a373e2c7b3ae = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x2e, 0xc9, 0x2f,
	0x4a, 0x4c, 0x4f, 0xd5, 0x2f, 0x4e, 0xce, 0x2f, 0x48, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x87, 0x0a, 0x2a, 0xcd, 0x61, 0xe4, 0x62, 0x0d, 0x06, 0x49, 0x08, 0x49, 0x70, 0xb1, 0x27,
	0xe7, 0x94, 0x16, 0x97, 0xa4, 0x16, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xb8, 0x42,
	0x32, 0x5c, 0x9c, 0x79, 0x89, 0xb9, 0xa9, 0xc5, 0x05, 0x89, 0xc9, 0xa9, 0x12, 0x4c, 0x60, 0x39,
	0x84, 0x80, 0x90, 0x16, 0x17, 0x6b, 0x4e, 0x62, 0x52, 0x6a, 0x8e, 0x04, 0xb3, 0x02, 0xa3, 0x06,
	0xb7, 0x91, 0x88, 0x1e, 0xd4, 0x68, 0x3d, 0xb0, 0xb1, 0x7a, 0x3e, 0x20, 0xb9, 0x20, 0x88, 0x12,
	0x29, 0x7d, 0x2e, 0x56, 0x30, 0x5f, 0x48, 0x80, 0x8b, 0x39, 0x3b, 0xb5, 0x12, 0x6a, 0x11, 0x88,
	0x29, 0x24, 0xc2, 0xc5, 0x5a, 0x96, 0x98, 0x53, 0x0a, 0xb3, 0x00, 0xc2, 0x71, 0x32, 0x39, 0xf1,
	0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x67, 0x3c, 0x96, 0x63, 0xe0,
	0x92, 0xcc, 0xcc, 0xd7, 0x2b, 0x2e, 0x49, 0x4c, 0xce, 0x2e, 0xca, 0xaf, 0x80, 0x78, 0x06, 0x66,
	0x61, 0x14, 0xcc, 0x53, 0x49, 0x6c, 0x60, 0x71, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xeb,
	0x65, 0x3f, 0x53, 0xfb, 0x00, 0x00, 0x00,
}

func (m *Scope) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Scope) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Scope) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Label != nil {
		{
			size, err := m.Label.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintScope(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Namespace) > 0 {
		i -= len(m.Namespace)
		copy(dAtA[i:], m.Namespace)
		i = encodeVarintScope(dAtA, i, uint64(len(m.Namespace)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Cluster) > 0 {
		i -= len(m.Cluster)
		copy(dAtA[i:], m.Cluster)
		i = encodeVarintScope(dAtA, i, uint64(len(m.Cluster)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Scope_Label) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Scope_Label) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Scope_Label) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintScope(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintScope(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintScope(dAtA []byte, offset int, v uint64) int {
	offset -= sovScope(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Scope) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Cluster)
	if l > 0 {
		n += 1 + l + sovScope(uint64(l))
	}
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovScope(uint64(l))
	}
	if m.Label != nil {
		l = m.Label.Size()
		n += 1 + l + sovScope(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Scope_Label) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovScope(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovScope(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovScope(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozScope(x uint64) (n int) {
	return sovScope(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Scope) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowScope
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
			return fmt.Errorf("proto: Scope: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Scope: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cluster", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScope
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthScope
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthScope
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cluster = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScope
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthScope
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthScope
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Label", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScope
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
				return ErrInvalidLengthScope
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthScope
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Label == nil {
				m.Label = &Scope_Label{}
			}
			if err := m.Label.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipScope(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthScope
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
func (m *Scope_Label) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowScope
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
			return fmt.Errorf("proto: Label: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Label: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScope
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthScope
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthScope
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScope
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthScope
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthScope
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipScope(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthScope
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
func skipScope(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowScope
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
					return 0, ErrIntOverflowScope
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
					return 0, ErrIntOverflowScope
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
				return 0, ErrInvalidLengthScope
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupScope
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthScope
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthScope        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowScope          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupScope = fmt.Errorf("proto: unexpected end of group")
)