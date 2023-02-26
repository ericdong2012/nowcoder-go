package main

/*
NC131 数据流中的中位数
https://www.nowcoder.com/practice/9be0172896bd43948f8a32fb954e1be1?tpId=117&tqId=37807&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=590&title=

描述
如何得到一个数据流中的中位数？
如果从数据流中读出奇数个数值，那么中位数就是所有数值排序之后位于中间的数值。
如果从数据流中读出偶数个数值，那么中位数就是所有数值排序之后中间两个数的平均值。
我们使用Insert()方法读取数据流，使用GetMedian()方法获取当前读取数据的中位数。

示例1
输入：
[5,2,3,4,1,6,7,0,8]
返回值：
"5.00 3.50 3.00 3.50 3.00 3.50 4.00 3.50 4.00"
说明：
数据流里面不断吐出的是5,2,3...,  则得到的平均数分别为5,(5+2)/2,3...


示例2
输入：
[1,1,1]
返回值：
"1.00 1.00 1.00"

*/

// 数组
var a []int
// 读取的数据个数
var count int

func Insert(num int) {
	n := len(a)
	count++
	// 空的， 不用排序，添加, 直接返回
	if n == 0 {
		a = append(a, num)
		return
	}
	// 非空(之前的坑定已经排好序了) 先插入, 后排序
	a = append(a, num)
	// 4，5，3
	// 从n开始加是因为n是之前的长度，对应现在的最后一位
	for i := n; i > 0; i-- {
		//  如果当前位小于前一位 交换
		if a[i] < a[i-1] {
			a[i], a[i-1] = a[i-1], a[i]
			//continue
		} else {
			// 找到了自己的位置 a[i] >= a[i-1]
			break
		}
	}
}

func GetMedian() float64 {
	// 做除法
	if count%2 == 0 {
		return float64((a[count/2] + a[count/2-1]) / 2)
	} else {
		return float64(a[count/2])
	}

}

func main() {
	//a := []int{5, 2, 3, 4, 1, 6, 7, 0, 8}
	//for i := len(a) - 1; i > 0; i-- {
	//	if a[i] < a[i-1] {
	//		a[i], a[i-1] = a[i-1], a[i]
	//		//continue
	//	}
	//	//break
	//}
	//
	//fmt.Printf("%+v", a)

}
