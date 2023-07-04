package logic

import (
	"context"

	"user/internal/svc"
	"user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryRecordNoticeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryRecordNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryRecordNoticeLogic {
	return &QueryRecordNoticeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryRecordNotice 查询消息记录返回消息类型 系统消息/用户消息
func (l *QueryRecordNoticeLogic) QueryRecordNotice(in *user.QueryRecordNoticeRequest) (*user.QueryRecordNoticeResponse, error) {
	dbNotice, err := l.svcCtx.RecordNotice.FindOneBySystemNoticeIdOrUserNoticeId(l.ctx, in.NoticeId)
	if err != nil {
		return nil, err
	}

	return &user.QueryRecordNoticeResponse{
		RecordNoticeCategory: dbNotice.RecordNoticeCategory,
	}, nil
}
