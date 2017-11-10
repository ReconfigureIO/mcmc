package main

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"xcl"
)

func main() {
	// Allocate a world for interacting with kernels
	world := xcl.NewWorld()
	defer world.Release()

	// Import the kernel.
	// Right now these two idenitifers are hard coded as an output from the build process
	krnl := world.Import("kernel_test").GetKernel("reconfigure_io_sdaccel_builder_stub_0_1")
	defer krnl.Release()

	// Allocate a buffer on the FPGA to store the return value of our computation
	// The output is a uint32, so we use 4 bytes to store it
	buff := world.Malloc(xcl.WriteOnly, 4)
	defer buff.Free()

	// Pass the arguments to the kernel

	// The first argument will be the seed.
	krnl.SetArg(0, rand.Uint32())
	// The second argument will be the number of iterations to do.
	krnl.SetArg(1, 2)
	// Set the pointer to the output buffer
	krnl.SetMemoryArg(2, buff)

	// Run the kernel with the supplied arguments
	krnl.Run(1, 1, 1)

	// Decode the byte slice given by the FPGA into the uint32 we're expecting
	var ret uint32
	err := binary.Read(buff.Reader(), binary.LittleEndian, &ret)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}

	// Just here to verify we can indeed generate random numbers on the CPU
	// side of things.
	fmt.Printf("%d, %d\n", rand.Int(), rand.Int())

	// Print the value given by the FPGA
	fmt.Printf("%d\n", ret)

	// Exit with an error if the value is not correct
	/*if ret != 3 {
		os.Exit(1)
	}*/
}
