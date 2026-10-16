"""
NC142 最长重复子串
https://www.nowcoder.com/practice/4fe306a84f084c249e4afad5edf889cc

描述
定义重复字符串是由两个相同的字符串首尾拼接而成，例如 abcabc 便是长度为 6 的一个重复字符串。
给定一个字符串，请返回其最长重复子串的长度。
若不存在任何重复字符子串，则返回 0

备注：字符串长度不超过 10000，且仅由小写字母组成

示例1
输入："ababc"
返回值：4
说明：abab 为最长的重复字符子串，长度为 4

示例2
输入："abcab"
返回值：0
"""

# 可以将两个字符串想像成两个连续的滑动窗口，并假设这个滑动窗口最大是字符串长度的一半，
# 通过比较两个窗口的内容是否相同，不相同的时候不断从左向右平移，完了之后还是不相同，
# 这时候就将窗口的大小调小一点，直到找到一个相同的，这个时候窗口的长度 ×2 就是最大字符串的大小

"""
public int solve (String a) {
    // write code here
    if (a == null || a.length() <= 1) return 0;
    char[] chars = a.toCharArray();
    int len = chars.length;
    int maxLen = chars.length / 2;
    for (int i = maxLen; i >= 1; --i){
        for (int j = 0; j <= len - 2 * i; ++j){
            if (check(chars, j, i))
                return 2 * i;
        }
    }
    return 0;
}

public boolean check(char[] chars, int start, int len){
    for (int i = start; i < start + len; ++i){
        if (chars[i] != chars[i + len])
              return false;
    }
    return true;
}
"""

# 滑动窗口
# 可以将两个字符串想像成两个连续的滑动窗口，并假设这个滑动窗口最大是字符串长度的一半，
# 通过比较两个窗口的内容是否相同，不相同的时候不断从左向右平移，完了之后还是不相同，
# 这时候就将窗口的大小调小一点，直到找到一个相同的，这个时候窗口的长度 ×2 就是最大字符串的大小
def solve2(a):
    n = len(a)
    if n < 2:
        # 长度小于 2 不存在重复子串，返回 0（与题目说明及文中 Java 参考实现一致）
        return 0
    res = 0
    # 从中间往前走(中间找不到，再往前走缩小窗口)
    # i 不能等于 0，否则后面 a[j] == a[j+i] 一直成立
    for i in range(n // 2, 0, -1):
        # 从前往后走
        for j in range(n - i):
            # 比如: abab，如果 j 和 j+i 相等，结果加 1
            # j + i 自动带上的滑动窗口 妙
            if a[j] == a[j + i]:
                res += 1
            else:
                # 比如: abac，如果 j 和 j+i 不等，结果要置 0 (当 i=2, j=1，此时 b 不等于 c)
                res = 0
            # 如果长度和索引相等，说明有重复的，而此时 i = n/2，res 已经是最大的，res * 2 直接返回
            # (比如: abab，i=2, j=1 发现 res == i)
            if res == i:
                return res * 2

    return res


if __name__ == "__main__":
    i = solve2("ababc")
    print(i)
