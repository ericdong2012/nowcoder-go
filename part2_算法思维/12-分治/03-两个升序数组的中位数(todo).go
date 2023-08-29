package _2_分治

/*

https://www.nowcoder.com/practice/b3b59248e61f499482eaba636305474b?tpId=196&tqId=40563&ru=/exam/oj

给定两个长度为 n 和 m 的升序数组（后一个数一定大于等于前一个数），请你找到这两个数组中全部元素的中位数。


[1,2,3,4,5],[6,7,8,9] => 5

[1,2,3,8,9],[4,5,6,7] => 5

*/

func Median(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	r1, r2 := 0, 0
	cnt, p1, p2 := 0, 0, 0
	// [1,2,3,4] [6,7,8,9]
	// [1,2,3,4,5] [6,7,8,9]
	for cnt <= (l1+l2)/2 {
		// 比右侧小前， 左侧一直往前走(p1) r1, r2一直在变化
		if p1 < l1 && p2 < l2 && nums1[p1] < nums2[p2] || p2 == l2 {
			// r1: 0   r2: 1   p1: 1
			// r1: 1   r2: 2   p1: 2
			// r1: 2   r2: 3   p1: 3
			// r1: 3   r2: 4   p1: 4    cnt: 4

			// 左侧有多的 比如 [1,2,3,4,5]
			// r1: 4   r2: 5   p1: 5    cnt: 5
			r1, r2 = r2, nums1[p1]
			p1++
		} else {
			// r1: 4  r2: 6   p2: 1  p1:4  cnt: 5
			r1, r2 = r2, nums2[p2]
			p2++
		}
		cnt++
	}

	// r1: 4  r2: 6   p2: 1  p1:4  cnt: 5
	if (l1+l2)%2 == 0 {
		return float64(r1+r2) / 2
	}

	return float64(r2)
}

//func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//	n, m := len(nums1), len(nums2)
//	left, right := (n+m+1)/2, (n+m+2)/2
//	return (float64(getKth(nums1, nums2, left)) + float64(getKth(nums1, nums2, right))) / 2
//}
//
//func getKth(nums1, nums2 []int, k int) int {
//	n, m := len(nums1), len(nums2)
//	if n > m {
//		return getKth(nums2, nums1, k)
//	}
//	if n == 0 {
//		return nums2[k-1]
//	}
//	if k == 1 {
//		return min(nums1[0], nums2[0])
//	}
//	i, j := min(n, k/2), min(m, k/2)
//	if nums1[i-1] > nums2[j-1] {
//		return getKth(nums1, nums2[j:], k-j)
//	} else {
//		return getKth(nums1[i:], nums2, k-i)
//	}
//}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
