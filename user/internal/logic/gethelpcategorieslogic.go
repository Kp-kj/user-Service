package logic

import (
	"context"

	"user/internal/svc"
	"user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHelpCategoriesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetHelpCategoriesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHelpCategoriesLogic {
	return &GetHelpCategoriesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetHelpCategories 获取帮助分类列表
func (l *GetHelpCategoriesLogic) GetHelpCategories(in *user.GetHelpCategoriesRequest) (*user.GetHelpCategoriesResponse, error) {

	// 获取帮助分类列表
	dbHelpCategories, err := l.svcCtx.HelpCategory.FindAll(l.ctx)
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	response := &user.GetHelpCategoriesResponse{
		HelpCategories: make([]*user.HelpCategory, 0, len(dbHelpCategories)),
	}

	for _, category := range dbHelpCategories {
		helpCategory := &user.HelpCategory{
			HelpCategoryId: category.HelpCategoryId,
			CreatedAt:      category.CreatedAt.Unix(), // 时间戳
			CategoryStatus: category.CategoryStatus,
		}
		response.HelpCategories = append(response.HelpCategories, helpCategory)
	}

	return response, nil
}
