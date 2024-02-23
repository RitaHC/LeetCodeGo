package main

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// An input string is valid if:

// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.

// Example 1:

// Input: s = "()"
// Output: true
// Example 2:

// Input: s = "()[]{}"
// Output: true
// Example 3:

// Input: s = "(]"
// Output: false

import (
	"fmt"
)

func isValid(s string) bool {

	// If string is empty or len is off
	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}

	// create a map of all
	// checking for a map by itself ensures that for very key, there is a value
	pairs := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	stack := []rune{}

	for _, r := range s {
		if _, ok := pairs[r]; ok {
			stack = append(stack, r)
		} else if len(stack) == 0 || pairs[stack[len(stack)-1]] != r {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0

}

func main() {
	fmt.Println("---------")
	fmt.Println(isValid("()"))
	fmt.Println("---------")
	fmt.Println(isValid("()[]{}"))
	fmt.Println("---------")
	fmt.Println(isValid("(]"))
	fmt.Println("---------")
	fmt.Println(isValid("(){}}{"))
}
