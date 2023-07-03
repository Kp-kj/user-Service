package logic

import (
	"context"

	"user/internal/svc"
	"user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditHelpCategoryTranslationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEditHelpCategoryTranslationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditHelpCategoryTranslationLogic {
	return &EditHelpCategoryTranslationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 编辑帮助分类翻译
func (l *EditHelpCategoryTranslationLogic) EditHelpCategoryTranslation(in *user.EditHelpCategoryTranslationRequest) (*user.EditHelpCategoryTranslationResponse, error) {
	// 修改帮助分类上下架状态
	err := l.svcCtx.HelpCategoryTranslation.EditHelpCategoryTranslation(l.ctx, in.HelpCategoryId, in.Language, in.CategoryName)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	return &user.EditHelpCategoryTranslationResponse{}, nil
}
