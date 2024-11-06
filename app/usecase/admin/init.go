package admin

import (
	"context"
	"time"

	mysqlrepo "blog-mandalika/app/repository/mysql"

	"github.com/Yureka-Teknologi-Cipta/yureka/response"
)

type adminUsecase struct {
	mysqlRepo      mysqlrepo.DBRepo
	contextTimeout time.Duration
}

type RepoInjection struct {
	MysqlRepo mysqlrepo.DBRepo
}

func NewAppAdminUsecase(r RepoInjection, timeout time.Duration) AdminUsecase {
	return &adminUsecase{
		mysqlRepo:      r.MysqlRepo,
		contextTimeout: timeout,
	}
}

type AdminUsecase interface {
	Login(ctx context.Context, options map[string]interface{}) response.Base
	CreateBlog(ctx context.Context, options map[string]interface{}) response.Base
}
