package main

import "fmt"

/*
NC118 数组中的逆序对
https://www.nowcoder.com/practice/96bd6684e04a44eb80e6a68efc0ec6c5?tpId=117&tqId=37763&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=590,5058,578&title=


描述
在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组,求出这个数组中的逆序对的总数P。并将P对1000000007取模的结果输出。 即输出P mod 1000000007
输入描述：题目保证输入的数组中没有的相同的数字

示例1
输入：
[1,2,3,4,5,6,7,0]
返回值：
7    [1, 0] / [2, 0] / [3, 0] / [4, 0] / [5, 0] / [6, 0] /[7, 0]

示例2
输入：
[1,2,3]
返回值：
0

*/

func InversePairs(data []int) int {
	// 归并排序
	/*
		递归划分整个区间为基本相等的左右两个区间
		合并两个有序区间
	*/
	temp := make([]int, len(data))
	return mergeSort(data, temp)
}

func mergeSort(nums []int, temp []int) int {
	if len(nums) < 2 {
		return 0
	}
	mid := len(nums) / 2
	left := nums[:mid]
	right := nums[mid:]
	cnt := mergeSort(left, temp[:mid]) % 1000000007
	cnt += mergeSort(right, temp[mid:]) % 1000000007
	temp = temp[:0]
	i, j := 0, 0
	lr, rr := len(left), len(right)
	for i < lr && j < rr {
		if left[i] <= right[j] {
			temp = append(temp, left[i])
			i++
		} else {
			temp = append(temp, right[j])
			cnt += (lr - i) % 1000000007
			j++
		}
	}
	if i == lr {
		for ; j < rr; j++ {
			temp = append(temp, right[j])
		}
		//tmp = append(tmp,nums[j:]...) // 不推荐使用，效率很低
	} else {
		for ; i < lr; i++ {
			temp = append(temp, left[i])
		}
		//tmp = append(tmp,nums[i:]...) //// 不推荐使用，效率很低
	}

	//for i := start; i < end; i++ {
	//  nums[i] = tmp[i - start]
	//}
	copy(nums, temp)
	return cnt % 1000000007
}

func main() {
	pairs := InversePairs([]int{1, 2, 3, 4, 5, 6, 7, 0})
	fmt.Println(pairs)
}
