package logic

import (
	"context"

	"user/internal/svc"
	"user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditHelpDocumentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEditHelpDocumentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditHelpDocumentLogic {
	return &EditHelpDocumentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// EditHelpDocument 编辑帮助文档
func (l *EditHelpDocumentLogic) EditHelpDocument(in *user.EditHelpDocumentRequest) (*user.EditHelpDocumentResponse, error) {
	// 修改帮助分类上下架状态
	err := l.svcCtx.HelpDocument.Edithelpdocument(l.ctx, in.HelpDocumentId, in.DocumentStatus)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	return &user.EditHelpDocumentResponse{IsSuccess: true}, nil
}
