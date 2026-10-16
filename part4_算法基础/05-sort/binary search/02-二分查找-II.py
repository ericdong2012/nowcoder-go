"""
NC105 二分查找-II
https://www.nowcoder.com/practice/4f470d1d3b734f8aaf2afb014185b395

描述
请实现有重复数字的升序数组的二分查找
给定一个元素有序的（升序）长度为 n 的整型数组 nums 和一个目标值 target，
写一个函数搜索 nums 中的第一个出现的 target，如果目标值存在返回下标，否则返回 -1

示例1
输入：[1,2,4,4,5],4
返回值：2

示例2
输入：[1,2,4,4,5],3
返回值：-1

示例3
输入：[1,1,1,1,1],1
返回值：0
"""


def search(nums, target):
    """write code here"""
    left, right = 0, len(nums) - 1
    stack = []
    while left <= right:
        # if nums[left] == target:
        #     stack.append(left)
        #     left += 1
        # elif nums[right] == target:
        #     stack.append(right)
        #     right -= 1
        # else:
        #     left += 1
        #     right -= 1
        mid = (left + right) // 2
        if nums[mid] > target:
            right = mid - 1
        elif nums[mid] < target:
            left = mid + 1
        else:
            stack.append(mid)
            # error
            # left += 1
            # right -= 1
            # 只要找最小的索引，不是找完整的，右边有也不是最小，所以只是右边移动
            right = mid - 1

    # 异常判断
    if len(stack) == 0:
        return -1

    # 例如 [1, 2, 4, 4, 4, 6] 4，上述拿到的不一定是最小的索引，所以还需要处理
    min_v = stack[0]
    # for i in range(len(stack)):
    for i in range(1, len(stack)):
        if stack[i] <= min_v:
            min_v = stack[i]
    return min_v


if __name__ == "__main__":
    i = search([1, 2, 4, 4, 5], 4)
    print(i)
