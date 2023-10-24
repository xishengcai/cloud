package cluster

import "github.com/gin-gonic/gin"

func Register(baseGroup *gin.RouterGroup) {
	router := baseGroup.Group("/cluster")
	//router.GET("/",List)
	router.GET("", List)
	router.POST("", Install)
	router.POST("/nodes", JoinNodes)
	router.POST("/upgrade", Upgrade)
}
