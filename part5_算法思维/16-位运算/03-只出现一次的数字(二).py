"""
NC227 只出现一次的数字(二)
https://www.nowcoder.com/practice/85cb47fc0c6c483fab7e5cefab54d9e5?tpId=117&tqId=39452&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=5074&title=


描述
给定一个整数数组，数组中有一个数出现了一次，其他数出现了三次，请找出只出现了一次的数。

示例1
输入：
[1]
返回值：
1

示例2
输入：
[1,2,2,2]
返回值：
1

"""


def singleNumber(nums):
    # write code here
    # 逐位统计：对每个二进制位，统计所有数字在该位上 1 的个数之和
    # 出现 3 次的数在该位的贡献是 3 的倍数，所以 sum % 3 就是只出现一次的数在该位的值
    arr = [0] * 32
    for v in nums:
        for i in range(32):
            arr[i] += (v >> i) & 1

    # NOTE: Go 原版用 int32 做符号位自然溢出处理负数结果。
    # Python int 是任意精度，需要手动处理符号位（第 31 位为 1 时结果是负数）。
    ans = 0
    for i in range(31, -1, -1):
        if arr[i] % 3 == 1:
            ans |= (1 << i)
    # 若符号位（第31位）为1，则是负数，需要将 32 位补码转换为 Python 负整数
    if arr[31] % 3 == 1:
        ans -= (1 << 32)
    return ans


# 哈希表解法（参考）
# def singleNumber(nums):
#     tempMap = {}
#     for i in range(len(nums)):
#         tempMap[nums[i]] = tempMap.get(nums[i], 0) + 1
#     for k, v in tempMap.items():
#         if v == 1:
#             return k
#     return -1


if __name__ == "__main__":
    number = singleNumber([2, 2, 1, 2])
    print(number)
