package common

// Sum an array.
func VectorSum(x [4]uint32) uint32 {
	var sum uint32 = 0
	for i := 0; i <= 3; i++ {
		sum = sum + x[i]
	}
	return sum
}

func matrixVector(x [4][4]uint32, a [4]uint32) [4]uint32 {
	b := [4]uint32{}
	go func() {
		for i := 0; i <= 3; i++ {
			for j := 0; j <= 3; j++ {
				b[i] = b[i] + a[i]*x[i][j]
			}
		}
	}()
	return b
}

func MatrixIterate(n int, x [4][4]uint32, a [4]uint32) [4]uint32 {
	b := a
	for i := 0; i < n+15; i++ {
		b = matrixVector(x, b)
	}
	return b
}
