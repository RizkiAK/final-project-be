package domain

import (
	"fmt"
	"time"
)

type Faq struct {
	ID        string     `json:"id"`
	Question  string     `json:"question"`
	Answer    string     `json:"answer"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (e *Faq) GenerateQuery(sql string, options map[string]interface{}) string {
	query := " WHERE deleted_at IS NULL "

	if id, ok := options["id"].(string); ok {
		query = query + "AND id = '" + id + "' "
	}

	if q, ok := options["q"].(string); ok {
		query = query + "AND question LIKE '%" + q + "%' "
	}

	// pagination & sorting
	offset, ok := options["offset"].(int64)
	if !ok {
		offset = 0
	}
	limit, ok := options["limit"].(int64)
	if !ok {
		limit = 10
	}
	sortBy, ok := options["sort"].(string)
	if !ok {
		sortBy = "created_at"
	}
	sortDir, ok := options["dir"].(string)
	if !ok {
		sortDir = "desc"
	}

	if sortBy != "" {
		query = query + "ORDER BY " + sortBy + " " + sortDir + " "
	}
	if limit != -1 {
		query = query + "LIMIT " + fmt.Sprint(limit) + " "
		query = query + "OFFSET " + fmt.Sprint(offset) + " "
	}
	query = sql + query

	return query
}
