package logic

import (
	"context"

	"user/internal/svc"
	"user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type OnlineCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOnlineCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OnlineCountLogic {
	return &OnlineCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// OnlineCount 在线人数
func (l *OnlineCountLogic) OnlineCount(in *user.Request) (*user.OnlineCountResponse, error) {
	// todo: add your logic here and delete this line
	// 使用缓存查询在线人数

	return &user.OnlineCountResponse{}, nil
}
