package service

import (
	"dev-producer/dao"
	"dev-producer/model"
	"dev-producer/tool"
	"time"
)

type MemberService struct {
}

//用户登录
func (ms *MemberService) Login(name string, password string) *model.Member {

	//1、使用用户名 + 密码 查询用户信息 如果存在用户 直接返回
	md := dao.MemberDao{tool.DbEngine}
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
