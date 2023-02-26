package part2

func twoSum(nums []int, target int) []int {
	// 保存值和索引
	record := make(map[int]int)
	for k, v := range nums {
		if _, ok := record[target-v]; ok {
			return []int{k, record[target-v]}
		}
		record[v] = k
	}
	return nil
}
