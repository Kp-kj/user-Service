package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ HelpdocumentModel = (*customHelpdocumentModel)(nil)

type (
	// HelpdocumentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHelpdocumentModel.
	HelpdocumentModel interface {
		helpdocumentModel
	}

	customHelpdocumentModel struct {
		*defaultHelpdocumentModel
	}
)

// NewHelpdocumentModel returns a model for the database table.
func NewHelpdocumentModel(conn sqlx.SqlConn) HelpdocumentModel {
	return &customHelpdocumentModel{
		defaultHelpdocumentModel: newHelpdocumentModel(conn),
	}
}
