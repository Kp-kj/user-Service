package logic

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"user/internal/svc"
	"user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdminLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminLoginLogic {
	return &AdminLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AdminLogin 管理员登录
func (l *AdminLoginLogic) AdminLogin(in *user.AdminLoginRequest) (*user.AdminLoginResponse, error) {

	//查询是否有管理员
	dbAdminUser, err := l.svcCtx.AdminUser.FindOneByAdminName(l.ctx, in.AdminName)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	// 解密 & 比较
	err = bcrypt.CompareHashAndPassword([]byte(dbAdminUser.AdminPasswd), []byte(in.AdminPasswd))
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	return &user.AdminLoginResponse{
		AdminUserId: dbAdminUser.AdminUserId,
	}, nil
}
