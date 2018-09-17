// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_d9acbd94fa0d6421, []int{0}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type ClientStatus struct {
	ReplicaSetUUID       string   `protobuf:"bytes,1,opt,name=ReplicaSetUUID,proto3" json:"ReplicaSetUUID,omitempty"`
	ReplicaSetName       string   `protobuf:"bytes,2,opt,name=ReplicaSetName,proto3" json:"ReplicaSetName,omitempty"`
	ReplicaSetVersion    uint64   `protobuf:"varint,3,opt,name=ReplicaSetVersion,proto3" json:"ReplicaSetVersion,omitempty"`
	RunningDBBackup      bool     `protobuf:"varint,4,opt,name=RunningDBBackup,proto3" json:"RunningDBBackup,omitempty"`
	RunningOplogBackup   bool     `protobuf:"varint,5,opt,name=RunningOplogBackup,proto3" json:"RunningOplogBackup,omitempty"`
	BackupType           string   `protobuf:"bytes,6,opt,name=BackupType,proto3" json:"BackupType,omitempty"`
	Compression          string   `protobuf:"bytes,7,opt,name=Compression,proto3" json:"Compression,omitempty"`
	Encrypted            string   `protobuf:"bytes,8,opt,name=Encrypted,proto3" json:"Encrypted,omitempty"`
	Destination          string   `protobuf:"bytes,9,opt,name=Destination,proto3" json:"Destination,omitempty"`
	Dirname              string   `protobuf:"bytes,10,opt,name=Dirname,proto3" json:"Dirname,omitempty"`
	Filename             string   `protobuf:"bytes,11,opt,name=Filename,proto3" json:"Filename,omitempty"`
	Started              int64    `protobuf:"varint,12,opt,name=Started,proto3" json:"Started,omitempty"`
	Finished             int64    `protobuf:"varint,13,opt,name=Finished,proto3" json:"Finished,omitempty"`
	StartOplogTs         int64    `protobuf:"varint,14,opt,name=StartOplogTs,proto3" json:"StartOplogTs,omitempty"`
	LastOplogTs          int64    `protobuf:"varint,15,opt,name=LastOplogTs,proto3" json:"LastOplogTs,omitempty"`
	LastError            string   `protobuf:"bytes,16,opt,name=LastError,proto3" json:"LastError,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientStatus) Reset()         { *m = ClientStatus{} }
func (m *ClientStatus) String() string { return proto.CompactTextString(m) }
func (*ClientStatus) ProtoMessage()    {}
func (*ClientStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_d9acbd94fa0d6421, []int{1}
}
func (m *ClientStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientStatus.Unmarshal(m, b)
}
func (m *ClientStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientStatus.Marshal(b, m, deterministic)
}
func (dst *ClientStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientStatus.Merge(dst, src)
}
func (m *ClientStatus) XXX_Size() int {
	return xxx_messageInfo_ClientStatus.Size(m)
}
func (m *ClientStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ClientStatus proto.InternalMessageInfo

func (m *ClientStatus) GetReplicaSetUUID() string {
	if m != nil {
		return m.ReplicaSetUUID
	}
	return ""
}

func (m *ClientStatus) GetReplicaSetName() string {
	if m != nil {
		return m.ReplicaSetName
	}
	return ""
}

func (m *ClientStatus) GetReplicaSetVersion() uint64 {
	if m != nil {
		return m.ReplicaSetVersion
	}
	return 0
}

func (m *ClientStatus) GetRunningDBBackup() bool {
	if m != nil {
		return m.RunningDBBackup
	}
	return false
}

func (m *ClientStatus) GetRunningOplogBackup() bool {
	if m != nil {
		return m.RunningOplogBackup
	}
	return false
}

func (m *ClientStatus) GetBackupType() string {
	if m != nil {
		return m.BackupType
	}
	return ""
}

func (m *ClientStatus) GetCompression() string {
	if m != nil {
		return m.Compression
	}
	return ""
}

func (m *ClientStatus) GetEncrypted() string {
	if m != nil {
		return m.Encrypted
	}
	return ""
}

func (m *ClientStatus) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

func (m *ClientStatus) GetDirname() string {
	if m != nil {
		return m.Dirname
	}
	return ""
}

func (m *ClientStatus) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *ClientStatus) GetStarted() int64 {
	if m != nil {
		return m.Started
	}
	return 0
}

func (m *ClientStatus) GetFinished() int64 {
	if m != nil {
		return m.Finished
	}
	return 0
}

func (m *ClientStatus) GetStartOplogTs() int64 {
	if m != nil {
		return m.StartOplogTs
	}
	return 0
}

func (m *ClientStatus) GetLastOplogTs() int64 {
	if m != nil {
		return m.LastOplogTs
	}
	return 0
}

func (m *ClientStatus) GetLastError() string {
	if m != nil {
		return m.LastError
	}
	return ""
}

type Client struct {
	Version              int32         `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	ID                   string        `protobuf:"bytes,2,opt,name=ID,proto3" json:"ID,omitempty"`
	NodeType             string        `protobuf:"bytes,3,opt,name=NodeType,proto3" json:"NodeType,omitempty"`
	NodeName             string        `protobuf:"bytes,4,opt,name=NodeName,proto3" json:"NodeName,omitempty"`
	ClusterID            string        `protobuf:"bytes,5,opt,name=ClusterID,proto3" json:"ClusterID,omitempty"`
	ReplicasetName       string        `protobuf:"bytes,6,opt,name=ReplicasetName,proto3" json:"ReplicasetName,omitempty"`
	ReplicasetID         string        `protobuf:"bytes,7,opt,name=ReplicasetID,proto3" json:"ReplicasetID,omitempty"`
	LastCommandSent      string        `protobuf:"bytes,8,opt,name=LastCommandSent,proto3" json:"LastCommandSent,omitempty"`
	LastSeen             int64         `protobuf:"varint,9,opt,name=LastSeen,proto3" json:"LastSeen,omitempty"`
	Status               *ClientStatus `protobuf:"bytes,10,opt,name=Status,proto3" json:"Status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Client) Reset()         { *m = Client{} }
func (m *Client) String() string { return proto.CompactTextString(m) }
func (*Client) ProtoMessage()    {}
func (*Client) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_d9acbd94fa0d6421, []int{2}
}
func (m *Client) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Client.Unmarshal(m, b)
}
func (m *Client) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Client.Marshal(b, m, deterministic)
}
func (dst *Client) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Client.Merge(dst, src)
}
func (m *Client) XXX_Size() int {
	return xxx_messageInfo_Client.Size(m)
}
func (m *Client) XXX_DiscardUnknown() {
	xxx_messageInfo_Client.DiscardUnknown(m)
}

var xxx_messageInfo_Client proto.InternalMessageInfo

func (m *Client) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Client) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Client) GetNodeType() string {
	if m != nil {
		return m.NodeType
	}
	return ""
}

func (m *Client) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *Client) GetClusterID() string {
	if m != nil {
		return m.ClusterID
	}
	return ""
}

func (m *Client) GetReplicasetName() string {
	if m != nil {
		return m.ReplicasetName
	}
	return ""
}

func (m *Client) GetReplicasetID() string {
	if m != nil {
		return m.ReplicasetID
	}
	return ""
}

func (m *Client) GetLastCommandSent() string {
	if m != nil {
		return m.LastCommandSent
	}
	return ""
}

func (m *Client) GetLastSeen() int64 {
	if m != nil {
		return m.LastSeen
	}
	return 0
}

func (m *Client) GetStatus() *ClientStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type StartBackupParams struct {
	BackupType           string   `protobuf:"bytes,1,opt,name=BackupType,proto3" json:"BackupType,omitempty"`
	Compression          string   `protobuf:"bytes,2,opt,name=Compression,proto3" json:"Compression,omitempty"`
	Encrypted            string   `protobuf:"bytes,3,opt,name=Encrypted,proto3" json:"Encrypted,omitempty"`
	Destination          string   `protobuf:"bytes,4,opt,name=Destination,proto3" json:"Destination,omitempty"`
	Dirname              string   `protobuf:"bytes,5,opt,name=Dirname,proto3" json:"Dirname,omitempty"`
	Filename             string   `protobuf:"bytes,6,opt,name=Filename,proto3" json:"Filename,omitempty"`
	StartOplogTs         int64    `protobuf:"varint,7,opt,name=StartOplogTs,proto3" json:"StartOplogTs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartBackupParams) Reset()         { *m = StartBackupParams{} }
func (m *StartBackupParams) String() string { return proto.CompactTextString(m) }
func (*StartBackupParams) ProtoMessage()    {}
func (*StartBackupParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_d9acbd94fa0d6421, []int{3}
}
func (m *StartBackupParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartBackupParams.Unmarshal(m, b)
}
func (m *StartBackupParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartBackupParams.Marshal(b, m, deterministic)
}
func (dst *StartBackupParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartBackupParams.Merge(dst, src)
}
func (m *StartBackupParams) XXX_Size() int {
	return xxx_messageInfo_StartBackupParams.Size(m)
}
func (m *StartBackupParams) XXX_DiscardUnknown() {
	xxx_messageInfo_StartBackupParams.DiscardUnknown(m)
}

var xxx_messageInfo_StartBackupParams proto.InternalMessageInfo

func (m *StartBackupParams) GetBackupType() string {
	if m != nil {
		return m.BackupType
	}
	return ""
}

func (m *StartBackupParams) GetCompression() string {
	if m != nil {
		return m.Compression
	}
	return ""
}

func (m *StartBackupParams) GetEncrypted() string {
	if m != nil {
		return m.Encrypted
	}
	return ""
}

func (m *StartBackupParams) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

func (m *StartBackupParams) GetDirname() string {
	if m != nil {
		return m.Dirname
	}
	return ""
}

func (m *StartBackupParams) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *StartBackupParams) GetStartOplogTs() int64 {
	if m != nil {
		return m.StartOplogTs
	}
	return 0
}

type BackupReplica struct {
	AgentID              string   `protobuf:"bytes,1,opt,name=AgentID,proto3" json:"AgentID,omitempty"`
	NodeType             string   `protobuf:"bytes,2,opt,name=NodeType,proto3" json:"NodeType,omitempty"`
	NodeName             string   `protobuf:"bytes,3,opt,name=NodeName,proto3" json:"NodeName,omitempty"`
	ClusterID            string   `protobuf:"bytes,4,opt,name=ClusterID,proto3" json:"ClusterID,omitempty"`
	ReplicasetName       string   `protobuf:"bytes,5,opt,name=ReplicasetName,proto3" json:"ReplicasetName,omitempty"`
	ReplicasetID         string   `protobuf:"bytes,6,opt,name=ReplicasetID,proto3" json:"ReplicasetID,omitempty"`
	StartedAt            int64    `protobuf:"varint,7,opt,name=StartedAt,proto3" json:"StartedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BackupReplica) Reset()         { *m = BackupReplica{} }
func (m *BackupReplica) String() string { return proto.CompactTextString(m) }
func (*BackupReplica) ProtoMessage()    {}
func (*BackupReplica) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_d9acbd94fa0d6421, []int{4}
}
func (m *BackupReplica) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BackupReplica.Unmarshal(m, b)
}
func (m *BackupReplica) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BackupReplica.Marshal(b, m, deterministic)
}
func (dst *BackupReplica) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BackupReplica.Merge(dst, src)
}
func (m *BackupReplica) XXX_Size() int {
	return xxx_messageInfo_BackupReplica.Size(m)
}
func (m *BackupReplica) XXX_DiscardUnknown() {
	xxx_messageInfo_BackupReplica.DiscardUnknown(m)
}

var xxx_messageInfo_BackupReplica proto.InternalMessageInfo

func (m *BackupReplica) GetAgentID() string {
	if m != nil {
		return m.AgentID
	}
	return ""
}

func (m *BackupReplica) GetNodeType() string {
	if m != nil {
		return m.NodeType
	}
	return ""
}

func (m *BackupReplica) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *BackupReplica) GetClusterID() string {
	if m != nil {
		return m.ClusterID
	}
	return ""
}

func (m *BackupReplica) GetReplicasetName() string {
	if m != nil {
		return m.ReplicasetName
	}
	return ""
}

func (m *BackupReplica) GetReplicasetID() string {
	if m != nil {
		return m.ReplicasetID
	}
	return ""
}

func (m *BackupReplica) GetStartedAt() int64 {
	if m != nil {
		return m.StartedAt
	}
	return 0
}

type StartBackupResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartBackupResponse) Reset()         { *m = StartBackupResponse{} }
func (m *StartBackupResponse) String() string { return proto.CompactTextString(m) }
func (*StartBackupResponse) ProtoMessage()    {}
func (*StartBackupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_d9acbd94fa0d6421, []int{5}
}
func (m *StartBackupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartBackupResponse.Unmarshal(m, b)
}
func (m *StartBackupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartBackupResponse.Marshal(b, m, deterministic)
}
func (dst *StartBackupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartBackupResponse.Merge(dst, src)
}
func (m *StartBackupResponse) XXX_Size() int {
	return xxx_messageInfo_StartBackupResponse.Size(m)
}
func (m *StartBackupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StartBackupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StartBackupResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Empty)(nil), "api.Empty")
	proto.RegisterType((*ClientStatus)(nil), "api.ClientStatus")
	proto.RegisterType((*Client)(nil), "api.Client")
	proto.RegisterType((*StartBackupParams)(nil), "api.StartBackupParams")
	proto.RegisterType((*BackupReplica)(nil), "api.BackupReplica")
	proto.RegisterType((*StartBackupResponse)(nil), "api.StartBackupResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ApiClient is the client API for Api service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ApiClient interface {
	GetClients(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Api_GetClientsClient, error)
	StartBackup(ctx context.Context, in *StartBackupParams, opts ...grpc.CallOption) (*StartBackupResponse, error)
}

type apiClient struct {
	cc *grpc.ClientConn
}

func NewApiClient(cc *grpc.ClientConn) ApiClient {
	return &apiClient{cc}
}

func (c *apiClient) GetClients(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Api_GetClientsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Api_serviceDesc.Streams[0], "/api.Api/GetClients", opts...)
	if err != nil {
		return nil, err
	}
	x := &apiGetClientsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Api_GetClientsClient interface {
	Recv() (*Client, error)
	grpc.ClientStream
}

type apiGetClientsClient struct {
	grpc.ClientStream
}

func (x *apiGetClientsClient) Recv() (*Client, error) {
	m := new(Client)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *apiClient) StartBackup(ctx context.Context, in *StartBackupParams, opts ...grpc.CallOption) (*StartBackupResponse, error) {
	out := new(StartBackupResponse)
	err := c.cc.Invoke(ctx, "/api.Api/StartBackup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiServer is the server API for Api service.
type ApiServer interface {
	GetClients(*Empty, Api_GetClientsServer) error
	StartBackup(context.Context, *StartBackupParams) (*StartBackupResponse, error)
}

func RegisterApiServer(s *grpc.Server, srv ApiServer) {
	s.RegisterService(&_Api_serviceDesc, srv)
}

func _Api_GetClients_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ApiServer).GetClients(m, &apiGetClientsServer{stream})
}

type Api_GetClientsServer interface {
	Send(*Client) error
	grpc.ServerStream
}

type apiGetClientsServer struct {
	grpc.ServerStream
}

func (x *apiGetClientsServer) Send(m *Client) error {
	return x.ServerStream.SendMsg(m)
}

func _Api_StartBackup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartBackupParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).StartBackup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Api/StartBackup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).StartBackup(ctx, req.(*StartBackupParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _Api_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Api",
	HandlerType: (*ApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartBackup",
			Handler:    _Api_StartBackup_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetClients",
			Handler:       _Api_GetClients_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_api_d9acbd94fa0d6421) }

var fileDescriptor_api_d9acbd94fa0d6421 = []byte{
	// 641 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x95, 0xdf, 0x6e, 0xd3, 0x3e,
	0x14, 0xc7, 0x97, 0xa6, 0x4d, 0x97, 0xd3, 0xfd, 0xf9, 0xcd, 0x3f, 0x81, 0xc2, 0x34, 0xa1, 0x2a,
	0x17, 0xa8, 0x93, 0xa6, 0x0a, 0x8d, 0x27, 0x68, 0x9b, 0x81, 0x2a, 0xc1, 0x98, 0xd2, 0x8d, 0x7b,
	0xd3, 0x1e, 0x15, 0x8b, 0xc6, 0xb1, 0x6c, 0x17, 0xb1, 0x37, 0xe3, 0x05, 0x78, 0x18, 0xae, 0x78,
	0x05, 0x64, 0xc7, 0x69, 0xd2, 0x6c, 0xac, 0xdc, 0xe5, 0x7c, 0xbf, 0xc7, 0xf1, 0x39, 0x3e, 0x1f,
	0x27, 0x10, 0x52, 0xc1, 0x86, 0x42, 0xe6, 0x3a, 0x27, 0x3e, 0x15, 0x2c, 0xee, 0x42, 0xe7, 0x2a,
	0x13, 0xfa, 0x3e, 0xfe, 0xd1, 0x86, 0x83, 0xc9, 0x8a, 0x21, 0xd7, 0x33, 0x4d, 0xf5, 0x5a, 0x91,
	0x57, 0x70, 0x94, 0xa2, 0x58, 0xb1, 0x39, 0x9d, 0xa1, 0xbe, 0xbb, 0x9b, 0x26, 0x91, 0xd7, 0xf7,
	0x06, 0x61, 0xda, 0x50, 0xb7, 0xf3, 0xae, 0x69, 0x86, 0x51, 0xab, 0x99, 0x67, 0x54, 0x72, 0x01,
	0x27, 0x95, 0xf2, 0x09, 0xa5, 0x62, 0x39, 0x8f, 0xfc, 0xbe, 0x37, 0x68, 0xa7, 0x0f, 0x0d, 0x32,
	0x80, 0xe3, 0x74, 0xcd, 0x39, 0xe3, 0xcb, 0x64, 0x3c, 0xa6, 0xf3, 0xaf, 0x6b, 0x11, 0xb5, 0xfb,
	0xde, 0x60, 0x3f, 0x6d, 0xca, 0x64, 0x08, 0xc4, 0x49, 0x1f, 0xc5, 0x2a, 0x5f, 0xba, 0xe4, 0x8e,
	0x4d, 0x7e, 0xc4, 0x21, 0x2f, 0x01, 0x8a, 0xa7, 0xdb, 0x7b, 0x81, 0x51, 0x60, 0x6b, 0xad, 0x29,
	0xa4, 0x0f, 0xbd, 0x49, 0x9e, 0x09, 0x89, 0xca, 0x56, 0xd8, 0xb5, 0x09, 0x75, 0x89, 0x9c, 0x41,
	0x78, 0xc5, 0xe7, 0xf2, 0x5e, 0x68, 0x5c, 0x44, 0xfb, 0xd6, 0xaf, 0x04, 0xb3, 0x3e, 0x41, 0xa5,
	0x19, 0xa7, 0xda, 0xac, 0x0f, 0x8b, 0xf5, 0x35, 0x89, 0x44, 0xd0, 0x4d, 0x98, 0xe4, 0xe6, 0xa8,
	0xc0, 0xba, 0x65, 0x48, 0x4e, 0x61, 0xff, 0x2d, 0x5b, 0xa1, 0xb5, 0x7a, 0xd6, 0xda, 0xc4, 0x66,
	0xd5, 0x4c, 0x53, 0x69, 0xf6, 0x3c, 0xe8, 0x7b, 0x03, 0x3f, 0x2d, 0xc3, 0x62, 0x15, 0x67, 0xea,
	0x0b, 0x2e, 0xa2, 0x43, 0x6b, 0x6d, 0x62, 0x12, 0xc3, 0x81, 0x4d, 0xb3, 0x27, 0x70, 0xab, 0xa2,
	0x23, 0xeb, 0x6f, 0x69, 0xa6, 0xe2, 0xf7, 0x54, 0x6d, 0x52, 0x8e, 0x6d, 0x4a, 0x5d, 0x32, 0x1d,
	0x9b, 0xf0, 0x4a, 0xca, 0x5c, 0x46, 0xff, 0x15, 0x1d, 0x6f, 0x84, 0xf8, 0x67, 0x0b, 0x82, 0x02,
	0x1d, 0x53, 0xe4, 0x37, 0x37, 0x5a, 0x43, 0x4b, 0x27, 0x2d, 0x43, 0x72, 0x04, 0xad, 0x69, 0xe2,
	0xd0, 0x68, 0x4d, 0x13, 0x53, 0xf4, 0x75, 0xbe, 0x40, 0x3b, 0x04, 0xbf, 0x68, 0xb5, 0x8c, 0x4b,
	0xcf, 0xc2, 0xd4, 0xae, 0x3c, 0x8b, 0xd1, 0x19, 0x84, 0x93, 0xd5, 0x5a, 0x69, 0x94, 0xd3, 0xc4,
	0x4e, 0x39, 0x4c, 0x2b, 0xa1, 0x06, 0xa3, 0x72, 0x30, 0x06, 0x5b, 0x30, 0x3a, 0xd5, 0x1c, 0x4b,
	0xa5, 0x4c, 0x13, 0x37, 0xe5, 0x2d, 0xcd, 0x20, 0x68, 0x7a, 0x9c, 0xe4, 0x59, 0x46, 0xf9, 0x62,
	0x86, 0x5c, 0xbb, 0x61, 0x37, 0x65, 0x53, 0xaf, 0x91, 0x66, 0x88, 0xc5, 0xbc, 0xfd, 0x74, 0x13,
	0x93, 0x73, 0x08, 0x8a, 0x0b, 0x65, 0x67, 0xdd, 0xbb, 0x3c, 0x19, 0x9a, 0x1b, 0x58, 0xbf, 0x69,
	0xa9, 0x4b, 0x88, 0x7f, 0x7b, 0x70, 0x62, 0x07, 0x53, 0xd0, 0x78, 0x43, 0x25, 0xcd, 0x54, 0x83,
	0x57, 0x6f, 0x17, 0xaf, 0xad, 0x1d, 0xbc, 0xfa, 0x3b, 0x78, 0x6d, 0x3f, 0xc9, 0x6b, 0xe7, 0xef,
	0xbc, 0x06, 0x0d, 0x5e, 0x9b, 0xe4, 0x75, 0x1f, 0x92, 0x17, 0xff, 0xf2, 0xe0, 0xb0, 0x68, 0xc5,
	0x9d, 0xbc, 0xd9, 0x6b, 0xb4, 0x44, 0xae, 0x37, 0x9f, 0x9b, 0x32, 0xdc, 0x02, 0xa6, 0xf5, 0x04,
	0x30, 0xfe, 0x53, 0xc0, 0xb4, 0x77, 0x03, 0xd3, 0xf9, 0x27, 0x60, 0x82, 0x47, 0x80, 0x39, 0x83,
	0xd0, 0x5d, 0xc9, 0x91, 0x76, 0xed, 0x56, 0x42, 0xfc, 0x0c, 0xfe, 0xaf, 0x0d, 0x37, 0x45, 0x25,
	0x72, 0xae, 0xf0, 0x52, 0x81, 0x3f, 0x12, 0x8c, 0x9c, 0x03, 0xbc, 0x43, 0x5d, 0x60, 0xa1, 0x08,
	0x58, 0x48, 0xec, 0x87, 0xf9, 0xb4, 0x57, 0x03, 0x26, 0xde, 0x7b, 0xed, 0x91, 0x11, 0xf4, 0x6a,
	0x2f, 0x22, 0xcf, 0xad, 0xff, 0x80, 0x9b, 0xd3, 0xa8, 0xa9, 0x97, 0x5b, 0xc6, 0x7b, 0xe3, 0x0b,
	0x78, 0xc1, 0xf2, 0xe1, 0x52, 0x8a, 0xf9, 0x10, 0xbf, 0xd3, 0x4c, 0xac, 0x50, 0x0d, 0x33, 0x54,
	0x8a, 0x2e, 0x51, 0x8d, 0x0f, 0x3f, 0xb8, 0xa7, 0x1b, 0xf3, 0x9b, 0xb8, 0xf1, 0x3e, 0x07, 0xf6,
	0x7f, 0xf1, 0xe6, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x30, 0xa5, 0x86, 0x3c, 0x06, 0x00,
	0x00,
}
