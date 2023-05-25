package logic

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
	"user/internal/model"
	"user/internal/svc"
	"user/user"
)

type AddUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserInfoLogic {
	return &AddUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddUserInfo 添加用户信息
func (l *AddUserInfoLogic) AddUserInfo(in *user.AddUserInfoRequest) (*user.AddUserInfoResponse, error) {

	userId, err := strconv.ParseInt(in.UserId, 10, 64)
	if err != nil {
		return nil, err
	}

	// 生成资料表id
	node, err := snowflake.NewNode(1)
	if err != nil {
		return nil, err
	}

	profile := &model.Profile{
		ProfileId:   node.Generate().Int64(),
		UserId:      userId,
		TwitterName: in.TwitterName,
		UserName:    in.UserName,
		Avatar:      in.UserAvatar,
	}
	_, err = l.svcCtx.Profile.Insert(l.ctx, profile)
	if err != nil {
		return &user.AddUserInfoResponse{
			IsSuccess: false,
		}, nil
	}
	return &user.AddUserInfoResponse{
		IsSuccess: true,
	}, nil
}
