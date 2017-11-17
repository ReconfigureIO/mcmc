package matrix

// Sum an array.
func VectorSum(x [8]uint32) uint32 {
	var sum uint32 = 0
	for i := 0; i <= 7; i++ {
		sum = sum + x[i]
	}
	return sum
}

func matrixVector(x [64]uint32, a [8]uint32) [8]uint32 {
	b := [8]uint32{}
	go func() {
		for i := 0; i <= 7; i++ {
			for j := 0; j <= 7; j++ {
				b[i] = b[i] + a[i]*x[8*i+j]
			}
		}
	}()
	return b
}

func MatrixIterate(n int, x [64]uint32, a [8]uint32) [8]uint32 {
	b := a
	for i := 0; i < n; i++ {
		b = matrixVector(x, b)
	}
	return b
}
