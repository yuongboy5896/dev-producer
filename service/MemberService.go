package service

import (
	"dev-producer/dao"
	"dev-producer/model"
	"dev-producer/param"
	"dev-producer/tool"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

type MemberService struct {
}

//用户登录
func (ms *MemberService) Login(name string, password string) *model.Member {

	//1、使用用户名 + 密码 查询用户信息 如果存在用户 直接返回
	md := dao.NewMemberDao()
	member := md.Query(name, password)
	if member.Id != 0 {
		return member
	}

	//2、用户信息不存在，作为新用户保存到数据库中
	user := model.Member{}
	user.UserName = name
	user.Password = tool.EncoderSha256(password)
	user.RegisterTime = time.Now().Unix()

	result := md.InsertMember(user)
	user.Id = result

	return &user
}

// 插入用户信息
func (ms *MemberService) AddUser(addUserParam param.AddUserParam) *model.Member {
	// 插入用户信息
	md := dao.NewMemberDao()

	user := model.Member{}
	user.UserName = addUserParam.Name
	user.Mobile = addUserParam.Mobile
	user.Email = addUserParam.Email
	user.Password = tool.EncoderSha256(addUserParam.Password)
	user.RegisterTime = time.Now().Unix()

	user.Id = md.InsertMember(user)
	return &user
}

// 查看用户名是否存在
func (ms *MemberService) ExistsUser(userName string) *model.Member {
	// 插入用户信息
	md := dao.NewMemberDao()
	member := md.QueryByName(userName)
	if member.Id != 0 {
		return member
	}
	return member
}

//用户手机号+验证码的登录
func (ms *MemberService) SmsLogin(loginparam param.SmsLoginParam) *model.Member {

	//1.获取到手机号和验证码

	//2.验证手机号+验证码是否正确
	md := dao.NewMemberDao()
	sms := md.ValidateSmsCode(loginparam.Phone, loginparam.Code)
	if sms.Id == 0 {
		return nil
	}

	//3、根据手机号member表中查询记录
	member := md.QueryByPhone(loginparam.Phone)
	if member.Id != 0 {
		return member
	}

	//4.新创建一个member记录，并保存
	user := model.Member{}
	user.UserName = loginparam.Phone
	user.Mobile = loginparam.Phone
	user.RegisterTime = time.Now().Unix()

	user.Id = md.InsertMember(user)

	return &user
}

func (ms *MemberService) Sendcode(phone string) bool {

	//1.产生一个验证码
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))

	//2.调用阿里云sdk 完成发送
	config := tool.GetConfig().Sms
	client, err := dysmsapi.NewClientWithAccessKey(config.RegionId, config.AppKey, config.AppSecret)
	if err != nil {
		//logger.Error(err.Error())
		println(err.Error())
		return false
	}

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.SignName = config.SignName
	request.TemplateCode = config.TemplateCode
	request.PhoneNumbers = phone
	par, err := json.Marshal(map[string]interface{}{
		"code": code,
	})
	request.TemplateParam = string(par)

	response, err := client.SendSms(request)
	fmt.Println(response)
	if err != nil {
		//logger.Error(err.Error())
		fmt.Println(err.Error())
		return false
	}

	//3.接收返回结果，并判断发送状态
	//短信验证码发送成功
	if response.Code == "OK" {
		//将验证码保存到数据库中
		smsCode := model.SmsCode{Phone: phone, Code: code, BizId: response.BizId, CreateTime: time.Now().Unix()}
		memberDao := dao.NewMemberDao()
		result := memberDao.InsertCode(smsCode)
		return result > 0
	}
	return false
}
