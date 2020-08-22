package main

import (
	"os"
	"passmgr/cli"
)

func main() {
	storageName := "default.enc"

	if len(os.Args) > 1 {
		storageName = os.Args[1]
	}
	cli.StartCli(storageName)
}
