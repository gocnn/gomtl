//go:build !darwin
// +build !darwin

package gomtl

import "errors"

// ComputePipelineState represents an object that contains a compiled compute pipeline.
//
// Referece: https://developer.apple.com/documentation/metal/mtlcomputepipelinestate
type ComputePipelineState struct {
	computePipelineState          uintptr
	MaxTotalThreadsPerThreadgroup uint
}

// NewComputePipelineStateWithFunction creates a new ComputePipelineState object with the specified compute function.
// It takes a Function object representing the compute function to be used and returns a ComputePipelineState object.
// An error is returned if the creation fails.
//
// Reference: https://developer.apple.com/documentation/metal/mtldevice/1433395-newcomputepipelinestatewithfunct
func (d Device) NewComputePipelineStateWithFunction(f Function) (ComputePipelineState, error) {
	return ComputePipelineState{}, errors.New("gomtl: Metal is only supported on macOS")
}
