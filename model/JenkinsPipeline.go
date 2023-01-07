package model

//
//缺少端口可能是多个;
type PipeLine struct {
	Id             int64  `xorm:"pk autoincr" json:"Id"`
	Pipename       string `xorm:"varchar(50)" json:"PipeName"`                  //发布流程
	PipeCode       string `xorm:"varchar(50)" json:"PipeCode"`                  //模块code
	TechnologyType string `xorm:"varchar(50)" json:"TechnologyType"`            //发布类型 技术类型 java node go C++
	EnvName        string `xorm:"varchar(50)" json:"EnvName"`                   //发布环境名字
	EnvCode        string `xorm:"varchar(50)" json:"EnvCode"`                   //发布环境编码
	NameSpace      string `xorm:"varchar(50)" json:"NameSpace" remarks:"命名空间"`  //发布环境的命名空间
	SshUrlToRepo   string `xorm:"varchar(200)" json:"SshUrlToRepo"`             //模块的url
	Branch         string `xorm:"varchar(200)" json:"Branch"`                   //模块的分支
	GitlabId       string `xorm:"varchar(50)" json:"GitlabId"`                  //模块的gitlab id
	ModuleName     string `xorm:"varchar(50)" json:"ModuleName" `               //模块的名字
	ModuleCode     string `xorm:"varchar(50)" json:"ModuleCode" remarks:"模块编码"` //模块的编码 remarks 用属性字段对应关系
	Department     string `xorm:"varchar(50)" json:"Department" `               //模块的名字
	ShowUrl        string `xorm:"varchar(50)" json:"ShowUrl"`                   //
	YamlId         int64  `xorm:"index" json:"YamlId"`                          // 发布模块Id 取名字取得不好
	RegistryId     int64  `xorm:"index" json:"RegistryId"`                      // 镜像仓库ID
	EnvId          int64  `xorm:"index" json:"EnvId"`                           // 环境Id
	EnvCommCloud   bool   `xorm:"int" json:"EnvCommCloud"`                      //是否共有云
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
	ModuleCode     string `xorm:"varchar(50)" json:"ModuleCode" `    //模块的英文 唯一编码
	ModuleName     string `xorm:"varchar(50)" json:"ModuleName"`     //模块的名称
	GitlabUrl      string `xorm:"varchar(200)" json:"GitlabUrl"`     //模块的gitlab url
	GitlabId       string `xorm:"varchar(200)" json:"GitlabId"`      //模块的gitlab id
	TechnologyType string `xorm:"varchar(50)" json:"TechnologyType"` //模块的技术类型 node java go C++
	SshUrlToRepo   string `xorm:"varchar(200)" json:"SshUrlToRepo"`
	HttpUrlToRepo  string `xorm:"varchar(200)" json:"HttpUrlToRepo"`
	ProjectType    string `xorm:"varchar(50)" json:"ProjectType"` //项目列表 镜像仓库用 web 微服务 物联网等分类
	ModulePort     int64  `xorm:"bigint" json:"ModulePort"`       //模块的端口
	ModuleUrl      string `xorm:"varchar(200)" json:"ModuleUrl"`  //模块的上下文地址 前端需要填
}

type DeployEnv struct {
	Id           int64  `xorm:"pk autoincr" json:"Id"`
	EnvName      string `xorm:"varchar(50)" json:"EnvName"`     //环境名称
	EnvCode      string `xorm:"varchar(50)" json:"EnvCode"`     //环境名称
	EnvIP        string `xorm:"varchar(50)" json:"EnvIP"`       //环境IP ,或者通过ssh 跳板机
	EnvType      string `xorm:"varchar(50)" json:"EnvType"`     // 环境类型 测试环境 生产环境
	EnvConn      string `xorm:"varchar(50)" json:"EnvConn"`     //连接方式 api ssh
	EnvConnPort  string `xorm:"varchar(50)" json:"EnvConnPort"` //连接方式 端口
	EnvCommCloud bool   `xorm:"int" json:"EnvCommCloud"`        //是否共有云
	EnvKey       string `xorm:"Text" json:"EnvKey"`             //key
	Desc         string `xorm:"varchar(200)" json:"Desc"`       //描述
}

type JobTemplate struct {
	Id      int64  `xorm:"pk autoincr" json:"Id"`
	JobName string `xorm:"varchar(50)" json:"JobName"`  //流水线名称
	JobCode string `xorm:"varchar(50)" json:"JobCode"`  //流水线编码
	JobFile string `xorm:"varchar(100)" json:"JobFile"` //流水线模版文件  未实现。。。
}

//零时接口
type ModuleForImageUrl struct {
	Id         int64  `xorm:"pk autoincr" json:"Id"`
	ModuleName string `xorm:"varchar(50) " json:"ModuleName"`         //模块名称
	ModuleCode string `xorm:"varchar(50) NOT NULL" json:"ModuleCode"` //模块名称
	ImageUrl   string `xorm:"varchar(300)" json:"ImageUrl"`           //镜像地址
	NameSpace  string `xorm:"varchar(50)" json:"NameSpace"`           //命名空间
	DeployEnv  string `xorm:"varchar(50)" json:"DeployEnv"`           //环境地址
}

//部署详细信息
type PipeLineInfo struct {
	PipeLine     `xorm:"extends"`
	DeployEnv    `xorm:"extends"`
	TemplateInfo `xorm:"extends"`
}

//当前jenkins 项目发布测试统计报
type JenkinsJob struct {
	Class    string `xorm:"varchar(50) " json:"_class"`
	Name     string `xorm:"pk varchar(50)" json:"name"`
	Url      string `xorm:"varchar(200) " json:"url"`
	Color    string `xorm:"varchar(20)" json:"color"`
	BuildNum int64  `xorm:"bigint" json:"buildnum"`
}
