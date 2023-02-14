package main

/*
NC82 滑动窗口的最大值
https://www.nowcoder.com/practice/1624bc35a45c42c0bc17d17fa0cba788?tpId=117&tqId=37784&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=5054&title=

描述
给定一个长度为 n 的数组 nums 和滑动窗口的大小 size ，找出所有滑动窗口里数值的最大值。

例如，如果输入数组{2,3,4,2,6,2,5,1}及滑动窗口的大小3，那么一共存在6个滑动窗口，他们的最大值分别为{4,4,6,6,6,5}；
针对数组{2,3,4,2,6,2,5,1}的滑动窗口有以下6个：
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

func maxInWindows(num []int, size int) []int {
	// write code here
	res := make([]int, 0, len(num)-size+1)
	deque := make([]int, 0)
	for i := 0; i < len(num); i++ {
		for len(deque) > 0 && num[i] > deque[len(deque)-1] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, num[i])
		if i >= size {
			if num[i-size] == deque[0] {
				deque = deque[1:]
			}
		}
		if i >= size-1 {
			res = append(res, deque[0])
		}
	}
	return res
}
