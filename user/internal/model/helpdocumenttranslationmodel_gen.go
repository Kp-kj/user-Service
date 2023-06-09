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
	helpdocumenttranslationFieldNames          = builder.RawFieldNames(&Helpdocumenttranslation{})
	helpdocumenttranslationRows                = strings.Join(helpdocumenttranslationFieldNames, ",")
	helpdocumenttranslationRowsExpectAutoSet   = strings.Join(stringx.Remove(helpdocumenttranslationFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	helpdocumenttranslationRowsWithPlaceHolder = strings.Join(stringx.Remove(helpdocumenttranslationFieldNames, "`helpDocumentTranslation_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	helpdocumenttranslationModel interface {
		Insert(ctx context.Context, data *Helpdocumenttranslation) (sql.Result, error)
		FindOne(ctx context.Context, helpDocumentTranslationId int64) (*Helpdocumenttranslation, error)
		FindOneByHelpDocumentIdLanguageCode(ctx context.Context, helpDocumentId int64, languageCode string) (*Helpdocumenttranslation, error)
		Update(ctx context.Context, data *Helpdocumenttranslation) error
		Delete(ctx context.Context, helpDocumentTranslationId int64) error
		DeleteByHelpDocumentIdLanguageCode(ctx context.Context, helpDocumentId int64, languageCode string) error
	}

	defaultHelpdocumenttranslationModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Helpdocumenttranslation struct {
		HelpDocumentTranslationId int64        `db:"helpDocumentTranslation_id"`
		CreatedAt                 time.Time    `db:"created_at"`
		UpdatedAt                 sql.NullTime `db:"updated_at"`
		DeletedAt                 sql.NullTime `db:"deleted_at"`
		HelpDocumentId            int64        `db:"helpDocument_id"`
		LanguageCode              string       `db:"language_code"`
		DocumentTitle             string       `db:"document_title"`
		DocumentContent           string       `db:"document_content"`
	}
)

func newHelpdocumenttranslationModel(conn sqlx.SqlConn) *defaultHelpdocumenttranslationModel {
	return &defaultHelpdocumenttranslationModel{
		conn:  conn,
		table: "`helpdocumenttranslation`",
	}
}

func (m *defaultHelpdocumenttranslationModel) Delete(ctx context.Context, helpDocumentTranslationId int64) error {
	query := fmt.Sprintf("delete from %s where `helpDocumentTranslation_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, helpDocumentTranslationId)
	return err
}

func (m *defaultHelpdocumenttranslationModel) DeleteByHelpDocumentIdLanguageCode(ctx context.Context, helpDocumentId int64, languageCode string) error {
	query := fmt.Sprintf("delete from %s where `helpDocument_id` = ? and `language_code` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, helpDocumentId, languageCode)
	return err
}


func (m *defaultHelpdocumenttranslationModel) FindOne(ctx context.Context, helpDocumentTranslationId int64) (*Helpdocumenttranslation, error) {
	query := fmt.Sprintf("select %s from %s where `helpDocumentTranslation_id` = ? limit 1", helpdocumenttranslationRows, m.table)
	var resp Helpdocumenttranslation
	err := m.conn.QueryRowCtx(ctx, &resp, query, helpDocumentTranslationId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultHelpdocumenttranslationModel) FindOneByHelpDocumentIdLanguageCode(ctx context.Context, helpDocumentId int64, languageCode string) (*Helpdocumenttranslation, error) {
	var resp Helpdocumenttranslation
	query := fmt.Sprintf("select %s from %s where `helpDocument_id` = ? and `language_code` = ? limit 1", helpdocumenttranslationRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, helpDocumentId, languageCode)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultHelpdocumenttranslationModel) Insert(ctx context.Context, data *Helpdocumenttranslation) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, helpdocumenttranslationRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.HelpDocumentTranslationId, data.DeletedAt, data.HelpDocumentId, data.LanguageCode, data.DocumentTitle, data.DocumentContent)
	return ret, err
}

func (m *defaultHelpdocumenttranslationModel) Update(ctx context.Context, newData *Helpdocumenttranslation) error {
	query := fmt.Sprintf("update %s set %s where `helpDocumentTranslation_id` = ?", m.table, helpdocumenttranslationRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.DeletedAt, newData.HelpDocumentId, newData.LanguageCode, newData.DocumentTitle, newData.DocumentContent, newData.HelpDocumentTranslationId)
	return err
}

func (m *defaultHelpdocumenttranslationModel) tableName() string {
	return m.table
}
