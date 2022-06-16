package dao

import (
	"dev-producer/model"
	"dev-producer/tool"
	"fmt"
)

type PipeTemporayDao struct {
	*tool.Orm
}

//实例化Dao对象
func NewPipeTemporayDao() *PipeTemporayDao {
	return &PipeTemporayDao{tool.DbEngine}
}

//模块的镜像仓库地址
func (ptd *PipeTemporayDao) InsertModuleForImageUrl(ModuleForImageUrl model.ModuleForImageUrl) int64 {
	result, err := ptd.InsertOne(&ModuleForImageUrl)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return result
}

//查询模块的镜像仓库地址是否存在
func (ptd *PipeTemporayDao) QueryByModuleForImageUrls(mfi model.ModuleForImageUrl) model.ModuleForImageUrl {
	var ModuleForImageUrl model.ModuleForImageUrl
	if _, err := ptd.Where(" ModuleCode  = ? ", mfi.ModuleCode).Get(&ModuleForImageUrl); err != nil {
		fmt.Println(err.Error())
	}
	return ModuleForImageUrl
}

//删除模块的镜像仓库地址
func (ptd *PipeTemporayDao) DeleteModuleForImageUrl(mfi model.ModuleForImageUrl) int64 {

	if _, err := ptd.Where(" ModuleCode  = ? ", mfi.ModuleCode).Delete(mfi); err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return 1
}

//更新模块的镜像仓库地址
func (ptd *PipeTemporayDao) UpdateModuleForImageUrl(mfi model.ModuleForImageUrl) int64 {

	if result, err := ptd.Where(" ModuleCode  = ? ", mfi.ModuleCode).Update(mfi); err != nil {
		fmt.Println(err.Error(), result)
		return 0
	}
	return 1
}
