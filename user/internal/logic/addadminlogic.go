package logic

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"golang.org/x/crypto/bcrypt"
	"user/internal/model"

	"user/internal/svc"
	"user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddAdminLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddAdminLogic {
	return &AddAdminLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddAdmin 添加管理员
func (l *AddAdminLogic) AddAdmin(in *user.AddAdminRequest) (*user.AddAdminResponse, error) {
	//查询是否有管理员
	_, err := l.svcCtx.AdminUser.FindOneByAdminName(l.ctx, in.AdminName)
	if err == nil {
		logx.Error(err)
		return nil, err
	}
	//生成管理员id
	node, err := snowflake.NewNode(1)
	if err != nil {
		return nil, err
	}

	//加密
	encryptionPwd, err := bcrypt.GenerateFromPassword([]byte(in.AdminPasswd), bcrypt.DefaultCost)

	adminUser := &model.AdminUser{
		AdminUserId: node.Generate().Int64(),
		AdminName:   in.AdminName,
		AdminPasswd: string(encryptionPwd),
		AdminStatus: 0, //0 正常 1 禁用
		AdminRole:   1, //1 管理员 2 超级管理员（暂时不用）
	}

	_, err = l.svcCtx.AdminUser.Insert(l.ctx, adminUser)
	if err != nil {
		return &user.AddAdminResponse{
			IsSuccess: false,
		}, nil
	}
	return &user.AddAdminResponse{
		IsSuccess: true,
	}, nil
}
