package main

import "fmt"

/*
NC28 最小覆盖子串
https://www.nowcoder.com/practice/c466d480d20c4c7c9d322d12ca7955ac?tpId=117&tqId=37732&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=585&title=

描述:
给出两个字符串 s 和 t，要求在 s 中找出最短的包含 t 中所有字符的连续子串。

例如：
S ="XDOYEZODEYXNZ"
T ="XYZ"T="XYZ"
找出的最短子串为"YXNZ"

注意：
如果 s 中没有包含 t 中所有字符的子串，返回空字符串 “”；
满足条件的子串可能有很多，但是题目保证满足条件的最短的子串唯一。

示例1
输入：
"XDOYEZODEYXNZ","XYZ"
返回值：
"YXNZ"

*/

// 该题和09-doublepointer 中的最小覆盖子串是同一题
// 滑动窗口， 哈希表
// 字母表是已知的，所以可以建立hash表
func minWindow(S string, T string) string {
	// hash + double point + sliding windows
	/*
		step 1：字符串仅包含大小写字母，则字符集是已知且有限的，那这种情况下我们可以考虑使用哈希表——只需要维护一个哈希表，里面是字符串T的字符为key值，初始化时当字符在T中出现一次则value值减1，后续如果找到就可以将其加回来。
		step 2：依次遍历字符串S，如果匹配则将哈希表中的相应的字符加1。
		step 3：在遍历过程中维护一个窗口，如果哈希表中所有元素都大于0，意味着已经找全了，则窗口收缩向左移动，找最小的窗口，如果不满足这个条件则窗口右移继续匹配。窗口移动的时候需要更新最小窗口，以取得最短子串。
		step 4：如果匹配到最后，窗口left（初始为-1）也没有右移，说明没有找到，返回空串即可。
		step 5：最后使用字符串截取函数，截取刚刚记录下的窗口即可得到符合条件的最短子串。
	*/
	lenS, lenT := len(S), len(T)
	// 构建基于T的hash 表
	windows := make(map[byte]int)
	for i := 0; i < lenT; i++ {
		windows[T[i]]++
	}

	// 左右指针
	left, right := 0, 0
	// windows中字母大于0 计数
	count := 0
	ans := ""
	for right < lenS {
		char := S[right]
		right++
		// windows中 已有会-1， 没有会直接赋值 -1
		windows[char]--
		if windows[char] >= 0 {
			count++
		}
		// 个数达到， ans 从最初的完整字符串一直在减小 "abcAbA" "bcAbA"
		for count == lenT {
			// 取较小值， ans 从最初的完整字符串一直在减小 "abcAbA" "bcAbA"
			if ans == "" || len(ans) > right-left {
				ans = S[left:right]
			}
			// 如果 windows[S[left]] = 0 ,  count-1, 最终的跳出条件在这里
			if windows[S[left]] == 0 {
				count--
			}
			// left 往右移动，windows相应的值加1,
			windows[S[left]]++
			left++
		}
	}
	return ans
}

func main() {
	window := minWindow("XDOYEZODEYXNZ", "XYZ")
	fmt.Println(window)
}
