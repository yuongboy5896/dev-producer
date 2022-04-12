package model

type VirtualMachine struct {
	Id       int64  `xorm:"pk autoincr" json:"id"`
	Member   string `xorm:"varchar(11)" json:"Member"`
	Disk     int64  `xorm:"bigint" json:"Disk"`
	OS       string `xorm:"varchar(50)" json:"OS"`
	IP       string `xorm:"varchar(50)" json:"IP"`
	Hostname string `xorm:"varchar(50)" json:"Hostname"`
	Remarks  string `xorm:"varchar(100)" json:"Remarks"`
	VmID     string `xorm:"pk varchar(50)" ` //虚拟机ID
}
type VcenterVm struct {
	Vm          string `xorm:"pk varchar(50)" json:"vm"`        //虚拟机ID
	Name        string `xorm:"varchar(50)"  json:"name"`        //虚拟机名称
	Mem         int64  `xorm:"bigint" json:"memory_size_MiB"`   //虚拟机内存Mib
	Power_state string `xorm:"varchar(50)"  json:"power_state"` //服务器开机状态
	Cpu         int64  `xorm:"bigint" json:"cpu_count"`         //虚拟机cpu核数

}

type VcenterSession struct {
	Value string `json:"value"`
}
type VcenterVmValue struct {
	Value []VcenterVm `json:"value"`
}

type VirtualMachineInfo struct {
	VirtualMachine `xorm:"extends"`
	VcenterVm      `xorm:"extends"`
}

func (VirtualMachineInfo) TableName() string {
	return "VirtualMachine"
}
