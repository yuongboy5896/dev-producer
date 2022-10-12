package service

import (
	"bytes"
	"crypto/tls"
	"dev-producer/dao"
	"dev-producer/model"
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
		return nil, nil
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
