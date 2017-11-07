package main

import (
	// Import the entire framework
	_ "sdaccel"

	aximemory "axi/memory"
	axiprotocol "axi/protocol"
)

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

	// Calculate the value
	val := a + b

	// Write it back to the pointer the host requests
	aximemory.WriteUInt32(
		memWriteAddr, memWriteData, memWriteResp, false, addr, val)
}
