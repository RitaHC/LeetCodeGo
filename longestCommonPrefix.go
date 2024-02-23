package main

// Write a function to find the longest common prefix string amongst an array of strings.

// If there is no common prefix, return an empty string "".

// Example 1:

// Input: strs = ["flower","flow","flight"]
// Output: "fl"
// Example 2:

// Input: strs = ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// Take the first string as a reference
	reference := strs[0]

	for i := 0; i < len(reference); i++ {
		// Iterate through each character of the reference string
		for j := 1; j < len(strs); j++ {
			// Check if the character matches in all other strings
			if i >= len(strs[j]) || reference[i] != strs[j][i] {
				return reference[:i]
			}
		}
	}

	return reference
}

func main() {
	fmt.Println("--------")
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println("--------")
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car", "dog"}))
}
