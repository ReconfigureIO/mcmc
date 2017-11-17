package matrix

// Sum an array.
func VectorSum(x [16]uint32) uint32 {
	var sum uint32 = 0
	for i := 0; i <= 15; i++ {
		sum = sum + x[i]
	}
	return sum
}

func matrixVector(x [256]uint32, a [16]uint32) [16]uint32 {
	b := [16]uint32{}
	go func() {
		for i := 0; i <= 15; i++ {
			for j := 0; j <= 15; j++ {
				b[i] = b[i] + a[i]*x[16*i+j]
			}
		}
	}()
	return b
}

func MatrixIterate(n int, x [256]uint32, a [16]uint32) [16]uint32 {
	b := a
	for i := 0; i < n; i++ {
		b = matrixVector(x, b)
	}
	return b
}
