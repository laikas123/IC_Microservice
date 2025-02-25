// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ic_server.proto

package ic_server

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Point struct {
	X                    int32    `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    int32    `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Point) Reset()         { *m = Point{} }
func (m *Point) String() string { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()    {}
func (*Point) Descriptor() ([]byte, []int) {
	return fileDescriptor_568ea1eef316f0e1, []int{0}
}

func (m *Point) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Point.Unmarshal(m, b)
}
func (m *Point) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Point.Marshal(b, m, deterministic)
}
func (m *Point) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Point.Merge(m, src)
}
func (m *Point) XXX_Size() int {
	return xxx_messageInfo_Point.Size(m)
}
func (m *Point) XXX_DiscardUnknown() {
	xxx_messageInfo_Point.DiscardUnknown(m)
}

var xxx_messageInfo_Point proto.InternalMessageInfo

func (m *Point) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Point) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

//our map will be a cartesian coordinate system
//0,0 is the bottom left corner and 200, 300 is the top right corner
type Rectangle struct {
	// One corner of the rectangle.
	Lo *Point `protobuf:"bytes,1,opt,name=lo,proto3" json:"lo,omitempty"`
	// The other corner of the rectangle.
	Hi                   *Point   `protobuf:"bytes,2,opt,name=hi,proto3" json:"hi,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Rectangle) Reset()         { *m = Rectangle{} }
func (m *Rectangle) String() string { return proto.CompactTextString(m) }
func (*Rectangle) ProtoMessage()    {}
func (*Rectangle) Descriptor() ([]byte, []int) {
	return fileDescriptor_568ea1eef316f0e1, []int{1}
}

func (m *Rectangle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rectangle.Unmarshal(m, b)
}
func (m *Rectangle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rectangle.Marshal(b, m, deterministic)
}
func (m *Rectangle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rectangle.Merge(m, src)
}
func (m *Rectangle) XXX_Size() int {
	return xxx_messageInfo_Rectangle.Size(m)
}
func (m *Rectangle) XXX_DiscardUnknown() {
	xxx_messageInfo_Rectangle.DiscardUnknown(m)
}

var xxx_messageInfo_Rectangle proto.InternalMessageInfo

func (m *Rectangle) GetLo() *Point {
	if m != nil {
		return m.Lo
	}
	return nil
}

func (m *Rectangle) GetHi() *Point {
	if m != nil {
		return m.Hi
	}
	return nil
}

//when we need to compare where the ice cream man is to where he's heading
//or if we need to figure out what's the most cost efficient route
type TwoPoints struct {
	First                *Point   `protobuf:"bytes,1,opt,name=first,proto3" json:"first,omitempty"`
	Second               *Point   `protobuf:"bytes,2,opt,name=second,proto3" json:"second,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TwoPoints) Reset()         { *m = TwoPoints{} }
func (m *TwoPoints) String() string { return proto.CompactTextString(m) }
func (*TwoPoints) ProtoMessage()    {}
func (*TwoPoints) Descriptor() ([]byte, []int) {
	return fileDescriptor_568ea1eef316f0e1, []int{2}
}

func (m *TwoPoints) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TwoPoints.Unmarshal(m, b)
}
func (m *TwoPoints) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TwoPoints.Marshal(b, m, deterministic)
}
func (m *TwoPoints) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TwoPoints.Merge(m, src)
}
func (m *TwoPoints) XXX_Size() int {
	return xxx_messageInfo_TwoPoints.Size(m)
}
func (m *TwoPoints) XXX_DiscardUnknown() {
	xxx_messageInfo_TwoPoints.DiscardUnknown(m)
}

var xxx_messageInfo_TwoPoints proto.InternalMessageInfo

func (m *TwoPoints) GetFirst() *Point {
	if m != nil {
		return m.First
	}
	return nil
}

func (m *TwoPoints) GetSecond() *Point {
	if m != nil {
		return m.Second
	}
	return nil
}

type LocationStatus struct {
	Usersonline          int32      `protobuf:"varint,1,opt,name=usersonline,proto3" json:"usersonline,omitempty"`
	Locationtoserve      *Rectangle `protobuf:"bytes,2,opt,name=locationtoserve,proto3" json:"locationtoserve,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *LocationStatus) Reset()         { *m = LocationStatus{} }
func (m *LocationStatus) String() string { return proto.CompactTextString(m) }
func (*LocationStatus) ProtoMessage()    {}
func (*LocationStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_568ea1eef316f0e1, []int{3}
}

func (m *LocationStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocationStatus.Unmarshal(m, b)
}
func (m *LocationStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocationStatus.Marshal(b, m, deterministic)
}
func (m *LocationStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocationStatus.Merge(m, src)
}
func (m *LocationStatus) XXX_Size() int {
	return xxx_messageInfo_LocationStatus.Size(m)
}
func (m *LocationStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_LocationStatus.DiscardUnknown(m)
}

var xxx_messageInfo_LocationStatus proto.InternalMessageInfo

func (m *LocationStatus) GetUsersonline() int32 {
	if m != nil {
		return m.Usersonline
	}
	return 0
}

func (m *LocationStatus) GetLocationtoserve() *Rectangle {
	if m != nil {
		return m.Locationtoserve
	}
	return nil
}

type Number struct {
	Anumber              int32    `protobuf:"varint,1,opt,name=anumber,proto3" json:"anumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Number) Reset()         { *m = Number{} }
func (m *Number) String() string { return proto.CompactTextString(m) }
func (*Number) ProtoMessage()    {}
func (*Number) Descriptor() ([]byte, []int) {
	return fileDescriptor_568ea1eef316f0e1, []int{4}
}

func (m *Number) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Number.Unmarshal(m, b)
}
func (m *Number) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Number.Marshal(b, m, deterministic)
}
func (m *Number) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Number.Merge(m, src)
}
func (m *Number) XXX_Size() int {
	return xxx_messageInfo_Number.Size(m)
}
func (m *Number) XXX_DiscardUnknown() {
	xxx_messageInfo_Number.DiscardUnknown(m)
}

var xxx_messageInfo_Number proto.InternalMessageInfo

func (m *Number) GetAnumber() int32 {
	if m != nil {
		return m.Anumber
	}
	return 0
}

func init() {
	proto.RegisterType((*Point)(nil), "main.Point")
	proto.RegisterType((*Rectangle)(nil), "main.Rectangle")
	proto.RegisterType((*TwoPoints)(nil), "main.TwoPoints")
	proto.RegisterType((*LocationStatus)(nil), "main.LocationStatus")
	proto.RegisterType((*Number)(nil), "main.Number")
}

func init() { proto.RegisterFile("ic_server.proto", fileDescriptor_568ea1eef316f0e1) }

var fileDescriptor_568ea1eef316f0e1 = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x4d, 0x8b, 0x1a, 0x41,
	0x10, 0xcd, 0x48, 0x34, 0x58, 0x8a, 0x93, 0x74, 0x02, 0x91, 0x78, 0x31, 0xe3, 0x25, 0xa7, 0x41,
	0x4c, 0x48, 0xc8, 0x55, 0x13, 0x24, 0x20, 0x61, 0x76, 0xdc, 0xfb, 0xd2, 0xb6, 0xa5, 0x36, 0xf4,
	0x74, 0x49, 0x77, 0x8f, 0xab, 0x7f, 0x7a, 0x7f, 0xc3, 0x32, 0x9f, 0xac, 0x1f, 0xec, 0xad, 0x5e,
	0xbf, 0x7a, 0xef, 0x55, 0x35, 0x05, 0xbe, 0x14, 0x0f, 0x16, 0xcd, 0x01, 0x4d, 0xb8, 0x37, 0xe4,
	0x88, 0xbd, 0x4d, 0xb8, 0xd4, 0xc1, 0x08, 0x9a, 0x11, 0x49, 0xed, 0x58, 0x17, 0xbc, 0x63, 0xdf,
	0x1b, 0x7a, 0xdf, 0x9a, 0xb1, 0x77, 0xcc, 0xd0, 0xa9, 0xdf, 0x28, 0xd0, 0x29, 0xf8, 0x0b, 0xed,
	0x18, 0x85, 0xe3, 0x7a, 0xab, 0x90, 0x0d, 0xa0, 0xa1, 0x28, 0xef, 0xec, 0x4c, 0x3a, 0x61, 0x66,
	0x12, 0xe6, 0x0e, 0x71, 0x43, 0x51, 0x46, 0xee, 0x64, 0x2e, 0xbc, 0x24, 0x77, 0x32, 0x58, 0x42,
	0xfb, 0xfe, 0x91, 0x72, 0x6c, 0xd9, 0x57, 0x68, 0x6e, 0xa4, 0xb1, 0xee, 0x96, 0x53, 0xc1, 0xb0,
	0x11, 0xb4, 0x2c, 0x0a, 0xd2, 0xeb, 0x5b, 0x86, 0x25, 0x15, 0x24, 0xd0, 0x5b, 0x90, 0xe0, 0x4e,
	0x92, 0x5e, 0x3a, 0xee, 0x52, 0xcb, 0x86, 0xd0, 0x49, 0x2d, 0x1a, 0x4b, 0x5a, 0x49, 0x8d, 0xe5,
	0x4e, 0x2f, 0x9f, 0xd8, 0x6f, 0xf0, 0x55, 0xa9, 0x71, 0x94, 0x7f, 0x4a, 0x99, 0xe0, 0x17, 0x09,
	0xf5, 0xb2, 0xf1, 0x65, 0x5f, 0x10, 0x40, 0xeb, 0x7f, 0x9a, 0xac, 0xd0, 0xb0, 0x3e, 0xbc, 0xe3,
	0x3a, 0x2f, 0xcb, 0x88, 0x0a, 0x4e, 0x9e, 0x3c, 0xf8, 0xf8, 0x6f, 0x36, 0xe3, 0x4a, 0xa4, 0x8a,
	0x3b, 0x32, 0x4b, 0x34, 0x07, 0x29, 0x90, 0xfd, 0x84, 0xde, 0x5d, 0x8a, 0xe6, 0x54, 0xcd, 0x6b,
	0xd9, 0xa7, 0x22, 0xef, 0x7c, 0x81, 0x2f, 0xdd, 0xe2, 0xb5, 0xc8, 0x09, 0xde, 0xb0, 0x1f, 0xf0,
	0xa1, 0x32, 0xc3, 0x3f, 0xd2, 0x3a, 0xae, 0x05, 0xb2, 0x72, 0xd4, 0xfa, 0x43, 0xaf, 0x54, 0x63,
	0x78, 0x5f, 0xab, 0xe6, 0xdc, 0x2e, 0xc8, 0x5a, 0x76, 0xd6, 0x73, 0xa5, 0xf8, 0x05, 0x9f, 0x6b,
	0x45, 0x35, 0x52, 0x64, 0x68, 0x23, 0xdd, 0xeb, 0xc2, 0xe9, 0x18, 0x06, 0x92, 0xc2, 0xad, 0xd9,
	0x8b, 0x10, 0x8f, 0x3c, 0xd9, 0x2b, 0xb4, 0xa1, 0xa1, 0xd4, 0xe1, 0x36, 0x95, 0x6b, 0x9c, 0xfa,
	0x71, 0x56, 0xcf, 0xb3, 0x3a, 0xca, 0x4e, 0x2f, 0xf2, 0x56, 0xad, 0xfc, 0x06, 0xbf, 0x3f, 0x07,
	0x00, 0x00, 0xff, 0xff, 0x8b, 0x80, 0x45, 0xc6, 0x96, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4


//Good thing to note is these interfaces and structs come with gRPC plugin being applied
// ICCalculatorServiceClient is the client API for ICCalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ICCalculatorServiceClient interface {
	QueryLocations(ctx context.Context, in *LocationStatus, opts ...grpc.CallOption) (*Number, error)
	CalculateDistance(ctx context.Context, in *TwoPoints, opts ...grpc.CallOption) (*Number, error)
	CalculateGasLoss(ctx context.Context, in *Number, opts ...grpc.CallOption) (*Number, error)
	CalculateLocationProfit(ctx context.Context, in *Number, opts ...grpc.CallOption) (*Number, error)
}

type iCCalculatorServiceClient struct {
	cc *grpc.ClientConn
}

func NewICCalculatorServiceClient(cc *grpc.ClientConn) ICCalculatorServiceClient {
	return &iCCalculatorServiceClient{cc}
}

func (c *iCCalculatorServiceClient) QueryLocations(ctx context.Context, in *LocationStatus, opts ...grpc.CallOption) (*Number, error) {
	out := new(Number)
	err := c.cc.Invoke(ctx, "/main.ICCalculatorService/QueryLocations", in, out, opts...)
	if err != nil {
		return nil, err
	}


	return out, nil
}

func (c *iCCalculatorServiceClient) CalculateDistance(ctx context.Context, in *TwoPoints, opts ...grpc.CallOption) (*Number, error) {
	out := new(Number)
	err := c.cc.Invoke(ctx, "/main.ICCalculatorService/CalculateDistance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iCCalculatorServiceClient) CalculateGasLoss(ctx context.Context, in *Number, opts ...grpc.CallOption) (*Number, error) {
	out := new(Number)
	err := c.cc.Invoke(ctx, "/main.ICCalculatorService/CalculateGasLoss", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iCCalculatorServiceClient) CalculateLocationProfit(ctx context.Context, in *Number, opts ...grpc.CallOption) (*Number, error) {
	out := new(Number)
	err := c.cc.Invoke(ctx, "/main.ICCalculatorService/CalculateLocationProfit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ICCalculatorServiceServer is the server API for ICCalculatorService service.
type ICCalculatorServiceServer interface {
	QueryLocations(context.Context, *LocationStatus) (*Number, error)
	CalculateDistance(context.Context, *TwoPoints) (*Number, error)
	CalculateGasLoss(context.Context, *Number) (*Number, error)
	CalculateLocationProfit(context.Context, *Number) (*Number, error)
}

// UnimplementedICCalculatorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedICCalculatorServiceServer struct {
}

func (*UnimplementedICCalculatorServiceServer) QueryLocations(ctx context.Context, req *LocationStatus) (*Number, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryLocations not implemented")
}
func (*UnimplementedICCalculatorServiceServer) CalculateDistance(ctx context.Context, req *TwoPoints) (*Number, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalculateDistance not implemented")
}
func (*UnimplementedICCalculatorServiceServer) CalculateGasLoss(ctx context.Context, req *Number) (*Number, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalculateGasLoss not implemented")
}
func (*UnimplementedICCalculatorServiceServer) CalculateLocationProfit(ctx context.Context, req *Number) (*Number, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalculateLocationProfit not implemented")
}

func RegisterICCalculatorServiceServer(s *grpc.Server, srv ICCalculatorServiceServer) {
	s.RegisterService(&_ICCalculatorService_serviceDesc, srv)
}

func _ICCalculatorService_QueryLocations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LocationStatus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ICCalculatorServiceServer).QueryLocations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.ICCalculatorService/QueryLocations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ICCalculatorServiceServer).QueryLocations(ctx, req.(*LocationStatus))
	}
	return interceptor(ctx, in, info, handler)
}

func _ICCalculatorService_CalculateDistance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TwoPoints)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ICCalculatorServiceServer).CalculateDistance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.ICCalculatorService/CalculateDistance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ICCalculatorServiceServer).CalculateDistance(ctx, req.(*TwoPoints))
	}
	return interceptor(ctx, in, info, handler)
}

func _ICCalculatorService_CalculateGasLoss_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Number)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ICCalculatorServiceServer).CalculateGasLoss(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.ICCalculatorService/CalculateGasLoss",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ICCalculatorServiceServer).CalculateGasLoss(ctx, req.(*Number))
	}
	return interceptor(ctx, in, info, handler)
}

func _ICCalculatorService_CalculateLocationProfit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Number)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ICCalculatorServiceServer).CalculateLocationProfit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.ICCalculatorService/CalculateLocationProfit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ICCalculatorServiceServer).CalculateLocationProfit(ctx, req.(*Number))
	}
	return interceptor(ctx, in, info, handler)
}

var _ICCalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "main.ICCalculatorService",
	HandlerType: (*ICCalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryLocations",
			Handler:    _ICCalculatorService_QueryLocations_Handler,
		},
		{
			MethodName: "CalculateDistance",
			Handler:    _ICCalculatorService_CalculateDistance_Handler,
		},
		{
			MethodName: "CalculateGasLoss",
			Handler:    _ICCalculatorService_CalculateGasLoss_Handler,
		},
		{
			MethodName: "CalculateLocationProfit",
			Handler:    _ICCalculatorService_CalculateLocationProfit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ic_server.proto",
}
