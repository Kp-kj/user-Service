package logic

import (
	"context"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"strconv"
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
	var helpCategoryId int64
	helpCategoryId = node.Generate().Int64()
	_, err = l.svcCtx.HelpCategory.Insert(l.ctx, &model.HelpCategory{
		HelpCategoryId: helpCategoryId,
		CategoryStatus: 1,
	})

	if err != nil {
		fmt.Println("err:")
		logx.Error(err)
		return nil, err
	}
	fmt.Println("helpCategoryId:", helpCategoryId)

	// 在这里使用 helpCategoryId
	return &user.CreateHelpCategoryResponse{
		HelpCategoryId: strconv.FormatInt(helpCategoryId, 10),
	}, nil
}
