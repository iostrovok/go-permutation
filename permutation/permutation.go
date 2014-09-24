package permutation

func Strings(line []string) [][]string {
	res := make([][]string, 0)

	if len(line) == 0 {
		return res
	}

	if len(line) == 1 {
		return [][]string{{line[0]}}
	}

	if len(line) == 2 {
		res = append(res, []string{line[0], line[1]})
		res = append(res, []string{line[1], line[0]})
		return res
	}

	for i := 0; i < len(line); i++ {
		small := []string{}
		small = append(small, line[:i]...)
		small = append(small, line[(i+1):]...)
		next := Strings(small)
		for _, v := range next {
			t := append([]string{line[i]}, v...)
			res = append(res, t)
		}
	}

	return res
}
