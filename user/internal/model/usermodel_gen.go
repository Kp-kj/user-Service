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
	userFieldNames          = builder.RawFieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, "`user_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	userModel interface {
		Insert(ctx context.Context, data *User) (sql.Result, error)
		FindOne(ctx context.Context, userId int64) (*User, error)
		FindOneByTwitterId(ctx context.Context, twitterId string) (*User, error)
		Update(ctx context.Context, data *User) error
		Delete(ctx context.Context, userId int64) error
	}

	defaultUserModel struct {
		conn  sqlx.SqlConn
		table string
	}

	User struct {
		UserId    int64        `db:"user_id"`
		CreatedAt time.Time    `db:"created_at"`
		UpdatedAt sql.NullTime `db:"updated_at"`
		DeletedAt sql.NullTime `db:"deleted_at"`
		TwitterId string       `db:"twitter_id"`
	}
)

func newUserModel(conn sqlx.SqlConn) *defaultUserModel {
	return &defaultUserModel{
		conn:  conn,
		table: "`user`",
	}
}

func (m *defaultUserModel) Delete(ctx context.Context, userId int64) error {
	query := fmt.Sprintf("delete from %s where `user_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, userId)
	return err
}

func (m *defaultUserModel) FindOne(ctx context.Context, userId int64) (*User, error) {
	query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", userRows, m.table)
	var resp User
	err := m.conn.QueryRowCtx(ctx, &resp, query, userId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindOneByTwitterId(ctx context.Context, twitterId string) (*User, error) {
	var resp User
	query := fmt.Sprintf("select %s from %s where `twitter_id` = ? limit 1", userRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, twitterId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) Insert(ctx context.Context, data *User) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, userRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserId, data.DeletedAt, data.TwitterId)
	return ret, err
}

func (m *defaultUserModel) Update(ctx context.Context, newData *User) error {
	query := fmt.Sprintf("update %s set %s where `user_id` = ?", m.table, userRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.DeletedAt, newData.TwitterId, newData.UserId)
	return err
}

func (m *defaultUserModel) tableName() string {
	return m.table
}
