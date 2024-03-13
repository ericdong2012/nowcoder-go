# -*- coding: utf-8 -*-
from typing import List

"""
实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
必须原地修改，只允许使用额外常数空间。

以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/next-permutation
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
"""
class Solution:
    def nextPermutation(self, nums: List[int]) -> List[int]:
        """
        https://leetcode.cn/problems/next-permutation/solution/xia-yi-ge-pai-lie-by-leetcode-solution/
        """
        def reverse(nums, start, end):
            while start < end:
                nums[start], nums[end] = nums[end], nums[start]
                start += 1
                end -= 1
            return nums

        n = len(nums)
        if n < 2: return nums

        # [4,5,3,1,2,6]
        # step1: 从后往前找 找满足nums[i - 1] <  nums[i] 的序列，记录i, 找到较小数（前面是降序， 突然变升序）,  注意： nums[i-1] 才是找到的数
        i = n - 1
        while i > 0 and nums[i - 1] >= nums[i]:
            i -= 1
        # i > 0, 在中间找到了合适的数
        if i > 0:
            # step2: 从后往前找，找nums[k] > nums[i-1], 找到较大数
            k = n - 1
            while nums[i - 1] >= nums[k] :
                k -= 1
            # print(k)
            # step3: 交换 nums[i - 1], nums[k]
            nums[i - 1], nums[k] = nums[k], nums[i - 1]

        # 找到最前都没有（i=0）， 比如： [4，3，2，1] , 反转
        reverse(nums, i, n - 1)
        return nums

s = Solution()
print(s.nextPermutation([1, 2, 4, 3, 6, 2]))


# go 解法
"""
func nextPermutation(nums []int)  []int{
	n := len(nums)
	if n < 2 {
		return nums
	}
	i:= n -1
	for i > 0 && nums[i-1] >= nums[i] {
		i--
	}

	if i > 0 {
		j := n -1
		for j >0 && nums[i-1] >= nums[j] {
			j --
		}
		nums[j], nums[i-1] = nums[i-1], nums[j]
	}

	reverse(nums[i:])
	return nums
}

func reverse(a []int) {
    for i, n := 0, len(a); i < n/2; i++ {
        a[i], a[n-1-i] = a[n-1-i], a[i]
    }
}

作者：LeetCode-Solution
链接：https://leetcode.cn/problems/next-permutation/solution/xia-yi-ge-pai-lie-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

"""