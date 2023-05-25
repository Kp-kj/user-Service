package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"user/internal/svc"
	"user/user"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *user.Request) (*user.Response, error) {
	// 测试连接
	return &user.Response{
		Pong: string("pong"),
	}, nil
}
