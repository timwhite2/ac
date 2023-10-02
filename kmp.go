package main

func computeNext(pattern string) []int {
	m := len(pattern)
	next := make([]int, m)
	length := 0
	i := 1

	for i < m {
		if pattern[i] == pattern[length] {
			length++
			next[i] = length
			i++
		} else {
			if length != 0 {
				length = next[length-1]
			} else {
				next[i] = 0
				i++
			}
		}
	}

	return next
}

func kmpSearch(text, pattern string) []int {
	m, n := len(pattern), len(text)
	if m == 0 {
		return []int{}
	}

	next := computeNext(pattern)
	result := make([]int, 0)
	i, j := 0, 0

	for i < n {
		if pattern[j] == text[i] {
			i++
			j++
		}

		if j == m {
			result = append(result, i-j)
			j = next[j-1]
		} else if i < n && pattern[j] != text[i] {
			if j != 0 {
				j = next[j-1]
			} else {
				i++
			}
		}
	}

	return result
}

func KMP(text string, addresses []string) map[int][]string {
	ret := make(map[int][]string)
	for _, address := range addresses {
		positions := kmpSearch(text, address)
		for _, j := range positions {
			ret[j] = append(ret[j], address)
		}
	}
	return ret
}
