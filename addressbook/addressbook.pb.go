// Code generated by protoc-gen-go. DO NOT EDIT.
// source: addressbook.proto

/*
Package addressbook is a generated protocol buffer package.

It is generated from these files:
	addressbook.proto

It has these top-level messages:
	Id
	Person
	Limit
	AddressBook
	MsgReply
*/
package addressbook

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

type Person_PhoneType int32

const (
	Person_MOBILE Person_PhoneType = 0
	Person_HOME   Person_PhoneType = 1
	Person_WORK   Person_PhoneType = 2
)

var Person_PhoneType_name = map[int32]string{
	0: "MOBILE",
	1: "HOME",
	2: "WORK",
}
var Person_PhoneType_value = map[string]int32{
	"MOBILE": 0,
	"HOME":   1,
	"WORK":   2,
}

func (x Person_PhoneType) String() string {
	return proto.EnumName(Person_PhoneType_name, int32(x))
}
func (Person_PhoneType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type Id struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *Id) Reset()                    { *m = Id{} }
func (m *Id) String() string            { return proto.CompactTextString(m) }
func (*Id) ProtoMessage()               {}
func (*Id) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Id) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Person struct {
	Name   string                `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Id     int32                 `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Email  string                `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Phones []*Person_PhoneNumber `protobuf:"bytes,4,rep,name=phones" json:"phones,omitempty"`
}

func (m *Person) Reset()                    { *m = Person{} }
func (m *Person) String() string            { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()               {}
func (*Person) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Person) GetPhones() []*Person_PhoneNumber {
	if m != nil {
		return m.Phones
	}
	return nil
}

type Person_PhoneNumber struct {
	Number string           `protobuf:"bytes,1,opt,name=number" json:"number,omitempty"`
	Type   Person_PhoneType `protobuf:"varint,2,opt,name=type,enum=addressbook.Person_PhoneType" json:"type,omitempty"`
}

func (m *Person_PhoneNumber) Reset()                    { *m = Person_PhoneNumber{} }
func (m *Person_PhoneNumber) String() string            { return proto.CompactTextString(m) }
func (*Person_PhoneNumber) ProtoMessage()               {}
func (*Person_PhoneNumber) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Person_PhoneNumber) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *Person_PhoneNumber) GetType() Person_PhoneType {
	if m != nil {
		return m.Type
	}
	return Person_MOBILE
}

type Limit struct {
	Start int32 `protobuf:"varint,1,opt,name=start" json:"start,omitempty"`
	End   int32 `protobuf:"varint,2,opt,name=end" json:"end,omitempty"`
}

func (m *Limit) Reset()                    { *m = Limit{} }
func (m *Limit) String() string            { return proto.CompactTextString(m) }
func (*Limit) ProtoMessage()               {}
func (*Limit) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Limit) GetStart() int32 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *Limit) GetEnd() int32 {
	if m != nil {
		return m.End
	}
	return 0
}

type AddressBook struct {
	People []*Person `protobuf:"bytes,1,rep,name=people" json:"people,omitempty"`
}

func (m *AddressBook) Reset()                    { *m = AddressBook{} }
func (m *AddressBook) String() string            { return proto.CompactTextString(m) }
func (*AddressBook) ProtoMessage()               {}
func (*AddressBook) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AddressBook) GetPeople() []*Person {
	if m != nil {
		return m.People
	}
	return nil
}

type MsgReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *MsgReply) Reset()                    { *m = MsgReply{} }
func (m *MsgReply) String() string            { return proto.CompactTextString(m) }
func (*MsgReply) ProtoMessage()               {}
func (*MsgReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *MsgReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Id)(nil), "addressbook.Id")
	proto.RegisterType((*Person)(nil), "addressbook.Person")
	proto.RegisterType((*Person_PhoneNumber)(nil), "addressbook.Person.PhoneNumber")
	proto.RegisterType((*Limit)(nil), "addressbook.Limit")
	proto.RegisterType((*AddressBook)(nil), "addressbook.AddressBook")
	proto.RegisterType((*MsgReply)(nil), "addressbook.MsgReply")
	proto.RegisterEnum("addressbook.Person_PhoneType", Person_PhoneType_name, Person_PhoneType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AddressServer service

type AddressServerClient interface {
	SavePerson(ctx context.Context, in *Person, opts ...grpc.CallOption) (*MsgReply, error)
	DelePerson(ctx context.Context, in *Id, opts ...grpc.CallOption) (*MsgReply, error)
	GetPersion(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Person, error)
	GetAddrBook(ctx context.Context, in *Limit, opts ...grpc.CallOption) (*AddressBook, error)
}

type addressServerClient struct {
	cc *grpc.ClientConn
}

func NewAddressServerClient(cc *grpc.ClientConn) AddressServerClient {
	return &addressServerClient{cc}
}

func (c *addressServerClient) SavePerson(ctx context.Context, in *Person, opts ...grpc.CallOption) (*MsgReply, error) {
	out := new(MsgReply)
	err := grpc.Invoke(ctx, "/addressbook.AddressServer/SavePerson", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressServerClient) DelePerson(ctx context.Context, in *Id, opts ...grpc.CallOption) (*MsgReply, error) {
	out := new(MsgReply)
	err := grpc.Invoke(ctx, "/addressbook.AddressServer/DelePerson", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressServerClient) GetPersion(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Person, error) {
	out := new(Person)
	err := grpc.Invoke(ctx, "/addressbook.AddressServer/GetPersion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressServerClient) GetAddrBook(ctx context.Context, in *Limit, opts ...grpc.CallOption) (*AddressBook, error) {
	out := new(AddressBook)
	err := grpc.Invoke(ctx, "/addressbook.AddressServer/GetAddrBook", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AddressServer service

type AddressServerServer interface {
	SavePerson(context.Context, *Person) (*MsgReply, error)
	DelePerson(context.Context, *Id) (*MsgReply, error)
	GetPersion(context.Context, *Id) (*Person, error)
	GetAddrBook(context.Context, *Limit) (*AddressBook, error)
}

func RegisterAddressServerServer(s *grpc.Server, srv AddressServerServer) {
	s.RegisterService(&_AddressServer_serviceDesc, srv)
}

func _AddressServer_SavePerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Person)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServerServer).SavePerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/addressbook.AddressServer/SavePerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServerServer).SavePerson(ctx, req.(*Person))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressServer_DelePerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServerServer).DelePerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/addressbook.AddressServer/DelePerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServerServer).DelePerson(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressServer_GetPersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServerServer).GetPersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/addressbook.AddressServer/GetPersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServerServer).GetPersion(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressServer_GetAddrBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Limit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServerServer).GetAddrBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/addressbook.AddressServer/GetAddrBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServerServer).GetAddrBook(ctx, req.(*Limit))
	}
	return interceptor(ctx, in, info, handler)
}

var _AddressServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "addressbook.AddressServer",
	HandlerType: (*AddressServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SavePerson",
			Handler:    _AddressServer_SavePerson_Handler,
		},
		{
			MethodName: "DelePerson",
			Handler:    _AddressServer_DelePerson_Handler,
		},
		{
			MethodName: "GetPersion",
			Handler:    _AddressServer_GetPersion_Handler,
		},
		{
			MethodName: "GetAddrBook",
			Handler:    _AddressServer_GetAddrBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "addressbook.proto",
}

func init() { proto.RegisterFile("addressbook.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x92, 0xc1, 0xab, 0xd3, 0x40,
	0x10, 0xc6, 0x93, 0x34, 0x8d, 0xef, 0x4d, 0xf0, 0x19, 0xc7, 0xa7, 0x84, 0x82, 0x58, 0x16, 0x0f,
	0x85, 0x42, 0xc5, 0x2a, 0x0a, 0x05, 0x0f, 0x16, 0x4b, 0x2d, 0xb6, 0xb6, 0xa4, 0x82, 0x5e, 0x53,
	0x32, 0xd4, 0xd0, 0x24, 0x1b, 0xb2, 0xb1, 0xd0, 0xab, 0x7f, 0xb5, 0x47, 0xd9, 0xcd, 0xe6, 0x91,
	0x40, 0x7a, 0x9b, 0xf9, 0xf2, 0x9b, 0x9d, 0x6f, 0xbf, 0x2c, 0x3c, 0x0d, 0xa3, 0xa8, 0x20, 0x21,
	0x0e, 0x9c, 0x9f, 0x26, 0x79, 0xc1, 0x4b, 0x8e, 0x6e, 0x43, 0x62, 0xf7, 0x60, 0xad, 0x22, 0xbc,
	0x03, 0x2b, 0x8e, 0x7c, 0x73, 0x68, 0x8e, 0xfa, 0x81, 0x15, 0x47, 0xec, 0xaf, 0x05, 0xce, 0x8e,
	0x0a, 0xc1, 0x33, 0x44, 0xb0, 0xb3, 0x30, 0x25, 0xf5, 0xf1, 0x36, 0x50, 0xb5, 0xc6, 0xad, 0x1a,
	0xc7, 0x7b, 0xe8, 0x53, 0x1a, 0xc6, 0x89, 0xdf, 0x53, 0x50, 0xd5, 0xe0, 0x47, 0x70, 0xf2, 0xdf,
	0x3c, 0x23, 0xe1, 0xdb, 0xc3, 0xde, 0xc8, 0x9d, 0xbe, 0x9a, 0x34, 0xbd, 0x54, 0xc7, 0x4f, 0x76,
	0x92, 0xf8, 0xfe, 0x27, 0x3d, 0x50, 0x11, 0x68, 0x7c, 0xf0, 0x0b, 0xdc, 0x86, 0x8c, 0x2f, 0xc0,
	0xc9, 0x54, 0xa5, 0x3d, 0xe8, 0x0e, 0xdf, 0x82, 0x5d, 0x5e, 0x72, 0x52, 0x3e, 0xee, 0xa6, 0x2f,
	0xaf, 0x9e, 0xfe, 0xe3, 0x92, 0x53, 0xa0, 0x50, 0x36, 0x86, 0xdb, 0x07, 0x09, 0x01, 0x9c, 0xcd,
	0x76, 0xbe, 0x5a, 0x2f, 0x3c, 0x03, 0x6f, 0xc0, 0xfe, 0xba, 0xdd, 0x2c, 0x3c, 0x53, 0x56, 0x3f,
	0xb7, 0xc1, 0x37, 0xcf, 0x62, 0x6f, 0xa0, 0xbf, 0x8e, 0xd3, 0xb8, 0x94, 0xd7, 0x13, 0x65, 0x58,
	0x94, 0x3a, 0xa0, 0xaa, 0x41, 0x0f, 0x7a, 0x94, 0xd5, 0x29, 0xc8, 0x92, 0xcd, 0xc0, 0xfd, 0x5c,
	0x79, 0x98, 0x73, 0x7e, 0xc2, 0x31, 0x38, 0x39, 0xf1, 0x3c, 0x91, 0xd9, 0xc9, 0xfb, 0x3f, 0xeb,
	0x70, 0x18, 0x68, 0x84, 0xbd, 0x86, 0x9b, 0x8d, 0x38, 0x06, 0x94, 0x27, 0x17, 0xf4, 0xe1, 0x51,
	0x4a, 0x42, 0x84, 0xc7, 0x3a, 0xf5, 0xba, 0x9d, 0xfe, 0x33, 0xe1, 0xb1, 0x5e, 0xb1, 0xa7, 0xe2,
	0x4c, 0x05, 0xce, 0x00, 0xf6, 0xe1, 0x99, 0xf4, 0xcf, 0xea, 0x5a, 0x31, 0x78, 0xde, 0x12, 0xeb,
	0x2d, 0xcc, 0xc0, 0x0f, 0x00, 0x5f, 0x28, 0xa9, 0x67, 0x9f, 0xb4, 0xb0, 0x55, 0x74, 0x7d, 0xee,
	0x3d, 0xc0, 0x92, 0x4a, 0x39, 0x16, 0x77, 0xcd, 0x75, 0x99, 0x60, 0x06, 0x7e, 0x02, 0x77, 0x49,
	0xa5, 0x74, 0xaf, 0xd2, 0xc1, 0x16, 0xa5, 0x82, 0x1e, 0xf8, 0x2d, 0xad, 0x91, 0x25, 0x33, 0x0e,
	0x8e, 0x7a, 0xbc, 0xef, 0xfe, 0x07, 0x00, 0x00, 0xff, 0xff, 0x91, 0x6d, 0xf6, 0x7d, 0xd1, 0x02,
	0x00, 0x00,
}
