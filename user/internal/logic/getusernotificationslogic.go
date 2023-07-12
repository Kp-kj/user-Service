package logic

import (
	"context"

	"user/internal/svc"
	"user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserNotificationsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserNotificationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserNotificationsLogic {
	return &GetUserNotificationsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户消息通知
func (l *GetUserNotificationsLogic) GetUserNotifications(in *user.GetUserNotificationsRequest) (*user.GetUserNotificationsResponse, error) {
	userNotifications, err := l.svcCtx.UserNotice.FindOne(l.ctx, in.UserNoticeId)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	return &user.GetUserNotificationsResponse{
		NoticeContent: userNotifications.NoticeContent,
		NoticeStatus:  userNotifications.NoticeStatus,
		UserId:        userNotifications.UserId,
	}, nil
}
