package Controller

import (
	"dev-producer/model"
	"dev-producer/service"
	"dev-producer/tool"

	"github.com/gin-gonic/gin"
)

type VirtualMachineControl struct {
}

func (vm *VirtualMachineControl) Router(engine *gin.Engine) {

	//添加虚拟机
	engine.POST("/api/addvm", vm.addVm)
	//获取虚拟机列表
	engine.GET("/api/getvmlist", vm.getvmlist)
	//删除虚拟机
	engine.POST("/api/deletevm", vm.deletevm)
	//删除虚拟机
	engine.POST("/api/updatevm", vm.updatevm)

}

func (vm *VirtualMachineControl) addVm(context *gin.Context) {

	//调用service添加虚拟机
	virtualMachineService := &service.VirtualMachineService{}

	//1、解析 虚拟机信息 传递参数
	var virtualMachine model.VirtualMachine
	println(context.Request.Body)
	err := tool.Decode(context.Request.Body, &virtualMachine)
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}
	//2.查询是否存在此虚拟机发防止多次提交
	resultvm := virtualMachineService.GetVirtualMachine(virtualMachine)
	if resultvm.IP != "" {
		tool.Failed(context, "已存在IP或者已存在服务器")
		return
	}

	//调用service添加虚拟机
	result := virtualMachineService.AddVirtualMachine(virtualMachine)
	if 0 == result {
		tool.Failed(context, "添加失败")
	}
	tool.Success(context, "添加成功")

}

func (vm *VirtualMachineControl) getvmlist(context *gin.Context) {
	//调用service功能获取服务器列表
	virtualMachineService := &service.VirtualMachineService{}
	virtualMachines, err := virtualMachineService.VirtualMachines()
	if err != nil {
		tool.Failed(context, "取服务器列表数据获取失败")
		return
	}
	tool.Success(context, virtualMachines)
}

func (vm *VirtualMachineControl) deletevm(context *gin.Context) {

	//调用service添加虚拟机
	virtualMachineService := &service.VirtualMachineService{}

	//1、解析 虚拟机信息 传递参数
	var virtualMachine model.VirtualMachine
	println(context.Request.Body)
	err := tool.Decode(context.Request.Body, &virtualMachine)
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}
	//删除操作
	result := virtualMachineService.DeleteVirtualMachine(virtualMachine)
	if result == 0 {
		tool.Failed(context, result)
		return
	}
	tool.Success(context, result)
}
func (vm *VirtualMachineControl) updatevm(context *gin.Context) {

	//调用service添加虚拟机
	virtualMachineService := &service.VirtualMachineService{}

	//1、解析 虚拟机信息 传递参数
	var virtualMachine model.VirtualMachine
	println(context.Request.Body)
	err := tool.Decode(context.Request.Body, &virtualMachine)
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}
	//更新数据
	result := virtualMachineService.UpdateVirtualMachine(virtualMachine)
	if result == 0 {
		tool.Failed(context, "更新失败")
		return
	}
	tool.Success(context, result)
}
