package main

import "fmt"

/*
NC125 和为K的连续子数组
https://www.nowcoder.com/practice/704c8388a82e42e58b7f5751ec943a11?tpId=117&tqId=37794&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=585&title=


描述
给定一个无序数组 arr, 其中元素可正、可负、可0。给定一个整数 k ，求 arr 所有连续子数组中累加和为k的最长连续子数组长度。保证至少存在一个合法的连续子数组。
[1,2,3]的连续子数组有[1,2]，[2,3]，[1,2,3] ，但是[1,3]不是

示例1
输入：
[1,-2,1,1,1],0
返回值：
3

示例2
输入：
[0,1,2,3],3
返回值：
3

*/

func maxlenEqualK(arr []int, k int) int {
	// write code here
	/*
	   // 定义一个整型数组，用于存放前缀和
	   int[] pre = new int[arr.length];
	   // 别忘了，初始化 pre 数组
	   pre[0] = arr[0];
	   for (int i = 1; i < pre.length; i++) {
	       pre[i] = pre[i - 1] + arr[i];
	   }

	   // 定义一个 HashMap，用于存放某个前缀和以及它所在的位置
	   HashMap<Integer, Integer> hashMap = new HashMap<>();
	   hashMap.put(0, -1);
	   // 定义一个整型变量，用于存放最终的返回结果
	   int res = 0;

	   // 说明：区间 [i, j] 的和等于 pre[j] - pre[i - 1]
	   for (int i = 0; i < pre.length; i++) {

	       int tmp = hashMap.getOrDefault(pre[i], -2);
	       if (tmp == -2) {
	           hashMap.put(pre[i], i);
	       }

	       int index = hashMap.getOrDefault(pre[i] - k, -2);
	       if (index != -2) {
	           res = Math.max(res, i - index);
	       }
	   }
	   return res;
	*/

	/*
	   if (arr.empty())
	       return 0;
	   unordered_map <int, int> hash; // 定义哈希表
	   hash[0] = -1; // 初始化哈希表
	   int n = arr.size(), res = 0, preSum = 0;
	   for (int i = 0; i < arr.size(); i ++) {
	       preSum += arr[i]; // 到当前位置的前缀和
	       if (hash.find(preSum) == hash.end()) // 哈希表中未找到
	           hash[preSum] = i;
	       int target = preSum - k; // 需要在哈希表中查找的值
	       if (hash.find(target) != hash.end()) // 找到
	           res = max(res, i - hash[target]); // 更新res
	   }
	   return res;
	*/

	// 暴力解法
	//count := 0
	//
	//for i := 0; i < len(arr); i++ {
	//	if len(arr)-i <= count {
	//		break
	//	}
	//	sum := 0
	//	for j := i; j < len(arr); j++ {
	//		sum += arr[j]
	//		if sum == k {
	//			if count <= j-i {
	//				count = j + 1 - i
	//			}
	//		}
	//	}
	//
	//}
	//return count

	// 前缀和 + hash, 代码有bug
	l := len(arr)
	if l == 0 {
		return 0
	}
	var (
		sum    int                 // 前缀和
		maxLen int                 // 累加和为k的最大子数组长度
		m      = make(map[int]int) // 某个sum出现的最早位置
	)
	m[0] = -1
	for i, n := range arr {
		sum += n

		// 数据就是i, j 之间的数据
		if j, ok := m[sum-k]; ok {
			// 如果比圆的maxlen 大，更新
			if i-j > maxLen {
				maxLen = i - j
			}
		}

		// 构建sum 为key, 索引为value 的hash表
		if _, ok := m[sum]; !ok {
			m[sum] = i
		}

	}
	return maxLen
}

func main() {
	//k := maxlenEqualK([]int{1, -2, 1, 1, 1}, 0)
	//k := maxlenEqualK([]int{1, -2, 1, 1, 1, -1, 1}, 2)
	k := maxlenEqualK([]int{1, -2, 1, 1}, 2)
	fmt.Println(k)
}
