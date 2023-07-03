package logic

import (
	"context"

	"user/internal/svc"
	"user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteHelpDocumentTranslationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteHelpDocumentTranslationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteHelpDocumentTranslationLogic {
	return &DeleteHelpDocumentTranslationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteHelpDocumentTranslation 删除帮助文档翻译
func (l *DeleteHelpDocumentTranslationLogic) DeleteHelpDocumentTranslation(in *user.DeleteHelpDocumentTranslationRequest) (*user.DeleteHelpDocumentTranslationResponse, error) {
	err := l.svcCtx.HelpDocumentTranslation.DeleteByHelpDocumentIdLanguageCode(l.ctx, in.HelpDocumentId, in.Language)
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	return &user.DeleteHelpDocumentTranslationResponse{
		IsSuccess: true,
	}, nil
}
