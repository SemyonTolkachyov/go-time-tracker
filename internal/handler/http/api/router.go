package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "go-time-tracker/api"
	"go-time-tracker/internal/middleware"
)

type HandlerRouter interface {
	AddRoutes(r *gin.RouterGroup)
	GetVersion() string
}

type Router struct {
	router *gin.Engine
}

func NewRouter() *Router {
	return &Router{router: gin.New()}
}

func (r *Router) WithSwagger() *Router {
	r.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

func (r *Router) WithHandler(h HandlerRouter) *Router {
	api := r.router.Group("/api/" + h.GetVersion())

	api.Use(gin.Recovery())
	api.Use(middleware.LoggingMiddleware())

	h.AddRoutes(api)

	return r
}
