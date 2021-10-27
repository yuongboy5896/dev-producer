package service

import (
	"dev-producer/dao"
	"dev-producer/model"
)

type PipeLineSimpleService struct {
}

/**
 * 获取获取 pipeline
 */
func (mis *PipeLineSimpleService) PipeLineSimples() ([]model.PipeLineSimple, error) {
	//数据库操作层
	pipeLine := dao.NewPipeLineSimpleDao()
	return pipeLine.QueryPipeLineSimples()
}

/*
* 添加pipeline
 */
func (mis *PipeLineSimpleService) AddPipeLineSimple(pipeLine model.PipeLineSimple) int64 {

	vmD := dao.NewPipeLineSimpleDao()

	result := vmD.InsertPipeLineSimple(pipeLine)

	return result
}

/*
* 查询pipeline
 */
func (mis *PipeLineSimpleService) GetPipeLineSimple(pipeLine model.PipeLineSimple) model.PipeLineSimple {

	vmD := dao.NewPipeLineSimpleDao()

	result := vmD.QueryByPipeLineSimples(pipeLine)

	return result
}

/*
* 删除pipeline
 */
func (mis *PipeLineSimpleService) DeletePipeLineSimple(pipeLine model.PipeLineSimple) int64 {

	vmD := dao.NewPipeLineSimpleDao()

	result := vmD.DeletePipeLineSimple(pipeLine)

	return result
}

/*
* 更新pipeline
 */
func (mis *PipeLineSimpleService) UpdatePipeLineSimple(pipeLine model.PipeLineSimple) int64 {

	vmD := dao.NewPipeLineSimpleDao()

	result := vmD.UpdatePipeLineSimple(pipeLine)

	return result
}
