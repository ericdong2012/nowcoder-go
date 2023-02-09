package main

import "fmt"

/*
BM69 把数字翻译成字符串
https://www.nowcoder.com/practice/046a55e6cd274cffb88fc32dba695668?tpId=295&tqId=1024831&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295


描述
有一种将字母编码成数字的方式：'a'->1, 'b->2', ... , 'z->26'。
我们把一个字符串编码成一串数字，再考虑逆向编译成字符串。
由于没有分隔符，数字编码成字母可能有多种编译结果，例如 11 既可以看做是两个 'a' 也可以看做是一个 'k' 。但 10 只可能是 'j' ，因为 0 不能编译成任何结果。
现在给一串数字，返回有多少种可能的译码结果

数据范围：字符串长度满足 0 < n \le 900<n≤90
进阶：空间复杂度 O(n)O(n)，时间复杂度 O(n)O(n)

示例1
输入：
"12"
返回值：
2
说明：
2种可能的译码结果（”ab” 或”l”）

示例2
输入：
"31717126241541717"
返回值：
192

说明：
192种可能的译码结果

*/

/*
对于普通数组1-9，译码方式只有一种，但是对于11-19，21-26，译码方式有可选择的两种方案，因此我们使用动态规划将两种方案累计。

step 1：用辅助数组dp表示前i个数的译码方法有多少种。
step 2：对于一个数，我们可以直接译码它，也可以将其与前面的1或者2组合起来译码：如果直接译码，则dp[i]=dp[i−1]如果组合译码，则dp[i]=dp[i−2]。
step 3：对于只有一种译码方式的，选上种dp[i−1]即可，对于满足两种译码方式（10，20不能）则是dp[i−1]+dp[i−2]
step 4：依次相加，最后的dp[length]即为所求答案。

  2 3 1 2 1 1  0  4  5
0 1 2 3 4 5 6  7  8  9  (2, 23, 231, 2312 ... )
1 1 2 2 4 6 10 6  6  6


class Solution:
    def solve(self , nums: str) -> int:
        #排除0
        if nums == "0":
            return 0
        #排除只有一种可能的10 和 20
        if nums == "10" or nums == "20":
            return 1
        #当0的前面不是1或2时，无法译码，0种
        for i in range(1, len(nums)):
            if nums[i] == '0':
                if nums[i - 1] != '1' and nums[i - 1] != '2':
                    return 0
        #辅助数组初始化为1
        dp = [1 for i in range(len(nums) + 1)]
        for i in range(2, len(nums) + 1):
            #在11-19，21-26之间的情况
            if (nums[i - 2] == '1' and nums[i - 1] != '0') or (nums[i - 2] == '2' and nums[i - 1] > '0' and nums[i - 1] < '7'):
                dp[i] = dp[i - 1] + dp[i - 2]
            else:
                dp[i] = dp[i - 1]

        return dp[len(nums)]
*/
func solve(nums string) int {
	// write code here
	if nums == "0" {
		return 0
	}

	//if nums == "10" || nums == "20" {
	//	return 1
	//}

	dp := make(map[int]int, len(nums))
	dp[0] = 1
	length := len(nums)
	for i := 1; i < length; i++ {
		if nums[i] != '0' {
			dp[i] = dp[i-1]
		}
		if nums[i-1] == '1' || (nums[i-1] == '2' && nums[i] <= '6') {
			if i > 1 {
				dp[i] += dp[i-2]
			} else {
				// i == 1 时才会进入
				dp[i] += 1
			}
		}
	}
	return dp[length-1]
}

func main() {
	//i := solve("231211045")    // 6
	i := solve("23121") // 6
	//i := solve("31717126241541717")   // 192
	//i := solve("10")    //1
	fmt.Println(i)
}
