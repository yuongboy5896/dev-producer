package Controller

import (
	"dev-producer/param"
	"dev-producer/service"
	"dev-producer/tool"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

type MemberConntroller struct {
}

func (mc *MemberConntroller) Router(engine *gin.Engine) {

	//login_pwd
	engine.POST("/api/login_pwd", mc.nameLogin)

}

func (mc *MemberConntroller) nameLogin(context *gin.Context) {

	//1、解析用户登录传递参数
	var loginParam param.LoginParam
	err := tool.Decode(context.Request.Body, &loginParam)
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}

	//2、验证验证码
	validate := tool.VertifyCaptcha(loginParam.Id, loginParam.Value)
	if !validate {
		tool.Failed(context, "验证码不正确，请重新验证")
		return
	}
	//3、登录
	ms := service.MemberService{}
	member := ms.Login(loginParam.Name, loginParam.Password)
	if member.Id != 0 {
		//用户信息保存到session
		sess, _ := json.Marshal(member)
		err = tool.SetSess(context, "user_"+fmt.Sprintf("%d", member.Id), sess)
		if err != nil {
			tool.Failed(context, "登录失败")
			return
		}
		tool.Success(context, &member)
		return
	}

	tool.Failed(context, "登录失败")

}
