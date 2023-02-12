package main

import "fmt"

/*
NC156 数组中只出现一次的数（其它数出现k次）
https://www.nowcoder.com/practice/5d3d74c3bf7f4e368e03096bb8857871?tpId=117&tqId=37866&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=5074&title=


描述
给定一个长度为 n 的整型数组 arr 和一个整数 k(k>1) 。
已知 arr 中只有 1 个数出现一次，其他的数都出现 k 次。
请返回只出现了 1 次的数。


进阶：时间复杂度 O(32n)，空间复杂度 O(1)


示例1
输入：
[5,4,1,1,5,1,5],3
返回值：
4

示例2
输入：
[2,2,1],2

返回值：
1

*/
func foundOnceNumber(arr []int, k int) int {
	// write code here
	binSum := make([]int, 32)
	for i := 0; i < 32; i++ {
		sum := 0
		for _, num := range arr {
			sum += (num >> i & 1)
		}
		binSum[i] = sum
	}
	res := 0
	for i := 0; i < 32; i++ {
		if binSum[i]%k != 0 {
			res += 1 << i
		}
	}
	return res
}

func main() {
	number := foundOnceNumber([]int{5, 4, 1, 1, 5, 1, 5}, 3)
	fmt.Println(number)
}
