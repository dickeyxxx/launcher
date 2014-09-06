package main

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUtil(t *testing.T) {
	Convey(".homeDir", t, func() {
		dir := homeDir()
		So(dir, ShouldEqual, os.Getenv("HOME"))
	})
}
