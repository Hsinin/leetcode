package main

import "fmt"

/*
 * @lc app=leetcode.cn id=136 lang=golang
 *
 * [136] 只出现一次的数字
 */

// @lc code=start
// func singleNumber(nums []int) int {
// 	tmp := make(map[int]int)
// 	var res int
// 	for _, v := range nums {
// 		tmp[v]++
// 	}
// 	for k, v := range tmp {
// 		if v == 1 {
// 			res = k
// 		}
// 	}
// 	return res
// }

// hash
// func singleNumber(nums []int) int {
// 	tmp := make(map[int]int)
// 	var res int
// 	for _, v := range nums {
// 		if _, ok := tmp[v]; ok {
// 			delete(tmp, v)
// 		} else {
// 			tmp[v] = 1
// 		}
// 	}
// 	for k, v := range tmp {
// 		res = k
// 	}
// 	return res
// }

func singleNumber(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res = res ^ nums[i]
	}
	return res
}

// @lc code=end
func main() {
	nums := []int{1, 1, 2, 3, 3}
	fmt.Println(singleNumber(nums))
}
