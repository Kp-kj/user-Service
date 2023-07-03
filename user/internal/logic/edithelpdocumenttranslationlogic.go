package logic

import (
	"context"

	"user/internal/svc"
	"user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditHelpDocumentTranslationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEditHelpDocumentTranslationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditHelpDocumentTranslationLogic {
	return &EditHelpDocumentTranslationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// EditHelpDocumentTranslation 编辑帮助文档翻译
func (l *EditHelpDocumentTranslationLogic) EditHelpDocumentTranslation(in *user.EditHelpDocumentTranslationRequest) (*user.EditHelpDocumentTranslationResponse, error) {
	// todo: add your logic here and delete this line

	return &user.EditHelpDocumentTranslationResponse{}, nil
}
