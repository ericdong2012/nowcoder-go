"""
NC49 最长的括号子串
https://www.nowcoder.com/practice/45fd68024a4c4e97a8d6c45fc61dc6ad

描述
给出一个长度为 n 的，仅包含字符 '(' 和 ')' 的字符串，计算最长的格式正确的括号子串的长度。

例1: 对于字符串 "(()" 来说，最长的格式正确的子串是 "()"，长度为 2。
例2: 对于字符串 ")()())" 来说，最长的格式正确的子串是 "()()"，长度为 4。

要求时间复杂度 O(n)，空间复杂度 O(n)。

示例1
输入："(()"
返回值：2

示例2
输入："(())"
返回值：4
"""

# 栈或者动态规划
#
# 栈
# step 1：可以使用栈来记录左括号下标。
# step 2：遍历字符串，左括号入栈，每次遇到右括号则弹出左括号的下标。
# step 3：然后长度则更新为当前下标与栈顶下标的距离。
# step 4：遇到不符合的括号，可能会使栈为空，因此需要使用 start 记录上一次结束的位置，
#         这样用当前下标减去 start 即可获取长度，即得到子串。
# step 5：循环中最后维护子串长度最大值。


def longest_valid_parentheses(s):
    # 记录索引
    stack = [-1]
    result = 0
    for i in range(len(s)):
        if s[i] == ')':
            # 先弹栈
            stack.pop()
            # 有元素
            if len(stack) != 0:
                result = max(result, i - stack[-1])
            else:
                # 没有元素
                stack.append(i)
        else:
            stack.append(i)

    return result


if __name__ == "__main__":
    # parentheses = longest_valid_parentheses(")(()")
    parentheses = longest_valid_parentheses("()()")
    print(parentheses)
