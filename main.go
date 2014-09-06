package main

import "os"

func main() {
	if isUpdateCheckNeeded() {
		updater, err := NewUpdater()
		must(err)
		err = updater.updateIfNeeded()
		must(err)
	}
	err := run(os.Args)
	must(err)
}
