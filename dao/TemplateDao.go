package dao

import (
	"dev-producer/model"
	"dev-producer/tool"
	"fmt"
)

type TemplateDao struct {
	*tool.Orm
}

//实例化Dao对象
func NewTemplateDao() *TemplateDao {
	return &TemplateDao{tool.DbEngine}
}

//从数据库中查询所有模版
func (Td *TemplateDao) QueryTemplates(daoPage *model.DaoPage) ([]model.Template, error) {
	var templates []model.Template
	if nil == daoPage {
		if err := Td.Engine.Find(&templates); err != nil {
			return nil, err
		}
	} else {
		if err := Td.Engine.Where("").Limit(daoPage.Pagenum, daoPage.Pagesize).Find(&templates); err != nil {
			return nil, err
		}
	}
	return templates, nil
}

//新发布模版的数据库插入操作
func (mid *TemplateDao) InsertTemplate(virtualMachine model.Template) int64 {
	result, err := mid.InsertOne(&virtualMachine)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return result
}

//查询发布模版是否存在
func (mid *TemplateDao) QueryByTemplate(mi model.Template) model.Template {
	var Template model.Template
	if _, err := mid.Where(" TemplateCode   = ? ", mi.TemplateCode).Get(&Template); err != nil {
		fmt.Println(err.Error())
	}
	return Template
}

//查询发布模版是否存在
func (mid *TemplateDao) QueryByIdTemplate(mi model.Template) model.Template {
	var Template model.Template
	if _, err := mid.Where(" Id  = ? ", mi.Id).Get(&Template); err != nil {
		fmt.Println(err.Error())
	}
	return Template
}

//删除发布模版
func (mid *TemplateDao) DeleteTemplate(mi model.Template) int64 {

	if _, err := mid.Where("  Id  = ? ", mi.Id).Delete(mi); err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return 1
}

//更新发布模版
func (mid *TemplateDao) UpdateTemplate(mi model.Template) int64 {

	if result, err := mid.Where("  Id  = ? ", mi.Id).Update(mi); err != nil {
		fmt.Println(err.Error(), result)
		return 0
	}
	return 1
}
