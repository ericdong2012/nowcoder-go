package main

import "fmt"

/*
BM1 反转链表
https://www.nowcoder.com/practice/75e878df47f24fdc9dc3e400ec6058ca?tpId=295&tqId=23286&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295

描述：
给定一个单链表的头结点pHead(该头节点是有值的，比如在下图，它的val是1)，长度为n，反转该链表后，返回新链表的表头。

数据范围： 0\leq n\leq10000≤n≤1000

要求：空间复杂度 O(1)O(1) ，时间复杂度 O(n)O(n)

如当输入链表{1,2,3}时，经反转后，原链表变为{3,2,1}，所以对应的输出为{3,2,1}



思路：
1. 双指针 / 手动栈
	1）pre指针指向已经反转好的链表的最后一个节点，最开始没有反转，所以指向nullptr
	2）cur指针指向待反转链表的第一个节点，最开始第一个节点待反转，所以指向head
	3）nex指针指向待反转链表的第二个节点，目的是保存链表，因为cur改变指向后，后面的链表则失效了，所以需要保存

	接下来，循环执行以下三个操作
	1）nex = cur->next, 保存作用
	2）cur->next = pre 未反转链表的第一个节点的下个指针指向已反转链表的最后一个节点
	3）pre = cur， cur = nex; 指针后移，操作下一个未反转链表的第一个节点
	循环条件，当然是cur != nullptr
	循环结束后，cur当然为nullptr，所以返回pre，即为反转后的头结点

2. 递归


tips:
	看题5分钟，思考不出来, 直接看题解(讨论 + 排行， 根据语言筛选)，优先看图解，然后再看代码，最后自己实现。

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(pHead *ListNode) *ListNode {
	// 3指针
	//if pHead == nil {
	//	return pHead
	//}
	//var pre *ListNode = nil
	//var cur *ListNode = pHead
	//// 用来保存结果，最终保存的是反转后链表的头节点
	//var next *ListNode
	//for cur != nil {
	//	next = cur.Next
	//	cur.Next = pre    //将原来cur.Next 指向前一个节点，实现反转
	//	pre, cur = cur, next
	//}
	//return next

	// 最优解， 单指针
	if pHead == nil || pHead.Next == nil {
		return pHead
	}

	var newHead *ListNode
	for pHead != nil {
		pNext := pHead.Next
		pHead.Next = newHead
		newHead = pHead
		pHead = pNext
	}
	return newHead

	// 递归
	//if pHead == nil || pHead.Next ==nil{
	//	return pHead
	//}
	//
	//var tmp = ReverseList(pHead.Next)
	//pHead.Next.Next = pHead
	//pHead.Next = nil
	//return tmp
}

func main() {
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

	result := ReverseList(&node)
	fmt.Printf("%+v\n", result)
	fmt.Printf("%+v\n", result.Next)
	fmt.Printf("%+v\n", result.Next.Next)
	fmt.Printf("%+v\n", result.Next.Next.Next)
	fmt.Printf("%+v\n", result.Next.Next.Next.Next)

}
