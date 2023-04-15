package main

import "fmt"

/*
BM2 链表内指定区间反转
https://www.nowcoder.com/practice/b58434e200a648c589ca2063f1faf58c?tpId=295&tqId=654&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295

描述
将一个节点数为 size 链表 m 位置到 n 位置之间的区间反转，要求时间复杂度 O(n)O(n)，空间复杂度 O(1)O(1)。
例如：
给出的链表为 1\to 2 \to 3 \to 4 \to 5 \to NULL1→2→3→4→5→NULL, m=2,n=4m=2,n=4,
返回 1\to 4\to 3\to 2\to 5\to NULL1→4→3→2→5→NULL.

数据范围： 链表长度 0 < size \le 10000<size≤1000，0 < m \le n \le size0<m≤n≤size，链表中每个节点的值满足 |val| \le 1000∣val∣≤1000
要求：时间复杂度 O(n)O(n) ，空间复杂度 O(n)O(n)
进阶：时间复杂度 O(n)O(n)，空间复杂度 O(1)O(1)

示例1
输入：
{1,2,3,4,5},2,4
返回值：
{1,4,3,2,5}

示例2
输入：
{5},1,1
返回值：
{5}


思路:
1. 将以前整个链表反转的几个指针，头和尾的位置做一定改变，同时记录定位的前一个位置和后一个位置
2. 反转指定区间
3. 拼接
*/

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	// write code here
	// 方式一：
	//if head == nil ||  head.Next == nil {
	//	return head
	//}
	//// 声明一个哑节点/头节点
	//dummy := &ListNode{Next: head}
	//pre := dummy
	//
	//// 拿到起始点的Pre
	//for i := 1; i < m; i++ {
	//	pre = pre.Next
	//}
	//
	//cur := pre.Next
	//for i := m; i < n; i++ {
	//	temp := cur.Next
	//	// ************begin*************
	//	// 将当前节点和后个节点反转  自己画图 1 2 3  -->  2 1 3
	//	cur.Next = temp.Next
	//	temp.Next = pre.Next
	//	// ************end***************
	//	// 拼接头
	//	pre.Next = temp
	//}
	//return dummy.Next

	// 方式二：
	if head == nil || head.Next == nil {
		return head
	}
	// 哑节点
	dummy := &ListNode{Next: head}
	// 开始和结束前的节点
	pre, end := dummy, dummy
	for i := 0; i < m-1; i++ {
		pre = pre.Next
	}

	// 开始节点
	start := pre.Next
	for i := 0; i < n; i++ {
		end = end.Next
	}
	// 结束节点的下一个节点
	next := end.Next

	end.Next = nil
	// 将开始前的节点指向反转后链表的头节点
	pre.Next = reverse(start)
	// 将开始节点(最后节点)指向反转后链表的尾节点的下一个节点
	start.Next = next

	// 返回哑节点的下一个节点，即整个链表的开头
	return dummy.Next

}

func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

func main02() {
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

	result := reverseBetween(&node, 2, 4)
	fmt.Printf("%+v\n", result)
	fmt.Printf("%+v\n", result.Next)
	fmt.Printf("%+v\n", result.Next.Next)
	fmt.Printf("%+v\n", result.Next.Next.Next)
	fmt.Printf("%+v\n", result.Next.Next.Next.Next)

}
