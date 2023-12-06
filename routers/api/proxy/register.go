package proxy

import "github.com/gin-gonic/gin"

func Register(baseGroup *gin.RouterGroup) {
	router := baseGroup.Group("/proxy")
	//router.GET("/",List)
	router.POST("", Install)
}
