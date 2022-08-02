package dao

import (
	"dev-producer/model"
	"dev-producer/tool"
	"fmt"
)

type ModuleInfoDao struct {
	*tool.Orm
}

//实例化Dao对象
func NewModuleInfoDao() *ModuleInfoDao {
	return &ModuleInfoDao{tool.DbEngine}
}

//从数据库中查询所有服务器列表
func (mid *ModuleInfoDao) QueryModuleInfos(daoPage *model.DaoPage) ([]model.ModuleInfo, error) {
	var virtualmachines []model.ModuleInfo
	if nil == daoPage {
		if err := mid.Engine.Find(&virtualmachines); err != nil {
			return nil, err
		}
	} else {
		if err := mid.Engine.Where("").Limit(daoPage.Pagenum, daoPage.Pagesize).Find(&virtualmachines); err != nil {
			return nil, err
		}
	}
	return virtualmachines, nil
}

//新应用模块的数据库插入操作
func (mid *ModuleInfoDao) InsertModuleInfo(virtualMachine model.ModuleInfo) int64 {
	result, err := mid.InsertOne(&virtualMachine)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return result
}

//查询应用模块是否存在
func (mid *ModuleInfoDao) QueryByModuleInfos(mi model.ModuleInfo) model.ModuleInfo {
	var moduleInfo model.ModuleInfo
	if _, err := mid.Where(" ModuleCode  = ? ", mi.ModuleCode).Get(&moduleInfo); err != nil {
		fmt.Println(err.Error())
	}
	return moduleInfo
}

//查询应用模块是否存在
func (mid *ModuleInfoDao) QueryByIdModuleInfos(mi model.ModuleInfo) model.ModuleInfo {
	var moduleInfo model.ModuleInfo
	if _, err := mid.Where(" Id  = ? ", mi.Id).Get(&moduleInfo); err != nil {
		fmt.Println(err.Error())
	}
	return moduleInfo
}

//删除应用模块
func (mid *ModuleInfoDao) DeleteModuleInfo(mi model.ModuleInfo) int64 {

	if _, err := mid.Where("  Id  = ? ", mi.Id).Delete(mi); err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return 1
}

//更新应用模块
func (mid *ModuleInfoDao) UpdateModuleInfo(mi model.ModuleInfo) int64 {

	if result, err := mid.Where("  Id  = ? ", mi.Id).Update(mi); err != nil {
		fmt.Println(err.Error(), result)
		return 0
	}
	return 1
}
