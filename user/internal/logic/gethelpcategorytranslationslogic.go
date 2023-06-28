package logic

import (
	"context"
	"user/internal/svc"
	"user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHelpCategoryTranslationsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetHelpCategoryTranslationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHelpCategoryTranslationsLogic {
	return &GetHelpCategoryTranslationsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetHelpCategoryTranslations 获取帮助分类翻译列表
func (l *GetHelpCategoryTranslationsLogic) GetHelpCategoryTranslations(in *user.GetHelpCategoryTranslationsRequest) (*user.GetHelpCategoryTranslationsResponse, error) {
	// 获取帮助分类翻译列表

	dbHelpCategoryTranslation, err := l.svcCtx.HelpCategoryTranslation.FindOneByHelpCategoryIdLanguageCode(l.ctx, in.HelpCategoryId, in.Language)

	if err != nil {
		logx.Error(err)
		return nil, err
	}

	return &user.GetHelpCategoryTranslationsResponse{
		HelpCategoryId: dbHelpCategoryTranslation.HelpCategoryId,
		Language:       dbHelpCategoryTranslation.LanguageCode,
		CategoryName:   dbHelpCategoryTranslation.CategoryName,
	}, nil
}
