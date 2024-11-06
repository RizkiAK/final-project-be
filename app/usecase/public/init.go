package public

import (
	"context"
	"time"

	mysqlrepo "blog-mandalika/app/repository/mysql"

	"github.com/Yureka-Teknologi-Cipta/yureka/response"
)

type publicUsecase struct {
	mysqlRepo      mysqlrepo.DBRepo
	contextTimeout time.Duration
}

type RepoInjection struct {
	MysqlRepo mysqlrepo.DBRepo
}

func NewAppPublicUsecase(r RepoInjection, timeout time.Duration) PublicUsecase {
	return &publicUsecase{
		mysqlRepo:      r.MysqlRepo,
		contextTimeout: timeout,
	}
}

type PublicUsecase interface {
	ListBlog(ctx context.Context, options map[string]interface{}) response.Base
	ListFAQ(ctx context.Context, options map[string]interface{}) response.Base
}
