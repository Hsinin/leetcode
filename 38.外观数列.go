/*
 * @lc app=leetcode.cn id=38 lang=golang
 *
 * [38] 外观数列
 */
package main
import (
	"fmt"
	"strconv"
) 
// @lc code=start
func countAndSay(n int) string {
	if n == 1 {
		return "11"
	} else {
		start, end := 0, 0
		var res string
		str := countAndSay(n-1)
		for end < len(str) {
			for end < len(str) && str[start] == str[end] {
				end++
			}
			res = res + strconv.Itoa(end-start) + string(str[start])
			start = end
		}
		return res
	}
}
// @lc code=end
func main()  {
	fmt.Println(countAndSay(6))
}

