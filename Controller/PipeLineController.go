package Controller

import (
	"dev-producer/model"
	"dev-producer/service"
	"dev-producer/tool"

	"github.com/gin-gonic/gin"
)

type PipeLineController struct {
}

func (pipeline *PipeLineController) Router(engine *gin.Engine) {
	//添加pipeline
	engine.POST("/api/addpl", pipeline.addpl)
	//获取pipeline
	engine.GET("/api/getpllist", pipeline.getpllist)
	//获取pipeline
	engine.GET("/api/getpllistbyid", pipeline.getpllistbyid)
	//删除pipeline
	engine.POST("/api/deletepl", pipeline.deletepl)
	//更新pipeline
	engine.POST("/api/updatepl", pipeline.updatepl)
}

func (pipeline *PipeLineController) addpl(context *gin.Context) {
	//调用service添加流水线
	pipeLineService := &service.PipeLineService{}

	//1、解析 pipeline信息 传递参数
	var pipeLine model.PipeLine
	println(context.Request.Body)
	err := tool.Decode(context.Request.Body, &pipeLine)
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}
	//2.查询是否存在此流水线发防止多次提交
	resultpl := pipeLineService.GetPipeLine(pipeLine)
	if resultpl.Pipename != "" {
		tool.Failed(context, "已存在Pipeline")
		return
	}

	//调用service添加流水线
	result := pipeLineService.AddPipeLine(pipeLine)
	if 0 == result {
		tool.Failed(context, "添加失败")
	}
	tool.Success(context, "添加成功")
}

func (pipeline *PipeLineController) getpllist(context *gin.Context) {
	pipeLineService := &service.PipeLineService{}
	pipeLines, err := pipeLineService.PipeLines()
	if err != nil {
		tool.Failed(context, "取服务器列表数据获取失败")
		return
	}
	tool.Success(context, pipeLines)
}
func (pipeline *PipeLineController) getpllistbyid(context *gin.Context) {
	pipeLineService := &service.PipeLineService{}
	pipeLines, err := pipeLineService.PipeLines()
	if err != nil {
		tool.Failed(context, "取服务器列表数据获取失败")
		return
	}
	tool.Success(context, pipeLines)
}

func (pipeline *PipeLineController) deletepl(context *gin.Context) {
	//调用service添加流水线
	pipeLineService := &service.PipeLineService{}

	//1、解析 pipeline信息 传递参数
	var pipeLine model.PipeLine
	println(context.Request.Body)
	err := tool.Decode(context.Request.Body, &pipeLine)
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}
	//删除操作
	result := pipeLineService.DeletePipeLine(pipeLine)
	if result == 0 {
		tool.Failed(context, result)
		return
	}
	tool.Success(context, result)
}

func (pipeline *PipeLineController) updatepl(context *gin.Context) {
	//调用service添加流水线
	pipeLineService := &service.PipeLineService{}

	//1、解析 pipeline信息 传递参数
	var pipeLine model.PipeLine
	println(context.Request.Body)
	err := tool.Decode(context.Request.Body, &pipeLine)
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}
	//更新数据
	result := pipeLineService.UpdatePipeLine(pipeLine)
	if result == 0 {
		tool.Failed(context, "更新失败")
		return
	}
	tool.Success(context, result)
}
