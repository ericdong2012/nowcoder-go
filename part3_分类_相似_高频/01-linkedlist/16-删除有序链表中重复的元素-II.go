package main

import "fmt"

/*
BM16 删除有序链表中重复的元素-II
https://www.nowcoder.com/practice/71cef9f8b5564579bf7ed93fbe0b2024?tpId=295&tqId=663&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295


描述
给出一个升序排序的链表，删除链表中的所有重复出现的元素，只保留原链表中只出现一次的元素。
例如：
给出的链表为1 \to 2\to 3\to 3\to 4\to 4\to51→2→3→3→4→4→5, 返回1\to 2\to51→2→5.
给出的链表为1\to1 \to 1\to 2 \to 31→1→1→2→3, 返回2\to 32→3.

示例1
输入：
{1,2,2}
返回值：
{1}

示例2
输入：
{}
返回值：
{}

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	// write code here
	if head == nil {
		return head
	}
	res := &ListNode{}
	// 保留原有的
	res.Next = head
	cur := res
	for cur.Next != nil && cur.Next.Next != nil {
		//遇到相邻两个节点值相同
		if cur.Next.Val == cur.Next.Next.Val {
			temp := cur.Next.Next.Val
			//将所有相同的都跳过
			for cur.Next != nil && cur.Next.Val == temp {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	//返回时去掉表头
	return res.Next
}

func main16() {
	// test
	node := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
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
	//dummyult := removeNthFromEnd(&node, 2)
	//dummyult := removeNthFromEnd(&node, 5)
	dummyult := deleteDuplicates(&node)
	fmt.Printf("%+v\n", dummyult)
	fmt.Printf("%+v\n", dummyult.Next)
	fmt.Printf("%+v\n", dummyult.Next.Next)
	fmt.Printf("%+v\n", dummyult.Next.Next.Next)
	fmt.Printf("%+v\n", dummyult.Next.Next.Next.Next)

}
