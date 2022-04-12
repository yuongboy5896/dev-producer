package dao

import (
	"dev-producer/model"
	"dev-producer/tool"
	"fmt"
)

type VcenterDao struct {
	*tool.Orm
}

func (Vc *VcenterDao) InsertVms(Vmlist []model.VcenterVm) int64 {
	result, err := Vc.InsertWithIGNORE(Vmlist)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return result
}
