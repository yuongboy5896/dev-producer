package dao

import (
	"dev-producer/model"
	"dev-producer/tool"
	"fmt"
)

type PipeLineDao struct {
	*tool.Orm
}

//实例化Dao对象
func NewPipeLineDao() *PipeLineDao {
	return &PipeLineDao{tool.DbEngine}
}

//从数据库中查询所有服务器列表
func (pld *PipeLineDao) QueryPipeLines() ([]model.PipeLine, error) {
	var pipeLine []model.PipeLine
	if err := pld.Engine.Find(&pipeLine); err != nil {
		return nil, err
	}
	return pipeLine, nil
}

// 从数据库中查询所有服务器列表
func (pld *PipeLineDao) QueryPipeLinesByGitlabeID(GitlabId int64) ([]model.PipeLine, error) {
	var pipeLine []model.PipeLine
	if err := pld.Where("GitlabId  = ? ", GitlabId).Find(&pipeLine); err != nil {
		return nil, err
	}
	return pipeLine, nil
}

// 从数据库中查询所有服务器列表
func (pld *PipeLineDao) QueryPipeLinesByID(GitlabId int64) (model.PipeLine, error) {
	var pipeLine model.PipeLine
	if _, err := pld.Where(" Id  = ? ", GitlabId).Get(&pipeLine); err != nil {
		fmt.Println(err.Error())
	}
	return pipeLine, nil
}

//新虚拟机的数据库插入操作
func (pld *PipeLineDao) InsertPipeLine(virtualMachine model.PipeLine) int64 {
	result, err := pld.InsertOne(&virtualMachine)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return result
}

//查询分支是否存在
func (pld *PipeLineDao) QueryByPipeLines(pl model.PipeLine) model.PipeLine {
	var virtualMachine model.PipeLine
	if _, err := pld.Where(" PipeCode  = ? ", pl.PipeCode).Get(&virtualMachine); err != nil {
		fmt.Println(err.Error())
	}
	return virtualMachine
}

//删除虚拟机
func (pld *PipeLineDao) DeletePipeLine(pl model.PipeLine) int64 {

	if _, err := pld.Where("  Id  = ? ", pl.Id).Delete(pl); err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return 1
}

//更新虚拟机
func (pld *PipeLineDao) UpdatePipeLine(pl model.PipeLine) int64 {

	if result, err := pld.Where("  Pipename  = ? ", pl.Pipename).Update(pl); err != nil {
		fmt.Println(err.Error(), result)
		return 0
	}
	return 1
}
