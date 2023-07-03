package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ RecordNoticeModel = (*customRecordNoticeModel)(nil)

type (
	// RecordNoticeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRecordNoticeModel.
	RecordNoticeModel interface {
		recordNoticeModel
	}

	customRecordNoticeModel struct {
		*defaultRecordNoticeModel
	}
)

// NewRecordNoticeModel returns a model for the database table.
func NewRecordNoticeModel(conn sqlx.SqlConn) RecordNoticeModel {
	return &customRecordNoticeModel{
		defaultRecordNoticeModel: newRecordNoticeModel(conn),
	}
}
