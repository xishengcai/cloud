package image

import "github.com/gin-gonic/gin"

func Register(baseGroup *gin.RouterGroup) {
	router := baseGroup.Group("images/")
	router.POST("/pull", pull)
	router.GET("/info", info)
}
