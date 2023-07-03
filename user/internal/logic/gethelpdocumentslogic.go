package logic

import (
	"context"

	"user/internal/svc"
	"user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHelpDocumentsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetHelpDocumentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHelpDocumentsLogic {
	return &GetHelpDocumentsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取帮助文档列表
func (l *GetHelpDocumentsLogic) GetHelpDocuments(in *user.GetHelpDocumentsRequest) (*user.GetHelpDocumentsResponse, error) {
	// todo: add your logic here and delete this line

	// 获取帮助分类列表
	dbHelpDocument, err := l.svcCtx.HelpDocument.FindAll(l.ctx)
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	response := &user.GetHelpDocumentsResponse{
		HelpDocuments: make([]*user.HelpDocument, 0, len(dbHelpDocument)),
	}

	for _, category := range dbHelpDocument {
		helpDocument := &user.HelpDocument{
			HelpDocumentId: category.HelpDocumentId, // 帮助文档id
			HelpCategoryId: category.HelpCategoryId, // 帮助分类id
			DocumentStatus: category.DocumentStatus, // 文档状态
		}
		response.HelpDocuments = append(response.HelpDocuments, helpDocument)
	}

	return response, nil
}
