package dao

import (
	"dev-producer/model"
	"dev-producer/tool"
	"fmt"
)

type TemplateInfoDao struct {
	*tool.Orm
}

//实例化Dao对象
func NewTemplateInfoDao() *TemplateInfoDao {
	return &TemplateInfoDao{tool.DbEngine}
}

//从数据库中查询所有模版
func (Td *TemplateInfoDao) QueryTemplateInfos(daoPage *model.DaoPage) ([]model.TemplateInfo, error) {
	var TemplateInfos []model.TemplateInfo
	if nil == daoPage {
		if err := Td.Engine.Find(&TemplateInfos); err != nil {
			return nil, err
		}
	} else {
		if err := Td.Engine.Where("").Limit(daoPage.Pagenum, daoPage.Pagesize).Find(&TemplateInfos); err != nil {
			return nil, err
		}
	}
	return TemplateInfos, nil
}

//新发布模版的数据库插入操作
func (mid *TemplateInfoDao) InsertTemplateInfo(virtualMachine model.TemplateInfo) int64 {
	result, err := mid.InsertOne(&virtualMachine)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return result
}

//查询发布模版是否存在
func (mid *TemplateInfoDao) QueryByTemplateInfo(mi model.TemplateInfo) model.TemplateInfo {
	var TemplateInfo model.TemplateInfo
	if _, err := mid.Where(" TemplateCode   = ? ", mi.TemplateCode).Get(&TemplateInfo); err != nil {
		fmt.Println(err.Error())
	}
	return TemplateInfo
}

//查询发布模版是否存在
func (mid *TemplateInfoDao) QueryByIdTemplateInfo(mi model.TemplateInfo) model.TemplateInfo {
	var TemplateInfo model.TemplateInfo
	if _, err := mid.Where(" Id  = ? ", mi.Id).Get(&TemplateInfo); err != nil {
		fmt.Println(err.Error())
	}
	return TemplateInfo
}

//删除发布模版
func (mid *TemplateInfoDao) DeleteTemplateInfo(mi model.TemplateInfo) int64 {

	if _, err := mid.Where("  Id  = ? ", mi.Id).Delete(mi); err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return 1
}

//更新发布模版
func (mid *TemplateInfoDao) UpdateTemplateInfo(mi model.TemplateInfo) int64 {

	if result, err := mid.Where("  Id  = ? ", mi.Id).Cols("TemplateName", "TemplateCode", "TemplateType", "TemplateText", "ReplaceText", "TemplateJekins").Update(mi); err != nil {
		fmt.Println(err.Error(), result)
		return 0
	}
	return 1
}
