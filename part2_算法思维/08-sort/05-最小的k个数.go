package main

/*
https://www.nowcoder.com/practice/6a296eb82cf844ca8539b57c23e6e9bf

给定一个长度为 n 的可能有重复值的数组，找出其中不去重的最小的 k 个数。例如数组元素是4,5,1,6,2,7,3,8这8个数字，则最小的4个数字是1,2,3,4(任意顺序皆可)。


示例1
输入：
[4,5,1,6,2,7,3,8],4
返回值：
[1,2,3,4]
说明：
返回最小的4个数即可，返回[1,3,2,4]也可以


示例2
输入：
[1],0
返回值：
[]


示例3
输入：
[0,1,2,1,2],3
返回值：
[0,1,1]

*/


func GetLeastNumbers_Solution( input []int ,  k int ) []int {
	if len(input) < k && k == 0 {
		return []int{}
	}
	for i:= 0; i < k ; i++{
		for k:=len(input) -1 ; k> i ; k-- {
			if input[k] < input[k-1] {
				input[k] , input[k-1] = input[k-1], input[k]
			}
		}
	}
	return input[:k]
}

//func GetLeastNumbers_Solution( input []int ,  k int ) []int {
//	// write code here
//	if len(input) < k{
//		return []int{}
//	}
//	sort.Ints(input)
//	if len(input) == k{
//		return input
//	}
//	result := make([]int, 0)
//	for i:=0; i<k; i++ {
//		result = append(result, input[i])
//	}
//	return  result
//}