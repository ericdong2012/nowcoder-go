package main

import "fmt"

/*
NC74 数字在升序数组中出现的次数
https://www.nowcoder.com/practice/70610bf967994b22bb1c26f9ae901fa2?tpId=117&tqId=37772&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=5058&title=

描述
给定一个长度为 n 的非降序数组和一个非负数整数 k ，要求统计 k 在数组中出现的次数

示例1
输入：
[1,2,3,3,3,3,4,5],3
返回值：
4

示例2
输入：
[1,3,4,5],6
返回值：
0

*/

func GetNumberOfK(data []int, k int) int {
	// write code here
	hash := make(map[int]int)
	for i := 0; i < len(data); i++ {
		if data[i] == k {
			hash[k] += 1
		}
	}
	return hash[k]
}

func main() {
	k := GetNumberOfK([]int{1, 2, 3, 3, 3, 3, 4, 5}, 3)
	fmt.Println(k)
}
