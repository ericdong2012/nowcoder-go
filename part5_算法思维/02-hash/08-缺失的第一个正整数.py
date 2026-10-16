"""
NC30 缺失的第一个正整数
https://www.nowcoder.com/practice/50ec6a5b0e4e45348544348278cdcee5

描述
给定一个无重复元素的整数数组 nums，请你找出其中没有出现的最小的正整数
进阶：空间复杂度 O(1)，时间复杂度 O(n)

示例1
输入：[1,0,2]
返回值：3

示例2
输入：[-2,3,4,1,5]
返回值：2

示例3
输入：[4,5,6,8,9]
返回值：1
"""

# 方法2
# 排序，看首位元素
# 如果首位大于等于 2，则最小的是 1
# 否则，以 1 为起始，最后一位为终止遍历，判断是否在数组中，不在，则该数就是最小正整数


def min_number_disappeared(nums):
    """write code here"""
    # 构造数据，生成一个长度是 5e5，值为 0 的数组
    temp = [0] * 500000
    # 更改 temp 中的 value
    temp[0] = 1
    for v in nums:
        # 大于 0 的数添加到新数组中，小于等于 0 的自动过滤，
        # 最终的结果是有的地方有 1，有的没有：[1, 1, 0, 0, 1，。。。。]
        if v > 0:
            temp[v] = 1
    # 遍历第一个 0，说明该值在原有 nums 中不存在
    for k, v in enumerate(temp):
        if v == 0:
            return k
    return 1

    """
    方法2
    def min_number_disappeared(nums):
        nums.sort()
        if nums[0] >= 2:
            return 1
        temp = {}
        for v in nums:
            temp[v] = 1
        for i in range(1, nums[-1]):
            if i not in temp:
                return i
        # 如果全部遍历完还没有找到，最大的加上 1
        return nums[-1] + 1
    """


if __name__ == "__main__":
    nums = [1, 23, 3, 4]
    disappeared = min_number_disappeared(nums)
    print(disappeared)
