//go:build darwin

package gomtl

// This file provides CGO flags for dynamic linking with MTL libraries (default).
//
// Reference: https://developer.apple.com/documentation/metal
/*
#cgo darwin CFLAGS: -x objective-c
#cgo darwin LDFLAGS: -framework Metal -framework CoreGraphics -framework Foundation
*/
import "C"
