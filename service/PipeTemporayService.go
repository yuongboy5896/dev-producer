package service


import (
	"dev-producer/dao"
	"dev-producer/model"
)

type PipeTemporayService struct {

}

/*
* 添加服务程序镜像地址
 */
 func (pts *PipeTemporayService) AddModuleForImageUrl(moduleForImageUrl model.ModuleForImageUrl) int64 {

	vmD := dao.NewPipeTemporayDao()

	result := vmD.InsertModuleForImageUrl(moduleForImageUrl)

	return result
}

/*
* 查询服务程序镜像地址
 */
func (pts *PipeTemporayService) GetModuleForImageUrl(moduleForImageUrl model.ModuleForImageUrl) model.ModuleForImageUrl {

	vmD := dao.NewPipeTemporayDao()

	result := vmD.QueryByModuleForImageUrls(moduleForImageUrl)

	return result
}

/*
* 更新服务程序镜像地址
 */
 func (mis *ModuleInfoService) UpdateModuleForImageUrl(moduleForImageUrl model.ModuleForImageUrl) int64 {

	vmD := dao.NewPipeTemporayDao()

	result := vmD.UpdateModuleForImageUrl(moduleForImageUrl)

	return result
}



/*
* 更新服务程序镜像地址
 */
 func (mis *PipeTemporayService) UpdateModuleInfo(moduleForImageUrl model.ModuleForImageUrl) int64 {

	vmD := dao.NewPipeTemporayDao()

	result := vmD.UpdateModuleForImageUrl(moduleForImageUrl)

	return result
}
