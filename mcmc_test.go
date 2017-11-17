package matrix

import (
	"fmt"
	. "github.com/ReconfigureIO/mcmc/matrix"
	"testing"
)

func TestVectorSum(t *testing.T) {
	a := [16]uint32{}
	for i := 0; i < 16; i++ {
		a[i] = 1
	}
	val := VectorSum(a)
	if val != 16 {
		t.Fail()
	}
}

func BenchmarkMatrixMult(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := [16]uint32{}
		x := [256]uint32{}
		val := MatrixIterate(1000, x, a)
		fmt.Sprintf("%d", val)
	}
}
