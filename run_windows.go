package main

import (
	"os"
	"os/exec"
)

func run(path string, args []string) error {
	cmd := exec.Command(path, args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
