package main

// Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

// Symbol       Value
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000

// I can be placed before V (5) and X (10) to make 4 and 9.
// X can be placed before L (50) and C (100) to make 40 and 90.
// C can be placed before D (500) and M (1000) to make 400 and 900.

import (
	"fmt"
)

func romanToInt(s string) int {
	// Create a map of all roman number and their numerical values
	know := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	lengthOfString := len(s)
	lastElement := s[len(s)-1 : lengthOfString]
	var result int
	result = know[lastElement]
	for i := len(s) - 1; i > 0; i-- {
		// if current number is smaller than number preceeding it
		if know[s[i:i+1]] <= know[s[i-1:i]] {
			// Then add the number to the result
			result += know[s[i-1:i]]
		} else {
			// if current number is bigger than preceeding number then deduct
			result -= know[s[i-1:i]]
		}
	}
	return result

}

func main() {
	fmt.Println("----------")
	fmt.Println(romanToInt("III"))
	fmt.Println("----------")
	fmt.Println(romanToInt("LVIII"))
	fmt.Println("----------")
	fmt.Println(romanToInt("MCMXCIV"))
}
