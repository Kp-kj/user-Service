package logic

import (
	"context"
	"database/sql"
	"github.com/bwmarrin/snowflake"
	"time"
	"user/internal/model"

	"user/internal/svc"
	"user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSystemNotificationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateSystemNotificationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSystemNotificationLogic {
	return &CreateSystemNotificationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 新增系统消息通知
func (l *CreateSystemNotificationLogic) CreateSystemNotification(in *user.CreateSystemNotificationRequest) (*user.CreateSystemNotificationResponse, error) {
	node, err := snowflake.NewNode(1) // 生成雪花ID
	if err != nil {
		return nil, err
	}

	var noticeStartTime sql.NullTime
	if in.NoticeStartTime != 0 {
		noticeStartTime = sql.NullTime{
			Time:  time.Unix(in.NoticeStartTime, 0),
			Valid: true,
		}
	} else {
		noticeStartTime = sql.NullTime{
			Time:  time.Time{},
			Valid: false,
		}
	}
	SystemNoticeId := node.Generate().Int64()
	_, err = l.svcCtx.SystemNotice.Insert(l.ctx, &model.SystemNotice{ // 插入数据
		SystemNoticeId:  SystemNoticeId,
		NoticeTitle:     in.NoticeTitle,    // 通知标题
		NoticeCategory:  in.NoticeCategory, // 通知分类
		NoticeContent:   in.NoticeContent,  // 通知内容
		NoticeStatus:    in.NoticeStatus,   // 通知状态
		NoticeStartTime: noticeStartTime,   // 通知开始时间
	})
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	return &user.CreateSystemNotificationResponse{
		SystemNoticeId: SystemNoticeId,
		IsSuccess:      true,
	}, nil
}
