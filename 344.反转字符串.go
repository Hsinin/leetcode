/*
 * @lc app=leetcode.cn id=344 lang=golang
 *
 * [344] 反转字符串
 */
package main

import "fmt"

// @lc code=start
func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		tmp := s[i]
		s[i] = s[len(s)-1-i]
		s[len(s)-1-i] = tmp
	}
}

// @lc code=end
func main() {
	// a := []byte{"h", "e", "l", "l", "o"}
	a := []byte{'h', 'e', 'l', 'l', 'o'}
	fmt.Println(a)
	reverseString(a)
	fmt.Println(a)
	fmt.Printf("%c", a[0])
}
