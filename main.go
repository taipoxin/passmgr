package main

import (
	"fmt"
	"os"
	"passmgr/cli"
)

func main() {
	fmt.Println(os.Args)
	storageName := "default.enc"

	if len(os.Args) > 1 {
		storageName = os.Args[1]
	}
	cli.StartCli(storageName)
}
