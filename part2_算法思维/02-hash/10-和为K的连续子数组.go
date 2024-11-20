package _3_hash

import "fmt"

/*
NC125 和为K的连续子数组
https://www.nowcoder.com/practice/704c8388a82e42e58b7f5751ec943a11?tpId=117&tqId=37794&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=585&title=


描述
给定一个无序数组 arr, 其中元素可正、可负、可0。给定一个整数k ，求 arr 所有连续子数组中 累加和为k 的 最长连续子数组长度。保证至少存在一个合法的连续子数组。
[1,2,3]的连续子数组有[1,2]，[2,3]，[1,2,3] ，但是[1,3]不是

示例1
输入：
[1,-2,1,1,1],0
返回值：
3

示例2
输入：
[0,1,2,3],3
返回值：
3

*/

/*
1. 连续
2. 和为k
3. 最长

前缀和 + hash

和前缀和中的   和大于等于k的最短子数组  比较， 一个最短，一个最长
和最长不含重复字符的子字符串比较  需要用到前缀和，hash

*/
func maxlenEqualK(arr []int, k int) int {
	// [1,-2,1,1,1], 0  => 3
	// 前缀和 + hash
	if len(arr) == 0 {
		return 0
	}

	sum := 0
	maxLen := 0
	temp := make(map[int]int, len(arr))
	// 此处很重要， 两个索引相减，长度需要加1
	temp[0] = -1
	for i, v := range arr {
		sum += v
		// 符合条件的数据在i, j 之间
		// j是之前元素的索引
		// i是sum加上v的索引
		if j, ok := temp[sum-k]; ok {
			maxLen = max(maxLen, i - j)
		}
		// 不存在, 构建sum为key, 索引为value 的temp hash表
		if _, ok := temp[sum]; !ok {
			temp[sum] = i
		}
	}
	return maxLen
}

func main() {
	//k := maxlenEqualK([]int{1, -2, 1, 1, 1}, 0)
	//k := maxlenEqualK([]int{1, -2, 1, 1, 1, -1, 1}, 2)
	k := maxlenEqualK([]int{1, -2, 1, 1}, 2)
	fmt.Println(k)
}
