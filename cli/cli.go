package cli

import (
	"bufio"
	"fmt"
	"os"
	"passmgr/storage"
	"strings"
)

func manipulateMessagePrint(storageName string) {
	fmt.Printf("Storage with name %v exists, what you want to do?\n", storageName)
	fmt.Println("1. Get storage data")
	fmt.Println("2. Change storage")
	fmt.Println("3. Change password")
	fmt.Println("4. Exit")
}

// pwd - cached value of entered successful password
func manipulateStorage(storageName string, pwd string) {
	manipulateMessagePrint(storageName)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		choose := scanner.Text()
		switch choose {
		case "1":
			pwd = storage.SelectStorage(storageName, pwd)
		case "2":
			pwd = storage.ChangeStorage(storageName, pwd)
		case "3":
			pwd = storage.ChangeStoragePwd(storageName, pwd)
		case "4":
			os.Exit(1)
		default:
			fmt.Println("Please try again")
		}
		manipulateMessagePrint(storageName)
	}
}

func createStorageCall(storageName string) string {
	fmt.Printf("Storage with name %v not exists you want to create it?\n", storageName)
	fmt.Println("Write answer: (y/n)")
	scanner := bufio.NewScanner(os.Stdin)
	var pwd string
	if scanner.Scan() {
		if answer := strings.TrimSpace(scanner.Text()); answer == "y" {
			data := storage.FillStorage()
			pwd = storage.SaveStorageArr(storageName, data, "")
		} else {
			os.Exit(1)
		}
	}
	return pwd
}

func StartCli(storageName string) {
	var storageExists bool
	if _, err := os.Stat(storageName); err == nil {
		storageExists = true
	}
	var pwd string
	if !storageExists {
		pwd = createStorageCall(storageName)
	}
	manipulateStorage(storageName, pwd)
}
