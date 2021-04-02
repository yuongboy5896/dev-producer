package Controller

import (
	"dev-producer/model"
	"dev-producer/service"
	"dev-producer/tool"

	"github.com/gin-gonic/gin"
)

type ClusterController struct {
}

func (clusterController *ClusterController) Router(engine *gin.Engine) {
	engine.POST("/api/importcluster", clusterController.ImportCluster)
}

//http://localhost:8090/api/importcluster
func (clusterController *ClusterController) ImportCluster(context *gin.Context) {

	//1、解析集群信息传递参数
	var importCluster model.ClusterInfo
	err := tool.Decode(context.Request.Body, &importCluster)
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}
	clusterInfoservice := service.ClusterInfoService{}
	clusterInfoservice.Import(&importCluster)
	tool.Success(context, "发送成功"+importCluster.Clustercode)
}

//http://localhost:8090/api/uploadconfig
