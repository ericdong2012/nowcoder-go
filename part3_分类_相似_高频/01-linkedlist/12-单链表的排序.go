package main

import "fmt"

/*
BM12 单链表的排序
https://www.nowcoder.com/practice/f23604257af94d939848729b1a5cda08?tpId=295&tqId=1008897&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295


描述
给定一个节点数为n的无序单链表，对其按升序排序。

示例1
输入：
{1,3,2,4,5}
返回值：
{1,2,3,4,5}

示例2
输入：
{-1,0,-2}
返回值：
{-2,-1,0}

*/

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func sortInList(head *ListNode) *ListNode {
	// write code here
	// 可将其拆分成两个链表，然后分别排序，最后转换成合并两个排序链表
	// 怎么找中间元素呢？我们也可以使用快慢双指针，快指针每次两步，慢指针每次一步，那么快指针到达链表尾的时候，慢指针正好走了快指针距离的一半，为中间元素。
	/*
		step 1：首先判断链表为空或者只有一个元素，直接就是有序的。
		step 2：准备三个指针，快指针right每次走两步，慢指针mid每次走一步，前序指针left每次跟在mid前一个位置。三个指针遍历链表，当快指针到达链表尾部的时候，慢指针mid刚好走了链表的一半，正好是中间位置。
		step 3：从left位置将链表断开，刚好分成两个子问题开始递归。
		step 4：将子问题得到的链表合并，参考合并两个有序链表。
	*/
	return mergeSort(head)
}

func mergeSort(head *ListNode) *ListNode {
	// 空节点
	if head == nil || head.Next == nil {
		return head
	}
	// 只有一个节点， tail == nil
	//if head.Next == tail {
	//
	//	return head
	//}

	//slow := head
	//fast := head
	//for fast != tail {
	//	fast = fast.Next
	//	slow = slow.Next
	//	if fast != tail {
	//		fast = fast.Next
	//	}
	//}
	//
	//l := mergeSort(head, slow)   //  1,     1 3 ,  5,   1 2 3
	//r := mergeSort(slow, tail)   //  3,     2,     4,   4 5

	var pre *ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 断开连接
	pre.Next = nil
	left := mergeSort(head)
	right := mergeSort(slow)
	return mergexx(left, right) //  mergexx(1, 3)
}

func mergexx(l, r *ListNode) *ListNode {
	p := &ListNode{}
	h := p
	for l != nil && r != nil {
		if l.Val < r.Val {
			h.Next = l
			l = l.Next
		} else {
			h.Next = r
			r = r.Next
		}
		h = h.Next
	}

	if l == nil {
		h.Next = r
	} else {
		h.Next = l
	}

	return p.Next
}

func main12() {
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
	result := sortInList(&node)
	fmt.Printf("%+v\n", result)
	fmt.Printf("%+v\n", result.Next)
	fmt.Printf("%+v\n", result.Next.Next)
	fmt.Printf("%+v\n", result.Next.Next.Next)
	fmt.Printf("%+v\n", result.Next.Next.Next.Next)
}
