package main

import "fmt"

/*
BM8 链表中倒数最后k个结点
https://www.nowcoder.com/practice/886370fe658f41b498d40fb34ae76ff9?tpId=295&tqId=1377477&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295


输入一个长度为 n 的链表，设链表中的元素的值为 ai ，返回该链表中倒数第k个节点。
如果该链表长度小于k，请返回一个长度为 0 的链表。

示例1
输入：
{1,2,3,4,5},2
返回值：
{4,5}

说明：
返回倒数第2个节点4，系统会打印后面所有的节点来比较。


示例2
输入：
{2},8
返回值：
{}

*/

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func FindKthToTail(pHead *ListNode, k int) *ListNode {
	// write code here
	if k <= 0 || pHead == nil {
		return nil
	}

	pre := pHead
	// 因为不是列表，所以需要自己去计算长度
	count := 1
	for pHead.Next != nil {
		pHead = pHead.Next
		count += 1
	}
	// 此时已经拉到count
	if count < k {
		return nil
	}
	//fmt.Println(count)
	// 此处不能正序找，报错
	// 反着找, pre 走到倒数节点的入口处
	for i := count; i > k; i-- {
		pre = pre.Next
	}
	return pre
}

func main08() {
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
	result := FindKthToTail(&node, 5)
	//result := FindKthToTail(&node, 2)
	//result := FindKthToTail(&node, 3)
	//result := FindKthToTail(&ListNode{}, 100)
	fmt.Printf("%+v\n", result)
	fmt.Printf("%+v\n", result.Next)
	fmt.Printf("%+v\n", result.Next.Next)
	fmt.Printf("%+v\n", result.Next.Next.Next)

}
