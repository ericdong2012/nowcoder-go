package main

import "fmt"

/*
NC213 除自身以外数组的乘积
https://www.nowcoder.com/practice/0786aa81c1c64c2a990e393fac811b45?tpId=117&tqId=39417&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=5059&title=


描述
给定一个长度为 n 的数组 nums ，返回一个数组 res，res[i]是nums数组中除了nums[i]本身以外其余所有元素的乘积，即：res[i] = nums[1] \times nums[2] \times ......\times nums[i-1] \times nums[i+1] ...... \times nums[n] \res[i]=nums[1]×nums[2]×......×nums[i−1]×nums[i+1]......×nums[n]

1.请不要使用除法，并且在 O(n) 时间复杂度内完成此题。
2.题目数据保证res数组的元素都在 32 位整数范围内
3.有O(1)空间复杂度的做法，返回的res数组不计入空间复杂度计算

数据范围：1 < n \le 100\1<n≤100
示例1
输入：
[1,2,3,4]
复制
返回值：
[24,12,8,6]
复制
说明：
res[0]=2*3*4=24
res[1]=1*3*4=12
res[2]=1*2*4=8
res[3]=1*2*3=6

*/

/*
利用左右前缀数组,对数据进行预处理,得到 每个值左侧和右侧的乘积,则除自身外的乘积为 左侧乘积 * 右侧乘积

class Solution:
    def timesExceptSelf(self , nums: List[int]) -> List[int]:
        # write code here
        l, r = [1 for _ in range(len(nums))], [1 for _ in range(len(nums))]
        for i in range(1, len(nums)):
            l[i] = l[i - 1] * nums[i - 1]
        for i in range(len(nums) - 2, -1, -1):
            r[i] = r[i + 1] * nums[i + 1]
        res = [0 for _ in range(len(nums))]
        for i in range(len(nums)):
            res[i] = l[i] * r[i]
        return res

*/

func timesExceptSelf(nums []int) []int {
	// write code here
	l, r := make([]int, len(nums)), make([]int, len(nums))
	for k, _ := range nums {
		l[k]++
		r[k]++
	}
	for i := 1; i < len(nums); i++ {
		l[i] = l[i-1] * nums[i-1]
	}

	for j := len(nums) - 2; j >= 0; j-- {
		r[j] = r[j+1] * nums[j+1]
	}
	res := make([]int, len(nums))
	for k, _ := range nums {
		res[k] = l[k] * r[k]
	}

	return res
}

func main() {
	self := timesExceptSelf([]int{1, 2, 3, 4})
	fmt.Println(self)
}
