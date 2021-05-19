/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
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
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func print(node *Node) {
	fmt.Print(node.Val, " ")
}

func setVal(val int, node *Node) {
	if node == nil {
		fmt.Println("set value to nil node!")
	}
	node.Val = val
}

func createNode(val int) *Node {
	return &Node{Val: val}
}

func main() {
	root := Node{Val: 3}
	root.Left = &Node{}
	setVal(4, root.Left)
	fmt.Println(root.Left.Val)
}

// func maxDepth(root *TreeNode) int {

// }

// @lc code=end
