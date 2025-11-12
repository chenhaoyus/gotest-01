package main

import (
	"fmt"
)

func main() {
	d := test01("((")

	fmt.Println(d)

	fmt.Println("====================最长公共前缀========================")
	// 测试用例
	testCases := [][]string{
		{"flower", "flow", "flight"}, // 期望: "fl"
		{"dog", "racecar", "car"},    // 期望: ""
		{"apple", "apple", "apple"},  // 期望: "apple"
		{"ab", "a"},                  // 期望: "a"
		{""},                         // 期望: ""
	}

	for _, strs := range testCases {
		result := test02(strs)
		fmt.Printf("test02(%v) = \"%s\"\n", strs, result)
	}

	fmt.Println("====================加一========================")
	// 测试用例
	testCases2 := [][]int{
		{1, 2, 3},    // 期望: "[1,2,4]"
		{4, 3, 2, 1}, // 期望: "[4,3,2,2]"
		{9},          // 期望: "[1,0]"
	}

	for _, strs := range testCases2 {
		result := test03(strs)
		fmt.Printf("test03(%v) = \"%v\"\n", strs, result)
	}

	fmt.Println("====================加一========================")
	v := test04([]int{1, 1, 2})

	fmt.Println(v)

}

func test01(s string) bool {
	fmt.Println("====================字符串有效性 {} () []========================")

	if len(s)%2 != 0 {
		return false
	}

	// 定义括号映射
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	var stack []rune = make([]rune, 0)

	for _, i := range s {
		if i == '{' || i == '(' || i == '[' {
			stack = append(stack, i)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[rune(i)] {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}

	if len(stack) != 0 {
		return false
	}

	return true

}

func test02(strs []string) string {

	for i := 0; i < len(strs[0]); i++ {

		for j := 0; j < len(strs); j++ {

			if i >= len(strs[j]) || strs[j][i] != strs[0][i] {

				return strs[0][:i]
			}
		}
	}

	return strs[0]

}

func test03(digits []int) []int {

	b := []int{}

	coun := 1
	for i := len(digits) - 1; i >= 0; i-- {

		tmp := digits[i] + coun

		if tmp >= 10 {
			coun = 1
			digits[i] = tmp % 10

			if i == 0 {
				b = append(b, 1)
				for _, s := range digits {
					b = append(b, s)
				}
				return b
			}
		} else {
			digits[i] = tmp
			coun = 0
		}
	}

	return digits
}

func test04(nums []int) int {

	var b []int = make([]int, len(nums))

	copy(b, nums)

	nums = nums[0:0]

	for _, c := range b {

		isexist := false

		for _, d := range nums {
			if d == c {
				isexist = true
				break
			}
		}

		if !isexist {
			nums = append(nums, c)
		}
	}

	fmt.Println(nums)

	return len(nums)
}
