package main

import (
	"PasswordGenerate/generate"
	"PasswordGenerate/leng"
	"fmt"
)

func main() {
	length, err := leng.GetLength()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	password := generate.GeneratePassword(length)
	fmt.Println("Password Generated:", password)
}
