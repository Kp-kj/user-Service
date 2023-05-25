// Code generated by goctl. DO NOT EDIT.

package model

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
	adminUserFieldNames          = builder.RawFieldNames(&AdminUser{})
	adminUserRows                = strings.Join(adminUserFieldNames, ",")
	adminUserRowsExpectAutoSet   = strings.Join(stringx.Remove(adminUserFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	adminUserRowsWithPlaceHolder = strings.Join(stringx.Remove(adminUserFieldNames, "`admin_user_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	adminUserModel interface {
		Insert(ctx context.Context, data *AdminUser) (sql.Result, error)
		FindOne(ctx context.Context, adminUserId int64) (*AdminUser, error)
		FindOneByAdminName(ctx context.Context, adminName string) (*AdminUser, error)
		Update(ctx context.Context, data *AdminUser) error
		Delete(ctx context.Context, adminUserId int64) error
	}

	defaultAdminUserModel struct {
		conn  sqlx.SqlConn
		table string
	}

	AdminUser struct {
		AdminUserId int64        `db:"admin_user_id"`
		CreatedAt   time.Time    `db:"created_at"`
		UpdatedAt   sql.NullTime `db:"updated_at"`
		DeletedAt   sql.NullTime `db:"deleted_at"`
		AdminName   string       `db:"admin_name"`
		AdminPasswd string       `db:"admin_passwd"`
		AdminStatus int64        `db:"admin_status"`
		AdminRole   int64        `db:"admin_role"`
	}
)

func newAdminUserModel(conn sqlx.SqlConn) *defaultAdminUserModel {
	return &defaultAdminUserModel{
		conn:  conn,
		table: "`admin_user`",
	}
}

func (m *defaultAdminUserModel) Delete(ctx context.Context, adminUserId int64) error {
	query := fmt.Sprintf("delete from %s where `admin_user_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, adminUserId)
	return err
}

func (m *defaultAdminUserModel) FindOne(ctx context.Context, adminUserId int64) (*AdminUser, error) {
	query := fmt.Sprintf("select %s from %s where `admin_user_id` = ? limit 1", adminUserRows, m.table)
	var resp AdminUser
	err := m.conn.QueryRowCtx(ctx, &resp, query, adminUserId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAdminUserModel) FindOneByAdminName(ctx context.Context, adminName string) (*AdminUser, error) {
	var resp AdminUser
	query := fmt.Sprintf("select %s from %s where `admin_name` = ? limit 1", adminUserRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, adminName)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAdminUserModel) Insert(ctx context.Context, data *AdminUser) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, adminUserRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.AdminUserId, data.DeletedAt, data.AdminName, data.AdminPasswd, data.AdminStatus, data.AdminRole)
	return ret, err
}

func (m *defaultAdminUserModel) Update(ctx context.Context, newData *AdminUser) error {
	query := fmt.Sprintf("update %s set %s where `admin_user_id` = ?", m.table, adminUserRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.DeletedAt, newData.AdminName, newData.AdminPasswd, newData.AdminStatus, newData.AdminRole, newData.AdminUserId)
	return err
}

func (m *defaultAdminUserModel) tableName() string {
	return m.table
}