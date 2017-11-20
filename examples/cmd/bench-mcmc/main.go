package main

import (
	"encoding/binary"
	// "math/rand"
	"testing"
	"xcl"

	"ReconfigureIO/reco-sdaccel/benchmarks"
)

const (
	MAX_BIT_WIDTH       = 16
	HISTOGRAM_BIT_WIDTH = 9
	HISTOGRAM_WIDTH     = 1 << 9
)

func main() {
	Process("Markov Chain Monte Carlo")
}

func Process(name string) {
	world := xcl.NewWorld()
	defer world.Release()

	program := world.Import("kernel_test")
	defer program.Release()

	krnl := program.GetKernel("reconfigure_io_sdaccel_builder_stub_0_1")
	defer krnl.Release()

	f := func(B *testing.B) {
		doit(world, krnl, B)
	}

	bm := testing.Benchmark(f)
	benchmarks.GipedaResults(name, bm)
}

func doit(world xcl.World, krnl *xcl.Kernel, B *testing.B) {
	B.SetBytes(4)
	B.ReportAllocs()

	// The data we'll send to the kernel for processing
	//seed := rand.Uint32()

	input := [64]uint32{}
	for i := 0; i < 64; i++ {
		/// this weird-ish hack gets us the identity matrix.
		if i/8 == i%8 {
			input[i] = 1
		}
	}

	inputBuff := world.Malloc(xcl.ReadOnly, uint(binary.Size(input)))
	defer inputBuff.Free()

	buff := world.Malloc(xcl.WriteOnly, 4)
	defer buff.Free()

	// set iterations to 1000
	krnl.SetArg(0, 1000)
	// set input length.
	krnl.SetArg(1, 64)
	krnl.SetMemoryArg(2, inputBuff)
	krnl.SetMemoryArg(3, buff)

	binary.Write(inputBuff.Writer(), binary.LittleEndian, &input)

	krnl.Run(1, 1, 1)
	B.StopTimer()
}
