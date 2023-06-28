package logic

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"time"
	"user/internal/model"
	"user/internal/svc"
	"user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateHelpCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateHelpCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateHelpCategoryLogic {
	return &CreateHelpCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateHelpCategory 创建帮助分类
func (l *CreateHelpCategoryLogic) CreateHelpCategory(in *user.CreateHelpCategoryRequest) (*user.CreateHelpCategoryResponse, error) {

	node, err := snowflake.NewNode(1) // 生成雪花ID
	if err != nil {
		return nil, err
	}

	_, err = l.svcCtx.HelpCategory.Insert(l.ctx, &model.HelpCategory{ // 插入数据
		HelpCategoryId: node.Generate().Int64(),
		CreatedAt:      time.Time{},
		CategoryStatus: 1,
	})
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	return &user.CreateHelpCategoryResponse{
		IsSuccess: true,
	}, nil
}
