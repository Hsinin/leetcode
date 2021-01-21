package main

import "fmt"

/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 */

// @lc code=start
func rotate(nums []int, k int) {
	j := k % len(nums)
	reverse(nums)
	reverse(nums[:j])
	reverse(nums[j:])
}

func reverse(array []int) {
	for i, n := 0, len(array); i < n/2; i++ {
		// tmp := array[i]
		// array[i] = array[n-1-i]
		// array[n-1-i] = tmp
		array[i], array[n-1-i] = array[n-1-i], array[i]
	}
}

// @lc code=end

func main() {
	// var a = []int{1, 2, 3, 4, 5}
	a := []int{1, 2, 3, 4, 5}
	rotate(a, 2)
	fmt.Println(a)
}
