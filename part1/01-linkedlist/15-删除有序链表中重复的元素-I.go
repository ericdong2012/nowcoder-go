package main

import "fmt"

/*
BM15 删除有序链表中重复的元素-I
https://www.nowcoder.com/practice/c087914fae584da886a0091e877f2c79?tpId=295&tqId=664&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295

描述
删除给出链表中的重复元素（链表中元素从小到大有序），使链表中的所有元素都只出现一次
例如：
给出的链表为1\to1\to21→1→2,返回1 \to 21→2.
给出的链表为1\to1\to 2 \to 3 \to 31→1→2→3→3,返回1\to 2 \to 31→2→3.

示例1
输入：
{1,1,2}
返回值：
{1,2}

示例2
输入：
{}
返回值：
{}

*/

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func deleteDuplicates1(head *ListNode) *ListNode {
	// 链表是有序的
	if head == nil {
		return head
	}
	pre := head
	for pre.Next != nil {
		if pre.Val == pre.Next.Val {
			pre.Next = pre.Next.Next
		} else {
			pre = pre.Next
		}
	}
	return head
}

func main15() {
	// test
	node := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	// 考虑边界情况
	//result := removeNthFromEnd(&node, 2)
	//result := removeNthFromEnd(&node, 5)
	result := deleteDuplicates1(&node)
	fmt.Printf("%+v\n", result)
	fmt.Printf("%+v\n", result.Next)
	fmt.Printf("%+v\n", result.Next.Next)
	fmt.Printf("%+v\n", result.Next.Next.Next)
	fmt.Printf("%+v\n", result.Next.Next.Next.Next)

}
