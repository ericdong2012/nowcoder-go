package _3_hash

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

1、特判，如果数组为 null 或者数组长度小于 3，返回 []。
2、对数组进行排序。
3、遍历排序后数组:
    1、若 num[i]>0：因为已经排序好，所以后面不可能有三个数加和等于 0，直接返回结果。
    2、对于重复元素(num[i] == num[i- 1] )：跳过
    3、令左指针 L=i+1，右指针R=n−1，
        1、当 nums[i]+nums[L]+nums[R]==0，执行循环，写入结果, 同时判断左界和右界是否和下一位置重复，去除重复解, 并将 L,R 移到下一位置
        2、若和大于 0，说明 nums[R] 太大，R 左移
        3、若和小于 0，说明 nums[L] 太小，L 右移

*/

//func threeSum(num []int) [][]int {
//	//1.三元组 非降序排列 -> 将num升序排列
//	//2.数求和问题 -> 双指针
//	if len(num) < 3 || num == nil {
//		return [][]int{}
//	}
//	//排序
//	sort.Ints(num)
//	n := len(num)
//	res := make([][]int, 0)
//	for i := 0; i < n-2; i++ {
//		if num[i] > 0 {
//			break //排好序后的数组，如果当前数字大于0, 则三数之和一定大于0， 后面两个数字一定比当前大
//		}
//		if i > 0 && num[i] == num[i-1] {
//			continue //去重
//		}
//		//双指针求和
//		low, high := i+1, n-1 // 一开始就是最大值和最小值
//		for low < high {
//			sum := num[i] + num[low] + num[high]
//			if sum == 0 {
//				//写入结果
//				res = append(res, []int{num[i], num[low], num[high]})
//				// 头部去重（如果后面一个数跟当前的数字相等，则代表有重复的结果生成，跳过
//				for low < high && num[low] == num[low+1] {
//					// continue
//					low++
//				}
//				low++
//				high--
//			} else if sum > 0 {
//				//太大前移
//				high--
//			} else {
//				//太小后移
//				low++
//			}
//		}
//	}
//	return res
//}

//	 func threeSum(nums []int) [][]int {
//    // 哈希表用于存储每个数字出现的次数
//    numCount := make(map[int]int)
//    for _, num := range nums {
//        numCount[num]++
//    }
//
//    var result [][]int
//    // 遍历数组中的每一个数字
//    for i, num1 := range nums {
//        // 如果num1已经被使用过，则跳过
//        if numCount[num1] == 0 {
//            continue
//        }
//        // 将num1从哈希表中删除
//        numCount[num1]--
//
//        // 在剩余的数字中寻找两个数与num1相加等于0
//        for j := i + 1; j < len(nums); j++ {
//            num2 := nums[j]
//            // 如果num2已经被使用过，则跳过
//            if numCount[num2] == 0 {
//                continue
//            }
//            // 将num2从哈希表中删除
//            numCount[num2]--
//
//            // 计算需要的第三个数
//            num3 := -num1 - num2
//            // 检查num3是否在哈希表中且未被使用过
//            if count, ok := numCount[num3]; ok && count > 0 {
//                result = append(result, []int{num1, num2, num3})
//            }
//
//            // 将num2重新添加到哈希表中
//            numCount[num2]++
//        }
//
//        // 将num1重新添加到哈希表中
//        numCount[num1]++
//    }
//
//    return result
//}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	resultMap := make(map[[3]int]struct{})
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			k := sort.SearchInts(nums[j+1:], -nums[i]-nums[j])
			if k < len(nums[j+1:]) && nums[i]+nums[j]+nums[j+k+1] == 0 {
				resultMap[[3]int{nums[i], nums[j], nums[j+k+1]}] = struct{}{}
			}
		}
	}

	result := [][]int{}
	for k := range resultMap {
		result = append(result, append([]int{}, k[0], k[1], k[2]))
	}
	return result
}

func main() {
	sum := threeSum([]int{-2, 0, 1, 2, 1, 2})
	fmt.Println(sum)
}
