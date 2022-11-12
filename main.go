package main

// 打开系统默认浏览器

import (
	"fmt"
	"os/exec"
	"project/kernel"
	"project/kernel/config"
	"project/kernel/controller"
	"runtime"
)
// 不同平台启动指令不同
var commands = map[string]string{
	"windows": "explorer",
	"darwin":  "open",
	"linux":   "xdg-open",
}




func main() {
	if runtime.GOOS != "windows" {
		panic("仅支持windows系统")
	}
	config.Init()

	go kernel.Server()

	_ = Open("http://127.0.0.1:8080")

	fmt.Println("请在浏览器打开 http://127.0.0.1:8080 ")

	<-controller.OpenChannel
}

func Open(uri string) error {
	// runtime.GOOS获取当前平台
	run, ok := commands[runtime.GOOS]
	if !ok {
		return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
	}

	cmd := exec.Command(run, uri)
	return cmd.Run()
}