package main

import "fmt"

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
var cnt int

func Insert(num int) {
	// 读取数据
	// 排序
	n := len(a)
	cnt++
	// 空的，添加，不用排序，直接返回
	if n == 0 {
		a = append(a, num)
		return
	}
	// 非空, 插入排序
	/*
		插入排序是排序中的一种方式，一旦一个无序数组开始排序，它前面部分就是已经排好的有序数组（一开始长度为0），
		而其后半部分则是需要排序的无序数组，插入排序的做法就是遍历后续需要排序的无序部分，对于每个元素，插入到前半部分有序数组中属于它的位置——即最后一个小于它的元素后

		int i = 0;
		//遍历找到插入点
		for(; i < val.size(); i++){
			if(num <= val.get(i))
				break;
		}
		//插入相应位置
		val.add(i, num);
	*/
	a = append(a, num)
	for i := n; i > 0; i++ {
		// 相当于是比较了最后一位， 如果不加continue, 则是将最小的找出来，同时最后的顺序也是没有问题的，考虑到元素是一个一个添加进来的，也没有问题，只是效率低一点
		if a[i] < a[i-1] {
			a[i], a[i-1] = a[i-1], a[i]
			continue
		}
		break
	}
}

func GetMedian() float64 {
	// 做除法
	if cnt%2 == 0 {
		return float64((a[cnt/2] + a[cnt/2-1]) / 2)
	} else {
		return float64(a[cnt/2])
	}

}

func main() {
	a := []int{5, 2, 3, 4, 1, 6, 7, 0, 8}
	for i := len(a) - 1; i > 0; i-- {
		if a[i] < a[i-1] {
			a[i], a[i-1] = a[i-1], a[i]
			//continue
		}
		//break
	}

	fmt.Printf("%+v", a)
}
