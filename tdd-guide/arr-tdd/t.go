package main

func Sum(numbers []int) (sum int) {
	for _, v := range numbers {
		sum += v
	}
	return
}

func SumAll(args ...[]int) []int {
	r := make([]int, len(args))
	for i := 0; i < len(r); i++ {
		r[i] = Sum(args[i])
	}
	return r
	// return []int{3, 9}
}

func SumAllTails(args ...[]int) []int {
	r := make([]int, len(args))

	for i, v := range args {
		if len(v) == 0 {
			r[i] = 0
			continue
		}
		r[i] = Sum(v[1:])
	}
	return r
}
