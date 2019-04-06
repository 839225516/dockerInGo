package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

/*
	IPC namespace 是用来隔离System V IPC 和 POSIX message queues
*/

func main() {
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC,
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

/*
	在宿主机打开一个shell
	# 查看现有的ipc Message Queues
	ipcs -q
	------ Message Queues --------
	key        msqid      owner      perms      used-bytes   messages

	在另一个shell 里运行go程序
	go run IPCnamespace.go

	ipcs -q
	------ Message Queues --------
	key        msqid      owner      perms      used-bytes   messages

	# 创建一个message Queues
	ipcmk -Q

	# 然后再查看一下
	ipcs -q


	# 再在宿主机的shell 查看是否有这个message queue
	ipcs -q

*/
