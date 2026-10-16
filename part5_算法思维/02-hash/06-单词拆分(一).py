"""
https://www.nowcoder.com/practice/c0d32c1ce5744472a01b2351a2c2767f

描述
给定一个字符串 s 和一个字符串数组 dic，在字符串的任意位置拆分任意次后得到的字符串集合
是否是给定字符串数组的子集。（即拆分后的字符串集合中的所有字符串都在 dic 数组中）

示例1
输入："nowcoder", ["now","coder","no","wcoder"]
返回值：true

示例4
输入："nownowcoder", ["now", "coder"]
返回值：true

说明：你可以重复使用 dic 数组中的字符串
"""


def word_div1(s, dic):
    # 对标的 dic 建立 hash 表，并初始化
    m = {v: True for v in dic}
    stack = [False] * (len(s) + 1)
    stack[0] = True
    # 滑动窗口
    for i in range(1, len(s) + 1):
        for j in range(i):
            # 调试看看
            # s[j:i] 这个窗口中的字符串在 m 中，也就是 dic 中，
            # 并且 stack[j] 存在，当 i 走到 3，stack[0]==true，m["now"] ==> stack[2] = true
            # 并且 stack[j] 存在，当 i 走到 10，stack[3]==true，m["coder"] ==> stack[10] = true
            if stack[j] and m.get(s[j:i], False):
                stack[i] = True
                # 优化时间
                break

    return stack[len(s)]


if __name__ == "__main__":
    # div = word_div1("nownowcoder", ["now", "coder"])
    div = word_div1("nownowcoder", ["n", "now", "owcoder", "coder"])
    print(div)
