/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第N个节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

import . "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type List struct {
	Head *ListNode
}

func (this *List) Append(data int) {
	//先储存数据
	node := &ListNode{Val: data}

	//判断是否为空
	if this.Head == nil {
		this.Head = node
	} else {
		list := this.Head
		for list.Next != nil {
			list = list.Next
		}
		list.Next = node
	}
}

func insertListNode(list *List, value int) {
	node := &ListNode{Val: value}
	if list.Head == nil {
		list.Head = node
	} else {
		head := list.Head
		for head.Next != nil {
			head = head.Next
		}
		head.Next = node
	}

}

func showListNode(list *List) []int {
	tmp := list.Head
	var res []int
	for tmp.Next != nil {
		res = append(res, tmp.Val)
		tmp = tmp.Next
	}
	res = append(res, tmp.Val)
	return res
}

func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	newHead := &ListNode{
		Next: head,
	}
	slow, fast := newHead, newHead
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return newHead.Next
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	tmp := &ListNode{0, head}
	slow := tmp
	fast := head
	// for {
	// 	fast = fast.Next
	// 	n--
	// 	if n == 0 {
	// 		break
	// 	}
	// }
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for ; fast != nil; fast = fast.Next {
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	// for {
	// 	slow = slow.Next
	// 	fast = fast.Next
	// 	if fast == nil {
	// 		slow.Next = slow.Next.Next
	// 		break
	// 	}
	// }
	return tmp.Next
}

func reverseList(head *ListNode) *ListNode {
	pre := &ListNode{Next: head}
	left, right := pre, head
	for right.Next != nil {
		right.Next = left
		left = left.Next
		right = right.Next
	}
	right.Next = left
	pre.Next.Next = nil
	return right
}

// @lc code=end
func main() {
	list := &List{}
	// Println(list.Head)
	insertListNode(list, 1)
	insertListNode(list, 2)
	insertListNode(list, 3)
	insertListNode(list, 4)
	insertListNode(list, 5)
	insertListNode(list, 6)
	reverseList(list.Head)
	// removeNthFromEnd1(list.Head, 6)
	Println(showListNode(list))
	// printListNode(list)
}
