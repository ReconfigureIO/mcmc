package matrix

// Sum an array.
func VectorSum(x [8]uint32) uint32 {
	var sum uint32 = 0
	for i := 0; i <= 7; i++ {
		sum = sum + x[i]
	}
	return sum
}

func MatrixIterate(n int, c <-chan uint32, a [8]uint32) [8]uint32 {
	x := [64]uint32{}
	for i := 0; i < 64; i++ {
		x[i] = <-c
	}
	b := a
	for i := 0; i < n; i++ {
		b = [8]uint32{}
		// TODO hard-code this w/ a code generator? go func() for each of them.
		// this will be ~64 times faster.
		go func() {
			b[0] = b[0] + a[0]*x[0] + a[1]*x[8] + a[2]*x[16] + a[3]*x[24] + a[4]*x[32] + a[5]*x[40] + a[6]*x[48] + a[7]*x[56]
		}()
		go func() {
			b[1] = b[1] + a[0]*x[1] + a[1]*x[9] + a[2]*x[17] + a[3]*x[25] + a[4]*x[33] + a[5]*x[41] + a[6]*x[49] + a[7]*x[57]
		}()
		go func() {
			b[2] = b[2] + a[0]*x[2] + a[1]*x[10] + a[2]*x[18] + a[3]*x[26] + a[4]*x[34] + a[5]*x[42] + a[6]*x[50] + a[7]*x[58]
		}()
		/*for i := 0; i <= 7; i++ {
			for j := 0; j <= 7; j++ {
				b[i] = b[i] + a[i]*x[8*i+j]
			}
		}*/
	}
	return b
}
