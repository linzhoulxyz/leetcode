package TwoSum

func twoSum(nums []int, target int) []int {
    var i, j int
    var found = false
    var result = []int{}
    for i = 0; i < len(nums) - 1; i++ {
        if found {
            break
        }
        
        for j = i + 1; j < len(nums); j++ {
            if nums[i] + nums[j] == target {
                found = true
                result = append(result, i, j)
                break;
            }
        }
    }
    
    return result
}