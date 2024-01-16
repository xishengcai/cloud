package image

import (
	"github.com/gin-gonic/gin"

	"github.com/xishengcai/cloud/pkg/app"
	"github.com/xishengcai/cloud/service/images"
)

// pull godoc
// @Summary image push to oss
// @Description image push to oss
// @Tags image
// @Accept  json
// @Produce  json
// @Param cluster body images.Pull true "pull Image, then push to OSS"
// @Failure 201 {object} app.Response
// @Router /v1/images/pull [post]
func pull(ctx *gin.Context) {
	k := &images.Pull{}
	app.HandleOperator(ctx, k, func(o app.Operator) {
		app.HandleServiceResultWithCode(ctx, 201, o.Run())
	})
}

// info godoc
// @Summary image
// @Description image pull info
// @Tags image
// @Accept  json
// @Produce  json
// @Failure 200 {object} app.Response
// @Router /v1/images/info [get]
func info(ctx *gin.Context) {
	k := &images.Info{}
	app.HandleOperator(ctx, k, func(o app.Operator) {
		app.HandleServiceResultWithCode(ctx, 201, o.Run())
	})
}
