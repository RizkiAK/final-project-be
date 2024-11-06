package mysqlrepo

import (
	"blog-mandalika/domain"
	"context"
	"database/sql"
)

type dbRepo struct {
	Conn       *sql.DB
	AdminTable string
	BlogTable  string
	FaqTable   string
}

func NewDBRepo(Conn *sql.DB) DBRepo {
	return &dbRepo{
		Conn:       Conn,
		AdminTable: "admin",
		BlogTable:  "blogs",
		FaqTable:   "faq",
	}
}

type DBRepo interface {
	// admin
	FetchOneAdmin(ctx context.Context, options map[string]interface{}) (domain.Admin, error)

	// blog
	FetchBlog(ctx context.Context, options map[string]interface{}) (*sql.Rows, error)
	CountBlog(ctx context.Context, options map[string]interface{}) int64
	CreateBlog(ctx context.Context, payload domain.Blog) error

	// FAQ
	FetchFAQ(ctx context.Context, options map[string]interface{}) (*sql.Rows, error)
	CountFAQ(ctx context.Context, options map[string]interface{}) int64
}
