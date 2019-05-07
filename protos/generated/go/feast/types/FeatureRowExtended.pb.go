// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feast/types/FeatureRowExtended.proto

package types

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Error struct {
	Cause                string   `protobuf:"bytes,1,opt,name=cause,proto3" json:"cause,omitempty"`
	Transform            string   `protobuf:"bytes,2,opt,name=transform,proto3" json:"transform,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	StackTrace           string   `protobuf:"bytes,4,opt,name=stackTrace,proto3" json:"stackTrace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_7823aa2c72575793, []int{0}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCause() string {
	if m != nil {
		return m.Cause
	}
	return ""
}

func (m *Error) GetTransform() string {
	if m != nil {
		return m.Transform
	}
	return ""
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Error) GetStackTrace() string {
	if m != nil {
		return m.StackTrace
	}
	return ""
}

type Attempt struct {
	Attempts             int32    `protobuf:"varint,1,opt,name=attempts,proto3" json:"attempts,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Attempt) Reset()         { *m = Attempt{} }
func (m *Attempt) String() string { return proto.CompactTextString(m) }
func (*Attempt) ProtoMessage()    {}
func (*Attempt) Descriptor() ([]byte, []int) {
	return fileDescriptor_7823aa2c72575793, []int{1}
}

func (m *Attempt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Attempt.Unmarshal(m, b)
}
func (m *Attempt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Attempt.Marshal(b, m, deterministic)
}
func (m *Attempt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Attempt.Merge(m, src)
}
func (m *Attempt) XXX_Size() int {
	return xxx_messageInfo_Attempt.Size(m)
}
func (m *Attempt) XXX_DiscardUnknown() {
	xxx_messageInfo_Attempt.DiscardUnknown(m)
}

var xxx_messageInfo_Attempt proto.InternalMessageInfo

func (m *Attempt) GetAttempts() int32 {
	if m != nil {
		return m.Attempts
	}
	return 0
}

func (m *Attempt) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type FeatureRowExtended struct {
	Row                  *FeatureRow          `protobuf:"bytes,1,opt,name=row,proto3" json:"row,omitempty"`
	LastAttempt          *Attempt             `protobuf:"bytes,2,opt,name=lastAttempt,proto3" json:"lastAttempt,omitempty"`
	FirstSeen            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=firstSeen,proto3" json:"firstSeen,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FeatureRowExtended) Reset()         { *m = FeatureRowExtended{} }
func (m *FeatureRowExtended) String() string { return proto.CompactTextString(m) }
func (*FeatureRowExtended) ProtoMessage()    {}
func (*FeatureRowExtended) Descriptor() ([]byte, []int) {
	return fileDescriptor_7823aa2c72575793, []int{2}
}

func (m *FeatureRowExtended) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeatureRowExtended.Unmarshal(m, b)
}
func (m *FeatureRowExtended) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeatureRowExtended.Marshal(b, m, deterministic)
}
func (m *FeatureRowExtended) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeatureRowExtended.Merge(m, src)
}
func (m *FeatureRowExtended) XXX_Size() int {
	return xxx_messageInfo_FeatureRowExtended.Size(m)
}
func (m *FeatureRowExtended) XXX_DiscardUnknown() {
	xxx_messageInfo_FeatureRowExtended.DiscardUnknown(m)
}

var xxx_messageInfo_FeatureRowExtended proto.InternalMessageInfo

func (m *FeatureRowExtended) GetRow() *FeatureRow {
	if m != nil {
		return m.Row
	}
	return nil
}

func (m *FeatureRowExtended) GetLastAttempt() *Attempt {
	if m != nil {
		return m.LastAttempt
	}
	return nil
}

func (m *FeatureRowExtended) GetFirstSeen() *timestamp.Timestamp {
	if m != nil {
		return m.FirstSeen
	}
	return nil
}

func init() {
	proto.RegisterType((*Error)(nil), "feast.types.Error")
	proto.RegisterType((*Attempt)(nil), "feast.types.Attempt")
	proto.RegisterType((*FeatureRowExtended)(nil), "feast.types.FeatureRowExtended")
}

func init() {
	proto.RegisterFile("feast/types/FeatureRowExtended.proto", fileDescriptor_7823aa2c72575793)
}

var fileDescriptor_7823aa2c72575793 = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x6b, 0xea, 0x40,
	0x10, 0xc6, 0xf1, 0xf9, 0xf2, 0x7c, 0x4e, 0x6e, 0x8b, 0x60, 0x08, 0xd2, 0x16, 0xe9, 0xc1, 0x5e,
	0x76, 0xc1, 0x82, 0xf4, 0x5a, 0xc1, 0x5e, 0x5b, 0xb6, 0x9e, 0x7a, 0x28, 0xac, 0x71, 0xb2, 0xb5,
	0x9a, 0x6c, 0xd8, 0x9d, 0x60, 0xfb, 0x67, 0xf5, 0x3f, 0x2c, 0xee, 0x6a, 0x8d, 0xd8, 0x5b, 0x66,
	0xbf, 0xdf, 0xcc, 0x7c, 0x99, 0x0f, 0xae, 0x73, 0x54, 0x8e, 0x04, 0x7d, 0x56, 0xe8, 0xc4, 0x03,
	0x2a, 0xaa, 0x2d, 0x4a, 0xb3, 0x9d, 0x7d, 0x10, 0x96, 0x4b, 0x5c, 0xf2, 0xca, 0x1a, 0x32, 0x2c,
	0xf6, 0x14, 0xf7, 0x54, 0x7a, 0xa9, 0x8d, 0xd1, 0x1b, 0x14, 0x5e, 0x5a, 0xd4, 0xb9, 0xa0, 0x55,
	0x81, 0x8e, 0x54, 0x51, 0x05, 0x3a, 0x1d, 0xfc, 0x3e, 0x33, 0xa8, 0xc3, 0x1a, 0xa2, 0x99, 0xb5,
	0xc6, 0xb2, 0x1e, 0x44, 0x99, 0xaa, 0x1d, 0x26, 0xad, 0xab, 0xd6, 0xa8, 0x2b, 0x43, 0xc1, 0x06,
	0xd0, 0x25, 0xab, 0x4a, 0x97, 0x1b, 0x5b, 0x24, 0x7f, 0xbc, 0x72, 0x7c, 0x60, 0x09, 0x74, 0x0a,
	0x74, 0x4e, 0x69, 0x4c, 0xda, 0x5e, 0x3b, 0x94, 0xec, 0x02, 0xc0, 0x91, 0xca, 0xd6, 0x73, 0xab,
	0x32, 0x4c, 0xfe, 0x7a, 0xb1, 0xf1, 0x32, 0x7c, 0x84, 0xce, 0x3d, 0x11, 0x16, 0x15, 0xb1, 0x14,
	0xfe, 0xab, 0xf0, 0xe9, 0xfc, 0xee, 0x48, 0xfe, 0xd4, 0x6c, 0x04, 0x11, 0xee, 0xdc, 0xf9, 0xd5,
	0xf1, 0x98, 0xf1, 0xc6, 0x9f, 0x73, 0xef, 0x5b, 0x06, 0x60, 0xf8, 0xd5, 0x02, 0x76, 0x7e, 0x30,
	0x76, 0x03, 0x6d, 0x6b, 0xb6, 0x7e, 0x6e, 0x3c, 0xee, 0x9f, 0xb4, 0x1f, 0x69, 0xb9, 0x63, 0xd8,
	0x04, 0xe2, 0x8d, 0x72, 0xb4, 0xb7, 0xb5, 0xdf, 0xd8, 0x3b, 0x69, 0xd9, 0x6b, 0xb2, 0x09, 0xb2,
	0x3b, 0xe8, 0xe6, 0x2b, 0xeb, 0xe8, 0x19, 0xb1, 0xf4, 0x67, 0x88, 0xc7, 0x29, 0x0f, 0xa1, 0xf0,
	0x43, 0x28, 0x7c, 0x7e, 0x08, 0x45, 0x1e, 0xe1, 0xe9, 0x2b, 0x34, 0x93, 0x9c, 0xf6, 0xcf, 0xfd,
	0x3f, 0xed, 0xfa, 0x5f, 0x26, 0x7a, 0x45, 0x6f, 0xf5, 0x82, 0x67, 0xa6, 0x10, 0xda, 0xbc, 0xe3,
	0x5a, 0x84, 0x48, 0xfd, 0x74, 0x27, 0x34, 0x96, 0x68, 0x15, 0xe1, 0x52, 0x68, 0x23, 0x1a, 0x61,
	0x2f, 0xfe, 0x79, 0xe0, 0xf6, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x99, 0x23, 0x9a, 0xf3, 0x56, 0x02,
	0x00, 0x00,
}