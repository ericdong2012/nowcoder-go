package main

import "fmt"

/*
NC186 两两交换链表的节点
https://www.nowcoder.com/practice/71f95c23810349f782a1aa6c9bd714b4?tpId=117&tqId=39353&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D196&difficulty=undefined&judgeStatus=undefined&tags=580&title=

描述
给你一个链表，请你两两交换相邻节点，你需要真正交换节点本身，而不是修改节点的值。


两两交换示例：
链表    ：1->2->3->4
交换后  ：2->1->4->3

链表    ：1->2->3
交换后   ：2->1->3

示例1
输入：
{1,2,3,4}
返回值：
{2,1,4,3}

示例2
输入：
{1,2,3}
返回值：
{2,1,3}

*/

type ListNode1 struct {
	Val  int
	Next *ListNode1
}

func swapLinkedPair(head *ListNode1) *ListNode1 {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	// 当前节点 -> 下下个节点反转后的结果, 1 -> 4
	head.Next = swapLinkedPair(newHead.Next)
	// 下一个节点 -> 当前节点   2 -> 1
	newHead.Next = head
	return newHead
}

// 需要调试
// 解法2
//func swapLinkedPair(head *ListNode1) *ListNode1 {
//	// write code here
//	// write code here
//	if head == nil || head.Next == nil {
//		return head
//	}
//	var ret = head.Next
//	for head != nil && head.Next != nil {
//		// 将head.next.next保存下来  1,2,3,4 中的3开始的节点保存下来
//		t := head.Next.Next
//		// head的下一位 反指 向了head（ret指向了head）
//		// 1,2,3,4 中的2指向了1 , ret变成了2,1,4    后面还有一串恶心的东西
//		// 3,4 中的4反指向了3， ret 中变成了2,1,4,3  后面还有一串恶心的东西
//		head.Next.Next = head
//		// 3,4都存在
//		if t != nil && t.Next != nil {
//			// 1指向了4, 原有的head 指向了head.next.next.next
//			head.Next = t.Next
//		} else if t != nil {
//			// t没有下一位，直接指向t, 比如1,2,3  1直接指向3
//			head.Next = t
//		} else {
//			// t为空，head 断开（3后面的断开）
//			head.Next = nil
//		}
//		// head前移2位
//		// 未移动前是1,4 , 移动后head跑到了3, 现在是3,4
//		head = t
//	}
//
//	return ret
//}

func main() {
	// test
	node := ListNode1{
		Val: 1,
		Next: &ListNode1{
			Val: 2,
			Next: &ListNode1{
				Val: 3,
				Next: &ListNode1{
					Val:  4,
					Next: nil,
				},
			},
		},
	}

	result := swapLinkedPair(&node)
	fmt.Printf("%+v\n", result)
	fmt.Printf("%+v\n", result.Next)
	fmt.Printf("%+v\n", result.Next.Next)
	fmt.Printf("%+v\n", result.Next.Next.Next)
}
