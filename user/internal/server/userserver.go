// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"user/internal/logic"
	"user/internal/svc"
	"user/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) Ping(ctx context.Context, in *user.Request) (*user.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}

func (s *UserServer) CheckTwitterId(ctx context.Context, in *user.CheckTwitterIdRequest) (*user.CheckTwitterIdResponse, error) {
	l := logic.NewCheckTwitterIdLogic(ctx, s.svcCtx)
	return l.CheckTwitterId(in)
}

func (s *UserServer) CreateUser(ctx context.Context, in *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	l := logic.NewCreateUserLogic(ctx, s.svcCtx)
	return l.CreateUser(in)
}

func (s *UserServer) CreateInvite(ctx context.Context, in *user.CreateInviteRequest) (*user.CreateInviteResponse, error) {
	l := logic.NewCreateInviteLogic(ctx, s.svcCtx)
	return l.CreateInvite(in)
}

func (s *UserServer) CheckTodayInvite(ctx context.Context, in *user.CheckTodayInviteRequest) (*user.CheckTodayInviteResponse, error) {
	l := logic.NewCheckTodayInviteLogic(ctx, s.svcCtx)
	return l.CheckTodayInvite(in)
}

func (s *UserServer) AddUserInfo(ctx context.Context, in *user.AddUserInfoRequest) (*user.AddUserInfoResponse, error) {
	l := logic.NewAddUserInfoLogic(ctx, s.svcCtx)
	return l.AddUserInfo(in)
}

func (s *UserServer) QueryUser(ctx context.Context, in *user.QueryUserRequest) (*user.QueryUserResponse, error) {
	l := logic.NewQueryUserLogic(ctx, s.svcCtx)
	return l.QueryUser(in)
}

func (s *UserServer) AddAdmin(ctx context.Context, in *user.AddAdminRequest) (*user.AddAdminResponse, error) {
	l := logic.NewAddAdminLogic(ctx, s.svcCtx)
	return l.AddAdmin(in)
}

func (s *UserServer) AdminLogin(ctx context.Context, in *user.AdminLoginRequest) (*user.AdminLoginResponse, error) {
	l := logic.NewAdminLoginLogic(ctx, s.svcCtx)
	return l.AdminLogin(in)
}

func (s *UserServer) RemoveAdmin(ctx context.Context, in *user.RemoveAdminRequest) (*user.RemoveAdminResponse, error) {
	l := logic.NewRemoveAdminLogic(ctx, s.svcCtx)
	return l.RemoveAdmin(in)
}
