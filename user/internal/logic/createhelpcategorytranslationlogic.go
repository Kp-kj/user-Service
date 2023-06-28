package logic

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"user/internal/model"

	"user/internal/svc"
	"user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateHelpCategoryTranslationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateHelpCategoryTranslationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateHelpCategoryTranslationLogic {
	return &CreateHelpCategoryTranslationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateHelpCategoryTranslation 创建帮助分类翻译
func (l *CreateHelpCategoryTranslationLogic) CreateHelpCategoryTranslation(in *user.CreateHelpCategoryTranslationRequest) (*user.CreateHelpCategoryTranslationResponse, error) {
	node, err := snowflake.NewNode(1) // 生成雪花ID
	if err != nil {
		return nil, err
	}

	_, err = l.svcCtx.HelpCategoryTranslation.Insert(l.ctx, &model.HelpCategoryTranslation{ // 插入数据
		HelpCategoryTranslationId: node.Generate().Int64(),
		HelpCategoryId:            in.HelpCategoryId,
		LanguageCode:              in.Language,
		CategoryName:              in.CategoryName,
	})
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	return &user.CreateHelpCategoryTranslationResponse{
		IsSuccess: true,
	}, nil
}
