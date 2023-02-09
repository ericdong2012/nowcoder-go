package main

import "fmt"

/*
BM87 合并两个有序的数组
https://www.nowcoder.com/practice/89865d4375634fc484f3a24b7fe65665?tpId=295&tqId=658&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295

描述
给出一个有序的整数数组 A 和有序的整数数组 B ，请将数组 B 合并到数组 A 中，变成一个有序的升序数组


注意：
1.保证 A 数组有足够的空间存放 B 数组的元素， A 和 B 中初始的元素数目分别为 m 和 n，A的数组空间大小为 m+n
2.不要返回合并的数组，将数组 B 的数据合并到 A 里面就好了，且后台会自动将合并后的数组 A 的内容打印出来，所以也不需要自己打印
3. A 数组在[0,m-1]的范围也是有序的
示例1
输入：
[4,5,6],[1,2,3]
复制
返回值：
[1,2,3,4,5,6]
复制
说明：
A数组为[4,5,6]，B数组为[1,2,3]，后台程序会预先将A扩容为[4,5,6,0,0,0]，B还是为[1,2,3]，m=3，n=3，传入到函数merge里面，然后请同学完成merge函数，将B的数据合并A里面，最后后台程序输出A数组
示例2
输入：
[1,2,3],[2,5,6]
复制
返回值：
[1,2,2,3,5,6]

*/

func merge(A []int, m int, B []int, n int) {
	// write code here
	// 实际运行会报错
	index := m + n - 1
	for m > 0 && n > 0 {
		// 先比较最后一位
		if A[m-1] >= B[n-1] {
			A[index] = A[m-1]
			m--
		} else {
			A[index] = B[n-1]
			n--
		}
		index--
	}
	//B数组依然有剩余
	for n > 0 {
		A[index] = B[n-1]
		n--
		index--
	}

	fmt.Println(A)

	//if m==0 || n == 0{
	//	copy(A,B)
	//}else {
	//	for k:=m+n-1;k>=0;k--{
	//		if A[m-1] > B[n-1] {
	//			A[k]=A[m-1]
	//			m--
	//			if m==0 {
	//				copy(A,B[:n])
	//				break
	//			}
	//		} else {
	//			A[k]=B[n-1]
	//			n--
	//			if n==0 {
	//				break
	//			}
	//		}
	//	}
	//}
}

func main() {
	merge([]int{1, 2, 3}, 3, []int{2, 5, 6}, 3)
}
