package main

import (
	"fmt"
	"sort"
)

/*
NC54 三数之和
https://www.nowcoder.com/practice/345e2ed5f81d4017bbb8cc6055b0b711?tpId=117&tqId=37751&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=590&title=

描述
给出一个有n个元素的数组S，S中是否有元素a,b,c满足a+b+c=0 找出数组S中所有满足条件的三元组。


注意：
三元组（a、b、c）中的元素可以按任意顺序排列。
解集中不能包含重复的三元组。

示例1
输入：
[-10,0,10,20,-10,-40]
返回值：
[[-10,-10,20],[-10,0,10]]

示例2
输入：
[-2,0,1,1,2]
返回值：
[[-2,0,2],[-2,1,1]]

示例3
输入：
[0,0]
返回值：
[]

*/

/*

1、特判，对于数组长度 n，如果数组为 null 或者数组长度小于 3，返回 []。
2、对数组进行排序。
3、遍历排序后数组：
    1、若 num[i]>0：因为已经排序好，所以后面不可能有三个数加和等于 0，直接返回结果。
    2、对于重复元素：跳过，避免出现重复解
    3、令左指针 L=i+1，右指针R=n−1，
        1、当 nums[i]+nums[L]+nums[R]==0，执行循环，写入结果，同时判断左界和右界是否和下一位置重复，去除重复解。并将 L,R 移到下一位置
        2、若和大于 0，说明 nums[R] 太大，R 左移
        3、若和小于 0，说明 nums[L] 太小，L 右移

*/

func threeSum(num []int) [][]int {
	//1.三元组 非降序排列》》将num升序排列
	//2.数求和问题》》双指针
	//if len(num) < 3 || num == nil {
	//	return [][]int{}
	//}
	//res := make([][]int, 0)
	////排序
	//sort.Ints(num)
	//n := len(num)
	//for i := 0; i < n-2; i++ {
	//	if num[i] > 0 {
	//		break //排好序后的数组，如果当前数字大于0, 则三数之和一定大于0
	//	}
	//	if i > 0 && num[i] == num[i-1] {
	//		continue //去重
	//	}
	//	//双指针求和
	//	low, high := i+1, n-1 //一开始就是最大值和最小值
	//	for low < high {
	//		sum := num[i] + num[low] + num[high]
	//		if sum == 0 {
	//			//写入结果
	//			res = append(res, []int{num[i], num[low], num[high]})
	//			// 头部去重（如果后面一个数跟当前的数字相等，则代表有重复的结果生成，跳过
	//			for low < high && num[low] == num[low+1] {
	//				low++
	//			}
	//			low++
	//			high--
	//		} else if sum > 0 {
	//			//太大前移
	//			high--
	//		} else {
	//			//太小后移
	//			low++
	//		}
	//	}
	//}
	//return res

	if len(num) < 3 || num == nil {
		return nil
	}
	res := make([][]int, 0)
	sort.Ints(num)
	n := len(num)
	for i := 0; i < n-1; i++ {
		if num[i] > 0 {
			return res
		}

		if i > 0 && num[i] == num[i-1] {
			continue
		}

		left, right := i+1, n-1
		for left < right {
			sums := num[i] + num[left] + num[right]
			if sums == 0 {
				res = append(res, []int{num[i], num[left], num[right]})

				for left < right && num[left] == num[left+1] {
					left++
				}
				left++
				right--
			} else if sums > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return res
}

func main() {
	sum := threeSum([]int{0, 0, 0})
	fmt.Println(sum)
}
