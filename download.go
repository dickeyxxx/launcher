package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

func hkURL() string {
	return "https://s3.amazonaws.com/dickeyxxx_dev/releases/hk_" + runtime.GOOS + "_" + runtime.GOARCH + ".gz"
}

func downloadHk(hkPath string) {
	fmt.Println("[loader.exe] hk.exe not found. Downloading hk to", hkPath)
	must(os.MkdirAll(filepath.Dir(hkPath), 0777))
	out, err := os.Create(hkPath)
	must(err)
	defer out.Close()
	if runtime.GOOS != "windows" {
		must(out.Chmod(0777))
	}
	resp, err := http.Get(hkURL())
	must(err)
	defer resp.Body.Close()
	uncompressed, err := gzip.NewReader(resp.Body)
	must(err)
	_, err = io.Copy(out, uncompressed)
	must(err)
}
