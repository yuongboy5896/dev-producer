package Controller

import (
	"dev-producer/model"
	"dev-producer/service"
	"dev-producer/tool"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PipeLineController struct {
}

func (pipeline *PipeLineController) Router(engine *gin.Engine) {
	//添加pipeline
	engine.POST("/api/addpl", pipeline.addpl)
	//获取pipeline
	engine.GET("/api/getpllist", pipeline.getpllist)
	// 根据gitlabeId获取pipeline
	engine.GET("/api/getpllistbygitlabid/:GitlabId", pipeline.getpllistbygitlabid)
	// 根据gitlabeId获取pipeline
	engine.GET("/api/getpllistbyid/:Id", pipeline.getpllistbyid)
	// 删除pipeline
	engine.DELETE("/api/deletepl/:Id", pipeline.deletepl)
	// 更新pipeline
	engine.PUT("/api/updatepl/:Id", pipeline.updatepl)
	// 发布pipeline
	engine.PUT("/api/publishplbyid/:Id", pipeline.publishplByid)
	// 发布pipeline
	engine.GET("/api/getjenkinsurl", pipeline.jenkinsUrl)
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
	//3  创建相对于的jenkins流水线
	jenkins := &service.JenkinsService{}

	//4 模版创建pipeline
	bcreate := jenkins.CreateJobFromTmp(pipeLine.PipeCode, pipeLine.TechnologyType, pipeLine)
	if !bcreate {
		tool.Failed(context, "添加jenkins pipeline 失败")
		return
	}
	//调用service添加流水线
	result := pipeLineService.AddPipeLine(pipeLine)
	if 0 == result {
		tool.Failed(context, "添加失败")
		return
	}
	tool.Success(context, "添加成功")
}

func (pipeline *PipeLineController) getpllist(context *gin.Context) {
	pipeLineService := &service.PipeLineService{}
	pipeLines, err := pipeLineService.PipeLines()
	if err != nil {
		tool.Failed(context, "取pipeline列表数据获取失败")
		return
	}
	tool.Success(context, pipeLines)
}
func (pipeline *PipeLineController) getpllistbygitlabid(context *gin.Context) {

	Id := context.Param("GitlabId")
	Id64, err := strconv.ParseInt(Id, 10, 64)
	pipeLineService := &service.PipeLineService{}
	pipeLines, err := pipeLineService.PipeLinesByGitLabId(Id64)
	if err != nil {
		tool.Failed(context, "取服务器列表数据获取失败")
		return
	}
	tool.Success(context, pipeLines)
}
func (pipeline *PipeLineController) getpllistbyid(context *gin.Context) {

	Id := context.Param("Id")
	Id64, err := strconv.ParseInt(Id, 10, 64)
	pipeLineService := &service.PipeLineService{}
	pipeLine, err := pipeLineService.PipeLinesById(Id64)
	if err != nil {
		tool.Failed(context, "取pipeline列表数据获取失败")
		return
	}
	tool.Success(context, pipeLine)
}

func (pipeline *PipeLineController) deletepl(context *gin.Context) {
	//调用service添加流水线
	pipeLineService := &service.PipeLineService{}

	//1、解析 pipeline信息 传递参数
	var pipeLine model.PipeLine
	Id := context.Param("Id")
	Id64, err := strconv.ParseInt(Id, 10, 64)
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}
	pipeLine.Id = Id64
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
	//4 模版创建pipeline
	jenkins := &service.JenkinsService{}
	bcreate := jenkins.CreateJobFromTmp(pipeLine.PipeCode, pipeLine.TechnologyType, pipeLine)
	if !bcreate {
		tool.Failed(context, "添加jenkins pipeline 失败")
		return
	}
	tool.Success(context, result)
}

func (pipeline *PipeLineController) publishplByid(context *gin.Context) {

	Id := context.Param("Id")
	Id64, err := strconv.ParseInt(Id, 10, 64)
	pipeLineService := &service.PipeLineService{}
	pipeLine, err := pipeLineService.PipeLinesById(Id64)
	if err != nil {
		tool.Failed(context, "取pipeline列表数据获取失败")
		return
	}

	// 镜像库地址 未实现

	pipeLineService.PublishPipeLine(pipeLine)

	tool.Success(context, pipeLine)
}

//
func (pipeline *PipeLineController) jenkinsUrl (context *gin.Context) {


	
	config := tool.GetConfig().JenkinsConfig
	url := "http://" + config.Addr + ":" + config.Port + "/job/"

	tool.Success(context, url)
}