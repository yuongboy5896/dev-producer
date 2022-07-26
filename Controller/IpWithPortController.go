package Controller

import (
	"dev-producer/model"
	"dev-producer/service"
	"dev-producer/tool"
	"net"
	"strings"

	"github.com/gin-gonic/gin"
)

type IpWithPortController struct {
}

func (ipWithPortController *IpWithPortController) Router(engine *gin.Engine) {
	engine.POST("/api/scanips", ipWithPortController.ScanIPs)
	engine.GET("/api/ipalive", ipWithPortController.ScanIPs)
}

/////

func (ipWithPortController *IpWithPortController) ScanIPs(context *gin.Context) {
	//
	//
	//1、解析IP段
	var ipalive model.IpAlive

	err := tool.Decode(context.Request.Body, &ipalive)
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}
	address := net.ParseIP(ipalive.Ip)
	if address == nil {
		tool.Failed(context, "ip地址格式不正确")
		return
	}
	//检查ip
	netWorkService := service.NetWorkService{}
	iplist := strings.Split(ipalive.Ip, ".")
	tmp := iplist[0] + "." + iplist[1] + "." + iplist[2] + "."
	//端口范围

	IpResult := netWorkService.ScanIP(tmp)
	if iplist == nil {
		tool.Failed(context, "未扫描到相关信息")
		return
	}
	tool.Success(context, IpResult)

}
func (ipWithPortController *IpWithPortController) Ipalive(context *gin.Context) {
	// 1、解析IP地址
	//

}
