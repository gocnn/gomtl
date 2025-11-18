//go:build !darwin
// +build !darwin

package gomtl

// CommandEncoder is an encoder that writes sequential GPU commands
// into a command buffer.
//
// Reference: https://developer.apple.com/documentation/metal/mtlcommandencoder
type CommandEncoder struct {
	commandEncoder uintptr
}

// EndEncoding declares that all command generation from this encoder is completed.
//
// Reference: https://developer.apple.com/documentation/metal/mtlcommandencoder/1458038-endencoding
func (ce CommandEncoder) EndEncoding() {
	panic("gomtl: Metal is only supported on macOS")
}

// ComputeCommandEncoder is an object for encoding commands in a compute pass.
//
// Reference: https://developer.apple.com/documentation/metal/mtlcomputecommandencoder
type ComputeCommandEncoder struct {
	CommandEncoder
}

// SetComputePipelineState sets the current compute pipeline state object for the compute command encoder.
//
// Reference: https://developer.apple.com/documentation/metal/mtlcomputecommandencoder/1443140-setcomputepipelinestate
func (cce ComputeCommandEncoder) SetComputePipelineState(cps ComputePipelineState) {
	panic("gomtl: Metal is only supported on macOS")
}

// SetBuffer sets a buffer for the compute function at a specified index.
//
// Reference: https://developer.apple.com/documentation/metal/mtlcomputecommandencoder/1443126-setbuffer
func (cce ComputeCommandEncoder) SetBuffer(buf Buffer, offset, index int) {
	panic("gomtl: Metal is only supported on macOS")
}

// DispatchThreads encodes a compute command using an arbitrarily sized grid.
//
// Reference: https://developer.apple.com/documentation/metal/mtlcomputecommandencoder/2866532-dispatchthreads
func (cce ComputeCommandEncoder) DispatchThreads(gridSize, threadgroupSize Size) {
	panic("gomtl: Metal is only supported on macOS")
}

// RenderCommandEncoder is an encoder that specifies graphics-rendering commands
// and executes graphics functions.
//
// Reference: https://developer.apple.com/documentation/metal/mtlrendercommandencoder.
type RenderCommandEncoder struct {
	CommandEncoder
}

// SetRenderPipelineState sets the current render pipeline state object.
//
// Reference: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1515811-setrenderpipelinestate.
func (rce RenderCommandEncoder) SetRenderPipelineState(rps RenderPipelineState) {
	panic("gomtl: Metal is only supported on macOS")
}

// SetVertexBuffer sets a buffer for the vertex shader function at an index
// in the buffer argument table with an offset that specifies the start of the data.
//
// Reference: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1515829-setvertexbuffer.
func (rce RenderCommandEncoder) SetVertexBuffer(buf Buffer, offset, index int) {
	panic("gomtl: Metal is only supported on macOS")
}

// SetVertexBytes sets a block of data for the vertex function.
//
// Reference: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1515846-setvertexbytes.
func (rce RenderCommandEncoder) SetVertexBytes(bytes uintptr, length uintptr, index int) {
	panic("gomtl: Metal is only supported on macOS")
}

// DrawPrimitives renders one instance of primitives using vertex data
// in contiguous array elements.
//
// Reference: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1516326-drawprimitives.
func (rce RenderCommandEncoder) DrawPrimitives(typ PrimitiveType, vertexStart, vertexCount int) {
	panic("gomtl: Metal is only supported on macOS")
}

// BlitCommandEncoder is an encoder that specifies resource copy
// and resource synchronization commands.
//
// Reference: https://developer.apple.com/documentation/metal/mtlblitcommandencoder.
type BlitCommandEncoder struct {
	CommandEncoder
}

// CopyFromTexture encodes a command to copy image data from a slice of
// a source texture into a slice of a destination texture.
//
// Reference: https://developer.apple.com/documentation/metal/mtlblitcommandencoder/1400754-copyfromtexture.
func (bce BlitCommandEncoder) CopyFromTexture(
	src Texture, srcSlice, srcLevel int, srcOrigin Origin, srcSize Size,
	dst Texture, dstSlice, dstLevel int, dstOrigin Origin,
) {
	panic("gomtl: Metal is only supported on macOS")
}

// SynchronizeResource flushes any copy of the specified resource from its corresponding
// Device caches and, if needed, invalidates any CPU caches.
//
// Reference: https://developer.apple.com/documentation/metal/mtlblitcommandencoder/1400775-synchronize.
func (bce BlitCommandEncoder) SynchronizeResource(resource Resource) {
	panic("gomtl: Metal is only supported on macOS")
}
