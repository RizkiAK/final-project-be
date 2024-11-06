package mysqlrepo

import (
	"blog-mandalika/domain"
	"context"
	"database/sql"
	"log"
)

func (r *dbRepo) FetchFAQ(ctx context.Context, options map[string]interface{}) (*sql.Rows, error) {
	sql := "SELECT * FROM faq"
	model := domain.Faq{}
	query := model.GenerateQuery(sql, options)
	rows, err := r.Conn.QueryContext(ctx, query)
	if err != nil {
		log.Println("Fetch faq", err)
		return nil, err
	}

	return rows, nil
}

func (r *dbRepo) CountFAQ(ctx context.Context, options map[string]interface{}) int64 {
	var count int64
	sql := "SELECT COUNT(*) FROM faq"
	model := domain.Faq{}
	query := model.GenerateQuery(sql, options)
	err := r.Conn.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		log.Println("Count FAQ", err)
		return 0
	}

	return count
}
