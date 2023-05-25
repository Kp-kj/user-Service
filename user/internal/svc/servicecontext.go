package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"user/internal/config"
	"user/internal/model"
)

type ServiceContext struct {
	Config         config.Config
	UserModel      model.UserModel           //用户
	Invitation     model.InvitationModel     //邀请关系
	InvitationTree model.InvitationTreeModel //邀请关系树
	Profile        model.ProfileModel        //用户资料
	AdminUser      model.AdminUserModel      //管理员
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		UserModel:      model.NewUserModel(sqlx.NewMysql(c.DB.DataSource)),
		Invitation:     model.NewInvitationModel(sqlx.NewMysql(c.DB.DataSource)),
		InvitationTree: model.NewInvitationTreeModel(sqlx.NewMysql(c.DB.DataSource)),
		Profile:        model.NewProfileModel(sqlx.NewMysql(c.DB.DataSource)),
		AdminUser:      model.NewAdminUserModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
