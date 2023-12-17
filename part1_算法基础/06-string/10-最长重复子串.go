package main

import "fmt"

/*
NC142 最长重复子串
https://www.nowcoder.com/practice/4fe306a84f084c249e4afad5edf889cc?tpId=117&tqId=37853&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=579&title=

描述
定义重复字符串是由两个相同的字符串首尾拼接而成， 例如 abcabc 便是长度为6的一个重复字符串。
给定一个字符串，请返回其最长重复子串的长度。
若不存在任何重复字符子串，则返回0

备注： 字符串长度不超过10000，且仅由小写字母组成

示例1
输入：
"ababc"
返回值：
4
说明：
abab为最长的重复字符子串，长度为4


示例2
输入：
"abcab"
返回值：
0
说明：
该字符串没有重复字符子串

*/

/*

public int solve (String a) {
    // write code here
    if (a == null || a.length() <= 1) return 0;
    char[] chars = a.toCharArray();
    int len = chars.length;
    int maxLen = chars.length / 2;
    for (int i = maxLen; i >= 1;--i){
        for (int j = 0; j <= len - 2 * i;++j){
            if (check(chars, j, i))
                return 2 * i;
        }
    }
    return 0;
}

public boolean check(char[] chars, int start, int len){
    for (int i = start;i < start + len;++i){
        if (chars[i] != chars[i +len])
              return false;
    }
    return true;
}

*/

// 二分思想
// 可以将两个字符串想像成两个连续的滑动窗口，并假设这个滑动窗口最大是字符串长度的一半，通过比较两个窗口的内容是否相同
// 不相同的时候不断从左向右平移，完了之后，还是不相同，这时候就将窗口的大小调小一点，直到找到一个相同的，这个时候窗口的长度×2 就是最大字符串的大小
func solve2(a string) int {
	n := len(a)
	if n < 2 {
		return n
	}
	res := 0
	// 从中间往前走(中间找不到，再往前走缩小窗口)
	// i 不能等于0， 否则后面 a[j] == a[j+i]  一直成立
	for i := n / 2; i > 0; i-- {
		// 从前往后走, 最多到 n - i
		for j := 0; j < n-i; j++ {
			// 比如: ababa,  如果 j 和 j+i 相等, 结果加1
			if a[j] == a[j+i] {
				res++
			} else {
				// 如果 j 和 j+i 不等，结果要置0 (比如: abac, 当 i= 2, j= 1, 此时 b 不等于 c)
				res = 0
			}
			// 如果长度和索引相等， 说明有重复的， 而此时 i = n /2, res已经是最大的, res * 2 直接返回 （ 比如: abab , i=2, j= 1  发现res == i ）
			if res == i {
				return res * 2
			}
		}
	}
	return res
}

func main() {
	i := solve2("ababc")
	fmt.Println(i)
}
