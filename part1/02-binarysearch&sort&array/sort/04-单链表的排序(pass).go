package main

/*
NC70 单链表的排序
https://www.nowcoder.com/practice/f23604257af94d939848729b1a5cda08?tpId=117&tqId=37817&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=590&title=

描述
给定一个节点数为n的无序单链表，对其按升序排序。

数据范围：0 < n \le 1000000<n≤100000
要求：时间复杂度 O(nlogn)O(nlogn)

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

type ListNode struct {
	Val  int
	Next *ListNode
}

// 见01-linkeelist/12

func sortInList(head *ListNode) *ListNode {
	// write code here
	return mergeSort(head)
}

func merge2(l, r *ListNode) *ListNode {
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

func mergeSort(head  *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var pre *ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	pre.Next = nil

	left := mergeSort(head)
	right := mergeSort(slow)
	return merge2(left, right)

}
