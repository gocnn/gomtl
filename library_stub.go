//go:build !darwin
// +build !darwin

package gomtl

import (
	"errors"
	"fmt"
)

// CompileOptions specifies optional compilation settings for
// the graphics or compute functions within a library.
//
// Reference: https://developer.apple.com/documentation/metal/mtlcompileoptions
type CompileOptions struct {
	// Indicates whether the compiler can perform optimizations for floating-point arithmetic that may violate the IEEE 754 standard.
	FastMathEnabled bool

	// Indicates whether the compiler should compile vertex shaders conservatively to generate consistent position calculations.
	PreserveInvariance bool

	// The language version used to interpret the library source code.
	LanguageVersion LanguageVersion
}

// Library represents a collection of compiled graphics or compute functions.
//
// Reference: https://developer.apple.com/documentation/metal/mtllibrary
type Library struct {
	library uintptr
}

// NewLibraryWithSource creates a new library that contains
// the functions stored in the specified source string.
//
// Reference: https://developer.apple.com/documentation/metal/mtldevice/1433431-newlibrarywithsource
func (d Device) NewLibraryWithSource(source string, optFns ...func(*CompileOptions)) (Library, error) {
	return Library{}, errors.New("gomtl: Metal is only supported on macOS")
}

// Function represents a programmable graphics or compute function executed by the GPU.
//
// Reference: https://developer.apple.com/documentation/metal/mtlfunction.
type Function struct {
	function uintptr
}

// NewFunctionWithName creates a new function object that represents a shader function in the library.
//
// Reference: https://developer.apple.com/documentation/metal/mtllibrary/1515524-newfunctionwithname
func (l Library) NewFunctionWithName(name string) (Function, error) {
	return Function{}, fmt.Errorf("gomtl: Metal is only supported on macOS")
}
