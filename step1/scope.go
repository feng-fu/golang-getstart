package main

import "fmt"

func calculate(args ...float64) float64 {
	count := len(args)
	if count == 0 {
		return 0
	}
	var sum float64
	for _, v := range args {
		sum += v
	}
	return sum / float64(count)
}

func bubble(list []float64) []float64 {
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
	return list
}

func fibonacci(end int) []int {
	var a = []int{}
	for i := 0; i < end; i++ {
		l := len(a)
		var v int
		if l == 0 || l == 1 {
			v = 0
		} else if l == 2 {
			v = 1
		} else {
			v = a[l-1] + a[l-2]
		}
		a = append(a, v)
	}
	return a
}

func plusTwo(y int) func(v int) int {
	return func(v int) int {
		return v + y
	}
}

func getMax(l []int) int {
	if len(l) == 0 {
		return 0
	}
	max := l[0]

	for _, v := range l {
		if v > max {
			max = v
		}
	}
	return max
}

func Map(f func(int) int, m []int) []int {
	r := []int{}
	for _, v := range m {
		r = append(r, f(v))
	}
	return r
}

func main() {
	// a := func() {
	// 	fmt.Println("Hello")
	// }

	// a()

	// fmt.Println(
	// 	calculate(10.2, 100, 2, 2.4, 12, 4),
	// 	calculate(), calculate(1, 0))
	// fmt.Println(bubble([]float64{100, 2, 311, 10, 20.2}))
	// fmt.Println(fibonacci(20))
	// p := plusTwo(1)
	// fmt.Printf("%v\n", p(2))

	// m := []int{1, 3, 4}
	// f := func(i int) int {
	// 	return i * i
	// }
	// fmt.Printf("%v", (Map(f, m)))

	var p *[]int = new([]int)
	var v []int = make([]int, 10, 100)

	fmt.Println(p, v)

	// 	fmt.Println(
	// 		getMax([]int{1, 2, 3, 2, 3, 4, 12, 423, 13, 23, 23, 2}))
}
