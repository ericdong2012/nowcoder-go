package _2_分治

/*
https://www.nowcoder.com/practice/65cfde9e5b9b4cf2b6bafa5f3ef33fa6?tpId=196&tqId=37081&ru=/exam/oj

合并 k 个升序的链表并将结果作为一个升序的链表返回其头节点。

输入：
[{1,2,3},{4,5,6,7}]
返回值：
{1,2,3,4,5,6,7}


输入：
[{1,2},{1,4,5},{6}]
返回值：
{1,1,2,4,5,6}

*/

type ListNode struct {
	Val  int
	Next *ListNode
}


// part2_算法思维/18-linkedlist/05-合并k个已排序的链表.go
// 和 合并两个有序数组， 合并区间对比、
// 主要思想就是 迭代
func mergeKLists(lists []*ListNode) *ListNode {
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

func merge2Lists(l1, l2 *ListNode) *ListNode {
	// 边界处理
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	// 比较两个链表的头结点，将小的作为新列表的头结点
	var head *ListNode
	if l1.Val < l2.Val {
		head = l1
		head.Next = merge2Lists(l1.Next, l2)
	} else {
		head = l2
		head.Next = merge2Lists(l1, l2.Next)
	}

	return head
}
