package basic

func LeastCommonMultiple(m int, n int) int {

	k := greatestCommonDivisor(m, n)

	return (m * n) / k
}

func greatestCommonDivisor(m int, n int) int {

	big := m
	little := n

	if big < n {
		big = n
		little = m
	}

	for {

		mod := big % little

		if mod == 0 {
			return little
		} else {
			big = little
			little = mod
		}

	}

}
