package Controller

import (
	"bufio"
	"context"
	"dev-producer/model"
	"dev-producer/tool"
	"fmt"
	"io"
	"sync"
	"time"

	"os/exec"

	"github.com/gin-gonic/gin"
)

type WebHookController struct {
}

func (webHookController *WebHookController) Router(engine *gin.Engine) {
	engine.POST("/api/webhook", webHookController.PipeLine)
}

//http://localhost:8090/api/importcluster
func (webHookController *WebHookController) PipeLine(cxt *gin.Context) {
	//1、解析集群信息传递参数
	var gilabWebRequest model.GilabWebRequest
	err := tool.Decode(cxt.Request.Body, &gilabWebRequest)
	if err != nil {
		tool.Failed(cxt, "参数解析失败")
		return
	}
	println(cxt.Request.Body)
	println("正在发布 start ")
	//cmd := exec.Command("cmd", "/C", "C:\\vue\\xidan\\IESPlatform\\update.bat")
	ctx, cancel := context.WithCancel(context.Background())
	tool.Success(cxt, "正在发布")
	go func(cancelFunc context.CancelFunc) {
		time.Sleep(30 * time.Minute)
		cancelFunc()
	}(cancel)
	if gilabWebRequest.Project_id == 129 {
		println("正在发布 ing ")
		Command(ctx, "C:\\vue\\xidan\\IESPlatform\\update.bat")
	}
	if gilabWebRequest.Project_id == 161 {
		Command(ctx, "D:\\projects2021\\vue\\iesplatform_sub\\update.bat")
	}
	println("正在发布 end ")
	tool.Success(cxt, "添加成功")
}

func read(ctx context.Context, wg *sync.WaitGroup, std io.ReadCloser) {
	reader := bufio.NewReader(std)
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			readString, err := reader.ReadString('\n')
			if err != nil || err == io.EOF {
				return
			}
			fmt.Print(readString)
		}
	}
}

func Command(ctx context.Context, cmd string) error {
	c := exec.CommandContext(ctx, "cmd", "/C", cmd) // windows
	//c := exec.CommandContext(ctx, "bash", "-c", cmd) // mac linux
	stdout, err := c.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := c.StderrPipe()
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	// 因为有2个任务, 一个需要读取stderr 另一个需要读取stdout
	wg.Add(2)
	go read(ctx, &wg, stderr)
	go read(ctx, &wg, stdout)
	// 这里一定要用start,而不是run 详情请看下面的图
	err = c.Start()
	// 等待任务结束
	wg.Wait()
	return err
}
