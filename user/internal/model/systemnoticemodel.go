package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SystemNoticeModel = (*customSystemNoticeModel)(nil)

type (
	// SystemNoticeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSystemNoticeModel.
	SystemNoticeModel interface {
		systemNoticeModel
	}

	customSystemNoticeModel struct {
		*defaultSystemNoticeModel
	}
)

// NewSystemNoticeModel returns a model for the database table.
func NewSystemNoticeModel(conn sqlx.SqlConn) SystemNoticeModel {
	return &customSystemNoticeModel{
		defaultSystemNoticeModel: newSystemNoticeModel(conn),
	}
}
