package main

func longestOnes(nums []int, k int) int {
    left := 0
	maxLen := 0

	for right := 1; right < len(nums); right++ {
		if nums[right] == 0 {
			k--
		}

		for k < 0 {
			if nums[left] == 0 {
				k++
			}
			left ++
		}

		currLen := right - left + 1
		if currLen > maxLen {
			maxLen = currLen
		}
	}

	return maxLen
}