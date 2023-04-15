package main

import "fmt"

/*
BM10 两个链表的第一个公共结点
https://www.nowcoder.com/practice/6ab1d9a29e88450685099d45c9e31e46?tpId=295&tqId=23257&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295


描述
输入两个无环的单向链表，找出它们的第一个公共结点，如果没有公共节点则返回空。（注意因为传入数据是链表，所以错误测试数据的提示是用其他方式显示的，保证传入数据是正确的）

例如，输入{1,2,3},{4,5},{6,7}时，两个无环的单向链表的结构如下图所示：
可以看到它们的第一个公共结点的结点值为6，所以返回结点值为6的结点。

输入描述：
输入分为是3段，第一段是第一个链表的非公共部分，第二段是第二个链表的非公共部分，第三段是第一个链表和第二个链表的公共部分。 后台会将这3个参数组装为两个链表，并将这两个链表对应的头节点传入到函数FindFirstCommonNode里面，用户得到的输入只有pHead1和pHead2。

返回值描述：
返回传入的pHead1和pHead2的第一个公共结点，后台会打印以该节点为头节点的链表。

示例1
输入：
{1,2,3},{4,5},{6,7}
返回值：
{6,7}

说明：
第一个参数{1,2,3}代表是第一个链表非公共部分，第二个参数{4,5}代表是第二个链表非公共部分，最后的{6,7}表示的是2个链表的公共部分
这3个参数最后在后台会组装成为2个两个无环的单链表，且是有公共节点的

示例2
输入：
{1},{2,3},{}
返回值：
{}

说明：
2个链表没有公共节点 ,返回null，后台打印{}

*/
//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func FindFirstCommonNode(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	// 本地跑的代码和线上结果不一样    这是比较坑的地方

	//找到公共点退出 或 走到链表末尾都nil退出
	/*
		首先，使用两个指针分别指向两个链表
		如果其中一个到达了终点，则把另一个链表头节点赋值给对应指针，
		直到最后两个指针指向同一个位置
	*/
	p1, p2 := pHead1, pHead2
	for p1 != p2 {
		if p1 != nil {
			p1 = p1.Next
		} else {
			p1 = pHead2
		}

		if p2 != nil {
			p2 = p2.Next
		} else {
			p2 = pHead1
		}
	}
	return p1

	// 方式二：map查找优化， 有bug
	/*
		思路核心：
		1、遍历第二个链表去匹配第一个
		2、用map优化查找速度
	*/

	//h := map[int]*ListNode{}
	//cur := pHead1
	//for cur != nil {
	//	h[cur.Val] = cur
	//	cur = cur.Next
	//}
	//cur2 := pHead2
	////fmt.Printf("1111: %+v", h)
	//for cur2 != nil {
	//	//fmt.Printf("2222: %+v", cur2)
	//	if h[cur2.Val] != nil {
	//		return cur2
	//	}
	//	cur2 = cur2.Next
	//}
	//return nil
}

func main10() {
	// test
	node := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val:  7,
						Next: nil,
					},
				},
			},
		},
	}

	node1 := ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val:  7,
					Next: nil,
				},
			},
		},
	}

	// 考虑边界情况
	//result := removeNthFromEnd(&node, 2)
	//result := removeNthFromEnd(&node, 5)
	result := FindFirstCommonNode(&node, &node1)
	fmt.Printf("%+v\n", result)
	//fmt.Printf("%+v\n", result.Next)
	//fmt.Printf("%+v\n", result.Next.Next)
	//fmt.Printf("%+v\n", result.Next.Next.Next)
	//fmt.Printf("%+v\n", result.Next.Next.Next.Next)

}
