package logic

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"user/internal/model"

	"user/internal/svc"
	"user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateNotificationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateNotificationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateNotificationLogic {
	return &CreateNotificationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 新增用户消息通知
func (l *CreateNotificationLogic) CreateNotification(in *user.CreateNotificationRequest) (*user.CreateNotificationResponse, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	userNotice := &model.UserNotice{
		UserNoticeId:  node.Generate().Int64(), // 用户通知ID
		UserId:        in.UserId,               // 用户ID
		NoticeContent: in.NoticeContent,        // 通知内容
		NoticeStatus:  1,                       // 通知状态 1 已读 2 下架
	}

	_, err = l.svcCtx.UserNotice.Insert(l.ctx, userNotice)
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	return &user.CreateNotificationResponse{
		UserNoticeId: userNotice.UserId,
		IsSuccess:    true,
	}, nil
}
