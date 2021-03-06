// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package position

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// PositionClient is the client API for Position service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PositionClient interface {
	Position(ctx context.Context, in *UserPositionFromUUIDRequest, opts ...grpc.CallOption) (*PositionResponse, error)
	AllPosition(ctx context.Context, in *AllPositionsRequest, opts ...grpc.CallOption) (*AllPositionsResponse, error)
}

type positionClient struct {
	cc grpc.ClientConnInterface
}

func NewPositionClient(cc grpc.ClientConnInterface) PositionClient {
	return &positionClient{cc}
}

func (c *positionClient) Position(ctx context.Context, in *UserPositionFromUUIDRequest, opts ...grpc.CallOption) (*PositionResponse, error) {
	out := new(PositionResponse)
	err := c.cc.Invoke(ctx, "/position.Position/Position", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *positionClient) AllPosition(ctx context.Context, in *AllPositionsRequest, opts ...grpc.CallOption) (*AllPositionsResponse, error) {
	out := new(AllPositionsResponse)
	err := c.cc.Invoke(ctx, "/position.Position/AllPosition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PositionServer is the server API for Position service.
// All implementations must embed UnimplementedPositionServer
// for forward compatibility
type PositionServer interface {
	Position(context.Context, *UserPositionFromUUIDRequest) (*PositionResponse, error)
	AllPosition(context.Context, *AllPositionsRequest) (*AllPositionsResponse, error)
	mustEmbedUnimplementedPositionServer()
}

// UnimplementedPositionServer must be embedded to have forward compatible implementations.
type UnimplementedPositionServer struct {
}

func (UnimplementedPositionServer) Position(context.Context, *UserPositionFromUUIDRequest) (*PositionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Position not implemented")
}
func (UnimplementedPositionServer) AllPosition(context.Context, *AllPositionsRequest) (*AllPositionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllPosition not implemented")
}
func (UnimplementedPositionServer) mustEmbedUnimplementedPositionServer() {}

// UnsafePositionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PositionServer will
// result in compilation errors.
type UnsafePositionServer interface {
	mustEmbedUnimplementedPositionServer()
}

func RegisterPositionServer(s grpc.ServiceRegistrar, srv PositionServer) {
	s.RegisterService(&_Position_serviceDesc, srv)
}

func _Position_Position_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPositionFromUUIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServer).Position(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/position.Position/Position",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServer).Position(ctx, req.(*UserPositionFromUUIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Position_AllPosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllPositionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServer).AllPosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/position.Position/AllPosition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServer).AllPosition(ctx, req.(*AllPositionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Position_serviceDesc = grpc.ServiceDesc{
	ServiceName: "position.Position",
	HandlerType: (*PositionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Position",
			Handler:    _Position_Position_Handler,
		},
		{
			MethodName: "AllPosition",
			Handler:    _Position_AllPosition_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "position/position.proto",
}
