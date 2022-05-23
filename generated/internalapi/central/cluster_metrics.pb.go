// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: internalapi/central/cluster_metrics.proto

package central

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

type ClusterMetrics struct {
	NodeCount            int64    `protobuf:"varint,1,opt,name=node_count,json=nodeCount,proto3" json:"node_count,omitempty"`
	CpuCapacity          int64    `protobuf:"varint,2,opt,name=cpu_capacity,json=cpuCapacity,proto3" json:"cpu_capacity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterMetrics) Reset()         { *m = ClusterMetrics{} }
func (m *ClusterMetrics) String() string { return proto.CompactTextString(m) }
func (*ClusterMetrics) ProtoMessage()    {}
func (*ClusterMetrics) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf4a2b5a7271bdbd, []int{0}
}
func (m *ClusterMetrics) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClusterMetrics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClusterMetrics.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClusterMetrics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterMetrics.Merge(m, src)
}
func (m *ClusterMetrics) XXX_Size() int {
	return m.Size()
}
func (m *ClusterMetrics) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterMetrics.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterMetrics proto.InternalMessageInfo

func (m *ClusterMetrics) GetNodeCount() int64 {
	if m != nil {
		return m.NodeCount
	}
	return 0
}

func (m *ClusterMetrics) GetCpuCapacity() int64 {
	if m != nil {
		return m.CpuCapacity
	}
	return 0
}

func (m *ClusterMetrics) MessageClone() proto.Message {
	return m.Clone()
}
func (m *ClusterMetrics) Clone() *ClusterMetrics {
	if m == nil {
		return nil
	}
	cloned := new(ClusterMetrics)
	*cloned = *m

	return cloned
}

func init() {
	proto.RegisterType((*ClusterMetrics)(nil), "central.ClusterMetrics")
}

func init() {
	proto.RegisterFile("internalapi/central/cluster_metrics.proto", fileDescriptor_cf4a2b5a7271bdbd)
}

var fileDescriptor_cf4a2b5a7271bdbd = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xcc, 0xcc, 0x2b, 0x49,
	0x2d, 0xca, 0x4b, 0xcc, 0x49, 0x2c, 0xc8, 0xd4, 0x4f, 0x4e, 0xcd, 0x2b, 0x29, 0x4a, 0xcc, 0xd1,
	0x4f, 0xce, 0x29, 0x2d, 0x2e, 0x49, 0x2d, 0x8a, 0xcf, 0x4d, 0x2d, 0x29, 0xca, 0x4c, 0x2e, 0xd6,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x4a, 0x2b, 0x05, 0x71, 0xf1, 0x39, 0x43, 0x54,
	0xf8, 0x42, 0x14, 0x08, 0xc9, 0x72, 0x71, 0xe5, 0xe5, 0xa7, 0xa4, 0xc6, 0x27, 0xe7, 0x97, 0xe6,
	0x95, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x30, 0x07, 0x71, 0x82, 0x44, 0x9c, 0x41, 0x02, 0x42, 0x8a,
	0x5c, 0x3c, 0xc9, 0x05, 0xa5, 0xf1, 0xc9, 0x89, 0x05, 0x89, 0xc9, 0x99, 0x25, 0x95, 0x12, 0x4c,
	0x60, 0x05, 0xdc, 0xc9, 0x05, 0xa5, 0xce, 0x50, 0x21, 0x27, 0xc9, 0x13, 0x8f, 0xe4, 0x18, 0x2f,
	0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc6, 0x63, 0x39, 0x86, 0x28, 0x98, 0x75, 0x49,
	0x6c, 0x60, 0xeb, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x31, 0x57, 0xbb, 0xc4, 0xab, 0x00,
	0x00, 0x00,
}

func (m *ClusterMetrics) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClusterMetrics) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClusterMetrics) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.CpuCapacity != 0 {
		i = encodeVarintClusterMetrics(dAtA, i, uint64(m.CpuCapacity))
		i--
		dAtA[i] = 0x10
	}
	if m.NodeCount != 0 {
		i = encodeVarintClusterMetrics(dAtA, i, uint64(m.NodeCount))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintClusterMetrics(dAtA []byte, offset int, v uint64) int {
	offset -= sovClusterMetrics(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClusterMetrics) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NodeCount != 0 {
		n += 1 + sovClusterMetrics(uint64(m.NodeCount))
	}
	if m.CpuCapacity != 0 {
		n += 1 + sovClusterMetrics(uint64(m.CpuCapacity))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovClusterMetrics(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClusterMetrics(x uint64) (n int) {
	return sovClusterMetrics(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClusterMetrics) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClusterMetrics
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
			return fmt.Errorf("proto: ClusterMetrics: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClusterMetrics: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeCount", wireType)
			}
			m.NodeCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClusterMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NodeCount |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CpuCapacity", wireType)
			}
			m.CpuCapacity = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClusterMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CpuCapacity |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipClusterMetrics(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClusterMetrics
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
func skipClusterMetrics(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClusterMetrics
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
					return 0, ErrIntOverflowClusterMetrics
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
					return 0, ErrIntOverflowClusterMetrics
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
				return 0, ErrInvalidLengthClusterMetrics
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupClusterMetrics
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthClusterMetrics
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthClusterMetrics        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClusterMetrics          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupClusterMetrics = fmt.Errorf("proto: unexpected end of group")
)