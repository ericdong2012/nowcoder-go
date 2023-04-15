package main

import "fmt"

/*
NC82 滑动窗口的最大值
https://www.nowcoder.com/practice/1624bc35a45c42c0bc17d17fa0cba788?tpId=117&tqId=37784&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=581,582,586&title=


描述
给定一个长度为 n 的数组 nums 和滑动窗口的大小 size ，找出所有滑动窗口里数值的最大值。

例如，如果输入数组{2,3,4,2,6,2,5,1}及滑动窗口的大小3，那么一共存在6个滑动窗口，他们的最大值分别为{4,4,6,6,6,5}； 针对数组{2,3,4,2,6,2,5,1}的滑动窗口有以下6个： {[2,3,4],2,6,2,5,1}， {2,[3,4,2],6,2,5,1}， {2,3,[4,2,6],2,5,1}， {2,3,4,[2,6,2],5,1}， {2,3,4,2,[6,2,5],1}， {2,3,4,2,6,[2,5,1]}。

数据范围： 1 \le size \le n \le 100001≤size≤n≤10000，数组中每个元素的值满足 |val| \le 10000∣val∣≤10000
要求：空间复杂度 O(n)O(n)，时间复杂度 O(n)O(n)


示例1
输入：
[2,3,4,2,6,2,5,1],3
复制
返回值：
[4,4,6,6,6,5]
复制
示例2
输入：
[9,10,9,-7,-3,8,2,-6],5
复制
返回值：
[10,10,9,8]
复制
示例3
输入：
[1,2,3,4],3
复制
返回值：
[3,4]

*/

/*
双向队列, 单调队列
*/

func maxInWindows(num []int, size int) []int {
	// write code here
	//l := len(num)
	//if l == 0 || size == 0 || size > l {
	//	return nil
	//}
	//
	////  定义返回的结果集
	//res := make([]int, 0)
	////  单调队列
	////  定义从大到小排序
	//list := make([]int, 0)
	//
	////  窗口右移，判断单调队列第一个元素是否等于被移除的元素值
	////  等于被移除的元素值，则单调队列第一个元素也移除
	//pop := func(val int) {
	//	if list[0] == val {
	//		list = list[1:]
	//	}
	//}
	//
	////  向单调队列压入一个数据
	////  单调队列:  大于等于左边元素 或者 小于等于右边元素，当前是从大到小的队列
	//push := func(val int) {
	//	for len(list) != 0 && list[len(list)-1] < val {
	//		// 删除最后一个元素
	//		list = list[:len(list)-1]
	//	}
	//	// 再压入数据
	//	list = append(list, val)
	//}
	//
	////  前size个元素进入单调队列
	//for i := 0; i < size; i++ {
	//	push(num[i])
	//}
	////  前size个元素为第一个窗口，res保存第一个窗口最大值list[0]
	//res = append(res, list[0])
	//
	////  处理后续元素
	//for i := size; i < len(num); i++ {
	//	// 处理窗口左侧被移除时 是否也需要从单调队列移除
	//	pop(num[i-size])
	//	// 处理当前窗口右侧新增的元素
	//	push(num[i])
	//
	//	// 取单调队列最大的元素值
	//	res = append(res, list[0])
	//}
	//
	//return res

	n := len(num)
	// 单调栈，存的是索引
	deque := make([]int, 0)
	// 结果
	res := make([]int, 0)
	for i := 0; i < n; i++ {
		for len(deque) > 0 && num[deque[len(deque)-1]] < num[i] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
		// 判断头部节点是否还在滑动窗口内
		if deque[0]+size <= i {
			deque = deque[1:]
		}
		if i+1 >= size {
			res = append(res, num[deque[0]])
		}
	}
	return res
}

func main() {
	L := maxInWindows([]int{2, 3, 4, 2, 6, 2, 5, 1}, 3)
	fmt.Println(L)
	//temp := []int{2, 3, 4, 2, 6, 2, 5, 1}
	//fmt.Println(temp[:len(temp)-1])
}
