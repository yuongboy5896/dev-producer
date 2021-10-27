package dao

import (
	"dev-producer/model"
	"dev-producer/tool"
	"fmt"
)

type PipeLineSimpleDao struct {
	*tool.Orm
}

//实例化Dao对象
func NewPipeLineSimpleDao() *PipeLineSimpleDao {
	return &PipeLineSimpleDao{tool.DbEngine}
}

//从数据库中查询所有服务器列表
func (pld *PipeLineSimpleDao) QueryPipeLineSimples() ([]model.PipeLineSimple, error) {
	var pipeLine []model.PipeLineSimple
	if err := pld.Engine.Find(&pipeLine); err != nil {
		return nil, err
	}
	return pipeLine, nil
}

//新虚拟机的数据库插入操作
func (pld *PipeLineSimpleDao) InsertPipeLineSimple(virtualMachine model.PipeLineSimple) int64 {
	result, err := pld.InsertOne(&virtualMachine)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return result
}

//查询虚拟机是否存在
func (pld *PipeLineSimpleDao) QueryByPipeLineSimples(pl model.PipeLineSimple) model.PipeLineSimple {
	var virtualMachine model.PipeLineSimple
	if _, err := pld.Where(" Pipename  = ? ", pl.Pipename).Get(&virtualMachine); err != nil {
		fmt.Println(err.Error())
	}
	return virtualMachine
}

//删除虚拟机
func (pld *PipeLineSimpleDao) DeletePipeLineSimple(pl model.PipeLineSimple) int64 {

	if _, err := pld.Where("  Pipename  = ? ", pl.Pipename).Delete(pl); err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return 1
}

//更新虚拟机
func (pld *PipeLineSimpleDao) UpdatePipeLineSimple(pl model.PipeLineSimple) int64 {

	if result, err := pld.Where("  Pipename  = ? ", pl.Pipename).Update(pl); err != nil {
		fmt.Println(err.Error(), result)
		return 0
	}
	return 1
}
