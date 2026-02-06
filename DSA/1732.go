package main

func largestAltitude(gain []int) int {
	maxAlt, currAlt := 0, 0

	for _, v := range gain {
		currAlt += v

		if maxAlt < currAlt {
			maxAlt = currAlt
		}
	}

	return maxAlt
}
