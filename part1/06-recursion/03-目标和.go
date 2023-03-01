package main

import "fmt"

/*
NC243 目标和
https://www.nowcoder.com/practice/7fc06e2162f048658252fac542fcb1e8?tpId=117&tqId=39588&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=591&title=


描述
给定一个整数数组nums和一个整数target，请你返回该数组能构成多少种不同的表达式等于target。
规则如下：
1.将数组里每个整数前面可以添加"+"或者"-"符号，组成一个表达式，例如[1,2]，可以变成”+1+2","+1-2","-1+2","-1-2"，这四种
2.只能添加"+"与"-"符号，不能添加其他的符号
3.如果构不成等于target的表达式，请返回0
4.保证返回的结果个数在整数范围内


示例1
输入：
[1,1,1,2],3
返回值：
3

说明：
-1 + 1 + 1 + 2 = 3
+1 - 1 + 1 + 2 = 3
+1 + 1 - 1 + 2 = 3

*/

/*

方式一： 递归
方式二： 动态规划   0-1背包问题

*/
func findTargetSumWays(nums []int, target int) int {
	// write code here
	n := len(nums)
	if n == 0 {
		return 0
	}
	count := 0
	var dfs func(idx int, rest int)
	dfs = func(idx int, rest int) {
		// 索引走到最后，剩余的数为0, 个数加一， 并返回
		if idx == n {
			if rest == 0 {
				count += 1
			}
			return
		}
		// 加号
		dfs(idx+1, rest-nums[idx])
		// 减号
		dfs(idx+1, rest+nums[idx])
	}
	dfs(0, target)
	return count
}

func main03() {
	ways := findTargetSumWays([]int{1, 1, 1, 2}, 3)
	fmt.Println(ways)
}
