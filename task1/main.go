package main

import "fmt"

//只出现一次数字

func oneNumber(nums []int) int {
	numMap := make(map[int]int)

	for _, num := range nums {
		v, ok := numMap[num]
		if ok {
			numMap[num] = v + 1
		} else {
			numMap[num] = 1
		}
	}

	for k, v := range numMap {
		if v == 1 {
			return k
		}
	}

	return -1
}

/*
*
判断是否是回文数
*/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var fx int = x

	var reverse int
	//d := 10
	for x > 0 {
		reverse = reverse*10 + x%10
		x = x / 10
	}

	fmt.Println("reverse:", reverse)

	return fx == reverse

}

func main() {
	b := isPalindrome(11211)
	fmt.Println(b)
}
