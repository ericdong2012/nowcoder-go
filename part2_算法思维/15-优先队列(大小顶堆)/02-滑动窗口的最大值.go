package _5_优先队列_大小顶堆_

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
