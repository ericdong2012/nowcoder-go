package main

/*
https://www.nowcoder.com/practice/b58434e200a648c589ca2063f1faf58c?tpId=295&tqId=654&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295



输入：
{1,2,3,4,5},2,4
返回值：
{1,4,3,2,5}



输入：
{5},1,1
返回值：
{5}

*/

/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

/**
 *
 * @param head ListNode类
 * @param m int整型
 * @param n int整型
 * @return ListNode类
 */
func reverseBetween( head *ListNode ,  m int ,  n int ) *ListNode {
	// write code here
	if head == nil || head.Next == nil {
		return head
	}
	// 哑节点
	dummy := &ListNode{Next: head}
	// 开始和结束前的节点
	pre, end := dummy, dummy
	for i := 0; i < m-1; i++ {
		pre = pre.Next
	}

	start := pre.Next
	for i := 0; i < n; i++ {
		end = end.Next
	}
	// 结束节点的下一个节点
	next := end.Next
	end.Next = nil

	// 将开始前的节点指向反转后链表的头节点
	pre.Next = reverse(start)
	// 将开始节点(最后节点)指向反转后链表的尾节点的下一个节点
	start.Next = next

	// 返回哑节点的下一个节点，即整个链表的开头
	return dummy.Next

}

func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}