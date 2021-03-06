package main

import (
	// Import the entire framework
	_ "github.com/ReconfigureIO/sdaccel"

	"github.com/ReconfigureIO/math/rand"
	. "github.com/ReconfigureIO/mcmc/matrix"
	aximemory "github.com/ReconfigureIO/sdaccel/axi/memory"
	axiprotocol "github.com/ReconfigureIO/sdaccel/axi/protocol"
)

// The kernel (this goes on the FPGA).
func Top(
	// The first set of arguments to this function can be any number
	// of Go primitive types and can be provided via `SetArg` on the host.

	// For this example, we have 3 arguments: the first is a seed from the
	// CPU.
	a uint32,
	inputLength uint32,
	inputAddr uintptr,
	addr uintptr,

	// The second set of arguments will be the ports for interacting with memory
	memReadAddr chan<- axiprotocol.Addr,
	memReadData <-chan axiprotocol.ReadData,

	memWriteAddr chan<- axiprotocol.Addr,
	memWriteData chan<- axiprotocol.WriteData,
	memWriteResp <-chan axiprotocol.WriteResp) {

	inputChannel := make(chan uint32)

	go aximemory.ReadBurstUInt32(
		memReadAddr, memReadData, true, inputAddr, inputLength, inputChannel)

	v := [8]uint32{}
	v[0] = 1
	v[1] = 4
	v[2] = 4
	v[3] = 1

	iter := int(a) // int(a >> 30)

	// matrix iterate can't read from memory?
	x := MatrixIterate(iter, inputChannel, v)

	// Calculate the value
	val := VectorSum(x) // + VectorSum(y)

	// Write it back to the pointer the host requests
	aximemory.WriteUInt32(
		memWriteAddr, memWriteData, memWriteResp, false, addr, val)
}
