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
	helpCategoryTranslationFieldNames          = builder.RawFieldNames(&HelpCategoryTranslation{})
	helpCategoryTranslationRows                = strings.Join(helpCategoryTranslationFieldNames, ",")
	helpCategoryTranslationRowsExpectAutoSet   = strings.Join(stringx.Remove(helpCategoryTranslationFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	helpCategoryTranslationRowsWithPlaceHolder = strings.Join(stringx.Remove(helpCategoryTranslationFieldNames, "`helpCategoryTranslation_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	helpCategoryTranslationModel interface {
		Insert(ctx context.Context, data *HelpCategoryTranslation) (sql.Result, error)
		FindOne(ctx context.Context, helpCategoryTranslationId int64) (*HelpCategoryTranslation, error)
		FindOneByHelpCategoryIdLanguageCode(ctx context.Context, helpCategoryId int64, languageCode string) (*HelpCategoryTranslation, error)
		Update(ctx context.Context, data *HelpCategoryTranslation) error
		Delete(ctx context.Context, helpCategoryTranslationId int64) error
		DeleteByHelpCategoryIdLanguageCode(ctx context.Context, helpCategoryId int64, languageCode string) error
		EditHelpCategoryTranslation(ctx context.Context, helpCategoryId int64, languageCode string, categoryName string) error
	}

	defaultHelpCategoryTranslationModel struct {
		conn  sqlx.SqlConn
		table string
	}

	HelpCategoryTranslation struct {
		HelpCategoryTranslationId int64        `db:"helpCategoryTranslation_id"`
		CreatedAt                 time.Time    `db:"created_at"`
		UpdatedAt                 sql.NullTime `db:"updated_at"`
		DeletedAt                 sql.NullTime `db:"deleted_at"`
		HelpCategoryId            int64        `db:"helpCategory_id"`
		LanguageCode              string       `db:"language_code"`
		CategoryName              string       `db:"category_name"`
	}
)

func newHelpCategoryTranslationModel(conn sqlx.SqlConn) *defaultHelpCategoryTranslationModel {
	return &defaultHelpCategoryTranslationModel{
		conn:  conn,
		table: "`helpCategoryTranslation`",
	}
}

func (m *defaultHelpCategoryTranslationModel) Delete(ctx context.Context, helpCategoryTranslationId int64) error {
	query := fmt.Sprintf("delete from %s where `helpCategoryTranslation_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, helpCategoryTranslationId)
	return err
}


// DeleteByHelpCategoryIdLanguageCode 根据帮助分类ID和语言代码删除
func (m *defaultHelpCategoryTranslationModel) DeleteByHelpCategoryIdLanguageCode(ctx context.Context, helpCategoryId int64, languageCode string) error {
	query := fmt.Sprintf("delete from %s where `helpCategory_id` = ? and `language_code` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, helpCategoryId, languageCode)
	return err
}



func (m *defaultHelpCategoryTranslationModel) FindOne(ctx context.Context, helpCategoryTranslationId int64) (*HelpCategoryTranslation, error) {
	query := fmt.Sprintf("select %s from %s where `helpCategoryTranslation_id` = ? limit 1", helpCategoryTranslationRows, m.table)
	var resp HelpCategoryTranslation
	err := m.conn.QueryRowCtx(ctx, &resp, query, helpCategoryTranslationId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// FindOneByHelpCategoryIdLanguageCode 根据帮助分类ID和语言代码查找
func (m *defaultHelpCategoryTranslationModel) FindOneByHelpCategoryIdLanguageCode(ctx context.Context, helpCategoryId int64, languageCode string) (*HelpCategoryTranslation, error) {
	query := fmt.Sprintf("select %s from %s where `helpCategory_id` = ? and `language_code` = ? limit 1", helpCategoryTranslationRows, m.table)
	var resp HelpCategoryTranslation
	err := m.conn.QueryRowCtx(ctx, &resp, query, helpCategoryId, languageCode)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

//EditHelpCategoryTranslation
func (m *defaultHelpCategoryTranslationModel) EditHelpCategoryTranslation(ctx context.Context, helpCategoryId int64, languageCode string, categoryName string) error {
	query := fmt.Sprintf("update %s set `category_name` = ? where `helpCategory_id` = ? and `language_code` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, categoryName, helpCategoryId, languageCode)
	return err
}

func (m *defaultHelpCategoryTranslationModel) Insert(ctx context.Context, data *HelpCategoryTranslation) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, helpCategoryTranslationRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.HelpCategoryTranslationId, data.DeletedAt, data.HelpCategoryId, data.LanguageCode, data.CategoryName)
	return ret, err
}

func (m *defaultHelpCategoryTranslationModel) Update(ctx context.Context, newData *HelpCategoryTranslation) error {
	query := fmt.Sprintf("update %s set %s where `helpCategoryTranslation_id` = ?", m.table, helpCategoryTranslationRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.DeletedAt, newData.HelpCategoryId, newData.LanguageCode, newData.CategoryName, newData.HelpCategoryTranslationId)
	return err
}

func (m *defaultHelpCategoryTranslationModel) tableName() string {
	return m.table
}
