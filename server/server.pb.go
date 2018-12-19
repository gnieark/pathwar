// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: server/server.proto

package server // import "pathwar.pw/server"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import entity "pathwar.pw/entity"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Void struct {
}

func (m *Void) Reset()         { *m = Void{} }
func (m *Void) String() string { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()    {}
func (*Void) Descriptor() ([]byte, []int) {
	return fileDescriptor_server_b6a4b7434f49796f, []int{0}
}
func (m *Void) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Void) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Void.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Void) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Void.Merge(dst, src)
}
func (m *Void) XXX_Size() int {
	return m.Size()
}
func (m *Void) XXX_DiscardUnknown() {
	xxx_messageInfo_Void.DiscardUnknown(m)
}

var xxx_messageInfo_Void proto.InternalMessageInfo

type AuthenticateInput struct {
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (m *AuthenticateInput) Reset()         { *m = AuthenticateInput{} }
func (m *AuthenticateInput) String() string { return proto.CompactTextString(m) }
func (*AuthenticateInput) ProtoMessage()    {}
func (*AuthenticateInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_server_b6a4b7434f49796f, []int{1}
}
func (m *AuthenticateInput) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuthenticateInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthenticateInput.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *AuthenticateInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticateInput.Merge(dst, src)
}
func (m *AuthenticateInput) XXX_Size() int {
	return m.Size()
}
func (m *AuthenticateInput) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticateInput.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticateInput proto.InternalMessageInfo

func (m *AuthenticateInput) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AuthenticateInput) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type AuthenticateOutput struct {
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (m *AuthenticateOutput) Reset()         { *m = AuthenticateOutput{} }
func (m *AuthenticateOutput) String() string { return proto.CompactTextString(m) }
func (*AuthenticateOutput) ProtoMessage()    {}
func (*AuthenticateOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_server_b6a4b7434f49796f, []int{2}
}
func (m *AuthenticateOutput) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuthenticateOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthenticateOutput.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *AuthenticateOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticateOutput.Merge(dst, src)
}
func (m *AuthenticateOutput) XXX_Size() int {
	return m.Size()
}
func (m *AuthenticateOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticateOutput.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticateOutput proto.InternalMessageInfo

func (m *AuthenticateOutput) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*Void)(nil), "pathwar.server.Void")
	proto.RegisterType((*AuthenticateInput)(nil), "pathwar.server.AuthenticateInput")
	proto.RegisterType((*AuthenticateOutput)(nil), "pathwar.server.AuthenticateOutput")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServerClient is the client API for Server service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServerClient interface {
	Authenticate(ctx context.Context, in *AuthenticateInput, opts ...grpc.CallOption) (*AuthenticateOutput, error)
	Ping(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Void, error)
	Session(ctx context.Context, in *Void, opts ...grpc.CallOption) (*entity.Session, error)
}

type serverClient struct {
	cc *grpc.ClientConn
}

func NewServerClient(cc *grpc.ClientConn) ServerClient {
	return &serverClient{cc}
}

func (c *serverClient) Authenticate(ctx context.Context, in *AuthenticateInput, opts ...grpc.CallOption) (*AuthenticateOutput, error) {
	out := new(AuthenticateOutput)
	err := c.cc.Invoke(ctx, "/pathwar.server.Server/Authenticate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverClient) Ping(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/pathwar.server.Server/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverClient) Session(ctx context.Context, in *Void, opts ...grpc.CallOption) (*entity.Session, error) {
	out := new(entity.Session)
	err := c.cc.Invoke(ctx, "/pathwar.server.Server/Session", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerServer is the server API for Server service.
type ServerServer interface {
	Authenticate(context.Context, *AuthenticateInput) (*AuthenticateOutput, error)
	Ping(context.Context, *Void) (*Void, error)
	Session(context.Context, *Void) (*entity.Session, error)
}

func RegisterServerServer(s *grpc.Server, srv ServerServer) {
	s.RegisterService(&_Server_serviceDesc, srv)
}

func _Server_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pathwar.server.Server/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).Authenticate(ctx, req.(*AuthenticateInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Server_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pathwar.server.Server/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).Ping(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Server_Session_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).Session(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pathwar.server.Server/Session",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).Session(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

var _Server_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pathwar.server.Server",
	HandlerType: (*ServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authenticate",
			Handler:    _Server_Authenticate_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Server_Ping_Handler,
		},
		{
			MethodName: "Session",
			Handler:    _Server_Session_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/server.proto",
}

func (m *Void) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Void) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *AuthenticateInput) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthenticateInput) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Username) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintServer(dAtA, i, uint64(len(m.Username)))
		i += copy(dAtA[i:], m.Username)
	}
	if len(m.Password) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintServer(dAtA, i, uint64(len(m.Password)))
		i += copy(dAtA[i:], m.Password)
	}
	return i, nil
}

func (m *AuthenticateOutput) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthenticateOutput) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Token) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintServer(dAtA, i, uint64(len(m.Token)))
		i += copy(dAtA[i:], m.Token)
	}
	return i, nil
}

func encodeVarintServer(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Void) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *AuthenticateInput) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Username)
	if l > 0 {
		n += 1 + l + sovServer(uint64(l))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovServer(uint64(l))
	}
	return n
}

func (m *AuthenticateOutput) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovServer(uint64(l))
	}
	return n
}

func sovServer(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozServer(x uint64) (n int) {
	return sovServer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Void) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowServer
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Void: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Void: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipServer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthServer
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AuthenticateInput) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowServer
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AuthenticateInput: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthenticateInput: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Username", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthServer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Username = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Password", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthServer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Password = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipServer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthServer
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AuthenticateOutput) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowServer
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AuthenticateOutput: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthenticateOutput: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthServer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipServer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthServer
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipServer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowServer
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowServer
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowServer
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthServer
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowServer
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipServer(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthServer = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowServer   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("server/server.proto", fileDescriptor_server_b6a4b7434f49796f) }

var fileDescriptor_server_b6a4b7434f49796f = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x4d, 0x4b, 0x33, 0x31,
	0x14, 0x85, 0x3b, 0xa5, 0x5f, 0x6f, 0x78, 0x2b, 0x36, 0x16, 0x2c, 0x83, 0x0c, 0x9a, 0x95, 0x54,
	0x9c, 0x80, 0xee, 0xdc, 0xd5, 0x9d, 0xba, 0x50, 0x2c, 0xb8, 0x70, 0x97, 0xb6, 0x71, 0x1a, 0xb4,
	0xb9, 0x61, 0x92, 0x58, 0xdc, 0xba, 0x72, 0x29, 0xf8, 0xa7, 0x5c, 0x16, 0xdc, 0xb8, 0x94, 0xd6,
	0x1f, 0x22, 0x93, 0x4c, 0x4b, 0xfd, 0x5c, 0xcd, 0x9c, 0x73, 0xcf, 0x7d, 0x66, 0x72, 0x6f, 0xd0,
	0x9a, 0xe6, 0xe9, 0x2d, 0x4f, 0xa9, 0x7f, 0xc4, 0x2a, 0x05, 0x03, 0x78, 0x45, 0x31, 0x33, 0x1c,
	0xb3, 0x34, 0xf6, 0x6e, 0xb8, 0x91, 0x00, 0x24, 0x37, 0x9c, 0x32, 0x25, 0x28, 0x93, 0x12, 0x0c,
	0x33, 0x02, 0xa4, 0xf6, 0xe9, 0xb0, 0xc9, 0xa5, 0x11, 0xe6, 0x8e, 0x6a, 0xae, 0xb5, 0x00, 0x99,
	0xbb, 0xbb, 0x89, 0x30, 0x43, 0xdb, 0x8b, 0xfb, 0x30, 0xa2, 0x09, 0x24, 0x40, 0x9d, 0xdd, 0xb3,
	0x57, 0x4e, 0x39, 0xe1, 0xde, 0x7c, 0x9c, 0x54, 0x50, 0xe9, 0x02, 0xc4, 0x80, 0x9c, 0xa0, 0x46,
	0xc7, 0x9a, 0x61, 0x86, 0xec, 0x33, 0xc3, 0x8f, 0xa4, 0xb2, 0x06, 0x87, 0xa8, 0x66, 0x35, 0x4f,
	0x25, 0x1b, 0xf1, 0x56, 0xb0, 0x19, 0x6c, 0xff, 0x3b, 0x5f, 0xe8, 0xac, 0xa6, 0x98, 0xd6, 0x63,
	0x48, 0x07, 0xad, 0xa2, 0xaf, 0xcd, 0x35, 0x69, 0x23, 0xbc, 0x0c, 0x3b, 0xb5, 0x26, 0xa3, 0x35,
	0x51, 0xd9, 0xc0, 0x35, 0x97, 0x39, 0xca, 0x8b, 0xbd, 0x87, 0x22, 0xaa, 0x74, 0xdd, 0x71, 0x31,
	0xa0, 0xff, 0xcb, 0x6d, 0x78, 0x2b, 0xfe, 0x3c, 0x8f, 0xf8, 0xdb, 0x1f, 0x86, 0xe4, 0xaf, 0x88,
	0xff, 0x2e, 0x69, 0xdd, 0xbf, 0xbc, 0x3f, 0x15, 0x31, 0xa9, 0x53, 0xb6, 0x54, 0x3c, 0x08, 0xda,
	0xb8, 0x83, 0x4a, 0x67, 0x42, 0x26, 0xb8, 0xf9, 0x95, 0x92, 0x8d, 0x24, 0xfc, 0xd1, 0x25, 0x75,
	0x47, 0xab, 0xe2, 0x32, 0x55, 0x59, 0xeb, 0x31, 0xaa, 0x76, 0xfd, 0xfc, 0x7f, 0xa1, 0xac, 0x2f,
	0x5c, 0xbf, 0xae, 0x38, 0x8f, 0x93, 0x55, 0x07, 0x42, 0xb8, 0x36, 0x5f, 0xe0, 0xe1, 0xce, 0xf3,
	0x34, 0x0a, 0x26, 0xd3, 0x28, 0x78, 0x9b, 0x46, 0xc1, 0xe3, 0x2c, 0x2a, 0x4c, 0x66, 0x51, 0xe1,
	0x75, 0x16, 0x15, 0x2e, 0x1b, 0x73, 0x86, 0x1a, 0xe7, 0x37, 0xa6, 0x57, 0x71, 0xfb, 0xdb, 0xff,
	0x08, 0x00, 0x00, 0xff, 0xff, 0x90, 0xa7, 0x3c, 0xc2, 0x49, 0x02, 0x00, 0x00,
}
