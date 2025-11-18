//go:build !darwin
// +build !darwin

package gomtl

import "errors"

// RenderPipelineColorAttachmentDescriptor represents a color attachment descriptor for a render pipeline.
type RenderPipelineColorAttachmentDescriptor struct {
	// PixelFormat is the pixel format of the color attachment's texture.
	PixelFormat PixelFormat
}

// RenderPipelineDescriptor represents a descriptor for a render pipeline.
type RenderPipelineDescriptor struct {
	// VertexFunction is a programmable function that processes individual vertices in a rendering pass.
	VertexFunction Function

	// FragmentFunction is a programmable function that processes individual fragments in a rendering pass.
	FragmentFunction Function

	// ColorAttachments is an array of attachments that store color data.
	ColorAttachments [1]RenderPipelineColorAttachmentDescriptor
}

// RenderPipelineState represents the state of a render pipeline.
type RenderPipelineState struct {
	renderPipelineState uintptr
}

// NewRenderPipelineStateWithDescriptor creates a new render pipeline state with the provided descriptor.
// It returns the created RenderPipelineState or an error if the creation fails.
func (d Device) NewRenderPipelineStateWithDescriptor(rpd RenderPipelineDescriptor) (RenderPipelineState, error) {
	return RenderPipelineState{}, errors.New("gomtl: Metal is only supported on macOS")
}
