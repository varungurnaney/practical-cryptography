package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func ASCIItoString(plaintextASCII []int) {

	var output bytes.Buffer
	for index, _ := range plaintextASCII {
		output.WriteString(string(plaintextASCII[index]))
	}
	fmt.Println(output.String())

}

func encrpytVigenere(data string, key string) {

	var encipherKeyASCII = make([]int, len(key), len(key))
	var plaintextASCII = make([]int, len(data), len(data))

	for i := 0; i < len(key); i++ {
		encipherKeyASCII[i] = int(key[i])
	}

	for j := 0; j < len(data); j++ {
		for k := 0; k < len(key); k++ {

			if (j % len(key)) == k {
				plaintextASCII[j] = (((int(data[j]) - 65) + (encipherKeyASCII[k] - 65)) % 26) + 65
			}
		}
	}

	ASCIItoString(plaintextASCII[:])

}
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
		//fmt.Println("Plaintext file being used: " + plaintextFile)

		/* open and read plaintext file */
		data, err := ioutil.ReadFile(plaintextFile)
		checkReadFile(err)
		//fmt.Print(len(string(data)))

		/* Convert String to Uppercase*/
		upperKey := strings.ToUpper(string(key))
		//fmt.Print("Key being used (in uppercase): " + upperKey + "\n")

		/* keep only a-zA-Z in plaintext file lext */
		regexExpr := regexp.MustCompile("[^[:alpha:]]")
		newData := regexExpr.ReplaceAllLiteralString(strings.ToUpper(string(data)), "")

		encrpytVigenere(newData, upperKey)

	}
}
