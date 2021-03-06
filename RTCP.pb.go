// Code generated by protoc-gen-go. DO NOT EDIT.
// source: RTCP.proto

package RTCP

import (
	fmt "fmt"
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

// HEP represents RTCP packet
type RTCPpkt struct {
	Version              uint32   `protobuf:"varint,1,opt,name=Version,proto3" json:"Version,omitempty"`
	Protocol             uint32   `protobuf:"varint,2,opt,name=Protocol,proto3" json:"Protocol,omitempty"`
	SrcIP                string   `protobuf:"bytes,3,opt,name=SrcIP,proto3" json:"SrcIP,omitempty"`
	DstIP                string   `protobuf:"bytes,4,opt,name=DstIP,proto3" json:"DstIP,omitempty"`
	SrcPort              uint32   `protobuf:"varint,5,opt,name=SrcPort,proto3" json:"SrcPort,omitempty"`
	DstPort              uint32   `protobuf:"varint,6,opt,name=DstPort,proto3" json:"DstPort,omitempty"`
	Tsec                 uint32   `protobuf:"varint,7,opt,name=Tsec,proto3" json:"Tsec,omitempty"`
	Tmsec                uint32   `protobuf:"varint,8,opt,name=Tmsec,proto3" json:"Tmsec,omitempty"`
	ProtoType            uint32   `protobuf:"varint,9,opt,name=ProtoType,proto3" json:"ProtoType,omitempty"`
	NodeID               uint32   `protobuf:"varint,10,opt,name=NodeID,proto3" json:"NodeID,omitempty"`
	NodePW               string   `protobuf:"bytes,11,opt,name=NodePW,proto3" json:"NodePW,omitempty"`
	Payload              string   `protobuf:"bytes,12,opt,name=Payload,proto3" json:"Payload,omitempty"`
	CID                  string   `protobuf:"bytes,13,opt,name=CID,proto3" json:"CID,omitempty"`
	Vlan                 uint32   `protobuf:"varint,14,opt,name=Vlan,proto3" json:"Vlan,omitempty"`
	ProtoString          string   `protobuf:"bytes,15,opt,name=ProtoString,proto3" json:"ProtoString,omitempty"`
	Timestamp            string   `protobuf:"bytes,16,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	HostTag              string   `protobuf:"bytes,17,opt,name=HostTag,proto3" json:"HostTag,omitempty"`
	NodeName             string   `protobuf:"bytes,18,opt,name=NodeName,proto3" json:"NodeName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RTCPpkt) Reset()         { *m = RTCPpkt{} }
func (m *RTCPpkt) String() string { return proto.CompactTextString(m) }
func (*RTCPpkt) ProtoMessage()    {}
func (*RTCPpkt) Descriptor() ([]byte, []int) {
	return fileDescriptor_a594914a6b0be9e0, []int{0}
}

func (m *RTCPpkt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RTCPpkt.Unmarshal(m, b)
}
func (m *RTCPpkt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RTCPpkt.Marshal(b, m, deterministic)
}
func (m *RTCPpkt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RTCPpkt.Merge(m, src)
}
func (m *RTCPpkt) XXX_Size() int {
	return xxx_messageInfo_RTCPpkt.Size(m)
}
func (m *RTCPpkt) XXX_DiscardUnknown() {
	xxx_messageInfo_RTCPpkt.DiscardUnknown(m)
}

var xxx_messageInfo_RTCPpkt proto.InternalMessageInfo

func (m *RTCPpkt) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *RTCPpkt) GetProtocol() uint32 {
	if m != nil {
		return m.Protocol
	}
	return 0
}

func (m *RTCPpkt) GetSrcIP() string {
	if m != nil {
		return m.SrcIP
	}
	return ""
}

func (m *RTCPpkt) GetDstIP() string {
	if m != nil {
		return m.DstIP
	}
	return ""
}

func (m *RTCPpkt) GetSrcPort() uint32 {
	if m != nil {
		return m.SrcPort
	}
	return 0
}

func (m *RTCPpkt) GetDstPort() uint32 {
	if m != nil {
		return m.DstPort
	}
	return 0
}

func (m *RTCPpkt) GetTsec() uint32 {
	if m != nil {
		return m.Tsec
	}
	return 0
}

func (m *RTCPpkt) GetTmsec() uint32 {
	if m != nil {
		return m.Tmsec
	}
	return 0
}

func (m *RTCPpkt) GetProtoType() uint32 {
	if m != nil {
		return m.ProtoType
	}
	return 0
}

func (m *RTCPpkt) GetNodeID() uint32 {
	if m != nil {
		return m.NodeID
	}
	return 0
}

func (m *RTCPpkt) GetNodePW() string {
	if m != nil {
		return m.NodePW
	}
	return ""
}

func (m *RTCPpkt) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func (m *RTCPpkt) GetCID() string {
	if m != nil {
		return m.CID
	}
	return ""
}

func (m *RTCPpkt) GetVlan() uint32 {
	if m != nil {
		return m.Vlan
	}
	return 0
}

func (m *RTCPpkt) GetProtoString() string {
	if m != nil {
		return m.ProtoString
	}
	return ""
}

func (m *RTCPpkt) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *RTCPpkt) GetHostTag() string {
	if m != nil {
		return m.HostTag
	}
	return ""
}

func (m *RTCPpkt) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func init() {
	proto.RegisterType((*RTCPpkt)(nil), "RTCPpkt")
}

func init() { proto.RegisterFile("RTCP.proto", fileDescriptor_a594914a6b0be9e0) }

var fileDescriptor_a594914a6b0be9e0 = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xcd, 0x6a, 0xeb, 0x30,
	0x10, 0x85, 0xc9, 0xcd, 0xaf, 0x27, 0x37, 0x6d, 0x2a, 0x4a, 0x19, 0x4a, 0x17, 0xa6, 0xab, 0xac,
	0xba, 0xe9, 0x23, 0xc4, 0x8b, 0x7a, 0x13, 0x84, 0x2d, 0xd2, 0xb5, 0xea, 0x88, 0x60, 0x6a, 0x5b,
	0x46, 0xd2, 0x26, 0x0f, 0xd9, 0x77, 0x2a, 0x33, 0xb2, 0x93, 0xee, 0xe6, 0xfb, 0x8e, 0xf1, 0x39,
	0x20, 0x80, 0x42, 0xed, 0xe5, 0x5b, 0xef, 0x6c, 0xb0, 0xaf, 0x3f, 0x53, 0x58, 0x12, 0xf6, 0xdf,
	0x41, 0x20, 0x2c, 0x8f, 0xc6, 0xf9, 0xda, 0x76, 0x38, 0x49, 0x27, 0xbb, 0x4d, 0x31, 0xa2, 0x78,
	0x86, 0x95, 0xa4, 0xcf, 0x2b, 0xdb, 0xe0, 0x3f, 0x8e, 0xae, 0x2c, 0x1e, 0x61, 0x5e, 0xba, 0x2a,
	0x97, 0x38, 0x4d, 0x27, 0xbb, 0xa4, 0x88, 0x40, 0x36, 0xf3, 0x21, 0x97, 0x38, 0x8b, 0x96, 0x81,
	0x1a, 0x4a, 0x57, 0x49, 0xeb, 0x02, 0xce, 0x63, 0xc3, 0x80, 0x94, 0x64, 0x3e, 0x70, 0xb2, 0x88,
	0xc9, 0x80, 0x42, 0xc0, 0x4c, 0x79, 0x53, 0xe1, 0x92, 0x35, 0xdf, 0xf4, 0x77, 0xd5, 0x92, 0x5c,
	0xb1, 0x8c, 0x20, 0x5e, 0x20, 0xe1, 0x55, 0xea, 0xd2, 0x1b, 0x4c, 0x38, 0xb9, 0x09, 0xf1, 0x04,
	0x8b, 0x83, 0x3d, 0x99, 0x3c, 0x43, 0xe0, 0x68, 0xa0, 0xd1, 0xcb, 0x4f, 0x5c, 0xf3, 0xd4, 0x81,
	0x68, 0x91, 0xd4, 0x97, 0xc6, 0xea, 0x13, 0xfe, 0xe7, 0x60, 0x44, 0xb1, 0x85, 0xe9, 0x3e, 0xcf,
	0x70, 0xc3, 0x96, 0x4e, 0xda, 0x78, 0x6c, 0x74, 0x87, 0x77, 0x71, 0x23, 0xdd, 0x22, 0x85, 0x35,
	0x97, 0x97, 0xc1, 0xd5, 0xdd, 0x19, 0xef, 0xf9, 0xeb, 0xbf, 0x8a, 0xf6, 0xaa, 0xba, 0x35, 0x3e,
	0xe8, 0xb6, 0xc7, 0x2d, 0xe7, 0x37, 0x41, 0xfd, 0x1f, 0xd6, 0x07, 0xa5, 0xcf, 0xf8, 0x10, 0xfb,
	0x07, 0xa4, 0xd7, 0xa0, 0x8d, 0x07, 0xdd, 0x1a, 0x14, 0x1c, 0x5d, 0xf9, 0x6b, 0xc1, 0xcf, 0xfa,
	0xfe, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xc5, 0xf1, 0x91, 0xb4, 0xe4, 0x01, 0x00, 0x00,
}
