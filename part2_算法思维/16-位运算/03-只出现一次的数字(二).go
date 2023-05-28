package main

import "fmt"

/*
NC227 只出现一次的数字(二)
https://www.nowcoder.com/practice/85cb47fc0c6c483fab7e5cefab54d9e5?tpId=117&tqId=39452&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=5074&title=


描述
给定一个整数数组，数组中有一个数出现了一次，其他数出现了三次，请找出只出现了一次的数。

示例1
输入：
[1]
返回值：
1

示例2
输入：
[1,2,2,2]
返回值：
1

*/

func singleNumber(nums []int) int {
	// write code here
	//temp := nums[0]
	//for i := 1; i < len(nums); i++ {
	//	temp = temp ^ nums[i]
	//	if temp == nums[0] {
	//		break
	//	}
	//}
	//return temp

	/*
		int res,sum;
		for(int i=31;i>=0;--i){
			int t=0;
			for(int j=0;j<nums.size();++j){
		        // 将 nums[j] 右移i位，然后除以2
				t+=(nums[j]>>i)&1;
			}
			t=t%3;
			sum=2*sum+t;
		}
		return sum;
	*/

	arr := make([]int32, 32)
	for _, v := range nums {
		t := int32(v)
		for i := 0; i < 32; i++ {
			arr[i] += (t >> i) & 1
		}
	}

	var ans int32
	for i := 31; i >= 0; i-- {
		if arr[i]%3 == 1 {
			ans |= (1 << i)
		}
	}
	return int(ans)
}

/*
func singleNumber(nums []int) int {
	tempMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		tempMap[nums[i]]++
	}
	for k, v := range tempMap {
		if v == 1 {
			return k
		}
	}

	return -1
}
*/
func main() {
	number := singleNumber([]int{2, 2, 1, 2})
	fmt.Println(number)
}
