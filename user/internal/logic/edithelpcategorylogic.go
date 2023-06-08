package logic

import (
	"context"

	"user/internal/svc"
	"user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditHelpCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEditHelpCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditHelpCategoryLogic {
	return &EditHelpCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// EditHelpCategory 修改帮助分类上下架状态
func (l *EditHelpCategoryLogic) EditHelpCategory(in *user.EditHelpCategoryRequest) (*user.EditHelpCategoryResponse, error) {

	// 修改帮助分类上下架状态
	err := l.svcCtx.HelpCategory.Edith(l.ctx, in.HelpCategoryId, in.CategoryStatus)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	return &user.EditHelpCategoryResponse{
		IsSuccess: true,
	}, nil
}
