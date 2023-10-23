package cluster

import "github.com/gin-gonic/gin"

func Register(baseGroup *gin.RouterGroup) {
	router := baseGroup.Group("/cluster")
	//router.GET("/",List)
	router.GET("", List)
	router.POST("", Install)
	router.POST("/slaves", InstallKubernetesSlave)
	router.POST("/joinMaster", JoinMaster)
	router.POST("/upgrade", Upgrade)
}
