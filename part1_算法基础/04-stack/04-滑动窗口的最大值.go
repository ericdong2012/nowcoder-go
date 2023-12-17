package main

import "fmt"

/*
NC82 滑动窗口的最大值
https://www.nowcoder.com/practice/1624bc35a45c42c0bc17d17fa0cba788?tpId=117&tqId=37784&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=581,582,586&title=


描述
给定一个长度为 n 的数组 nums 和滑动窗口的大小 size ，找出所有滑动窗口里数值的最大值。

例如，如果输入数组{2,3,4,2,6,2,5,1}及滑动窗口的大小3，那么一共存在6个滑动窗口，他们的最大值分别为{4,4,6,6,6,5}； 针对数组{2,3,4,2,6,2,5,1}的滑动窗口有以下6个：
{[2,3,4],2,6,2,5,1}， {2,[3,4,2],6,2,5,1}， {2,3,[4,2,6],2,5,1}， {2,3,4,[2,6,2],5,1}， {2,3,4,2,[6,2,5],1}， {2,3,4,2,6,[2,5,1]}。

示例1
输入：
[2,3,4,2,6,2,5,1],3
返回值：
[4,4,6,6,6,5]

示例2
输入：
[9,10,9,-7,-3,8,2,-6],5
返回值：
[10,10,9,8]

示例3
输入：
[1,2,3,4],3
返回值：
[3,4]

*/

/*
双向队列, 单调队列
*/

func maxInWindows(num []int, size int) []int {
	// 单调栈
	//if size > len(num) || size==0 {
	//	return nil
	//}
	//n := len(num)
	//// 存的是索引
	//deque := make([]int, 0)
	//// 结果
	//res := make([]int, 0)
	//for i := 0; i < n; i++ {
	//	// 不断移除队列中比当前元素小的元素, 最终将当前窗口的最大值的索引存起来
	//	for len(deque) > 0 && num[deque[len(deque)-1]] < num[i] {
	//		deque = deque[:len(deque)-1]
	//	}
	//	deque = append(deque, i)
	//	// 判断头部节点是否还在滑动窗口内 ??? 比如: [1, -1], 1
	//	//if deque[0]+size <= i {
	//	//	deque = deque[1:]
	//	//}
	//	// i 到了 滑动窗口外面的下一个元素，同时队列中保存的索引是 当前索引 - 长度，也就是滑动窗口中第一个元素就是最大值
	//	if  i >= size &&  deque[0] == i-size{
	//		deque = deque[1:]
	//	}
	//	// 长度达到 size
	//	if i+1 >= size {
	//		res = append(res, num[deque[0]])
	//	}
	//}
	//return res

	// 选择排序的变换, 只是要固定长度
	if size > len(num) || size == 0 {
		return nil
	}
	var res []int
	// 1. 要固定长度， 最后的等于号不要忘
	// 2. 内部起点，终点要想清楚
	for j := 0; j <= len(num)-size; j++ {
		maxValue := num[j]
		// for i := j + 1; i < len(num)-size+1; i++ {
		// 想清楚起点，终点, 一定是加上
		//for i := j +1 ; i < j + size ; i++ {
		for i := j ; i < j + size ; i++ {
			if maxValue < num[i] {
				maxValue = num[i]
			}
		}

		res = append(res, maxValue)
	}

	return res
}

func main() {
	L := maxInWindows([]int{2, 3, 4, 2, 6, 2, 5, 1}, 3)
	fmt.Println(L)
	//temp := []int{2, 3, 4, 2, 6, 2, 5, 1}
	//fmt.Println(temp[:len(temp)-1])
}
