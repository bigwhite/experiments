package main

import (
	"log"
	"time"

	"github.com/sevlyar/go-daemon"
)

func main() {
	cntxt := &daemon.Context{
		PidFileName: "example.pid",
		PidFilePerm: 0644,
		LogFileName: "example.log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
	}

	d, err := cntxt.Reborn()
	if err != nil {
		log.Fatal("无法运行：", err)
	}
	if d != nil {
		return
	}
	defer cntxt.Release()

	log.Print("守护进程已启动")

	// 守护进程逻辑
	for {
		// ... 执行任务 ...
		time.Sleep(time.Second * 30)
	}
}
