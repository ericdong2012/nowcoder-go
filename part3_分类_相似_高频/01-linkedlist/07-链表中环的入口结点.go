package main

import "fmt"

/*
BM7 链表中环的入口结点
https://www.nowcoder.com/practice/253d2c59ec3e4bc68da16833f79a38e4?tpId=295&tqId=23449&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295

描述
给一个长度为n链表，若其中包含环，请找出该链表的环的入口结点，否则，返回null。

数据范围： n\le10000n≤10000，1<=结点值<=100001<=结点值<=10000
要求：空间复杂度 O(1)O(1)，时间复杂度 O(n)O(n)

例如，输入{1,2},{3,4,5}时，对应的环形链表如下图所示：
可以看到环的入口结点的结点值为3，所以返回结点值为3的结点。
输入描述：
输入分为2段，第一段是入环前的链表部分，第二段是链表环的部分，后台会根据第二段是否为空将这两段组装成一个无环或者有环单链表

返回值描述：
返回链表的环的入口结点即可，我们后台程序会打印这个结点对应的结点值；若没有，则返回对应编程语言的空结点即可。


示例1
输入：
{1,2},{3,4,5}
返回值：
3

说明：
返回环形链表入口结点，我们后台程序会打印该环形链表入口结点对应的结点值，即3


示例2
输入：
{1},{}

返回值：
"null"

说明：
没有环，返回对应编程语言的空结点，后台程序会打印"null"


示例3
输入：
{},{2}
返回值：
2

说明：
环的部分只有一个结点，所以返回该环形链表入口结点，后台程序打印该结点对应的结点值，即2
*/

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func EntryNodeOfLoop(pHead *ListNode) *ListNode {

	/*
		快慢指针
		map / hash

	*/
	// 方式一
	//if pHead == nil || pHead.Next == nil {
	//	return nil
	//}
	//// 快慢指针
	//fast := pHead
	//last := pHead
	//
	//for fast != nil && fast.Next != nil {
	//	fast = fast.Next.Next
	//	last = last.Next
	//	if fast == last {
	//		break
	//	}
	//}
	//// 第一次相遇
	//if last == nil || fast == nil {
	//	// 无环
	//	return nil
	//}
	//last = pHead
	//for last != fast {
	//	last = last.Next
	//	fast = fast.Next
	//}
	//return last

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

	// 方式三： hash表
	//hit := make(map[*ListNode]bool)
	//for {
	//	if pHead == nil {
	//		return nil
	//	}
	//	if _, ok := hit[pHead]; ok {
	//		return pHead
	//	} else {
	//		hit[pHead] = true
	//		pHead = pHead.Next
	//	}
	//}
}

func main07() {
	//// test
	//// 画图
	node1 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}

	//node2 := ListNode{
	//	Val: 3,
	//	Next: &ListNode{
	//		Val: 4,
	//		Next: &ListNode{
	//			Val:  5,
	//			Next: nil,
	//		},
	//	},
	//}

	result := EntryNodeOfLoop(&node1)
	fmt.Printf("%+v\n", result)

}
