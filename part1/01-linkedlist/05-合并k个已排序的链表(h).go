package main

import "fmt"

/*
BM5 合并k个已排序的链表
https://www.nowcoder.com/practice/65cfde9e5b9b4cf2b6bafa5f3ef33fa6?tpId=295&tqId=724&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295


描述
合并 k 个升序的链表并将结果作为一个升序的链表返回其头节点。

数据范围：节点总数满足 0 \le n \le 10^50≤n≤10
链表个数满足 1 \le k \le 10^5 \1≤k≤10  每个链表的长度满足 1 \le len \le 200 \1≤len≤200  ，每个节点的值满足 |val| <= 1000∣val∣<=1000
要求：   时间复杂度 O(nlogk)O(nlogk)


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

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func mergeKLists(lists []*ListNode) *ListNode {
	// write code here
	/*
		方法一： 双指针，分治，归并(在合并两个排序链表的基础上，加上分治， 变成合并k个已排序的链表)
		方法二： 优先队列
	*/

	//if lists == nil || len(lists) == 0 {
	//	return nil
	//}
	//return merge(lists, 0, len(lists)-1)

	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	cur := merge2Lists(lists[0], lists[1])
	for i := 2; i < len(lists); i++ {
		cur = merge2Lists(cur, lists[i])
	}
	return cur
}

//func merge(lists []*ListNode, start, end int) *ListNode {
//	if start == end {
//		return lists[start]
//	}
//	if start > end {
//		return nil
//	}
//	mid := start + (end-start)/2
//	left := merge(lists, start, mid)
//	right := merge(lists, mid+1, end)
//	return merge2Lists(left, right)
//}

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
	if l1 != nil {
		head.Next = l1
	}
	if l2 != nil {
		head.Next = l2
	}
	return dummyHead.Next
}

//func mergeKLists(lists []*ListNode) *ListNode {
//	// write code here
//	switch n := len(lists); n {
//	case 0:
//		return nil
//	case 1:
//		return lists[0]
//	case 2:
//		return merge(lists[0], lists[1])
//	default:
//		// 分治
//		div := n / 2
//		return merge(mergeKLists(lists[:div]), mergeKLists(lists[div:]))
//	}
//}
//
//func merge(list1, list2 *ListNode) *ListNode {
//	head := &ListNode{}
//	cur := head
//	for list1 != nil && list2 != nil {
//		if list1.Val < list2.Val {
//			cur.Next = list1
//			list1 = list1.Next
//		} else {
//			cur.Next = list2
//			list2 = list2.Next
//		}
//		cur = cur.Next
//	}
//	if list2 != nil {
//		cur.Next = list2
//	}
//	cur.Next = list1
//	return head.Next
//
//}

/* python 版本解法
 # 辅助数组
 def mergeKLists(self , lists ):
        # write code here
        # 将所有链表元素存放到list中，排序后再转换为链表
        tmp = []
        for head in lists:
            while head:
                # 将链表结点存放到tmp中
                tmp.append(head.val)
                head = head.next
        if not tmp:
            return None
        # tmp进行排序
        tmp.sort()
        res = ListNode(-1)
        cur = res
        # 根据tmp生成新的链表
        for i in range(len(tmp)):
            cur.next = ListNode(tmp[i])
            cur = cur.next
        return res.next

	# 优先队列
	  def mergeKLists(self , lists ):
        # write code here
        dummy = ListNode(0)
        p = dummy
        head = []
        for i in range(len(lists)):
            if lists[i] :
                heapq.heappush(head, (lists[i].val, i))
                lists[i] = lists[i].next
        while head:
            val, idx = heapq.heappop(head)
            p.next = ListNode(val)
            p = p.next
            if lists[idx]:
                heapq.heappush(head, (lists[idx].val, idx))
                lists[idx] = lists[idx].next
        return dummy.next


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
