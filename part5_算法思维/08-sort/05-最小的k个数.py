"""
https://www.nowcoder.com/practice/6a296eb82cf844ca8539b57c23e6e9bf

给定一个长度为 n 的可能有重复值的数组，找出其中不去重的最小的 k 个数。例如数组元素是4,5,1,6,2,7,3,8这8个数字，则最小的4个数字是1,2,3,4(任意顺序皆可)。


示例1
输入：
[4,5,1,6,2,7,3,8],4
返回值：
[1,2,3,4]
说明：
返回最小的4个数即可，返回[1,3,2,4]也可以


示例2
输入：
[1],0
返回值：
[]


示例3
输入：
[0,1,2,1,2],3
返回值：
[0,1,1]
"""


def get_least_numbers_solution(input_arr, k):
    # NOTE: 原始 Go 条件 `len(input) < k && k == 0` 逻辑有误（&& 应为 ||），
    # 当 len < k 且 k != 0 时不会返回 []，会越界。修正为 ||。
    if len(input_arr) < k or k == 0:
        return []

    # for i in range(k):
    #     for k in range(len(input_arr) - 1, i, -1):
    #         if input_arr[k] < input_arr[k - 1]:
    #             input_arr[k], input_arr[k - 1] = input_arr[k - 1], input_arr[k]

    for i in range(len(input_arr)):
        for j in range(i + 1, len(input_arr)):
            if input_arr[i] > input_arr[j]:
                input_arr[i], input_arr[j] = input_arr[j], input_arr[i]
    return input_arr[:k]


# def get_least_numbers_solution(input_arr, k):
#     # write code here
#     if len(input_arr) < k:
#         return []
#     input_arr.sort()
#     if len(input_arr) == k:
#         return input_arr
#     result = []
#     for i in range(k):
#         result.append(input_arr[i])
#     return result


if __name__ == "__main__":
    print(get_least_numbers_solution([4, 5, 1, 6, 2, 7, 3, 8], 4))
