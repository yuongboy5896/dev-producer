package service

import (
	"dev-producer/dao"
	"dev-producer/model"
)

type ModuleInfoService struct {
}

/**
 * 获取获取 服务程序
 */
func (mis *ModuleInfoService) ModuleInfos() ([]model.ModuleInfo, error) {
	//数据库操作层
	virtualmachineDao := dao.NewModuleInfoDao()
	return virtualmachineDao.QueryModuleInfos()
}

/*
* 添加服务程序
 */
func (mis *ModuleInfoService) AddModuleInfo(moduleInfo model.ModuleInfo) int64 {

	vmD := dao.NewModuleInfoDao()

	result := vmD.InsertModuleInfo(moduleInfo)

	return result
}

/*
* 查询服务程序
 */
func (mis *ModuleInfoService) GetModuleInfo(moduleInfo model.ModuleInfo) model.ModuleInfo {

	vmD := dao.NewModuleInfoDao()

	result := vmD.QueryByModuleInfos(moduleInfo)

	return result
}

/*
* 查询服务程序根据
 */
func (mis *ModuleInfoService) GetModuleInfoById(moduleInfo model.ModuleInfo) model.ModuleInfo {

	vmD := dao.NewModuleInfoDao()

	result := vmD.QueryByIdModuleInfos(moduleInfo)

	return result
}

/*
* 删除服务程序
 */
func (mis *ModuleInfoService) DeleteModuleInfo(moduleInfo model.ModuleInfo) int64 {

	vmD := dao.NewModuleInfoDao()

	result := vmD.DeleteModuleInfo(moduleInfo)

	return result
}

/*
* 更新服务程序
 */
func (mis *ModuleInfoService) UpdateModuleInfo(moduleInfo model.ModuleInfo) int64 {

	vmD := dao.NewModuleInfoDao()

	result := vmD.UpdateModuleInfo(moduleInfo)

	return result
}
