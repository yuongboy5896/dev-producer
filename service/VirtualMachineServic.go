package service

import (
	"dev-producer/dao"
	"dev-producer/model"
)

type VirtualMachineService struct {
}

/**
 * 获取获取虚拟机列表
 */
func (fcs *VirtualMachineService) VirtualMachines() ([]model.VirtualMachine, error) {

	//数据库操作层
	virtualmachineDao := dao.NewVirtualMachineDao()
	return virtualmachineDao.QueryVirtualMachines()
}

/*
* 添加虚拟机
 */
func (fcs *VirtualMachineService) AddVirtualMachine(virtualMachine model.VirtualMachine) int64 {

	vmD := dao.NewVirtualMachineDao()

	result := vmD.InsertVirtualMachine(virtualMachine)

	return result
}

/*
* 查询虚拟机
 */
func (fcs *VirtualMachineService) GetVirtualMachine(virtualMachine model.VirtualMachine) model.VirtualMachine {

	vmD := dao.NewVirtualMachineDao()

	result := vmD.QueryByVirtualMachines(virtualMachine)

	return result
}

/*
* 删除虚拟机
 */
func (fcs *VirtualMachineService) DeleteVirtualMachine(virtualMachine model.VirtualMachine) int64 {

	vmD := dao.NewVirtualMachineDao()

	result := vmD.DeleteVirtualMachine(virtualMachine)

	return result
}

/*
* 更新虚拟机
 */
func (fcs *VirtualMachineService) UpdateVirtualMachine(virtualMachine model.VirtualMachine) int64 {

	vmD := dao.NewVirtualMachineDao()

	result := vmD.UpdateVirtualMachine(virtualMachine)

	return result
}
