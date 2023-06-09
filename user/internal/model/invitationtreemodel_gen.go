// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
	"strings"
	"time"
)

var (
	invitationTreeFieldNames          = builder.RawFieldNames(&InvitationTree{})
	invitationTreeRows                = strings.Join(invitationTreeFieldNames, ",")
	invitationTreeRowsExpectAutoSet   = strings.Join(stringx.Remove(invitationTreeFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	invitationTreeRowsWithPlaceHolder = strings.Join(stringx.Remove(invitationTreeFieldNames, "`invitation_tree_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	invitationTreeModel interface {
		Insert(ctx context.Context, data *InvitationTree) (sql.Result, error)
		FindOne(ctx context.Context, invitationTreeId int64) (*InvitationTree, error)
		FindOneByInvitationTreeId(ctx context.Context, invitationTreeId int64) (*InvitationTree, error)
		Update(ctx context.Context, data *InvitationTree) error
		Delete(ctx context.Context, invitationTreeId int64) error
		FindParentID(ctx context.Context, invitationTreeId int64)(*InvitationTree, error)
	}

	defaultInvitationTreeModel struct {
		conn  sqlx.SqlConn
		table string
	}

	InvitationTree struct {
		InvitationTreeId int64        `db:"invitation_tree_id"`
		CreatedAt        time.Time    `db:"created_at"`
		UpdatedAt        sql.NullTime `db:"updated_at"`
		DeletedAt        sql.NullTime `db:"deleted_at"`
		UserId           int64        `db:"user_id"`
		ParentId         int64        `db:"parent_id"`
		Depth            int64        `db:"depth"`
	}
)

func newInvitationTreeModel(conn sqlx.SqlConn) *defaultInvitationTreeModel {
	return &defaultInvitationTreeModel{
		conn:  conn,
		table: "`invitation_tree`",
	}
}


func (m *defaultInvitationTreeModel) Delete(ctx context.Context, invitationTreeId int64) error {
	query := fmt.Sprintf("delete from %s where `invitation_tree_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, invitationTreeId)
	return err
}

func (m *defaultInvitationTreeModel) FindOne(ctx context.Context, invitationTreeId int64) (*InvitationTree, error) {
	query := fmt.Sprintf("select %s from %s where `invitation_tree_id` = ? limit 1", invitationTreeRows, m.table)
	var resp InvitationTree
	err := m.conn.QueryRowCtx(ctx, &resp, query, invitationTreeId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultInvitationTreeModel) FindOneByInvitationTreeId(ctx context.Context, invitationTreeId int64) (*InvitationTree, error) {
	var resp InvitationTree
	query := fmt.Sprintf("select %s from %s where `invitation_tree_id` = ? limit 1", invitationTreeRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, invitationTreeId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultInvitationTreeModel) Insert(ctx context.Context, data *InvitationTree) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, invitationTreeRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.InvitationTreeId, data.DeletedAt, data.UserId, data.ParentId, data.Depth)
	return ret, err
}

func (m *defaultInvitationTreeModel) Update(ctx context.Context, newData *InvitationTree) error {
	query := fmt.Sprintf("update %s set %s where `invitation_tree_id` = ?", m.table, invitationTreeRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.DeletedAt, newData.UserId, newData.ParentId, newData.Depth, newData.InvitationTreeId)
	return err
}

// FindParentID 根据父节点id查找
func (m *defaultInvitationTreeModel) FindParentID(ctx context.Context, invitationTreeId int64) (*InvitationTree, error) {
	var resp InvitationTree
	query := fmt.Sprintf("select %s from %s where `invitation_tree_id` = ?", invitationTreeRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, invitationTreeId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultInvitationTreeModel) tableName() string {
	return m.table
}
