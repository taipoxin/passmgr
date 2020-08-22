package auth

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func engagePwd(pass string) string {
	much := 32 - len(pass)
	res := pass + strings.Repeat("*", much)
	return res
}

func CreatePwd() string {
	fmt.Println("Please, create master password for storage")
	var pwd string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		txt := scanner.Text()
		if len(txt) < 6 {
			fmt.Println("Pwd too short, please try again")
			pwd = ""
			continue
		}
		if len(txt) > 32 {
			fmt.Println("Pwd too long, please try again")
			pwd = ""
			continue
		}
		if pwd == txt {
			fmt.Println("Okey, saving...")
			break
		}
		if len(pwd) != 0 && pwd != txt {
			fmt.Println("Passwords not matching, try again")
			pwd = ""
			continue
		}
		pwd = txt
		fmt.Println("Write password again")
	}
	pwd = engagePwd(pwd)
	return pwd
}

func AuthPwd() string {
	var pwd string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		txt := scanner.Text()
		if len(txt) < 6 {
			fmt.Println("Please try again")
		}
		pwd = txt
		break
	}
	pwd = engagePwd(pwd)
	return pwd
}
