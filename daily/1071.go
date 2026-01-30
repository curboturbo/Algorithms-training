package main

func getPi(s string) []int {
	n := len(s)
	pi := make([]int, n)
	for i := 1; i < n; i++ {
		j := pi[i-1]
		for j > 0 && s[i] != s[j] {
			j = pi[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		pi[i] = j
	}
	return pi
}

func getMinBase(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}
	pi := getPi(s)
	lastPi := pi[n-1]
	unitLen := n - lastPi
	if n%unitLen == 0 {
		return s[:unitLen]
	}
	return s
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func gcdOfStrings(str1 string, str2 string) string {
	base1 := getMinBase(str1)
	base2 := getMinBase(str2)
	if base1 != base2 {
		return ""
	}
	g := gcd(len(str1), len(str2))
	return str1[:g]
}