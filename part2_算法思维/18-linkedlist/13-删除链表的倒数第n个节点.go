package main

import "fmt"

/*
BM9 删除链表的倒数第n个节点
https://www.nowcoder.com/practice/f95dcdafbde44b22a6d741baf71653f6?tpId=295&tqId=727&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295

描述
给定一个链表, 删除链表的倒数第 n 个节点并返回链表的头指针
例如
给出的链表为: 1→2→3→4→5, n=2
删除了链表的倒数第n个节点之后, 链表变为 1→2→3→5


示例1
输入：
{1,2}, 2
返回值：
{2}

*/

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

// 解法2
// 去leetcode上写，nowcoder上已经没有这题了
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n <= 0 || head == nil {
		return nil
	}
	// 需要先将head保存下来，后面会做移动操作
	pre := head
	result := head
	// 统计节点长度
	count := 1
	for pre.Next != nil {
		pre = pre.Next
		count += 1
	}
	// 长度等于指定的数
	if count == n {
		return head.Next
	}
	// 长度小于指定的数
	if count < n {
		return head
	}
	// 长度大于指定的数
	// 找要删除的那个节点
	// 倒数第n个，对应的索引是count - n , 想要删除第n个， 要停留在前一位count-n-1，然后更改Next, 删除倒数第n个节点
	// 但是result会往前走一位， 下面的i只能小于count-n-1， 才能保证result 往下走一步，刚好停在count-n-1
	// 结论： 提前2步就停
	for i := 0; i < count-n-1; i++ {
		result = result.Next
	}
	// 跳过中间要删除的那个
	result.Next = result.Next.Next

	return head
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
