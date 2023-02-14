package main

import "fmt"

/*
NC75 数组中只出现一次的两个数字
https://www.nowcoder.com/practice/389fc1c3d3be4479a154f63f495abff8?tpId=117&tqId=37773&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=5074&title=


描述
一个整型数组里除了两个数字只出现一次，其他的数字都出现了两次。请写程序找出这两个只出现一次的数字。

数据范围：数组长度 2\le n \le 10002≤n≤1000，数组中每个数的大小 0 < val \le 10000000<val≤1000000
要求：空间复杂度 O(1)，时间复杂度 O(n)
提示：输出时按非降序排列。

示例1
输入：
[1,4,1,6]
返回值：
[4,6]

说明：
返回的结果中较小的数排在前面

示例2
输入：
[1,2,3,3,2,9]
返回值：
[1,9]

*/

func FindNumsAppearOnce(array []int) []int {
	// write code here

	// 异或操作，相同为0，不同为1
	/*	先把所有的值进行异或，得到的结果就是两个不同的值异或的结果
		找到这个值中有1的位，例如6(110),4(100),异或010
		第二位为1(异或结果 010 )， 根据这个位置来对所有的值进行分组，相同的值肯定分到同一组，不同的值不在同一组
		6和4必定在不同组,相同的元素必在同一组
	*/

	//先把所有的值进行异或，得到的结果就是两个不同的值异或的结果
	ans := 0
	for _, v := range array {
		ans ^= v
	}

	// 拿到最后一个标识(完整的 -  高位的)
	flag := ans - (ans & (ans - 1))
	// 通过下面的操作将两个数据分开
	a, b := 0, 0
	for _, v := range array {
		// flag = 2
		if (flag & v) != 0 {
			// 0 ^ 6
			a ^= v
		} else {
			// 0 ^ 1 ^ 4 ^ 1
			b ^= v
		}
	}
	if a > b {
		return []int{b, a}
	}
	return []int{a, b}
}
func main() {
	once := FindNumsAppearOnce([]int{1, 4, 1, 6})
	fmt.Println(once)
}
