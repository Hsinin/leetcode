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

func preOrder(node *Node) {
	if node != nil {
		print(node)
		preOrder(node.Left)
		preOrder(node.Right)
	}
}

func midOrder(node *Node) {
	if node != nil {
		midOrder(node.Left)
		print(node)
		midOrder(node.Right)
	}
}

func postOrder(node *Node) {
	if node != nil {
		postOrder(node.Left)
		postOrder(node.Right)
		print(node)
	}
}

// 层次遍历 BFS
func bfsOrder(node *Node) {
	nodes := []*Node{node}
	result := []int{}
	for len(nodes) > 0 {
		curNode := nodes[0]
		nodes = nodes[1:]
		result = append(result, curNode.Val)
		if curNode.Left != nil {
			nodes = append(nodes, curNode.Left)
		}
		if curNode.Right != nil {
			nodes = append(nodes, curNode.Right)
		}
	}
	for _, v := range result {
		fmt.Print(v, " ")
	}
}

func main() {
	root := createNode(6)
	root.Left = createNode(3)
	root.Left.Left = createNode(1)
	root.Left.Right = createNode(5)
	root.Right = createNode(7)
	// fmt.Println(root.Left.Right)
	preOrder(root)
	fmt.Println("")
	midOrder(root)
	fmt.Println("")
	postOrder(root)
	// print(&root)
	fmt.Println("")
	bfsOrder(root)
}

// func maxDepth(root *TreeNode) int {

// }

// @lc code=end
