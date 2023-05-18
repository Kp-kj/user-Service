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
	profilesFieldNames          = builder.RawFieldNames(&Profiles{})
	profilesRows                = strings.Join(profilesFieldNames, ",")
	profilesRowsExpectAutoSet   = strings.Join(stringx.Remove(profilesFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	profilesRowsWithPlaceHolder = strings.Join(stringx.Remove(profilesFieldNames, "`profileId`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	profilesModel interface {
		Insert(ctx context.Context, data *Profiles) (sql.Result, error)
		FindOne(ctx context.Context, profileId int64) (*Profiles, error)
		Update(ctx context.Context, data *Profiles) error
		Delete(ctx context.Context, profileId int64) error
	}

	defaultProfilesModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Profiles struct {
		ProfileId int64        `db:"profileId"`
		CreatedAt time.Time    `db:"created_at"`
		UpdatedAt time.Time    `db:"updated_at"`
		DeletedAt sql.NullTime `db:"deleted_at"`
		UserId    int64        `db:"userId"`
		UserName  string       `db:"userName"`
		Avatar    string       `db:"avatar"`
	}
)

func newProfilesModel(conn sqlx.SqlConn) *defaultProfilesModel {
	return &defaultProfilesModel{
		conn:  conn,
		table: "`profiles`",
	}
}

func (m *defaultProfilesModel) Delete(ctx context.Context, profileId int64) error {
	query := fmt.Sprintf("delete from %s where `profileId` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, profileId)
	return err
}

func (m *defaultProfilesModel) FindOne(ctx context.Context, profileId int64) (*Profiles, error) {
	query := fmt.Sprintf("select %s from %s where `profileId` = ? limit 1", profilesRows, m.table)
	var resp Profiles
	err := m.conn.QueryRowCtx(ctx, &resp, query, profileId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultProfilesModel) Insert(ctx context.Context, data *Profiles) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, profilesRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.ProfileId, data.DeletedAt, data.UserId, data.UserName, data.Avatar)
	return ret, err
}

func (m *defaultProfilesModel) Update(ctx context.Context, data *Profiles) error {
	query := fmt.Sprintf("update %s set %s where `profileId` = ?", m.table, profilesRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.DeletedAt, data.UserId, data.UserName, data.Avatar, data.ProfileId)
	return err
}

func (m *defaultProfilesModel) tableName() string {
	return m.table
}
