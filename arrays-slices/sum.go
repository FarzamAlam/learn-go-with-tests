package arrays

func Sum(numbers []int) int {
	var sum int
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func SumAll(numbers ...[]int) []int {
	var res []int
	for _, number := range numbers {
		sum := 0
		for _, n := range number {
			sum += n
		}
		res = append(res, sum)
	}
	return res
}
