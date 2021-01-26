/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] 两个数组的交集 II
 */

package main

import "fmt"

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {
	tmp := make(map[int]int)
	for _, v := range nums1 {
		// if _, ok := tmp[v]; ok {
		// 	tmp[v]++
		// } else {
		// 	tmp[v] = 1
		// }
		tmp[v]++
	}
	res := []int{}
	for _, v := range nums2 {
		if tmp[v] > 0 {
			tmp[v]--
			res = append(res, v)
		}
	}
	return res
}

// @lc code=end

func main() {
	a := []int{4, 9, 5, 4, 4}
	b := []int{9, 4, 9, 8, 4, 6}
	// intersect(a, b)
	fmt.Println(intersect(a, b))
}
