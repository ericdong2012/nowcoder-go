package main

import "fmt"

/*
BM11 链表相加(二)
https://www.nowcoder.com/practice/c56f6c70fb3f4849bc56e33ff2a50b6b?tpId=295&tqId=1008772&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295


描述
假设链表中每一个节点的值都在 0 - 9 之间，那么链表整体就可以代表一个整数。
给定两个这种链表，请生成代表两个整数相加值的结果链表。

例如：链表 1 为 9->3->7，链表 2 为 6->3，最后生成新的结果链表为 1->0->0->0。

示例1
输入：
[9,3,7],[6,3]
返回值：
{1,0,0,0}


示例2
输入：
[0],[6,3]
返回值：
{6,3}

*/

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func reverseLink(head1 *ListNode) *ListNode {
	pre := (*ListNode)(nil)
	current := head1
	for current != nil {
		next := current.Next
		current.Next = pre
		pre = current
		current = next
	}
	return pre
}

func addInList(head1 *ListNode, head2 *ListNode) *ListNode {
	// write code here
	/*
		因为两个链表长度可能不一样，从高位到低位运算无法实现，所以首先想到要反转两个单链表，这题因为没有要求要维持原链表结构不变，直接在原链表上翻转即可
		翻转链表之后，依次对两个链表的对齐节点进行运算即可，因为两个链表的长度可能不一样，这里有个小技巧，较短的链表高位都当作0即可，计算可能产生进位，用一个临时变量保存是否有进位
	*/
	t1 := reverseLink(head1)
	t2 := reverseLink(head2)

	virtualNode := &ListNode{Next: nil}
	p := virtualNode

	carry := 0
	for t1 != nil || t2 != nil {
		// 关键点
		v := carry
		if t1 != nil {
			v += t1.Val
			t1 = t1.Next
		}
		if t2 != nil {
			v += t2.Val
			t2 = t2.Next
		}
		if v >= 10 {
			v -= 10
			carry = 1
		} else {
			carry = 0
		}
		newNode := &ListNode{Val: v}
		p.Next = newNode
		p = p.Next

	}
	// 较长链的最后一个节点, 刚好有进位, 需要在整个链的前面加上一个1
	// 1, 2, 4, 9
	// 2, 3, 6
	// 3, 5, 0, 0, 1
	if carry == 1 {
		p.Next = &ListNode{Val: 1}
	}
	// 将链表反转回来
	result := reverseLink(virtualNode.Next)
	return result

}

func main11() {
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
	result := addInList(&node, &node1)
	fmt.Printf("%+v\n", result)
	fmt.Printf("%+v\n", result.Next)
	fmt.Printf("%+v\n", result.Next.Next)
	fmt.Printf("%+v\n", result.Next.Next.Next)
	fmt.Printf("%+v\n", result.Next.Next.Next.Next)

}
