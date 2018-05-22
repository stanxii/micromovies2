// Code generated by protoc-gen-go. DO NOT EDIT.
// source: users.proto

package pb

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

type NewUserRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	LastName             string   `protobuf:"bytes,2,opt,name=LastName" json:"LastName,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=Email" json:"Email,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=Password" json:"Password,omitempty"`
	Role                 string   `protobuf:"bytes,5,opt,name=Role" json:"Role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewUserRequest) Reset()         { *m = NewUserRequest{} }
func (m *NewUserRequest) String() string { return proto.CompactTextString(m) }
func (*NewUserRequest) ProtoMessage()    {}
func (*NewUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_4e3f99c9794f712d, []int{0}
}
func (m *NewUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewUserRequest.Unmarshal(m, b)
}
func (m *NewUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewUserRequest.Marshal(b, m, deterministic)
}
func (dst *NewUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewUserRequest.Merge(dst, src)
}
func (m *NewUserRequest) XXX_Size() int {
	return xxx_messageInfo_NewUserRequest.Size(m)
}
func (m *NewUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewUserRequest proto.InternalMessageInfo

func (m *NewUserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NewUserRequest) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *NewUserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *NewUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *NewUserRequest) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

type NewUserResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id" json:"Id,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=Err" json:"Err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewUserResponse) Reset()         { *m = NewUserResponse{} }
func (m *NewUserResponse) String() string { return proto.CompactTextString(m) }
func (*NewUserResponse) ProtoMessage()    {}
func (*NewUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_4e3f99c9794f712d, []int{1}
}
func (m *NewUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewUserResponse.Unmarshal(m, b)
}
func (m *NewUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewUserResponse.Marshal(b, m, deterministic)
}
func (dst *NewUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewUserResponse.Merge(dst, src)
}
func (m *NewUserResponse) XXX_Size() int {
	return xxx_messageInfo_NewUserResponse.Size(m)
}
func (m *NewUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NewUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NewUserResponse proto.InternalMessageInfo

func (m *NewUserResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *NewUserResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*NewUserRequest)(nil), "pb.NewUserRequest")
	proto.RegisterType((*NewUserResponse)(nil), "pb.NewUserResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for User service

type UserClient interface {
	NewUser(ctx context.Context, in *NewUserRequest, opts ...grpc.CallOption) (*NewUserResponse, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) NewUser(ctx context.Context, in *NewUserRequest, opts ...grpc.CallOption) (*NewUserResponse, error) {
	out := new(NewUserResponse)
	err := grpc.Invoke(ctx, "/pb.User/NewUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserServer interface {
	NewUser(context.Context, *NewUserRequest) (*NewUserResponse, error)
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_NewUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).NewUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.User/NewUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).NewUser(ctx, req.(*NewUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewUser",
			Handler:    _User_NewUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users.proto",
}

func init() { proto.RegisterFile("users.proto", fileDescriptor_users_4e3f99c9794f712d) }

var fileDescriptor_users_4e3f99c9794f712d = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x2d, 0x4e, 0x2d,
	0x2a, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x6a, 0x63, 0xe4, 0xe2,
	0xf3, 0x4b, 0x2d, 0x0f, 0x2d, 0x4e, 0x2d, 0x0a, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12,
	0xe2, 0x62, 0xf1, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85,
	0xa4, 0xb8, 0x38, 0x7c, 0x12, 0x8b, 0x4b, 0xc0, 0xe2, 0x4c, 0x60, 0x71, 0x38, 0x5f, 0x48, 0x84,
	0x8b, 0xd5, 0x35, 0x37, 0x31, 0x33, 0x47, 0x82, 0x19, 0x2c, 0x01, 0xe1, 0x80, 0x74, 0x04, 0x24,
	0x16, 0x17, 0x97, 0xe7, 0x17, 0xa5, 0x48, 0xb0, 0x40, 0x74, 0xc0, 0xf8, 0x20, 0x1b, 0x82, 0xf2,
	0x73, 0x52, 0x25, 0x58, 0x21, 0x36, 0x80, 0xd8, 0x4a, 0xc6, 0x5c, 0xfc, 0x70, 0x77, 0x14, 0x17,
	0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0xf1, 0x71, 0x31, 0x79, 0xa6, 0x40, 0x9d, 0xc1, 0xe4, 0x99, 0x22,
	0x24, 0xc0, 0xc5, 0xec, 0x5a, 0x54, 0x04, 0xb5, 0x1f, 0xc4, 0x34, 0xb2, 0xe1, 0x62, 0x01, 0xe9,
	0x10, 0x32, 0xe1, 0x62, 0x87, 0x6a, 0x16, 0x12, 0xd2, 0x2b, 0x48, 0xd2, 0x43, 0xf5, 0x91, 0x94,
	0x30, 0x8a, 0x18, 0xc4, 0x74, 0x25, 0x86, 0x24, 0x36, 0x70, 0x30, 0x18, 0x03, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x22, 0x2d, 0x3e, 0x55, 0x15, 0x01, 0x00, 0x00,
}
