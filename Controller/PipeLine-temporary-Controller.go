package Controller

import (
	"dev-producer/model"
	"dev-producer/service"
	"dev-producer/tool"

	"github.com/gin-gonic/gin"
)

type PipeTemporayController struct {
}

func (pipetemporay *PipeTemporayController) Router(engine *gin.Engine) {

	//添加pipeline
	engine.POST("/api/addpt", pipetemporay.addpt)
}

func (pipetemporay *PipeTemporayController) addpt(context *gin.Context) {
	//调用service添加 服务模块
	PipeTemporayService := &service.PipeTemporayService{}

	//1、解析 服务信息 传递参数
	var ModuleForImageUrl model.ModuleForImageUrl
	println(context.Request.Body)
	err := tool.Decode(context.Request.Body, &ModuleForImageUrl)
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}
	//2.查询是否存在此服务发防止多次提交
	resulpfi := PipeTemporayService.GetModuleForImageUrl(ModuleForImageUrl)
	if resulpfi.ModuleCode != "" {
		//省略一个接口搞定update 模块信息
		result := PipeTemporayService.UpdateModuleInfo(ModuleForImageUrl)
		if 0 == result {
			tool.Failed(context, "镜像地址更新失败")
		}
		tool.Success(context, "镜像地址更新成功")
		return
	}
	//调用service添加服务
	result := PipeTemporayService.AddModuleForImageUrl(ModuleForImageUrl)
	if 0 == result {
		tool.Failed(context, "添加失败")
	}
	tool.Success(context, "添加成功")
}
