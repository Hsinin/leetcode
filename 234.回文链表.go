package main

import . "fmt"

/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	slow := &ListNode{Next: head}
	fast := slow
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	var pre *ListNode
	cur := slow.Next
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	slow.Next = pre

	for slow.Next != nil {
		if head.Val == slow.Next.Val {
			head = head.Next
			slow = slow.Next
		} else {
			return false
		}
	}
	return true
}

func creatList(head *ListNode, arr []int) *ListNode {
	// var head *ListNode
	// last = head

	// head = &ListNode{}
	// var last *ListNode
	// last = head

	// 头指针
	// last := head
	// last.Val = arr[0]
	// for i := 1; i < len(arr); i++ {
	// 	tmp := &ListNode{Val: arr[i]}
	// 	last.Next = tmp
	// 	last = last.Next
	// }
	// return head

	// 头结点
	last := head
	for i := 0; i < len(arr); i++ {
		tmp := &ListNode{Val: arr[i]}
		last.Next = tmp
		last = last.Next
	}
	return head.Next
}

func printList(head *ListNode) {
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	Println(res)
}

// @lc code=end
func main() {
	// var head *ListNode
	// head = &ListNode{}
	head := &ListNode{}
	arr := []int{}

	head1 := creatList(head, arr)
	printList(head1)
	// Println(isPalindrome(head1).Val)
	Println(isPalindrome(head1))
	// printList(isPalindrome(head1))
	// Println(head)
}
