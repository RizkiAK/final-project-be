package http

import (
	"blog-mandalika/app/delivery/http/middleware"

	"github.com/gin-gonic/gin"
)

func (handler *routeHandler) handlerPublicRoute(m *middleware.AppMiddleware, route *gin.RouterGroup) {
	blog := route.Group("/blog")
	blog.GET("/", handler.ListBlog)
	blog.GET("/admin", m.AdminAuth(), handler.ListBlog)

	faq := route.Group("/faq")
	faq.GET("/", handler.ListFAQ)
}

func (r *routeHandler) ListBlog(c *gin.Context) {
	ctx := c.Request.Context()

	options := map[string]interface{}{
		"query": c.Request.URL.Query(),
	}

	response := r.PublicUsecase.ListBlog(ctx, options)
	c.JSON(response.Status, response)
}

func (r *routeHandler) ListFAQ(c *gin.Context) {
	ctx := c.Request.Context()

	options := map[string]interface{}{
		"query": c.Request.URL.Query(),
	}

	response := r.PublicUsecase.ListFAQ(ctx, options)
	c.JSON(response.Status, response)
}
