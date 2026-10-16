"""
NC200 移动 0
https://www.nowcoder.com/practice/102586387caa4afcbad6f96affce9780

描述
给定一个数组，请你实现将所有 0 移动到数组末尾并且不改变其他数字的相对顺序。

示例1
输入：[1,2,0,3]
返回值：[1,2,3,0]

示例2
输入：[1,2,3]
返回值：[1,2,3]

示例3
输入：[0,0]
返回值：[0,0]
"""


def move_zeroes(nums):
    """write code here"""
    # 零的个数
    # x = 0
    # for i, v in enumerate(nums):
    #     if v == 0:
    #         x += 1
    #     elif x != 0:
    #         # 例如 走到 3，将前面的地方用当前非零数填充
    #         nums[i - x] = nums[i]
    #         # 将当前地方用 0 填充
    #         nums[i] = 0
    # return nums

    # 解法2
    if nums is None:
        return None

    j = 0
    for i in range(len(nums)):
        # 如果 nums[i] != 0，i、j 数据位置不会更改，还在原来的位置；等于 0，j 落后，等下次就会被交换
        if nums[i] != 0:
            nums[j], nums[i] = nums[i], nums[j]
            j += 1

    return nums


if __name__ == "__main__":
    zeroes = move_zeroes([1, 2, 0, 0, 3])
    print(zeroes)
