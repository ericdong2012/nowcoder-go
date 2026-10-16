"""
NC82 滑动窗口的最大值
https://www.nowcoder.com/practice/1624bc35a45c42c0bc17d17fa0cba788

描述
给定一个长度为 n 的数组 nums 和滑动窗口的大小 size，找出所有滑动窗口里数值的最大值。

示例1
输入：[2,3,4,2,6,2,5,1],3
返回值：[4,4,6,6,6,5]
"""

# 双向队列，单调队列
# 可以和 part1 中的比较


def max_in_windows(num, size):
    """write code here"""
    # n = len(num)
    # # 单调栈，存的是索引(数据较大的索引)
    # deque = []
    # # 结果
    # res = []
    # for i in range(n):
    #     # deque 中存的是数据较大的索引，否则一直缩短
    #     while len(deque) > 0 and num[deque[-1]] < num[i]:
    #         deque.pop()
    #     deque.append(i)
    #     # 判断头部节点是否还在滑动窗口内
    #     if deque[0] + size <= i:
    #         deque.pop(0)
    #     # 满足了一个滑动窗口
    #     if i + 1 >= size:
    #         res.append(num[deque[0]])
    # return res

    # 选择排序的变换
    if size > len(num) or size == 0:
        return None
    res = []
    max_value = 0
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
