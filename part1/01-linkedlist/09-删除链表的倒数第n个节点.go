package main

import "fmt"

/*
BM9 删除链表的倒数第n个节点
https://www.nowcoder.com/practice/f95dcdafbde44b22a6d741baf71653f6?tpId=295&tqId=727&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295

描述
给定一个链表，删除链表的倒数第 n 个节点并返回链表的头指针
例如，
给出的链表为: 1→2→3→4→5, n=2
删除了链表的倒数第 n个节点之后, 链表变为 1→2→3→5


示例1
输入：
{1,2},2
返回值：
{2}

*/

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// write code here
	//if head == nil {
	//	return nil
	//}
	//node := head
	//result := head
	//begin := result
	//size := 0
	//for node != nil {
	//	size++
	//	node = node.Next
	//}
	//nSize := size - n
	//// 特殊处理当nSize==0 也就是删除头结点
	//if nSize == 0 {
	//	return result.Next
	//}
	//for i := 0; i < nSize-1; i++ {
	//	result = result.Next
	//}
	//// 删除第n个节点信息
	//result.Next = result.Next.Next
	//
	//return begin

	if n <= 0 || head == nil {
		return nil
	}

	count := 1
	pre := head
	result := pre
	for head.Next != nil {
		head = head.Next
		count += 1
	}

	if count == n {
		result = result.Next
		return result
	}

	if count < n {
		return result
	}
	// pre 往前走了count -n -1 , 此时位于 count - n
	for i := 0; i < count-n-1; i++ {
		pre = pre.Next
	}
	// 跳过中间要删除的那个
	pre.Next = pre.Next.Next

	return result
}

func main09() {
	// test
	node := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
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
	result := removeNthFromEnd(&node, 6)
	//result := FindKthToTail(&node, 2)
	//result := FindKthToTail(&node, 3)
	//result := FindKthToTail(&ListNode{}, 100)
	fmt.Printf("%+v\n", result)
	fmt.Printf("%+v\n", result.Next)
	fmt.Printf("%+v\n", result.Next.Next)
	fmt.Printf("%+v\n", result.Next.Next.Next)
	fmt.Printf("%+v\n", result.Next.Next.Next.Next)

}
