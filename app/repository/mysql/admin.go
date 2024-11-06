package mysqlrepo

import (
	"blog-mandalika/domain"
	"context"
	"log"
)

func (r *dbRepo) FetchOneAdmin(ctx context.Context, options map[string]interface{}) (domain.Admin, error) {
	sql := "SELECT * FROM admin"
	model := domain.Admin{}
	query := model.GenerateQuery(sql, options)
	row := r.Conn.QueryRowContext(ctx, query)
	if row.Err() != nil {
		log.Println("FetchOne Admin", row.Err())
		return domain.Admin{}, row.Err()
	}

	err := row.Scan(&model.ID, &model.Email, &model.Password, &model.CreatedAt, &model.UpdatedAt, &model.DeletedAt)
	if err != nil {
		return domain.Admin{}, err
	}
	return model, nil
}
