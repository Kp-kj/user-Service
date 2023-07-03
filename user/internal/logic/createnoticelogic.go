package logic

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"user/internal/model"

	"user/internal/svc"
	"user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateNoticeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateNoticeLogic {
	return &CreateNoticeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateNotice  创建通知记录
func (l *CreateNoticeLogic) CreateNotice(in *user.CreateNoticeRequest) (*user.CreateNoticeResponse, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	recordNotice := &model.RecordNotice{
		RecordNoticeId:       node.Generate().Int64(), // 通知记录ID
		UserId:               in.UserId,               // 用户ID
		SystemNoticeId:       in.SystemNoticeId,       // 系统通知ID
		UserNoticeId:         in.UserNoticeId,         // 用户通知ID
		RecordNoticeCategory: in.RecordNoticeCategory, // 通知记录类型 0-系统通知 1-用户通知
		RecordNoticeStatus:   in.RecordNoticeStatus,   // 通知记录状态  0-未读 1-已读
	}

	_, err = l.svcCtx.RecordNotice.Insert(l.ctx, recordNotice)
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	return &user.CreateNoticeResponse{
		IsSuccess: true,
	}, nil
}
