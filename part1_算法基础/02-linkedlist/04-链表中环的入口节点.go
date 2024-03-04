package main

/*
https://www.nowcoder.com/practice/253d2c59ec3e4bc68da16833f79a38e4?tpId=295&tqId=23449&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295

*/

func EntryNodeOfLoop(pHead *ListNode) *ListNode{
	// 方式一  快慢指针
	/*
		看题解，看图， 有链表，有环   从相遇节点到入口节点的距离 == 从头节点到入口节点的距离 （画图，数学分析得来）
	*/
	if pHead == nil {
		return pHead
	}
	fast, slow := pHead, pHead
	for fast != nil && fast.Next != nil {
		//1、先用快慢指针判断链表是否有环，如果有环找到快慢指针相遇的节点，注意该节点不一定是环的入口节点
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			// 2. 此时相遇，通过数学分析从相遇节点到入口节点的距离 == 从头节点到入口节点的距离， 所以 让慢指针回到表头，此时快慢执行继续往下走，但快慢指针再相遇就是入口处
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