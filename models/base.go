package models

type BaseInformer interface {
	Validate()
}

type BaseInfo struct {
	Token string `json:"token"`
	*PageInfo
}

type PageInfo struct {
	Page     int `form:"page,default=1"`
	PageSize int `form:"pageSize,default=10"`
}

func (p PageInfo) GetOffset() int {
	return (p.Page - 1) * p.PageSize
}

func (baseInfo *BaseInfo) SetPageInfo(page, pageSize int) {
	baseInfo.Page = page
	baseInfo.PageSize = pageSize
}

type ShortAccountInfo struct {
	KeystoneUserIDSub   string `json:"keystoneUserIdSub"`
	KeystoneUserNameSub string `json:"keystoneUserNameSub"`
}

type SubUserResponse struct {
	Status int    `json:"status"`
	ResMsg string `json:"resMsg"`
	Data   struct {
		Total int                `json:"total"`
		Items []ShortAccountInfo `json:"items"`
	} `json:"data"`
}
