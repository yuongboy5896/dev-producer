package Controller

import (
	"dev-producer/model"
	"dev-producer/service"
	"dev-producer/tool"

	"github.com/gin-gonic/gin"
)

type PipeLineSimpleController struct {
}

func (pipeline *PipeLineSimpleController) Router(engine *gin.Engine) {
	//添加pipeline
	engine.POST("/api/addpls", pipeline.addpls)
	//获取pipeline
	engine.GET("/api/getplslist", pipeline.getplslist)
	//删除pipeline
	engine.POST("/api/deleteplss", pipeline.deletepls)
	//更新pipeline
	engine.POST("/api/updateplss", pipeline.updatepls)
}

func (pipeline *PipeLineSimpleController) addpls(context *gin.Context) {
	//调用service添加虚拟机
	pipeLineService := &service.PipeLineSimpleService{}

	//1、解析 pipeline信息 传递参数
	var pipeLine model.PipeLineSimple
	println(context.Request.Body)
	err := tool.Decode(context.Request.Body, &pipeLine)
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}
	//2.查询是否存在此虚拟机发防止多次提交
	resultpl := pipeLineService.GetPipeLineSimple(pipeLine)
	if resultpl.Pipename != "" {
		tool.Failed(context, "已存在Pipeline")
		return
	}

	//调用service添加虚拟机
	result := pipeLineService.AddPipeLineSimple(pipeLine)
	if 0 == result {
		tool.Failed(context, "添加失败")
	}
	tool.Success(context, "添加成功")
}

func (pipeline *PipeLineSimpleController) getplslist(context *gin.Context) {
	pipeLineService := &service.PipeLineSimpleService{}
	pipeLines, err := pipeLineService.PipeLineSimples()
	if err != nil {
		tool.Failed(context, "取服务器列表数据获取失败")
		return
	}
	tool.Success(context, pipeLines)
}

func (pipeline *PipeLineSimpleController) deletepls(context *gin.Context) {
	//调用service添加虚拟机
	pipeLineService := &service.PipeLineSimpleService{}

	//1、解析 pipeline信息 传递参数
	var pipeLine model.PipeLineSimple
	println(context.Request.Body)
	err := tool.Decode(context.Request.Body, &pipeLine)
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}
	//删除操作
	result := pipeLineService.DeletePipeLineSimple(pipeLine)
	if result == 0 {
		tool.Failed(context, result)
		return
	}
	tool.Success(context, result)
}

func (pipeline *PipeLineSimpleController) updatepls(context *gin.Context) {
	//调用service添加虚拟机
	pipeLineService := &service.PipeLineSimpleService{}

	//1、解析 pipeline信息 传递参数
	var pipeLine model.PipeLineSimple
	println(context.Request.Body)
	err := tool.Decode(context.Request.Body, &pipeLine)
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}
	//更新数据
	result := pipeLineService.UpdatePipeLineSimple(pipeLine)
	if result == 0 {
		tool.Failed(context, "更新失败")
		return
	}
	tool.Success(context, result)
}
