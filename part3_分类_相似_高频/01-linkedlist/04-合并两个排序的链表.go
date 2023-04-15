package main

import "fmt"

/*
BM4 合并两个排序的链表
https://www.nowcoder.com/practice/d8b6b4358f774294a89de2a6ac4d9337?tpId=295&tqId=23267&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295


输入两个递增的链表，单个链表的长度为n，合并这两个链表并使新链表中的节点仍然是递增排序的。
数据范围： 0 \le n \le 10000≤n≤1000，-1000 \le 节点值 \le 1000−1000≤节点值≤1000
要求：空间复杂度 O(1)O(1)，时间复杂度 O(n)O(n)

如输入{1,3,5},{2,4,6}时，合并后的链表为{1,2,3,4,5,6}，所以对应的输出为{1,2,3,4,5,6}，转换过程如下图所示：
或输入{-1,2,4},{1,3,4}时，合并后的链表为{-1,1,2,3,4,4}，所以对应的输出为{-1,1,2,3,4,4}，转换过程如下图所示：


示例1
输入：
{1,3,5},{2,4,6}
返回值：
{1,2,3,4,5,6}


示例2
输入：
{},{}
返回值：
{}

示例3
输入：
{-1,2,4},{1,3,4}
返回值：
{-1,1,2,3,4,4}

*/

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

/**
 *
 * @param pHead1 ListNode类
 * @param pHead2 ListNode类
 * @return ListNode类
 */
func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// 哑节点，虚拟头节点
	vhead := &ListNode{Val: -1}
	cur := vhead
	for pHead1 != nil && pHead2 != nil {
		if pHead1.Val <= pHead2.Val {
			cur.Next = pHead1
			pHead1 = pHead1.Next
		} else {
			cur.Next = pHead2
			pHead2 = pHead2.Next
		}
		cur = cur.Next
	}
	if pHead1 != nil {
		cur.Next = pHead1
	} else {
		cur.Next = pHead2
	}
	return vhead.Next
}

func main04() {
	// test
	// 画图
	node1 := ListNode{
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

	node2 := ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 7,
					Next: &ListNode{
						Val:  9,
						Next: nil,
					},
				},
			},
		},
	}

	result := Merge(&node1, &node2)
	fmt.Printf("%+v\n", result)
	fmt.Printf("%+v\n", result.Next)
	fmt.Printf("%+v\n", result.Next.Next)
	fmt.Printf("%+v\n", result.Next.Next.Next)
	fmt.Printf("%+v\n", result.Next.Next.Next.Next)
	fmt.Printf("%+v\n", result.Next.Next.Next.Next.Next)
	fmt.Printf("%+v\n", result.Next.Next.Next.Next.Next.Next)
	fmt.Printf("%+v\n", result.Next.Next.Next.Next.Next.Next.Next)
	fmt.Printf("%+v\n", result.Next.Next.Next.Next.Next.Next.Next.Next)
	fmt.Printf("%+v\n", result.Next.Next.Next.Next.Next.Next.Next.Next.Next)

}
