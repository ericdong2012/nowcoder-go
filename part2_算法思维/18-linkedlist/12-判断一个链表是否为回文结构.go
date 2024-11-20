package main

import "fmt"

/*
BM13 判断一个链表是否为回文结构
https://www.nowcoder.com/practice/3fed228444e740c8be66232ce8b87c2f?tpId=295&tqId=1008769&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295

描述
给定一个链表，请判断该链表是否为回文结构。
回文是指该字符串正序逆序完全一致。

示例1
输入：
{1}
返回值：
true

示例2
输入：
{2,1}
返回值：
false

说明：
2->1


示例3
输入：
{1,2,2,1}
返回值：
true

说明：
1->2->2->1
*/

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}


// 直接将链表中的值 用数组存起来，然后前后比较，妙啊
func isPail( head *ListNode ) bool {
	// write code here
	if head == nil {
		return false
	}

	arr := make([]int, 0)
	for head != nil  {
		arr = append(arr, head.Val)
		head = head.Next
	}

	i := 0
	j := len(arr) - 1
	for i < j {
		if arr[i] != arr[j] {
			return false
		}
		i++
		j--
	}
	return true
}

//func isPail( head *ListNode ) bool {
//	// write code here
//	temp:=&ListNode{Next:head}
//	slow,fast:=temp,temp
//	for fast!=nil && fast.Next!=nil{
//		slow=slow.Next
//		fast=fast.Next.Next
//	}
//	mid:=slow
//	var pre,tmp *ListNode
//	cur:=mid
//	for cur!=nil{
//		tmp=cur.Next
//		cur.Next=pre
//		pre=cur
//		cur=tmp
//	}
//	tail:=pre
//	for head!=nil && tail!=nil{
//		if head.Val != tail.Val{
//			return false
//		}
//		head=head.Next
//		tail=tail.Next
//	}
//	return true
//}

func main13() {
	// test
	node := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		},
	}

	// 考虑边界情况
	//result := removeNthFromEnd(&node, 2)
	//result := removeNthFromEnd(&node, 5)
	result := isPail(&node)
	fmt.Printf("%+v\n", result)

}
