package main

/*
BM7 链表中环的入口结点
https://www.nowcoder.com/practice/253d2c59ec3e4bc68da16833f79a38e4?tpId=295&tqId=23449&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295

给一个长度为n链表，若其中包含环，请找出该链表的环的入口结点，否则，返回null。

*/

func EntryNodeOfLoop(pHead *ListNode) *ListNode{
	/*
	    看官方题解

	    1 ---  6  ---  3 - 5 - 2 - 4
                       | - - - - - |

	    1. 链表有环
		判断链表中是否有环，主要思想是利用环没有末尾NULL，然后快慢双指针相遇就代表有环

	    2. 找入口处
	   （参考入口节点示意图，数学分析得来）从相遇节点到入口节点的距离 == 从头节点到入口节点的距离

	*/
	if pHead == nil {
		return pHead
	}
	// 快慢双指针
	fast, slow := pHead, pHead
	for fast != nil && fast.Next != nil {
		// 走一步，走两步
		slow = slow.Next
		fast = fast.Next.Next
		// 相等（相遇），说明有环，但是该位置不一定是入口节点， 此时相遇点是2
		if fast == slow {
			// 通过数学分析  从相遇节点到入口节点的距离 == 从头节点到入口节点的距离
			// 让慢指针回到表头， 此时快慢执行继续往下走，调整为相同步数
			// 再相遇就是入口处， 也就是3这个位置
			slow = pHead
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}

			return fast
		}
	}

	return nil
}