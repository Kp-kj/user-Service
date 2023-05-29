package logic

import (
	"context"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
	"user/internal/model"
	"user/internal/svc"
	"user/user"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	var dbUser model.User
	dbUser.TwitterId = in.TwitterId
	// 雪花ID生成
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
	}
	dbUser.UserId = node.Generate().Int64()
	_, err = l.svcCtx.UserModel.Insert(l.ctx, &dbUser) // 保存到数据库
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &user.CreateUserResponse{
		UserId: strconv.FormatInt(dbUser.UserId, 10),
	}, nil
}
