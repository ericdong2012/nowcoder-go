package main

import "fmt"

/*

NC79 丑数
https://www.nowcoder.com/practice/6aa9e04fc3794f68acf8778237ba065b?tpId=117&tqId=37779&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=5058&title=

描述
把只包含质因子2、3和5的数称作丑数（Ugly Number）。例如6、8都是丑数，但14不是，因为它包含质因子7。 习惯上我们把1当做是第一个丑数。求按从小到大的顺序的第 n 个丑数。

示例1
输入:
7
返回值:
8

*/

/*
   题目的意思是： 一个数 等于 2x 3y 5z （次方） 相乘的结果
   1, 2, 3, 4, 5, 6, 8, 9, 10, 12      7,11,13,14

   传参是个数，结果是最后的那个值
*/

func GetUglyNumber(index int) int {

	// write code here
	// 最小堆
	/*
					step 1：第一个丑数1加入数组。
					step 2：使用i、j、k三个索引表示该数字有无被乘2、乘3、乘5。
					step 3：后续继续找n−1个丑数，每次取当前丑数索引乘2、乘3、乘5的最小值加入数组，并计数。
					step 4：若是该丑数为相应索引乘上某个数字，则对应的索引往后一位。

		丑数				 1  2  3  4  5  6  8  9
				   索引   0  1  2
				乘2索引   0  1   2
				乘3索引   0  1   2
				乘5索引   0  1   2

					class Solution:
					    def GetUglyNumber_Solution(self , index: int) -> int:
					        #排除0
					        if index == 0:
					            return 0
					        #按顺序记录丑数
					        num = []
					        num.append(1)
					        #记录这是第几个丑数
					        count = 1;
					        #分别代表要乘上2 3 5的下标
					        i = 0
					        j = 0
					        k = 0
					        while count < index:
					            #找到三个数中最小的丑数
					            num.append(min(num[i] * 2, min(num[j] * 3, num[k] * 5)));
					            count += 1
					            #由2与已知丑数相乘得到的丑数，那该下标及之前的在2这里都用不上了
					            if num[count - 1] == num[i] * 2:
					                i += 1
					            #由3与已知丑数相乘得到的丑数，那该下标及之前的在3这里都用不上了
					            if num[count - 1] == num[j] * 3:
					                j += 1
					            #由5与已知丑数相乘得到的丑数，那该下标及之前的在5这里都用不上了
					            if num[count - 1] == num[k] * 5:
					                k += 1
					        return num[count - 1]
	*/

	if index <= 1 {
		return index
	}
	num := make([]int, index)
	num[0] = 1
	//i，j，k分别为三个队列的指针，temp从队列头选出来的最小数
	i := 0
	j := 0
	k := 0
	temp := 0
	for ii := 1; ii < index; ii++ {
		// 第一次是 2、3、5比较, 得到最小的是2
		/*
		 画图

		 ii = 1
		 i  = 1
		 num = [1, 2]

		*/

		//num = append(num, min3(num[i]*2, min3(num[j]*3, num[k]*5)))
		// 最小堆
		// 选出三个队列头最小的数
		A, B, C := num[i]*2, num[j]*3, num[k]*5
		temp = min3(A, min3(B, C))

		// 这三个if有可能进入一个或者多个，进入多个是三个队列头最小的数有多个的情况
		if temp == A {
			i++
		}
		if temp == B {
			j++
		}
		if temp == C {
			k++
		}

		num[ii] = temp
		/*
			num[1] = 2, i=1, j=0, k=0
			num[2] = 3, i=1, j=1, k=0
			num[3] = 4, i=1, j=1, k=1

			num[4] = min(4,6,10) = 4 , i=2, j=1, k=1
		*/
	}

	return num[index-1]
	//return temp
}

func min3(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	solution := GetUglyNumber(7)
	fmt.Println(solution)
}
