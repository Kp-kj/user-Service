package logic

import (
	"context"

	"user/internal/svc"
	"user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteHelpCategoryTranslationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteHelpCategoryTranslationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteHelpCategoryTranslationLogic {
	return &DeleteHelpCategoryTranslationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteHelpCategoryTranslation 删除帮助分类翻译
func (l *DeleteHelpCategoryTranslationLogic) DeleteHelpCategoryTranslation(in *user.DeleteHelpCategoryTranslationRequest) (*user.DeleteHelpCategoryTranslationResponse, error) {
	err := l.svcCtx.HelpCategoryTranslation.DeleteByHelpCategoryIdLanguageCode(l.ctx, in.HelpCategoryId, in.Language)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	return &user.DeleteHelpCategoryTranslationResponse{
		IsSuccess: true,
	}, nil
}
