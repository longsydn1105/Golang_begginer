package main

func accumulation(nums []int) int {

	summation := 0

	for _, num := range nums {
		summation += num
	}

	return summation
}

func pivotIndex(nums []int) int {
	totalWeightOnLeft := 0
	totalWeightOnRight := accumulation(nums)

	for idx, currentWeight := range nums {
		totalWeightOnRight -= currentWeight

		if totalWeightOnLeft == totalWeightOnRight {
			return idx
		}

		totalWeightOnLeft += currentWeight
	}

	return -1

}
