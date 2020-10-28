// You are given two arbitrarily large numbers,
// stored one digit at a time in a slice.
// The first must be added to the second,
// and the second must be reversed before addition.
//
// The goal is to calculate the sum of those two sets of values.
//
// IMPORTANT NOTE:
// - The input can be any lengths (i.e: it can be 20+ digits long).
// - num1 and num2 can be different lengths.
//
// Sample Inputs:
// num1 = 123456
// num2 = 123456
//
// Sample Output:
// Result: 777777
//
// We would also like to see a demonstration of appropriate unit tests
// for this functionality.

package main

// Add adds to numbers represented by int slices, and returns the sum represented as a string
// The first must be added to the second, and the second must be reversed before addition.
func Add(leftNums []int, rightNums []int) string {

	/*

		To make implementation more intuitive, we actually reverse the first number instead of second number.
		We iterate over the total number slice while calculating sum from least significant to most significant
		digit. Then we simply post process. Reverse the total number, pop off any leading zeros in total,
		and preserve the state of given input numbers.
		Finally we return string representation of total number.

		leftNums =   98999  ->  99989 ->  98999
		rightNums =  10101  ->  10101 ->  10101

		totalNums =  000000 -> 001901 -> 109100

	*/

	//fmt.Println(LogVal("leftNums", leftNums))
	//fmt.Println(LogVal("rightNums", rightNums))

	// pre processing steps
	leftLen, rightLen := len(leftNums), len(rightNums)
	totalNums := make([]int, MaxInt(leftLen, rightLen)+1)
	leftNums = reverseSlice(leftNums)

	var carry int
	for i, _ := range totalNums {
		var left, right int

		// check bounds for left arr
		if i < leftLen {
			left = leftNums[i]
		}

		// check bounds for right arr
		if i < rightLen {
			right = rightNums[i]
		}

		temp := left + right + carry
		carry = temp / 10
		totalNums[i] = temp % 10
	}

	// post processing steps
	totalNums = reverseSlice(totalNums)
	totalNums = popLeadingZeros(totalNums)
	//fmt.Println(LogVal("totalNums", totalNums))

	// preserve leftNums
	leftNums = reverseSlice(leftNums)

	// string conversion
	totalStr := intSliceToString(totalNums)

	return totalStr
}

/*
func main() {
	num1 := []int{}
	num2 := []int{}

	num1length := 4
	for i := 1; i <= num1length; i++ {
		num1 = append(num1, i)
	}

	num2length := 4
	for i := 1; i <= num2length; i++ {
		num2 = append(num2, i)
	}

	fmt.Println(LogVal("num1", num1))
	fmt.Println(LogVal("num2", num2))

	result := Add(num1, num2)

	fmt.Println(LogVal("Result", result))
}
*/
