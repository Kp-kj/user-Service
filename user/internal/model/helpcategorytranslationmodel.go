package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ HelpCategoryTranslationModel = (*customHelpCategoryTranslationModel)(nil)

type (
	// HelpCategoryTranslationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHelpCategoryTranslationModel.
	HelpCategoryTranslationModel interface {
		helpCategoryTranslationModel
	}

	customHelpCategoryTranslationModel struct {
		*defaultHelpCategoryTranslationModel
	}
)

// NewHelpCategoryTranslationModel returns a model for the database table.
func NewHelpCategoryTranslationModel(conn sqlx.SqlConn) HelpCategoryTranslationModel {
	return &customHelpCategoryTranslationModel{
		defaultHelpCategoryTranslationModel: newHelpCategoryTranslationModel(conn),
	}
}
