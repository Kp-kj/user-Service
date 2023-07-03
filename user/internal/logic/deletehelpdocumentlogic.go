package logic

import (
	"context"

	"user/internal/svc"
	"user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteHelpDocumentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteHelpDocumentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteHelpDocumentLogic {
	return &DeleteHelpDocumentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除帮助文档
func (l *DeleteHelpDocumentLogic) DeleteHelpDocument(in *user.DeleteHelpDocumentRequest) (*user.DeleteHelpDocumentResponse, error) {
	err := l.svcCtx.HelpDocument.Delete(l.ctx, in.HelpDocumentId)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	return &user.DeleteHelpDocumentResponse{
		IsSuccess: true,
	}, nil
}
