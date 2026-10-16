"""
最长的括号子串
https://www.nowcoder.com/practice/45fd68024a4c4e97a8d6c45fc61dc6ad

描述
给出一个长度为 n 的，仅包含字符 '(' 和 ')' 的字符串，计算最长的格式正确的括号子串的长度。

示例1
输入："(()"
返回值：2

示例2
输入："(())"
返回值：4
"""


def longest_valid_parentheses(s):
    """write code here"""
    # stack = []
    # count = 0
    # for i in range(len(s)):
    #     if s[i] == '(':
    #         stack.append(s[i])
    #     else:
    #         if len(stack) != 0:
    #             stack.pop()
    #             count += 2
    # return count

    # if len(s) == 0:
    #     return 0
    # stack = [0] * (len(s) + 1)
    # top = 0
    # stack[top] = -1
    # res = 0
    # for i in range(len(s)):
    #     if s[i] == '(':
    #         top += 1
    #         stack[top] = i
    #     else:
    #         top -= 1
    #         if top == -1:
    #             top += 1
    #             stack[top] = i
    #         else:
    #             res = max(res, i - stack[top])
    # return res

    if len(s) <= 1:
        return 0
    # 用来存索引，需要有个元素，不然会报错，比如: )....
    stack = [-1]
    result = 0
    for i in range(len(s)):
        if s[i] == ')':
            # 先弹栈
            stack.pop()
            # 有元素
            if len(stack) != 0:
                result = max(result, i - stack[-1])
                # result += 2
            else:
                # 没有元素
                stack.append(i)
        else:
            stack.append(i)

    return result


if __name__ == "__main__":
    parentheses = longest_valid_parentheses("()()")
    print(parentheses)
