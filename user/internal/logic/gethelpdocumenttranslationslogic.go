package logic

import (
	"context"

	"user/internal/svc"
	"user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHelpDocumentTranslationsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetHelpDocumentTranslationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHelpDocumentTranslationsLogic {
	return &GetHelpDocumentTranslationsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetHelpDocumentTranslations 获取帮助文档翻译列表
func (l *GetHelpDocumentTranslationsLogic) GetHelpDocumentTranslations(in *user.GetHelpDocumentTranslationsRequest) (*user.GetHelpDocumentTranslationsResponse, error) {
	// todo: add your logic here and delete this line

	// 获取帮助分类翻译列表
	dbHelpDocumentTranslation, err := l.svcCtx.HelpDocumentTranslation.FindOneByHelpDocumentIdLanguageCode(l.ctx, in.HelpDocumentId, in.Language)

	if err != nil {
		logx.Error(err)
		return nil, err
	}

	return &user.GetHelpDocumentTranslationsResponse{
		HelpDocumentId:  dbHelpDocumentTranslation.HelpDocumentId,
		DocumentContent: dbHelpDocumentTranslation.DocumentContent,
		DocumentTitle:   dbHelpDocumentTranslation.DocumentTitle,
		Language:        dbHelpDocumentTranslation.LanguageCode,
	}, nil
}
