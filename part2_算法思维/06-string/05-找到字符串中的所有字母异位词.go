package main

import "reflect"

/*
找到字符串中的所有字母异位词

给定一个字符串 s 和一个非空字符串 p，找出 s 中的所有是 p 的 Anagrams 字符串的子串，返回这些子串的起始索引。

解题思路
这道题是一道考“滑动窗口”的题目。和第 3 题，第 76 题，第 567 题类似的。解法也是用 freq[256] 记录每个字符的出现的频次次数。
滑动窗口左边界往右滑动的时候，划过去的元素释放次数(即次数 ++)，滑动窗口右边界往右滑动的时候，划过去的元素消耗次数(即次数 --)。
右边界和左边界相差 len(p) 的时候，需要判断每个元素是否都用过一遍了。具体做法是每经过一个符合规范的元素，count 就 --，count 初始值是 len(p)，
当每个元素都符合规范的时候，右边界和左边界相差 len(p) 的时候，count 也会等于 0 。
当区间内有不符合规范的元素(freq < 0 或者是不存在的元素)，那么当区间达到 len(p) 的时候，count 无法减少到 0，
区间右移动的时候，左边界又会开始 count ++，只有当左边界移出了这些不合规范的元素以后，才可能出现 count = 0 的情况，即找到 Anagrams 的情况。

*/

// 滑动窗口
//func findAnagrams(s string, p string) []int {
//	var freq [256]int
//	result := []int{}
//	if len(s) == 0 || len(s) < len(p) {
//		return result
//	}
//	for i := 0; i < len(p); i++ {
//		freq[p[i]-'a']++
//	}
//	left, right, count := 0, 0, len(p)
//
//	for right < len(s) {
//		if freq[s[right]-'a'] >= 1 {
//			count--
//		}
//		freq[s[right]-'a']--
//		right++
//		if count == 0 {
//			result = append(result, left)
//		}
//		if right-left == len(p) {
//			if freq[s[left]-'a'] >= 0 {
//				count++
//			}
//			freq[s[left]-'a']++
//			left++
//		}
//
//	}
//	return result
//}

func findAnagrams(s string, p string) []int {
	pCount := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		pCount[p[i]]++
	}

	var result []int
	sCount := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		// 将端点加入窗口
		sCount[s[i]]++

		// 如果当前窗口长度大于模式串长度, 将左端点移除窗口; 如果左端点中的元素 个数为0, 则将该端点删除
		if i >= len(p) {
			sCount[s[i-len(p)]]--
			if sCount[s[i-len(p)]] == 0 {
				delete(sCount, s[i-len(p)])
			}
		}

		// 判断窗口内的字符与模式串是否相同
		if reflect.DeepEqual(sCount, pCount) {
			result = append(result, i-len(p)+1)
		}
	}

	return result
}
