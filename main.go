package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
)

func parent() {
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	}
}

func child() {
	err := syscall.Mount("rootfs", "rootfs", "", syscall.MS_BIND, "")
	if err != nil {
		return
	}
	err = os.MkdirAll("rootfs/oldrootfs", 0700)
	if err != nil {
		return
	}
	err = syscall.PivotRoot("rootfs", "rootfs/oldrootfs")
	if err != nil {
		return
	}
	err = os.Chdir("/")
	if err != nil {
		return
	}
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal("Error in child")
	}
}

func main() {
	switch os.Args[1] {
	case "run":
		parent()
	case "child":
		child()
	default:
		panic("Something went wrong!")
	}
}
