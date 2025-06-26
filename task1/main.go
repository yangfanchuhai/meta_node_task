package main

import (
	"fmt"
	"sort"
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

// 加一
func plusOne(digits []int) []int {
	addNum := 1
	for i := len(digits) - 1; i >= 0; i-- {
		if addNum == 0 {
			break
		}

		nowNum := digits[i] + addNum
		digits[i] = nowNum % 10
		addNum = nowNum / 10
	}

	if addNum != 0 {
		digits = append([]int{addNum}, digits...)
	}
	return digits
}

// 删除排序数组中的重复项
func removeDuplicates(nums []int) int {
	var i, j int
	i = 0
	j = 1
	for ; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

// 两数之和，时间复杂度仅为o(n)，充分利用map数据结构的hash feature，以时间换空间，空间复杂度为o(n)
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		otherNum := target - nums[i]
		if v, ok := m[otherNum]; ok {
			return []int{v, otherNum}
		} else {
			m[nums[i]] = target - nums[i]
		}
	}
	return []int{}
}

// 合并区间
func merge(intervals [][]int) [][]int {
	// 排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var res [][]int = [][]int{}
	var cur []int = intervals[0]
	//res = append(res, cur)
	for i := 1; i < len(intervals); i++ {
		if cur[1] >= intervals[i][0] {
			cur[1] = max(cur[1], intervals[i][1])
		} else {
			res = append(res, cur)
			cur = intervals[i]
		}
	}
	res = append(res, cur)
	return res
}

func main() {
	intervals := [][]int{{1, 4}, {4, 5}}
	res := merge(intervals)
	fmt.Println(res)
}
