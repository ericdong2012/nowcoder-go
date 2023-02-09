package main

import "fmt"

/*
NC71 旋转数组的最小数字
https://www.nowcoder.com/practice/9f3231a991af4f55b95579b44b7a01ba?tpId=117&tqId=37768&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=5058&title=

描述
有一个长度为 n 的非降序数组，比如[1,2,3,4,5]，将它进行旋转，即把一个数组最开始的若干个元素搬到数组的末尾，变成一个旋转数组，比如变成了[3,4,5,1,2]，或者[4,5,1,2,3]这样的。请问，给定这样一个旋转数组，求数组中的最小值。


示例1
输入：
[3,4,5,1,2]
返回值：
1

示例2
输入：
[3,100,200,3]
返回值：
3

*/

func minNumberInRotateArray(rotateArray []int) int {
	// write code here
	left, right := 0, len(rotateArray)-1
	minV := 0
	for left <= right {
		if rotateArray[left] >= rotateArray[right] && minV <= rotateArray[right] {
			left++
			minV = rotateArray[right]
		} else {
			right--
			minV = rotateArray[left]
		}
	}
	return minV
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	array := minNumberInRotateArray([]int{3, 4, 5, 1, 2})
	fmt.Println(array)
}
