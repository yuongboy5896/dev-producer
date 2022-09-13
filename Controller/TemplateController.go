package Controller

import (
	"dev-producer/model"
	"dev-producer/service"
	"dev-producer/tool"
	"strconv"

	"github.com/gin-gonic/gin"
)

//模版管理
type TemplateInfoControl struct {
}

func (vm *TemplateInfoControl) Router(engine *gin.Engine) {
	//添加模版
	engine.POST("/api/addtc", vm.addTc)
	//修改模版
	engine.POST("/api/updatetc", vm.updateTc)
	//删除模版
	engine.POST("/api/deltc", vm.delTc)
	//获取模版列表
	engine.POST("/api/gettclist", vm.getTclist)
	// 获取ModuleInfo
	engine.GET("/api/gettc/:Id", vm.gettem)
}

func (vm *TemplateInfoControl) addTc(context *gin.Context) {
	//调用service添加虚拟机
	TemplateInfoService := &service.TemplateInfoService{}

	//1、解析 模版信息 传递参数
	var TemplateInfo model.TemplateInfo
	println(context.Request.Body)
	err := tool.Decode(context.Request.Body, &TemplateInfo)
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}
	//2.查询是否存在此模版信息发防止多次提交
	tem := TemplateInfoService.GetTemplateInfo(TemplateInfo)
	if tem.TemplateCode != "" {
		tool.Failed(context, "模版信息已存在")
		return
	}
	//3.调用service添加虚拟机
	result := TemplateInfoService.AddTemplateInfo(TemplateInfo)
	if 0 == result {
		tool.Failed(context, "添加失败")
	}
	tool.Success(context, "添加成功")

}

func (vm *TemplateInfoControl) updateTc(context *gin.Context) {
	//调用service添加虚拟机
	TemplateInfoService := &service.TemplateInfoService{}

	//1、解析 虚拟机信息 传递参数
	var TemplateInfo model.TemplateInfo
	println(context.Request.Body)
	err := tool.Decode(context.Request.Body, &TemplateInfo)
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}
	//更新数据
	result := TemplateInfoService.UpdateTemplateInfo(TemplateInfo)
	if result == 0 {
		tool.Failed(context, "更新失败")
		return
	}
	tool.Success(context, result)
}

func (vm *TemplateInfoControl) delTc(context *gin.Context) {
	//调用service添加虚拟机
	TemplateInfoService := &service.TemplateInfoService{}

	//1、解析 模版信息 传递参数
	var TemplateInfo model.TemplateInfo
	println(context.Request.Body)
	err := tool.Decode(context.Request.Body, &TemplateInfo)
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}
	//删除操作
	result := TemplateInfoService.DeleteTemplateInfo(TemplateInfo)
	if result == 0 {
		tool.Failed(context, result)
		return
	}
	tool.Success(context, result)
}

func (vm *TemplateInfoControl) getTclist(context *gin.Context) {
	// 解析分页
	daoPage, err := tool.PaseUrl(context)
	if err != nil {
		tool.Failed(context, "取服务器列表数据获取失败")
		return
	}
	//调用service功能获取服务器列表
	TemplateInfoService := &service.TemplateInfoService{}
	TemplateInfos, err := TemplateInfoService.TemplateInfolist(&daoPage)
	if err != nil {
		tool.Failed(context, "取服务器列表数据获取失败")
		return
	}
	tool.Success(context, TemplateInfos)

}
func (vm *TemplateInfoControl) gettem(context *gin.Context) {
	//调用service添加服务
	TemplateInfoService := &service.TemplateInfoService{}

	//1、解析 服务信息 传递参数
	var TemplateInfo model.TemplateInfo

	Id := context.Param("Id")
	Id64, err := strconv.ParseInt(Id, 10, 64)
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}
	TemplateInfo.Id = Id64
	TemplateInfo = TemplateInfoService.GetTemplateInfoById(TemplateInfo)
	if TemplateInfo.TemplateCode == "" {
		tool.Failed(context, "获取模块信息失败")
		return
	}
	tool.Success(context, TemplateInfo)
}
