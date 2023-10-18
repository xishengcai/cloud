package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	EventCode int         `json:"eventCode"`
	ResMsg    interface{} `json:"resMsg"`
	Data      interface{} `json:"data"`
}

type ListData struct {
	Total int64       `json:"total"`
	Items interface{} `json:"items"`
}

type ResultRaw interface {
	GetResponse() Response
	ResMsg() string
	HasError() bool
	Error() error
	HandleData(opt func(data interface{}) interface{})
}

var _ ResultRaw = &ServiceResult{}

type ServiceResult struct {
	data interface{}
	err  error
}

func NewServiceResult(data interface{}, err error) *ServiceResult {
	return &ServiceResult{
		data: data,
		err:  err,
	}
}

func NewServiceResultWitRawData(data interface{}, err error) *ServiceResult {
	// check data. Avoid null-pointer exceptions
	if data == nil {
		data = &ListData{
			Total: 0,
			Items: []interface{}{},
		}
	}

	return &ServiceResult{
		data: data,
		err:  err,
	}
}

func NewServiceResultWithArray(i *InterfaceArray, err error) *ServiceResult {
	// check data. Avoid null-pointer exceptions
	data := i.convertToDataList()
	return NewServiceResultWitRawData(data, err)
}

func (s *ServiceResult) HandleData(opt func(data interface{}) interface{}) {
	listData := s.data.(ListData)
	items := opt(listData.Items)
	s.data = ListData{listData.Total, items}
}

func (s *ServiceResult) ResMsg() string {
	if s.HasError() {
		return s.err.Error()
	}
	return ""
}

func (s *ServiceResult) HasError() bool {
	return s.err != nil
}

func (s *ServiceResult) Error() error {
	return s.err
}

func (s *ServiceResult) Check() error {
	return s.err
}

func (s *ServiceResult) GetResponse() Response {
	response := Response{Data: s.data}
	if s.HasError() {
		response.ResMsg = s.ResMsg()
		response.EventCode = 1
	} else {
		response.ResMsg = "操作成功"
	}
	return response
}

func HandleServiceResultWithCode(ctx *gin.Context, status int, rr ResultRaw) {
	ctx.JSON(status, rr.GetResponse())
}

func HandleServiceResult(ctx *gin.Context, rr ResultRaw) {
	res := rr.GetResponse()
	ctx.JSON(http.StatusOK, res)
}

type InterfaceArray struct {
	Total int
	Items []interface{}
}

func (i InterfaceArray) convertToDataList() *ListData {
	return &ListData{
		Total: int64(i.Total),
		Items: i.Items,
	}
}
