package dao

import (
	"dev-producer/model"
	"dev-producer/tool"
	"fmt"
)

type NetWrokDao struct {
	*tool.Orm
}

//插入ip信息
func (NWD *NetWrokDao) InsertNetWrokDao(netWrokInfo *model.IpAlive) int64 {
	result, err := NWD.InsertOne(netWrokInfo)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return result
}

//查询ip信息是否存在
func (ded *DeployEnvDao) QueryByNetWrok(IA model.IpAlive) model.IpAlive {
	var netWrokInfo model.IpAlive
	if _, err := ded.Where(" ip  = ? ", IA.Ip).Get(&netWrokInfo); err != nil {
		fmt.Println(err.Error())
	}
	return netWrokInfo
}

//更新ip信息
func (ded *DeployEnvDao) UpdateNetWrok(IA model.IpAlive) int64 {

	if result, err := ded.Where(" ip  = ? ", IA.Ip).Update(IA); err != nil {
		fmt.Println(err.Error(), result)
		return 0
	}
	return 1
}
