//go:build !darwin
// +build !darwin

package gomtl

// CommandQueue represents a queue that organizes the order
// in which command buffers are executed by the GPU.
//
// Reference: https://developer.apple.com/documentation/metal/mtlcommandqueue
type CommandQueue struct {
	commandQueue uintptr
}

// NewCommandQueue creates a new command queue for the specified device.
// The command queue is used to submit rendering and computation commands to the GPU.
//
// Reference: https://developer.apple.com/documentation/metal/mtldevice/1433388-newcommandqueue
func (d Device) NewCommandQueue() CommandQueue {
	panic("gomtl: Metal is only supported on macOS")
}

// CommandBuffer returns a command buffer from the command queue that maintains strong references to resources.
//
// Reference: https://developer.apple.com/documentation/metal/mtlcommandqueue/1508686-commandbuffer
func (cq CommandQueue) CommandBuffer() CommandBuffer {
	panic("gomtl: Metal is only supported on macOS")
}
