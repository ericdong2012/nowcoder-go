package main

import "fmt"

/*
BM14 链表的奇偶重排
https://www.nowcoder.com/practice/02bf49ea45cd486daa031614f9bd6fc3?tpId=295&tqId=1073463&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295

描述
给定一个单链表，请设定一个函数，将链表的奇数位节点和偶数位节点分别放在一起，重排后输出。
注意是节点的编号而非节点的数值

示例1
输入：
{1,2,3,4,5,6}
返回值：
{1,3,5,2,4,6}

说明：
1->2->3->4->5->6->NULL
重排后为
1->3->5->2->4->6->NULL

示例2
输入：
{1,4,6,3,7}
返回值：
{1,6,7,4,3}
说明：
1->4->6->3->7->NULL
重排后为
1->6->7->4->3->NULL
奇数位节点有1,6,7，偶数位节点有4,3。重排后为1,6,7,4,3

*/

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func oddEvenList(head *ListNode) *ListNode {
	// write code here
	if head == nil || head.Next == nil {
		return head
	}
	// 奇数  1,3,2,5,4
	oddHead := head
	// 偶数  3, 2, 5, 4
	evenHead := head.Next

	even := evenHead
	next := head.Next.Next
	count := 0
	for next != nil {
		count++
		if count%2 == 1 {
			oddHead.Next = next
			oddHead = oddHead.Next
			// 1，2，5，4
			// 1, 2, 4
		} else {
			evenHead.Next = next
			evenHead = evenHead.Next
			// 3，5，4
		}
		next = next.Next
	}

	// 偶数断开，丢掉后面多余的  3，5
	evenHead.Next = nil
	// 奇数指向偶数   1,2,4  3,5
	oddHead.Next = even

	return head
}

func main14() {
	// test
	node := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
		},
	}

	// 考虑边界情况
	//result := removeNthFromEnd(&node, 2)
	//result := removeNthFromEnd(&node, 5)
	result := oddEvenList(&node)
	fmt.Printf("%+v\n", result)
	fmt.Printf("%+v\n", result.Next)
	fmt.Printf("%+v\n", result.Next.Next)
	fmt.Printf("%+v\n", result.Next.Next.Next)
	fmt.Printf("%+v\n", result.Next.Next.Next.Next)

}
