// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.20.3
// source: quizzes/quiz_service.proto

package quizzes

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Quizzes_Update_FullMethodName = "/quizzes.Quizzes/Update"
	Quizzes_Answer_FullMethodName = "/quizzes.Quizzes/Answer"
	Quizzes_Get_FullMethodName    = "/quizzes.Quizzes/Get"
)

// QuizzesClient is the client API for Quizzes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QuizzesClient interface {
	Update(ctx context.Context, in *QuizUpdateInput, opts ...grpc.CallOption) (*Quiz, error)
	Answer(ctx context.Context, in *QuizAnswerInput, opts ...grpc.CallOption) (*QuizAnswer, error)
	Get(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Quiz, error)
}

type quizzesClient struct {
	cc grpc.ClientConnInterface
}

func NewQuizzesClient(cc grpc.ClientConnInterface) QuizzesClient {
	return &quizzesClient{cc}
}

func (c *quizzesClient) Update(ctx context.Context, in *QuizUpdateInput, opts ...grpc.CallOption) (*Quiz, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Quiz)
	err := c.cc.Invoke(ctx, Quizzes_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quizzesClient) Answer(ctx context.Context, in *QuizAnswerInput, opts ...grpc.CallOption) (*QuizAnswer, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QuizAnswer)
	err := c.cc.Invoke(ctx, Quizzes_Answer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quizzesClient) Get(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Quiz, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Quiz)
	err := c.cc.Invoke(ctx, Quizzes_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuizzesServer is the server API for Quizzes service.
// All implementations should embed UnimplementedQuizzesServer
// for forward compatibility
type QuizzesServer interface {
	Update(context.Context, *QuizUpdateInput) (*Quiz, error)
	Answer(context.Context, *QuizAnswerInput) (*QuizAnswer, error)
	Get(context.Context, *Id) (*Quiz, error)
}

// UnimplementedQuizzesServer should be embedded to have forward compatible implementations.
type UnimplementedQuizzesServer struct {
}

func (UnimplementedQuizzesServer) Update(context.Context, *QuizUpdateInput) (*Quiz, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedQuizzesServer) Answer(context.Context, *QuizAnswerInput) (*QuizAnswer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Answer not implemented")
}
func (UnimplementedQuizzesServer) Get(context.Context, *Id) (*Quiz, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

// UnsafeQuizzesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QuizzesServer will
// result in compilation errors.
type UnsafeQuizzesServer interface {
	mustEmbedUnimplementedQuizzesServer()
}

func RegisterQuizzesServer(s grpc.ServiceRegistrar, srv QuizzesServer) {
	s.RegisterService(&Quizzes_ServiceDesc, srv)
}

func _Quizzes_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuizUpdateInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuizzesServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Quizzes_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuizzesServer).Update(ctx, req.(*QuizUpdateInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Quizzes_Answer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuizAnswerInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuizzesServer).Answer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Quizzes_Answer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuizzesServer).Answer(ctx, req.(*QuizAnswerInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Quizzes_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuizzesServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Quizzes_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuizzesServer).Get(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

// Quizzes_ServiceDesc is the grpc.ServiceDesc for Quizzes service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Quizzes_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "quizzes.Quizzes",
	HandlerType: (*QuizzesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Update",
			Handler:    _Quizzes_Update_Handler,
		},
		{
			MethodName: "Answer",
			Handler:    _Quizzes_Answer_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Quizzes_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "quizzes/quiz_service.proto",
}
