package model

type VirtualMachine struct {
	Id       int64  `xorm:"pk autoincr" json:"id"`
	Member   string `xorm:"varchar(11)" json:"Member"`
	Cpu      int64  `xorm:"bigint" json:"Cpu"`
	Mem      int64  `xorm:"bigint" json:"Mem"`
	Disk     int64  `xorm:"bigint" json:"Disk"`
	OS       string `xorm:"varchar(50)" json:"OS"`
	IP       string `xorm:"varchar(50)" json:"IP"`
	Hostname string `xorm:"varchar(50)" json:"Hostname"`
	Remarks  string `xorm:"varchar(100)" json:"Remarks"`
}
