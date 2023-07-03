package logic

import (
	"context"
	"user/internal/model"

	"user/internal/svc"
	"user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditSystemNotificationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEditSystemNotificationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditSystemNotificationLogic {
	return &EditSystemNotificationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 编辑系统消息通知
func (l *EditSystemNotificationLogic) EditSystemNotification(in *user.EditSystemNotificationRequest) (*user.EditSystemNotificationResponse, error) {
	// 修改帮助分类上下架状态
	dbSystemNotification := &model.SystemNotice{
		NoticeTitle:    in.NoticeTitle,    // 消息标题
		NoticeContent:  in.NoticeContent,  // 消息内容
		NoticeCategory: in.NoticeCategory, // 消息分类
		NoticeStatus:   in.NoticeStatus,   // 消息状态
		SystemNoticeId: in.SystemNoticeId, // 消息id
	}

	err := l.svcCtx.SystemNotice.Update(l.ctx, dbSystemNotification)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	return &user.EditSystemNotificationResponse{
		IsSuccess: true,
	}, nil
}

//1675749205508886528
