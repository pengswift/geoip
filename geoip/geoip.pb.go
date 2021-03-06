// Code generated by protoc-gen-go.
// source: geoip.proto
// DO NOT EDIT!

/*
Package geoip is a generated protocol buffer package.

It is generated from these files:
	geoip.proto

It has these top-level messages:
	GeoIPRequest
	GeoIPResponse
*/
package geoip

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

type GeoIPRequest struct {
	Ip string `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
}

func (m *GeoIPRequest) Reset()         { *m = GeoIPRequest{} }
func (m *GeoIPRequest) String() string { return proto.CompactTextString(m) }
func (*GeoIPRequest) ProtoMessage()    {}

type GeoIPResponse struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GeoIPResponse) Reset()         { *m = GeoIPResponse{} }
func (m *GeoIPResponse) String() string { return proto.CompactTextString(m) }
func (*GeoIPResponse) ProtoMessage()    {}

func init() {
	proto.RegisterType((*GeoIPRequest)(nil), "geoip.GeoIPRequest")
	proto.RegisterType((*GeoIPResponse)(nil), "geoip.GeoIPResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for GeoIPService service

type GeoIPServiceClient interface {
	QueryCountry(ctx context.Context, in *GeoIPRequest, opts ...grpc.CallOption) (*GeoIPResponse, error)
	QuerySubdivision(ctx context.Context, in *GeoIPRequest, opts ...grpc.CallOption) (*GeoIPResponse, error)
	QueryCity(ctx context.Context, in *GeoIPRequest, opts ...grpc.CallOption) (*GeoIPResponse, error)
}

type geoIPServiceClient struct {
	cc *grpc.ClientConn
}

func NewGeoIPServiceClient(cc *grpc.ClientConn) GeoIPServiceClient {
	return &geoIPServiceClient{cc}
}

func (c *geoIPServiceClient) QueryCountry(ctx context.Context, in *GeoIPRequest, opts ...grpc.CallOption) (*GeoIPResponse, error) {
	out := new(GeoIPResponse)
	err := grpc.Invoke(ctx, "/geoip.GeoIPService/QueryCountry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoIPServiceClient) QuerySubdivision(ctx context.Context, in *GeoIPRequest, opts ...grpc.CallOption) (*GeoIPResponse, error) {
	out := new(GeoIPResponse)
	err := grpc.Invoke(ctx, "/geoip.GeoIPService/QuerySubdivision", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoIPServiceClient) QueryCity(ctx context.Context, in *GeoIPRequest, opts ...grpc.CallOption) (*GeoIPResponse, error) {
	out := new(GeoIPResponse)
	err := grpc.Invoke(ctx, "/geoip.GeoIPService/QueryCity", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GeoIPService service

type GeoIPServiceServer interface {
	QueryCountry(context.Context, *GeoIPRequest) (*GeoIPResponse, error)
	QuerySubdivision(context.Context, *GeoIPRequest) (*GeoIPResponse, error)
	QueryCity(context.Context, *GeoIPRequest) (*GeoIPResponse, error)
}

func RegisterGeoIPServiceServer(s *grpc.Server, srv GeoIPServiceServer) {
	s.RegisterService(&_GeoIPService_serviceDesc, srv)
}

func _GeoIPService_QueryCountry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GeoIPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(GeoIPServiceServer).QueryCountry(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _GeoIPService_QuerySubdivision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GeoIPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(GeoIPServiceServer).QuerySubdivision(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _GeoIPService_QueryCity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GeoIPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(GeoIPServiceServer).QueryCity(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _GeoIPService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "geoip.GeoIPService",
	HandlerType: (*GeoIPServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryCountry",
			Handler:    _GeoIPService_QueryCountry_Handler,
		},
		{
			MethodName: "QuerySubdivision",
			Handler:    _GeoIPService_QuerySubdivision_Handler,
		},
		{
			MethodName: "QueryCity",
			Handler:    _GeoIPService_QueryCity_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
