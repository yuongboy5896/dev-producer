package Controller

import (
	"dev-producer/service"
	"dev-producer/tool"
	"strconv"

	"github.com/gin-gonic/gin"
)

type K8sController struct {
}

func (k8sController *K8sController) Router(engine *gin.Engine) {
	//engine.GET("/api/getdeploys", k8sController.GetDeploy)

	engine.GET("/api/getnamesapces", k8sController.GetNameSpaces)
}

//http://localhost:8090/api/getdeploys?namespace=default&deploy=test&clusterid=111
func (k8sController *K8sController) GetDeploy(context *gin.Context) {

	clusterid, exist := context.GetQuery("clusterid")
	if !exist {
		tool.Failed(context, "参数解析失败"+clusterid)
		return
	}
	namespace, exist := context.GetQuery("namespace")
	if !exist {
		namespace = ""
	}
	deploy, exist := context.GetQuery("deploy")
	if !exist {
		deploy = ""
	}
	if true {
		tool.Success(context, "发送成功"+deploy+"/t"+namespace)
		return
	}

}

//http://localhost:8090/api/getnamesapces?clusterid=
func (k8sController *K8sController) GetNameSpaces(context *gin.Context) {
	clusterid, exist := context.GetQuery("clusterid")
	if !exist {
		tool.Failed(context, "参数解析失败"+clusterid)
		return
	}

	//
	//调用service添加服务
	Id64, err := strconv.ParseInt(clusterid, 10, 64)
	if err != nil {
		tool.Failed(context, "获取集群信息解析")
		return
	}
	k8sApiService := &service.K8sApiService{}
	items, err := k8sApiService.GetNameSpaces(Id64)
	if err != nil {
		tool.Failed(context, "获取集群信息失败")
		return
	}
	tool.Success(context, items)
	return
}
