package _3_hash

import (
	"fmt"
	"strings"
)

/*
https://www.nowcoder.com/practice/bb40e0aae84a4491a35f93b322516cb5?tpId=117&tqId=39349&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj&difficulty=undefined&judgeStatus=undefined&tags=&title=%E5%8D%95%E8%AF%8D%E6%8B%86%E5%88%86

描述
给定一个字符串 s 和一个字符串数组 dic, 在字符串的任意位置拆分任意次后得到的字符串集合 是否 是给定字符串数组的子集。（即拆分后的字符串集合中的所有字符串都在 dic 数组中），

示例1
输入：
"nowcoder", ["now","coder","no","wcoder"]
返回值：
["no wcoder", "now coder"]

示例2
输入：
"nowcoder", ["now","wcoder"]
返回值：
[]

示例3
输入：
"nowcoder", ["nowcoder"]
返回值：
["nowcoder"]

示例4
输入：
"nownowcoder", ["now", "coder"]
返回值：
["now now coder"]

说明：
你可以重复使用 dic 数组中的字符串

*/

// 思路： 以dic为基础，做排列组合（全排列），然后看是否等于s, 等于存到数组中
// 很多边界没有考虑情况
//func wordDiv2(s string, dic []string) []string {
//	arr := []string{}
//	for i := 0; i < len(dic); i++ {
//		for j := 0; j < len(dic); j++ {
//			if i == j && i != 0 {
//				continue
//			}
//			if len(dic) > 1 {
//				arr = append(arr, dic[i]+" "+dic[j])
//			} else {
//				arr = append(arr, dic[i])
//			}
//		}
//	}
//	result := []string{}
//	for i := 0; i < len(arr); i++ {
//		strs := strings.Split(arr[i], " ")
//		if len(strs) > 1 && strs[0]+strs[1] == s {
//			result = append(result, arr[i])
//		} else if len(strs) == 1 && strs[0] == s {
//			result = append(result, arr[i])
//		}
//	}
//	return result
//}

/*
nowcoder1  nowcoder1    ==> nowcoder1
nowcoder11  nowcoder1 1  ==> nowcoder1 1 1
*/

// 解法1
func wordDiv( s string ,  dic []string ) []string {
	// write code here
	m := make(map[string]bool)
	for _, v := range dic{
		m[v] = true
	}
	var (
		ans []string
		path []string
		dfs func (idx int)
	)
	// 递归 +　回溯
	dfs = func(idx int) {
		if idx == len(s) {
			ans = append(ans, strings.Join(path, " "))
			return
		}
		for i := idx + 1; i <= len(s); i++ {
			if m[s[idx:i]] {
				path = append(path, s[idx:i])
				dfs(i)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return ans
}

// 解法2
func wordBreak(s string, wordDict []string) []string {
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}
	n := len(s)
	ans := []string{}

	var dfs func(start, end int, sb []string)
	dfs = func(start, end int, sb []string) {
		// 最终，做拼接
		if start >= n {
			ans = append(ans, strings.Join(sb, " "))
			return
		}
		// 弹栈
		if end >= n {
			return
		}
		// 滑动窗口，当字符串符合要求，将下次递归的首位更新
		tmp := s[start : end+1]
		if wordSet[tmp] {
			dfs(end+1, end+1, append(sb, tmp))
		}
		dfs(start, end+1, sb)
	}

	dfs(0, 0, []string{})
	return ans
}

/*
解法2扩展 - 字典树
class Trie:  # trie类
    def __init__(self):
        self.children = dict()  # 用dict存储类
        self.words = list()

class Solution:
    def wordBreak(self, s: str, wordDict: List[str]) -> List[str]:
        n = len(s)
        ans = []
        root = Trie()  # 初始化trie树
        for word in wordDict:  # 构建trie树
            cur = root
            for c in word:
                if c not in cur.children:
                    cur.children[c] = Trie()
                cur = cur.children[c]
            cur.words.append(word)

        def dfs(i, node, sb):
            if i == n:
                if node == root:#回到根
                    ans.append(" ".join(sb))
                return
            if s[i] not in node.children:#如果不匹配立刻结束
                return
            node = node.children[s[i]]
            if node.words:#断开
                dfs(i + 1, root, sb + [node.words[0]])
            dfs(i + 1, node, sb)#继续

        node = root
        dfs(0, node, [])
        return ans

作者：cheturtle
链接：https://leetcode.cn/problems/word-break-ii/solutions/2768416/python-mei-ju-fen-ge-dian-35ms-by-nrib8z-3fg1/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/

func main() {
	div1 := wordDiv("nownowcoder", []string{"now", "coder"})
	div2 := wordBreak("nownowcoder", []string{"n", "now", "owcoder", "coder"})
	fmt.Println(div1)
	fmt.Println(div2)
}
