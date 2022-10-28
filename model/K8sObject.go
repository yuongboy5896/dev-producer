package model

type Namespaces struct {
	Items []NameSpaceItem `json:"items"`
}

type NameSpaceItem struct {
	Metadata Namesmetadata   `json:"metadata"`
	Status   NameSpaceStatus `json:"status"`
}
type Namesmetadata struct {
	Name string `json:"name"`
}

type NameSpaceStatus struct {
	Phase string `json:"phase"`
}


///用于判断是否存在模块 简单后期添加  修改
type NameSpaceGetDeploy struct {
	//Metadata string   `json:"metadata"`
	Status   string `json:"status"`
}