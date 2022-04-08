package dao

import (
	"dev-producer/model"
	"dev-producer/tool"
	"fmt"
)

type VirtualMachineDao struct {
	*tool.Orm
}

//实例化Dao对象
func NewVirtualMachineDao() *VirtualMachineDao {
	return &VirtualMachineDao{tool.DbEngine}
}

//从数据库中查询所有服务器列表
func (vmd *VirtualMachineDao) QueryVirtualMachines() ([]model.VirtualMachine, error) {
	
	var virtualmachines []model.VirtualMachine
	if err := vmd.Engine.Find(&virtualmachines); err != nil {
		return nil, err
	}
	return virtualmachines, nil
}

//新虚拟机的数据库插入操作
func (vmd *VirtualMachineDao) InsertVirtualMachine(virtualMachine model.VirtualMachine) int64 {
	result, err := vmd.InsertOne(&virtualMachine)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return result
}

//查询虚拟机是否存在
func (vmd *VirtualMachineDao) QueryByVirtualMachines(vm model.VirtualMachine) model.VirtualMachine {
	var virtualMachine model.VirtualMachine
	if _, err := vmd.Where(" i_p  = ? ", vm.IP).Get(&virtualMachine); err != nil {
		fmt.Println(err.Error())
	}
	return virtualMachine
}

//删除虚拟机
func (vmd *VirtualMachineDao) DeleteVirtualMachine(vm model.VirtualMachine) int64 {

	if _, err := vmd.Where(" i_p  = ? ", vm.IP).Delete(vm); err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return 1
}

//更新虚拟机
func (vmd *VirtualMachineDao) UpdateVirtualMachine(vm model.VirtualMachine) int64 {

	if result, err := vmd.Where(" i_p  = ? ", vm.IP).Update(vm); err != nil {
		fmt.Println(err.Error(), result)
		return 0
	}
	return 1
}
