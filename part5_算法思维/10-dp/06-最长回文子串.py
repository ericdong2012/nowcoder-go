"""
NC17 最长回文子串
https://www.nowcoder.com/practice/b4525d1d84934cf280439aeecc36f4af?tpId=117&tqId=37789

描述
对于长度为n的一个字符串A（仅包含数字，大小写英文字母），请设计一个高效算法，计算其中最长回文子串的长度。
要求: 空间复杂度 O(1)，时间复杂度 O(n^2)
进阶: 空间复杂度 O(n)，时间复杂度 O(n)

示例1
输入："ababc"
返回值：3

示例2
输入："abbba"
返回值：5
"""

# 中心扩散法
#
# 回文串，有着左右对称的特征，从首尾一起访问，遇到的元素都是相同的。
# 但是我们这里正是要找最长的回文串，并不事先知道长度，怎么办？判断回文的过程是从首尾到中间，
# 那我们找最长回文串可以逆着来，从中间延伸到首尾，这就是中心扩展法
#
# step 1：遍历字符串每个字符。
# step 2：以每次遍历到的字符为中心（分奇数长度和偶数长度两种情况），不断向两边扩展。
# step 3：如果两边都是相同的就是回文，不断扩大到最大长度即是以这个字符（或偶数两个）为中心的最长回文子串。
# step 4：我们比较完每个字符为中心的最长回文子串，取最大值即可。


# 中心扩散法
def getLongestPalindrome1(A):
    # "baabccc"
    # 看运行过程 最后一个元素，倒数第二个元素...
    maxLen = 1
    for i in range(len(A) - 1, -1, -1):
        # 两种情况：可能是类似 aba 的字符串，也可能是类似 abba 的情况
        # 偶数
        left = run2(A, i, i)
        # 奇数
        right = run2(A, i, i + 1)
        maxLen = max2(max2(maxLen, left), right)
    return maxLen


def run2(s, begin, end):
    while begin >= 0 and end < len(s) and s[begin] == s[end]:
        begin -= 1
        end += 1
    return end - begin - 1


def max2(a, b):
    if a < b:
        return b
    return a


if __name__ == "__main__":
    print(getLongestPalindrome1("ababc"))
