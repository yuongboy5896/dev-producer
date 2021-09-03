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
func (mid *ModuleInfoDao) QueryModuleInfos() ([]model.ModuleInfo, error) {
	var virtualmachines []model.ModuleInfo
	if err := mid.Engine.Find(&virtualmachines); err != nil {
		return nil, err
	}
	return virtualmachines, nil
}

//新虚拟机的数据库插入操作
func (mid *ModuleInfoDao) InsertModuleInfo(virtualMachine model.ModuleInfo) int64 {
	result, err := mid.InsertOne(&virtualMachine)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return result
}

//查询虚拟机是否存在
func (mid *ModuleInfoDao) QueryByModuleInfos(mi model.ModuleInfo) model.ModuleInfo {
	var virtualMachine model.ModuleInfo
	if _, err := mid.Where(" ModuleCode  = ? ", mi.ModuleCode).Get(&virtualMachine); err != nil {
		fmt.Println(err.Error())
	}
	return virtualMachine
}

//删除虚拟机
func (mid *ModuleInfoDao) DeleteModuleInfo(mi model.ModuleInfo) int64 {

	if _, err := mid.Where("  ModuleCode  = ? ", mi.ModuleCode).Delete(mi); err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return 1
}

//更新虚拟机
func (mid *ModuleInfoDao) UpdateModuleInfo(mi model.ModuleInfo) int64 {

	if result, err := mid.Where("  ModuleCode  = ? ", mi.ModuleCode).Update(mi); err != nil {
		fmt.Println(err.Error(), result)
		return 0
	}
	return 1
}
