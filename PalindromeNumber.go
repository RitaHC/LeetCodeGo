package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isPalindrome(x int) bool {
	// convert int ot string
	strNum := strconv.Itoa(x)
	// convert the slice to slice
	strSlice := strings.Split(strNum, "")

	lastIndex := len(strSlice) - 1
	// Reverse the Slice
	var reverseString string
	for lastIndex >= 0 {
		// concatinate the str
		reverseString = reverseString + strSlice[lastIndex]
		lastIndex--
	}
	reverseNum, _ := strconv.Atoi(reverseString)
	// Now check if the reverse num is same as the original number
	if reverseNum == x {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println("----------")
	fmt.Println(isPalindrome(121))
	fmt.Println("----------")
	fmt.Println(isPalindrome(-121))
	fmt.Println("----------")
	fmt.Println(isPalindrome(10))
}
