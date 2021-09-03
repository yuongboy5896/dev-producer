package Controller

import (
	"dev-producer/service"
	"dev-producer/tool"

	"github.com/gin-gonic/gin"
)

type DeploymentController struct {
}

func (deployment *DeploymentController) Router(engine *gin.Engine) {
	//添加deployment
	engine.POST("/api/addpl", deployment.adddeploy)
	//获取deployment
	engine.GET("/api/getpllist", deployment.getdeploylist)
	//删除deployment
	engine.POST("/api/deletepl", deployment.deletedeploy)
	//更新deployment
	engine.POST("/api/updatepl", deployment.updatedeploy)

}

func (deployment *DeploymentController) adddeploy(context *gin.Context) {

}

func (deployment *DeploymentController) getdeploylist(context *gin.Context) {
	//调用service功能获取服务器列表
	virtualMachineService := &service.VirtualMachineService{}
	virtualMachines, err := virtualMachineService.VirtualMachines()
	if err != nil {
		tool.Failed(context, "取服务器列表数据获取失败")
		return
	}
	tool.Success(context, virtualMachines)
}

func (deployment *DeploymentController) deletedeploy(context *gin.Context) {

}

func (deployment *DeploymentController) updatedeploy(context *gin.Context) {

}
