package param

type K8sGetDeployParam struct {
	ModuleCode string `json:"modulecode"`
	NameSpace  string `json:"namespace"`
	EnvId      int64  `json:"envid"`
}
