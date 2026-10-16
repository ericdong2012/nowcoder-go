"""
直方图最大矩形
https://www.nowcoder.com/practice/bfaac4eebe5947af80912d1bcfcf2baa

给定一个数组heights，长度为n，height[i]是在第i点的高度，那么height[i]表示的直方图，能够形成的最大矩形是多少?
1.每个直方图宽度都为1
2.直方图都是相邻的
3.如果不能形成矩形，返回0即可
4.保证返回的结果不会超过231-1

示例1
输入：
[3,4,7,8,1,2]
返回值：
14

示例2
输入：
[1,7,3,2,4,5,8,2,7]
返回值：
16
"""

# 和接雨水，盛水最多的容器 比较看
# 和移掉K位数字后最小     对比看


def largest_rectangle_area(height):
    """write code here"""
    # 单调栈 单调递增，出栈时计算最大面积
    if len(height) < 1:
        return 0
    if len(height) == 1:
        return height[0]
    # 加0作用是弹出栈所有的元素。
    height = height + [0]
    # 装索引
    stack = []
    # stack.append(-1)  # 最边上设置为-1保证不能出栈 保证栈非空
    # 结果
    res = 0
    # [3,4,7,8,1,2]
    for i in range(len(height)):
        # 出栈逻辑, 前一个值 大于 当前值, 说明开始降序
        # 反复的进入该循环，计算最大值
        while len(stack) != 0 and height[stack[-1]] > height[i]:
            # 最难理解的地方
            # 高度是前一个的，不停往前走
            hei = height[stack[-1]]
            # 先弹栈，看看左边还有没有
            stack.pop()
            if len(stack) == 0:
                wid = i
            else:
                # 当前存在比弹出的还小的， 所以左边界是更矮的那个索引加1
                # 宽度 = i- 左边界
                wid = i - (stack[-1] + 1)

            res = max(res, hei * wid)
        # 入栈 入下标
        stack.append(i)

    return res


if __name__ == "__main__":
    area = largest_rectangle_area([3, 4, 7, 8, 1, 2])
    print(area)
