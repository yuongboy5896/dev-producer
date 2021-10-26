package model

//员工数据结构体定义
type Member struct {
	Id           int64   `xorm:"pk autoincr" json:"id"`
	UserName     string  `xorm:"varchar(128)" json:"user_name"`
	Mobile       string  `xorm:"varchar(11)" json:"mobile"`
	Email        string  `xorm:"varchar(255)" json:"email"`
	Password     string  `xorm:"varchar(255)" json:"password"`
	RegisterTime int64   `xorm:"bigint" json:"register_time"`
	Avatar       string  `xorm:"varchar(255)" json:"avatar"`
	Balance      float64 `xorm:"double" json:"balance"`
	IsActive     int8    `xorm:"tinyint" json:"is_active"`
	BU           string  `xorm:"varchar(10)" json:"BU"`
}
