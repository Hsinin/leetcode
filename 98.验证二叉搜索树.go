/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func createNode(value int) *TreeNode {
	return &TreeNode{Val: value}
}

func midOrder(node *TreeNode) {
	if node != nil {
		midOrder(node.Left)
		fmt.Print(node.Val)
		midOrder(node.Right)
	} 
}

func midOrderStack(root *TreeNode) {
	stack := []*TreeNode{}
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack) - 1]
		fmt.Print(root.Val, " ")
		stack = stack[:len(stack) - 1]
		root = root.Right
	}
}


func isValidBST(root *TreeNode) bool {
	// var result []int
	tmp := math.MinInt64
	stack := []*TreeNode{}
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack) - 1]
		
		// if len(result) == 0 || root.Val > result[len(result)-1] {
		// 	result = append(result, root.Val)
		// } else {
		// 	return false
		// }
		if root.Val < tmp {
			return false
		}
		stack = stack[:len(stack) - 1]
		tmp = root.Val
		root = root.Right
	}
	return true
}
// @lc code=end


func main() {
	root := createNode(6)
	root.Left = createNode(3)
	root.Left.Left = createNode(1)
	root.Left.Right = createNode(4)
	root.Right = createNode(9)
	root.Right.Left = createNode(7)
	root.Right.Right = createNode(11)
	fmt.Print(isValidBST(root))
}