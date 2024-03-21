package series

func All(n int, s string) []string {
	series := []string{}

	for i := 0; i+n <= len(s); i++ {
		series = append(series, s[i:i+n])
	}

	return series
}

func UnsafeFirst(n int, s string) string {
	return All(n, s)[0]
}

func First(n int, s string) (first string, ok bool) {
	val := All(n, s)
	if len(val) == 0 {
		return "", false
	}
	return val[0], true
}
