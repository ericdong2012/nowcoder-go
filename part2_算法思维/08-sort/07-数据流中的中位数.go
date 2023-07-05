package main

/*
NC131 数据流中的中位数
https://www.nowcoder.com/practice/9be0172896bd43948f8a32fb954e1be1?tpId=117&tqId=37807&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=581,582,586&title=


描述
如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，那么中位数就是所有数值排序之后位于中间的数值。如果从数据流中读出偶数个数值，那么中位数就是所有数值排序之后中间两个数的平均值。我们使用Insert()方法读取数据流，使用GetMedian()方法获取当前读取数据的中位数。
进阶： 空间复杂度 O(n)  ， 时间复杂度 O(nlogn)


示例1
输入：
[5,2,3,4,1,6,7,0,8]
返回值：
"5.00 3.50 3.00 3.50 3.00 3.50 4.00 3.50 4.00 "
说明：
数据流里面不断吐出的是5,2,3...,则得到的平均数分别为5,(5+2)/2,3...


示例2
输入：
[1,1,1]=
返回值：
"1.00 1.00 1.00 "

*/

/*

本题是对顶堆算法的一个经典题，所谓对顶堆算法，实际上就是使用一个小顶堆和一个大顶堆
对于求动态中位数的问题，我们可以利用堆的性质，可以在logn的时间复杂度取出最大值或最小值，那么我们可以用大顶堆维护前一半的数，用小顶堆维护后面一半的数
我们规定，若元素总数为奇数，那么大顶堆的元素个数要比小顶对的个数多1，这样，当总数为奇数时，中位数就是大顶堆的堆顶元素；当总数为偶数时，中位数就大顶堆堆顶和小顶堆堆顶元素的平均值

该题 heap&hash&stack 中的sort 13-数据流中的中位数.go

*/

var a []int
var cnt int

func Insert(num int) {
	aLen := len(a)
	cnt++
	if aLen == 0 {
		a = append(a, num)
		return
	}

	a = append(a, num)
	for i := aLen; i > 0; i-- {
		if a[i] < a[i-1] {
			a[i], a[i-1] = a[i-1], a[i]
			continue
		}
		break
	}
}
func GetMedian() float64 {
	if cnt%2 == 0 {
		return float64(a[cnt/2]+a[cnt/2-1]) / 2.0
	} else {
		return float64(a[cnt/2])
	}
}