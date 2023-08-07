package _5_前缀和

/*
NC213 除自身以外数组的乘积
https://www.nowcoder.com/practice/0786aa81c1c64c2a990e393fac811b45?tpId=196&tqId=39406&ru=/exam/oj

知识点
数组
前缀和

描述
给定一个长度为 n 的数组 nums ，返回一个数组 res，res[i]是nums数组中除了nums[i]本身以外其余所有元素的乘积，即：
res[i]=nums[1]×nums[2]×......×nums[i−1]×nums[i+1]......×nums[n]
1.请不要使用除法，并且在 O(n) 时间复杂度内完成此题。
2.题目数据保证res数组的元素都在 32 位整数范围内。
3.有O(1)空间复杂度的做法，返回的res数组不计入空间复杂度计算。

输入：
[1,2,3,4]
返回值：
[24,12,8,6]

说明：
res[0]=2*3*4=24
res[1]=1*3*4=12
res[2]=1*2*4=8
res[3]=1*2*3=6
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
