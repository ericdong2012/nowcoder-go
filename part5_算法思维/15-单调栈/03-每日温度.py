"""
NC208 每日温度
https://www.nowcoder.com/practice/1f54e163e6944cc7b8759cc09e9c78d8?tpId=117&tqId=39397&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=581,582,586&title=

描述
根据往后 n 天的天气预报，计算每一天需要等待几天才会出现一次更高的气温，如果往后都没有更高的气温，则用 0 补位。
例如往后三天的气温是 [1,2,3] ， 则输出 [1,1,0]

示例1
输入：
[1,2,3]
返回值：
[1,1,0]

示例2
输入：
[2,4,5,9,10,0,9]
返回值：
[1,1,1,1,0,1,0]

示例3
输入：
[3,1,4]
返回值：
[2,1,0]
"""


# def temperatures(daily_temperatures):
#     # 保存索引
#     stack = []
#     # 保存最终结果
#     res = [0] * len(daily_temperatures)
#     for i in range(len(daily_temperatures)):
#         # 前一位比后一位小
#         # 看 main 中的多个测试结果
#         while len(stack) != 0 and daily_temperatures[stack[-1]] < daily_temperatures[i]:
#             index = stack[-1]
#             stack.pop()
#             res[index] = i - index
#         stack.append(i)
#     return res


def temperatures(daily_temperatures):
    """write code here"""
    res = [0] * len(daily_temperatures)

    for i in range(len(daily_temperatures)):
        for j in range(i + 1, len(daily_temperatures)):
            # 在后面，找到一个比当前大的， 记录索引，相差的距离
            if daily_temperatures[j] > daily_temperatures[i]:
                res[i] = j - i
                break
            # 如果走到了最后, 则 记录索引，0
            if j == len(daily_temperatures) - 1:
                res[i] = 0
                break

    return res


if __name__ == "__main__":
    # ints = temperatures([1,2,3])
    # ints = temperatures([4,5,3])
    # ints = temperatures([3,1,4])
    ints = temperatures([3, 1, 2])
    print(ints)
