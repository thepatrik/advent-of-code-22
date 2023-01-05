package slice

func Sum(array []int) int {
	sum := 0
	for _, v := range array {
		sum += v
	}
	return sum
}
