package main

/*
题目大意 #
给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。你可以假设除了整数 0 之外，这个整数不会以零开头。

解题思路 #
给出一个数组，代表一个十进制数，数组的 0 下标是十进制数的高位。要求计算这个十进制数加一以后的结果。
简单的模拟题。从数组尾部开始往前扫，逐位进位即可。最高位如果还有进位需要在数组里面第 0 位再插入一个 1 。

*/

func plusOne(digits []int) []int {
	// 最后一位不为9
	if digits[len(digits)-1] != 9 {
		digits[len(digits)-1] += 1
		return digits
	}

	// 最后一位为9, 将最后一位置为0， 判断其他位
	digits[len(digits)-1] = 0
	// 倒序，从倒数第二位开始遍历
	for i := len(digits) - 2; i >= 0; i-- {
		digits[i]++
		// 倒数第二位加1后不是10
		if digits[i] != 10 {
			return digits
		}
		// 倒数第二位加1后是10，继续遍历
		digits[i] = 0
	}
	// 全是9, 其他位已经在上面处理过，变成0, 在最前面添加1
	// 第一位变成1， 并在后面加上0
	digits[0] = 1
	digits = append(digits, 0)
	return digits
}
