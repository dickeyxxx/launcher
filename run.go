package main

import (
	"os"
	"os/exec"
	"runtime"
	"syscall"
)

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
