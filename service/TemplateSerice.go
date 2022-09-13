package service

import (
	"dev-producer/dao"
	"dev-producer/model"
)

type TemplateInfoService struct {
}

/**
 * 获取获取 TemplateInfo
 */
func (mis *TemplateInfoService) TemplateInfolist(daoPage *model.DaoPage) ([]model.TemplateInfo, error) {
	//数据库操作层
	TemplateInfoDao := dao.NewTemplateInfoDao()
	return TemplateInfoDao.QueryTemplateInfos(daoPage)
}

/*
* 添加发布模版
 */
func (mis *TemplateInfoService) AddTemplateInfo(TemplateInfo model.TemplateInfo) int64 {

	vmD := dao.NewTemplateInfoDao()

	result := vmD.InsertTemplateInfo(TemplateInfo)

	return result
}

/*
* 查询发布模版
 */
func (mis *TemplateInfoService) GetTemplateInfo(TemplateInfo model.TemplateInfo) model.TemplateInfo {

	vmD := dao.NewTemplateInfoDao()

	result := vmD.QueryByTemplateInfo(TemplateInfo)

	return result
}

/*
* 查询发布模版根据
 */
func (mis *TemplateInfoService) GetTemplateInfoById(TemplateInfo model.TemplateInfo) model.TemplateInfo {

	vmD := dao.NewTemplateInfoDao()

	result := vmD.QueryByIdTemplateInfo(TemplateInfo)

	return result
}

/*
* 删除发布模版
 */
func (mis *TemplateInfoService) DeleteTemplateInfo(TemplateInfo model.TemplateInfo) int64 {

	vmD := dao.NewTemplateInfoDao()

	result := vmD.DeleteTemplateInfo(TemplateInfo)

	return result
}

/*
* 更新发布模版
 */
func (mis *TemplateInfoService) UpdateTemplateInfo(TemplateInfo model.TemplateInfo) int64 {

	vmD := dao.NewTemplateInfoDao()

	result := vmD.UpdateTemplateInfo(TemplateInfo)

	return result
}
