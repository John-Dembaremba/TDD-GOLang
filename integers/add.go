package integers

func Add(data []int) int {
	results := 0
	for _, number := range data {
		results += number
	}
	return results
}
