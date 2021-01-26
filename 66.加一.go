/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	a := [len(digits) + 1]int{}
	digits[len(digits)-1]++
	for i := len(digits) - 1; i >= 0; i-- {
		a[i+1] = digits[i] % 10
		if digits[i] > 9 {
			a[i]++
		}
	}
}

// @lc code=end
