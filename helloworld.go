package main

import (
	"fmt"
	"os"

)

func main() {
	if (len(os.Args[1:]) != 2) 
	{
	fmt.Println(" Usage: vigenere-encrypt <key> <plaintext filename\n")
	} 
	else 
	{
		key := os.Args[1]
		/* the assignment says that the key should be all uppercase chars less than 32 chars */
		if (len(key) > 32) {
			fmt.Println(" Key should be uppercase chars less than 32 chars \n")
		  }
	}	

}
