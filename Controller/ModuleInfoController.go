package Controller

import (
	"dev-producer/model"
	"dev-producer/service"
	"dev-producer/tool"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ModuleInfoController struct {
}

func (moduleInfo *ModuleInfoController) Router(engine *gin.Engine) {

	//添加ModuleInfo
	engine.POST("/api/addmi", moduleInfo.addmi)
	//获取moduleInfo
	engine.GET("/api/getmilist", moduleInfo.getmilist)
	//删除moduleInfo
	engine.DELETE("/api/deletemi/:Id", moduleInfo.deletemi)
	// 更新moduleInfo
	engine.PUT("/api/updatemi/:Id", moduleInfo.updatemi)
	// 获取ModuleInfo
	engine.GET("/api/getmi/:Id", moduleInfo.getmi)

}

func (mi *ModuleInfoController) addmi(context *gin.Context) {

	//调用service添加 服务模块
	moduleInfoService := &service.ModuleInfoService{}

	//1、解析 服务信息 传递参数
	var moduleInfo model.ModuleInfo
	println(context.Request.Body)
	err := tool.Decode(context.Request.Body, &moduleInfo)
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}
	//2.查询是否存在此服务发防止多次提交
	resultmi := moduleInfoService.GetModuleInfo(moduleInfo)
	if resultmi.ModuleCode != "" {
		tool.Failed(context, "已存在服务模块")
		return
	}
	//设置gitlabid
	gitlabService := &service.GitlabService{}
	for j := 1; j < 20; j++ {
		page := fmt.Sprintf("%d", j)
		gitlabProjects, err := gitlabService.GitlabProject(page)
		//http://192.168.48.15:8080/thpower-energy-cloud/enterpriseplatform
		//http://gitlab.thpower.com:8080/thpower-energy-cloud/enterpriseplatform.git
		if err != nil {
			tool.Failed(context, "添加失败")
			return
		}
		web_url := strings.Split(moduleInfo.GitlabUrl, "/")
		strsize := len(web_url)
		lasturl := web_url[strsize-2] + "/" + web_url[strsize-1]
		for i := 0; i < len(gitlabProjects); i++ {
			fmt.Println(gitlabProjects[i].Web_url)
			if find := strings.Contains(gitlabProjects[i].Web_url, lasturl); find {
				moduleInfo.HttpUrlToRepo = gitlabProjects[i].Http_url_to_repo
				moduleInfo.SshUrlToRepo = gitlabProjects[i].Ssh_url_to_repo
				//moduleInfo.GitlabId = gitlabProjects[i].Id
				str := fmt.Sprintf("%d", gitlabProjects[i].Id)
				moduleInfo.GitlabId = str
				break
			}
		}
	}

	//调用service添加服务
	result := moduleInfoService.AddModuleInfo(moduleInfo)
	if 0 == result {
		tool.Failed(context, "添加失败")
	}
	tool.Success(context, "添加成功")

}

func (mi *ModuleInfoController) getmilist(context *gin.Context) {
	// 解析分页
	daoPage, err := tool.PaseUrl(context)
	if err != nil {
		tool.Failed(context, "取服务器列表数据获取失败")
		return
	}
	//调用service功能获取服务器列表
	moduleInfoService := &service.ModuleInfoService{}
	moduleInfos, err := moduleInfoService.ModuleInfos(&daoPage)
	if err != nil {
		tool.Failed(context, "取服务器列表数据获取失败")
		return
	}
	tool.Success(context, moduleInfos)
}

func (mi *ModuleInfoController) deletemi(context *gin.Context) {

	//调用service添加服务
	moduleInfoService := &service.ModuleInfoService{}

	//1、解析 服务信息 传递参数
	var moduleInfo model.ModuleInfo
	println(context.Request.Body)
	Id := context.Param("Id")
	Id64, err := strconv.ParseInt(Id, 10, 64)
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}
	moduleInfo.Id = Id64

	//删除操作
	result := moduleInfoService.DeleteModuleInfo(moduleInfo)
	if result == 0 {
		tool.Failed(context, result)
		return
	}
	tool.Success(context, result)
}
func (mi *ModuleInfoController) updatemi(context *gin.Context) {

	//调用service添加服务
	moduleInfoService := &service.ModuleInfoService{}

	//1、解析 服务信息 传递参数
	var moduleInfo model.ModuleInfo
	println(context.Request.Body)
	err := tool.Decode(context.Request.Body, &moduleInfo)
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}
	Id := context.Param("Id")
	Id64, err := strconv.ParseInt(Id, 10, 64)
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}
	moduleInfo.Id = Id64

	//更新数据
	result := moduleInfoService.UpdateModuleInfo(moduleInfo)
	if result == 0 {
		tool.Failed(context, "更新失败")
		return
	}
	tool.Success(context, result)
}
func (mi *ModuleInfoController) getmi(context *gin.Context) {
	//调用service添加服务
	moduleInfoService := &service.ModuleInfoService{}

	//1、解析 服务信息 传递参数
	var moduleInfo model.ModuleInfo

	Id := context.Param("Id")
	Id64, err := strconv.ParseInt(Id, 10, 64)
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}
	moduleInfo.Id = Id64
	moduleInfo = moduleInfoService.GetModuleInfoById(moduleInfo)
	if moduleInfo.ModuleName == "" {
		tool.Failed(context, "获取模块信息失败")
		return
	}
	tool.Success(context, moduleInfo)
}
