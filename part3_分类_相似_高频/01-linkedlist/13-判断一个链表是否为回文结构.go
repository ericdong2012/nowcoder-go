package main

import "fmt"

/*
BM13 判断一个链表是否为回文结构
https://www.nowcoder.com/practice/3fed228444e740c8be66232ce8b87c2f?tpId=295&tqId=1008769&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295

描述
给定一个链表，请判断该链表是否为回文结构。
回文是指该字符串正序逆序完全一致。

示例1
输入：
{1}
返回值：
true

示例2
输入：
{2,1}
返回值：
false

说明：
2->1


示例3
输入：
{1,2,2,1}
返回值：
true

说明：
1->2->2->1
*/

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func isPail(head *ListNode) bool {
	// write code here
	pre := &ListNode{}
	pre.Next = head

	var tail *ListNode
	for head != nil {
		next := head.Next
		head.Next = tail
		tail = head
		head = next
	}
	// 原有的，反转的比对
	start := pre.Next
	for start != nil && tail != nil {
		if start.Val != tail.Val {
			return false
		}
		start = start.Next
		tail = tail.Next
	}
	return true
}

func main13() {
	// test
	node := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		},
	}

	// 考虑边界情况
	//result := removeNthFromEnd(&node, 2)
	//result := removeNthFromEnd(&node, 5)
	result := isPail(&node)
	fmt.Printf("%+v\n", result)

}
