package logic

import (
	"context"

	"user/internal/svc"
	"user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryHelpCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryHelpCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryHelpCategoryLogic {
	return &QueryHelpCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryHelpCategory 查询帮助分类
func (l *QueryHelpCategoryLogic) QueryHelpCategory(in *user.QueryHelpCategoryRequest) (*user.QueryHelpCategoryResponse, error) {
	//初始化返回值
	var totalCount int64

	var TotalCategorys []*user.TotalCategory

	//当两个都为空情况下查询所有
	if in.CategoryName == "" && in.IsShow == 2 {
		// 先查所有分类
		categoryList, err := l.svcCtx.HelpCategory.FindAll(l.ctx)
		if err != nil {
			return nil, err
		}
		//根据每页条数计算出 totalCount, totalPage, thisCount
		totalCount = int64(len(categoryList))
		//看下有没有当lastId
		//有就从lastId开始取，没有就从0开始取
		// 没有lastId
		if in.LastId == 0 {
			// 只遍历 categorylist前30条
			for i, v := range categoryList {
				// 如果i小于30
				if i < 30 {
					// 调用获取帮助文档翻译接口  这将获取到帮助文档的名称
					data, err := l.svcCtx.HelpCategoryTranslation.FindOneByHelpCategoryIdLanguageCode(l.ctx, v.HelpCategoryId, "")
					if err != nil {
						return nil, err
					}
					// 往TotalCategorys里面添加数据
					TotalCategorys = append(TotalCategorys, &user.TotalCategory{
						CategoryId:   v.HelpCategoryId,
						CategoryName: data.CategoryName,
						IsShow:       v.CategoryStatus,
					})
				}
			}

			return &user.QueryHelpCategoryResponse{
				TotalCount:     totalCount,
				TotalCategorys: TotalCategorys,
			}, nil
		} else {
			// 有lastId
			// 则categoryList中lastId开始的后面的30条
			// 记录下标的变量
			index := 0
			flage := false
			for i, v := range categoryList {
				v.HelpCategoryId = in.LastId // 从lastId开始
				index = i                    // 记录下标
				flage = true                 // 记录是否有lastId
				//如果i-inde小于30
				if flage {
					if i-index < 30 {
						// 调用获取帮助文档翻译接口  这将获取到帮助文档的名称
						data, err := l.svcCtx.HelpCategoryTranslation.FindOneByHelpCategoryIdLanguageCode(l.ctx, v.HelpCategoryId, "")
						if err != nil {
							return nil, err
						}
						// 往TotalCategorys里面添加数据
						TotalCategorys = append(TotalCategorys, &user.TotalCategory{
							CategoryId:   v.HelpCategoryId,
							CategoryName: data.CategoryName,
							IsShow:       v.CategoryStatus,
						})
					}
				}
			}
		}
		return &user.QueryHelpCategoryResponse{
			TotalCount:     totalCount,
			TotalCategorys: TotalCategorys,
		}, nil
	}

	// 两个都不为空
	if in.CategoryName != "" && in.IsShow != 2 {
		// 先查所有分类
		categoryList, err := l.svcCtx.HelpCategory.FindAll(l.ctx)
		if err != nil {
			return nil, err
		}
		//遍历categoryList 找出符合条件的
		for _, v := range categoryList {
			if v.CategoryStatus == in.IsShow {
				data, err := l.svcCtx.HelpCategoryTranslation.FindOne(l.ctx, v.HelpCategoryId)
				if err != nil {
					return nil, err
				}
				if data.CategoryName == in.CategoryName {
					// 往TotalCategorys里面添加数据
					TotalCategorys = append(TotalCategorys, &user.TotalCategory{
						CategoryId:   v.HelpCategoryId,
						CategoryName: data.CategoryName,
						IsShow:       v.CategoryStatus,
					})
				}
			}
		}
		return &user.QueryHelpCategoryResponse{
			TotalCount:     totalCount,
			TotalCategorys: TotalCategorys,
		}, nil
	}

	// 只有分类名称
	if in.CategoryName != "" {
		// 先查所有分类
		categoryList, err := l.svcCtx.HelpCategory.FindAll(l.ctx)
		if err != nil {
			return nil, err
		}
		//遍历categoryList 找出符合条件的
		for _, v := range categoryList {
			data, err := l.svcCtx.HelpCategoryTranslation.FindOne(l.ctx, v.HelpCategoryId)
			if err != nil {
				return nil, err
			}
			if data.CategoryName == in.CategoryName {
				// 往TotalCategorys里面添加数据
				TotalCategorys = append(TotalCategorys, &user.TotalCategory{
					CategoryId:   v.HelpCategoryId,
					CategoryName: data.CategoryName,
					IsShow:       v.CategoryStatus,
				})
			}

		}
		return &user.QueryHelpCategoryResponse{
			TotalCount:     totalCount,
			TotalCategorys: TotalCategorys,
		}, nil
	}

	// 只有是否显示
	if in.IsShow != 2 {
		// 先查所有分类
		categoryList, err := l.svcCtx.HelpCategory.FindAll(l.ctx)
		if err != nil {
			return nil, err
		}
		//遍历categoryList 找出符合条件的
		for _, v := range categoryList {
			if v.CategoryStatus == in.IsShow {
				data, err := l.svcCtx.HelpCategoryTranslation.FindOne(l.ctx, v.HelpCategoryId)
				if err != nil {
					return nil, err
				}
				// 往TotalCategorys里面添加数据
				TotalCategorys = append(TotalCategorys, &user.TotalCategory{
					CategoryId:   v.HelpCategoryId,
					CategoryName: data.CategoryName,
					IsShow:       v.CategoryStatus,
				})

			}
		}
		return &user.QueryHelpCategoryResponse{
			TotalCount:     totalCount,
			TotalCategorys: TotalCategorys,
		}, nil
	}
	return &user.QueryHelpCategoryResponse{
		TotalCount:     totalCount,
		TotalCategorys: TotalCategorys,
	}, nil
}
