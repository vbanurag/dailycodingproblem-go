package day318

// NumberOfValidPlayListsRec is a recursive solution.
// Runs in O(N) worst case time.
func NumberOfValidPlayListsRec(n, m, b int) int {
	if b >= m {
		return 0
	}

	if n == 0 {
		return 1
	}

	mp := m

	if b != 0 {
		mp--
	}

	return m * NumberOfValidPlayListsRec(n-1, mp, max(0, b-1))
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// NumberOfValidPlayListsIterative is an iterative solution.
// Runs in O(N) worst case time.
func NumberOfValidPlayListsIterative(n, m, b int) int {
	if b >= m {
		return 0
	}

	total := 1
	available := m

	for i := 0; i <= b; i++ {
		total *= available
		available--
	}

	if available < m {
		available++
	}

	for i := b + 1; i < n; i++ {
		total *= available
	}

	return total
}

// NumberOfValidPlayListsFormula is a mathematical formula.
// Technically, runs in O(1) if factorial and exponential functions are constant.
// Formula: M!/(M-B!) * (M-B)^(N-B)
func NumberOfValidPlayListsFormula(n, m, b int) int {
	return binomial(m, m-b) * pow(m-b, n-b)
}

func binomial(a, b int) int {
	result := 1

	for ; a > b; a-- {
		result *= a
	}

	return result
}

func pow(n, exp int) int {
	result := 1

	for i := 0; i < exp; i++ {
		result *= n
	}

	return result
}
