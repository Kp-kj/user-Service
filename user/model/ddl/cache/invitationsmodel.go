package cache

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ InvitationsModel = (*customInvitationsModel)(nil)

type (
	// InvitationsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customInvitationsModel.
	InvitationsModel interface {
		invitationsModel
	}

	customInvitationsModel struct {
		*defaultInvitationsModel
	}
)

// NewInvitationsModel returns a model for the database table.
func NewInvitationsModel(conn sqlx.SqlConn) InvitationsModel {
	return &customInvitationsModel{
		defaultInvitationsModel: newInvitationsModel(conn),
	}
}
