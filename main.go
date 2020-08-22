package main

import (
	"bufio"
	"fmt"
	"os"
	"passmgr/storage"
)

func main() {

	fmt.Println("Hello user, what are you want to do?")
	fmt.Println("1. Auth default storage")
	fmt.Println("2. Create new storage")
	fmt.Println("3. Use custom storage")
	fmt.Println("4. Delete storage")
	scanner := bufio.NewScanner(os.Stdin)
	var check bool
	for !check && scanner.Scan() {

		choose := scanner.Text()
		switch choose {
		case "1":
			check = true
			storage.SelectStorage("default.enc")
		case "2":
			check = true
		case "3":
			check = true
		case "4":
			check = true
		default:
			fmt.Println("Please try again")
		}

	}

}
