package cluster

import (
	"github.com/xishengcai/cloud/models"
	"github.com/xishengcai/cloud/pkg/app"
	"github.com/xishengcai/cloud/service/kubernetes"

	"github.com/gin-gonic/gin"
)

// InstallKubernetes godoc
// @Summary install cluster
// @Description install cluster master
// @Tags install cluster master
// @Accept  json
// @Produce  json
// @Param cluster body models.Kubernetes true "install cluster master"
// @Failure 201 {object} app.Response
// @Router /v1/cluster/masters [post]
func InstallKubernetes(ctx *gin.Context) {
	k := &kubernetes.InstallKuber{}
	app.HandleOperator(ctx, k, func(o app.Operator) {
		app.HandleServiceResult(ctx, 201, o.Run())
	})
}

// InstallKubernetesSlave godoc
// @Summary install cluster slave
// @Description install cluster slave
// @Tags install cluster slave
// @Accept  json
// @Produce  json
// @Param cluster body InstallSlave true "install cluster slave"
// @Failure 201 {object} app.Response
// @Router /v1/cluster/slaves [post]
func InstallKubernetesSlave(ctx *gin.Context) {
	k := &kubernetes.InstallSlave{}
	app.HandleOperator(ctx, k, func(o app.Operator) {
		app.HandleServiceResult(ctx, 201, o.Run())
	})
}

// InstallSlave batch join slave to  k8s
type InstallSlave struct {
	Nodes   []models.Host `json:"nodes"`
	Master  models.Host   `json:"master"`
	Version string        `json:"version"`
	DryRun  bool          `json:"dryRun,omitempty"`
}

func JoinMaster(ctx *gin.Context) {
	k := &kubernetes.JoinMaster{}
	app.HandleOperator(ctx, k, func(o app.Operator) {
		app.HandleServiceResult(ctx, 201, o.Run())
	})
}

// Upgrade godoc
// @Summary upgrade k8s
// @Description install cluster slave
// @Tags upgrade
// @Accept  json
// @Produce  json
// @Param upgrade body  kubernetes.Upgrade true  "k8s all nodes"
// @Failure 201 {object} app.Response
// @Router /v1/cluster/upgrade [post]
func Upgrade(ctx *gin.Context) {
	u := &kubernetes.Upgrade{}
	app.HandleOperator(ctx, u, func(o app.Operator) {
		app.HandleServiceResult(ctx, 200, o.Run())
	})
}
