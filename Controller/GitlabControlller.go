package Controller

import (
	"dev-producer/service"
	"dev-producer/tool"

	"github.com/gin-gonic/gin"
)

type GitlabController struct {
}

func (Gitlab *GitlabController) Router(engine *gin.Engine) {

	//添加GitlabProjects
	engine.GET("/api/getgplist", Gitlab.getgplist)
	//获取分支信息按gitlabprojectiD
	engine.GET("/api/getgrlist/:Id", Gitlab.getgrlist)

}

func (gc *GitlabController) getgplist(context *gin.Context) {

	page := context.Query("page")

	if page == "" {
		page = "1"
	}
	gitlabService := &service.GitlabService{}
	moduleInfos, err := gitlabService.GitlabProject(page)
	if err != nil {
		tool.Failed(context, "取GitlabProjects列表数据获取失败")
		return
	}
	tool.Success(context, moduleInfos)

}
func (gc *GitlabController) getgrlist(context *gin.Context) {

	//gitProID := context.Query("gitlabproid")
	gitProID := context.Param("Id")
	if gitProID == "" {
		tool.Failed(context, "未获取导参数")
		return
	}
	gitlabService := &service.GitlabService{}
	gitlabBranchs, err := gitlabService.GitlabBranch(gitProID)
	if err != nil {
		tool.Failed(context, "取分支信息数据获取失败")
		return
	}
	tool.Success(context, gitlabBranchs)

}
