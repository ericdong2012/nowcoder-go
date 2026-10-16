"""
NC37 合并区间
https://www.nowcoder.com/practice/69f4e5b7ad284a478777cb2a17fb5e6a

描述
给出一组区间，请合并所有重叠的区间。
请保证合并后的区间按区间起点升序排列。

要求：空间复杂度 O(n)，时间复杂度 O(nlogn)
进阶：空间复杂度 O(val)，时间复杂度O(val)

示例1
输入：
[[10,30],[20,60],[80,100],[150,180]]
返回值：
[[10,60],[80,100],[150,180]]

示例2
输入：
[[0,10],[10,20]]
返回值：
[[0,20]]
"""


class Interval:
    def __init__(self, start=0, end=0):
        self.start = start
        self.end = end


def merge(intervals):
    if len(intervals) == 0:
        return None
    # 将所有数据按起点升序排列
    intervals.sort(key=lambda x: x.start)

    res = []
    res.append(intervals[0])
    # 判断首位，末尾，考虑是否合并
    for i in range(1, len(intervals)):
        """
        几种可能的情形:
        [10, 20] [20, 60]
        [10, 30] [20, 40]
        [10, 30] [20, 25]

        [10, 30] [40, 50]
        """

        # 有交叉
        if res[-1].end >= intervals[i].start:
            # 只需更新原有的，不需要删除或者添加
            res[-1].end = max(res[-1].end, intervals[i].end)
        else:
            # 无交叉
            res.append(intervals[i])

    return res

    # 思想很好，但是会超时
    # if len(intervals) == 0:
    #     return None
    #
    # intervals.sort(key=lambda x: x.start)
    #
    # i = 0
    # while i < len(intervals) - 1:
    #     # 有交叉, 判断下末尾，并移除掉后面的一位
    #     if intervals[i].end >= intervals[i + 1].start:
    #         intervals[i].end = max(intervals[i].end, intervals[i + 1].end)
    #         intervals.pop(i + 1)
    #     else:
    #         # 无交叉，保持原样
    #         i += 1
    #
    # return intervals


if __name__ == "__main__":
    ivs = [Interval(10, 30), Interval(20, 60), Interval(80, 100), Interval(150, 180)]
    result = merge(ivs)
    print([(iv.start, iv.end) for iv in result])
