package mysqlrepo

import (
	"blog-mandalika/domain"
	"context"
	"database/sql"
	"log"
)

func (r *dbRepo) FetchBlog(ctx context.Context, options map[string]interface{}) (*sql.Rows, error) {
	sql := "SELECT * FROM blogs"
	model := domain.Faq{}
	query := model.GenerateQuery(sql, options)
	rows, err := r.Conn.QueryContext(ctx, query)
	if err != nil {
		log.Println("Fetch blogs", err)
		return nil, err
	}

	return rows, nil
}

func (r *dbRepo) CreateBlog(ctx context.Context, payload domain.Blog) error {
	sql := "INSERT INTO blogs (id, title, content, media_url, category, created_at) VALUES (?, ?, ?, ?, ?, NOW())"
	row := r.Conn.QueryRowContext(ctx, sql, payload.ID, payload.Title, payload.Content, payload.MediaUrl, payload.Category)

	if row.Err() != nil {
		log.Println("Create blogs", row.Err())
		return row.Err()
	}

	return nil
}

func (r *dbRepo) CountBlog(ctx context.Context, options map[string]interface{}) int64 {
	var count int64
	sql := "SELECT COUNT(*) FROM blogs"
	model := domain.Faq{}
	query := model.GenerateQuery(sql, options)
	err := r.Conn.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		log.Println("Count blogs", err)
		return 0
	}

	return count
}
