package main

import "fmt"

/*
NC26 括号生成
https://www.nowcoder.com/practice/c9addb265cdf4cdd92c092c655d164ca?tpId=117&tqId=37748&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=591&title=

描述
给出n对括号，请编写一个函数来生成所有的由n对括号组成的合法组合。
例如，给出n=3，解集为：
"((()))", "(()())", "(())()", "()()()", "()(())"

数据范围：1 \le n \le 81≤n≤8
示例1
输入：
1
复制
返回值：
["()"]
复制
示例2
输入：
2
复制
返回值：
["(())","()()"]
*/

func generateParenthesis(n int) []string {
	res := []string{}
	// lRemain 是左括号使用次数，rRemain是右括号使用次数
	var dfs func(lRemain int, rRemain int, path string)
	dfs = func(lRemain int, rRemain int, path string) {
		if 2*n == len(path) {
			res = append(res, path)
			return
		}
		if lRemain < n {
			dfs(lRemain+1, rRemain, path+"(")
		}
		// 右括号还有剩余，且左括号使用次数大于右括号使用次数，此时可以使用右括号
		if lRemain > rRemain && rRemain < n {
			dfs(lRemain, rRemain+1, path+")")
		}
	}

	dfs(0, 0, "")
	return res
}

func main() {
	parenthesis := generateParenthesis(3)
	fmt.Println(parenthesis)
}
