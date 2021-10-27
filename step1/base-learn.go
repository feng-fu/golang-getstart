package main

import (
	"fmt"
)

func mystring() {
	s := "hello"
	// s[0] = 'c'
	c := []rune(s)
	c[0] = 'c'
	s2 := string(c)
	fmt.Println("current value is ", 'c', s2)
}

func myfunc() {
	i := 0
Here:
	fmt.Println(i)
	i++
	if i < 100 {
		goto Here
	}
}

func myloop() {
	// J:
	for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			if i > j {
				// break J
				break
			}
			fmt.Println(i)
		}
	}
}

func myrange() {
	list := []string{"a", "b"}
	for index, value := range list {
		fmt.Println(index, value)
	}
}

func myswitch(i int) {
	switch i {
	case 0:
		fallthrough
	case 1:
		fmt.Println("current i is", i)
	case 2:
		fallthrough
	default:
		fmt.Println("not matched ", i)
	}
}

func myarray() {
	// a := []int{1, 2, 3, 4, 5}
	// s1 := a[1:2:4]
	// // s1[0] = 100
	// fmt.Printf("%v, %v, %v \n", len(a), cap(a), a)
	// fmt.Printf("%v, %v, %v \n", len(s1), cap(s1), s1)

	// var array [100]int
	// slice := array[0:100]

	// slice[98] = 1
	// slice[99] = 2
	s0 := []int{0, 0}
	s1 := append(s0, 2)
	s2 := append(s1, 3, 5, 7)
	s3 := append(s2, s0...)
	s2[0] = 100
	s0[0] = 300
	fmt.Println(s0, len(s0), s1, s2, s3)
}

func mymap() {
	m := make(map[string]string)
	m["hello"] = "Jack"
	m["test"] = "Hallo"
	fmt.Println(m)
}

func myfor(start int) int {
	sum := start

	// for k,v:=range map {
	// 	fmt.Println("map's key:",k)
	// 	fmt.Println("map's val:",v)
	// }
	for sum < 1000 {
		sum += sum
	}
	return sum
}

func myswitch2() {
	integer := 4
	switch integer {
	case 4:
		fmt.Println("The integer was <= 4")
		fallthrough
	case 5:
		fmt.Println("The integer was <= 5")
		fallthrough
	case 6:
		fmt.Println("The integer was <= 6")
		fallthrough
	case 7:
		fmt.Println("The integer was <= 7")
		fallthrough
	case 8:
		fmt.Println("The integer was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
}

func myf(arg ...int) int {
	var sum int
	for _, v := range arg {
		sum += v
	}
	return sum
}

func mydefer() {
	for i := 0; i < 5; i++ {
		if i == 1 {
			panic("not allowed 1")
		}
		defer fmt.Printf("%d \n", i)
	}
}

func mystruct() {
	type Skills []string

	type Human struct {
		name   string
		age    int
		weight int
		phone  string
	}

	type Student struct {
		Human
		Skills
		int
		speciality string
		phone      string
	}

	jane := Student{Human: Human{"Jane", 35, 100, "1355888"}, speciality: "Chinese", phone: "1312222"}

	fmt.Println("Her name is ", jane.name)
	fmt.Println("Her age is ", jane.age)
	fmt.Println("Her weight is ", jane.weight)
	fmt.Println("her speciality is ", jane.speciality)

	fmt.Println("her skills are ", jane.Skills)

	jane.Skills = []string{"anatomy"}

	fmt.Println("her skills are ", jane.Skills)
	fmt.Println("She acquired two new ones")

	jane.Skills = append(jane.Skills, "physics", "golang")

	fmt.Println("her skills now are ", jane.Skills)

	jane.int = 3
	fmt.Println("Her perferred number is", jane.int)

	fmt.Println("Jane's study phone is ", jane.phone)
	fmt.Println("Jane's work phone is", jane.Human.phone)
}

func main() {
	// mystring()
	// myfunc()
	// myloop()
	// myrange()
	// myswitch(0)
	// myswitch(1)
	// myswitch(2)
	// myswitch(3)
	// myarray()
	// mymap()
	// fmt.Println(myfor(100))
	// myswitch2()
	// fmt.Println(myf(1, 2, 3, 4, 5, 6, 7, 8))
	// mydefer()
	mystruct()
}
