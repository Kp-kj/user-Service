package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserNoticeModel = (*customUserNoticeModel)(nil)

type (
	// UserNoticeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserNoticeModel.
	UserNoticeModel interface {
		userNoticeModel
	}

	customUserNoticeModel struct {
		*defaultUserNoticeModel
	}
)

// NewUserNoticeModel returns a model for the database table.
func NewUserNoticeModel(conn sqlx.SqlConn) UserNoticeModel {
	return &customUserNoticeModel{
		defaultUserNoticeModel: newUserNoticeModel(conn),
	}
}
