package main 

func longestSubsequence(nums []int) int {
	maxLen := 0
	tails := make([]int, 0, len(nums))

	for bit := 0; bit < 31; bit++ {
		tails = tails[:0]
		for _, num := range nums {
			if (num>>bit)&1 == 0 {
				continue
			}

			l, r := 0, len(tails)
			for l < r {
				mid := (l + r) >> 1
				if tails[mid] < num {
					l = mid + 1
				} else {
					r = mid
				}
			}

			if l == len(tails) {
				tails = append(tails, num)
			} else {
				tails[l] = num
			}
		}
		if len(tails) > maxLen {
			maxLen = len(tails)
		}
	}
	return maxLen
}