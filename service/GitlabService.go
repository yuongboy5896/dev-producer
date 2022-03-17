package service

import (
	"bytes"
	"dev-producer/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GitlabService struct {
}

/**
 * 获取获取 gitlab 项目
 */
func (gls *GitlabService) GitlabProject(page string) ([]model.GitlabProject, error) {
	//通过http 请求
	resp, err := http.Get("https://git.thpyun.com/api/v4/projects?access_token=1nqcj74S22wx8C-TsNC-&page=" + page)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	//fmt.Println(string(body))
	//fmt.Println(resp.StatusCode)
	if resp.StatusCode == 200 {
		fmt.Println("ok")
	}
	var gitlabProjectlist []model.GitlabProject
	reader := bytes.NewReader(body)
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(&gitlabProjectlist); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return gitlabProjectlist, nil
}

/**
 * 获取获取 gitlab 项目branch
 */
func (gls *GitlabService) GitlabBranch(projectId string) ([]model.GitlabBranch, error) {
	//通过http 请求
	resp, err := http.Get("https://git.thpyun.com/api/v4/projects/" + projectId + "/repository/branches?access_token=1nqcj74S22wx8C-TsNC-")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	//fmt.Println(string(body))
	//fmt.Println(resp.StatusCode)
	if resp.StatusCode == 200 {
		fmt.Println("ok")
	}
	var gitlabBranchlist []model.GitlabBranch
	reader := bytes.NewReader(body)
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(&gitlabBranchlist); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return gitlabBranchlist, nil
}
