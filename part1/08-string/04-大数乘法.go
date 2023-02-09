package main

import (
	"fmt"
)

/*
NC10 大数乘法
https://www.nowcoder.com/practice/c4c488d4d40d4c4e9824c3650f7d5571?tpId=117&tqId=37843&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=579&title=

描述
以字符串的形式读入两个数字，编写一个函数计算它们的乘积，以字符串形式返回。
数据范围： 读入的数字大小满足 0 \le n \le 10^{1000}0≤n≤101000
要求：空间复杂度 O(n)，时间复杂度 O(n^2)

示例1
输入：
"11","99"
返回值：
"1089"

说明：
11*99=1089

示例2
输入：
"1","0"
返回值：
"0"

*/

/*
剑指offer题目

遍历法
	将乘法转换成加法

状态机
*/

func solve3(s string, t string) string {
	/*
		将两个字符串转化为整型数组形式
		逐位相乘, 中间 i * j 会有重复出现, 如89 * 12, 会有res = [0, 8*1, 8*2+9*1, 9*2]
		从最低位, 也就是数组res最右边元素开始处理进位
		接着从最高位开始收集结果即可
	*/
	numS, numT := []byte(s), []byte(t)
	tempRes := make([]int, len(s)+len(t))
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(t); j++ {
			tempRes[i+j+1] += int(numS[i]-'0') * int(numT[j]-'0')
		}
	}
	// 处理进位
	for i := len(tempRes) - 1; i > 0; i-- {
		if tempRes[i] >= 10 {
			tempRes[i-1] += tempRes[i] / 10
			tempRes[i] %= 10
		}
	}
	// 去除前导0, 有些极端情况
	k := 0
	for ; k < len(tempRes); k++ {
		if tempRes[k] != 0 {
			break
		}
	}
	// 全是0
	if k == len(tempRes) {
		return "0"
	}
	tempRes = tempRes[k:]

	// 构建字符串
	res := make([]rune, len(tempRes))
	for i, n := range tempRes {
		res[i] = rune(n + '0')
	}
	return string(res)

}

func main() {
	s := solve3("1", "1")
	fmt.Println(s)
}
