package main

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// You can return the answer in any order.

// Example 1:

// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
// Example 2:

// Input: nums = [3,2,4], target = 6
// Output: [1,2]
// Example 3:

// Input: nums = [3,3], target = 6
// Output: [0,1]

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	// Loop over the nums slice
	for i1, num := range nums {
		desirableNum := target - num
		// fmt.Println(i, desirableNum)
		// now check if the desirable number exist in the slice
		// skip the current index from being analysized for checking the next desirable number
		for i2 := i1 + 1; i2 < len(nums); i2++ {
			num2 := nums[i2]
			if desirableNum == num2 {
				fmt.Println("Desirable Number 1: ", num, " Desirable Number 2: ", num2)
				return []int{i1, i2}
			}
			// Once the result is found, break the code
		}
	}
	return []int{}
}

func main() {
	fmt.Println("-------------")
	nums1 := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums1, 9))
	fmt.Println("-------------")
	nums2 := []int{3, 2, 4}
	fmt.Println(twoSum(nums2, 6))
	fmt.Println("-------------")
	nums3 := []int{3, 3}
	fmt.Println(twoSum(nums3, 6))
}
