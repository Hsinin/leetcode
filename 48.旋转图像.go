/*
 * @lc app=leetcode.cn id=48 lang=golang
 *
 * [48] 旋转图像
 */
package main

import "fmt"

// @lc code=start
func rotate(matrix [][]int) {
	// count := len(matrix[0]) - 1
	count := len(matrix)
	for i := 0; i < count/2; i++ {
		for j := i; j < count-1-i; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[count-1-j][i]
			matrix[count-1-j][i] = matrix[count-1-i][count-1-j]
			matrix[count-1-i][count-1-j] = matrix[j][count-1-i]
			matrix[j][count-1-i] = tmp
		}
	}
}

// @lc code=end
func main() {
	arr := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("%v ", arr[i][j])
		}
		fmt.Println()
	}
	rotate(arr)
	fmt.Println("****************")
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("%v ", arr[i][j])
		}
		fmt.Println()
	}
}
