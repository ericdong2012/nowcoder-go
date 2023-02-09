package main

import "fmt"

/*
BM75 编辑距离(一)
https://www.nowcoder.com/practice/6a1483b5be1547b1acd7940f867be0da?tpId=295&tqId=2294660&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295


描述
给定两个字符串 str1 和 str2 ，请你算出将 str1 转为 str2 的最少操作数。
你可以对字符串进行3种操作：
1.插入一个字符
2.删除一个字符
3.修改一个字符。

字符串长度满足 1 \le n \le 1000 \1≤n≤1000  ，保证字符串中只出现小写英文字母。
示例1
输入：
"nowcoder","new"
复制
返回值：
6
复制
说明：
"nowcoder"=>"newcoder"(将'o'替换为'e')，修改操作1次
"nowcoder"=>"new"(删除"coder")，删除操作5次
示例2
输入：
"intention","execution"
复制
返回值：
5
复制
说明：
一种方案为:
因为2个长度都是9，后面的4个后缀的长度都为"tion"，于是从"inten"到"execu"逐个修改即可
示例3
输入：
"now","nowcoder"
复制
返回值：
5


*/

/*
非常标准的动态规划

二维数组计算个数
建立dp方程

	0 n o w c o d e r
0
n
e
w

step 1：初始条件： 假设第二个字符串为空，那很明显第一个字符串子串每增加一个字符，编辑距离就加1，这步操作是删除；同理，假设第一个字符串为空，那第二个字符串每增加一个字符，编剧距离就加1，这步操作是添加。
step 2：状态转移： 状态转移肯定是将dp矩阵填满，那就遍历第一个字符串的每个长度，对应第二个字符串的每个长度。如果遍历到str1[i]和 str2[j]的位置，这两个字符相同，
这多出来的字符就不用操作，操作次数与两个子串的前一个相同，因此有dp[i][j]=dp[i−1][j−1]；
如果这两个字符不相同，那么这两个字符需要编辑，但是此时的最短的距离不一定是修改这最后一位，也有可能是删除某个字符或者增加某个字符，
因此我们选取这三种情况的最小值增加一个编辑距离，即dp[i][j]=min(dp[i−1][j−1],min(dp[i−1][j],dp[i][j−1]))+1

class Solution:
    def editDistance(self , str1: str, str2: str) -> int:
        n1 = len(str1)
        n2 = len(str2)
        #dp[i][j]表示到str1[i]和str2[j]为止的子串需要的编辑距离
        dp = [[0] * (n2 + 1) for i in range(n1 + 1)]
        #初始化边界
        for i in range(1, n1 + 1):
            dp[i][0] = dp[i - 1][0] + 1
        for i in range(1, n2 + 1):
            dp[0][i] = dp[0][i - 1] + 1
        #遍历第一个字符串的每个位置
        for i in range(1, n1 + 1):
            #对应第二个字符串每个位置
            for j in range(1, n2 + 1):
                #若是字符相同，此处不用编辑
                if str1[i - 1] == str2[j - 1]:
                    #直接等于二者前一个的距离
                    dp[i][j] = dp[i - 1][j - 1]
                else:
                    #选取最小的距离加上此处编辑距离1
                    dp[i][j] = min(dp[i - 1][j - 1], min(dp[i - 1][j], dp[i][j - 1])) + 1
        return dp[n1][n2]

*/

func editDistance(str1 string, str2 string) int {
	// write code here
	m, n := len(str1), len(str2)
	dp := make([][]int, m+1)

	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i < m+1; i++ {
		dp[i][0] = dp[i-1][0] + 1
	}

	for i := 1; i < n+1; i++ {
		dp[0][i] = dp[0][i-1] + 1
	}

	for i := 1; i < len(str1)+1; i++ {
		for j := 1; j < len(str2)+1; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min4(dp[i-1][j-1], min4(dp[i-1][j], dp[i][j-1])) + 1
			}
		}
	}
	return dp[m][n]
}

func min4(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	distance := editDistance("nowcoder", "new")
	fmt.Println(distance)
}
