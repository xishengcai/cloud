package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/xishengcai/cloud/docs"
	"github.com/xishengcai/cloud/pkg/middleware"
	"github.com/xishengcai/cloud/routers/api/cluster"
	"github.com/xishengcai/cloud/routers/api/image"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	baseRoute := r.Group("/api/v1/")
	cluster.Register(baseRoute)
	image.Register(baseRoute)

	agg := r.Group("/")
	agg.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
