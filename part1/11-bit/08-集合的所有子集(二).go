package main

import (
	"fmt"
	"sort"
)

/*
NC221 集合的所有子集(二)
https://www.nowcoder.com/practice/a3dfd4bc8ae74fad9bc65d5ced7ae813?tpId=117&tqId=39446&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=5074&title=


描述
给定一个整数数组 nums ，其中可能包含重复元素，请你返回这个数组的所有可能子集。
返回的答案中不能包含重复的子集，将答案按字典序进行排序。

示例1
输入：
[1,2]
返回值：
[[],[1],[1,2],[2]]

示例2
输入：
[1]
返回值：
[[],[1]]

*/

// 对比02-binarysearch&sort&array/array/03-集合的所有子集
// 区别在于上述是没有重复元素，该题是有重复元素

func subsets(nums []int) [][]int {
	// write code here
	/*
		递归， 回溯， 访问标记
		举例说明，对于数组[1,2,2]，将其标记为[1,2,2']。

		排序后，对于相邻的相等元素：
		如果我们访问了2，再访问2'; 将子集[2, 2']加入结果。继续。
		如果我们访问了2'，却还没有访问2，我们是不希望有这种情况的，得跳过。

		class Solution {
		public:
		    vector<vector<int>> res;
		    vector<int> tmp;
		    bitset<10> visited;
		    void backTrace(vector<int>& nums, int start){
		        res.push_back(tmp);
		        for(size_t i = start; i < nums.size(); i++){
		            if(i > 0 && nums[i] == nums[i-1] && !visited[i-1]) continue;
		            tmp.push_back(nums[i]);
		            visited[i] = 1;
		            backTrace(nums, i + 1);
		            tmp.pop_back();
		            visited[i] = 0;
		        }
		    }
			vector<vector<int> > subsets(vector<int>& nums) {
				// write code here
				if(nums.empty()) return res;
				sort(nums.begin(), nums.end());
				backTrace(nums, 0);
				return res;
			}
		}

		递归剪枝

		相比于方法一，我们引入了set集合
		该集合相当于对nums中的元素进行了去重操作，因此相当于数组的长度n的值有一定程度变小，优化了时间开销。

		class Solution:
		    def subsets(self , nums: List[int]) -> List[List[int]]:
		        # write code here
		        res=[]
		        nums.sort()
		        def backtrace(start,path):
		            res.append(path)                    # 结果加入res中
		            add=set()
		            for i in range(start,len(nums)):    # 起点
		                if nums[i] not in add:          # 判断是否已经存在，元素去重
		                    add.add(nums[i])
		                    cur_path=path+[nums[i]]
		                    backtrace(1+i, cur_path)
		        backtrace(0,[])
		        return res

		时间复杂度：O(n×2n)个状态，每个状态需要O(n)的时间代价来构造
		空间复杂度：O(n)，递归深度的空间开销
	*/
	sort.Ints(nums)
	//使用深度优先搜索
	var res [][]int
	var DFS func(start int, temp []int)
	DFS = func(start int, temp []int) {
		res = append(res, temp)
		for i := start; i < len(nums); i++ {
			if i > start && nums[i-1] == nums[i] {
				continue
			}
			temp = append(temp, nums[i])
			DFS(i+1, temp)
			//回溯
			//var t []int
			//t = append(t, temp[:len(temp)-1]...)
			//temp = t
			temp = temp[:len(temp)-1]
		}
	}
	DFS(0, []int{})
	return res
}

func main() {
	i := subsets([]int{1, 2})
	fmt.Println(i)
}
