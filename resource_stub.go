//go:build !darwin
// +build !darwin

package gomtl

// Buffer is a memory allocation for storing unformatted data
// that is accessible to the GPU.
//
// Reference: https://developer.apple.com/documentation/metal/mtlbuffer
type Buffer struct {
	buffer uintptr
}

// resource implements the Resource interface.
func (b *Buffer) resource() uintptr { return b.buffer }

// Contents returns a pointer to the contents of the buffer.
func (b *Buffer) Contents() uintptr {
	return 0
}

// NewBufferWithLength creates a new buffer with the specified length.
//
// Reference: https://developer.apple.com/documentation/metal/mtldevice/1433375-newbufferwithlength
func (d Device) NewBufferWithLength(length uintptr, opt ResourceOptions) Buffer {
	panic("gomtl: Metal is only supported on macOS")
}

// NewBufferWithBytes creates a new buffer of a given length and initializes its contents by copying existing data into it.
//
// Reference: https://developer.apple.com/documentation/metal/mtldevice/1433429-newbufferwithbytes
func (d Device) NewBufferWithBytes(bytes uintptr, length uintptr, opt ResourceOptions) Buffer {
	panic("gomtl: Metal is only supported on macOS")
}

// TextureDescriptor configures new Texture objects.
//
// Reference: https://developer.apple.com/documentation/metal/mtltexturedescriptor
type TextureDescriptor struct {
	PixelFormat PixelFormat
	Width       uint
	Height      uint
	StorageMode StorageMode
}

// Texture is a memory allocation for storing formatted
// image data that is accessible to the GPU.
//
// Reference: https://developer.apple.com/documentation/metal/mtltexture
type Texture struct {
	texture uintptr

	// Width is the width of the texture image for the base level mipmap, in pixels.
	Width uint

	// Height is the height of the texture image for the base level mipmap, in pixels.
	Height uint
}

// NewTextureWithDescriptor creates a new texture with the provided descriptor using the device.
func (d Device) NewTextureWithDescriptor(td TextureDescriptor) Texture {
	panic("gomtl: Metal is only supported on macOS")
}

// resource implements the Resource interface.
func (t Texture) resource() uintptr { return t.texture }

// ReplaceRegion copies a block of pixels into a section of texture slice 0.
//
// Reference: https://developer.apple.com/documentation/metal/mtltexture/1515464-replaceregion
func (t Texture) ReplaceRegion(region Region, level int, pixelBytes *byte, bytesPerRow uintptr) {
	panic("gomtl: Metal is only supported on macOS")
}

// GetBytes copies a block of pixels from the storage allocation of texture
// slice zero into system memory at a specified address.
//
// Reference: https://developer.apple.com/documentation/metal/mtltexture/1515751-getbytes
func (t Texture) GetBytes(pixelBytes *byte, bytesPerRow uintptr, region Region, level int) {
	panic("gomtl: Metal is only supported on macOS")
}
