package service

import (
	"bytes"
	"crypto/tls"
	"dev-producer/dao"
	"dev-producer/model"
	"dev-producer/tool"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type VcenterService struct {
}

/**
* 获取vCenter Session
 */
func (Vcs *VcenterService) GetSession() (string, error) {
	//跳过证书认证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Timeout:   time.Second * 10,
		Transport: tr,
	}
	url := "https://192.168.49.252/rest/com/vmware/cis/session"
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return "", err
	}
	username := "administrator@vsphere.local"
	password := "Gkht@123"
	req.SetBasicAuth(username, password)
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		fmt.Println("ok")
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	var session model.VcenterSession
	reader := bytes.NewReader(body)
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(&session); err != nil {
		fmt.Println(err)
		return "", err
	}
	return session.Value, nil

}

/**
* 获取vCenter vwm列表
 */

func (Vcs *VcenterService) GetVmlist(SessionID string, states string) error {
	//跳过证书认证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Timeout:   time.Second * 10,
		Transport: tr,
	}
	url := "https://192.168.49.252/rest/vcenter/vm?filter.power_states=" + states
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("vmware-api-session-id", SessionID)
	response, err := client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		fmt.Println("ok")
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	var vcenterVmlist model.VcenterVmValue
	reader := bytes.NewReader(body)
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(&vcenterVmlist); err != nil {
		fmt.Println(err)
		return err
	}
	Vcs.AddVm(vcenterVmlist.Value)
	return nil

}

func (Vcs *VcenterService) AddVm(Vmlist []model.VcenterVm) int64 {
	vcD := dao.VcenterDao{tool.DbEngine}
	result := vcD.InsertVms(Vmlist)
	if result != 0 {

	}
	return result
}
