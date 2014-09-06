package main

import (
	"compress/gzip"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"runtime"
	"strings"
)

func homeDir() string {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	return user.HomeDir
}

func fileExists(path string) (bool, error) {
	var err error
	if runtime.GOOS == "windows" {
		// Windows doesn't seem to like using os.Stat
		_, err = os.Open(path)
	} else {
		_, err = os.Stat(path)
	}
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func downloadGzip(url string, filepath string) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	uncompressed, err := gzip.NewReader(resp.Body)
	if err != nil {
		return err
	}
	_, err = io.Copy(out, uncompressed)
	if err != nil {
		return err
	}
	return nil
}

func makeExecutable(filepath string) error {
	if runtime.GOOS == "windows" {
		return nil
	}
	return os.Chmod(filepath, 0777)
}

func getUrlAsString(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	version, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(version)), nil
}
