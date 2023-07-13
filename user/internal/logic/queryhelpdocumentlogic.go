package logic

import (
	"context"

	"user/internal/svc"
	"user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryHelpDocumentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryHelpDocumentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryHelpDocumentLogic {
	return &QueryHelpDocumentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//
func (l *QueryHelpDocumentLogic) QueryHelpDocument(in *user.QueryHelpDocumentRequest) (*user.QueryHelpDocumentResponse, error) {
	//初始化返回数据
	//var TotalCount int64
	//var List []*user.HelpDocumentData

	// 如果标题和状态都为空，则查询所有
	if in.DocumentTitle == "" && in.IsShow == 3 {
		//查找所有分类内容

	}

	return &user.QueryHelpDocumentResponse{}, nil
}
