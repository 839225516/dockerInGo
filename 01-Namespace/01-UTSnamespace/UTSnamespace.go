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

/*
	在linux上执行go run UTSnamespace.go

	// 输出一个当前的PID
	# echo $$
	3054

	// 在另一个shell里用 pstree -pl 查看进程关系
           ├─sshd(1226)───sshd(2995)─┬─bash(2997)───go(3034)─┬─UTSnamespace(3051)─┬─sh(3054)
           │                         │                       │                    ├─{UTSnamespace}(3052)
           │                         │                       │                    └─{UTSnamespace}(3053)
           │                         │                       ├─{go}(3035)
           │                         │                       ├─{go}(3036)
           │                         │                       ├─{go}(3038)
           │                         │                       └─{go}(3039)
           │                         └─bash(3063)───pstree(3080)

	// 由于UST namespace 对 hostname做了隔离，所有在这个环境内修改不影响外部主机
	# hostname -b dockertest
	# hostname
	dockertest
*/
