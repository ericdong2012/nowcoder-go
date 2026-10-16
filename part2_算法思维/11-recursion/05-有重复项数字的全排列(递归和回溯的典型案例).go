package main

import (
	"fmt"
	"sort"
)

/*
NC42 有重复项数字的全排列
https://www.nowcoder.com/practice/a43a2b986ef34843ac4fdd9159b69863?tpId=117&tqId=37739&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=591&title=


描述
给出一组可能包含重复项的数字，返回该组数字的所有排列。结果以字典序升序排列。
要求：空间复杂度 O(n!)，时间复杂度 O(n!)

示例1
输入：
[1,1,2]
返回值：
[[1,1,2],[1,2,1],[2,1,1]]

示例2
输入：
[0,1]
返回值：
[[0,1],[1,0]]
*/


/*
参考官方，精华题解

思路： 递归 +　回溯
我们就使用临时数组去组装一个排列的情况：每当我们选取一个数组元素以后，就确定了其位置，只需要对数组中剩下的元素进行全排列添加在该元素后面，那么给剩余部分进行全排列就是一个子问题，因此可以使用递归。
终止条件： 临时数组中选取了n个元素，已经形成了一种排列情况了，可以将其加入输出数组中。
返回值：   每一层给上一层返回的就是本层级在临时数组中添加的元素，递归到末尾的时候就能添加全部元素。
本级任务： 每一级都需要选择一个不重复元素加入到临时数组末尾（遍历数组选择）。

回溯的思想也与　没有重复数字的全排列　类似，对于数组[1,2,2,3]，如果事先在临时数组中加入了1，后续子问题只能是[2,2,3]的全排列接在1后面，对于2开头的分支达不到，因此也需要回溯：将临时数组刚刚加入的数字pop掉，同时vis修改为没有加入，这样才能正常进入别的分支。
//字典，标记为使用过
vis[i] =  true;
//加入数组
temp.add(num[i]);
recursion(res, num, temp, vis);
//回溯
vis[i] =  false;
temp.remove(temp.size() - 1);

具体做法：
step1：先对数组按照字典序排序，获取第一个排列情况。
step2：准备一个数组暂存递归过程中组装的排列情况。使用额外的vis数组用于记录哪些位置的数字被加入了。
step3：每次递归从头遍历数组，获取数字加入：首先根据vis数组，已经加入的元素不能再次加入了；同时，如果当前的元素num[i]与同一层的前一个元素num[i-1]相同且num[i-1]已经用，也不需要将其纳入。
step4：进入下一层递归前将vis数组当前位置标记为使用过。
step5：回溯的时候需要修改vis数组当前位置标记，同时去掉刚刚加入数组的元素，
step6：临时数组长度到达原数组长度就是一种排列情况。 break条件


数据结构：　数组 + 字典
*/

// 解法1
func permuteUnique( num []int ) [][]int {
	// write code here
	n := len(num)
	if n == 0 {
		return [][]int{}
	}
	sort.Ints(num)

	var helper func(cur []int)
	var ans [][]int
	var used = make(map[int]bool, n)
	helper = func(cur []int) {
		// break条件
		if len(cur) == n {
			target := make([]int, n)
			copy(target, cur)
			ans = append(ans, target)
			return
		}
		for i:=0; i<n; i++{
			// 如果使用过，不用了
			if used[i] == true {
				continue
			//	如果当前值和之前的值一样，不用了
			} else if i > 0 && used[i-1] != true && num[i] == num[i-1] {
				continue
			} else {
				// 标记使用，添加结果到临时数组
				used[i] = true
				cur = append(cur, num[i])
				// 递归
				helper(cur)
				// 回溯
				used[i] = false
				cur = cur[:len(cur)-1]
			}
		}
	}
	helper(nil)
	return ans
}

// 解法2
//func permuteUnique(num []int) [][]int {
//	// write code here
//
//	if len(num) == 0 {
//		return nil
//	}
//	sort.Ints(num)
//	result := [][]int{}
//	dfs1(num, 0, &result)
//	return result
//}
//
//func dfs1(nums []int, index int, result *[][]int) {
//	// break条件，长度够了，添加到结果中
//	if index == len(nums)-1 {
//		temp := make([]int, len(nums))
//		copy(temp, nums)
//		*result = append(*result, temp)
//		return
//	}
//	// 注意起始点和终止点
//	for i := index; i < len(nums); i++ {
//		// 剪枝, 下一次循环的，位置不同，字符相同，跳过
//		if i != index && nums[i] == nums[index] {
//			continue
//		}
//		// 交换
//		nums[i], nums[index] = nums[index], nums[i]
//		// 递归处理下一个位置
//		dfs1(nums, index+1, result)
//	}
//
//	// 回溯，nums数据还原
//	for i := len(nums) - 1; i > index; i-- {
//		nums[i], nums[index] = nums[index], nums[i]
//	}
//}

func main09() {
	unique := permuteUnique([]int{1, 1, 2})
	fmt.Println(unique)
}

/*

class Solution:
    def recursion(self, res:List[List[int]], num:List[int], temp:List[int], vis:List[int]):
        # 临时数组满了加入输出
        if len(temp) == len(num):
            res.append(temp.copy())
            return
        # 遍历所有元素选取一个加入
        for i in range(len(num)):
            # 如果该元素已经被加入了则不需要再加入了
            if vis[i]==1:
                continue
            # 当前的元素num[i]与同一层的前一个元素num[i-1]相同且num[i-1]已经用过了
            if i > 0 and num[i - 1] == num[i] and not vis[i - 1]:
                continue

            #标记为使用过
            vis[i] = 1
            #加入数组
            temp.append(num[i])
            self.recursion(res, num, temp, vis)
            # 回溯
            vis[i] = 0
            temp.pop()

    def permuteUnique(self , num: List[int]) -> List[List[int]]:
        #先按字典序排序
        num.sort()
        #标记每个位置的元素是否被使用过
        vis = [0]*len(num)
        temp = list()
        res = list(list())
        # 递归获取
        self.recursion(res, num, temp, vis)
        return res

*/

