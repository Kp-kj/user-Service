package logic

import (
	"context"
	"fmt"
	"user/internal/svc"
	"user/user"

	"github.com/zeromicro/go-zero/core/logx"
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

// Ping ...
func (l *PingLogic) Ping(in *user.Request) (*user.Response, error) {

	fmt.Println(in.Ping)

	fmt.Println("pong")
	return &user.Response{
		Pong: "pong",
	}, nil
}
