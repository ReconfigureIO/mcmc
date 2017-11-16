package matrix

import (
	"fmt"
	. "github.com/ReconfigureIO/mcmc/matrix"
	"testing"
)

func TestVectorSum(t *testing.T) {
	a := [4]uint32{
		1, 4, 4, 1,
	}
	val := VectorSum(a)
	if val != 10 {
		t.Fail()
	}
}

func BenchmarkMatrixMult(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := [4]uint32{}
		x := [4][4]uint32{}
		val := MatrixIterate(100, x, a)
		fmt.Sprintf("%d", val)
	}
}
