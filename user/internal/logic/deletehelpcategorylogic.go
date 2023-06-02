package logic

import (
	"context"

	"user/internal/svc"
	"user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteHelpCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteHelpCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteHelpCategoryLogic {
	return &DeleteHelpCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteHelpCategory 删除帮助分类
func (l *DeleteHelpCategoryLogic) DeleteHelpCategory(in *user.DeleteHelpCategoryRequest) (*user.DeleteHelpCategoryResponse, error) {
	err := l.svcCtx.HelpCategory.Delete(l.ctx, in.HelpCategoryId)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	return &user.DeleteHelpCategoryResponse{
		IsSuccess: true,
	}, nil
}
