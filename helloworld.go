package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func checkReadFile(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	if len(os.Args[1:]) != 2 {
		fmt.Println(" Usage: vigenere-encrypt <key> <plaintext file> \n")

	} else {

		key := os.Args[1]
		if len(key) > 32 {
			fmt.Println(" Key should be uppercase and of 32 characters \n")
		}

		plaintextFile := os.Args[2]
		fmt.Println("Plaintext file being used: " + plaintextFile)

		/* open and read plaintext file */
		data, err := ioutil.ReadFile(plaintextFile)
		checkReadFile(err)

	}
}
