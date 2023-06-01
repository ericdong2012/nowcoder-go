package _5_优先队列_大小顶堆_

import "sort"

var numArr = []int{}

func insert(num int) {
	numArr = append(numArr, num)
}

func getMedian() float64 {
	length := len(numArr)
	if length == 0 {
		return float64(-1)
	}

	sort.Ints(numArr)
	if len(numArr)%2 == 0 {
		return float64(numArr[length/2]+numArr[length/2-1]) / 2
	} else {
		return float64(numArr[length/2])
	}
}
