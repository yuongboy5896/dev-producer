package param

type LoginParam struct {
	Name     string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
	Id       string `json:"id"`       // 验证码id
	Value    string `json:"value"`    // 验证码输入值
}

type AddUserParam struct {
	Name     string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
	Email    string `json:"email"`    // 邮箱
	Mobile   string `json:"mobile"`   // 电话
}
