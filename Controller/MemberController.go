package Controller

import (
	"dev-producer/param"
	"dev-producer/service"
	"dev-producer/tool"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

type MemberController struct {
}

func (mc *MemberController) Router(engine *gin.Engine) {

	engine.GET("/api/sendcode", mc.sendSmsCode)

	engine.POST("/api/login_sms", mc.smsLogin)

	engine.GET("/api/captcha", mc.captcha)

	//获取验证码
	engine.POST("/api/vertifycha", mc.vertifyCaptcha)
	//login_pwd
	engine.POST("/api/login_pwd", mc.nameLogin)
	//头像上传
	//engine.POST("/api/upload/avator", mc.uploadAvator)

	//添加用户
	engine.POST("/api/adduser", mc.addUser)

}

//头像上传
//func (mc *MemberController) uploadAvator(context *gin.Context) {
//	return
//}

func (mc *MemberController) nameLogin(context *gin.Context) {

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

//生成验证码
func (mc *MemberController) captcha(context *gin.Context) {
	tool.GenerateCaptcha(context)
}

//验证验证码是否正确
func (mc *MemberController) vertifyCaptcha(context *gin.Context) {
	var captcha tool.CaptchaResult
	err := tool.Decode(context.Request.Body, &captcha)
	if err != nil {
		tool.Failed(context, " 参数解析失败 ")
		return
	}

	result := tool.VertifyCaptcha(captcha.Id, captcha.VertifyValue)
	if result {
		fmt.Println("验证通过")
	} else {
		fmt.Println("验证失败")
	}
}

// http://localhost:8090/api/sendcode?phone=13167582436
func (mc *MemberController) sendSmsCode(context *gin.Context) {
	//发送验证码
	phone, exist := context.GetQuery("phone")
	if !exist {
		tool.Failed(context, "参数解析失败")
		return
	}

	ms := service.MemberService{}
	isSend := ms.Sendcode(phone)
	if isSend {
		tool.Success(context, "发送成功")
		return
	}

	tool.Failed(context, "发送失败")
}

//手机号+短信  登录的方法
func (mc *MemberController) smsLogin(context *gin.Context) {

	var smsLoginParam param.SmsLoginParam
	err := tool.Decode(context.Request.Body, &smsLoginParam)

	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}

	//完成手机+验证码登录
	us := service.MemberService{}
	member := us.SmsLogin(smsLoginParam)

	if member != nil {
		sess, _ := json.Marshal(member)
		err = tool.SetSess(context, "user_"+ string(member.Id), sess)
		if err != nil {
			tool.Failed(context, "登录失败")
		}
		tool.Success(context, member)
		return
	}
	tool.Failed(context, "登录失败")
}

//用户 密码 邮箱 电话 注册 的 方法
func (mc *MemberController) addUser(context *gin.Context) {
	//1、解析用户登录传递参数
	var addUserParam param.AddUserParam
	err := tool.Decode(context.Request.Body, &addUserParam)
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}
	// 2.查看用户名是否存在
	us := service.MemberService{}
	user := us.ExistsUser(addUserParam.Name)

	if user != nil {
		tool.Failed(context, "用户已存在")
		return
	}

	//3 添加用户信息
	member := us.AddUser(addUserParam)

	if member != nil {
		sess, _ := json.Marshal(member)
		err = tool.SetSess(context, "user_"+ string(member.Id), sess)
		if err != nil {
			tool.Failed(context, "注册失败")
		}
		tool.Success(context, member)
		return
	}
	tool.Failed(context, "注册失败")

}
