// Code generated by goctl. DO NOT EDIT.

package cache

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	invitationsFieldNames          = builder.RawFieldNames(&Invitations{})
	invitationsRows                = strings.Join(invitationsFieldNames, ",")
	invitationsRowsExpectAutoSet   = strings.Join(stringx.Remove(invitationsFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	invitationsRowsWithPlaceHolder = strings.Join(stringx.Remove(invitationsFieldNames, "`invitationId`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	invitationsModel interface {
		Insert(ctx context.Context, data *Invitations) (sql.Result, error)
		FindOne(ctx context.Context, invitationId int64) (*Invitations, error)
		Update(ctx context.Context, data *Invitations) error
		Delete(ctx context.Context, invitationId int64) error
	}

	defaultInvitationsModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Invitations struct {
		InvitationId int64        `db:"invitationId"`
		CreatedAt    time.Time    `db:"created_at"`
		UpdatedAt    time.Time    `db:"updated_at"`
		DeletedAt    sql.NullTime `db:"deleted_at"`
		InviterId    int64        `db:"inviterId"`
		InviteeId    int64        `db:"inviteeId"`
	}
)

func newInvitationsModel(conn sqlx.SqlConn) *defaultInvitationsModel {
	return &defaultInvitationsModel{
		conn:  conn,
		table: "`invitations`",
	}
}

func (m *defaultInvitationsModel) Delete(ctx context.Context, invitationId int64) error {
	query := fmt.Sprintf("delete from %s where `invitationId` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, invitationId)
	return err
}

func (m *defaultInvitationsModel) FindOne(ctx context.Context, invitationId int64) (*Invitations, error) {
	query := fmt.Sprintf("select %s from %s where `invitationId` = ? limit 1", invitationsRows, m.table)
	var resp Invitations
	err := m.conn.QueryRowCtx(ctx, &resp, query, invitationId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultInvitationsModel) Insert(ctx context.Context, data *Invitations) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, invitationsRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.InvitationId, data.DeletedAt, data.InviterId, data.InviteeId)
	return ret, err
}

func (m *defaultInvitationsModel) Update(ctx context.Context, data *Invitations) error {
	query := fmt.Sprintf("update %s set %s where `invitationId` = ?", m.table, invitationsRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.DeletedAt, data.InviterId, data.InviteeId, data.InvitationId)
	return err
}

func (m *defaultInvitationsModel) tableName() string {
	return m.table
}
