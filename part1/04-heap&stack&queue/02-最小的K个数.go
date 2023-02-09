package main

import (
	"fmt"
	"sort"
)

/*
NC119 最小的K个数
https://www.nowcoder.com/practice/6a296eb82cf844ca8539b57c23e6e9bf?tpId=117&tqId=37765&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=581,582,586&title=

描述
给定一个长度为 n 的可能有重复值的数组，找出其中不去重的最小的 k 个数。例如数组元素是4,5,1,6,2,7,3,8这8个数字，则最小的4个数字是1,2,3,4(任意顺序皆可)。
数据范围：0\le k,n \le 100000≤k,n≤10000，数组中每个数的大小0 \le val \le 10000≤val≤1000
要求：空间复杂度 O(n)O(n) ，时间复杂度 O(nlogn)O(nlogn)

示例1
输入：
[4,5,1,6,2,7,3,8],4
复制
返回值：
[1,2,3,4]
复制
说明：
返回最小的4个数即可，返回[1,3,2,4]也可以
示例2
输入：
[1],0
复制
返回值：
[]
复制
示例3
输入：
[0,1,2,1,2],3
复制
返回值：
[0,1,1]

*/

func GetLeastNumbers_Solution(input []int, k int) []int {
	// write code here
	if len(input) < k && k == 0 {
		return []int{}
	}
	sort.Ints(input)
	if len(input) == k {
		return input
	}
	//result := make([]int, 0)
	//for i := 0; i < k; i++ {
	//	result = append(result, input[i])
	//}
	return input[:k]
}

func main() {
	solution := GetLeastNumbers_Solution([]int{4, 5, 1, 6, 2, 7, 3, 8}, 4)
	fmt.Println(solution)
}
