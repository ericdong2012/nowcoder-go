package _3_hash

/*
题⽬⼤意
在数组中找到 2 个数之和等于给定值的数字，结果返回 2 个数字在数组中的下标。


解题思路
这道题最优的做法时间复杂度是 O(n)。
顺序扫描数组，对每⼀个元素，在 map 中找能组合给定值的另⼀半数字，如果找到了，直接返回 2 个
数字的下标即可。如果找不到，就把这个数字存⼊ map 中，等待扫到“另⼀半”数字的时候，再取出来返
回结果。

*/

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[nums[i]] = i
	}

	return nil
}
