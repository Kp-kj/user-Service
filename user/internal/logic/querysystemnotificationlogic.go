package logic

import (
	"context"

	"user/internal/svc"
	"user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuerySystemNotificationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQuerySystemNotificationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySystemNotificationLogic {
	return &QuerySystemNotificationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询系统消息通知
func (l *QuerySystemNotificationLogic) QuerySystemNotification(in *user.QuerySystemNotificationRequest) (*user.QuerySystemNotificationResponse, error) {
	systemNotification, err := l.svcCtx.SystemNotice.FindOne(l.ctx, in.SystemNoticeId)
	if err != nil {
		return nil, err
	}
	noticeStartTime := int64(0) // 默认值，当时间不存在时使用
	if systemNotification.NoticeStartTime.Valid {
		noticeStartTime = systemNotification.NoticeStartTime.Time.Unix()
	}
	return &user.QuerySystemNotificationResponse{
		NoticeTitle:     systemNotification.NoticeTitle,
		NoticeStartTime: noticeStartTime,
		NoticeContent:   systemNotification.NoticeContent,
		NoticeStatus:    systemNotification.NoticeStatus,
		NoticeCategory:  systemNotification.NoticeCategory,
		SystemNoticeId:  systemNotification.SystemNoticeId,
	}, nil
}
