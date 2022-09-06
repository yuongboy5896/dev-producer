package service

import (
	"dev-producer/dao"
	"dev-producer/model"
)

type TemplateService struct {
}

/**
 * 获取获取 Template
 */
func (mis *TemplateService) Templatelist(daoPage *model.DaoPage) ([]model.Template, error) {
	//数据库操作层
	templateDao := dao.NewTemplateDao()
	return templateDao.QueryTemplates(daoPage)
}

/*
* 添加发布模版
 */
func (mis *TemplateService) AddTemplate(Template model.Template) int64 {

	vmD := dao.NewTemplateDao()

	result := vmD.InsertTemplate(Template)

	return result
}

/*
* 查询发布模版
 */
func (mis *TemplateService) GetTemplate(Template model.Template) model.Template {

	vmD := dao.NewTemplateDao()

	result := vmD.QueryByTemplate(Template)

	return result
}

/*
* 查询发布模版根据
 */
func (mis *TemplateService) GetTemplateById(Template model.Template) model.Template {

	vmD := dao.NewTemplateDao()

	result := vmD.QueryByIdTemplate(Template)

	return result
}

/*
* 删除发布模版
 */
func (mis *TemplateService) DeleteTemplate(Template model.Template) int64 {

	vmD := dao.NewTemplateDao()

	result := vmD.DeleteTemplate(Template)

	return result
}

/*
* 更新发布模版
 */
func (mis *TemplateService) UpdateTemplate(Template model.Template) int64 {

	vmD := dao.NewTemplateDao()

	result := vmD.UpdateTemplate(Template)

	return result
}
