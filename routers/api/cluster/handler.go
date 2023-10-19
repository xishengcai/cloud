package cluster

import (
	"github.com/xishengcai/cloud/models"
	"github.com/xishengcai/cloud/pkg/app"
	"github.com/xishengcai/cloud/service/kubernetes"

	"github.com/gin-gonic/gin"
)

// Install godoc
// @Summary install cluster
// @Description install cluster master
// @Tags k8s cluster
// @Accept  json
// @Produce  json
// @Param cluster body clusterParam true "install cluster"
// @Failure 201 {object} app.Response
// @Router /v1/cluster [post]
func Install(ctx *gin.Context) {
	k := &kubernetes.Cluster{}
	app.HandleOperator(ctx, k, func(o app.Operator) {
		app.HandleServiceResultWithCode(ctx, 201, o.Run())
	})
}

// List godoc
// @Summary list cluster
// @Description list cluster
// @Tags k8s cluster
// @Accept  json
// @Produce  json
// @param   page   query  int    false "page number, optional"
// @param   pageSize  query  int     false  "page size, optional"
// @Failure 200 {object} app.Response
// @Router /v1/cluster [get]
func List(ctx *gin.Context) {
	k := &kubernetes.List{}
	app.HandleOperator(ctx, k, func(o app.Operator) {
		app.HandleServiceResult(ctx, o.Run())
	})
}

// InstallKubernetesSlave godoc
// @Summary install cluster slave
// @Description install cluster slave
// @Tags k8s cluster
// @Accept  json
// @Produce  json
// @Param cluster body InstallSlave true "install cluster slave"
// @Failure 201 {object} app.Response
// @Router /v1/cluster/slaves [post]
func InstallKubernetesSlave(ctx *gin.Context) {
	k := &kubernetes.InstallSlave{}
	app.HandleOperator(ctx, k, func(o app.Operator) {
		app.HandleServiceResult(ctx, o.Run())
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
		app.HandleServiceResult(ctx, o.Run())
	})
}

// Upgrade godoc
// @Summary upgrade k8s
// @Description install cluster slave
// @Tags k8s cluster
// @Accept  json
// @Produce  json
// @Param upgrade body  kubernetes.Upgrade true  "k8s all nodes"
// @Failure 201 {object} app.Response
// @Router /v1/cluster/upgrade [post]
func Upgrade(ctx *gin.Context) {
	u := &kubernetes.Upgrade{}
	app.HandleOperator(ctx, u, func(o app.Operator) {
		app.HandleServiceResult(ctx, o.Run())
	})
}
