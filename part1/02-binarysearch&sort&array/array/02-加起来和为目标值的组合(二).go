package main

import "sort"

/*
NC46 加起来和为目标值的组合(二)
https://www.nowcoder.com/practice/75e6cd5b85ab41c6a7c43359a74e869a?tpId=117&tqId=37742&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=590,5058,578&title=


描述
给出一组候选数 c 和一个目标数 t ，找出候选数中起来和等于 t 的所有组合。
c 中的每个数字在一个组合中只能使用一次。

注意：
1. 题目中所有的数字（包括目标数 t ）都是正整数
2. 组合中的数字 ( a_1, a_2, … , a_k)
3. 结果中不能包含重复的组合
4. 组合之间的排序按照索引从小到大依次比较，小的排在前面，如果索引相同的情况下数值相同，则比较下一个索引。

要求：空间复杂度 O(n!） ， 时间复杂度 O(n!)

示例1
输入：
[100,10,20,70,60,10,50],80
返回值：
[[10,10,60],[10,20,50],[10,70],[20,60]]

说明：
给定的候选数集是[100,10,20,70,60,10,50]，目标数是80


示例2
输入：
[2],1
返回值：
[]

*/

func combinationSum2(num []int, target int) [][]int {
	// write code here
	sort.Ints(num)
	n := len(num)
	result := [][]int{}
	visited := make([]bool, n)
	var tb func(temp []int, start, t int)
	tb = func(temp []int, start, t int) {
		if t == target {
			tempCopy := make([]int, len(temp))
			copy(tempCopy, temp)
			result = append(result, tempCopy)
			return
		}
		if start >= n || t > target {
			return
		}
		for i := start; i < n; i++ {
			if i > 0 && num[i] == num[i-1] && !visited[i-1] {
				continue
			}
			temp = append(temp, num[i])
			visited[i] = true
			tb(temp, i+1, t+num[i])
			// 将数据复原
			visited[i] = false
			temp = temp[:len(temp)-1]
		}
	}
	tb([]int{}, 0, 0)
	return result
}
