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

func (l *EditHelpCategoryLogic) EditHelpCategory(in *user.EditHelpCategoryRequest) (*user.EditHelpCategoryResponse, error) {

	// TODO:  完成修改 help_category 逻辑

	// EditHelpCategory

	// 先查出来再修改

	// 1. 先查出来

	return &user.EditHelpCategoryResponse{}, nil
}
