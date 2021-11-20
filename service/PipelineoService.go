package service

import (
	"dev-producer/dao"
	"dev-producer/model"
)

type PipeLineService struct {
}

/**
 * 获取获取 pipeline
 */
func (mis *PipeLineService) PipeLines() ([]model.PipeLine, error) {
	//数据库操作层
	pipeLine := dao.NewPipeLineDao()
	return pipeLine.QueryPipeLines()
}

/**
 * 获取获取 pipeline
 * 根据模块Id
 */
func (mis *PipeLineService) PipeLinesById(Id int64) ([]model.PipeLine, error) {
	//数据库操作层
	pipeLine := dao.NewPipeLineDao()
	return pipeLine.QueryPipeLines()
}

/*
* 添加pipeline
 */
func (mis *PipeLineService) AddPipeLine(pipeLine model.PipeLine) int64 {

	vmD := dao.NewPipeLineDao()

	result := vmD.InsertPipeLine(pipeLine)

	return result
}

/*
* 查询pipeline
 */
func (mis *PipeLineService) GetPipeLine(pipeLine model.PipeLine) model.PipeLine {

	vmD := dao.NewPipeLineDao()

	result := vmD.QueryByPipeLines(pipeLine)

	return result
}

/*
* 删除pipeline
 */
func (mis *PipeLineService) DeletePipeLine(pipeLine model.PipeLine) int64 {

	vmD := dao.NewPipeLineDao()

	result := vmD.DeletePipeLine(pipeLine)

	return result
}

/*
* 更新pipeline
 */
func (mis *PipeLineService) UpdatePipeLine(pipeLine model.PipeLine) int64 {

	vmD := dao.NewPipeLineDao()

	result := vmD.UpdatePipeLine(pipeLine)

	return result
}
