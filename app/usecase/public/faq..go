package public

import (
	"blog-mandalika/domain"
	"context"
	"log"
	"math"
	"net/http"
	"net/url"

	"github.com/Yureka-Teknologi-Cipta/yureka/response"

	yurekahelpers "github.com/Yureka-Teknologi-Cipta/yureka/helpers"
)

func (u *publicUsecase) ListFAQ(ctx context.Context, options map[string]interface{}) response.Base {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	query := options["query"].(url.Values)

	page, limit, offset := yurekahelpers.GetLimitOffset(query)

	fetchOptions := map[string]interface{}{
		"limit":  limit,
		"offset": offset,
	}

	// filtering
	if query.Get("sort") != "" {
		fetchOptions["sort"] = query.Get("sort")
	}

	if query.Get("dir") != "" {
		fetchOptions["dir"] = query.Get("dir")
	}

	if query.Get("q") != "" {
		fetchOptions["q"] = query.Get("q")
	}

	// count first
	counts := u.mysqlRepo.CountFAQ(ctx, fetchOptions)

	if counts == 0 {
		return response.Success(map[string]interface{}{
			"List":      []interface{}{},
			"Limit":     limit,
			"Page":      page,
			"TotalData": counts,
			"TotalPage": math.Ceil(float64(counts) / float64(limit)),
		})
	}

	// check ticket list
	cur, err := u.mysqlRepo.FetchFAQ(ctx, fetchOptions)

	if err != nil {
		return response.Success(map[string]interface{}{
			"List":      []interface{}{},
			"Limit":     limit,
			"Page":      page,
			"TotalData": counts,
			"TotalPage": math.Ceil(float64(counts) / float64(limit)),
		})
	}

	defer cur.Close()

	list := make([]interface{}, 0)
	for cur.Next() {
		row := domain.Faq{}
		err := cur.Scan(
			&row.ID, &row.Question, &row.Answer, &row.CreatedAt, &row.UpdatedAt, &row.DeletedAt,
		)
		if err != nil {
			log.Println("FetchFAQ Scan ", err)
			return response.Error(http.StatusBadRequest, err.Error())
		}
		list = append(list, row)
	}

	return response.Success(map[string]interface{}{
		"List":      list,
		"Limit":     limit,
		"Page":      page,
		"TotalData": counts,
		"TotalPage": math.Ceil(float64(counts) / float64(limit)),
	})
}
