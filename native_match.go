package main

func naiveStringMatch(text, pattern string) []int {
	var matches []int

	n := len(text)
	m := len(pattern)

	for i := 0; i <= n-m; i++ {
		j := 0
		for j < m && text[i+j] == pattern[j] {
			j++
		}
		if j == m {
			matches = append(matches, i)
		}
	}
	return matches
}

func NaiveMatch(text string, addresses []string) map[int][]string {
	ret := make(map[int][]string)
	for _, pattern := range addresses {
		positions := naiveStringMatch(text, pattern)
		for _, j := range positions {
			ret[j] = append(ret[j], pattern)
		}
	}
	return ret
}
