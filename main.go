package main

import (
	// Import the entire framework
	_ "sdaccel"

	aximemory "axi/memory"
	axiprotocol "axi/protocol"
	"github.com/ReconfigureIO/math/rand"
)

func VectorSum(x [4]uint32) uint32 {
	var sum uint32 = 0
	for i := 0; i <= 3; i++ {
		sum = sum + x[i]
	}
	return sum
}

func MatrixVector(x [4][4]uint32, a [4]uint32) [4]uint32 {
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
		b = MatrixVector(x, b)
	}
	return b
}

// The kernel (this goes on the FPGA).
func Top(
	// The first set of arguments to this function can be any number
	// of Go primitive types and can be provided via `SetArg` on the host.

	// For this example, we have 3 arguments: the first is a seed from the
	// CPU.
	a uint32,
	b uint32,
	addr uintptr,

	// The second set of arguments will be the ports for interacting with memory
	memReadAddr chan<- axiprotocol.Addr,
	memReadData <-chan axiprotocol.ReadData,

	memWriteAddr chan<- axiprotocol.Addr,
	memWriteData chan<- axiprotocol.WriteData,
	memWriteResp <-chan axiprotocol.WriteResp) {

	// Since we're not reading anything from memory, disable those reads
	go axiprotocol.ReadDisable(memReadAddr, memReadData)

	m := [4][4]uint32{}
	m[0][0] = 1
	m[1][1] = 1
	m[2][2] = 1
	m[3][3] = 1

	v := [4]uint32{}
	v[0] = 1
	v[1] = 4
	v[2] = 4
	v[3] = 1

	iter := int(a >> 30)

	// matrix iterate can't read from memory?
	x := MatrixIterate(iter, m, v)
	y := MatrixIterate(iter, m, v)

	//outputChannel := make(chan uint32)
	//rand.RandUint32(a, outputChannel)
	//msg := <-outputChannel

	// Calculate the value
	val := VectorSum(x) + VectorSum(y)

	// Write it back to the pointer the host requests
	aximemory.WriteUInt32(
		memWriteAddr, memWriteData, memWriteResp, false, addr, val)
}
