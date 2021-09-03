package model

type Department struct {
	Id            int64  `xorm:"pk autoincr" json:"Id"`
	AppName       string `xorm:"varchar(50)" json:"Name"`
	ImagesUrl     string `xorm:"varchar(50)" json:"ImagesUrl"`
	GitlabUrl     string `xorm:"varchar(200)" json:"GitlabUrl"`
	ContainerPort int32  `xorm:"int32" json:"ContainerPort"`
	Cmd           string `xorm:"varchar(200)" json:"Cmd"`
	Type          string `xorm:"varchar(200)" json:"Type"`
}
