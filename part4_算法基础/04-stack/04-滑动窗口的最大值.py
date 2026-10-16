"""
NC82 滑动窗口的最大值
https://www.nowcoder.com/practice/1624bc35a45c42c0bc17d17fa0cba788

描述
给定一个长度为 n 的数组 nums 和滑动窗口的大小 size，找出所有滑动窗口里数值的最大值。

例如，如果输入数组 {2,3,4,2,6,2,5,1} 及滑动窗口的大小 3，那么一共存在 6 个滑动窗口，
他们的最大值分别为 {4,4,6,6,6,5}。

示例1
输入：[2,3,4,2,6,2,5,1],3
返回值：[4,4,6,6,6,5]
"""

# 双向队列，单调队列


def max_in_windows(num, size):
    """
    write code here
    单调栈
    """
    # 单调栈
    # if size > len(num) or size == 0:
    #     return None
    # n = len(num)
    # # 存的是索引
    # deque = []
    # # 结果
    # res = []
    # for i in range(n):
    #     # 不断移除队列中比当前元素小的元素，最终将当前窗口的最大值的索引存起来
    #     while len(deque) > 0 and num[deque[-1]] < num[i]:
    #         deque.pop()
    #     deque.append(i)
    #     # 判断头部节点是否还在滑动窗口内 ??? 比如: [1, -1], 1
    #     # if deque[0] + size <= i:
    #     #     deque.pop(0)
    #     # i 到了滑动窗口外面的下一个元素，同时队列中保存的索引是当前索引 - 长度，
    #     # 也就是滑动窗口中第一个元素就是最大值
    #     if i >= size and deque[0] == i - size:
    #         deque.pop(0)
    #     # 长度达到 size
    #     if i + 1 >= size:
    #         res.append(num[deque[0]])
    # return res

    # 选择排序的变换，只是要固定长度
    if size > len(num) or size == 0:
        return None
    res = []
    # 1. 要固定长度，最后的等于号不要忘
    # 2. 内部起点，终点要想清楚
    for j in range(len(num) - size + 1):
        max_value = num[j]
        # for i in range(j + 1, len(num) - size + 1):
        # 想清楚起点，终点，一定是加上
        # for i in range(j + 1, j + size):
        for i in range(j, j + size):
            if max_value < num[i]:
                max_value = num[i]
        res.append(max_value)

    return res


if __name__ == "__main__":
    L = max_in_windows([2, 3, 4, 2, 6, 2, 5, 1], 3)
    print(L)
    # temp = [2, 3, 4, 2, 6, 2, 5, 1]
    # print(temp[:-1])
