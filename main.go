package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"syscall"
)

func main() {
	if isUpdateCheckNeeded() {
		updater, err := NewUpdater()
		if err != nil {
			fmt.Println("Error updating")
		}
		err = updater.updateIfNeeded()
		if err != nil {
			fmt.Println("Error updating")
		}
	}
	err := run(os.Args)
	if err != nil {
		panic(err)
	}
}

func run(args []string) error {
	if runtime.GOOS == "windows" {
		return runWindows(args)
	}
	env := os.Environ()
	return syscall.Exec(hkPath(), args, env)
}

func runWindows(args []string) error {
	cmd := exec.Command(hkPath(), args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
