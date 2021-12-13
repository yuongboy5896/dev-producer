package Controller

import (
	"github.com/gin-gonic/gin"
)

type RepositoryController struct {
}

func (repository *RepositoryController) Router(engine *gin.Engine) {
	// 添加adddr
	engine.POST("/api/adddr", repository.adddr)
	// 获取镜像仓库
	engine.GET("/api/getddrbyid/:Id", repository.getddrbyid)
	// 获取获取镜像仓库列表
	engine.GET("/api/getadddrlist", repository.getdrlist)
	// 获取获取镜像仓库
	engine.GET("/api/getadddrlist", repository.getdrlist)
	// 删除获取镜像仓库
	engine.DELETE("/api/deletedr/:Id", repository.deletedr)
	// 更新获取镜像仓库
	engine.PUT("/api/updatedr/:Id", repository.updatedr)

}

// 添加镜像仓库
func (repository *RepositoryController) adddr(context *gin.Context) {

}

// 添加镜像仓库列表
func (repository *RepositoryController) getdrlist(context *gin.Context) {

}

// 删除镜像仓库列表
func (repository *RepositoryController) deletedr(context *gin.Context) {

}

// 更新镜像仓库列表
func (repository *RepositoryController) updatedr(context *gin.Context) {

}

// 更新镜像仓库
func (repository *RepositoryController) getddrbyid(context *gin.Context) {

}
