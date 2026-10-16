package main

import "fmt"

/*
BM3 链表中的节点每k个一组翻转
https://www.nowcoder.com/practice/b49c3dc907814e9bbfa8437c251b028e?tpId=295&tqId=722&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295


描述
将给出的链表中的节点每 k 个一组翻转，返回翻转后的链表
如果链表中的节点数不是 k 的倍数，将最后剩下的节点保持原样
你不能更改节点中的值，只能更改节点本身。

示例1
输入：
{1,2,3,4,5},2
返回值：
{2,1,4,3,5}

示例2
输入：
{},1
返回值：
{}

*/


// 对比看 单链表的排序
// 对比看 两两交换链表的节点
// 找到尾节点，交换
func reverseKGroup(head *ListNode, k int) *ListNode {
	// write code here
	dummy := &ListNode{Next: head}
	pre := dummy

	for head != nil {
		tail := pre
		// 滑动窗口, 大小为k，找尾节点
		for i := 0; i < k; i++ {
			// tail往前走
			tail = tail.Next
			// 长度不足k
			if tail == nil {
				return dummy.Next
			}
		}
		// 保存下一组的起始节点
		next := tail.Next
		// 拿到反转后的头和尾(之前链表的尾和头)
		head1, tail1 := myReverse(head, tail)
		// 做拼接
		pre.Next = head1
		tail1.Next = next
		// 将pre往前移动, 移动到反转后的尾节点， 将head往前移动到下一组的起始节点(pre/tail1的下一个节点)
		pre = tail1
		head = next
	}
	return dummy.Next
}

// 真正难理解的是这里
// 参考题解中的c++代码， 画图理解
func myReverse(head, tail *ListNode) (*ListNode, *ListNode) {
	//pre := &ListNode{Next: nil}
	//1. 头节点(哑节点)，而不是尾节点
	pre := tail.Next
	cur := head
	// pre 走到tail 位置，退出
	for pre != tail {
		// 保存下一位结果
		next := cur.Next
		// 将cur的下一位断开，指向pre
		cur.Next = pre
		// 将pre移动到cur, 将cur移动到next
		pre = cur
		cur = next
	}

	return tail, head
}

func main03() {
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

	result := reverseKGroup(&node, 3)
	fmt.Printf("%+v\n", result)
	fmt.Printf("%+v\n", result.Next)
	fmt.Printf("%+v\n", result.Next.Next)
	fmt.Printf("%+v\n", result.Next.Next.Next)
	fmt.Printf("%+v\n", result.Next.Next.Next.Next)

}
