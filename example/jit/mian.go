//go:build darwin

package main

import (
	_ "embed"
	"fmt"
	"log"
	"unsafe"

	"github.com/gocnn/gomtl"
)

//go:embed add.metal
var metalSource string

func main() {
	// Initialize Metal
	device, err := gomtl.CreateSystemDefaultDevice()
	if err != nil {
		log.Fatal(err)
	}
	library, err := device.NewLibraryWithSource(metalSource)
	if err != nil {
		log.Fatal(err)
	}

	function, err := library.NewFunctionWithName("vector_add")
	if err != nil {
		log.Fatal(err)
	}

	pipelineState, err := device.NewComputePipelineStateWithFunction(function)
	if err != nil {
		log.Fatal(err)
	}

	// Prepare data
	const n = 1024
	a := make([]float32, n)
	b := make([]float32, n)
	for i := 0; i < n; i++ {
		a[i] = float32(i)
		b[i] = float32(i * 2)
	}

	// Create buffers
	bufferA := device.NewBufferWithBytes(unsafe.Pointer(&a[0]), uintptr(n*4), gomtl.ResourceStorageModeShared)
	bufferB := device.NewBufferWithBytes(unsafe.Pointer(&b[0]), uintptr(n*4), gomtl.ResourceStorageModeShared)
	bufferResult := device.NewBufferWithLength(uintptr(n*4), gomtl.ResourceStorageModeShared)

	// Execute compute shader
	commandQueue := device.NewCommandQueue()
	commandBuffer := commandQueue.CommandBuffer()
	encoder := commandBuffer.ComputeCommandEncoder()

	encoder.SetComputePipelineState(pipelineState)
	encoder.SetBuffer(bufferA, 0, 0)
	encoder.SetBuffer(bufferB, 0, 1)
	encoder.SetBuffer(bufferResult, 0, 2)

	gridSize := gomtl.Size{Width: n, Height: 1, Depth: 1}
	threadgroupSize := gomtl.Size{Width: 64, Height: 1, Depth: 1}
	encoder.DispatchThreads(gridSize, threadgroupSize)
	encoder.EndEncoding()

	commandBuffer.Commit()
	commandBuffer.WaitUntilCompleted()

	// Verify results
	results := (*[1024]float32)(bufferResult.Contents())[:n:n]
	for i := 0; i < 5; i++ {
		expected := a[i] + b[i]
		fmt.Printf("%.0f + %.0f = %.0f\n", a[i], b[i], results[i])
		if results[i] != expected {
			log.Fatalf("Incorrect result at %d: got %.1f, expected %.1f", i, results[i], expected)
		}
	}
	fmt.Printf("✓ Vector addition completed successfully (%d elements)\n", n)
}
