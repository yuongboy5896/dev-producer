package model

// 网络段IP扫描
type IpAlive struct {
	Id     int64  `xorm:"pk autoincr" json:"id"`
	Ip     string `xorm:"varchar(50)" json:"ip"`
	Port   string `xorm:"varchar(50)" json:"prot"`
	Status int    `xorm:"bigint" json:"status"` // 1为开启 0为关闭

}
