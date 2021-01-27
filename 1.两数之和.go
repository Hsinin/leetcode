package main

import "fmt"

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// 输入：nums = [2,7,11,15], target = 9
// 输出：[0,1]

// 输入：nums = [3, 3], target = 6

// @lc code=start
func twoSum(nums []int, target int) []int {
	tmp := make(map[int]int)
	for i, v := range nums {
		// 分离开数组的元素和map的value
		if key, ok := tmp[target-v]; ok {
			return []int{key, i}
		}
		tmp[v] = i
	}
	return nil
}

// @lc code=end
func main() {
	nums := []int{0, 4, 3, 1}
	fmt.Println(twoSum(nums, 0))
}
