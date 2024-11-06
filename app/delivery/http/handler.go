package http

import (
	"blog-mandalika/app/delivery/http/middleware"
	"blog-mandalika/app/usecase/admin"
	"blog-mandalika/app/usecase/public"

	"github.com/gin-gonic/gin"
)

type routeHandler struct {
	AdminUsecase  admin.AdminUsecase
	PublicUsecase public.PublicUsecase
}

func NewRouteHandler(ginEngine *gin.Engine, admin admin.AdminUsecase, public public.PublicUsecase) {
	middle := middleware.InitMiddleware()
	handler := &routeHandler{
		AdminUsecase:  admin,
		PublicUsecase: public,
	}

	api := ginEngine.Group("/v1/mandalika")

	adminAPI := api.Group("/admin")
	publicAPI := api.Group("/public")

	handler.handlerAdminRoute(middle, adminAPI)
	handler.handlerPublicRoute(middle, publicAPI)
}
