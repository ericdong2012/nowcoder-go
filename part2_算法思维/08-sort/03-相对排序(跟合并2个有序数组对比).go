package main

import "sort"

/*
题⽬⼤意

给你两个数组，arr1 和 arr2，
arr2 中的元素各不相同
arr2 中的每个元素都出现在 arr1 中
对 arr1 中的元素进⾏排序，使 arr1 中项的相对顺序和 arr2 中的相对顺序相同。未在 arr2 中出现过的
元素需要按照升序放在 arr1 的末尾。

提示：
arr1.length, arr2.length <= 1000
0 <= arr1[i], arr2[i] <= 1000
arr2 中的元素 arr2[i] 各不相同
arr2 中的每个元素 arr2[i] 都出现在 arr1 中


解题思路
给出 2 个数组 A 和 B，A 中包含 B 中的所有元素。要求 A 按照 B 的元素顺序排序，B 中没有的元
素再接着排在后⾯，从⼩到⼤排序。
这⼀题有多种解法。⼀种暴⼒的解法就是依照题意，先把 A 中的元素都统计频次放在 map 中，然
后 按照 B 的顺序输出，接着再把剩下的元素排序接在后⾯。还有⼀种桶排序的思想，由于题⽬限
定了 A 的⼤⼩是 1000，这个数量级很⼩，所以可以⽤ 1001 个桶装所有的数，把数都放在桶⾥，
这样默认就已经排好序了。接下来的做法和前⾯暴⼒解法差不多，按照频次输出。B 中以外的元素
就按照桶的顺序依次输出即可

https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1122.Relative-Sort-Array

*/

func relativeSortArray(arr1 []int, arr2 []int) []int {
	// 使用哈希表记录 arr2 中出现过的数字及其对应的下标
	//index := make(map[int]int)
	//for i, v := range arr2 {
	//	index[v] = i
	//}
	//
	//// 定义排序函数
	//sort.Slice(arr1, func(i, j int) bool {
	//	x, y := arr1[i], arr1[j]
	//	// 如果 x 和 y 都出现在了 arr2 中，按照它们在 arr2 中的顺序排序
	//	if ix, ok := index[x]; ok {
	//		if iy, ok := index[y]; ok {
	//			return ix < iy
	//		}
	//		return true
	//	}
	//	// 如果 x 和 y 中只有一个出现在了 arr2 中，则将出现在 arr2 中的放在前面
	//	if _, ok := index[y]; ok {
	//		return false
	//	}
	//	// 如果 x 和 y 都没出现在 arr2 中，直接比较它们的值
	//	return x < y
	//})
	//
	//return arr1

	// 对arr2 做hash
	index := make(map[int]int)
	for i, v := range arr2 {
		index[v] = i
	}
	// 拿到不在arr2中的arr1中数
	notIn := []int{}
	for _, v := range arr1 {
		if _, ok := index[v]; !ok {
			notIn = append(notIn, v)
		}
	}
	//	对arr1中数，计数
	count := make(map[int]int)
	for _, v := range arr2 {
		count[v]++
	}

	// 处理结果
	res := []int{}
	for v, _ := range index {
		temp := []int{}
		for i:= 0 ; i < count[v]; i ++ {
			temp = append(temp, v)
		}
		res = append(res, temp...)
	}

	// 将剩余的排序，并添加进结果中
	sort.Ints(notIn)
	res = append(res, notIn...)

	return res
}

// 解法⼀ 桶排序，时间复杂度 O(n^2)
func relativeSortArray1(A, B []int) []int {
	count := [1001]int{}
	for _, a := range A {
		count[a]++
	}
	res := make([]int, 0, len(A))
	for _, b := range B {
		for count[b] > 0 {
			res = append(res, b)
			count[b]--
		}
	}
	for i := 0; i < 1001; i++ {
		for count[i] > 0 {
			res = append(res, i)
			count[i]--
		}
	}
	return res
}

// 解法⼆ 模拟，时间复杂度 O(n^2)
func relativeSortArray2(arr1 []int, arr2 []int) []int {
	leftover, m, res := []int{}, make(map[int]int), []int{}
	for _, v := range arr1 {
		m[v]++
	}

	for _, s := range arr2 {
		count := m[s]
		for i := 0; i < count; i++ {
			res = append(res, s)
		}
		m[s] = 0
	}
	for v, count := range m {
		for i := 0; i < count; i++ {
			leftover = append(leftover, v)
		}
	}
	sort.Ints(leftover)
	res = append(res, leftover...)
	return res
}
