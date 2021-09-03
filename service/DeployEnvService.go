package service

import (
	"dev-producer/dao"
	"dev-producer/model"
)

type DeployEnvService struct {
}

/**
 * 获取获取 服务程序
 */
func (mis *DeployEnvService) DeployEnvs() ([]model.DeployEnv, error) {
	//数据库操作层
	virtualmachineDao := dao.NewDeployEnvDao()
	return virtualmachineDao.QueryDeployEnvs()
}

/*
* 添加服务程序
 */
func (mis *DeployEnvService) AddDeployEnv(moduleInfo model.DeployEnv) int64 {

	vmD := dao.NewDeployEnvDao()

	result := vmD.InsertDeployEnv(moduleInfo)

	return result
}

/*
* 查询服务程序
 */
func (mis *DeployEnvService) GetDeployEnv(moduleInfo model.DeployEnv) model.DeployEnv {

	vmD := dao.NewDeployEnvDao()

	result := vmD.QueryByDeployEnvs(moduleInfo)

	return result
}

/*
* 删除服务程序
 */
func (mis *DeployEnvService) DeleteDeployEnv(moduleInfo model.DeployEnv) int64 {

	vmD := dao.NewDeployEnvDao()

	result := vmD.DeleteDeployEnv(moduleInfo)

	return result
}

/*
* 更新服务程序
 */
func (mis *DeployEnvService) UpdateDeployEnv(moduleInfo model.DeployEnv) int64 {

	vmD := dao.NewDeployEnvDao()

	result := vmD.UpdateDeployEnv(moduleInfo)

	return result
}
