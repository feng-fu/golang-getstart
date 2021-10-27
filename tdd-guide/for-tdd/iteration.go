package iteration

func Repeat(s string, t int) string {
	repeated := ""
	for i := 0; i < t; i++ {
		repeated += s
	}
	return repeated
}
