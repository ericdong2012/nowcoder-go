package main

/*
NC36 在两个长度相等的排序数组中找到上中位数
https://www.nowcoder.com/practice/6fbe70f3a51d44fa9395cfc49694404f?tpId=117&tqId=37808&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=5058&title=

描述
给定两个递增数组arr1和arr2，已知两个数组的长度都为N，求两个数组中所有数的上中位数。
上中位数：假设递增序列长度为n，为第n/2个数

要求：时间复杂度 O(n)，空间复杂度 O(1)
进阶：时间复杂度为O(logN)，空间复杂度为O(1)

示例1
输入：
[1,2,3,4], [3,4,5,6]
返回：
3
说明：
总共有8个数，上中位数是第4小的数，所以返回3。

*/
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findMedianinTwoSortedAray(arr1 []int, arr2 []int) int {
	// write code here
	// 方式一： 两个合并， 排序， 算长度，算中位数是第几个，直接取值
	/*
		    temp := []int{}
			temp = append(temp, arr1...)
			temp = append(temp, arr2...)
			sort.Ints(temp)

			return temp[len(temp) / 2 -1]
	*/

	// 方式二：
	if len(arr1) == 1 {
		return Min(arr1[0], arr2[0])
	}

	l1 := 0
	r1 := len(arr1) - 1
	l2 := 0
	r2 := len(arr2) - 1

	for l1 < r1 {
		pos := (r1 - l1) % 2
		mid1 := (l1 + r1) / 2
		mid2 := (l2 + r2) / 2
		if arr1[mid1] > arr2[mid2] {
			l2 = mid2 + pos
			r1 = mid1
		} else if arr1[mid1] < arr2[mid2] {
			l1 = mid1 + pos
			r2 = mid2
		} else {
			return arr1[mid1]
		}
	}
	return Min(arr1[l1], arr2[l2])
}
