// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/iam/v2beta/rules.proto

package v2beta // import "github.com/chef/automate/components/automate-gateway/api/iam/v2beta"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import request "github.com/chef/automate/components/automate-gateway/api/iam/v2beta/request"
import response "github.com/chef/automate/components/automate-gateway/api/iam/v2beta/response"
import _ "github.com/chef/automate/components/automate-grpc/protoc-gen-policy/api"
import _ "github.com/chef/automate/components/automate-grpc/protoc-gen-policy/iam"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RulesClient is the client API for Rules service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RulesClient interface {
	CreateRule(ctx context.Context, in *request.CreateRuleReq, opts ...grpc.CallOption) (*response.CreateRuleResp, error)
	UpdateRule(ctx context.Context, in *request.UpdateRuleReq, opts ...grpc.CallOption) (*response.UpdateRuleResp, error)
	GetRule(ctx context.Context, in *request.GetRuleReq, opts ...grpc.CallOption) (*response.GetRuleResp, error)
	ListRules(ctx context.Context, in *request.ListRulesReq, opts ...grpc.CallOption) (*response.ListRulesResp, error)
	ListRulesForProject(ctx context.Context, in *request.ListRulesForProjectReq, opts ...grpc.CallOption) (*response.ListRulesForProjectResp, error)
	DeleteRule(ctx context.Context, in *request.DeleteRuleReq, opts ...grpc.CallOption) (*response.DeleteRuleResp, error)
}

type rulesClient struct {
	cc *grpc.ClientConn
}

func NewRulesClient(cc *grpc.ClientConn) RulesClient {
	return &rulesClient{cc}
}

func (c *rulesClient) CreateRule(ctx context.Context, in *request.CreateRuleReq, opts ...grpc.CallOption) (*response.CreateRuleResp, error) {
	out := new(response.CreateRuleResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2beta.Rules/CreateRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rulesClient) UpdateRule(ctx context.Context, in *request.UpdateRuleReq, opts ...grpc.CallOption) (*response.UpdateRuleResp, error) {
	out := new(response.UpdateRuleResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2beta.Rules/UpdateRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rulesClient) GetRule(ctx context.Context, in *request.GetRuleReq, opts ...grpc.CallOption) (*response.GetRuleResp, error) {
	out := new(response.GetRuleResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2beta.Rules/GetRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rulesClient) ListRules(ctx context.Context, in *request.ListRulesReq, opts ...grpc.CallOption) (*response.ListRulesResp, error) {
	out := new(response.ListRulesResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2beta.Rules/ListRules", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rulesClient) ListRulesForProject(ctx context.Context, in *request.ListRulesForProjectReq, opts ...grpc.CallOption) (*response.ListRulesForProjectResp, error) {
	out := new(response.ListRulesForProjectResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2beta.Rules/ListRulesForProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rulesClient) DeleteRule(ctx context.Context, in *request.DeleteRuleReq, opts ...grpc.CallOption) (*response.DeleteRuleResp, error) {
	out := new(response.DeleteRuleResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2beta.Rules/DeleteRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RulesServer is the server API for Rules service.
type RulesServer interface {
	CreateRule(context.Context, *request.CreateRuleReq) (*response.CreateRuleResp, error)
	UpdateRule(context.Context, *request.UpdateRuleReq) (*response.UpdateRuleResp, error)
	GetRule(context.Context, *request.GetRuleReq) (*response.GetRuleResp, error)
	ListRules(context.Context, *request.ListRulesReq) (*response.ListRulesResp, error)
	ListRulesForProject(context.Context, *request.ListRulesForProjectReq) (*response.ListRulesForProjectResp, error)
	DeleteRule(context.Context, *request.DeleteRuleReq) (*response.DeleteRuleResp, error)
}

func RegisterRulesServer(s *grpc.Server, srv RulesServer) {
	s.RegisterService(&_Rules_serviceDesc, srv)
}

func _Rules_CreateRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.CreateRuleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RulesServer).CreateRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2beta.Rules/CreateRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RulesServer).CreateRule(ctx, req.(*request.CreateRuleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rules_UpdateRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.UpdateRuleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RulesServer).UpdateRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2beta.Rules/UpdateRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RulesServer).UpdateRule(ctx, req.(*request.UpdateRuleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rules_GetRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetRuleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RulesServer).GetRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2beta.Rules/GetRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RulesServer).GetRule(ctx, req.(*request.GetRuleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rules_ListRules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.ListRulesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RulesServer).ListRules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2beta.Rules/ListRules",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RulesServer).ListRules(ctx, req.(*request.ListRulesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rules_ListRulesForProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.ListRulesForProjectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RulesServer).ListRulesForProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2beta.Rules/ListRulesForProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RulesServer).ListRulesForProject(ctx, req.(*request.ListRulesForProjectReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rules_DeleteRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.DeleteRuleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RulesServer).DeleteRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2beta.Rules/DeleteRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RulesServer).DeleteRule(ctx, req.(*request.DeleteRuleReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Rules_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.api.iam.v2beta.Rules",
	HandlerType: (*RulesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRule",
			Handler:    _Rules_CreateRule_Handler,
		},
		{
			MethodName: "UpdateRule",
			Handler:    _Rules_UpdateRule_Handler,
		},
		{
			MethodName: "GetRule",
			Handler:    _Rules_GetRule_Handler,
		},
		{
			MethodName: "ListRules",
			Handler:    _Rules_ListRules_Handler,
		},
		{
			MethodName: "ListRulesForProject",
			Handler:    _Rules_ListRulesForProject_Handler,
		},
		{
			MethodName: "DeleteRule",
			Handler:    _Rules_DeleteRule_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "components/automate-gateway/api/iam/v2beta/rules.proto",
}

func init() {
	proto.RegisterFile("components/automate-gateway/api/iam/v2beta/rules.proto", fileDescriptor_rules_d7803f9c364b5491)
}

var fileDescriptor_rules_d7803f9c364b5491 = []byte{
	// 559 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x3d, 0x8f, 0xd4, 0x3c,
	0x10, 0x80, 0x95, 0xbb, 0xf7, 0xf6, 0xe5, 0xcc, 0xc7, 0xed, 0x79, 0x25, 0x08, 0xd1, 0x49, 0x88,
	0x54, 0xb0, 0xc7, 0x26, 0xd2, 0xf1, 0x51, 0x6c, 0x01, 0x88, 0xe3, 0xa3, 0xa1, 0x40, 0x27, 0x41,
	0x41, 0xe7, 0xcd, 0x0e, 0x59, 0xa3, 0x4d, 0xec, 0x8b, 0x1d, 0xd0, 0x09, 0x21, 0xa1, 0x48, 0x34,
	0xdb, 0x22, 0xf1, 0x4f, 0xd2, 0xd1, 0xd1, 0x52, 0x41, 0x49, 0xcb, 0x0f, 0x41, 0xb1, 0xf7, 0x92,
	0x40, 0x92, 0x95, 0x97, 0x32, 0x99, 0x67, 0xc6, 0xf3, 0xd8, 0x63, 0xa3, 0x3b, 0x01, 0x8b, 0x38,
	0x8b, 0x21, 0x96, 0xc2, 0x27, 0xa9, 0x64, 0x11, 0x91, 0x30, 0x0a, 0x89, 0x84, 0xb7, 0xe4, 0xc4,
	0x27, 0x9c, 0xfa, 0x94, 0x44, 0xfe, 0x9b, 0x83, 0x09, 0x48, 0xe2, 0x27, 0xe9, 0x1c, 0x84, 0xc7,
	0x13, 0x26, 0x19, 0xde, 0x0b, 0x66, 0xf0, 0xca, 0x3b, 0xcd, 0xf0, 0x08, 0xa7, 0x1e, 0x25, 0x91,
	0xa7, 0x49, 0x67, 0x2f, 0x64, 0x2c, 0x9c, 0x83, 0x2a, 0x40, 0xe2, 0x98, 0x49, 0x22, 0x29, 0x8b,
	0x97, 0xb9, 0xce, 0xdd, 0x75, 0xd6, 0x84, 0xe3, 0x14, 0x84, 0xac, 0xaf, 0xed, 0xdc, 0x5b, 0x2b,
	0x5f, 0x70, 0x16, 0x0b, 0xf8, 0xa3, 0xc0, 0xfd, 0xd6, 0x02, 0x09, 0x0f, 0x7c, 0x15, 0x0f, 0x46,
	0x21, 0xc4, 0x23, 0xce, 0xe6, 0x34, 0x38, 0xe9, 0x50, 0x58, 0xa7, 0x42, 0xd1, 0x4d, 0xa3, 0xc2,
	0xc1, 0x67, 0x84, 0xb6, 0x8e, 0x8a, 0x9e, 0xf0, 0x37, 0x0b, 0xa1, 0xc3, 0x04, 0x88, 0x84, 0xe2,
	0x1b, 0xef, 0x7b, 0xab, 0xb6, 0xd6, 0xab, 0xc8, 0x23, 0x38, 0x76, 0x6e, 0x98, 0xc3, 0x82, 0xbb,
	0x41, 0x96, 0xdb, 0xe7, 0x10, 0x22, 0xa9, 0x9c, 0x8d, 0xd5, 0x7e, 0x64, 0xb9, 0x7d, 0x06, 0xf7,
	0x02, 0x45, 0x2d, 0x72, 0xfb, 0x2c, 0xda, 0xa6, 0x24, 0xd2, 0xa1, 0x45, 0x6e, 0x63, 0xdc, 0x2f,
	0x3f, 0xc7, 0x1a, 0xca, 0x7e, 0xfc, 0xfa, 0xb4, 0x31, 0x70, 0x77, 0x1b, 0x63, 0xa1, 0x02, 0x9b,
	0x63, 0x6b, 0x88, 0x7f, 0x5a, 0x08, 0x3d, 0xe7, 0x53, 0x43, 0x9d, 0x8a, 0x34, 0xd0, 0xa9, 0xc3,
	0x82, 0xbb, 0x49, 0x96, 0xdb, 0xbb, 0x68, 0xa7, 0xd2, 0x19, 0xbf, 0xa3, 0xd3, 0xf7, 0xda, 0x29,
	0x55, 0xe8, 0x22, 0xb7, 0xfb, 0xe8, 0x42, 0x25, 0x51, 0xc4, 0xff, 0x16, 0xd3, 0xa4, 0xea, 0xdf,
	0x76, 0x2e, 0x36, 0xc4, 0x7c, 0x55, 0xb3, 0xb4, 0xfb, 0x6a, 0xa1, 0xff, 0x9f, 0x80, 0x54, 0x6a,
	0xd7, 0x56, 0x77, 0xbb, 0xc4, 0x0a, 0xaf, 0xeb, 0x86, 0xe4, 0xf2, 0x8c, 0x5a, 0xa5, 0xb6, 0xf0,
	0x66, 0x08, 0xb2, 0xc3, 0x68, 0x07, 0x9f, 0xaf, 0xfe, 0x85, 0x20, 0xb5, 0x0e, 0xee, 0xd0, 0xc1,
	0x5f, 0x2c, 0xb4, 0xfd, 0x94, 0x0a, 0xa9, 0x07, 0x70, 0xb8, 0xba, 0xbb, 0x12, 0x2c, 0x4c, 0xf6,
	0x8d, 0x59, 0xc1, 0xdd, 0x17, 0x2d, 0xf3, 0xd6, 0xc3, 0xff, 0x25, 0x40, 0xa6, 0xcd, 0x69, 0xeb,
	0xe3, 0x9a, 0xd6, 0x9c, 0x0a, 0xed, 0x30, 0xc0, 0xcd, 0x59, 0xc3, 0x1f, 0x37, 0xd0, 0xa0, 0x5c,
	0xe9, 0x31, 0x4b, 0x9e, 0x25, 0xec, 0x35, 0x04, 0x12, 0xdf, 0x32, 0x6c, 0xae, 0x4a, 0x29, 0x94,
	0x6e, 0xff, 0x43, 0x96, 0xe0, 0xee, 0x07, 0x2b, 0xcb, 0x6d, 0x07, 0xd9, 0xca, 0x8e, 0xeb, 0x80,
	0x3e, 0x8f, 0x16, 0xd7, 0xcb, 0xe8, 0x52, 0x21, 0xd7, 0x82, 0x75, 0x9a, 0x5f, 0xc5, 0x57, 0xea,
	0xe6, 0xa7, 0xa9, 0xea, 0x00, 0x97, 0xfb, 0xf0, 0xdd, 0x42, 0xe8, 0x21, 0xcc, 0xc1, 0xec, 0xaa,
	0x55, 0xa4, 0xc1, 0x55, 0xab, 0xc3, 0x82, 0xbb, 0x74, 0xc5, 0x55, 0x9b, 0x2a, 0xd4, 0xe4, 0xaa,
	0x69, 0x52, 0xcf, 0xe6, 0xb0, 0x63, 0x36, 0x1f, 0x3c, 0x7a, 0x79, 0x18, 0x52, 0x39, 0x4b, 0x27,
	0x5e, 0xc0, 0x22, 0xbf, 0x68, 0xb2, 0x7c, 0x61, 0x7d, 0xf3, 0x87, 0x7f, 0xd2, 0x53, 0xcf, 0xec,
	0xcd, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf7, 0xc0, 0xa3, 0x0a, 0xe1, 0x06, 0x00, 0x00,
}
