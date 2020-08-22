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
	fmt.Println("2. Add value (in progress)")
	fmt.Println("3. Change storage (in progress)")
	fmt.Println("4. Change password")
	fmt.Println("5. Exit")
}

func manipulateStorage(storageName string) {
	manipulateMessagePrint(storageName)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		choose := scanner.Text()
		switch choose {
		case "1":
			storage.SelectStorage(storageName)
		case "2":
		case "3":
		case "4":
			storage.ChangeStoragePwd(storageName)
		case "5":
			os.Exit(1)
		default:
			fmt.Println("Please try again")
		}
		manipulateMessagePrint(storageName)
	}
}

func createStorageCall(storageName string) {
	fmt.Printf("Storage with name %v not exists you want to create it?\n", storageName)
	fmt.Println("Write answer: (y/n)")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		if answer := strings.TrimSpace(scanner.Text()); answer == "y" {
			data := storage.FillStorage()
			storage.SaveStorageArr(storageName, data)
		} else {
			os.Exit(1)
		}
	}
}

func StartCli(storageName string) {
	var storageExists bool
	if _, err := os.Stat(storageName); err == nil {
		storageExists = true
	}
	if !storageExists {
		createStorageCall(storageName)
	}
	manipulateStorage(storageName)
}
