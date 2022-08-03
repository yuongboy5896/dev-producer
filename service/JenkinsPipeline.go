package service

import (
	"context"
	"dev-producer/model"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/bndr/gojenkins"
)

type JenkinsService struct {
}

/**
*根据jenkins模块job 创建多个job 为了查找日志方便
 */
func (Js *JenkinsService) CreateJobFromTmp(NewJob string, JobType string, pipeline model.PipeLine) bool {

	ctx := context.Background()
	// 连接方式未封装 写到配置未做
	jenkins := gojenkins.CreateJenkins(nil, "http://192.168.48.37:8080/", "thpower", "1qaz2wsx")
	_, err := jenkins.Init(ctx)
	if err != nil {
		log.Printf("连接Jenkins失败, %v\n", err)

		return false
	}
	log.Println("Jenkins连接成功")
	// 模版上传未做
	file, err := os.Open("./config/" + JobType + ".xml")
	if err != nil {
		fmt.Println("读文件失败", err)
		return false
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("读内容失败", err)
		return false
	}
	fmt.Println(string(content))
	configString := string(content)

	configString = strings.Replace(configString, "##GITURL##", pipeline.SshUrlToRepo, -1)    //代码地址
	configString = strings.Replace(configString, "##MODULENAME##", pipeline.ModuleName, -1)  //模块中文描述
	configString = strings.Replace(configString, "##BRANCH##", pipeline.Branch, -1)          //代码分支
	configString = strings.Replace(configString, "##DEPLOY##", pipeline.ModuleCode, -1)      //模块英文名称
	configString = strings.Replace(configString, "##ENV##", pipeline.EnvName, -1)            //环境地址
	configString = strings.Replace(configString, "##NAMESPACE##", pipeline.NameSpace, -1)    //命名空间
	configString = strings.Replace(configString, "##IMAGEULR##", pipeline.ShowUrl, -1)       //上传镜像地址
	configString = strings.Replace(configString, "git.thpyun.com", "gitlab.thpower.com", -1) // 要服务器2222端口  临时
	configString = strings.Replace(configString, "192.168.48.15", "gitlab.thpower.com", -1)  // 要服务器2222端口  临时
	if 1 == pipeline.EnvCommCloud {
		str := strings.Split(pipeline.PipeCode, "-")
		if len(str) > 0 && "java" == JobType {
			configString = strings.Replace(configString, "install", "package", -1)
			configString = strings.Replace(configString, "build-portal.sh", "build-portal-"+str[0]+".sh", -1)
			configString = strings.Replace(configString, "deploy-portal.sh", "deploy-portal-"+str[0]+".sh", -1)
		} else if find := strings.Contains(JobType, "vue"); len(str) > 0 && find {
			configString = strings.Replace(configString, "build-thpws.sh", "build-thpws-"+str[0]+".sh", -1)
			configString = strings.Replace(configString, "deploy-portal.sh", "deploy-portal-"+str[0]+".sh", -1)
		}

	}

	var del bool
	getjob, err := jenkins.GetJob(ctx, NewJob)

	if getjob != nil {
		del, err = jenkins.DeleteJob(ctx, NewJob)
		if err != nil && !del {
			//panic(err)
			return false
		}

	}
	job, err := jenkins.CreateJobInFolder(ctx, configString, NewJob)
	if err != nil {

		//panic(err)
		return false
	}
	if job != nil {
		fmt.Println("")
	}
	return true

}
