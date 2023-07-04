package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"user/internal/config"
	"user/internal/model"
)

type ServiceContext struct {
	Config                  config.Config
	UserModel               model.UserModel                    //用户
	Invitation              model.InvitationModel              //邀请关系
	InvitationTree          model.InvitationTreeModel          //邀请关系树
	Profile                 model.ProfileModel                 //用户资料
	AdminUser               model.AdminUserModel               //管理员
	Black                   model.BlackModel                   //黑名单
	HelpCategory            model.HelpCategoryModel            //帮助分类
	HelpCategoryTranslation model.HelpCategoryTranslationModel //帮助分类翻译
	HelpDocument            model.HelpdocumentModel            //帮助文档
	HelpDocumentTranslation model.HelpdocumenttranslationModel //帮助文档翻译
	RecordNotice            model.RecordNoticeModel            //通知记录
	SystemNotice            model.SystemNoticeModel            //系统通知
	UserNotice              model.UserNoticeModel              //用户通知

}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                  c,
		UserModel:               model.NewUserModel(sqlx.NewMysql(c.DB.DataSource)),
		Invitation:              model.NewInvitationModel(sqlx.NewMysql(c.DB.DataSource)),
		InvitationTree:          model.NewInvitationTreeModel(sqlx.NewMysql(c.DB.DataSource)),
		Profile:                 model.NewProfileModel(sqlx.NewMysql(c.DB.DataSource)),
		AdminUser:               model.NewAdminUserModel(sqlx.NewMysql(c.DB.DataSource)),
		Black:                   model.NewBlackModel(sqlx.NewMysql(c.DB.DataSource)),
		HelpCategory:            model.NewHelpCategoryModel(sqlx.NewMysql(c.DB.DataSource)),
		HelpCategoryTranslation: model.NewHelpCategoryTranslationModel(sqlx.NewMysql(c.DB.DataSource)),
		HelpDocument:            model.NewHelpdocumentModel(sqlx.NewMysql(c.DB.DataSource)),
		HelpDocumentTranslation: model.NewHelpdocumenttranslationModel(sqlx.NewMysql(c.DB.DataSource)),
		RecordNotice:            model.NewRecordNoticeModel(sqlx.NewMysql(c.DB.DataSource)),
		SystemNotice:            model.NewSystemNoticeModel(sqlx.NewMysql(c.DB.DataSource)),
		UserNotice:              model.NewUserNoticeModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
