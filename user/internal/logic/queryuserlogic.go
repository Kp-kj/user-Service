package logic

import (
	"context"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
	"user/internal/svc"
	"user/user"
)

type QueryUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserLogic {
	return &QueryUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryUser 查询用户资料
func (l *QueryUserLogic) QueryUser(in *user.QueryUserRequest) (*user.QueryUserResponse, error) {

	// 将 in.UserId 转为 int64 类型
	userId, err := strconv.ParseInt(in.UserId, 10, 64)
	if err != nil {
		return nil, err
	}

	userData, err := l.svcCtx.Profile.FindOneByUserId(l.ctx, userId)
	if err != nil {
		return nil, err
	}

	return &user.QueryUserResponse{
		TwitterName: userData.TwitterName,
		UserName:    userData.UserName,
		UserAvatar:  userData.Avatar,
		IsNew:       userData.IsNew,
	}, nil
}
