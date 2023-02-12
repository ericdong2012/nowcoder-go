package main

import "fmt"

/*
NC115 栈和排序
https://www.nowcoder.com/practice/95cb356556cf430f912e7bdf1bc2ec8f?tpId=117&tqId=37839&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=590&title=


描述
给你一个 1 到 n 的排列和一个栈，并按照排列顺序入栈
你要在不打乱入栈顺序的情况下，仅利用入栈和出栈两种操作，输出字典序最大的出栈序列
排列：指 1 到 n 每个数字出现且仅出现一次

进阶：空间复杂度 O(n) ，时间复杂度 O(n)


示例1
输入：
[2,1,5,3,4]
返回值：
[5,4,3,1,2]

说明：
操作       栈     结果
2 入栈；[2]       []
1 入栈；[2\1]     []
5 入栈；[2\1\5]   []
5 出栈；[2\1]     [5]
3 入栈；[2\1\3]   [5]
4 入栈；[2\1\3\4] [5]
4 出栈；[2\1\3]   [5,4]
3 出栈；[2\1]     [5,4,3]
1 出栈；[2]       [5,4,3,1]
2 出栈；[]        [5,4,3,1,2]


示例2
输入：
[1,2,3,4,5]
返回值：
[5,4,3,2,1]

*/

func solve1(a []int) []int {
	if len(a) == 0 {
		return []int{}
	}

	max := make([]int, len(a))
	// 找到最大值
	curMax := a[len(a)-1]
	max[len(a)-1] = curMax
	// 构造一个数组  [5,5,5,5,4]
	for i := len(a) - 2; i >= 0; i-- {
		if a[i] > curMax {
			curMax = a[i]
		}
		max[i] = curMax
	}

	var stack []int
	var result []int
	// 顺序入站，如果当前入站的元素比将要入站的剩余元素都要大，那么这个元素出站
	for i := 0; i < len(a); i++ {
		// 此处容易出错， 可能连续好几个数字都是一样的， 等于最大的数字
		for len(stack) != 0 && stack[len(stack)-1] >= max[i] {
			result = append(result, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, a[i])
	}
	// 此时resutl 中有[5,4], stack: [3,1,2]
	// 将 stack 中剩余的逆序输出
	for i := len(stack) - 1; i >= 0; i-- {
		result = append(result, stack[i])
	}
	return result
}

func main() {
	//ints := solve1([]int{2, 1, 5, 3, 4})
	ints := solve1([]int{1, 2, 3, 4, 5})
	fmt.Println(ints)
}
