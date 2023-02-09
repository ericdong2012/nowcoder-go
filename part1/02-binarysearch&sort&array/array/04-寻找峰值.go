package main

/*
NC107 寻找峰值
https://www.nowcoder.com/practice/fcf87540c4f347bcb4cf720b5b350c76?tpId=117&tqId=37831&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=590,5058,578&title=


描述
给定一个长度为n的数组nums，请你找到峰值并返回其索引。数组可能包含多个峰值，在这种情况下，返回任何一个所在位置即可。
1.峰值元素是指其值严格大于左右相邻值的元素。严格大于即不能有等于
2.假设 nums[-1] = nums[n] = -\infty−∞
3.对于所有有效的 i 都有 nums[i] != nums[i + 1]
4.你可以使用O(logN)的时间复杂度实现此问题吗？


如输入[2,4,1,2,7,8,4]时，会形成两个山峰，一个是索引为1，峰值为4的山峰，另一个是索引为5，峰值为8的山峰，如下图所示：

示例1
输入：
[2,4,1,2,7,8,4]
返回值：
1
说明：
4和8都是峰值元素，返回4的索引1或者8的索引5都可以

示例2
输入：
[1,2,3,1]
返回值：
2
说明：
3 是峰值元素，返回其索引 2

*/

func findPeakElement(nums []int) int {
	// write code here
	// 分治，动态规划，递归 本质没有什么区别 +  二分查找
	/*
		class Solution:
		    def findPeakElement(self , nums: List[int]) -> int:
		        left = 0
		        right = len(nums) - 1
		        # 二分法
				# 查找某一区间的最大值
		        while left < right:
		            mid = int((left + right) / 2)
		            # 右边是往下，不一定有坡峰
		            if nums[mid] > nums[mid+1]:
		                right = mid
		            # 右边是往上，一定能找到波峰
		            else:
		                left = mid + 1
		        # 其中一个波峰
		        return right
	*/
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := int((right + left) / 2)
		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return right

}
