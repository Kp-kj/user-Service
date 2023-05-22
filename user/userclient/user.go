// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userclient

import (
	"context"

	"user/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CheckTwitterIdRequest  = user.CheckTwitterIdRequest
	CheckTwitterIdResponse = user.CheckTwitterIdResponse
	CreateUserRequest      = user.CreateUserRequest
	CreateUserResponse     = user.CreateUserResponse
	ErrorResponse          = user.ErrorResponse
	Request                = user.Request
	Response               = user.Response
	ResponseMessage        = user.ResponseMessage

	User interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
		CheckTwitterId(ctx context.Context, in *CheckTwitterIdRequest, opts ...grpc.CallOption) (*CheckTwitterIdResponse, error)
		CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}

func (m *defaultUser) CheckTwitterId(ctx context.Context, in *CheckTwitterIdRequest, opts ...grpc.CallOption) (*CheckTwitterIdResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.CheckTwitterId(ctx, in, opts...)
}

func (m *defaultUser) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.CreateUser(ctx, in, opts...)
}
