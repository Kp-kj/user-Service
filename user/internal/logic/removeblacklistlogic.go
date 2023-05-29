package logic

import (
	"context"

	"user/internal/svc"
	"user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveBlackListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveBlackListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveBlackListLogic {
	return &RemoveBlackListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// RemoveBlackList 删除黑名单
func (l *RemoveBlackListLogic) RemoveBlackList(in *user.RemoveBlackListRequest) (*user.RemoveBlackListResponse, error) {
	// 不用判断是否存在，Api层用协程并发调用，不存在打印log返回报错
	err := l.svcCtx.Black.Delete(l.ctx, in.UserId)
	if err != nil {
		logx.Error(err)
		return &user.RemoveBlackListResponse{IsSuccess: false}, err
	}

	return &user.RemoveBlackListResponse{
		IsSuccess: true,
	}, nil
}
