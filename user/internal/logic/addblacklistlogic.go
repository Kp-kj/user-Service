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

type AddBlackListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddBlackListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddBlackListLogic {
	return &AddBlackListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddBlackList 添加黑名单
func (l *AddBlackListLogic) AddBlackList(in *user.AddBlackListRequest) (*user.AddBlackListResponse, error) {

	//生成管理员id
	node, err := snowflake.NewNode(1)
	if err != nil {
		return nil, err
	}

	black := &model.Black{
		BlackId:     node.Generate().Int64(),     // 黑名单ID
		UserId:      in.UserId,                   // 用户ID
		BlackUserId: in.UserId,                   // 黑名单用户ID   TODO:这里重复了，需要修改。应该删掉一个
		StartTime:   time.Now(),                  // 开始时间
		EndTime:     time.Now().AddDate(0, 0, 7), // 结束时间
		BlackType:   0,                           //  0 永久 1 临时
	}

	_, err = l.svcCtx.Black.Insert(l.ctx, black) // 插入黑名单
	if err != nil {
		logx.Error(err)
		return &user.AddBlackListResponse{IsSuccess: false}, err
	}

	return &user.AddBlackListResponse{
		IsSuccess: true,
	}, nil
}
