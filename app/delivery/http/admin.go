package http

import (
	"blog-mandalika/app/delivery/http/middleware"
	"blog-mandalika/domain"
	"net/http"

	"github.com/Yureka-Teknologi-Cipta/yureka/response"
	"github.com/gin-gonic/gin"
)

func (handler *routeHandler) handlerAdminRoute(m *middleware.AppMiddleware, route *gin.RouterGroup) {
	route.POST("/login", handler.Login)

	blog := route.Group("/blog")
	blog.POST("/create", m.AdminAuth(), handler.CreateBlog)
}

func (r *routeHandler) Login(c *gin.Context) {
	ctx := c.Request.Context()

	payload := domain.LoginAdminRequest{}
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Error(http.StatusBadRequest, "invalid json data"))
		return
	}

	options := map[string]interface{}{
		"payload": payload,
	}

	response := r.AdminUsecase.Login(ctx, options)
	c.JSON(response.Status, response)
}

func (r *routeHandler) CreateBlog(c *gin.Context) {
	ctx := c.Request.Context()

	adminID := c.MustGet("id").(string)

	payload := domain.CreateBlogRequest{}
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Error(http.StatusBadRequest, "invalid json data"))
		return
	}

	options := map[string]interface{}{
		"admin_id": adminID,
		"payload":  payload,
	}

	response := r.AdminUsecase.CreateBlog(ctx, options)
	c.JSON(response.Status, response)
}
