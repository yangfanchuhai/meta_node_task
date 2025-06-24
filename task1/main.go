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

func main() {
	var nums = []int{2, 2, 1}
	fmt.Println(oneNumber(nums))
}
