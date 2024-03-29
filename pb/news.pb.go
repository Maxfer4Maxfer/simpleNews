// Code generated by protoc-gen-go. DO NOT EDIT.
// source: news.proto

package pb

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

type News struct {
	Uuid                 string               `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Title                string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *News) Reset()         { *m = News{} }
func (m *News) String() string { return proto.CompactTextString(m) }
func (*News) ProtoMessage()    {}
func (*News) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c0382e93bed6d84, []int{0}
}

func (m *News) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_News.Unmarshal(m, b)
}
func (m *News) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_News.Marshal(b, m, deterministic)
}
func (m *News) XXX_Merge(src proto.Message) {
	xxx_messageInfo_News.Merge(m, src)
}
func (m *News) XXX_Size() int {
	return xxx_messageInfo_News.Size(m)
}
func (m *News) XXX_DiscardUnknown() {
	xxx_messageInfo_News.DiscardUnknown(m)
}

var xxx_messageInfo_News proto.InternalMessageInfo

func (m *News) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *News) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *News) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type NewsResponce struct {
	News                 *News    `protobuf:"bytes,1,opt,name=news,proto3" json:"news,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewsResponce) Reset()         { *m = NewsResponce{} }
func (m *NewsResponce) String() string { return proto.CompactTextString(m) }
func (*NewsResponce) ProtoMessage()    {}
func (*NewsResponce) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c0382e93bed6d84, []int{1}
}

func (m *NewsResponce) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewsResponce.Unmarshal(m, b)
}
func (m *NewsResponce) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewsResponce.Marshal(b, m, deterministic)
}
func (m *NewsResponce) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewsResponce.Merge(m, src)
}
func (m *NewsResponce) XXX_Size() int {
	return xxx_messageInfo_NewsResponce.Size(m)
}
func (m *NewsResponce) XXX_DiscardUnknown() {
	xxx_messageInfo_NewsResponce.DiscardUnknown(m)
}

var xxx_messageInfo_NewsResponce proto.InternalMessageInfo

func (m *NewsResponce) GetNews() *News {
	if m != nil {
		return m.News
	}
	return nil
}

func (m *NewsResponce) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type SaveNewsResponce struct {
	Err                  string   `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaveNewsResponce) Reset()         { *m = SaveNewsResponce{} }
func (m *SaveNewsResponce) String() string { return proto.CompactTextString(m) }
func (*SaveNewsResponce) ProtoMessage()    {}
func (*SaveNewsResponce) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c0382e93bed6d84, []int{2}
}

func (m *SaveNewsResponce) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveNewsResponce.Unmarshal(m, b)
}
func (m *SaveNewsResponce) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveNewsResponce.Marshal(b, m, deterministic)
}
func (m *SaveNewsResponce) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveNewsResponce.Merge(m, src)
}
func (m *SaveNewsResponce) XXX_Size() int {
	return xxx_messageInfo_SaveNewsResponce.Size(m)
}
func (m *SaveNewsResponce) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveNewsResponce.DiscardUnknown(m)
}

var xxx_messageInfo_SaveNewsResponce proto.InternalMessageInfo

func (m *SaveNewsResponce) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type AllNewsResponce struct {
	News                 []*News  `protobuf:"bytes,1,rep,name=news,proto3" json:"news,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AllNewsResponce) Reset()         { *m = AllNewsResponce{} }
func (m *AllNewsResponce) String() string { return proto.CompactTextString(m) }
func (*AllNewsResponce) ProtoMessage()    {}
func (*AllNewsResponce) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c0382e93bed6d84, []int{3}
}

func (m *AllNewsResponce) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllNewsResponce.Unmarshal(m, b)
}
func (m *AllNewsResponce) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllNewsResponce.Marshal(b, m, deterministic)
}
func (m *AllNewsResponce) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllNewsResponce.Merge(m, src)
}
func (m *AllNewsResponce) XXX_Size() int {
	return xxx_messageInfo_AllNewsResponce.Size(m)
}
func (m *AllNewsResponce) XXX_DiscardUnknown() {
	xxx_messageInfo_AllNewsResponce.DiscardUnknown(m)
}

var xxx_messageInfo_AllNewsResponce proto.InternalMessageInfo

func (m *AllNewsResponce) GetNews() []*News {
	if m != nil {
		return m.News
	}
	return nil
}

func (m *AllNewsResponce) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type DeleteAllNewsResponce struct {
	Err                  string   `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteAllNewsResponce) Reset()         { *m = DeleteAllNewsResponce{} }
func (m *DeleteAllNewsResponce) String() string { return proto.CompactTextString(m) }
func (*DeleteAllNewsResponce) ProtoMessage()    {}
func (*DeleteAllNewsResponce) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c0382e93bed6d84, []int{4}
}

func (m *DeleteAllNewsResponce) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteAllNewsResponce.Unmarshal(m, b)
}
func (m *DeleteAllNewsResponce) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteAllNewsResponce.Marshal(b, m, deterministic)
}
func (m *DeleteAllNewsResponce) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteAllNewsResponce.Merge(m, src)
}
func (m *DeleteAllNewsResponce) XXX_Size() int {
	return xxx_messageInfo_DeleteAllNewsResponce.Size(m)
}
func (m *DeleteAllNewsResponce) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteAllNewsResponce.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteAllNewsResponce proto.InternalMessageInfo

func (m *DeleteAllNewsResponce) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*News)(nil), "pb.News")
	proto.RegisterType((*NewsResponce)(nil), "pb.NewsResponce")
	proto.RegisterType((*SaveNewsResponce)(nil), "pb.SaveNewsResponce")
	proto.RegisterType((*AllNewsResponce)(nil), "pb.AllNewsResponce")
	proto.RegisterType((*DeleteAllNewsResponce)(nil), "pb.DeleteAllNewsResponce")
}

func init() { proto.RegisterFile("news.proto", fileDescriptor_2c0382e93bed6d84) }

var fileDescriptor_2c0382e93bed6d84 = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x8f, 0x41, 0x4b, 0x86, 0x30,
	0x1c, 0xc6, 0x99, 0x5a, 0xe4, 0xdf, 0x20, 0x19, 0x05, 0x22, 0x41, 0x22, 0x1d, 0xec, 0x32, 0xc1,
	0x2e, 0x9d, 0x02, 0xa1, 0x73, 0x07, 0xeb, 0x0b, 0xb8, 0xfa, 0x27, 0xc6, 0x74, 0xc3, 0xcd, 0xfc,
	0xfa, 0xe1, 0xd4, 0xa4, 0xe2, 0x85, 0xf7, 0xf6, 0xec, 0xd9, 0x8f, 0xe7, 0xb7, 0x01, 0xf4, 0x38,
	0x69, 0xa6, 0x06, 0x69, 0x24, 0x75, 0x14, 0x8f, 0x6f, 0x1a, 0x29, 0x1b, 0x81, 0xb9, 0x6d, 0xf8,
	0xf8, 0x91, 0x9b, 0xb6, 0x43, 0x6d, 0xea, 0x4e, 0x2d, 0x50, 0xfa, 0x09, 0xde, 0x33, 0x4e, 0x9a,
	0x52, 0xf0, 0xc6, 0xb1, 0x7d, 0x8f, 0x48, 0x42, 0x32, 0xbf, 0xb2, 0x99, 0x5e, 0xc2, 0x89, 0x69,
	0x8d, 0xc0, 0xc8, 0xb1, 0xe5, 0x72, 0xa0, 0x0f, 0xe0, 0xff, 0x8c, 0x44, 0x6e, 0x42, 0xb2, 0xa0,
	0x88, 0xd9, 0xa2, 0x61, 0x9b, 0x86, 0xbd, 0x6e, 0x44, 0xb5, 0xc3, 0xe9, 0x23, 0x9c, 0xcf, 0xae,
	0x0a, 0xb5, 0x92, 0xfd, 0x1b, 0xd2, 0x6b, 0xf0, 0xe6, 0xe7, 0x5a, 0x67, 0x50, 0x9c, 0x31, 0xc5,
	0x99, 0xbd, 0xb7, 0x2d, 0x0d, 0xc1, 0xc5, 0x61, 0x58, 0xdd, 0x73, 0x4c, 0x6f, 0x21, 0x7c, 0xa9,
	0xbf, 0xf0, 0xd7, 0xc6, 0x4a, 0x91, 0x9d, 0x2a, 0xe1, 0xa2, 0x14, 0xe2, 0x80, 0xc8, 0x3d, 0x4a,
	0x74, 0x07, 0x57, 0x4f, 0x28, 0xd0, 0xe0, 0xdf, 0xa1, 0x7f, 0x36, 0x7e, 0x6a, 0xbf, 0x7c, 0xff,
	0x1d, 0x00, 0x00, 0xff, 0xff, 0x9c, 0xd8, 0xc2, 0x0d, 0x79, 0x01, 0x00, 0x00,
}
