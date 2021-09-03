package Controller

import (
	"github.com/gin-gonic/gin"
)

type PipeLineController struct {
}

func (pipeline *PipeLineController) Router(engine *gin.Engine) {
	//添加pipeline
	engine.POST("/api/addpl", pipeline.addpl)
	//获取pipeline
	engine.GET("/api/getpllist", pipeline.getpllist)
	//删除pipeline
	engine.POST("/api/deletepl", pipeline.deletepl)
	//更新pipeline
	engine.POST("/api/updatepl", pipeline.updatepl)


	//添加ModuleInfo
	engine.POST("/api/addmi", pipeline.addmi)
	//获取pipeline
	engine.GET("/api/getmilist", pipeline.getmilist)
	//删除pipeline
	engine.POST("/api/deletemi", pipeline.deletemi)
	//更新pipeline
	engine.POST("/api/updatemi", pipeline.updatemi)

}

func (pipeline *PipeLineController) addpl(context *gin.Context) {

}

func (pipeline *PipeLineController) getpllist(context *gin.Context) {

}

func (pipeline *PipeLineController) deletepl(context *gin.Context) {

}

func (pipeline *PipeLineController) updatepl(context *gin.Context) {

}



func (pipeline *PipeLineController) addmi(context *gin.Context) {

}

func (pipeline *PipeLineController) getmilist(context *gin.Context) {

}

func (pipeline *PipeLineController) deletemi(context *gin.Context) {

}

func (pipeline *PipeLineController) updatemi(context *gin.Context) {

}
