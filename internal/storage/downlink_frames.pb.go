// Code generated by protoc-gen-go. DO NOT EDIT.
// source: downlink_frames.proto

package storage

import (
	fmt "fmt"
	gw "github.com/brocaar/chirpstack-api/go/v3/gw"
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

type DownlinkFrames struct {
	// Token.
	Token uint32 `protobuf:"varint,1,opt,name=token,proto3" json:"token,omitempty"`
	// DevEUI.
	DevEui []byte `protobuf:"bytes,2,opt,name=dev_eui,json=devEui,proto3" json:"dev_eui,omitempty"`
	// Multicast Group ID.
	MulticastGroupId []byte `protobuf:"bytes,3,opt,name=multicast_group_id,json=multicastGroupId,proto3" json:"multicast_group_id,omitempty"`
	// Downlink frames.
	DownlinkFrames []*gw.DownlinkFrame `protobuf:"bytes,4,rep,name=downlink_frames,json=downlinkFrames,proto3" json:"downlink_frames,omitempty"`
	// Routing Profile ID.
	RoutingProfileId []byte `protobuf:"bytes,5,opt,name=routing_profile_id,json=routingProfileId,proto3" json:"routing_profile_id,omitempty"`
	// Downlink frame-counter (full).
	FCnt uint32 `protobuf:"varint,6,opt,name=f_cnt,json=fCnt,proto3" json:"f_cnt,omitempty"`
	// Encrypted FOpts (LoRaWAN 1.1).
	EncryptedFopts bool `protobuf:"varint,7,opt,name=encrypted_fopts,json=encryptedFopts,proto3" json:"encrypted_fopts,omitempty"`
	// Network session encryption key (for FOpts).
	NwkSEncKey           []byte   `protobuf:"bytes,8,opt,name=nwk_s_enc_key,json=nwkSEncKey,proto3" json:"nwk_s_enc_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownlinkFrames) Reset()         { *m = DownlinkFrames{} }
func (m *DownlinkFrames) String() string { return proto.CompactTextString(m) }
func (*DownlinkFrames) ProtoMessage()    {}
func (*DownlinkFrames) Descriptor() ([]byte, []int) {
	return fileDescriptor_a06c04f39a283b6b, []int{0}
}

func (m *DownlinkFrames) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownlinkFrames.Unmarshal(m, b)
}
func (m *DownlinkFrames) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownlinkFrames.Marshal(b, m, deterministic)
}
func (m *DownlinkFrames) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownlinkFrames.Merge(m, src)
}
func (m *DownlinkFrames) XXX_Size() int {
	return xxx_messageInfo_DownlinkFrames.Size(m)
}
func (m *DownlinkFrames) XXX_DiscardUnknown() {
	xxx_messageInfo_DownlinkFrames.DiscardUnknown(m)
}

var xxx_messageInfo_DownlinkFrames proto.InternalMessageInfo

func (m *DownlinkFrames) GetToken() uint32 {
	if m != nil {
		return m.Token
	}
	return 0
}

func (m *DownlinkFrames) GetDevEui() []byte {
	if m != nil {
		return m.DevEui
	}
	return nil
}

func (m *DownlinkFrames) GetMulticastGroupId() []byte {
	if m != nil {
		return m.MulticastGroupId
	}
	return nil
}

func (m *DownlinkFrames) GetDownlinkFrames() []*gw.DownlinkFrame {
	if m != nil {
		return m.DownlinkFrames
	}
	return nil
}

func (m *DownlinkFrames) GetRoutingProfileId() []byte {
	if m != nil {
		return m.RoutingProfileId
	}
	return nil
}

func (m *DownlinkFrames) GetFCnt() uint32 {
	if m != nil {
		return m.FCnt
	}
	return 0
}

func (m *DownlinkFrames) GetEncryptedFopts() bool {
	if m != nil {
		return m.EncryptedFopts
	}
	return false
}

func (m *DownlinkFrames) GetNwkSEncKey() []byte {
	if m != nil {
		return m.NwkSEncKey
	}
	return nil
}

func init() {
	proto.RegisterType((*DownlinkFrames)(nil), "storage.DownlinkFrames")
}

func init() { proto.RegisterFile("downlink_frames.proto", fileDescriptor_a06c04f39a283b6b) }

var fileDescriptor_a06c04f39a283b6b = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xdf, 0x4a, 0xf3, 0x30,
	0x18, 0xc6, 0xe9, 0xfe, 0x75, 0x64, 0xdf, 0xba, 0xcf, 0xa8, 0x18, 0x3c, 0xaa, 0x9e, 0xd8, 0x83,
	0x51, 0x41, 0xcf, 0x3c, 0xd5, 0x4d, 0x86, 0x27, 0x52, 0x2f, 0x20, 0xd4, 0xe6, 0x6d, 0x08, 0xed,
	0xde, 0x94, 0x34, 0x5d, 0xe8, 0x9d, 0x79, 0x79, 0xd2, 0x76, 0x0c, 0xe6, 0x61, 0x9e, 0xdf, 0x03,
	0xcf, 0x2f, 0x2f, 0xb9, 0x16, 0xda, 0x61, 0xa9, 0xb0, 0xe0, 0xb9, 0x49, 0xf7, 0x50, 0xc7, 0x95,
	0xd1, 0x56, 0x53, 0xbf, 0xb6, 0xda, 0xa4, 0x12, 0x6e, 0x17, 0xd2, 0x3d, 0x4a, 0x37, 0xa4, 0xf7,
	0x3f, 0x23, 0x12, 0xbc, 0x1d, 0xfb, 0xdb, 0xbe, 0x4e, 0xaf, 0xc8, 0xd4, 0xea, 0x02, 0x90, 0x79,
	0xa1, 0x17, 0x2d, 0x93, 0xe1, 0x41, 0x6f, 0x88, 0x2f, 0xe0, 0xc0, 0xa1, 0x51, 0x6c, 0x14, 0x7a,
	0xd1, 0xbf, 0x64, 0x26, 0xe0, 0xb0, 0x69, 0x14, 0x5d, 0x13, 0xba, 0x6f, 0x4a, 0xab, 0xb2, 0xb4,
	0xb6, 0x5c, 0x1a, 0xdd, 0x54, 0x5c, 0x09, 0x36, 0xee, 0x3b, 0xff, 0x4f, 0xe4, 0xbd, 0x03, 0x3b,
	0x41, 0x5f, 0xc8, 0xea, 0x8f, 0x1e, 0x9b, 0x84, 0xe3, 0x68, 0xf1, 0x74, 0x11, 0x4b, 0x17, 0x9f,
	0x99, 0x24, 0x81, 0x38, 0x17, 0x5b, 0x13, 0x6a, 0x74, 0x63, 0x15, 0x4a, 0x5e, 0x19, 0x9d, 0xab,
	0x12, 0xba, 0xa5, 0xe9, 0xb0, 0x74, 0x24, 0x9f, 0x03, 0xd8, 0x09, 0x7a, 0x49, 0xa6, 0x39, 0xcf,
	0xd0, 0xb2, 0x59, 0xff, 0x8d, 0x49, 0xfe, 0x8a, 0x96, 0x3e, 0x90, 0x15, 0x60, 0x66, 0xda, 0xca,
	0x82, 0xe0, 0xb9, 0xae, 0x6c, 0xcd, 0xfc, 0xd0, 0x8b, 0xe6, 0x49, 0x70, 0x8a, 0xb7, 0x5d, 0x4a,
	0xef, 0xc8, 0x12, 0x5d, 0xc1, 0x6b, 0x0e, 0x98, 0xf1, 0x02, 0x5a, 0x36, 0xef, 0x67, 0x08, 0xba,
	0xe2, 0x6b, 0x83, 0xd9, 0x07, 0xb4, 0xdf, 0xb3, 0xfe, 0x82, 0xcf, 0xbf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xa7, 0x19, 0xfa, 0x04, 0x70, 0x01, 0x00, 0x00,
}
