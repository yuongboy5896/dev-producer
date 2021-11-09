package dao

import (
	"dev-producer/model"
	"dev-producer/tool"
	"fmt"
)

type DeployEnvDao struct {
	*tool.Orm
}

//实例化Dao对象
func NewDeployEnvDao() *DeployEnvDao {
	return &DeployEnvDao{tool.DbEngine}
}

//从数据库中查询所有服务器列表
func (ded *DeployEnvDao) QueryDeployEnvs() ([]model.DeployEnv, error) {
	var virtualmachines []model.DeployEnv
	if err := ded.Engine.Find(&virtualmachines); err != nil {
		return nil, err
	}
	return virtualmachines, nil
}

//新虚拟机的数据库插入操作
func (ded *DeployEnvDao) InsertDeployEnv(virtualMachine model.DeployEnv) int64 {
	result, err := ded.InsertOne(&virtualMachine)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return result
}

//查询虚拟机是否存在
func (ded *DeployEnvDao) QueryByDeployEnvs(de model.DeployEnv) model.DeployEnv {
	var virtualMachine model.DeployEnv
	if _, err := ded.Where(" Id  = ? ", de.Id).Get(&virtualMachine); err != nil {
		fmt.Println(err.Error())
	}
	return virtualMachine
}

//删除虚拟机
func (ded *DeployEnvDao) DeleteDeployEnv(de model.DeployEnv) int64 {

	if _, err := ded.Where(" Id  = ? ", de.Id).Delete(de); err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return 1
}

//更新虚拟机
func (ded *DeployEnvDao) UpdateDeployEnv(de model.DeployEnv) int64 {

	if result, err := ded.Where(" Id  = ? ", de.Id).Update(de); err != nil {
		fmt.Println(err.Error(), result)
		return 0
	}
	return 1
}
