package main

import (
	"fmt"
	"strings"
)

func isPalindrome(x string) bool {
	x = strings.ReplaceAll(x, " ", "")
	x = strings.ToLower(x)

	reversed := ""

	for i := len(x) - 1; i >= 0; i-- {
		reversed += string(x[i])

	}
	return x == reversed

}
func main() {
	example := "kofi"
	if isPalindrome(example) {
		fmt.Println(example, "is a palindrome")
	} else {
		fmt.Println(example, "is not a palindrome")
	}
}
