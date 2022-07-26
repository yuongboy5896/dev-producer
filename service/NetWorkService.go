package service

import (
	"dev-producer/dao"
	"dev-producer/model"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"
)

//type IpAlive struct {
//	Ip    string `json:"ip"`
//	Alive int    `json:"alive"` // 1为开启 0为关闭
//}

type NetWorkService struct {
}

var wg sync.WaitGroup

var lock sync.Mutex

func (Nws *NetWorkService) ScanIP(ip string) []model.IpAlive {

	start := time.Now()
	//ip := "192.168.48."

	ipslist := make([]model.IpAlive, 0)
	wg.Add(254)
	for i := 1; i <= 254; i++ {
		//fmt.Println(ip + strconv.Itoa(i))
		true_ip := ip + strconv.Itoa(i)
		go pingips(true_ip, &ipslist)
	}
	wg.Wait()
	
	cost := time.Since(start)
	fmt.Println("执行时间:", cost)
	//for i =0 ;ipslist
	networkDao := dao.NewNetWrokDao()
	for _, ipalive := range ipslist {
		netWrokInfo := networkDao.QueryByNetWrok(ipalive)
		if netWrokInfo.Ip == "" {
			result := networkDao.InsertNetWrokDao(&ipalive)
			if result == 0 {
				fmt.Println("插入IP相关信息失败")
			}
		} else {
			result := networkDao.UpdateNetWrok(ipalive)
			if result == 0 {
				fmt.Println("更新IP相关信息失败败")
			}
		}
	}

	return ipslist
}

func (Nws *NetWorkService) IpAlive(IP string) []model.IpAlive {

	start := time.Now()
	ipslist := make([]model.IpAlive, 0)
	pingips(IP, &ipslist)
	cost := time.Since(start)
	fmt.Println("执行时间:", cost)
	return ipslist
}

// 适用于linux mac
func pingips(ip string, ips *[]model.IpAlive) {
	var beaf = "false"
	Command := fmt.Sprintf("ping -c 1 %s  > /dev/null && echo true || echo false", ip)
	output, err := exec.Command("/bin/sh", "-c", Command).Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	real_ip := strings.TrimSpace(string(output))

	if real_ip == beaf {
		fmt.Printf("IP: %s  失败\n", ip)
		ipAlive := model.IpAlive{}
		ipAlive.Ip = ip
		ipAlive.Status = 0
		(*ips) = append((*ips), ipAlive)

	} else {
		lock.Lock()
		ipAlive := model.IpAlive{}
		ipAlive.Ip = ip
		ipAlive.Status = 1
		(*ips) = append((*ips), ipAlive)
		lock.Unlock()
		fmt.Printf("IP: %s  成功 ping通\n", ip)
	}
	wg.Done()

}
