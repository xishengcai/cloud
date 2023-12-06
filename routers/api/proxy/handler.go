package proxy

import (
	"github.com/xishengcai/cloud/pkg/app"
	"github.com/xishengcai/cloud/service/proxy"

	"github.com/gin-gonic/gin"
)

// Install godoc
// @Summary install proxy
// @Description install proxy
// @Tags proxy
// @Accept  json
// @Produce  json
// @Param cluster body proxy.Install true "install proxy"
// @Failure 201 {object} app.Response
// @Router /v1/proxy [post]
func Install(ctx *gin.Context) {
	k := &proxy.Install{}
	app.HandleOperator(ctx, k, func(o app.Operator) {
		app.HandleServiceResultWithCode(ctx, 201, o.Run())
	})
}
