package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ InvitationTreeModel = (*customInvitationTreeModel)(nil)

type (
	// InvitationTreeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customInvitationTreeModel.
	InvitationTreeModel interface {
		invitationTreeModel
	}

	customInvitationTreeModel struct {
		*defaultInvitationTreeModel
	}
)

// NewInvitationTreeModel returns a model for the database table.
func NewInvitationTreeModel(conn sqlx.SqlConn) InvitationTreeModel {
	return &customInvitationTreeModel{
		defaultInvitationTreeModel: newInvitationTreeModel(conn),
	}
}
