package main

import "fmt"

/*
NC115 栈和排序
https://www.nowcoder.com/practice/95cb356556cf430f912e7bdf1bc2ec8f?tpId=117&tqId=37839&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=581,582,586&title=


描述
给你一个 1 到 n 的排列和一个栈，并按照排列顺序入栈
你要在不打乱入栈顺序的情况下，仅利用入栈和出栈两种操作，输出字典序最大的出栈序列
排列：指 1 到 n 每个数字出现且仅出现一次

示例1
输入：
[2,1,5,3,4]
返回值：
[5,4,3,1,2]

说明：
操作    栈        结果
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

/*

因为给定的是1-n的数据，所有最大值必定是n，我们可以设置一个dp数组表示这个最大值是否被访问过，
如果这个最大值已经被用过出栈了，我们需要递减以更新最大值，如果这个最大值没有被访问过可以直接比较栈顶元素和它的大小，当栈顶较大时，出栈即可。


读入数据：第i个数出栈满足字典序最大,一定是i+1到n中最大的一个数
用一个数组maxs存i-n之间最大的数
按照读入顺序入栈,如果当前入栈第i个数字比将要入栈的剩余元素都要大 那么这个元素出栈

*/

func solve1(a []int) []int {
	if len(a) == 0 {
		return []int{}
	}

	max := make([]int, len(a))
	// 找到最大值
	curMax := a[len(a)-1]
	max[len(a)-1] = curMax
	for i := len(a) - 2; i >= 0; i-- {
		if a[i] > curMax {
			curMax = a[i]
		}
		max[i] = curMax
	}

	var stack []int
	var rst []int

	// 顺序入站，如果当前入站的元素比将要入站的剩余元素都要大，那么这个元素出站
	for i := 0; i < len(a); i++ {
		for len(stack) != 0 && stack[len(stack)-1] >= max[i] {
			rst = append(rst, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, a[i])
	}

	// 将 stack 中剩余的逆序输出
	for i := len(stack) - 1; i >= 0; i-- {
		rst = append(rst, stack[i])
	}
	return rst
}
func main() {
	//ints := solve1([]int{2, 1, 5, 3, 4})
	ints := solve1([]int{1, 2, 3, 4, 5})
	fmt.Println(ints)
}
