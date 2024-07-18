package iteration

func Repeat(ch string, times int) string {
	var result string
	for i := 0; i < times; i++ {
		result += ch
	}

	return result
}
