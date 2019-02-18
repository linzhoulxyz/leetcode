package twosum

func twoSum(nums []int, target int) []int {
	var i, j int
	var found = false
	var result = []int{}
	for i = 0; i < len(nums)-1; i++ {
		if found {
			break
		}

		for j = i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				found = true
				result = append(result, i, j)
				break
			}
		}
	}

	return result
}

// 优化版本
func twoSum2(nums []int, target int) []int {
	maps := map[int]int{}
	for i, a := range nums {
		b := target - a
		if j, ok := maps[b]; ok {
			return []int{j, i}
		}

		maps[a] = i
	}

	return nil
}
