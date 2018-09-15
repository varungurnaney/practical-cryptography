// vigenere-keylength <ciphertext-filename>

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

func showUsage() {
	fmt.Print("\n* Vigenere Keylength Guessing Tool *\n\n")
	fmt.Print("A tool for guessing the keylength of vigenere cipher given the ciphertext.\n")
	fmt.Print("usage: vigenere-keylength <ciphertext-filename>\n\n")
}

func check(e error) {
	if e != nil {
		showUsage()
		panic(e)
	}
}

func sanitize(ip string) string {
	op := []rune{}
	for _, v := range ip {
		if 65 <= v && v <= 90 {
			op = append(op, v)
		} else if 97 <= v && v <= 122 {
			op = append(op, v-32)
		}
	}
	return string(op)
}

func getIndex(ctr []int, l int, v int) int {
	j := 0
	for j := 0; j < l; j++ {
		if ctr[j] == v {
			return j
		}
	}
	return j
}

func gcd(x int, y int) int {
	var n int
	for j := 1; j <= x && j <= y; j++ {
		if x%j == 0 && y%j == 0 {
			n = j
		}
	}
	return n
}

func guessKeyLength(cipherText string) int {
	cipherLen := len(cipherText)
	var ctr []int
	var sortedCtr []int
	maxKeyLen := 100
	ctr = make([]int, maxKeyLen, maxKeyLen)
	for i := 1; i < maxKeyLen; i++ {
		for j := 0; j < cipherLen; j++ {
			if (j + i) >= cipherLen {
				break
			}
			x := string(cipherText[j])
			y := string(cipherText[j+i])
			if x == y {
				ctr[i] += 1
			}
		}
	}
	ctrLen := len(ctr)
	sortedCtr = make([]int, ctrLen, ctrLen)
	for j := 0; j < ctrLen; j++ {
		sortedCtr[j] = ctr[j]
	}
	sort.Ints(sortedCtr)
	index := make([]int, 4, 4)
	for j := 0; j < 4; j++ {
		index[j] = getIndex(ctr, ctrLen, sortedCtr[len(ctr)-(j+1)])
	}
	keyLen := gcd(gcd(index[0], index[1]), gcd(index[2], index[3]))
	if keyLen == 1 {
		sort.Ints(index)
		keyLen = index[0] + 1
	}
	return keyLen
}

func main() {
	if len(os.Args) < 2 {
		showUsage()
		os.Exit(1)
	}
	cipherTextFile, err := ioutil.ReadFile(os.Args[1])
	check(err)
	cipherText := sanitize(string(cipherTextFile))
	//fmt.Print("\nKEY: " + skey)
	fmt.Print("\nCiphertext:\t" + cipherText)
	keyLength := guessKeyLength(cipherText)
	fmt.Print("\nPredicted Key Length: ", keyLength, "\n\n")
}
