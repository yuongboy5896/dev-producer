package Controller

import (
	"dev-producer/model"
	"dev-producer/service"
	"dev-producer/tool"

	"github.com/gin-gonic/gin"
)

type DeployEnvController struct {
}

func (moduleInfo *DeployEnvController) Router(engine *gin.Engine) {

	//添加DeployEnv
	engine.POST("/api/addde", moduleInfo.addde)
	//获取moduleInfo
	engine.GET("/api/getdelist", moduleInfo.getdelist)
	//删除moduleInfo
	engine.POST("/api/deletede", moduleInfo.deletede)
	//更新moduleInfo
	engine.POST("/api/updatede", moduleInfo.updatede)

}

func (mi *DeployEnvController) addde(context *gin.Context) {

	//调用service添加 服务模块
	moduleInfoService := &service.DeployEnvService{}

	//1、解析 服务信息 传递参数
	var moduleInfo model.DeployEnv
	println(context.Request.Body)
	err := tool.Decode(context.Request.Body, &moduleInfo)
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}
	//2.查询是否存在此服务发防止多次提交
	resultmi := moduleInfoService.GetDeployEnv(moduleInfo)
	if resultmi.EnvIP != "" {
		tool.Failed(context, "已存在服务模块")
		return
	}

	//调用service添加服务
	result := moduleInfoService.AddDeployEnv(moduleInfo)
	if 0 == result {
		tool.Failed(context, "添加失败")
	}
	tool.Success(context, "添加成功")

}

func (mi *DeployEnvController) getdelist(context *gin.Context) {
	//调用service功能获取服务器列表
	moduleInfoService := &service.DeployEnvService{}
	moduleInfos, err := moduleInfoService.DeployEnvs()
	if err != nil {
		tool.Failed(context, "取服务器列表数据获取失败")
		return
	}
	tool.Success(context, moduleInfos)
}

func (mi *DeployEnvController) deletede(context *gin.Context) {

	//调用service添加服务
	moduleInfoService := &service.DeployEnvService{}

	//1、解析 服务信息 传递参数
	var moduleInfo model.DeployEnv
	println(context.Request.Body)
	err := tool.Decode(context.Request.Body, &moduleInfo)
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}
	//删除操作
	result := moduleInfoService.DeleteDeployEnv(moduleInfo)
	if result == 0 {
		tool.Failed(context, result)
		return
	}
	tool.Success(context, result)
}
func (mi *DeployEnvController) updatede(context *gin.Context) {

	//调用service添加服务
	moduleInfoService := &service.DeployEnvService{}

	//1、解析 服务信息 传递参数
	var moduleInfo model.DeployEnv
	println(context.Request.Body)
	err := tool.Decode(context.Request.Body, &moduleInfo)
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}
	//更新数据
	result := moduleInfoService.UpdateDeployEnv(moduleInfo)
	if result == 0 {
		tool.Failed(context, "更新失败")
		return
	}
	tool.Success(context, result)
}
