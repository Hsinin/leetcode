/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */
package main

import (
	"fmt"
	"math"
)

// @lc code=start

func reverse(x int) int {
	newNum := 0
	for {
		tmp := x % 10
		x = x / 10
		if newNum > int((math.Pow(2, 31)-1)/10) ||
			(newNum == int((math.Pow(2, 31)-1)/10) && tmp > 7) {
			return 0
		}
		if newNum < (-int(math.Pow(2, 31))/10) ||
			(newNum == (-int(math.Pow(2, 31))/10) && tmp < -8) {
			return 0
		}
		newNum = newNum*10 + tmp
		if x == 0 {
			break
		}
	}
	return newNum
}

// @lc code=end
func main() {
	// 2147483647
	fmt.Println(int(math.Pow(2, 31)))
	fmt.Println(reverse(8463847412))
}
