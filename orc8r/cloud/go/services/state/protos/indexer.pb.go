// Code generated by protoc-gen-go. DO NOT EDIT.
// source: indexer.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protos "magma/orc8r/lib/go/protos"
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

type GetIndexerInfoRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetIndexerInfoRequest) Reset()         { *m = GetIndexerInfoRequest{} }
func (m *GetIndexerInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetIndexerInfoRequest) ProtoMessage()    {}
func (*GetIndexerInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b06a290ab031ed6, []int{0}
}

func (m *GetIndexerInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetIndexerInfoRequest.Unmarshal(m, b)
}
func (m *GetIndexerInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetIndexerInfoRequest.Marshal(b, m, deterministic)
}
func (m *GetIndexerInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetIndexerInfoRequest.Merge(m, src)
}
func (m *GetIndexerInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetIndexerInfoRequest.Size(m)
}
func (m *GetIndexerInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetIndexerInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetIndexerInfoRequest proto.InternalMessageInfo

type GetIndexerInfoResponse struct {
	// version of the indexer's current implementation
	Version uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// state_types is the type of states the indexer will index
	// utilized for pre-filtering
	StateTypes           []string `protobuf:"bytes,2,rep,name=state_types,json=stateTypes,proto3" json:"state_types,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetIndexerInfoResponse) Reset()         { *m = GetIndexerInfoResponse{} }
func (m *GetIndexerInfoResponse) String() string { return proto.CompactTextString(m) }
func (*GetIndexerInfoResponse) ProtoMessage()    {}
func (*GetIndexerInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b06a290ab031ed6, []int{1}
}

func (m *GetIndexerInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetIndexerInfoResponse.Unmarshal(m, b)
}
func (m *GetIndexerInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetIndexerInfoResponse.Marshal(b, m, deterministic)
}
func (m *GetIndexerInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetIndexerInfoResponse.Merge(m, src)
}
func (m *GetIndexerInfoResponse) XXX_Size() int {
	return xxx_messageInfo_GetIndexerInfoResponse.Size(m)
}
func (m *GetIndexerInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetIndexerInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetIndexerInfoResponse proto.InternalMessageInfo

func (m *GetIndexerInfoResponse) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *GetIndexerInfoResponse) GetStateTypes() []string {
	if m != nil {
		return m.StateTypes
	}
	return nil
}

type IndexRequest struct {
	// states to reindex
	States []*protos.State `protobuf:"bytes,1,rep,name=states,proto3" json:"states,omitempty"`
	// network_id of the states
	NetworkId string `protobuf:"bytes,2,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	// reporter_hwid is the hwid of the gateway that reported the states
	ReporterHwid         string   `protobuf:"bytes,3,opt,name=reporter_hwid,json=reporterHwid,proto3" json:"reporter_hwid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IndexRequest) Reset()         { *m = IndexRequest{} }
func (m *IndexRequest) String() string { return proto.CompactTextString(m) }
func (*IndexRequest) ProtoMessage()    {}
func (*IndexRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b06a290ab031ed6, []int{2}
}

func (m *IndexRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexRequest.Unmarshal(m, b)
}
func (m *IndexRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexRequest.Marshal(b, m, deterministic)
}
func (m *IndexRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexRequest.Merge(m, src)
}
func (m *IndexRequest) XXX_Size() int {
	return xxx_messageInfo_IndexRequest.Size(m)
}
func (m *IndexRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IndexRequest proto.InternalMessageInfo

func (m *IndexRequest) GetStates() []*protos.State {
	if m != nil {
		return m.States
	}
	return nil
}

func (m *IndexRequest) GetNetworkId() string {
	if m != nil {
		return m.NetworkId
	}
	return ""
}

func (m *IndexRequest) GetReporterHwid() string {
	if m != nil {
		return m.ReporterHwid
	}
	return ""
}

type IndexResponse struct {
	// state_errors are errors experienced trying to index specific pieces of state.
	StateErrors          []*protos.IDAndError `protobuf:"bytes,1,rep,name=state_errors,json=stateErrors,proto3" json:"state_errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *IndexResponse) Reset()         { *m = IndexResponse{} }
func (m *IndexResponse) String() string { return proto.CompactTextString(m) }
func (*IndexResponse) ProtoMessage()    {}
func (*IndexResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b06a290ab031ed6, []int{3}
}

func (m *IndexResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexResponse.Unmarshal(m, b)
}
func (m *IndexResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexResponse.Marshal(b, m, deterministic)
}
func (m *IndexResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexResponse.Merge(m, src)
}
func (m *IndexResponse) XXX_Size() int {
	return xxx_messageInfo_IndexResponse.Size(m)
}
func (m *IndexResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IndexResponse proto.InternalMessageInfo

func (m *IndexResponse) GetStateErrors() []*protos.IDAndError {
	if m != nil {
		return m.StateErrors
	}
	return nil
}

type PrepareReindexRequest struct {
	// indexer_id being reindexed
	IndexerId string `protobuf:"bytes,1,opt,name=indexer_id,json=indexerId,proto3" json:"indexer_id,omitempty"`
	// from_version is the indexer's current (actual) version
	FromVersion uint32 `protobuf:"varint,2,opt,name=from_version,json=fromVersion,proto3" json:"from_version,omitempty"`
	// to_version is the indexer's future (desired) version
	ToVersion uint32 `protobuf:"varint,3,opt,name=to_version,json=toVersion,proto3" json:"to_version,omitempty"`
	// is_first is true iff this is the first time this indexer is being reindexed
	IsFirst              bool     `protobuf:"varint,4,opt,name=is_first,json=isFirst,proto3" json:"is_first,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrepareReindexRequest) Reset()         { *m = PrepareReindexRequest{} }
func (m *PrepareReindexRequest) String() string { return proto.CompactTextString(m) }
func (*PrepareReindexRequest) ProtoMessage()    {}
func (*PrepareReindexRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b06a290ab031ed6, []int{4}
}

func (m *PrepareReindexRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrepareReindexRequest.Unmarshal(m, b)
}
func (m *PrepareReindexRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrepareReindexRequest.Marshal(b, m, deterministic)
}
func (m *PrepareReindexRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrepareReindexRequest.Merge(m, src)
}
func (m *PrepareReindexRequest) XXX_Size() int {
	return xxx_messageInfo_PrepareReindexRequest.Size(m)
}
func (m *PrepareReindexRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrepareReindexRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrepareReindexRequest proto.InternalMessageInfo

func (m *PrepareReindexRequest) GetIndexerId() string {
	if m != nil {
		return m.IndexerId
	}
	return ""
}

func (m *PrepareReindexRequest) GetFromVersion() uint32 {
	if m != nil {
		return m.FromVersion
	}
	return 0
}

func (m *PrepareReindexRequest) GetToVersion() uint32 {
	if m != nil {
		return m.ToVersion
	}
	return 0
}

func (m *PrepareReindexRequest) GetIsFirst() bool {
	if m != nil {
		return m.IsFirst
	}
	return false
}

type PrepareReindexResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrepareReindexResponse) Reset()         { *m = PrepareReindexResponse{} }
func (m *PrepareReindexResponse) String() string { return proto.CompactTextString(m) }
func (*PrepareReindexResponse) ProtoMessage()    {}
func (*PrepareReindexResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b06a290ab031ed6, []int{5}
}

func (m *PrepareReindexResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrepareReindexResponse.Unmarshal(m, b)
}
func (m *PrepareReindexResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrepareReindexResponse.Marshal(b, m, deterministic)
}
func (m *PrepareReindexResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrepareReindexResponse.Merge(m, src)
}
func (m *PrepareReindexResponse) XXX_Size() int {
	return xxx_messageInfo_PrepareReindexResponse.Size(m)
}
func (m *PrepareReindexResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PrepareReindexResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PrepareReindexResponse proto.InternalMessageInfo

type CompleteReindexRequest struct {
	// indexer_id being reindexed
	IndexerId string `protobuf:"bytes,1,opt,name=indexer_id,json=indexerId,proto3" json:"indexer_id,omitempty"`
	// from_version is the indexer's current (actual) version
	FromVersion uint32 `protobuf:"varint,2,opt,name=from_version,json=fromVersion,proto3" json:"from_version,omitempty"`
	// to_version is the indexer's future (desired) version
	ToVersion            uint32   `protobuf:"varint,3,opt,name=to_version,json=toVersion,proto3" json:"to_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CompleteReindexRequest) Reset()         { *m = CompleteReindexRequest{} }
func (m *CompleteReindexRequest) String() string { return proto.CompactTextString(m) }
func (*CompleteReindexRequest) ProtoMessage()    {}
func (*CompleteReindexRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b06a290ab031ed6, []int{6}
}

func (m *CompleteReindexRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompleteReindexRequest.Unmarshal(m, b)
}
func (m *CompleteReindexRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompleteReindexRequest.Marshal(b, m, deterministic)
}
func (m *CompleteReindexRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompleteReindexRequest.Merge(m, src)
}
func (m *CompleteReindexRequest) XXX_Size() int {
	return xxx_messageInfo_CompleteReindexRequest.Size(m)
}
func (m *CompleteReindexRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CompleteReindexRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CompleteReindexRequest proto.InternalMessageInfo

func (m *CompleteReindexRequest) GetIndexerId() string {
	if m != nil {
		return m.IndexerId
	}
	return ""
}

func (m *CompleteReindexRequest) GetFromVersion() uint32 {
	if m != nil {
		return m.FromVersion
	}
	return 0
}

func (m *CompleteReindexRequest) GetToVersion() uint32 {
	if m != nil {
		return m.ToVersion
	}
	return 0
}

type CompleteReindexResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CompleteReindexResponse) Reset()         { *m = CompleteReindexResponse{} }
func (m *CompleteReindexResponse) String() string { return proto.CompactTextString(m) }
func (*CompleteReindexResponse) ProtoMessage()    {}
func (*CompleteReindexResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b06a290ab031ed6, []int{7}
}

func (m *CompleteReindexResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompleteReindexResponse.Unmarshal(m, b)
}
func (m *CompleteReindexResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompleteReindexResponse.Marshal(b, m, deterministic)
}
func (m *CompleteReindexResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompleteReindexResponse.Merge(m, src)
}
func (m *CompleteReindexResponse) XXX_Size() int {
	return xxx_messageInfo_CompleteReindexResponse.Size(m)
}
func (m *CompleteReindexResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CompleteReindexResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CompleteReindexResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GetIndexerInfoRequest)(nil), "magma.orc8r.state.GetIndexerInfoRequest")
	proto.RegisterType((*GetIndexerInfoResponse)(nil), "magma.orc8r.state.GetIndexerInfoResponse")
	proto.RegisterType((*IndexRequest)(nil), "magma.orc8r.state.IndexRequest")
	proto.RegisterType((*IndexResponse)(nil), "magma.orc8r.state.IndexResponse")
	proto.RegisterType((*PrepareReindexRequest)(nil), "magma.orc8r.state.PrepareReindexRequest")
	proto.RegisterType((*PrepareReindexResponse)(nil), "magma.orc8r.state.PrepareReindexResponse")
	proto.RegisterType((*CompleteReindexRequest)(nil), "magma.orc8r.state.CompleteReindexRequest")
	proto.RegisterType((*CompleteReindexResponse)(nil), "magma.orc8r.state.CompleteReindexResponse")
}

func init() { proto.RegisterFile("indexer.proto", fileDescriptor_2b06a290ab031ed6) }

var fileDescriptor_2b06a290ab031ed6 = []byte{
	// 475 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xad, 0x13, 0x68, 0x92, 0x49, 0x52, 0xc4, 0x4a, 0x4d, 0xb6, 0x96, 0xa2, 0x1a, 0x73, 0x31,
	0x3d, 0xb8, 0xa8, 0xb9, 0x20, 0x6e, 0x7c, 0x63, 0xc1, 0x01, 0xb9, 0x88, 0x03, 0x17, 0x2b, 0xd4,
	0x93, 0xb0, 0x40, 0xbc, 0x66, 0x76, 0x69, 0xe8, 0x85, 0x3f, 0x81, 0xf8, 0xbf, 0xc8, 0xbb, 0xeb,
	0x2a, 0x49, 0x5d, 0xa9, 0xa7, 0x9e, 0x92, 0x9d, 0xf7, 0x66, 0xe6, 0xcd, 0xcc, 0x93, 0x61, 0x28,
	0x8a, 0x1c, 0x7f, 0x23, 0xc5, 0x25, 0x49, 0x2d, 0xd9, 0xfd, 0xe5, 0x6c, 0xb1, 0x9c, 0xc5, 0x92,
	0xce, 0x9e, 0x50, 0xac, 0xf4, 0x4c, 0xa3, 0x3f, 0x31, 0x8f, 0x63, 0x83, 0xab, 0x63, 0x85, 0x74,
	0x2e, 0xce, 0x70, 0xfa, 0x78, 0x6a, 0x33, 0x7c, 0xbe, 0x09, 0x57, 0x29, 0x16, 0x09, 0xc7, 0xb0,
	0xff, 0x06, 0x75, 0x62, 0xeb, 0x27, 0xc5, 0x5c, 0xa6, 0xf8, 0xf3, 0x17, 0x2a, 0x1d, 0x9e, 0xc2,
	0x68, 0x1b, 0x50, 0xa5, 0x2c, 0x14, 0x32, 0x0e, 0x9d, 0x73, 0x24, 0x25, 0x64, 0xc1, 0xbd, 0xc0,
	0x8b, 0x86, 0x69, 0xfd, 0x64, 0x87, 0xd0, 0x37, 0xb5, 0x33, 0x7d, 0x51, 0xa2, 0xe2, 0xad, 0xa0,
	0x1d, 0xf5, 0x52, 0x30, 0xa1, 0x8f, 0x55, 0x24, 0xfc, 0x03, 0x03, 0x53, 0xd1, 0x35, 0x61, 0x47,
	0xb0, 0x6b, 0x50, 0xc5, 0xbd, 0xa0, 0x1d, 0xf5, 0x4f, 0x58, 0xbc, 0x3e, 0xda, 0x69, 0x05, 0xa5,
	0x8e, 0xc1, 0x26, 0x00, 0x05, 0xea, 0x95, 0xa4, 0xef, 0x99, 0xc8, 0x79, 0x2b, 0xf0, 0xa2, 0x5e,
	0xda, 0x73, 0x91, 0x24, 0x67, 0x0f, 0x61, 0x48, 0x58, 0x4a, 0xd2, 0x48, 0xd9, 0xd7, 0x95, 0xc8,
	0x79, 0xdb, 0x30, 0x06, 0x75, 0xf0, 0xed, 0x4a, 0xe4, 0xe1, 0x3b, 0x18, 0xba, 0xfe, 0x6e, 0x96,
	0xa7, 0x30, 0xb0, 0x8a, 0x91, 0x48, 0x52, 0x2d, 0x63, 0xbc, 0x21, 0x23, 0x79, 0xf9, 0xac, 0xc8,
	0x5f, 0x55, 0x78, 0x6a, 0xc7, 0x33, 0xff, 0x55, 0xf8, 0xd7, 0x83, 0xfd, 0x0f, 0x84, 0xe5, 0x8c,
	0x30, 0x45, 0xb1, 0x3e, 0xd6, 0x04, 0xc0, 0x5d, 0xac, 0x92, 0xea, 0x59, 0xa9, 0x2e, 0x92, 0xe4,
	0xec, 0x01, 0x0c, 0xe6, 0x24, 0x97, 0x59, 0xbd, 0xc5, 0x96, 0xd9, 0x62, 0xbf, 0x8a, 0x7d, 0x72,
	0x9b, 0x9c, 0x00, 0x68, 0x79, 0x49, 0x68, 0x1b, 0x42, 0x4f, 0xcb, 0x1a, 0x3e, 0x80, 0xae, 0x50,
	0xd9, 0x5c, 0x90, 0xd2, 0xfc, 0x4e, 0xe0, 0x45, 0xdd, 0xb4, 0x23, 0xd4, 0xeb, 0xea, 0x19, 0x72,
	0x18, 0x6d, 0x8b, 0xb2, 0xb3, 0x86, 0x17, 0x30, 0x7a, 0x21, 0x97, 0xe5, 0x0f, 0xd4, 0xb7, 0xad,
	0x37, 0x3c, 0x80, 0xf1, 0x95, 0xd6, 0x56, 0xd5, 0xc9, 0xbf, 0x36, 0x74, 0x9c, 0xcb, 0xd8, 0x02,
	0xf6, 0x36, 0x3d, 0xc7, 0xa2, 0xf8, 0x8a, 0xd7, 0xe3, 0x46, 0xbf, 0xfa, 0x8f, 0x6e, 0xc0, 0x74,
	0x8b, 0xd8, 0x61, 0xef, 0xe1, 0xae, 0x01, 0xd8, 0x61, 0x43, 0xd6, 0xba, 0x43, 0xfd, 0xe0, 0x7a,
	0xc2, 0x65, 0xb5, 0x05, 0xec, 0x6d, 0xae, 0xbc, 0x51, 0x76, 0xa3, 0x55, 0x1a, 0x65, 0x5f, 0x73,
	0xbf, 0x1d, 0xf6, 0x0d, 0xee, 0x6d, 0xad, 0x91, 0x35, 0xe5, 0x37, 0x5f, 0xd9, 0x3f, 0xba, 0x09,
	0xb5, 0xee, 0xf5, 0xbc, 0xfb, 0x79, 0xd7, 0x7e, 0x2e, 0xbe, 0xd8, 0xdf, 0xe9, 0xff, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x55, 0x6f, 0x42, 0x41, 0x86, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// IndexerClient is the client API for Indexer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IndexerClient interface {
	// GetIndexerInfo gets info about this indexer.
	GetIndexerInfo(ctx context.Context, in *GetIndexerInfoRequest, opts ...grpc.CallOption) (*GetIndexerInfoResponse, error)
	// Index a set of states by forwarding to locally-registered indexers.
	Index(ctx context.Context, in *IndexRequest, opts ...grpc.CallOption) (*IndexResponse, error)
	// PrepareReindex of a particular indexer.
	PrepareReindex(ctx context.Context, in *PrepareReindexRequest, opts ...grpc.CallOption) (*PrepareReindexResponse, error)
	// CompleteReindex of a particular indexer.
	CompleteReindex(ctx context.Context, in *CompleteReindexRequest, opts ...grpc.CallOption) (*CompleteReindexResponse, error)
}

type indexerClient struct {
	cc grpc.ClientConnInterface
}

func NewIndexerClient(cc grpc.ClientConnInterface) IndexerClient {
	return &indexerClient{cc}
}

func (c *indexerClient) GetIndexerInfo(ctx context.Context, in *GetIndexerInfoRequest, opts ...grpc.CallOption) (*GetIndexerInfoResponse, error) {
	out := new(GetIndexerInfoResponse)
	err := c.cc.Invoke(ctx, "/magma.orc8r.state.Indexer/GetIndexerInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexerClient) Index(ctx context.Context, in *IndexRequest, opts ...grpc.CallOption) (*IndexResponse, error) {
	out := new(IndexResponse)
	err := c.cc.Invoke(ctx, "/magma.orc8r.state.Indexer/Index", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexerClient) PrepareReindex(ctx context.Context, in *PrepareReindexRequest, opts ...grpc.CallOption) (*PrepareReindexResponse, error) {
	out := new(PrepareReindexResponse)
	err := c.cc.Invoke(ctx, "/magma.orc8r.state.Indexer/PrepareReindex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexerClient) CompleteReindex(ctx context.Context, in *CompleteReindexRequest, opts ...grpc.CallOption) (*CompleteReindexResponse, error) {
	out := new(CompleteReindexResponse)
	err := c.cc.Invoke(ctx, "/magma.orc8r.state.Indexer/CompleteReindex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IndexerServer is the server API for Indexer service.
type IndexerServer interface {
	// GetIndexerInfo gets info about this indexer.
	GetIndexerInfo(context.Context, *GetIndexerInfoRequest) (*GetIndexerInfoResponse, error)
	// Index a set of states by forwarding to locally-registered indexers.
	Index(context.Context, *IndexRequest) (*IndexResponse, error)
	// PrepareReindex of a particular indexer.
	PrepareReindex(context.Context, *PrepareReindexRequest) (*PrepareReindexResponse, error)
	// CompleteReindex of a particular indexer.
	CompleteReindex(context.Context, *CompleteReindexRequest) (*CompleteReindexResponse, error)
}

// UnimplementedIndexerServer can be embedded to have forward compatible implementations.
type UnimplementedIndexerServer struct {
}

func (*UnimplementedIndexerServer) GetIndexerInfo(ctx context.Context, req *GetIndexerInfoRequest) (*GetIndexerInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIndexerInfo not implemented")
}
func (*UnimplementedIndexerServer) Index(ctx context.Context, req *IndexRequest) (*IndexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Index not implemented")
}
func (*UnimplementedIndexerServer) PrepareReindex(ctx context.Context, req *PrepareReindexRequest) (*PrepareReindexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrepareReindex not implemented")
}
func (*UnimplementedIndexerServer) CompleteReindex(ctx context.Context, req *CompleteReindexRequest) (*CompleteReindexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteReindex not implemented")
}

func RegisterIndexerServer(s *grpc.Server, srv IndexerServer) {
	s.RegisterService(&_Indexer_serviceDesc, srv)
}

func _Indexer_GetIndexerInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIndexerInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexerServer).GetIndexerInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.state.Indexer/GetIndexerInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexerServer).GetIndexerInfo(ctx, req.(*GetIndexerInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Indexer_Index_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexerServer).Index(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.state.Indexer/Index",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexerServer).Index(ctx, req.(*IndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Indexer_PrepareReindex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrepareReindexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexerServer).PrepareReindex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.state.Indexer/PrepareReindex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexerServer).PrepareReindex(ctx, req.(*PrepareReindexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Indexer_CompleteReindex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteReindexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexerServer).CompleteReindex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.state.Indexer/CompleteReindex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexerServer).CompleteReindex(ctx, req.(*CompleteReindexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Indexer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.orc8r.state.Indexer",
	HandlerType: (*IndexerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetIndexerInfo",
			Handler:    _Indexer_GetIndexerInfo_Handler,
		},
		{
			MethodName: "Index",
			Handler:    _Indexer_Index_Handler,
		},
		{
			MethodName: "PrepareReindex",
			Handler:    _Indexer_PrepareReindex_Handler,
		},
		{
			MethodName: "CompleteReindex",
			Handler:    _Indexer_CompleteReindex_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "indexer.proto",
}
