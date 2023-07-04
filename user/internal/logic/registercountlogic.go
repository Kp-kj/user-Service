package logic

import (
	"context"

	"user/internal/svc"
	"user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterCountLogic {
	return &RegisterCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// RegisterCount 注册人数
func (l *RegisterCountLogic) RegisterCount(in *user.Request) (*user.RegisterCountResponse, error) {
	count, err := l.svcCtx.UserModel.Count(l.ctx)
	if err != nil {
		return nil, err
	}
	return &user.RegisterCountResponse{
		Count: count,
	}, nil
}
