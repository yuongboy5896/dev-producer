package service

import (
	"bytes"
	"crypto/tls"
	"dev-producer/dao"
	"dev-producer/model"
	"dev-producer/param"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetHttpsSkip(url, token string) ([]byte, error) {

	// 创建各类对象
	var client *http.Client
	var request *http.Request
	var resp *http.Response
	var body []byte
	var err error

	//`这里请注意，使用 InsecureSkipVerify: true 来跳过证书验证`
	client = &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}}

	// 获取 request请求
	request, err = http.NewRequest("GET", url, nil)

	if err != nil {
		log.Println("GetHttpSkip Request Error:", err)
		return nil, nil
	}

	// 加入 token
	request.Header.Add("Authorization", token)
	resp, err = client.Do(request)
	if err != nil {
		log.Println("GetHttpSkip Response Error:", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	defer client.CloseIdleConnections()
	return body, nil
}

type K8sApiService struct {
}

/**
 * 获取namesapces list
 */
func (kas *K8sApiService) GetNameSpaces(EnvID int64) ([]model.NameSpaceItem, error) {
	//获取环境信息
	var deployEnv = model.DeployEnv{}
	deployEnv.Id = EnvID
	vmD := dao.NewDeployEnvDao()
	result := vmD.QueryByDeployEnvs(deployEnv)
	if result.EnvKey == "" {
		log.Println("EnvKey is null ")
	}

	//通过http 请求
	body, err := GetHttpsSkip("https://"+result.EnvIP+":"+result.EnvConnPort+"/api/v1/namespaces", "Bearer "+result.EnvKey)
	if err != nil {
		fmt.Println(err)
	}

	var k8sNamespaces model.Namespaces
	reader := bytes.NewReader(body)
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(&k8sNamespaces); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return k8sNamespaces.Items, nil
}

/*
* 创建根据yaml创建
 */
func (kas *K8sApiService) CreateFromYaml(pipeline model.PipeLine) (string, error) {
	//1.获取模版信息
	var TemplateInfo model.TemplateInfo
	TemplateInfo.Id = pipeline.YamlId
	vmD := dao.NewTemplateInfoDao()

	result := vmD.QueryByIdTemplateInfo(TemplateInfo)
	//替换相关

	println(result.TemplateText)

	///创建
	return "", nil

}

/*
* 获取deploy信息 暂时 其他类型后期优化
*
 */
func (kas *K8sApiService) GetDeployInfo(DeployParam param.K8sGetDeployParam) (bool, error) {
	//获取环境信息
	var deployEnv = model.DeployEnv{}
	deployEnv.Id = int64(DeployParam.EnvId)
	vmD := dao.NewDeployEnvDao()
	result := vmD.QueryByDeployEnvs(deployEnv)
	if result.EnvKey == "" {
		log.Println("EnvKey is null ")
	}

	//通过http 请求
	//https://192.168.2.114:6443/apis/apps/v1/namespaces/kube-system/deployments/coredns/status
	body, err := GetHttpsSkip("https://"+result.EnvIP+":"+result.EnvConnPort+"/api/v1/namespaces/"+DeployParam.NameSpace+"/deployments/"+DeployParam.ModuleCode+"status", "Bearer "+result.EnvKey)
	if err != nil {
		fmt.Println(err)
	}

	var nameSpaceGetDeploy model.NameSpaceGetDeploy
	reader := bytes.NewReader(body)
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(&nameSpaceGetDeploy); err != nil {
		fmt.Println(err)
		return false, err
	}
	if nameSpaceGetDeploy.Status == "Failure" {
		return false, err
	}
	return true, nil
}
