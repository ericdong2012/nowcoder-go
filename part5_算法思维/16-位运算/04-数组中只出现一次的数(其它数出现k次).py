"""
NC156 数组中只出现一次的数（其它数出现k次）
https://www.nowcoder.com/practice/5d3d74c3bf7f4e368e03096bb8857871?tpId=117&tqId=37866&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=5074&title=


描述
给定一个长度为 n 的整型数组 arr 和一个整数 k(k>1) 。
已知 arr 中只有 1 个数出现一次，其他的数都出现 k 次。
请返回只出现了 1 次的数。

进阶：时间复杂度 O(32n)，空间复杂度 O(1)


示例1
输入：
[5,4,1,1,5,1,5],3
返回值：
4

示例2
输入：
[2,2,1],2
返回值：
1

"""


def foundOnceNumber(arr, k):
    # write code here
    # 逐位统计：对每个二进制位，统计所有数字在该位上 1 的个数之和
    # 出现 k 次的数在该位的贡献是 k 的倍数，所以 sum % k 就是只出现一次的数在该位的值
    binSum = [0] * 32
    for i in range(32):
        s = 0
        for num in arr:
            s += (num >> i) & 1
        binSum[i] = s
    res = 0
    for i in range(32):
        if binSum[i] % k != 0:
            res += 1 << i
    return res


# 哈希表解法（参考）
# def foundOnceNumber(arr, k):
#     tempMap = {}
#     for i in range(len(arr)):
#         tempMap[arr[i]] = tempMap.get(arr[i], 0) + 1
#     for i, v in tempMap.items():
#         if v % k != 0:
#             return i
#     return -1


if __name__ == "__main__":
    number = foundOnceNumber([5, 4, 1, 1, 5, 1, 5], 3)
    print(number)
