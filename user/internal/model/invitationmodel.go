package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ InvitationModel = (*customInvitationModel)(nil)

type (
	// InvitationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customInvitationModel.
	InvitationModel interface {
		invitationModel
	}

	customInvitationModel struct {
		*defaultInvitationModel
	}
)

// NewInvitationModel returns a model for the database table.
func NewInvitationModel(conn sqlx.SqlConn) InvitationModel {
	return &customInvitationModel{
		defaultInvitationModel: newInvitationModel(conn),
	}
}
