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

func TestMatrixMult(t *testing.T) {
	a := [8]uint32{1, 1, 1, 1, 1, 1, 1, 1}
	x := [64]uint32{}
	c := make(chan uint32, 64)
	go func() {
		for i := 0; i < 64; i++ {
			if i%8 == i/8 {
				x[i] = 1
			}
			c <- x[i]
		}
	}()
	val := MatrixIterate(100, c, a)
	if val != [8]uint32{1, 1, 1, 1, 1, 1, 1, 1} {
		fmt.Print("%d", val)
		t.Fail()
	}
}

func BenchmarkMatrixMult(b *testing.B) {
	b.SetBytes(4 * 64)
	for i := 0; i < b.N; i++ {
		a := [8]uint32{}
		x := [64]uint32{}
		c := make(chan uint32, 64)
		go func() {
			for i := 0; i < 64; i++ {
				c <- x[i]
			}
		}()
		val := MatrixIterate(1000, c, a)
		fmt.Sprintf("%d", val)
	}
}
