// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pengguna.proto

package pengguna

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type PenggunaData struct {
	IdPengguna           int64    `protobuf:"varint,1,opt,name=id_pengguna,json=idPengguna,proto3" json:"id_pengguna,omitempty"`
	NamaPengguna         string   `protobuf:"bytes,2,opt,name=nama_pengguna,json=namaPengguna,proto3" json:"nama_pengguna,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	NomorTelepon         string   `protobuf:"bytes,4,opt,name=nomor_telepon,json=nomorTelepon,proto3" json:"nomor_telepon,omitempty"`
	Alamat               string   `protobuf:"bytes,5,opt,name=alamat,proto3" json:"alamat,omitempty"`
	Email                string   `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	Balance              float32  `protobuf:"fixed32,7,opt,name=balance,proto3" json:"balance,omitempty"`
	JenisKelamin         string   `protobuf:"bytes,8,opt,name=jenis_kelamin,json=jenisKelamin,proto3" json:"jenis_kelamin,omitempty"`
	TanggalLahir         string   `protobuf:"bytes,9,opt,name=tanggal_lahir,json=tanggalLahir,proto3" json:"tanggal_lahir,omitempty"`
	ErrorsMessage        []string `protobuf:"bytes,10,rep,name=errors_message,json=errorsMessage,proto3" json:"errors_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PenggunaData) Reset()         { *m = PenggunaData{} }
func (m *PenggunaData) String() string { return proto.CompactTextString(m) }
func (*PenggunaData) ProtoMessage()    {}
func (*PenggunaData) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4d756069402c44b, []int{0}
}

func (m *PenggunaData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PenggunaData.Unmarshal(m, b)
}
func (m *PenggunaData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PenggunaData.Marshal(b, m, deterministic)
}
func (m *PenggunaData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PenggunaData.Merge(m, src)
}
func (m *PenggunaData) XXX_Size() int {
	return xxx_messageInfo_PenggunaData.Size(m)
}
func (m *PenggunaData) XXX_DiscardUnknown() {
	xxx_messageInfo_PenggunaData.DiscardUnknown(m)
}

var xxx_messageInfo_PenggunaData proto.InternalMessageInfo

func (m *PenggunaData) GetIdPengguna() int64 {
	if m != nil {
		return m.IdPengguna
	}
	return 0
}

func (m *PenggunaData) GetNamaPengguna() string {
	if m != nil {
		return m.NamaPengguna
	}
	return ""
}

func (m *PenggunaData) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *PenggunaData) GetNomorTelepon() string {
	if m != nil {
		return m.NomorTelepon
	}
	return ""
}

func (m *PenggunaData) GetAlamat() string {
	if m != nil {
		return m.Alamat
	}
	return ""
}

func (m *PenggunaData) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *PenggunaData) GetBalance() float32 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *PenggunaData) GetJenisKelamin() string {
	if m != nil {
		return m.JenisKelamin
	}
	return ""
}

func (m *PenggunaData) GetTanggalLahir() string {
	if m != nil {
		return m.TanggalLahir
	}
	return ""
}

func (m *PenggunaData) GetErrorsMessage() []string {
	if m != nil {
		return m.ErrorsMessage
	}
	return nil
}

func init() {
	proto.RegisterType((*PenggunaData)(nil), "pengguna.penggunaData")
}

func init() { proto.RegisterFile("pengguna.proto", fileDescriptor_e4d756069402c44b) }

var fileDescriptor_e4d756069402c44b = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xcd, 0x4a, 0x33, 0x31,
	0x14, 0x86, 0xbf, 0x69, 0xbf, 0xfe, 0x9d, 0xfe, 0x28, 0x41, 0x4a, 0xe8, 0xc6, 0xa1, 0x22, 0xcc,
	0xaa, 0x0b, 0xbd, 0x00, 0x11, 0xba, 0x53, 0x41, 0x46, 0xf7, 0xc3, 0x69, 0x7b, 0x88, 0xd1, 0x4c,
	0x32, 0x24, 0x51, 0xef, 0xd2, 0x1b, 0xf0, 0x66, 0x64, 0x92, 0x99, 0xd6, 0x8d, 0x1b, 0x5d, 0x3e,
	0xcf, 0x79, 0x79, 0x0f, 0x9c, 0x04, 0x66, 0x15, 0x69, 0x21, 0x5e, 0x35, 0xae, 0x2a, 0x6b, 0xbc,
	0x61, 0xc3, 0x96, 0x97, 0x1f, 0x1d, 0x98, 0xb4, 0xb0, 0x46, 0x8f, 0xec, 0x14, 0xc6, 0x72, 0x57,
	0xb4, 0x8a, 0x27, 0x69, 0x92, 0x75, 0x73, 0x90, 0xbb, 0xfb, 0xc6, 0xb0, 0x33, 0x98, 0x6a, 0x2c,
	0xf1, 0x10, 0xe9, 0xa4, 0x49, 0x36, 0xca, 0x27, 0xb5, 0xdc, 0x87, 0x16, 0x30, 0xac, 0xd0, 0xb9,
	0x77, 0x63, 0x77, 0xbc, 0x1b, 0xe6, 0x7b, 0x0e, 0x05, 0xa6, 0x34, 0xb6, 0xf0, 0xa4, 0xa8, 0x32,
	0x9a, 0xff, 0x6f, 0x0a, 0x6a, 0xf9, 0x18, 0x1d, 0x9b, 0x43, 0x1f, 0x15, 0x96, 0xe8, 0x79, 0x2f,
	0x4c, 0x1b, 0x62, 0x27, 0xd0, 0xa3, 0x12, 0xa5, 0xe2, 0xfd, 0xa0, 0x23, 0x30, 0x0e, 0x83, 0x0d,
	0x2a, 0xd4, 0x5b, 0xe2, 0x83, 0x34, 0xc9, 0x3a, 0x79, 0x8b, 0xf5, 0xb2, 0x67, 0xd2, 0xd2, 0x15,
	0x2f, 0xa4, 0xb0, 0x94, 0x9a, 0x0f, 0xe3, 0xb2, 0x20, 0x6f, 0xa2, 0xab, 0x43, 0x1e, 0xb5, 0x10,
	0xa8, 0x0a, 0x85, 0x4f, 0xd2, 0xf2, 0x51, 0x0c, 0x35, 0xf2, 0xb6, 0x76, 0xec, 0x1c, 0x66, 0x64,
	0xad, 0xb1, 0xae, 0x28, 0xc9, 0x39, 0x14, 0xc4, 0x21, 0xed, 0x66, 0xa3, 0x7c, 0x1a, 0xed, 0x5d,
	0x94, 0x17, 0x9f, 0x09, 0x1c, 0xb5, 0xa7, 0x79, 0x20, 0xfb, 0x26, 0xb7, 0xc4, 0xae, 0x60, 0x6c,
	0x34, 0xed, 0x8f, 0x33, 0x5f, 0x1d, 0x9e, 0xe3, 0xdb, 0xe9, 0x17, 0x3f, 0xf8, 0xe5, 0x3f, 0x76,
	0x0d, 0x53, 0x65, 0x84, 0xd4, 0x7f, 0xa8, 0x58, 0xc3, 0xb1, 0x25, 0x21, 0x9d, 0x27, 0xfb, 0xfb,
	0x96, 0x4d, 0x3f, 0xfc, 0x9f, 0xcb, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x92, 0x68, 0x18, 0xb1,
	0x51, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PenggunaServiceClient is the client API for PenggunaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PenggunaServiceClient interface {
	OnePengguna(ctx context.Context, in *PenggunaData, opts ...grpc.CallOption) (*PenggunaData, error)
	LoginPengguna(ctx context.Context, in *PenggunaData, opts ...grpc.CallOption) (*PenggunaData, error)
	RegisterPengguna(ctx context.Context, in *PenggunaData, opts ...grpc.CallOption) (*PenggunaData, error)
}

type penggunaServiceClient struct {
	cc *grpc.ClientConn
}

func NewPenggunaServiceClient(cc *grpc.ClientConn) PenggunaServiceClient {
	return &penggunaServiceClient{cc}
}

func (c *penggunaServiceClient) OnePengguna(ctx context.Context, in *PenggunaData, opts ...grpc.CallOption) (*PenggunaData, error) {
	out := new(PenggunaData)
	err := c.cc.Invoke(ctx, "/pengguna.penggunaService/onePengguna", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *penggunaServiceClient) LoginPengguna(ctx context.Context, in *PenggunaData, opts ...grpc.CallOption) (*PenggunaData, error) {
	out := new(PenggunaData)
	err := c.cc.Invoke(ctx, "/pengguna.penggunaService/loginPengguna", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *penggunaServiceClient) RegisterPengguna(ctx context.Context, in *PenggunaData, opts ...grpc.CallOption) (*PenggunaData, error) {
	out := new(PenggunaData)
	err := c.cc.Invoke(ctx, "/pengguna.penggunaService/registerPengguna", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PenggunaServiceServer is the server API for PenggunaService service.
type PenggunaServiceServer interface {
	OnePengguna(context.Context, *PenggunaData) (*PenggunaData, error)
	LoginPengguna(context.Context, *PenggunaData) (*PenggunaData, error)
	RegisterPengguna(context.Context, *PenggunaData) (*PenggunaData, error)
}

func RegisterPenggunaServiceServer(s *grpc.Server, srv PenggunaServiceServer) {
	s.RegisterService(&_PenggunaService_serviceDesc, srv)
}

func _PenggunaService_OnePengguna_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PenggunaData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PenggunaServiceServer).OnePengguna(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pengguna.penggunaService/OnePengguna",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PenggunaServiceServer).OnePengguna(ctx, req.(*PenggunaData))
	}
	return interceptor(ctx, in, info, handler)
}

func _PenggunaService_LoginPengguna_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PenggunaData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PenggunaServiceServer).LoginPengguna(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pengguna.penggunaService/LoginPengguna",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PenggunaServiceServer).LoginPengguna(ctx, req.(*PenggunaData))
	}
	return interceptor(ctx, in, info, handler)
}

func _PenggunaService_RegisterPengguna_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PenggunaData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PenggunaServiceServer).RegisterPengguna(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pengguna.penggunaService/RegisterPengguna",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PenggunaServiceServer).RegisterPengguna(ctx, req.(*PenggunaData))
	}
	return interceptor(ctx, in, info, handler)
}

var _PenggunaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pengguna.penggunaService",
	HandlerType: (*PenggunaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "onePengguna",
			Handler:    _PenggunaService_OnePengguna_Handler,
		},
		{
			MethodName: "loginPengguna",
			Handler:    _PenggunaService_LoginPengguna_Handler,
		},
		{
			MethodName: "registerPengguna",
			Handler:    _PenggunaService_RegisterPengguna_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pengguna.proto",
}
