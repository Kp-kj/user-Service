package logic

import (
	"context"

	"user/internal/svc"
	"user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecordNoticeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRecordNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecordNoticeLogic {
	return &RecordNoticeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RecordNoticeLogic) RecordNotice(in *user.RecordNoticeRequest) (*user.RecordNoticeResponse, error) {
	// 根据userid, RecordNoticeCategory和RecordNoticeStatus，如果给出LastNoticeId则根据Limit和sort返回指定数据
	// 如果没有给出LastNoticeId则根据Limit和sort返回指定数据

	// 先判断是否存在 lastNoticeId
	if in.LastNoticeId == "" {
		// 没有 lastNoticeId
		// 根据userid、RecordNoticeCategory和RecordNoticeStatus，根据Limit和sort返回指定数据
		recordNoticeData, err := l.svcCtx.RecordNotice.FindOneByUserIdAndRecordNoticeCategoryAndRecordNoticeStatusOrderLimitAndSort(
			l.ctx, in.UserId,
			in.RecordNoticeCategory,
			in.RecordNoticeStatus,
			in.Limit,
			in.Sort) // 根据用户ID、通知记录类型和通知记录状态查询通知记录并根据时间排序
		if err != nil {
			return nil, err
		}

		// 将类型转换为 user.Notice 类型的切片
		noticeSlice := make([]*user.Notice, len(recordNoticeData))
		for i, recordNotice := range recordNoticeData {
			noticeSlice[i] = &user.Notice{
				UserId:               recordNotice.UserId,               // 用户ID
				SystemNoticeId:       recordNotice.SystemNoticeId,       // 系统通知ID
				UserNoticeId:         recordNotice.UserNoticeId,         // 通知ID
				RecordNoticeCategory: recordNotice.RecordNoticeCategory, // 通知记录类型 0-系统通知 1-用户通知
				RecordNoticeStatus:   recordNotice.RecordNoticeStatus,   // 通知记录状态  0-未读 1-已读
			}
		}

		return &user.RecordNoticeResponse{
			Notice: noticeSlice,
		}, nil

	} else {
		// 处理有 lastNoticeId 的情况
		recordNoticeData, err := l.svcCtx.RecordNotice.FindByUserIdAndRecordNoticeCategoryAndRecordNoticeStatusOrderLimitAndSortWithLastNoticeID(
			l.ctx, in.UserId,
			in.RecordNoticeCategory,
			in.RecordNoticeStatus,
			in.LastNoticeId,
			in.Limit,
			in.Sort,
		) // 根据用户ID、通知记录类型和通知记录状态查询通知记录并根据时间排序

		if err != nil {
			return nil, err
		}

		noticeSlice := make([]*user.Notice, len(recordNoticeData))
		for i, recordNotice := range recordNoticeData {
			noticeSlice[i] = &user.Notice{
				UserId:               recordNotice.UserId,               // 用户ID
				SystemNoticeId:       recordNotice.SystemNoticeId,       // 系统通知ID
				UserNoticeId:         recordNotice.UserNoticeId,         // 通知ID
				RecordNoticeCategory: recordNotice.RecordNoticeCategory, // 通知记录类型 0-系统通知 1-用户通知
				RecordNoticeStatus:   recordNotice.RecordNoticeStatus,   // 通知记录状态  0-未读 1-已读
			}
		}
		return &user.RecordNoticeResponse{
			Notice: noticeSlice,
		}, nil
	}
}
