/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start

package main

import "fmt"

// [0,1,0,3,12]
// [1,3,12,0,0]

// [1,0,1]
// [4,2,4,0,0,3,0,5,1,0]

// 快指针i -> 找非0
// 慢指针j -> 最左边的0
// func moveZeroes(nums []int) {
// 	i, j := 1, 0
// 	for {
// 		if i == len(nums) {
// 			break
// 		}
// 		if nums[i] != 0 {
// 			if nums[j] == 0 {
// 				nums[i], nums[j] = nums[j], nums[i]
// 			}
// 			j++
// 		} else {
// 			if nums[j] != 0 {
// 				j++
// 			}
// 		}
// 		i++
// 	}
// }

func moveZeroes(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

// @lc code=end

func main() {
	a := []int{4, 2, 4, 0, 0, 3, 0, 5, 1, 0}
	// a := []int{2, 1}
	// a := []int{2, 0}
	// a := []int{1, 0, 1}
	moveZeroes(a)
	fmt.Println(a)
}
