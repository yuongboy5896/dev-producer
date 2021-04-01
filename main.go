package main

import (
	"dev-producer/Controller"
	"dev-producer/tool"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg, err := tool.ParseConfig("./config/app.json")
	if err != nil {
		println(err)
	}
	//
	_, err = tool.OrmEngine(cfg)
	if err != nil {
		println(err)
		return
	}

	gin.SetMode(cfg.AppMode)
	app := gin.Default()
	registerRouter(app)
	app.Run(cfg.AppHost + ":" + cfg.AppPort)

	fmt.Printf("start")
}

//路由设置
func registerRouter(router *gin.Engine) {
	new(Controller.HelloController).Router(router)
	new(Controller.K8sController).Router(router)
	new(Controller.ClusterController).Router(router)
}
