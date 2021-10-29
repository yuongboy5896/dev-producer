package model

type PipeLine struct {
	Id             int64  `xorm:"pk autoincr" json:"Id"`
	Pipename       string `xorm:"varchar(50)" json:"PipeName"`       //发布流程
	PipeCode       string `xorm:"varchar(50)" json:"PipeCode"`       //模块code
	TechnologyType string `xorm:"varchar(50)" json:"TechnologyType"` //发布类型 技术类型 java node go C++
	EnvName        string `xorm:"varchar(50)" json:"EnvName"`        //发布环境名字
	SshUrlToRepo   string `xorm:"varchar(200)" json:"SshUrlToRepo"`  //模块的url
	Branch         string `xorm:"varchar(200)" json:"Branch"`        //模块的分支
	ModuleName     string `xorm:"varchar(50)" json:"ModuleName"`     //模块的名字
	ModuleCode     string `xorm:"varchar(50)" json:"ModuleCode"`     //模块的编码
	Department     string `xorm:"varchar(50)" json:"Department"`     //模块的名字
	ShowUrl        string `xorm:"varchar(50)" json:"ShowUrl"`
}

type PipeLineSimple struct {
	Id       int64  `xorm:"pk autoincr" json:"Id"`
	Pipename string `xorm:"varchar(50)" json:"PipeName"` //发布流程 //用户填写
	PipeCode string `xorm:"varchar(50)" json:"PipeCode"` //模块code  //用户填写
	Branch   string `xorm:"varchar(200)" json:"Branch"`  //模块的分支
	ShowUrl  string `xorm:"varchar(50)" json:"ShowUrl"`
	ModuleId int64  `xorm:"bigint" json:"ModuleId"`
	EnvId    int64  `xorm:"bigint" json:"EnvId"`
}

type PipeLineHistory struct {
	Id        int64  `xorm:"pk autoincr" json:"Id"`
	PipeId    int64  `xorm:"bigint" json:"PipeId"`
	ImagesUrl string `xorm:"varchar(50)" json:"ImagesUrl"`
	Time      string `xorm:"varchar(200)" json:"Time"`
}

type ModuleInfo struct {
	Id             int64  `xorm:"pk autoincr" json:"Id"`             //自增iD
	ModuleCode     string `xorm:"varchar(50)" json:"ModuleCode"`     //模块的英文 唯一编码
	ModuleName     string `xorm:"varchar(50)" json:"ModuleName"`     //模块的名称
	GitlabUrl      string `xorm:"varchar(200)" json:"GitlabUrl"`     //模块的gitlab url
	GitlabId       int64  `xorm:"varchar(200)" json:"GitlabId"`      //模块的gitlab id
	TechnologyType string `xorm:"varchar(50)" json:"TechnologyType"` //模块的技术类型 node java go C++
	SshUrlToRepo   string `xorm:"varchar(200)" json:"SshUrlToRepo"`
	HttpUrlToRepo  string `xorm:"varchar(200)" json:"HttpUrlToRepo"`
	ProjectType    string `xorm:"varchar(50)" json:"ProjectType"` //项目列表 镜像仓库用 web 微服务 物联网等分类
}

type DeployEnv struct {
	Id          int64  `xorm:"pk autoincr" json:"Id"`
	EnvName     string `xorm:"varchar(50)" json:"EnvName"`     //环境名称
	EnvIP       string `xorm:"varchar(50)" json:"EnvIP"`       //环境IP ,或者通过ssh 跳板机
	EnvType     string `xorm:"varchar(50)" json:"EnvType"`     // 环境类型 测试环境 生产环境
	EnvConn     string `xorm:"varchar(50)" json:"EnvConn"`     //连接方式 api ssh
	EnvConnPort string `xorm:"varchar(50)" json:"EnvConnPort"` //连接方式 端口
	Desc        string `xorm:"varchar(200)" json:"Desc"`       //描述
}
