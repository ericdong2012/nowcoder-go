"""
NC104 比较版本号
https://www.nowcoder.com/practice/2b317e02f14247a49ffdbdba315459e7

描述
牛客项目发布项目版本时会有版本号，比如 1.02.11，2.14.4 等等
现在给你 2 个版本号 version1 和 version2，请你比较他们的大小
版本号是由修订号组成，修订号与修订号之间由一个 "." 连接。1 个修订号可能有多位数字组成，
修订号可能包含前导 0，且是合法的。例如，1.02.11，0.1，0.2 都是合法的版本号
每个版本号至少包含 1 个修订号。

比较规则：
一. 比较版本号时，请按从左到右的顺序依次比较它们的修订号。比较修订号时，只需比较忽略任何前导零后的整数值。
    比如 "0.1" 和 "0.01" 的版本号是相等的
二. 如果版本号没有指定某个下标处的修订号，则该修订号视为 0。例如，"1.1" 的版本号小于 "1.1.1"。
    因为 "1.1" 的版本号相当于 "1.1.0"，第 3 位修订号的下标为 0，小于 1
三. version1 > version2 返回 1，如果 version1 < version2 返回 -1，不然返回 0。

示例1
输入："1.1","2.1"
返回值：-1
"""

# 忽略 "."，不停遍历，处理好进位，最后比较大小
# 直接当字符串，直接加，最后转成数字比较大小


def compare(version1, version2):
    i, j = 0, 0
    while i < len(version1) or j < len(version2):
        num1 = 0
        while i < len(version1) and version1[i] != '.':
            v = int(version1[i])
            num1 = num1 * 10 + v
            i += 1

        num2 = 0
        while j < len(version2) and version2[j] != '.':
            v = int(version2[j])
            num2 = num2 * 10 + v
            j += 1

        if num1 > num2:
            return 1
        elif num1 < num2:
            return -1

        i += 1
        j += 1

    return 0
