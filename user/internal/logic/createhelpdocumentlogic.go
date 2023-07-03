package logic

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"user/internal/model"
	"user/internal/svc"
	"user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateHelpDocumentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateHelpDocumentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateHelpDocumentLogic {
	return &CreateHelpDocumentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建帮助文档
func (l *CreateHelpDocumentLogic) CreateHelpDocument(in *user.CreateHelpDocumentRequest) (*user.CreateHelpDocumentResponse, error) {
	node, err := snowflake.NewNode(1) // 生成雪花ID
	if err != nil {
		return nil, err
	}

	_, err = l.svcCtx.HelpDocument.Insert(l.ctx, &model.Helpdocument{
		HelpDocumentId: node.Generate().Int64(), //雪花ID
		HelpCategoryId: in.HelpCategoryId,       //帮助分类ID
		DocumentStatus: in.DocumentStatus,       //上架状态 0 下架 1 上架
	})
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	return &user.CreateHelpDocumentResponse{
		IsSuccess: true,
	}, nil
}
