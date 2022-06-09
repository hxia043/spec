package main

import (
	"aes/crypto"
	"fmt"
)

func main() {
	initKey := []byte("40xtLlrOPI7m5mGCZx9oomxfPu20PnWXpZOuXmotmnPqT39xXm2BXGinQQVi3SyQ5YBE4DbgeHOXdAHuxfj0F1DuObHD66UD4B1iUeqFXjOPaNZGNUVXV0pkE30zv3lB")
	plainText := []byte("abcdefghijklmnopqrstuvwxyz")

	key, iv := crypto.GenerateKeyandIv(initKey)

	// encrypt
	// padding block
	data := crypto.BlockPadding(plainText, len(iv))

	encryptData, err := crypto.Encrypt(data, key, iv)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(encryptData)

	// decrypt
	decryptData, err := crypto.Decrypt(encryptData, key, iv)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(decryptData))
}
