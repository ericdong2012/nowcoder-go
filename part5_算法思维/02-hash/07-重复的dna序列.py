"""
https://www.nowcoder.com/practice/fe9099e5308042a8af2f7aabdb3719fe

描述
所有的 DNA 序列都是由 'A', 'C', 'G', 'T' 字符串组成的，例如 'ACTGGGC'
请你实现一个函数找出所有的目标子串，目标子串的定义是：长度等于 10，且在 DNA 序列中出现次数超过 1 次的子串
（允许两个子串有重合的部分）。
（注：返回的所有目标子串的顺序必须与原 DNA 序列的顺序一致）

保证序列中只出现 'A', 'C', 'G', 'T'

示例1
输入："AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
返回值：["AAAAACCCCC","CCCCCAAAAA"]

示例2
输入："AAAAAAAAAAA"
返回值：["AAAAAAAAAA"]
"""


def repeated_dna(dna):
    """write code here"""
    # 用哈希表 tmp 储存子串和位置信息，遍历一遍字符串，统计每个子串出现的次数
    # 然后为了保证重复子串的顺序，再遍历一遍字符串，对于出现次数大于 1 的子串保存到 res 中，
    # 同时在哈希表中删除这一子串的次数信息，防止重复

    # 储存长度为 10 的子串和出现的次数
    temp = {}
    for i in range(len(dna) - 10 + 1):
        # 长度为 10
        s = dna[i:i + 10]
        temp[s] = temp.get(s, 0) + 1
    res = []
    # 本来可以这样写
    # for k, v in temp.items():
    #     if v > 1:
    #         res.append(k)
    # 但是题目要求，在原 DNA 序列中，"AAAAACCCCC" 要比 "CCCCCAAAAA" 先出现，
    # 所以再遍历原有的字符串，而不是直接遍历 hash
    for i in range(len(dna) - 10 + 1):
        s = dna[i:i + 10]
        if temp.get(s, 0) > 1:
            res.append(s)
            # 防止特例，只有一个字母
            del temp[s]

    return res


if __name__ == "__main__":
    dna = repeated_dna("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT")
    print(dna)
