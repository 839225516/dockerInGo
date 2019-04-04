package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

/*
	UTS namespace 主要用来隔离 nodename和domainname两个系统标识
	在 uts namespace里，每个namespace都有自己的hostname
*/

func main() {
	// 指定被fork出来的新进程内的初始命令为sh
	cmd := exec.Command("sh")

	// 使用CLONE_NEWUTS这个标识创建一个 UTS namespace
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS,
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

}
