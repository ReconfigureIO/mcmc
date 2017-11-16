package matrix

import (
	"fmt"
	. "github.com/ReconfigureIO/mcmc/matrix"
	"testing"
)

func TestVectorSum(t *testing.T) {
	a := [8]uint32{}
	for i := 0; i < 8; i++ {
		a[i] = 1
	}
	val := VectorSum(a)
	if val != 8 {
		t.Fail()
	}
}

func BenchmarkMatrixMult(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := [8]uint32{}
		x := [8][8]uint32{}
		val := MatrixIterate(100, x, a)
		fmt.Sprintf("%d", val)
	}
}
