package main

/*
NC144 懂二进制
https://www.nowcoder.com/practice/120e406db3fd46f09d55d59093f13dd8?tpId=117&tqId=37855&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=5074&title=


描述
世界上有10种人，一种懂二进制，一种不懂。那么你知道两个int32整数m和n的二进制表达，有多少个位(bit)不同么？

示例1
输入：
3,5
返回值：
2

说明：
3的二进制为11， 5的二进制为101，总共有2位不同

示例2
输入：
1999,2299
返回值：
7

*/

func countBitDiff(m int, n int) int {
	// write code here
	/*
		先异或，得到所有不同位 （r = m^n）
		数r有多少位是1即可

		class Solution:
		    def countBitDiff(self , m: int, n: int) -> int:
		        res=0;
		        r = m^n
		        while r:
		            r &= (r-1)
		            res += 1
		        return res


	*/
	res := 0
	temp := m ^ n
	for temp > 0 {
		res++
		// 最后一位1
		temp &= (temp - 1)
	}
	return res
}
