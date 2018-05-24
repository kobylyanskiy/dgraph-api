// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dgraph.proto

package dgraph

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

type Agent struct {
	Codename             string   `protobuf:"bytes,1,opt,name=codename" json:"codename,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Agent) Reset()         { *m = Agent{} }
func (m *Agent) String() string { return proto.CompactTextString(m) }
func (*Agent) ProtoMessage()    {}
func (*Agent) Descriptor() ([]byte, []int) {
	return fileDescriptor_dgraph_3f3d270d4c407406, []int{0}
}
func (m *Agent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Agent.Unmarshal(m, b)
}
func (m *Agent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Agent.Marshal(b, m, deterministic)
}
func (dst *Agent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Agent.Merge(dst, src)
}
func (m *Agent) XXX_Size() int {
	return xxx_messageInfo_Agent.Size(m)
}
func (m *Agent) XXX_DiscardUnknown() {
	xxx_messageInfo_Agent.DiscardUnknown(m)
}

var xxx_messageInfo_Agent proto.InternalMessageInfo

func (m *Agent) GetCodename() string {
	if m != nil {
		return m.Codename
	}
	return ""
}

type Result struct {
	Result               string   `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_dgraph_3f3d270d4c407406, []int{1}
}
func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (dst *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(dst, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*Agent)(nil), "dgraph.Agent")
	proto.RegisterType((*Result)(nil), "dgraph.Result")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DgraphServiceClient is the client API for DgraphService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DgraphServiceClient interface {
	AddAgent(ctx context.Context, in *Agent, opts ...grpc.CallOption) (*Result, error)
}

type dgraphServiceClient struct {
	cc *grpc.ClientConn
}

func NewDgraphServiceClient(cc *grpc.ClientConn) DgraphServiceClient {
	return &dgraphServiceClient{cc}
}

func (c *dgraphServiceClient) AddAgent(ctx context.Context, in *Agent, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/dgraph.DgraphService/AddAgent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DgraphServiceServer is the server API for DgraphService service.
type DgraphServiceServer interface {
	AddAgent(context.Context, *Agent) (*Result, error)
}

func RegisterDgraphServiceServer(s *grpc.Server, srv DgraphServiceServer) {
	s.RegisterService(&_DgraphService_serviceDesc, srv)
}

func _DgraphService_AddAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Agent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DgraphServiceServer).AddAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dgraph.DgraphService/AddAgent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DgraphServiceServer).AddAgent(ctx, req.(*Agent))
	}
	return interceptor(ctx, in, info, handler)
}

var _DgraphService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dgraph.DgraphService",
	HandlerType: (*DgraphServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddAgent",
			Handler:    _DgraphService_AddAgent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dgraph.proto",
}

func init() { proto.RegisterFile("dgraph.proto", fileDescriptor_dgraph_3f3d270d4c407406) }

var fileDescriptor_dgraph_3f3d270d4c407406 = []byte{
	// 134 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x49, 0x2f, 0x4a,
	0x2c, 0xc8, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0x94, 0xb9, 0x58,
	0x1d, 0xd3, 0x53, 0xf3, 0x4a, 0x84, 0xa4, 0xb8, 0x38, 0x92, 0xf3, 0x53, 0x52, 0xf3, 0x12, 0x73,
	0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xe0, 0x7c, 0x25, 0x05, 0x2e, 0xb6, 0xa0, 0xd4,
	0xe2, 0xd2, 0x9c, 0x12, 0x21, 0x31, 0x2e, 0xb6, 0x22, 0x30, 0x0b, 0xaa, 0x06, 0xca, 0x33, 0xb2,
	0xe1, 0xe2, 0x75, 0x01, 0x1b, 0x18, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a, 0xa4, 0xcd, 0xc5,
	0xe1, 0x98, 0x92, 0x02, 0x31, 0x9a, 0x57, 0x0f, 0x6a, 0x35, 0x98, 0x2b, 0xc5, 0x07, 0xe3, 0x42,
	0xcc, 0x54, 0x62, 0x48, 0x62, 0x03, 0xbb, 0xc9, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xad, 0x71,
	0x76, 0xf5, 0xa3, 0x00, 0x00, 0x00,
}