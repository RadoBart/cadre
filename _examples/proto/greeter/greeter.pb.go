// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/greeter/greeter.proto

// package definition

package greeter

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// messages' definitions
type GreetingRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetingRequest) Reset()         { *m = GreetingRequest{} }
func (m *GreetingRequest) String() string { return proto.CompactTextString(m) }
func (*GreetingRequest) ProtoMessage()    {}
func (*GreetingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0df690f529f06ab8, []int{0}
}
func (m *GreetingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetingRequest.Unmarshal(m, b)
}
func (m *GreetingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetingRequest.Marshal(b, m, deterministic)
}
func (m *GreetingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetingRequest.Merge(m, src)
}
func (m *GreetingRequest) XXX_Size() int {
	return xxx_messageInfo_GreetingRequest.Size(m)
}
func (m *GreetingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetingRequest proto.InternalMessageInfo

func (m *GreetingRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GreetingResponse struct {
	Greeting             string   `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetingResponse) Reset()         { *m = GreetingResponse{} }
func (m *GreetingResponse) String() string { return proto.CompactTextString(m) }
func (*GreetingResponse) ProtoMessage()    {}
func (*GreetingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0df690f529f06ab8, []int{1}
}
func (m *GreetingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetingResponse.Unmarshal(m, b)
}
func (m *GreetingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetingResponse.Marshal(b, m, deterministic)
}
func (m *GreetingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetingResponse.Merge(m, src)
}
func (m *GreetingResponse) XXX_Size() int {
	return xxx_messageInfo_GreetingResponse.Size(m)
}
func (m *GreetingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetingResponse proto.InternalMessageInfo

func (m *GreetingResponse) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

func init() {
	proto.RegisterType((*GreetingRequest)(nil), "example.GreetingRequest")
	proto.RegisterType((*GreetingResponse)(nil), "example.GreetingResponse")
}

func init() { proto.RegisterFile("proto/greeter/greeter.proto", fileDescriptor_0df690f529f06ab8) }

var fileDescriptor_0df690f529f06ab8 = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x2f, 0x4a, 0x4d, 0x2d, 0x49, 0x2d, 0x82, 0xd1, 0x7a, 0x60, 0x51, 0x21, 0xf6,
	0xd4, 0x8a, 0xc4, 0xdc, 0x82, 0x9c, 0x54, 0x25, 0x55, 0x2e, 0x7e, 0x77, 0x90, 0x4c, 0x66, 0x5e,
	0x7a, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x10, 0x17, 0x4b, 0x5e, 0x62, 0x6e, 0xaa,
	0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x98, 0xad, 0xa4, 0xc7, 0x25, 0x80, 0x50, 0x56, 0x5c,
	0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x24, 0xc5, 0xc5, 0x91, 0x0e, 0x15, 0x83, 0xaa, 0x85, 0xf3, 0x8d,
	0x02, 0xb8, 0xf8, 0xdc, 0x21, 0x16, 0x06, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a, 0xd9, 0x71,
	0xb1, 0x06, 0x27, 0x56, 0x7a, 0x64, 0x0a, 0x49, 0xe8, 0x41, 0xed, 0xd6, 0x43, 0xb3, 0x58, 0x4a,
	0x12, 0x8b, 0x0c, 0xc4, 0x2e, 0x25, 0x06, 0x27, 0xe3, 0x28, 0xc3, 0xf4, 0xcc, 0x12, 0xbd, 0xdc,
	0xfc, 0x94, 0xd4, 0xa2, 0xbc, 0x92, 0x32, 0xbd, 0xd4, 0x52, 0xfd, 0xf4, 0x7c, 0xfd, 0xe4, 0xc4,
	0x94, 0xa2, 0x54, 0xfd, 0x78, 0xa8, 0x3e, 0x98, 0x27, 0xad, 0xa1, 0x74, 0x12, 0x1b, 0xd8, 0xb7,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x57, 0x78, 0xab, 0x0c, 0x01, 0x00, 0x00,
}