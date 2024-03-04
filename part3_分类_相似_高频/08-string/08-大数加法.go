package main

import (
	"fmt"
	"strconv"
)

/*
NC1 大数加法
https://www.nowcoder.com/practice/11ae12e8c6fe48f883cad618c2e81475?tpId=117&tqId=37842&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=579&title=

描述
以字符串的形式读入两个数字，编写一个函数计算它们的和，以字符串形式返回。

字符串仅由'0'~‘9’构成
时间复杂度 O(n)

示例1
输入：
"1","99"
返回值：
"100"
说明：
1+99=100

示例2
输入：
"114514",""
返回值：
"114514"

*/

/*

模拟法或者栈

    设定 i，j 两指针分别指向 s，t 尾部，模拟人工加法；
    计算进位： 计算 carry = tmp // 10，代表当前位相加是否产生进位；
    添加当前位： 计算 tmp = n1 + n2 + carry，并将当前位 tmp % 10 添加至 res 头部；
    索引溢出处理： 当指针 i或j 走过数字首部后，给 n1，n2 赋值为 0，相当于给 s，t 中长度较短的数字前面填 0，以便后续计算。
    当遍历完 s，t 后跳出循环，并根据 carry 值决定是否在头部添加进位 1，最终返回 res 即可。

class Solution:
    def solve(self , s , t ):
        # write code here
        res = ""
        i, j, carry = len(s) - 1, len(t) - 1, 0
        while i >= 0&&j >= 0:
            n1 = int(s[i]) if i >= 0 else 0
            n2 = int(t[j]) if j >= 0 else 0
            tmp = n1 + n2 + carry
            carry = tmp // 10
            res = str(tmp % 10) + res
            i, j = i - 1, j - 1
        return "1" + res if carry else res

*/

func solve1(s string, t string) string {
	// write code here
	res := ""
	i, j := len(s)-1, len(t)-1
	var carry int
	for i >= 0 || j >= 0 {
		var n1, n2 int
		if i >= 0 {
			n1 = int(s[i] - '0')
		} else {
			// 每个数位数不同，位数走完，需要用0填充
			n1 = 0
		}

		if j >= 0 {
			n2 = int(t[j] - '0')
		} else {
			n2 = 0
		}
        // 某一位的两数相加得到的值
		temp := n1 + n2 + carry
		// 超过10才有进位
		carry = temp / 10
		// strconv.Itoa(temp%10)在前是因为 是从个位到十位到百位
		res = strconv.Itoa(temp%10) + res
		j--
		i--
	}

	if carry != 0 {
		return "1" + res
	} else {
		return res
	}
}

func main() {
	s := solve1("1", "99")
	fmt.Println(s)
}
