/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
 */
package main

import (
	"fmt"
	// "math"
)

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createNode(value int) *TreeNode {
	return &TreeNode{Val: value}
}

func midOrder(root *TreeNode) []int {
	stack := []*TreeNode{}
	res := []int{}
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		res = append(res, root.Val)
		stack = stack[:len(stack)-1]
		root = root.Right
	}
	return res
}

func check(p, q *TreeNode) bool {
	if q == nil && p == nil {
		return true
	}
	if q == nil || p == nil {
		return false
	}
	return q.Val == p.Val && check(p.Left, q.Right) && check(p.Right, q.Left)

}
func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}

// @lc code=end

func main() {
	// root := createNode(1)
	// root.Left = createNode(2)
	// root.Left.Left = createNode(3)
	// root.Left.Right = createNode(4)
	// root.Right = createNode(2)
	// root.Right.Left = createNode(4)
	// root.Right.Right = createNode(3)

	// [1,2,2,2,null,2]
	root := createNode(1)
	root.Left = createNode(2)
	root.Left.Left = createNode(2)
	// root.Left.Right = createNode(4)
	root.Right = createNode(2)
	root.Right.Left = createNode(2)
	fmt.Println(isSymmetric(root))

}
