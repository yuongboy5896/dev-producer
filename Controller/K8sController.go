package Controller

import (
	"dev-producer/tool"

	"github.com/gin-gonic/gin"
)

type K8sController struct {
}

func (k8sController *K8sController) Router(engine *gin.Engine) {
	engine.GET("/api/getdeploys", k8sController.GetDeploy)
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
