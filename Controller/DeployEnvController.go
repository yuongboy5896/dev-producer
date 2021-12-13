package Controller

import (
	"dev-producer/model"
	"dev-producer/service"
	"dev-producer/tool"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeployEnvController struct {
}

func (deployEnv *DeployEnvController) Router(engine *gin.Engine) {

	//添加DeployEnv
	engine.POST("/api/addde", deployEnv.addde)
	//获取deployEnv
	engine.GET("/api/getdelist", deployEnv.getdelist)
	//删除deployEnv
	engine.DELETE("/api/deletede/:Id", deployEnv.deletede)
	//更新deployEnv
	engine.PUT("/api/updatede/:Id", deployEnv.updatede)
	//获取信息deployEnv
	engine.GET("/api/getde/:Id", deployEnv.getde)

}

func (mi *DeployEnvController) addde(context *gin.Context) {

	//调用service添加 服务模块
	deployEnvService := &service.DeployEnvService{}

	//1、解析 服务信息 传递参数
	var deployEnv model.DeployEnv
	println(context.Request.Body)
	err := tool.Decode(context.Request.Body, &deployEnv)
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}
	//2.查询是否存在此服务发防止多次提交
	resultmi := deployEnvService.GetDeployEnv(deployEnv)
	if resultmi.EnvIP != "" {
		tool.Failed(context, "已存在服务模块")
		return
	}

	//调用service添加服务
	result := deployEnvService.AddDeployEnv(deployEnv)
	if 0 == result {
		tool.Failed(context, "添加失败")
	}
	tool.Success(context, "添加成功")

}

func (mi *DeployEnvController) getdelist(context *gin.Context) {
	//调用service功能获取服务器列表
	deployEnvService := &service.DeployEnvService{}
	deployEnvs, err := deployEnvService.DeployEnvs()
	if err != nil {
		tool.Failed(context, "取服务器列表数据获取失败")
		return
	}
	tool.Success(context, deployEnvs)
}

func (mi *DeployEnvController) deletede(context *gin.Context) {

	//调用service添加服务
	deployEnvService := &service.DeployEnvService{}
	//1、解析 环境信息 传递参数
	var deployEnv model.DeployEnv
	Id := context.Param("Id")
	Id64, err := strconv.ParseInt(Id, 10, 64)
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}
	deployEnv.Id = Id64

	//删除操作
	result := deployEnvService.DeleteDeployEnv(deployEnv)
	if result == 0 {
		tool.Failed(context, result)
		return
	}
	tool.Success(context, result)
}
func (mi *DeployEnvController) updatede(context *gin.Context) {

	//调用service添加服务
	deployEnvService := &service.DeployEnvService{}

	//1、解析 服务信息 传递参数
	var deployEnv model.DeployEnv
	println(context.Request.Body)
	err := tool.Decode(context.Request.Body, &deployEnv)
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}
	//更新数据
	result := deployEnvService.UpdateDeployEnv(deployEnv)
	if result == 0 {
		tool.Failed(context, "更新失败")
		return
	}
	tool.Success(context, result)
}

func (mi *DeployEnvController) getde(context *gin.Context) {

	//调用service添加服务
	deployEnvService := &service.DeployEnvService{}

	//1、解析 环境信息 传递参数
	var deployEnv model.DeployEnv

	Id := context.Param("Id")
	Id64, err := strconv.ParseInt(Id, 10, 64)
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}
	deployEnv.Id = Id64
	//2. 获取信息
	deployEnv = deployEnvService.GetDeployEnv(deployEnv)
	if deployEnv.EnvIP == "" {
		tool.Failed(context, "获取模块信息失败")
		return
	}
	tool.Success(context, deployEnv)
}
