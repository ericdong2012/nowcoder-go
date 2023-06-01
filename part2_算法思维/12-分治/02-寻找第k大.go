package _2_分治

/*
https://www.nowcoder.com/practice/e016ad9b7f0b45048c58a9f27ba618bf?tpId=196&tqId=37124&ru=/exam/oj


有一个整数数组，请你根据快速排序的思路，找出数组中第 k 大的数。
给定一个整数数组 a ,同时给定它的大小n和要找的 k ，请返回第 k 大的数(包括重复的元素，不用去重)，保证答案存在。

输入：
[1,3,5,2,2],5,3
返回值：
2

*/

/*
参考：
https://www.cnblogs.com/lesroad/p/10621074.html#_label3
https://blog.csdn.net/wjinjie/article/details/117480013

*/

func findKth(a []int, n int, K int) int {
	// write code here
	//if len(a) < K || K == 0 {
	//	return -1
	//}
	//// 构建一下小顶堆
	//heap := heap.Init()//构建一下小顶堆，存储最大的k个数据，最后返回 top对应的数据即可
	////将前k个元素入堆
	//for i := 0; i < K; i++ {
	//	heap.Push(a[i])
	//}
	////如果待加入堆的元素 大于堆中的数据，首先将堆顶元素出堆，再将新元素入堆
	//for i := K; i < len(a); i++ {
	//	if a[i] > heap.top() {
	//		heap.pop()
	//		heap.push(a[i])
	//	}
	//}
	////返回堆顶元素
	//return heap.top()

	mid, i := a[0], 1
	first, tail := 0, n-1
	for first < tail {
		if a[i] < mid {
			a[i], a[tail] = a[tail], a[i]
			tail--
		} else {
			a[i], a[first] = a[first], a[i]
			first++
			i++
		}
	}
	a[first] = mid

	if K-1 == first {
		return mid
	}
	if K-1 < first {
		return findKth(a[:first], first, K)
	}

	return findKth(a[first+1:], n-first-1, K-first-1)
}
