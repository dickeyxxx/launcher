package main

import "os/user"

func homeDir() string {
	user, err := user.Current()
	must(err)
	return user.HomeDir
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
