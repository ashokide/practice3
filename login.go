package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var emailregex string = "^[a-z0-9._%+\\-]+@[a-z0-9.\\-]+\\.[a-z]{2,4}$"

func isValidEmail(email string) bool {
	reg := regexp.MustCompile(emailregex)
	return reg.MatchString(email)
}

func main() {
	email := ""
	pass := ""
	fmt.Println("Enter Email : ")
	fmt.Scan(&email)
	fmt.Println("Enter Password : ")
	fmt.Scan(&pass)
	if isValidEmail(email) {
		file, err := os.OpenFile("login.txt", os.O_RDONLY, os.ModePerm)
		defer file.Close()
		if err != nil {
			fmt.Println("Error In Database")
			os.Exit(0)
		}
		data := []string{}
		buf := bufio.NewScanner(file)
		for buf.Scan() {
			data = append(data, buf.Text())
		}
		if data[0] == email {
			if data[1] == pass {
				fmt.Println("Logged In")
				fmt.Println("Welcome User.....")
			} else {
				fmt.Println("Password incorrect")
			}
		} else {
			fmt.Println("Email/User not available")
		}
	} else {
		fmt.Println("Invalid Email")
	}
}
