package main

import "fmt"

/*
NC79 丑数
https://www.nowcoder.com/practice/6aa9e04fc3794f68acf8778237ba065b?tpId=117&tqId=37779&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=5058&title=

描述
把只包含质因子2、3和5的数称作丑数（Ugly Number）。例如6、8都是丑数，但14不是，因为它包含质因子7。 习惯上我们把1当做是第一个丑数。求按从小到大的顺序的第 n个丑数。

示例1
输入：
7
返回值：
8
*/

/*
2x3y5z
*/

func GetUglyNumber_Solution(index int) int {
	// write code here
	// 最小堆
	// 动态规划
	/*
			step 1：第一个丑数1加入数组。
			step 2：使用i、j、k三个索引表示该数字有无被乘2、乘3、乘5.
			step 3：后续继续找n−1n-1n−1个丑数，每次取当前丑数索引乘2、乘3、乘5的最小值加入数组，并计数。
			step 4：若是该丑数为相应索引乘上某个数字，则对应的索引往后一位。

				 1  2  3  4  5  6
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
	num := make([]int, 1)
	num[0] = 1
	count := 1
	i := 0
	j := 0
	k := 0

	for count < index {
		num = append(num, min3(num[i]*2, min3(num[j]*3, num[k]*5)))
		count++
		// 2与已知丑数相乘得到的丑数，那该下标及之前的在2这里都用不上了
		if num[count-1] == num[i]*2 {
			i++
		}
		// 3与已知丑数相乘得到的丑数，那该下标及之前的在3这里都用不上了
		if num[count-1] == num[j]*3 {
			j++
		}
		// 5与已知丑数相乘得到的丑数，那该下标及之前的在3这里都用不上了
		if num[count-1] == num[k]*5 {
			k++
		}
	}
	return num[count-1]
}

func min3(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	solution := GetUglyNumber_Solution(7)
	fmt.Println(solution)
}
