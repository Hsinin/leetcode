package main

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */
// strs = ["flower","flow","flight"]
// @lc code=start

func commonPrefix(str1 string, str2 string) string {
	var min, max, res string
	// var min, max string
	if len(str1) > len(str2) {
		min, max = str2, str1
	} else {
		min, max = str1, str2
	}
	if max[:len(min)] == min {
		return min
	}
	for i := 0; i < len(min); i++ {
		if min[i] != max[i] {
			res = min[:i]
			break
		}
	}
	return res
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	comm := strs[0]
	for i := 1; i < len(strs); i++ {
		comm = commonPrefix(comm, strs[i])
	}
	return comm
}

// @lc code=end
func main() {
	// strs := []string{"flower", "flow", "flight"}
	strs := []string{}
	fmt.Println(longestCommonPrefix(strs))
}
