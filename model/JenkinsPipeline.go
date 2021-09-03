package model

type PipeLine struct {
	Id          int64  `xorm:"pk autoincr" json:"Id"`
	Pipename    string `xorm:"varchar(50)" json:"PipeName"`
	PipeType    string `xorm:"varchar(50)" json:"PipeType"`
	Environment string `xorm:"varchar(50)" json:"Environment"`
	GitlabUrl   string `xorm:"varchar(200)" json:"GitlabUrl"`
	Branch      string `xorm:"varchar(200)" json:"Branch"`
	Module      string `xorm:"varchar(50)" json:"Module"`
	Department  string `xorm:"varchar(50)" json:"Department"`
	ShowUrl     string `xorm:"varchar(50)" json:"ShowUrl"`
}

type PipeLineHistory struct {
	Id        int64  `xorm:"pk autoincr" json:"Id"`
	PipeId    int64  `xorm:"bigint" json:"PipeId"`
	ImagesUrl string `xorm:"varchar(50)" json:"ImagesUrl"`
	Time      string `xorm:"varchar(200)" json:"Time"`
}

type ModuleInfo struct {
	Id             int64  `xorm:"pk autoincr" json:"Id"`
	ModuleCode     string `xorm:"varchar(50)" json:"ModuleCode"`
	ModuleName     string `xorm:"varchar(50)" json:"ModuleName"`
	GitlabUrl      string `xorm:"varchar(200)" json:"GitlabUrl"`
	GitlabId       int64  `xorm:"varchar(200)" json:"GitlabId"`
	TechnologyType string `xorm:"varchar(50)" json:"TechnologyType"`
	SshUrlToRepo   string `xorm:"varchar(200)" json:"SshUrlToRepo"`
	HttpUrlToRepo  string `xorm:"varchar(200)" json:"HttpUrlToRepo"`
}

type DeployEnv struct {
	Id      int64  `xorm:"pk autoincr" json:"Id"`
	EnvName string `xorm:"varchar(50)" json:"EnvName"`
	EnvIP   string `xorm:"varchar(50)" json:"EnvIP"`
	EnvType string `xorm:"varchar(50)" json:"EnvType"`
}
