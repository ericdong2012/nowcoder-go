package main

import "fmt"

/*
NC244 对链表进行插入排序
https://www.nowcoder.com/practice/cc6c61215dfb446f8eccea3663e3d8db?tpId=117&tqId=39589&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=590&title=

描述
对链表进行插入排序，
从链表第一个元素开始可以视为部分已排序，每次操作从链表中移除一个元素，然后原地将移除的元素插入到已排好序的部分。
例如输入{2,4,1}时，对应的返回值为{1,2,4}，转换过程如下图所示。

示例1
输入：
{1,2,3}
返回值：
{1,2,3}

示例2
输入：
{2,4,1}
返回值：
{1,2,4}

*/

type TreeNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	// write code here
	/*
		插入排序思想,当前值 <= val 时继续迭代, 直到遇到较大值,由于链表结构,需要重置 next 节点,若是数组要进行移位操作

		插入排序是一种最简单直观的排序算法，它的工作原理是通过构建有序序列，对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入。
		插入排序和冒泡排序一样，也有一种优化算法，叫做拆半插入

		class Solution:
		    def MySort(self , arr: List[int]) -> List[int]:
		        # write code here
		        #插入排序
		        for i in range(1, len(arr)):
		            pre_index=i-1
		            cur=arr[i]
		            while pre_index>=0 and arr[pre_index]>cur:
		                arr[pre_index + 1] = arr[pre_index]
		                pre_index-=1
		            arr[pre_index + 1] = cur
		        return arr


		func insertionSort(arr []int) {
			var i, j, current int
			for i = 1; i < len(arr); i++ {
				current = arr[i]
				j = i - 1
				// 如果前一个比当前的大，交换，继续往前找，直到前一个比当期小
				for j >= 0 && arr[j] > current {
					arr[j+1] = arr[j]
					j = j - 1
				}
				arr[j+1] = current
			}
		}
	*/

	dummy := &ListNode{Val: -10000}
	for head != nil {
		pre := dummy
		// pre 链上的结果小于head 上的结果，pre前进
		for pre.Next != nil && pre.Next.Val < head.Val {
			pre = pre.Next
		}
		// 保存新链的数据
		newNext := pre.Next
		// 链接
		pre.Next = head
		// 保存原有链的数据
		oldNext := head.Next
		// 链接原有pre的下一个节点
		head.Next = newNext
		// head 移动
		head = oldNext
	}

	return dummy.Next
}

func main08() {
	node := ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
	}

	dummyult := insertionSortList(&node)
	fmt.Printf("%+v\n", dummyult)
	fmt.Printf("%+v\n", dummyult.Next)
	fmt.Printf("%+v\n", dummyult.Next.Next)
}
