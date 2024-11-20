package main

import "fmt"

/*
BM5 合并k个已排序的链表
https://www.nowcoder.com/practice/65cfde9e5b9b4cf2b6bafa5f3ef33fa6?tpId=295&tqId=724&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295

描述
合并 k 个升序的链表并将结果作为一个升序的链表返回其头节点。

示例1
输入：
[{1,2,3},{4,5,6,7}]
返回值：
{1,2,3,4,5,6,7}

示例2
输入：
[{1,2},{1,4,5},{6}]
返回值：
{1,1,2,4,5,6}

*/


// part2_算法思维/12-分治/01-合并K个已排序的链表.go
func mergeKLists(lists []*ListNode) *ListNode {
	// write code here
	/*
		方法一： 在合并两个排序链表的基础上，加上分治， 变成合并k个已排序的链表
		方法二： 优先队列
	*/
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}
	// 将最前面的两个先合并
	cur := merge2Lists(lists[0], lists[1])
	// 将后面的和前面的合并
	for i := 2; i < len(lists); i++ {
		cur = merge2Lists(cur, lists[i])
	}

	return cur
}

// 合并2个链表  解法1
//func merge2Lists(l1, l2 *ListNode) *ListNode {
//	// 边界处理
//	if l1 == nil {
//		return l2
//	}
//	if l2 == nil {
//		return l1
//	}
//	// 比较两个链表的头结点，将小的作为新列表的头结点
//	var head *ListNode
//	if l1.Val < l2.Val {
//		head = l1
//		head.Next = merge2Lists(l1.Next, l2)
//	} else {
//		head = l2
//		head.Next = merge2Lists(l1, l2.Next)
//	}
//
//	return head
//}

// 合并2个链表  解法2
func merge2Lists(l1, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	head := dummyHead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head.Next = l1
			l1 = l1.Next
		} else {
			head.Next = l2
			l2 = l2.Next
		}
		head = head.Next
	}
	// 对剩余的做处理（l1, l2中至少有一个为空，剩余的是大数）
	if l1 != nil {
		head.Next = l1
	}
	if l2 != nil {
		head.Next = l2
	}
	return dummyHead.Next
}


/* python  优先队列解法
from queue import PriorityQueue
class Solution:
    def mergeKLists(self , lists: List[ListNode]) -> ListNode:
        #小顶堆
        pq = PriorityQueue()
        #遍历所有链表第一个元素
        for i in range(len(lists)):
            #不为空则加入小顶堆
            if lists[i] != None:
                pq.put((lists[i].val, i))
                lists[i] = lists[i].next
        #加一个表头
        res = ListNode(-1)
        head = res
        #直到小顶堆为空
        while not pq.empty():
            #取出最小的元素
            val, idx = pq.get()
            #连接
            head.next = ListNode(val)
            head = head.next
            if lists[idx] != None:
                #每次取出链表的后一个元素加入小顶堆
                pq.put((lists[idx].val, idx))
                lists[idx] = lists[idx].next
        return res.next
*/

func main05() {
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

	result := mergeKLists([]*ListNode{&node1, &node2})
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
