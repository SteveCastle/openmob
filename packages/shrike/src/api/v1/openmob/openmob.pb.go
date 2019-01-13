// Code generated by protoc-gen-go. DO NOT EDIT.
// source: openmob.proto

package openmob

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Cause struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cause) Reset()         { *m = Cause{} }
func (m *Cause) String() string { return proto.CompactTextString(m) }
func (*Cause) ProtoMessage()    {}
func (*Cause) Descriptor() ([]byte, []int) {
	return fileDescriptor_openmob_c053bd7fba7a1daa, []int{0}
}
func (m *Cause) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cause.Unmarshal(m, b)
}
func (m *Cause) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cause.Marshal(b, m, deterministic)
}
func (dst *Cause) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cause.Merge(dst, src)
}
func (m *Cause) XXX_Size() int {
	return xxx_messageInfo_Cause.Size(m)
}
func (m *Cause) XXX_DiscardUnknown() {
	xxx_messageInfo_Cause.DiscardUnknown(m)
}

var xxx_messageInfo_Cause proto.InternalMessageInfo

func (m *Cause) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Cause) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

type CreateCauseRequest struct {
	Item                 *Cause   `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCauseRequest) Reset()         { *m = CreateCauseRequest{} }
func (m *CreateCauseRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCauseRequest) ProtoMessage()    {}
func (*CreateCauseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_openmob_c053bd7fba7a1daa, []int{1}
}
func (m *CreateCauseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCauseRequest.Unmarshal(m, b)
}
func (m *CreateCauseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCauseRequest.Marshal(b, m, deterministic)
}
func (dst *CreateCauseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCauseRequest.Merge(dst, src)
}
func (m *CreateCauseRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCauseRequest.Size(m)
}
func (m *CreateCauseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCauseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCauseRequest proto.InternalMessageInfo

func (m *CreateCauseRequest) GetItem() *Cause {
	if m != nil {
		return m.Item
	}
	return nil
}

type CreateCauseResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCauseResponse) Reset()         { *m = CreateCauseResponse{} }
func (m *CreateCauseResponse) String() string { return proto.CompactTextString(m) }
func (*CreateCauseResponse) ProtoMessage()    {}
func (*CreateCauseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_openmob_c053bd7fba7a1daa, []int{2}
}
func (m *CreateCauseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCauseResponse.Unmarshal(m, b)
}
func (m *CreateCauseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCauseResponse.Marshal(b, m, deterministic)
}
func (dst *CreateCauseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCauseResponse.Merge(dst, src)
}
func (m *CreateCauseResponse) XXX_Size() int {
	return xxx_messageInfo_CreateCauseResponse.Size(m)
}
func (m *CreateCauseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCauseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCauseResponse proto.InternalMessageInfo

func (m *CreateCauseResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CreateCausesRequest struct {
	Items                []*Cause `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCausesRequest) Reset()         { *m = CreateCausesRequest{} }
func (m *CreateCausesRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCausesRequest) ProtoMessage()    {}
func (*CreateCausesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_openmob_c053bd7fba7a1daa, []int{3}
}
func (m *CreateCausesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCausesRequest.Unmarshal(m, b)
}
func (m *CreateCausesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCausesRequest.Marshal(b, m, deterministic)
}
func (dst *CreateCausesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCausesRequest.Merge(dst, src)
}
func (m *CreateCausesRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCausesRequest.Size(m)
}
func (m *CreateCausesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCausesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCausesRequest proto.InternalMessageInfo

func (m *CreateCausesRequest) GetItems() []*Cause {
	if m != nil {
		return m.Items
	}
	return nil
}

type CreateCausesResponse struct {
	Ids                  []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCausesResponse) Reset()         { *m = CreateCausesResponse{} }
func (m *CreateCausesResponse) String() string { return proto.CompactTextString(m) }
func (*CreateCausesResponse) ProtoMessage()    {}
func (*CreateCausesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_openmob_c053bd7fba7a1daa, []int{4}
}
func (m *CreateCausesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCausesResponse.Unmarshal(m, b)
}
func (m *CreateCausesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCausesResponse.Marshal(b, m, deterministic)
}
func (dst *CreateCausesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCausesResponse.Merge(dst, src)
}
func (m *CreateCausesResponse) XXX_Size() int {
	return xxx_messageInfo_CreateCausesResponse.Size(m)
}
func (m *CreateCausesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCausesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCausesResponse proto.InternalMessageInfo

func (m *CreateCausesResponse) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type GetCauseRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCauseRequest) Reset()         { *m = GetCauseRequest{} }
func (m *GetCauseRequest) String() string { return proto.CompactTextString(m) }
func (*GetCauseRequest) ProtoMessage()    {}
func (*GetCauseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_openmob_c053bd7fba7a1daa, []int{5}
}
func (m *GetCauseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCauseRequest.Unmarshal(m, b)
}
func (m *GetCauseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCauseRequest.Marshal(b, m, deterministic)
}
func (dst *GetCauseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCauseRequest.Merge(dst, src)
}
func (m *GetCauseRequest) XXX_Size() int {
	return xxx_messageInfo_GetCauseRequest.Size(m)
}
func (m *GetCauseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCauseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCauseRequest proto.InternalMessageInfo

func (m *GetCauseRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetCauseResponse struct {
	Item                 *Cause   `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCauseResponse) Reset()         { *m = GetCauseResponse{} }
func (m *GetCauseResponse) String() string { return proto.CompactTextString(m) }
func (*GetCauseResponse) ProtoMessage()    {}
func (*GetCauseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_openmob_c053bd7fba7a1daa, []int{6}
}
func (m *GetCauseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCauseResponse.Unmarshal(m, b)
}
func (m *GetCauseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCauseResponse.Marshal(b, m, deterministic)
}
func (dst *GetCauseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCauseResponse.Merge(dst, src)
}
func (m *GetCauseResponse) XXX_Size() int {
	return xxx_messageInfo_GetCauseResponse.Size(m)
}
func (m *GetCauseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCauseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetCauseResponse proto.InternalMessageInfo

func (m *GetCauseResponse) GetItem() *Cause {
	if m != nil {
		return m.Item
	}
	return nil
}

type ListCauseRequest struct {
	Limit                int32    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	NotCompleted         bool     `protobuf:"varint,2,opt,name=not_completed,json=notCompleted,proto3" json:"not_completed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListCauseRequest) Reset()         { *m = ListCauseRequest{} }
func (m *ListCauseRequest) String() string { return proto.CompactTextString(m) }
func (*ListCauseRequest) ProtoMessage()    {}
func (*ListCauseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_openmob_c053bd7fba7a1daa, []int{7}
}
func (m *ListCauseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCauseRequest.Unmarshal(m, b)
}
func (m *ListCauseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCauseRequest.Marshal(b, m, deterministic)
}
func (dst *ListCauseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCauseRequest.Merge(dst, src)
}
func (m *ListCauseRequest) XXX_Size() int {
	return xxx_messageInfo_ListCauseRequest.Size(m)
}
func (m *ListCauseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCauseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListCauseRequest proto.InternalMessageInfo

func (m *ListCauseRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListCauseRequest) GetNotCompleted() bool {
	if m != nil {
		return m.NotCompleted
	}
	return false
}

type ListCauseResponse struct {
	Items                []*Cause `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListCauseResponse) Reset()         { *m = ListCauseResponse{} }
func (m *ListCauseResponse) String() string { return proto.CompactTextString(m) }
func (*ListCauseResponse) ProtoMessage()    {}
func (*ListCauseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_openmob_c053bd7fba7a1daa, []int{8}
}
func (m *ListCauseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCauseResponse.Unmarshal(m, b)
}
func (m *ListCauseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCauseResponse.Marshal(b, m, deterministic)
}
func (dst *ListCauseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCauseResponse.Merge(dst, src)
}
func (m *ListCauseResponse) XXX_Size() int {
	return xxx_messageInfo_ListCauseResponse.Size(m)
}
func (m *ListCauseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCauseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListCauseResponse proto.InternalMessageInfo

func (m *ListCauseResponse) GetItems() []*Cause {
	if m != nil {
		return m.Items
	}
	return nil
}

type DeleteCauseRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteCauseRequest) Reset()         { *m = DeleteCauseRequest{} }
func (m *DeleteCauseRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteCauseRequest) ProtoMessage()    {}
func (*DeleteCauseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_openmob_c053bd7fba7a1daa, []int{9}
}
func (m *DeleteCauseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCauseRequest.Unmarshal(m, b)
}
func (m *DeleteCauseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCauseRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteCauseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCauseRequest.Merge(dst, src)
}
func (m *DeleteCauseRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteCauseRequest.Size(m)
}
func (m *DeleteCauseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCauseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCauseRequest proto.InternalMessageInfo

func (m *DeleteCauseRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteCauseResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteCauseResponse) Reset()         { *m = DeleteCauseResponse{} }
func (m *DeleteCauseResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteCauseResponse) ProtoMessage()    {}
func (*DeleteCauseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_openmob_c053bd7fba7a1daa, []int{10}
}
func (m *DeleteCauseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCauseResponse.Unmarshal(m, b)
}
func (m *DeleteCauseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCauseResponse.Marshal(b, m, deterministic)
}
func (dst *DeleteCauseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCauseResponse.Merge(dst, src)
}
func (m *DeleteCauseResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteCauseResponse.Size(m)
}
func (m *DeleteCauseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCauseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCauseResponse proto.InternalMessageInfo

type UpdateCauseRequest struct {
	Item                 *Cause   `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateCauseRequest) Reset()         { *m = UpdateCauseRequest{} }
func (m *UpdateCauseRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateCauseRequest) ProtoMessage()    {}
func (*UpdateCauseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_openmob_c053bd7fba7a1daa, []int{11}
}
func (m *UpdateCauseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCauseRequest.Unmarshal(m, b)
}
func (m *UpdateCauseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCauseRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateCauseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCauseRequest.Merge(dst, src)
}
func (m *UpdateCauseRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateCauseRequest.Size(m)
}
func (m *UpdateCauseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCauseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCauseRequest proto.InternalMessageInfo

func (m *UpdateCauseRequest) GetItem() *Cause {
	if m != nil {
		return m.Item
	}
	return nil
}

type UpdateCauseResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateCauseResponse) Reset()         { *m = UpdateCauseResponse{} }
func (m *UpdateCauseResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateCauseResponse) ProtoMessage()    {}
func (*UpdateCauseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_openmob_c053bd7fba7a1daa, []int{12}
}
func (m *UpdateCauseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCauseResponse.Unmarshal(m, b)
}
func (m *UpdateCauseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCauseResponse.Marshal(b, m, deterministic)
}
func (dst *UpdateCauseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCauseResponse.Merge(dst, src)
}
func (m *UpdateCauseResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateCauseResponse.Size(m)
}
func (m *UpdateCauseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCauseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCauseResponse proto.InternalMessageInfo

type UpdateCausesRequest struct {
	Items                []*Cause `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateCausesRequest) Reset()         { *m = UpdateCausesRequest{} }
func (m *UpdateCausesRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateCausesRequest) ProtoMessage()    {}
func (*UpdateCausesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_openmob_c053bd7fba7a1daa, []int{13}
}
func (m *UpdateCausesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCausesRequest.Unmarshal(m, b)
}
func (m *UpdateCausesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCausesRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateCausesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCausesRequest.Merge(dst, src)
}
func (m *UpdateCausesRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateCausesRequest.Size(m)
}
func (m *UpdateCausesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCausesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCausesRequest proto.InternalMessageInfo

func (m *UpdateCausesRequest) GetItems() []*Cause {
	if m != nil {
		return m.Items
	}
	return nil
}

type UpdateCausesResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateCausesResponse) Reset()         { *m = UpdateCausesResponse{} }
func (m *UpdateCausesResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateCausesResponse) ProtoMessage()    {}
func (*UpdateCausesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_openmob_c053bd7fba7a1daa, []int{14}
}
func (m *UpdateCausesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCausesResponse.Unmarshal(m, b)
}
func (m *UpdateCausesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCausesResponse.Marshal(b, m, deterministic)
}
func (dst *UpdateCausesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCausesResponse.Merge(dst, src)
}
func (m *UpdateCausesResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateCausesResponse.Size(m)
}
func (m *UpdateCausesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCausesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCausesResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Cause)(nil), "openmob.v1.Cause")
	proto.RegisterType((*CreateCauseRequest)(nil), "openmob.v1.CreateCauseRequest")
	proto.RegisterType((*CreateCauseResponse)(nil), "openmob.v1.CreateCauseResponse")
	proto.RegisterType((*CreateCausesRequest)(nil), "openmob.v1.CreateCausesRequest")
	proto.RegisterType((*CreateCausesResponse)(nil), "openmob.v1.CreateCausesResponse")
	proto.RegisterType((*GetCauseRequest)(nil), "openmob.v1.GetCauseRequest")
	proto.RegisterType((*GetCauseResponse)(nil), "openmob.v1.GetCauseResponse")
	proto.RegisterType((*ListCauseRequest)(nil), "openmob.v1.ListCauseRequest")
	proto.RegisterType((*ListCauseResponse)(nil), "openmob.v1.ListCauseResponse")
	proto.RegisterType((*DeleteCauseRequest)(nil), "openmob.v1.DeleteCauseRequest")
	proto.RegisterType((*DeleteCauseResponse)(nil), "openmob.v1.DeleteCauseResponse")
	proto.RegisterType((*UpdateCauseRequest)(nil), "openmob.v1.UpdateCauseRequest")
	proto.RegisterType((*UpdateCauseResponse)(nil), "openmob.v1.UpdateCauseResponse")
	proto.RegisterType((*UpdateCausesRequest)(nil), "openmob.v1.UpdateCausesRequest")
	proto.RegisterType((*UpdateCausesResponse)(nil), "openmob.v1.UpdateCausesResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CauseServiceClient is the client API for CauseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CauseServiceClient interface {
	CreateCause(ctx context.Context, in *CreateCauseRequest, opts ...grpc.CallOption) (*CreateCauseResponse, error)
	// Bulk version of CreateCause
	CreateCauses(ctx context.Context, in *CreateCausesRequest, opts ...grpc.CallOption) (*CreateCausesResponse, error)
	GetCause(ctx context.Context, in *GetCauseRequest, opts ...grpc.CallOption) (*GetCauseResponse, error)
	ListCause(ctx context.Context, in *ListCauseRequest, opts ...grpc.CallOption) (*ListCauseResponse, error)
	DeleteCause(ctx context.Context, in *DeleteCauseRequest, opts ...grpc.CallOption) (*DeleteCauseResponse, error)
	UpdateCause(ctx context.Context, in *UpdateCauseRequest, opts ...grpc.CallOption) (*UpdateCauseResponse, error)
	UpdateCauses(ctx context.Context, in *UpdateCausesRequest, opts ...grpc.CallOption) (*UpdateCausesResponse, error)
}

type causeServiceClient struct {
	cc *grpc.ClientConn
}

func NewCauseServiceClient(cc *grpc.ClientConn) CauseServiceClient {
	return &causeServiceClient{cc}
}

func (c *causeServiceClient) CreateCause(ctx context.Context, in *CreateCauseRequest, opts ...grpc.CallOption) (*CreateCauseResponse, error) {
	out := new(CreateCauseResponse)
	err := c.cc.Invoke(ctx, "/openmob.v1.CauseService/CreateCause", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *causeServiceClient) CreateCauses(ctx context.Context, in *CreateCausesRequest, opts ...grpc.CallOption) (*CreateCausesResponse, error) {
	out := new(CreateCausesResponse)
	err := c.cc.Invoke(ctx, "/openmob.v1.CauseService/CreateCauses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *causeServiceClient) GetCause(ctx context.Context, in *GetCauseRequest, opts ...grpc.CallOption) (*GetCauseResponse, error) {
	out := new(GetCauseResponse)
	err := c.cc.Invoke(ctx, "/openmob.v1.CauseService/GetCause", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *causeServiceClient) ListCause(ctx context.Context, in *ListCauseRequest, opts ...grpc.CallOption) (*ListCauseResponse, error) {
	out := new(ListCauseResponse)
	err := c.cc.Invoke(ctx, "/openmob.v1.CauseService/ListCause", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *causeServiceClient) DeleteCause(ctx context.Context, in *DeleteCauseRequest, opts ...grpc.CallOption) (*DeleteCauseResponse, error) {
	out := new(DeleteCauseResponse)
	err := c.cc.Invoke(ctx, "/openmob.v1.CauseService/DeleteCause", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *causeServiceClient) UpdateCause(ctx context.Context, in *UpdateCauseRequest, opts ...grpc.CallOption) (*UpdateCauseResponse, error) {
	out := new(UpdateCauseResponse)
	err := c.cc.Invoke(ctx, "/openmob.v1.CauseService/UpdateCause", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *causeServiceClient) UpdateCauses(ctx context.Context, in *UpdateCausesRequest, opts ...grpc.CallOption) (*UpdateCausesResponse, error) {
	out := new(UpdateCausesResponse)
	err := c.cc.Invoke(ctx, "/openmob.v1.CauseService/UpdateCauses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CauseServiceServer is the server API for CauseService service.
type CauseServiceServer interface {
	CreateCause(context.Context, *CreateCauseRequest) (*CreateCauseResponse, error)
	// Bulk version of CreateCause
	CreateCauses(context.Context, *CreateCausesRequest) (*CreateCausesResponse, error)
	GetCause(context.Context, *GetCauseRequest) (*GetCauseResponse, error)
	ListCause(context.Context, *ListCauseRequest) (*ListCauseResponse, error)
	DeleteCause(context.Context, *DeleteCauseRequest) (*DeleteCauseResponse, error)
	UpdateCause(context.Context, *UpdateCauseRequest) (*UpdateCauseResponse, error)
	UpdateCauses(context.Context, *UpdateCausesRequest) (*UpdateCausesResponse, error)
}

func RegisterCauseServiceServer(s *grpc.Server, srv CauseServiceServer) {
	s.RegisterService(&_CauseService_serviceDesc, srv)
}

func _CauseService_CreateCause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCauseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CauseServiceServer).CreateCause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openmob.v1.CauseService/CreateCause",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CauseServiceServer).CreateCause(ctx, req.(*CreateCauseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CauseService_CreateCauses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCausesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CauseServiceServer).CreateCauses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openmob.v1.CauseService/CreateCauses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CauseServiceServer).CreateCauses(ctx, req.(*CreateCausesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CauseService_GetCause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCauseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CauseServiceServer).GetCause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openmob.v1.CauseService/GetCause",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CauseServiceServer).GetCause(ctx, req.(*GetCauseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CauseService_ListCause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCauseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CauseServiceServer).ListCause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openmob.v1.CauseService/ListCause",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CauseServiceServer).ListCause(ctx, req.(*ListCauseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CauseService_DeleteCause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCauseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CauseServiceServer).DeleteCause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openmob.v1.CauseService/DeleteCause",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CauseServiceServer).DeleteCause(ctx, req.(*DeleteCauseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CauseService_UpdateCause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCauseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CauseServiceServer).UpdateCause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openmob.v1.CauseService/UpdateCause",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CauseServiceServer).UpdateCause(ctx, req.(*UpdateCauseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CauseService_UpdateCauses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCausesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CauseServiceServer).UpdateCauses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openmob.v1.CauseService/UpdateCauses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CauseServiceServer).UpdateCauses(ctx, req.(*UpdateCausesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CauseService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openmob.v1.CauseService",
	HandlerType: (*CauseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCause",
			Handler:    _CauseService_CreateCause_Handler,
		},
		{
			MethodName: "CreateCauses",
			Handler:    _CauseService_CreateCauses_Handler,
		},
		{
			MethodName: "GetCause",
			Handler:    _CauseService_GetCause_Handler,
		},
		{
			MethodName: "ListCause",
			Handler:    _CauseService_ListCause_Handler,
		},
		{
			MethodName: "DeleteCause",
			Handler:    _CauseService_DeleteCause_Handler,
		},
		{
			MethodName: "UpdateCause",
			Handler:    _CauseService_UpdateCause_Handler,
		},
		{
			MethodName: "UpdateCauses",
			Handler:    _CauseService_UpdateCauses_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "openmob.proto",
}

func init() { proto.RegisterFile("openmob.proto", fileDescriptor_openmob_c053bd7fba7a1daa) }

var fileDescriptor_openmob_c053bd7fba7a1daa = []byte{
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x55, 0x52, 0x0c, 0xf5, 0x38, 0x2d, 0xc9, 0x24, 0x2d, 0xc1, 0x14, 0x08, 0x0b, 0x15, 0x55,
	0x25, 0x62, 0xb5, 0x9c, 0x28, 0xa8, 0x07, 0x82, 0xc4, 0x05, 0x2e, 0x46, 0x5c, 0x40, 0x08, 0x39,
	0xf1, 0x12, 0xad, 0x70, 0xbc, 0x26, 0xde, 0xe4, 0x82, 0xb8, 0xf0, 0x17, 0xf8, 0x69, 0xfc, 0x01,
	0x0e, 0xfc, 0x10, 0xe4, 0xf5, 0xba, 0xd9, 0xf5, 0x07, 0x2a, 0xdc, 0xbc, 0x33, 0xf3, 0xde, 0xdb,
	0xd9, 0xf7, 0x64, 0xd8, 0xe1, 0x09, 0x8d, 0x17, 0x7c, 0x3a, 0x4e, 0x96, 0x5c, 0x70, 0x84, 0xe2,
	0xb8, 0x3e, 0x71, 0x0f, 0xe6, 0x9c, 0xcf, 0x23, 0xea, 0x05, 0x09, 0xf3, 0x82, 0x38, 0xe6, 0x22,
	0x10, 0x8c, 0xc7, 0x69, 0x3e, 0x49, 0x1e, 0x81, 0x35, 0x09, 0x56, 0x29, 0xc5, 0x5d, 0x68, 0xb3,
	0x70, 0xd8, 0x1a, 0xb5, 0x8e, 0x6c, 0xbf, 0xcd, 0x42, 0x1c, 0x80, 0x25, 0x98, 0x88, 0xe8, 0xb0,
	0x2d, 0x4b, 0xf9, 0x81, 0x3c, 0x05, 0x9c, 0x2c, 0x69, 0x20, 0xa8, 0x04, 0xf9, 0xf4, 0xcb, 0x8a,
	0xa6, 0x02, 0x0f, 0xe1, 0x0a, 0x13, 0x74, 0x21, 0xd1, 0xce, 0x69, 0x6f, 0xbc, 0x51, 0x1f, 0xe7,
	0x73, 0xb2, 0x4d, 0x0e, 0xa1, 0x6f, 0x80, 0xd3, 0x84, 0xc7, 0x55, 0x65, 0x72, 0x6e, 0x8c, 0xa5,
	0x85, 0xc8, 0x43, 0xb0, 0x32, 0x96, 0x74, 0xd8, 0x1a, 0x6d, 0xd5, 0xab, 0xe4, 0x7d, 0x72, 0x04,
	0x03, 0x13, 0xaf, 0x74, 0xba, 0xb0, 0xc5, 0xc2, 0x1c, 0x6e, 0xfb, 0xd9, 0x27, 0xb9, 0x07, 0xd7,
	0x5f, 0x52, 0x61, 0xac, 0x52, 0xbe, 0xcc, 0x13, 0xe8, 0x6e, 0x46, 0x14, 0xd1, 0x25, 0xd7, 0x7d,
	0x0d, 0xdd, 0x57, 0x2c, 0x35, 0xe9, 0x07, 0x60, 0x45, 0x6c, 0xc1, 0x84, 0xc4, 0x5a, 0x7e, 0x7e,
	0xc0, 0xfb, 0xb0, 0x13, 0x73, 0xf1, 0x71, 0xc6, 0x17, 0x49, 0x44, 0x05, 0x0d, 0xe5, 0x9b, 0x6f,
	0xfb, 0x9d, 0x98, 0x8b, 0x49, 0x51, 0x23, 0xcf, 0xa0, 0xa7, 0xd1, 0xa9, 0xab, 0x5c, 0xfa, 0x51,
	0x1e, 0x00, 0xbe, 0xa0, 0x19, 0xd1, 0x5f, 0xb7, 0xdd, 0x83, 0xbe, 0x31, 0x95, 0xab, 0x64, 0xae,
	0xbf, 0x4d, 0xc2, 0xff, 0x74, 0x7d, 0x0f, 0xfa, 0x06, 0x58, 0x71, 0x9e, 0x1b, 0xe5, 0x7f, 0x77,
	0x79, 0x1f, 0x06, 0x26, 0x3e, 0xe7, 0x3d, 0xfd, 0x65, 0x41, 0x47, 0x96, 0xde, 0xd0, 0xe5, 0x9a,
	0xcd, 0x28, 0xce, 0xc1, 0xd1, 0xe2, 0x80, 0x77, 0x0c, 0xc6, 0x4a, 0x96, 0xdd, 0xbb, 0x8d, 0x7d,
	0x75, 0xf1, 0x1b, 0xdf, 0x7f, 0xfe, 0xfe, 0xd1, 0xee, 0x11, 0xdb, 0x5b, 0x9f, 0x78, 0xb3, 0xac,
	0x75, 0x26, 0x17, 0xc5, 0x08, 0x3a, 0x7a, 0xee, 0xb0, 0x89, 0xa9, 0xd8, 0xd5, 0x1d, 0x35, 0x0f,
	0x28, 0xad, 0x9b, 0x52, 0xab, 0x4f, 0x76, 0x2f, 0xb4, 0xbc, 0xe9, 0x2a, 0xfa, 0x7c, 0xd6, 0x3a,
	0xc6, 0x0f, 0xb0, 0x5d, 0x04, 0x13, 0x6f, 0xe9, 0x44, 0xa5, 0x44, 0xbb, 0x07, 0xf5, 0x4d, 0xa5,
	0xb0, 0x2f, 0x15, 0xba, 0xa8, 0x29, 0x7c, 0x65, 0xe1, 0x37, 0x7c, 0x0f, 0xf6, 0x45, 0xda, 0xd0,
	0xa0, 0x28, 0x67, 0xda, 0xbd, 0xdd, 0xd0, 0x55, 0x0a, 0x3d, 0xa9, 0xe0, 0xe0, 0xe6, 0xbd, 0xf0,
	0x13, 0x38, 0x5a, 0xcc, 0x4c, 0x4b, 0xaa, 0x29, 0x35, 0x2d, 0xa9, 0xcb, 0xa7, 0x5a, 0xe2, 0xb8,
	0xbc, 0xc4, 0x1c, 0x1c, 0x2d, 0x23, 0xa6, 0x4e, 0x35, 0xd0, 0xa6, 0x4e, 0x5d, 0x66, 0x95, 0xf5,
	0x6e, 0x9d, 0xf5, 0x7a, 0x18, 0xb1, 0x89, 0xa9, 0xde, 0xfa, 0xba, 0x1c, 0x17, 0xd6, 0xbb, 0x55,
	0xeb, 0x9f, 0xdb, 0xef, 0xae, 0x29, 0xf4, 0xf4, 0xaa, 0xfc, 0x8b, 0x3f, 0xfe, 0x13, 0x00, 0x00,
	0xff, 0xff, 0xb9, 0xe7, 0x93, 0xb1, 0x00, 0x06, 0x00, 0x00,
}
