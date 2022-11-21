package Controller

import (
	"dev-producer/model"
	"dev-producer/param"
	"dev-producer/service"
	"dev-producer/tool"
	"strconv"

	"github.com/gin-gonic/gin"
)

type K8sController struct {
}

func (k8sController *K8sController) Router(engine *gin.Engine) {
	//engine.GET("/api/getdeploys", k8sController.GetDeploy)
	//
	engine.GET("/api/getnamesapces", k8sController.GetNameSpaces)
	//
	engine.POST("/api/createfromyaml", k8sController.CreateFromYaml)
	//
	engine.GET("/api/getdeploystatus", k8sController.GetDeployStatus)
}

//http://localhost:8090/api/getdeploys?namespace=default&deploy=test&envid=111
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

//部署yaml
func (K8sController *K8sController) CreateFromYaml(context *gin.Context) {
	//模版信息 是从项目信息中的编号
	//1、解析项目传递参数
	var pipelineParam model.PipeLine

	err := tool.Decode(context.Request.Body, &pipelineParam)
	k8sApiService := &service.K8sApiService{}
	k8sApiService.CreateFromYaml(pipelineParam)
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}
	tool.Success(context, "部署成功")
}

//获取项目状态
func (K8sController *K8sController) GetDeployStatus(context *gin.Context) {
	//以 模块编号为准
	ModuleCode, exist := context.GetQuery("modulecode")
	if !exist {
		tool.Failed(context, "参数解析失败"+ModuleCode)
		return
	}
	NameSpace, exist := context.GetQuery("namespace")
	if !exist {
		tool.Failed(context, "参数解析失败"+NameSpace)
		return
	}
	envid, exist := context.GetQuery("envid")
	if !exist {
		tool.Failed(context, "参数解析失败"+envid)
		return
	}
	var getDeployParam param.K8sGetDeployParam
	getDeployParam.ModuleCode = ModuleCode
	getDeployParam.NameSpace = NameSpace
	Id64, err := strconv.ParseInt(envid, 10, 64)
	getDeployParam.EnvId = Id64
	k8sApiService := &service.K8sApiService{}
	is, err := k8sApiService.GetDeployInfo(getDeployParam)
	if err != nil {
		tool.Failed(context, "调用k8s失败")
		return
	}
	if !is {
		tool.Failed(context, "模块不存在")
		return
	}
	tool.Success(context, "模块存在")
}

