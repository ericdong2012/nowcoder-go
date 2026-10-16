"""
NC202 长度最小的连续子数组
https://www.nowcoder.com/practice/10dd5f8c5d984aa3bd69788d86aaef23

描述
给定一个数组 nums 和一个正整数 target，找出满足和大于等于 target 的长度最短的连续子数组并返回其长度，
如果不存在这种子数组则返回 0。

示例1
输入：[1,2,4,4,1,1,1],9
返回值：3

示例2
输入：[1,4,4,4,1,1,1],3
返回值：1
"""

# 二分+前缀和
#
# 双指针+窗口
# def minSubarray(self, nums: List[int], target: int) -> int:
#     # write code here
#     # 左右指针移动滑窗求长度
#     left = 0          # 左指针
#     right = 0         # 右指针
#     sums = 0          # 左右指针滑窗内部的值之和
#     ans = len(nums)   # 这个是用来比较滑窗内长短的值，可以设为无限大，这里我设为 nums 长度
#     while right < len(nums):     # 右指针先走，走完整个 nums
#         while right < len(nums) and sums < target:
#             sums += nums[right]
#             right += 1
#         while left < right and sums >= target:
#             ans = min(ans, right - left)
#             sums -= nums[left]
#             left += 1
#     return ans


def _min4(a, b):
    if a < b:
        return a
    else:
        return b


def min_subarray(nums, target):
    """write code here"""
    left, right, sums = 0, 0, 0
    ans = len(nums)

    while right < len(nums):
        while right < len(nums) and sums < target:
            sums += nums[right]
            right += 1
        while right <= len(nums) and sums >= target:
            # 最短长度
            ans = _min4(ans, right - left)
            # 总和中减去最左侧的数字，往右移动
            sums -= nums[left]
            left += 1

    return ans


if __name__ == "__main__":
    subarray = min_subarray([1, 2, 4, 4, 1, 1, 1], 9)
    # subarray = min_subarray([1, 1, 1, 2, 4], 7)
    print(subarray)
