package logic

import (
	"context"

	"user/internal/svc"
	"user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveAdminLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveAdminLogic {
	return &RemoveAdminLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// RemoveAdmin 删除管理员
func (l *RemoveAdminLogic) RemoveAdmin(in *user.RemoveAdminRequest) (*user.RemoveAdminResponse, error) {
	err := l.svcCtx.AdminUser.Delete(l.ctx, in.AdminUserId)
	if err != nil {
		logx.Error()
		return nil, err
	}
	return &user.RemoveAdminResponse{
		IsSuccess: true,
	}, nil
}
