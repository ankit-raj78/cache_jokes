package api

import (
	"github.com/gin-gonic/gin"
)

// NewRouter builds the HTTP router with middleware and routes.
func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	v1 := r.Group("/v1")
	{
		v1.GET("/health", Health)
		v1.POST("/recommend", Recommend)
		v1.POST("/provision", Provision)
	}
	return r
}
