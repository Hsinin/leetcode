/*
 * @lc app=leetcode.cn id=217 lang=golang
 *
 * [217] 存在重复元素
 */

package main

import "fmt"

// @lc code=start
// func containsDuplicate(nums []int) bool {
// 	tmp := make(map[int]int)
// 	var res bool
// 	for i := 0; i < len(nums); i++ {
// 		if tmp[nums[i]] == 0 {
// 			tmp[nums[i]] = i + 1
// 			if i == len(nums) {
// 				res = false
// 			}
// 		} else {
// 			res = true
// 		}
// 	}
// 	return res
// }

func containsDuplicate(nums []int) bool {
	set := make(map[int]int)
	for i, v := range nums {
		if _, ok := set[v]; ok {
			return true
		}
		set[v] = i
	}
	return false
}

// @lc code=end
func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(containsDuplicate(a))
}
