package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ HelpdocumenttranslationModel = (*customHelpdocumenttranslationModel)(nil)

type (
	// HelpdocumenttranslationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHelpdocumenttranslationModel.
	HelpdocumenttranslationModel interface {
		helpdocumenttranslationModel
	}

	customHelpdocumenttranslationModel struct {
		*defaultHelpdocumenttranslationModel
	}
)

// NewHelpdocumenttranslationModel returns a model for the database table.
func NewHelpdocumenttranslationModel(conn sqlx.SqlConn) HelpdocumenttranslationModel {
	return &customHelpdocumenttranslationModel{
		defaultHelpdocumenttranslationModel: newHelpdocumenttranslationModel(conn),
	}
}
