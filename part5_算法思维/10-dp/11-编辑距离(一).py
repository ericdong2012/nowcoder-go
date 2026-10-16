"""
BM75 编辑距离(一)
https://www.nowcoder.com/practice/6a1483b5be1547b1acd7940f867be0da?tpId=295&tqId=2294660

描述
给定两个字符串 str1 和 str2 ，请你算出将 str1 转为 str2 的最少操作数。
你可以对字符串进行3种操作：
1.插入一个字符
2.删除一个字符
3.修改一个字符。
保证字符串中只出现小写英文字母。

示例1
输入："nowcoder", "new"
返回值：6

示例2
输入："intention", "execution"
返回值：5
"""

# 建立dp方程  先画这个方程
#    0 n o w c o d e r
# 0   0 1 2 3 4 5 6 7 8
# n   1 0 1 2 3 4 5 6 7
# e   2 1 1 2 3 4 5 5 6
# w   3 2 2 1 2 3 4 5 6
#
# step 1：初始条件：假设第二个字符串为空，那很明显第一个字符串子串每增加一个字符，编辑距离就加1，这步操作是删除；
#         同理，假设第一个字符串为空，那第二个字符串每增加一个字符，编剧距离就加1，这步操作是添加。
# step 2：状态转移：遍历第一个字符串的每个长度，对应第二个字符串的每个长度。
#         如果遍历到str1[i]和 str2[j]的位置，这两个字符相同，这多出来的字符就不用操作，操作次数与两个子串的前一个相同，
#              因此有dp[i][j]=dp[i-1][j-1]；
#         如果这两个字符不相同，那么这两个字符需要编辑，选取这三种情况的最小值增加一个编辑距离，
#              即dp[i][j]=min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1]))+1


def editDistance(str1, str2):
    # write code here
    m, n = len(str1), len(str2)
    # 初始化一个二维数组
    dp = [[0] * (n + 1) for _ in range(m + 1)]
    # 思考为啥要这样初始化: 因为修改一个要 +1
    # 第0列初始化
    for i in range(1, m + 1):
        dp[i][0] = dp[i - 1][0] + 1
    # 第0行初始化
    for i in range(1, n + 1):
        dp[0][i] = dp[0][i - 1] + 1

    for i in range(1, m + 1):
        for j in range(1, n + 1):
            if str1[i - 1] == str2[j - 1]:
                dp[i][j] = dp[i - 1][j - 1]
            else:
                # 最小距离就是1， dp[i-1][j-1] dp[i-1][j] dp[i][j-1] 有可能为0，所以需要加上1
                dp[i][j] = min4(dp[i - 1][j - 1], min4(dp[i - 1][j], dp[i][j - 1])) + 1

    return dp[m][n]


def min4(a, b):
    if a < b:
        return a
    return b


def main13():
    distance = editDistance("nowcoder", "new")
    print(distance)


if __name__ == "__main__":
    main13()
