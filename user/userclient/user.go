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
	AddAdminRequest          = user.AddAdminRequest
	AddAdminResponse         = user.AddAdminResponse
	AddBlackListRequest      = user.AddBlackListRequest
	AddBlackListResponse     = user.AddBlackListResponse
	AddUserInfoRequest       = user.AddUserInfoRequest
	AddUserInfoResponse      = user.AddUserInfoResponse
	AdminLoginRequest        = user.AdminLoginRequest
	AdminLoginResponse       = user.AdminLoginResponse
	CheckTodayInviteRequest  = user.CheckTodayInviteRequest
	CheckTodayInviteResponse = user.CheckTodayInviteResponse
	CheckTwitterIdRequest    = user.CheckTwitterIdRequest
	CheckTwitterIdResponse   = user.CheckTwitterIdResponse
	CreateInviteRequest      = user.CreateInviteRequest
	CreateInviteResponse     = user.CreateInviteResponse
	CreateUserRequest        = user.CreateUserRequest
	CreateUserResponse       = user.CreateUserResponse
	QueryBlackListRequest    = user.QueryBlackListRequest
	QueryBlackListResponse   = user.QueryBlackListResponse
	QueryUserRequest         = user.QueryUserRequest
	QueryUserResponse        = user.QueryUserResponse
	RemoveAdminRequest       = user.RemoveAdminRequest
	RemoveAdminResponse      = user.RemoveAdminResponse
	RemoveBlackListRequest   = user.RemoveBlackListRequest
	RemoveBlackListResponse  = user.RemoveBlackListResponse
	Request                  = user.Request
	Response                 = user.Response

	User interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
		CheckTwitterId(ctx context.Context, in *CheckTwitterIdRequest, opts ...grpc.CallOption) (*CheckTwitterIdResponse, error)
		CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
		CreateInvite(ctx context.Context, in *CreateInviteRequest, opts ...grpc.CallOption) (*CreateInviteResponse, error)
		CheckTodayInvite(ctx context.Context, in *CheckTodayInviteRequest, opts ...grpc.CallOption) (*CheckTodayInviteResponse, error)
		AddUserInfo(ctx context.Context, in *AddUserInfoRequest, opts ...grpc.CallOption) (*AddUserInfoResponse, error)
		QueryUser(ctx context.Context, in *QueryUserRequest, opts ...grpc.CallOption) (*QueryUserResponse, error)
		AddAdmin(ctx context.Context, in *AddAdminRequest, opts ...grpc.CallOption) (*AddAdminResponse, error)
		AdminLogin(ctx context.Context, in *AdminLoginRequest, opts ...grpc.CallOption) (*AdminLoginResponse, error)
		RemoveAdmin(ctx context.Context, in *RemoveAdminRequest, opts ...grpc.CallOption) (*RemoveAdminResponse, error)
		AddBlackList(ctx context.Context, in *AddBlackListRequest, opts ...grpc.CallOption) (*AddBlackListResponse, error)
		QueryBlackList(ctx context.Context, in *QueryBlackListRequest, opts ...grpc.CallOption) (*QueryBlackListResponse, error)
		RemoveBlackList(ctx context.Context, in *RemoveBlackListRequest, opts ...grpc.CallOption) (*RemoveBlackListResponse, error)
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

func (m *defaultUser) CreateInvite(ctx context.Context, in *CreateInviteRequest, opts ...grpc.CallOption) (*CreateInviteResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.CreateInvite(ctx, in, opts...)
}

func (m *defaultUser) CheckTodayInvite(ctx context.Context, in *CheckTodayInviteRequest, opts ...grpc.CallOption) (*CheckTodayInviteResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.CheckTodayInvite(ctx, in, opts...)
}

func (m *defaultUser) AddUserInfo(ctx context.Context, in *AddUserInfoRequest, opts ...grpc.CallOption) (*AddUserInfoResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.AddUserInfo(ctx, in, opts...)
}

func (m *defaultUser) QueryUser(ctx context.Context, in *QueryUserRequest, opts ...grpc.CallOption) (*QueryUserResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.QueryUser(ctx, in, opts...)
}

func (m *defaultUser) AddAdmin(ctx context.Context, in *AddAdminRequest, opts ...grpc.CallOption) (*AddAdminResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.AddAdmin(ctx, in, opts...)
}

func (m *defaultUser) AdminLogin(ctx context.Context, in *AdminLoginRequest, opts ...grpc.CallOption) (*AdminLoginResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.AdminLogin(ctx, in, opts...)
}

func (m *defaultUser) RemoveAdmin(ctx context.Context, in *RemoveAdminRequest, opts ...grpc.CallOption) (*RemoveAdminResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.RemoveAdmin(ctx, in, opts...)
}

func (m *defaultUser) AddBlackList(ctx context.Context, in *AddBlackListRequest, opts ...grpc.CallOption) (*AddBlackListResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.AddBlackList(ctx, in, opts...)
}

func (m *defaultUser) QueryBlackList(ctx context.Context, in *QueryBlackListRequest, opts ...grpc.CallOption) (*QueryBlackListResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.QueryBlackList(ctx, in, opts...)
}

func (m *defaultUser) RemoveBlackList(ctx context.Context, in *RemoveBlackListRequest, opts ...grpc.CallOption) (*RemoveBlackListResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.RemoveBlackList(ctx, in, opts...)
}
