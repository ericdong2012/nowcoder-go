package main

import "fmt"

/*
NC208 每日温度
https://www.nowcoder.com/practice/1f54e163e6944cc7b8759cc09e9c78d8?tpId=117&tqId=39397&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=581,582,586&title=


描述
根据往后 n 天的天气预报，计算每一天需要等待几天才会出现一次更高的气温，如果往后都没有更高的气温，则用 0 补位。

例如往后三天的气温是 [1,2,3] ， 则输出 [1,1,0]

数据范围： 1 \le n \le 10^5 \1≤n≤10
5
   ，每天的温度会满足  0 \le dailyTemperatures[i] \le 1000 \0≤dailyTemperatures[i]≤1000
示例1
输入：
[1,2,3]
复制
返回值：
[1,1,0]
复制
示例2
输入：
[2,4,5,9,10,0,9]
复制
返回值：
[1,1,1,1,0,1,0]
复制
示例3
输入：
[3,1,4]
复制
返回值：
[2,1,0]

*/

// 单调栈
func temperatures(dailyTemperatures []int) []int {
	// write code here
	stack := []int{}
	res := make([]int, len(dailyTemperatures))
	for i := 0; i < len(dailyTemperatures); i++ {
		for len(stack) != 0 && dailyTemperatures[stack[len(stack)-1]] < dailyTemperatures[i] {
			index := stack[len(stack)-1]
			stack = stack[:len(stack) - 1]
			res[index] = i - index
		}
		stack = append(stack, i)
	}
	return res
}

func main() {
	ints := temperatures([]int{1, 2, 3})
	fmt.Println(ints)
}
