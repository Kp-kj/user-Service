package logic

import (
	"context"

	"user/internal/svc"
	"user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryBlackListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryBlackListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryBlackListLogic {
	return &QueryBlackListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryBlackList 查询黑名单
func (l *QueryBlackListLogic) QueryBlackList(in *user.QueryBlackListRequest) (*user.QueryBlackListResponse, error) {
	// 查询黑名单ByUserId
	dbBlack, err := l.svcCtx.Black.FindOneByBlackId(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}

	return &user.QueryBlackListResponse{
		BlackUserId: dbBlack.BlackUserId,
	}, nil
}
