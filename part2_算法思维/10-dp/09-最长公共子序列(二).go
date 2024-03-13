package main

import "fmt"

/*
BM65 最长公共子序列(二)  顺序可以是非连续的，但是顺序要一致，要将数据保存下来
https://www.nowcoder.com/practice/6d29638c85bb4ffd80c020fe244baf11?tpId=295&tqId=991075&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295

描述
给定两个字符串str1和str2，输出两个字符串的最长公共子序列。如果最长公共子序列为空，则返回"-1"。目前给出的数据，仅仅会存在一个最长的公共子序列

示例1
输入：
"1A2C3D4B56","B1D23A456A"
返回值：
"123456"

示例2
输入：
"abc","def"
返回值：
"-1"

示例3
输入：
"abc","abc"
返回值：
"abc"

示例4
输入：
"ab",""
返回值：
"-1"

*/

/*
二维数组
dp方程，找规律， max

动态规划 + (递归)栈
step 1：优先检查特殊情况。
step 2：获取最长公共子序列的长度可以使用动态规划，我们以dp[i][j]表示在s1中以iii结尾，s2中以jjj结尾的字符串的最长公共子序列长度。
step 3：遍历两个字符串的所有位置，开始状态转移：若是iii位与jjj位的字符相等，则该问题可以变成1+dp[i−1][j−1]，即到此处为止最长公共子序列长度由前面的结果加1。
step 4：若是不相等，说明到此处为止的子串，最后一位不可能同时属于最长公共子序列，毕竟它们都不相同，因此我们考虑换成两个子问题，dp[i][j−1]或者dp[i−1][j]，我们取较大的一个就可以了。
step 5：得到最长长度后，获取不需要第二个辅助数组b，直接从dp数组最后一位开始，每次比较当前位置与其左、上、左上的关系，然后将符合要求的字符加入栈中，符合要求即来自dp表格左上方的字符。
step 6：最后将栈中的字符拼接即可得到最长公共子序列，注意检查子序列是否为空。

class Solution:
    def LCS(self , s1: str, s2: str) -> str:
        #特殊情况
        if s1 is None or s2 is None:
            return "-1"
        len1 = len(s1)
        len2 = len(s2)
        #dp[i][j]表示第一个字符串到第i位，第二个字符串到第j位为止的最长公共子序列长度
        dp = [[0] * (len2 + 1) for i in range(len1 + 1)]
        #遍历两个字符串每个位置求的最长长度
        for i in range(1, len1 + 1):
            for j in range(1, len2 + 1):
                #遇到两个字符相等
                if s1[i - 1] == s2[j -1]:
                    #来自于对角线
                    dp[i][j] = dp[i - 1][j - 1] + 1
                #遇到的两个字符不同
                else:
                    #来自左边或者上方的最大值
                    dp[i][j] = max(dp[i - 1][j], dp[i][j - 1])

        #从动态规划数组末尾开始
        i = len1
        j = len2
        s = []
        while dp[i][j] != 0:
            #来自于左方向
            if dp[i][j] == dp[i - 1][j]:
                i = i - 1
            #来自于上方向
            elif dp[i][j] == dp[i][j - 1]:
                j = j - 1
            #来自于左上方向
            elif dp[i][j] > dp[i - 1][j - 1]:
                i = i - 1
                j = j - 1
                #只有左上方向才是字符相等的情况，入栈，逆序使用
                s.append(s1[i])

        res = ""
        #拼接子序列
        while len(s) != 0:
            res += s[-1]
            s.pop()
        #如果两个完全不同，返回字符串为空，则要改成-1
        if res is None or res == "":
            return "-1"
        else:
            return res

时间复杂度：O(n^2)，最坏复杂度为构造辅助数组dp两层循环
空间复杂度：O(n^2)，辅助二维数组dp与栈空间最大为O(n^2)
*/


// 跟背包问题，兑换零钱问题 是一类解决方案
func LCS2(s1 string, s2 string) string {
	/*
		   1 A 2 C 3 D 4 B 5 6
		B
		1
		D
		2
		3
		A
		4
		5
		6
		A
	*/
	if len(s2) == 0 {
		return "-1"
	}
	dp := make([]string, len(s2))
	var t, t1 string
	for i := 0; i < len(s1); i++ {
		t = ""
		for j := 0; j < len(s2); j++ {
			t1 = dp[j]
			// 执行到最后， dp的最后一位是所有复合要求的字符
			// 字符相等
			if s2[j] == s1[i] {
				dp[j] = t + string(s2[j])
			} else {
				// 如果当前长度小于前一个长度，会将当前长度更新, 此处更新比较妙
				// 调试看
				// 1, 1, 1, 1, 1, 1, 1 。。。
				// 1, 1, 1A, 1A, 1A ...
				// 1, 12 ,1A, 1A, 1A ...
				// ...
				// 1, 123, 123, 123, 123...

				// j > 0 是因为 有dp[j-1]
				if j > 0 && len(dp[j]) < len(dp[j-1]) {
					dp[j] = dp[j-1]
				}
			}
			t = t1
		}
	}
	// 如果最后一位是空 == dp[ len(dp) - 1 ]
	if dp[len(s2)-1] == "" {
		return "-1"
	}
	// 返回最后一位
	return dp[len(s2)-1]
}

func main() {
	lcs2 := LCS2("1A2C3D4B56", "B1D23A456A")
	fmt.Println(lcs2)
}
