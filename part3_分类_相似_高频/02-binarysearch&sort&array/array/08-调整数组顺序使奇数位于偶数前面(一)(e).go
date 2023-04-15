package main

/*
NC77 调整数组顺序使奇数位于偶数前面(一)
https://www.nowcoder.com/practice/ef1f53ef31ca408cada5093c8780f44b?tpId=117&tqId=37776&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=590,5058,578&title=


描述
输入一个长度为 n 整数数组，实现一个函数来调整该数组中数字的顺序，使得所有的奇数位于数组的前面部分，所有的偶数位于数组的后面部分，并保证奇数和奇数，偶数和偶数之间的相对位置不变。

示例1
输入：
[1,2,3,4]
返回值：
[1,3,2,4]

示例2
输入：
[2,4,6,5,7]
返回值：
[5,7,2,4,6]

*/
func reOrderArray(array []int) []int {
	// write code here
	/*
		class Solution:
		    def reOrderArray(self , array ):
		        # write code here
		        odd = []
		        even = []
		        for i in range(len(array)):
		            if array[i] % 2 ==0:
		                # 偶数存储
		                even.append(array[i])
		            else:
		                # 奇数存储
		                odd.append(array[i])
		        array.clear()
		        # 组合两个数组
		        array = odd+even
		        return array

	*/

	odd := []int{}
	even := []int{}
	for i := 0; i < len(array); i++ {
		if array[i]%2 == 0 {
			odd = append(odd, array[i])
		} else {
			even = append(even, array[i])
		}
	}

	even = append(even, odd...)
	return even

}
