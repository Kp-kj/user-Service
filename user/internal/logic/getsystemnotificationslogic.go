package logic

import (
	"context"
	"user/internal/svc"
	"user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSystemNotificationsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSystemNotificationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSystemNotificationsLogic {
	return &GetSystemNotificationsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取系统消息通知列表
func (l *GetSystemNotificationsLogic) GetSystemNotifications(in *user.GetSystemNotificationsRequest) (*user.GetSystemNotificationsResponse, error) {
	dbSystemNotifications, err := l.svcCtx.SystemNotice.FindOneByNoticeCategoryNoticeStatus(l.ctx, in.NoticeCategory, in.NoticeStatus)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	response := &user.GetSystemNotificationsResponse{
		SystemNotifications: make([]*user.SystemNotification, len(dbSystemNotifications)),
	}
	for _, notification := range dbSystemNotifications {
		noticeStartTime := int64(0) // 默认值，当时间不存在时使用
		if notification.NoticeStartTime.Valid {
			noticeStartTime = notification.NoticeStartTime.Time.Unix()
		}

		systemNotification := &user.SystemNotification{
			SystemNoticeId:  notification.SystemNoticeId, // 通知id
			NoticeTitle:     notification.NoticeTitle,    // 通知标题
			NoticeContent:   notification.NoticeContent,  // 通知内容
			NoticeStatus:    notification.NoticeStatus,   // 通知状态
			NoticeCategory:  notification.NoticeCategory, // 通知类别
			NoticeStartTime: noticeStartTime,             // 通知开始时间
		}
		response.SystemNotifications = append(response.SystemNotifications, systemNotification)
	}

	return response, nil
}
