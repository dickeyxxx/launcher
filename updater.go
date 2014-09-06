package main

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"time"
)

func hkURL(channel, version string) string {
	return "https://dickeyxxx_dev.s3.amazonaws.com/hk/" + channel + "/" + version + "/hk_" + runtime.GOOS + "_" + runtime.GOARCH + ".gz"
}

func hkPath() string {
	hkPath := filepath.Join(homeDir(), ".hk", "hk")
	if runtime.GOOS == "windows" {
		hkPath = hkPath + ".exe"
	}
	return hkPath
}

func updateLogPath() string {
	return filepath.Join(homeDir(), ".hk", "update.log")
}

func isUpdateCheckNeeded() bool {
	if f, err := os.Stat(updateLogPath()); err == nil {
		return f.ModTime().Add(10 * time.Second).Before(time.Now())
	}
	return true
}

type Updater struct {
	logger *LogFile
}

func NewUpdater() (*Updater, error) {
	logger, err := NewLogFile(updateLogPath())
	if err != nil {
		logger.Println(err)
		return nil, err
	}
	return &Updater{logger}, nil
}

func (u *Updater) updateIfNeeded() error {
	channel := "release"
	updateNeeded, latest, err := u.checkForUpdate()
	if err != nil {
		u.logger.Println(err)
		return err
	}
	if updateNeeded {
		if err = u.update(channel, latest); err != nil {
			u.logger.Println(err)
			return err
		}
		u.logger.Println("Updated to", latest, "on", channel)
	}
	return nil
}

func (u *Updater) checkForUpdate() (bool, string, error) {
	u.logger.Println("Checking latest version...")
	latest, err := getUrlAsString("https://dickeyxxx_dev.s3.amazonaws.com/hk/release/VERSION")
	if err != nil {
		return false, "", err
	}
	u.logger.Println("Latest version:", latest)
	current, err := getCurrentVersion()
	if err != nil {
		u.logger.Println("Error getting latest version:", err)
	} else {
		u.logger.Println("Current version:", current)
	}
	return latest != current, latest, nil
}

func getCurrentVersion() (string, error) {
	var stdout bytes.Buffer
	cmd := exec.Command(hkPath(), "version")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}
	r := regexp.MustCompile(`hk/([\d\.]+-\w+).*`)
	version := r.FindStringSubmatch(stdout.String())[1]
	return string(version), nil
}

func (u *Updater) update(channel, version string) error {
	url := hkURL(channel, version)
	u.logger.Println("Downloading", url)
	if err := downloadGzip(url, hkPath()); err != nil {
		return err
	}
	if err := makeExecutable(hkPath()); err != nil {
		return err
	}
	return nil
}
