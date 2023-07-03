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
	systemNoticeFieldNames          = builder.RawFieldNames(&SystemNotice{})
	systemNoticeRows                = strings.Join(systemNoticeFieldNames, ",")
	systemNoticeRowsExpectAutoSet   = strings.Join(stringx.Remove(systemNoticeFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	systemNoticeRowsWithPlaceHolder = strings.Join(stringx.Remove(systemNoticeFieldNames, "`systemNotice_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	systemNoticeModel interface {
		Insert(ctx context.Context, data *SystemNotice) (sql.Result, error)
		FindOne(ctx context.Context, systemNoticeId int64) (*SystemNotice, error)
		Update(ctx context.Context, data *SystemNotice) error
		Delete(ctx context.Context, systemNoticeId int64) error
		FindOneByNoticeCategoryNoticeStatus(ctx context.Context, noticeCategory int64, noticeStatus int64) ([]*SystemNotice, error)
	}

	defaultSystemNoticeModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SystemNotice struct {
		SystemNoticeId  int64        `db:"systemNotice_id"`
		CreatedAt       time.Time    `db:"created_at"`
		UpdatedAt       sql.NullTime `db:"updated_at"`
		DeletedAt       sql.NullTime `db:"deleted_at"`
		NoticeTitle     string       `db:"notice_title"`
		NoticeContent   string       `db:"notice_content"`
		NoticeCategory  int64        `db:"notice_category"`
		NoticeStatus    int64        `db:"notice_status"`
		NoticeStartTime sql.NullTime `db:"notice_start_time"`
	}
)

func newSystemNoticeModel(conn sqlx.SqlConn) *defaultSystemNoticeModel {
	return &defaultSystemNoticeModel{
		conn:  conn,
		table: "`systemNotice`",
	}
}

func (m *defaultSystemNoticeModel) Delete(ctx context.Context, systemNoticeId int64) error {
	query := fmt.Sprintf("delete from %s where `systemNotice_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, systemNoticeId)
	return err
}

func (m *defaultSystemNoticeModel) FindOne(ctx context.Context, systemNoticeId int64) (*SystemNotice, error) {
	query := fmt.Sprintf("select %s from %s where `systemNotice_id` = ? limit 1", systemNoticeRows, m.table)
	var resp SystemNotice
	err := m.conn.QueryRowCtx(ctx, &resp, query, systemNoticeId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

//FindOneByNoticeCategoryNoticeStatus
func (m *defaultSystemNoticeModel) FindOneByNoticeCategoryNoticeStatus(ctx context.Context, noticeCategory int64, noticeStatus int64) ([]*SystemNotice, error) {
	var resp []*SystemNotice
	var query string

	if noticeCategory == 3 {
		// Retrieve all notices regardless of the notice category
		query = fmt.Sprintf("SELECT %s FROM %s WHERE `notice_status` = ?", systemNoticeRows, m.table)
	} else {
		query = fmt.Sprintf("SELECT %s FROM %s WHERE `notice_category` = ? AND `notice_status` = ?", systemNoticeRows, m.table)
	}

	err := m.conn.QueryRowsCtx(ctx, &resp, query, noticeCategory, noticeStatus)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}


func (m *defaultSystemNoticeModel) Insert(ctx context.Context, data *SystemNotice) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, systemNoticeRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.SystemNoticeId, data.DeletedAt, data.NoticeTitle, data.NoticeContent, data.NoticeCategory, data.NoticeStatus, data.NoticeStartTime)
	return ret, err
}

func (m *defaultSystemNoticeModel) Update(ctx context.Context, data *SystemNotice) error {
	query := fmt.Sprintf("update %s set %s where `systemNotice_id` = ?", m.table, systemNoticeRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.DeletedAt, data.NoticeTitle, data.NoticeContent, data.NoticeCategory, data.NoticeStatus, data.NoticeStartTime, data.SystemNoticeId)
	return err
}

func (m *defaultSystemNoticeModel) tableName() string {
	return m.table
}
