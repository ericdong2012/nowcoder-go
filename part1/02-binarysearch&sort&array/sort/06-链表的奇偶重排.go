package main

/*
NC133 链表的奇偶重排
https://www.nowcoder.com/practice/02bf49ea45cd486daa031614f9bd6fc3?tpId=117&tqId=37845&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=590&title=

描述
给定一个单链表，请设定一个函数，将链表的奇数位节点和偶数位节点分别放在一起，重排后输出。
注意是节点的编号而非节点的数值。


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

//type TreeNode struct {
//	Val  int
//	Next *ListNode
//}

func oddEvenList(head *ListNode) *ListNode {
	// write code here
	if head == nil || head.Next == nil {
		return head
	}
	even, evenHead := head.Next, head.Next
	odd := head
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}
