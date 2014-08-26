package main

import (
	"os"
	"syscall"
)

func run(path string, args []string) error {
	env := os.Environ()
	return syscall.Exec(path, args, env)
}
