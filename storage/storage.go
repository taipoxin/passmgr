package storage

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"passmgr/aes"
	"passmgr/auth"
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

func SaveStorageArr(storage string, data []string) {
	// arr to string casting
	var datastr string
	for _, el := range data {
		datastr += el + "\n"
	}
	datastr = datastr[:len(datastr)-2]

	SaveStorage(storage, datastr)
}

func SaveStorage(storage string, data string) {
	fmt.Println("Now save you storage data")
	pwd := auth.CreatePwd()
	encdata := aes.Encrypt(pwd, data)
	fmt.Println("Encrypted!")
	// write
	err := ioutil.WriteFile(storage, encdata, 0777)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Saved to", storage)
}

func SelectStorage(storage string) {
	res := ReadStorage(storage)
	if res == nil || len(res) == 0 {
		fmt.Println("Now storage is empty")
		data := FillStorage()
		SaveStorageArr(storage, data)
		res = ReadStorage(storage)
		fmt.Println("Now storage collect some data")
	}
	var decrypted []byte
	for {
		fmt.Println("To read data write your pass")
		pwd := auth.AuthPwd()
		data, err := aes.Decrypt(pwd, res)
		if err != nil {
			fmt.Println("Error in decryption, try again")
			continue
		}
		decrypted = data
		break
	}
	fmt.Println("This is your data:")
	strs := strings.Split(string(decrypted), "\n")
	for i, p := range strs {
		fmt.Printf("%v. %v\n", i+1, p)
	}
	fmt.Println("\n")
}

func ChangeStoragePwd(storageName string) {
	res := ReadStorage(storageName)
	fmt.Println("To change password write your old password")
	var decrypted []byte
	for {
		pwd := auth.AuthPwd()
		data, err := aes.Decrypt(pwd, res)
		if err != nil {
			fmt.Println("Error in decryption, try again")
			continue
		}
		decrypted = data
		break
	}
	SaveStorage(storageName, string(decrypted))

}
