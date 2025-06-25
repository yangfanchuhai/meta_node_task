package main

import (
	"fmt"
	"strings"
)

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

func validStr(str string) bool {
	prefix := "({["
	//suffix := ")}]"

	m := make(map[byte]byte)
	m[byte('(')] = byte(')')
	m[byte('{')] = byte('}')
	m[byte('[')] = byte(']')
	stack := []byte{}

	for _, s := range str {
		if strings.Contains(prefix, string(s)) {
			stack = append(stack, byte(s))
		} else {
			if len(stack) == 0 {
				return false
			}

			if m[stack[len(stack)-1]] != byte(s) {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

// 最长公共前缀
func longestCommonPrefix(strs []string) string {
	commonPrefixSlice := []uint8{}
	for i := 0; i < len(strs[0]); i++ {
		char := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != char {
				return string(commonPrefixSlice)
			}
		}
		commonPrefixSlice = append(commonPrefixSlice, char)
	}
	return string(commonPrefixSlice)
}

func main() {
	strings := []string{"flower", "afjlow", ""}
	fmt.Println(longestCommonPrefix(strings))
}
