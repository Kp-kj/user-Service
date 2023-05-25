package logic

import (
	"context"
	"strconv"
	"user/user"

	"user/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckTodayInviteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckTodayInviteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckTodayInviteLogic {
	return &CheckTodayInviteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckTodayInviteLogic) CheckTodayInvite(in *user.CheckTodayInviteRequest) (*user.CheckTodayInviteResponse, error) {
	userId, err := strconv.ParseInt(in.UserId, 10, 64)
	if err != nil {
		return nil, err
	}
	isCheckTodayInvite, err := l.svcCtx.Invitation.CheckTodayInvite(l.ctx, userId)
	if err != nil {
		return nil, err
	}
	return &user.CheckTodayInviteResponse{
		IsInvite: isCheckTodayInvite,
	}, nil
}
