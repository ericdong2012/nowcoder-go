package main

import "fmt"

/*
NC91 最长上升子序列(三)
https://www.nowcoder.com/practice/9cf027bf54714ad889d4f30ff0ae5481?tpId=117&tqId=37796&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=5058&title=

描述
给定数组 arr ，设长度为 n ，输出 arr 的最长上升子序列。（如果有多个答案，请输出其中 按数值(注：区别于按单个字符的ASCII码值)进行比较的 字典序最小的那个）

示例1
输入：
[2,1,5,3,6,4,8,9,7]
返回值：
[1,3,4,8,9]

示例2
输入：
[1,2,8,6,4]
返回值：
[1,2,4]

说明：
其最长递增子序列有3个，（1，2，8）、（1，2，6）、（1，2，4）其中第三个 按数值进行比较的字典序 最小，故答案为（1，2，4）
*/

/*
两步走：
第一步——求最长递增子序列长度
第二步——求字典序靠前的子序列


对于第一步，有两种解法：
动态规划，时间复杂度为O(n^2)，会超时
贪心+二分，时间复杂度为O(nlogn)

下面说说贪心+二分的解法，举例说明基本思路，假设数组arr为[2, 3, 1, 2, 3]，vec数组里面存放递增子序列，maxLen数组里存放以元素i结尾的最大递增子序列长度，那么遍历数组arr并执行如下更新规则:

初始情况下，vec为[2]，maxLen[1]
接下来遇到3，由于vec最后一个元素小于3，直接更新，vec为[2,3]，maxLen[1,2]
接下来遇到1，由于vec最后的元素大于1, 我们在vec中查找大于等于1的第一个元素的下标，并用1替换之，此时vec为[1,3], maxLen[1,2,1]
接下来遇到2，由于vec最后的元素大于2，我们在vec中查找大于等于2的第一个元素的下标，并用2替换之，此时vec为[1,2], maxLen[1,2,1,2]
接下来遇到3，由于vec最后一个元素小于3，直接更新，vec为[1,2,3]，maxLen为[1,2,1,2,3]
此时vec的大小就是整个序列中最长递增子序列的长度（但是vec不一定是本题的最终解）



对于第二步，假设我们原始数组是arr1，得到的maxLen为[1,2,3,1,3]，最终输出结果为res（字典序最小的最长递增子序列），
res的最后一个元素在arr1中位置无庸置疑是maxLen[i]==3对应的下标，那么到底是arr1[2]还是arr1[4]呢？如果是arr1[2]，那么arr1[2]<arr1[4]，则maxLen[4]==4，
与已知条件相悖。因此我们应该取arr1[4]放在res的最后一个位置。

*/

func LIS(arr []int) []int {
	if len(arr) == 0 {
		return nil
	}
	// 记录复合要求的切片长度
	maxLen := make([]int, len(arr))
	// 记录切片
	res := []int{arr[0]}
	maxLen[0] = 1
	for i := 1; i < len(arr); i++ {
		// [1,2,8,6,4] [1]
		if arr[i] > res[len(res)-1] {
			// [1,2]   [1,2,0,0,0]
			// [1,2,8] [1,2,3,0,0]
			res = append(res, arr[i])
			maxLen[i] = len(res)
		} else {
			// 6
			// res: [1,2,8]
			// 在res中查找大于等于 arr[i] 的第一个元素的下标，并用arr[i] 替换它
			l, r := 0, len(res)-1
			for l < r {
				// l + (右减左除以2)  ??? 为什么不直接是右减左除以2
				mid := l+(r-l)>>1
				if res[mid] >= arr[i] {
					r = mid
				} else {
					l = mid+1
				}
			}
			// 替换res[l]
			res[l] = arr[i]
			// 记录复合要求的切片长度为 左 + 1
			maxLen[i] = l+1
		}
	}
	// 处理最小字典序问题
	for i, j := len(maxLen)-1, len(res); j > 0; i-- {
		// [1,2，3，3，3]  [1,2,8] / [1,2,6] /[1,2,4]
		// maxLen中的元素 == res长度时
		if maxLen[i] == j {
			j--
			res[j] = arr[i]
		}
	}
	return res
}

func main() {
	fmt.Printf("%+v", LIS([]int{1, 2, 8, 6, 4}))
}
