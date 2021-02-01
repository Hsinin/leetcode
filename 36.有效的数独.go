/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] 有效的数独
 */

package main

import "fmt"

// 3x3宫格序号：i/3 + (j/3)*3 => i/3*3 + j/3
// 同一行：j
// 同一列：i

// @lc code=start

// func IsContain(items []int, item int) bool {
// 	for _, eachItem := range items {
// 		if eachItem == item {
// 			return true
// 		}
// 	}
// 	return false
// }

// 方法一：二维数组
// func isValidSudoku(board [][]byte) bool {
// 	arrBox := make([][]int, 9)
// 	for i := 0; i < 9; i++ {
// 		arrBox[i] = make([]int, 9)
// 	}
// 	arrRow := make([][]int, 9)
// 	for i := 0; i < 9; i++ {
// 		arrRow[i] = make([]int, 9)
// 	}
// 	arrCol := make([][]int, 9)
// 	for i := 0; i < 9; i++ {
// 		arrCol[i] = make([]int, 9)
// 	}

// 	for i := 0; i < len(board); i++ {
// 		for j := 0; j < len(board[i]); j++ {
// 			value := int(board[i][j]) - 48
// 			if value > 0 {
// 				tmp := i/3*3 + j/3
// 				// 任一arr中存在
// 				if IsContain(arrBox[tmp], value) || IsContain(arrRow[i], value) || IsContain(arrCol[j], value) {
// 					return false
// 				} else {
// 					arrBox[tmp] = append(arrBox[tmp], value)
// 					arrRow[i] = append(arrRow[i], value)
// 					arrCol[j] = append(arrCol[j], value)
// 				}
// 			}
// 		}
// 	}
// 	return true
// }

// 方法二：hash表
func isValidSudoku(board [][]byte) bool {
	var row, col, sbox [9][9]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := board[i][j] - '1'
				index_box := (i/3)*3 + j/3
				if row[i][num] == 1 {
					return false
				} else {
					row[i][num] = 1
				}
				if col[j][num] == 1 {
					return false
				} else {
					col[j][num] = 1
				}
				if sbox[index_box][num] == 1 {
					return false
				} else {
					sbox[index_box][num] = 1
				}
			}
		}
	}
	return true
}

// 作者：BtO1bO8Y66
// 链接：https://leetcode-cn.com/problems/valid-sudoku/solution/shu-zu-mo-fang-guan-fang-jie-fa-by-bto1bo8y66/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

// @lc code=end

func main() {
	arr := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '5'},
	}
	// for i := 0; i < len(arr); i++ {
	// 	for j := 0; j < len(arr[i]); j++ {
	// 		fmt.Printf("%v ", string(arr[i][j]))
	// 	}
	// 	fmt.Println()
	// }
	fmt.Println(isValidSudoku(arr))
}
