package main

import (
	"fmt"
	"strings"
)

// reverseSlice returns the larger of x or y.
func reverseSlice(arr []int) []int {
	for left, right := 0, len(arr)-1; left < right; left, right = left+1, right-1 {
		arr[left], arr[right] = arr[right], arr[left]
	}
	return arr
}

// intSliceToString returns string representation of a number by converting the integer slice representation to a string
func intSliceToString(nums []int) string {
	return strings.Trim(strings.Replace(fmt.Sprint(nums), " ", "", -1), "[]")
}

// maxInt returns the larger of x or y.
func maxInt(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// popLeadingZeros returns the number after popping off any leading zeros in the number
func popLeadingZeros(nums []int) []int {
	//fmt.Println(logVal("popping leading zeros from int array", nums))

	var idx int
	for i, val := range nums {
		// found the non-zero digit
		if val != 0 {
			idx = i
			break
		}
		// all zeros so far
		idx = i + 1
	}
	nums = nums[idx:]
	return nums
}

/*
// logVal is a helper function to help debug code
func logVal(arg string, val interface{}) string {
	return fmt.Sprintf("%s: \t\t %+v", arg, val)
}
*/
