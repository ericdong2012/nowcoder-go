"""
NC52 有效括号序列
https://www.nowcoder.com/practice/37548e94a270412c8b9fb85643c8ccc2

描述
给出一个仅包含字符 '(', ')', '{', '}', '[' 和 ']' 的字符串，判断给出的字符串是否是合法的括号序列
括号必须以正确的顺序关闭，"()" 和 "()[]{}" 都是合法的括号序列，但 "(]" 和 "([)]" 不合法。

要求：空间复杂度 O(n)，时间复杂度 O(n)

示例1
输入："()[]{}"
返回值：true

示例3
输入："([)]"
返回值：false
"""


def is_valid(s):
    stack = []
    for c in s:
        if c == '(':
            # 以下 case 只压栈
            stack.append(')')
        elif c == '[':
            stack.append(']')
        elif c == '{':
            stack.append('}')
        else:
            # s 是个 "]" 等特殊情况
            if len(stack) == 0:
                return False
            # 当前字符跟 stack 中最后一个做比较，相等则弹栈，否则 false
            # last_v = stack[-1]
            # if c == last_v:
            #     stack.pop()
            # else:
            #     return False
            last_v = stack.pop()
            if c != last_v:
                return False

    # 全部弹栈完成，如果为空，说明是符合要求的
    return len(stack) == 0


if __name__ == "__main__":
    # valid = is_valid("()[]{}")
    # valid = is_valid("[()]{}")
    valid = is_valid("[([[)]{}")
    # valid = is_valid("[{{}}]")
    print(valid)
