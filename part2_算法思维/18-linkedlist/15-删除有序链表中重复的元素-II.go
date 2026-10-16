package main

import "fmt"

/*
BM16 删除有序链表中重复的元素-II
https://www.nowcoder.com/practice/71cef9f8b5564579bf7ed93fbe0b2024?tpId=295&tqId=663&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295

描述
给出一个升序排序的链表，删除链表中的所有重复出现的元素，只保留原链表中只出现一次的元素。


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
	// res的下一位指向了head
	res.Next = head
	// res的指针，cur是干活的, 是一个活动Node
	cur := res
	for cur.Next != nil && cur.Next.Next != nil {
		// 遇到相邻两个节点值相同
		if cur.Next.Val == cur.Next.Next.Val {
			temp := cur.Next.Next.Val
			// 将所有相同的都跳过， 中心扩散
			// 1,2,2 第一次进入删除掉了第一个2， 1指向了后一个2， 第二次进入 cur.Next是后一个2，依然满足条件进入循环，cur.Next = nil， 删除后一个2，无法再次进入，跳出循环
			for cur.Next != nil && cur.Next.Val == temp {
				// 跳过cur.Next  （ cur,  cur.Next,  cur.Next.Next ）
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	//返回时去掉表头, 初始化时res的下一位指向了head
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
