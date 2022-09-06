package Controller

import (
	"github.com/gin-gonic/gin"
)

//模版管理
type TemplateControl struct {
}

func (vm *TemplateControl) Router(engine *gin.Engine) {
	//添加模版
	engine.POST("/api/addtc", vm.addTc)
	//修改模版
	engine.POST("/api/updatetc", vm.updateTc)
	//删除模版
	engine.POST("/api/deltc", vm.delTc)
	//获取模版列表
	engine.POST("/api/gettclist", vm.getTclist)
}

func (vm *TemplateControl) addTc(context *gin.Context) {

}


func (vm *TemplateControl) updateTc(context *gin.Context) {

}

func (vm *TemplateControl) delTc(context *gin.Context) {

}

func (vm *TemplateControl) getTclist(context *gin.Context) {

}