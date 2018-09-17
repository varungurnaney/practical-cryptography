package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func determineKeyLength(cipherData string) int {
	var counter = make([]int, 100, 100)
	for i := 1; i < 100; i++ {
		for j := 0; j < len(cipherData); j++ {
			if (j + i) >= len(cipherData) {
				break
			}
			temp1 := string(cipherData[j])
			temp2 := string(cipherData[j+i])
			if temp1 == temp2 {
				counter[i]++
			}
		}
	}
	var sCounter = make([]int, len(counter), len(counter))
	for index := range counter {
		sCounter[index] = counter[index]
	}
	/*for j := 0; j < len(counter); j++ {
		sCounter[j] = counter[j]
	}*/
	sort.Ints(sCounter)
	arr := make([]int, 4, 4)
	for j := 0; j < 4; j++ {
		arr[j] = getIndex(counter, len(counter), sCounter[len(counter)-(j+1)])
	}
	finalKeySize := hcf(hcf(arr[0], arr[1]), hcf(arr[2], arr[3]))
	if finalKeySize == 1 {
		sort.Ints(arr)
		finalKeySize = arr[0] + 1
	}
	return finalKeySize
}

func getIndex(counter []int, len int, val int) int {
	i := 0
	for i := 0; i < len; i++ {
		if counter[i] == val {
			return i
		}
	}
	return i
}

func hcf(num1 int, num2 int) int {
	var a int
	for i := 1; i <= num1 && i <= num2; i++ {
		if num1%i == 0 && num2%i == 0 {
			a = i
		}
	}
	return a
}

func main() {
	if len(os.Args) < 2 {
		fmt.Print("usage: vigenere-keylength <ciphertext-filename>\n\n")
	} else {
		cipherData, err := ioutil.ReadFile(os.Args[1])
		check(err)
		regexExpr := regexp.MustCompile("[^[:alpha:]]")
		newData := regexExpr.ReplaceAllLiteralString(strings.ToUpper(string(cipherData)), "")

		//fmt.Print("\nKEY: " + skey)
		keyLength := determineKeyLength(newData)
		fmt.Print("\n The  Key Length: ", keyLength, "\n\n")
	}
}
