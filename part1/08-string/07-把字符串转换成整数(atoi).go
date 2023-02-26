package main

import (
	"math"
	"strings"
)

/*
NC100 把字符串转换成整数(atoi)
https://www.nowcoder.com/practice/d11471c3bf2d40f38b66bb12785df47f?tpId=117&tqId=37754&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=579&title=

描述
写一个函数 StrToInt，实现把字符串转换成整数这个功能。不能使用 atoi 或者其他类似的库函数。传入的字符串可能有以下部分组成:
1.若干空格
2.（可选）一个符号字符（'+' 或 '-'）
3. 数字，字母，符号，空格组成的字符串表达式
4. 若干空格

转换算法如下:
1.去掉无用的前导空格
2.第一个非空字符为+或者-号时，作为该整数的正负号，如果没有符号，默认为正数
3.判断整数的有效部分：
	3.1 确定符号位之后，与之后面尽可能多的连续数字组合起来成为有效整数数字，如果没有有效的整数部分，那么直接返回0
	3.2 将字符串前面的整数部分取出，后面可能会存在存在多余的字符(字母，符号，空格等)，这些字符可以被忽略，它们对于函数不应该造成影响
	3.3  整数超过 32 位有符号整数范围 [−231,  231 − 1] ，需要截断这个整数，使其保持在这个范围内。具体来说，小于 −231的整数应该被调整为 −231 ，大于 231 − 1 的整数应该被调整为 231 − 1
4.去掉无用的后导空格

数据范围:
1.0 <=字符串长度<= 100
2.字符串由英文字母（大写和小写）、数字（0-9）、' '、'+'、'-' 和 '.' 组成

*/

func StrToInt(s string) int {
	// write code here
	/*
		既然是将字符串转化为数字，那我们可以遍历字符串，一个字符串，一个字符地检查，然后取出掉无用的，取出数字，利用如下代码，一个数字一个数字地转换，前面的扩大十倍加上后面一位。
		res = res * 10 + sign * (c - '0');

		具体做法:
		step 1：首先越过前导空格，以及越过前导空格后什么都没有(空串)就返回0.
		step 2：然后检查符号，没有符号默认为正数。
		step 3：再在后续遍历的时候，将数字字符转换成字符，遇到非数字则结束转换(题目意思是字符串前面的数字)。 同时与Int型最大最小值比较，检查越界情况。

	*/

	// 去除前导空格和后导空格
	s = strings.Trim(s, " ")
	if len(s) == 0 {
		return 0
	}
	// 符号
	sign := 1
	if s[0] == '+' {
		s = s[1:]
	} else if s[0] == '-' {
		sign = -1
		s = s[1:]
	}

	s1 := []byte(s)
	res := 0
	for _, v := range s1 {
		// 如果是数字字符
		if v >= '0' && v <= '9' {
			//前面的扩大十倍加上后面一位
			res = res*10 + int(v-'0')*sign
			// 边界判断
			if sign == 1 && res > math.MaxInt32 {
				return math.MaxInt32
			}
			if sign == -1 && res < math.MinInt32 {
				return math.MinInt32
			}
		} else {
			// 碰到非数字字符， 直接结束
			break
		}
	}
	return res
}
