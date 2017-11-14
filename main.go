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

// TODO check it sums to 1.
func MatrixVector(x [4][4]uint32, a [4]uint32) [4]uint32 {
	b := [4]uint32{}
	for i := 0; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			b[i] = b[i] + a[i]*x[i][j]
		}
	}
	return b
}

func MatrixIterate(n int, x [4][4]uint32, a [4]uint32) [4]uint32 {
	b := a
	for i := 0; i < n; i++ {
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

	// TODO we should see the RNG on the CPU
	//
	// Next TODO: a PRNG

	// The second set of arguments will be the ports for interacting with memory
	memReadAddr chan<- axiprotocol.Addr,
	memReadData <-chan axiprotocol.ReadData,

	memWriteAddr chan<- axiprotocol.Addr,
	memWriteData chan<- axiprotocol.WriteData,
	memWriteResp <-chan axiprotocol.WriteResp) {

	// Since we're not reading anything from memory, disable those reads
	go axiprotocol.ReadDisable(memReadAddr, memReadData)

	m := [4][4]uint32{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
	v := [4]uint32{
		1, 4, 4, 1,
	}

	iter := a >> 30
	x := MatrixIterate(iter, m, v)

	//outputChannel := make(chan uint32)
	//rand.RandUint32(a, outputChannel)
	//msg := <-outputChannel

	// Calculate the value
	val := a + b + VectorSum(x)

	// Write it back to the pointer the host requests
	aximemory.WriteUInt32(
		memWriteAddr, memWriteData, memWriteResp, false, addr, val)
}
