package main

func twoSum(nums []int, target int) []int {
	set := make(map[int]int, 0)

	for i, _ := range nums {
		if si, ok := set[target-nums[i]]; ok {
			return []int{si, i}
		}
		set[nums[i]] = i
	}
	return []int{}
}
