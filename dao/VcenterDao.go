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
	//不支持官方不支持 IGORE 方法，已修改xorm。xorm已停止更新，后期换gorm
	//result, err := Vc.InsertWithIGNORE(Vmlist)
	result, err := Vc.Insert(Vmlist)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return result
}
