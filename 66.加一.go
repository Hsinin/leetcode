/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */
package main

import "fmt"

// @lc code=start
func plusOne(digits []int) []int {
	// a := [len(digits) + 1]int{}
	a := make([]int, len(digits)+1)
	flag := 0
	digits[len(digits)-1]++
	for i := len(digits) - 1; i >= 0; i-- {
		a[i+1] = digits[i] % 10
		if digits[i] > 9 {
			if i > 0 {
				digits[i-1]++
			} else {
				a[0] = 1
				flag = 1
			}
		}
	}
	if flag == 1 {
		return a
	} else {
		return a[1:]
	}
}

// @lc code=end
func main() {
	// arrayList := []int{1, 2, 3, 4}
	arrayList := []int{9, 9, 9}
	fmt.Println(plusOne(arrayList))
}
