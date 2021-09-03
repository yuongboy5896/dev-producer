package param

type VirtualMachineParam struct {
	Member   string `json:"name"`     //用户名
	Cpu      int64  `json:"cpu"`      //处理器
	Mem      int64  `json:"mem"`      // 内存
	Disk     int64  `json:"disk"`     // 磁盘
	OS       string `json:"os"`       // 操作系统
	IP       string `json:"ip"`       // ip
	Hostname string `json:"hostname"` // 服务器hostname
	Remarks  string `json:"remarks"`  // 备注
}
