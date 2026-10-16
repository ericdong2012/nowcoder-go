"""
可以和part1/04-stack/滑动窗口的最大值 对比(用的单调栈)
"""


def max_value(temp):
    result = temp[0]
    for i in range(len(temp)):
        if temp[i] > result:
            result = temp[i]

    return result


def max_window(nums, k):
    if len(nums) == 0 or len(nums) < k:
        return []
    stack = []
    # 123456  3
    for i in range(len(nums) - k + 1):
        result = max_value(nums[i:i + k])
        stack.append(result)

    return stack


if __name__ == "__main__":
    print(max_window([1, 2, 3, 4, 5, 6], 3))
