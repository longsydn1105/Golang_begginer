package main

func longestSubarray(nums []int) int {
	left, maxLen, k := 0, 0, 1

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			k--
		}

		for k < 0 {
			if nums[left] == 0 {
				k++
			}
			left++
		}
		currLen := right - left + 1
		if maxLen < currLen {
			maxLen = currLen
		}
	}

	if k == len(nums) {
		return 0
	}

	if k == 1 {
		return len(nums) - 1
	}
	return maxLen
}
