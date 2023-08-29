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
        Do not return anything, modify nums in-place instead.
        """

        """
        标准的“下一个排列”算法可以描述为：

        1. 从后向前查找第一个相邻升序的元素对 (i,j)，满足 A[i] < A[j]。此时 [j,end) 必然是降序
        2. 在 [j,end) 从后向前查找第一个满足 A[i] < A[k] 的 k。 A[i]、A[k] 分别就是上文所说的「小数」、「大数」
        3. 将 A[i] 与 A[k] 交换
        4. 可以断定这时 [j,end) 必然是降序，逆置 [j,end)，使其升序
        5. 如果在步骤 1 找不到符合的相邻元素对，说明当前 [begin,end) 为一个降序顺序，则直接跳到步骤 4
        该方法支持数据重复，且在 C++ STL 中被采用

        作者：imageslr
        链接：https://leetcode-cn.com/problems/next-permutation/solution/xia-yi-ge-pai-lie-suan-fa-xiang-jie-si-lu-tui-dao-/
        来源：力扣（LeetCode）
        著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

        """

        # if len(nums) <= 1:
        #     return
        # i, j, k = len(nums) - 2, len(nums) - 1, len(nums) - 1
        # # find: A[i] < A[j]
        # while i >= 0 and nums[i] >= nums[j]:
        #     i -= 1
        #     j -= 1
        # if i >= 0:
        #     # 不是最后一个排列
        #     # find: A[i] < A[k]
        #     while nums[i] >= nums[k]:
        #         k -= 1
        #
        # # swap A[i], A[k]
        # nums[i], nums[k] = nums[k], nums[i]
        #
        # # reverse  A[j:end]
        # for i in range(j, len(nums) - 1):
        #     if i < j:
        #         i, j = i + 1, j - 1
        #         nums[i], nums[j] = nums[j], nums[i]
        #
        # return nums

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
        # step1: 找满足nums[i - 1] <  nums[i] 的序列，记录i, 找到较小数,  注意： nums[i-1] 才是找到的数
        # 从后往前找，直到 nums[i - 1] < nums[i]
        i = n - 1
        while i > 0 and nums[i - 1] >= nums[i]:
            i -= 1
        # 找到最前都没有， 比如： [4，3，2，1] ， 反转
        if i > 0:
            # step2: 从后往前找，找nums[k] > nums[i-1], 找到较大数，交换nums[i - 1], nums[k]

            k = n - 1
            while nums[i - 1] >= nums[k] :
                k -= 1
            # print(k)
            nums[i - 1], nums[k] = nums[k], nums[i - 1]

        # step3: 反转
        reverse(nums, i, n - 1)
        return nums

s = Solution()
print(s.nextPermutation([1, 2, 4, 3, 6, 2]))


"""
func nextPermutation(nums []int) {
    #     n := len(nums)
    #     i := n - 2
    #     for i >= 0 && nums[i] >= nums[i+1] {
    #         i--
    #     }
    #     if i >= 0 {
    #         j := n - 1
    #         for j >= 0 && nums[i] >= nums[j] {
    #             j--
    #         }
    #         nums[i], nums[j] = nums[j], nums[i]
    #     }
    #     reverse(nums[i+1:])


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