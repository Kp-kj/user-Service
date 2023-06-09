package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ BlackModel = (*customBlackModel)(nil)

type (
	// BlackModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBlackModel.
	BlackModel interface {
		blackModel
	}

	customBlackModel struct {
		*defaultBlackModel
	}
)

// NewBlackModel returns a model for the database table.
func NewBlackModel(conn sqlx.SqlConn) BlackModel {
	return &customBlackModel{
		defaultBlackModel: newBlackModel(conn),
	}
}
