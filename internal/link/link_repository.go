package link

import (
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {
	return &repository{db: db}
}

func (r *repository) CreateLink(ctx context.Context, link *Link) (*Link, error) {
	var persistedLink Link

	query := "INSERT INTO links(url, code) VALUES ($1, $2) RETURNING id, url, code, created_at"
	err := r.db.QueryRowContext(ctx, query, link.URL, link.Code).
		Scan(
			&persistedLink.ID,
			&persistedLink.URL,
			&persistedLink.Code,
			&persistedLink.CreatedAt,
		)

	if err != nil {
		return &Link{}, err
	}

	return &persistedLink, nil
}

func (r *repository) GetLinkByCode(ctx context.Context, code string) (*Link, error) {
	link := Link{}

	query := "SELECT id, url, code, created_at FROM links WHERE code = $1"
	err := r.db.QueryRowContext(ctx, query, code).
		Scan(
			&link.ID,
			&link.URL,
			&link.Code,
			&link.CreatedAt,
		)

	if err != nil {
		return &Link{}, err
	}

	return &link, nil
}
