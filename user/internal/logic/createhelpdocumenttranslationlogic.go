package logic

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"user/internal/model"

	"user/internal/svc"
	"user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateHelpDocumentTranslationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateHelpDocumentTranslationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateHelpDocumentTranslationLogic {
	return &CreateHelpDocumentTranslationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateHelpDocumentTranslationLogic) CreateHelpDocumentTranslation(in *user.CreateHelpDocumentTranslationRequest) (*user.CreateHelpDocumentTranslationResponse, error) {
	node, err := snowflake.NewNode(1) // 生成雪花ID
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.HelpDocumentTranslation.Insert(l.ctx, &model.Helpdocumenttranslation{ // 插入数据
		HelpDocumentTranslationId: node.Generate().Int64(),
		HelpDocumentId:            in.HelpDocumentId,  // 帮助文档ID
		LanguageCode:              in.Language,        // 语言编码
		DocumentTitle:             in.DocumentTitle,   // 文档标题
		DocumentContent:           in.DocumentContent, // 文档内容
	})
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	return &user.CreateHelpDocumentTranslationResponse{
		IsSuccess: true,
	}, nil
}
