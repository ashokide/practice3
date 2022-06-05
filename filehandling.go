package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("File Handling....")
	data := []byte("Erode\nTirupur\nChennai\nSalem")
	fileInfo, err := os.Stat("cities.txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File Not Exists")
			err = os.WriteFile("cities.txt", data, os.ModeAppend)
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
			fmt.Println("File Written")
		}
	} else {
		fmt.Println("File Exists")
		fmt.Println("File Info : ")
		fmt.Println(fileInfo)
	}

}
