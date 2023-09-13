package models

import (
	"gorm.io/gorm"

	"github.com/xishengcai/cloud/pkg/setting"
)

type IDs struct {
	IDs []string `json:"ids"`
}

// RestPage 分页查询
// page  设置起始页、每页条数,
// name  查询目标表的名称
// query 查询条件,
// dest  查询结果绑定的结构体,
// bind  绑定表结构对应的结构体
func restPage(page Page, name string, query interface{}, dest interface{}, bind interface{}) (int64, error) {
	if page.PageNum > 0 && page.PageSize > 0 {
		offset := (page.PageNum - 1) * page.PageSize
		setting.DB.Offset(offset).Limit(page.PageSize).Table(name).Where(query).Order("created desc").Find(dest)
	}
	res := setting.DB.Table(name).Where(query).Order("created desc").Find(bind)
	return res.RowsAffected, res.Error
}

type Page struct {
	PageNum  int `form:"pageNum,default=1"`
	PageSize int `form:"pageSize,default=10"`
}

func (p Page) Offset() int {
	if p.PageNum < 1 {
		p.PageNum = 1
	}
	return (p.PageNum - 1) * p.PageSize

}

func (p Page) Limit() int {
	if p.PageSize <= 0 {
		return 10
	}
	return p.PageSize
}

func GetDB(tableName string) DBHelp {
	return DBHelp{
		Table: tableName,
		DB:    setting.DB.Table(tableName),
	}
}

type DBHelp struct {
	Table string
	*gorm.DB
}

func (db DBHelp) Find(id string, dest interface{}) (interface{}, error) {
	err := db.Where("id = ?", id).First(&dest).Error
	return dest, err
}

func (db DBHelp) Query(page Page, result interface{}) (res interface{}, total int64, err error) {
	err = db.Count(&total).
		Limit(page.PageSize).
		Offset(page.Offset()).
		Order("updated_time desc").
		Find(&result).Error
	res = result
	return
}

func (db DBHelp) Del(ids []string) error {
	return db.Where("id in ?", ids).Delete(db.Table).Error
}
