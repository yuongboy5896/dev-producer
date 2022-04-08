package model

type VirtualMachine struct {
	Id       int64  `xorm:"pk autoincr" json:"id"`
	Member   string `xorm:"varchar(11)" json:"Member"`
	Cpu      int64  `xorm:"bigint" json:"Cpu"`
	Mem      int64  `xorm:"bigint" json:"Mem"` //兆
	Disk     int64  `xorm:"bigint" json:"Disk"`
	OS       string `xorm:"varchar(50)" json:"OS"`
	IP       string `xorm:"varchar(50)" json:"IP"`
	Hostname string `xorm:"varchar(50)" json:"Hostname"`
	Remarks  string `xorm:"varchar(100)" json:"Remarks"`
}

type VcenterSession struct {
	Value string `json:"value"`
}
type VcenterVmValue struct {
	Value []VcenterVm `json:"value"`
}

type VcenterVm struct {
	Mem         int64  `json:"memory_size_MiB"` //虚拟机内存Mib
	Vm          string `json:"vm"`              //虚拟机ID
	Name        string `json:"name"`            //虚拟机名称
	Power_state string `json:"power_state"`     //服务器开机状态
	Cpu         int64  `json:"cpu_count"`       //虚拟机cpu核数

}
