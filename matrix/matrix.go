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

		//done := make(chan bool)

		b = [8]uint32{}

		go func() {
			b[0] = b[0] + a[0]*x[0] + a[1]*x[8] + a[2]*x[16] + a[3]*x[24] + a[4]*x[32] + a[5]*x[40] + a[6]*x[48] + a[7]*x[56]
			//done <- true
		}()
		go func() {
			b[1] = b[1] + a[0]*x[1] + a[1]*x[9] + a[2]*x[17] + a[3]*x[25] + a[4]*x[33] + a[5]*x[41] + a[6]*x[49] + a[7]*x[57]
			//done <- true
		}()
		go func() {
			b[2] = b[2] + a[0]*x[2] + a[1]*x[10] + a[2]*x[18] + a[3]*x[26] + a[4]*x[34] + a[5]*x[42] + a[6]*x[50] + a[7]*x[58]
			//done <- true
		}()
		go func() {
			b[3] = b[3] + a[0]*x[3] + a[1]*x[11] + a[2]*x[19] + a[3]*x[27] + a[4]*x[35] + a[5]*x[43] + a[6]*x[51] + a[7]*x[59]
			//done <- true
		}()
		go func() {
			b[4] = b[4] + a[0]*x[4] + a[1]*x[12] + a[2]*x[20] + a[3]*x[28] + a[4]*x[36] + a[5]*x[44] + a[6]*x[52] + a[7]*x[60]
			//done <- true
		}()
		go func() {
			b[5] = b[5] + a[0]*x[5] + a[1]*x[13] + a[2]*x[21] + a[3]*x[29] + a[4]*x[37] + a[5]*x[45] + a[6]*x[53] + a[7]*x[61]
			//done <- true
		}()
		go func() {
			b[6] = b[6] + a[0]*x[6] + a[1]*x[14] + a[2]*x[22] + a[3]*x[30] + a[4]*x[38] + a[5]*x[46] + a[6]*x[54] + a[7]*x[62]
			//done <- true
		}()
		go func() {
			b[7] = b[7] + a[0]*x[7] + a[1]*x[15] + a[2]*x[23] + a[3]*x[31] + a[4]*x[39] + a[5]*x[47] + a[6]*x[55] + a[7]*x[63]
			//done <- true
		}()
		//for i := 0; i < 7; i++ {
		//	<-done
		//}
	}
	return b
}
