package logic

import (
	"context"

	"user/internal/svc"
	"user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetUserList 返回用户列表
func (l *GetUserListLogic) GetUserList(in *user.Request) (*user.UserListResponse, error) {
	// todo: add your logic here and delete this line

	userList, err := l.svcCtx.UserModel.FindAll(l.ctx)

	if err != nil {
		return nil, err
	}
	response := &user.UserListResponse{
		UserResponseLists: make([]*user.UserResponseList, len(userList)),
	}
	return response, nil
}
