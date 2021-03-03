/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
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

func creatList(head *ListNode, arr []int) *ListNode {
	last := head
	for i := 0; i < len(arr); i++ {
		tmp := &ListNode{Val: arr[i], Next: nil}
		last.Next = tmp
		last = last.Next
	}
	return head.Next
}

func printList(head *ListNode) {
	res := []int{}
	// for head.Next != nil {
	// 	head = head.Next
	// 	res = append(res, head.Val)
	// }
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	Println(res)
}

func appendNode(head *ListNode, val int) *ListNode {
	last := head
	for last.Next != nil {
		last = last.Next
	}
	tmp := &ListNode{Val: val, Next: nil}
	last.Next = tmp
	return head
}

func reverseList(head *ListNode) *ListNode {
	// pre := &ListNode{}
	var pre *ListNode
	cur := head
	for cur != nil {
		t := cur.Next
		cur.Next = pre
		pre = cur
		cur = t
	}
	return pre
}

func reverseList1(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	newList := head
	var tmp *ListNode
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tmp = &ListNode{Val: l1.Val}
			l1 = l1.Next
		} else {
			tmp = &ListNode{Val: l2.Val}
			l2 = l2.Next
		}
		newList.Next = tmp
		newList = newList.Next
	}
	if l1 == nil {
		newList.Next = l2
	} else {
		newList.Next = l1
	}
	return head
}

func isPalindrome(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil || fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast.Next != nil {
		fast = fast.Next
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
	return head
}

// @lc code=end
func main() {
	head := &ListNode{}
	arr := []int{1, 2, 3, 4, 5, 6}

	printList(creatList(head, arr))
	// printList(isPalindrome(head))

	// head := &ListNode{}
	// arr := []int{2, 3}
	// creatList(head, arr)

	// head1 := &ListNode{}
	// arr1 := []int{1, 2, 3, 4}
	// creatList(head1, arr1)

	// // printList(head)
	// // printList(head1)
	// printList(mergeTwoLists(head.Next, head1.Next))
}
