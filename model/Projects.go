package model

type ProjectInfo struct {
	Projectid     int64  `xorm:"pk autoincr" json:"Projectid"`
	Projectcode   string `xorm:"varchar(50)" json:"Projectcode"`
	ProjectName   string `xorm:"varchar(30)" json:"ProjectName"`
	Projectdepson string `xorm:"varchar(50)" json:"Projectdepson"`
	Projectconfig string `xorm:"varchar(50)" json:"Projectconfig"`
	Create_time   int64  `xorm:"bigint" json:"create_time"`
}
