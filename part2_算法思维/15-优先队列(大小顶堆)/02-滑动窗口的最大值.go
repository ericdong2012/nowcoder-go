package _5_优先队列_大小顶堆_


// 可以和part1/04-stack/滑动窗口的最大值 对比(用的单调栈)
func maxValue(temp []int) int {
	result := temp[0]
	for i := 0; i < len(temp); i++ {
		if temp[i] > result {
			result = temp[i]
		}
	}

	return result
}

func maxWindow(nums []int, k int) []int {
	if len(nums) == 0 || len(nums) < k {
		return []int{}
	}
	stack := []int{}
	// 123456  3
	for i := 0; i <= len(nums)-k; i++ {
		result := maxValue(nums[i : i+k])
		stack = append(stack, result)
	}

	return stack
}
