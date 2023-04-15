package main

import "fmt"

/*
NC202 长度最小的连续子数组
https://www.nowcoder.com/practice/10dd5f8c5d984aa3bd69788d86aaef23?tpId=117&tqId=39391&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=5058&title=

描述
给定一个数组 nums 和一个正整数 target , 找出满足和大于等于 target 的长度最短的连续子数组并返回其长度，如果不存在这种子数组则返回0。

示例1
输入：
[1,2,4,4,1,1,1],9
返回值：
3

示例2
输入：
[1,4,4,4,1,1,1],3
返回值：
1

*/

/*
二分+前缀和

双指针+窗口
def minSubarray(self , nums: List[int], target: int) -> int:
    # write code here
    #左右指针移动滑窗求长度
    left = 0          #左指针
    right = 0         #右指针
    sums = 0          #左右指针滑窗内部的值之和
    ans = len(nums)   #这个是用来比较滑窗内长短的值，可以设为无限大，这里我设为nums长度
    while right < len(nums):     #右指针先走，走完整个nums
        while right < len(nums) and sums < target:
        #这里再次限定了右指针行动范围，同时讨论sums没有大于target的情况
        #这种时候就要把右指针右移，扩大标定范围，同时将sums加上右移添加的值
            sums += nums[right]
            right += 1
        while left <right and sums >= target:
        #这里则是使得保证左指针小于右指针，且当sums是大于等于target时的情况
        #这种情况是满足要求的，但是需要比较所含范围的长短
            ans  = min(ans,right-left) #比较长短，最后就已经可以得到ans
            sums -= nums[left] #如果长了就右移左指针，减去值
            left += 1
    return ans

*/

func minSubarray(nums []int, target int) int {
	// write code here
	left, right, sums := 0, 0, 0
	ans := len(nums)

	for right < len(nums) {
		for right < len(nums) && sums < target {
			sums += nums[right]
			right += 1
		}
		for right <= len(nums) && sums >= target {
			// 最短长度
			ans = min4(ans, right-left)
			// 总和中减去最左侧的数字，往右移动
			sums -= nums[left]
			left += 1
		}
	}

	return ans
}

func min4(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	subarray := minSubarray([]int{1, 2, 4, 4, 1, 1, 1}, 9)
	//subarray := minSubarray([]int{1, 1, 1, 2, 4}, 7)
	fmt.Println(subarray)
}
