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

func SumAllTails(slices ...[]int) []int {
	var result []int
	for _, slice := range slices {
		if len(slice) == 0 {
			result = append(result, 0)
		} else {
			result = append(result, Sum(slice[1:]))
		}
	}

	return result
}
