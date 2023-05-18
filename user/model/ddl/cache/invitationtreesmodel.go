package cache

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ InvitationTreesModel = (*customInvitationTreesModel)(nil)

type (
	// InvitationTreesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customInvitationTreesModel.
	InvitationTreesModel interface {
		invitationTreesModel
	}

	customInvitationTreesModel struct {
		*defaultInvitationTreesModel
	}
)

// NewInvitationTreesModel returns a model for the database table.
func NewInvitationTreesModel(conn sqlx.SqlConn) InvitationTreesModel {
	return &customInvitationTreesModel{
		defaultInvitationTreesModel: newInvitationTreesModel(conn),
	}
}
