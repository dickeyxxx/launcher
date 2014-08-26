package main

import (
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	hkPath := filepath.Join(homeDir(), ".hk", "hk")
	if runtime.GOOS == "windows" {
		hkPath = hkPath + ".exe"
	}
	exists, err := fileExists(hkPath)
	must(err)
	if !exists {
		downloadHk(hkPath)
	}
	err = run(hkPath, os.Args)
	must(err)
}
