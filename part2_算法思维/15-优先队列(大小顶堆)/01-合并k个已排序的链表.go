package _5_优先队列_大小顶堆_

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

func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	}
	interval := 1
	for interval < n {
		// 如果是2个 0,1 merge即可
		// 如果是3个，0,1 merge完成后，  跳出然后 0，2 merge  所以是interval * 2
		// 如果是4个，0,1 merge完成后   2，3 merge, 跳出 0,2 merge
		for i := 0; i < n-interval; i += interval * 2 {
			lists[i] = merge2Lists(lists[i], lists[i+interval])
		}
		interval *= 2
	}
	return lists[0]
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
