package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"user/internal/svc"
	"user/user"
)

type CheckTwitterIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckTwitterIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckTwitterIdLogic {
	return &CheckTwitterIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckTwitterIdLogic) CheckTwitterId(in *user.CheckTwitterIdRequest) (*user.CheckTwitterIdResponse, error) {
	userData, err := l.svcCtx.UserModel.FindOneByTwitterId(l.ctx, in.TwitterId)
	if err != nil {
		// 处理错误，例如记录日志或返回特定的错误响应
		fmt.Println(err)
		return nil, err
	}

	return &user.CheckTwitterIdResponse{
		UserId: string(userData.UserId),
	}, nil
}
