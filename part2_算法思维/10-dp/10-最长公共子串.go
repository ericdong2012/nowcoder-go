package main

import (
	"fmt"
	"strings"
)

/*
BM66 最长公共子串   顺序是连续的
https://www.nowcoder.com/practice/f33f5adc55f444baa0e0ca87ad8a6aac?tpId=295&tqId=991150&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295


描述
给定两个字符串str1和str2, 输出两个字符串的最长公共子串
题目保证str1和str2的最长公共子串存在且唯一。

要求： 空间复杂度 O(n^2)，时间复杂度 O(n^2)

示例1
输入：
"1AB2345CD","12345EF"
返回值：
"2345"

*/

/*
双指针 + 中心化扩散
class Solution:
    def LCS(self , str1: str, str2: str) -> str:
        # 让str1为较长的字符串
        if len(str1) < len(str2):
            str1, str2 = str2, str1
        res = ''
        max_len = 0
        # 遍历str1的长度
        for i in range(len(str1)):
            # 查找是否存在
            if str1[i - max_len : i + 1] in str2:
                res = str1[i - max_len : i + 1]
                max_len += 1
        return res


动态规划
step 1：我们可以用dp[i][j]dp[i][j]dp[i][j]表示在str1中以第iii个字符结尾在str2中以第jjj个字符结尾时的公共子串长度，
step 2：遍历两个字符串填充dp数组，转移方程为：如果遍历到的该位两个字符相等，则此时长度等于两个前一位长度+1，dp[i][j]=dp[i−1][j−1]+1dp[i][j] = dp[i - 1][j - 1] + 1dp[i][j]=dp[i−1][j−1]+1，如果遍历到该位时两个字符不相等，则置为0，因为这是子串，必须连续相等，断开要重新开始。
step 3：每次更新dp[i][j]dp[i][j]dp[i][j]后，我们维护最大值，并更新该子串结束位置。
step 4：最后根据最大值结束位置即可截取出子串。


class Solution:
    def LCS(self , str1: str, str2: str) -> str:
        #dp[i][j]表示到str1第i个个到str2第j个为止的公共子串长度
        dp = [[0] * (len(str2) + 1) for i in range(len(str1) + 1)]
        max = 0
        pos = 0
        for i in range(1, len(str1) + 1):
            for j in range(1, len(str2) + 1):
                #如果该两位相同
                if str1[i - 1] == str2[j - 1]:
                    #则增加长度
                    dp[i][j] = dp[i - 1][j - 1] + 1
                else:
                    #该位置为0
                    dp[i][j] = 0
                #更新最大长度
                if dp[i][j] > max:
                    max = dp[i][j]
                    pos = i - 1
        return str1[pos - max + 1: pos + 1]

*/

func LCS(str1 string, str2 string) string {
	// write code here
	//dp := make([][]int, len(str1))
	//for i := 0; i < len(str1); i++ {
	//	dp[i] = make([]int, len(str2))
	//}
	//maxV := 0
	//pos := 0
	//for i := 1; i < len(str1); i++ {
	//	for j := 1; j < len(str2); j++ {
	//		if str1[i-1] == str2[j-1] {
	//			dp[i][j] = dp[i-1][j-1] + 1
	//		} else {
	//			dp[i][j] = 0
	//		}
	//
	//		if dp[i][j] > maxV {
	//			maxV = dp[i][j]
	//			pos = i - 1
	//		}
	//	}
	//}
	//
	//return str1[pos-maxV+1 : pos+1]
	if len(str1) < len(str2) {
		str1, str2 = str2, str1
	}

	res := ""
	pos := 0
	for i := 0; i < len(str1); i++ {
		// 非固定的滑动窗口
		// 双指针 + 中心扩散， 将头部固定，尾部不断扩张
		if strings.Contains(str2, str1[i-pos:i+1]) {
			res = str1[i-pos : i+1]
			pos += 1
		}
	}

	return res
}

func main10() {
	lcs := LCS("1AB2345CD", "12345EF")
	//lcs := LCS("22222", "22222")
	fmt.Println(lcs)
}
