package main

import (
	"os"
	"os/exec"
	"runtime"
	"syscall"
)

func run(path string, args []string) error {
	if runtime.GOOS == "windows" {
		return runWindows(path, args)
	}
	env := os.Environ()
	return syscall.Exec(path, args, env)
}

func runWindows(path string, args []string) error {
	cmd := exec.Command(path, args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
