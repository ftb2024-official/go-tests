package arraysandslices

func Sum(nums []int) int {
	var sum int
	for _, v := range nums {
		sum += v
	}

	return sum
}

func SumAll(slices ...[]int) []int {
	result := make([]int, len(slices))
	for i, sl := range slices {
		result[i] = Sum(sl)
	}

	// var result []int
	// for _, slice := range slices {
	// 	result = append(result, Sum(slice))
	// }

	return result
}
