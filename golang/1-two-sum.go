package main

func twoSum(nums []int, target int) []int {
	lenNums := len(nums)
	for idx1, num := range nums {
		if idx1 >= lenNums {
			break
		}

		slicedIndex := idx1 + 1
		complementary := target - num
		for idx2, num := range nums[slicedIndex:] {
			if num == complementary {
				return []int{idx1, idx2 + slicedIndex}
			}
		}
	}

	return []int{}
}
