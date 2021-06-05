package basic

func LoopPow(x float64, n float64) float64 {

	intn := int(n)

	for i := 1; i < intn; i++ {
		x = x * x
	}
	return x
}
