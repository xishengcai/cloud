package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"k8s.io/klog/v2"
)

type Operator interface {
	Run() ResultRaw
	Validate() error
}

type Action func(p Operator)

// HandleOperator 操作接口operator
func HandleOperator(ctx *gin.Context, o Operator, action Action) {
	if err := Bind(ctx, o); err != nil {
		return
	}
	action(o)
}

func HandleParameterBindError(ctx *gin.Context, err error) {
	klog.Errorf("参数绑定失败:%v", err)

	ctx.JSON(http.StatusOK, Response{
		ResMsg: "参数解析失败: " + err.Error(),
	})
}

func Bind(ctx *gin.Context, operator Operator) error {
	errReturn := func(err error) error {
		HandleParameterBindError(ctx, err)
		return err
	}

	var err error
	// header parameter
	err = ctx.ShouldBindHeader(operator)
	if err != nil {
		return errReturn(err)
	}

	// url query parameter
	err = ctx.ShouldBindQuery(operator)
	if err != nil {
		return errReturn(err)
	}

	// url middle path parameter
	err = ctx.ShouldBindUri(operator)
	if err != nil {
		return errReturn(err)
	}

	// json body parameter
	// Get请求中如果没有json body 会报EOF错误
	// 判断如果没有form就忽略
	// 坑：https://studygolang.com/articles/17745
	err = ctx.ShouldBind(operator)
	if err != nil && err.Error() != "EOF" {
		return errReturn(err)
	}

	val := validator.New()
	if err := val.Struct(operator); err != nil {
		return errReturn(err)
	}

	return nil
}
