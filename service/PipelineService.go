package service

import (
	"dev-producer/dao"
	"dev-producer/model"
	"fmt"
	"io/ioutil"
	"net/http"
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
func (mis *PipeLineService) PipeLinesByGitLabId(Id int64) ([]model.PipeLine, error) {
	//数据库操作层
	pipeLine := dao.NewPipeLineDao()
	return pipeLine.QueryPipeLinesByGitlabeID(Id)
}

/**
 * 获取获取 pipeline
 * 根据Id
 */
func (mis *PipeLineService) PipeLinesById(Id int64) (model.PipeLine, error) {
	//数据库操作层
	pipeLine := dao.NewPipeLineDao()
	return pipeLine.QueryPipeLinesByID(Id)
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

/*
* 发布pipeline
 */
func (mis *PipeLineService) PublishPipeLine(pipeLine model.PipeLine) int64 {

	// 镜像仓库
	// 通过http 请求
	resp, err := http.Get("http://192.168.48.37:8080/job/paramete_test2" +
		"/buildWithParameters?token=123456&branch=" + pipeLine.Branch + "&repository=" +
		pipeLine.SshUrlToRepo + "&images=192.168.48.36")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	fmt.Println(string(body))
	if resp.StatusCode == 200 {
		fmt.Println("ok")
	}
	// 发布程序
	return 1
}
