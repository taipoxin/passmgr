package storage

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"passmgr/aes"
	"passmgr/auth"
	"strconv"
	"strings"
)

func createStorage(storagename string) {
	file, err := os.Create(storagename)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()

}

func ReadStorage(storage string) []byte {
	dat, err := ioutil.ReadFile(storage)
	if err != nil {
		fmt.Println("Unable to read storage, creating it")
		createStorage(storage)
		return nil
	}
	return dat
}

func FillStorage() []string {
	fmt.Println("Now you can fill storage with name-pass values")
	fmt.Println("Write for example 'user pass', then press enter")
	fmt.Println("When your want to stop, write 's' and press Enter")
	var storageArr []string = make([]string, 0, 10)
	scanner := bufio.NewScanner(os.Stdin)
	// var end bool
	for scanner.Scan() {
		str := scanner.Text()
		if str == "s" {
			break
		}
		str = strings.TrimSpace(str)
		storageArr = append(storageArr, str)
	}
	fmt.Println("Your storage is:")
	for i, p := range storageArr {
		fmt.Printf("%v. %v\n", i+1, p)
	}
	return storageArr
}

func SaveStorageArr(storage string, data []string, pwd string) string {
	// arr to string casting
	var datastr string
	for _, el := range data {
		datastr += el + "\n"
	}
	datastr = datastr[:len(datastr)-2]

	return SaveStorage(storage, datastr, pwd)
}

func SaveStorage(storage string, data string, pwd string) string {
	fmt.Println("Now save you storage data")
	if pwd == "" {
		pwd = auth.CreatePwd()
	}
	encdata := aes.Encrypt(pwd, data)
	fmt.Println("Encrypted!")
	// write
	err := ioutil.WriteFile(storage, encdata, 0777)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Saved to", storage)
	return pwd
}

func AuthAndDecrypt(res []byte, pwd string) ([]byte, string) {
	var decrypted []byte
	for {
		if pwd == "" {
			fmt.Println("To read data write your pass")
			pwd = auth.AuthPwd()
		}
		data, err := aes.Decrypt(pwd, res)
		if err != nil {
			fmt.Println("Error in decryption, try again")
			pwd = ""
			continue
		}
		decrypted = data
		break
	}
	return decrypted, pwd
}

func decryptedStrToArr(decrypted []byte) []string {
	strs := strings.Split(string(decrypted), "\n")
	return strs
}

func SelectStorage(storage string, pwd string) string {
	res := ReadStorage(storage)
	if res == nil || len(res) == 0 {
		fmt.Println("Now storage is empty")
		data := FillStorage()
		SaveStorageArr(storage, data, "")
		res = ReadStorage(storage)
		fmt.Println("Now storage collect some data")
	}
	decrypted, pwd := AuthAndDecrypt(res, pwd)
	fmt.Println("This is your data:")
	strs := decryptedStrToArr(decrypted)
	for i, p := range strs {
		fmt.Printf("%v. %v\n", i+1, p)
	}
	fmt.Println("\n")
	return pwd
}

func ChangeStoragePwd(storageName string, pwd string) string {
	res := ReadStorage(storageName)
	var decrypted []byte
	for {
		if pwd == "" {
			fmt.Println("To change password write your old password")
			pwd = auth.AuthPwd()
		}
		data, err := aes.Decrypt(pwd, res)
		if err != nil {
			fmt.Println("Error in decryption, try again")
			pwd = ""
			continue
		}
		decrypted = data
		break
	}
	pwd = SaveStorage(storageName, string(decrypted), "")
	return pwd
}

func manageText() {
	fmt.Println("\nYou can manage storage using commands")
	fmt.Println("add value1 value2 - add new value to the end")
	fmt.Println("del 2 - delete 2nd string")
	fmt.Println("set 5 3 - change 5 and 3 index strings to 3 and 5")
	fmt.Println("exit s - with save")
	fmt.Println("exit w - without save")
	fmt.Println()
}

func remove(i int, s []string) ([]string, error) {
	if i >= len(s) || i < 0 {
		return nil, errors.New("Invalid index param")
	}
	// Remove the element at index i from a.
	copy(s[i:], s[i+1:])     // Shift a[i+1:] left one index.
	s[len(s)-1] = ""         // Erase last element (write zero value).
	return s[:len(s)-1], nil // Truncate slice.
}

func messageOK(strs []string) {
	fmt.Println("Operation successfully ended")
	for i, p := range strs {
		fmt.Printf("%v. %v\n", i+1, p)
	}
}

func ChangeStorage(storageName string, pwd string) string {
	res := ReadStorage(storageName)
	decrypted, pwd := AuthAndDecrypt(res, pwd)
	fmt.Println("This is current storage data:")
	strs := decryptedStrToArr(decrypted)
	for i, p := range strs {
		fmt.Printf("%v. %v\n", i+1, p)
	}

	manageText()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		choose := scanner.Text()
		args := strings.Split(choose, " ")
		if len(args) < 2 {
			fmt.Println("Please try again")
			continue
		}
		switch args[0] {
		case "add":
			args = args[1:]
			val := strings.Join(args, " ")
			strs = append(strs, val)
			messageOK(strs)

		case "del":
			pos, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("Please try again")
				continue
			}
			strs, err = remove(pos-1, strs)
			if err != nil {
				fmt.Println("Please try again")
				continue
			}
			messageOK(strs)

		case "set":
			if len(args) < 3 {
				fmt.Println("Please try again")
				continue
			}
			pos1, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("Please try again")
				continue
			}
			pos2, err := strconv.Atoi(args[2])
			if err != nil {
				fmt.Println("Please try again")
				continue
			}
			if (pos1 < 1 || pos1 > len(strs)) ||
				(pos2 < 1 || pos2 > len(strs)) {
				fmt.Println("Please try again")
				continue
			}
			strs[pos1-1], strs[pos2-1] = strs[pos2-1], strs[pos1-1]
			// cache := strs[pos1]
			// strs[pos1] = strs[pos2]
			// strs[pos2] = cache
			messageOK(strs)

		case "exit":
			if args[1] == "s" {
				pwd := SaveStorageArr(storageName, strs, pwd)
				return pwd
			}
			if args[1] == "w" {
				return pwd
			}

		default:
			fmt.Println("Please try again")
		}
		manageText()
	}
	return pwd
}
