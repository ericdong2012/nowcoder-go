package main

import (
	"fmt"
)

/*
NC27 集合的所有子集(一)
https://www.nowcoder.com/practice/c333d551eb6243e0b4d92e37a06fbfc9?tpId=117&tqId=37731&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=590,5058,578&title=


描述
现在有一个没有重复元素的整数集合S，求S的所有子集
注意：
你给出的子集中的元素必须按升序排列
给出的解集中不能出现重复的元素

数据范围：1 \le n \le 51≤n≤5，集合中的任意元素满足 |val| \le 10∣val∣≤10
要求：空间复杂度 O(n!)，时间复杂度 O(n!)

示例1
输入：
[1,2,3]
返回值：
[[],[1],[2],[3],[1,2],[1,3],[2,3],[1,2,3]]

示例2
输入：
[]
返回值：
[]

*/

func subsets(A []int) [][]int {
	// write code here
	/*
		每个元素只有两种状态，选择或者不选择，选择当前第i个元素，计算出其的所有组合结果，或者是不选择该元素，返回其组合的结果值。还是首先把集合内的元素进行排序，添加一个空子集，随后进行判断

		class Solution:
		    def subsets(self , A ):
		        # write code here
		        res=[[]]
		        for item in A:
		            res+=[i+[item] for i in res]
		        return res
	*/

	//sort.Ints(A)
	res := [][]int{}
	res = append(res, []int{})
	for i := 0; i < len(A); i++ {
		n := len(res)
		for j := 0; j < n; j++ {
			var subSet []int
			subSet = append(subSet, res[j]...)
			subSet = append(subSet, A[i])
			res = append(res, subSet)
		}
	}
	return res
}

func main() {
	i := subsets([]int{1, 2, 3})
	fmt.Println(i)
}
