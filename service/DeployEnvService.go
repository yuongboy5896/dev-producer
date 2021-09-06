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
func (mis *DeployEnvService) AddDeployEnv(deployEnv model.DeployEnv) int64 {

	vmD := dao.NewDeployEnvDao()

	result := vmD.InsertDeployEnv(deployEnv)

	return result
}

/*
* 查询服务程序
 */
func (mis *DeployEnvService) GetDeployEnv(deployEnv model.DeployEnv) model.DeployEnv {

	vmD := dao.NewDeployEnvDao()

	result := vmD.QueryByDeployEnvs(deployEnv)

	return result
}

/*
* 删除服务程序
 */
func (mis *DeployEnvService) DeleteDeployEnv(deployEnv model.DeployEnv) int64 {

	vmD := dao.NewDeployEnvDao()

	result := vmD.DeleteDeployEnv(deployEnv)

	return result
}

/*
* 更新服务程序
 */
func (mis *DeployEnvService) UpdateDeployEnv(deployEnv model.DeployEnv) int64 {

	vmD := dao.NewDeployEnvDao()

	result := vmD.UpdateDeployEnv(deployEnv)

	return result
}
